package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosCustomConfiguration 
type IosCustomConfiguration struct {
    DeviceConfiguration
    // Payload. (UTF8 encoded byte array)
    payload []byte
    // Payload file name (.mobileconfig
    payloadFileName *string
    // Name that is displayed to the user.
    payloadName *string
}
// NewIosCustomConfiguration instantiates a new IosCustomConfiguration and sets the default values.
func NewIosCustomConfiguration()(*IosCustomConfiguration) {
    m := &IosCustomConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.iosCustomConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosCustomConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosCustomConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosCustomConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosCustomConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["payload"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayload(val)
        }
        return nil
    }
    res["payloadFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadFileName(val)
        }
        return nil
    }
    res["payloadName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPayloadName(val)
        }
        return nil
    }
    return res
}
// GetPayload gets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosCustomConfiguration) GetPayload()([]byte) {
    return m.payload
}
// GetPayloadFileName gets the payloadFileName property value. Payload file name (.mobileconfig
func (m *IosCustomConfiguration) GetPayloadFileName()(*string) {
    return m.payloadFileName
}
// GetPayloadName gets the payloadName property value. Name that is displayed to the user.
func (m *IosCustomConfiguration) GetPayloadName()(*string) {
    return m.payloadName
}
// Serialize serializes information the current object
func (m *IosCustomConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("payload", m.GetPayload())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadFileName", m.GetPayloadFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("payloadName", m.GetPayloadName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPayload sets the payload property value. Payload. (UTF8 encoded byte array)
func (m *IosCustomConfiguration) SetPayload(value []byte)() {
    m.payload = value
}
// SetPayloadFileName sets the payloadFileName property value. Payload file name (.mobileconfig
func (m *IosCustomConfiguration) SetPayloadFileName(value *string)() {
    m.payloadFileName = value
}
// SetPayloadName sets the payloadName property value. Name that is displayed to the user.
func (m *IosCustomConfiguration) SetPayloadName(value *string)() {
    m.payloadName = value
}
