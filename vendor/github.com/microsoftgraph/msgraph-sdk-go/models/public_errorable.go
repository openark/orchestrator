package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PublicErrorable 
type PublicErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetDetails()([]PublicErrorDetailable)
    GetInnerError()(PublicInnerErrorable)
    GetMessage()(*string)
    GetOdataType()(*string)
    GetTarget()(*string)
    SetCode(value *string)()
    SetDetails(value []PublicErrorDetailable)()
    SetInnerError(value PublicInnerErrorable)()
    SetMessage(value *string)()
    SetOdataType(value *string)()
    SetTarget(value *string)()
}
