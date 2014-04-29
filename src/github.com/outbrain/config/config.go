package config

import (
    "encoding/json"
    "os"
    
	"github.com/outbrain/log"
)

type Configuration struct {
	InstanceRefreshSeconds	uint
	MySQLTopologyUser		string
	MySQLTopologyPassword	string
	MySQLOrchestratorHost	string
	MySQLOrchestratorPort	uint
	MySQLOrchestratorDatabase	string
	MySQLOrchestratorUser		string
	MySQLOrchestratorPassword	string
	SlaveLagQuery				string
	DiscoverByShowSlaveHosts	bool
	InstanceUpToDateSeconds		uint
	DiscoveryPollSeconds		int
}	

var Config *Configuration = NewConfiguration()

func NewConfiguration() *Configuration {
	return &Configuration{
		InstanceRefreshSeconds:		60,
		InstanceUpToDateSeconds:	60,
		DiscoverByShowSlaveHosts:	false,
		DiscoveryPollSeconds:		5,
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
