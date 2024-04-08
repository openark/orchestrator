package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppleManagedIdentityProvider 
type AppleManagedIdentityProvider struct {
    IdentityProviderBase
    // The certificate data, which is a long string of text from the certificate. Can be null.
    certificateData *string
    // The Apple developer identifier. Required.
    developerId *string
    // The Apple key identifier. Required.
    keyId *string
    // The Apple service identifier. Required.
    serviceId *string
}
// NewAppleManagedIdentityProvider instantiates a new AppleManagedIdentityProvider and sets the default values.
func NewAppleManagedIdentityProvider()(*AppleManagedIdentityProvider) {
    m := &AppleManagedIdentityProvider{
        IdentityProviderBase: *NewIdentityProviderBase(),
    }
    odataTypeValue := "#microsoft.graph.appleManagedIdentityProvider"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAppleManagedIdentityProviderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAppleManagedIdentityProviderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAppleManagedIdentityProvider(), nil
}
// GetCertificateData gets the certificateData property value. The certificate data, which is a long string of text from the certificate. Can be null.
func (m *AppleManagedIdentityProvider) GetCertificateData()(*string) {
    return m.certificateData
}
// GetDeveloperId gets the developerId property value. The Apple developer identifier. Required.
func (m *AppleManagedIdentityProvider) GetDeveloperId()(*string) {
    return m.developerId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AppleManagedIdentityProvider) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentityProviderBase.GetFieldDeserializers()
    res["certificateData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateData(val)
        }
        return nil
    }
    res["developerId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeveloperId(val)
        }
        return nil
    }
    res["keyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyId(val)
        }
        return nil
    }
    res["serviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceId(val)
        }
        return nil
    }
    return res
}
// GetKeyId gets the keyId property value. The Apple key identifier. Required.
func (m *AppleManagedIdentityProvider) GetKeyId()(*string) {
    return m.keyId
}
// GetServiceId gets the serviceId property value. The Apple service identifier. Required.
func (m *AppleManagedIdentityProvider) GetServiceId()(*string) {
    return m.serviceId
}
// Serialize serializes information the current object
func (m *AppleManagedIdentityProvider) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentityProviderBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("certificateData", m.GetCertificateData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("developerId", m.GetDeveloperId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyId", m.GetKeyId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceId", m.GetServiceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCertificateData sets the certificateData property value. The certificate data, which is a long string of text from the certificate. Can be null.
func (m *AppleManagedIdentityProvider) SetCertificateData(value *string)() {
    m.certificateData = value
}
// SetDeveloperId sets the developerId property value. The Apple developer identifier. Required.
func (m *AppleManagedIdentityProvider) SetDeveloperId(value *string)() {
    m.developerId = value
}
// SetKeyId sets the keyId property value. The Apple key identifier. Required.
func (m *AppleManagedIdentityProvider) SetKeyId(value *string)() {
    m.keyId = value
}
// SetServiceId sets the serviceId property value. The Apple service identifier. Required.
func (m *AppleManagedIdentityProvider) SetServiceId(value *string)() {
    m.serviceId = value
}
