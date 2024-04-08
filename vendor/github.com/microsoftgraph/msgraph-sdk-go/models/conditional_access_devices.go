package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessDevices 
type ConditionalAccessDevices struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them.
    deviceFilter ConditionalAccessFilterable
    // The OdataType property
    odataType *string
}
// NewConditionalAccessDevices instantiates a new conditionalAccessDevices and sets the default values.
func NewConditionalAccessDevices()(*ConditionalAccessDevices) {
    m := &ConditionalAccessDevices{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessDevicesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessDevicesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessDevices(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessDevices) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceFilter gets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them.
func (m *ConditionalAccessDevices) GetDeviceFilter()(ConditionalAccessFilterable) {
    return m.deviceFilter
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessDevices) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessFilterFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceFilter(val.(ConditionalAccessFilterable))
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
func (m *ConditionalAccessDevices) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ConditionalAccessDevices) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("deviceFilter", m.GetDeviceFilter())
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
func (m *ConditionalAccessDevices) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceFilter sets the deviceFilter property value. Filter that defines the dynamic-device-syntax rule to include/exclude devices. A filter can use device properties (such as extension attributes) to include/exclude them.
func (m *ConditionalAccessDevices) SetDeviceFilter(value ConditionalAccessFilterable)() {
    m.deviceFilter = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessDevices) SetOdataType(value *string)() {
    m.odataType = value
}
