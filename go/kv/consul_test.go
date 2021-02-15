package kv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strings"
	"testing"

	consulapi "github.com/hashicorp/consul/api"
)

const consulTestDefaultDatacenter = "dc1"

type consulTestServerOp struct {
	Method       string
	URL          string
	Request      interface{}
	Response     interface{}
	ResponseCode int
}

// sortTxnKVOps sort TxnOps by op.KV.Key to resolve random test failures
func sortTxnKVOps(txnOps []*consulapi.TxnOp) []*consulapi.TxnOp {
	sort.Slice(txnOps, func(a, b int) bool {
		return txnOps[a].KV.Key < txnOps[b].KV.Key
	})
	return txnOps
}

func buildConsulTestServer(t *testing.T, testOps []consulTestServerOp) *httptest.Server {
	handlerFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestBytes, _ := ioutil.ReadAll(r.Body)
		requestBody := strings.TrimSpace(string(requestBytes))

		for _, testOp := range testOps {
			if r.Method != testOp.Method || r.URL.String() != testOp.URL {
				continue
			}
			if testOp.ResponseCode == 0 {
				testOp.ResponseCode = http.StatusOK
			}
			if strings.HasPrefix(r.URL.String(), "/v1/catalog/datacenters") {
				w.WriteHeader(testOp.ResponseCode)
				json.NewEncoder(w).Encode(testOp.Response)
				return
			} else if strings.HasPrefix(r.URL.String(), "/v1/kv") && testOp.Response != nil {
				w.WriteHeader(testOp.ResponseCode)
				json.NewEncoder(w).Encode(testOp.Response)
				return
			} else if strings.HasPrefix(r.URL.String(), "/v1/txn") {
				var txnOps consulapi.TxnOps
				if err := json.Unmarshal(requestBytes, &txnOps); err != nil {
					t.Fatalf("Unable to unmarshal json request body: %v", err)
					continue
				}
				testOpRequest := sortTxnKVOps(testOp.Request.(consulapi.TxnOps))
				if testOp.Response != nil && reflect.DeepEqual(testOpRequest, sortTxnKVOps(txnOps)) {
					w.WriteHeader(testOp.ResponseCode)
					json.NewEncoder(w).Encode(testOp.Response)
					return
				}
			}
		}

		t.Fatalf("No requests matched setup. Got method %s, Path %s, body %s", r.Method, r.URL.String(), requestBody)
		w.WriteHeader(http.StatusNotFound)
		_, _ = fmt.Fprintln(w, "")
	})
	return httptest.NewServer(handlerFunc)
}
