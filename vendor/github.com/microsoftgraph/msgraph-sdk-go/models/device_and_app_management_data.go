package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAndAppManagementData exported Data
type DeviceAndAppManagementData struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Not yet documented
    content []byte
    // The OdataType property
    odataType *string
}
// NewDeviceAndAppManagementData instantiates a new deviceAndAppManagementData and sets the default values.
func NewDeviceAndAppManagementData()(*DeviceAndAppManagementData) {
    m := &DeviceAndAppManagementData{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceAndAppManagementDataFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAndAppManagementDataFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAndAppManagementData(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceAndAppManagementData) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContent gets the content property value. Not yet documented
func (m *DeviceAndAppManagementData) GetContent()([]byte) {
    return m.content
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAndAppManagementData) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
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
func (m *DeviceAndAppManagementData) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DeviceAndAppManagementData) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteByteArrayValue("content", m.GetContent())
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
func (m *DeviceAndAppManagementData) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContent sets the content property value. Not yet documented
func (m *DeviceAndAppManagementData) SetContent(value []byte)() {
    m.content = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceAndAppManagementData) SetOdataType(value *string)() {
    m.odataType = value
}
