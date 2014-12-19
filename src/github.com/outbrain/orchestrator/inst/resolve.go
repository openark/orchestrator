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

package inst

import (
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/config"
	"github.com/pmylund/go-cache"
	"net"
	"strings"
	"time"
)

func init() {
	if config.Config.ExpiryHostnameResolvesMinutes < 1 {
		config.Config.ExpiryHostnameResolvesMinutes = 1
	}
}

var hostnameResolvesLightweightCache = cache.New(time.Duration(config.Config.ExpiryHostnameResolvesMinutes)*time.Minute, time.Minute)

// GetCNAME resolves an IP or hostname into a normalized valid CNAME
func GetCNAME(hostname string) (string, error) {
	res, err := net.LookupCNAME(hostname)
	if err != nil {
		return hostname, err
	}
	res = strings.TrimRight(res, ".")
	return res, nil
}

func resolveHostname(hostname string) (string, error) {
	switch strings.ToLower(config.Config.HostnameResolveMethod) {
	case "none":
		return hostname, nil
	case "cname":
		return GetCNAME(hostname)
	}
	return hostname, nil
}

// Attempt to resolve a hostname. This may returned a database cached hostname or otherwise
// it may resolve the hostname via CNAME
func ResolveHostname(hostname string) (string, error) {

	// First go to lightweight cache
	if resolvedHostname, found := hostnameResolvesLightweightCache.Get(hostname); found {
		return resolvedHostname.(string), nil
	}

	// Unfound? Go to database cache
	// Database cache actually has SHORTER expiry, and serves just for the sake of a service restart, to help
	// warm up the cache
	if resolvedHostname, _ := ReadResolvedHostname(hostname); resolvedHostname != "" {
		hostnameResolvesLightweightCache.Set(hostname, resolvedHostname, 0)
		return resolvedHostname, nil
	}

	// Unfound: resolve!
	log.Debugf("Unfound: %s", hostname)
	resolvedHostname, err := resolveHostname(hostname)
	if err != nil {
		UpdateResolvedHostname(hostname, hostname)
		return hostname, err
	}
	// Cache result
	log.Debugf("Cache %s as %s", hostname, resolvedHostname)
	UpdateResolvedHostname(hostname, resolvedHostname)
	return resolvedHostname, nil
}

// UpdateResolvedHostname will store the given resolved hostname in cache
func UpdateResolvedHostname(hostname string, resolvedHostname string) {
	hostnameResolvesLightweightCache.Set(hostname, resolvedHostname, 0)
	WriteResolvedHostname(hostname, resolvedHostname)
}
