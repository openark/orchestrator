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
	"github.com/rcrowley/go-metrics"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/ssl"
)

// Error3159 is the error given by MySQL if TLS is required for connections
const Error3159 = "Error 3159:"

// ErrorTLSNotSupported is the error given by MySQL if the client requests TLS and the server does not support it.
const ErrorTLSNotSupported = "TLS requested but server does not support TLS"

// Track if a TLS has already been configured for topology
var topologyTLSConfigured bool

// Track if a TLS has already been configured for Orchestrator
var orchestratorTLSConfigured bool

var requireTLSCache = cache.New(time.Duration(config.Config.TLSCacheTTLFactor*config.Config.InstancePollSeconds)*time.Second, time.Second)

var readInstanceTLSCounter = metrics.NewCounter()
var writeInstanceTLSCounter = metrics.NewCounter()
var readInstanceTLSCacheCounter = metrics.NewCounter()
var writeInstanceTLSCacheCounter = metrics.NewCounter()

func init() {
	metrics.Register("instance_tls.read", readInstanceTLSCounter)
	metrics.Register("instance_tls.write", writeInstanceTLSCounter)
	metrics.Register("instance_tls.read_cache", readInstanceTLSCacheCounter)
	metrics.Register("instance_tls.write_cache", writeInstanceTLSCacheCounter)
}

// Does the host require TLS? return true if it does.
// - determine this by storing in the database the appropriate value
// - use a cache to avoid a lot of backend queries
// - catch TLS requires and TLS not supported errors when connecting
//   and adjust the settings as appropriate
// the mysqlURI is the one _without_ TLS configuration settings
func requiresTLS(host string, port int, mysqlURI string) bool {
	const (
		selectQuery = `
			select
				required
			from
				database_instance_tls
			where
				hostname = ?
				and port = ?
		`
		insertOrUpdateQuery = `
			insert into
			database_instance_tls (
				hostname, port, required
			) values (
				?, ?, ?
			)
			on duplicate key update
			required=values(required)
		`
	)

	var (
		requiredInBackend      int
		requiredWhenConnecting int
		foundInBackend         bool
		foundInCache           bool
		value                  interface{}
	)
	cacheKey := fmt.Sprintf("%s:%d", host, port)

	// Check the value from the cache and return that if available
	if value, foundInCache = requireTLSCache.Get(cacheKey); foundInCache {
		readInstanceTLSCacheCounter.Inc(1)

		return value.(int) != 0
	}

	// Check from the backend the expected value to use
	// - note: this value may be wrong as the server configuration may have changed.
	err := QueryOrchestrator(selectQuery, sqlutils.Args(host, port), func(m sqlutils.RowMap) error {
		foundInBackend = true
		requiredInBackend = m.GetInt("required")
		return nil
	})
	readInstanceTLSCounter.Inc(1)
	if err != nil {
		log.Errore(err)
		return requiredInBackend != 0 // this value may be wrong!
	}

	// Try to connect to the instance using the normal uri
	db, _, _ := sqlutils.GetDB(mysqlURI)
	err = db.Ping()
	if err != nil {
		// TLS required message
		if strings.Contains(err.Error(), Error3159) {
			requiredWhenConnecting = 1
		} else if strings.Contains(err.Error(), ErrorTLSNotSupported) {
			requiredWhenConnecting = 0
		}
	}

	// update the cache
	requireTLSCache.Set(fmt.Sprintf("%s:%d", host, port), requiredWhenConnecting, cache.DefaultExpiration)
	writeInstanceTLSCacheCounter.Inc(1)

	// figure out if we need to update the backend database
	if !foundInBackend || requiredInBackend != requiredWhenConnecting {
		_, err = ExecOrchestrator(insertOrUpdateQuery, host, port, requiredWhenConnecting)
		if err != nil {
			log.Errore(err)
			return requiredWhenConnecting != 0
		}

		writeInstanceTLSCounter.Inc(1)
	}

	return requiredWhenConnecting != 0
}

// ForgetInstanceTLSCache removes the special TLS configuration cache entry
func ForgetInstanceTLSCache(host string, port int) {
	requireTLSCache.Delete(fmt.Sprintf("%s:%d", host, port))
}

// SetupMySQLTopologyTLS creates a TLS configuration from the config supplied CA, Certificate, and Private key.
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
				return "", log.Errorf("Can't setup TLS key pairs for %s: %s", uri, err)
			}
		}
		if err = mysql.RegisterTLSConfig("topology", tlsConfig); err != nil {
			return "", log.Errorf("Can't register mysql TLS config for topology: %s", err)
		}
		topologyTLSConfigured = true
	}
	return fmt.Sprintf("%s&tls=topology", uri), nil
}

// SetupMySQLOrchestratorTLS creates a TLS configuration from the config supplied CA, Certificate, and Private key.
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
