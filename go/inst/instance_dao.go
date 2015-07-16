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

package inst

import (
	"database/sql"
	"fmt"
	"github.com/outbrain/golib/log"
	"github.com/outbrain/golib/math"
	"github.com/outbrain/golib/sqlutils"
	"github.com/outbrain/orchestrator/go/config"
	"github.com/outbrain/orchestrator/go/db"
	"regexp"
	"sort"
	"strings"
	"time"
)

const sqlThreadPollDuration = 200 * time.Millisecond

const backendDBConcurrency = 20

var instanceReadChan = make(chan bool, backendDBConcurrency)
var instanceWriteChan = make(chan bool, backendDBConcurrency)

var detachPattern *regexp.Regexp

// Max concurrency for bulk topology operations
const topologyConcurrency = 100

var topologyConcurrencyChan = make(chan bool, topologyConcurrency)

// InstancesByCountSlaveHosts is a sortable type for Instance
type InstancesByCountSlaveHosts [](*Instance)

func (this InstancesByCountSlaveHosts) Len() int      { return len(this) }
func (this InstancesByCountSlaveHosts) Swap(i, j int) { this[i], this[j] = this[j], this[i] }
func (this InstancesByCountSlaveHosts) Less(i, j int) bool {
	return len(this[i].SlaveHosts) < len(this[j].SlaveHosts)
}

func init() {
	detachPattern, _ = regexp.Compile(`//([^/:]+):([\d]+)`) // e.g. `//binlog.01234:567890`
}

// ExecuteOnTopology will execute given function while maintaining concurrency limit
// on topology servers. It is safe in the sense that we will not leak tokens.
func ExecuteOnTopology(f func()) {
	topologyConcurrencyChan <- true
	defer func() { <-topologyConcurrencyChan }()
	f()
}

// ExecDBWriteFunc chooses how to execute a write onto the database: whether synchronuously or not
func ExecDBWriteFunc(f func() error) error {
	instanceWriteChan <- true
	defer func() { <-instanceWriteChan }()
	res := f()
	return res
}

// ExecInstance executes a given query on the given MySQL topology instance
func ExecInstance(instanceKey *InstanceKey, query string, args ...interface{}) (sql.Result, error) {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.Exec(db, query, args...)
	return res, err
}

// ExecInstanceNoPrepare executes a given query on the given MySQL topology instance, without using prepared statements
func ExecInstanceNoPrepare(instanceKey *InstanceKey, query string, args ...interface{}) (sql.Result, error) {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.ExecNoPrepare(db, query, args...)
	return res, err
}

// ScanInstanceRow executes a read-a-single-row query on a given MySQL topology instance
func ScanInstanceRow(instanceKey *InstanceKey, query string, dest ...interface{}) error {
	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return err
	}
	err = db.QueryRow(query).Scan(dest...)
	return err
}

