package inst

import (
	"fmt"
	"errors"
	"time"
	"strconv"
	"strings"
	"math"
	"database/sql"
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/config"
	"github.com/outbrain/log"
)

func ExecInstance(instanceKey *InstanceKey, query string, args ...interface{}) (sql.Result, error) {
	db,	err	:=	sqlutils.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return nil, err
	}
	res, err := sqlutils.Exec(db, query, args...)
	return res, err
}


func ScanInstanceRow(instanceKey *InstanceKey, query string, dest ...interface{}) error {
	db,	err	:=	sqlutils.OpenTopology(instanceKey.Hostname, instanceKey.Port)
	if err != nil {
		return err
	}
	err = db.QueryRow(query).Scan(dest...)
	return err
}

func ReadTopologyInstance(instanceKey *InstanceKey) (*Instance, error) {
	instance := NewInstance()
	instanceFound := false;
    foundBySlaveHosts := false

	db,	err	:=	sqlutils.OpenTopology(instanceKey.Hostname, instanceKey.Port)

    if err != nil {goto Cleanup}

   	instance.Key = *instanceKey
    err = db.QueryRow("select @@global.server_id, @@global.version, @@global.binlog_format, @@global.log_bin, @@global.log_slave_updates").Scan(
       	&instance.ServerID, &instance.Version, &instance.Binlog_format, &instance.LogBinEnabled, &instance.LogSlaveUpdatesEnabled)
    if err != nil {goto Cleanup}
    instanceFound = true
    err = sqlutils.QueryRowsMap(db, "show slave status", func(m sqlutils.RowMap) error {
		instance.Slave_IO_Running = (m["Slave_IO_Running"] == "Yes")
      	instance.Slave_SQL_Running = (m["Slave_SQL_Running"] == "Yes")
       	instance.ReadBinlogCoordinates.LogFile = m["Master_Log_File"]
       	instance.ReadBinlogCoordinates.LogPos, _ = strconv.ParseInt(m["Read_Master_Log_Pos"], 10, 0)
       	instance.ExecBinlogCoordinates.LogFile = m["Relay_Master_Log_File"]
       	instance.ExecBinlogCoordinates.LogPos, _ = strconv.ParseInt(m["Exec_Master_Log_Pos"], 10, 0)
       	var err error
       	instance.MasterKey.Hostname, err = GetCNAME(m["Master_Host"])
       	if err != nil {log.Errore(err)}
       	instance.MasterKey.Port, err = strconv.Atoi(m["Master_Port"])
       	if err != nil {return err}
       	if config.Config.SlaveLagQuery == "" {
       		instance.SecondsBehindMaster = m.GetIntD("Seconds_Behind_Master", math.MaxInt32)
        }
       	if err != nil {return err}
       	return err
   	})
    if err != nil {goto Cleanup}

	if config.Config.SlaveLagQuery != "" {
		err = db.QueryRow(config.Config.SlaveLagQuery).Scan(&instance.SecondsBehindMaster)
	    if err != nil {goto Cleanup}
	}
        
    err = sqlutils.QueryRowsMap(db, "show master status", func(m sqlutils.RowMap) error {
    	var err error
       	instance.SelfBinlogCoordinates.LogFile = m["File"]
       	instance.SelfBinlogCoordinates.LogPos, err = strconv.ParseInt(m["Position"], 10, 0)
       	return err
   	})
    if err != nil {goto Cleanup}
        
    // Get slaves, either by SHOW SLAVE HOSTS or via PROCESSLIST
    if config.Config.DiscoverByShowSlaveHosts {
        err = sqlutils.QueryRowsMap(db, `show slave hosts`, 
        		func(m sqlutils.RowMap) error {
        			cname, err := GetCNAME(m["Host"])
        			port, err := strconv.Atoi(m["Port"])
        			if err == nil {
	        			slaveKey := InstanceKey{Hostname: cname, Port: port}
						instance.AddSlaveKey(&slaveKey)
						foundBySlaveHosts = true
					}
					return err
		       	})
		
        if err != nil {goto Cleanup}
    } 
    if !foundBySlaveHosts {
       	// Discover by processlist
        err = sqlutils.QueryRowsMap(db, `
        	select 
        		substring_index(host, ':', 1) as slave_hostname 
        	from 
        		information_schema.processlist 
        	where 
        		command='Binlog Dump'`, 
        		func(m sqlutils.RowMap) error {
        			cname, err := GetCNAME(m["slave_hostname"])
        			if err != nil {return err}
        			slaveKey := InstanceKey{Hostname: cname, Port: instance.Key.Port}
					instance.AddSlaveKey(&slaveKey)
					return err
		       	})
			
        if err != nil {goto Cleanup}
	}
    if err != nil {goto Cleanup}

	instance.ClusterName, err = ReadClusterNameByMaster(&instance.Key, &instance.MasterKey)
    if err != nil {goto Cleanup}

	Cleanup:
	if instanceFound {
		_ = WriteInstance(instance, err)
	}
	if err	!=	nil	{
		log.Errore(err)
	}
	return instance, err
}


