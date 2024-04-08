package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows81GeneralConfiguration 
type Windows81GeneralConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account.
    accountsBlockAddingNonMicrosoftAccountEmail *bool
    // Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
    applyOnlyToWindows81 *bool
    // Indicates whether or not to block auto fill.
    browserBlockAutofill *bool
    // Indicates whether or not to block automatic detection of Intranet sites.
    browserBlockAutomaticDetectionOfIntranetSites *bool
    // Indicates whether or not to block enterprise mode access.
    browserBlockEnterpriseModeAccess *bool
    // Indicates whether or not to Block the user from using JavaScript.
    browserBlockJavaScript *bool
    // Indicates whether or not to block plug-ins.
    browserBlockPlugins *bool
    // Indicates whether or not to block popups.
    browserBlockPopups *bool
    // Indicates whether or not to Block the user from sending the do not track header.
    browserBlockSendingDoNotTrackHeader *bool
    // Indicates whether or not to block a single word entry on Intranet sites.
    browserBlockSingleWordEntryOnIntranetSites *bool
    // The enterprise mode site list location. Could be a local file, local network or http location.
    browserEnterpriseModeSiteListLocation *string
    // Possible values for internet site security level.
    browserInternetSecurityLevel *InternetSiteSecurityLevel
    // Possible values for site security level.
    browserIntranetSecurityLevel *SiteSecurityLevel
    // The logging report location.
    browserLoggingReportLocation *string
    // Indicates whether or not to require a firewall.
    browserRequireFirewall *bool
    // Indicates whether or not to require fraud warning.
    browserRequireFraudWarning *bool
    // Indicates whether or not to require high security for restricted sites.
    browserRequireHighSecurityForRestrictedSites *bool
    // Indicates whether or not to require the user to use the smart screen filter.
    browserRequireSmartScreen *bool
    // Possible values for site security level.
    browserTrustedSitesSecurityLevel *SiteSecurityLevel
    // Indicates whether or not to block data roaming.
    cellularBlockDataRoaming *bool
    // Indicates whether or not to block diagnostic data submission.
    diagnosticsBlockDataSubmission *bool
    // Indicates whether or not to Block the user from using a pictures password and pin.
    passwordBlockPicturePasswordAndPin *bool
    // Password expiration in days.
    passwordExpirationDays *int32
    // The number of character sets required in the password.
    passwordMinimumCharacterSetCount *int32
    // The minimum password length.
    passwordMinimumLength *int32
    // The minutes of inactivity before the screen times out.
    passwordMinutesOfInactivityBeforeScreenTimeout *int32
    // The number of previous passwords to prevent re-use of. Valid values 0 to 24
    passwordPreviousPasswordBlockCount *int32
    // Possible values of required passwords.
    passwordRequiredType *RequiredPasswordType
    // The number of sign in failures before factory reset.
    passwordSignInFailureCountBeforeFactoryReset *int32
    // Indicates whether or not to require encryption on a mobile device.
    storageRequireDeviceEncryption *bool
    // Indicates whether or not to require automatic updates.
    updatesRequireAutomaticUpdates *bool
    // Possible values for Windows user account control settings.
    userAccountControlSettings *WindowsUserAccountControlSettings
    // The work folders url.
    workFoldersUrl *string
}
// NewWindows81GeneralConfiguration instantiates a new Windows81GeneralConfiguration and sets the default values.
func NewWindows81GeneralConfiguration()(*Windows81GeneralConfiguration) {
    m := &Windows81GeneralConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows81GeneralConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows81GeneralConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows81GeneralConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows81GeneralConfiguration(), nil
}
// GetAccountsBlockAddingNonMicrosoftAccountEmail gets the accountsBlockAddingNonMicrosoftAccountEmail property value. Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account.
func (m *Windows81GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmail()(*bool) {
    return m.accountsBlockAddingNonMicrosoftAccountEmail
}
// GetApplyOnlyToWindows81 gets the applyOnlyToWindows81 property value. Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
func (m *Windows81GeneralConfiguration) GetApplyOnlyToWindows81()(*bool) {
    return m.applyOnlyToWindows81
}
// GetBrowserBlockAutofill gets the browserBlockAutofill property value. Indicates whether or not to block auto fill.
func (m *Windows81GeneralConfiguration) GetBrowserBlockAutofill()(*bool) {
    return m.browserBlockAutofill
}
// GetBrowserBlockAutomaticDetectionOfIntranetSites gets the browserBlockAutomaticDetectionOfIntranetSites property value. Indicates whether or not to block automatic detection of Intranet sites.
func (m *Windows81GeneralConfiguration) GetBrowserBlockAutomaticDetectionOfIntranetSites()(*bool) {
    return m.browserBlockAutomaticDetectionOfIntranetSites
}
// GetBrowserBlockEnterpriseModeAccess gets the browserBlockEnterpriseModeAccess property value. Indicates whether or not to block enterprise mode access.
func (m *Windows81GeneralConfiguration) GetBrowserBlockEnterpriseModeAccess()(*bool) {
    return m.browserBlockEnterpriseModeAccess
}
// GetBrowserBlockJavaScript gets the browserBlockJavaScript property value. Indicates whether or not to Block the user from using JavaScript.
func (m *Windows81GeneralConfiguration) GetBrowserBlockJavaScript()(*bool) {
    return m.browserBlockJavaScript
}
// GetBrowserBlockPlugins gets the browserBlockPlugins property value. Indicates whether or not to block plug-ins.
func (m *Windows81GeneralConfiguration) GetBrowserBlockPlugins()(*bool) {
    return m.browserBlockPlugins
}
// GetBrowserBlockPopups gets the browserBlockPopups property value. Indicates whether or not to block popups.
func (m *Windows81GeneralConfiguration) GetBrowserBlockPopups()(*bool) {
    return m.browserBlockPopups
}
// GetBrowserBlockSendingDoNotTrackHeader gets the browserBlockSendingDoNotTrackHeader property value. Indicates whether or not to Block the user from sending the do not track header.
func (m *Windows81GeneralConfiguration) GetBrowserBlockSendingDoNotTrackHeader()(*bool) {
    return m.browserBlockSendingDoNotTrackHeader
}
// GetBrowserBlockSingleWordEntryOnIntranetSites gets the browserBlockSingleWordEntryOnIntranetSites property value. Indicates whether or not to block a single word entry on Intranet sites.
func (m *Windows81GeneralConfiguration) GetBrowserBlockSingleWordEntryOnIntranetSites()(*bool) {
    return m.browserBlockSingleWordEntryOnIntranetSites
}
// GetBrowserEnterpriseModeSiteListLocation gets the browserEnterpriseModeSiteListLocation property value. The enterprise mode site list location. Could be a local file, local network or http location.
func (m *Windows81GeneralConfiguration) GetBrowserEnterpriseModeSiteListLocation()(*string) {
    return m.browserEnterpriseModeSiteListLocation
}
// GetBrowserInternetSecurityLevel gets the browserInternetSecurityLevel property value. Possible values for internet site security level.
func (m *Windows81GeneralConfiguration) GetBrowserInternetSecurityLevel()(*InternetSiteSecurityLevel) {
    return m.browserInternetSecurityLevel
}
// GetBrowserIntranetSecurityLevel gets the browserIntranetSecurityLevel property value. Possible values for site security level.
func (m *Windows81GeneralConfiguration) GetBrowserIntranetSecurityLevel()(*SiteSecurityLevel) {
    return m.browserIntranetSecurityLevel
}
// GetBrowserLoggingReportLocation gets the browserLoggingReportLocation property value. The logging report location.
func (m *Windows81GeneralConfiguration) GetBrowserLoggingReportLocation()(*string) {
    return m.browserLoggingReportLocation
}
// GetBrowserRequireFirewall gets the browserRequireFirewall property value. Indicates whether or not to require a firewall.
func (m *Windows81GeneralConfiguration) GetBrowserRequireFirewall()(*bool) {
    return m.browserRequireFirewall
}
// GetBrowserRequireFraudWarning gets the browserRequireFraudWarning property value. Indicates whether or not to require fraud warning.
func (m *Windows81GeneralConfiguration) GetBrowserRequireFraudWarning()(*bool) {
    return m.browserRequireFraudWarning
}
// GetBrowserRequireHighSecurityForRestrictedSites gets the browserRequireHighSecurityForRestrictedSites property value. Indicates whether or not to require high security for restricted sites.
func (m *Windows81GeneralConfiguration) GetBrowserRequireHighSecurityForRestrictedSites()(*bool) {
    return m.browserRequireHighSecurityForRestrictedSites
}
// GetBrowserRequireSmartScreen gets the browserRequireSmartScreen property value. Indicates whether or not to require the user to use the smart screen filter.
func (m *Windows81GeneralConfiguration) GetBrowserRequireSmartScreen()(*bool) {
    return m.browserRequireSmartScreen
}
// GetBrowserTrustedSitesSecurityLevel gets the browserTrustedSitesSecurityLevel property value. Possible values for site security level.
func (m *Windows81GeneralConfiguration) GetBrowserTrustedSitesSecurityLevel()(*SiteSecurityLevel) {
    return m.browserTrustedSitesSecurityLevel
}
// GetCellularBlockDataRoaming gets the cellularBlockDataRoaming property value. Indicates whether or not to block data roaming.
func (m *Windows81GeneralConfiguration) GetCellularBlockDataRoaming()(*bool) {
    return m.cellularBlockDataRoaming
}
// GetDiagnosticsBlockDataSubmission gets the diagnosticsBlockDataSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *Windows81GeneralConfiguration) GetDiagnosticsBlockDataSubmission()(*bool) {
    return m.diagnosticsBlockDataSubmission
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows81GeneralConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["accountsBlockAddingNonMicrosoftAccountEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountsBlockAddingNonMicrosoftAccountEmail(val)
        }
        return nil
    }
    res["applyOnlyToWindows81"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyOnlyToWindows81(val)
        }
        return nil
    }
    res["browserBlockAutofill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserBlockAutofill(val)
        }
        return nil
    }
    res["browserBlockAutomaticDetectionOfIntranetSites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserBlockAutomaticDetectionOfIntranetSites(val)
        }
        return nil
    }
    res["browserBlockEnterpriseModeAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserBlockEnterpriseModeAccess(val)
        }
        return nil
    }
    res["browserBlockJavaScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserBlockJavaScript(val)
        }
        return nil
    }
    res["browserBlockPlugins"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserBlockPlugins(val)
        }
        return nil
    }
    res["browserBlockPopups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserBlockPopups(val)
        }
        return nil
    }
    res["browserBlockSendingDoNotTrackHeader"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserBlockSendingDoNotTrackHeader(val)
        }
        return nil
    }
    res["browserBlockSingleWordEntryOnIntranetSites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserBlockSingleWordEntryOnIntranetSites(val)
        }
        return nil
    }
    res["browserEnterpriseModeSiteListLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserEnterpriseModeSiteListLocation(val)
        }
        return nil
    }
    res["browserInternetSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInternetSiteSecurityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserInternetSecurityLevel(val.(*InternetSiteSecurityLevel))
        }
        return nil
    }
    res["browserIntranetSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSiteSecurityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserIntranetSecurityLevel(val.(*SiteSecurityLevel))
        }
        return nil
    }
    res["browserLoggingReportLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserLoggingReportLocation(val)
        }
        return nil
    }
    res["browserRequireFirewall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserRequireFirewall(val)
        }
        return nil
    }
    res["browserRequireFraudWarning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserRequireFraudWarning(val)
        }
        return nil
    }
    res["browserRequireHighSecurityForRestrictedSites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserRequireHighSecurityForRestrictedSites(val)
        }
        return nil
    }
    res["browserRequireSmartScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserRequireSmartScreen(val)
        }
        return nil
    }
    res["browserTrustedSitesSecurityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSiteSecurityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrowserTrustedSitesSecurityLevel(val.(*SiteSecurityLevel))
        }
        return nil
    }
    res["cellularBlockDataRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockDataRoaming(val)
        }
        return nil
    }
    res["diagnosticsBlockDataSubmission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticsBlockDataSubmission(val)
        }
        return nil
    }
    res["passwordBlockPicturePasswordAndPin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordBlockPicturePasswordAndPin(val)
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
    res["passwordMinutesOfInactivityBeforeScreenTimeout"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordMinutesOfInactivityBeforeScreenTimeout(val)
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
    res["passwordSignInFailureCountBeforeFactoryReset"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordSignInFailureCountBeforeFactoryReset(val)
        }
        return nil
    }
    res["storageRequireDeviceEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageRequireDeviceEncryption(val)
        }
        return nil
    }
    res["updatesRequireAutomaticUpdates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdatesRequireAutomaticUpdates(val)
        }
        return nil
    }
    res["userAccountControlSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsUserAccountControlSettings)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountControlSettings(val.(*WindowsUserAccountControlSettings))
        }
        return nil
    }
    res["workFoldersUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkFoldersUrl(val)
        }
        return nil
    }
    return res
}
// GetPasswordBlockPicturePasswordAndPin gets the passwordBlockPicturePasswordAndPin property value. Indicates whether or not to Block the user from using a pictures password and pin.
func (m *Windows81GeneralConfiguration) GetPasswordBlockPicturePasswordAndPin()(*bool) {
    return m.passwordBlockPicturePasswordAndPin
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Password expiration in days.
func (m *Windows81GeneralConfiguration) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows81GeneralConfiguration) GetPasswordMinimumCharacterSetCount()(*int32) {
    return m.passwordMinimumCharacterSetCount
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. The minimum password length.
func (m *Windows81GeneralConfiguration) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. The minutes of inactivity before the screen times out.
func (m *Windows81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.passwordMinutesOfInactivityBeforeScreenTimeout
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent re-use of. Valid values 0 to 24
func (m *Windows81GeneralConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows81GeneralConfiguration) GetPasswordRequiredType()(*RequiredPasswordType) {
    return m.passwordRequiredType
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. The number of sign in failures before factory reset.
func (m *Windows81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.passwordSignInFailureCountBeforeFactoryReset
}
// GetStorageRequireDeviceEncryption gets the storageRequireDeviceEncryption property value. Indicates whether or not to require encryption on a mobile device.
func (m *Windows81GeneralConfiguration) GetStorageRequireDeviceEncryption()(*bool) {
    return m.storageRequireDeviceEncryption
}
// GetUpdatesRequireAutomaticUpdates gets the updatesRequireAutomaticUpdates property value. Indicates whether or not to require automatic updates.
func (m *Windows81GeneralConfiguration) GetUpdatesRequireAutomaticUpdates()(*bool) {
    return m.updatesRequireAutomaticUpdates
}
// GetUserAccountControlSettings gets the userAccountControlSettings property value. Possible values for Windows user account control settings.
func (m *Windows81GeneralConfiguration) GetUserAccountControlSettings()(*WindowsUserAccountControlSettings) {
    return m.userAccountControlSettings
}
// GetWorkFoldersUrl gets the workFoldersUrl property value. The work folders url.
func (m *Windows81GeneralConfiguration) GetWorkFoldersUrl()(*string) {
    return m.workFoldersUrl
}
// Serialize serializes information the current object
func (m *Windows81GeneralConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountsBlockAddingNonMicrosoftAccountEmail", m.GetAccountsBlockAddingNonMicrosoftAccountEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserBlockAutofill", m.GetBrowserBlockAutofill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserBlockAutomaticDetectionOfIntranetSites", m.GetBrowserBlockAutomaticDetectionOfIntranetSites())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserBlockEnterpriseModeAccess", m.GetBrowserBlockEnterpriseModeAccess())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserBlockJavaScript", m.GetBrowserBlockJavaScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserBlockPlugins", m.GetBrowserBlockPlugins())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserBlockPopups", m.GetBrowserBlockPopups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserBlockSendingDoNotTrackHeader", m.GetBrowserBlockSendingDoNotTrackHeader())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserBlockSingleWordEntryOnIntranetSites", m.GetBrowserBlockSingleWordEntryOnIntranetSites())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("browserEnterpriseModeSiteListLocation", m.GetBrowserEnterpriseModeSiteListLocation())
        if err != nil {
            return err
        }
    }
    if m.GetBrowserInternetSecurityLevel() != nil {
        cast := (*m.GetBrowserInternetSecurityLevel()).String()
        err = writer.WriteStringValue("browserInternetSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetBrowserIntranetSecurityLevel() != nil {
        cast := (*m.GetBrowserIntranetSecurityLevel()).String()
        err = writer.WriteStringValue("browserIntranetSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("browserLoggingReportLocation", m.GetBrowserLoggingReportLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserRequireFirewall", m.GetBrowserRequireFirewall())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserRequireFraudWarning", m.GetBrowserRequireFraudWarning())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserRequireHighSecurityForRestrictedSites", m.GetBrowserRequireHighSecurityForRestrictedSites())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("browserRequireSmartScreen", m.GetBrowserRequireSmartScreen())
        if err != nil {
            return err
        }
    }
    if m.GetBrowserTrustedSitesSecurityLevel() != nil {
        cast := (*m.GetBrowserTrustedSitesSecurityLevel()).String()
        err = writer.WriteStringValue("browserTrustedSitesSecurityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockDataRoaming", m.GetCellularBlockDataRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("diagnosticsBlockDataSubmission", m.GetDiagnosticsBlockDataSubmission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("passwordBlockPicturePasswordAndPin", m.GetPasswordBlockPicturePasswordAndPin())
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
        err = writer.WriteInt32Value("passwordMinutesOfInactivityBeforeScreenTimeout", m.GetPasswordMinutesOfInactivityBeforeScreenTimeout())
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
    if m.GetPasswordRequiredType() != nil {
        cast := (*m.GetPasswordRequiredType()).String()
        err = writer.WriteStringValue("passwordRequiredType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRequireDeviceEncryption", m.GetStorageRequireDeviceEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("updatesRequireAutomaticUpdates", m.GetUpdatesRequireAutomaticUpdates())
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountControlSettings() != nil {
        cast := (*m.GetUserAccountControlSettings()).String()
        err = writer.WriteStringValue("userAccountControlSettings", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("workFoldersUrl", m.GetWorkFoldersUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountsBlockAddingNonMicrosoftAccountEmail sets the accountsBlockAddingNonMicrosoftAccountEmail property value. Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account.
func (m *Windows81GeneralConfiguration) SetAccountsBlockAddingNonMicrosoftAccountEmail(value *bool)() {
    m.accountsBlockAddingNonMicrosoftAccountEmail = value
}
// SetApplyOnlyToWindows81 sets the applyOnlyToWindows81 property value. Value indicating whether this policy only applies to Windows 8.1. This property is read-only.
func (m *Windows81GeneralConfiguration) SetApplyOnlyToWindows81(value *bool)() {
    m.applyOnlyToWindows81 = value
}
// SetBrowserBlockAutofill sets the browserBlockAutofill property value. Indicates whether or not to block auto fill.
func (m *Windows81GeneralConfiguration) SetBrowserBlockAutofill(value *bool)() {
    m.browserBlockAutofill = value
}
// SetBrowserBlockAutomaticDetectionOfIntranetSites sets the browserBlockAutomaticDetectionOfIntranetSites property value. Indicates whether or not to block automatic detection of Intranet sites.
func (m *Windows81GeneralConfiguration) SetBrowserBlockAutomaticDetectionOfIntranetSites(value *bool)() {
    m.browserBlockAutomaticDetectionOfIntranetSites = value
}
// SetBrowserBlockEnterpriseModeAccess sets the browserBlockEnterpriseModeAccess property value. Indicates whether or not to block enterprise mode access.
func (m *Windows81GeneralConfiguration) SetBrowserBlockEnterpriseModeAccess(value *bool)() {
    m.browserBlockEnterpriseModeAccess = value
}
// SetBrowserBlockJavaScript sets the browserBlockJavaScript property value. Indicates whether or not to Block the user from using JavaScript.
func (m *Windows81GeneralConfiguration) SetBrowserBlockJavaScript(value *bool)() {
    m.browserBlockJavaScript = value
}
// SetBrowserBlockPlugins sets the browserBlockPlugins property value. Indicates whether or not to block plug-ins.
func (m *Windows81GeneralConfiguration) SetBrowserBlockPlugins(value *bool)() {
    m.browserBlockPlugins = value
}
// SetBrowserBlockPopups sets the browserBlockPopups property value. Indicates whether or not to block popups.
func (m *Windows81GeneralConfiguration) SetBrowserBlockPopups(value *bool)() {
    m.browserBlockPopups = value
}
// SetBrowserBlockSendingDoNotTrackHeader sets the browserBlockSendingDoNotTrackHeader property value. Indicates whether or not to Block the user from sending the do not track header.
func (m *Windows81GeneralConfiguration) SetBrowserBlockSendingDoNotTrackHeader(value *bool)() {
    m.browserBlockSendingDoNotTrackHeader = value
}
// SetBrowserBlockSingleWordEntryOnIntranetSites sets the browserBlockSingleWordEntryOnIntranetSites property value. Indicates whether or not to block a single word entry on Intranet sites.
func (m *Windows81GeneralConfiguration) SetBrowserBlockSingleWordEntryOnIntranetSites(value *bool)() {
    m.browserBlockSingleWordEntryOnIntranetSites = value
}
// SetBrowserEnterpriseModeSiteListLocation sets the browserEnterpriseModeSiteListLocation property value. The enterprise mode site list location. Could be a local file, local network or http location.
func (m *Windows81GeneralConfiguration) SetBrowserEnterpriseModeSiteListLocation(value *string)() {
    m.browserEnterpriseModeSiteListLocation = value
}
// SetBrowserInternetSecurityLevel sets the browserInternetSecurityLevel property value. Possible values for internet site security level.
func (m *Windows81GeneralConfiguration) SetBrowserInternetSecurityLevel(value *InternetSiteSecurityLevel)() {
    m.browserInternetSecurityLevel = value
}
// SetBrowserIntranetSecurityLevel sets the browserIntranetSecurityLevel property value. Possible values for site security level.
func (m *Windows81GeneralConfiguration) SetBrowserIntranetSecurityLevel(value *SiteSecurityLevel)() {
    m.browserIntranetSecurityLevel = value
}
// SetBrowserLoggingReportLocation sets the browserLoggingReportLocation property value. The logging report location.
func (m *Windows81GeneralConfiguration) SetBrowserLoggingReportLocation(value *string)() {
    m.browserLoggingReportLocation = value
}
// SetBrowserRequireFirewall sets the browserRequireFirewall property value. Indicates whether or not to require a firewall.
func (m *Windows81GeneralConfiguration) SetBrowserRequireFirewall(value *bool)() {
    m.browserRequireFirewall = value
}
// SetBrowserRequireFraudWarning sets the browserRequireFraudWarning property value. Indicates whether or not to require fraud warning.
func (m *Windows81GeneralConfiguration) SetBrowserRequireFraudWarning(value *bool)() {
    m.browserRequireFraudWarning = value
}
// SetBrowserRequireHighSecurityForRestrictedSites sets the browserRequireHighSecurityForRestrictedSites property value. Indicates whether or not to require high security for restricted sites.
func (m *Windows81GeneralConfiguration) SetBrowserRequireHighSecurityForRestrictedSites(value *bool)() {
    m.browserRequireHighSecurityForRestrictedSites = value
}
// SetBrowserRequireSmartScreen sets the browserRequireSmartScreen property value. Indicates whether or not to require the user to use the smart screen filter.
func (m *Windows81GeneralConfiguration) SetBrowserRequireSmartScreen(value *bool)() {
    m.browserRequireSmartScreen = value
}
// SetBrowserTrustedSitesSecurityLevel sets the browserTrustedSitesSecurityLevel property value. Possible values for site security level.
func (m *Windows81GeneralConfiguration) SetBrowserTrustedSitesSecurityLevel(value *SiteSecurityLevel)() {
    m.browserTrustedSitesSecurityLevel = value
}
// SetCellularBlockDataRoaming sets the cellularBlockDataRoaming property value. Indicates whether or not to block data roaming.
func (m *Windows81GeneralConfiguration) SetCellularBlockDataRoaming(value *bool)() {
    m.cellularBlockDataRoaming = value
}
// SetDiagnosticsBlockDataSubmission sets the diagnosticsBlockDataSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *Windows81GeneralConfiguration) SetDiagnosticsBlockDataSubmission(value *bool)() {
    m.diagnosticsBlockDataSubmission = value
}
// SetPasswordBlockPicturePasswordAndPin sets the passwordBlockPicturePasswordAndPin property value. Indicates whether or not to Block the user from using a pictures password and pin.
func (m *Windows81GeneralConfiguration) SetPasswordBlockPicturePasswordAndPin(value *bool)() {
    m.passwordBlockPicturePasswordAndPin = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Password expiration in days.
func (m *Windows81GeneralConfiguration) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows81GeneralConfiguration) SetPasswordMinimumCharacterSetCount(value *int32)() {
    m.passwordMinimumCharacterSetCount = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. The minimum password length.
func (m *Windows81GeneralConfiguration) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. The minutes of inactivity before the screen times out.
func (m *Windows81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.passwordMinutesOfInactivityBeforeScreenTimeout = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent re-use of. Valid values 0 to 24
func (m *Windows81GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows81GeneralConfiguration) SetPasswordRequiredType(value *RequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. The number of sign in failures before factory reset.
func (m *Windows81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.passwordSignInFailureCountBeforeFactoryReset = value
}
// SetStorageRequireDeviceEncryption sets the storageRequireDeviceEncryption property value. Indicates whether or not to require encryption on a mobile device.
func (m *Windows81GeneralConfiguration) SetStorageRequireDeviceEncryption(value *bool)() {
    m.storageRequireDeviceEncryption = value
}
// SetUpdatesRequireAutomaticUpdates sets the updatesRequireAutomaticUpdates property value. Indicates whether or not to require automatic updates.
func (m *Windows81GeneralConfiguration) SetUpdatesRequireAutomaticUpdates(value *bool)() {
    m.updatesRequireAutomaticUpdates = value
}
// SetUserAccountControlSettings sets the userAccountControlSettings property value. Possible values for Windows user account control settings.
func (m *Windows81GeneralConfiguration) SetUserAccountControlSettings(value *WindowsUserAccountControlSettings)() {
    m.userAccountControlSettings = value
}
// SetWorkFoldersUrl sets the workFoldersUrl property value. The work folders url.
func (m *Windows81GeneralConfiguration) SetWorkFoldersUrl(value *string)() {
    m.workFoldersUrl = value
}
