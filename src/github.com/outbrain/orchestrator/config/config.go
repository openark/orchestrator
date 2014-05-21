package config

import (
    "encoding/json"
    "os"
    
	"github.com/outbrain/log"
)

type Configuration struct {
	MySQLTopologyUser		string
	MySQLTopologyPassword	string
	MySQLOrchestratorHost	string
	MySQLOrchestratorPort	uint
	MySQLOrchestratorDatabase	string
	MySQLOrchestratorUser		string
	MySQLOrchestratorPassword	string
	SlaveLagQuery				string		// custom query to check on slave lg (e.g. heartbeat table)
	SlaveStartPostWaitMilliseconds	int		// Time to wait after START SLAVE before re-readong instance (give slave chance to connect to master)
	DiscoverByShowSlaveHosts	bool		// Attempt SHOW SLAVE HOSTS before PROCESSLIST
	InstancePollSeconds			uint		// Number of seconds between instance reads
	UnseenInstanceForgetHours	uint		// Number of hours after which an unseen instance is forgotten
	DiscoveryPollSeconds		int			// Auto/continuous discovery of instances sleep time between polls
	ReasonableReplicationLagSeconds	int		// Abvoe this value is considered a problem
	HTTPAuthUser		string				// Username for HTTP Basic authentication (blank disables authentication)
	HTTPAuthPassword	string				// Password for HTTP Basic authentication
}	

var Config *Configuration = NewConfiguration()

func NewConfiguration() *Configuration {
	return &Configuration {
		InstancePollSeconds:		60,
		UnseenInstanceForgetHours:	240,
		SlaveStartPostWaitMilliseconds: 1000,
		DiscoverByShowSlaveHosts:	false,
		DiscoveryPollSeconds:		5,
		ReasonableReplicationLagSeconds: 10,
		HTTPAuthUser		: "",
		HTTPAuthPassword	: "",
	}
}

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
	}
	return Config, err
}

func Read(file_names ...string) *Configuration {
	for _, file_name := range file_names {
		read(file_name)
	}
	return Config
}

func ForceRead(file_name string) *Configuration {
	_, err := read(file_name)
	if err != nil {
		log.Fatal("Cannot read config file:", file_name, err)
	}
	return Config
}
