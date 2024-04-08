package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CloudApplicationEvidenceable 
type CloudApplicationEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppId()(*int64)
    GetDisplayName()(*string)
    GetInstanceId()(*int64)
    GetInstanceName()(*string)
    GetSaasAppId()(*int64)
    SetAppId(value *int64)()
    SetDisplayName(value *string)()
    SetInstanceId(value *int64)()
    SetInstanceName(value *string)()
    SetSaasAppId(value *int64)()
}