// ReadTopologyInstance connects to a topology MySQL instance and reads its configuration and
// replication status. It writes read info into orchestrator's backend.
func ReadTopologyInstance(instanceKey *InstanceKey) (*Instance, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Unexpected error: %+v", err)
		}
	}()

	instance := NewInstance()
	instanceFound := false
	foundByShowSlaveHosts := false
	longRunningProcesses := []Process{}
	resolvedHostname := ""
	isMaxScale := false
	slaveStatusFound := false
	var resolveErr error

	// Before we even begin anything, declare that we're about to cehck this instance.If anything goes wrong network-wise,
	// this is our source of truth in terms of instance being inaccessible
	_ = UpdateInstanceLastAttemptedCheck(instanceKey)

	db, err := db.OpenTopology(instanceKey.Hostname, instanceKey.Port)

	if err != nil {
		goto Cleanup
	}

	instance.Key = *instanceKey

	{
		// Is this MaxScale? (a proxy, not a real server)
		err = sqlutils.QueryRowsMap(db, "show variables like 'maxscale%'", func(m sqlutils.RowMap) error {
			variableName := m.GetString("Variable_name")
			if variableName == "MAXSCALE_VERSION" {
				instance.Version = m.GetString("value") + "-maxscale"
				instance.ServerID = 0
				instance.Uptime = 0
				instance.Binlog_format = "INHERIT"
				instance.ReadOnly = true
				instance.LogBinEnabled = true
				instance.LogSlaveUpdatesEnabled = true
				resolvedHostname = instance.Key.Hostname
				UpdateResolvedHostname(resolvedHostname, resolvedHostname)
				isMaxScale = true
			}
			return nil
		})
		if err != nil {
			// The query should not error (even if it's not maxscale)
			goto Cleanup
		}
	}

	if !isMaxScale {
		var mysqlHostname, mysqlReportHost string
		err = db.QueryRow("select @@global.hostname, ifnull(@@global.report_host, ''), @@global.server_id, @@global.version, @@global.read_only, @@global.binlog_format, @@global.log_bin, @@global.log_slave_updates").Scan(
			&mysqlHostname, &mysqlReportHost, &instance.ServerID, &instance.Version, &instance.ReadOnly, &instance.Binlog_format, &instance.LogBinEnabled, &instance.LogSlaveUpdatesEnabled)
		if err != nil {
			goto Cleanup
		}
		switch strings.ToLower(config.Config.MySQLHostnameResolveMethod) {
		case "none":
			resolvedHostname = instance.Key.Hostname
		case "default", "hostname", "@@hostname":
			resolvedHostname = mysqlHostname
		case "report_host", "@@report_host":
			if mysqlReportHost == "" {
				err = fmt.Errorf("MySQLHostnameResolveMethod configured to use @@report_host but %+v has NULL/empty @@report_host", instanceKey)
				goto Cleanup
			}
			resolvedHostname = mysqlReportHost
		default:
			resolvedHostname = instance.Key.Hostname
		}

		err = db.QueryRow("select variable_value from information_schema.global_status where variable_name='Uptime'").Scan(&instance.Uptime)
		if err != nil {
			goto Cleanup
		}
		// @@gtid_mode only available in Orcale MySQL >= 5.6
		_ = db.QueryRow("select @@global.gtid_mode = 'ON'").Scan(&instance.SupportsOracleGTID)
	}
	if resolvedHostname != instance.Key.Hostname {
		UpdateResolvedHostname(instance.Key.Hostname, resolvedHostname)
		instance.Key.Hostname = resolvedHostname
	}
	if config.Config.DataCenterPattern != "" {
		if pattern, err := regexp.Compile(config.Config.DataCenterPattern); err == nil {
			match := pattern.FindStringSubmatch(instance.Key.Hostname)
			if len(match) != 0 {
				instance.DataCenter = match[1]
			}
		}
	}
	if config.Config.PhysicalEnvironmentPattern != "" {
		if pattern, err := regexp.Compile(config.Config.PhysicalEnvironmentPattern); err == nil {
			match := pattern.FindStringSubmatch(instance.Key.Hostname)
			if len(match) != 0 {
				instance.PhysicalEnvironment = match[1]
			}
		}
	}

	err = sqlutils.QueryRowsMap(db, "show slave status", func(m sqlutils.RowMap) error {
		instance.Slave_IO_Running = (m.GetString("Slave_IO_Running") == "Yes")
		instance.Slave_SQL_Running = (m.GetString("Slave_SQL_Running") == "Yes")
		instance.ReadBinlogCoordinates.LogFile = m.GetString("Master_Log_File")
		instance.ReadBinlogCoordinates.LogPos = m.GetInt64("Read_Master_Log_Pos")
		instance.ExecBinlogCoordinates.LogFile = m.GetString("Relay_Master_Log_File")
		instance.ExecBinlogCoordinates.LogPos = m.GetInt64("Exec_Master_Log_Pos")
		instance.RelaylogCoordinates.LogFile = m.GetString("Relay_Log_File")
		instance.RelaylogCoordinates.LogPos = m.GetInt64("Relay_Log_Pos")
		instance.RelaylogCoordinates.Type = RelayLog
		instance.LastSQLError = m.GetString("Last_SQL_Error")
		instance.LastIOError = m.GetString("Last_IO_Error")
		instance.SQLDelay = m.GetUintD("SQL_Delay", 0)
		instance.UsingOracleGTID = (m.GetIntD("Auto_Position", 0) == 1)
		instance.UsingMariaDBGTID = (m.GetStringD("Using_Gtid", "No") != "No")
		instance.HasReplicationFilters = ((m.GetStringD("Replicate_Do_DB", "") != "") || (m.GetStringD("Replicate_Ignore_DB", "") != "") || (m.GetStringD("Replicate_Do_Table", "") != "") || (m.GetStringD("Replicate_Ignore_Table", "") != "") || (m.GetStringD("Replicate_Wild_Do_Table", "") != "") || (m.GetStringD("Replicate_Wild_Ignore_Table", "") != ""))

		masterKey, err := NewInstanceKeyFromStrings(m.GetString("Master_Host"), m.GetString("Master_Port"))
		if err != nil {
			log.Errore(err)
		}
		masterKey.Hostname, resolveErr = ResolveHostname(masterKey.Hostname)
		if resolveErr != nil {
			log.Errore(resolveErr)
		}
		instance.MasterKey = *masterKey
		instance.SecondsBehindMaster = m.GetNullInt64("Seconds_Behind_Master")
		// And until told otherwise:
		instance.SlaveLagSeconds = instance.SecondsBehindMaster
		// Not breaking the flow even on error
		slaveStatusFound = true
		return nil
	})
	if err != nil {
		goto Cleanup
	}
	if isMaxScale && !slaveStatusFound {
		err = fmt.Errorf("No 'SHOW SLAVE STATUS' output found for a MaxScale instance: %+v", instanceKey)
		goto Cleanup
	}

	instance.UsingPseudoGTID = false
	if config.Config.DetectPseudoGTIDQuery != "" && !isMaxScale {
		resultData, err := sqlutils.QueryResultData(db, config.Config.DetectPseudoGTIDQuery)
		if err == nil {
			if len(resultData) > 0 {
				if len(resultData[0]) > 0 {
					if resultData[0][0].Valid && resultData[0][0].String == "1" {
						instance.UsingPseudoGTID = true
					}
				}
			}
		}
	}

	if instance.LogBinEnabled {
		err = sqlutils.QueryRowsMap(db, "show master status", func(m sqlutils.RowMap) error {
			var err error
			instance.SelfBinlogCoordinates.LogFile = m.GetString("File")
			instance.SelfBinlogCoordinates.LogPos = m.GetInt64("Position")
			return err
		})
		if err != nil {
			goto Cleanup
		}
	}

	instanceFound = true

	// -------------------------------------------------------------------------
	// Anything after this point does not affect the fact the instance is found.
	// No `goto Cleanup` after this point.
	// -------------------------------------------------------------------------

	// Get slaves, either by SHOW SLAVE HOSTS or via PROCESSLIST
	// MaxScale does not support PROCESSLIST, so SHOW SLAVE HOSTS is the only option
	if config.Config.DiscoverByShowSlaveHosts || isMaxScale {
		err = sqlutils.QueryRowsMap(db, `show slave hosts`,
			func(m sqlutils.RowMap) error {
				slaveKey, err := NewInstanceKeyFromStrings(m.GetString("Host"), m.GetString("Port"))
				slaveKey.Hostname, resolveErr = ResolveHostname(slaveKey.Hostname)
				if err == nil {
					instance.AddSlaveKey(slaveKey)
					foundByShowSlaveHosts = true
				}
				return err
			})

		if err != nil {
			log.Errore(err)
		}
	}
	if !foundByShowSlaveHosts && !isMaxScale {
		// Either not configured to read SHOW SLAVE HOSTS or nothing was there.
		// Discover by processlist
		err = sqlutils.QueryRowsMap(db, `
        	select 
        		substring_index(host, ':', 1) as slave_hostname 
        	from 
        		information_schema.processlist 
        	where 
        		command='Binlog Dump'`,
			func(m sqlutils.RowMap) error {
				cname, resolveErr := ResolveHostname(m.GetString("slave_hostname"))
				if resolveErr != nil {
					log.Errore(resolveErr)
				}
				slaveKey := InstanceKey{Hostname: cname, Port: instance.Key.Port}
				instance.AddSlaveKey(&slaveKey)
				return err
			})

		if err != nil {
			log.Errore(err)
		}
	}

	if config.Config.ReadLongRunningQueries && !isMaxScale {
		// Get long running processes
		err := sqlutils.QueryRowsMap(db, `
				  select 
				    id,
				    user,
				    host,
				    db,
				    command,
				    time,
				    state,
				    left(processlist.info, 1024) as info,
				    now() - interval time second as started_at
				  from 
				    information_schema.processlist 
				  where
				    time > 60
				    and command != 'Sleep'
				    and id != connection_id()
				    and user != 'system user' 
				    and command != 'Binlog dump'
				    and user != 'event_scheduler'
				  order by
				    time desc
        		`,
			func(m sqlutils.RowMap) error {
				process := Process{}
				process.Id = m.GetInt64("id")
				process.User = m.GetString("user")
				process.Host = m.GetString("host")
				process.Db = m.GetString("db")
				process.Command = m.GetString("command")
				process.Time = m.GetInt64("time")
				process.State = m.GetString("state")
				process.Info = m.GetString("info")
				process.StartedAt = m.GetString("started_at")

				longRunningProcesses = append(longRunningProcesses, process)
				return nil
			})

		if err != nil {
			log.Errore(err)
		}
	}

	if config.Config.SlaveLagQuery != "" && !isMaxScale {
		err := db.QueryRow(config.Config.SlaveLagQuery).Scan(&instance.SlaveLagSeconds)
		if err != nil {
			instance.SlaveLagSeconds = instance.SecondsBehindMaster
			log.Errore(err)
		}
	}

	instance.ClusterName, instance.ReplicationDepth, instance.IsCoMaster, err = ReadClusterNameByMaster(&instance.Key, &instance.MasterKey)
	if err != nil {
		log.Errore(err)
	}
	if instance.ReplicationDepth == 0 && config.Config.DetectClusterAliasQuery != "" && !isMaxScale {
		// Only need to do on masters
		clusterAlias := ""
		err := db.QueryRow(config.Config.DetectClusterAliasQuery).Scan(&clusterAlias)
		if err != nil {
			clusterAlias = ""
			log.Errore(err)
		}
		if clusterAlias != "" {
			err := SetClusterAlias(instance.ClusterName, clusterAlias)
			if err != nil {
				log.Errore(err)
			}
		}
	}

Cleanup:
	if instanceFound {
		_ = writeInstance(instance, instanceFound, err)
		WriteLongRunningProcesses(&instance.Key, longRunningProcesses)
	} else {
		_ = UpdateInstanceLastChecked(&instance.Key)
	}
	if err != nil {
		log.Errore(err)
	}
	return instance, err
}

// ReadClusterNameByMaster will return the cluster name for a given instance by looking at its master
// and getting it from there.
// It is a non-recursive function and so-called-recursion is performed upon periodic reading of
// instances.
func ReadClusterNameByMaster(instanceKey *InstanceKey, masterKey *InstanceKey) (clusterName string, replicationDepth uint, isCoMaster bool, err error) {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return "", 0, false, log.Errore(err)
	}

	clusterNameByInstanceKey := fmt.Sprintf("%s:%d", instanceKey.Hostname, instanceKey.Port)
	var masterMasterKey InstanceKey
	err = db.QueryRow(`
       	select 
       		if (
       			max(cluster_name) != '',
       			max(cluster_name),
	       		ifnull(concat(max(hostname), ':', max(port)), '')
	       	) as cluster_name,
	       	ifnull(max(replication_depth)+1, 0) as replication_depth,
	       	ifnull(max(master_host), '') as master_master_host,
	       	ifnull(max(master_port), 0) as master_master_port
       	from database_instance 
		 	where hostname=? and port=?
		 	group by hostname, port`,
		masterKey.Hostname, masterKey.Port).Scan(&clusterName, &replicationDepth, &masterMasterKey.Hostname, &masterMasterKey.Port)

	if err != nil && err != sql.ErrNoRows {
		return "", 0, false, log.Errore(err)
	}
	if clusterName == "" {
		clusterName = clusterNameByInstanceKey
	}
	isCoMaster = false
	if masterMasterKey.Equals(instanceKey) {
		isCoMaster = true
		if clusterName == clusterNameByInstanceKey {
			// circular replication. Avoid infinite ++ on replicationDepth
			replicationDepth = 0
		} // While the other stays "1"
	}
	return clusterName, replicationDepth, isCoMaster, nil
}

