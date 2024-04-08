package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDefenderAdvancedThreatProtectionConfiguration 
type WindowsDefenderAdvancedThreatProtectionConfiguration struct {
    DeviceConfiguration
    // Windows Defender AdvancedThreatProtection 'Allow Sample Sharing' Rule
    allowSampleSharing *bool
    // Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
    enableExpeditedTelemetryReporting *bool
}
// NewWindowsDefenderAdvancedThreatProtectionConfiguration instantiates a new WindowsDefenderAdvancedThreatProtectionConfiguration and sets the default values.
func NewWindowsDefenderAdvancedThreatProtectionConfiguration()(*WindowsDefenderAdvancedThreatProtectionConfiguration) {
    m := &WindowsDefenderAdvancedThreatProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsDefenderAdvancedThreatProtectionConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsDefenderAdvancedThreatProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDefenderAdvancedThreatProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDefenderAdvancedThreatProtectionConfiguration(), nil
}
// GetAllowSampleSharing gets the allowSampleSharing property value. Windows Defender AdvancedThreatProtection 'Allow Sample Sharing' Rule
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetAllowSampleSharing()(*bool) {
    return m.allowSampleSharing
}
// GetEnableExpeditedTelemetryReporting gets the enableExpeditedTelemetryReporting property value. Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetEnableExpeditedTelemetryReporting()(*bool) {
    return m.enableExpeditedTelemetryReporting
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["allowSampleSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowSampleSharing(val)
        }
        return nil
    }
    res["enableExpeditedTelemetryReporting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableExpeditedTelemetryReporting(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowSampleSharing", m.GetAllowSampleSharing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enableExpeditedTelemetryReporting", m.GetEnableExpeditedTelemetryReporting())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowSampleSharing sets the allowSampleSharing property value. Windows Defender AdvancedThreatProtection 'Allow Sample Sharing' Rule
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetAllowSampleSharing(value *bool)() {
    m.allowSampleSharing = value
}
// SetEnableExpeditedTelemetryReporting sets the enableExpeditedTelemetryReporting property value. Expedite Windows Defender Advanced Threat Protection telemetry reporting frequency.
func (m *WindowsDefenderAdvancedThreatProtectionConfiguration) SetEnableExpeditedTelemetryReporting(value *bool)() {
    m.enableExpeditedTelemetryReporting = value
}
