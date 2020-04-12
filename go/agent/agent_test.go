package agent

import (
	"database/sql"
	"flag"
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
	"github.com/martini-contrib/render"
	"github.com/openark/golib/log"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/db"
	"github.com/github/orchestrator/go/inst"
	mysqldrv "github.com/go-sql-driver/mysql"
	"github.com/openark/golib/sqlutils"
	"github.com/ory/dockertest"
	. "gopkg.in/check.v1"
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

var agentsCount = 4

func init() {
	log.SetLevel(log.DEBUG)
}

var testname = flag.String("testname", "", "test names to run")

func Test(t *testing.T) { TestingT(t) }

type AgentTestSuite struct {
	testAgents map[string]*testAgent
	pool       *dockertest.Pool
	containers []*dockertest.Resource
}

var _ = Suite(&AgentTestSuite{})

type testAgent struct {
	agent                  *Agent
	agentSeedStageStatus   *SeedStageState
	agentSeedMetadata      *SeedMetadata
	agentServer            *httptest.Server
	agentMux               *martini.ClassicMartini
	agentMySQLUseGTID      bool
	agentMySQLContainerIP  string
	agentMySQLContainerID  string
	agentContainerResource *dockertest.Resource
}

func (s *AgentTestSuite) SetUpTest(c *C) {
	if len(*testname) > 0 {
		if c.TestName() != fmt.Sprintf("AgentTestSuite.%s", *testname) {
			c.Skip("skipping test due to not matched testname")
			log.Info("skipping test due to not matched testname")
		}
	}
	log.Infof("Setting up test data for test %s", c.TestName())
	if _, err := db.ExecOrchestrator("TRUNCATE TABLE host_agent;"); err != nil {
		log.Fatalf("Unable to truncate host_agent table: %s", err)
	}
	if _, err := db.ExecOrchestrator("TRUNCATE TABLE agent_seed;"); err != nil {
		log.Fatalf("Unable to truncate agent_seed table: %s", err)
	}
	if _, err := db.ExecOrchestrator("TRUNCATE TABLE agent_seed_state;"); err != nil {
		log.Fatalf("Unable to truncate agent_seed_state table: %s", err)
	}

	for _, testAgent := range s.testAgents {
		agentdb, _, err := sqlutils.GetDB(fmt.Sprintf("root:secret@(localhost:%s)/mysql", testAgent.agentContainerResource.GetPort("3306/tcp")))
		if err != nil {
			log.Fatalf("Unable get agent container: %s", err)
		}
		if _, err = agentdb.Exec("DROP database if exists test_repl"); err != nil {
			log.Fatalf("Unable to drop agent test DB: %s", err)
		}
		if _, err = agentdb.Exec("STOP SLAVE;"); err != nil {
			log.Fatalf("Unable to execute stop slave: %s", err)
		}
		if _, err = agentdb.Exec("RESET SLAVE ALL;"); err != nil {
			log.Fatalf("Unable to execute RESET SLAVE: %s", err)
		}
		if _, err = agentdb.Exec("RESET MASTER;"); err != nil {
			log.Fatalf("Unable to execute RESET MASTER: %s", err)
		}
		if _, err = agentdb.Exec("Create database test_repl"); err != nil {
			log.Fatalf("Unable to create agent test DB: %s", err)
		}
		if _, err = agentdb.Exec("Create table test_repl.test(id int);"); err != nil {
			log.Fatalf("Unable create agent test table: %s", err)
		}
		if _, err = agentdb.Exec("insert into test_repl.test(id) VALUES (1), (2), (3), (4);"); err != nil {
			log.Fatalf("Unable to insert data to agent test table: %s", err)
		}

		if err = sqlutils.QueryRowsMap(agentdb, "show master status", func(m sqlutils.RowMap) error {
			var err error
			testAgent.agentSeedMetadata.LogFile = m.GetString("File")
			testAgent.agentSeedMetadata.LogPos = m.GetInt64("Position")
			testAgent.agentSeedMetadata.GtidExecuted = m.GetString("Executed_Gtid_Set")
			return err
		}); err != nil {
			log.Fatalf("Unable to get agent binlog position: %s", err)
		}

		mysqlDatabases := map[string]*MySQLDatabase{
			"sakila": {
				Engines: []Engine{InnoDB},
				Size:    0,
			},
		}
		availiableSeedMethods := map[SeedMethod]*SeedMethodOpts{
			Mydumper: {
				BackupSide:       Target,
				SupportedEngines: []Engine{ROCKSDB, MRG_MYISAM, CSV, BLACKHOLE, InnoDB, MEMORY, ARCHIVE, MyISAM, FEDERATED, TokuDB},
				BackupToDatadir:  false,
			},
		}
		testAgent.agent = &Agent{
			Info: &Info{
				Hostname:  testAgent.agent.Info.Hostname,
				Port:      testAgent.agent.Info.Port,
				MySQLPort: testAgent.agent.Info.MySQLPort,
				Token:     "token",
			},
			Data: &Data{
				LocalSnapshotsHosts: testAgent.agent.Data.LocalSnapshotsHosts,
				SnaphostHosts:       testAgent.agent.Data.SnaphostHosts,
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
				BackupDirDiskUsed:     0,
				BackupDirDiskFree:     10000,
				MySQLRunning:          true,
				MySQLDatadir:          "/var/lib/mysql",
				MySQLDatadirDiskUsed:  10,
				MySQLDatadirDiskFree:  10000,
				MySQLDatabases:        mysqlDatabases,
				AvailiableSeedMethods: availiableSeedMethods,
			},
		}
	}
}

func (s *AgentTestSuite) SetUpSuite(c *C) {
	var testAgents = make(map[string]*testAgent)
	s.testAgents = testAgents
	s.containers = []*dockertest.Resource{}
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	s.pool = pool

	// pulls an image, creates a container based on it and runs it
	log.Info("Creating docker container for Orchestrator DB")
	resource, err := pool.Run("mysql", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	s.containers = append(s.containers, resource)

	// configure Orchestrator
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
	config.Config.MySQLTopologyUser = "orc_topology"
	config.Config.MySQLTopologyPassword = "orc_topologypassword@"
	config.Config.ReplicationCredentialsQuery = "select username, password from orchestrator_meta.replication;"
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
	if err != nil {
		log.Fatalf("Unable to create orchestrator DB: %s", err)
	}

	// init few functions nessesary for test process
	inst.InitializeInstanceDao()
	InitHttpClient()
	go func() {
		for seededAgent := range SeededAgents {
			instanceKey := &inst.InstanceKey{Hostname: seededAgent.Info.Hostname, Port: int(seededAgent.Info.MySQLPort)}
			log.Infof("%+v", instanceKey)
		}
	}()

	// create mocks for agents
	log.Info("Creating Orchestrator agents mocks")
	for i := 1; i <= agentsCount; i++ {
		mysqlDatabases := map[string]*MySQLDatabase{
			"sakila": {
				Engines: []Engine{InnoDB},
				Size:    0,
			},
		}
		availiableSeedMethods := map[SeedMethod]*SeedMethodOpts{
			Mydumper: {
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
				MySQLDatabases:        mysqlDatabases,
				AvailiableSeedMethods: availiableSeedMethods,
			},
		}
		if err := s.createTestAgent(agent); err != nil {
			log.Fatalf("unable to create test agent: %s", err)
		}
	}
}

func (s *AgentTestSuite) createTestAgent(agent *Agent) error {
	agentAddress := fmt.Sprintf("%s:%d", agent.Info.Hostname, agent.Info.Port)
	m := martini.Classic()
	m.Use(render.Renderer())
	testAgent := &testAgent{
		agent:                agent,
		agentSeedStageStatus: &SeedStageState{},
		agentSeedMetadata:    &SeedMetadata{},
		agentMySQLUseGTID:    true,
	}
	m.Get("/api/get-agent-data", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.JSON(200, testAgent.agent.Data)
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
	m.Get("/api/post-seed-cmd/:seedID", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.Text(202, "Done")
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
	s.testAgents[agent.Info.Hostname] = testAgent
	return s.createTestAgentsMySQLServers(testAgent)
}

func (s *AgentTestSuite) createTestAgentsMySQLServers(agent *testAgent) error {
	var testdb *sql.DB
	rand.Seed(time.Now().UnixNano())
	serverID := rand.Intn(100000000)

	dockerCmd := []string{"mysqld", fmt.Sprintf("--server-id=%d", serverID), "--log-bin=/var/lib/mysql/mysql-bin"}
	log.Infof("Creating docker container for agent %s", agent.agent.Info.Hostname)

	if agent.agentMySQLUseGTID {
		dockerCmd = append(dockerCmd, "--enforce-gtid-consistency=ON", "--gtid-mode=ON")
	}
	resource, err := s.pool.RunWithOptions(&dockertest.RunOptions{Repository: "mysql", Tag: "5.7", Env: []string{"MYSQL_ROOT_PASSWORD=secret"}, Cmd: dockerCmd, CapAdd: []string{"NET_ADMIN", "NET_RAW"}})
	if err != nil {
		return fmt.Errorf("Could not connect to docker: %s", err)
	}

	agent.agentMySQLContainerIP = resource.Container.NetworkSettings.Networks["bridge"].IPAddress
	agent.agentMySQLContainerID = resource.Container.ID
	agent.agentContainerResource = resource

	s.containers = append(s.containers, resource)

	cmd := exec.Command("docker", "exec", "-i", agent.agentMySQLContainerID, "apt-get", "update")
	if err := cmd.Run(); err != nil {
		return err
	}
	cmd = exec.Command("docker", "exec", "-i", agent.agentMySQLContainerID, "apt-get", "install", "-y", "iptables")
	if err := cmd.Run(); err != nil {
		return err
	}

	if err := s.pool.Retry(func() error {
		var err error
		testdb, _, err = sqlutils.GetDB(fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp")))
		mysqldrv.SetLogger(logger.New(ioutil.Discard, "discard", 1))
		if err != nil {
			return err
		}
		return testdb.Ping()
	}); err != nil {
		return fmt.Errorf("Could not connect to docker: %s", err)
	}

	if _, err = testdb.Exec("Create database orchestrator_meta;"); err != nil {
		return err
	}
	if _, err = testdb.Exec("Create table orchestrator_meta.replication (`username` varchar(128) CHARACTER SET ascii NOT NULL DEFAULT '',`password` varchar(128) CHARACTER SET ascii NOT NULL DEFAULT '',PRIMARY KEY (`username`,`password`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;"); err != nil {
		return err
	}
	if _, err = testdb.Exec("CREATE USER `slave`@`%` IDENTIFIED BY 'slavepassword@';"); err != nil {
		return err
	}
	if _, err = testdb.Exec("GRANT REPLICATION SLAVE ON *.* TO `slave`@`%`"); err != nil {
		return err
	}
	if _, err = testdb.Exec("CREATE USER `orc_topology`@`%` IDENTIFIED BY 'orc_topologypassword@';"); err != nil {
		return err
	}
	if _, err = testdb.Exec("GRANT ALL PRIVILEGES ON *.* TO `orc_topology`@`%`"); err != nil {
		return err
	}
	if _, err = testdb.Exec("Insert into orchestrator_meta.replication(username, password) VALUES ('slave', 'slavepassword@');"); err != nil {
		return err
	}

	if _, err = testdb.Exec("Create database test_repl"); err != nil {
		return err
	}
	if _, err = testdb.Exec("Create table test_repl.test(id int);"); err != nil {
		return err
	}
	if _, err = testdb.Exec("insert into test_repl.test(id) VALUES (1), (2), (3), (4);"); err != nil {
		return err
	}

	if err = sqlutils.QueryRowsMap(testdb, "show master status", func(m sqlutils.RowMap) error {
		var err error
		agent.agentSeedMetadata.LogFile = m.GetString("File")
		agent.agentSeedMetadata.LogPos = m.GetInt64("Position")
		agent.agentSeedMetadata.GtidExecuted = m.GetString("Executed_Gtid_Set")
		return err
	}); err != nil {
		return err
	}

	mysqlPort, err := strconv.Atoi(resource.GetPort("3306/tcp"))
	if err != nil {
		return err
	}
	agent.agent.Info.MySQLPort = mysqlPort

	return nil
}

// this is needed in order to redirect from host machine ip 172.0.0.x to container ip for replication to work
func (s *AgentTestSuite) createIPTablesRulesForReplication(targetAgent *testAgent, sourceAgent *testAgent) error {
	cmd := exec.Command("docker", "exec", "-i", targetAgent.agentMySQLContainerID, "iptables", "-t", "nat", "-I", "OUTPUT", "-p", "tcp", "-o", "eth0", "--dport", fmt.Sprintf("%d", sourceAgent.agent.Info.MySQLPort), "-j", "DNAT", "--to-destination", fmt.Sprintf("%s:3306", sourceAgent.agentMySQLContainerIP))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", string(out))
	}
	cmd = exec.Command("docker", "exec", "-i", targetAgent.agentMySQLContainerID, "/bin/sh", "-c", fmt.Sprintf("echo %s %s >> /etc/hosts", sourceAgent.agentMySQLContainerIP, sourceAgent.agent.Info.Hostname))
	out, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", string(out))
	}
	return nil
}

func (s *AgentTestSuite) TearDownSuite(c *C) {
	log.Info("Teardown test data")
	for _, container := range s.containers {
		if err := s.pool.Purge(container); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	}
	for _, testAgent := range s.testAgents {
		testAgent.agentServer.Close()
	}
}

func (s *AgentTestSuite) registerAgents(c *C) {
	for _, testAgent := range s.testAgents {
		hostname, err := RegisterAgent(testAgent.agent.Info)
		c.Assert(err, IsNil)
		c.Assert(hostname, Equals, testAgent.agent.Info.Hostname)
	}
}

func (s *AgentTestSuite) getSeedAgents(c *C, targetTestAgent *testAgent, sourceTestAgent *testAgent) (targetAgent *Agent, sourceAgent *Agent) {
	s.registerAgents(c)
	targetAgent, err := ReadAgent(targetTestAgent.agent.Info.Hostname)
	c.Assert(err, IsNil)

	sourceAgent, err = ReadAgent(sourceTestAgent.agent.Info.Hostname)
	c.Assert(err, IsNil)

	return targetAgent, sourceAgent
}

func (s *AgentTestSuite) readSeed(c *C, seedID int64, targetHostname string, sourceHostname string, seedMethod SeedMethod, backupSide SeedSide, status SeedStatus, stage SeedStage, retries int) *Seed {
	seed, err := ReadSeed(seedID)
	c.Assert(err, IsNil)
	c.Assert(seed.TargetHostname, Equals, targetHostname)
	c.Assert(seed.SourceHostname, Equals, sourceHostname)
	c.Assert(seed.SeedMethod, Equals, seedMethod)
	c.Assert(seed.BackupSide, Equals, backupSide)
	c.Assert(seed.Status, Equals, status)
	c.Assert(seed.Stage, Equals, stage)
	c.Assert(seed.Retries, Equals, retries)
	return seed
}

func (s *AgentTestSuite) readSeedStageStates(c *C, seed *Seed, stateRecords int, targetTestAgent *testAgent, sourceTestAgent *testAgent) {
	seedStates, err := seed.ReadSeedStageStates()
	c.Assert(err, IsNil)
	for _, seedState := range seedStates[:stateRecords] {
		for _, agent := range []*testAgent{targetTestAgent, sourceTestAgent} {
			if seedState.Hostname == agent.agent.Info.Hostname {
				c.Assert(seedState.SeedID, Equals, agent.agentSeedStageStatus.SeedID)
				c.Assert(seedState.Stage, Equals, agent.agentSeedStageStatus.Stage)
				c.Assert(seedState.Status, Equals, agent.agentSeedStageStatus.Status)
			}
		}
	}
}

func (s *AgentTestSuite) TestAgentRegistration(c *C) {
	s.registerAgents(c)
}

func (s *AgentTestSuite) TestReadAgents(c *C) {
	s.registerAgents(c)

	registeredAgents, err := ReadAgents()
	c.Assert(err, IsNil)
	c.Assert(registeredAgents, HasLen, 4)

	for _, testAgent := range s.testAgents {
		for _, registeredAgent := range registeredAgents {
			if registeredAgent.Info.Port == testAgent.agent.Info.Port && registeredAgent.Info.Hostname == testAgent.agent.Info.Hostname {
				c.Assert(registeredAgent.Info, DeepEquals, testAgent.agent.Info)
				c.Assert(registeredAgent.Data, DeepEquals, testAgent.agent.Data)
			}
		}
	}
}

func (s *AgentTestSuite) TestReadAgentsInfo(c *C) {
	s.registerAgents(c)

	registeredAgents, err := ReadAgentsInfo()
	c.Assert(err, IsNil)
	c.Assert(registeredAgents, HasLen, 4)

	for _, testAgent := range s.testAgents {
		for _, registeredAgent := range registeredAgents {
			if registeredAgent.Info.Port == testAgent.agent.Info.Port && registeredAgent.Info.Hostname == testAgent.agent.Info.Hostname {
				c.Assert(registeredAgent.Info, DeepEquals, testAgent.agent.Info)
				c.Assert(registeredAgent.Data, DeepEquals, &Data{})
			}
		}
	}
}

func (s *AgentTestSuite) TestReadAgent(c *C) {
	testAgent := s.testAgents["agent1"]

	s.registerAgents(c)

	registeredAgent, err := ReadAgent(testAgent.agent.Info.Hostname)
	c.Assert(err, IsNil)
	c.Assert(registeredAgent.Info, DeepEquals, testAgent.agent.Info)
	c.Assert(registeredAgent.Data, DeepEquals, testAgent.agent.Data)
}

func (s *AgentTestSuite) TestReadAgentInfo(c *C) {
	testAgent := s.testAgents["agent1"]

	s.registerAgents(c)

	registeredAgent, err := ReadAgentInfo(testAgent.agent.Info.Hostname)
	c.Assert(err, IsNil)
	c.Assert(registeredAgent.Info, DeepEquals, testAgent.agent.Info)
	c.Assert(registeredAgent.Data, DeepEquals, &Data{})
}

func (s *AgentTestSuite) TestReadOutdatedAgents(c *C) {
	config.Config.AgentPollMinutes = 2

	s.registerAgents(c)

	outdatedAgents := []*testAgent{s.testAgents["agent1"]}
	upToDateAgents := []*testAgent{s.testAgents["agent2"], s.testAgents["agent3"], s.testAgents["agent4"]}

	for _, outdatedAgent := range outdatedAgents {
		db.ExecOrchestrator(fmt.Sprintf("UPDATE host_agent SET last_checked = NOW() - interval 60 minute WHERE hostname='%s'", outdatedAgent.agent.Info.Hostname))
	}

	for _, upToDateAgent := range upToDateAgents {
		db.ExecOrchestrator(fmt.Sprintf("UPDATE host_agent SET last_checked = NOW() WHERE hostname='%s'", upToDateAgent.agent.Info.Hostname))
	}

	outdatedOrchestratorAgents, err := ReadOutdatedAgents()
	c.Assert(err, IsNil)
	c.Assert(outdatedOrchestratorAgents, HasLen, 1)
	c.Assert(outdatedOrchestratorAgents[0].Info, DeepEquals, outdatedAgents[0].agent.Info)
	c.Assert(outdatedOrchestratorAgents[0].Data, DeepEquals, &Data{})
}

func (s *AgentTestSuite) TestReadAgentsHosts(c *C) {
	s.registerAgents(c)

	agents, err := ReadAgentsHosts("agent1")

	c.Assert(err, IsNil)
	c.Assert(agents, HasLen, 1)
	c.Assert(agents[0], Equals, "agent1")
}

func (s *AgentTestSuite) TestReadAgentsPaged(c *C) {
	s.registerAgents(c)

	agents, err := ReadAgentsPaged(0)

	c.Assert(err, IsNil)
	c.Assert(agents, HasLen, 4)
}

func (s *AgentTestSuite) TestReadAgentsInStatusPaged(c *C) {
	s.registerAgents(c)

	agents, err := ReadAgentsInStatusPaged(Active, 0)

	c.Assert(err, IsNil)
	c.Assert(agents, HasLen, 4)
}

func (s *AgentTestSuite) TestUpdateAgent(c *C) {
	testAgent := s.testAgents["agent1"]

	s.registerAgents(c)

	registeredAgent, err := ReadAgentInfo(testAgent.agent.Info.Hostname)
	c.Assert(err, IsNil)

	testAgent.agent.Data.LocalSnapshotsHosts = []string{"127.0.0.10", "127.0.0.12"}
	registeredAgent.Status = Inactive

	err = registeredAgent.updateAgentStatus()
	c.Assert(err, IsNil)

	err = registeredAgent.UpdateAgent()
	c.Assert(err, IsNil)

	c.Assert(registeredAgent.Info, DeepEquals, testAgent.agent.Info)
	c.Assert(registeredAgent.Data, DeepEquals, testAgent.agent.Data)
	c.Assert(registeredAgent.Status, Equals, Active)
}

func (s *AgentTestSuite) TestUpdateAgentFailed(c *C) {
	testAgent := s.testAgents["agent1"]

	s.registerAgents(c)
	testAgent.agentServer.Close()

	registeredAgent, err := ReadAgentInfo(testAgent.agent.Info.Hostname)
	c.Assert(err, IsNil)

	err = registeredAgent.UpdateAgent()
	c.Assert(err, NotNil)

	registeredAgent, err = ReadAgentInfo(testAgent.agent.Info.Hostname)
	c.Assert(err, IsNil)
	c.Assert(registeredAgent.Status, Equals, Inactive)
}

func (s *AgentTestSuite) TestForgetLongUnseenAgents(c *C) {
	config.Config.UnseenAgentForgetHours = 1
	testAgent := s.testAgents["agent1"]

	s.registerAgents(c)
	db.ExecOrchestrator(fmt.Sprintf("UPDATE host_agent SET last_seen = last_seen - interval 2 hour WHERE hostname='%s'", testAgent.agent.Info.Hostname))

	err := ForgetLongUnseenAgents()
	c.Assert(err, IsNil)

	_, err = ReadAgentInfo(testAgent.agent.Info.Hostname)
	c.Assert(err, NotNil)
}

func (s *AgentTestSuite) TestReadSeeds(c *C) {
	targetTestAgent1 := s.testAgents["agent1"]
	sourceTestAgent1 := s.testAgents["agent2"]
	targetTestAgent2 := s.testAgents["agent3"]
	sourceTestAgent2 := s.testAgents["agent4"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent1, sourceAgent1 := s.getSeedAgents(c, targetTestAgent1, sourceTestAgent1)
	targetAgent2, sourceAgent2 := s.getSeedAgents(c, targetTestAgent2, sourceTestAgent2)

	_, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	c.Assert(err, IsNil)

	_, err = NewSeed("Mydumper", targetAgent2, sourceAgent2)
	c.Assert(err, IsNil)

	seeds, err := ReadSeeds()
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 2)
}

func (s *AgentTestSuite) TestReadSeedsPaged(c *C) {
	targetTestAgent1 := s.testAgents["agent1"]
	sourceTestAgent1 := s.testAgents["agent2"]
	targetTestAgent2 := s.testAgents["agent3"]
	sourceTestAgent2 := s.testAgents["agent4"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent1, sourceAgent1 := s.getSeedAgents(c, targetTestAgent1, sourceTestAgent1)
	targetAgent2, sourceAgent2 := s.getSeedAgents(c, targetTestAgent2, sourceTestAgent2)

	_, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	c.Assert(err, IsNil)

	_, err = NewSeed("Mydumper", targetAgent2, sourceAgent2)
	c.Assert(err, IsNil)

	seeds, err := ReadSeedsPaged(0)
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 2)
}

func (s *AgentTestSuite) TestReadSeedsForTargetAgentPaged(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)

	seeds, err := ReadSeedsForTargetAgentPaged(targetAgent, 0)
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 1)
	c.Assert(seeds[0].TargetHostname, Equals, targetAgent.Info.Hostname)
}

func (s *AgentTestSuite) TestReadSeedsForSourceAgentPaged(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)

	seeds, err := ReadSeedsForSourceAgentPaged(sourceAgent, 0)
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 1)
	c.Assert(seeds[0].SourceHostname, Equals, sourceAgent.Info.Hostname)
}

