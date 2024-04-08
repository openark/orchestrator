package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceUpdateMessageViewpointable 
type ServiceUpdateMessageViewpointable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsArchived()(*bool)
    GetIsFavorited()(*bool)
    GetIsRead()(*bool)
    GetOdataType()(*string)
    SetIsArchived(value *bool)()
    SetIsFavorited(value *bool)()
    SetIsRead(value *bool)()
    SetOdataType(value *string)()
}
