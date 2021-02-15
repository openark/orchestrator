package kv

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"testing"

	consulapi "github.com/hashicorp/consul/api"
	"github.com/openark/orchestrator/go/config"
)

func TestGroupKVPairsByPrefix(t *testing.T) {
	config.Config.ConsulMaxKVsPerTransaction = 10
	config.Config.KVClusterMasterPrefix = "mysql/master"

	// make 100 KVs for 20 clusters
	kvPairs := consulapi.KVPairs{}
	var kvs int
	for kvs < 100 {
		kvPairs = append(kvPairs,
			&consulapi.KVPair{
				Key:   fmt.Sprintf("%s/cluster%d", config.Config.KVClusterMasterPrefix, kvs),
				Value: []byte("mysql.example.com:3306"),
			},
			&consulapi.KVPair{
				Key:   fmt.Sprintf("%s/cluster%d/hostname", config.Config.KVClusterMasterPrefix, kvs),
				Value: []byte("mysql.example.com"),
			},
			&consulapi.KVPair{
				Key:   fmt.Sprintf("%s/cluster%d/ipv4", config.Config.KVClusterMasterPrefix, kvs),
				Value: []byte("10.20.30.40"),
			},
			&consulapi.KVPair{
				Key:   fmt.Sprintf("%s/cluster%d/ipv6", config.Config.KVClusterMasterPrefix, kvs),
				Value: []byte("fdf0:7a53:0b88:d147:xxxx:xxxx:xxxx:xxxx"),
			},
			&consulapi.KVPair{
				Key:   fmt.Sprintf("%s/cluster%d/port", config.Config.KVClusterMasterPrefix, kvs),
				Value: []byte("3306"),
			},
		)
		kvs += 5
	}

	grouped := groupKVPairsByPrefix(kvPairs)
	if len(grouped) != 10 {
		t.Fatalf("expected 10 groups, got %d: %v", len(grouped), grouped)
	}
	if len(grouped[0]) != config.Config.ConsulMaxKVsPerTransaction {
		t.Fatalf("expected %d KVPairs in first group, got %d: %v", config.Config.ConsulMaxKVsPerTransaction, len(grouped[0]), grouped[0])
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
			if count != config.ConsulKVsPerCluster {
				t.Fatalf("expected group to contain %d x %s keys, found: %d", config.ConsulKVsPerCluster, cluster, count)
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
		}); err == nil || err.Error() != "test error" {
			t.Fatalf("Expected %q error from .PutKVPairs(), got: %q", "test error", err.Error())
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
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test"},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test2"},
				},
			},
			Response: &consulapi.TxnResponse{
				Results: consulapi.TxnResults{
					{
						KV: &consulapi.KVPair{Key: "test", Value: []byte("test")},
					},
					{
						KV: &consulapi.KVPair{Key: "test2", Value: []byte("not-equal")},
					},
				},
			},
		},
		{
			Method: "PUT",
			URL:    "/v1/txn?dc=dc1",
			Request: consulapi.TxnOps{
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test"},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "doesnt-exist"},
				},
			},
			Response: &consulapi.TxnResponse{
				Errors: consulapi.TxnErrors{
					{
						OpIndex: 1,
						What:    `key "doesnt-exist" doesn't exist`,
					},
				},
			},
			ResponseCode: http.StatusConflict, // PUT /v1/txn returns a HTTP 409 code on txn failure
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
		{
			Method: "PUT",
			URL:    "/v1/txn?dc=dc1",
			Request: consulapi.TxnOps{
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVSet, Key: "test", Value: []byte("test")},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVSet, Key: "doesnt-exist", Value: []byte("test")},
				},
			},
			Response: &consulapi.TxnResponse{
				Results: consulapi.TxnResults{
					{
						KV: &consulapi.KVPair{Key: "test", Value: []byte("test")},
					},
					{
						KV: &consulapi.KVPair{Key: "doesnt-exist", Value: []byte("test")},
					},
				},
			},
		},
	})
	defer server.Close()

	var wg sync.WaitGroup
	config.Config.ConsulAddress = server.URL
	store := NewConsulTxnStore().(*consulTxnStore)
	respChan := make(chan updateDatacenterKVPairsResponse)

	t.Run("success-cached", func(t *testing.T) {
		wg.Add(1)
		cacheKey := getConsulKVCacheKey(consulTestDefaultDatacenter, "cached")
		store.kvCache.SetDefault(cacheKey, "cached") // pre-cache the 'cached' key-value
		defer store.kvCache.Flush()

		kvPairs := []*consulapi.KVPair{
			{Key: "cached", Value: []byte("cached")}, // already cached key-value
			{Key: "test", Value: []byte("test")},     // already correct on consul server
			{Key: "test2", Value: []byte("test")},    // not equal on consul server
		}

		go store.updateDatacenterKVPairs(&wg, consulTestDefaultDatacenter, kvPairs, respChan)
		resp := <-respChan
		if resp.err != nil {
			t.Fatalf(".updateDatacenterKVPairs() should not return an error, got: %v", resp.err)
		}
		if resp.skipped != 1 || resp.existing != 1 || resp.written != 1 || resp.failed != 0 {
			t.Fatalf("expected: existing/skipped/written=1 and failed=0, got: skipped=%d, existing=%d, written=%d, failed=%d",
				resp.skipped, resp.existing, resp.written, resp.failed,
			)
		}

		for _, pair := range kvPairs {
			cacheKey := getConsulKVCacheKey(consulTestDefaultDatacenter, pair.Key)
			if cached, found := store.kvCache.Get(cacheKey); !found || cached != string(pair.Value) {
				t.Fatalf("expected cache key %q to equal %q, got %v", cacheKey, string(pair.Value), cached)
			}
		}
	})

	t.Run("success-missing-kv", func(t *testing.T) {
		wg.Add(1)
		kvPairs := []*consulapi.KVPair{
			{Key: "test", Value: []byte("test")},         // already correct on consul server
			{Key: "doesnt-exist", Value: []byte("test")}, // does not exist on consul server
		}
		go store.updateDatacenterKVPairs(&wg, consulTestDefaultDatacenter, kvPairs, respChan)
		resp := <-respChan

		if resp.err != nil {
			t.Fatalf(".updateDatacenterKVPairs() should not return an error, got: %v", resp.err)
		}
		if resp.skipped != 0 || resp.existing != 0 || resp.written != 2 || resp.failed != 0 { // confirm all KVs are updated if one does not exist
			t.Fatalf("expected: existing/skipped/failed=0 and written=2, got: skipped=%d, existing=%d, written=%d, failed=%d",
				resp.skipped, resp.existing, resp.written, resp.failed,
			)
		}
	})
}

