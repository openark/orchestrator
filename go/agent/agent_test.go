package agent

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
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

// If you want to create multiple agent instance, the simplest way is to use multiple loopback aliases. On
// modern linux distributions you probably don't need to do anything, but on Mac OS X, you will need to create the aliases with
// sudo ifconfig lo0 alias 127.0.0.2 up
// sudo ifconfig lo0 alias 127.0.0.3 up
// for each agent

func init() {
	log.SetLevel(log.DEBUG)
}

type testAgent struct {
	agent                *Agent
	agentSeedStageStatus *SeedStageState
	agentSeedMetadata    *SeedMetadata
	agentServer          *httptest.Server
	agentMux             *martini.ClassicMartini
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
	config.MarkConfigurationLoaded()

	go func() {
		for seededAgent := range SeededAgents {
			instanceKey := &inst.InstanceKey{Hostname: seededAgent.Info.Hostname, Port: int(seededAgent.Info.MySQLPort)}
			log.Infof("%+v", instanceKey)
		}
	}()

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
				Hostname: fmt.Sprintf("127.0.0.%d", i),
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

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	registeredAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)
	expectStructsEquals(t, registeredAgent.Info, testAgents[1].agent.Info)
	expectStructsEquals(t, registeredAgent.Data, testAgents[1].agent.Data)
}

func TestReadAgentInfo(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	registeredAgent, err := ReadAgentInfo("127.0.0.1")
	test.S(t).ExpectNil(err)
	expectStructsEquals(t, registeredAgent.Info, testAgents[1].agent.Info)
	expectStructsEquals(t, registeredAgent.Data, &Data{})
}

func TestReadOutdatedAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	config.Config.AgentPollMinutes = 2

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	db.ExecOrchestrator("UPDATE host_agent SET last_checked = NOW() WHERE hostname='127.0.0.2'")

	outdatedAgents, err := ReadOutdatedAgents()
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(1, len(outdatedAgents))
	expectStructsEquals(t, outdatedAgents[0].Info, testAgents[1].agent.Info)
	expectStructsEquals(t, outdatedAgents[0].Data, &Data{})
}

func TestUpdateAgent(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	registeredAgent, err := ReadAgentInfo("127.0.0.1")
	test.S(t).ExpectNil(err)

	testAgents[1].agent.Data.LocalSnapshotsHosts = []string{"127.0.0.10", "127.0.0.12"}
	registeredAgent.Status = Inactive
	registeredAgent.updateAgentStatus()

	err = registeredAgent.UpdateAgent()
	test.S(t).ExpectNil(err)

	expectStructsEquals(t, registeredAgent.Info, testAgents[1].agent.Info)
	expectStructsEquals(t, registeredAgent.Data, testAgents[1].agent.Data)
	test.S(t).ExpectEquals(registeredAgent.Status, Active)
}

func TestUpdateAgentFailed(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	testAgents[1].agentServer.Close()

	registeredAgent, err := ReadAgentInfo("127.0.0.1")
	test.S(t).ExpectNil(err)

	err = registeredAgent.UpdateAgent()
	test.S(t).ExpectNotNil(err)

	registeredAgent, err = ReadAgentInfo("127.0.0.1")
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(registeredAgent.Status, Inactive)
}

func TestForgetLongUnseenAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)
	config.Config.UnseenAgentForgetHours = 1
	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	db.ExecOrchestrator("UPDATE host_agent SET last_seen = last_seen - interval 2 hour WHERE hostname='127.0.0.1'")

	err = ForgetLongUnseenAgents()
	test.S(t).ExpectNil(err)

	_, err = ReadAgentInfo("127.0.0.1")
	test.S(t).ExpectNotNil(err)
}

func TestNewSeed(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))
}

