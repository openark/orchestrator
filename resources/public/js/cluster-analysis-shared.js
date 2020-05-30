var interestingAnalysis = {
	"DeadMaster" : true,
	"DeadMasterAndSlaves" : true,
	"DeadMasterAndSomeSlaves" : true,
	"DeadMasterWithoutSlaves" : true,
	"UnreachableMasterWithStaleSlaves": true,
	"UnreachableMasterWithLaggingReplicas": true,
	"UnreachableMaster" : true,
	"LockedSemiSyncMaster" : true,
	"AllMasterSlavesNotReplicating" : true,
	"AllMasterSlavesNotReplicatingOrDead" : true,
	"AllMasterSlavesStale" : true,
	"DeadCoMaster" : true,
	"DeadCoMasterAndSomeSlaves" : true,
	"DeadIntermediateMaster" : true,
	"DeadIntermediateMasterWithSingleSlaveFailingToConnect" : true,
	"DeadIntermediateMasterWithSingleSlave" : true,
	"DeadIntermediateMasterAndSomeSlaves" : true,
	"DeadIntermediateMasterAndSlaves" : true,
	"AllIntermediateMasterSlavesFailingToConnectOrDead" : true,
	"AllIntermediateMasterSlavesNotReplicating" : true,
	"UnreachableIntermediateMasterWithLaggingReplicas": true,
	"UnreachableIntermediateMaster" : true,
	"BinlogServerFailingToConnectToMaster" : true,
};

var errorMapping = {
	"in_maintenance": {
		"badge": "label-info",
		"description": "In maintenance"
	},
	"last_check_invalid": {
		"badge": "label-fatal",
		"description": "Last check invalid"
	},
	"not_recently_checked": {
		"badge": "label-stale",
		"description": "Not recently checked (stale)"
	},
	"not_replicating": {
		"badge": "label-danger",
		"description": "Not replicating"
	},
	"replication_lag": {
		"badge": "label-warning",
		"description": "Replication lag"
	},
	"errant_gtid": {
		"badge": "label-errant",
		"description": "Errant GTID"
	}
};
