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

