package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosStoreAppAssignmentSettings 
type IosStoreAppAssignmentSettings struct {
    MobileAppAssignmentSettings
    // The VPN Configuration Id to apply for this app.
    vpnConfigurationId *string
}
// NewIosStoreAppAssignmentSettings instantiates a new IosStoreAppAssignmentSettings and sets the default values.
func NewIosStoreAppAssignmentSettings()(*IosStoreAppAssignmentSettings) {
    m := &IosStoreAppAssignmentSettings{
        MobileAppAssignmentSettings: *NewMobileAppAssignmentSettings(),
    }
    odataTypeValue := "#microsoft.graph.iosStoreAppAssignmentSettings"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosStoreAppAssignmentSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosStoreAppAssignmentSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosStoreAppAssignmentSettings(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosStoreAppAssignmentSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileAppAssignmentSettings.GetFieldDeserializers()
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
// GetVpnConfigurationId gets the vpnConfigurationId property value. The VPN Configuration Id to apply for this app.
func (m *IosStoreAppAssignmentSettings) GetVpnConfigurationId()(*string) {
    return m.vpnConfigurationId
}
// Serialize serializes information the current object
func (m *IosStoreAppAssignmentSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileAppAssignmentSettings.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("vpnConfigurationId", m.GetVpnConfigurationId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetVpnConfigurationId sets the vpnConfigurationId property value. The VPN Configuration Id to apply for this app.
func (m *IosStoreAppAssignmentSettings) SetVpnConfigurationId(value *string)() {
    m.vpnConfigurationId = value
}
