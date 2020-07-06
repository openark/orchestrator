package http

import (
	"net/http"
	"strconv"

	"github.com/openark/orchestrator/go/inst"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type instanceCollectorector struct {
}

var (
	cl = &instanceCollectorector{}

	descUptime = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "uptime_seconds"),
		"instance uptime seconds",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descReadOnly = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "read_only"),
		"instance is read only",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descDetachedMaster = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "detached_master"),
		"instance is detached master",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descReplicationLag = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "replication_lag_seconds"),
		"instance replication lag",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descSecondsBehindMaster = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "behind_master_seconds"),
		"instance seconds behind master",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descSecondsSinceLastSeen = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "since_last_seen_seconds"),
		"instance since last seen",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descLastDiscoveryLatency = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "last_discovery_latency_seconds"),
		"instance last discovery latency",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descReplicationIOThreadState = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "replication_io_thread_state"),
		"instance replication io thread state",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descReplicationSQLThreadState = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "replication_sql_thread_state"),
		"instance replication sql thread state",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descIsMaster = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "is_master"),
		"instance master status",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
	descGtidErrant = prometheus.NewDesc(
		prometheus.BuildFQName("mysql_orchestrator", "instance", "gtid_errant_found"),
		"instance gtid errant transactions",
		[]string{"mysql_instance", "cluster_alias"},
		nil,
	)
)

func RegisterMetrics() http.Handler {
	prometheus.MustRegister(cl)
	// disable compression, the martini will do it for us if requested by client
	return promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{DisableCompression: true})
}

func (e *instanceCollectorector) Describe(ch chan<- *prometheus.Desc) {
	ch <- descReplicationLag
}

func (e *instanceCollectorector) Collect(ch chan<- prometheus.Metric) {
	instances, err := inst.ReadAllInstances()
	if err != nil {
		return
	}

	for _, i := range instances {
		ch <- prometheus.MustNewConstMetric(descUptime, prometheus.GaugeValue, float64(i.Uptime), i.Key.Hostname, i.SuggestedClusterAlias)
		ch <- prometheus.MustNewConstMetric(descReadOnly, prometheus.GaugeValue, btof64(i.ReadOnly), i.Key.Hostname, i.SuggestedClusterAlias)
		ch <- prometheus.MustNewConstMetric(descDetachedMaster, prometheus.GaugeValue, btof64(i.IsDetachedMaster), i.Key.Hostname, i.SuggestedClusterAlias)
		{
			var replicationLag = 0.0
			if i.SlaveLagSeconds.Valid {
				replicationLag = float64(i.SlaveLagSeconds.Int64)
			}
			ch <- prometheus.MustNewConstMetric(descReplicationLag, prometheus.GaugeValue, replicationLag, i.Key.Hostname, i.SuggestedClusterAlias)
		}
		{
			var behind = 0.0
			if i.SecondsBehindMaster.Valid {
				behind = float64(i.SecondsBehindMaster.Int64)
			}
			ch <- prometheus.MustNewConstMetric(descSecondsBehindMaster, prometheus.GaugeValue, behind, i.Key.Hostname, i.SuggestedClusterAlias)
		}
		ch <- prometheus.MustNewConstMetric(descLastDiscoveryLatency, prometheus.GaugeValue, i.LastDiscoveryLatency.Seconds(), i.Key.Hostname, i.SuggestedClusterAlias)
		{
			var seen = 0.0
			if i.SecondsSinceLastSeen.Valid {
				seen = float64(i.SecondsSinceLastSeen.Int64)
			}
			ch <- prometheus.MustNewConstMetric(descSecondsSinceLastSeen, prometheus.GaugeValue, seen, i.Key.Hostname, i.SuggestedClusterAlias)
		}
		// NOTE: other possibility would be to encode state in label and generate four metrics
		ch <- prometheus.MustNewConstMetric(descReplicationIOThreadState, prometheus.GaugeValue, float64(int(i.ReplicationIOThreadState)), i.Key.Hostname, i.SuggestedClusterAlias)
		ch <- prometheus.MustNewConstMetric(descReplicationSQLThreadState, prometheus.GaugeValue, float64(int(i.ReplicationSQLThreadState)), i.Key.Hostname, i.SuggestedClusterAlias)
		{
			var master = (i.Key.Hostname + ":" + strconv.Itoa(i.Key.Port)) == i.ClusterName
			ch <- prometheus.MustNewConstMetric(descIsMaster, prometheus.GaugeValue, btof64(master), i.Key.Hostname, i.SuggestedClusterAlias)
		}
		ch <- prometheus.MustNewConstMetric(descGtidErrant, prometheus.GaugeValue, btof64(i.GtidErrant != ""), i.Key.Hostname, i.SuggestedClusterAlias)
	}
}

func btof64(v bool) float64 {
	if v {
		return 1.0
	} else {
		return 0.0
	}
}