// readInstanceRow reads a single instance row from the orchestrator backend database.
func readInstanceRow(m sqlutils.RowMap) *Instance {
	instance := NewInstance()

	instance.Key.Hostname = m.GetString("hostname")
	instance.Key.Port = m.GetInt("port")
	instance.Uptime = m.GetUint("uptime")
	instance.ServerID = m.GetUint("server_id")
	instance.Version = m.GetString("version")
	instance.ReadOnly = m.GetBool("read_only")
	instance.Binlog_format = m.GetString("binlog_format")
	instance.LogBinEnabled = m.GetBool("log_bin")
	instance.LogSlaveUpdatesEnabled = m.GetBool("log_slave_updates")
	instance.MasterKey.Hostname = m.GetString("master_host")
	instance.MasterKey.Port = m.GetInt("master_port")
	instance.Slave_SQL_Running = m.GetBool("slave_sql_running")
	instance.Slave_IO_Running = m.GetBool("slave_io_running")
	instance.HasReplicationFilters = m.GetBool("has_replication_filters")
	instance.UsingOracleGTID = m.GetBool("oracle_gtid")
	instance.UsingMariaDBGTID = m.GetBool("mariadb_gtid")
	instance.UsingPseudoGTID = m.GetBool("pseudo_gtid")
	instance.SelfBinlogCoordinates.LogFile = m.GetString("binary_log_file")
	instance.SelfBinlogCoordinates.LogPos = m.GetInt64("binary_log_pos")
	instance.ReadBinlogCoordinates.LogFile = m.GetString("master_log_file")
	instance.ReadBinlogCoordinates.LogPos = m.GetInt64("read_master_log_pos")
	instance.ExecBinlogCoordinates.LogFile = m.GetString("relay_master_log_file")
	instance.ExecBinlogCoordinates.LogPos = m.GetInt64("exec_master_log_pos")
	instance.RelaylogCoordinates.LogFile = m.GetString("relay_log_file")
	instance.RelaylogCoordinates.LogPos = m.GetInt64("relay_log_pos")
	instance.RelaylogCoordinates.Type = RelayLog
	instance.LastSQLError = m.GetString("last_sql_error")
	instance.LastIOError = m.GetString("last_io_error")
	instance.SecondsBehindMaster = m.GetNullInt64("seconds_behind_master")
	instance.SlaveLagSeconds = m.GetNullInt64("slave_lag_seconds")
	instance.SQLDelay = m.GetUint("sql_delay")
	slaveHostsJSON := m.GetString("slave_hosts")
	instance.ClusterName = m.GetString("cluster_name")
	instance.DataCenter = m.GetString("data_center")
	instance.PhysicalEnvironment = m.GetString("physical_environment")
	instance.ReplicationDepth = m.GetUint("replication_depth")
	instance.IsCoMaster = m.GetBool("is_co_master")
	instance.IsUpToDate = (m.GetUint("seconds_since_last_checked") <= config.Config.InstancePollSeconds)
	instance.IsRecentlyChecked = (m.GetUint("seconds_since_last_checked") <= config.Config.InstancePollSeconds*5)
	instance.IsLastCheckValid = m.GetBool("is_last_check_valid")
	instance.SecondsSinceLastSeen = m.GetNullInt64("seconds_since_last_seen")
	instance.IsCandidate = m.GetBool("is_candidate")
	instance.UnresolvedHostname = m.GetString("unresolved_hostname")

	instance.ReadSlaveHostsFromJson(slaveHostsJSON)
	return instance
}

// readInstancesByCondition is a generic function to read instances from the backend database
func readInstancesByCondition(condition string, sort string) ([](*Instance), error) {
	readFunc := func() ([](*Instance), error) {
		instances := [](*Instance){}

		db, err := db.OpenOrchestrator()
		if err != nil {
			return instances, log.Errore(err)
		}

		if sort == "" {
			sort = `hostname, port`
		}
		query := fmt.Sprintf(`
		select 
			*,
			timestampdiff(second, last_checked, now()) as seconds_since_last_checked,
			(last_checked <= last_seen) is true as is_last_check_valid,
			timestampdiff(second, last_seen, now()) as seconds_since_last_seen,
			candidate_database_instance.last_suggested is not null as is_candidate,
			ifnull(unresolved_hostname, '') as unresolved_hostname 
		from 
			database_instance 
			left join candidate_database_instance using (hostname, port)
			left join hostname_unresolve using (hostname)
		where
			%s
		order by
			%s
			`, condition, sort)

		err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
			instance := readInstanceRow(m)
			instances = append(instances, instance)
			return nil
		})
		if err != nil {
			return instances, log.Errore(err)
		}
		err = PopulateInstancesAgents(instances)
		if err != nil {
			return instances, log.Errore(err)
		}
		return instances, err
	}
	instanceReadChan <- true
	instances, err := readFunc()
	<-instanceReadChan
	return instances, err
}

// ReadInstance reads an instance from the orchestrator backend database
func ReadInstance(instanceKey *InstanceKey) (*Instance, bool, error) {
	condition := fmt.Sprintf(`
			hostname = '%s'
			and port = %d
		`, instanceKey.Hostname, instanceKey.Port)
	instances, err := readInstancesByCondition(condition, "")
	// We know there will be at most one (hostname & port are PK)
	// And we expect to find one
	if len(instances) == 0 {
		return nil, false, err
	}
	if err != nil {
		return instances[0], false, err
	}
	return instances[0], true, nil
}

// ReadClusterInstances reads all instances of a given cluster
func ReadClusterInstances(clusterName string) ([](*Instance), error) {
	if strings.Index(clusterName, "'") >= 0 {
		return [](*Instance){}, log.Errorf("Invalid cluster name: %s", clusterName)
	}
	condition := fmt.Sprintf(`cluster_name = '%s'`, clusterName)
	return readInstancesByCondition(condition, "")
}

// ReadSlaveInstances reads slaves of a given master
func ReadSlaveInstances(masterKey *InstanceKey) ([](*Instance), error) {
	condition := fmt.Sprintf(`
			master_host = '%s'
			and master_port = %d
		`, masterKey.Hostname, masterKey.Port)
	return readInstancesByCondition(condition, "")
}

// ReadUnseenInstances reads all instances which were not recently seen
func ReadUnseenInstances() ([](*Instance), error) {
	condition := fmt.Sprintf(`
			last_seen < last_checked
		`)
	return readInstancesByCondition(condition, "")
}

// ReadProblemInstances reads all instances with problems
func ReadProblemInstances() ([](*Instance), error) {
	condition := fmt.Sprintf(`
			(last_seen < last_checked)
			or (not ifnull(timestampdiff(second, last_checked, now()) <= %d, false))
			or (not slave_sql_running)
			or (not slave_io_running)
			or (abs(cast(seconds_behind_master as signed) - cast(sql_delay as signed)) > %d)
			or (abs(cast(slave_lag_seconds as signed) - cast(sql_delay as signed)) > %d)
		`, config.Config.InstancePollSeconds, config.Config.ReasonableReplicationLagSeconds, config.Config.ReasonableReplicationLagSeconds)
	return readInstancesByCondition(condition, "")
}

// SearchInstances reads all instances qualifying for some searchString
func SearchInstances(searchString string) ([](*Instance), error) {
	searchString = strings.TrimSpace(searchString)
	condition := fmt.Sprintf(`
			locate('%s', hostname) > 0
			or locate('%s', cluster_name) > 0
			or locate('%s', version) > 0
			or locate('%s', concat(hostname, ':', port)) > 0
			or concat(server_id, '') = '%s'
			or concat(port, '') = '%s'
		`, searchString, searchString, searchString, searchString, searchString, searchString)
	return readInstancesByCondition(condition, `replication_depth asc, num_slave_hosts desc, cluster_name, hostname, port`)
}

// FindInstances reads all instances whose name matches given pattern
func FindInstances(regexpPattern string) ([](*Instance), error) {
	condition := fmt.Sprintf(`
			hostname rlike '%s'
		`, regexpPattern)
	return readInstancesByCondition(condition, `replication_depth asc, num_slave_hosts desc, cluster_name, hostname, port`)
}

