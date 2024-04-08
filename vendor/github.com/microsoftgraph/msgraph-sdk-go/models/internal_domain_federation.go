package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InternalDomainFederation 
type InternalDomainFederation struct {
    SamlOrWsFedProvider
    // URL of the endpoint used by active clients when authenticating with federated domains set up for single sign-on in Azure Active Directory (Azure AD). Corresponds to the ActiveLogOnUri property of the Set-MsolDomainFederationSettings MSOnline v1 PowerShell cmdlet.
    activeSignInUri *string
    // Determines whether Azure AD accepts the MFA performed by the federated IdP when a federated user accesses an application that is governed by a conditional access policy that requires MFA. The possible values are: acceptIfMfaDoneByFederatedIdp, enforceMfaByFederatedIdp, rejectMfaByFederatedIdp, unknownFutureValue. For more information, see federatedIdpMfaBehavior values.
    federatedIdpMfaBehavior *FederatedIdpMfaBehavior
    // If true, when SAML authentication requests are sent to the federated SAML IdP, Azure AD will sign those requests using the OrgID signing key. If false (default), the SAML authentication requests sent to the federated IdP are not signed.
    isSignedAuthenticationRequestRequired *bool
    // Fallback token signing certificate that is used to sign tokens when the primary signing certificate expires. Formatted as Base64 encoded strings of the public portion of the federated IdP's token signing certificate. Needs to be compatible with the X509Certificate2 class. Much like the signingCertificate, the nextSigningCertificate property is used if a rollover is required outside of the auto-rollover update, a new federation service is being set up, or if the new token signing certificate is not present in the federation properties after the federation service certificate has been updated.
    nextSigningCertificate *string
    // Sets the preferred behavior for the sign-in prompt. The possible values are: translateToFreshPasswordAuthentication, nativeSupport, disabled, unknownFutureValue.
    promptLoginBehavior *PromptLoginBehavior
    // Provides status and timestamp of the last update of the signing certificate.
    signingCertificateUpdateStatus SigningCertificateUpdateStatusable
    // URI that clients are redirected to when they sign out of Azure AD services. Corresponds to the LogOffUri property of the Set-MsolDomainFederationSettings MSOnline v1 PowerShell cmdlet.
    signOutUri *string
}
// NewInternalDomainFederation instantiates a new InternalDomainFederation and sets the default values.
func NewInternalDomainFederation()(*InternalDomainFederation) {
    m := &InternalDomainFederation{
        SamlOrWsFedProvider: *NewSamlOrWsFedProvider(),
    }
    odataTypeValue := "#microsoft.graph.internalDomainFederation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateInternalDomainFederationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInternalDomainFederationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInternalDomainFederation(), nil
}
// GetActiveSignInUri gets the activeSignInUri property value. URL of the endpoint used by active clients when authenticating with federated domains set up for single sign-on in Azure Active Directory (Azure AD). Corresponds to the ActiveLogOnUri property of the Set-MsolDomainFederationSettings MSOnline v1 PowerShell cmdlet.
func (m *InternalDomainFederation) GetActiveSignInUri()(*string) {
    return m.activeSignInUri
}
// GetFederatedIdpMfaBehavior gets the federatedIdpMfaBehavior property value. Determines whether Azure AD accepts the MFA performed by the federated IdP when a federated user accesses an application that is governed by a conditional access policy that requires MFA. The possible values are: acceptIfMfaDoneByFederatedIdp, enforceMfaByFederatedIdp, rejectMfaByFederatedIdp, unknownFutureValue. For more information, see federatedIdpMfaBehavior values.
func (m *InternalDomainFederation) GetFederatedIdpMfaBehavior()(*FederatedIdpMfaBehavior) {
    return m.federatedIdpMfaBehavior
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternalDomainFederation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SamlOrWsFedProvider.GetFieldDeserializers()
    res["activeSignInUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActiveSignInUri(val)
        }
        return nil
    }
    res["federatedIdpMfaBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFederatedIdpMfaBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFederatedIdpMfaBehavior(val.(*FederatedIdpMfaBehavior))
        }
        return nil
    }
    res["isSignedAuthenticationRequestRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSignedAuthenticationRequestRequired(val)
        }
        return nil
    }
    res["nextSigningCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextSigningCertificate(val)
        }
        return nil
    }
    res["promptLoginBehavior"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePromptLoginBehavior)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPromptLoginBehavior(val.(*PromptLoginBehavior))
        }
        return nil
    }
    res["signingCertificateUpdateStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSigningCertificateUpdateStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningCertificateUpdateStatus(val.(SigningCertificateUpdateStatusable))
        }
        return nil
    }
    res["signOutUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignOutUri(val)
        }
        return nil
    }
    return res
}
// GetIsSignedAuthenticationRequestRequired gets the isSignedAuthenticationRequestRequired property value. If true, when SAML authentication requests are sent to the federated SAML IdP, Azure AD will sign those requests using the OrgID signing key. If false (default), the SAML authentication requests sent to the federated IdP are not signed.
func (m *InternalDomainFederation) GetIsSignedAuthenticationRequestRequired()(*bool) {
    return m.isSignedAuthenticationRequestRequired
}
// GetNextSigningCertificate gets the nextSigningCertificate property value. Fallback token signing certificate that is used to sign tokens when the primary signing certificate expires. Formatted as Base64 encoded strings of the public portion of the federated IdP's token signing certificate. Needs to be compatible with the X509Certificate2 class. Much like the signingCertificate, the nextSigningCertificate property is used if a rollover is required outside of the auto-rollover update, a new federation service is being set up, or if the new token signing certificate is not present in the federation properties after the federation service certificate has been updated.
func (m *InternalDomainFederation) GetNextSigningCertificate()(*string) {
    return m.nextSigningCertificate
}
// GetPromptLoginBehavior gets the promptLoginBehavior property value. Sets the preferred behavior for the sign-in prompt. The possible values are: translateToFreshPasswordAuthentication, nativeSupport, disabled, unknownFutureValue.
func (m *InternalDomainFederation) GetPromptLoginBehavior()(*PromptLoginBehavior) {
    return m.promptLoginBehavior
}
// GetSigningCertificateUpdateStatus gets the signingCertificateUpdateStatus property value. Provides status and timestamp of the last update of the signing certificate.
func (m *InternalDomainFederation) GetSigningCertificateUpdateStatus()(SigningCertificateUpdateStatusable) {
    return m.signingCertificateUpdateStatus
}
// GetSignOutUri gets the signOutUri property value. URI that clients are redirected to when they sign out of Azure AD services. Corresponds to the LogOffUri property of the Set-MsolDomainFederationSettings MSOnline v1 PowerShell cmdlet.
func (m *InternalDomainFederation) GetSignOutUri()(*string) {
    return m.signOutUri
}
// Serialize serializes information the current object
func (m *InternalDomainFederation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SamlOrWsFedProvider.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("activeSignInUri", m.GetActiveSignInUri())
        if err != nil {
            return err
        }
    }
    if m.GetFederatedIdpMfaBehavior() != nil {
        cast := (*m.GetFederatedIdpMfaBehavior()).String()
        err = writer.WriteStringValue("federatedIdpMfaBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isSignedAuthenticationRequestRequired", m.GetIsSignedAuthenticationRequestRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("nextSigningCertificate", m.GetNextSigningCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetPromptLoginBehavior() != nil {
        cast := (*m.GetPromptLoginBehavior()).String()
        err = writer.WriteStringValue("promptLoginBehavior", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("signingCertificateUpdateStatus", m.GetSigningCertificateUpdateStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signOutUri", m.GetSignOutUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActiveSignInUri sets the activeSignInUri property value. URL of the endpoint used by active clients when authenticating with federated domains set up for single sign-on in Azure Active Directory (Azure AD). Corresponds to the ActiveLogOnUri property of the Set-MsolDomainFederationSettings MSOnline v1 PowerShell cmdlet.
func (m *InternalDomainFederation) SetActiveSignInUri(value *string)() {
    m.activeSignInUri = value
}
// SetFederatedIdpMfaBehavior sets the federatedIdpMfaBehavior property value. Determines whether Azure AD accepts the MFA performed by the federated IdP when a federated user accesses an application that is governed by a conditional access policy that requires MFA. The possible values are: acceptIfMfaDoneByFederatedIdp, enforceMfaByFederatedIdp, rejectMfaByFederatedIdp, unknownFutureValue. For more information, see federatedIdpMfaBehavior values.
func (m *InternalDomainFederation) SetFederatedIdpMfaBehavior(value *FederatedIdpMfaBehavior)() {
    m.federatedIdpMfaBehavior = value
}
// SetIsSignedAuthenticationRequestRequired sets the isSignedAuthenticationRequestRequired property value. If true, when SAML authentication requests are sent to the federated SAML IdP, Azure AD will sign those requests using the OrgID signing key. If false (default), the SAML authentication requests sent to the federated IdP are not signed.
func (m *InternalDomainFederation) SetIsSignedAuthenticationRequestRequired(value *bool)() {
    m.isSignedAuthenticationRequestRequired = value
}
// SetNextSigningCertificate sets the nextSigningCertificate property value. Fallback token signing certificate that is used to sign tokens when the primary signing certificate expires. Formatted as Base64 encoded strings of the public portion of the federated IdP's token signing certificate. Needs to be compatible with the X509Certificate2 class. Much like the signingCertificate, the nextSigningCertificate property is used if a rollover is required outside of the auto-rollover update, a new federation service is being set up, or if the new token signing certificate is not present in the federation properties after the federation service certificate has been updated.
func (m *InternalDomainFederation) SetNextSigningCertificate(value *string)() {
    m.nextSigningCertificate = value
}
// SetPromptLoginBehavior sets the promptLoginBehavior property value. Sets the preferred behavior for the sign-in prompt. The possible values are: translateToFreshPasswordAuthentication, nativeSupport, disabled, unknownFutureValue.
func (m *InternalDomainFederation) SetPromptLoginBehavior(value *PromptLoginBehavior)() {
    m.promptLoginBehavior = value
}
// SetSigningCertificateUpdateStatus sets the signingCertificateUpdateStatus property value. Provides status and timestamp of the last update of the signing certificate.
func (m *InternalDomainFederation) SetSigningCertificateUpdateStatus(value SigningCertificateUpdateStatusable)() {
    m.signingCertificateUpdateStatus = value
}
// SetSignOutUri sets the signOutUri property value. URI that clients are redirected to when they sign out of Azure AD services. Corresponds to the LogOffUri property of the Set-MsolDomainFederationSettings MSOnline v1 PowerShell cmdlet.
func (m *InternalDomainFederation) SetSignOutUri(value *string)() {
    m.signOutUri = value
}
