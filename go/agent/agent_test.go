package agent

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

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
	agentServer          *httptest.Server
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
		agent := &Agent{
			Info: &Info{
				Hostname: fmt.Sprintf("127.0.0.%d", i),
				Port:     3002 + i,
				Token:    "token",
			},
			Data: &Data{},
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
	m.Get("/api/get-agent-data", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.JSON(200, agent.Data)
	})
	testServer := httptest.NewUnstartedServer(m)
	listener, _ := net.Listen("tcp", agentAddress)
	testServer.Listener = listener
	testServer.Start()
	testAgent := &testAgent{
		agent:                agent,
		agentSeedStageStatus: &SeedStageState{},
		agentServer:          testServer,
	}
	return testAgent
}

func registerAgent(t *testing.T, testAgent *testAgent, agentData []byte) (string, error) {
	err := json.Unmarshal(agentData, testAgent.agent.Data)
	if err != nil {
		return "", err
	}
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
	agentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	hostname, err := registerAgent(t, testAgents[1], agentData)
	if err != nil {
		log.Fatale(err)
	}
	test.S(t).ExpectEquals(hostname, testAgents[1].agent.Info.Hostname)
}

func TestReadAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)
	firstAgentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[1], firstAgentData); err != nil {
		log.Fatale(err)
	}
	secondAgentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.2"],"SnaphostHosts":["localhost 127.0.0.2"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":4000000,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[2], secondAgentData); err != nil {
		log.Fatale(err)
	}
	registeredAgents, err := ReadAgents()
	if err != nil {
		log.Fatale(err)
	}
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
	firstAgentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[1], firstAgentData); err != nil {
		log.Fatale(err)
	}
	secondAgentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.2"],"SnaphostHosts":["localhost 127.0.0.2"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":4000000,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[2], secondAgentData); err != nil {
		log.Fatale(err)
	}
	registeredAgents, err := ReadAgentsInfo()
	if err != nil {
		log.Fatale(err)
	}
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
	agentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[1], agentData); err != nil {
		log.Fatale(err)
	}
	registeredAgent, err := ReadAgent("127.0.0.1")
	if err != nil {
		log.Fatale(err)
	}
	expectStructsEquals(t, registeredAgent.Info, testAgents[1].agent.Info)
	expectStructsEquals(t, registeredAgent.Data, testAgents[1].agent.Data)
}

func TestReadAgentInfo(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)
	agentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[1], agentData); err != nil {
		log.Fatale(err)
	}
	registeredAgent, err := ReadAgentInfo("127.0.0.1")
	if err != nil {
		log.Fatale(err)
	}
	expectStructsEquals(t, registeredAgent.Info, testAgents[1].agent.Info)
	expectStructsEquals(t, registeredAgent.Data, &Data{})
}

func TestReadOutdatedAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 2)
	defer teardownTestCase(t)
	config.Config.AgentPollMinutes = 2
	agentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[1], agentData); err != nil {
		log.Fatale(err)
	}
	secondAgentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.2"],"SnaphostHosts":["localhost 127.0.0.2"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":4000000,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[2], secondAgentData); err != nil {
		log.Fatale(err)
	}
	db.ExecOrchestrator("UPDATE host_agent SET last_checked = NOW() WHERE hostname='127.0.0.2'")
	outdatedAgents, err := ReadOutdatedAgents()
	if err != nil {
		log.Fatale(err)
	}
	test.S(t).ExpectEquals(1, len(outdatedAgents))
	expectStructsEquals(t, outdatedAgents[0].Info, testAgents[1].agent.Info)
	expectStructsEquals(t, outdatedAgents[0].Data, &Data{})
}

func TestUpdateAgent(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)
	agentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[1], agentData); err != nil {
		log.Fatale(err)
	}
	registeredAgent, err := ReadAgentInfo("127.0.0.1")
	if err != nil {
		log.Fatale(err)
	}

	testAgents[1].agent.Data.LocalSnapshotsHosts = []string{"127.0.0.10", "127.0.0.12"}
	registeredAgent.Status = Inactive
	registeredAgent.updateAgentStatus()

	if err := registeredAgent.UpdateAgent(); err != nil {
		log.Fatale(err)
	}
	expectStructsEquals(t, registeredAgent.Info, testAgents[1].agent.Info)
	expectStructsEquals(t, registeredAgent.Data, testAgents[1].agent.Data)
	test.S(t).ExpectEquals(registeredAgent.Status, Active)
}

func TestUpdateAgentFailed(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)
	agentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[1], agentData); err != nil {
		log.Fatale(err)
	}
	testAgents[1].agentServer.Close()
	registeredAgent, err := ReadAgentInfo("127.0.0.1")
	if err != nil {
		log.Fatale(err)
	}
	err = registeredAgent.UpdateAgent()
	test.S(t).ExpectNotNil(err)

	registeredAgent, err = ReadAgentInfo("127.0.0.1")
	if err != nil {
		log.Fatale(err)
	}
	test.S(t).ExpectEquals(registeredAgent.Status, Inactive)
}

func TestForgetLongUnseenAgents(t *testing.T) {
	teardownTestCase := prepareTestData(t, 1)
	defer teardownTestCase(t)
	config.Config.UnseenAgentForgetHours = 1
	agentData := []byte(`{"LocalSnapshotsHosts":["127.0.0.1"],"SnaphostHosts":["localhost 127.0.0.1"],"LogicalVolumes":[],"MountPoint":{"Path":"/tmp","Device":"","LVPath":"","FileSystem":"","IsMounted":false,"DiskUsage":0},"BackupDir":"/tmp/bkp","BackupDirDiskFree":38551957504,"MySQLRunning":true,"MySQLDatadir":"/var/lib/mysql/","MySQLDatadirDiskUsed":195537620,"MySQLDatadirDiskFree":38551957504,"MySQLVersion":"5.7","MySQLDatabases":{"akila":{"Engines":["InnoDB"],"Size":15138816},"sakila":{"Engines":["InnoDB"],"Size":15138816},"world":{"Engines":["InnoDB"],"Size":802816}},"AvailiableSeedMethods":{"Mydumper":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Mysqldump":{"BackupSide":"Target","SupportedEngines":["ROCKSDB","MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED","TokuDB"],"BackupToDatadir":false},"Xtrabackup":{"BackupSide":"Source","SupportedEngines":["MRG_MYISAM","CSV","BLACKHOLE","InnoDB","MEMORY","ARCHIVE","MyISAM","FEDERATED"],"BackupToDatadir":true}}}`)
	if _, err := registerAgent(t, testAgents[1], agentData); err != nil {
		log.Fatale(err)
	}
	db.ExecOrchestrator("UPDATE host_agent SET last_seen = last_seen - interval 2 hour WHERE hostname='127.0.0.1'")
	if err := ForgetLongUnseenAgents(); err != nil {
		log.Fatale(err)
	}
	_, err := ReadAgentInfo("127.0.0.1")
	test.S(t).ExpectNotNil(err)
}
