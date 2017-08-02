/*
   Copyright 2017 Shlomi Noach, GitHub Inc.

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

package orcraft

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/github/orchestrator/go/config"

	"github.com/openark/golib/log"
)

var httpClient *http.Client

func setupHttpClient() error {
	httpTimeout := time.Duration(config.ActiveNodeExpireSeconds) * time.Second
	dialTimeout := func(network, addr string) (net.Conn, error) {
		return net.DialTimeout(network, addr, httpTimeout)
	}
	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: config.Config.MySQLOrchestratorSSLSkipVerify},
		Dial:            dialTimeout,
		ResponseHeaderTimeout: httpTimeout,
	}
	httpClient = &http.Client{Transport: httpTransport}

	return nil
}

// PeerAPI generates the API path of a raft peer
func peerAPI(peer string) string {
	protocol := "http"
	if config.Config.UseSSL {
		protocol = "https"
	}
	api := fmt.Sprintf("%s://%s%s/api", protocol, peer, config.Config.ListenAddress)
	return api
}

// readResponse returns the body of an HTTP response
func HttpGet(peer string, path string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", peerAPI(peer), path)
	res, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return body, log.Errorf("HttpGet: got %d status on %s", res.StatusCode, url)
	}

	return body, nil
}
