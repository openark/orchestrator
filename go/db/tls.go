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

package db

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
	"github.com/patrickmn/go-cache"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/ssl"
)

const Error3159 = "Error 3159: Connections using insecure transport are prohibited while --require_secure_transport=ON."

// Track if a TLS has already been configured for topology
var topologyTLSConfigured bool = false

// Track if a TLS has already been configured for Orchestrator
var orchestratorTLSConfigured bool = false

var requireTLSCache *cache.Cache = cache.New(time.Duration(config.Config.TLSCacheTTL)*time.Second, time.Second)

func requiresTLS(host string, port int, mysql_uri string) bool {
	var required int = 0
	var first_time bool = true
	var found bool = false
	var value interface{}

	if value, found = requireTLSCache.Get(fmt.Sprintf("%s:%d", host, port)); found {
		required = value.(int)
	}

	if !found {
		query := `
			select
				required
			from
				database_instance_tls
			where
				hostname = ?
				and port = ?
				`
		err := QueryOrchestrator(query, sqlutils.Args(host, port), func(m sqlutils.RowMap) error {
			required = m.GetInt("required")
			first_time = false
			return nil
		})
		if err != nil {
			log.Errore(err)
			return required != 0
		}

		if first_time {
			db, _, _ := sqlutils.GetDB(mysql_uri)
			err = db.Ping()
			if err != nil && strings.Contains(err.Error(), Error3159) {
				required = 1
			} else {
				required = 0
			}

			insert := `
				insert into
					database_instance_tls
				set
					hostname = ?,
					port = ?,
					required = ?
					`
			_, err = ExecOrchestrator(insert, host, port, required)
			if err != nil {
				log.Errore(err)
				return required != 0
			}
		}

		requireTLSCache.Set(fmt.Sprintf("%s:%d", host, port), required, cache.DefaultExpiration)
	}

	return required != 0
}

// Create a TLS configuration from the config supplied CA, Certificate, and Private key.
// Register the TLS config with the mysql drivers as the "topology" config
// Modify the supplied URI to call the TLS config
func SetupMySQLTopologyTLS(uri string) (string, error) {
	if !topologyTLSConfigured {
		tlsConfig, err := ssl.NewTLSConfig(config.Config.MySQLTopologySSLCAFile, !config.Config.MySQLTopologySSLSkipVerify)
		// Drop to TLS 1.0 for talking to MySQL
		tlsConfig.MinVersion = tls.VersionTLS10
		if err != nil {
			return "", log.Fatalf("Can't create TLS configuration for Topology connection %s: %s", uri, err)
		}
		tlsConfig.InsecureSkipVerify = config.Config.MySQLTopologySSLSkipVerify

		if config.Config.MySQLTopologyUseMutualTLS ||
			config.Config.MySQLTopologySSLCertFile != "" ||
			config.Config.MySQLTopologySSLPrivateKeyFile != "" {
			if err = ssl.AppendKeyPair(tlsConfig, config.Config.MySQLTopologySSLCertFile, config.Config.MySQLTopologySSLPrivateKeyFile); err != nil {
				return "", log.Fatalf("Can't setup TLS key pairs for %s: %s", uri, err)
			}
		}
		if err = mysql.RegisterTLSConfig("topology", tlsConfig); err != nil {
			return "", log.Fatalf("Can't register mysql TLS config for topology: %s", err)
		}
		topologyTLSConfigured = true
	}
	return fmt.Sprintf("%s&tls=topology", uri), nil
}

// Create a TLS configuration from the config supplied CA, Certificate, and Private key.
// Register the TLS config with the mysql drivers as the "orchestrator" config
// Modify the supplied URI to call the TLS config
func SetupMySQLOrchestratorTLS(uri string) (string, error) {
	if !orchestratorTLSConfigured {
		tlsConfig, err := ssl.NewTLSConfig(config.Config.MySQLOrchestratorSSLCAFile, true)
		// Drop to TLS 1.0 for talking to MySQL
		tlsConfig.MinVersion = tls.VersionTLS10
		if err != nil {
			return "", log.Fatalf("Can't create TLS configuration for Orchestrator connection %s: %s", uri, err)
		}
		tlsConfig.InsecureSkipVerify = config.Config.MySQLOrchestratorSSLSkipVerify
		if err = ssl.AppendKeyPair(tlsConfig, config.Config.MySQLOrchestratorSSLCertFile, config.Config.MySQLOrchestratorSSLPrivateKeyFile); err != nil {
			return "", log.Fatalf("Can't setup TLS key pairs for %s: %s", uri, err)
		}
		if err = mysql.RegisterTLSConfig("orchestrator", tlsConfig); err != nil {
			return "", log.Fatalf("Can't register mysql TLS config for orchestrator: %s", err)
		}
		orchestratorTLSConfigured = true
	}
	return fmt.Sprintf("%s&tls=orchestrator", uri), nil
}
