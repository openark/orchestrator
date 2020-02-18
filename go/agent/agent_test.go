package agent

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"os/exec"
	"strconv"
	"testing"
	"time"

	logger "log"

	"github.com/go-martini/martini"
	"github.com/google/go-cmp/cmp"
	"github.com/martini-contrib/render"
	"github.com/openark/golib/log"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/db"
	"github.com/github/orchestrator/go/inst"
	mysqldrv "github.com/go-sql-driver/mysql"
	"github.com/openark/golib/sqlutils"
	test "github.com/openark/golib/tests"
	"github.com/ory/dockertest"
)

// before running add following to your /etc/hosts file depending on number of agents you plan to use
// 127.0.0.2 agent1
// 127.0.0.3 agent2
// ...
// 127.0.0.n agentn

// also if you are using Mac OS X, you will need to create the aliases with
// sudo ifconfig lo0 alias 127.0.0.2 up
// sudo ifconfig lo0 alias 127.0.0.3 up
// for each agent

func init() {
	log.SetLevel(log.DEBUG)
}

type testAgent struct {
	agent                 *Agent
	agentSeedStageStatus  *SeedStageState
	agentSeedMetadata     *SeedMetadata
	agentServer           *httptest.Server
	agentMux              *martini.ClassicMartini
	agentMySQLContainerIP string
	agentMySQLContainerID string
}

var testAgents = make(map[int]*testAgent)

