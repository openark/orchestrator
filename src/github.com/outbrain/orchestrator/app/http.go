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
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"

	nethttp "net/http"
	"strings"

	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/config"
	"github.com/outbrain/orchestrator/http"
	"github.com/outbrain/orchestrator/inst"
	"github.com/outbrain/orchestrator/logic"
)

// Http starts serving
func Http(discovery bool) {
	martini.Env = martini.Prod
	if config.Config.ServeAgentsHttp {
		go agentsHttp()
	}
	standardHttp(discovery)
}

// standardHttp starts serving standard HTTP (api/web) requests, to be used by normal clients
func standardHttp(discovery bool) {
	m := martini.Classic()

	switch strings.ToLower(config.Config.AuthenticationMethod) {
	case "basic":
		{
			if config.Config.HTTPAuthUser == "" {
				// Still allowed; may be disallowed in future versions
				log.Warning("AuthenticationMethod is configured as 'basic' but HTTPAuthUser undefined. Running without authentication.")
			}
			m.Use(auth.Basic(config.Config.HTTPAuthUser, config.Config.HTTPAuthPassword))
		}
	case "multi":
		{
			if config.Config.HTTPAuthUser == "" {
				// Still allowed; may be disallowed in future versions
				log.Fatal("AuthenticationMethod is configured as 'multi' but HTTPAuthUser undefined")
			}

			m.Use(auth.BasicFunc(func(username, password string) bool {
				if username == "readonly" {
					// Will be treated as "read-only"
					return true
				}
				return auth.SecureCompare(username, config.Config.HTTPAuthUser) && auth.SecureCompare(password, config.Config.HTTPAuthPassword)
			}))
		}
	default:
		{
			// We inject a dummy User object because we have function signatures with User argument in api.go
			m.Map(auth.User(""))
		}
	}

	m.Use(gzip.All())
	// Render html templates from templates directory
	m.Use(render.Renderer(render.Options{
		Directory:       "resources",
		Layout:          "templates/layout",
		HTMLContentType: "text/html",
	}))
	m.Use(martini.Static("resources/public"))

	inst.SetMaintenanceOwner(orchestrator.ThisHostname)

	log.Info("Starting HTTP")

	if discovery {
		go orchestrator.ContinuousDiscovery()
	}
	inst.ReadClusterAliases()

	http.API.RegisterRequests(m)
	http.Web.RegisterRequests(m)

	// Serve
	if err := nethttp.ListenAndServe(config.Config.ListenAddress, m); err != nil {
		log.Fatale(err)
	}
}

// agentsHttp startes serving agents API requests
func agentsHttp() {
	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(render.Renderer())

	log.Info("Starting agents HTTP")

	go orchestrator.ContinuousAgentsPoll()

	http.AgentsAPI.RegisterRequests(m)

	// Serve
	if config.Config.AgentsUseSSL {
		log.Info("Serving via SSL")
		err := nethttp.ListenAndServeTLS(":3001", config.Config.SSLCertFile, config.Config.SSLPrivateKeyFile, m)
		if err != nil {
			log.Fatale(err)
		}
	} else {
		nethttp.ListenAndServe(":3001", m)
	}
}
