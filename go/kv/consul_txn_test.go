package kv

import (
	"net/http"
	"strings"
	"sync"
	"testing"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/openark/orchestrator/go/config"
)

func TestGroupKVPairsByPrefix(t *testing.T) {
	// batch 9 x KVPairs into 2 x transactions
	config.Config.ConsulMaxKVsPerTransaction = 6
	grouped := groupKVPairsByPrefix(consulapi.KVPairs{
		{Key: "mysql/master/cluster1/hostname", Value: []byte("test")},
		{Key: "mysql/master/cluster1/ipv4", Value: []byte("test")},
		{Key: "mysql/master/cluster1/port", Value: []byte("test")},
		{Key: "mysql/master/cluster2/hostname", Value: []byte("test")},
		{Key: "mysql/master/cluster2/ipv4", Value: []byte("test")},
		{Key: "mysql/master/cluster2/port", Value: []byte("test")},
		{Key: "mysql/master/cluster3/hostname", Value: []byte("test")},
		{Key: "mysql/master/cluster3/ipv4", Value: []byte("test")},
		{Key: "mysql/master/cluster3/port", Value: []byte("test")},
	})
	if len(grouped) != 2 {
		t.Fatalf("expected 2 groups, got %d: %v", len(grouped), grouped)
	}
	if len(grouped[0]) != 6 {
		t.Fatalf("expected 6 KVPairs in first group, got %d: %v", len(grouped[0]), grouped[0])
	}
	if len(grouped[1]) != 3 {
		t.Fatalf("expected 3 KVPairs in second group, got %d: %v", len(grouped[1]), grouped[1])
	}

	// check KVs for a cluster are in a single group
	clusterCounts := map[string]map[int]int{}
	for i, group := range grouped {
		for _, kvPair := range group {
			s := strings.Split(kvPair.Key, "/")
			clusterName := s[2]
			if _, ok := clusterCounts[clusterName]; ok {
				clusterCounts[clusterName][i]++
			} else {
				clusterCounts[clusterName] = map[int]int{i: 1}
			}
		}
	}
	for cluster, groups := range clusterCounts {
		if len(groups) != 1 {
			t.Fatalf("expected %s to be in a single group, found it in %d group(s): %v", cluster, len(groups), groups)
		}
		for _, count := range groups {
			if count != 3 {
				t.Fatalf("expected group to contain 3 x %s keys, found: %d", cluster, count)
			}
		}
	}
}

func TestConsulTxnStorePutKVPairs(t *testing.T) {
	server := buildConsulTestServer(t, []consulTestServerOp{
		{
			Method:   "PUT",
			URL:      "/v1/kv/kv1",
			Request:  "test",
			Response: &consulapi.KVPair{Key: "kv1", Value: []byte("test")},
		},
		{
			Method: "PUT",
			URL:    "/v1/txn",
			Request: consulapi.TxnOps{
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVSet, Key: "kv1", Value: []byte("test")},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVSet, Key: "kv2", Value: []byte("test2")},
				},
			},
			Response: &consulapi.TxnResponse{
				Results: consulapi.TxnResults{
					{
						KV: &consulapi.KVPair{Key: "kv1", Value: []byte("test")},
					},
					{
						KV: &consulapi.KVPair{Key: "kv2", Value: []byte("test2")},
					},
				},
			},
		},
		{
			Method: "PUT",
			URL:    "/v1/txn",
			Request: consulapi.TxnOps{
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVSet, Key: "fail1", Value: []byte("test")},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVSet, Key: "fail2", Value: []byte("test2")},
				},
			},
			Response: &consulapi.TxnResponse{
				Errors: consulapi.TxnErrors{
					{
						What: "test error",
					},
				},
			},
			ResponseCode: http.StatusConflict, // PUT /v1/txn returns a HTTP 409 code on txn failure
		},
	})
	defer server.Close()
	config.Config.ConsulAddress = server.URL
	store := NewConsulTxnStore()

	t.Run("single-kv", func(t *testing.T) {
		if err := store.PutKVPairs([]*KVPair{
			{Key: "kv1", Value: "test"},
		}); err != nil {
			t.Fatalf("Unable to run .PutKVPairs(): %v", err)
		}
	})

	t.Run("multi-kv", func(t *testing.T) {
		if err := store.PutKVPairs([]*KVPair{
			{Key: "kv1", Value: "test"},
			{Key: "kv2", Value: "test2"},
		}); err != nil {
			t.Fatalf("Unable to run .PutKVPairs(): %v", err)
		}
	})

	t.Run("multi-kv-failed", func(t *testing.T) {
		if err := store.PutKVPairs([]*KVPair{
			{Key: "fail1", Value: "test"},
			{Key: "fail2", Value: "test2"},
		}); err == nil || err.Error() != "consul txn error: test error" {
			t.Fatalf("Expected %q error from .PutKVPairs(), got: %q", "consul txn error: test error", err.Error())
		}
	})
}

func TestConsulTxnStoreUpdateDatacenterKVPairs(t *testing.T) {
	server := buildConsulTestServer(t, []consulTestServerOp{
		{
			Method: "PUT",
			URL:    "/v1/txn?dc=dc1",
			Request: consulapi.TxnOps{
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test"}, // exists
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test2"}, // does not exist
				},
			},
			Response: &consulapi.TxnResponse{
				Results: consulapi.TxnResults{
					{
						KV: &consulapi.KVPair{Key: "test", Value: []byte("test")},
					},
				},
			},
		},
		{
			Method: "PUT",
			URL:    "/v1/txn?dc=dc1",
			Request: consulapi.TxnOps{
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVSet, Key: "test2", Value: []byte("test")},
				},
			},
			Response: &consulapi.TxnResponse{
				Results: consulapi.TxnResults{
					{
						KV: &consulapi.KVPair{Key: "test2", Value: []byte("test")},
					},
				},
			},
		},
	})
	defer server.Close()
	config.Config.ConsulAddress = server.URL

	store := NewConsulTxnStore().(*consulTxnStore)
	cacheKey := getConsulKVCacheKey(consulTestDefaultDatacenter, "cached")
	store.kvCache.SetDefault(cacheKey, "cached") // pre-cache the 'cached' key-value

	var wg sync.WaitGroup
	wg.Add(1)
	kvPairs := []*consulapi.KVPair{
		{Key: "cached", Value: []byte("cached")}, // already cached key-value
		{Key: "test", Value: []byte("test")},     // already correct on consul server
		{Key: "test2", Value: []byte("test")},    // does not exist on consul server
	}
	skipped, existing, written, failed, err := store.updateDatacenterKVPairs(&wg, consulTestDefaultDatacenter, kvPairs)
	if err != nil {
		t.Fatalf(".updateDatacenterKVPairs() should not return an error, got: %v", err)
	}
	if skipped != 1 || existing != 1 || written != 1 || failed != 0 {
		t.Fatalf("expected: existing/skipped/written=1 and failed=0, got: skipped=%d, existing=%d, written=%d, failed=%d",
			skipped, existing, written, failed,
		)
	}

	for _, pair := range kvPairs {
		cacheKey := getConsulKVCacheKey(consulTestDefaultDatacenter, pair.Key)
		if cached, found := store.kvCache.Get(cacheKey); !found || cached != string(pair.Value) {
			t.Fatalf("expected cache key %q to equal %q, got %v", cacheKey, string(pair.Value), cached)
		}
	}
}
