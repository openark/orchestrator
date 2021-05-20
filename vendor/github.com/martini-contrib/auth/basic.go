package auth

import (
	"encoding/base64"
	"github.com/go-martini/martini"
	"net/http"
	"strings"
)

// User is the authenticated username that was extracted from the request.
type User string

// BasicRealm is used when setting the WWW-Authenticate response header.
var BasicRealm = "Authorization Required"

// Basic returns a Handler that authenticates via Basic Auth. Writes a http.StatusUnauthorized
// if authentication fails.
func Basic(username string, password string) martini.Handler {
	var siteAuth = base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		auth := req.Header.Get("Authorization")
		if !SecureCompare(auth, "Basic "+siteAuth) {
			unauthorized(res)
			return
		}
		c.Map(User(username))
	}
}

// BasicFunc returns a Handler that authenticates via Basic Auth using the provided function.
// The function should return true for a valid username/password combination.
func BasicFunc(authfn func(string, string) bool) martini.Handler {
	return func(res http.ResponseWriter, req *http.Request, c martini.Context) {
		auth := req.Header.Get("Authorization")
		if len(auth) < 6 || auth[:6] != "Basic " {
			unauthorized(res)
			return
		}
		b, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			unauthorized(res)
			return
		}
		tokens := strings.SplitN(string(b), ":", 2)
		if len(tokens) != 2 || !authfn(tokens[0], tokens[1]) {
			unauthorized(res)
			return
		}
		c.Map(User(tokens[0]))
	}
}

func unauthorized(res http.ResponseWriter) {
	res.Header().Set("WWW-Authenticate", "Basic realm=\""+BasicRealm+"\"")
	http.Error(res, "Not Authorized", http.StatusUnauthorized)
}