func (s *AgentTestSuite) TestReadSeedsInStatusPaged(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)

	seeds, err := ReadSeedsInStatusPaged(Scheduled, 0)
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 1)
	seeds, err = ReadSeedsInStatusPaged(Running, 0)
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 0)
}

func (s *AgentTestSuite) TestReadSeedsForAgent(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)

	seeds, err := ReadSeedsForAgent(sourceAgent, "")
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 1)
	c.Assert(seeds[0].SourceHostname, Equals, sourceAgent.Info.Hostname)
	seeds, err = ReadSeedsForAgent(targetAgent, "")
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 1)
	c.Assert(seeds[0].TargetHostname, Equals, targetAgent.Info.Hostname)
}

func (s *AgentTestSuite) TestNewSeed(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))
}

func (s *AgentTestSuite) TestNewSeedWrongSeedMethod(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("test", targetAgent, sourceAgent)
	c.Assert(err, NotNil)
}

func (s *AgentTestSuite) TestNewSeedSeedItself(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent1"]

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, NotNil)
}

func (s *AgentTestSuite) TestNewSeedUnsupportedSeedMethod(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Mysqldump", targetAgent, sourceAgent)
	c.Assert(err, NotNil)
}

func (s *AgentTestSuite) TestNewSeedUnsupportedSeedMethodForDB(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	sourceTestAgent.agent.Data.AvailiableSeedMethods[Xtrabackup] = &SeedMethodOpts{
		BackupSide:       Target,
		SupportedEngines: []Engine{MRG_MYISAM, CSV, BLACKHOLE, InnoDB, MEMORY, ARCHIVE, MyISAM, FEDERATED, TokuDB},
	}
	sourceTestAgent.agent.Data.MySQLDatabases["test"] = &MySQLDatabase{
		Engines: []Engine{ROCKSDB},
		Size:    0,
	}

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Xtrabackup", targetAgent, sourceAgent)
	c.Assert(err, NotNil)
}

