package inst

import (
	"fmt"
	"errors"
	"strconv"
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/orchestrator/db"
	"github.com/outbrain/log"
)


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
	db,	err	:=	db.OpenOrchestrator()
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
		maintenanceToken, _ = res.LastInsertId()
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
	} 
	return err		 
}

