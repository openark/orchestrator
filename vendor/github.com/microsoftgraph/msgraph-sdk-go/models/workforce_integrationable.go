package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkforceIntegrationable 
type WorkforceIntegrationable interface {
    ChangeTrackedEntityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiVersion()(*int32)
    GetDisplayName()(*string)
    GetEncryption()(WorkforceIntegrationEncryptionable)
    GetIsActive()(*bool)
    GetSupportedEntities()(*WorkforceIntegrationSupportedEntities)
    GetUrl()(*string)
    SetApiVersion(value *int32)()
    SetDisplayName(value *string)()
    SetEncryption(value WorkforceIntegrationEncryptionable)()
    SetIsActive(value *bool)()
    SetSupportedEntities(value *WorkforceIntegrationSupportedEntities)()
    SetUrl(value *string)()
}
