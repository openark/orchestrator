package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CrossCloudAzureActiveDirectoryTenantable 
type CrossCloudAzureActiveDirectoryTenantable interface {
    IdentitySourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudInstance()(*string)
    GetDisplayName()(*string)
    GetTenantId()(*string)
    SetCloudInstance(value *string)()
    SetDisplayName(value *string)()
    SetTenantId(value *string)()
}
