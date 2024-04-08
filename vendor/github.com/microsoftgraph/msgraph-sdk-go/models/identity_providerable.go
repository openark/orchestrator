package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityProviderable 
type IdentityProviderable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetName()(*string)
    GetType()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetName(value *string)()
    SetType(value *string)()
}
