package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateAuthenticationMethodConfiguration 
type X509CertificateAuthenticationMethodConfiguration struct {
    AuthenticationMethodConfiguration
    // Defines strong authentication configurations. This configuration includes the default authentication mode and the different rules for strong authentication bindings.
    authenticationModeConfiguration X509CertificateAuthenticationModeConfigurationable
    // Defines fields in the X.509 certificate that map to attributes of the Azure AD user object in order to bind the certificate to the user. The priority of the object determines the order in which the binding is carried out. The first binding that matches will be used and the rest ignored.
    certificateUserBindings []X509CertificateUserBindingable
    // A collection of groups that are enabled to use the authentication method.
    includeTargets []AuthenticationMethodTargetable
}
// NewX509CertificateAuthenticationMethodConfiguration instantiates a new X509CertificateAuthenticationMethodConfiguration and sets the default values.
func NewX509CertificateAuthenticationMethodConfiguration()(*X509CertificateAuthenticationMethodConfiguration) {
    m := &X509CertificateAuthenticationMethodConfiguration{
        AuthenticationMethodConfiguration: *NewAuthenticationMethodConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.x509CertificateAuthenticationMethodConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateX509CertificateAuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateX509CertificateAuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateAuthenticationMethodConfiguration(), nil
}
// GetAuthenticationModeConfiguration gets the authenticationModeConfiguration property value. Defines strong authentication configurations. This configuration includes the default authentication mode and the different rules for strong authentication bindings.
func (m *X509CertificateAuthenticationMethodConfiguration) GetAuthenticationModeConfiguration()(X509CertificateAuthenticationModeConfigurationable) {
    return m.authenticationModeConfiguration
}
// GetCertificateUserBindings gets the certificateUserBindings property value. Defines fields in the X.509 certificate that map to attributes of the Azure AD user object in order to bind the certificate to the user. The priority of the object determines the order in which the binding is carried out. The first binding that matches will be used and the rest ignored.
func (m *X509CertificateAuthenticationMethodConfiguration) GetCertificateUserBindings()([]X509CertificateUserBindingable) {
    return m.certificateUserBindings
}
// GetFieldDeserializers the deserialization information for the current model
func (m *X509CertificateAuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodConfiguration.GetFieldDeserializers()
    res["authenticationModeConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateX509CertificateAuthenticationModeConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationModeConfiguration(val.(X509CertificateAuthenticationModeConfigurationable))
        }
        return nil
    }
    res["certificateUserBindings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateX509CertificateUserBindingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]X509CertificateUserBindingable, len(val))
            for i, v := range val {
                res[i] = v.(X509CertificateUserBindingable)
            }
            m.SetCertificateUserBindings(res)
        }
        return nil
    }
    res["includeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodTargetable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodTargetable)
            }
            m.SetIncludeTargets(res)
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. A collection of groups that are enabled to use the authentication method.
func (m *X509CertificateAuthenticationMethodConfiguration) GetIncludeTargets()([]AuthenticationMethodTargetable) {
    return m.includeTargets
}
// Serialize serializes information the current object
func (m *X509CertificateAuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("authenticationModeConfiguration", m.GetAuthenticationModeConfiguration())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateUserBindings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateUserBindings()))
        for i, v := range m.GetCertificateUserBindings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certificateUserBindings", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuthenticationModeConfiguration sets the authenticationModeConfiguration property value. Defines strong authentication configurations. This configuration includes the default authentication mode and the different rules for strong authentication bindings.
func (m *X509CertificateAuthenticationMethodConfiguration) SetAuthenticationModeConfiguration(value X509CertificateAuthenticationModeConfigurationable)() {
    m.authenticationModeConfiguration = value
}
// SetCertificateUserBindings sets the certificateUserBindings property value. Defines fields in the X.509 certificate that map to attributes of the Azure AD user object in order to bind the certificate to the user. The priority of the object determines the order in which the binding is carried out. The first binding that matches will be used and the rest ignored.
func (m *X509CertificateAuthenticationMethodConfiguration) SetCertificateUserBindings(value []X509CertificateUserBindingable)() {
    m.certificateUserBindings = value
}
// SetIncludeTargets sets the includeTargets property value. A collection of groups that are enabled to use the authentication method.
func (m *X509CertificateAuthenticationMethodConfiguration) SetIncludeTargets(value []AuthenticationMethodTargetable)() {
    m.includeTargets = value
}