func (s *AgentTestSuite) TestNewSeedAgentHadActiveSeed(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)

	_, err = NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, NotNil)
}

func (s *AgentTestSuite) TestNewSeedNotEnoughSpaceInMySQLDatadir(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]
	targetTestAgent.agent.Data.MySQLDatadirDiskFree = 10
	sourceTestAgent.agent.Data.MySQLDatadirDiskUsed = 1000

	// register agents to Orchestrator
	s.registerAgents(c)
	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	_, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, NotNil)
}

func (s *AgentTestSuite) TestReadSeed(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)
}

func (s *AgentTestSuite) TestReadActiveSeeds(c *C) {
	targetTestAgent1 := s.testAgents["agent1"]
	sourceTestAgent1 := s.testAgents["agent2"]
	targetTestAgent2 := s.testAgents["agent3"]
	sourceTestAgent2 := s.testAgents["agent4"]

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent1, sourceAgent1 := s.getSeedAgents(c, targetTestAgent1, sourceTestAgent1)
	targetAgent2, sourceAgent2 := s.getSeedAgents(c, targetTestAgent2, sourceTestAgent2)

	seedID, err := NewSeed("Mydumper", targetAgent1, sourceAgent1)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	seedID, err = NewSeed("Mydumper", targetAgent2, sourceAgent2)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(2))

	completedSeed := s.readSeed(c, seedID, targetAgent2.Info.Hostname, sourceAgent2.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)
	completedSeed.Status = Completed
	completedSeed.updateSeedData()

	seeds, err := ReadActiveSeeds()
	c.Assert(err, IsNil)
	c.Assert(seeds, HasLen, 1)

	c.Assert(seeds[0].TargetHostname, Equals, targetTestAgent1.agent.Info.Hostname)
	c.Assert(seeds[0].SourceHostname, Equals, sourceTestAgent1.agent.Info.Hostname)
	c.Assert(seeds[0].SeedMethod, Equals, Mydumper)
	c.Assert(seeds[0].BackupSide, Equals, Target)
	c.Assert(seeds[0].Status, Equals, Scheduled)
	c.Assert(seeds[0].Stage, Equals, Prepare)
	c.Assert(seeds[0].Retries, Equals, 0)
}

