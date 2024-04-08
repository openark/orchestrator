package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResultInfo 
type ResultInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The result code.
    code *int32
    // The message.
    message *string
    // The OdataType property
    odataType *string
    // The result sub-code.
    subcode *int32
}
// NewResultInfo instantiates a new resultInfo and sets the default values.
func NewResultInfo()(*ResultInfo) {
    m := &ResultInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResultInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResultInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResultInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResultInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCode gets the code property value. The result code.
func (m *ResultInfo) GetCode()(*int32) {
    return m.code
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResultInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["code"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
    res["subcode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubcode(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The message.
func (m *ResultInfo) GetMessage()(*string) {
    return m.message
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ResultInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetSubcode gets the subcode property value. The result sub-code.
func (m *ResultInfo) GetSubcode()(*int32) {
    return m.subcode
}
// Serialize serializes information the current object
func (m *ResultInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("code", m.GetCode())
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
        err := writer.WriteInt32Value("subcode", m.GetSubcode())
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
func (m *ResultInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCode sets the code property value. The result code.
func (m *ResultInfo) SetCode(value *int32)() {
    m.code = value
}
// SetMessage sets the message property value. The message.
func (m *ResultInfo) SetMessage(value *string)() {
    m.message = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ResultInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSubcode sets the subcode property value. The result sub-code.
func (m *ResultInfo) SetSubcode(value *int32)() {
    m.subcode = value
}