// filterOSCInstances will filter the given list such that only slaves fit for OSC control remain.
func filterOSCInstances(instances [](*Instance)) [](*Instance) {
	result := [](*Instance){}
	for _, instance := range instances {
		skipThisHost := false
		for _, filter := range config.Config.OSCIgnoreHostnameFilters {
			if matched, _ := regexp.MatchString(filter, instance.Key.Hostname); matched {
				skipThisHost = true
			}
		}
		if instance.IsMaxScale() {
			skipThisHost = true
		}

		if !instance.IsLastCheckValid {
			skipThisHost = true
		}
		if !skipThisHost {
			result = append(result, instance)
		}
	}
	return result
}

// GetClusterOSCSlaves returns a heuristic list of slaves which are fit as controll slaves for an OSC operation.
// These would be intermediate masters
func GetClusterOSCSlaves(clusterName string) ([](*Instance), error) {
	intermediateMasters := [](*Instance){}
	result := [](*Instance){}
	var err error
	if strings.Index(clusterName, "'") >= 0 {
		return [](*Instance){}, log.Errorf("Invalid cluster name: %s", clusterName)
	}
	{
		// Pick up to two busiest IMs
		condition := fmt.Sprintf(`
			replication_depth = 1
			and num_slave_hosts > 0
			and cluster_name = '%s'
		`, clusterName)
		intermediateMasters, err = readInstancesByCondition(condition, "")
		if err != nil {
			return result, err
		}
		sort.Sort(sort.Reverse(InstancesByCountSlaveHosts(intermediateMasters)))
		intermediateMasters = filterOSCInstances(intermediateMasters)
		intermediateMasters = intermediateMasters[0:math.MinInt(2, len(intermediateMasters))]
		result = append(result, intermediateMasters...)
	}
	{
		// Get 2 slaves of found IMs, if possible
		if len(intermediateMasters) == 1 {
			// Pick 2 slaves for this IM
			slaves, err := ReadSlaveInstances(&(intermediateMasters[0].Key))
			if err != nil {
				return result, err
			}
			sort.Sort(sort.Reverse(InstancesByCountSlaveHosts(slaves)))
			slaves = filterOSCInstances(slaves)
			slaves = slaves[0:math.MinInt(2, len(slaves))]
			result = append(result, slaves...)

		}
		if len(intermediateMasters) == 2 {
			// Pick one slave from each IM (should be possible)
			for _, im := range intermediateMasters {
				slaves, err := ReadSlaveInstances(&im.Key)
				if err != nil {
					return result, err
				}
				sort.Sort(sort.Reverse(InstancesByCountSlaveHosts(slaves)))
				slaves = filterOSCInstances(slaves)
				if len(slaves) > 0 {
					result = append(result, slaves[0])
				}
			}
		}
	}
	{
		// Get 2 3rd tier slaves, if possible
		condition := fmt.Sprintf(`
			replication_depth = 3
			and cluster_name = '%s'
		`, clusterName)
		slaves, err := readInstancesByCondition(condition, "")
		if err != nil {
			return result, err
		}
		sort.Sort(sort.Reverse(InstancesByCountSlaveHosts(slaves)))
		slaves = filterOSCInstances(slaves)
		slaves = slaves[0:math.MinInt(2, len(slaves))]
		result = append(result, slaves...)
	}
	{
		// Get 2 1st tier leaf slaves, if possible
		condition := fmt.Sprintf(`
			replication_depth = 1
			and num_slave_hosts = 0
			and cluster_name = '%s'
		`, clusterName)
		slaves, err := readInstancesByCondition(condition, "")
		if err != nil {
			return result, err
		}
		slaves = filterOSCInstances(slaves)
		slaves = slaves[0:math.MinInt(2, len(slaves))]
		result = append(result, slaves...)
	}

	return result, nil
}

// GetClusterHeuristicLag returns a heuristic lag for a cluster, based on its OSC slaves
func GetClusterHeuristicLag(clusterName string) (int64, error) {
	instances, err := GetClusterOSCSlaves(clusterName)
	if err != nil {
		return 0, err
	}
	var maxLag int64
	for _, clusterInstance := range instances {
		if clusterInstance.SlaveLagSeconds.Valid && clusterInstance.SlaveLagSeconds.Int64 > maxLag {
			maxLag = clusterInstance.SlaveLagSeconds.Int64
		}
	}
	return maxLag, nil

}

// updateInstanceClusterName
func updateInstanceClusterName(instance *Instance) error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}
		_, err = sqlutils.Exec(db, `
			update 
				database_instance 
			set 
				cluster_name=? 
			where 
				hostname=? and port=?
        	`, instance.ClusterName, instance.Key.Hostname, instance.Key.Port,
		)
		if err != nil {
			return log.Errore(err)
		}
		AuditOperation("update-cluster-name", &instance.Key, fmt.Sprintf("set to %s", instance.ClusterName))
		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// ReviewUnseenInstances reviews instances that have not been seen (suposedly dead) and updates some of their data
func ReviewUnseenInstances() error {
	instances, err := ReadUnseenInstances()
	if err != nil {
		return log.Errore(err)
	}
	operations := 0
	for _, instance := range instances {
		instance := instance

		masterHostname, err := ResolveHostname(instance.MasterKey.Hostname)
		if err != nil {
			log.Errore(err)
			continue
		}
		instance.MasterKey.Hostname = masterHostname
		clusterName, replicationDepth, isCoMaster, err := ReadClusterNameByMaster(&instance.Key, &instance.MasterKey)

		if err != nil {
			log.Errore(err)
		} else if clusterName != instance.ClusterName {
			instance.ClusterName = clusterName
			instance.ReplicationDepth = replicationDepth
			instance.IsCoMaster = isCoMaster
			updateInstanceClusterName(instance)
			operations++
		}
	}

	AuditOperation("review-unseen-instances", nil, fmt.Sprintf("Operations: %d", operations))
	return err
}

// readUnseenMasterKeys will read list of masters that have never been seen, and yet whose slaves
// seem to be replicating.
func readUnseenMasterKeys() ([]InstanceKey, error) {
	res := []InstanceKey{}
	db, err := db.OpenOrchestrator()
	if err != nil {
		return res, log.Errore(err)
	}
	err = sqlutils.QueryRowsMap(db, `
			SELECT DISTINCT
			    slave_instance.master_host, slave_instance.master_port
			FROM
			    database_instance slave_instance
			        LEFT JOIN
			    hostname_resolve ON (slave_instance.master_host = hostname_resolve.hostname)
			        LEFT JOIN
			    database_instance master_instance ON (
			    	COALESCE(hostname_resolve.resolved_hostname, slave_instance.master_host) = master_instance.hostname
			    	and slave_instance.master_port = master_instance.port)
			WHERE
			    master_instance.last_checked IS NULL
			    and slave_instance.master_host != ''
			    and slave_instance.master_host != '_'
			    and slave_instance.master_port > 0
			    and slave_instance.slave_io_running = 1
			`, func(m sqlutils.RowMap) error {
		instanceKey, _ := NewInstanceKeyFromStrings(m.GetString("master_host"), m.GetString("master_port"))
		// we ignore the error. It can be expected that we are unable to resolve the hostname.
		// Maybe that's how we got here in the first place!
		res = append(res, *instanceKey)

		return nil
	})
	if err != nil {
		return res, log.Errore(err)
	}

	return res, nil
}

// InjectUnseenMasters will review masters of instances that are known to be replicating, yet which are not listed
// in database_instance. Since their slaves are listed as replicating, we can assume that such masters actually do
// exist: we shall therefore inject them with minimal details into the database_instance table.
func InjectUnseenMasters() error {

	unseenMasterKeys, err := readUnseenMasterKeys()
	if err != nil {
		return err
	}

	operations := 0
	for _, masterKey := range unseenMasterKeys {
		masterKey := masterKey
		clusterName := fmt.Sprintf("%s:%d", masterKey.Hostname, masterKey.Port)
		// minimal details:
		instance := Instance{Key: masterKey, Version: "Unknown", ClusterName: clusterName}
		if err := writeInstance(&instance, false, nil); err == nil {
			operations++
		}
	}

	AuditOperation("inject-unseen-masters", nil, fmt.Sprintf("Operations: %d", operations))
	return err
}

