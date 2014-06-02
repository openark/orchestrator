// 
package app

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/auth"
	
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/logic"
	"github.com/outbrain/orchestrator/http"
	"github.com/outbrain/log"
)


// Http starts serving HTTP (api/web) requests 
func Http(discovery bool) {
	m := martini.Classic()
	if config.Config.HTTPAuthUser != "" {
		m.Use(auth.Basic(config.Config.HTTPAuthUser, config.Config.HTTPAuthPassword))
    }
	
	// Render html templates from templates directory
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

	// Serve
	m.Run()
}