func (s *AgentTestSuite) TestReadActiveSeedsForAgent(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	for _, agent := range []*Agent{targetAgent, sourceAgent} {
		seeds, err := ReadActiveSeedsForAgent(agent)
		c.Assert(err, IsNil)
		c.Assert(seeds, HasLen, 1)
		c.Assert(seeds[0].TargetHostname, Equals, targetTestAgent.agent.Info.Hostname)
		c.Assert(seeds[0].SourceHostname, Equals, sourceTestAgent.agent.Info.Hostname)
		c.Assert(seeds[0].SeedMethod, Equals, Mydumper)
		c.Assert(seeds[0].BackupSide, Equals, Target)
		c.Assert(seeds[0].Status, Equals, Scheduled)
		c.Assert(seeds[0].Stage, Equals, Prepare)
		c.Assert(seeds[0].Retries, Equals, 0)
	}

	activeSeed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)
	activeSeed.Status = Completed
	activeSeed.updateSeedData()

	for _, agent := range []*Agent{targetAgent, sourceAgent} {
		seeds, err := ReadActiveSeedsForAgent(agent)
		c.Assert(err, IsNil)
		c.Assert(seeds, HasLen, 0)
	}

}

func (s *AgentTestSuite) TestReedSeedAgents(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	seedTargetAgent, seedSourceAgent, err := seed.readSeadAgents(false)
	c.Assert(err, IsNil)
	c.Assert(targetAgent.Info, DeepEquals, seedTargetAgent.Info)
	c.Assert(sourceAgent.Info, DeepEquals, seedSourceAgent.Info)
}

