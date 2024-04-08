package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FederatedIdentityCredentialable 
type FederatedIdentityCredentialable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudiences()([]string)
    GetDescription()(*string)
    GetIssuer()(*string)
    GetName()(*string)
    GetSubject()(*string)
    SetAudiences(value []string)()
    SetDescription(value *string)()
    SetIssuer(value *string)()
    SetName(value *string)()
    SetSubject(value *string)()
}
