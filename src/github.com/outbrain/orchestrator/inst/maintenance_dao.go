package inst

import (
	"fmt"
	"errors"
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/orchestrator/db"
	"github.com/outbrain/log"
)


func ReadActiveMaintenance() ([]Maintenance, error) {
	res := []Maintenance{}
	query := fmt.Sprintf(`
		select 
			database_instance_maintenance_id,
			hostname,
			port,
			begin_timestamp,
			timestampdiff(second, begin_timestamp, now()) as seconds_elapsed,
			maintenance_active,
			owner,
			reason
		from 
			database_instance_maintenance
		where
			maintenance_active = 1
		order by
			database_instance_maintenance_id
		`)
	db,	err	:=	db.OpenOrchestrator()
    if err != nil {goto Cleanup}
    
    err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
    	maintenance := Maintenance{}
    	maintenance.MaintenanceId = m.GetUint("database_instance_maintenance_id")
    	maintenance.Key.Hostname = m.GetString("hostname")
    	maintenance.Key.Port = m.GetInt("port")
    	maintenance.BeginTimestamp = m.GetString("begin_timestamp") 
    	maintenance.SecondsElapsed = m.GetUint("seconds_elapsed") 
    	maintenance.IsActive = m.GetBool("maintenance_active") 
    	maintenance.Owner = m.GetString("owner") 
    	maintenance.Reason = m.GetString("reason") 

    	res = append(res, maintenance)
    	return err       	
   	})
	Cleanup:

	if err	!=	nil	{
		log.Errore(err)
	}
	return res, err

}


func BeginMaintenance(instanceKey *InstanceKey, owner string, reason string) (int64, error) {
	db,	err	:=	db.OpenOrchestrator()
	var maintenanceToken int64 = 0
	if err != nil {return maintenanceToken, log.Errore(err)}
	
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
	if err != nil {return maintenanceToken, log.Errore(err)}	
	
	if affected, _ := res.RowsAffected(); affected == 0 {
		err = errors.New(fmt.Sprintf("Cannot begin maintenance for instance: %+v", instanceKey))
	} else {
		// success
		maintenanceToken, _ = res.LastInsertId()
		AuditOperation("begin-maintenance", instanceKey, fmt.Sprintf("maintenanceToken: %d", maintenanceToken))
	}
	return maintenanceToken, err		 
}



func EndMaintenanceByKey(instanceKey *InstanceKey) error {
	db,	err	:=	db.OpenOrchestrator()
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
	} else {
		// success
		AuditOperation("end-maintenance", instanceKey, "")
	}
	return err		 
}


func EndMaintenance(maintenanceToken int64) error {
	db,	err	:=	db.OpenOrchestrator()
	if err != nil {return log.Errore(err)}
	
	res, err := sqlutils.Exec(db, `
			update
				database_instance_maintenance
			set  
				maintenance_active = NULL,
				end_timestamp = NOW()
			where
				database_instance_maintenance_id = ? 
			`,
			maintenanceToken,
		 )
	if err != nil {return log.Errore(err)}
	if affected, _ := res.RowsAffected(); affected == 0 {
		err = errors.New(fmt.Sprintf("Instance is not in maintenance mode; token = %+v", maintenanceToken))
	} else {
		// success
		AuditOperation("end-maintenance", nil, fmt.Sprintf("maintenanceToken: %d", maintenanceToken))
	}
	return err		 
}

