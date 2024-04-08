package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VmMetadataable 
type VmMetadataable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCloudProvider()(*VmCloudProvider)
    GetOdataType()(*string)
    GetResourceId()(*string)
    GetSubscriptionId()(*string)
    GetVmId()(*string)
    SetCloudProvider(value *VmCloudProvider)()
    SetOdataType(value *string)()
    SetResourceId(value *string)()
    SetSubscriptionId(value *string)()
    SetVmId(value *string)()
}
