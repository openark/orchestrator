package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsMobileMSI 
type WindowsMobileMSI struct {
    MobileLobApp
    // The command line.
    commandLine *string
    // A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature.
    ignoreVersionDetection *bool
    // The product code.
    productCode *string
    // The product version of Windows Mobile MSI Line of Business (LoB) app.
    productVersion *string
}
// NewWindowsMobileMSI instantiates a new WindowsMobileMSI and sets the default values.
func NewWindowsMobileMSI()(*WindowsMobileMSI) {
    m := &WindowsMobileMSI{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsMobileMSI"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsMobileMSIFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsMobileMSIFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsMobileMSI(), nil
}
// GetCommandLine gets the commandLine property value. The command line.
func (m *WindowsMobileMSI) GetCommandLine()(*string) {
    return m.commandLine
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsMobileMSI) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["commandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommandLine(val)
        }
        return nil
    }
    res["ignoreVersionDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreVersionDetection(val)
        }
        return nil
    }
    res["productCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductCode(val)
        }
        return nil
    }
    res["productVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductVersion(val)
        }
        return nil
    }
    return res
}
// GetIgnoreVersionDetection gets the ignoreVersionDetection property value. A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature.
func (m *WindowsMobileMSI) GetIgnoreVersionDetection()(*bool) {
    return m.ignoreVersionDetection
}
// GetProductCode gets the productCode property value. The product code.
func (m *WindowsMobileMSI) GetProductCode()(*string) {
    return m.productCode
}
// GetProductVersion gets the productVersion property value. The product version of Windows Mobile MSI Line of Business (LoB) app.
func (m *WindowsMobileMSI) GetProductVersion()(*string) {
    return m.productVersion
}
// Serialize serializes information the current object
func (m *WindowsMobileMSI) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("commandLine", m.GetCommandLine())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ignoreVersionDetection", m.GetIgnoreVersionDetection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productCode", m.GetProductCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productVersion", m.GetProductVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCommandLine sets the commandLine property value. The command line.
func (m *WindowsMobileMSI) SetCommandLine(value *string)() {
    m.commandLine = value
}
// SetIgnoreVersionDetection sets the ignoreVersionDetection property value. A boolean to control whether the app's version will be used to detect the app after it is installed on a device. Set this to true for Windows Mobile MSI Line of Business (LoB) apps that use a self update feature.
func (m *WindowsMobileMSI) SetIgnoreVersionDetection(value *bool)() {
    m.ignoreVersionDetection = value
}
// SetProductCode sets the productCode property value. The product code.
func (m *WindowsMobileMSI) SetProductCode(value *string)() {
    m.productCode = value
}
// SetProductVersion sets the productVersion property value. The product version of Windows Mobile MSI Line of Business (LoB) app.
func (m *WindowsMobileMSI) SetProductVersion(value *string)() {
    m.productVersion = value
}
