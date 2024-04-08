package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10CustomConfiguration 
type Windows10CustomConfiguration struct {
    DeviceConfiguration
    // OMA settings. This collection can contain a maximum of 1000 elements.
    omaSettings []OmaSettingable
}
// NewWindows10CustomConfiguration instantiates a new Windows10CustomConfiguration and sets the default values.
func NewWindows10CustomConfiguration()(*Windows10CustomConfiguration) {
    m := &Windows10CustomConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10CustomConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10CustomConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10CustomConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10CustomConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10CustomConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *Windows10CustomConfiguration) GetOmaSettings()([]OmaSettingable) {
    return m.omaSettings
}
// Serialize serializes information the current object
func (m *Windows10CustomConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *Windows10CustomConfiguration) SetOmaSettings(value []OmaSettingable)() {
    m.omaSettings = value
}
