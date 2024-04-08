package odataerrors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MainError 
type MainError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The code property
    code *string
    // The details property
    details []ErrorDetailsable
    // The innererror property
    innererror InnerErrorable
    // The message property
    message *string
    // The target property
    target *string
}
// NewMainError instantiates a new MainError and sets the default values.
func NewMainError()(*MainError) {
    m := &MainError{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMainErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMainErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMainError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MainError) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The code property
func (m *MainError) GetCode()(*string) {
    return m.code
}
// GetDetails gets the details property value. The details property
func (m *MainError) GetDetails()([]ErrorDetailsable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MainError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetCollectionOfObjectValues(CreateErrorDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ErrorDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(ErrorDetailsable)
            }
            m.SetDetails(res)
        }
        return nil
    }
    res["innererror"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInnerErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnererror(val.(InnerErrorable))
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
// GetInnererror gets the innererror property value. The innererror property
func (m *MainError) GetInnererror()(InnerErrorable) {
    return m.innererror
}
// GetMessage gets the message property value. The message property
func (m *MainError) GetMessage()(*string) {
    return m.message
}
// GetTarget gets the target property value. The target property
func (m *MainError) GetTarget()(*string) {
    return m.target
}
// Serialize serializes information the current object
func (m *MainError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteObjectValue("innererror", m.GetInnererror())
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
func (m *MainError) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The code property
func (m *MainError) SetCode(value *string)() {
    m.code = value
}
// SetDetails sets the details property value. The details property
func (m *MainError) SetDetails(value []ErrorDetailsable)() {
    m.details = value
}
// SetInnererror sets the innererror property value. The innererror property
func (m *MainError) SetInnererror(value InnerErrorable)() {
    m.innererror = value
}
// SetMessage sets the message property value. The message property
func (m *MainError) SetMessage(value *string)() {
    m.message = value
}
// SetTarget sets the target property value. The target property
func (m *MainError) SetTarget(value *string)() {
    m.target = value
}
