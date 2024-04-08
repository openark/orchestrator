package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SamlOrWsFedProvider 
type SamlOrWsFedProvider struct {
    IdentityProviderBase
    // Issuer URI of the federation server.
    issuerUri *string
    // URI of the metadata exchange endpoint used for authentication from rich client applications.
    metadataExchangeUri *string
    // URI that web-based clients are directed to when signing in to Azure Active Directory (Azure AD) services.
    passiveSignInUri *string
    // Preferred authentication protocol. The possible values are: wsFed, saml, unknownFutureValue.
    preferredAuthenticationProtocol *AuthenticationProtocol
    // Current certificate used to sign tokens passed to the Microsoft identity platform. The certificate is formatted as a Base64 encoded string of the public portion of the federated IdP's token signing certificate and must be compatible with the X509Certificate2 class.   This property is used in the following scenarios:  if a rollover is required outside of the autorollover update a new federation service is being set up  if the new token signing certificate isn't present in the federation properties after the federation service certificate has been updated.   Azure AD updates certificates via an autorollover process in which it attempts to retrieve a new certificate from the federation service metadata, 30 days before expiry of the current certificate. If a new certificate isn't available, Azure AD monitors the metadata daily and will update the federation settings for the domain when a new certificate is available.
    signingCertificate *string
}
// NewSamlOrWsFedProvider instantiates a new SamlOrWsFedProvider and sets the default values.
func NewSamlOrWsFedProvider()(*SamlOrWsFedProvider) {
    m := &SamlOrWsFedProvider{
        IdentityProviderBase: *NewIdentityProviderBase(),
    }
    odataTypeValue := "#microsoft.graph.samlOrWsFedProvider"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSamlOrWsFedProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSamlOrWsFedProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.internalDomainFederation":
                        return NewInternalDomainFederation(), nil
                    case "#microsoft.graph.samlOrWsFedExternalDomainFederation":
                        return NewSamlOrWsFedExternalDomainFederation(), nil
                }
            }
        }
    }
    return NewSamlOrWsFedProvider(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SamlOrWsFedProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityProviderBase.GetFieldDeserializers()
    res["issuerUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuerUri(val)
        }
        return nil
    }
    res["metadataExchangeUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMetadataExchangeUri(val)
        }
        return nil
    }
    res["passiveSignInUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassiveSignInUri(val)
        }
        return nil
    }
    res["preferredAuthenticationProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationProtocol)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredAuthenticationProtocol(val.(*AuthenticationProtocol))
        }
        return nil
    }
    res["signingCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSigningCertificate(val)
        }
        return nil
    }
    return res
}
// GetIssuerUri gets the issuerUri property value. Issuer URI of the federation server.
func (m *SamlOrWsFedProvider) GetIssuerUri()(*string) {
    return m.issuerUri
}
// GetMetadataExchangeUri gets the metadataExchangeUri property value. URI of the metadata exchange endpoint used for authentication from rich client applications.
func (m *SamlOrWsFedProvider) GetMetadataExchangeUri()(*string) {
    return m.metadataExchangeUri
}
// GetPassiveSignInUri gets the passiveSignInUri property value. URI that web-based clients are directed to when signing in to Azure Active Directory (Azure AD) services.
func (m *SamlOrWsFedProvider) GetPassiveSignInUri()(*string) {
    return m.passiveSignInUri
}
// GetPreferredAuthenticationProtocol gets the preferredAuthenticationProtocol property value. Preferred authentication protocol. The possible values are: wsFed, saml, unknownFutureValue.
func (m *SamlOrWsFedProvider) GetPreferredAuthenticationProtocol()(*AuthenticationProtocol) {
    return m.preferredAuthenticationProtocol
}
// GetSigningCertificate gets the signingCertificate property value. Current certificate used to sign tokens passed to the Microsoft identity platform. The certificate is formatted as a Base64 encoded string of the public portion of the federated IdP's token signing certificate and must be compatible with the X509Certificate2 class.   This property is used in the following scenarios:  if a rollover is required outside of the autorollover update a new federation service is being set up  if the new token signing certificate isn't present in the federation properties after the federation service certificate has been updated.   Azure AD updates certificates via an autorollover process in which it attempts to retrieve a new certificate from the federation service metadata, 30 days before expiry of the current certificate. If a new certificate isn't available, Azure AD monitors the metadata daily and will update the federation settings for the domain when a new certificate is available.
func (m *SamlOrWsFedProvider) GetSigningCertificate()(*string) {
    return m.signingCertificate
}
// Serialize serializes information the current object
func (m *SamlOrWsFedProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityProviderBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("issuerUri", m.GetIssuerUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("metadataExchangeUri", m.GetMetadataExchangeUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passiveSignInUri", m.GetPassiveSignInUri())
        if err != nil {
            return err
        }
    }
    if m.GetPreferredAuthenticationProtocol() != nil {
        cast := (*m.GetPreferredAuthenticationProtocol()).String()
        err = writer.WriteStringValue("preferredAuthenticationProtocol", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signingCertificate", m.GetSigningCertificate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIssuerUri sets the issuerUri property value. Issuer URI of the federation server.
func (m *SamlOrWsFedProvider) SetIssuerUri(value *string)() {
    m.issuerUri = value
}
// SetMetadataExchangeUri sets the metadataExchangeUri property value. URI of the metadata exchange endpoint used for authentication from rich client applications.
func (m *SamlOrWsFedProvider) SetMetadataExchangeUri(value *string)() {
    m.metadataExchangeUri = value
}
// SetPassiveSignInUri sets the passiveSignInUri property value. URI that web-based clients are directed to when signing in to Azure Active Directory (Azure AD) services.
func (m *SamlOrWsFedProvider) SetPassiveSignInUri(value *string)() {
    m.passiveSignInUri = value
}
// SetPreferredAuthenticationProtocol sets the preferredAuthenticationProtocol property value. Preferred authentication protocol. The possible values are: wsFed, saml, unknownFutureValue.
func (m *SamlOrWsFedProvider) SetPreferredAuthenticationProtocol(value *AuthenticationProtocol)() {
    m.preferredAuthenticationProtocol = value
}
// SetSigningCertificate sets the signingCertificate property value. Current certificate used to sign tokens passed to the Microsoft identity platform. The certificate is formatted as a Base64 encoded string of the public portion of the federated IdP's token signing certificate and must be compatible with the X509Certificate2 class.   This property is used in the following scenarios:  if a rollover is required outside of the autorollover update a new federation service is being set up  if the new token signing certificate isn't present in the federation properties after the federation service certificate has been updated.   Azure AD updates certificates via an autorollover process in which it attempts to retrieve a new certificate from the federation service metadata, 30 days before expiry of the current certificate. If a new certificate isn't available, Azure AD monitors the metadata daily and will update the federation settings for the domain when a new certificate is available.
func (m *SamlOrWsFedProvider) SetSigningCertificate(value *string)() {
    m.signingCertificate = value
}
