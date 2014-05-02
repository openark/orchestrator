// 
package orchestrator

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/outbrain/http"
	"github.com/outbrain/orch"
	"github.com/outbrain/log"
)

func Http(discovery bool) {
	m := martini.Classic()
	// render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Directory: "resources",
		Layout: "templates/layout",	
		HTMLContentType: "text/html",
	}))
	m.Use(martini.Static("resources/public"))
	
	log.Info("Started HTTP")
	
	if discovery {
		go orchestrator.ContinuousDiscovery()
	}


	http.API.RegisterRequests(m)
	http.Web.RegisterRequests(m)

	m.Run()
}

