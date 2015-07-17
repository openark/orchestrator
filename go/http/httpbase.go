/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

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

package http

import (
	"github.com/martini-contrib/auth"
	"github.com/outbrain/orchestrator/go/config"
	"net/http"
	"strings"
)

func getProxyAuthUser(req *http.Request) string {
	for _, user := range req.Header[config.Config.AuthUserHeader] {
		return user
	}
	return ""
}

// isAuthorizedForAction checks req to see whether authenticated user has write-privileges.
// This depends on configured authentication method.
func isAuthorizedForAction(req *http.Request, user auth.User) bool {
	if config.Config.ReadOnly {
		return false
	}

	switch strings.ToLower(config.Config.AuthenticationMethod) {
	case "basic":
		{
			// The mere fact we're here means the user has passed authentication
			return true
		}
	case "multi":
		{
			if string(user) == "readonly" {
				// read only
				return false
			}
			// passed authentication ==> writeable
			return true
		}
	case "proxy":
		{
			authUser := getProxyAuthUser(req)
			for _, configPowerAuthUser := range config.Config.PowerAuthUsers {
				if configPowerAuthUser == "*" || configPowerAuthUser == authUser {
					return true
				}
			}
			return false
		}
	default:
		{
			// Default: no authentication method
			return true
		}
	}
}

// getUserId returns the authenticated user id, if available, depending on authertication method.
func getUserId(req *http.Request, user auth.User) string {
	if config.Config.ReadOnly {
		return ""
	}

	switch strings.ToLower(config.Config.AuthenticationMethod) {
	case "basic":
		{
			return string(user)
		}
	case "multi":
		{
			return string(user)
		}
	case "proxy":
		{
			return getProxyAuthUser(req)
		}
	default:
		{
			return ""
		}
	}
}