func (s *AgentTestSuite) TestProcessSeeds(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// create iptables rules for replication
	err := s.createIPTablesRulesForReplication(targetTestAgent, sourceTestAgent)
	c.Assert(err, IsNil)

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// ProcessSeeds. Orchestrator will ask agents to start Prepare stage. So it's state would change from Scheduled to Running
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Prepare, 0)

	// simulate that prepare stage is running on both agents
	for _, agent := range []*testAgent{targetTestAgent, sourceTestAgent} {
		agent.agentSeedStageStatus.SeedID = seedID
		agent.agentSeedStageStatus.Stage = Prepare
		agent.agentSeedStageStatus.Hostname = agent.agent.Info.Hostname
		agent.agentSeedStageStatus.Timestamp = time.Now()
		agent.agentSeedStageStatus.Status = Running
		agent.agentSeedStageStatus.Details = "processing prepare stage"
	}

	// check that SeedStageStates in Orchestator DB are the same, that were read from agents
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. As seed is in status Running, Orchestrator will ask agent about seed stage status. As it's not Completed or Errored, it will just record SeedStageState in DB.
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Prepare, 0)
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)

	// now simulate that target agent completed Prepare stage
	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Details = "completed prepare stage"

	// ProcessSeeds. We have stage: Prepare and status: Running as only target agent completed Prepare stage, but in SeedStageStates we will have one Completed record and one Running
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Prepare, 0)
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)

	// now simulate that source agent completed Prepare stage
	sourceTestAgent.agentSeedStageStatus.Status = Completed
	sourceTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	sourceTestAgent.agentSeedStageStatus.Details = "completed prepare stage"

	// ProcessSeeds. Now both agents completed Prepare stage, so Orchestrator will move Seed to Backup stage with status Scheduled. And we will have SeedStageState records with stage Prepare and status Completed
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Backup, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. Orchestrator will ask agent to start Backup stage. So it's state would change from Scheduled to Running
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Backup, 0)

	// simulate that Backup stage is running on target agent(as seedSide for Mydumper method is Target)
	targetTestAgent.agentSeedStageStatus.Stage = Backup
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "running backup stage"
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. TargetAgent is still running Backup Stage, so seed state won't change and we will have one SeedStageState record from target agent
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Backup, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// now simulate that target agent completed Backup stage
	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Details = "completed backup stage"

	// ProcessSeeds. Target agent had completed Backup stage, so Orchestrator will move Seed to Restore stage with status Scheduled. And we will have one SeedStageState record with stage Backup and status Completed
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Restore, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. Orchestrator will ask agent to start Restore stage. So it's state would change from Scheduled to Running
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Restore, 0)

	// simulate that Backup stage is running on target agent(as restore is always processed by targetAgent)
	targetTestAgent.agentSeedStageStatus.Stage = Restore
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "running restore stage"
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. TargetAgent is still running Restore Stage, so seed state won't change and we will have one SeedStageState record from target agent
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Restore, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// now simulate that target agent completed Restore stage
	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Details = "completed restore stage"

	//register agents one more time to update mysqlport
	s.registerAgents(c)

	// update targetAgent with source agent replication pos\gtid
	targetTestAgent.agentSeedMetadata = sourceTestAgent.agentSeedMetadata

	// ProcessSeeds. Target agent had completed Restore stage, so Orchestrator will move Seed to ConnectSlave stage with status Scheduled. And we will have one SeedStageState record with stage Restore and status Completed
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, ConnectSlave, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. Orchestrator will ask agents to start ConnectSlave stage. So it's state would change from Scheduled to Running
	ProcessSeeds()
	targetTestAgent.agentSeedStageStatus.Stage = ConnectSlave
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "running connectSlave stage"
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, ConnectSlave, 0)

	// targetAgent completed ConnectSlave stage
	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Details = "completed connectSlave stage"

	// Read seed one more time. It's state should be cleanup and status Scheduled and we should have one SeedStageState record with stage ConnectSlave and status Completed
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Cleanup, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. Orchestrator will ask agents to start Cleanup stage. So it's state would change from Scheduled to Running
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Cleanup, 0)

	// simulate that cleanup stage is running on both agents
	for _, agent := range []*testAgent{targetTestAgent, sourceTestAgent} {
		agent.agentSeedStageStatus.Stage = Cleanup
		agent.agentSeedStageStatus.Timestamp = time.Now()
		agent.agentSeedStageStatus.Status = Running
		agent.agentSeedStageStatus.Details = "processing cleanup stage"
	}

	// check that SeedStageStates in Orchestator DB are the same, that were read from agents
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. Both agents are still running Cleanup Stage, so seed state won't change and we will have 2 SeedStageState record from each agent
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Cleanup, 0)
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)

	// now simulate that source agent completed Cleanup stage
	sourceTestAgent.agentSeedStageStatus.Status = Completed
	sourceTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	sourceTestAgent.agentSeedStageStatus.Details = "completed cleanup stage"

	// ProcessSeeds. We have stage: Cleanup and status: Running as only source agent completed Cleanup stage, but in SeedStageStates we will have one Completed record and one Running
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Cleanup, 0)
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)

	// now simulate that target agent completed Prepare stage
	targetTestAgent.agentSeedStageStatus.Status = Completed
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Details = "completed cleanup stage"

	// ProcessSeeds. This is the last stage, so we will have seed in completed Status.
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Completed, Cleanup, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)
}