// ForgetUnseenInstancesDifferentlyResolved will purge instances which are invalid, and whose hostname
// appears on the hostname_resolved table; this means some time in the past their hostname was unresovled, and now
// resovled to a different value; the old hostname is never accessed anymore and the old entry should be removed.
func ForgetUnseenInstancesDifferentlyResolved() error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	sqlResult, err := sqlutils.Exec(db, `
		DELETE FROM 
			database_instance
		USING
		    hostname_resolve
		    JOIN database_instance ON (hostname_resolve.hostname = database_instance.hostname)
		WHERE
		    hostname_resolve.hostname != hostname_resolve.resolved_hostname
		    AND (last_checked <= last_seen) IS NOT TRUE
		`,
	)
	if err != nil {
		return log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	if err != nil {
		return log.Errore(err)
	}
	AuditOperation("forget-unseen-differently-resolved", nil, fmt.Sprintf("Forgotten instances: %d", rows))
	return err
}

// readUnknownMasterHostnameResolves will figure out the resolved hostnames of master-hosts which cannot be found.
// It uses the hostname_resolve_history table to heuristically guess the correct hostname (based on "this was the
// last time we saw this hostname and it resolves into THAT")
func readUnknownMasterHostnameResolves() (map[string]string, error) {
	res := make(map[string]string)
	db, err := db.OpenOrchestrator()
	if err != nil {
		return res, log.Errore(err)
	}
	err = sqlutils.QueryRowsMap(db, `
			SELECT DISTINCT
			    slave_instance.master_host, hostname_resolve_history.resolved_hostname
			FROM
			    database_instance slave_instance
			LEFT JOIN hostname_resolve ON (slave_instance.master_host = hostname_resolve.hostname)
			LEFT JOIN database_instance master_instance ON (
			    COALESCE(hostname_resolve.resolved_hostname, slave_instance.master_host) = master_instance.hostname
			    and slave_instance.master_port = master_instance.port
			) LEFT JOIN hostname_resolve_history ON (slave_instance.master_host = hostname_resolve_history.hostname)
			WHERE
			    master_instance.last_checked IS NULL
			    and slave_instance.master_host != ''
			    and slave_instance.master_host != '_'
			    and slave_instance.master_port > 0
			`, func(m sqlutils.RowMap) error {
		res[m.GetString("master_host")] = m.GetString("resolved_hostname")
		return nil
	})
	if err != nil {
		return res, log.Errore(err)
	}

	return res, nil
}

// ResolveUnknownMasterHostnameResolves fixes missing hostname resolves based on hostname_resolve_history
// The use case is slaves replicating from some unknown-hostname which cannot be otherwise found. This could
// happen due to an expire unresolve together with clearing up of hostname cache.
func ResolveUnknownMasterHostnameResolves() error {

	hostnameResolves, err := readUnknownMasterHostnameResolves()
	if err != nil {
		return err
	}
	for hostname, resolvedHostname := range hostnameResolves {
		UpdateResolvedHostname(hostname, resolvedHostname)
	}

	AuditOperation("resolve-unknown-masters", nil, fmt.Sprintf("Num resolved hostnames: %d", len(hostnameResolves)))
	return err
}

// ReadCountMySQLSnapshots is a utility method to return registered number of snapshots for a given list of hosts
func ReadCountMySQLSnapshots(hostnames []string) (map[string]int, error) {
	res := make(map[string]int)
	if !config.Config.ServeAgentsHttp {
		return res, nil
	}
	query := fmt.Sprintf(`
		select 
			hostname,
			count_mysql_snapshots
		from 
			host_agent
		where
			hostname in (%s)
		order by
			hostname
		`, sqlutils.InClauseStringValues(hostnames))
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		res[m.GetString("hostname")] = m.GetInt("count_mysql_snapshots")
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err
}

// PopulateInstancesAgents will fill in extra data acquired from agents for given instances
// At current this is the number of snapshots.
// This isn't too pretty; it's a push-into-instance-data-that-belongs-to-agent thing.
// Originally the need was to visually present the number of snapshots per host on the web/cluster page, which
// indeed proves to be useful in our experience.
func PopulateInstancesAgents(instances [](*Instance)) error {
	if len(instances) == 0 {
		return nil
	}
	hostnames := []string{}
	for _, instance := range instances {
		hostnames = append(hostnames, instance.Key.Hostname)
	}
	agentsCountMySQLSnapshots, err := ReadCountMySQLSnapshots(hostnames)
	if err != nil {
		return err
	}
	for _, instance := range instances {
		if count, ok := agentsCountMySQLSnapshots[instance.Key.Hostname]; ok {
			instance.CountMySQLSnapshots = count
		}
	}

	return nil
}

// ReadClusters reads names of all known clusters
func ReadClusters() ([]string, error) {
	clusterNames := []string{}

	db, err := db.OpenOrchestrator()
	if err != nil {
		return clusterNames, log.Errore(err)
	}

	query := fmt.Sprintf(`
		select 
			cluster_name
		from 
			database_instance 
		group by
			cluster_name`)

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		clusterNames = append(clusterNames, m.GetString("cluster_name"))
		return nil
	})

	return clusterNames, err
}

// ReadClusterInfo reads some info about a given cluster
func ReadClusterInfo(clusterName string) (*ClusterInfo, error) {
	clusterInfo := &ClusterInfo{}

	db, err := db.OpenOrchestrator()
	if err != nil {
		return clusterInfo, log.Errore(err)
	}

	query := fmt.Sprintf(`
		select 
			cluster_name,
			count(*) as count_instances
		from 
			database_instance 
		where
			cluster_name='%s'
		group by
			cluster_name`, clusterName)

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		clusterInfo.ClusterName = m.GetString("cluster_name")
		clusterInfo.CountInstances = m.GetUint("count_instances")
		ApplyClusterAlias(clusterInfo)
		clusterInfo.ReadRecoveryInfo()
		return nil
	})
	if err != nil {
		return clusterInfo, err
	}
	clusterInfo.HeuristicLag, err = GetClusterHeuristicLag(clusterName)

	return clusterInfo, err
}

// ReadClustersInfo reads names of all known clusters and some aggregated info
func ReadClustersInfo() ([]ClusterInfo, error) {
	clusters := []ClusterInfo{}

	db, err := db.OpenOrchestrator()
	if err != nil {
		return clusters, log.Errore(err)
	}

	query := fmt.Sprintf(`
		select 
			cluster_name,
			count(*) as count_instances
		from 
			database_instance 
		group by
			cluster_name`)

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		clusterInfo := ClusterInfo{
			ClusterName:    m.GetString("cluster_name"),
			CountInstances: m.GetUint("count_instances"),
		}
		ApplyClusterAlias(&clusterInfo)
		clusterInfo.ReadRecoveryInfo()

		clusters = append(clusters, clusterInfo)
		return nil
	})

	return clusters, err
}

// ReadOutdatedInstanceKeys reads and returns keys for all instances that are not up to date (i.e.
// pre-configured time has passed since they were last cheked)
// But we also check for the case where an attempt at instance checking has been made, that hasn't
// resulted in an actual check! This can happen when TCP/IP connections are hung, in which case the "check"
// never returns. In such case we multiply interval by a factor, so as not to open too many connections on
// the instance.
func ReadOutdatedInstanceKeys() ([]InstanceKey, error) {
	res := []InstanceKey{}
	query := fmt.Sprintf(`
		select 
			hostname, port 
		from 
			database_instance 
		where
			if (
				last_attempted_check <= last_checked,
				last_checked < now() - interval %d second,
				last_checked < now() - interval (%d * 2) second
			)
			`,
		config.Config.InstancePollSeconds, config.Config.InstancePollSeconds)
	db, err := db.OpenOrchestrator()
	if err != nil {
		goto Cleanup
	}

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		instanceKey, merr := NewInstanceKeyFromStrings(m.GetString("hostname"), m.GetString("port"))
		if merr != nil {
			log.Errore(merr)
		} else {
			res = append(res, *instanceKey)
		}
		// We don;t return an error because we want to keep filling the outdated instances list.
		return nil
	})
Cleanup:

	if err != nil {
		log.Errore(err)
	}
	return res, err

}

