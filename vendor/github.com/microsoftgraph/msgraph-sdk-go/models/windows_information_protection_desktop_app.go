package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionDesktopApp 
type WindowsInformationProtectionDesktopApp struct {
    WindowsInformationProtectionApp
    // The binary name.
    binaryName *string
    // The high binary version.
    binaryVersionHigh *string
    // The lower binary version.
    binaryVersionLow *string
}
// NewWindowsInformationProtectionDesktopApp instantiates a new WindowsInformationProtectionDesktopApp and sets the default values.
func NewWindowsInformationProtectionDesktopApp()(*WindowsInformationProtectionDesktopApp) {
    m := &WindowsInformationProtectionDesktopApp{
        WindowsInformationProtectionApp: *NewWindowsInformationProtectionApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsInformationProtectionDesktopApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsInformationProtectionDesktopAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionDesktopAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionDesktopApp(), nil
}
// GetBinaryName gets the binaryName property value. The binary name.
func (m *WindowsInformationProtectionDesktopApp) GetBinaryName()(*string) {
    return m.binaryName
}
// GetBinaryVersionHigh gets the binaryVersionHigh property value. The high binary version.
func (m *WindowsInformationProtectionDesktopApp) GetBinaryVersionHigh()(*string) {
    return m.binaryVersionHigh
}
// GetBinaryVersionLow gets the binaryVersionLow property value. The lower binary version.
func (m *WindowsInformationProtectionDesktopApp) GetBinaryVersionLow()(*string) {
    return m.binaryVersionLow
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionDesktopApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsInformationProtectionApp.GetFieldDeserializers()
    res["binaryName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBinaryName(val)
        }
        return nil
    }
    res["binaryVersionHigh"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBinaryVersionHigh(val)
        }
        return nil
    }
    res["binaryVersionLow"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBinaryVersionLow(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionDesktopApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsInformationProtectionApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("binaryName", m.GetBinaryName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("binaryVersionHigh", m.GetBinaryVersionHigh())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("binaryVersionLow", m.GetBinaryVersionLow())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBinaryName sets the binaryName property value. The binary name.
func (m *WindowsInformationProtectionDesktopApp) SetBinaryName(value *string)() {
    m.binaryName = value
}
// SetBinaryVersionHigh sets the binaryVersionHigh property value. The high binary version.
func (m *WindowsInformationProtectionDesktopApp) SetBinaryVersionHigh(value *string)() {
    m.binaryVersionHigh = value
}
// SetBinaryVersionLow sets the binaryVersionLow property value. The lower binary version.
func (m *WindowsInformationProtectionDesktopApp) SetBinaryVersionLow(value *string)() {
    m.binaryVersionLow = value
}
