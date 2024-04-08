package authentication

import (
	u "net/url"
	"strings"
)

// AllowedHostsValidator Maintains a list of valid hosts and allows authentication providers to check whether a host is valid before authenticating a request
type AllowedHostsValidator struct {
	validHosts map[string]bool
}

// NewAllowedHostsValidator creates a new AllowedHostsValidator object with provided values.
func NewAllowedHostsValidator(validHosts []string) AllowedHostsValidator {
	result := AllowedHostsValidator{}
	result.SetAllowedHosts(validHosts)
	return result
}

// GetAllowedHosts returns the list of valid hosts.
func (v *AllowedHostsValidator) GetAllowedHosts() map[string]bool {
	hosts := make(map[string]bool, len(v.validHosts))
	for host := range v.validHosts {
		hosts[host] = true
	}
	return hosts
}

// SetAllowedHosts sets the list of valid hosts.
func (v *AllowedHostsValidator) SetAllowedHosts(hosts []string) {
	v.validHosts = make(map[string]bool, len(hosts))
	if len(hosts) > 0 {
		for _, host := range hosts {
			v.validHosts[strings.ToLower(host)] = true
		}
	}
}

// IsValidHost returns true if the host is valid.
func (v *AllowedHostsValidator) IsUrlHostValid(uri *u.URL) bool {
	if uri == nil {
		return false
	}
	host := uri.Hostname()
	if host == "" {
		return false
	}
	return len(v.validHosts) == 0 || v.validHosts[strings.ToLower(host)]
}
