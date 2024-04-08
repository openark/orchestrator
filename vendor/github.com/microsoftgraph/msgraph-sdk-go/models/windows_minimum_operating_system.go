package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsMinimumOperatingSystem the minimum operating system required for a Windows mobile app.
type WindowsMinimumOperatingSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Windows version 10.0 or later.
    v10_0 *bool
    // Windows version 8.0 or later.
    v8_0 *bool
    // Windows version 8.1 or later.
    v8_1 *bool
}
// NewWindowsMinimumOperatingSystem instantiates a new windowsMinimumOperatingSystem and sets the default values.
func NewWindowsMinimumOperatingSystem()(*WindowsMinimumOperatingSystem) {
    m := &WindowsMinimumOperatingSystem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsMinimumOperatingSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsMinimumOperatingSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["v10_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV100(val)
        }
        return nil
    }
    res["v8_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV80(val)
        }
        return nil
    }
    res["v8_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV81(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WindowsMinimumOperatingSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetV100 gets the v10_0 property value. Windows version 10.0 or later.
func (m *WindowsMinimumOperatingSystem) GetV100()(*bool) {
    return m.v10_0
}
// GetV80 gets the v8_0 property value. Windows version 8.0 or later.
func (m *WindowsMinimumOperatingSystem) GetV80()(*bool) {
    return m.v8_0
}
// GetV81 gets the v8_1 property value. Windows version 8.1 or later.
func (m *WindowsMinimumOperatingSystem) GetV81()(*bool) {
    return m.v8_1
}
// Serialize serializes information the current object
func (m *WindowsMinimumOperatingSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_0", m.GetV100())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_0", m.GetV80())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_1", m.GetV81())
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
func (m *WindowsMinimumOperatingSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WindowsMinimumOperatingSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetV100 sets the v10_0 property value. Windows version 10.0 or later.
func (m *WindowsMinimumOperatingSystem) SetV100(value *bool)() {
    m.v10_0 = value
}
// SetV80 sets the v8_0 property value. Windows version 8.0 or later.
func (m *WindowsMinimumOperatingSystem) SetV80(value *bool)() {
    m.v8_0 = value
}
// SetV81 sets the v8_1 property value. Windows version 8.1 or later.
func (m *WindowsMinimumOperatingSystem) SetV81(value *bool)() {
    m.v8_1 = value
}