func TestConsulTxnStoreDistributePairs(t *testing.T) {
	server := buildConsulTestServer(t, []consulTestServerOp{
		{
			Method:   "GET",
			URL:      "/v1/catalog/datacenters",
			Response: []string{"dc1"},
		},
		{
			Method: "PUT",
			URL:    "/v1/txn?dc=dc1",
			Request: consulapi.TxnOps{
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test/cluster1"},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test/cluster1/hostname"},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test/cluster1/ipv4"},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test/cluster1/ipv6"},
				},
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVGet, Key: "test/cluster1/port"},
				},
			},
			Response: &consulapi.TxnResponse{
				Results: consulapi.TxnResults{
					{
						KV: &consulapi.KVPair{Key: "test/cluster1", Value: []byte("not-equal")},
					},
					{
						KV: &consulapi.KVPair{Key: "test/cluster1/hostname", Value: []byte("mysql.example.com")},
					},
					{
						KV: &consulapi.KVPair{Key: "test/cluster1/ipv4", Value: []byte("10.20.30.40")},
					},
					{
						KV: &consulapi.KVPair{Key: "test/cluster1/ipv6", Value: []byte("fdf0:7a53:0b88:d147:xxxx:xxxx:xxxx:xxxx")},
					},
					{
						KV: &consulapi.KVPair{Key: "test/cluster1/port", Value: []byte("3306")},
					},
				},
			},
		},
		{
			Method: "PUT",
			URL:    "/v1/txn?dc=dc1",
			Request: consulapi.TxnOps{
				{
					KV: &consulapi.KVTxnOp{Verb: consulapi.KVSet, Key: "test/cluster1", Value: []byte("mysql.example.com:3306")},
				},
			},
			Response: &consulapi.TxnResponse{
				Results: consulapi.TxnResults{
					{
						KV: &consulapi.KVPair{Key: "test/cluster1", Value: []byte("mysql.example.com:3306")},
					},
				},
			},
		},
	})
	defer server.Close()
	config.Config.ConsulAddress = server.URL
	config.Config.ConsulCrossDataCenterDistribution = true
	config.Config.KVClusterMasterPrefix = "test"

	store := NewConsulTxnStore()
	if err := store.DistributePairs([]*KVPair{
		{Key: "test/cluster1", Value: "mysql.example.com:3306"},
		{Key: "test/cluster1/hostname", Value: "mysql.example.com"},
		{Key: "test/cluster1/ipv4", Value: "10.20.30.40"},
		{Key: "test/cluster1/ipv6", Value: "fdf0:7a53:0b88:d147:xxxx:xxxx:xxxx:xxxx"},
		{Key: "test/cluster1/port", Value: "3306"},
	}); err != nil {
		t.Fatal(err)
	}
}
