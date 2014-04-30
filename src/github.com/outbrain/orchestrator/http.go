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
	m.Use(render.Renderer())
	
	log.Info("Started HTTP")
	
	if discovery {
		go orchestrator.ContinuousDiscovery()
	}

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", map[string]interface{}{"name": "shushu", "PI":3.14, "complex": map[string]string{"title":"rollback"}})
	})

	http.API.RegisterRequests(m)

	m.Run()
}