func (s *AgentTestSuite) TestProcessSeedsErrorOnPrepare(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	config.Config.MaxRetriesForSeedStage = 2

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)
	ProcessSeeds()

	// One agent is OK, another had error
	for _, agent := range []*testAgent{targetTestAgent, sourceTestAgent} {
		agent.agentSeedStageStatus.SeedID = seedID
		agent.agentSeedStageStatus.Stage = Prepare
		agent.agentSeedStageStatus.Hostname = agent.agent.Info.Hostname
		agent.agentSeedStageStatus.Timestamp = time.Now()
	}
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "processing prepare stage"
	sourceTestAgent.agentSeedStageStatus.Status = Error
	sourceTestAgent.agentSeedStageStatus.Details = "error processing prepare stage"

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Error, Prepare, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. After error our seed will be moved back to Prepare stage Scheduled status and retries increased
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 1)

	// so our SeedStages should also be in Scheduled status
	sourceTestAgent.agentSeedStageStatus.Status = Scheduled
	targetTestAgent.agentSeedStageStatus.Status = Scheduled
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)
}

func (s *AgentTestSuite) TestProcessSeedsErrorOnBackup(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	config.Config.MaxRetriesForSeedStage = 2

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Scheduled and stage to Backup
	seed.Status = Scheduled
	seed.Stage = Backup
	seed.updateSeedData()
	ProcessSeeds()

	// As Mydumper is executed on target, set in status to Error
	targetTestAgent.agentSeedStageStatus.SeedID = seedID
	targetTestAgent.agentSeedStageStatus.Stage = Backup
	targetTestAgent.agentSeedStageStatus.Hostname = targetTestAgent.agent.Info.Hostname
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Error
	targetTestAgent.agentSeedStageStatus.Details = "error processing backup stage"

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Error, Backup, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. After error our seed will be moved back to Prepare stage Scheduled status and retries increased (due to backup stage is always moved to scheduled)
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 1)

	// so our SeedStage should also be in Scheduled status
	targetTestAgent.agentSeedStageStatus.Status = Scheduled
	targetTestAgent.agentSeedStageStatus.Stage = Prepare
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)
}

