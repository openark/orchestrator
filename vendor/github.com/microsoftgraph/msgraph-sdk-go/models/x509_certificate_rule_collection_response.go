package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// X509CertificateRuleCollectionResponse 
type X509CertificateRuleCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []X509CertificateRuleable
}
// NewX509CertificateRuleCollectionResponse instantiates a new X509CertificateRuleCollectionResponse and sets the default values.
func NewX509CertificateRuleCollectionResponse()(*X509CertificateRuleCollectionResponse) {
    m := &X509CertificateRuleCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateX509CertificateRuleCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateX509CertificateRuleCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewX509CertificateRuleCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *X509CertificateRuleCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateX509CertificateRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]X509CertificateRuleable, len(val))
            for i, v := range val {
                res[i] = v.(X509CertificateRuleable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *X509CertificateRuleCollectionResponse) GetValue()([]X509CertificateRuleable) {
    return m.value
}
// Serialize serializes information the current object
func (m *X509CertificateRuleCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *X509CertificateRuleCollectionResponse) SetValue(value []X509CertificateRuleable)() {
    m.value = value
}
