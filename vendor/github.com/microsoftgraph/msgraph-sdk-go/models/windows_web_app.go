package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsWebApp 
type WindowsWebApp struct {
    MobileApp
    // Indicates the Windows web app URL. Example: 'https://www.contoso.com'
    appUrl *string
}
// NewWindowsWebApp instantiates a new WindowsWebApp and sets the default values.
func NewWindowsWebApp()(*WindowsWebApp) {
    m := &WindowsWebApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsWebApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsWebAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsWebAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsWebApp(), nil
}
// GetAppUrl gets the appUrl property value. Indicates the Windows web app URL. Example: 'https://www.contoso.com'
func (m *WindowsWebApp) GetAppUrl()(*string) {
    return m.appUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsWebApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["appUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppUrl(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsWebApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appUrl", m.GetAppUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppUrl sets the appUrl property value. Indicates the Windows web app URL. Example: 'https://www.contoso.com'
func (m *WindowsWebApp) SetAppUrl(value *string)() {
    m.appUrl = value
}
