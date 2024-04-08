package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingFloatingPoint 
type OmaSettingFloatingPoint struct {
    OmaSetting
    // Value.
    value *float32
}
// NewOmaSettingFloatingPoint instantiates a new OmaSettingFloatingPoint and sets the default values.
func NewOmaSettingFloatingPoint()(*OmaSettingFloatingPoint) {
    m := &OmaSettingFloatingPoint{
        OmaSetting: *NewOmaSetting(),
    }
    odataTypeValue := "#microsoft.graph.omaSettingFloatingPoint"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOmaSettingFloatingPointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingFloatingPointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingFloatingPoint(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingFloatingPoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OmaSetting.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. Value.
func (m *OmaSettingFloatingPoint) GetValue()(*float32) {
    return m.value
}
// Serialize serializes information the current object
func (m *OmaSettingFloatingPoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OmaSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat32Value("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. Value.
func (m *OmaSettingFloatingPoint) SetValue(value *float32)() {
    m.value = value
}
