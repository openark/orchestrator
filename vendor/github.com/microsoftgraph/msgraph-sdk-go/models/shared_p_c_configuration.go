package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedPCConfiguration 
type SharedPCConfiguration struct {
    DeviceConfiguration
    // Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false.
    accountManagerPolicy SharedPCAccountManagerPolicyable
    // Type of accounts that are allowed to share the PC.
    allowedAccounts *SharedPCAllowedAccountType
    // Specifies whether local storage is allowed on a shared PC.
    allowLocalStorage *bool
    // Disables the account manager for shared PC mode.
    disableAccountManager *bool
    // Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and later, this policy will be applied without setting Enabled to true.
    disableEduPolicies *bool
    // Specifies whether the default shared PC power policies should be disabled.
    disablePowerPolicies *bool
    // Disables the requirement to sign in whenever the device wakes up from sleep mode.
    disableSignInOnResume *bool
    // Enables shared PC mode and applies the shared pc policies.
    enabled *bool
    // Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0 prevents the sleep timeout from occurring.
    idleTimeBeforeSleepInSeconds *int32
    // Specifies the display text for the account shown on the sign-in screen which launches the app specified by SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set.
    kioskAppDisplayName *string
    // Specifies the application user model ID of the app to use with assigned access.
    kioskAppUserModelId *string
    // Specifies the daily start time of maintenance hour.
    maintenanceStartTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
}
// NewSharedPCConfiguration instantiates a new SharedPCConfiguration and sets the default values.
func NewSharedPCConfiguration()(*SharedPCConfiguration) {
    m := &SharedPCConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.sharedPCConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSharedPCConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedPCConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedPCConfiguration(), nil
}
// GetAccountManagerPolicy gets the accountManagerPolicy property value. Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false.
func (m *SharedPCConfiguration) GetAccountManagerPolicy()(SharedPCAccountManagerPolicyable) {
    return m.accountManagerPolicy
}
// GetAllowedAccounts gets the allowedAccounts property value. Type of accounts that are allowed to share the PC.
func (m *SharedPCConfiguration) GetAllowedAccounts()(*SharedPCAllowedAccountType) {
    return m.allowedAccounts
}
// GetAllowLocalStorage gets the allowLocalStorage property value. Specifies whether local storage is allowed on a shared PC.
func (m *SharedPCConfiguration) GetAllowLocalStorage()(*bool) {
    return m.allowLocalStorage
}
// GetDisableAccountManager gets the disableAccountManager property value. Disables the account manager for shared PC mode.
func (m *SharedPCConfiguration) GetDisableAccountManager()(*bool) {
    return m.disableAccountManager
}
// GetDisableEduPolicies gets the disableEduPolicies property value. Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and later, this policy will be applied without setting Enabled to true.
func (m *SharedPCConfiguration) GetDisableEduPolicies()(*bool) {
    return m.disableEduPolicies
}
// GetDisablePowerPolicies gets the disablePowerPolicies property value. Specifies whether the default shared PC power policies should be disabled.
func (m *SharedPCConfiguration) GetDisablePowerPolicies()(*bool) {
    return m.disablePowerPolicies
}
// GetDisableSignInOnResume gets the disableSignInOnResume property value. Disables the requirement to sign in whenever the device wakes up from sleep mode.
func (m *SharedPCConfiguration) GetDisableSignInOnResume()(*bool) {
    return m.disableSignInOnResume
}
// GetEnabled gets the enabled property value. Enables shared PC mode and applies the shared pc policies.
func (m *SharedPCConfiguration) GetEnabled()(*bool) {
    return m.enabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedPCConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["accountManagerPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharedPCAccountManagerPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountManagerPolicy(val.(SharedPCAccountManagerPolicyable))
        }
        return nil
    }
    res["allowedAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSharedPCAllowedAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedAccounts(val.(*SharedPCAllowedAccountType))
        }
        return nil
    }
    res["allowLocalStorage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLocalStorage(val)
        }
        return nil
    }
    res["disableAccountManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAccountManager(val)
        }
        return nil
    }
    res["disableEduPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableEduPolicies(val)
        }
        return nil
    }
    res["disablePowerPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisablePowerPolicies(val)
        }
        return nil
    }
    res["disableSignInOnResume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableSignInOnResume(val)
        }
        return nil
    }
    res["enabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["idleTimeBeforeSleepInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdleTimeBeforeSleepInSeconds(val)
        }
        return nil
    }
    res["kioskAppDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskAppDisplayName(val)
        }
        return nil
    }
    res["kioskAppUserModelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKioskAppUserModelId(val)
        }
        return nil
    }
    res["maintenanceStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintenanceStartTime(val)
        }
        return nil
    }
    return res
}
// GetIdleTimeBeforeSleepInSeconds gets the idleTimeBeforeSleepInSeconds property value. Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0 prevents the sleep timeout from occurring.
func (m *SharedPCConfiguration) GetIdleTimeBeforeSleepInSeconds()(*int32) {
    return m.idleTimeBeforeSleepInSeconds
}
// GetKioskAppDisplayName gets the kioskAppDisplayName property value. Specifies the display text for the account shown on the sign-in screen which launches the app specified by SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set.
func (m *SharedPCConfiguration) GetKioskAppDisplayName()(*string) {
    return m.kioskAppDisplayName
}
// GetKioskAppUserModelId gets the kioskAppUserModelId property value. Specifies the application user model ID of the app to use with assigned access.
func (m *SharedPCConfiguration) GetKioskAppUserModelId()(*string) {
    return m.kioskAppUserModelId
}
// GetMaintenanceStartTime gets the maintenanceStartTime property value. Specifies the daily start time of maintenance hour.
func (m *SharedPCConfiguration) GetMaintenanceStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.maintenanceStartTime
}
// Serialize serializes information the current object
func (m *SharedPCConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accountManagerPolicy", m.GetAccountManagerPolicy())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedAccounts() != nil {
        cast := (*m.GetAllowedAccounts()).String()
        err = writer.WriteStringValue("allowedAccounts", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowLocalStorage", m.GetAllowLocalStorage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAccountManager", m.GetDisableAccountManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableEduPolicies", m.GetDisableEduPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disablePowerPolicies", m.GetDisablePowerPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableSignInOnResume", m.GetDisableSignInOnResume())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("idleTimeBeforeSleepInSeconds", m.GetIdleTimeBeforeSleepInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskAppDisplayName", m.GetKioskAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("kioskAppUserModelId", m.GetKioskAppUserModelId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("maintenanceStartTime", m.GetMaintenanceStartTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountManagerPolicy sets the accountManagerPolicy property value. Specifies how accounts are managed on a shared PC. Only applies when disableAccountManager is false.
func (m *SharedPCConfiguration) SetAccountManagerPolicy(value SharedPCAccountManagerPolicyable)() {
    m.accountManagerPolicy = value
}
// SetAllowedAccounts sets the allowedAccounts property value. Type of accounts that are allowed to share the PC.
func (m *SharedPCConfiguration) SetAllowedAccounts(value *SharedPCAllowedAccountType)() {
    m.allowedAccounts = value
}
// SetAllowLocalStorage sets the allowLocalStorage property value. Specifies whether local storage is allowed on a shared PC.
func (m *SharedPCConfiguration) SetAllowLocalStorage(value *bool)() {
    m.allowLocalStorage = value
}
// SetDisableAccountManager sets the disableAccountManager property value. Disables the account manager for shared PC mode.
func (m *SharedPCConfiguration) SetDisableAccountManager(value *bool)() {
    m.disableAccountManager = value
}
// SetDisableEduPolicies sets the disableEduPolicies property value. Specifies whether the default shared PC education environment policies should be disabled. For Windows 10 RS2 and later, this policy will be applied without setting Enabled to true.
func (m *SharedPCConfiguration) SetDisableEduPolicies(value *bool)() {
    m.disableEduPolicies = value
}
// SetDisablePowerPolicies sets the disablePowerPolicies property value. Specifies whether the default shared PC power policies should be disabled.
func (m *SharedPCConfiguration) SetDisablePowerPolicies(value *bool)() {
    m.disablePowerPolicies = value
}
// SetDisableSignInOnResume sets the disableSignInOnResume property value. Disables the requirement to sign in whenever the device wakes up from sleep mode.
func (m *SharedPCConfiguration) SetDisableSignInOnResume(value *bool)() {
    m.disableSignInOnResume = value
}
// SetEnabled sets the enabled property value. Enables shared PC mode and applies the shared pc policies.
func (m *SharedPCConfiguration) SetEnabled(value *bool)() {
    m.enabled = value
}
// SetIdleTimeBeforeSleepInSeconds sets the idleTimeBeforeSleepInSeconds property value. Specifies the time in seconds that a device must sit idle before the PC goes to sleep. Setting this value to 0 prevents the sleep timeout from occurring.
func (m *SharedPCConfiguration) SetIdleTimeBeforeSleepInSeconds(value *int32)() {
    m.idleTimeBeforeSleepInSeconds = value
}
// SetKioskAppDisplayName sets the kioskAppDisplayName property value. Specifies the display text for the account shown on the sign-in screen which launches the app specified by SetKioskAppUserModelId. Only applies when KioskAppUserModelId is set.
func (m *SharedPCConfiguration) SetKioskAppDisplayName(value *string)() {
    m.kioskAppDisplayName = value
}
// SetKioskAppUserModelId sets the kioskAppUserModelId property value. Specifies the application user model ID of the app to use with assigned access.
func (m *SharedPCConfiguration) SetKioskAppUserModelId(value *string)() {
    m.kioskAppUserModelId = value
}
// SetMaintenanceStartTime sets the maintenanceStartTime property value. Specifies the daily start time of maintenance hour.
func (m *SharedPCConfiguration) SetMaintenanceStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.maintenanceStartTime = value
}
