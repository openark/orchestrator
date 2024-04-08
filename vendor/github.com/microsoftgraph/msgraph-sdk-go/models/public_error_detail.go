package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PublicErrorDetail 
type PublicErrorDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error code.
    code *string
    // The error message.
    message *string
    // The OdataType property
    odataType *string
    // The target of the error.
    target *string
}
// NewPublicErrorDetail instantiates a new publicErrorDetail and sets the default values.
func NewPublicErrorDetail()(*PublicErrorDetail) {
    m := &PublicErrorDetail{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePublicErrorDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublicErrorDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicErrorDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PublicErrorDetail) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The error code.
func (m *PublicErrorDetail) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublicErrorDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCode(val)
        }
        return nil
    }
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val)
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
    res["target"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The error message.
func (m *PublicErrorDetail) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PublicErrorDetail) GetOdataType()(*string) {
    return m.odataType
}
// GetTarget gets the target property value. The target of the error.
func (m *PublicErrorDetail) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *PublicErrorDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("message", m.GetMessage())
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
        err := writer.WriteStringValue("target", m.GetTarget())
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
func (m *PublicErrorDetail) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The error code.
func (m *PublicErrorDetail) SetCode(value *string)() {
    m.code = value
}
// SetMessage sets the message property value. The error message.
func (m *PublicErrorDetail) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PublicErrorDetail) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTarget sets the target property value. The target of the error.
func (m *PublicErrorDetail) SetTarget(value *string)() {
    m.target = value
}