func ReadClusterNameByMaster(instanceKey *InstanceKey, masterKey *InstanceKey) (string, error) {
	db,	err	:=	sqlutils.OpenOrchestrator()
	if	err	!=	nil	{
		return "", log.Errore(err)
	}

	var clusterName string
    err = db.QueryRow(`
       	select 
       		if (
       			cluster_name != '',
       			cluster_name,
	       		ifnull(concat(max(hostname), ':', max(port)), '')
	       	) as cluster_name
       	from database_instance 
		 	where hostname=? and port=?`, 
		masterKey.Hostname, masterKey.Port).Scan(
		 	&clusterName,
		)
    if err != nil {return "", log.Errore(err)}
    if clusterName == "" {
    	return fmt.Sprintf("%s:%d", instanceKey.Hostname, instanceKey.Port), nil
    }
  	return clusterName, err
}


func ReadInstance(instanceKey *InstanceKey) (*Instance, bool, error) {
	db,	err	:=	sqlutils.OpenOrchestrator()
	if	err	!=	nil	{
		return nil, false, log.Errore(err)
	}
	instance := NewInstance()
	instance.Key = *instanceKey

	var slaveHostsJson string
    err = db.QueryRow(`
       	select 
       		server_id,
			version,
			binlog_format,
			log_bin, 
			log_slave_updates,
			binary_log_file,
			binary_log_pos,
			master_host,
			master_port,
			slave_sql_running,
			slave_io_running,
			master_log_file,
			read_master_log_pos,
			relay_master_log_file,
			exec_master_log_pos,
			seconds_behind_master,
			slave_hosts,
			cluster_name,
			ifnull(timestampdiff(second, last_seen, now()) <= ?, false) as is_up_to_date,
			is_last_seen_valid
		 from database_instance 
		 	where hostname=? and port=?`, 
		 config.Config.InstanceUpToDateSeconds, instanceKey.Hostname, instanceKey.Port).Scan(
		 	&instance.ServerID,
		 	&instance.Version,
		 	&instance.Binlog_format,
		 	&instance.LogBinEnabled,
		 	&instance.LogSlaveUpdatesEnabled,
		 	&instance.SelfBinlogCoordinates.LogFile,
		 	&instance.SelfBinlogCoordinates.LogPos,
		 	&instance.MasterKey.Hostname,
		 	&instance.MasterKey.Port,
		 	&instance.Slave_SQL_Running,
		 	&instance.Slave_IO_Running,
		 	&instance.ReadBinlogCoordinates.LogFile,
		 	&instance.ReadBinlogCoordinates.LogPos,
		 	&instance.ExecBinlogCoordinates.LogFile,
		 	&instance.ExecBinlogCoordinates.LogPos,
		 	&instance.SecondsBehindMaster,
		 	&slaveHostsJson,
		 	&instance.ClusterName,
		 	&instance.IsUpToDate,
		 	&instance.IsLastSeenValid,
		)
	if err == sql.ErrNoRows {log.Infof("No entry for %+v", instanceKey); return instance, false, err}	
		
    if err != nil {log.Error("error on", instanceKey, err); return instance, false, err}
    instance.ReadSlaveHostsFromJson(slaveHostsJson)
    
	return instance, true, err
}


