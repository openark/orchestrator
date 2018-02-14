package http

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/openark/golib/log"

	"github.com/github/orchestrator/go/raft"
	"github.com/go-martini/martini"
)

func raftReverseProxy(w http.ResponseWriter, r *http.Request, c martini.Context) {
	if !orcraft.IsRaftEnabled() {
		// No raft, so no reverse proxy to the leader
		return
	}
	if orcraft.IsLeader() {
		// I am the leader. I will handle the request directly.
		return
	}
	if orcraft.GetLeader() == "" {
		return
	}

	url, err := url.Parse(orcraft.LeaderURI.Get())
	if err != nil {
		log.Errore(err)
		return
	}
	r.Header.Del("Accept-Encoding")
	proxy := httputil.NewSingleHostReverseProxy(url)
	proxy.ServeHTTP(w, r)
}