func (s *AgentTestSuite) TestProcessSeedsErrorOnRestore(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	config.Config.MaxRetriesForSeedStage = 2

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Scheduled and stage to Restore
	seed.Status = Scheduled
	seed.Stage = Restore
	seed.updateSeedData()
	ProcessSeeds()

	// As Restore is always executed on target, set in status to Error
	targetTestAgent.agentSeedStageStatus.SeedID = seedID
	targetTestAgent.agentSeedStageStatus.Stage = Restore
	targetTestAgent.agentSeedStageStatus.Hostname = targetTestAgent.agent.Info.Hostname
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Error
	targetTestAgent.agentSeedStageStatus.Details = "error processing restore stage"

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Error, Restore, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. After error our seed will be moved back to Restore stage Scheduled status and retries increased (due to backup stage is always moved to scheduled)
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Restore, 1)

	// so our SeedStage should also be in Scheduled status
	targetTestAgent.agentSeedStageStatus.Status = Scheduled
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)
}

func (s *AgentTestSuite) TestProcessSeedsErrorOnCleanup(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	config.Config.MaxRetriesForSeedStage = 2

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Running and stage to Cleanup
	seed.Status = Running
	seed.Stage = Cleanup
	seed.updateSeedData()

	// One agent is OK, another had error
	for _, agent := range []*testAgent{targetTestAgent, sourceTestAgent} {
		agent.agentSeedStageStatus.SeedID = seedID
		agent.agentSeedStageStatus.Stage = Cleanup
		agent.agentSeedStageStatus.Hostname = agent.agent.Info.Hostname
		agent.agentSeedStageStatus.Timestamp = time.Now()
	}
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "processing cleanup stage"
	sourceTestAgent.agentSeedStageStatus.Status = Error
	sourceTestAgent.agentSeedStageStatus.Details = "error processing cleanup stage"

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Error, Cleanup, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. After error our seed will be moved back to Cleanup stage Scheduled status and retries increased
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Cleanup, 1)

	// so our SeedStages should also be in Scheduled status
	sourceTestAgent.agentSeedStageStatus.Status = Scheduled
	targetTestAgent.agentSeedStageStatus.Status = Scheduled
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)
}

func (s *AgentTestSuite) TestProcessSeedsErrorOnConnectSlave(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	config.Config.MaxRetriesForSeedStage = 2

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Running and stage to ConnectSlave
	seed.Status = Running
	seed.Stage = ConnectSlave
	seed.updateSeedData()

	targetTestAgent.agentSeedStageStatus.SeedID = seedID
	targetTestAgent.agentSeedStageStatus.Stage = ConnectSlave
	targetTestAgent.agentSeedStageStatus.Hostname = targetTestAgent.agent.Info.Hostname
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Error
	targetTestAgent.agentSeedStageStatus.Details = "error processing ConnectSlave stage"

	// As ConnectSlave requires agent metadata, set it
	targetTestAgent.agentSeedMetadata = &SeedMetadata{
		LogFile:      "",
		LogPos:       10,
		GtidExecuted: "",
	}

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Error, ConnectSlave, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. After error our seed will be moved back to ConnectSlave stage
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, ConnectSlave, 1)

	// so our SeedStage should also be in Scheduled status
	targetTestAgent.agentSeedStageStatus.Status = Scheduled
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)
}

func (s *AgentTestSuite) TestProcessSeedsFailed(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// MaxRetriesForStage set to 0, so it will be failed after first error
	config.Config.MaxRetriesForSeedStage = 0

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Scheduled and stage to Restore
	seed.Status = Scheduled
	seed.Stage = Restore
	seed.updateSeedData()
	ProcessSeeds()

	// As Restore is always executed on target, set in status to Error
	targetTestAgent.agentSeedStageStatus.SeedID = seedID
	targetTestAgent.agentSeedStageStatus.Stage = Restore
	targetTestAgent.agentSeedStageStatus.Hostname = targetTestAgent.agent.Info.Hostname
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Error
	targetTestAgent.agentSeedStageStatus.Details = "error processing restore stage"

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Error, Restore, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. After error our seed will be Failed
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Failed, Restore, 0)

	// so our SeedStage should also be in Scheduled status
	targetTestAgent.agentSeedStageStatus.Status = Failed
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	activeSeeds, err := ReadActiveSeeds()
	c.Assert(err, IsNil)
	c.Assert(activeSeeds, HasLen, 0)
}

func (s *AgentTestSuite) TestAbortSeedScheduled(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	err = seed.AbortSeed()
	c.Assert(err, IsNil)

	activeSeeds, err := ReadActiveSeeds()
	c.Assert(err, IsNil)
	c.Assert(activeSeeds, HasLen, 0)
}

func (s *AgentTestSuite) TestAbortSeedCompleted(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// MaxRetriesForStage set to 0, so it will be failed after first error
	config.Config.MaxRetriesForSeedStage = 0

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Completed and stage to Cleanup
	seed.Status = Completed
	seed.Stage = Cleanup
	seed.updateSeedData()

	err = seed.AbortSeed()
	c.Assert(err, NotNil)
}

func (s *AgentTestSuite) TestAbortSeedRestore(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// MaxRetriesForStage set to 0, so it will be failed after first error
	config.Config.MaxRetriesForSeedStage = 0

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Scheduled and stage to Restore
	seed.Status = Scheduled
	seed.Stage = Restore
	seed.updateSeedData()
	ProcessSeeds()

	// As Restore is always executed on target, set in status to Error
	targetTestAgent.agentSeedStageStatus.SeedID = seedID
	targetTestAgent.agentSeedStageStatus.Stage = Restore
	targetTestAgent.agentSeedStageStatus.Hostname = targetTestAgent.agent.Info.Hostname
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "processing restore stage"

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Restore, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	err = seed.AbortSeed()
	c.Assert(err, IsNil)

	activeSeeds, err := ReadActiveSeeds()
	c.Assert(err, IsNil)
	c.Assert(activeSeeds, HasLen, 0)
}

