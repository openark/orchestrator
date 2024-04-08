package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailClusterEvidenceable 
type MailClusterEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClusterBy()(*string)
    GetClusterByValue()(*string)
    GetEmailCount()(*int64)
    GetNetworkMessageIds()([]string)
    GetQuery()(*string)
    GetUrn()(*string)
    SetClusterBy(value *string)()
    SetClusterByValue(value *string)()
    SetEmailCount(value *int64)()
    SetNetworkMessageIds(value []string)()
    SetQuery(value *string)()
    SetUrn(value *string)()
}
