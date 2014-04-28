package inst

import (
	"fmt"
	"errors"
	"log"
	"time"
	"strconv"
	"database/sql"
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/config"
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
    foundBySlaveHosts := false

	db,	err	:=	sqlutils.OpenTopology(instanceKey.Hostname, instanceKey.Port)

    if err != nil {goto Cleanup}

   	instance.Key = *instanceKey
    err = db.QueryRow("select @@global.server_id, @@global.version, @@global.binlog_format, @@global.log_bin, @@global.log_slave_updates").Scan(
       	&instance.ServerID, &instance.Version, &instance.Binlog_format, &instance.LogBinEnabled, &instance.LogSlaveUpdatesEnabled)
    if err != nil {goto Cleanup}
    err = sqlutils.QueryRowsMap(db, "show slave status", func(m map[string]string) {
		instance.Slave_IO_Running = (m["Slave_IO_Running"] == "Yes")
      	instance.Slave_SQL_Running = (m["Slave_SQL_Running"] == "Yes")
       	instance.ReadBinlogCoordinates.LogFile = m["Master_Log_File"]
       	instance.ReadBinlogCoordinates.LogPos, _ = strconv.ParseInt(m["Read_Master_Log_Pos"], 10, 0)
       	instance.ExecBinlogCoordinates.LogFile = m["Relay_Master_Log_File"]
       	instance.ExecBinlogCoordinates.LogPos, _ = strconv.ParseInt(m["Exec_Master_Log_Pos"], 10, 0)
       	instance.Master_Host, _ = GetCNAME(m["Master_Host"])
       	instance.Master_Port, _ = strconv.Atoi(m["Master_Port"])
       	if config.Config.SlaveLagQuery == "" {
        	instance.SecondsBehindMaster, _ = strconv.Atoi(m["Seconds_Behind_Master"])
        }
   	})
    if err != nil {goto Cleanup}

	if config.Config.SlaveLagQuery != "" {
		err = db.QueryRow(config.Config.SlaveLagQuery).Scan(&instance.Version, &instance.SecondsBehindMaster)
	    if err != nil {goto Cleanup}
	}
        
    err = sqlutils.QueryRowsMap(db, "show master status", func(m map[string]string) {
       	instance.SelfBinlogCoordinates.LogFile = m["File"]
       	instance.SelfBinlogCoordinates.LogPos, _ = strconv.ParseInt(m["Position"], 10, 0)
   	})
    if err != nil {goto Cleanup}
        
    // Get slaves, either by SHOW SLAVE HOSTS or via PROCESSLIST
    if config.Config.DiscoverByShowSlaveHosts {
        err = sqlutils.QueryRowsMap(db, `show slave hosts`, 
        		func(m map[string]string) {
        			cname, err := GetCNAME(m["Host"])
        			port, err := strconv.Atoi(m["Port"])
        			if err == nil {
	        			slaveKey := InstanceKey{Hostname: cname, Port: port}
						instance.AddSlaveKey(&slaveKey)
						foundBySlaveHosts = true
					}
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
        		func(m map[string]string) {
        			cname, _ := GetCNAME(m["slave_hostname"])
        			slaveKey := InstanceKey{Hostname: cname, Port: instance.Key.Port}
					instance.AddSlaveKey(&slaveKey)
		       	})
			
        if err != nil {goto Cleanup}
	}
		        
    if err != nil {goto Cleanup}

	Cleanup:
	_ = WriteInstance(instance, err)
	if err	!=	nil	{
		log.Println(err)
	}
	return instance, err
}

func ReadInstance(instanceKey *InstanceKey) (*Instance, bool, error) {
	db,	err	:=	sqlutils.OpenOrchestrator()

	instance := NewInstance()
	instance.Key = *instanceKey
	if	err	!=	nil	{
		log.Println(err)
	} else {
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
			 	&instance.Master_Host,
			 	&instance.Master_Port,
			 	&instance.Slave_SQL_Running,
			 	&instance.Slave_IO_Running,
			 	&instance.ReadBinlogCoordinates.LogFile,
			 	&instance.ReadBinlogCoordinates.LogPos,
			 	&instance.ExecBinlogCoordinates.LogFile,
			 	&instance.ExecBinlogCoordinates.LogPos,
			 	&instance.SecondsBehindMaster,
			 	&slaveHostsJson,
			 	&instance.IsUpToDate,
			 	&instance.IsLastSeenValid,
			 	)
		if err == sql.ErrNoRows {log.Println("No entry for ", instanceKey); return instance, false, err}	
		
        if	err	!=	nil	{log.Println("error on", instanceKey, err); return instance, true, err}
        instance.ReadSlaveHostsFromJson(slaveHostsJson)
	}
	return instance, true, err
}