// writeInstance stores an instance in the orchestrator backend
func writeInstance(instance *Instance, instanceWasActuallyFound bool, lastError error) error {

	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		insertIgnore := ""
		onDuplicateKeyUpdate := ""
		if instanceWasActuallyFound {
			// Data is real
			onDuplicateKeyUpdate = `
				on duplicate key update
	        		last_checked=VALUES(last_checked),
	        		last_attempted_check=VALUES(last_attempted_check),
	        		uptime=VALUES(uptime),
	        		server_id=VALUES(server_id),
					version=VALUES(version),
					read_only=VALUES(read_only),
					binlog_format=VALUES(binlog_format),
					log_bin=VALUES(log_bin),
					log_slave_updates=VALUES(log_slave_updates),
					binary_log_file=VALUES(binary_log_file),
					binary_log_pos=VALUES(binary_log_pos),
					master_host=VALUES(master_host),
					master_port=VALUES(master_port),
					slave_sql_running=VALUES(slave_sql_running),
					slave_io_running=VALUES(slave_io_running),
					has_replication_filters=VALUES(has_replication_filters),
					oracle_gtid=VALUES(oracle_gtid),
					mariadb_gtid=VALUES(mariadb_gtid),
					pseudo_gtid=values(pseudo_gtid),
					master_log_file=VALUES(master_log_file),
					read_master_log_pos=VALUES(read_master_log_pos),
					relay_master_log_file=VALUES(relay_master_log_file),
					exec_master_log_pos=VALUES(exec_master_log_pos),
					relay_log_file=VALUES(relay_log_file),
					relay_log_pos=VALUES(relay_log_pos),
					last_sql_error=VALUES(last_sql_error),
					last_io_error=VALUES(last_io_error),
					seconds_behind_master=VALUES(seconds_behind_master),
					slave_lag_seconds=VALUES(slave_lag_seconds),
					sql_delay=VALUES(sql_delay),
					num_slave_hosts=VALUES(num_slave_hosts),
					slave_hosts=VALUES(slave_hosts),
					cluster_name=VALUES(cluster_name),
					data_center=VALUES(data_center),
					physical_environment=values(physical_environment),
					replication_depth=VALUES(replication_depth),
					is_co_master=VALUES(is_co_master)
				`
		} else {
			// Scenario: some slave reported a master of his; but the master cannot be contacted.
			// We might still want to have that master written down
			// Use with caution
			insertIgnore = `ignore`
		}
		insertQuery := fmt.Sprintf(`
        	insert %s into database_instance (
        		hostname,
        		port,
        		last_checked,
        		last_attempted_check,
        		uptime,
        		server_id,
				version,
				read_only,
				binlog_format,
				log_bin,
				log_slave_updates,
				binary_log_file,
				binary_log_pos,
				master_host,
				master_port,
				slave_sql_running,
				slave_io_running,
				has_replication_filters,
				oracle_gtid,
				mariadb_gtid,
				pseudo_gtid,
				master_log_file,
				read_master_log_pos,
				relay_master_log_file,
				exec_master_log_pos,
				relay_log_file,
				relay_log_pos,
				last_sql_error,
				last_io_error,
				seconds_behind_master,
				slave_lag_seconds,
				sql_delay,
				num_slave_hosts,
				slave_hosts,
				cluster_name,
				data_center,
				physical_environment,
				replication_depth,
				is_co_master
			) values (?, ?, NOW(), NOW(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
			%s
			`, insertIgnore, onDuplicateKeyUpdate)

		_, err = sqlutils.Exec(db, insertQuery,
			instance.Key.Hostname,
			instance.Key.Port,
			instance.Uptime,
			instance.ServerID,
			instance.Version,
			instance.ReadOnly,
			instance.Binlog_format,
			instance.LogBinEnabled,
			instance.LogSlaveUpdatesEnabled,
			instance.SelfBinlogCoordinates.LogFile,
			instance.SelfBinlogCoordinates.LogPos,
			instance.MasterKey.Hostname,
			instance.MasterKey.Port,
			instance.Slave_SQL_Running,
			instance.Slave_IO_Running,
			instance.HasReplicationFilters,
			instance.UsingOracleGTID,
			instance.UsingMariaDBGTID,
			instance.UsingPseudoGTID,
			instance.ReadBinlogCoordinates.LogFile,
			instance.ReadBinlogCoordinates.LogPos,
			instance.ExecBinlogCoordinates.LogFile,
			instance.ExecBinlogCoordinates.LogPos,
			instance.RelaylogCoordinates.LogFile,
			instance.RelaylogCoordinates.LogPos,
			instance.LastSQLError,
			instance.LastIOError,
			instance.SecondsBehindMaster,
			instance.SlaveLagSeconds,
			instance.SQLDelay,
			len(instance.SlaveHosts),
			instance.GetSlaveHostsAsJson(),
			instance.ClusterName,
			instance.DataCenter,
			instance.PhysicalEnvironment,
			instance.ReplicationDepth,
			instance.IsCoMaster,
		)
		if err != nil {
			return log.Errore(err)
		}

		if instanceWasActuallyFound && lastError == nil {
			sqlutils.Exec(db, `
        	update database_instance set last_seen = NOW() where hostname=? and port=?
        	`, instance.Key.Hostname, instance.Key.Port,
			)
		} else {
			log.Debugf("writeInstance: will not update database_instance due to error: %+v", lastError)
		}
		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// UpdateInstanceLastChecked updates the last_check timestamp in the orchestrator backed database
// for a given instance
func UpdateInstanceLastChecked(instanceKey *InstanceKey) error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		_, err = sqlutils.Exec(db, `
        	update 
        		database_instance 
        	set
        		last_checked = NOW()
			where 
				hostname = ?
				and port = ?`,
			instanceKey.Hostname,
			instanceKey.Port,
		)
		if err != nil {
			return log.Errore(err)
		}

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// UpdateInstanceLastAttemptedCheck updates the last_attempted_check timestamp in the orchestrator backed database
// for a given instance.
// This is used as a failsafe mechanism in case access to the instance gets hung (it happens), in which case
// the entire ReadTopology gets stuck (and no, connection timeout nor driver timeouts don't help. Don't look at me,
// the world is a harsh place to live in).
// And so we make sure to note down *before* we even attempt to access the instance; and this raises a red flag when we
// wish to access the instance again: if last_attempted_check is *newer* than last_checked, that's bad news and means
// we have a "hanging" issue.
func UpdateInstanceLastAttemptedCheck(instanceKey *InstanceKey) error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		_, err = sqlutils.Exec(db, `
        	update 
        		database_instance 
        	set
        		last_attempted_check = NOW()
			where 
				hostname = ?
				and port = ?`,
			instanceKey.Hostname,
			instanceKey.Port,
		)
		if err != nil {
			return log.Errore(err)
		}

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// ForgetInstance removes an instance entry from the orchestrator backed database.
// It may be auto-rediscovered through topology or requested for discovery by multiple means.
func ForgetInstance(instanceKey *InstanceKey) error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	_, err = sqlutils.Exec(db, `
			delete 
				from database_instance 
			where 
				hostname = ? and port = ?`,
		instanceKey.Hostname,
		instanceKey.Port,
	)
	AuditOperation("forget", instanceKey, "")
	return err
}

// ForgetLongUnseenInstances will remove entries of all instacnes that have long since been last seen.
func ForgetLongUnseenInstances() error {
	db, err := db.OpenOrchestrator()
	if err != nil {
		return log.Errore(err)
	}

	sqlResult, err := sqlutils.Exec(db, `
			delete 
				from database_instance 
			where 
				last_seen < NOW() - interval ? hour`,
		config.Config.UnseenInstanceForgetHours,
	)
	if err != nil {
		return log.Errore(err)
	}
	rows, err := sqlResult.RowsAffected()
	if err != nil {
		return log.Errore(err)
	}
	AuditOperation("forget-unseen", nil, fmt.Sprintf("Forgotten instances: %d", rows))
	return err
}

// RefreshTopologyInstance will synchronuously re-read topology instance
func RefreshTopologyInstance(instanceKey *InstanceKey) (*Instance, error) {
	_, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return nil, err
	}

	inst, found, err := ReadInstance(instanceKey)
	if err != nil || !found {
		return nil, err
	}

	return inst, nil
}

