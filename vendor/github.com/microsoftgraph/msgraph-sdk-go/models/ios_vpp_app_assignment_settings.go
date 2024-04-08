package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppAppAssignmentSettings 
type IosVppAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // Whether or not to use device licensing.
    useDeviceLicensing *bool
    // The VPN Configuration Id to apply for this app.
    vpnConfigurationId *string
}
// NewIosVppAppAssignmentSettings instantiates a new IosVppAppAssignmentSettings and sets the default values.
func NewIosVppAppAssignmentSettings()(*IosVppAppAssignmentSettings) {
    m := &IosVppAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.iosVppAppAssignmentSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosVppAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosVppAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosVppAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosVppAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
    res["useDeviceLicensing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUseDeviceLicensing(val)
        }
        return nil
    }
    res["vpnConfigurationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVpnConfigurationId(val)
        }
        return nil
    }
    return res
}
// GetUseDeviceLicensing gets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *IosVppAppAssignmentSettings) GetUseDeviceLicensing()(*bool) {
    return m.useDeviceLicensing
}
// GetVpnConfigurationId gets the vpnConfigurationId property value. The VPN Configuration Id to apply for this app.
func (m *IosVppAppAssignmentSettings) GetVpnConfigurationId()(*string) {
    return m.vpnConfigurationId
}
// Serialize serializes information the current object
func (m *IosVppAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("useDeviceLicensing", m.GetUseDeviceLicensing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("vpnConfigurationId", m.GetVpnConfigurationId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUseDeviceLicensing sets the useDeviceLicensing property value. Whether or not to use device licensing.
func (m *IosVppAppAssignmentSettings) SetUseDeviceLicensing(value *bool)() {
    m.useDeviceLicensing = value
}
// SetVpnConfigurationId sets the vpnConfigurationId property value. The VPN Configuration Id to apply for this app.
func (m *IosVppAppAssignmentSettings) SetVpnConfigurationId(value *string)() {
    m.vpnConfigurationId = value
}