func TestNewSeedWrongSeedMethod(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	_, err = NewSeed("test", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedSeedItself(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedUnsupportedSeedMethod(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mysqldump", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedUnsupportedSeedMethodForDB(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	testAgents[2].agent.Data.AvailiableSeedMethods[Xtrabackup] = &SeedMethodOpts{
		BackupSide:       Target,
		SupportedEngines: []Engine{MRG_MYISAM, CSV, BLACKHOLE, InnoDB, MEMORY, ARCHIVE, MyISAM, FEDERATED, TokuDB},
	}
	testAgents[2].agent.Data.MySQLDatabases["test"] = &MySQLDatabase{
		Engines: []Engine{ROCKSDB},
		Size:    0,
	}
	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Xtrabackup", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedSourceAgentMySQLVersionLessThanTarget(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	testAgents[1].agent.Data.MySQLVersion = "5.6.40"
	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestNewSeedAgentHadActiveSeed(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}
func TestNewSeedNotEnoughSpaceInMySQLDatadir(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	testAgents[1].agent.Data.MySQLDatadirDiskFree = 10
	testAgents[2].agent.Data.MySQLDatadirDiskUsed = 1000

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	test.S(t).ExpectNotNil(err)
}

func TestReadActiveSeeds(t *testing.T) {
	teardownTestCase := prepareTestData(t, 4)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[3])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[4])
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	targetAgent2, err := ReadAgent("127.0.0.3")
	test.S(t).ExpectNil(err)

	sourceAgent2, err := ReadAgent("127.0.0.4")
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
			test.S(t).ExpectEquals(seed.TargetHostname, testAgents[1].agent.Info.Hostname)
			test.S(t).ExpectEquals(seed.SourceHostname, testAgents[2].agent.Info.Hostname)
		} else {
			test.S(t).ExpectEquals(seed.TargetHostname, testAgents[3].agent.Info.Hostname)
			test.S(t).ExpectEquals(seed.SourceHostname, testAgents[4].agent.Info.Hostname)
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

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[3])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[4])
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	targetAgent2, err := ReadAgent("127.0.0.3")
	test.S(t).ExpectNil(err)

	sourceAgent2, err := ReadAgent("127.0.0.4")
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
			test.S(t).ExpectEquals(seed.TargetHostname, testAgents[1].agent.Info.Hostname)
			test.S(t).ExpectEquals(seed.SourceHostname, testAgents[2].agent.Info.Hostname)
		} else {
			test.S(t).ExpectEquals(seed.TargetHostname, testAgents[3].agent.Info.Hostname)
			test.S(t).ExpectEquals(seed.SourceHostname, testAgents[4].agent.Info.Hostname)
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
	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	seed, err := ReadSeed(seedID)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seed.TargetHostname, testAgents[1].agent.Info.Hostname)
	test.S(t).ExpectEquals(seed.SourceHostname, testAgents[2].agent.Info.Hostname)
	test.S(t).ExpectEquals(seed.SeedMethod, Mydumper)
	test.S(t).ExpectEquals(seed.BackupSide, Target)
	test.S(t).ExpectEquals(seed.Status, Started)
	test.S(t).ExpectEquals(seed.Stage, Prepare)
	test.S(t).ExpectEquals(seed.Retries, 0)
}

