package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateUserBindingCollectionResponse 
type X509CertificateUserBindingCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []X509CertificateUserBindingable
}
// NewX509CertificateUserBindingCollectionResponse instantiates a new X509CertificateUserBindingCollectionResponse and sets the default values.
func NewX509CertificateUserBindingCollectionResponse()(*X509CertificateUserBindingCollectionResponse) {
    m := &X509CertificateUserBindingCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateX509CertificateUserBindingCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateX509CertificateUserBindingCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateUserBindingCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *X509CertificateUserBindingCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateX509CertificateUserBindingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]X509CertificateUserBindingable, len(val))
            for i, v := range val {
                res[i] = v.(X509CertificateUserBindingable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *X509CertificateUserBindingCollectionResponse) GetValue()([]X509CertificateUserBindingable) {
    return m.value
}
// Serialize serializes information the current object
func (m *X509CertificateUserBindingCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *X509CertificateUserBindingCollectionResponse) SetValue(value []X509CertificateUserBindingable)() {
    m.value = value
}