func prepareTestData(t *testing.T, agentsNumber int) func(t *testing.T) {
	log.Info("Setting up test data")
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	log.Info("Creating docker container for Orchestrator DB")
	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	config.Config.MySQLOrchestratorHost = "127.0.0.1"
	port, _ := strconv.Atoi(resource.GetPort("3306/tcp"))
	config.Config.MySQLOrchestratorPort = uint(port)
	config.Config.MySQLOrchestratorUser = "root"
	config.Config.MySQLOrchestratorPassword = "secret"
	config.Config.MySQLOrchestratorDatabase = "orchestrator"
	config.Config.AgentsServerPort = ":3001"
	config.Config.ServeAgentsHttp = true
	config.Config.AuditToSyslog = false
	config.Config.EnableSyslog = false
	config.Config.MySQLConnectTimeoutSeconds = 600
	config.Config.MySQLConnectionLifetimeSeconds = 600
	config.Config.MySQLOrchestratorReadTimeoutSeconds = 600
	config.Config.MySQLTopologyReadTimeoutSeconds = 600
	config.Config.MySQLTopologySSLSkipVerify = true
	config.Config.MySQLTopologyUseMixedTLS = true
	config.Config.HostnameResolveMethod = "none"
	config.Config.MySQLHostnameResolveMethod = "none"
	falseFlag := false
	trueFlag := true
	config.RuntimeCLIFlags.Noop = &falseFlag
	config.RuntimeCLIFlags.SkipUnresolve = &trueFlag
	config.Config.SkipMaxScaleCheck = true

	config.MarkConfigurationLoaded()

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var testdb *sql.DB
		var err error
		testdb, _, err = sqlutils.GetDB(fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
		mysqldrv.SetLogger(logger.New(ioutil.Discard, "discard", 1))
		if err != nil {
			return err
		}
		return testdb.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	log.Info("Creating Orchestrator DB")
	_, err = db.OpenOrchestrator()
	//inst.initializeInstanceDao()
	if err != nil {
		log.Fatalf("Unable to create orchestrator DB: %s", err)
	}
	inst.InitializeInstanceDao()
	InitHttpClient()

	go func() {
		for seededAgent := range SeededAgents {
			instanceKey := &inst.InstanceKey{Hostname: seededAgent.Info.Hostname, Port: int(seededAgent.Info.MySQLPort)}
			log.Infof("%+v", instanceKey)
		}
	}()

	log.Info("Creating Orchestrator agents mocks")
	for i := 1; i <= agentsNumber; i++ {
		mysqlDatabases := map[string]*MySQLDatabase{
			"sakila": &MySQLDatabase{
				Engines: []Engine{InnoDB},
				Size:    0,
			},
		}
		availiableSeedMethods := map[SeedMethod]*SeedMethodOpts{
			Mydumper: &SeedMethodOpts{
				BackupSide:       Target,
				SupportedEngines: []Engine{ROCKSDB, MRG_MYISAM, CSV, BLACKHOLE, InnoDB, MEMORY, ARCHIVE, MyISAM, FEDERATED, TokuDB},
			},
		}
		agent := &Agent{
			Info: &Info{
				Hostname: fmt.Sprintf("agent%d", i),
				Port:     3002 + i,
				Token:    "token",
			},
			Data: &Data{
				LocalSnapshotsHosts: []string{fmt.Sprintf("127.0.0.%d", i)},
				SnaphostHosts:       []string{fmt.Sprintf("127.0.0.%d", i), "localhost"},
				LogicalVolumes:      []*LogicalVolume{},
				MountPoint: &Mount{
					Path:       "/tmp",
					Device:     "",
					LVPath:     "",
					FileSystem: "",
					IsMounted:  false,
					DiskUsage:  0,
				},
				BackupDir:             "/tmp/bkp",
				BackupDirDiskFree:     10000,
				MySQLRunning:          true,
				MySQLDatadir:          "/var/lib/mysql",
				MySQLDatadirDiskUsed:  10,
				MySQLDatadirDiskFree:  10000,
				MySQLVersion:          "5.7.25",
				MySQLDatabases:        mysqlDatabases,
				AvailiableSeedMethods: availiableSeedMethods,
			},
		}
		testAgent := createTestAgent(t, agent)
		testAgents[i] = testAgent
	}

	return func(t *testing.T) {
		log.Info("Teardown test data")
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
		for _, testAgent := range testAgents {
			testAgent.agentServer.Close()
		}
	}
}

func createTestAgent(t *testing.T, agent *Agent) *testAgent {
	agentAddress := fmt.Sprintf("%s:%d", agent.Info.Hostname, agent.Info.Port)
	m := martini.Classic()
	m.Use(render.Renderer())
	testAgent := &testAgent{
		agent:                agent,
		agentSeedStageStatus: &SeedStageState{},
		agentSeedMetadata:    &SeedMetadata{},
	}
	m.Get("/api/get-agent-data", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.JSON(200, agent.Data)
	})
	m.Get("/api/prepare/:seedID/:seedMethod/:seedSide", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.Text(202, "Started")
	})
	m.Get("/api/backup/:seedID/:seedMethod/:seedHost/:mysqlPort", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.Text(202, "Started")
	})
	m.Get("/api/restore/:seedID/:seedMethod", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.Text(202, "Started")
	})
	m.Get("/api/cleanup/:seedID/:seedMethod/:seedSide", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.Text(202, "Started")
	})
	m.Get("/api/get-metadata/:seedID/:seedMethod", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.JSON(200, testAgent.agentSeedMetadata)
	})
	m.Get("/api/seed-stage-state/:seedID/:seedStage", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.JSON(200, testAgent.agentSeedStageStatus)
	})
	m.Get("/api/abort-seed-stage/:seedID/:seedStage", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.Text(200, "killed")
	})
	testServer := httptest.NewUnstartedServer(m)
	listener, _ := net.Listen("tcp", agentAddress)
	testServer.Listener = listener
	testServer.Start()
	testAgent.agentServer = testServer
	testAgent.agentMux = m
	return testAgent
}

func registerAgent(t *testing.T, testAgent *testAgent) (string, error) {
	hostname, err := RegisterAgent(testAgent.agent.Info)
	if err != nil {
		return "", err
	}
	return hostname, nil
}

func expectStructsEquals(t *testing.T, actual, value interface{}) {
	if cmp.Equal(actual, value) {
		return
	}
	t.Errorf("Expected:\n[[[%+v]]]\n- got:\n[[[%+v]]]", value, actual)
}

func TestAgentRegistration(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	hostname, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(hostname, testAgents[1].agent.Info.Hostname)
}

func TestReadAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	registeredAgents, err := ReadAgents()
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(2, len(registeredAgents))

	testAgents := []*Agent{testAgents[1].agent, testAgents[2].agent}
	for _, testAgent := range testAgents {
		for _, registeredAgent := range registeredAgents {
			if registeredAgent.Info.Port == testAgent.Info.Port && registeredAgent.Info.Hostname == testAgent.Info.Hostname {
				expectStructsEquals(t, registeredAgent.Info, testAgent.Info)
				expectStructsEquals(t, registeredAgent.Data, testAgent.Data)
			}
		}
	}
}

func TestReadAgentsInfo(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	registeredAgents, err := ReadAgentsInfo()
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(2, len(registeredAgents))

	testAgents := []*Agent{testAgents[1].agent, testAgents[2].agent}
	for _, testAgent := range testAgents {
		for _, registeredAgent := range registeredAgents {
			if registeredAgent.Info.Port == testAgent.Info.Port && registeredAgent.Info.Hostname == testAgent.Info.Hostname {
				expectStructsEquals(t, registeredAgent.Info, testAgent.Info)
				expectStructsEquals(t, registeredAgent.Data, &Data{})
			}
		}
	}
}

func TestReadAgent(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	testAgent := testAgents[1]

	_, err := registerAgent(t, testAgent)
	test.S(t).ExpectNil(err)

	registeredAgent, err := ReadAgent(testAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)
	expectStructsEquals(t, registeredAgent.Info, testAgent.agent.Info)
	expectStructsEquals(t, registeredAgent.Data, testAgent.agent.Data)
}

func TestReadAgentInfo(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	testAgent := testAgents[1]

	_, err := registerAgent(t, testAgent)
	test.S(t).ExpectNil(err)

	registeredAgent, err := ReadAgentInfo(testAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)
	expectStructsEquals(t, registeredAgent.Info, testAgent.agent.Info)
	expectStructsEquals(t, registeredAgent.Data, &Data{})
}

func TestReadOutdatedAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	config.Config.AgentPollMinutes = 2

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	db.ExecOrchestrator(fmt.Sprintf("UPDATE host_agent SET last_checked = NOW() WHERE hostname='%s'", sourceTestAgent.agent.Info.Hostname))
	db.ExecOrchestrator(fmt.Sprintf("UPDATE host_agent SET last_checked = NOW() - interval 60 minute WHERE hostname='%s'", targetTestAgent.agent.Info.Hostname))

	outdatedAgents, err := ReadOutdatedAgents()
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(1, len(outdatedAgents))
	expectStructsEquals(t, outdatedAgents[0].Info, targetTestAgent.agent.Info)
	expectStructsEquals(t, outdatedAgents[0].Data, &Data{})
}