func (s *AgentTestSuite) TestProcessAborting(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	// MaxRetriesForStage set to 0, so it will be failed after first error
	config.Config.MaxRetriesForSeedStage = 0

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Aborting and stage to Restore
	seed.Status = Aborting
	seed.Stage = Restore
	seed.updateSeedData()
	ProcessSeeds()

	activeSeeds, err := ReadActiveSeeds()
	c.Assert(err, IsNil)
	c.Assert(activeSeeds, HasLen, 0)
}

func (s *AgentTestSuite) TestAbortSeedCleanup(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	config.Config.MaxRetriesForSeedStage = 2

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)

	// change seed status to Running and stage to Cleanup
	seed.Status = Running
	seed.Stage = Cleanup
	seed.updateSeedData()

	// One agent is OK, another had error
	for _, agent := range []*testAgent{targetTestAgent, sourceTestAgent} {
		agent.agentSeedStageStatus.SeedID = seedID
		agent.agentSeedStageStatus.Stage = Cleanup
		agent.agentSeedStageStatus.Hostname = agent.agent.Info.Hostname
		agent.agentSeedStageStatus.Timestamp = time.Now()
		agent.agentSeedStageStatus.Status = Running
		agent.agentSeedStageStatus.Details = "processing cleanup stage"
	}

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Cleanup, 0)
	s.readSeedStageStates(c, seed, 2, targetTestAgent, sourceTestAgent)

	err = seed.AbortSeed()
	c.Assert(err, IsNil)

	activeSeeds, err := ReadActiveSeeds()
	c.Assert(err, IsNil)
	c.Assert(activeSeeds, HasLen, 0)
}

func (s *AgentTestSuite) TestProcessSeedsStaleBackup(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	config.Config.MaxRetriesForSeedStage = 2
	config.Config.SeedBackupStaleFailMinutes = 1

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)
	ProcessSeeds()

	// change seed status to Scheduled and stage to Backup
	seed.Status = Scheduled
	seed.Stage = Backup
	seed.updateSeedData()
	ProcessSeeds()

	// As Mydumper is executed on target, set that backup stage is running on targetAgent
	targetTestAgent.agentSeedStageStatus.SeedID = seedID
	targetTestAgent.agentSeedStageStatus.Stage = Backup
	targetTestAgent.agentSeedStageStatus.Hostname = targetTestAgent.agent.Info.Hostname
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "processing backup stage"

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Backup, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// simulate that seed is in process and BackupDirDiskUsed is increasing on targetAgents
	targetTestAgent.agent.Data.BackupDirDiskUsed = 10
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Backup, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// simulate that seed is in process and BackupDirDiskUsed is not increasing on targetAgents. Increase timestamp to simulate stale
	targetTestAgent.agent.Data.BackupDirDiskUsed = 10
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now().Add(5 * time.Minute)
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Error, Backup, 0)
	targetTestAgent.agentSeedStageStatus.Status = Error
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. After error our seed will be moved back to Prepare stage Scheduled status and retries increased (due to backup stage is always moved to scheduled)
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 1)

	// so our SeedStage should also be in Scheduled status
	targetTestAgent.agentSeedStageStatus.Status = Scheduled
	targetTestAgent.agentSeedStageStatus.Stage = Prepare
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)
}

func (s *AgentTestSuite) TestProcessSeedsStaleRestore(c *C) {
	targetTestAgent := s.testAgents["agent1"]
	sourceTestAgent := s.testAgents["agent2"]

	config.Config.MaxRetriesForSeedStage = 2
	config.Config.SeedBackupStaleFailMinutes = 1

	// register agents to Orchestrator
	s.registerAgents(c)

	targetAgent, sourceAgent := s.getSeedAgents(c, targetTestAgent, sourceTestAgent)

	seedID, err := NewSeed("Mydumper", targetAgent, sourceAgent)
	c.Assert(err, IsNil)
	c.Assert(seedID, Equals, int64(1))

	// Orchestrator registered seed. It's will have first stage - Prepare and status Scheduled
	seed := s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Prepare, 0)
	ProcessSeeds()

	// change seed status to Scheduled and stage to Restore
	seed.Status = Scheduled
	seed.Stage = Restore
	seed.updateSeedData()
	ProcessSeeds()

	// Restore stage is running on targetAgent
	targetTestAgent.agentSeedStageStatus.SeedID = seedID
	targetTestAgent.agentSeedStageStatus.Stage = Restore
	targetTestAgent.agentSeedStageStatus.Hostname = targetTestAgent.agent.Info.Hostname
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now()
	targetTestAgent.agentSeedStageStatus.Status = Running
	targetTestAgent.agentSeedStageStatus.Details = "processing restore stage"

	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Restore, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// simulate that seed is in process and MySQLDatadirDiskUsed is increasing on targetAgents
	targetTestAgent.agent.Data.MySQLDatadirDiskUsed = 10
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Running, Restore, 0)
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// simulate that seed is in process and MySQLDatadirDiskUsed is not increasing on targetAgents. Increase timestamp to simulate stale
	targetTestAgent.agent.Data.MySQLDatadirDiskUsed = 10
	targetTestAgent.agentSeedStageStatus.Timestamp = time.Now().Add(5 * time.Minute)
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Error, Restore, 0)
	targetTestAgent.agentSeedStageStatus.Status = Error
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)

	// ProcessSeeds. After error our seed will be moved back to Restore stage Scheduled status and retries increased (due to backup stage is always moved to scheduled)
	ProcessSeeds()
	seed = s.readSeed(c, seedID, targetAgent.Info.Hostname, sourceAgent.Info.Hostname, Mydumper, Target, Scheduled, Restore, 1)

	// so our SeedStage should also be in Scheduled status
	targetTestAgent.agentSeedStageStatus.Status = Scheduled
	targetTestAgent.agentSeedStageStatus.Stage = Restore
	s.readSeedStageStates(c, seed, 1, targetTestAgent, sourceTestAgent)
}
