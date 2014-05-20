package inst

import (

)


type Audit struct {
	AuditId				int64
	AuditTimestamp		string
	AuditType			string
	AuditInstanceKey	InstanceKey
	Message				string
}
