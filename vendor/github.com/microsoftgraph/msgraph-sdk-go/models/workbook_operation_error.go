package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookOperationError 
type WorkbookOperationError struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The error code.
    code *string
    // The innerError property
    innerError WorkbookOperationErrorable
    // The error message.
    message *string
    // The OdataType property
    odataType *string
}
// NewWorkbookOperationError instantiates a new workbookOperationError and sets the default values.
func NewWorkbookOperationError()(*WorkbookOperationError) {
    m := &WorkbookOperationError{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkbookOperationErrorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookOperationErrorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookOperationError(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookOperationError) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The error code.
func (m *WorkbookOperationError) GetCode()(*string) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookOperationError) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["innerError"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookOperationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInnerError(val.(WorkbookOperationErrorable))
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
    return res
}
// GetInnerError gets the innerError property value. The innerError property
func (m *WorkbookOperationError) GetInnerError()(WorkbookOperationErrorable) {
    return m.innerError
}
// GetMessage gets the message property value. The error message.
func (m *WorkbookOperationError) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkbookOperationError) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *WorkbookOperationError) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("code", m.GetCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("innerError", m.GetInnerError())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookOperationError) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The error code.
func (m *WorkbookOperationError) SetCode(value *string)() {
    m.code = value
}
// SetInnerError sets the innerError property value. The innerError property
func (m *WorkbookOperationError) SetInnerError(value WorkbookOperationErrorable)() {
    m.innerError = value
}
// SetMessage sets the message property value. The error message.
func (m *WorkbookOperationError) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookOperationError) SetOdataType(value *string)() {
    m.odataType = value
}