func ReadClusterInstances(clusterName string) ([](*Instance), error) {
	instances := [](*Instance){}

	db,	err	:=	sqlutils.OpenOrchestrator()
	if	err	!=	nil	{
		return instances, log.Errore(err)
	}
	if strings.Index(clusterName, "'") >= 0 {
		return instances, log.Errorf("Invalid cluster name: %s", clusterName)	
	}

	query := fmt.Sprintf(`
		select 
			*,
			ifnull(timestampdiff(second, last_seen, now()) <= %d, false) as is_up_to_date,
			is_last_seen_valid
		from 
			database_instance 
		where
			cluster_name = '%s'
		order by
			hostname, port`, config.Config.InstanceUpToDateSeconds, clusterName)

    err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
    	instance := NewInstance()
    	
    	instance.Key.Hostname = m.GetString("hostname")
    	instance.Key.Port = m.GetInt("port")
		instance.Slave_IO_Running = (m.GetString("Slave_IO_Running") == "Yes")
      	instance.Slave_SQL_Running = (m.GetString("Slave_SQL_Running") == "Yes")
       	instance.ReadBinlogCoordinates.LogFile = m.GetString("Master_Log_File")
       	instance.ReadBinlogCoordinates.LogPos = m.GetInt64("Read_Master_Log_Pos")
       	instance.ExecBinlogCoordinates.LogFile = m.GetString("Relay_Master_Log_File")
       	instance.ExecBinlogCoordinates.LogPos = m.GetInt64("Exec_Master_Log_Pos")
    	instance.ServerID = m.GetUint("server_id")
	 	instance.Version = m.GetString("version")
	 	instance.Binlog_format = m.GetString("binlog_format")
	 	instance.LogBinEnabled = m.GetBool("log_bin")
	 	instance.LogSlaveUpdatesEnabled = m.GetBool("log_slave_updates")
	 	instance.SelfBinlogCoordinates.LogFile = m.GetString("binary_log_file")
	 	instance.SelfBinlogCoordinates.LogPos = m.GetInt64("binary_log_pos")
	 	instance.MasterKey.Hostname = m.GetString("master_host")
	 	instance.MasterKey.Port = m.GetInt("master_port")
	 	instance.Slave_SQL_Running = m.GetBool("slave_sql_running")
	 	instance.Slave_IO_Running = m.GetBool("slave_io_running")
	 	instance.ReadBinlogCoordinates.LogFile = m.GetString("master_log_file")
	 	instance.ReadBinlogCoordinates.LogPos = m.GetInt64("read_master_log_pos")
	 	instance.ExecBinlogCoordinates.LogFile = m.GetString("relay_master_log_file")
	 	instance.ExecBinlogCoordinates.LogPos = m.GetInt64("exec_master_log_pos")
	 	instance.SecondsBehindMaster = m.GetInt("seconds_behind_master")
	 	slaveHostsJson := m.GetString("slave_hosts")
	 	instance.ClusterName = m.GetString("cluster_name")
	 	instance.IsUpToDate = m.GetBool("is_up_to_date")
	 	instance.IsLastSeenValid = m.GetBool("is_last_seen_valid")
	 	
	 	instance.ReadSlaveHostsFromJson(slaveHostsJson)

    	instances = append(instances, instance)
    	return err       	
   	})

	return instances, err
}



func ReadClusters() ([]string, error) {
	clusterNames := []string{}

	db,	err	:=	sqlutils.OpenOrchestrator()
	if	err	!=	nil	{
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


func ReadOutdatedInstanceKeys() ([]InstanceKey, error) {
	res := []InstanceKey{}
	query := fmt.Sprintf(`
		select 
			hostname, port 
		from 
			database_instance 
		where
			timestampdiff(second, last_checked, now()) >= %d
    		and ifnull(timestampdiff(second, last_seen, now()) <= %d, false) = false `, 
    		config.Config.InstanceReattemptSeconds, config.Config.InstanceUpToDateSeconds)
	db,	err	:=	sqlutils.OpenOrchestrator()
    if err != nil {goto Cleanup}
    
    err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		cname, err := GetCNAME(m["hostname"])
	    if err != nil {return err}
		port, err := strconv.Atoi(m["port"])
	    if err != nil {return err}
    	instanceKey := InstanceKey{Hostname: cname, Port: port}
    	res = append(res, instanceKey)
    	return err       	
   	})
	Cleanup:

	if err	!=	nil	{
		log.Errore(err)
	}
	return res, err

}



