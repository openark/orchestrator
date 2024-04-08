package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidWorkProfileCompliancePolicy 
type AndroidWorkProfileCompliancePolicy struct {
    DeviceCompliancePolicy
    // Require that devices have enabled device threat protection.
    deviceThreatProtectionEnabled *bool
    // Device threat protection levels for the Device Threat Protection API.
    deviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Minimum Android security patch level.
    minAndroidSecurityPatchLevel *string
    // Maximum Android version.
    osMaximumVersion *string
    // Minimum Android version.
    osMinimumVersion *string
    // Number of days before the password expires. Valid values 1 to 365
    passwordExpirationDays *int32
    // Minimum password length. Valid values 4 to 16
    passwordMinimumLength *int32
    // Minutes of inactivity before a password is required.
    passwordMinutesOfInactivityBeforeLock *int32
    // Number of previous passwords to block. Valid values 1 to 24
    passwordPreviousPasswordBlockCount *int32
    // Require a password to unlock device.
    passwordRequired *bool
    // Android required password type.
    passwordRequiredType *AndroidRequiredPasswordType
    // Devices must not be jailbroken or rooted.
    securityBlockJailbrokenDevices *bool
    // Disable USB debugging on Android devices.
    securityDisableUsbDebugging *bool
    // Require that devices disallow installation of apps from unknown sources.
    securityPreventInstallAppsFromUnknownSources *bool
    // Require the device to pass the Company Portal client app runtime integrity check.
    securityRequireCompanyPortalAppIntegrity *bool
    // Require Google Play Services to be installed and enabled on the device.
    securityRequireGooglePlayServices *bool
    // Require the device to pass the SafetyNet basic integrity check.
    securityRequireSafetyNetAttestationBasicIntegrity *bool
    // Require the device to pass the SafetyNet certified device check.
    securityRequireSafetyNetAttestationCertifiedDevice *bool
    // Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
    securityRequireUpToDateSecurityProviders *bool
    // Require the Android Verify apps feature is turned on.
    securityRequireVerifyApps *bool
    // Require encryption on Android devices.
    storageRequireEncryption *bool
}
// NewAndroidWorkProfileCompliancePolicy instantiates a new AndroidWorkProfileCompliancePolicy and sets the default values.
func NewAndroidWorkProfileCompliancePolicy()(*AndroidWorkProfileCompliancePolicy) {
    m := &AndroidWorkProfileCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.androidWorkProfileCompliancePolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidWorkProfileCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidWorkProfileCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidWorkProfileCompliancePolicy(), nil
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidWorkProfileCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    return m.deviceThreatProtectionEnabled
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *AndroidWorkProfileCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    return m.deviceThreatProtectionRequiredSecurityLevel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidWorkProfileCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["deviceThreatProtectionEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionEnabled(val)
        }
        return nil
    }
    res["deviceThreatProtectionRequiredSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceThreatProtectionLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceThreatProtectionRequiredSecurityLevel(val.(*DeviceThreatProtectionLevel))
        }
        return nil
    }
    res["minAndroidSecurityPatchLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinAndroidSecurityPatchLevel(val)
        }
        return nil
    }
    res["osMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMaximumVersion(val)
        }
        return nil
    }
    res["osMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMinimumVersion(val)
        }
        return nil
    }
    res["passwordExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordExpirationDays(val)
        }
        return nil
    }
    res["passwordMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumLength(val)
        }
        return nil
    }
    res["passwordMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeLock(val)
        }
        return nil
    }
    res["passwordPreviousPasswordBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPreviousPasswordBlockCount(val)
        }
        return nil
    }
    res["passwordRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequired(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAndroidRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*AndroidRequiredPasswordType))
        }
        return nil
    }
    res["securityBlockJailbrokenDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityBlockJailbrokenDevices(val)
        }
        return nil
    }
    res["securityDisableUsbDebugging"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityDisableUsbDebugging(val)
        }
        return nil
    }
    res["securityPreventInstallAppsFromUnknownSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityPreventInstallAppsFromUnknownSources(val)
        }
        return nil
    }
    res["securityRequireCompanyPortalAppIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireCompanyPortalAppIntegrity(val)
        }
        return nil
    }
    res["securityRequireGooglePlayServices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireGooglePlayServices(val)
        }
        return nil
    }
    res["securityRequireSafetyNetAttestationBasicIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireSafetyNetAttestationBasicIntegrity(val)
        }
        return nil
    }
    res["securityRequireSafetyNetAttestationCertifiedDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireSafetyNetAttestationCertifiedDevice(val)
        }
        return nil
    }
    res["securityRequireUpToDateSecurityProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireUpToDateSecurityProviders(val)
        }
        return nil
    }
    res["securityRequireVerifyApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecurityRequireVerifyApps(val)
        }
        return nil
    }
    res["storageRequireEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageRequireEncryption(val)
        }
        return nil
    }
    return res
}
// GetMinAndroidSecurityPatchLevel gets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidWorkProfileCompliancePolicy) GetMinAndroidSecurityPatchLevel()(*string) {
    return m.minAndroidSecurityPatchLevel
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidWorkProfileCompliancePolicy) GetOsMaximumVersion()(*string) {
    return m.osMaximumVersion
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidWorkProfileCompliancePolicy) GetOsMinimumVersion()(*string) {
    return m.osMinimumVersion
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidWorkProfileCompliancePolicy) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidWorkProfileCompliancePolicy) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidWorkProfileCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    return m.passwordMinutesOfInactivityBeforeLock
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidWorkProfileCompliancePolicy) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidWorkProfileCompliancePolicy) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Android required password type.
func (m *AndroidWorkProfileCompliancePolicy) GetPasswordRequiredType()(*AndroidRequiredPasswordType) {
    return m.passwordRequiredType
}
// GetSecurityBlockJailbrokenDevices gets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityBlockJailbrokenDevices()(*bool) {
    return m.securityBlockJailbrokenDevices
}
// GetSecurityDisableUsbDebugging gets the securityDisableUsbDebugging property value. Disable USB debugging on Android devices.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityDisableUsbDebugging()(*bool) {
    return m.securityDisableUsbDebugging
}
// GetSecurityPreventInstallAppsFromUnknownSources gets the securityPreventInstallAppsFromUnknownSources property value. Require that devices disallow installation of apps from unknown sources.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityPreventInstallAppsFromUnknownSources()(*bool) {
    return m.securityPreventInstallAppsFromUnknownSources
}
// GetSecurityRequireCompanyPortalAppIntegrity gets the securityRequireCompanyPortalAppIntegrity property value. Require the device to pass the Company Portal client app runtime integrity check.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityRequireCompanyPortalAppIntegrity()(*bool) {
    return m.securityRequireCompanyPortalAppIntegrity
}
// GetSecurityRequireGooglePlayServices gets the securityRequireGooglePlayServices property value. Require Google Play Services to be installed and enabled on the device.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityRequireGooglePlayServices()(*bool) {
    return m.securityRequireGooglePlayServices
}
// GetSecurityRequireSafetyNetAttestationBasicIntegrity gets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool) {
    return m.securityRequireSafetyNetAttestationBasicIntegrity
}
// GetSecurityRequireSafetyNetAttestationCertifiedDevice gets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool) {
    return m.securityRequireSafetyNetAttestationCertifiedDevice
}
// GetSecurityRequireUpToDateSecurityProviders gets the securityRequireUpToDateSecurityProviders property value. Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityRequireUpToDateSecurityProviders()(*bool) {
    return m.securityRequireUpToDateSecurityProviders
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidWorkProfileCompliancePolicy) GetSecurityRequireVerifyApps()(*bool) {
    return m.securityRequireVerifyApps
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidWorkProfileCompliancePolicy) GetStorageRequireEncryption()(*bool) {
    return m.storageRequireEncryption
}
// Serialize serializes information the current object
func (m *AndroidWorkProfileCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("deviceThreatProtectionEnabled", m.GetDeviceThreatProtectionEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceThreatProtectionRequiredSecurityLevel() != nil {
        cast := (*m.GetDeviceThreatProtectionRequiredSecurityLevel()).String()
        err = writer.WriteStringValue("deviceThreatProtectionRequiredSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minAndroidSecurityPatchLevel", m.GetMinAndroidSecurityPatchLevel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMaximumVersion", m.GetOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osMinimumVersion", m.GetOsMinimumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordExpirationDays", m.GetPasswordExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinimumLength", m.GetPasswordMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeLock", m.GetPasswordMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordPreviousPasswordBlockCount", m.GetPasswordPreviousPasswordBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordRequired", m.GetPasswordRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityBlockJailbrokenDevices", m.GetSecurityBlockJailbrokenDevices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityDisableUsbDebugging", m.GetSecurityDisableUsbDebugging())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityPreventInstallAppsFromUnknownSources", m.GetSecurityPreventInstallAppsFromUnknownSources())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireCompanyPortalAppIntegrity", m.GetSecurityRequireCompanyPortalAppIntegrity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireGooglePlayServices", m.GetSecurityRequireGooglePlayServices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireSafetyNetAttestationBasicIntegrity", m.GetSecurityRequireSafetyNetAttestationBasicIntegrity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireSafetyNetAttestationCertifiedDevice", m.GetSecurityRequireSafetyNetAttestationCertifiedDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireUpToDateSecurityProviders", m.GetSecurityRequireUpToDateSecurityProviders())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("securityRequireVerifyApps", m.GetSecurityRequireVerifyApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRequireEncryption", m.GetStorageRequireEncryption())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceThreatProtectionEnabled sets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidWorkProfileCompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    m.deviceThreatProtectionEnabled = value
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *AndroidWorkProfileCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    m.deviceThreatProtectionRequiredSecurityLevel = value
}
// SetMinAndroidSecurityPatchLevel sets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidWorkProfileCompliancePolicy) SetMinAndroidSecurityPatchLevel(value *string)() {
    m.minAndroidSecurityPatchLevel = value
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidWorkProfileCompliancePolicy) SetOsMaximumVersion(value *string)() {
    m.osMaximumVersion = value
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidWorkProfileCompliancePolicy) SetOsMinimumVersion(value *string)() {
    m.osMinimumVersion = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidWorkProfileCompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidWorkProfileCompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidWorkProfileCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    m.passwordMinutesOfInactivityBeforeLock = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidWorkProfileCompliancePolicy) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidWorkProfileCompliancePolicy) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Android required password type.
