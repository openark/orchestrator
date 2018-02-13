package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/openark/golib/log"

	"github.com/github/orchestrator/go/raft"
	"github.com/go-martini/martini"
)

var reverseProxy = func(w http.ResponseWriter, r *http.Request, c martini.Context) {
	if !orcraft.IsRaftEnabled() {
		// No raft, so no reverse proxy to the leader
		return
	}
	if orcraft.IsLeader() {
		// I am the leader. I will handle the request directly.
		return
	}
	url, err := url.Parse(orcraft.LeaderURI.Get())
	if err != nil {
		log.Errore(err)
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(url)
	log.Debugf("................reverse proxy to %s", url)
	proxy.ServeHTTP(w, r)
}

// All returns a Handler that adds gzip compression to all requests
func RaftReverseProxy() martini.Handler {
	return func(w http.ResponseWriter, r *http.Request, c martini.Context) {
		reverseProxy(w, r, c)
	}
}
