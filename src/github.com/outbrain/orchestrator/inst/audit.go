package inst

import (

)

// Audit presents a single audit entry (namely in the database)
type Audit struct {
	AuditId				int64
	AuditTimestamp		string
	AuditType			string
	AuditInstanceKey	InstanceKey
	Message				string
}
