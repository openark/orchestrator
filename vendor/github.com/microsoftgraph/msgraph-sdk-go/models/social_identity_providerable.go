package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SocialIdentityProviderable 
type SocialIdentityProviderable interface {
    IdentityProviderBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetIdentityProviderType()(*string)
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetIdentityProviderType(value *string)()
}