func TestUpdateAgent(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	testAgent := testAgents[1]

	_, err := registerAgent(t, testAgent)
	test.S(t).ExpectNil(err)

	registeredAgent, err := ReadAgentInfo(testAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	testAgent.agent.Data.LocalSnapshotsHosts = []string{"127.0.0.10", "127.0.0.12"}
	registeredAgent.Status = Inactive
	registeredAgent.updateAgentStatus()

	err = registeredAgent.UpdateAgent()
	test.S(t).ExpectNil(err)

	expectStructsEquals(t, registeredAgent.Info, testAgent.agent.Info)
	expectStructsEquals(t, registeredAgent.Data, testAgent.agent.Data)
	test.S(t).ExpectEquals(registeredAgent.Status, Active)
}

func TestUpdateAgentFailed(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	testAgent := testAgents[1]

	_, err := registerAgent(t, testAgent)
	test.S(t).ExpectNil(err)

	testAgent.agentServer.Close()

	registeredAgent, err := ReadAgentInfo(testAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	err = registeredAgent.UpdateAgent()
	test.S(t).ExpectNotNil(err)

	registeredAgent, err = ReadAgentInfo(testAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(registeredAgent.Status, Inactive)
}

func TestForgetLongUnseenAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	testAgent := testAgents[1]
	config.Config.UnseenAgentForgetHours = 1
	_, err := registerAgent(t, testAgent)
	test.S(t).ExpectNil(err)

	db.ExecOrchestrator(fmt.Sprintf("UPDATE host_agent SET last_seen = last_seen - interval 2 hour WHERE hostname='%s'", testAgent.agent.Info.Hostname))

	err = ForgetLongUnseenAgents()
	test.S(t).ExpectNil(err)

	_, err = ReadAgentInfo(testAgent.agent.Info.Hostname)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeed(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))
}

func TestNewSeedWrongSeedMethod(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("test", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedSeedItself(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedUnsupportedSeedMethod(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mysqldump", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedUnsupportedSeedMethodForDB(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	sourceTestAgent.agent.Data.AvailiableSeedMethods[Xtrabackup] = &SeedMethodOpts{
		BackupSide:       Target,
		SupportedEngines: []Engine{MRG_MYISAM, CSV, BLACKHOLE, InnoDB, MEMORY, ARCHIVE, MyISAM, FEDERATED, TokuDB},
	}
	sourceTestAgent.agent.Data.MySQLDatabases["test"] = &MySQLDatabase{
		Engines: []Engine{ROCKSDB},
		Size:    0,
	}
	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Xtrabackup", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedSourceAgentMySQLVersionLessThanTarget(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	targetTestAgent.agent.Data.MySQLVersion = "5.6.40"
	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedAgentHadActiveSeed(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}
func TestNewSeedNotEnoughSpaceInMySQLDatadir(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	targetTestAgent.agent.Data.MySQLDatadirDiskFree = 10
	sourceTestAgent.agent.Data.MySQLDatadirDiskUsed = 1000

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestReadActiveSeeds(t *testing.T) {
	teardownTestCase := prepareTestData(t, 4)
	defer teardownTestCase(t)

	targetTestAgent1 := testAgents[1]
	sourceTestAgent1 := testAgents[2]
	targetTestAgent2 := testAgents[3]
	sourceTestAgent2 := testAgents[4]

	_, err := registerAgent(t, targetTestAgent1)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent1)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, targetTestAgent2)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent2)
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent(targetTestAgent1.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent(sourceTestAgent1.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	targetAgent2, err := ReadAgent(targetTestAgent2.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent2, err := ReadAgent(sourceTestAgent2.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	seedID, err = NewSeed("Mydumper", targetAgent2, sourceAgent2)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(2))

	seeds, err := ReadActiveSeeds()
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(len(seeds), 2)

	for _, seed := range seeds {
		if seed.SeedID == 1 {
			test.S(t).ExpectEquals(seed.TargetHostname, targetTestAgent1.agent.Info.Hostname)
			test.S(t).ExpectEquals(seed.SourceHostname, sourceTestAgent1.agent.Info.Hostname)
		} else {
			test.S(t).ExpectEquals(seed.TargetHostname, targetTestAgent2.agent.Info.Hostname)
			test.S(t).ExpectEquals(seed.SourceHostname, sourceTestAgent2.agent.Info.Hostname)
		}
		test.S(t).ExpectEquals(seed.SeedMethod, Mydumper)
		test.S(t).ExpectEquals(seed.BackupSide, Target)
		test.S(t).ExpectEquals(seed.Status, Started)
		test.S(t).ExpectEquals(seed.Stage, Prepare)
		test.S(t).ExpectEquals(seed.Retries, 0)
	}
}

func TestReadRecentSeeds(t *testing.T) {
	teardownTestCase := prepareTestData(t, 4)
	defer teardownTestCase(t)

	targetTestAgent1 := testAgents[1]
	sourceTestAgent1 := testAgents[2]
	targetTestAgent2 := testAgents[3]
	sourceTestAgent2 := testAgents[4]

	_, err := registerAgent(t, targetTestAgent1)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent1)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, targetTestAgent2)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent2)
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent(targetTestAgent1.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent(sourceTestAgent1.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	targetAgent2, err := ReadAgent(targetTestAgent2.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent2, err := ReadAgent(sourceTestAgent2.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	seedID, err = NewSeed("Mydumper", targetAgent2, sourceAgent2)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(2))

	seeds, err := ReadRecentSeeds()
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(len(seeds), 2)

	for _, seed := range seeds {
		if seed.SeedID == 1 {
			test.S(t).ExpectEquals(seed.TargetHostname, targetTestAgent1.agent.Info.Hostname)
			test.S(t).ExpectEquals(seed.SourceHostname, sourceTestAgent1.agent.Info.Hostname)
		} else {
			test.S(t).ExpectEquals(seed.TargetHostname, targetTestAgent2.agent.Info.Hostname)
			test.S(t).ExpectEquals(seed.SourceHostname, sourceTestAgent2.agent.Info.Hostname)
		}
		test.S(t).ExpectEquals(seed.SeedMethod, Mydumper)
		test.S(t).ExpectEquals(seed.BackupSide, Target)
		test.S(t).ExpectEquals(seed.Status, Started)
		test.S(t).ExpectEquals(seed.Stage, Prepare)
		test.S(t).ExpectEquals(seed.Retries, 0)
	}
}

func TestReadSeed(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	seed, err := ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.TargetHostname, targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectEquals(seed.SourceHostname, sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectEquals(seed.SeedMethod, Mydumper)
	test.S(t).ExpectEquals(seed.BackupSide, Target)
	test.S(t).ExpectEquals(seed.Status, Started)
	test.S(t).ExpectEquals(seed.Stage, Prepare)
	test.S(t).ExpectEquals(seed.Retries, 0)
}

func TestReadRecentSeedsForAgentInStatus(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	for _, agent := range testAgents {
		seeds, err := ReadRecentSeedsForAgentInStatus(agent.agent, Started, "limit 1")
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(len(seeds), 1)
		test.S(t).ExpectEquals(seeds[0].TargetHostname, targetTestAgent.agent.Info.Hostname)
		test.S(t).ExpectEquals(seeds[0].SourceHostname, sourceTestAgent.agent.Info.Hostname)
		test.S(t).ExpectEquals(seeds[0].SeedMethod, Mydumper)
		test.S(t).ExpectEquals(seeds[0].BackupSide, Target)
		test.S(t).ExpectEquals(seeds[0].Status, Started)
		test.S(t).ExpectEquals(seeds[0].Stage, Prepare)
		test.S(t).ExpectEquals(seeds[0].Retries, 0)
	}

	for _, agent := range testAgents {
		seeds, err := ReadRecentSeedsForAgentInStatus(agent.agent, Running, "limit 1")
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(len(seeds), 0)
	}

}

func TestReadActiveSeedsForAgent(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	for _, agent := range testAgents {
		seeds, err := ReadActiveSeedsForAgent(agent.agent)
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(len(seeds), 1)
		test.S(t).ExpectEquals(seeds[0].TargetHostname, targetTestAgent.agent.Info.Hostname)
		test.S(t).ExpectEquals(seeds[0].SourceHostname, sourceTestAgent.agent.Info.Hostname)
		test.S(t).ExpectEquals(seeds[0].SeedMethod, Mydumper)
		test.S(t).ExpectEquals(seeds[0].BackupSide, Target)
		test.S(t).ExpectEquals(seeds[0].Status, Started)
		test.S(t).ExpectEquals(seeds[0].Stage, Prepare)
		test.S(t).ExpectEquals(seeds[0].Retries, 0)
	}

}

func TestGetSeedAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	seed, err := ReadSeed(seedID)
	test.S(t).ExpectNil(err)

	targetAgent, sourceAgent, err = seed.GetSeedAgents()
	test.S(t).ExpectNil(err)
	expectStructsEquals(t, targetAgent.Info, targetTestAgent.agent.Info)
	expectStructsEquals(t, sourceAgent.Info, sourceTestAgent.agent.Info)
}

func createAgentsMySQLServers(t *testing.T, agent *testAgent, useGTID bool, createSlaveUser bool) (func(t *testing.T), error) {
	log.Info("Setting agents MySQL")
	var testdb *sql.DB
	rand.Seed(time.Now().UnixNano())
	serverID := rand.Intn(100000000)
	pool, err := dockertest.NewPool("")

	dockerCmd := []string{"mysqld", fmt.Sprintf("--server-id=%d", serverID), "--log-bin=/var/lib/mysql/mysql-bin"}
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	log.Info("Creating docker container for agent")
	// pulls an image, creates a container based on it and runs it
	if useGTID {
		dockerCmd = append(dockerCmd, "--enforce-gtid-consistency=ON", "--gtid-mode=ON")
	}
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{Repository: "mysql", Tag: "5.7", Env: []string{"MYSQL_ROOT_PASSWORD=secret"}, Cmd: dockerCmd, CapAdd: []string{"NET_ADMIN", "NET_RAW"}})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	agent.agentMySQLContainerIP = resource.Container.NetworkSettings.Networks["bridge"].IPAddress
	agent.agentMySQLContainerID = resource.Container.ID

	teardownFunc := func(t *testing.T) {
		log.Info("Teardown MySQL for agent")
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}

	cmd := exec.Command("docker", "exec", "-i", agent.agentMySQLContainerID, "apt-get", "update")
	if err := cmd.Run(); err != nil {
		return teardownFunc, err
	}
	cmd = exec.Command("docker", "exec", "-i", agent.agentMySQLContainerID, "apt-get", "install", "-y", "iptables")
	if err := cmd.Run(); err != nil {
		return teardownFunc, err
	}

	if err := pool.Retry(func() error {
		var err error
		testdb, _, err = sqlutils.GetDB(fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
		mysqldrv.SetLogger(logger.New(ioutil.Discard, "discard", 1))
		if err != nil {
			return err
		}
		return testdb.Ping()
	}); err != nil {
		return teardownFunc, fmt.Errorf("Could not connect to docker: %s", err)
	}

	if createSlaveUser {
		if _, err = testdb.Exec("Create database orchestrator_meta;"); err != nil {
			return teardownFunc, err
		}
		if _, err = testdb.Exec("Create table orchestrator_meta.replication (`username` varchar(128) CHARACTER SET ascii NOT NULL DEFAULT '',`password` varchar(128) CHARACTER SET ascii NOT NULL DEFAULT '',PRIMARY KEY (`username`,`password`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;"); err != nil {
			return teardownFunc, err
		}
		if _, err = testdb.Exec("CREATE USER `slave`@`%` IDENTIFIED BY 'slavepassword@';"); err != nil {
			return teardownFunc, err
		}
		if _, err = testdb.Exec("GRANT REPLICATION SLAVE ON *.* TO `slave`@`%`"); err != nil {
			return teardownFunc, err
		}
		if _, err = testdb.Exec("CREATE USER `orc_topology`@`%` IDENTIFIED BY 'orc_topologypassword@';"); err != nil {
			return teardownFunc, err
		}
		if _, err = testdb.Exec("GRANT ALL PRIVILEGES ON *.* TO `orc_topology`@`%`"); err != nil {
			return teardownFunc, err
		}
		if _, err = testdb.Exec("Insert into orchestrator_meta.replication(username, password) VALUES ('slave', 'slavepassword@');"); err != nil {
			return teardownFunc, err
		}
	}

	if _, err = testdb.Exec("Create database test_repl"); err != nil {
		return teardownFunc, err
	}
	if _, err = testdb.Exec("Create table test_repl.test(id int);"); err != nil {
		return teardownFunc, err
	}
	if _, err = testdb.Exec("insert into test_repl.test(id) VALUES (1), (2), (3), (4);"); err != nil {
		return teardownFunc, err
	}

	if err = sqlutils.QueryRowsMap(testdb, "show master status", func(m sqlutils.RowMap) error {
		var err error
		agent.agentSeedMetadata.LogFile = m.GetString("File")
		agent.agentSeedMetadata.LogPos = m.GetInt64("Position")
		agent.agentSeedMetadata.GtidExecuted = m.GetString("Executed_Gtid_Set")
		return err
	}); err != nil {
		return teardownFunc, err
	}

	mysqlPort, err := strconv.Atoi(resource.GetPort("3306/tcp"))
	if err != nil {
		return teardownFunc, err
	}
	agent.agent.Info.MySQLPort = mysqlPort

	return func(t *testing.T) {
		log.Info("Teardown MySQL for agent")
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}, nil
}

// this is needed in order to redirect from host machine ip 172.0.0.x to container ip for replication to work
func createIPTablesRulesForReplication(t *testing.T, targetAgent *testAgent, sourceAgent *testAgent) error {
	cmd := exec.Command("docker", "exec", "-i", targetAgent.agentMySQLContainerID, "iptables", "-t", "nat", "-I", "OUTPUT", "-p", "tcp", "-o", "eth0", "--dport", fmt.Sprintf("%d", sourceAgent.agent.Info.MySQLPort), "-j", "DNAT", "--to-destination", fmt.Sprintf("%s:3306", sourceAgent.agentMySQLContainerIP))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", string(out))
	}
	// iptables -t nat -I OUTPUT -p tcp -o eth0 --dport 33447 -j DNAT --to-destination 172.17.0.4:3306
	cmd = exec.Command("docker", "exec", "-i", targetAgent.agentMySQLContainerID, "/bin/sh", "-c", fmt.Sprintf("echo %s %s >> /etc/hosts", sourceAgent.agentMySQLContainerIP, sourceAgent.agent.Info.Hostname))
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", string(out))
	}
	return nil
}

func TestProcessSeeds(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	targetTestAgent := testAgents[1]
	sourceTestAgent := testAgents[2]

	_, err := registerAgent(t, targetTestAgent)
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, sourceTestAgent)
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent(sourceTestAgent.agent.Info.Hostname)
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	ProcessSeeds()
	seed, err := ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Prepare)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err := seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Prepare)
		test.S(t).ExpectEquals(seedState.Status, Running)
	}

	for _, agent := range testAgents {
		agent.agentSeedStageStatus.SeedID = int64(seedID)
		agent.agentSeedStageStatus.Stage = Prepare
		agent.agentSeedStageStatus.Hostname = agent.agent.Info.Hostname
		agent.agentSeedStageStatus.Timestamp = time.Now()
		agent.agentSeedStageStatus.Status = Running
		agent.agentSeedStageStatus.Details = "processing prepare stage"
	}
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Prepare)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:2] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Prepare)
		test.S(t).ExpectEquals(seedState.Status, Running)
		test.S(t).ExpectEquals(seedState.Details, "processing prepare stage")
	}

	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Details = "completed prepare stage"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Prepare)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:2] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Prepare)
		if seedState.Hostname == targetTestAgent.agent.Info.Hostname {
			test.S(t).ExpectEquals(seedState.Status, Completed)
			test.S(t).ExpectEquals(seedState.Details, "completed prepare stage")
		}
		if seedState.Hostname == sourceTestAgent.agent.Info.Hostname {
			test.S(t).ExpectEquals(seedState.Status, Running)
			test.S(t).ExpectEquals(seedState.Details, "processing prepare stage")
		}
	}

	sourceTestAgent.agentSeedStageStatus.Status = Completed
	sourceTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	sourceTestAgent.agentSeedStageStatus.Details = "completed prepare stage"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Started)
	test.S(t).ExpectEquals(seed.Stage, Backup)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Prepare)
		test.S(t).ExpectEquals(seedState.Status, Completed)
		test.S(t).ExpectEquals(seedState.Details, "completed prepare stage")
	}

	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Backup)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Backup)
		test.S(t).ExpectEquals(seedState.Status, Running)
	}

	targetTestAgent.agentSeedStageStatus.Stage = Backup
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "running backup stage"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Backup)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Backup)
		test.S(t).ExpectEquals(seedState.Status, Running)
		test.S(t).ExpectEquals(seedState.Details, "running backup stage")
	}

	targetTestAgent.agentSeedStageStatus.Stage = Backup
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Details = "completed backup stage"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Started)
	test.S(t).ExpectEquals(seed.Stage, Restore)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Backup)
		test.S(t).ExpectEquals(seedState.Status, Completed)
		test.S(t).ExpectEquals(seedState.Details, "completed backup stage")
	}

	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Restore)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Restore)
		test.S(t).ExpectEquals(seedState.Status, Running)
	}

	targetTestAgent.agentSeedStageStatus.Stage = Restore
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "running restore stage"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Restore)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Restore)
		test.S(t).ExpectEquals(seedState.Status, Running)
		test.S(t).ExpectEquals(seedState.Details, "running restore stage")
	}

	targetTestAgent.agentSeedStageStatus.Stage = Restore
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Details = "completed restore stage"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Started)
	test.S(t).ExpectEquals(seed.Stage, Cleanup)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Restore)
		test.S(t).ExpectEquals(seedState.Status, Completed)
		test.S(t).ExpectEquals(seedState.Details, "completed restore stage")
	}

	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Cleanup)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:2] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Cleanup)
		test.S(t).ExpectEquals(seedState.Status, Running)
	}

	for _, agent := range testAgents {
		agent.agentSeedStageStatus.Stage = Cleanup
		agent.agentSeedStageStatus.Timestamp = time.Now()
		agent.agentSeedStageStatus.Status = Running
		agent.agentSeedStageStatus.Details = "processing cleanup stage"
	}
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Cleanup)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:2] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Cleanup)
		test.S(t).ExpectEquals(seedState.Status, Running)
		test.S(t).ExpectEquals(seedState.Details, "processing cleanup stage")
	}

	sourceTestAgent.agentSeedStageStatus.Status = Completed
	sourceTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	sourceTestAgent.agentSeedStageStatus.Details = "completed cleanup stage"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Running)
	test.S(t).ExpectEquals(seed.Stage, Cleanup)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:2] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Cleanup)
		if seedState.Hostname == targetTestAgent.agent.Info.Hostname {
			test.S(t).ExpectEquals(seedState.Status, Running)
			test.S(t).ExpectEquals(seedState.Details, "processing cleanup stage")
		}
		if seedState.Hostname == sourceTestAgent.agent.Info.Hostname {
			test.S(t).ExpectEquals(seedState.Status, Completed)
			test.S(t).ExpectEquals(seedState.Details, "completed cleanup stage")
		}
	}

	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Details = "completed cleanup stage"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Started)
	test.S(t).ExpectEquals(seed.Stage, ConnectSlave)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, Cleanup)
		test.S(t).ExpectEquals(seedState.Status, Completed)
		test.S(t).ExpectEquals(seedState.Details, "completed cleanup stage")
	}

	// create mysql containers for agents
	teardownTargetAgent, err := createAgentsMySQLServers(t, targetTestAgent, true, true)
	if err != nil {
		teardownTargetAgent(t)
		log.Fatale(err)
	}
	defer teardownTargetAgent(t)

	teardownSourceAgent, err := createAgentsMySQLServers(t, sourceTestAgent, true, true)
	if err != nil {
		teardownSourceAgent(t)
		log.Fatale(err)
	}
	defer teardownSourceAgent(t)

	err = createIPTablesRulesForReplication(t, targetTestAgent, sourceTestAgent)
	if err != nil {
		log.Errore(err)
		return
	}

	//register then one more time to update mysqlport
	_, err = RegisterAgent(targetTestAgent.agent.Info)
	test.S(t).ExpectNil(err)
	_, err = RegisterAgent(sourceTestAgent.agent.Info)
	test.S(t).ExpectNil(err)

	config.Config.MySQLTopologyUser = "orc_topology"
	config.Config.MySQLTopologyPassword = "orc_topologypassword@"
	config.Config.ReplicationCredentialsQuery = "select username, password from orchestrator_meta.replication;"
	ProcessSeeds()
	seed, err = ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.Status, Completed)
	test.S(t).ExpectEquals(seed.Stage, ConnectSlave)
	test.S(t).ExpectEquals(seed.Retries, 0)
	seedStates, err = seed.ReadSeedStageStates()
	test.S(t).ExpectNil(err)
	for _, seedState := range seedStates[:1] {
		test.S(t).ExpectEquals(seedState.SeedID, int64(1))
		test.S(t).ExpectEquals(seedState.Stage, ConnectSlave)
		test.S(t).ExpectEquals(seedState.Status, Completed)
		test.S(t).ExpectEquals(seedState.Details, "Seed completed")
	}

}

func TestProcessSeeds2(t *testing.T) {
	fnc, err := createAgentsMySQLServers(t, &testAgent{}, true, true)
	defer fnc(t)
	if err != nil {
		log.Errore(err)
	}
}
