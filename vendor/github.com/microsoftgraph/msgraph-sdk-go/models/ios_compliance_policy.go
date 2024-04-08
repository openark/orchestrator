package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosCompliancePolicy 
type IosCompliancePolicy struct {
    DeviceCompliancePolicy
    // Require that devices have enabled device threat protection .
    deviceThreatProtectionEnabled *bool
    // Device threat protection levels for the Device Threat Protection API.
    deviceThreatProtectionRequiredSecurityLevel *DeviceThreatProtectionLevel
    // Indicates whether or not to require a managed email profile.
    managedEmailProfileRequired *bool
    // Maximum IOS version.
    osMaximumVersion *string
    // Minimum IOS version.
    osMinimumVersion *string
    // Indicates whether or not to block simple passcodes.
    passcodeBlockSimple *bool
    // Number of days before the passcode expires. Valid values 1 to 65535
    passcodeExpirationDays *int32
    // The number of character sets required in the password.
    passcodeMinimumCharacterSetCount *int32
    // Minimum length of passcode. Valid values 4 to 14
    passcodeMinimumLength *int32
    // Minutes of inactivity before a passcode is required.
    passcodeMinutesOfInactivityBeforeLock *int32
    // Number of previous passcodes to block. Valid values 1 to 24
    passcodePreviousPasscodeBlockCount *int32
    // Indicates whether or not to require a passcode.
    passcodeRequired *bool
    // Possible values of required passwords.
    passcodeRequiredType *RequiredPasswordType
    // Devices must not be jailbroken or rooted.
    securityBlockJailbrokenDevices *bool
}
// NewIosCompliancePolicy instantiates a new IosCompliancePolicy and sets the default values.
func NewIosCompliancePolicy()(*IosCompliancePolicy) {
    m := &IosCompliancePolicy{
        DeviceCompliancePolicy: *NewDeviceCompliancePolicy(),
    }
    odataTypeValue := "#microsoft.graph.iosCompliancePolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosCompliancePolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosCompliancePolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosCompliancePolicy(), nil
}
// GetDeviceThreatProtectionEnabled gets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection .
func (m *IosCompliancePolicy) GetDeviceThreatProtectionEnabled()(*bool) {
    return m.deviceThreatProtectionEnabled
}
// GetDeviceThreatProtectionRequiredSecurityLevel gets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *IosCompliancePolicy) GetDeviceThreatProtectionRequiredSecurityLevel()(*DeviceThreatProtectionLevel) {
    return m.deviceThreatProtectionRequiredSecurityLevel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosCompliancePolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["managedEmailProfileRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedEmailProfileRequired(val)
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
    res["passcodeBlockSimple"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeBlockSimple(val)
        }
        return nil
    }
    res["passcodeExpirationDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeExpirationDays(val)
        }
        return nil
    }
    res["passcodeMinimumCharacterSetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinimumCharacterSetCount(val)
        }
        return nil
    }
    res["passcodeMinimumLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinimumLength(val)
        }
        return nil
    }
    res["passcodeMinutesOfInactivityBeforeLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeMinutesOfInactivityBeforeLock(val)
        }
        return nil
    }
    res["passcodePreviousPasscodeBlockCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodePreviousPasscodeBlockCount(val)
        }
        return nil
    }
    res["passcodeRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeRequired(val)
        }
        return nil
    }
    res["passcodeRequiredType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequiredPasswordType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasscodeRequiredType(val.(*RequiredPasswordType))
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
    return res
}
// GetManagedEmailProfileRequired gets the managedEmailProfileRequired property value. Indicates whether or not to require a managed email profile.
func (m *IosCompliancePolicy) GetManagedEmailProfileRequired()(*bool) {
    return m.managedEmailProfileRequired
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Maximum IOS version.
func (m *IosCompliancePolicy) GetOsMaximumVersion()(*string) {
    return m.osMaximumVersion
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Minimum IOS version.
func (m *IosCompliancePolicy) GetOsMinimumVersion()(*string) {
    return m.osMinimumVersion
}
// GetPasscodeBlockSimple gets the passcodeBlockSimple property value. Indicates whether or not to block simple passcodes.
func (m *IosCompliancePolicy) GetPasscodeBlockSimple()(*bool) {
    return m.passcodeBlockSimple
}
// GetPasscodeExpirationDays gets the passcodeExpirationDays property value. Number of days before the passcode expires. Valid values 1 to 65535
func (m *IosCompliancePolicy) GetPasscodeExpirationDays()(*int32) {
    return m.passcodeExpirationDays
}
// GetPasscodeMinimumCharacterSetCount gets the passcodeMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *IosCompliancePolicy) GetPasscodeMinimumCharacterSetCount()(*int32) {
    return m.passcodeMinimumCharacterSetCount
}
// GetPasscodeMinimumLength gets the passcodeMinimumLength property value. Minimum length of passcode. Valid values 4 to 14
func (m *IosCompliancePolicy) GetPasscodeMinimumLength()(*int32) {
    return m.passcodeMinimumLength
}
// GetPasscodeMinutesOfInactivityBeforeLock gets the passcodeMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a passcode is required.
func (m *IosCompliancePolicy) GetPasscodeMinutesOfInactivityBeforeLock()(*int32) {
    return m.passcodeMinutesOfInactivityBeforeLock
}
// GetPasscodePreviousPasscodeBlockCount gets the passcodePreviousPasscodeBlockCount property value. Number of previous passcodes to block. Valid values 1 to 24
func (m *IosCompliancePolicy) GetPasscodePreviousPasscodeBlockCount()(*int32) {
    return m.passcodePreviousPasscodeBlockCount
}
// GetPasscodeRequired gets the passcodeRequired property value. Indicates whether or not to require a passcode.
func (m *IosCompliancePolicy) GetPasscodeRequired()(*bool) {
    return m.passcodeRequired
}
// GetPasscodeRequiredType gets the passcodeRequiredType property value. Possible values of required passwords.
func (m *IosCompliancePolicy) GetPasscodeRequiredType()(*RequiredPasswordType) {
    return m.passcodeRequiredType
}
// GetSecurityBlockJailbrokenDevices gets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *IosCompliancePolicy) GetSecurityBlockJailbrokenDevices()(*bool) {
    return m.securityBlockJailbrokenDevices
}
// Serialize serializes information the current object
func (m *IosCompliancePolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("managedEmailProfileRequired", m.GetManagedEmailProfileRequired())
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
        err = writer.WriteBoolValue("passcodeBlockSimple", m.GetPasscodeBlockSimple())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeExpirationDays", m.GetPasscodeExpirationDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinimumCharacterSetCount", m.GetPasscodeMinimumCharacterSetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinimumLength", m.GetPasscodeMinimumLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodeMinutesOfInactivityBeforeLock", m.GetPasscodeMinutesOfInactivityBeforeLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passcodePreviousPasscodeBlockCount", m.GetPasscodePreviousPasscodeBlockCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passcodeRequired", m.GetPasscodeRequired())
        if err != nil {
            return err
        }
    }
    if m.GetPasscodeRequiredType() != nil {
        cast := (*m.GetPasscodeRequiredType()).String()
        err = writer.WriteStringValue("passcodeRequiredType", &cast)
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
    return nil
}
// SetDeviceThreatProtectionEnabled sets the deviceThreatProtectionEnabled property value. Require that devices have enabled device threat protection .
func (m *IosCompliancePolicy) SetDeviceThreatProtectionEnabled(value *bool)() {
    m.deviceThreatProtectionEnabled = value
}
// SetDeviceThreatProtectionRequiredSecurityLevel sets the deviceThreatProtectionRequiredSecurityLevel property value. Device threat protection levels for the Device Threat Protection API.
func (m *IosCompliancePolicy) SetDeviceThreatProtectionRequiredSecurityLevel(value *DeviceThreatProtectionLevel)() {
    m.deviceThreatProtectionRequiredSecurityLevel = value
}
// SetManagedEmailProfileRequired sets the managedEmailProfileRequired property value. Indicates whether or not to require a managed email profile.
func (m *IosCompliancePolicy) SetManagedEmailProfileRequired(value *bool)() {
    m.managedEmailProfileRequired = value
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Maximum IOS version.
func (m *IosCompliancePolicy) SetOsMaximumVersion(value *string)() {
    m.osMaximumVersion = value
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Minimum IOS version.
func (m *IosCompliancePolicy) SetOsMinimumVersion(value *string)() {
    m.osMinimumVersion = value
}
// SetPasscodeBlockSimple sets the passcodeBlockSimple property value. Indicates whether or not to block simple passcodes.
func (m *IosCompliancePolicy) SetPasscodeBlockSimple(value *bool)() {
    m.passcodeBlockSimple = value
}
// SetPasscodeExpirationDays sets the passcodeExpirationDays property value. Number of days before the passcode expires. Valid values 1 to 65535
func (m *IosCompliancePolicy) SetPasscodeExpirationDays(value *int32)() {
    m.passcodeExpirationDays = value
}
// SetPasscodeMinimumCharacterSetCount sets the passcodeMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *IosCompliancePolicy) SetPasscodeMinimumCharacterSetCount(value *int32)() {
    m.passcodeMinimumCharacterSetCount = value
}
// SetPasscodeMinimumLength sets the passcodeMinimumLength property value. Minimum length of passcode. Valid values 4 to 14
func (m *IosCompliancePolicy) SetPasscodeMinimumLength(value *int32)() {
    m.passcodeMinimumLength = value
}
// SetPasscodeMinutesOfInactivityBeforeLock sets the passcodeMinutesOfInactivityBeforeLock property value. Minutes of inactivity before a passcode is required.
func (m *IosCompliancePolicy) SetPasscodeMinutesOfInactivityBeforeLock(value *int32)() {
    m.passcodeMinutesOfInactivityBeforeLock = value
}
// SetPasscodePreviousPasscodeBlockCount sets the passcodePreviousPasscodeBlockCount property value. Number of previous passcodes to block. Valid values 1 to 24
func (m *IosCompliancePolicy) SetPasscodePreviousPasscodeBlockCount(value *int32)() {
    m.passcodePreviousPasscodeBlockCount = value
}
// SetPasscodeRequired sets the passcodeRequired property value. Indicates whether or not to require a passcode.
func (m *IosCompliancePolicy) SetPasscodeRequired(value *bool)() {
    m.passcodeRequired = value
}
// SetPasscodeRequiredType sets the passcodeRequiredType property value. Possible values of required passwords.
func (m *IosCompliancePolicy) SetPasscodeRequiredType(value *RequiredPasswordType)() {
    m.passcodeRequiredType = value
}
// SetSecurityBlockJailbrokenDevices sets the securityBlockJailbrokenDevices property value. Devices must not be jailbroken or rooted.
func (m *IosCompliancePolicy) SetSecurityBlockJailbrokenDevices(value *bool)() {
    m.securityBlockJailbrokenDevices = value
}
