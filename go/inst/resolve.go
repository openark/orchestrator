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
	"errors"
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/pmylund/go-cache"
	"net"
	"regexp"
	"strings"
	"time"
)

type HostnameResolve struct {
	hostname         string
	resolvedHostname string
}

func init() {
	if config.Config.ExpiryHostnameResolvesMinutes < 1 {
		config.Config.ExpiryHostnameResolvesMinutes = 1
	}
}

var hostnameResolvesLightweightCache = cache.New(time.Duration(config.Config.ExpiryHostnameResolvesMinutes)*time.Minute, time.Minute)
var hostnameResolvesLightweightCacheLoadedOnceFromDB bool = false

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
	case "default":
		return hostname, nil
	case "cname":
		return GetCNAME(hostname)
	}
	return hostname, nil
}

// Attempt to resolve a hostname. This may returned a database cached hostname or otherwise
// it may resolve the hostname via CNAME
func ResolveHostname(hostname string) (string, error) {
	hostname = strings.TrimSpace(hostname)
	if hostname == "" {
		return hostname, errors.New("Will not resolve empty hostname")
	}
	// First go to lightweight cache
	if resolvedHostname, found := hostnameResolvesLightweightCache.Get(hostname); found {
		return resolvedHostname.(string), nil
	}

	if !hostnameResolvesLightweightCacheLoadedOnceFromDB {
		// A continuous-discovery will first make sure to load all resolves from DB.
		// However cli does not do so.
		// Anyway, it seems like the cache was not loaded from DB. Before doing real resolves,
		// let's try and get the resolved hostname from database.
		if resolvedHostname, err := ReadResolvedHostname(hostname); err == nil && resolvedHostname != "" {
			hostnameResolvesLightweightCache.Set(hostname, resolvedHostname, 0)
			return resolvedHostname, nil
		}
	}

	// Unfound: resolve!
	log.Debugf("Hostname unresolved yet: %s", hostname)
	resolvedHostname, err := resolveHostname(hostname)
	if config.Config.RejectHostnameResolvePattern != "" {
		// Reject, don't even cache
		if matched, _ := regexp.MatchString(config.Config.RejectHostnameResolvePattern, resolvedHostname); matched {
			return hostname, fmt.Errorf("Resolved hostname is rejected: %s", resolvedHostname)
		}
	}

	if err != nil {
		// Problem. What we'll do is cache the hostname for just one minute, so as to avoid flooding requests
		// on one hand, yet make it refresh shortly on the other hand. Anyway do not write to database.
		hostnameResolvesLightweightCache.Set(hostname, resolvedHostname, time.Minute)
		return hostname, err
	}
	// Good result! Cache it, also to DB
	log.Debugf("Cache hostname resolve %s as %s", hostname, resolvedHostname)
	UpdateResolvedHostname(hostname, resolvedHostname)
	return resolvedHostname, nil
}

// UpdateResolvedHostname will store the given resolved hostname in cache
// Returns false when the key already existed with same resolved value (similar
// to AFFECTED_ROWS() in mysql)
func UpdateResolvedHostname(hostname string, resolvedHostname string) bool {
	if existingResolvedHostname, found := hostnameResolvesLightweightCache.Get(hostname); found && (existingResolvedHostname == resolvedHostname) {
		return false
	}
	hostnameResolvesLightweightCache.Set(hostname, resolvedHostname, 0)
	if strings.ToLower(config.Config.HostnameResolveMethod) != "none" {
		WriteResolvedHostname(hostname, resolvedHostname)
	}
	return true
}

func LoadHostnameResolveCacheFromDatabase() error {
	allHostnamesResolves, err := readAllHostnameResolves()
	if err != nil {
		return err
	}
	for _, hostnameResolve := range allHostnamesResolves {
		hostnameResolvesLightweightCache.Set(hostnameResolve.hostname, hostnameResolve.resolvedHostname, 0)
	}
	hostnameResolvesLightweightCacheLoadedOnceFromDB = true
	return nil
}

func ResetHostnameResolveCache() error {
	err := deleteHostnameResolves()
	hostnameResolvesLightweightCache.Flush()
	hostnameResolvesLightweightCacheLoadedOnceFromDB = false
	return err
}

func HostnameResolveCache() (map[string]*cache.Item, error) {
	return hostnameResolvesLightweightCache.Items(), nil
}

func UnresolveHostname(instanceKey *InstanceKey) (InstanceKey, bool, error) {
	unresolvedHostname, err := readUnresolvedHostname(instanceKey.Hostname)
	if err != nil {
		return *instanceKey, false, log.Errore(err)
	}
	if unresolvedHostname == instanceKey.Hostname {
		// unchanged. Nothing to do
		return *instanceKey, false, nil
	}
	// We unresovled to a different hostname. We will now re-resolve to double-check!
	unresolvedKey := &InstanceKey{Hostname: unresolvedHostname, Port: instanceKey.Port}

	instance, err := ReadTopologyInstance(unresolvedKey)
	if err != nil {
		return *instanceKey, false, log.Errore(err)
	}
	if instance.Key.Hostname != instanceKey.Hostname {
		// Resolve(Unresolve(hostname)) != hostname ==> Bad; reject
		if *config.RuntimeCLIFlags.SkipUnresolveCheck {
			return *instanceKey, false, nil
		}
		return *instanceKey, false, log.Errorf("Error unresolving; hostname=%s, unresolved=%s, re-resolved=%s; mismatch. Skip/ignore with --skip-unresolve-check", instanceKey.Hostname, unresolvedKey.Hostname, instance.Key.Hostname)
	}
	return *unresolvedKey, true, nil
}

func RegisterHostnameUnresolve(instanceKey *InstanceKey, unresolvedHostname string) (err error) {
	return WriteHostnameUnresolve(instanceKey, unresolvedHostname)
}
