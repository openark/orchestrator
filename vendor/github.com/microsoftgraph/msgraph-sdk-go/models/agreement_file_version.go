package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgreementFileVersion 
type AgreementFileVersion struct {
    AgreementFileProperties
}
// NewAgreementFileVersion instantiates a new agreementFileVersion and sets the default values.
func NewAgreementFileVersion()(*AgreementFileVersion) {
    m := &AgreementFileVersion{
        AgreementFileProperties: *NewAgreementFileProperties(),
    }
    return m
}
// CreateAgreementFileVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFileVersionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgreementFileVersion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFileVersion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AgreementFileProperties.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AgreementFileVersion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AgreementFileProperties.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
