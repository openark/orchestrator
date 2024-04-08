package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidCompliancePolicy 
type AndroidCompliancePolicy struct {
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
// NewAndroidCompliancePolicy instantiates a new AndroidCompliancePolicy and sets the default values.
func NewAndroidCompliancePolicy()(*AndroidCompliancePolicy) {
    m := &AndroidCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.androidCompliancePolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidCompliancePolicy(), nil
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection.
func (m *AndroidCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    return m.deviceThreatProtectionEnabled
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *AndroidCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    return m.deviceThreatProtectionRequiredSecurityLevel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *AndroidCompliancePolicy) GetMinAndroidSecurityPatchLevel()(*string) {
    return m.minAndroidSecurityPatchLevel
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidCompliancePolicy) GetOsMaximumVersion()(*string) {
    return m.osMaximumVersion
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidCompliancePolicy) GetOsMinimumVersion()(*string) {
    return m.osMinimumVersion
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidCompliancePolicy) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidCompliancePolicy) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidCompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    return m.passwordMinutesOfInactivityBeforeLock
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidCompliancePolicy) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidCompliancePolicy) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Android required password type.
func (m *AndroidCompliancePolicy) GetPasswordRequiredType()(*AndroidRequiredPasswordType) {
    return m.passwordRequiredType
}
// GetSecurityBlockJailbrokenDevices gets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AndroidCompliancePolicy) GetSecurityBlockJailbrokenDevices()(*bool) {
    return m.securityBlockJailbrokenDevices
}
// GetSecurityDisableUsbDebugging gets the securityDisableUsbDebugging property value. Disable USB debugging on Android devices.
func (m *AndroidCompliancePolicy) GetSecurityDisableUsbDebugging()(*bool) {
    return m.securityDisableUsbDebugging
}
// GetSecurityPreventInstallAppsFromUnknownSources gets the securityPreventInstallAppsFromUnknownSources property value. Require that devices disallow installation of apps from unknown sources.
func (m *AndroidCompliancePolicy) GetSecurityPreventInstallAppsFromUnknownSources()(*bool) {
    return m.securityPreventInstallAppsFromUnknownSources
}
// GetSecurityRequireCompanyPortalAppIntegrity gets the securityRequireCompanyPortalAppIntegrity property value. Require the device to pass the Company Portal client app runtime integrity check.
func (m *AndroidCompliancePolicy) GetSecurityRequireCompanyPortalAppIntegrity()(*bool) {
    return m.securityRequireCompanyPortalAppIntegrity
}
// GetSecurityRequireGooglePlayServices gets the securityRequireGooglePlayServices property value. Require Google Play Services to be installed and enabled on the device.
func (m *AndroidCompliancePolicy) GetSecurityRequireGooglePlayServices()(*bool) {
    return m.securityRequireGooglePlayServices
}
// GetSecurityRequireSafetyNetAttestationBasicIntegrity gets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidCompliancePolicy) GetSecurityRequireSafetyNetAttestationBasicIntegrity()(*bool) {
    return m.securityRequireSafetyNetAttestationBasicIntegrity
}
// GetSecurityRequireSafetyNetAttestationCertifiedDevice gets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidCompliancePolicy) GetSecurityRequireSafetyNetAttestationCertifiedDevice()(*bool) {
    return m.securityRequireSafetyNetAttestationCertifiedDevice
}
// GetSecurityRequireUpToDateSecurityProviders gets the securityRequireUpToDateSecurityProviders property value. Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
func (m *AndroidCompliancePolicy) GetSecurityRequireUpToDateSecurityProviders()(*bool) {
    return m.securityRequireUpToDateSecurityProviders
}
// GetSecurityRequireVerifyApps gets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidCompliancePolicy) GetSecurityRequireVerifyApps()(*bool) {
    return m.securityRequireVerifyApps
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidCompliancePolicy) GetStorageRequireEncryption()(*bool) {
    return m.storageRequireEncryption
}
// Serialize serializes information the current object
func (m *AndroidCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AndroidCompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    m.deviceThreatProtectionEnabled = value
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *AndroidCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    m.deviceThreatProtectionRequiredSecurityLevel = value
}
// SetMinAndroidSecurityPatchLevel sets the minAndroidSecurityPatchLevel property value. Minimum Android security patch level.
func (m *AndroidCompliancePolicy) SetMinAndroidSecurityPatchLevel(value *string)() {
    m.minAndroidSecurityPatchLevel = value
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Android version.
func (m *AndroidCompliancePolicy) SetOsMaximumVersion(value *string)() {
    m.osMaximumVersion = value
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Android version.
func (m *AndroidCompliancePolicy) SetOsMinimumVersion(value *string)() {
    m.osMinimumVersion = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires. Valid values 1 to 365
func (m *AndroidCompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum password length. Valid values 4 to 16
func (m *AndroidCompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *AndroidCompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    m.passwordMinutesOfInactivityBeforeLock = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 1 to 24
func (m *AndroidCompliancePolicy) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock device.
func (m *AndroidCompliancePolicy) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Android required password type.
func (m *AndroidCompliancePolicy) SetPasswordRequiredType(value *AndroidRequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetSecurityBlockJailbrokenDevices sets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *AndroidCompliancePolicy) SetSecurityBlockJailbrokenDevices(value *bool)() {
    m.securityBlockJailbrokenDevices = value
}
// SetSecurityDisableUsbDebugging sets the securityDisableUsbDebugging property value. Disable USB debugging on Android devices.
func (m *AndroidCompliancePolicy) SetSecurityDisableUsbDebugging(value *bool)() {
    m.securityDisableUsbDebugging = value
}
// SetSecurityPreventInstallAppsFromUnknownSources sets the securityPreventInstallAppsFromUnknownSources property value. Require that devices disallow installation of apps from unknown sources.
func (m *AndroidCompliancePolicy) SetSecurityPreventInstallAppsFromUnknownSources(value *bool)() {
    m.securityPreventInstallAppsFromUnknownSources = value
}
// SetSecurityRequireCompanyPortalAppIntegrity sets the securityRequireCompanyPortalAppIntegrity property value. Require the device to pass the Company Portal client app runtime integrity check.
func (m *AndroidCompliancePolicy) SetSecurityRequireCompanyPortalAppIntegrity(value *bool)() {
    m.securityRequireCompanyPortalAppIntegrity = value
}
// SetSecurityRequireGooglePlayServices sets the securityRequireGooglePlayServices property value. Require Google Play Services to be installed and enabled on the device.
func (m *AndroidCompliancePolicy) SetSecurityRequireGooglePlayServices(value *bool)() {
    m.securityRequireGooglePlayServices = value
}
// SetSecurityRequireSafetyNetAttestationBasicIntegrity sets the securityRequireSafetyNetAttestationBasicIntegrity property value. Require the device to pass the SafetyNet basic integrity check.
func (m *AndroidCompliancePolicy) SetSecurityRequireSafetyNetAttestationBasicIntegrity(value *bool)() {
    m.securityRequireSafetyNetAttestationBasicIntegrity = value
}
// SetSecurityRequireSafetyNetAttestationCertifiedDevice sets the securityRequireSafetyNetAttestationCertifiedDevice property value. Require the device to pass the SafetyNet certified device check.
func (m *AndroidCompliancePolicy) SetSecurityRequireSafetyNetAttestationCertifiedDevice(value *bool)() {
    m.securityRequireSafetyNetAttestationCertifiedDevice = value
}
// SetSecurityRequireUpToDateSecurityProviders sets the securityRequireUpToDateSecurityProviders property value. Require the device to have up to date security providers. The device will require Google Play Services to be enabled and up to date.
func (m *AndroidCompliancePolicy) SetSecurityRequireUpToDateSecurityProviders(value *bool)() {
    m.securityRequireUpToDateSecurityProviders = value
}
// SetSecurityRequireVerifyApps sets the securityRequireVerifyApps property value. Require the Android Verify apps feature is turned on.
func (m *AndroidCompliancePolicy) SetSecurityRequireVerifyApps(value *bool)() {
    m.securityRequireVerifyApps = value
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on Android devices.
func (m *AndroidCompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    m.storageRequireEncryption = value
}