func (m *AndroidWorkProfileCompliancePolicy) SetPasswordRequiredType(value *AndroidRequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetSecurityBlockJailbrokenDevices sets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityBlockJailbrokenDevices(value *bool)() {
    m.securityBlockJailbrokenDevices = value
}
// SetSecurityDisableUsbDebugging sets the securityDisableUsbDebugging property value. Disable USB debugging on Android devices.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityDisableUsbDebugging(value *bool)() {
    m.securityDisableUsbDebugging = value
}
// SetSecurityPreventInstallAppsFromUnknownSources sets the securityPreventInstallAppsFromUnknownSources property value. Require that devices disallow installation of apps from unknown sources.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityPreventInstallAppsFromUnknownSources(value *bool)() {
    m.securityPreventInstallAppsFromUnknownSources = value
}
// SetSecurityRequireCompanyPortalAppIntegrity sets the securityRequireCompanyPortalAppIntegrity property value. Require the device to pass the Company Portal client app runtime integrity check.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityRequireCompanyPortalAppIntegrity(value *bool)() {
    m.securityRequireCompanyPortalAppIntegrity = value
}
// SetSecurityRequireGooglePlayServices sets the securityRequireGooglePlayServices property value. Require Google Play Services to be installed and enabled on the device.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityRequireGooglePlayServices(value *bool)() {
    m.securityRequireGooglePlayServices = value
}
// SetSecurityRequireSafetyNetAttestationBasicIntegrity sets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)() {
    m.securityRequireSafetyNetAttestationBasicIntegrity = value
}
// SetSecurityRequireSafetyNetAttestationCertifiedDevice sets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)() {
    m.securityRequireSafetyNetAttestationCertifiedDevice = value
}
// SetSecurityRequireUpToDateSecurityProviders sets the securityRequireUpToDateSecurityProviders property value. Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityRequireUpToDateSecurityProviders(value *bool)() {
    m.securityRequireUpToDateSecurityProviders = value
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidWorkProfileCompliancePolicy) SetSecurityRequireVerifyApps(value *bool)() {
    m.securityRequireVerifyApps = value
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidWorkProfileCompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    m.storageRequireEncryption = value
}
