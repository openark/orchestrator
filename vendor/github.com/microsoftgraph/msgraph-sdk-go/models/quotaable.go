package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Quotaable 
type Quotaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeleted()(*int64)
    GetOdataType()(*string)
    GetRemaining()(*int64)
    GetState()(*string)
    GetStoragePlanInformation()(StoragePlanInformationable)
    GetTotal()(*int64)
    GetUsed()(*int64)
    SetDeleted(value *int64)()
    SetOdataType(value *string)()
    SetRemaining(value *int64)()
    SetState(value *string)()
    SetStoragePlanInformation(value StoragePlanInformationable)()
    SetTotal(value *int64)()
    SetUsed(value *int64)()
}
