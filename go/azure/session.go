package azure

import (
	"github.com/martini-contrib/sessions"
)

func GetUsername(session sessions.Session) string {
	username := session.Get("Username")
	if username == nil {
		return ""
	}

	return username.(string)
}

func GetUserRole(session sessions.Session) string {
	role := session.Get("Role")
	if role == nil {
		return ""
	}

	return role.(string)
}

func IsUserLogged(session sessions.Session) bool {
	username := GetUsername(session)
	role := GetUserRole(session)

	if username != "" && role != "" {
		return true
	}

	return false
}
