package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81CustomConfiguration 
type WindowsPhone81CustomConfiguration struct {
    DeviceConfiguration
    // OMA settings. This collection can contain a maximum of 1000 elements.
    omaSettings []OmaSettingable
}
// NewWindowsPhone81CustomConfiguration instantiates a new WindowsPhone81CustomConfiguration and sets the default values.
func NewWindowsPhone81CustomConfiguration()(*WindowsPhone81CustomConfiguration) {
    m := &WindowsPhone81CustomConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81CustomConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsPhone81CustomConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81CustomConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhone81CustomConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81CustomConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["omaSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOmaSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OmaSettingable, len(val))
            for i, v := range val {
                res[i] = v.(OmaSettingable)
            }
            m.SetOmaSettings(res)
        }
        return nil
    }
    return res
}
// GetOmaSettings gets the omaSettings property value. OMA settings. This collection can contain a maximum of 1000 elements.
func (m *WindowsPhone81CustomConfiguration) GetOmaSettings()([]OmaSettingable) {
    return m.omaSettings
}
// Serialize serializes information the current object
func (m *WindowsPhone81CustomConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetOmaSettings() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOmaSettings()))
        for i, v := range m.GetOmaSettings() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("omaSettings", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOmaSettings sets the omaSettings property value. OMA settings. This collection can contain a maximum of 1000 elements.
func (m *WindowsPhone81CustomConfiguration) SetOmaSettings(value []OmaSettingable)() {
    m.omaSettings = value
}
