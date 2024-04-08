package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateAuthenticationMethodConfigurationCollectionResponse 
type X509CertificateAuthenticationMethodConfigurationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []X509CertificateAuthenticationMethodConfigurationable
}
// NewX509CertificateAuthenticationMethodConfigurationCollectionResponse instantiates a new X509CertificateAuthenticationMethodConfigurationCollectionResponse and sets the default values.
func NewX509CertificateAuthenticationMethodConfigurationCollectionResponse()(*X509CertificateAuthenticationMethodConfigurationCollectionResponse) {
    m := &X509CertificateAuthenticationMethodConfigurationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateX509CertificateAuthenticationMethodConfigurationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateX509CertificateAuthenticationMethodConfigurationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateAuthenticationMethodConfigurationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *X509CertificateAuthenticationMethodConfigurationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateX509CertificateAuthenticationMethodConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]X509CertificateAuthenticationMethodConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(X509CertificateAuthenticationMethodConfigurationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *X509CertificateAuthenticationMethodConfigurationCollectionResponse) GetValue()([]X509CertificateAuthenticationMethodConfigurationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *X509CertificateAuthenticationMethodConfigurationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *X509CertificateAuthenticationMethodConfigurationCollectionResponse) SetValue(value []X509CertificateAuthenticationMethodConfigurationable)() {
    m.value = value
}
