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

const Error3159 = "Error 3159:"
const Error1045 = "Access denied for user"

// This cache exists to prevent re-registering tls.Conf instances while also
// allowing them to be interrogated in tests.
var tlsConfigCache *cache.Cache = cache.New(cache.NoExpiration, time.Second)
var requireTLSCache *cache.Cache = cache.New(time.Duration(config.Config.TLSCacheTTLFactor*config.Config.InstancePollSeconds)*time.Second, time.Second)

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

func requiresTLS(host string, port int, mysql_uri string) bool {
	cacheKey := fmt.Sprintf("%s:%d", host, port)

	if value, found := requireTLSCache.Get(cacheKey); found {
		readInstanceTLSCacheCounter.Inc(1)
		return value.(bool)
	}

	required := false
	db, _, _ := sqlutils.GetDB(mysql_uri)
	if err := db.Ping(); err != nil && (strings.Contains(err.Error(), Error3159) || strings.Contains(err.Error(), Error1045)) {
		required = true
	}

	query := `
			insert into
				database_instance_tls (
					hostname, port, required
				) values (
					?, ?, ?
				)
				on duplicate key update
					required=values(required)
				`
	if _, err := ExecOrchestrator(query, host, port, required); err != nil {
		log.Errore(err)
	}
	writeInstanceTLSCounter.Inc(1)

	requireTLSCache.Set(cacheKey, required, cache.DefaultExpiration)
	writeInstanceTLSCacheCounter.Inc(1)

	return required
}

// Create a TLS configuration from the config supplied CA, Certificate, and Private key.
// Register the TLS config with the mysql drivers as the "topology" config
// Modify the supplied URI to call the TLS config
func SetupMySQLTopologyTLS(uri string) (string, error) {
	c := connectionConfig{
		uri:              uri,
		caCertFile:       config.Config.MySQLTopologySSLCAFile,
		useClientAuth:    config.Config.MySQLTopologyUseMutualTLS,
		clientCertFile:   config.Config.MySQLTopologySSLCertFile,
		clientKeyFile:    config.Config.MySQLTopologySSLPrivateKeyFile,
		verifyClientCert: !config.Config.MySQLTopologySSLSkipVerify,
		verifyServerCert: !config.Config.MySQLTopologySSLSkipVerify,
	}

	return setupTls("topology", c)
}

// Create a TLS configuration from the config supplied CA, Certificate, and Private key.
// Register the TLS config with the mysql drivers as the "orchestrator" config
// Modify the supplied URI to call the TLS config
func SetupMySQLOrchestratorTLS(uri string) (string, error) {
	c := connectionConfig{
		uri:              uri,
		caCertFile:       config.Config.MySQLOrchestratorSSLCAFile,
		useClientAuth:    config.Config.MySQLOrchestratorUseMutualTLS,
		clientCertFile:   config.Config.MySQLOrchestratorSSLCertFile,
		clientKeyFile:    config.Config.MySQLOrchestratorSSLPrivateKeyFile,
		verifyClientCert: !config.Config.MySQLOrchestratorSSLSkipVerify,
		verifyServerCert: !config.Config.MySQLOrchestratorSSLSkipVerify,
	}
	return setupTls("orchestrator", c)
}

func setupTls(name string, c connectionConfig) (string, error) {
	if _, found := tlsConfigCache.Get(c.uri); !found {
		tlsConfig, err := ssl.NewTLSConfig(c.caCertFile, c.verifyClientCert)
		// Drop to TLS 1.0 for talking to MySQL
		tlsConfig.MinVersion = tls.VersionTLS10
		if err != nil {
			return "", log.Fatalf("Can't create TLS configuration for %s connection %s: %s", name, c.uri, err)
		}
		tlsConfig.InsecureSkipVerify = !c.verifyServerCert

		if c.useClientAuth && c.verifyServerCert && c.clientCertFile != "" && c.clientKeyFile != "" {
			if err = ssl.AppendKeyPair(tlsConfig, c.clientCertFile, c.clientKeyFile); err != nil {
				return "", log.Fatalf("Can't setup TLS key pairs for %s: %s", c.uri, err)
			}
		}

		if err = mysql.RegisterTLSConfig(name, tlsConfig); err != nil {
			return "", log.Fatalf("Can't register mysql TLS config for %s: %s", name, err)
		}
		tlsConfigCache.Set(c.uri, tlsConfig, cache.NoExpiration)
	}
	return fmt.Sprintf("%s&tls=%s", c.uri, name), nil
}

type connectionConfig struct {
	uri              string
	caCertFile       string
	useClientAuth    bool
	clientCertFile   string
	clientKeyFile    string
	verifyClientCert bool
	verifyServerCert bool
}
