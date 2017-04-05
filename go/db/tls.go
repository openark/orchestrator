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
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/openark/golib/log"

	"github.com/github/orchestrator/go/config"
	"github.com/github/orchestrator/go/ssl"
)

const Error3159 = "Error 3159: Connections using insecure transport are prohibited while --require_secure_transport=ON."

const (
	noName            string = ""         // explicitly empty
	topologyName      string = "topology" // for MySQLTopologyUseMutualTLS
	tlsWarningSeconds        = 60
)

// configKey contains the hostname and port to connect to.
type configKey struct {
	host string
	port int
}

// configValue contains the TLS configuration information for a configKey
type configValue struct {
	required   bool
	configured bool
}

// global structure for managing TLS configuration either for Mutual
// TLS configuration or an individual setup
type TLSConfig struct {
	lock                   sync.Mutex
	config                 map[configKey]configValue
	lastWarningTime        time.Time
	topologyConfigured     bool // Track if the MutualTLS configuration has been configured
	orchestratorConfigured bool // Track if a TLS has already been configured for Orchestrator
}

var tc *TLSConfig = NewTLSConfig() // setup new TLS config

// return a new TLS configuration structure
func NewTLSConfig() *TLSConfig {
	return &TLSConfig{
		config: make(map[configKey]configValue),
	}
}

// NewConfigKey returns a configKey based on the given host/port
func NewConfigKey(host string, port int) configKey {
	return configKey{
		host: host,
		port: port,
	}
}

// isRequired returns true if the host is known to require a
// TLS configuration.  Lookup in tc.config
func (tc *TLSConfig) isRequired(host string, port int) bool {
	key := NewConfigKey(host, port)
	v, _ := tc.config[key]

	return v.required
}

// isConfigured returns true if the host is already confiured for TLS
func (tc *TLSConfig) isConfigured(host string, port int) bool {
	key := NewConfigKey(host, port)
	v, _ := tc.config[key]

	return v.configured
}

// setConfigured marks the host port as configured
func (tc *TLSConfig) setConfigured(host string, port int) {
	key := NewConfigKey(host, port)
	if v, found := tc.config[key]; found {
		v.configured = true
		tc.config[key] = v
	} else {
		log.Fatalf("setConfigured: key %+v not found", key)
	}
}

// setRequired registers the fact the host requires TLS connectivity.
func (tc *TLSConfig) setRequired(host string, port int, required bool) {
	key := NewConfigKey(host, port)

	v, _ := tc.config[key]

	if v.required != required {
		v.required = required
	}

	if required {
		// if true set the new value
		tc.config[key] = v
	} else {
		// if false delete the entry to keep the map small
		delete(tc.config, key)
	}
}

// CatchTLSRequirementError checks for a TLS requirement error and if found
// remembers this for a later connection.
func CatchTLSRequirementError(host string, port int, err error) {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	if err != nil && !tc.isRequired(host, port) && strings.Contains(err.Error(), Error3159) {
		log.Infof("CatchTLSRequirementError: Instance %s:%d requires TLS, adjusting config for next discovery", host, port)
		tc.setRequired(host, port, true)
		// FIXME - need to throw the existing non-TLS pool handle away
	}
}

// configName returns the name used in the TLS configuration
func configName(host string, port int) string {
	if config.Config.MySQLTopologyUseMutualTLS {
		return topologyName
	}

	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%s:%d", host, port)))

	return hex.EncodeToString(hasher.Sum(nil))
}

// ConfigurationName creates a TLS configuration if required and returns a
// TLS configuration name to the caller.
// * if MySQLTopologyUseMutualTLS is configured then use the required
//   CA, Certificate, and Private key and register the "topology" config
// * if an individual configuration is required use the required CA,
//   use Certificate and Private key (if provided) and register with a
//   name based on the hostname/port.
// * if TLS is required for the host and no CA is provided generating a
//   periodic warning so at least the CA is configured.
// * TODO - this configuration is not dynamic and won't handle orchestrator
//   reconfiguration via SIGHUP
func (tc *TLSConfig) ConfigurationName(host string, port int) (string, error) {
	var name string
	tc.lock.Lock()
	defer tc.lock.Unlock()

	// no configuration required?
	if !config.Config.MySQLTopologyUseMutualTLS && !tc.isRequired(host, port) {
		return noName, nil
	}

	name = configName(host, port)

	// already configured?
	if config.Config.MySQLTopologyUseMutualTLS {
		if tc.topologyConfigured {
			return name, nil
		}
	} else {
		if tc.isConfigured(host, port) {
			return name, nil
		}

		// If the CAFile is not configured then warn at most once a minute.
		if config.Config.MySQLTopologySSLCAFile == "" && time.Since(tc.lastWarningTime) > (tlsWarningSeconds*time.Second) {
			log.Infof("Minimal TLS configuration is required to access %s:%d. Please set MySQLTopologySSLCAFile appropriately.", host, port)
			tc.lastWarningTime = time.Now() // we're in a lock so need for extra locking
			return noName, nil              // the uri is likely to fail but we can't do any better.
		}
	}

	// setup the TLS configuration
	tlsConfig, err := ssl.NewTLSConfig(config.Config.MySQLTopologySSLCAFile, !config.Config.MySQLTopologySSLSkipVerify)
	if err != nil {
		return noName, log.Fatalf("Can't create TLS configuration for topology connection <%s:%d>: %s", host, port, err)
	}

	tlsConfig.MinVersion = tls.VersionTLS10 // Drop to TLS 1.0 for talking to MySQL
	tlsConfig.InsecureSkipVerify = config.Config.MySQLTopologySSLSkipVerify

	// If we want MutualTLS or we provide TLS cert configuration
	// then attempt to use it and fail if it does not work
	if config.Config.MySQLTopologyUseMutualTLS ||
		config.Config.MySQLTopologySSLCertFile != "" ||
		config.Config.MySQLTopologySSLPrivateKeyFile != "" {
		if err = ssl.AppendKeyPair(tlsConfig, config.Config.MySQLTopologySSLCertFile, config.Config.MySQLTopologySSLPrivateKeyFile); err != nil {
			return noName, log.Fatalf("Can't setup TLS key pairs for <%s:%d>: %s", host, port, err)
		}
	}

	// try to register config
	if err = mysql.RegisterTLSConfig(name, tlsConfig); err != nil {
		return noName, log.Fatalf("Can't register mysql TLS config for <%s:%d> (%q): %s", host, port, name, err)
	}

	// record the config has been made depending on usage
	if config.Config.MySQLTopologyUseMutualTLS {
		tc.topologyConfigured = true
	} else {
		tc.setConfigured(host, port)
	}
	return name, nil
}

// Create a TLS configuration from the config supplied CA, Certificate, and Private key.
// Register the TLS config with the mysql drivers as the "orchestrator" config
// Modify the supplied URI to call the TLS config
func SetupMySQLOrchestratorTLS(uri string) (string, error) {
	tc.lock.Lock()
	defer tc.lock.Unlock()

	if !tc.orchestratorConfigured {
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
		tc.orchestratorConfigured = true
	}
	return fmt.Sprintf("%s&tls=orchestrator", uri), nil
}
