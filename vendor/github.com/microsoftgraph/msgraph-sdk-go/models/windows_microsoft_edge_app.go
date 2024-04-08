package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsMicrosoftEdgeApp 
type WindowsMicrosoftEdgeApp struct {
    MobileApp
    // The enum to specify the channels for Microsoft Edge apps.
    channel *MicrosoftEdgeChannel
    // The language locale to use when the Edge app displays text to the user.
    displayLanguageLocale *string
}
// NewWindowsMicrosoftEdgeApp instantiates a new WindowsMicrosoftEdgeApp and sets the default values.
func NewWindowsMicrosoftEdgeApp()(*WindowsMicrosoftEdgeApp) {
    m := &WindowsMicrosoftEdgeApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.windowsMicrosoftEdgeApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsMicrosoftEdgeAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsMicrosoftEdgeAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsMicrosoftEdgeApp(), nil
}
// GetChannel gets the channel property value. The enum to specify the channels for Microsoft Edge apps.
func (m *WindowsMicrosoftEdgeApp) GetChannel()(*MicrosoftEdgeChannel) {
    return m.channel
}
// GetDisplayLanguageLocale gets the displayLanguageLocale property value. The language locale to use when the Edge app displays text to the user.
func (m *WindowsMicrosoftEdgeApp) GetDisplayLanguageLocale()(*string) {
    return m.displayLanguageLocale
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsMicrosoftEdgeApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["channel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMicrosoftEdgeChannel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannel(val.(*MicrosoftEdgeChannel))
        }
        return nil
    }
    res["displayLanguageLocale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayLanguageLocale(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsMicrosoftEdgeApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChannel() != nil {
        cast := (*m.GetChannel()).String()
        err = writer.WriteStringValue("channel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayLanguageLocale", m.GetDisplayLanguageLocale())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChannel sets the channel property value. The enum to specify the channels for Microsoft Edge apps.
func (m *WindowsMicrosoftEdgeApp) SetChannel(value *MicrosoftEdgeChannel)() {
    m.channel = value
}
// SetDisplayLanguageLocale sets the displayLanguageLocale property value. The language locale to use when the Edge app displays text to the user.
func (m *WindowsMicrosoftEdgeApp) SetDisplayLanguageLocale(value *string)() {
    m.displayLanguageLocale = value
}
