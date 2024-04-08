package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10CompliancePolicy 
type Windows10CompliancePolicy struct {
    DeviceCompliancePolicy
    // Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled
    bitLockerEnabled *bool
    // Require devices to be reported as healthy by Windows Device Health Attestation.
    codeIntegrityEnabled *bool
    // Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is enabled.
    earlyLaunchAntiMalwareDriverEnabled *bool
    // Maximum Windows Phone version.
    mobileOsMaximumVersion *string
    // Minimum Windows Phone version.
    mobileOsMinimumVersion *string
    // Maximum Windows 10 version.
    osMaximumVersion *string
    // Minimum Windows 10 version.
    osMinimumVersion *string
    // Indicates whether or not to block simple password.
    passwordBlockSimple *bool
    // The password expiration in days.
    passwordExpirationDays *int32
    // The number of character sets required in the password.
    passwordMinimumCharacterSetCount *int32
    // The minimum password length.
    passwordMinimumLength *int32
    // Minutes of inactivity before a password is required.
    passwordMinutesOfInactivityBeforeLock *int32
    // The number of previous passwords to prevent re-use of.
    passwordPreviousPasswordBlockCount *int32
    // Require a password to unlock Windows device.
    passwordRequired *bool
    // Require a password to unlock an idle device.
    passwordRequiredToUnlockFromIdle *bool
    // Possible values of required passwords.
    passwordRequiredType *RequiredPasswordType
    // Require devices to be reported as healthy by Windows Device Health Attestation.
    requireHealthyDeviceReport *bool
    // Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled.
    secureBootEnabled *bool
    // Require encryption on windows devices.
    storageRequireEncryption *bool
}
// NewWindows10CompliancePolicy instantiates a new Windows10CompliancePolicy and sets the default values.
func NewWindows10CompliancePolicy()(*Windows10CompliancePolicy) {
    m := &Windows10CompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.windows10CompliancePolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10CompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10CompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10CompliancePolicy(), nil
}
// GetBitLockerEnabled gets the bitLockerEnabled property value. Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled
func (m *Windows10CompliancePolicy) GetBitLockerEnabled()(*bool) {
    return m.bitLockerEnabled
}
// GetCodeIntegrityEnabled gets the codeIntegrityEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation.
func (m *Windows10CompliancePolicy) GetCodeIntegrityEnabled()(*bool) {
    return m.codeIntegrityEnabled
}
// GetEarlyLaunchAntiMalwareDriverEnabled gets the earlyLaunchAntiMalwareDriverEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is enabled.
func (m *Windows10CompliancePolicy) GetEarlyLaunchAntiMalwareDriverEnabled()(*bool) {
    return m.earlyLaunchAntiMalwareDriverEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10CompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceCompliancePolicy.GetFieldDeserializers()
    res["bitLockerEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerEnabled(val)
        }
        return nil
    }
    res["codeIntegrityEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeIntegrityEnabled(val)
        }
        return nil
    }
    res["earlyLaunchAntiMalwareDriverEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEarlyLaunchAntiMalwareDriverEnabled(val)
        }
        return nil
    }
    res["mobileOsMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileOsMaximumVersion(val)
        }
        return nil
    }
    res["mobileOsMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileOsMinimumVersion(val)
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
    res["passwordBlockSimple"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockSimple(val)
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
    res["passwordMinimumCharacterSetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinimumCharacterSetCount(val)
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
    res["passwordRequiredToUnlockFromIdle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredToUnlockFromIdle(val)
        }
        return nil
    }
    res["passwordRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequiredType(val.(*RequiredPasswordType))
        }
        return nil
    }
    res["requireHealthyDeviceReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequireHealthyDeviceReport(val)
        }
        return nil
    }
    res["secureBootEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBootEnabled(val)
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
// GetMobileOsMaximumVersion gets the mobileOsMaximumVersion property value. Maximum Windows Phone version.
func (m *Windows10CompliancePolicy) GetMobileOsMaximumVersion()(*string) {
    return m.mobileOsMaximumVersion
}
// GetMobileOsMinimumVersion gets the mobileOsMinimumVersion property value. Minimum Windows Phone version.
func (m *Windows10CompliancePolicy) GetMobileOsMinimumVersion()(*string) {
    return m.mobileOsMinimumVersion
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum Windows 10 version.
func (m *Windows10CompliancePolicy) GetOsMaximumVersion()(*string) {
    return m.osMaximumVersion
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum Windows 10 version.
func (m *Windows10CompliancePolicy) GetOsMinimumVersion()(*string) {
    return m.osMinimumVersion
}
// GetPasswordBlockSimple gets the passwordBlockSimple property value. Indicates whether or not to block simple password.
func (m *Windows10CompliancePolicy) GetPasswordBlockSimple()(*bool) {
    return m.passwordBlockSimple
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. The password expiration in days.
func (m *Windows10CompliancePolicy) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows10CompliancePolicy) GetPasswordMinimumCharacterSetCount()(*int32) {
    return m.passwordMinimumCharacterSetCount
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. The minimum password length.
func (m *Windows10CompliancePolicy) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeLock gets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *Windows10CompliancePolicy) GetPasswordMinutesOfInactivityBeforeLock()(*int32) {
    return m.passwordMinutesOfInactivityBeforeLock
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent re-use of.
func (m *Windows10CompliancePolicy) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequired gets the passwordRequired property value. Require a password to unlock Windows device.
func (m *Windows10CompliancePolicy) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredToUnlockFromIdle gets the passwordRequiredToUnlockFromIdle property value. Require a password to unlock an idle device.
func (m *Windows10CompliancePolicy) GetPasswordRequiredToUnlockFromIdle()(*bool) {
    return m.passwordRequiredToUnlockFromIdle
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows10CompliancePolicy) GetPasswordRequiredType()(*RequiredPasswordType) {
    return m.passwordRequiredType
}
// GetRequireHealthyDeviceReport gets the requireHealthyDeviceReport property value. Require devices to be reported as healthy by Windows Device Health Attestation.
func (m *Windows10CompliancePolicy) GetRequireHealthyDeviceReport()(*bool) {
    return m.requireHealthyDeviceReport
}
// GetSecureBootEnabled gets the secureBootEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled.
func (m *Windows10CompliancePolicy) GetSecureBootEnabled()(*bool) {
    return m.secureBootEnabled
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Require encryption on windows devices.
func (m *Windows10CompliancePolicy) GetStorageRequireEncryption()(*bool) {
    return m.storageRequireEncryption
}
// Serialize serializes information the current object
func (m *Windows10CompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceCompliancePolicy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("bitLockerEnabled", m.GetBitLockerEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("codeIntegrityEnabled", m.GetCodeIntegrityEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("earlyLaunchAntiMalwareDriverEnabled", m.GetEarlyLaunchAntiMalwareDriverEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobileOsMaximumVersion", m.GetMobileOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobileOsMinimumVersion", m.GetMobileOsMinimumVersion())
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
        err = writer.WriteBoolValue("passwordBlockSimple", m.GetPasswordBlockSimple())
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
        err = writer.WriteInt32Value("passwordMinimumCharacterSetCount", m.GetPasswordMinimumCharacterSetCount())
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
    {
        err = writer.WriteBoolValue("passwordRequiredToUnlockFromIdle", m.GetPasswordRequiredToUnlockFromIdle())
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
        err = writer.WriteBoolValue("requireHealthyDeviceReport", m.GetRequireHealthyDeviceReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("secureBootEnabled", m.GetSecureBootEnabled())
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
// SetBitLockerEnabled sets the bitLockerEnabled property value. Require devices to be reported healthy by Windows Device Health Attestation - bit locker is enabled
func (m *Windows10CompliancePolicy) SetBitLockerEnabled(value *bool)() {
    m.bitLockerEnabled = value
}
// SetCodeIntegrityEnabled sets the codeIntegrityEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation.
func (m *Windows10CompliancePolicy) SetCodeIntegrityEnabled(value *bool)() {
    m.codeIntegrityEnabled = value
}
// SetEarlyLaunchAntiMalwareDriverEnabled sets the earlyLaunchAntiMalwareDriverEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation - early launch antimalware driver is enabled.
func (m *Windows10CompliancePolicy) SetEarlyLaunchAntiMalwareDriverEnabled(value *bool)() {
    m.earlyLaunchAntiMalwareDriverEnabled = value
}
// SetMobileOsMaximumVersion sets the mobileOsMaximumVersion property value. Maximum Windows Phone version.
func (m *Windows10CompliancePolicy) SetMobileOsMaximumVersion(value *string)() {
    m.mobileOsMaximumVersion = value
}
// SetMobileOsMinimumVersion sets the mobileOsMinimumVersion property value. Minimum Windows Phone version.
func (m *Windows10CompliancePolicy) SetMobileOsMinimumVersion(value *string)() {
    m.mobileOsMinimumVersion = value
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum Windows 10 version.
func (m *Windows10CompliancePolicy) SetOsMaximumVersion(value *string)() {
    m.osMaximumVersion = value
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum Windows 10 version.
func (m *Windows10CompliancePolicy) SetOsMinimumVersion(value *string)() {
    m.osMinimumVersion = value
}
// SetPasswordBlockSimple sets the passwordBlockSimple property value. Indicates whether or not to block simple password.
func (m *Windows10CompliancePolicy) SetPasswordBlockSimple(value *bool)() {
    m.passwordBlockSimple = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. The password expiration in days.
func (m *Windows10CompliancePolicy) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows10CompliancePolicy) SetPasswordMinimumCharacterSetCount(value *int32)() {
    m.passwordMinimumCharacterSetCount = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. The minimum password length.
func (m *Windows10CompliancePolicy) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeLock sets the passwordMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a password is required.
func (m *Windows10CompliancePolicy) SetPasswordMinutesOfInactivityBeforeLock(value *int32)() {
    m.passwordMinutesOfInactivityBeforeLock = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent re-use of.
func (m *Windows10CompliancePolicy) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequired sets the passwordRequired property value. Require a password to unlock Windows device.
func (m *Windows10CompliancePolicy) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredToUnlockFromIdle sets the passwordRequiredToUnlockFromIdle property value. Require a password to unlock an idle device.
func (m *Windows10CompliancePolicy) SetPasswordRequiredToUnlockFromIdle(value *bool)() {
    m.passwordRequiredToUnlockFromIdle = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows10CompliancePolicy) SetPasswordRequiredType(value *RequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetRequireHealthyDeviceReport sets the requireHealthyDeviceReport property value. Require devices to be reported as healthy by Windows Device Health Attestation.
func (m *Windows10CompliancePolicy) SetRequireHealthyDeviceReport(value *bool)() {
    m.requireHealthyDeviceReport = value
}
// SetSecureBootEnabled sets the secureBootEnabled property value. Require devices to be reported as healthy by Windows Device Health Attestation - secure boot is enabled.
func (m *Windows10CompliancePolicy) SetSecureBootEnabled(value *bool)() {
    m.secureBootEnabled = value
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Require encryption on windows devices.
func (m *Windows10CompliancePolicy) SetStorageRequireEncryption(value *bool)() {
    m.storageRequireEncryption = value
}