func ReadMaintenanceInstanceKeys() ([]InstanceKey, error) {
	res := []InstanceKey{}
	query := fmt.Sprintf(`
		select 
			hostname, port 
		from 
			database_instance_maintenance
		where
			maintenance_active = 1
		order by
			database_instance_maintenance_id
		`)
	db,	err	:=	sqlutils.OpenOrchestrator()
    if err != nil {goto Cleanup}
    
    err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
		cname, err := GetCNAME(m["hostname"])
	    if err != nil {return err}
		port, err := strconv.Atoi(m["port"])
	    if err != nil {return err}
    	instanceKey := InstanceKey{Hostname: cname, Port: port}
    	res = append(res, instanceKey)
    	return err       	
   	})
	Cleanup:

	if err	!=	nil	{
		log.Errore(err)
	}
	return res, err

}



func WriteInstance(instance *Instance, lastError error) error {
	db,	err	:=	sqlutils.OpenOrchestrator()
	if err != nil {return log.Errore(err)}
	
	_, err = sqlutils.Exec(db, `
        	replace into database_instance (
        		hostname,
        		port,
        		last_checked,
        		server_id,
				version,
				binlog_format,
				log_bin,
				log_slave_updates,
				binary_log_file,
				binary_log_pos,
				master_host,
				master_port,
				slave_sql_running,
				slave_io_running,
				master_log_file,
				read_master_log_pos,
				relay_master_log_file,
				exec_master_log_pos,
				seconds_behind_master,
				num_slave_hosts,
				slave_hosts,
				cluster_name
			) values (?, ?, NOW(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			instance.Key.Hostname, 
		 	instance.Key.Port,
		 	instance.ServerID,
		 	instance.Version,
		 	instance.Binlog_format,
		 	instance.LogBinEnabled,
		 	instance.LogSlaveUpdatesEnabled,
			instance.SelfBinlogCoordinates.LogFile,
			instance.SelfBinlogCoordinates.LogPos,
		 	instance.MasterKey.Hostname,
		 	instance.MasterKey.Port,
		 	instance.Slave_SQL_Running,
		 	instance.Slave_IO_Running,
		 	instance.ReadBinlogCoordinates.LogFile,
		 	instance.ReadBinlogCoordinates.LogPos,
		 	instance.ExecBinlogCoordinates.LogFile,
		 	instance.ExecBinlogCoordinates.LogPos,
		 	instance.SecondsBehindMaster,
		 	len(instance.SlaveHosts),
		 	instance.GetSlaveHostsAsJson(),
		 	instance.ClusterName,
		 	)
    if err != nil {return log.Errore(err)}
	
	if lastError == nil {
		sqlutils.Exec(db, `
        	update database_instance set last_seen = NOW(), is_last_seen_valid = 1 where hostname=? and port=?
        	`, instance.Key.Hostname, instance.Key.Port,
        )
	}
	return nil
}


func ForgetInstance(instanceKey *InstanceKey) error {
	db,	err	:=	sqlutils.OpenOrchestrator()
	if err != nil {return log.Errore(err)}
	
	_, err = sqlutils.Exec(db, `
			delete 
				from database_instance 
			where 
				hostname = ? and port = ?`,
			instanceKey.Hostname, 
		 	instanceKey.Port,
		 )
	return err		 
}


func BeginMaintenance(instanceKey *InstanceKey, owner string, reason string) error {
	db,	err	:=	sqlutils.OpenOrchestrator()
	if err != nil {return log.Errore(err)}
	
	res, err := sqlutils.Exec(db, `
			insert ignore
				into database_instance_maintenance (
					hostname, port, maintenance_active, begin_timestamp, end_timestamp, owner, reason
				) VALUES (
					?, ?, 1, NOW(), NULL, ?, ?
				)
			`,
			instanceKey.Hostname, 
		 	instanceKey.Port,
		 	owner, 
		 	reason,
		 )
	if err != nil {return log.Errore(err)}	
	
	if affected, _ := res.RowsAffected(); affected == 0 {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance for instance: %+v", instanceKey))
	} 
	return err		 
}



func EndMaintenance(instanceKey *InstanceKey) error {
	db,	err	:=	sqlutils.OpenOrchestrator()
	if err != nil {return log.Errore(err)}
	
	res, err := sqlutils.Exec(db, `
			update
				database_instance_maintenance
			set  
				maintenance_active = NULL,
				end_timestamp = NOW()
			where
				hostname = ? 
				and port = ?
				and maintenance_active = 1
			`,
			instanceKey.Hostname, 
		 	instanceKey.Port,
		 )
	if err != nil {return log.Errore(err)}
	if affected, _ := res.RowsAffected(); affected == 0 {
		err = errors.New(fmt.Sprintf("Instance is not in maintenance mode: %+v", instanceKey))
	} 
	return err		 
}


func StopSlaveNicely(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {return instance, log.Errore(err)}
	
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	
	_, err = ExecInstance(instanceKey, `stop slave io_thread`)
	
	for up_to_date := false; !up_to_date; {
		instance, err = ReadTopologyInstance(instanceKey)
		if err != nil {return instance, log.Errore(err)}
		
		if instance.SQLThreadUpToDate() {
			up_to_date = true
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}
	_, err = ExecInstance(instanceKey, `stop slave`)
	if err != nil {return instance, log.Errore(err)}
	
	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}


func StopSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {return instance, log.Errore(err)}
	
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	_, err = ExecInstance(instanceKey, `stop slave`)
	if err != nil {return instance, log.Errore(err)}
	instance, err = ReadTopologyInstance(instanceKey)

	log.Infof("Stopped slave on %+v, Self:%+v, Exec:%+v", *instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates) 
	return instance, err
}


func StartSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {return instance, log.Errore(err)}
	
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	
	_, err = ExecInstance(instanceKey, `start slave`)
	if err != nil {return instance, log.Errore(err)}
	log.Infof("Started slave on %+v", instanceKey) 
	
	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}


func StartSlaveUntilMasterCoordinates(instanceKey *InstanceKey, masterCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {return instance, log.Errore(err)}
	
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	if instance.SlaveRunning() {
		return instance, errors.New(fmt.Sprintf("slave already running: %+v", instanceKey))
	}

	log.Infof("Will start slave on %+v until coordinates: %+v", instanceKey, masterCoordinates) 

	_, err = ExecInstance(instanceKey, fmt.Sprintf("start slave until master_log_file='%s', master_log_pos=%d", 
		masterCoordinates.LogFile, masterCoordinates.LogPos))
	if err != nil {return instance, log.Errore(err)}
	
		
	for up_to_date := false; !up_to_date; {
		instance, err = ReadTopologyInstance(instanceKey)
		if err != nil {return instance, log.Errore(err)}
		
		switch {
			case instance.ExecBinlogCoordinates.SmallerThan(masterCoordinates): time.Sleep(200 * time.Millisecond)
			case instance.ExecBinlogCoordinates.Equals(masterCoordinates): up_to_date = true
			case masterCoordinates.SmallerThan(&instance.ExecBinlogCoordinates): return instance, errors.New(fmt.Sprintf("Start SLAVE UNTIL is past coordinates: %+v", instanceKey))
		}
	}
		
	instance, err = StopSlave(instanceKey)
	if err != nil {return instance, log.Errore(err)} 
	
	return instance, err
}


func ChangeMasterTo(instanceKey *InstanceKey, masterKey *InstanceKey, masterBinlogCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {return instance, log.Errore(err)}
	
	if instance.SlaveRunning() {
		return instance, errors.New(fmt.Sprintf("Cannot change master on: %+v because slave is running", instanceKey))
	}
	
	_, err = ExecInstance(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_log_file='%s', master_log_pos=%d", 
		masterKey.Hostname, masterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos))
	if err != nil {return instance, log.Errore(err)}
	log.Infof("Changed master on %+v to: %+v, %+v", instanceKey, masterKey, masterBinlogCoordinates) 
	
	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}


func MasterPosWait(instanceKey *InstanceKey, binlogCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if err != nil {return instance, log.Errore(err)}
	
	_, err = ExecInstance(instanceKey, fmt.Sprintf("select master_pos_wait('%s', %d)", 
		binlogCoordinates.LogFile, binlogCoordinates.LogPos))
	if err != nil {return instance, log.Errore(err)}
	log.Infof("Instance %+v has reached coordinates: %+v", instanceKey, binlogCoordinates) 
	
	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}
