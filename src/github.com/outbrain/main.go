package main 

import (
	"fmt"
	"github.com/outbrain/log"
	"github.com/outbrain/inst"
	"github.com/outbrain/config"
)

func main() {
	config.Read("/etc/orchestrator.conf.json", "conf/orchestrator.conf.json", "orchestrator.conf.json")
	i := &inst.Instance{}
	log.Debug("I'm running");
	log.Debug("I'm running...", *i);
	log.Debug("config: ", config.Config);
	
	fmt.Println(fmt.Sprintf("the filesi: %s", "/tmp/f.txt"))
}
