package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningStepable 
type ProvisioningStepable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDetails()(DetailsInfoable)
    GetName()(*string)
    GetOdataType()(*string)
    GetProvisioningStepType()(*ProvisioningStepType)
    GetStatus()(*ProvisioningResult)
    SetDescription(value *string)()
    SetDetails(value DetailsInfoable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetProvisioningStepType(value *ProvisioningStepType)()
    SetStatus(value *ProvisioningResult)()
}