// RefreshTopologyInstances will do a blocking (though concurrent) refresh of all given instances
func RefreshTopologyInstances(instances [](*Instance)) {
	// use concurrency but wait for all to complete
	barrier := make(chan InstanceKey)
	for _, instance := range instances {
		instance := instance
		go func() {
			// Signal completed slave
			defer func() { barrier <- instance.Key }()
			// Wait your turn to read a slave
			ExecuteOnTopology(func() {
				log.Debugf("... reading instance: %+v", instance.Key)
				ReadTopologyInstance(&instance.Key)
			})
		}()
	}
	for _ = range instances {
		<-barrier
	}
}

// RefreshInstanceSlaveHosts is a workaround for a bug in MySQL where
// SHOW SLAVE HOSTS continues to present old, long disconnected slaves.
// It turns out issuing a couple FLUSH commands mitigates the problem.
func RefreshInstanceSlaveHosts(instanceKey *InstanceKey) (*Instance, error) {
	_, _ = ExecInstance(instanceKey, `flush error logs`)
	_, _ = ExecInstance(instanceKey, `flush error logs`)

	instance, err := ReadTopologyInstance(instanceKey)
	return instance, err
}

// FlushBinaryLogs attempts a 'FLUSH BINARY LOGS' statement on the given instance.
func FlushBinaryLogs(instanceKey *InstanceKey) error {

	_, err := ExecInstance(instanceKey, `flush binary logs`)
	if err != nil {
		return log.Errore(err)
	}

	log.Infof("flush-binary-logs on %+v", *instanceKey)
	AuditOperation("flush-binary-logs", instanceKey, "success")

	return nil
}

// StopSlaveNicely stops a slave such that SQL_thread and IO_thread are aligned (i.e.
// SQL_thread consumes all relay log entries)
// It will actually START the sql_thread even if the slave is completely stopped.
func StopSlaveNicely(instanceKey *InstanceKey, timeout time.Duration) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}

	_, err = ExecInstanceNoPrepare(instanceKey, `stop slave io_thread`)
	_, err = ExecInstanceNoPrepare(instanceKey, `start slave sql_thread`)

	if instance.SQLDelay == 0 {
		// Otherwise we don't bother.
		startTime := time.Now()
		for upToDate := false; !upToDate; {
			if timeout > 0 && time.Since(startTime) >= timeout {
				// timeout
				return nil, log.Errorf("StopSlaveNicely timeout on %+v", *instanceKey)
			}
			instance, err = ReadTopologyInstance(instanceKey)
			if err != nil {
				return instance, log.Errore(err)
			}

			if instance.SQLThreadUpToDate() {
				upToDate = true
			} else {
				time.Sleep(sqlThreadPollDuration)
			}
		}
	}
	_, err = ExecInstanceNoPrepare(instanceKey, `stop slave`)
	if err != nil {
		return instance, log.Errore(err)
	}

	instance, err = ReadTopologyInstance(instanceKey)
	log.Infof("Stopped slave nicely on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)
	return instance, err
}

// StopSlavesNicely will attemt to stop all given slaves nicely, up to timeout
func StopSlavesNicely(slaves [](*Instance), timeout time.Duration) [](*Instance) {
	refreshedSlaves := [](*Instance){}

	log.Debugf("Stopping %d slaves nicely", len(slaves))
	// use concurrency but wait for all to complete
	barrier := make(chan *Instance)
	for _, slave := range slaves {
		slave := slave
		go func() {
			updatedSlave := &slave
			// Signal completed slave
			defer func() { barrier <- *updatedSlave }()
			// Wait your turn to read a slave
			ExecuteOnTopology(func() {
				StopSlaveNicely(&slave.Key, timeout)
				slave, _ = StopSlave(&slave.Key)
				updatedSlave = &slave
			})
		}()
	}
	for _ = range slaves {
		refreshedSlaves = append(refreshedSlaves, <-barrier)
	}
	return refreshedSlaves
}

// StopSlave stops replication on a given instance
func StopSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}
	_, err = ExecInstanceNoPrepare(instanceKey, `stop slave`)
	if err != nil {
		return instance, log.Errore(err)
	}
	instance, err = ReadTopologyInstance(instanceKey)

	log.Infof("Stopped slave on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)
	return instance, err
}

// StartSlave starts replication on a given instance
func StartSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}

	_, err = ExecInstanceNoPrepare(instanceKey, `start slave`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Started slave on %+v", instanceKey)
	if config.Config.SlaveStartPostWaitMilliseconds > 0 {
		time.Sleep(time.Duration(config.Config.SlaveStartPostWaitMilliseconds) * time.Millisecond)
	}

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// StartSlaves will do concurrent start-slave
func StartSlaves(slaves [](*Instance)) {
	// use concurrency but wait for all to complete
	log.Debugf("Starting %d slaves", len(slaves))
	barrier := make(chan InstanceKey)
	for _, instance := range slaves {
		instance := instance
		go func() {
			// Signal compelted slave
			defer func() { barrier <- instance.Key }()
			// Wait your turn to read a slave
			ExecuteOnTopology(func() { StartSlave(&instance.Key) })
		}()
	}
	for _ = range slaves {
		<-barrier
	}
}

// StartSlaveUntilMasterCoordinates issuesa START SLAVE UNTIL... statement on given instance
func StartSlaveUntilMasterCoordinates(instanceKey *InstanceKey, masterCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}
	if instance.SlaveRunning() {
		return instance, fmt.Errorf("slave already running: %+v", instanceKey)
	}

	log.Infof("Will start slave on %+v until coordinates: %+v", instanceKey, masterCoordinates)

	// MariaDB has a bug: a CHANGE MASTER TO statement does not work properly with prepared statement... :P
	// See https://mariadb.atlassian.net/browse/MDEV-7640
	// This is the reason for ExecInstanceNoPrepare
	_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("start slave until master_log_file='%s', master_log_pos=%d",
		masterCoordinates.LogFile, masterCoordinates.LogPos))
	if err != nil {
		return instance, log.Errore(err)
	}

	for upToDate := false; !upToDate; {
		instance, err = ReadTopologyInstance(instanceKey)
		if err != nil {
			return instance, log.Errore(err)
		}

		switch {
		case instance.ExecBinlogCoordinates.SmallerThan(masterCoordinates):
			time.Sleep(sqlThreadPollDuration)
		case instance.ExecBinlogCoordinates.Equals(masterCoordinates):
			upToDate = true
		case masterCoordinates.SmallerThan(&instance.ExecBinlogCoordinates):
			return instance, fmt.Errorf("Start SLAVE UNTIL is past coordinates: %+v", instanceKey)
		}
	}

	instance, err = StopSlave(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	return instance, err
}

