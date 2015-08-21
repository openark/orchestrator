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
	nethttp "net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/gzip"
	"github.com/martini-contrib/render"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/http"
	"github.com/outbrain/orchestrator/go/inst"
	"github.com/outbrain/orchestrator/go/logic"
)

// acceptSignals registers for OS signals
func acceptSignals() {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGHUP)
	go func() {
		for sig := range c {
			switch sig {
			case syscall.SIGHUP:
				log.Debugf("Received SIGHUP. Reloading configuration")
				config.Reload()
			}
		}
	}()
}

// Http starts serving
func Http(discovery bool) {
	acceptSignals()

	martini.Env = martini.Prod
	if config.Config.ServeAgentsHttp {
		go agentsHttp()
	}
	standardHttp(discovery)
}

// standardHttp starts serving HTTP or HTTPS (api/web) requests, to be used by normal clients
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
	if config.Config.UseMutualTLS {
		m.Use(http.VerifyOUs(config.Config.SSLValidOUs))
	}

	inst.SetMaintenanceOwner(logic.ThisHostname)

	if discovery {
		log.Info("Starting Discovery")
		go logic.ContinuousDiscovery()
	}
	inst.ReadClusterAliases()

	log.Info("Registering endpoints")
	http.API.RegisterRequests(m)
	http.Web.RegisterRequests(m)

	// Serve
	if config.Config.UseSSL {
		log.Info("Starting HTTPS listener")
		tlsConfig, err := http.NewTLSConfig(config.Config.SSLCAFile, config.Config.UseMutualTLS)
		if err != nil {
			log.Fatale(err)
		}
		if err = http.AppendKeyPair(tlsConfig, config.Config.SSLCertFile, config.Config.SSLPrivateKeyFile); err != nil {
			log.Fatale(err)
		}
		if err = http.ListenAndServeTLS(config.Config.ListenAddress, m, tlsConfig); err != nil {
			log.Fatale(err)
		}
	} else {
		log.Info("Starting HTTP listener")
		if err := nethttp.ListenAndServe(config.Config.ListenAddress, m); err != nil {
			log.Fatale(err)
		}
	}
	log.Info("Web server started")
}

// agentsHttp startes serving agents HTTP or HTTPS API requests
func agentsHttp() {
	m := martini.Classic()
	m.Use(gzip.All())
	m.Use(render.Renderer())
	if config.Config.AgentsUseMutualTLS {
		m.Use(http.VerifyOUs(config.Config.AgentSSLValidOUs))
	}

	log.Info("Starting agents listener")

	go logic.ContinuousAgentsPoll()

	http.AgentsAPI.RegisterRequests(m)

	// Serve
	if config.Config.AgentsUseSSL {
		log.Info("Starting agent HTTPS listener")
		tlsConfig, err := http.NewTLSConfig(config.Config.AgentSSLCAFile, config.Config.AgentsUseMutualTLS)
		if err != nil {
			log.Fatale(err)
		}
		if err = http.AppendKeyPair(tlsConfig, config.Config.AgentSSLCertFile, config.Config.AgentSSLPrivateKeyFile); err != nil {
			log.Fatale(err)
		}
		if err = http.ListenAndServeTLS(config.Config.AgentsServerPort, m, tlsConfig); err != nil {
			log.Fatale(err)
		}
	} else {
		log.Info("Starting agent HTTP listener")
		if err := nethttp.ListenAndServe(config.Config.AgentsServerPort, m); err != nil {
			log.Fatale(err)
		}
	}
	log.Info("Agent server started")
}
