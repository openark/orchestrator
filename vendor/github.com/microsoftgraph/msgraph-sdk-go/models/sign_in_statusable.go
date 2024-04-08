package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SignInStatusable 
type SignInStatusable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalDetails()(*string)
    GetErrorCode()(*int32)
    GetFailureReason()(*string)
    GetOdataType()(*string)
    SetAdditionalDetails(value *string)()
    SetErrorCode(value *int32)()
    SetFailureReason(value *string)()
    SetOdataType(value *string)()
}
