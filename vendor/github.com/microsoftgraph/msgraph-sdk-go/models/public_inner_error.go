package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PublicInnerError 
type PublicInnerError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error code.
    code *string
    // A collection of error details.
    details []PublicErrorDetailable
    // The error message.
    message *string
    // The OdataType property
    odataType *string
    // The target of the error.
    target *string
}
// NewPublicInnerError instantiates a new publicInnerError and sets the default values.
func NewPublicInnerError()(*PublicInnerError) {
    m := &PublicInnerError{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePublicInnerErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublicInnerErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPublicInnerError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PublicInnerError) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The error code.
func (m *PublicInnerError) GetCode()(*string) {
    return m.code
}
// GetDetails gets the details property value. A collection of error details.
func (m *PublicInnerError) GetDetails()([]PublicErrorDetailable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublicInnerError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePublicErrorDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PublicErrorDetailable, len(val))
            for i, v := range val {
                res[i] = v.(PublicErrorDetailable)
            }
            m.SetDetails(res)
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
func (m *PublicInnerError) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PublicInnerError) GetOdataType()(*string) {
    return m.odataType
}
// GetTarget gets the target property value. The target of the error.
func (m *PublicInnerError) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *PublicInnerError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    if m.GetDetails() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDetails()))
        for i, v := range m.GetDetails() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("details", cast)
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
func (m *PublicInnerError) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The error code.
func (m *PublicInnerError) SetCode(value *string)() {
    m.code = value
}
// SetDetails sets the details property value. A collection of error details.
func (m *PublicInnerError) SetDetails(value []PublicErrorDetailable)() {
    m.details = value
}
// SetMessage sets the message property value. The error message.
func (m *PublicInnerError) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PublicInnerError) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTarget sets the target property value. The target of the error.
func (m *PublicInnerError) SetTarget(value *string)() {
    m.target = value
}