func TestReadRecentSeedsForAgentInStatus(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	for _, agent := range testAgents {
		seeds, err := ReadRecentSeedsForAgentInStatus(agent.agent, Started, "limit 1")
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(len(seeds), 1)
		test.S(t).ExpectEquals(seeds[0].TargetHostname, testAgents[1].agent.Info.Hostname)
		test.S(t).ExpectEquals(seeds[0].SourceHostname, testAgents[2].agent.Info.Hostname)
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

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	for _, agent := range testAgents {
		seeds, err := ReadActiveSeedsForAgent(agent.agent)
		test.S(t).ExpectNil(err)
		test.S(t).ExpectEquals(len(seeds), 1)
		test.S(t).ExpectEquals(seeds[0].TargetHostname, testAgents[1].agent.Info.Hostname)
		test.S(t).ExpectEquals(seeds[0].SourceHostname, testAgents[2].agent.Info.Hostname)
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

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	test.S(t).ExpectNil(err)
	test.S(t).ExpectEquals(seedID, int64(1))

	seed, err := ReadSeed(seedID)
	test.S(t).ExpectNil(err)

	targetAgent, sourceAgent, err := seed.GetSeedAgents()
	test.S(t).ExpectNil(err)
	expectStructsEquals(t, targetAgent.Info, testAgents[1].agent.Info)
	expectStructsEquals(t, sourceAgent.Info, testAgents[2].agent.Info)
}

func TestProcessSeeds(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)

	_, err := registerAgent(t, testAgents[1])
	test.S(t).ExpectNil(err)

	_, err = registerAgent(t, testAgents[2])
	test.S(t).ExpectNil(err)

	targetAgent1, err := ReadAgent("127.0.0.1")
	test.S(t).ExpectNil(err)

	sourceAgent1, err := ReadAgent("127.0.0.2")
	test.S(t).ExpectNil(err)

	seedID, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
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

	testAgents[1].agentSeedStageStatus.Status = Completed
	testAgents[1].agentSeedStageStatus.Timestamp = time.Now()
	testAgents[1].agentSeedStageStatus.Details = "completed prepare stage"
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
		if seedState.Hostname == "127.0.0.1" {
			test.S(t).ExpectEquals(seedState.Status, Completed)
			test.S(t).ExpectEquals(seedState.Details, "completed prepare stage")
		}
		if seedState.Hostname == "127.0.0.2" {
			test.S(t).ExpectEquals(seedState.Status, Running)
			test.S(t).ExpectEquals(seedState.Details, "processing prepare stage")
		}
	}

	testAgents[2].agentSeedStageStatus.Status = Completed
	testAgents[2].agentSeedStageStatus.Timestamp = time.Now()
	testAgents[2].agentSeedStageStatus.Details = "completed prepare stage"
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

	testAgents[1].agentSeedStageStatus.Stage = Backup
	testAgents[1].agentSeedStageStatus.Timestamp = time.Now()
	testAgents[1].agentSeedStageStatus.Status = Running
	testAgents[1].agentSeedStageStatus.Details = "running backup stage"
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

	testAgents[1].agentSeedStageStatus.Stage = Backup
	testAgents[1].agentSeedStageStatus.Timestamp = time.Now()
	testAgents[1].agentSeedStageStatus.Status = Completed
	testAgents[1].agentSeedStageStatus.Details = "completed backup stage"
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

	testAgents[1].agentSeedStageStatus.Stage = Restore
	testAgents[1].agentSeedStageStatus.Timestamp = time.Now()
	testAgents[1].agentSeedStageStatus.Status = Running
	testAgents[1].agentSeedStageStatus.Details = "running restore stage"
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

	testAgents[1].agentSeedStageStatus.Stage = Restore
	testAgents[1].agentSeedStageStatus.Timestamp = time.Now()
	testAgents[1].agentSeedStageStatus.Status = Completed
	testAgents[1].agentSeedStageStatus.Details = "completed restore stage"
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

	testAgents[2].agentSeedStageStatus.Status = Completed
	testAgents[2].agentSeedStageStatus.Timestamp = time.Now()
	testAgents[2].agentSeedStageStatus.Details = "completed cleanup stage"
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
		if seedState.Hostname == "127.0.0.1" {
			test.S(t).ExpectEquals(seedState.Status, Running)
			test.S(t).ExpectEquals(seedState.Details, "processing cleanup stage")
		}
		if seedState.Hostname == "127.0.0.2" {
			test.S(t).ExpectEquals(seedState.Status, Completed)
			test.S(t).ExpectEquals(seedState.Details, "completed cleanup stage")
		}
	}

	testAgents[1].agentSeedStageStatus.Status = Completed
	testAgents[1].agentSeedStageStatus.Timestamp = time.Now()
	testAgents[1].agentSeedStageStatus.Details = "completed cleanup stage"
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

}
