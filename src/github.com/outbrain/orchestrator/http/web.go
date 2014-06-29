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

package http

import (
	"strconv"
	"fmt"	
	"errors"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"

	"github.com/outbrain/orchestrator/inst"
)

// HttpWeb is the web requests server, mapping each request to a web page
type HttpWeb struct{}

var Web HttpWeb = HttpWeb{}



func (this *HttpWeb) getInstanceKey(host string, port string) (inst.InstanceKey, error) {
	instanceKey := inst.InstanceKey{Hostname: host}
	var err error

	if instanceKey.Port, err = strconv.Atoi(port); err != nil {
		return instanceKey, errors.New(fmt.Sprintf("Invalid port: %s", port))
	}
	return instanceKey, err
}


func (this *HttpWeb) Cluster(params martini.Params, r render.Render) {
	r.HTML(200, "templates/cluster", map[string]interface{}{
		"title": "cluster", 
		"activePage": "cluster", 
		"clusterName": params["clusterName"],
		"autoshow_problems": true,
		})
}


func (this *HttpWeb) Search(params martini.Params, r render.Render, req *http.Request) {
	searchString := params["searchString"]
	if searchString == "" {
		searchString = req.URL.Query().Get("s");
	}
	r.HTML(200, "templates/search", map[string]interface{}{
		"title": "search", 
		"activePage": "search", 
		"searchString": searchString, 
		"autoshow_problems": false,
		})
}


func (this *HttpWeb) Discover(params martini.Params, r render.Render) {

	r.HTML(200, "templates/discover", map[string]interface{}{
		"title": "discover", 
		"activePage": "discover", 
		"autoshow_problems": false,
		})
}


func (this *HttpWeb) Audit(params martini.Params, r render.Render) {
	page, err := strconv.Atoi(params["page"])
	if err != nil { page = 0 }

	r.HTML(200, "templates/audit", map[string]interface{}{
		"title": "audit", 
		"activePage": "audit", 
		"autoshow_problems": false,
		"page": page,
		})
}


func (this *HttpWeb) Home(params martini.Params, r render.Render) {

	r.HTML(200, "templates/home", map[string]interface{}{
		"title": "home", 
		"activePage": "home", 
		"autoshow_problems": false,
		})
}


func (this *HttpWeb) About(params martini.Params, r render.Render) {

	r.HTML(200, "templates/about", map[string]interface{}{
		"title": "about", 
		"activePage": "home", 
		"autoshow_problems": false,
		})
}



func (this *HttpWeb) FAQ(params martini.Params, r render.Render) {

	r.HTML(200, "templates/faq", map[string]interface{}{
		"title": "FAQ", 
		"activePage": "home", 
		"autoshow_problems": false,
		})
}



// RegisterRequests makes for the de-facto list of known Web calls
func (this *HttpWeb) RegisterRequests(m *martini.ClassicMartini) {
	m.Get("/", this.About) 
	m.Get("/web", this.About) 
	m.Get("/web/home", this.About) 
	m.Get("/web/about", this.About) 
	m.Get("/web/faq", this.FAQ) 
	m.Get("/web/cluster/:clusterName", this.Cluster) 
	m.Get("/web/search/:searchString", this.Search) 
	m.Get("/web/search", this.Search) 
	m.Get("/web/discover", this.Discover) 
	m.Get("/web/audit", this.Audit) 
	m.Get("/web/audit/:page", this.Audit) 
}