func WriteInstance(instance *Instance, lastError error) error {
	db,	err	:=	sqlutils.OpenOrchestrator()
	if	err	!=	nil	{log.Println(err); return err}
	
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
				slave_hosts
			) values (?, ?, NOW(), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
			instance.Key.Hostname, 
		 	instance.Key.Port,
		 	instance.ServerID,
		 	instance.Version,
		 	instance.Binlog_format,
		 	instance.LogBinEnabled,
		 	instance.LogSlaveUpdatesEnabled,
			instance.SelfBinlogCoordinates.LogFile,
			instance.SelfBinlogCoordinates.LogPos,
		 	instance.Master_Host,
		 	instance.Master_Port,
		 	instance.Slave_SQL_Running,
		 	instance.Slave_IO_Running,
		 	instance.ReadBinlogCoordinates.LogFile,
		 	instance.ReadBinlogCoordinates.LogPos,
		 	instance.ExecBinlogCoordinates.LogFile,
		 	instance.ExecBinlogCoordinates.LogPos,
		 	instance.SecondsBehindMaster,
		 	instance.GetSlaveHostsAsJson(),
		 	)
	if	err	!=	nil	{log.Println(err); return err}
	
	if lastError == nil {
		sqlutils.Exec(db, `
        	update database_instance set last_seen = NOW(), is_last_seen_valid = 1 where hostname=? and port=?
        	`, instance.Key.Hostname, instance.Key.Port,
        )
	}
	return nil
}


func StopSlaveNicely(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err}
	
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	
	_, err = ExecInstance(instanceKey, `stop slave io_thread`)
	
	for up_to_date := false; !up_to_date; {
		instance, err = ReadTopologyInstance(instanceKey)
		if	err	!=	nil	{log.Println(err); return instance, err}
		
		if instance.SQLThreadUpToDate() {
			up_to_date = true
		} else {
			time.Sleep(200 * time.Millisecond)
		}
	}
	_, err = ExecInstance(instanceKey, `stop slave`)
	if	err	!=	nil	{log.Println(err); return instance, err}
	
	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}


func StopSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err}
	
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	_, err = ExecInstance(instanceKey, `stop slave`)
	if	err	!=	nil	{log.Println(err); return instance, err}
	instance, err = ReadTopologyInstance(instanceKey)

	log.Println(fmt.Sprintf("Stopped slave on %+v, Self:%+v, Exec:%+v", instanceKey, instance.SelfBinlogCoordinates, instance.ExecBinlogCoordinates)) 
	return instance, err
}


func StartSlave(instanceKey *InstanceKey) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err}
	
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	
	_, err = ExecInstance(instanceKey, `start slave`)
	if	err	!=	nil	{log.Println(err); return instance, err}
	log.Println(fmt.Sprintf("Started slave on %+v", instanceKey)) 
	
	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}


func StartSlaveUntilMasterCoordinates(instanceKey *InstanceKey, masterCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err}
	
	if !instance.IsSlave() {
		return instance, errors.New(fmt.Sprintf("instance is not a slave: %+v", instanceKey))
	}
	if instance.SlaveRunning() {
		return instance, errors.New(fmt.Sprintf("slave already running: %+v", instanceKey))
	}

	log.Println(fmt.Sprintf("Will start slave on %+v until coordinates: %+v", instanceKey, masterCoordinates)) 

	_, err = ExecInstance(instanceKey, fmt.Sprintf("start slave until master_log_file='%s', master_log_pos=%d", 
		masterCoordinates.LogFile, masterCoordinates.LogPos))
	if	err	!=	nil	{log.Println(err); return instance, err}
	
		
	for up_to_date := false; !up_to_date; {
		instance, err = ReadTopologyInstance(instanceKey)
		if	err	!=	nil	{log.Println(err); return instance, err}
		
		switch {
			case instance.ExecBinlogCoordinates.SmallerThan(masterCoordinates): time.Sleep(200 * time.Millisecond)
			case instance.ExecBinlogCoordinates.Equals(masterCoordinates): up_to_date = true
			case masterCoordinates.SmallerThan(&instance.ExecBinlogCoordinates): return instance, errors.New(fmt.Sprintf("Start SLAVE UNTIL is past coordinates: %+v", instanceKey))
		}
	}
		
	instance, err = StopSlave(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err} 
	
	return instance, err
}


func ChangeMasterTo(instanceKey *InstanceKey, masterKey *InstanceKey, masterBinlogCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err}
	
	if instance.SlaveRunning() {
		return instance, errors.New(fmt.Sprintf("Cannot change master on: %+v because slave is running", instanceKey))
	}
	
	_, err = ExecInstance(instanceKey, fmt.Sprintf("change master to master_host='%s', master_port=%d, master_log_file='%s', master_log_pos=%d", 
		masterKey.Hostname, masterKey.Port, masterBinlogCoordinates.LogFile, masterBinlogCoordinates.LogPos))
	if	err	!=	nil	{log.Println(err); return instance, err}
	log.Println(fmt.Sprintf("Changed master on %+v to: %+v, %+v", instanceKey, masterKey, masterBinlogCoordinates)) 
	
	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}


func MasterPosWait(instanceKey *InstanceKey, binlogCoordinates *BinlogCoordinates) (*Instance, error) {
	instance, err := ReadTopologyInstance(instanceKey)
	if	err	!=	nil	{log.Println(err); return instance, err}
	
	_, err = ExecInstance(instanceKey, fmt.Sprintf("select master_pos_wait('%s', %d)", 
		binlogCoordinates.LogFile, binlogCoordinates.LogPos))
	if	err	!=	nil	{log.Println(err); return instance, err}
	log.Println(fmt.Sprintf("Instance %+v has reached coordinates: %+v", instanceKey, binlogCoordinates)) 
	
	instance, err = ReadTopologyInstance(instanceKey)
	return instance, err
}
