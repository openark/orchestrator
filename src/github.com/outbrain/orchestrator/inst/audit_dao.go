package inst

import (
	"fmt"
	"github.com/outbrain/sqlutils"
	"github.com/outbrain/orchestrator/db"
	"github.com/outbrain/log"
)


func AuditOperation(auditType string, instanceKey *InstanceKey, message string) error {
	db,	err	:=	db.OpenOrchestrator()
	if err != nil {return log.Errore(err)}
	
	if instanceKey == nil {
		instanceKey = &InstanceKey{}
	}
	
	_, err = sqlutils.Exec(db, `
			insert 
				into audit (
					audit_timestamp, audit_type, hostname, port, message
				) VALUES (
					NOW(), ?, ?, ?, ?
				)
			`,
			auditType,
			instanceKey.Hostname, 
		 	instanceKey.Port,
		 	message,
		 )
	if err != nil {return log.Errore(err)}
	
	return err
}


func ReadRecentAudit() ([]Audit, error) {
	res := []Audit{}
	query := fmt.Sprintf(`
		select 
			audit_id,
			audit_timestamp,
			audit_type,
			hostname,
			port,
			message
		from 
			audit
		order by
			audit_timestamp desc
		limit 100
		`)
	db,	err	:=	db.OpenOrchestrator()
    if err != nil {goto Cleanup}
    
    err = sqlutils.QueryRowsMap(db, query, func(m sqlutils.RowMap) error {
    	audit := Audit{}
    	audit.AuditId = m.GetInt64("audit_id")
    	audit.AuditTimestamp = m.GetString("audit_timestamp") 
    	audit.AuditType = m.GetString("audit_type")
    	audit.AuditInstanceKey.Hostname = m.GetString("hostname")
    	audit.AuditInstanceKey.Port = m.GetInt("port")
    	audit.Message = m.GetString("message") 

    	res = append(res, audit)
    	return err       	
   	})
	Cleanup:

	if err	!=	nil	{
		log.Errore(err)
	}
	return res, err

}
