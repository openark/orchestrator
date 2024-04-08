package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharePointIdentitySetable 
type SharePointIdentitySetable interface {
    IdentitySetable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroup()(Identityable)
    GetSiteGroup()(SharePointIdentityable)
    GetSiteUser()(SharePointIdentityable)
    SetGroup(value Identityable)()
    SetSiteGroup(value SharePointIdentityable)()
    SetSiteUser(value SharePointIdentityable)()
}
