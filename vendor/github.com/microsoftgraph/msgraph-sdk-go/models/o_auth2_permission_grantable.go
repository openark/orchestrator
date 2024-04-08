package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OAuth2PermissionGrantable 
type OAuth2PermissionGrantable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetConsentType()(*string)
    GetPrincipalId()(*string)
    GetResourceId()(*string)
    GetScope()(*string)
    SetClientId(value *string)()
    SetConsentType(value *string)()
    SetPrincipalId(value *string)()
    SetResourceId(value *string)()
    SetScope(value *string)()
}
