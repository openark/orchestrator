package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InternalDomainFederationable 
type InternalDomainFederationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    SamlOrWsFedProviderable
    GetActiveSignInUri()(*string)
    GetFederatedIdpMfaBehavior()(*FederatedIdpMfaBehavior)
    GetIsSignedAuthenticationRequestRequired()(*bool)
    GetNextSigningCertificate()(*string)
    GetPromptLoginBehavior()(*PromptLoginBehavior)
    GetSigningCertificateUpdateStatus()(SigningCertificateUpdateStatusable)
    GetSignOutUri()(*string)
    SetActiveSignInUri(value *string)()
    SetFederatedIdpMfaBehavior(value *FederatedIdpMfaBehavior)()
    SetIsSignedAuthenticationRequestRequired(value *bool)()
    SetNextSigningCertificate(value *string)()
    SetPromptLoginBehavior(value *PromptLoginBehavior)()
    SetSigningCertificateUpdateStatus(value SigningCertificateUpdateStatusable)()
    SetSignOutUri(value *string)()
}
