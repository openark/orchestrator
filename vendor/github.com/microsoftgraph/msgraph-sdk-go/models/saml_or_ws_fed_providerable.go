package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SamlOrWsFedProviderable 
type SamlOrWsFedProviderable interface {
    IdentityProviderBaseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIssuerUri()(*string)
    GetMetadataExchangeUri()(*string)
    GetPassiveSignInUri()(*string)
    GetPreferredAuthenticationProtocol()(*AuthenticationProtocol)
    GetSigningCertificate()(*string)
    SetIssuerUri(value *string)()
    SetMetadataExchangeUri(value *string)()
    SetPassiveSignInUri(value *string)()
    SetPreferredAuthenticationProtocol(value *AuthenticationProtocol)()
    SetSigningCertificate(value *string)()
}
