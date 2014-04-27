package config

import (
    "encoding/json"
    "os"
    "log"
    "fmt"
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
}	

var Config *Configuration = &Configuration{}

func read(file_name string) (*Configuration, error) {
	file, err := os.Open(file_name)
	if err == nil {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(Config)
		if err == nil {
			log.Println("Read config: ", file_name)
		} else {
	  		log.Fatal(fmt.Sprintf("Cannot read config file %s: ", file_name), err)
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
		log.Fatal(fmt.Sprintf("Cannot read config file %s: ", file_name), err)
	}
	return Config
}
