// 
package main

import (
	"flag"
	
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/outbrain/config"
	"github.com/outbrain/http"
	"github.com/outbrain/orch"
	"github.com/outbrain/log"
)

func main() {
	configFile := flag.String("config", "", "config file name")
	verbose := flag.Bool("verbose", false, "verbose")
	debug := flag.Bool("debug", false, "debug mode (very verbose)")
	discovery := flag.Bool("discovery", false, "auto discovery mode")
	flag.Parse();
	
	log.SetLevel(log.ERROR)
	if *verbose {
		log.SetLevel(log.INFO)
	}
	if *debug {
		log.SetLevel(log.DEBUG)
	}
	
	if len(flag.Args()) > 0 {
		log.Critical("Unexpected arguments:", flag.Args()) 
	}

	if len(*configFile) > 0 {
		config.ForceRead(*configFile)
	} else {
		config.Read("/etc/orchestrator.conf.json", "conf/orchestrator.conf.json", "orchestrator.conf.json")
	}


	m := martini.Classic()
	// render html templates from templates directory
	m.Use(render.Renderer())
	
	log.Info("Started")
	
	if *discovery {
		go orchestrator.ContinuousDiscovery()
	}

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", map[string]interface{}{"name": "shushu", "PI":3.14, "complex": map[string]string{"title":"rollback"}})
	})

	http.API.RegisterRequests(m)

	m.Run()
}

