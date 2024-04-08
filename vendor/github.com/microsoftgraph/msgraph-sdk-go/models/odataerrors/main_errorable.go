package odataerrors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MainErrorable 
type MainErrorable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCode()(*string)
    GetDetails()([]ErrorDetailsable)
    GetInnererror()(InnerErrorable)
    GetMessage()(*string)
    GetTarget()(*string)
    SetCode(value *string)()
    SetDetails(value []ErrorDetailsable)()
    SetInnererror(value InnerErrorable)()
    SetMessage(value *string)()
    SetTarget(value *string)()
}
