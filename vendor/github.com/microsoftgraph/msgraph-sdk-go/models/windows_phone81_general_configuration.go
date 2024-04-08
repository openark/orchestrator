package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsPhone81GeneralConfiguration 
type WindowsPhone81GeneralConfiguration struct {
    DeviceConfiguration
    // Value indicating whether this policy only applies to Windows Phone 8.1. This property is read-only.
    applyOnlyToWindowsPhone81 *bool
    // Indicates whether or not to block copy paste.
    appsBlockCopyPaste *bool
    // Indicates whether or not to block bluetooth.
    bluetoothBlocked *bool
    // Indicates whether or not to block camera.
    cameraBlocked *bool
    // Indicates whether or not to block Wi-Fi tethering. Has no impact if Wi-Fi is blocked.
    cellularBlockWifiTethering *bool
    // Possible values of the compliance app list.
    compliantAppListType *AppListType
    // List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
    compliantAppsList []AppListItemable
    // Indicates whether or not to block diagnostic data submission.
    diagnosticDataBlockSubmission *bool
    // Indicates whether or not to block custom email accounts.
    emailBlockAddingAccounts *bool
    // Indicates whether or not to block location services.
    locationServicesBlocked *bool
    // Indicates whether or not to block using a Microsoft Account.
    microsoftAccountBlocked *bool
    // Indicates whether or not to block Near-Field Communication.
    nfcBlocked *bool
    // Indicates whether or not to block syncing the calendar.
    passwordBlockSimple *bool
    // Number of days before the password expires.
    passwordExpirationDays *int32
    // Number of character sets a password must contain.
    passwordMinimumCharacterSetCount *int32
    // Minimum length of passwords.
    passwordMinimumLength *int32
    // Minutes of inactivity before screen timeout.
    passwordMinutesOfInactivityBeforeScreenTimeout *int32
    // Number of previous passwords to block. Valid values 0 to 24
    passwordPreviousPasswordBlockCount *int32
    // Indicates whether or not to require a password.
    passwordRequired *bool
    // Possible values of required passwords.
    passwordRequiredType *RequiredPasswordType
    // Number of sign in failures allowed before factory reset.
    passwordSignInFailureCountBeforeFactoryReset *int32
    // Indicates whether or not to block screenshots.
    screenCaptureBlocked *bool
    // Indicates whether or not to block removable storage.
    storageBlockRemovableStorage *bool
    // Indicates whether or not to require encryption.
    storageRequireEncryption *bool
    // Indicates whether or not to block the web browser.
    webBrowserBlocked *bool
    // Indicates whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
    wifiBlockAutomaticConnectHotspots *bool
    // Indicates whether or not to block Wi-Fi.
    wifiBlocked *bool
    // Indicates whether or not to block Wi-Fi hotspot reporting. Has no impact if Wi-Fi is blocked.
    wifiBlockHotspotReporting *bool
    // Indicates whether or not to block the Windows Store.
    windowsStoreBlocked *bool
}
// NewWindowsPhone81GeneralConfiguration instantiates a new WindowsPhone81GeneralConfiguration and sets the default values.
func NewWindowsPhone81GeneralConfiguration()(*WindowsPhone81GeneralConfiguration) {
    m := &WindowsPhone81GeneralConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windowsPhone81GeneralConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsPhone81GeneralConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsPhone81GeneralConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsPhone81GeneralConfiguration(), nil
}
// GetApplyOnlyToWindowsPhone81 gets the applyOnlyToWindowsPhone81 property value. Value indicating whether this policy only applies to Windows Phone 8.1. This property is read-only.
func (m *WindowsPhone81GeneralConfiguration) GetApplyOnlyToWindowsPhone81()(*bool) {
    return m.applyOnlyToWindowsPhone81
}
// GetAppsBlockCopyPaste gets the appsBlockCopyPaste property value. Indicates whether or not to block copy paste.
func (m *WindowsPhone81GeneralConfiguration) GetAppsBlockCopyPaste()(*bool) {
    return m.appsBlockCopyPaste
}
// GetBluetoothBlocked gets the bluetoothBlocked property value. Indicates whether or not to block bluetooth.
func (m *WindowsPhone81GeneralConfiguration) GetBluetoothBlocked()(*bool) {
    return m.bluetoothBlocked
}
// GetCameraBlocked gets the cameraBlocked property value. Indicates whether or not to block camera.
func (m *WindowsPhone81GeneralConfiguration) GetCameraBlocked()(*bool) {
    return m.cameraBlocked
}
// GetCellularBlockWifiTethering gets the cellularBlockWifiTethering property value. Indicates whether or not to block Wi-Fi tethering. Has no impact if Wi-Fi is blocked.
func (m *WindowsPhone81GeneralConfiguration) GetCellularBlockWifiTethering()(*bool) {
    return m.cellularBlockWifiTethering
}
// GetCompliantAppListType gets the compliantAppListType property value. Possible values of the compliance app list.
func (m *WindowsPhone81GeneralConfiguration) GetCompliantAppListType()(*AppListType) {
    return m.compliantAppListType
}
// GetCompliantAppsList gets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *WindowsPhone81GeneralConfiguration) GetCompliantAppsList()([]AppListItemable) {
    return m.compliantAppsList
}
// GetDiagnosticDataBlockSubmission gets the diagnosticDataBlockSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *WindowsPhone81GeneralConfiguration) GetDiagnosticDataBlockSubmission()(*bool) {
    return m.diagnosticDataBlockSubmission
}
// GetEmailBlockAddingAccounts gets the emailBlockAddingAccounts property value. Indicates whether or not to block custom email accounts.
func (m *WindowsPhone81GeneralConfiguration) GetEmailBlockAddingAccounts()(*bool) {
    return m.emailBlockAddingAccounts
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsPhone81GeneralConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["applyOnlyToWindowsPhone81"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyOnlyToWindowsPhone81(val)
        }
        return nil
    }
    res["appsBlockCopyPaste"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsBlockCopyPaste(val)
        }
        return nil
    }
    res["bluetoothBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlocked(val)
        }
        return nil
    }
    res["cameraBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCameraBlocked(val)
        }
        return nil
    }
    res["cellularBlockWifiTethering"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockWifiTethering(val)
        }
        return nil
    }
    res["compliantAppListType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppListType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantAppListType(val.(*AppListType))
        }
        return nil
    }
    res["compliantAppsList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppListItemable, len(val))
            for i, v := range val {
                res[i] = v.(AppListItemable)
            }
            m.SetCompliantAppsList(res)
        }
        return nil
    }
    res["diagnosticDataBlockSubmission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticDataBlockSubmission(val)
        }
        return nil
    }
    res["emailBlockAddingAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailBlockAddingAccounts(val)
        }
        return nil
    }
    res["locationServicesBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationServicesBlocked(val)
        }
        return nil
    }
    res["microsoftAccountBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftAccountBlocked(val)
        }
        return nil
    }
    res["nfcBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNfcBlocked(val)
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
    res["screenCaptureBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    res["storageBlockRemovableStorage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageBlockRemovableStorage(val)
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
    res["webBrowserBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebBrowserBlocked(val)
        }
        return nil
    }
    res["wifiBlockAutomaticConnectHotspots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBlockAutomaticConnectHotspots(val)
        }
        return nil
    }
    res["wifiBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBlocked(val)
        }
        return nil
    }
    res["wifiBlockHotspotReporting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBlockHotspotReporting(val)
        }
        return nil
    }
    res["windowsStoreBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsStoreBlocked(val)
        }
        return nil
    }
    return res
}
// GetLocationServicesBlocked gets the locationServicesBlocked property value. Indicates whether or not to block location services.
func (m *WindowsPhone81GeneralConfiguration) GetLocationServicesBlocked()(*bool) {
    return m.locationServicesBlocked
}
// GetMicrosoftAccountBlocked gets the microsoftAccountBlocked property value. Indicates whether or not to block using a Microsoft Account.
func (m *WindowsPhone81GeneralConfiguration) GetMicrosoftAccountBlocked()(*bool) {
    return m.microsoftAccountBlocked
}
// GetNfcBlocked gets the nfcBlocked property value. Indicates whether or not to block Near-Field Communication.
func (m *WindowsPhone81GeneralConfiguration) GetNfcBlocked()(*bool) {
    return m.nfcBlocked
}
// GetPasswordBlockSimple gets the passwordBlockSimple property value. Indicates whether or not to block syncing the calendar.
func (m *WindowsPhone81GeneralConfiguration) GetPasswordBlockSimple()(*bool) {
    return m.passwordBlockSimple
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. Number of days before the password expires.
func (m *WindowsPhone81GeneralConfiguration) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. Number of character sets a password must contain.
func (m *WindowsPhone81GeneralConfiguration) GetPasswordMinimumCharacterSetCount()(*int32) {
    return m.passwordMinimumCharacterSetCount
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. Minimum length of passwords.
func (m *WindowsPhone81GeneralConfiguration) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before screen timeout.
func (m *WindowsPhone81GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.passwordMinutesOfInactivityBeforeScreenTimeout
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *WindowsPhone81GeneralConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequired gets the passwordRequired property value. Indicates whether or not to require a password.
func (m *WindowsPhone81GeneralConfiguration) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Possible values of required passwords.
func (m *WindowsPhone81GeneralConfiguration) GetPasswordRequiredType()(*RequiredPasswordType) {
    return m.passwordRequiredType
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset.
func (m *WindowsPhone81GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.passwordSignInFailureCountBeforeFactoryReset
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to block screenshots.
func (m *WindowsPhone81GeneralConfiguration) GetScreenCaptureBlocked()(*bool) {
    return m.screenCaptureBlocked
}
// GetStorageBlockRemovableStorage gets the storageBlockRemovableStorage property value. Indicates whether or not to block removable storage.
func (m *WindowsPhone81GeneralConfiguration) GetStorageBlockRemovableStorage()(*bool) {
    return m.storageBlockRemovableStorage
}
// GetStorageRequireEncryption gets the storageRequireEncryption property value. Indicates whether or not to require encryption.
func (m *WindowsPhone81GeneralConfiguration) GetStorageRequireEncryption()(*bool) {
    return m.storageRequireEncryption
}
// GetWebBrowserBlocked gets the webBrowserBlocked property value. Indicates whether or not to block the web browser.
func (m *WindowsPhone81GeneralConfiguration) GetWebBrowserBlocked()(*bool) {
    return m.webBrowserBlocked
}
// GetWifiBlockAutomaticConnectHotspots gets the wifiBlockAutomaticConnectHotspots property value. Indicates whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
func (m *WindowsPhone81GeneralConfiguration) GetWifiBlockAutomaticConnectHotspots()(*bool) {
    return m.wifiBlockAutomaticConnectHotspots
}
// GetWifiBlocked gets the wifiBlocked property value. Indicates whether or not to block Wi-Fi.
func (m *WindowsPhone81GeneralConfiguration) GetWifiBlocked()(*bool) {
    return m.wifiBlocked
}
// GetWifiBlockHotspotReporting gets the wifiBlockHotspotReporting property value. Indicates whether or not to block Wi-Fi hotspot reporting. Has no impact if Wi-Fi is blocked.
func (m *WindowsPhone81GeneralConfiguration) GetWifiBlockHotspotReporting()(*bool) {
    return m.wifiBlockHotspotReporting
}
// GetWindowsStoreBlocked gets the windowsStoreBlocked property value. Indicates whether or not to block the Windows Store.
func (m *WindowsPhone81GeneralConfiguration) GetWindowsStoreBlocked()(*bool) {
    return m.windowsStoreBlocked
}
// Serialize serializes information the current object
func (m *WindowsPhone81GeneralConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("appsBlockCopyPaste", m.GetAppsBlockCopyPaste())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlocked", m.GetBluetoothBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cameraBlocked", m.GetCameraBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockWifiTethering", m.GetCellularBlockWifiTethering())
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppListType() != nil {
        cast := (*m.GetCompliantAppListType()).String()
        err = writer.WriteStringValue("compliantAppListType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetCompliantAppsList() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCompliantAppsList()))
        for i, v := range m.GetCompliantAppsList() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("compliantAppsList", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("diagnosticDataBlockSubmission", m.GetDiagnosticDataBlockSubmission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("emailBlockAddingAccounts", m.GetEmailBlockAddingAccounts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("locationServicesBlocked", m.GetLocationServicesBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftAccountBlocked", m.GetMicrosoftAccountBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("nfcBlocked", m.GetNfcBlocked())
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
        err = writer.WriteInt32Value("passwordSignInFailureCountBeforeFactoryReset", m.GetPasswordSignInFailureCountBeforeFactoryReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageBlockRemovableStorage", m.GetStorageBlockRemovableStorage())
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
    {
        err = writer.WriteBoolValue("webBrowserBlocked", m.GetWebBrowserBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiBlockAutomaticConnectHotspots", m.GetWifiBlockAutomaticConnectHotspots())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiBlocked", m.GetWifiBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wifiBlockHotspotReporting", m.GetWifiBlockHotspotReporting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsStoreBlocked", m.GetWindowsStoreBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplyOnlyToWindowsPhone81 sets the applyOnlyToWindowsPhone81 property value. Value indicating whether this policy only applies to Windows Phone 8.1. This property is read-only.
func (m *WindowsPhone81GeneralConfiguration) SetApplyOnlyToWindowsPhone81(value *bool)() {
    m.applyOnlyToWindowsPhone81 = value
}
// SetAppsBlockCopyPaste sets the appsBlockCopyPaste property value. Indicates whether or not to block copy paste.
func (m *WindowsPhone81GeneralConfiguration) SetAppsBlockCopyPaste(value *bool)() {
    m.appsBlockCopyPaste = value
}
// SetBluetoothBlocked sets the bluetoothBlocked property value. Indicates whether or not to block bluetooth.
func (m *WindowsPhone81GeneralConfiguration) SetBluetoothBlocked(value *bool)() {
    m.bluetoothBlocked = value
}
// SetCameraBlocked sets the cameraBlocked property value. Indicates whether or not to block camera.
func (m *WindowsPhone81GeneralConfiguration) SetCameraBlocked(value *bool)() {
    m.cameraBlocked = value
}
// SetCellularBlockWifiTethering sets the cellularBlockWifiTethering property value. Indicates whether or not to block Wi-Fi tethering. Has no impact if Wi-Fi is blocked.
func (m *WindowsPhone81GeneralConfiguration) SetCellularBlockWifiTethering(value *bool)() {
    m.cellularBlockWifiTethering = value
}
// SetCompliantAppListType sets the compliantAppListType property value. Possible values of the compliance app list.
func (m *WindowsPhone81GeneralConfiguration) SetCompliantAppListType(value *AppListType)() {
    m.compliantAppListType = value
}
// SetCompliantAppsList sets the compliantAppsList property value. List of apps in the compliance (either allow list or block list, controlled by CompliantAppListType). This collection can contain a maximum of 10000 elements.
func (m *WindowsPhone81GeneralConfiguration) SetCompliantAppsList(value []AppListItemable)() {
    m.compliantAppsList = value
}
// SetDiagnosticDataBlockSubmission sets the diagnosticDataBlockSubmission property value. Indicates whether or not to block diagnostic data submission.
func (m *WindowsPhone81GeneralConfiguration) SetDiagnosticDataBlockSubmission(value *bool)() {
    m.diagnosticDataBlockSubmission = value
}
// SetEmailBlockAddingAccounts sets the emailBlockAddingAccounts property value. Indicates whether or not to block custom email accounts.
func (m *WindowsPhone81GeneralConfiguration) SetEmailBlockAddingAccounts(value *bool)() {
    m.emailBlockAddingAccounts = value
}
// SetLocationServicesBlocked sets the locationServicesBlocked property value. Indicates whether or not to block location services.
func (m *WindowsPhone81GeneralConfiguration) SetLocationServicesBlocked(value *bool)() {
    m.locationServicesBlocked = value
}
// SetMicrosoftAccountBlocked sets the microsoftAccountBlocked property value. Indicates whether or not to block using a Microsoft Account.
func (m *WindowsPhone81GeneralConfiguration) SetMicrosoftAccountBlocked(value *bool)() {
    m.microsoftAccountBlocked = value
}
// SetNfcBlocked sets the nfcBlocked property value. Indicates whether or not to block Near-Field Communication.
func (m *WindowsPhone81GeneralConfiguration) SetNfcBlocked(value *bool)() {
    m.nfcBlocked = value
}
// SetPasswordBlockSimple sets the passwordBlockSimple property value. Indicates whether or not to block syncing the calendar.
func (m *WindowsPhone81GeneralConfiguration) SetPasswordBlockSimple(value *bool)() {
    m.passwordBlockSimple = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. Number of days before the password expires.
func (m *WindowsPhone81GeneralConfiguration) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. Number of character sets a password must contain.
func (m *WindowsPhone81GeneralConfiguration) SetPasswordMinimumCharacterSetCount(value *int32)() {
    m.passwordMinimumCharacterSetCount = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. Minimum length of passwords.
func (m *WindowsPhone81GeneralConfiguration) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. Minutes of inactivity before screen timeout.
func (m *WindowsPhone81GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.passwordMinutesOfInactivityBeforeScreenTimeout = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. Number of previous passwords to block. Valid values 0 to 24
func (m *WindowsPhone81GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequired sets the passwordRequired property value. Indicates whether or not to require a password.
func (m *WindowsPhone81GeneralConfiguration) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Possible values of required passwords.
func (m *WindowsPhone81GeneralConfiguration) SetPasswordRequiredType(value *RequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. Number of sign in failures allowed before factory reset.
func (m *WindowsPhone81GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.passwordSignInFailureCountBeforeFactoryReset = value
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to block screenshots.
func (m *WindowsPhone81GeneralConfiguration) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// SetStorageBlockRemovableStorage sets the storageBlockRemovableStorage property value. Indicates whether or not to block removable storage.
func (m *WindowsPhone81GeneralConfiguration) SetStorageBlockRemovableStorage(value *bool)() {
    m.storageBlockRemovableStorage = value
}
// SetStorageRequireEncryption sets the storageRequireEncryption property value. Indicates whether or not to require encryption.
func (m *WindowsPhone81GeneralConfiguration) SetStorageRequireEncryption(value *bool)() {
    m.storageRequireEncryption = value
}
// SetWebBrowserBlocked sets the webBrowserBlocked property value. Indicates whether or not to block the web browser.
func (m *WindowsPhone81GeneralConfiguration) SetWebBrowserBlocked(value *bool)() {
    m.webBrowserBlocked = value
}
// SetWifiBlockAutomaticConnectHotspots sets the wifiBlockAutomaticConnectHotspots property value. Indicates whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
func (m *WindowsPhone81GeneralConfiguration) SetWifiBlockAutomaticConnectHotspots(value *bool)() {
    m.wifiBlockAutomaticConnectHotspots = value
}
// SetWifiBlocked sets the wifiBlocked property value. Indicates whether or not to block Wi-Fi.
func (m *WindowsPhone81GeneralConfiguration) SetWifiBlocked(value *bool)() {
    m.wifiBlocked = value
}
// SetWifiBlockHotspotReporting sets the wifiBlockHotspotReporting property value. Indicates whether or not to block Wi-Fi hotspot reporting. Has no impact if Wi-Fi is blocked.
func (m *WindowsPhone81GeneralConfiguration) SetWifiBlockHotspotReporting(value *bool)() {
    m.wifiBlockHotspotReporting = value
}
// SetWindowsStoreBlocked sets the windowsStoreBlocked property value. Indicates whether or not to block the Windows Store.
func (m *WindowsPhone81GeneralConfiguration) SetWindowsStoreBlocked(value *bool)() {
    m.windowsStoreBlocked = value
}