// ChangeMasterTo changes the given instance's master according to given input.
func ChangeMasterTo(instanceKey *InstanceKey, masterKey *InstanceKey, masterBinlogCoordinates *BinlogCoordinates, skipUnresolve bool) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("ChangeMasterTo: Cannot change master on: %+v because slave is running", *instanceKey)
	}
	log.Debugf("ChangeMasterTo: will attempt changing master on %+v to %+v, %+v", *instanceKey, *masterKey, *masterBinlogCoordinates)
	changeToMasterKey := masterKey
	if !skipUnresolve {
		unresolvedMasterKey, nameUnresolved, err := UnresolveHostname(masterKey)
		if err != nil {
			log.Debugf("ChangeMasterTo: aborting operation on %+v due to resolving error on %+v: %+v", *instanceKey, *masterKey, err)
			return instance, err
		}
		if nameUnresolved {
			log.Debugf("ChangeMasterTo: Unresolved %+v into %+v", *masterKey, unresolvedMasterKey)
		}
		changeToMasterKey = &unresolvedMasterKey
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting CHANGE MASTER TO operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	if instance.UsingMariaDBGTID {
		// MariaDB has a bug: a CHANGE MASTER TO statement does not work properly with prepared statement... :P
		// See https://mariadb.atlassian.net/browse/MDEV-7640
		// This is the reason for ExecInstanceNoPrepare
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d",
			changeToMasterKey.Hostname, changeToMasterKey.Port))
	} else if instance.UsingOracleGTID {
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_auto_position=1",
			changeToMasterKey.Hostname, changeToMasterKey.Port))
	} else {
		_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_log_file='%s', master_log_pos=%d",
			changeToMasterKey.Hostname, changeToMasterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos))
	}
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("ChangeMasterTo: Changed master on %+v to: %+v, %+v", *instanceKey, changeToMasterKey, masterBinlogCoordinates)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// ResetSlave resets a slave, breaking the replication
func ResetSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("Cannot reset slave on: %+v because slave is running", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reset-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	// MySQL's RESET SLAVE is done correctly; however SHOW SLAVE STATUS still returns old hostnames etc
	// and only resets till after next restart. This leads to orchestrator still thinking the instance replicates
	// from old host. We therefore forcibly modify the hostname.
	// RESET SLAVE ALL command solves this, but only as of 5.6.3
	_, err = ExecInstanceNoPrepare(instanceKey, `change master to master_host='_'`)
	if err != nil {
		return instance, log.Errore(err)
	}
	_, err = ExecInstanceNoPrepare(instanceKey, `reset slave /*!50603 all */`)
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Reset slave %+v", instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// SkipQuery skip a single query in a failed replication instance
func SkipQuery(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if !instance.IsSlave() {
		return instance, fmt.Errorf("instance is not a slave: %+v", instanceKey)
	}
	if instance.Slave_SQL_Running {
		return instance, fmt.Errorf("Slave_SQL_is running on %+v", instanceKey)
	}
	if instance.LastSQLError == "" {
		return instance, fmt.Errorf("No SQL error on %+v", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting skip-query operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	log.Debugf("Skipping one query on %+v", instanceKey)
	_, err = ExecInstance(instanceKey, `set global sql_slave_skip_counter := 1`)
	if err != nil {
		return instance, log.Errore(err)
	}
	return StartSlave(instanceKey)
}

// DetachSlave detaches a slave from replication; forcibly corrupting the binlog coordinates (though in such way
// that is reversible)
func DetachSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("Cannot detach slave on: %+v because slave is running", instanceKey)
	}

	detachedCoordinatesSubmatch := detachPattern.FindStringSubmatch(instance.ExecBinlogCoordinates.LogFile)
	if len(detachedCoordinatesSubmatch) != 0 {
		return instance, fmt.Errorf("Cannot (need not) detach slave on: %+v because slave is already detached", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting detach-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	detachedCoordinates := BinlogCoordinates{LogFile: fmt.Sprintf("//%s:%d", instance.ExecBinlogCoordinates.LogFile, instance.ExecBinlogCoordinates.LogPos), LogPos: instance.ExecBinlogCoordinates.LogPos}
	// Encode the current coordinates within the log file name, in such way that replication is broken, but info can still be resurrected
	_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf(`change master to master_log_file='%s', master_log_pos=%d`, detachedCoordinates.LogFile, detachedCoordinates.LogPos))
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Detach slave %+v", instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// ReattachSlave restores a detahced slave back into replication
func ReattachSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if instance.SlaveRunning() {
		return instance, fmt.Errorf("Cannot (need not) reattach slave on: %+v because slave is running", instanceKey)
	}

	detachedCoordinatesSubmatch := detachPattern.FindStringSubmatch(instance.ExecBinlogCoordinates.LogFile)
	if len(detachedCoordinatesSubmatch) == 0 {
		return instance, fmt.Errorf("Cannot reattach slave on: %+v because slave is not detached", instanceKey)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting reattach-slave operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstanceNoPrepare(instanceKey, fmt.Sprintf(`change master to master_log_file='%s', master_log_pos=%s`, detachedCoordinatesSubmatch[1], detachedCoordinatesSubmatch[2]))
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Reattach slave %+v", instanceKey)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// MasterPosWait issues a MASTER_POS_WAIT() an given instance according to given coordinates.
func MasterPosWait(instanceKey *InstanceKey, binlogCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	_, err = ExecInstance(instanceKey, fmt.Sprintf("select master_pos_wait('%s', %d)",
		binlogCoordinates.LogFile, binlogCoordinates.LogPos))
	if err != nil {
		return instance, log.Errore(err)
	}
	log.Infof("Instance %+v has reached coordinates: %+v", instanceKey, binlogCoordinates)

	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}

// SetReadOnly sets or clears the instance's global read_only variable
func SetReadOnly(instanceKey *InstanceKey, readOnly bool) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting set-read-only operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstance(instanceKey, fmt.Sprintf("set global read_only = %t", readOnly))
	if err != nil {
		return instance, log.Errore(err)
	}
	instance, err = ReadTopologyInstance(instanceKey)

	log.Infof("instance %+v read_only: %t", instanceKey, readOnly)
	AuditOperation("read-only", instanceKey, fmt.Sprintf("set as %t", readOnly))

	return instance, err
}

// KillQuery stops replication on a given instance
func KillQuery(instanceKey *InstanceKey, process int64) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	if *config.RuntimeCLIFlags.Noop {
		return instance, fmt.Errorf("noop: aborting kill-query operation on %+v; signalling error but nothing went wrong.", *instanceKey)
	}

	_, err = ExecInstance(instanceKey, fmt.Sprintf(`kill query %d`, process))
	if err != nil {
		return instance, log.Errore(err)
	}

	instance, err = ReadTopologyInstance(instanceKey)
	if err != nil {
		return instance, log.Errore(err)
	}

	log.Infof("Killed query on %+v", *instanceKey)
	AuditOperation("kill-query", instanceKey, fmt.Sprintf("Killed query %d", process))
	return instance, err
}

// SnapshotTopologies records topology graph for all existing topologies
func SnapshotTopologies() error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		_, err = sqlutils.Exec(db, `
        	insert into 
        		database_instance_topology_history (snapshot_unix_timestamp,
        			hostname, port, master_host, master_port, cluster_name)
        	select
        		UNIX_TIMESTAMP(NOW()),
        		hostname, port, master_host, master_port, cluster_name
			from
				database_instance
				`,
		)
		if err != nil {
			return log.Errore(err)
		}

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// ReadHistoryClusterInstances reads (thin) instances from history
func ReadHistoryClusterInstances(clusterName string, historyTimestampPattern string) ([](*Instance), error) {
	instances := [](*Instance){}

	db, err := db.OpenOrchestrator()
	if err != nil {
		return instances, log.Errore(err)
	}

	query := fmt.Sprintf(`
		select 
			*
		from 
			database_instance_topology_history 
		where
			snapshot_unix_timestamp rlike '%s'
			and cluster_name ='%s' 
		order by
			hostname, port`, historyTimestampPattern, clusterName)

	err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		instance := NewInstance()

		instance.Key.Hostname = m.GetString("hostname")
		instance.Key.Port = m.GetInt("port")
		instance.MasterKey.Hostname = m.GetString("master_host")
		instance.MasterKey.Port = m.GetInt("master_port")
		instance.ClusterName = m.GetString("cluster_name")

		instances = append(instances, instance)
		return nil
	})
	if err != nil {
		return instances, log.Errore(err)
	}
	return instances, err
}

// RegisterCandidateInstance markes a given instance as suggested for successoring a master in the event of failover.
func RegisterCandidateInstance(instanceKey *InstanceKey) error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		_, err = sqlutils.Exec(db, `
        	insert into candidate_database_instance (
        		hostname,
        		port,
        		last_suggested)
        	values (?, ?, NOW())
        	on duplicate key update
        		hostname=values(hostname),
        		port=values(port),
        		last_suggested=now()
				`, instanceKey.Hostname, instanceKey.Port,
		)
		if err != nil {
			return log.Errore(err)
		}

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// RegisterCandidateInstance markes a given instance as suggested for successoring a master in the event of failover.
func ExpireCandidateInstances() error {
	writeFunc := func() error {
		db, err := db.OpenOrchestrator()
		if err != nil {
			return log.Errore(err)
		}

		_, err = sqlutils.Exec(db, `
        	delete from candidate_database_instance 
				where last_suggested < NOW() - INTERVAL ? MINUTE
				`, config.Config.CandidateInstanceExpireMinutes,
		)
		if err != nil {
			return log.Errore(err)
		}

		return nil
	}
	return ExecDBWriteFunc(writeFunc)
}

// ReadClusterCandidateInstances reads cluster instances which are also marked as candidates
func ReadClusterCandidateInstances(clusterName string) ([](*Instance), error) {
	condition := fmt.Sprintf(`
			cluster_name = '%s'
			and (hostname, port) in (select hostname, port from candidate_database_instance)
			`, clusterName)
	return readInstancesByCondition(condition, "")
}
