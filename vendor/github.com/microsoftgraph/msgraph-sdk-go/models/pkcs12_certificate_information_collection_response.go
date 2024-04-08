package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Pkcs12CertificateInformationCollectionResponse 
type Pkcs12CertificateInformationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []Pkcs12CertificateInformationable
}
// NewPkcs12CertificateInformationCollectionResponse instantiates a new Pkcs12CertificateInformationCollectionResponse and sets the default values.
func NewPkcs12CertificateInformationCollectionResponse()(*Pkcs12CertificateInformationCollectionResponse) {
    m := &Pkcs12CertificateInformationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreatePkcs12CertificateInformationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePkcs12CertificateInformationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPkcs12CertificateInformationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Pkcs12CertificateInformationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePkcs12CertificateInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Pkcs12CertificateInformationable, len(val))
            for i, v := range val {
                res[i] = v.(Pkcs12CertificateInformationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *Pkcs12CertificateInformationCollectionResponse) GetValue()([]Pkcs12CertificateInformationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *Pkcs12CertificateInformationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Pkcs12CertificateInformationCollectionResponse) SetValue(value []Pkcs12CertificateInformationable)() {
    m.value = value
}
