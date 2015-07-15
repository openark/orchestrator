/*
   Copyright 2014 Outbrain Inc.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package config

import (
	"code.google.com/p/gcfg"
	"encoding/json"
	"os"

	"github.com/outbrain/golib/log"
)

// Configuration makes for orchestrator configuration input, which can be provided by user via JSON formatted file.
// Some of the parameteres have reasonable default values, and some (like database credentials) are
// strictly expected from user.
type Configuration struct {
	Debug                                      bool // set debug mode (similar to --debug option)
	ListenAddress                              string
	MySQLTopologyUser                          string
	MySQLTopologyPassword                      string // my.cnf style configuration file from where to pick credentials. Expecting `user`, `password` under `[client]` section
	MySQLTopologyCredentialsConfigFile         string
	MySQLTopologyMaxPoolConnections            int // Max concurrent connections on any topology instance
	MySQLOrchestratorHost                      string
	MySQLOrchestratorPort                      uint
	MySQLOrchestratorDatabase                  string
	MySQLOrchestratorUser                      string
	MySQLOrchestratorPassword                  string
	MySQLOrchestratorCredentialsConfigFile     string // my.cnf style configuration file from where to pick credentials. Expecting `user`, `password` under `[client]` section
	MySQLConnectTimeoutSeconds                 int    // Number of seconds before connection is aborted (driver-side)
	DefaultInstancePort                        int    // In case port was not specified on command line
	SkipOrchestratorDatabaseUpdate             bool   // When false, orchestrator will attempt to create & update all tables in backend database; when true, this is skipped. It makes sense to skip on command-line invocations and to enable for http or occasional invocations, or just after upgrades
	SlaveLagQuery                              string // custom query to check on slave lg (e.g. heartbeat table)
	SlaveStartPostWaitMilliseconds             int    // Time to wait after START SLAVE before re-readong instance (give slave chance to connect to master)
	DiscoverByShowSlaveHosts                   bool   // Attempt SHOW SLAVE HOSTS before PROCESSLIST
	InstancePollSeconds                        uint   // Number of seconds between instance reads
	ReadLongRunningQueries                     bool   // Whether orchestrator should read and record current long running executing queries.
	UnseenInstanceForgetHours                  uint   // Number of hours after which an unseen instance is forgotten
	SnapshotTopologiesIntervalHours            uint   // Interval in hour between snapshot-topologies invocation. Default: 0 (disabled)
	DiscoveryPollSeconds                       uint   // Auto/continuous discovery of instances sleep time between polls
	InstanceBulkOperationsWaitTimeoutSeconds   uint   // Time to wait on a single instance when doing bulk (many instances) operation
	ActiveNodeExpireSeconds                    uint   // Maximum time to wait for active node to send keepalive before attempting to take over as active node.
	HostnameResolveMethod                      string // Method by which to "normalize" hostname ("none"/"default"/"cname")
	MySQLHostnameResolveMethod                 string // Method by which to "normalize" hostname via MySQL server. ("none"/"@@hostname"/"@@report_host"; default "@@hostname")
	ExpiryHostnameResolvesMinutes              int    // Number of minutes after which to expire hostname-resolves
	RejectHostnameResolvePattern               string // Regexp pattern for resolved hostname that will not be accepted (not cached, not written to db). This is done to avoid storing wrong resolves due to network glitches.
	ReasonableReplicationLagSeconds            int    // Above this value is considered a problem
	VerifyReplicationFilters                   bool   // Include replication filters check before approving topology refactoring
	MaintenanceOwner                           string // (Default) name of maintenance owner to use if none provided
	ReasonableMaintenanceReplicationLagSeconds int    // Above this value move-up and move-below are blocked
	MaintenanceExpireMinutes                   uint   // Minutes after which a maintenance flag is considered stale and is cleared
	MaintenancePurgeDays                       uint   // Days after which maintenance entries are purged from the database
	CandidateInstanceExpireMinutes             uint   // Minutes after which a suggestion to use an instance as a candidate slave (to be preferably promoted on master failover) is expired.
	AuditLogFile                               string // Name of log file for audit operations. Disabled when empty.
	AuditPageSize                              int
	RemoveTextFromHostnameDisplay              string // Text to strip off the hostname on cluster/clusters pages
	ReadOnly                                   bool
	AuthenticationMethod                       string            // Type of autherntication to use, if any. "" for none, "basic" for BasicAuth, "multi" for advanced BasicAuth, "proxy" for forwarded credentials via reverse proxy
	HTTPAuthUser                               string            // Username for HTTP Basic authentication (blank disables authentication)
	HTTPAuthPassword                           string            // Password for HTTP Basic authentication
	AuthUserHeader                             string            // HTTP header indicating auth user, when AuthenticationMethod is "proxy"
	PowerAuthUsers                             []string          // On AuthenticationMethod == "proxy", list of users that can make changes. All others are read-only.
	ClusterNameToAlias                         map[string]string // map between regex matching cluster name to a human friendly alias
	DetectClusterAliasQuery                    string            // Optional query (executed on topology instance) that returns the alias of a cluster. Query will only be executed on cluster master (though until the topology's master is resovled it may execute on other/all slaves). If provided, must return one row, one column
	DataCenterPattern                          string            // Regexp pattern with one group, extracting the datacenter name from the hostname
	PhysicalEnvironmentPattern                 string            // Regexp pattern with one group, extracting physical environment info from hostname (e.g. combination of datacenter & prod/dev env)
	PromotionIgnoreHostnameFilters             []string          // Orchestrator will not promote slaves with hostname matching pattern (via -c recovery; for example, avoid promoting dev-dedicated machines)
	ServeAgentsHttp                            bool              // Spawn another HTTP interface dedicated for orcehstrator-agent
	AgentsUseSSL                               bool              // When "true" orchestrator will listen on agents port with SSL as well as connect to agents via SSL
	SSLSkipVerify                              bool              // When using SSL, should we ignore SSL certification error
	SSLPrivateKeyFile                          string            // Name of SSL private key file, applies only when AgentsUseSSL = true
	SSLCertFile                                string            // Name of SSL certification file, applies only when AgentsUseSSL = true
	HttpTimeoutSeconds                         int               // Number of idle seconds before HTTP GET request times out (when accessing orchestrator-agent)
	AgentPollMinutes                           uint              // Minutes between agent polling
	UnseenAgentForgetHours                     uint              // Number of hours after which an unseen agent is forgotten
	StaleSeedFailMinutes                       uint              // Number of minutes after which a stale (no progress) seed is considered failed.
	SeedAcceptableBytesDiff                    int64             // Difference in bytes between seed source & target data size that is still considered as successful copy
	PseudoGTIDPattern                          string            // Pattern to look for in binary logs that makes for a unique entry (pseudo GTID). When empty, Pseudo-GTID based refactoring is disabled.
	PseudoGTIDMonotonicHint                    string            // subtring in Pseudo-GTID entry which indicates Pseudo-GTID entries are expected to be monotonically increasing
	DetectPseudoGTIDQuery                      string            // Optional query which is used to authoritatively decide whether pseudo gtid is enabled on instance
	BinlogEventsChunkSize                      int               // Chunk size (X) for SHOW BINLOG|RELAYLOG EVENTS LIMIT ?,X statements. Smaller means less locking and mroe work to be done
	BufferBinlogEvents                         bool              // Should we used buffered read on SHOW BINLOG|RELAYLOG EVENTS -- releases the database lock sooner (recommended)
	FailureDetectionPeriodBlockMinutes         int               // The time for which an instance's failure discovery is kept "active", so as to avoid concurrent "discoveries" of the instance's failure; this preceeds any recovery process, if any.
	RecoveryPeriodBlockMinutes                 int               // The time for which an instance's recovery is kept "active", so as to avoid concurrent recoveries on smae instance as well as flapping
	RecoveryIgnoreHostnameFilters              []string          // Recovery analysis will completely ignore hosts matching given patterns
	RecoverMasterClusterFilters                []string          // Only do master recovery on clusters matching these regexp patterns (of course the ".*" pattern matches everything)
	RecoverIntermediateMasterClusterFilters    []string          // Only do IM recovery on clusters matching these regexp patterns (of course the ".*" pattern matches everything)
	OnFailureDetectionProcesses                []string          // Processes to execute when detecting a failover scenario (before making a decision whether to failover or not). May and should use some of these placeholders: {failureType}, {failureDescription}, {failedHost}, {failureCluster}, {failureClusterAlias}, {failedPort}, {successorHost}, {successorPort}, {countSlaves}, {slaveHosts}
	PreFailoverProcesses                       []string          // Processes to execute before doing a failover (aborting operation should any once of them exits with non-zero code; order of execution undefined). May and should use some of these placeholders: {failureType}, {failureDescription}, {failedHost}, {failureCluster}, {failureClusterAlias}, {failedPort}, {successorHost}, {successorPort}, {countSlaves}, {slaveHosts}
	PostFailoverProcesses                      []string          // Processes to execute after doing a failover (order of execution undefined). May and should use some of these placeholders: {failureType}, {failureDescription}, {failedHost}, {failureCluster}, {failureClusterAlias}, {failedPort}, {successorHost}, {successorPort}, {countSlaves}, {slaveHosts}
	PostMasterFailoverProcesses                []string          // Processes to execute after doing a master failover (order of execution undefined). Uses same placeholders as PostFailoverProcesses
	PostIntermediateMasterFailoverProcesses    []string          // Processes to execute after doing a master failover (order of execution undefined). Uses same placeholders as PostFailoverProcesses
	OSCIgnoreHostnameFilters                   []string          // OSC slaves recommendation will ignore slave hostnames matching given patterns
}

var Config *Configuration = NewConfiguration()
var readFileNames []string

func NewConfiguration() *Configuration {
	return &Configuration{
		Debug:                                      false,
		ListenAddress:                              ":3000",
		MySQLOrchestratorPort:                      3306,
		MySQLTopologyMaxPoolConnections:            3,
		MySQLConnectTimeoutSeconds:                 5,
		DefaultInstancePort:                        3306,
		SkipOrchestratorDatabaseUpdate:             false,
		InstancePollSeconds:                        60,
		ReadLongRunningQueries:                     true,
		UnseenInstanceForgetHours:                  240,
		SnapshotTopologiesIntervalHours:            0,
		SlaveStartPostWaitMilliseconds:             1000,
		DiscoverByShowSlaveHosts:                   false,
		DiscoveryPollSeconds:                       5,
		InstanceBulkOperationsWaitTimeoutSeconds:   10,
		ActiveNodeExpireSeconds:                    60,
		HostnameResolveMethod:                      "cname",
		MySQLHostnameResolveMethod:                 "@@hostname",
		ExpiryHostnameResolvesMinutes:              60,
		RejectHostnameResolvePattern:               "",
		ReasonableReplicationLagSeconds:            10,
		VerifyReplicationFilters:                   false,
		MaintenanceOwner:                           "orchestrator",
		ReasonableMaintenanceReplicationLagSeconds: 20,
		MaintenanceExpireMinutes:                   10,
		MaintenancePurgeDays:                       365,
		CandidateInstanceExpireMinutes:             60,
		AuditLogFile:                               "",
		AuditPageSize:                              20,
		RemoveTextFromHostnameDisplay:              "",
		ReadOnly:                                   false,
		AuthenticationMethod:                       "basic",
		HTTPAuthUser:                               "",
		HTTPAuthPassword:                           "",
		AuthUserHeader:                             "X-Forwarded-User",
		PowerAuthUsers:                             []string{"*"},
		ClusterNameToAlias:                         make(map[string]string),
		DetectClusterAliasQuery:                    "",
		DataCenterPattern:                          "",
		PhysicalEnvironmentPattern:                 "",
		PromotionIgnoreHostnameFilters:             []string{},
		ServeAgentsHttp:                            false,
		AgentsUseSSL:                               false,
		SSLSkipVerify:                              false,
		SSLPrivateKeyFile:                          "",
		SSLCertFile:                                "",
		HttpTimeoutSeconds:                         60,
		AgentPollMinutes:                           60,
		UnseenAgentForgetHours:                     6,
		StaleSeedFailMinutes:                       60,
		SeedAcceptableBytesDiff:                    8192,
		PseudoGTIDPattern:                          "",
		PseudoGTIDMonotonicHint:                    "",
		DetectPseudoGTIDQuery:                      "",
		BinlogEventsChunkSize:                      10000,
		BufferBinlogEvents:                         true,
		FailureDetectionPeriodBlockMinutes:         60,
		RecoveryPeriodBlockMinutes:                 60,
		RecoveryIgnoreHostnameFilters:              []string{},
		RecoverMasterClusterFilters:                []string{},
		RecoverIntermediateMasterClusterFilters:    []string{},
		OnFailureDetectionProcesses:                []string{},
		PreFailoverProcesses:                       []string{},
		PostMasterFailoverProcesses:                []string{},
		PostIntermediateMasterFailoverProcesses:    []string{},
		PostFailoverProcesses:                      []string{},
		OSCIgnoreHostnameFilters:                   []string{},
	}
}

// read reads configuration from given file, or silently skips if the file does not exist.
// If the file does exist, then it is expected to be in valid JSON format or the function bails out.
func read(file_name string) (*Configuration, error) {
	file, err := os.Open(file_name)
	if err == nil {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(Config)
		if err == nil {
			log.Infof("Read config: %s", file_name)
		} else {
			log.Fatal("Cannot read config file:", file_name, err)
		}
		if Config.MySQLOrchestratorCredentialsConfigFile != "" {
			mySQLConfig := struct {
				Client struct {
					User     string
					Password string
				}
			}{}
			err := gcfg.ReadFileInto(&mySQLConfig, Config.MySQLOrchestratorCredentialsConfigFile)
			if err != nil {
				log.Fatalf("Failed to parse gcfg data from file: %+v", err)
			} else {
				log.Debugf("Parsed orchestrator credentials from %s", Config.MySQLOrchestratorCredentialsConfigFile)
				Config.MySQLOrchestratorUser = mySQLConfig.Client.User
				Config.MySQLOrchestratorPassword = mySQLConfig.Client.Password
			}
		}
		if Config.MySQLTopologyCredentialsConfigFile != "" {
			mySQLConfig := struct {
				Client struct {
					User     string
					Password string
				}
			}{}
			err := gcfg.ReadFileInto(&mySQLConfig, Config.MySQLTopologyCredentialsConfigFile)
			if err != nil {
				log.Fatalf("Failed to parse gcfg data from file: %+v", err)
			} else {
				log.Debugf("Parsed topology credentials from %s", Config.MySQLOrchestratorCredentialsConfigFile)
				Config.MySQLTopologyUser = mySQLConfig.Client.User
				Config.MySQLTopologyPassword = mySQLConfig.Client.Password
			}
		}

	}
	return Config, err
}

// Read reads configuration from zero, either, some or all given files, in order of input.
// A file can override configuration provided in previous file.
func Read(file_names ...string) *Configuration {
	for _, file_name := range file_names {
		read(file_name)
	}
	readFileNames = file_names
	return Config
}

// ForceRead reads configuration from given file name or bails out if it fails
func ForceRead(file_name string) *Configuration {
	_, err := read(file_name)
	if err != nil {
		log.Fatal("Cannot read config file:", file_name, err)
	}
	readFileNames = []string{file_name}
	return Config
}

func Reload() *Configuration {
	for _, file_name := range readFileNames {
		read(file_name)
	}
	return Config
}
