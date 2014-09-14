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

	"strings"	
	nethttp "net/http" 
	
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/logic"
	"github.com/outbrain/orchestrator/http"
	"github.com/outbrain/log"
)


// Http starts serving
func Http(discovery bool) {
	if config.Config.ServeAgentsHttp {
		go agentsHttp()
	}
	standardHttp(discovery)
}

// standardHttp starts serving standard HTTP (api/web) requests, to be used by normal clients 
func standardHttp(discovery bool) {
	m := martini.Classic()
	if strings.ToLower(config.Config.AuthenticationMethod) == "basic" && config.Config.HTTPAuthUser == "" {
		log.Warning("AuthenticationMethod is configured as 'basic' but not HTTPAuthUser defined. Running without autherntication.")
	}
	if strings.ToLower(config.Config.AuthenticationMethod) == "basic" && config.Config.HTTPAuthUser != "" {
		m.Use(auth.Basic(config.Config.HTTPAuthUser, config.Config.HTTPAuthPassword))
    }
	
	// Render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Directory: "resources",
		Layout: "templates/layout",	
		HTMLContentType: "text/html",
	}))
	m.Use(martini.Static("resources/public"))
	
	log.Info("Starting HTTP")
	
	if discovery {
		go orchestrator.ContinuousDiscovery()
	}

	http.API.RegisterRequests(m)
	http.Web.RegisterRequests(m)

	// Serve
	
	nethttp.ListenAndServe(config.Config.ListenAddress, m)
}



// agentsHttp startes serving agents API requests
func agentsHttp() {
	m := martini.Classic()
	m.Use(render.Renderer())
	
	log.Info("Starting agents HTTP")
	
	go orchestrator.ContinuousAgentsPoll()
	
	
	http.AgentsAPI.RegisterRequests(m)
	
	// Serve
	nethttp.ListenAndServe(":3001", m)
}


