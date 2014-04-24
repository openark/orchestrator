package main 

import (
	"fmt"
	"log"
	"github.com/outbrain/inst"
	"github.com/outbrain/config"
)

func main() {
	config.Read("/etc/orchestrator.conf.json", "conf/orchestrator.conf.json", "orchestrator.conf.json")
	i := &instance.Instance{}
	log.Println("I'm running");
	fmt.Println("I'm running...", i);
	fmt.Println("config: ", config.Config);
}
