package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AgreementFileData 
type AgreementFileData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Data that represents the terms of use PDF document. Read-only.
    data []byte
    // The OdataType property
    odataType *string
}
// NewAgreementFileData instantiates a new agreementFileData and sets the default values.
func NewAgreementFileData()(*AgreementFileData) {
    m := &AgreementFileData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAgreementFileDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAgreementFileDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAgreementFileData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AgreementFileData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetData gets the data property value. Data that represents the terms of use PDF document. Read-only.
func (m *AgreementFileData) GetData()([]byte) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AgreementFileData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AgreementFileData) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AgreementFileData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("data", m.GetData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AgreementFileData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetData sets the data property value. Data that represents the terms of use PDF document. Read-only.
func (m *AgreementFileData) SetData(value []byte)() {
    m.data = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AgreementFileData) SetOdataType(value *string)() {
    m.odataType = value
}
