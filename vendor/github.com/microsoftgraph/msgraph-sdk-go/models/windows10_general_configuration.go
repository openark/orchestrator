package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10GeneralConfiguration 
type Windows10GeneralConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account.
    accountsBlockAddingNonMicrosoftAccountEmail *bool
    // Indicates whether or not to block the user from selecting an AntiTheft mode preference (Windows 10 Mobile only).
    antiTheftModeBlocked *bool
    // State Management Setting.
    appsAllowTrustedAppsSideloading *StateManagementSetting
    // Indicates whether or not to disable the launch of all apps from Windows Store that came pre-installed or were downloaded.
    appsBlockWindowsStoreOriginatedApps *bool
    // Specify a list of allowed Bluetooth services and profiles in hex formatted strings.
    bluetoothAllowedServices []string
    // Whether or not to Block the user from using bluetooth advertising.
    bluetoothBlockAdvertising *bool
    // Whether or not to Block the user from using bluetooth discoverable mode.
    bluetoothBlockDiscoverableMode *bool
    // Whether or not to Block the user from using bluetooth.
    bluetoothBlocked *bool
    // Whether or not to block specific bundled Bluetooth peripherals to automatically pair with the host device.
    bluetoothBlockPrePairing *bool
    // Whether or not to Block the user from accessing the camera of the device.
    cameraBlocked *bool
    // Whether or not to Block the user from using data over cellular while roaming.
    cellularBlockDataWhenRoaming *bool
    // Whether or not to Block the user from using VPN over cellular.
    cellularBlockVpn *bool
    // Whether or not to Block the user from using VPN when roaming over cellular.
    cellularBlockVpnWhenRoaming *bool
    // Whether or not to Block the user from doing manual root certificate installation.
    certificatesBlockManualRootCertificateInstallation *bool
    // Whether or not to block Connected Devices Service which enables discovery and connection to other devices, remote messaging, remote app sessions and other cross-device experiences.
    connectedDevicesServiceBlocked *bool
    // Whether or not to Block the user from using copy paste.
    copyPasteBlocked *bool
    // Whether or not to Block the user from using Cortana.
    cortanaBlocked *bool
    // Whether or not to block end user access to Defender.
    defenderBlockEndUserAccess *bool
    // Possible values of Cloud Block Level
    defenderCloudBlockLevel *DefenderCloudBlockLevelType
    // Number of days before deleting quarantined malware. Valid values 0 to 90
    defenderDaysBeforeDeletingQuarantinedMalware *int32
    // Gets or sets Defender’s actions to take on detected Malware per threat level.
    defenderDetectedMalwareActions DefenderDetectedMalwareActionsable
    // File extensions to exclude from scans and real time protection.
    defenderFileExtensionsToExclude []string
    // Files and folder to exclude from scans and real time protection.
    defenderFilesAndFoldersToExclude []string
    // Possible values for monitoring file activity.
    defenderMonitorFileActivity *DefenderMonitorFileActivity
    // Processes to exclude from scans and real time protection.
    defenderProcessesToExclude []string
    // Possible values for prompting user for samples submission.
    defenderPromptForSampleSubmission *DefenderPromptForSampleSubmission
    // Indicates whether or not to require behavior monitoring.
    defenderRequireBehaviorMonitoring *bool
    // Indicates whether or not to require cloud protection.
    defenderRequireCloudProtection *bool
    // Indicates whether or not to require network inspection system.
    defenderRequireNetworkInspectionSystem *bool
    // Indicates whether or not to require real time monitoring.
    defenderRequireRealTimeMonitoring *bool
    // Indicates whether or not to scan archive files.
    defenderScanArchiveFiles *bool
    // Indicates whether or not to scan downloads.
    defenderScanDownloads *bool
    // Indicates whether or not to scan incoming mail messages.
    defenderScanIncomingMail *bool
    // Indicates whether or not to scan mapped network drives during full scan.
    defenderScanMappedNetworkDrivesDuringFullScan *bool
    // Max CPU usage percentage during scan. Valid values 0 to 100
    defenderScanMaxCpu *int32
    // Indicates whether or not to scan files opened from a network folder.
    defenderScanNetworkFiles *bool
    // Indicates whether or not to scan removable drives during full scan.
    defenderScanRemovableDrivesDuringFullScan *bool
    // Indicates whether or not to scan scripts loaded in Internet Explorer browser.
    defenderScanScriptsLoadedInInternetExplorer *bool
    // Possible values for system scan type.
    defenderScanType *DefenderScanType
    // The time to perform a daily quick scan.
    defenderScheduledQuickScanTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The defender time for the system scan.
    defenderScheduledScanTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The signature update interval in hours. Specify 0 not to check. Valid values 0 to 24
    defenderSignatureUpdateIntervalInHours *int32
    // Possible values for a weekly schedule.
    defenderSystemScanSchedule *WeeklySchedule
    // State Management Setting.
    developerUnlockSetting *StateManagementSetting
    // Indicates whether or not to Block the user from resetting their phone.
    deviceManagementBlockFactoryResetOnMobile *bool
    // Indicates whether or not to Block the user from doing manual un-enrollment from device management.
    deviceManagementBlockManualUnenroll *bool
    // Allow the device to send diagnostic and usage telemetry data, such as Watson.
    diagnosticsDataSubmissionMode *DiagnosticDataSubmissionMode
    // Allow users to change Start pages on Edge. Use the EdgeHomepageUrls to specify the Start pages that the user would see by default when they open Edge.
    edgeAllowStartPagesModification *bool
    // Indicates whether or not to prevent access to about flags on Edge browser.
    edgeBlockAccessToAboutFlags *bool
    // Block the address bar dropdown functionality in Microsoft Edge. Disable this settings to minimize network connections from Microsoft Edge to Microsoft services.
    edgeBlockAddressBarDropdown *bool
    // Indicates whether or not to block auto fill.
    edgeBlockAutofill *bool
    // Block Microsoft compatibility list in Microsoft Edge. This list from Microsoft helps Edge properly display sites with known compatibility issues.
    edgeBlockCompatibilityList *bool
    // Indicates whether or not to block developer tools in the Edge browser.
    edgeBlockDeveloperTools *bool
    // Indicates whether or not to Block the user from using the Edge browser.
    edgeBlocked *bool
    // Indicates whether or not to block extensions in the Edge browser.
    edgeBlockExtensions *bool
    // Indicates whether or not to block InPrivate browsing on corporate networks, in the Edge browser.
    edgeBlockInPrivateBrowsing *bool
    // Indicates whether or not to Block the user from using JavaScript.
    edgeBlockJavaScript *bool
    // Block the collection of information by Microsoft for live tile creation when users pin a site to Start from Microsoft Edge.
    edgeBlockLiveTileDataCollection *bool
    // Indicates whether or not to Block password manager.
    edgeBlockPasswordManager *bool
    // Indicates whether or not to block popups.
    edgeBlockPopups *bool
    // Indicates whether or not to block the user from using the search suggestions in the address bar.
    edgeBlockSearchSuggestions *bool
    // Indicates whether or not to Block the user from sending the do not track header.
    edgeBlockSendingDoNotTrackHeader *bool
    // Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer. Note: the name of this property is misleading; the property is obsolete, use EdgeSendIntranetTrafficToInternetExplorer instead.
    edgeBlockSendingIntranetTrafficToInternetExplorer *bool
    // Clear browsing data on exiting Microsoft Edge.
    edgeClearBrowsingDataOnExit *bool
    // Possible values to specify which cookies are allowed in Microsoft Edge.
    edgeCookiePolicy *EdgeCookiePolicy
    // Block the Microsoft web page that opens on the first use of Microsoft Edge. This policy allows enterprises, like those enrolled in zero emissions configurations, to block this page.
    edgeDisableFirstRunPage *bool
    // Indicates the enterprise mode site list location. Could be a local file, local network or http location.
    edgeEnterpriseModeSiteListLocation *string
    // The first run URL for when Edge browser is opened for the first time.
    edgeFirstRunUrl *string
    // The list of URLs for homepages shodwn on MDM-enrolled devices on Edge browser.
    edgeHomepageUrls []string
    // Indicates whether or not to Require the user to use the smart screen filter.
    edgeRequireSmartScreen *bool
    // Allows IT admins to set a default search engine for MDM-Controlled devices. Users can override this and change their default search engine provided the AllowSearchEngineCustomization policy is not set.
    edgeSearchEngine EdgeSearchEngineBaseable
    // Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer.
    edgeSendIntranetTrafficToInternetExplorer *bool
    // Enable favorites sync between Internet Explorer and Microsoft Edge. Additions, deletions, modifications and order changes to favorites are shared between browsers.
    edgeSyncFavoritesWithInternetExplorer *bool
    // Endpoint for discovering cloud printers.
    enterpriseCloudPrintDiscoveryEndPoint *string
    // Maximum number of printers that should be queried from a discovery endpoint. This is a mobile only setting. Valid values 1 to 65535
    enterpriseCloudPrintDiscoveryMaxLimit *int32
    // OAuth resource URI for printer discovery service as configured in Azure portal.
    enterpriseCloudPrintMopriaDiscoveryResourceIdentifier *string
    // Authentication endpoint for acquiring OAuth tokens.
    enterpriseCloudPrintOAuthAuthority *string
    // GUID of a client application authorized to retrieve OAuth tokens from the OAuth Authority.
    enterpriseCloudPrintOAuthClientIdentifier *string
    // OAuth resource URI for print service as configured in the Azure portal.
    enterpriseCloudPrintResourceIdentifier *string
    // Indicates whether or not to enable device discovery UX.
    experienceBlockDeviceDiscovery *bool
    // Indicates whether or not to allow the error dialog from displaying if no SIM card is detected.
    experienceBlockErrorDialogWhenNoSIM *bool
    // Indicates whether or not to enable task switching on the device.
    experienceBlockTaskSwitcher *bool
    // Indicates whether or not to block DVR and broadcasting.
    gameDvrBlocked *bool
    // Indicates whether or not to Block the user from using internet sharing.
    internetSharingBlocked *bool
    // Indicates whether or not to Block the user from location services.
    locationServicesBlocked *bool
    // Specify whether to show a user-configurable setting to control the screen timeout while on the lock screen of Windows 10 Mobile devices. If this policy is set to Allow, the value set by lockScreenTimeoutInSeconds is ignored.
    lockScreenAllowTimeoutConfiguration *bool
    // Indicates whether or not to block action center notifications over lock screen.
    lockScreenBlockActionCenterNotifications *bool
    // Indicates whether or not the user can interact with Cortana using speech while the system is locked.
    lockScreenBlockCortana *bool
    // Indicates whether to allow toast notifications above the device lock screen.
    lockScreenBlockToastNotifications *bool
    // Set the duration (in seconds) from the screen locking to the screen turning off for Windows 10 Mobile devices. Supported values are 11-1800. Valid values 11 to 1800
    lockScreenTimeoutInSeconds *int32
    // Disables the ability to quickly switch between users that are logged on simultaneously without logging off.
    logonBlockFastUserSwitching *bool
    // Indicates whether or not to Block a Microsoft account.
    microsoftAccountBlocked *bool
    // Indicates whether or not to Block Microsoft account settings sync.
    microsoftAccountBlockSettingsSync *bool
    // If set, proxy settings will be applied to all processes and accounts in the device. Otherwise, it will be applied to the user account that’s enrolled into MDM.
    networkProxyApplySettingsDeviceWide *bool
    // Address to the proxy auto-config (PAC) script you want to use.
    networkProxyAutomaticConfigurationUrl *string
    // Disable automatic detection of settings. If enabled, the system will try to find the path to a proxy auto-config (PAC) script.
    networkProxyDisableAutoDetect *bool
    // Specifies manual proxy server settings.
    networkProxyServer Windows10NetworkProxyServerable
    // Indicates whether or not to Block the user from using near field communication.
    nfcBlocked *bool
    // Gets or sets a value allowing IT admins to prevent apps and features from working with files on OneDrive.
    oneDriveDisableFileSync *bool
    // Specify whether PINs or passwords such as '1111' or '1234' are allowed. For Windows 10 desktops, it also controls the use of picture passwords.
    passwordBlockSimple *bool
    // The password expiration in days. Valid values 0 to 730
    passwordExpirationDays *int32
    // The number of character sets required in the password.
    passwordMinimumCharacterSetCount *int32
    // The minimum password length. Valid values 4 to 16
    passwordMinimumLength *int32
    // The minutes of inactivity before the screen times out.
    passwordMinutesOfInactivityBeforeScreenTimeout *int32
    // The number of previous passwords to prevent reuse of. Valid values 0 to 50
    passwordPreviousPasswordBlockCount *int32
    // Indicates whether or not to require the user to have a password.
    passwordRequired *bool
    // Possible values of required passwords.
    passwordRequiredType *RequiredPasswordType
    // Indicates whether or not to require a password upon resuming from an idle state.
    passwordRequireWhenResumeFromIdleState *bool
    // The number of sign in failures before factory reset. Valid values 0 to 999
    passwordSignInFailureCountBeforeFactoryReset *int32
    // A http or https Url to a jpg, jpeg or png image that needs to be downloaded and used as the Desktop Image or a file Url to a local image on the file system that needs to used as the Desktop Image.
    personalizationDesktopImageUrl *string
    // A http or https Url to a jpg, jpeg or png image that neeeds to be downloaded and used as the Lock Screen Image or a file Url to a local image on the file system that needs to be used as the Lock Screen Image.
    personalizationLockScreenImageUrl *string
    // State Management Setting.
    privacyAdvertisingId *StateManagementSetting
    // Indicates whether or not to allow the automatic acceptance of the pairing and privacy user consent dialog when launching apps.
    privacyAutoAcceptPairingAndConsentPrompts *bool
    // Indicates whether or not to block the usage of cloud based speech services for Cortana, Dictation, or Store applications.
    privacyBlockInputPersonalization *bool
    // Indicates whether or not to Block the user from reset protection mode.
    resetProtectionModeBlocked *bool
    // Specifies what level of safe search (filtering adult content) is required
    safeSearchFilter *SafeSearchFilterType
    // Indicates whether or not to Block the user from taking Screenshots.
    screenCaptureBlocked *bool
    // Specifies if search can use diacritics.
    searchBlockDiacritics *bool
    // Specifies whether to use automatic language detection when indexing content and properties.
    searchDisableAutoLanguageDetection *bool
    // Indicates whether or not to disable the search indexer backoff feature.
    searchDisableIndexerBackoff *bool
    // Indicates whether or not to block indexing of WIP-protected items to prevent them from appearing in search results for Cortana or Explorer.
    searchDisableIndexingEncryptedItems *bool
    // Indicates whether or not to allow users to add locations on removable drives to libraries and to be indexed.
    searchDisableIndexingRemovableDrive *bool
    // Specifies minimum amount of hard drive space on the same drive as the index location before indexing stops.
    searchEnableAutomaticIndexSizeManangement *bool
    // Indicates whether or not to block remote queries of this computer’s index.
    searchEnableRemoteQueries *bool
    // Indicates whether or not to block access to Accounts in Settings app.
    settingsBlockAccountsPage *bool
    // Indicates whether or not to block the user from installing provisioning packages.
    settingsBlockAddProvisioningPackage *bool
    // Indicates whether or not to block access to Apps in Settings app.
    settingsBlockAppsPage *bool
    // Indicates whether or not to block the user from changing the language settings.
    settingsBlockChangeLanguage *bool
    // Indicates whether or not to block the user from changing power and sleep settings.
    settingsBlockChangePowerSleep *bool
    // Indicates whether or not to block the user from changing the region settings.
    settingsBlockChangeRegion *bool
    // Indicates whether or not to block the user from changing date and time settings.
    settingsBlockChangeSystemTime *bool
    // Indicates whether or not to block access to Devices in Settings app.
    settingsBlockDevicesPage *bool
    // Indicates whether or not to block access to Ease of Access in Settings app.
    settingsBlockEaseOfAccessPage *bool
    // Indicates whether or not to block the user from editing the device name.
    settingsBlockEditDeviceName *bool
    // Indicates whether or not to block access to Gaming in Settings app.
    settingsBlockGamingPage *bool
    // Indicates whether or not to block access to Network & Internet in Settings app.
    settingsBlockNetworkInternetPage *bool
    // Indicates whether or not to block access to Personalization in Settings app.
    settingsBlockPersonalizationPage *bool
    // Indicates whether or not to block access to Privacy in Settings app.
    settingsBlockPrivacyPage *bool
    // Indicates whether or not to block the runtime configuration agent from removing provisioning packages.
    settingsBlockRemoveProvisioningPackage *bool
    // Indicates whether or not to block access to Settings app.
    settingsBlockSettingsApp *bool
    // Indicates whether or not to block access to System in Settings app.
    settingsBlockSystemPage *bool
    // Indicates whether or not to block access to Time & Language in Settings app.
    settingsBlockTimeLanguagePage *bool
    // Indicates whether or not to block access to Update & Security in Settings app.
    settingsBlockUpdateSecurityPage *bool
    // Indicates whether or not to block multiple users of the same app to share data.
    sharedUserAppDataAllowed *bool
    // Indicates whether or not users can override SmartScreen Filter warnings about potentially malicious websites.
    smartScreenBlockPromptOverride *bool
    // Indicates whether or not users can override the SmartScreen Filter warnings about downloading unverified files
    smartScreenBlockPromptOverrideForFiles *bool
    // This property will be deprecated in July 2019 and will be replaced by property SmartScreenAppInstallControl. Allows IT Admins to control whether users are allowed to install apps from places other than the Store.
    smartScreenEnableAppInstallControl *bool
    // Indicates whether or not to block the user from unpinning apps from taskbar.
    startBlockUnpinningAppsFromTaskbar *bool
    // Type of start menu app list visibility.
    startMenuAppListVisibility *WindowsStartMenuAppListVisibilityType
    // Enabling this policy hides the change account setting from appearing in the user tile in the start menu.
    startMenuHideChangeAccountSettings *bool
    // Enabling this policy hides the most used apps from appearing on the start menu and disables the corresponding toggle in the Settings app.
    startMenuHideFrequentlyUsedApps *bool
    // Enabling this policy hides hibernate from appearing in the power button in the start menu.
    startMenuHideHibernate *bool
    // Enabling this policy hides lock from appearing in the user tile in the start menu.
    startMenuHideLock *bool
    // Enabling this policy hides the power button from appearing in the start menu.
    startMenuHidePowerButton *bool
    // Enabling this policy hides recent jump lists from appearing on the start menu/taskbar and disables the corresponding toggle in the Settings app.
    startMenuHideRecentJumpLists *bool
    // Enabling this policy hides recently added apps from appearing on the start menu and disables the corresponding toggle in the Settings app.
    startMenuHideRecentlyAddedApps *bool
    // Enabling this policy hides 'Restart/Update and Restart' from appearing in the power button in the start menu.
    startMenuHideRestartOptions *bool
    // Enabling this policy hides shut down/update and shut down from appearing in the power button in the start menu.
    startMenuHideShutDown *bool
    // Enabling this policy hides sign out from appearing in the user tile in the start menu.
    startMenuHideSignOut *bool
    // Enabling this policy hides sleep from appearing in the power button in the start menu.
    startMenuHideSleep *bool
    // Enabling this policy hides switch account from appearing in the user tile in the start menu.
    startMenuHideSwitchAccount *bool
    // Enabling this policy hides the user tile from appearing in the start menu.
    startMenuHideUserTile *bool
    // This policy setting allows you to import Edge assets to be used with startMenuLayoutXml policy. Start layout can contain secondary tile from Edge app which looks for Edge local asset file. Edge local asset would not exist and cause Edge secondary tile to appear empty in this case. This policy only gets applied when startMenuLayoutXml policy is modified. The value should be a UTF-8 Base64 encoded byte array.
    startMenuLayoutEdgeAssetsXml []byte
    // Allows admins to override the default Start menu layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in a UTF8 encoded byte array format.
    startMenuLayoutXml []byte
    // Type of display modes for the start menu.
    startMenuMode *WindowsStartMenuModeType
    // Generic visibility state.
    startMenuPinnedFolderDocuments *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderDownloads *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderFileExplorer *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderHomeGroup *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderMusic *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderNetwork *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderPersonalFolder *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderPictures *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderSettings *VisibilitySetting
    // Generic visibility state.
    startMenuPinnedFolderVideos *VisibilitySetting
    // Indicates whether or not to Block the user from using removable storage.
    storageBlockRemovableStorage *bool
    // Indicating whether or not to require encryption on a mobile device.
    storageRequireMobileDeviceEncryption *bool
    // Indicates whether application data is restricted to the system drive.
    storageRestrictAppDataToSystemVolume *bool
    // Indicates whether the installation of applications is restricted to the system drive.
    storageRestrictAppInstallToSystemVolume *bool
    // Whether the device is required to connect to the network.
    tenantLockdownRequireNetworkDuringOutOfBoxExperience *bool
    // Indicates whether or not to Block the user from USB connection.
    usbBlocked *bool
    // Indicates whether or not to Block the user from voice recording.
    voiceRecordingBlocked *bool
    // Indicates whether or not user's localhost IP address is displayed while making phone calls using the WebRTC
    webRtcBlockLocalhostIpAddress *bool
    // Indicating whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
    wiFiBlockAutomaticConnectHotspots *bool
    // Indicates whether or not to Block the user from using Wi-Fi.
    wiFiBlocked *bool
    // Indicates whether or not to Block the user from using Wi-Fi manual configuration.
    wiFiBlockManualConfiguration *bool
    // Specify how often devices scan for Wi-Fi networks. Supported values are 1-500, where 100 = default, and 500 = low frequency. Valid values 1 to 500
    wiFiScanInterval *int32
    // Allows IT admins to block experiences that are typically for consumers only, such as Start suggestions, Membership notifications, Post-OOBE app install and redirect tiles.
    windowsSpotlightBlockConsumerSpecificFeatures *bool
    // Allows IT admins to turn off all Windows Spotlight features
    windowsSpotlightBlocked *bool
    // Block suggestions from Microsoft that show after each OS clean install, upgrade or in an on-going basis to introduce users to what is new or changed
    windowsSpotlightBlockOnActionCenter *bool
    // Block personalized content in Windows spotlight based on user’s device usage.
    windowsSpotlightBlockTailoredExperiences *bool
    // Block third party content delivered via Windows Spotlight
    windowsSpotlightBlockThirdPartyNotifications *bool
    // Block Windows Spotlight Windows welcome experience
    windowsSpotlightBlockWelcomeExperience *bool
    // Allows IT admins to turn off the popup of Windows Tips.
    windowsSpotlightBlockWindowsTips *bool
    // Allows IT admind to set a predefined default search engine for MDM-Controlled devices
    windowsSpotlightConfigureOnLockScreen *WindowsSpotlightEnablementSettings
    // Indicates whether or not to block automatic update of apps from Windows Store.
    windowsStoreBlockAutoUpdate *bool
    // Indicates whether or not to Block the user from using the Windows store.
    windowsStoreBlocked *bool
    // Indicates whether or not to enable Private Store Only.
    windowsStoreEnablePrivateStoreOnly *bool
    // Indicates whether or not to allow other devices from discovering this PC for projection.
    wirelessDisplayBlockProjectionToThisDevice *bool
    // Indicates whether or not to allow user input from wireless display receiver.
    wirelessDisplayBlockUserInputFromReceiver *bool
    // Indicates whether or not to require a PIN for new devices to initiate pairing.
    wirelessDisplayRequirePinForPairing *bool
}
// NewWindows10GeneralConfiguration instantiates a new Windows10GeneralConfiguration and sets the default values.
func NewWindows10GeneralConfiguration()(*Windows10GeneralConfiguration) {
    m := &Windows10GeneralConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10GeneralConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10GeneralConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10GeneralConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10GeneralConfiguration(), nil
}
// GetAccountsBlockAddingNonMicrosoftAccountEmail gets the accountsBlockAddingNonMicrosoftAccountEmail property value. Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account.
func (m *Windows10GeneralConfiguration) GetAccountsBlockAddingNonMicrosoftAccountEmail()(*bool) {
    return m.accountsBlockAddingNonMicrosoftAccountEmail
}
// GetAntiTheftModeBlocked gets the antiTheftModeBlocked property value. Indicates whether or not to block the user from selecting an AntiTheft mode preference (Windows 10 Mobile only).
func (m *Windows10GeneralConfiguration) GetAntiTheftModeBlocked()(*bool) {
    return m.antiTheftModeBlocked
}
// GetAppsAllowTrustedAppsSideloading gets the appsAllowTrustedAppsSideloading property value. State Management Setting.
func (m *Windows10GeneralConfiguration) GetAppsAllowTrustedAppsSideloading()(*StateManagementSetting) {
    return m.appsAllowTrustedAppsSideloading
}
// GetAppsBlockWindowsStoreOriginatedApps gets the appsBlockWindowsStoreOriginatedApps property value. Indicates whether or not to disable the launch of all apps from Windows Store that came pre-installed or were downloaded.
func (m *Windows10GeneralConfiguration) GetAppsBlockWindowsStoreOriginatedApps()(*bool) {
    return m.appsBlockWindowsStoreOriginatedApps
}
// GetBluetoothAllowedServices gets the bluetoothAllowedServices property value. Specify a list of allowed Bluetooth services and profiles in hex formatted strings.
func (m *Windows10GeneralConfiguration) GetBluetoothAllowedServices()([]string) {
    return m.bluetoothAllowedServices
}
// GetBluetoothBlockAdvertising gets the bluetoothBlockAdvertising property value. Whether or not to Block the user from using bluetooth advertising.
func (m *Windows10GeneralConfiguration) GetBluetoothBlockAdvertising()(*bool) {
    return m.bluetoothBlockAdvertising
}
// GetBluetoothBlockDiscoverableMode gets the bluetoothBlockDiscoverableMode property value. Whether or not to Block the user from using bluetooth discoverable mode.
func (m *Windows10GeneralConfiguration) GetBluetoothBlockDiscoverableMode()(*bool) {
    return m.bluetoothBlockDiscoverableMode
}
// GetBluetoothBlocked gets the bluetoothBlocked property value. Whether or not to Block the user from using bluetooth.
func (m *Windows10GeneralConfiguration) GetBluetoothBlocked()(*bool) {
    return m.bluetoothBlocked
}
// GetBluetoothBlockPrePairing gets the bluetoothBlockPrePairing property value. Whether or not to block specific bundled Bluetooth peripherals to automatically pair with the host device.
func (m *Windows10GeneralConfiguration) GetBluetoothBlockPrePairing()(*bool) {
    return m.bluetoothBlockPrePairing
}
// GetCameraBlocked gets the cameraBlocked property value. Whether or not to Block the user from accessing the camera of the device.
func (m *Windows10GeneralConfiguration) GetCameraBlocked()(*bool) {
    return m.cameraBlocked
}
// GetCellularBlockDataWhenRoaming gets the cellularBlockDataWhenRoaming property value. Whether or not to Block the user from using data over cellular while roaming.
func (m *Windows10GeneralConfiguration) GetCellularBlockDataWhenRoaming()(*bool) {
    return m.cellularBlockDataWhenRoaming
}
// GetCellularBlockVpn gets the cellularBlockVpn property value. Whether or not to Block the user from using VPN over cellular.
func (m *Windows10GeneralConfiguration) GetCellularBlockVpn()(*bool) {
    return m.cellularBlockVpn
}
// GetCellularBlockVpnWhenRoaming gets the cellularBlockVpnWhenRoaming property value. Whether or not to Block the user from using VPN when roaming over cellular.
func (m *Windows10GeneralConfiguration) GetCellularBlockVpnWhenRoaming()(*bool) {
    return m.cellularBlockVpnWhenRoaming
}
// GetCertificatesBlockManualRootCertificateInstallation gets the certificatesBlockManualRootCertificateInstallation property value. Whether or not to Block the user from doing manual root certificate installation.
func (m *Windows10GeneralConfiguration) GetCertificatesBlockManualRootCertificateInstallation()(*bool) {
    return m.certificatesBlockManualRootCertificateInstallation
}
// GetConnectedDevicesServiceBlocked gets the connectedDevicesServiceBlocked property value. Whether or not to block Connected Devices Service which enables discovery and connection to other devices, remote messaging, remote app sessions and other cross-device experiences.
func (m *Windows10GeneralConfiguration) GetConnectedDevicesServiceBlocked()(*bool) {
    return m.connectedDevicesServiceBlocked
}
// GetCopyPasteBlocked gets the copyPasteBlocked property value. Whether or not to Block the user from using copy paste.
func (m *Windows10GeneralConfiguration) GetCopyPasteBlocked()(*bool) {
    return m.copyPasteBlocked
}
// GetCortanaBlocked gets the cortanaBlocked property value. Whether or not to Block the user from using Cortana.
func (m *Windows10GeneralConfiguration) GetCortanaBlocked()(*bool) {
    return m.cortanaBlocked
}
// GetDefenderBlockEndUserAccess gets the defenderBlockEndUserAccess property value. Whether or not to block end user access to Defender.
func (m *Windows10GeneralConfiguration) GetDefenderBlockEndUserAccess()(*bool) {
    return m.defenderBlockEndUserAccess
}
// GetDefenderCloudBlockLevel gets the defenderCloudBlockLevel property value. Possible values of Cloud Block Level
func (m *Windows10GeneralConfiguration) GetDefenderCloudBlockLevel()(*DefenderCloudBlockLevelType) {
    return m.defenderCloudBlockLevel
}
// GetDefenderDaysBeforeDeletingQuarantinedMalware gets the defenderDaysBeforeDeletingQuarantinedMalware property value. Number of days before deleting quarantined malware. Valid values 0 to 90
func (m *Windows10GeneralConfiguration) GetDefenderDaysBeforeDeletingQuarantinedMalware()(*int32) {
    return m.defenderDaysBeforeDeletingQuarantinedMalware
}
// GetDefenderDetectedMalwareActions gets the defenderDetectedMalwareActions property value. Gets or sets Defender’s actions to take on detected Malware per threat level.
func (m *Windows10GeneralConfiguration) GetDefenderDetectedMalwareActions()(DefenderDetectedMalwareActionsable) {
    return m.defenderDetectedMalwareActions
}
// GetDefenderFileExtensionsToExclude gets the defenderFileExtensionsToExclude property value. File extensions to exclude from scans and real time protection.
func (m *Windows10GeneralConfiguration) GetDefenderFileExtensionsToExclude()([]string) {
    return m.defenderFileExtensionsToExclude
}
// GetDefenderFilesAndFoldersToExclude gets the defenderFilesAndFoldersToExclude property value. Files and folder to exclude from scans and real time protection.
func (m *Windows10GeneralConfiguration) GetDefenderFilesAndFoldersToExclude()([]string) {
    return m.defenderFilesAndFoldersToExclude
}
// GetDefenderMonitorFileActivity gets the defenderMonitorFileActivity property value. Possible values for monitoring file activity.
func (m *Windows10GeneralConfiguration) GetDefenderMonitorFileActivity()(*DefenderMonitorFileActivity) {
    return m.defenderMonitorFileActivity
}
// GetDefenderProcessesToExclude gets the defenderProcessesToExclude property value. Processes to exclude from scans and real time protection.
func (m *Windows10GeneralConfiguration) GetDefenderProcessesToExclude()([]string) {
    return m.defenderProcessesToExclude
}
// GetDefenderPromptForSampleSubmission gets the defenderPromptForSampleSubmission property value. Possible values for prompting user for samples submission.
func (m *Windows10GeneralConfiguration) GetDefenderPromptForSampleSubmission()(*DefenderPromptForSampleSubmission) {
    return m.defenderPromptForSampleSubmission
}
// GetDefenderRequireBehaviorMonitoring gets the defenderRequireBehaviorMonitoring property value. Indicates whether or not to require behavior monitoring.
func (m *Windows10GeneralConfiguration) GetDefenderRequireBehaviorMonitoring()(*bool) {
    return m.defenderRequireBehaviorMonitoring
}
// GetDefenderRequireCloudProtection gets the defenderRequireCloudProtection property value. Indicates whether or not to require cloud protection.
func (m *Windows10GeneralConfiguration) GetDefenderRequireCloudProtection()(*bool) {
    return m.defenderRequireCloudProtection
}
// GetDefenderRequireNetworkInspectionSystem gets the defenderRequireNetworkInspectionSystem property value. Indicates whether or not to require network inspection system.
func (m *Windows10GeneralConfiguration) GetDefenderRequireNetworkInspectionSystem()(*bool) {
    return m.defenderRequireNetworkInspectionSystem
}
// GetDefenderRequireRealTimeMonitoring gets the defenderRequireRealTimeMonitoring property value. Indicates whether or not to require real time monitoring.
func (m *Windows10GeneralConfiguration) GetDefenderRequireRealTimeMonitoring()(*bool) {
    return m.defenderRequireRealTimeMonitoring
}
// GetDefenderScanArchiveFiles gets the defenderScanArchiveFiles property value. Indicates whether or not to scan archive files.
func (m *Windows10GeneralConfiguration) GetDefenderScanArchiveFiles()(*bool) {
    return m.defenderScanArchiveFiles
}
// GetDefenderScanDownloads gets the defenderScanDownloads property value. Indicates whether or not to scan downloads.
func (m *Windows10GeneralConfiguration) GetDefenderScanDownloads()(*bool) {
    return m.defenderScanDownloads
}
// GetDefenderScanIncomingMail gets the defenderScanIncomingMail property value. Indicates whether or not to scan incoming mail messages.
func (m *Windows10GeneralConfiguration) GetDefenderScanIncomingMail()(*bool) {
    return m.defenderScanIncomingMail
}
// GetDefenderScanMappedNetworkDrivesDuringFullScan gets the defenderScanMappedNetworkDrivesDuringFullScan property value. Indicates whether or not to scan mapped network drives during full scan.
func (m *Windows10GeneralConfiguration) GetDefenderScanMappedNetworkDrivesDuringFullScan()(*bool) {
    return m.defenderScanMappedNetworkDrivesDuringFullScan
}
// GetDefenderScanMaxCpu gets the defenderScanMaxCpu property value. Max CPU usage percentage during scan. Valid values 0 to 100
func (m *Windows10GeneralConfiguration) GetDefenderScanMaxCpu()(*int32) {
    return m.defenderScanMaxCpu
}
// GetDefenderScanNetworkFiles gets the defenderScanNetworkFiles property value. Indicates whether or not to scan files opened from a network folder.
func (m *Windows10GeneralConfiguration) GetDefenderScanNetworkFiles()(*bool) {
    return m.defenderScanNetworkFiles
}
// GetDefenderScanRemovableDrivesDuringFullScan gets the defenderScanRemovableDrivesDuringFullScan property value. Indicates whether or not to scan removable drives during full scan.
func (m *Windows10GeneralConfiguration) GetDefenderScanRemovableDrivesDuringFullScan()(*bool) {
    return m.defenderScanRemovableDrivesDuringFullScan
}
// GetDefenderScanScriptsLoadedInInternetExplorer gets the defenderScanScriptsLoadedInInternetExplorer property value. Indicates whether or not to scan scripts loaded in Internet Explorer browser.
func (m *Windows10GeneralConfiguration) GetDefenderScanScriptsLoadedInInternetExplorer()(*bool) {
    return m.defenderScanScriptsLoadedInInternetExplorer
}
// GetDefenderScanType gets the defenderScanType property value. Possible values for system scan type.
func (m *Windows10GeneralConfiguration) GetDefenderScanType()(*DefenderScanType) {
    return m.defenderScanType
}
// GetDefenderScheduledQuickScanTime gets the defenderScheduledQuickScanTime property value. The time to perform a daily quick scan.
func (m *Windows10GeneralConfiguration) GetDefenderScheduledQuickScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.defenderScheduledQuickScanTime
}
// GetDefenderScheduledScanTime gets the defenderScheduledScanTime property value. The defender time for the system scan.
func (m *Windows10GeneralConfiguration) GetDefenderScheduledScanTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.defenderScheduledScanTime
}
// GetDefenderSignatureUpdateIntervalInHours gets the defenderSignatureUpdateIntervalInHours property value. The signature update interval in hours. Specify 0 not to check. Valid values 0 to 24
func (m *Windows10GeneralConfiguration) GetDefenderSignatureUpdateIntervalInHours()(*int32) {
    return m.defenderSignatureUpdateIntervalInHours
}
// GetDefenderSystemScanSchedule gets the defenderSystemScanSchedule property value. Possible values for a weekly schedule.
func (m *Windows10GeneralConfiguration) GetDefenderSystemScanSchedule()(*WeeklySchedule) {
    return m.defenderSystemScanSchedule
}
// GetDeveloperUnlockSetting gets the developerUnlockSetting property value. State Management Setting.
func (m *Windows10GeneralConfiguration) GetDeveloperUnlockSetting()(*StateManagementSetting) {
    return m.developerUnlockSetting
}
// GetDeviceManagementBlockFactoryResetOnMobile gets the deviceManagementBlockFactoryResetOnMobile property value. Indicates whether or not to Block the user from resetting their phone.
func (m *Windows10GeneralConfiguration) GetDeviceManagementBlockFactoryResetOnMobile()(*bool) {
    return m.deviceManagementBlockFactoryResetOnMobile
}
// GetDeviceManagementBlockManualUnenroll gets the deviceManagementBlockManualUnenroll property value. Indicates whether or not to Block the user from doing manual un-enrollment from device management.
func (m *Windows10GeneralConfiguration) GetDeviceManagementBlockManualUnenroll()(*bool) {
    return m.deviceManagementBlockManualUnenroll
}
// GetDiagnosticsDataSubmissionMode gets the diagnosticsDataSubmissionMode property value. Allow the device to send diagnostic and usage telemetry data, such as Watson.
func (m *Windows10GeneralConfiguration) GetDiagnosticsDataSubmissionMode()(*DiagnosticDataSubmissionMode) {
    return m.diagnosticsDataSubmissionMode
}
// GetEdgeAllowStartPagesModification gets the edgeAllowStartPagesModification property value. Allow users to change Start pages on Edge. Use the EdgeHomepageUrls to specify the Start pages that the user would see by default when they open Edge.
func (m *Windows10GeneralConfiguration) GetEdgeAllowStartPagesModification()(*bool) {
    return m.edgeAllowStartPagesModification
}
// GetEdgeBlockAccessToAboutFlags gets the edgeBlockAccessToAboutFlags property value. Indicates whether or not to prevent access to about flags on Edge browser.
func (m *Windows10GeneralConfiguration) GetEdgeBlockAccessToAboutFlags()(*bool) {
    return m.edgeBlockAccessToAboutFlags
}
// GetEdgeBlockAddressBarDropdown gets the edgeBlockAddressBarDropdown property value. Block the address bar dropdown functionality in Microsoft Edge. Disable this settings to minimize network connections from Microsoft Edge to Microsoft services.
func (m *Windows10GeneralConfiguration) GetEdgeBlockAddressBarDropdown()(*bool) {
    return m.edgeBlockAddressBarDropdown
}
// GetEdgeBlockAutofill gets the edgeBlockAutofill property value. Indicates whether or not to block auto fill.
func (m *Windows10GeneralConfiguration) GetEdgeBlockAutofill()(*bool) {
    return m.edgeBlockAutofill
}
// GetEdgeBlockCompatibilityList gets the edgeBlockCompatibilityList property value. Block Microsoft compatibility list in Microsoft Edge. This list from Microsoft helps Edge properly display sites with known compatibility issues.
func (m *Windows10GeneralConfiguration) GetEdgeBlockCompatibilityList()(*bool) {
    return m.edgeBlockCompatibilityList
}
// GetEdgeBlockDeveloperTools gets the edgeBlockDeveloperTools property value. Indicates whether or not to block developer tools in the Edge browser.
func (m *Windows10GeneralConfiguration) GetEdgeBlockDeveloperTools()(*bool) {
    return m.edgeBlockDeveloperTools
}
// GetEdgeBlocked gets the edgeBlocked property value. Indicates whether or not to Block the user from using the Edge browser.
func (m *Windows10GeneralConfiguration) GetEdgeBlocked()(*bool) {
    return m.edgeBlocked
}
// GetEdgeBlockExtensions gets the edgeBlockExtensions property value. Indicates whether or not to block extensions in the Edge browser.
func (m *Windows10GeneralConfiguration) GetEdgeBlockExtensions()(*bool) {
    return m.edgeBlockExtensions
}
// GetEdgeBlockInPrivateBrowsing gets the edgeBlockInPrivateBrowsing property value. Indicates whether or not to block InPrivate browsing on corporate networks, in the Edge browser.
func (m *Windows10GeneralConfiguration) GetEdgeBlockInPrivateBrowsing()(*bool) {
    return m.edgeBlockInPrivateBrowsing
}
// GetEdgeBlockJavaScript gets the edgeBlockJavaScript property value. Indicates whether or not to Block the user from using JavaScript.
func (m *Windows10GeneralConfiguration) GetEdgeBlockJavaScript()(*bool) {
    return m.edgeBlockJavaScript
}
// GetEdgeBlockLiveTileDataCollection gets the edgeBlockLiveTileDataCollection property value. Block the collection of information by Microsoft for live tile creation when users pin a site to Start from Microsoft Edge.
func (m *Windows10GeneralConfiguration) GetEdgeBlockLiveTileDataCollection()(*bool) {
    return m.edgeBlockLiveTileDataCollection
}
// GetEdgeBlockPasswordManager gets the edgeBlockPasswordManager property value. Indicates whether or not to Block password manager.
func (m *Windows10GeneralConfiguration) GetEdgeBlockPasswordManager()(*bool) {
    return m.edgeBlockPasswordManager
}
// GetEdgeBlockPopups gets the edgeBlockPopups property value. Indicates whether or not to block popups.
func (m *Windows10GeneralConfiguration) GetEdgeBlockPopups()(*bool) {
    return m.edgeBlockPopups
}
// GetEdgeBlockSearchSuggestions gets the edgeBlockSearchSuggestions property value. Indicates whether or not to block the user from using the search suggestions in the address bar.
func (m *Windows10GeneralConfiguration) GetEdgeBlockSearchSuggestions()(*bool) {
    return m.edgeBlockSearchSuggestions
}
// GetEdgeBlockSendingDoNotTrackHeader gets the edgeBlockSendingDoNotTrackHeader property value. Indicates whether or not to Block the user from sending the do not track header.
func (m *Windows10GeneralConfiguration) GetEdgeBlockSendingDoNotTrackHeader()(*bool) {
    return m.edgeBlockSendingDoNotTrackHeader
}
// GetEdgeBlockSendingIntranetTrafficToInternetExplorer gets the edgeBlockSendingIntranetTrafficToInternetExplorer property value. Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer. Note: the name of this property is misleading; the property is obsolete, use EdgeSendIntranetTrafficToInternetExplorer instead.
func (m *Windows10GeneralConfiguration) GetEdgeBlockSendingIntranetTrafficToInternetExplorer()(*bool) {
    return m.edgeBlockSendingIntranetTrafficToInternetExplorer
}
// GetEdgeClearBrowsingDataOnExit gets the edgeClearBrowsingDataOnExit property value. Clear browsing data on exiting Microsoft Edge.
func (m *Windows10GeneralConfiguration) GetEdgeClearBrowsingDataOnExit()(*bool) {
    return m.edgeClearBrowsingDataOnExit
}
// GetEdgeCookiePolicy gets the edgeCookiePolicy property value. Possible values to specify which cookies are allowed in Microsoft Edge.
func (m *Windows10GeneralConfiguration) GetEdgeCookiePolicy()(*EdgeCookiePolicy) {
    return m.edgeCookiePolicy
}
// GetEdgeDisableFirstRunPage gets the edgeDisableFirstRunPage property value. Block the Microsoft web page that opens on the first use of Microsoft Edge. This policy allows enterprises, like those enrolled in zero emissions configurations, to block this page.
func (m *Windows10GeneralConfiguration) GetEdgeDisableFirstRunPage()(*bool) {
    return m.edgeDisableFirstRunPage
}
// GetEdgeEnterpriseModeSiteListLocation gets the edgeEnterpriseModeSiteListLocation property value. Indicates the enterprise mode site list location. Could be a local file, local network or http location.
func (m *Windows10GeneralConfiguration) GetEdgeEnterpriseModeSiteListLocation()(*string) {
    return m.edgeEnterpriseModeSiteListLocation
}
// GetEdgeFirstRunUrl gets the edgeFirstRunUrl property value. The first run URL for when Edge browser is opened for the first time.
func (m *Windows10GeneralConfiguration) GetEdgeFirstRunUrl()(*string) {
    return m.edgeFirstRunUrl
}
// GetEdgeHomepageUrls gets the edgeHomepageUrls property value. The list of URLs for homepages shodwn on MDM-enrolled devices on Edge browser.
func (m *Windows10GeneralConfiguration) GetEdgeHomepageUrls()([]string) {
    return m.edgeHomepageUrls
}
// GetEdgeRequireSmartScreen gets the edgeRequireSmartScreen property value. Indicates whether or not to Require the user to use the smart screen filter.
func (m *Windows10GeneralConfiguration) GetEdgeRequireSmartScreen()(*bool) {
    return m.edgeRequireSmartScreen
}
// GetEdgeSearchEngine gets the edgeSearchEngine property value. Allows IT admins to set a default search engine for MDM-Controlled devices. Users can override this and change their default search engine provided the AllowSearchEngineCustomization policy is not set.
func (m *Windows10GeneralConfiguration) GetEdgeSearchEngine()(EdgeSearchEngineBaseable) {
    return m.edgeSearchEngine
}
// GetEdgeSendIntranetTrafficToInternetExplorer gets the edgeSendIntranetTrafficToInternetExplorer property value. Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer.
func (m *Windows10GeneralConfiguration) GetEdgeSendIntranetTrafficToInternetExplorer()(*bool) {
    return m.edgeSendIntranetTrafficToInternetExplorer
}
// GetEdgeSyncFavoritesWithInternetExplorer gets the edgeSyncFavoritesWithInternetExplorer property value. Enable favorites sync between Internet Explorer and Microsoft Edge. Additions, deletions, modifications and order changes to favorites are shared between browsers.
func (m *Windows10GeneralConfiguration) GetEdgeSyncFavoritesWithInternetExplorer()(*bool) {
    return m.edgeSyncFavoritesWithInternetExplorer
}
// GetEnterpriseCloudPrintDiscoveryEndPoint gets the enterpriseCloudPrintDiscoveryEndPoint property value. Endpoint for discovering cloud printers.
func (m *Windows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryEndPoint()(*string) {
    return m.enterpriseCloudPrintDiscoveryEndPoint
}
// GetEnterpriseCloudPrintDiscoveryMaxLimit gets the enterpriseCloudPrintDiscoveryMaxLimit property value. Maximum number of printers that should be queried from a discovery endpoint. This is a mobile only setting. Valid values 1 to 65535
func (m *Windows10GeneralConfiguration) GetEnterpriseCloudPrintDiscoveryMaxLimit()(*int32) {
    return m.enterpriseCloudPrintDiscoveryMaxLimit
}
// GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier gets the enterpriseCloudPrintMopriaDiscoveryResourceIdentifier property value. OAuth resource URI for printer discovery service as configured in Azure portal.
func (m *Windows10GeneralConfiguration) GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier()(*string) {
    return m.enterpriseCloudPrintMopriaDiscoveryResourceIdentifier
}
// GetEnterpriseCloudPrintOAuthAuthority gets the enterpriseCloudPrintOAuthAuthority property value. Authentication endpoint for acquiring OAuth tokens.
func (m *Windows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthAuthority()(*string) {
    return m.enterpriseCloudPrintOAuthAuthority
}
// GetEnterpriseCloudPrintOAuthClientIdentifier gets the enterpriseCloudPrintOAuthClientIdentifier property value. GUID of a client application authorized to retrieve OAuth tokens from the OAuth Authority.
func (m *Windows10GeneralConfiguration) GetEnterpriseCloudPrintOAuthClientIdentifier()(*string) {
    return m.enterpriseCloudPrintOAuthClientIdentifier
}
// GetEnterpriseCloudPrintResourceIdentifier gets the enterpriseCloudPrintResourceIdentifier property value. OAuth resource URI for print service as configured in the Azure portal.
func (m *Windows10GeneralConfiguration) GetEnterpriseCloudPrintResourceIdentifier()(*string) {
    return m.enterpriseCloudPrintResourceIdentifier
}
// GetExperienceBlockDeviceDiscovery gets the experienceBlockDeviceDiscovery property value. Indicates whether or not to enable device discovery UX.
func (m *Windows10GeneralConfiguration) GetExperienceBlockDeviceDiscovery()(*bool) {
    return m.experienceBlockDeviceDiscovery
}
// GetExperienceBlockErrorDialogWhenNoSIM gets the experienceBlockErrorDialogWhenNoSIM property value. Indicates whether or not to allow the error dialog from displaying if no SIM card is detected.
func (m *Windows10GeneralConfiguration) GetExperienceBlockErrorDialogWhenNoSIM()(*bool) {
    return m.experienceBlockErrorDialogWhenNoSIM
}
// GetExperienceBlockTaskSwitcher gets the experienceBlockTaskSwitcher property value. Indicates whether or not to enable task switching on the device.
func (m *Windows10GeneralConfiguration) GetExperienceBlockTaskSwitcher()(*bool) {
    return m.experienceBlockTaskSwitcher
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10GeneralConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["antiTheftModeBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAntiTheftModeBlocked(val)
        }
        return nil
    }
    res["appsAllowTrustedAppsSideloading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStateManagementSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsAllowTrustedAppsSideloading(val.(*StateManagementSetting))
        }
        return nil
    }
    res["appsBlockWindowsStoreOriginatedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppsBlockWindowsStoreOriginatedApps(val)
        }
        return nil
    }
    res["bluetoothAllowedServices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBluetoothAllowedServices(res)
        }
        return nil
    }
    res["bluetoothBlockAdvertising"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockAdvertising(val)
        }
        return nil
    }
    res["bluetoothBlockDiscoverableMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockDiscoverableMode(val)
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
    res["bluetoothBlockPrePairing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBluetoothBlockPrePairing(val)
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
    res["cellularBlockDataWhenRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockDataWhenRoaming(val)
        }
        return nil
    }
    res["cellularBlockVpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockVpn(val)
        }
        return nil
    }
    res["cellularBlockVpnWhenRoaming"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCellularBlockVpnWhenRoaming(val)
        }
        return nil
    }
    res["certificatesBlockManualRootCertificateInstallation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificatesBlockManualRootCertificateInstallation(val)
        }
        return nil
    }
    res["connectedDevicesServiceBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedDevicesServiceBlocked(val)
        }
        return nil
    }
    res["copyPasteBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopyPasteBlocked(val)
        }
        return nil
    }
    res["cortanaBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCortanaBlocked(val)
        }
        return nil
    }
    res["defenderBlockEndUserAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderBlockEndUserAccess(val)
        }
        return nil
    }
    res["defenderCloudBlockLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderCloudBlockLevelType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderCloudBlockLevel(val.(*DefenderCloudBlockLevelType))
        }
        return nil
    }
    res["defenderDaysBeforeDeletingQuarantinedMalware"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDaysBeforeDeletingQuarantinedMalware(val)
        }
        return nil
    }
    res["defenderDetectedMalwareActions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDefenderDetectedMalwareActionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderDetectedMalwareActions(val.(DefenderDetectedMalwareActionsable))
        }
        return nil
    }
    res["defenderFileExtensionsToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderFileExtensionsToExclude(res)
        }
        return nil
    }
    res["defenderFilesAndFoldersToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderFilesAndFoldersToExclude(res)
        }
        return nil
    }
    res["defenderMonitorFileActivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderMonitorFileActivity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderMonitorFileActivity(val.(*DefenderMonitorFileActivity))
        }
        return nil
    }
    res["defenderProcessesToExclude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderProcessesToExclude(res)
        }
        return nil
    }
    res["defenderPromptForSampleSubmission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderPromptForSampleSubmission)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderPromptForSampleSubmission(val.(*DefenderPromptForSampleSubmission))
        }
        return nil
    }
    res["defenderRequireBehaviorMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderRequireBehaviorMonitoring(val)
        }
        return nil
    }
    res["defenderRequireCloudProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderRequireCloudProtection(val)
        }
        return nil
    }
    res["defenderRequireNetworkInspectionSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderRequireNetworkInspectionSystem(val)
        }
        return nil
    }
    res["defenderRequireRealTimeMonitoring"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderRequireRealTimeMonitoring(val)
        }
        return nil
    }
    res["defenderScanArchiveFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanArchiveFiles(val)
        }
        return nil
    }
    res["defenderScanDownloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanDownloads(val)
        }
        return nil
    }
    res["defenderScanIncomingMail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanIncomingMail(val)
        }
        return nil
    }
    res["defenderScanMappedNetworkDrivesDuringFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanMappedNetworkDrivesDuringFullScan(val)
        }
        return nil
    }
    res["defenderScanMaxCpu"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanMaxCpu(val)
        }
        return nil
    }
    res["defenderScanNetworkFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanNetworkFiles(val)
        }
        return nil
    }
    res["defenderScanRemovableDrivesDuringFullScan"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanRemovableDrivesDuringFullScan(val)
        }
        return nil
    }
    res["defenderScanScriptsLoadedInInternetExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanScriptsLoadedInInternetExplorer(val)
        }
        return nil
    }
    res["defenderScanType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderScanType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScanType(val.(*DefenderScanType))
        }
        return nil
    }
    res["defenderScheduledQuickScanTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScheduledQuickScanTime(val)
        }
        return nil
    }
    res["defenderScheduledScanTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderScheduledScanTime(val)
        }
        return nil
    }
    res["defenderSignatureUpdateIntervalInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSignatureUpdateIntervalInHours(val)
        }
        return nil
    }
    res["defenderSystemScanSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeeklySchedule)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSystemScanSchedule(val.(*WeeklySchedule))
        }
        return nil
    }
    res["developerUnlockSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStateManagementSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeveloperUnlockSetting(val.(*StateManagementSetting))
        }
        return nil
    }
    res["deviceManagementBlockFactoryResetOnMobile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementBlockFactoryResetOnMobile(val)
        }
        return nil
    }
    res["deviceManagementBlockManualUnenroll"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceManagementBlockManualUnenroll(val)
        }
        return nil
    }
    res["diagnosticsDataSubmissionMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDiagnosticDataSubmissionMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiagnosticsDataSubmissionMode(val.(*DiagnosticDataSubmissionMode))
        }
        return nil
    }
    res["edgeAllowStartPagesModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeAllowStartPagesModification(val)
        }
        return nil
    }
    res["edgeBlockAccessToAboutFlags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockAccessToAboutFlags(val)
        }
        return nil
    }
    res["edgeBlockAddressBarDropdown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockAddressBarDropdown(val)
        }
        return nil
    }
    res["edgeBlockAutofill"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockAutofill(val)
        }
        return nil
    }
    res["edgeBlockCompatibilityList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockCompatibilityList(val)
        }
        return nil
    }
    res["edgeBlockDeveloperTools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockDeveloperTools(val)
        }
        return nil
    }
    res["edgeBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlocked(val)
        }
        return nil
    }
    res["edgeBlockExtensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockExtensions(val)
        }
        return nil
    }
    res["edgeBlockInPrivateBrowsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockInPrivateBrowsing(val)
        }
        return nil
    }
    res["edgeBlockJavaScript"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockJavaScript(val)
        }
        return nil
    }
    res["edgeBlockLiveTileDataCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockLiveTileDataCollection(val)
        }
        return nil
    }
    res["edgeBlockPasswordManager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockPasswordManager(val)
        }
        return nil
    }
    res["edgeBlockPopups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockPopups(val)
        }
        return nil
    }
    res["edgeBlockSearchSuggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockSearchSuggestions(val)
        }
        return nil
    }
    res["edgeBlockSendingDoNotTrackHeader"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockSendingDoNotTrackHeader(val)
        }
        return nil
    }
    res["edgeBlockSendingIntranetTrafficToInternetExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeBlockSendingIntranetTrafficToInternetExplorer(val)
        }
        return nil
    }
    res["edgeClearBrowsingDataOnExit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeClearBrowsingDataOnExit(val)
        }
        return nil
    }
    res["edgeCookiePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEdgeCookiePolicy)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeCookiePolicy(val.(*EdgeCookiePolicy))
        }
        return nil
    }
    res["edgeDisableFirstRunPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeDisableFirstRunPage(val)
        }
        return nil
    }
    res["edgeEnterpriseModeSiteListLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeEnterpriseModeSiteListLocation(val)
        }
        return nil
    }
    res["edgeFirstRunUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeFirstRunUrl(val)
        }
        return nil
    }
    res["edgeHomepageUrls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEdgeHomepageUrls(res)
        }
        return nil
    }
    res["edgeRequireSmartScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeRequireSmartScreen(val)
        }
        return nil
    }
    res["edgeSearchEngine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEdgeSearchEngineBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeSearchEngine(val.(EdgeSearchEngineBaseable))
        }
        return nil
    }
    res["edgeSendIntranetTrafficToInternetExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeSendIntranetTrafficToInternetExplorer(val)
        }
        return nil
    }
    res["edgeSyncFavoritesWithInternetExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeSyncFavoritesWithInternetExplorer(val)
        }
        return nil
    }
    res["enterpriseCloudPrintDiscoveryEndPoint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseCloudPrintDiscoveryEndPoint(val)
        }
        return nil
    }
    res["enterpriseCloudPrintDiscoveryMaxLimit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseCloudPrintDiscoveryMaxLimit(val)
        }
        return nil
    }
    res["enterpriseCloudPrintMopriaDiscoveryResourceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier(val)
        }
        return nil
    }
    res["enterpriseCloudPrintOAuthAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseCloudPrintOAuthAuthority(val)
        }
        return nil
    }
    res["enterpriseCloudPrintOAuthClientIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseCloudPrintOAuthClientIdentifier(val)
        }
        return nil
    }
    res["enterpriseCloudPrintResourceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnterpriseCloudPrintResourceIdentifier(val)
        }
        return nil
    }
    res["experienceBlockDeviceDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExperienceBlockDeviceDiscovery(val)
        }
        return nil
    }
    res["experienceBlockErrorDialogWhenNoSIM"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExperienceBlockErrorDialogWhenNoSIM(val)
        }
        return nil
    }
    res["experienceBlockTaskSwitcher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExperienceBlockTaskSwitcher(val)
        }
        return nil
    }
    res["gameDvrBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGameDvrBlocked(val)
        }
        return nil
    }
    res["internetSharingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetSharingBlocked(val)
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
    res["lockScreenAllowTimeoutConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenAllowTimeoutConfiguration(val)
        }
        return nil
    }
    res["lockScreenBlockActionCenterNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockActionCenterNotifications(val)
        }
        return nil
    }
    res["lockScreenBlockCortana"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockCortana(val)
        }
        return nil
    }
    res["lockScreenBlockToastNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenBlockToastNotifications(val)
        }
        return nil
    }
    res["lockScreenTimeoutInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLockScreenTimeoutInSeconds(val)
        }
        return nil
    }
    res["logonBlockFastUserSwitching"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonBlockFastUserSwitching(val)
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
    res["microsoftAccountBlockSettingsSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftAccountBlockSettingsSync(val)
        }
        return nil
    }
    res["networkProxyApplySettingsDeviceWide"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkProxyApplySettingsDeviceWide(val)
        }
        return nil
    }
    res["networkProxyAutomaticConfigurationUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkProxyAutomaticConfigurationUrl(val)
        }
        return nil
    }
    res["networkProxyDisableAutoDetect"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkProxyDisableAutoDetect(val)
        }
        return nil
    }
    res["networkProxyServer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindows10NetworkProxyServerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkProxyServer(val.(Windows10NetworkProxyServerable))
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
    res["oneDriveDisableFileSync"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneDriveDisableFileSync(val)
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
    res["passwordRequireWhenResumeFromIdleState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRequireWhenResumeFromIdleState(val)
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
    res["personalizationDesktopImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalizationDesktopImageUrl(val)
        }
        return nil
    }
    res["personalizationLockScreenImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalizationLockScreenImageUrl(val)
        }
        return nil
    }
    res["privacyAdvertisingId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStateManagementSetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyAdvertisingId(val.(*StateManagementSetting))
        }
        return nil
    }
    res["privacyAutoAcceptPairingAndConsentPrompts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyAutoAcceptPairingAndConsentPrompts(val)
        }
        return nil
    }
    res["privacyBlockInputPersonalization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyBlockInputPersonalization(val)
        }
        return nil
    }
    res["resetProtectionModeBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResetProtectionModeBlocked(val)
        }
        return nil
    }
    res["safeSearchFilter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSafeSearchFilterType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafeSearchFilter(val.(*SafeSearchFilterType))
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
    res["searchBlockDiacritics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchBlockDiacritics(val)
        }
        return nil
    }
    res["searchDisableAutoLanguageDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchDisableAutoLanguageDetection(val)
        }
        return nil
    }
    res["searchDisableIndexerBackoff"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchDisableIndexerBackoff(val)
        }
        return nil
    }
    res["searchDisableIndexingEncryptedItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchDisableIndexingEncryptedItems(val)
        }
        return nil
    }
    res["searchDisableIndexingRemovableDrive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchDisableIndexingRemovableDrive(val)
        }
        return nil
    }
    res["searchEnableAutomaticIndexSizeManangement"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchEnableAutomaticIndexSizeManangement(val)
        }
        return nil
    }
    res["searchEnableRemoteQueries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSearchEnableRemoteQueries(val)
        }
        return nil
    }
    res["settingsBlockAccountsPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockAccountsPage(val)
        }
        return nil
    }
    res["settingsBlockAddProvisioningPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockAddProvisioningPackage(val)
        }
        return nil
    }
    res["settingsBlockAppsPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockAppsPage(val)
        }
        return nil
    }
    res["settingsBlockChangeLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockChangeLanguage(val)
        }
        return nil
    }
    res["settingsBlockChangePowerSleep"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockChangePowerSleep(val)
        }
        return nil
    }
    res["settingsBlockChangeRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockChangeRegion(val)
        }
        return nil
    }
    res["settingsBlockChangeSystemTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockChangeSystemTime(val)
        }
        return nil
    }
    res["settingsBlockDevicesPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockDevicesPage(val)
        }
        return nil
    }
    res["settingsBlockEaseOfAccessPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockEaseOfAccessPage(val)
        }
        return nil
    }
    res["settingsBlockEditDeviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockEditDeviceName(val)
        }
        return nil
    }
    res["settingsBlockGamingPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockGamingPage(val)
        }
        return nil
    }
    res["settingsBlockNetworkInternetPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockNetworkInternetPage(val)
        }
        return nil
    }
    res["settingsBlockPersonalizationPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockPersonalizationPage(val)
        }
        return nil
    }
    res["settingsBlockPrivacyPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockPrivacyPage(val)
        }
        return nil
    }
    res["settingsBlockRemoveProvisioningPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockRemoveProvisioningPackage(val)
        }
        return nil
    }
    res["settingsBlockSettingsApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockSettingsApp(val)
        }
        return nil
    }
    res["settingsBlockSystemPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockSystemPage(val)
        }
        return nil
    }
    res["settingsBlockTimeLanguagePage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockTimeLanguagePage(val)
        }
        return nil
    }
    res["settingsBlockUpdateSecurityPage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockUpdateSecurityPage(val)
        }
        return nil
    }
    res["sharedUserAppDataAllowed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedUserAppDataAllowed(val)
        }
        return nil
    }
    res["smartScreenBlockPromptOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenBlockPromptOverride(val)
        }
        return nil
    }
    res["smartScreenBlockPromptOverrideForFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenBlockPromptOverrideForFiles(val)
        }
        return nil
    }
    res["smartScreenEnableAppInstallControl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenEnableAppInstallControl(val)
        }
        return nil
    }
    res["startBlockUnpinningAppsFromTaskbar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartBlockUnpinningAppsFromTaskbar(val)
        }
        return nil
    }
    res["startMenuAppListVisibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsStartMenuAppListVisibilityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuAppListVisibility(val.(*WindowsStartMenuAppListVisibilityType))
        }
        return nil
    }
    res["startMenuHideChangeAccountSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideChangeAccountSettings(val)
        }
        return nil
    }
    res["startMenuHideFrequentlyUsedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideFrequentlyUsedApps(val)
        }
        return nil
    }
    res["startMenuHideHibernate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideHibernate(val)
        }
        return nil
    }
    res["startMenuHideLock"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideLock(val)
        }
        return nil
    }
    res["startMenuHidePowerButton"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHidePowerButton(val)
        }
        return nil
    }
    res["startMenuHideRecentJumpLists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideRecentJumpLists(val)
        }
        return nil
    }
    res["startMenuHideRecentlyAddedApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideRecentlyAddedApps(val)
        }
        return nil
    }
    res["startMenuHideRestartOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideRestartOptions(val)
        }
        return nil
    }
    res["startMenuHideShutDown"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideShutDown(val)
        }
        return nil
    }
    res["startMenuHideSignOut"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideSignOut(val)
        }
        return nil
    }
    res["startMenuHideSleep"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideSleep(val)
        }
        return nil
    }
    res["startMenuHideSwitchAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideSwitchAccount(val)
        }
        return nil
    }
    res["startMenuHideUserTile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuHideUserTile(val)
        }
        return nil
    }
    res["startMenuLayoutEdgeAssetsXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuLayoutEdgeAssetsXml(val)
        }
        return nil
    }
    res["startMenuLayoutXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuLayoutXml(val)
        }
        return nil
    }
    res["startMenuMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsStartMenuModeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuMode(val.(*WindowsStartMenuModeType))
        }
        return nil
    }
    res["startMenuPinnedFolderDocuments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderDocuments(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderDownloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderDownloads(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderFileExplorer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderFileExplorer(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderHomeGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderHomeGroup(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderMusic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderMusic(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderNetwork(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderPersonalFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderPersonalFolder(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderPictures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderPictures(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderSettings(val.(*VisibilitySetting))
        }
        return nil
    }
    res["startMenuPinnedFolderVideos"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseVisibilitySetting)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartMenuPinnedFolderVideos(val.(*VisibilitySetting))
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
    res["storageRequireMobileDeviceEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageRequireMobileDeviceEncryption(val)
        }
        return nil
    }
    res["storageRestrictAppDataToSystemVolume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageRestrictAppDataToSystemVolume(val)
        }
        return nil
    }
    res["storageRestrictAppInstallToSystemVolume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStorageRestrictAppInstallToSystemVolume(val)
        }
        return nil
    }
    res["tenantLockdownRequireNetworkDuringOutOfBoxExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantLockdownRequireNetworkDuringOutOfBoxExperience(val)
        }
        return nil
    }
    res["usbBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsbBlocked(val)
        }
        return nil
    }
    res["voiceRecordingBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVoiceRecordingBlocked(val)
        }
        return nil
    }
    res["webRtcBlockLocalhostIpAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebRtcBlockLocalhostIpAddress(val)
        }
        return nil
    }
    res["wiFiBlockAutomaticConnectHotspots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiBlockAutomaticConnectHotspots(val)
        }
        return nil
    }
    res["wiFiBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiBlocked(val)
        }
        return nil
    }
    res["wiFiBlockManualConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiBlockManualConfiguration(val)
        }
        return nil
    }
    res["wiFiScanInterval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWiFiScanInterval(val)
        }
        return nil
    }
    res["windowsSpotlightBlockConsumerSpecificFeatures"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSpotlightBlockConsumerSpecificFeatures(val)
        }
        return nil
    }
    res["windowsSpotlightBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSpotlightBlocked(val)
        }
        return nil
    }
    res["windowsSpotlightBlockOnActionCenter"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSpotlightBlockOnActionCenter(val)
        }
        return nil
    }
    res["windowsSpotlightBlockTailoredExperiences"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSpotlightBlockTailoredExperiences(val)
        }
        return nil
    }
    res["windowsSpotlightBlockThirdPartyNotifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSpotlightBlockThirdPartyNotifications(val)
        }
        return nil
    }
    res["windowsSpotlightBlockWelcomeExperience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSpotlightBlockWelcomeExperience(val)
        }
        return nil
    }
    res["windowsSpotlightBlockWindowsTips"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSpotlightBlockWindowsTips(val)
        }
        return nil
    }
    res["windowsSpotlightConfigureOnLockScreen"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindowsSpotlightEnablementSettings)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsSpotlightConfigureOnLockScreen(val.(*WindowsSpotlightEnablementSettings))
        }
        return nil
    }
    res["windowsStoreBlockAutoUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsStoreBlockAutoUpdate(val)
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
    res["windowsStoreEnablePrivateStoreOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsStoreEnablePrivateStoreOnly(val)
        }
        return nil
    }
    res["wirelessDisplayBlockProjectionToThisDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWirelessDisplayBlockProjectionToThisDevice(val)
        }
        return nil
    }
    res["wirelessDisplayBlockUserInputFromReceiver"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWirelessDisplayBlockUserInputFromReceiver(val)
        }
        return nil
    }
    res["wirelessDisplayRequirePinForPairing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWirelessDisplayRequirePinForPairing(val)
        }
        return nil
    }
    return res
}
// GetGameDvrBlocked gets the gameDvrBlocked property value. Indicates whether or not to block DVR and broadcasting.
func (m *Windows10GeneralConfiguration) GetGameDvrBlocked()(*bool) {
    return m.gameDvrBlocked
}
// GetInternetSharingBlocked gets the internetSharingBlocked property value. Indicates whether or not to Block the user from using internet sharing.
func (m *Windows10GeneralConfiguration) GetInternetSharingBlocked()(*bool) {
    return m.internetSharingBlocked
}
// GetLocationServicesBlocked gets the locationServicesBlocked property value. Indicates whether or not to Block the user from location services.
func (m *Windows10GeneralConfiguration) GetLocationServicesBlocked()(*bool) {
    return m.locationServicesBlocked
}
// GetLockScreenAllowTimeoutConfiguration gets the lockScreenAllowTimeoutConfiguration property value. Specify whether to show a user-configurable setting to control the screen timeout while on the lock screen of Windows 10 Mobile devices. If this policy is set to Allow, the value set by lockScreenTimeoutInSeconds is ignored.
func (m *Windows10GeneralConfiguration) GetLockScreenAllowTimeoutConfiguration()(*bool) {
    return m.lockScreenAllowTimeoutConfiguration
}
// GetLockScreenBlockActionCenterNotifications gets the lockScreenBlockActionCenterNotifications property value. Indicates whether or not to block action center notifications over lock screen.
func (m *Windows10GeneralConfiguration) GetLockScreenBlockActionCenterNotifications()(*bool) {
    return m.lockScreenBlockActionCenterNotifications
}
// GetLockScreenBlockCortana gets the lockScreenBlockCortana property value. Indicates whether or not the user can interact with Cortana using speech while the system is locked.
func (m *Windows10GeneralConfiguration) GetLockScreenBlockCortana()(*bool) {
    return m.lockScreenBlockCortana
}
// GetLockScreenBlockToastNotifications gets the lockScreenBlockToastNotifications property value. Indicates whether to allow toast notifications above the device lock screen.
func (m *Windows10GeneralConfiguration) GetLockScreenBlockToastNotifications()(*bool) {
    return m.lockScreenBlockToastNotifications
}
// GetLockScreenTimeoutInSeconds gets the lockScreenTimeoutInSeconds property value. Set the duration (in seconds) from the screen locking to the screen turning off for Windows 10 Mobile devices. Supported values are 11-1800. Valid values 11 to 1800
func (m *Windows10GeneralConfiguration) GetLockScreenTimeoutInSeconds()(*int32) {
    return m.lockScreenTimeoutInSeconds
}
// GetLogonBlockFastUserSwitching gets the logonBlockFastUserSwitching property value. Disables the ability to quickly switch between users that are logged on simultaneously without logging off.
func (m *Windows10GeneralConfiguration) GetLogonBlockFastUserSwitching()(*bool) {
    return m.logonBlockFastUserSwitching
}
// GetMicrosoftAccountBlocked gets the microsoftAccountBlocked property value. Indicates whether or not to Block a Microsoft account.
func (m *Windows10GeneralConfiguration) GetMicrosoftAccountBlocked()(*bool) {
    return m.microsoftAccountBlocked
}
// GetMicrosoftAccountBlockSettingsSync gets the microsoftAccountBlockSettingsSync property value. Indicates whether or not to Block Microsoft account settings sync.
func (m *Windows10GeneralConfiguration) GetMicrosoftAccountBlockSettingsSync()(*bool) {
    return m.microsoftAccountBlockSettingsSync
}
// GetNetworkProxyApplySettingsDeviceWide gets the networkProxyApplySettingsDeviceWide property value. If set, proxy settings will be applied to all processes and accounts in the device. Otherwise, it will be applied to the user account that’s enrolled into MDM.
func (m *Windows10GeneralConfiguration) GetNetworkProxyApplySettingsDeviceWide()(*bool) {
    return m.networkProxyApplySettingsDeviceWide
}
// GetNetworkProxyAutomaticConfigurationUrl gets the networkProxyAutomaticConfigurationUrl property value. Address to the proxy auto-config (PAC) script you want to use.
func (m *Windows10GeneralConfiguration) GetNetworkProxyAutomaticConfigurationUrl()(*string) {
    return m.networkProxyAutomaticConfigurationUrl
}
// GetNetworkProxyDisableAutoDetect gets the networkProxyDisableAutoDetect property value. Disable automatic detection of settings. If enabled, the system will try to find the path to a proxy auto-config (PAC) script.
func (m *Windows10GeneralConfiguration) GetNetworkProxyDisableAutoDetect()(*bool) {
    return m.networkProxyDisableAutoDetect
}
// GetNetworkProxyServer gets the networkProxyServer property value. Specifies manual proxy server settings.
func (m *Windows10GeneralConfiguration) GetNetworkProxyServer()(Windows10NetworkProxyServerable) {
    return m.networkProxyServer
}
// GetNfcBlocked gets the nfcBlocked property value. Indicates whether or not to Block the user from using near field communication.
func (m *Windows10GeneralConfiguration) GetNfcBlocked()(*bool) {
    return m.nfcBlocked
}
// GetOneDriveDisableFileSync gets the oneDriveDisableFileSync property value. Gets or sets a value allowing IT admins to prevent apps and features from working with files on OneDrive.
func (m *Windows10GeneralConfiguration) GetOneDriveDisableFileSync()(*bool) {
    return m.oneDriveDisableFileSync
}
// GetPasswordBlockSimple gets the passwordBlockSimple property value. Specify whether PINs or passwords such as '1111' or '1234' are allowed. For Windows 10 desktops, it also controls the use of picture passwords.
func (m *Windows10GeneralConfiguration) GetPasswordBlockSimple()(*bool) {
    return m.passwordBlockSimple
}
// GetPasswordExpirationDays gets the passwordExpirationDays property value. The password expiration in days. Valid values 0 to 730
func (m *Windows10GeneralConfiguration) GetPasswordExpirationDays()(*int32) {
    return m.passwordExpirationDays
}
// GetPasswordMinimumCharacterSetCount gets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows10GeneralConfiguration) GetPasswordMinimumCharacterSetCount()(*int32) {
    return m.passwordMinimumCharacterSetCount
}
// GetPasswordMinimumLength gets the passwordMinimumLength property value. The minimum password length. Valid values 4 to 16
func (m *Windows10GeneralConfiguration) GetPasswordMinimumLength()(*int32) {
    return m.passwordMinimumLength
}
// GetPasswordMinutesOfInactivityBeforeScreenTimeout gets the passwordMinutesOfInactivityBeforeScreenTimeout property value. The minutes of inactivity before the screen times out.
func (m *Windows10GeneralConfiguration) GetPasswordMinutesOfInactivityBeforeScreenTimeout()(*int32) {
    return m.passwordMinutesOfInactivityBeforeScreenTimeout
}
// GetPasswordPreviousPasswordBlockCount gets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent reuse of. Valid values 0 to 50
func (m *Windows10GeneralConfiguration) GetPasswordPreviousPasswordBlockCount()(*int32) {
    return m.passwordPreviousPasswordBlockCount
}
// GetPasswordRequired gets the passwordRequired property value. Indicates whether or not to require the user to have a password.
func (m *Windows10GeneralConfiguration) GetPasswordRequired()(*bool) {
    return m.passwordRequired
}
// GetPasswordRequiredType gets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows10GeneralConfiguration) GetPasswordRequiredType()(*RequiredPasswordType) {
    return m.passwordRequiredType
}
// GetPasswordRequireWhenResumeFromIdleState gets the passwordRequireWhenResumeFromIdleState property value. Indicates whether or not to require a password upon resuming from an idle state.
func (m *Windows10GeneralConfiguration) GetPasswordRequireWhenResumeFromIdleState()(*bool) {
    return m.passwordRequireWhenResumeFromIdleState
}
// GetPasswordSignInFailureCountBeforeFactoryReset gets the passwordSignInFailureCountBeforeFactoryReset property value. The number of sign in failures before factory reset. Valid values 0 to 999
func (m *Windows10GeneralConfiguration) GetPasswordSignInFailureCountBeforeFactoryReset()(*int32) {
    return m.passwordSignInFailureCountBeforeFactoryReset
}
// GetPersonalizationDesktopImageUrl gets the personalizationDesktopImageUrl property value. A http or https Url to a jpg, jpeg or png image that needs to be downloaded and used as the Desktop Image or a file Url to a local image on the file system that needs to used as the Desktop Image.
func (m *Windows10GeneralConfiguration) GetPersonalizationDesktopImageUrl()(*string) {
    return m.personalizationDesktopImageUrl
}
// GetPersonalizationLockScreenImageUrl gets the personalizationLockScreenImageUrl property value. A http or https Url to a jpg, jpeg or png image that neeeds to be downloaded and used as the Lock Screen Image or a file Url to a local image on the file system that needs to be used as the Lock Screen Image.
func (m *Windows10GeneralConfiguration) GetPersonalizationLockScreenImageUrl()(*string) {
    return m.personalizationLockScreenImageUrl
}
// GetPrivacyAdvertisingId gets the privacyAdvertisingId property value. State Management Setting.
func (m *Windows10GeneralConfiguration) GetPrivacyAdvertisingId()(*StateManagementSetting) {
    return m.privacyAdvertisingId
}
// GetPrivacyAutoAcceptPairingAndConsentPrompts gets the privacyAutoAcceptPairingAndConsentPrompts property value. Indicates whether or not to allow the automatic acceptance of the pairing and privacy user consent dialog when launching apps.
func (m *Windows10GeneralConfiguration) GetPrivacyAutoAcceptPairingAndConsentPrompts()(*bool) {
    return m.privacyAutoAcceptPairingAndConsentPrompts
}
// GetPrivacyBlockInputPersonalization gets the privacyBlockInputPersonalization property value. Indicates whether or not to block the usage of cloud based speech services for Cortana, Dictation, or Store applications.
func (m *Windows10GeneralConfiguration) GetPrivacyBlockInputPersonalization()(*bool) {
    return m.privacyBlockInputPersonalization
}
// GetResetProtectionModeBlocked gets the resetProtectionModeBlocked property value. Indicates whether or not to Block the user from reset protection mode.
func (m *Windows10GeneralConfiguration) GetResetProtectionModeBlocked()(*bool) {
    return m.resetProtectionModeBlocked
}
// GetSafeSearchFilter gets the safeSearchFilter property value. Specifies what level of safe search (filtering adult content) is required
func (m *Windows10GeneralConfiguration) GetSafeSearchFilter()(*SafeSearchFilterType) {
    return m.safeSearchFilter
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether or not to Block the user from taking Screenshots.
func (m *Windows10GeneralConfiguration) GetScreenCaptureBlocked()(*bool) {
    return m.screenCaptureBlocked
}
// GetSearchBlockDiacritics gets the searchBlockDiacritics property value. Specifies if search can use diacritics.
func (m *Windows10GeneralConfiguration) GetSearchBlockDiacritics()(*bool) {
    return m.searchBlockDiacritics
}
// GetSearchDisableAutoLanguageDetection gets the searchDisableAutoLanguageDetection property value. Specifies whether to use automatic language detection when indexing content and properties.
func (m *Windows10GeneralConfiguration) GetSearchDisableAutoLanguageDetection()(*bool) {
    return m.searchDisableAutoLanguageDetection
}
// GetSearchDisableIndexerBackoff gets the searchDisableIndexerBackoff property value. Indicates whether or not to disable the search indexer backoff feature.
func (m *Windows10GeneralConfiguration) GetSearchDisableIndexerBackoff()(*bool) {
    return m.searchDisableIndexerBackoff
}
// GetSearchDisableIndexingEncryptedItems gets the searchDisableIndexingEncryptedItems property value. Indicates whether or not to block indexing of WIP-protected items to prevent them from appearing in search results for Cortana or Explorer.
func (m *Windows10GeneralConfiguration) GetSearchDisableIndexingEncryptedItems()(*bool) {
    return m.searchDisableIndexingEncryptedItems
}
// GetSearchDisableIndexingRemovableDrive gets the searchDisableIndexingRemovableDrive property value. Indicates whether or not to allow users to add locations on removable drives to libraries and to be indexed.
func (m *Windows10GeneralConfiguration) GetSearchDisableIndexingRemovableDrive()(*bool) {
    return m.searchDisableIndexingRemovableDrive
}
// GetSearchEnableAutomaticIndexSizeManangement gets the searchEnableAutomaticIndexSizeManangement property value. Specifies minimum amount of hard drive space on the same drive as the index location before indexing stops.
func (m *Windows10GeneralConfiguration) GetSearchEnableAutomaticIndexSizeManangement()(*bool) {
    return m.searchEnableAutomaticIndexSizeManangement
}
// GetSearchEnableRemoteQueries gets the searchEnableRemoteQueries property value. Indicates whether or not to block remote queries of this computer’s index.
func (m *Windows10GeneralConfiguration) GetSearchEnableRemoteQueries()(*bool) {
    return m.searchEnableRemoteQueries
}
// GetSettingsBlockAccountsPage gets the settingsBlockAccountsPage property value. Indicates whether or not to block access to Accounts in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockAccountsPage()(*bool) {
    return m.settingsBlockAccountsPage
}
// GetSettingsBlockAddProvisioningPackage gets the settingsBlockAddProvisioningPackage property value. Indicates whether or not to block the user from installing provisioning packages.
func (m *Windows10GeneralConfiguration) GetSettingsBlockAddProvisioningPackage()(*bool) {
    return m.settingsBlockAddProvisioningPackage
}
// GetSettingsBlockAppsPage gets the settingsBlockAppsPage property value. Indicates whether or not to block access to Apps in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockAppsPage()(*bool) {
    return m.settingsBlockAppsPage
}
// GetSettingsBlockChangeLanguage gets the settingsBlockChangeLanguage property value. Indicates whether or not to block the user from changing the language settings.
func (m *Windows10GeneralConfiguration) GetSettingsBlockChangeLanguage()(*bool) {
    return m.settingsBlockChangeLanguage
}
// GetSettingsBlockChangePowerSleep gets the settingsBlockChangePowerSleep property value. Indicates whether or not to block the user from changing power and sleep settings.
func (m *Windows10GeneralConfiguration) GetSettingsBlockChangePowerSleep()(*bool) {
    return m.settingsBlockChangePowerSleep
}
// GetSettingsBlockChangeRegion gets the settingsBlockChangeRegion property value. Indicates whether or not to block the user from changing the region settings.
func (m *Windows10GeneralConfiguration) GetSettingsBlockChangeRegion()(*bool) {
    return m.settingsBlockChangeRegion
}
// GetSettingsBlockChangeSystemTime gets the settingsBlockChangeSystemTime property value. Indicates whether or not to block the user from changing date and time settings.
func (m *Windows10GeneralConfiguration) GetSettingsBlockChangeSystemTime()(*bool) {
    return m.settingsBlockChangeSystemTime
}
// GetSettingsBlockDevicesPage gets the settingsBlockDevicesPage property value. Indicates whether or not to block access to Devices in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockDevicesPage()(*bool) {
    return m.settingsBlockDevicesPage
}
// GetSettingsBlockEaseOfAccessPage gets the settingsBlockEaseOfAccessPage property value. Indicates whether or not to block access to Ease of Access in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockEaseOfAccessPage()(*bool) {
    return m.settingsBlockEaseOfAccessPage
}
// GetSettingsBlockEditDeviceName gets the settingsBlockEditDeviceName property value. Indicates whether or not to block the user from editing the device name.
func (m *Windows10GeneralConfiguration) GetSettingsBlockEditDeviceName()(*bool) {
    return m.settingsBlockEditDeviceName
}
// GetSettingsBlockGamingPage gets the settingsBlockGamingPage property value. Indicates whether or not to block access to Gaming in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockGamingPage()(*bool) {
    return m.settingsBlockGamingPage
}
// GetSettingsBlockNetworkInternetPage gets the settingsBlockNetworkInternetPage property value. Indicates whether or not to block access to Network & Internet in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockNetworkInternetPage()(*bool) {
    return m.settingsBlockNetworkInternetPage
}
// GetSettingsBlockPersonalizationPage gets the settingsBlockPersonalizationPage property value. Indicates whether or not to block access to Personalization in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockPersonalizationPage()(*bool) {
    return m.settingsBlockPersonalizationPage
}
// GetSettingsBlockPrivacyPage gets the settingsBlockPrivacyPage property value. Indicates whether or not to block access to Privacy in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockPrivacyPage()(*bool) {
    return m.settingsBlockPrivacyPage
}
// GetSettingsBlockRemoveProvisioningPackage gets the settingsBlockRemoveProvisioningPackage property value. Indicates whether or not to block the runtime configuration agent from removing provisioning packages.
func (m *Windows10GeneralConfiguration) GetSettingsBlockRemoveProvisioningPackage()(*bool) {
    return m.settingsBlockRemoveProvisioningPackage
}
// GetSettingsBlockSettingsApp gets the settingsBlockSettingsApp property value. Indicates whether or not to block access to Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockSettingsApp()(*bool) {
    return m.settingsBlockSettingsApp
}
// GetSettingsBlockSystemPage gets the settingsBlockSystemPage property value. Indicates whether or not to block access to System in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockSystemPage()(*bool) {
    return m.settingsBlockSystemPage
}
// GetSettingsBlockTimeLanguagePage gets the settingsBlockTimeLanguagePage property value. Indicates whether or not to block access to Time & Language in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockTimeLanguagePage()(*bool) {
    return m.settingsBlockTimeLanguagePage
}
// GetSettingsBlockUpdateSecurityPage gets the settingsBlockUpdateSecurityPage property value. Indicates whether or not to block access to Update & Security in Settings app.
func (m *Windows10GeneralConfiguration) GetSettingsBlockUpdateSecurityPage()(*bool) {
    return m.settingsBlockUpdateSecurityPage
}
// GetSharedUserAppDataAllowed gets the sharedUserAppDataAllowed property value. Indicates whether or not to block multiple users of the same app to share data.
func (m *Windows10GeneralConfiguration) GetSharedUserAppDataAllowed()(*bool) {
    return m.sharedUserAppDataAllowed
}
// GetSmartScreenBlockPromptOverride gets the smartScreenBlockPromptOverride property value. Indicates whether or not users can override SmartScreen Filter warnings about potentially malicious websites.
func (m *Windows10GeneralConfiguration) GetSmartScreenBlockPromptOverride()(*bool) {
    return m.smartScreenBlockPromptOverride
}
// GetSmartScreenBlockPromptOverrideForFiles gets the smartScreenBlockPromptOverrideForFiles property value. Indicates whether or not users can override the SmartScreen Filter warnings about downloading unverified files
func (m *Windows10GeneralConfiguration) GetSmartScreenBlockPromptOverrideForFiles()(*bool) {
    return m.smartScreenBlockPromptOverrideForFiles
}
// GetSmartScreenEnableAppInstallControl gets the smartScreenEnableAppInstallControl property value. This property will be deprecated in July 2019 and will be replaced by property SmartScreenAppInstallControl. Allows IT Admins to control whether users are allowed to install apps from places other than the Store.
func (m *Windows10GeneralConfiguration) GetSmartScreenEnableAppInstallControl()(*bool) {
    return m.smartScreenEnableAppInstallControl
}
// GetStartBlockUnpinningAppsFromTaskbar gets the startBlockUnpinningAppsFromTaskbar property value. Indicates whether or not to block the user from unpinning apps from taskbar.
func (m *Windows10GeneralConfiguration) GetStartBlockUnpinningAppsFromTaskbar()(*bool) {
    return m.startBlockUnpinningAppsFromTaskbar
}
// GetStartMenuAppListVisibility gets the startMenuAppListVisibility property value. Type of start menu app list visibility.
func (m *Windows10GeneralConfiguration) GetStartMenuAppListVisibility()(*WindowsStartMenuAppListVisibilityType) {
    return m.startMenuAppListVisibility
}
// GetStartMenuHideChangeAccountSettings gets the startMenuHideChangeAccountSettings property value. Enabling this policy hides the change account setting from appearing in the user tile in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideChangeAccountSettings()(*bool) {
    return m.startMenuHideChangeAccountSettings
}
// GetStartMenuHideFrequentlyUsedApps gets the startMenuHideFrequentlyUsedApps property value. Enabling this policy hides the most used apps from appearing on the start menu and disables the corresponding toggle in the Settings app.
func (m *Windows10GeneralConfiguration) GetStartMenuHideFrequentlyUsedApps()(*bool) {
    return m.startMenuHideFrequentlyUsedApps
}
// GetStartMenuHideHibernate gets the startMenuHideHibernate property value. Enabling this policy hides hibernate from appearing in the power button in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideHibernate()(*bool) {
    return m.startMenuHideHibernate
}
// GetStartMenuHideLock gets the startMenuHideLock property value. Enabling this policy hides lock from appearing in the user tile in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideLock()(*bool) {
    return m.startMenuHideLock
}
// GetStartMenuHidePowerButton gets the startMenuHidePowerButton property value. Enabling this policy hides the power button from appearing in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHidePowerButton()(*bool) {
    return m.startMenuHidePowerButton
}
// GetStartMenuHideRecentJumpLists gets the startMenuHideRecentJumpLists property value. Enabling this policy hides recent jump lists from appearing on the start menu/taskbar and disables the corresponding toggle in the Settings app.
func (m *Windows10GeneralConfiguration) GetStartMenuHideRecentJumpLists()(*bool) {
    return m.startMenuHideRecentJumpLists
}
// GetStartMenuHideRecentlyAddedApps gets the startMenuHideRecentlyAddedApps property value. Enabling this policy hides recently added apps from appearing on the start menu and disables the corresponding toggle in the Settings app.
func (m *Windows10GeneralConfiguration) GetStartMenuHideRecentlyAddedApps()(*bool) {
    return m.startMenuHideRecentlyAddedApps
}
// GetStartMenuHideRestartOptions gets the startMenuHideRestartOptions property value. Enabling this policy hides 'Restart/Update and Restart' from appearing in the power button in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideRestartOptions()(*bool) {
    return m.startMenuHideRestartOptions
}
// GetStartMenuHideShutDown gets the startMenuHideShutDown property value. Enabling this policy hides shut down/update and shut down from appearing in the power button in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideShutDown()(*bool) {
    return m.startMenuHideShutDown
}
// GetStartMenuHideSignOut gets the startMenuHideSignOut property value. Enabling this policy hides sign out from appearing in the user tile in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideSignOut()(*bool) {
    return m.startMenuHideSignOut
}
// GetStartMenuHideSleep gets the startMenuHideSleep property value. Enabling this policy hides sleep from appearing in the power button in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideSleep()(*bool) {
    return m.startMenuHideSleep
}
// GetStartMenuHideSwitchAccount gets the startMenuHideSwitchAccount property value. Enabling this policy hides switch account from appearing in the user tile in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideSwitchAccount()(*bool) {
    return m.startMenuHideSwitchAccount
}
// GetStartMenuHideUserTile gets the startMenuHideUserTile property value. Enabling this policy hides the user tile from appearing in the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuHideUserTile()(*bool) {
    return m.startMenuHideUserTile
}
// GetStartMenuLayoutEdgeAssetsXml gets the startMenuLayoutEdgeAssetsXml property value. This policy setting allows you to import Edge assets to be used with startMenuLayoutXml policy. Start layout can contain secondary tile from Edge app which looks for Edge local asset file. Edge local asset would not exist and cause Edge secondary tile to appear empty in this case. This policy only gets applied when startMenuLayoutXml policy is modified. The value should be a UTF-8 Base64 encoded byte array.
func (m *Windows10GeneralConfiguration) GetStartMenuLayoutEdgeAssetsXml()([]byte) {
    return m.startMenuLayoutEdgeAssetsXml
}
// GetStartMenuLayoutXml gets the startMenuLayoutXml property value. Allows admins to override the default Start menu layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in a UTF8 encoded byte array format.
func (m *Windows10GeneralConfiguration) GetStartMenuLayoutXml()([]byte) {
    return m.startMenuLayoutXml
}
// GetStartMenuMode gets the startMenuMode property value. Type of display modes for the start menu.
func (m *Windows10GeneralConfiguration) GetStartMenuMode()(*WindowsStartMenuModeType) {
    return m.startMenuMode
}
// GetStartMenuPinnedFolderDocuments gets the startMenuPinnedFolderDocuments property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderDocuments()(*VisibilitySetting) {
    return m.startMenuPinnedFolderDocuments
}
// GetStartMenuPinnedFolderDownloads gets the startMenuPinnedFolderDownloads property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderDownloads()(*VisibilitySetting) {
    return m.startMenuPinnedFolderDownloads
}
// GetStartMenuPinnedFolderFileExplorer gets the startMenuPinnedFolderFileExplorer property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderFileExplorer()(*VisibilitySetting) {
    return m.startMenuPinnedFolderFileExplorer
}
// GetStartMenuPinnedFolderHomeGroup gets the startMenuPinnedFolderHomeGroup property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderHomeGroup()(*VisibilitySetting) {
    return m.startMenuPinnedFolderHomeGroup
}
// GetStartMenuPinnedFolderMusic gets the startMenuPinnedFolderMusic property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderMusic()(*VisibilitySetting) {
    return m.startMenuPinnedFolderMusic
}
// GetStartMenuPinnedFolderNetwork gets the startMenuPinnedFolderNetwork property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderNetwork()(*VisibilitySetting) {
    return m.startMenuPinnedFolderNetwork
}
// GetStartMenuPinnedFolderPersonalFolder gets the startMenuPinnedFolderPersonalFolder property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderPersonalFolder()(*VisibilitySetting) {
    return m.startMenuPinnedFolderPersonalFolder
}
// GetStartMenuPinnedFolderPictures gets the startMenuPinnedFolderPictures property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderPictures()(*VisibilitySetting) {
    return m.startMenuPinnedFolderPictures
}
// GetStartMenuPinnedFolderSettings gets the startMenuPinnedFolderSettings property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderSettings()(*VisibilitySetting) {
    return m.startMenuPinnedFolderSettings
}
// GetStartMenuPinnedFolderVideos gets the startMenuPinnedFolderVideos property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) GetStartMenuPinnedFolderVideos()(*VisibilitySetting) {
    return m.startMenuPinnedFolderVideos
}
// GetStorageBlockRemovableStorage gets the storageBlockRemovableStorage property value. Indicates whether or not to Block the user from using removable storage.
func (m *Windows10GeneralConfiguration) GetStorageBlockRemovableStorage()(*bool) {
    return m.storageBlockRemovableStorage
}
// GetStorageRequireMobileDeviceEncryption gets the storageRequireMobileDeviceEncryption property value. Indicating whether or not to require encryption on a mobile device.
func (m *Windows10GeneralConfiguration) GetStorageRequireMobileDeviceEncryption()(*bool) {
    return m.storageRequireMobileDeviceEncryption
}
// GetStorageRestrictAppDataToSystemVolume gets the storageRestrictAppDataToSystemVolume property value. Indicates whether application data is restricted to the system drive.
func (m *Windows10GeneralConfiguration) GetStorageRestrictAppDataToSystemVolume()(*bool) {
    return m.storageRestrictAppDataToSystemVolume
}
// GetStorageRestrictAppInstallToSystemVolume gets the storageRestrictAppInstallToSystemVolume property value. Indicates whether the installation of applications is restricted to the system drive.
func (m *Windows10GeneralConfiguration) GetStorageRestrictAppInstallToSystemVolume()(*bool) {
    return m.storageRestrictAppInstallToSystemVolume
}
// GetTenantLockdownRequireNetworkDuringOutOfBoxExperience gets the tenantLockdownRequireNetworkDuringOutOfBoxExperience property value. Whether the device is required to connect to the network.
func (m *Windows10GeneralConfiguration) GetTenantLockdownRequireNetworkDuringOutOfBoxExperience()(*bool) {
    return m.tenantLockdownRequireNetworkDuringOutOfBoxExperience
}
// GetUsbBlocked gets the usbBlocked property value. Indicates whether or not to Block the user from USB connection.
func (m *Windows10GeneralConfiguration) GetUsbBlocked()(*bool) {
    return m.usbBlocked
}
// GetVoiceRecordingBlocked gets the voiceRecordingBlocked property value. Indicates whether or not to Block the user from voice recording.
func (m *Windows10GeneralConfiguration) GetVoiceRecordingBlocked()(*bool) {
    return m.voiceRecordingBlocked
}
// GetWebRtcBlockLocalhostIpAddress gets the webRtcBlockLocalhostIpAddress property value. Indicates whether or not user's localhost IP address is displayed while making phone calls using the WebRTC
func (m *Windows10GeneralConfiguration) GetWebRtcBlockLocalhostIpAddress()(*bool) {
    return m.webRtcBlockLocalhostIpAddress
}
// GetWiFiBlockAutomaticConnectHotspots gets the wiFiBlockAutomaticConnectHotspots property value. Indicating whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
func (m *Windows10GeneralConfiguration) GetWiFiBlockAutomaticConnectHotspots()(*bool) {
    return m.wiFiBlockAutomaticConnectHotspots
}
// GetWiFiBlocked gets the wiFiBlocked property value. Indicates whether or not to Block the user from using Wi-Fi.
func (m *Windows10GeneralConfiguration) GetWiFiBlocked()(*bool) {
    return m.wiFiBlocked
}
// GetWiFiBlockManualConfiguration gets the wiFiBlockManualConfiguration property value. Indicates whether or not to Block the user from using Wi-Fi manual configuration.
func (m *Windows10GeneralConfiguration) GetWiFiBlockManualConfiguration()(*bool) {
    return m.wiFiBlockManualConfiguration
}
// GetWiFiScanInterval gets the wiFiScanInterval property value. Specify how often devices scan for Wi-Fi networks. Supported values are 1-500, where 100 = default, and 500 = low frequency. Valid values 1 to 500
func (m *Windows10GeneralConfiguration) GetWiFiScanInterval()(*int32) {
    return m.wiFiScanInterval
}
// GetWindowsSpotlightBlockConsumerSpecificFeatures gets the windowsSpotlightBlockConsumerSpecificFeatures property value. Allows IT admins to block experiences that are typically for consumers only, such as Start suggestions, Membership notifications, Post-OOBE app install and redirect tiles.
func (m *Windows10GeneralConfiguration) GetWindowsSpotlightBlockConsumerSpecificFeatures()(*bool) {
    return m.windowsSpotlightBlockConsumerSpecificFeatures
}
// GetWindowsSpotlightBlocked gets the windowsSpotlightBlocked property value. Allows IT admins to turn off all Windows Spotlight features
func (m *Windows10GeneralConfiguration) GetWindowsSpotlightBlocked()(*bool) {
    return m.windowsSpotlightBlocked
}
// GetWindowsSpotlightBlockOnActionCenter gets the windowsSpotlightBlockOnActionCenter property value. Block suggestions from Microsoft that show after each OS clean install, upgrade or in an on-going basis to introduce users to what is new or changed
func (m *Windows10GeneralConfiguration) GetWindowsSpotlightBlockOnActionCenter()(*bool) {
    return m.windowsSpotlightBlockOnActionCenter
}
// GetWindowsSpotlightBlockTailoredExperiences gets the windowsSpotlightBlockTailoredExperiences property value. Block personalized content in Windows spotlight based on user’s device usage.
func (m *Windows10GeneralConfiguration) GetWindowsSpotlightBlockTailoredExperiences()(*bool) {
    return m.windowsSpotlightBlockTailoredExperiences
}
// GetWindowsSpotlightBlockThirdPartyNotifications gets the windowsSpotlightBlockThirdPartyNotifications property value. Block third party content delivered via Windows Spotlight
func (m *Windows10GeneralConfiguration) GetWindowsSpotlightBlockThirdPartyNotifications()(*bool) {
    return m.windowsSpotlightBlockThirdPartyNotifications
}
// GetWindowsSpotlightBlockWelcomeExperience gets the windowsSpotlightBlockWelcomeExperience property value. Block Windows Spotlight Windows welcome experience
func (m *Windows10GeneralConfiguration) GetWindowsSpotlightBlockWelcomeExperience()(*bool) {
    return m.windowsSpotlightBlockWelcomeExperience
}
// GetWindowsSpotlightBlockWindowsTips gets the windowsSpotlightBlockWindowsTips property value. Allows IT admins to turn off the popup of Windows Tips.
func (m *Windows10GeneralConfiguration) GetWindowsSpotlightBlockWindowsTips()(*bool) {
    return m.windowsSpotlightBlockWindowsTips
}
// GetWindowsSpotlightConfigureOnLockScreen gets the windowsSpotlightConfigureOnLockScreen property value. Allows IT admind to set a predefined default search engine for MDM-Controlled devices
func (m *Windows10GeneralConfiguration) GetWindowsSpotlightConfigureOnLockScreen()(*WindowsSpotlightEnablementSettings) {
    return m.windowsSpotlightConfigureOnLockScreen
}
// GetWindowsStoreBlockAutoUpdate gets the windowsStoreBlockAutoUpdate property value. Indicates whether or not to block automatic update of apps from Windows Store.
func (m *Windows10GeneralConfiguration) GetWindowsStoreBlockAutoUpdate()(*bool) {
    return m.windowsStoreBlockAutoUpdate
}
// GetWindowsStoreBlocked gets the windowsStoreBlocked property value. Indicates whether or not to Block the user from using the Windows store.
func (m *Windows10GeneralConfiguration) GetWindowsStoreBlocked()(*bool) {
    return m.windowsStoreBlocked
}
// GetWindowsStoreEnablePrivateStoreOnly gets the windowsStoreEnablePrivateStoreOnly property value. Indicates whether or not to enable Private Store Only.
func (m *Windows10GeneralConfiguration) GetWindowsStoreEnablePrivateStoreOnly()(*bool) {
    return m.windowsStoreEnablePrivateStoreOnly
}
// GetWirelessDisplayBlockProjectionToThisDevice gets the wirelessDisplayBlockProjectionToThisDevice property value. Indicates whether or not to allow other devices from discovering this PC for projection.
func (m *Windows10GeneralConfiguration) GetWirelessDisplayBlockProjectionToThisDevice()(*bool) {
    return m.wirelessDisplayBlockProjectionToThisDevice
}
// GetWirelessDisplayBlockUserInputFromReceiver gets the wirelessDisplayBlockUserInputFromReceiver property value. Indicates whether or not to allow user input from wireless display receiver.
func (m *Windows10GeneralConfiguration) GetWirelessDisplayBlockUserInputFromReceiver()(*bool) {
    return m.wirelessDisplayBlockUserInputFromReceiver
}
// GetWirelessDisplayRequirePinForPairing gets the wirelessDisplayRequirePinForPairing property value. Indicates whether or not to require a PIN for new devices to initiate pairing.
func (m *Windows10GeneralConfiguration) GetWirelessDisplayRequirePinForPairing()(*bool) {
    return m.wirelessDisplayRequirePinForPairing
}
// Serialize serializes information the current object
func (m *Windows10GeneralConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("antiTheftModeBlocked", m.GetAntiTheftModeBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetAppsAllowTrustedAppsSideloading() != nil {
        cast := (*m.GetAppsAllowTrustedAppsSideloading()).String()
        err = writer.WriteStringValue("appsAllowTrustedAppsSideloading", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("appsBlockWindowsStoreOriginatedApps", m.GetAppsBlockWindowsStoreOriginatedApps())
        if err != nil {
            return err
        }
    }
    if m.GetBluetoothAllowedServices() != nil {
        err = writer.WriteCollectionOfStringValues("bluetoothAllowedServices", m.GetBluetoothAllowedServices())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockAdvertising", m.GetBluetoothBlockAdvertising())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bluetoothBlockDiscoverableMode", m.GetBluetoothBlockDiscoverableMode())
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
        err = writer.WriteBoolValue("bluetoothBlockPrePairing", m.GetBluetoothBlockPrePairing())
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
        err = writer.WriteBoolValue("cellularBlockDataWhenRoaming", m.GetCellularBlockDataWhenRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockVpn", m.GetCellularBlockVpn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cellularBlockVpnWhenRoaming", m.GetCellularBlockVpnWhenRoaming())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("certificatesBlockManualRootCertificateInstallation", m.GetCertificatesBlockManualRootCertificateInstallation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("connectedDevicesServiceBlocked", m.GetConnectedDevicesServiceBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("copyPasteBlocked", m.GetCopyPasteBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("cortanaBlocked", m.GetCortanaBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderBlockEndUserAccess", m.GetDefenderBlockEndUserAccess())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderCloudBlockLevel() != nil {
        cast := (*m.GetDefenderCloudBlockLevel()).String()
        err = writer.WriteStringValue("defenderCloudBlockLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderDaysBeforeDeletingQuarantinedMalware", m.GetDefenderDaysBeforeDeletingQuarantinedMalware())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("defenderDetectedMalwareActions", m.GetDefenderDetectedMalwareActions())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderFileExtensionsToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderFileExtensionsToExclude", m.GetDefenderFileExtensionsToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderFilesAndFoldersToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderFilesAndFoldersToExclude", m.GetDefenderFilesAndFoldersToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderMonitorFileActivity() != nil {
        cast := (*m.GetDefenderMonitorFileActivity()).String()
        err = writer.WriteStringValue("defenderMonitorFileActivity", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefenderProcessesToExclude() != nil {
        err = writer.WriteCollectionOfStringValues("defenderProcessesToExclude", m.GetDefenderProcessesToExclude())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderPromptForSampleSubmission() != nil {
        cast := (*m.GetDefenderPromptForSampleSubmission()).String()
        err = writer.WriteStringValue("defenderPromptForSampleSubmission", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderRequireBehaviorMonitoring", m.GetDefenderRequireBehaviorMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderRequireCloudProtection", m.GetDefenderRequireCloudProtection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderRequireNetworkInspectionSystem", m.GetDefenderRequireNetworkInspectionSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderRequireRealTimeMonitoring", m.GetDefenderRequireRealTimeMonitoring())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderScanArchiveFiles", m.GetDefenderScanArchiveFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderScanDownloads", m.GetDefenderScanDownloads())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderScanIncomingMail", m.GetDefenderScanIncomingMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderScanMappedNetworkDrivesDuringFullScan", m.GetDefenderScanMappedNetworkDrivesDuringFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderScanMaxCpu", m.GetDefenderScanMaxCpu())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderScanNetworkFiles", m.GetDefenderScanNetworkFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderScanRemovableDrivesDuringFullScan", m.GetDefenderScanRemovableDrivesDuringFullScan())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderScanScriptsLoadedInInternetExplorer", m.GetDefenderScanScriptsLoadedInInternetExplorer())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderScanType() != nil {
        cast := (*m.GetDefenderScanType()).String()
        err = writer.WriteStringValue("defenderScanType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("defenderScheduledQuickScanTime", m.GetDefenderScheduledQuickScanTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("defenderScheduledScanTime", m.GetDefenderScheduledScanTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defenderSignatureUpdateIntervalInHours", m.GetDefenderSignatureUpdateIntervalInHours())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderSystemScanSchedule() != nil {
        cast := (*m.GetDefenderSystemScanSchedule()).String()
        err = writer.WriteStringValue("defenderSystemScanSchedule", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeveloperUnlockSetting() != nil {
        cast := (*m.GetDeveloperUnlockSetting()).String()
        err = writer.WriteStringValue("developerUnlockSetting", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceManagementBlockFactoryResetOnMobile", m.GetDeviceManagementBlockFactoryResetOnMobile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceManagementBlockManualUnenroll", m.GetDeviceManagementBlockManualUnenroll())
        if err != nil {
            return err
        }
    }
    if m.GetDiagnosticsDataSubmissionMode() != nil {
        cast := (*m.GetDiagnosticsDataSubmissionMode()).String()
        err = writer.WriteStringValue("diagnosticsDataSubmissionMode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeAllowStartPagesModification", m.GetEdgeAllowStartPagesModification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockAccessToAboutFlags", m.GetEdgeBlockAccessToAboutFlags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockAddressBarDropdown", m.GetEdgeBlockAddressBarDropdown())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockAutofill", m.GetEdgeBlockAutofill())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockCompatibilityList", m.GetEdgeBlockCompatibilityList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockDeveloperTools", m.GetEdgeBlockDeveloperTools())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlocked", m.GetEdgeBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockExtensions", m.GetEdgeBlockExtensions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockInPrivateBrowsing", m.GetEdgeBlockInPrivateBrowsing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockJavaScript", m.GetEdgeBlockJavaScript())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockLiveTileDataCollection", m.GetEdgeBlockLiveTileDataCollection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockPasswordManager", m.GetEdgeBlockPasswordManager())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockPopups", m.GetEdgeBlockPopups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockSearchSuggestions", m.GetEdgeBlockSearchSuggestions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockSendingDoNotTrackHeader", m.GetEdgeBlockSendingDoNotTrackHeader())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeBlockSendingIntranetTrafficToInternetExplorer", m.GetEdgeBlockSendingIntranetTrafficToInternetExplorer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeClearBrowsingDataOnExit", m.GetEdgeClearBrowsingDataOnExit())
        if err != nil {
            return err
        }
    }
    if m.GetEdgeCookiePolicy() != nil {
        cast := (*m.GetEdgeCookiePolicy()).String()
        err = writer.WriteStringValue("edgeCookiePolicy", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeDisableFirstRunPage", m.GetEdgeDisableFirstRunPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("edgeEnterpriseModeSiteListLocation", m.GetEdgeEnterpriseModeSiteListLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("edgeFirstRunUrl", m.GetEdgeFirstRunUrl())
        if err != nil {
            return err
        }
    }
    if m.GetEdgeHomepageUrls() != nil {
        err = writer.WriteCollectionOfStringValues("edgeHomepageUrls", m.GetEdgeHomepageUrls())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeRequireSmartScreen", m.GetEdgeRequireSmartScreen())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("edgeSearchEngine", m.GetEdgeSearchEngine())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeSendIntranetTrafficToInternetExplorer", m.GetEdgeSendIntranetTrafficToInternetExplorer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("edgeSyncFavoritesWithInternetExplorer", m.GetEdgeSyncFavoritesWithInternetExplorer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enterpriseCloudPrintDiscoveryEndPoint", m.GetEnterpriseCloudPrintDiscoveryEndPoint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("enterpriseCloudPrintDiscoveryMaxLimit", m.GetEnterpriseCloudPrintDiscoveryMaxLimit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enterpriseCloudPrintMopriaDiscoveryResourceIdentifier", m.GetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enterpriseCloudPrintOAuthAuthority", m.GetEnterpriseCloudPrintOAuthAuthority())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enterpriseCloudPrintOAuthClientIdentifier", m.GetEnterpriseCloudPrintOAuthClientIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("enterpriseCloudPrintResourceIdentifier", m.GetEnterpriseCloudPrintResourceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("experienceBlockDeviceDiscovery", m.GetExperienceBlockDeviceDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("experienceBlockErrorDialogWhenNoSIM", m.GetExperienceBlockErrorDialogWhenNoSIM())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("experienceBlockTaskSwitcher", m.GetExperienceBlockTaskSwitcher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("gameDvrBlocked", m.GetGameDvrBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("internetSharingBlocked", m.GetInternetSharingBlocked())
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
        err = writer.WriteBoolValue("lockScreenAllowTimeoutConfiguration", m.GetLockScreenAllowTimeoutConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockActionCenterNotifications", m.GetLockScreenBlockActionCenterNotifications())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockCortana", m.GetLockScreenBlockCortana())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("lockScreenBlockToastNotifications", m.GetLockScreenBlockToastNotifications())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("lockScreenTimeoutInSeconds", m.GetLockScreenTimeoutInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("logonBlockFastUserSwitching", m.GetLogonBlockFastUserSwitching())
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
        err = writer.WriteBoolValue("microsoftAccountBlockSettingsSync", m.GetMicrosoftAccountBlockSettingsSync())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("networkProxyApplySettingsDeviceWide", m.GetNetworkProxyApplySettingsDeviceWide())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("networkProxyAutomaticConfigurationUrl", m.GetNetworkProxyAutomaticConfigurationUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("networkProxyDisableAutoDetect", m.GetNetworkProxyDisableAutoDetect())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("networkProxyServer", m.GetNetworkProxyServer())
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
        err = writer.WriteBoolValue("oneDriveDisableFileSync", m.GetOneDriveDisableFileSync())
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
        err = writer.WriteBoolValue("passwordRequireWhenResumeFromIdleState", m.GetPasswordRequireWhenResumeFromIdleState())
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
        err = writer.WriteStringValue("personalizationDesktopImageUrl", m.GetPersonalizationDesktopImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personalizationLockScreenImageUrl", m.GetPersonalizationLockScreenImageUrl())
        if err != nil {
            return err
        }
    }
    if m.GetPrivacyAdvertisingId() != nil {
        cast := (*m.GetPrivacyAdvertisingId()).String()
        err = writer.WriteStringValue("privacyAdvertisingId", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("privacyAutoAcceptPairingAndConsentPrompts", m.GetPrivacyAutoAcceptPairingAndConsentPrompts())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("privacyBlockInputPersonalization", m.GetPrivacyBlockInputPersonalization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("resetProtectionModeBlocked", m.GetResetProtectionModeBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetSafeSearchFilter() != nil {
        cast := (*m.GetSafeSearchFilter()).String()
        err = writer.WriteStringValue("safeSearchFilter", &cast)
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
        err = writer.WriteBoolValue("searchBlockDiacritics", m.GetSearchBlockDiacritics())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("searchDisableAutoLanguageDetection", m.GetSearchDisableAutoLanguageDetection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("searchDisableIndexerBackoff", m.GetSearchDisableIndexerBackoff())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("searchDisableIndexingEncryptedItems", m.GetSearchDisableIndexingEncryptedItems())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("searchDisableIndexingRemovableDrive", m.GetSearchDisableIndexingRemovableDrive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("searchEnableAutomaticIndexSizeManangement", m.GetSearchEnableAutomaticIndexSizeManangement())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("searchEnableRemoteQueries", m.GetSearchEnableRemoteQueries())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockAccountsPage", m.GetSettingsBlockAccountsPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockAddProvisioningPackage", m.GetSettingsBlockAddProvisioningPackage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockAppsPage", m.GetSettingsBlockAppsPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockChangeLanguage", m.GetSettingsBlockChangeLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockChangePowerSleep", m.GetSettingsBlockChangePowerSleep())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockChangeRegion", m.GetSettingsBlockChangeRegion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockChangeSystemTime", m.GetSettingsBlockChangeSystemTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockDevicesPage", m.GetSettingsBlockDevicesPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockEaseOfAccessPage", m.GetSettingsBlockEaseOfAccessPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockEditDeviceName", m.GetSettingsBlockEditDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockGamingPage", m.GetSettingsBlockGamingPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockNetworkInternetPage", m.GetSettingsBlockNetworkInternetPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockPersonalizationPage", m.GetSettingsBlockPersonalizationPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockPrivacyPage", m.GetSettingsBlockPrivacyPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockRemoveProvisioningPackage", m.GetSettingsBlockRemoveProvisioningPackage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockSettingsApp", m.GetSettingsBlockSettingsApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockSystemPage", m.GetSettingsBlockSystemPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockTimeLanguagePage", m.GetSettingsBlockTimeLanguagePage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockUpdateSecurityPage", m.GetSettingsBlockUpdateSecurityPage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sharedUserAppDataAllowed", m.GetSharedUserAppDataAllowed())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenBlockPromptOverride", m.GetSmartScreenBlockPromptOverride())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenBlockPromptOverrideForFiles", m.GetSmartScreenBlockPromptOverrideForFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenEnableAppInstallControl", m.GetSmartScreenEnableAppInstallControl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startBlockUnpinningAppsFromTaskbar", m.GetStartBlockUnpinningAppsFromTaskbar())
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuAppListVisibility() != nil {
        cast := (*m.GetStartMenuAppListVisibility()).String()
        err = writer.WriteStringValue("startMenuAppListVisibility", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideChangeAccountSettings", m.GetStartMenuHideChangeAccountSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideFrequentlyUsedApps", m.GetStartMenuHideFrequentlyUsedApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideHibernate", m.GetStartMenuHideHibernate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideLock", m.GetStartMenuHideLock())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHidePowerButton", m.GetStartMenuHidePowerButton())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideRecentJumpLists", m.GetStartMenuHideRecentJumpLists())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideRecentlyAddedApps", m.GetStartMenuHideRecentlyAddedApps())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideRestartOptions", m.GetStartMenuHideRestartOptions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideShutDown", m.GetStartMenuHideShutDown())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideSignOut", m.GetStartMenuHideSignOut())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideSleep", m.GetStartMenuHideSleep())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideSwitchAccount", m.GetStartMenuHideSwitchAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("startMenuHideUserTile", m.GetStartMenuHideUserTile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("startMenuLayoutEdgeAssetsXml", m.GetStartMenuLayoutEdgeAssetsXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("startMenuLayoutXml", m.GetStartMenuLayoutXml())
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuMode() != nil {
        cast := (*m.GetStartMenuMode()).String()
        err = writer.WriteStringValue("startMenuMode", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderDocuments() != nil {
        cast := (*m.GetStartMenuPinnedFolderDocuments()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderDocuments", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderDownloads() != nil {
        cast := (*m.GetStartMenuPinnedFolderDownloads()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderDownloads", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderFileExplorer() != nil {
        cast := (*m.GetStartMenuPinnedFolderFileExplorer()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderFileExplorer", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderHomeGroup() != nil {
        cast := (*m.GetStartMenuPinnedFolderHomeGroup()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderHomeGroup", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderMusic() != nil {
        cast := (*m.GetStartMenuPinnedFolderMusic()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderMusic", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderNetwork() != nil {
        cast := (*m.GetStartMenuPinnedFolderNetwork()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderNetwork", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderPersonalFolder() != nil {
        cast := (*m.GetStartMenuPinnedFolderPersonalFolder()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderPersonalFolder", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderPictures() != nil {
        cast := (*m.GetStartMenuPinnedFolderPictures()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderPictures", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderSettings() != nil {
        cast := (*m.GetStartMenuPinnedFolderSettings()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderSettings", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStartMenuPinnedFolderVideos() != nil {
        cast := (*m.GetStartMenuPinnedFolderVideos()).String()
        err = writer.WriteStringValue("startMenuPinnedFolderVideos", &cast)
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
        err = writer.WriteBoolValue("storageRequireMobileDeviceEncryption", m.GetStorageRequireMobileDeviceEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRestrictAppDataToSystemVolume", m.GetStorageRestrictAppDataToSystemVolume())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("storageRestrictAppInstallToSystemVolume", m.GetStorageRestrictAppInstallToSystemVolume())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("tenantLockdownRequireNetworkDuringOutOfBoxExperience", m.GetTenantLockdownRequireNetworkDuringOutOfBoxExperience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usbBlocked", m.GetUsbBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("voiceRecordingBlocked", m.GetVoiceRecordingBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("webRtcBlockLocalhostIpAddress", m.GetWebRtcBlockLocalhostIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wiFiBlockAutomaticConnectHotspots", m.GetWiFiBlockAutomaticConnectHotspots())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wiFiBlocked", m.GetWiFiBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wiFiBlockManualConfiguration", m.GetWiFiBlockManualConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("wiFiScanInterval", m.GetWiFiScanInterval())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsSpotlightBlockConsumerSpecificFeatures", m.GetWindowsSpotlightBlockConsumerSpecificFeatures())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsSpotlightBlocked", m.GetWindowsSpotlightBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsSpotlightBlockOnActionCenter", m.GetWindowsSpotlightBlockOnActionCenter())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsSpotlightBlockTailoredExperiences", m.GetWindowsSpotlightBlockTailoredExperiences())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsSpotlightBlockThirdPartyNotifications", m.GetWindowsSpotlightBlockThirdPartyNotifications())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsSpotlightBlockWelcomeExperience", m.GetWindowsSpotlightBlockWelcomeExperience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsSpotlightBlockWindowsTips", m.GetWindowsSpotlightBlockWindowsTips())
        if err != nil {
            return err
        }
    }
    if m.GetWindowsSpotlightConfigureOnLockScreen() != nil {
        cast := (*m.GetWindowsSpotlightConfigureOnLockScreen()).String()
        err = writer.WriteStringValue("windowsSpotlightConfigureOnLockScreen", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsStoreBlockAutoUpdate", m.GetWindowsStoreBlockAutoUpdate())
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
    {
        err = writer.WriteBoolValue("windowsStoreEnablePrivateStoreOnly", m.GetWindowsStoreEnablePrivateStoreOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wirelessDisplayBlockProjectionToThisDevice", m.GetWirelessDisplayBlockProjectionToThisDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wirelessDisplayBlockUserInputFromReceiver", m.GetWirelessDisplayBlockUserInputFromReceiver())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("wirelessDisplayRequirePinForPairing", m.GetWirelessDisplayRequirePinForPairing())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountsBlockAddingNonMicrosoftAccountEmail sets the accountsBlockAddingNonMicrosoftAccountEmail property value. Indicates whether or not to Block the user from adding email accounts to the device that are not associated with a Microsoft account.
func (m *Windows10GeneralConfiguration) SetAccountsBlockAddingNonMicrosoftAccountEmail(value *bool)() {
    m.accountsBlockAddingNonMicrosoftAccountEmail = value
}
// SetAntiTheftModeBlocked sets the antiTheftModeBlocked property value. Indicates whether or not to block the user from selecting an AntiTheft mode preference (Windows 10 Mobile only).
func (m *Windows10GeneralConfiguration) SetAntiTheftModeBlocked(value *bool)() {
    m.antiTheftModeBlocked = value
}
// SetAppsAllowTrustedAppsSideloading sets the appsAllowTrustedAppsSideloading property value. State Management Setting.
func (m *Windows10GeneralConfiguration) SetAppsAllowTrustedAppsSideloading(value *StateManagementSetting)() {
    m.appsAllowTrustedAppsSideloading = value
}
// SetAppsBlockWindowsStoreOriginatedApps sets the appsBlockWindowsStoreOriginatedApps property value. Indicates whether or not to disable the launch of all apps from Windows Store that came pre-installed or were downloaded.
func (m *Windows10GeneralConfiguration) SetAppsBlockWindowsStoreOriginatedApps(value *bool)() {
    m.appsBlockWindowsStoreOriginatedApps = value
}
// SetBluetoothAllowedServices sets the bluetoothAllowedServices property value. Specify a list of allowed Bluetooth services and profiles in hex formatted strings.
func (m *Windows10GeneralConfiguration) SetBluetoothAllowedServices(value []string)() {
    m.bluetoothAllowedServices = value
}
// SetBluetoothBlockAdvertising sets the bluetoothBlockAdvertising property value. Whether or not to Block the user from using bluetooth advertising.
func (m *Windows10GeneralConfiguration) SetBluetoothBlockAdvertising(value *bool)() {
    m.bluetoothBlockAdvertising = value
}
// SetBluetoothBlockDiscoverableMode sets the bluetoothBlockDiscoverableMode property value. Whether or not to Block the user from using bluetooth discoverable mode.
func (m *Windows10GeneralConfiguration) SetBluetoothBlockDiscoverableMode(value *bool)() {
    m.bluetoothBlockDiscoverableMode = value
}
// SetBluetoothBlocked sets the bluetoothBlocked property value. Whether or not to Block the user from using bluetooth.
func (m *Windows10GeneralConfiguration) SetBluetoothBlocked(value *bool)() {
    m.bluetoothBlocked = value
}
// SetBluetoothBlockPrePairing sets the bluetoothBlockPrePairing property value. Whether or not to block specific bundled Bluetooth peripherals to automatically pair with the host device.
func (m *Windows10GeneralConfiguration) SetBluetoothBlockPrePairing(value *bool)() {
    m.bluetoothBlockPrePairing = value
}
// SetCameraBlocked sets the cameraBlocked property value. Whether or not to Block the user from accessing the camera of the device.
func (m *Windows10GeneralConfiguration) SetCameraBlocked(value *bool)() {
    m.cameraBlocked = value
}
// SetCellularBlockDataWhenRoaming sets the cellularBlockDataWhenRoaming property value. Whether or not to Block the user from using data over cellular while roaming.
func (m *Windows10GeneralConfiguration) SetCellularBlockDataWhenRoaming(value *bool)() {
    m.cellularBlockDataWhenRoaming = value
}
// SetCellularBlockVpn sets the cellularBlockVpn property value. Whether or not to Block the user from using VPN over cellular.
func (m *Windows10GeneralConfiguration) SetCellularBlockVpn(value *bool)() {
    m.cellularBlockVpn = value
}
// SetCellularBlockVpnWhenRoaming sets the cellularBlockVpnWhenRoaming property value. Whether or not to Block the user from using VPN when roaming over cellular.
func (m *Windows10GeneralConfiguration) SetCellularBlockVpnWhenRoaming(value *bool)() {
    m.cellularBlockVpnWhenRoaming = value
}
// SetCertificatesBlockManualRootCertificateInstallation sets the certificatesBlockManualRootCertificateInstallation property value. Whether or not to Block the user from doing manual root certificate installation.
func (m *Windows10GeneralConfiguration) SetCertificatesBlockManualRootCertificateInstallation(value *bool)() {
    m.certificatesBlockManualRootCertificateInstallation = value
}
// SetConnectedDevicesServiceBlocked sets the connectedDevicesServiceBlocked property value. Whether or not to block Connected Devices Service which enables discovery and connection to other devices, remote messaging, remote app sessions and other cross-device experiences.
func (m *Windows10GeneralConfiguration) SetConnectedDevicesServiceBlocked(value *bool)() {
    m.connectedDevicesServiceBlocked = value
}
// SetCopyPasteBlocked sets the copyPasteBlocked property value. Whether or not to Block the user from using copy paste.
func (m *Windows10GeneralConfiguration) SetCopyPasteBlocked(value *bool)() {
    m.copyPasteBlocked = value
}
// SetCortanaBlocked sets the cortanaBlocked property value. Whether or not to Block the user from using Cortana.
func (m *Windows10GeneralConfiguration) SetCortanaBlocked(value *bool)() {
    m.cortanaBlocked = value
}
// SetDefenderBlockEndUserAccess sets the defenderBlockEndUserAccess property value. Whether or not to block end user access to Defender.
func (m *Windows10GeneralConfiguration) SetDefenderBlockEndUserAccess(value *bool)() {
    m.defenderBlockEndUserAccess = value
}
// SetDefenderCloudBlockLevel sets the defenderCloudBlockLevel property value. Possible values of Cloud Block Level
func (m *Windows10GeneralConfiguration) SetDefenderCloudBlockLevel(value *DefenderCloudBlockLevelType)() {
    m.defenderCloudBlockLevel = value
}
// SetDefenderDaysBeforeDeletingQuarantinedMalware sets the defenderDaysBeforeDeletingQuarantinedMalware property value. Number of days before deleting quarantined malware. Valid values 0 to 90
func (m *Windows10GeneralConfiguration) SetDefenderDaysBeforeDeletingQuarantinedMalware(value *int32)() {
    m.defenderDaysBeforeDeletingQuarantinedMalware = value
}
// SetDefenderDetectedMalwareActions sets the defenderDetectedMalwareActions property value. Gets or sets Defender’s actions to take on detected Malware per threat level.
func (m *Windows10GeneralConfiguration) SetDefenderDetectedMalwareActions(value DefenderDetectedMalwareActionsable)() {
    m.defenderDetectedMalwareActions = value
}
// SetDefenderFileExtensionsToExclude sets the defenderFileExtensionsToExclude property value. File extensions to exclude from scans and real time protection.
func (m *Windows10GeneralConfiguration) SetDefenderFileExtensionsToExclude(value []string)() {
    m.defenderFileExtensionsToExclude = value
}
// SetDefenderFilesAndFoldersToExclude sets the defenderFilesAndFoldersToExclude property value. Files and folder to exclude from scans and real time protection.
func (m *Windows10GeneralConfiguration) SetDefenderFilesAndFoldersToExclude(value []string)() {
    m.defenderFilesAndFoldersToExclude = value
}
// SetDefenderMonitorFileActivity sets the defenderMonitorFileActivity property value. Possible values for monitoring file activity.
func (m *Windows10GeneralConfiguration) SetDefenderMonitorFileActivity(value *DefenderMonitorFileActivity)() {
    m.defenderMonitorFileActivity = value
}
// SetDefenderProcessesToExclude sets the defenderProcessesToExclude property value. Processes to exclude from scans and real time protection.
func (m *Windows10GeneralConfiguration) SetDefenderProcessesToExclude(value []string)() {
    m.defenderProcessesToExclude = value
}
// SetDefenderPromptForSampleSubmission sets the defenderPromptForSampleSubmission property value. Possible values for prompting user for samples submission.
func (m *Windows10GeneralConfiguration) SetDefenderPromptForSampleSubmission(value *DefenderPromptForSampleSubmission)() {
    m.defenderPromptForSampleSubmission = value
}
// SetDefenderRequireBehaviorMonitoring sets the defenderRequireBehaviorMonitoring property value. Indicates whether or not to require behavior monitoring.
func (m *Windows10GeneralConfiguration) SetDefenderRequireBehaviorMonitoring(value *bool)() {
    m.defenderRequireBehaviorMonitoring = value
}
// SetDefenderRequireCloudProtection sets the defenderRequireCloudProtection property value. Indicates whether or not to require cloud protection.
func (m *Windows10GeneralConfiguration) SetDefenderRequireCloudProtection(value *bool)() {
    m.defenderRequireCloudProtection = value
}
// SetDefenderRequireNetworkInspectionSystem sets the defenderRequireNetworkInspectionSystem property value. Indicates whether or not to require network inspection system.
func (m *Windows10GeneralConfiguration) SetDefenderRequireNetworkInspectionSystem(value *bool)() {
    m.defenderRequireNetworkInspectionSystem = value
}
// SetDefenderRequireRealTimeMonitoring sets the defenderRequireRealTimeMonitoring property value. Indicates whether or not to require real time monitoring.
func (m *Windows10GeneralConfiguration) SetDefenderRequireRealTimeMonitoring(value *bool)() {
    m.defenderRequireRealTimeMonitoring = value
}
// SetDefenderScanArchiveFiles sets the defenderScanArchiveFiles property value. Indicates whether or not to scan archive files.
func (m *Windows10GeneralConfiguration) SetDefenderScanArchiveFiles(value *bool)() {
    m.defenderScanArchiveFiles = value
}
// SetDefenderScanDownloads sets the defenderScanDownloads property value. Indicates whether or not to scan downloads.
func (m *Windows10GeneralConfiguration) SetDefenderScanDownloads(value *bool)() {
    m.defenderScanDownloads = value
}
// SetDefenderScanIncomingMail sets the defenderScanIncomingMail property value. Indicates whether or not to scan incoming mail messages.
func (m *Windows10GeneralConfiguration) SetDefenderScanIncomingMail(value *bool)() {
    m.defenderScanIncomingMail = value
}
// SetDefenderScanMappedNetworkDrivesDuringFullScan sets the defenderScanMappedNetworkDrivesDuringFullScan property value. Indicates whether or not to scan mapped network drives during full scan.
func (m *Windows10GeneralConfiguration) SetDefenderScanMappedNetworkDrivesDuringFullScan(value *bool)() {
    m.defenderScanMappedNetworkDrivesDuringFullScan = value
}
// SetDefenderScanMaxCpu sets the defenderScanMaxCpu property value. Max CPU usage percentage during scan. Valid values 0 to 100
func (m *Windows10GeneralConfiguration) SetDefenderScanMaxCpu(value *int32)() {
    m.defenderScanMaxCpu = value
}
// SetDefenderScanNetworkFiles sets the defenderScanNetworkFiles property value. Indicates whether or not to scan files opened from a network folder.
func (m *Windows10GeneralConfiguration) SetDefenderScanNetworkFiles(value *bool)() {
    m.defenderScanNetworkFiles = value
}
// SetDefenderScanRemovableDrivesDuringFullScan sets the defenderScanRemovableDrivesDuringFullScan property value. Indicates whether or not to scan removable drives during full scan.
func (m *Windows10GeneralConfiguration) SetDefenderScanRemovableDrivesDuringFullScan(value *bool)() {
    m.defenderScanRemovableDrivesDuringFullScan = value
}
// SetDefenderScanScriptsLoadedInInternetExplorer sets the defenderScanScriptsLoadedInInternetExplorer property value. Indicates whether or not to scan scripts loaded in Internet Explorer browser.
func (m *Windows10GeneralConfiguration) SetDefenderScanScriptsLoadedInInternetExplorer(value *bool)() {
    m.defenderScanScriptsLoadedInInternetExplorer = value
}
// SetDefenderScanType sets the defenderScanType property value. Possible values for system scan type.
func (m *Windows10GeneralConfiguration) SetDefenderScanType(value *DefenderScanType)() {
    m.defenderScanType = value
}
// SetDefenderScheduledQuickScanTime sets the defenderScheduledQuickScanTime property value. The time to perform a daily quick scan.
func (m *Windows10GeneralConfiguration) SetDefenderScheduledQuickScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.defenderScheduledQuickScanTime = value
}
// SetDefenderScheduledScanTime sets the defenderScheduledScanTime property value. The defender time for the system scan.
func (m *Windows10GeneralConfiguration) SetDefenderScheduledScanTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.defenderScheduledScanTime = value
}
// SetDefenderSignatureUpdateIntervalInHours sets the defenderSignatureUpdateIntervalInHours property value. The signature update interval in hours. Specify 0 not to check. Valid values 0 to 24
func (m *Windows10GeneralConfiguration) SetDefenderSignatureUpdateIntervalInHours(value *int32)() {
    m.defenderSignatureUpdateIntervalInHours = value
}
// SetDefenderSystemScanSchedule sets the defenderSystemScanSchedule property value. Possible values for a weekly schedule.
func (m *Windows10GeneralConfiguration) SetDefenderSystemScanSchedule(value *WeeklySchedule)() {
    m.defenderSystemScanSchedule = value
}
// SetDeveloperUnlockSetting sets the developerUnlockSetting property value. State Management Setting.
func (m *Windows10GeneralConfiguration) SetDeveloperUnlockSetting(value *StateManagementSetting)() {
    m.developerUnlockSetting = value
}
// SetDeviceManagementBlockFactoryResetOnMobile sets the deviceManagementBlockFactoryResetOnMobile property value. Indicates whether or not to Block the user from resetting their phone.
func (m *Windows10GeneralConfiguration) SetDeviceManagementBlockFactoryResetOnMobile(value *bool)() {
    m.deviceManagementBlockFactoryResetOnMobile = value
}
// SetDeviceManagementBlockManualUnenroll sets the deviceManagementBlockManualUnenroll property value. Indicates whether or not to Block the user from doing manual un-enrollment from device management.
func (m *Windows10GeneralConfiguration) SetDeviceManagementBlockManualUnenroll(value *bool)() {
    m.deviceManagementBlockManualUnenroll = value
}
// SetDiagnosticsDataSubmissionMode sets the diagnosticsDataSubmissionMode property value. Allow the device to send diagnostic and usage telemetry data, such as Watson.
func (m *Windows10GeneralConfiguration) SetDiagnosticsDataSubmissionMode(value *DiagnosticDataSubmissionMode)() {
    m.diagnosticsDataSubmissionMode = value
}
// SetEdgeAllowStartPagesModification sets the edgeAllowStartPagesModification property value. Allow users to change Start pages on Edge. Use the EdgeHomepageUrls to specify the Start pages that the user would see by default when they open Edge.
func (m *Windows10GeneralConfiguration) SetEdgeAllowStartPagesModification(value *bool)() {
    m.edgeAllowStartPagesModification = value
}
// SetEdgeBlockAccessToAboutFlags sets the edgeBlockAccessToAboutFlags property value. Indicates whether or not to prevent access to about flags on Edge browser.
func (m *Windows10GeneralConfiguration) SetEdgeBlockAccessToAboutFlags(value *bool)() {
    m.edgeBlockAccessToAboutFlags = value
}
// SetEdgeBlockAddressBarDropdown sets the edgeBlockAddressBarDropdown property value. Block the address bar dropdown functionality in Microsoft Edge. Disable this settings to minimize network connections from Microsoft Edge to Microsoft services.
func (m *Windows10GeneralConfiguration) SetEdgeBlockAddressBarDropdown(value *bool)() {
    m.edgeBlockAddressBarDropdown = value
}
// SetEdgeBlockAutofill sets the edgeBlockAutofill property value. Indicates whether or not to block auto fill.
func (m *Windows10GeneralConfiguration) SetEdgeBlockAutofill(value *bool)() {
    m.edgeBlockAutofill = value
}
// SetEdgeBlockCompatibilityList sets the edgeBlockCompatibilityList property value. Block Microsoft compatibility list in Microsoft Edge. This list from Microsoft helps Edge properly display sites with known compatibility issues.
func (m *Windows10GeneralConfiguration) SetEdgeBlockCompatibilityList(value *bool)() {
    m.edgeBlockCompatibilityList = value
}
// SetEdgeBlockDeveloperTools sets the edgeBlockDeveloperTools property value. Indicates whether or not to block developer tools in the Edge browser.
func (m *Windows10GeneralConfiguration) SetEdgeBlockDeveloperTools(value *bool)() {
    m.edgeBlockDeveloperTools = value
}
// SetEdgeBlocked sets the edgeBlocked property value. Indicates whether or not to Block the user from using the Edge browser.
func (m *Windows10GeneralConfiguration) SetEdgeBlocked(value *bool)() {
    m.edgeBlocked = value
}
// SetEdgeBlockExtensions sets the edgeBlockExtensions property value. Indicates whether or not to block extensions in the Edge browser.
func (m *Windows10GeneralConfiguration) SetEdgeBlockExtensions(value *bool)() {
    m.edgeBlockExtensions = value
}
// SetEdgeBlockInPrivateBrowsing sets the edgeBlockInPrivateBrowsing property value. Indicates whether or not to block InPrivate browsing on corporate networks, in the Edge browser.
func (m *Windows10GeneralConfiguration) SetEdgeBlockInPrivateBrowsing(value *bool)() {
    m.edgeBlockInPrivateBrowsing = value
}
// SetEdgeBlockJavaScript sets the edgeBlockJavaScript property value. Indicates whether or not to Block the user from using JavaScript.
func (m *Windows10GeneralConfiguration) SetEdgeBlockJavaScript(value *bool)() {
    m.edgeBlockJavaScript = value
}
// SetEdgeBlockLiveTileDataCollection sets the edgeBlockLiveTileDataCollection property value. Block the collection of information by Microsoft for live tile creation when users pin a site to Start from Microsoft Edge.
func (m *Windows10GeneralConfiguration) SetEdgeBlockLiveTileDataCollection(value *bool)() {
    m.edgeBlockLiveTileDataCollection = value
}
// SetEdgeBlockPasswordManager sets the edgeBlockPasswordManager property value. Indicates whether or not to Block password manager.
func (m *Windows10GeneralConfiguration) SetEdgeBlockPasswordManager(value *bool)() {
    m.edgeBlockPasswordManager = value
}
// SetEdgeBlockPopups sets the edgeBlockPopups property value. Indicates whether or not to block popups.
func (m *Windows10GeneralConfiguration) SetEdgeBlockPopups(value *bool)() {
    m.edgeBlockPopups = value
}
// SetEdgeBlockSearchSuggestions sets the edgeBlockSearchSuggestions property value. Indicates whether or not to block the user from using the search suggestions in the address bar.
func (m *Windows10GeneralConfiguration) SetEdgeBlockSearchSuggestions(value *bool)() {
    m.edgeBlockSearchSuggestions = value
}
// SetEdgeBlockSendingDoNotTrackHeader sets the edgeBlockSendingDoNotTrackHeader property value. Indicates whether or not to Block the user from sending the do not track header.
func (m *Windows10GeneralConfiguration) SetEdgeBlockSendingDoNotTrackHeader(value *bool)() {
    m.edgeBlockSendingDoNotTrackHeader = value
}
// SetEdgeBlockSendingIntranetTrafficToInternetExplorer sets the edgeBlockSendingIntranetTrafficToInternetExplorer property value. Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer. Note: the name of this property is misleading; the property is obsolete, use EdgeSendIntranetTrafficToInternetExplorer instead.
func (m *Windows10GeneralConfiguration) SetEdgeBlockSendingIntranetTrafficToInternetExplorer(value *bool)() {
    m.edgeBlockSendingIntranetTrafficToInternetExplorer = value
}
// SetEdgeClearBrowsingDataOnExit sets the edgeClearBrowsingDataOnExit property value. Clear browsing data on exiting Microsoft Edge.
func (m *Windows10GeneralConfiguration) SetEdgeClearBrowsingDataOnExit(value *bool)() {
    m.edgeClearBrowsingDataOnExit = value
}
// SetEdgeCookiePolicy sets the edgeCookiePolicy property value. Possible values to specify which cookies are allowed in Microsoft Edge.
func (m *Windows10GeneralConfiguration) SetEdgeCookiePolicy(value *EdgeCookiePolicy)() {
    m.edgeCookiePolicy = value
}
// SetEdgeDisableFirstRunPage sets the edgeDisableFirstRunPage property value. Block the Microsoft web page that opens on the first use of Microsoft Edge. This policy allows enterprises, like those enrolled in zero emissions configurations, to block this page.
func (m *Windows10GeneralConfiguration) SetEdgeDisableFirstRunPage(value *bool)() {
    m.edgeDisableFirstRunPage = value
}
// SetEdgeEnterpriseModeSiteListLocation sets the edgeEnterpriseModeSiteListLocation property value. Indicates the enterprise mode site list location. Could be a local file, local network or http location.
func (m *Windows10GeneralConfiguration) SetEdgeEnterpriseModeSiteListLocation(value *string)() {
    m.edgeEnterpriseModeSiteListLocation = value
}
// SetEdgeFirstRunUrl sets the edgeFirstRunUrl property value. The first run URL for when Edge browser is opened for the first time.
func (m *Windows10GeneralConfiguration) SetEdgeFirstRunUrl(value *string)() {
    m.edgeFirstRunUrl = value
}
// SetEdgeHomepageUrls sets the edgeHomepageUrls property value. The list of URLs for homepages shodwn on MDM-enrolled devices on Edge browser.
func (m *Windows10GeneralConfiguration) SetEdgeHomepageUrls(value []string)() {
    m.edgeHomepageUrls = value
}
// SetEdgeRequireSmartScreen sets the edgeRequireSmartScreen property value. Indicates whether or not to Require the user to use the smart screen filter.
func (m *Windows10GeneralConfiguration) SetEdgeRequireSmartScreen(value *bool)() {
    m.edgeRequireSmartScreen = value
}
// SetEdgeSearchEngine sets the edgeSearchEngine property value. Allows IT admins to set a default search engine for MDM-Controlled devices. Users can override this and change their default search engine provided the AllowSearchEngineCustomization policy is not set.
func (m *Windows10GeneralConfiguration) SetEdgeSearchEngine(value EdgeSearchEngineBaseable)() {
    m.edgeSearchEngine = value
}
// SetEdgeSendIntranetTrafficToInternetExplorer sets the edgeSendIntranetTrafficToInternetExplorer property value. Indicates whether or not to switch the intranet traffic from Edge to Internet Explorer.
func (m *Windows10GeneralConfiguration) SetEdgeSendIntranetTrafficToInternetExplorer(value *bool)() {
    m.edgeSendIntranetTrafficToInternetExplorer = value
}
// SetEdgeSyncFavoritesWithInternetExplorer sets the edgeSyncFavoritesWithInternetExplorer property value. Enable favorites sync between Internet Explorer and Microsoft Edge. Additions, deletions, modifications and order changes to favorites are shared between browsers.
func (m *Windows10GeneralConfiguration) SetEdgeSyncFavoritesWithInternetExplorer(value *bool)() {
    m.edgeSyncFavoritesWithInternetExplorer = value
}
// SetEnterpriseCloudPrintDiscoveryEndPoint sets the enterpriseCloudPrintDiscoveryEndPoint property value. Endpoint for discovering cloud printers.
func (m *Windows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryEndPoint(value *string)() {
    m.enterpriseCloudPrintDiscoveryEndPoint = value
}
// SetEnterpriseCloudPrintDiscoveryMaxLimit sets the enterpriseCloudPrintDiscoveryMaxLimit property value. Maximum number of printers that should be queried from a discovery endpoint. This is a mobile only setting. Valid values 1 to 65535
func (m *Windows10GeneralConfiguration) SetEnterpriseCloudPrintDiscoveryMaxLimit(value *int32)() {
    m.enterpriseCloudPrintDiscoveryMaxLimit = value
}
// SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier sets the enterpriseCloudPrintMopriaDiscoveryResourceIdentifier property value. OAuth resource URI for printer discovery service as configured in Azure portal.
func (m *Windows10GeneralConfiguration) SetEnterpriseCloudPrintMopriaDiscoveryResourceIdentifier(value *string)() {
    m.enterpriseCloudPrintMopriaDiscoveryResourceIdentifier = value
}
// SetEnterpriseCloudPrintOAuthAuthority sets the enterpriseCloudPrintOAuthAuthority property value. Authentication endpoint for acquiring OAuth tokens.
func (m *Windows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthAuthority(value *string)() {
    m.enterpriseCloudPrintOAuthAuthority = value
}
// SetEnterpriseCloudPrintOAuthClientIdentifier sets the enterpriseCloudPrintOAuthClientIdentifier property value. GUID of a client application authorized to retrieve OAuth tokens from the OAuth Authority.
func (m *Windows10GeneralConfiguration) SetEnterpriseCloudPrintOAuthClientIdentifier(value *string)() {
    m.enterpriseCloudPrintOAuthClientIdentifier = value
}
// SetEnterpriseCloudPrintResourceIdentifier sets the enterpriseCloudPrintResourceIdentifier property value. OAuth resource URI for print service as configured in the Azure portal.
func (m *Windows10GeneralConfiguration) SetEnterpriseCloudPrintResourceIdentifier(value *string)() {
    m.enterpriseCloudPrintResourceIdentifier = value
}
// SetExperienceBlockDeviceDiscovery sets the experienceBlockDeviceDiscovery property value. Indicates whether or not to enable device discovery UX.
func (m *Windows10GeneralConfiguration) SetExperienceBlockDeviceDiscovery(value *bool)() {
    m.experienceBlockDeviceDiscovery = value
}
// SetExperienceBlockErrorDialogWhenNoSIM sets the experienceBlockErrorDialogWhenNoSIM property value. Indicates whether or not to allow the error dialog from displaying if no SIM card is detected.
func (m *Windows10GeneralConfiguration) SetExperienceBlockErrorDialogWhenNoSIM(value *bool)() {
    m.experienceBlockErrorDialogWhenNoSIM = value
}
// SetExperienceBlockTaskSwitcher sets the experienceBlockTaskSwitcher property value. Indicates whether or not to enable task switching on the device.
func (m *Windows10GeneralConfiguration) SetExperienceBlockTaskSwitcher(value *bool)() {
    m.experienceBlockTaskSwitcher = value
}
// SetGameDvrBlocked sets the gameDvrBlocked property value. Indicates whether or not to block DVR and broadcasting.
func (m *Windows10GeneralConfiguration) SetGameDvrBlocked(value *bool)() {
    m.gameDvrBlocked = value
}
// SetInternetSharingBlocked sets the internetSharingBlocked property value. Indicates whether or not to Block the user from using internet sharing.
func (m *Windows10GeneralConfiguration) SetInternetSharingBlocked(value *bool)() {
    m.internetSharingBlocked = value
}
// SetLocationServicesBlocked sets the locationServicesBlocked property value. Indicates whether or not to Block the user from location services.
func (m *Windows10GeneralConfiguration) SetLocationServicesBlocked(value *bool)() {
    m.locationServicesBlocked = value
}
// SetLockScreenAllowTimeoutConfiguration sets the lockScreenAllowTimeoutConfiguration property value. Specify whether to show a user-configurable setting to control the screen timeout while on the lock screen of Windows 10 Mobile devices. If this policy is set to Allow, the value set by lockScreenTimeoutInSeconds is ignored.
func (m *Windows10GeneralConfiguration) SetLockScreenAllowTimeoutConfiguration(value *bool)() {
    m.lockScreenAllowTimeoutConfiguration = value
}
// SetLockScreenBlockActionCenterNotifications sets the lockScreenBlockActionCenterNotifications property value. Indicates whether or not to block action center notifications over lock screen.
func (m *Windows10GeneralConfiguration) SetLockScreenBlockActionCenterNotifications(value *bool)() {
    m.lockScreenBlockActionCenterNotifications = value
}
// SetLockScreenBlockCortana sets the lockScreenBlockCortana property value. Indicates whether or not the user can interact with Cortana using speech while the system is locked.
func (m *Windows10GeneralConfiguration) SetLockScreenBlockCortana(value *bool)() {
    m.lockScreenBlockCortana = value
}
// SetLockScreenBlockToastNotifications sets the lockScreenBlockToastNotifications property value. Indicates whether to allow toast notifications above the device lock screen.
func (m *Windows10GeneralConfiguration) SetLockScreenBlockToastNotifications(value *bool)() {
    m.lockScreenBlockToastNotifications = value
}
// SetLockScreenTimeoutInSeconds sets the lockScreenTimeoutInSeconds property value. Set the duration (in seconds) from the screen locking to the screen turning off for Windows 10 Mobile devices. Supported values are 11-1800. Valid values 11 to 1800
func (m *Windows10GeneralConfiguration) SetLockScreenTimeoutInSeconds(value *int32)() {
    m.lockScreenTimeoutInSeconds = value
}
// SetLogonBlockFastUserSwitching sets the logonBlockFastUserSwitching property value. Disables the ability to quickly switch between users that are logged on simultaneously without logging off.
func (m *Windows10GeneralConfiguration) SetLogonBlockFastUserSwitching(value *bool)() {
    m.logonBlockFastUserSwitching = value
}
// SetMicrosoftAccountBlocked sets the microsoftAccountBlocked property value. Indicates whether or not to Block a Microsoft account.
func (m *Windows10GeneralConfiguration) SetMicrosoftAccountBlocked(value *bool)() {
    m.microsoftAccountBlocked = value
}
// SetMicrosoftAccountBlockSettingsSync sets the microsoftAccountBlockSettingsSync property value. Indicates whether or not to Block Microsoft account settings sync.
func (m *Windows10GeneralConfiguration) SetMicrosoftAccountBlockSettingsSync(value *bool)() {
    m.microsoftAccountBlockSettingsSync = value
}
// SetNetworkProxyApplySettingsDeviceWide sets the networkProxyApplySettingsDeviceWide property value. If set, proxy settings will be applied to all processes and accounts in the device. Otherwise, it will be applied to the user account that’s enrolled into MDM.
func (m *Windows10GeneralConfiguration) SetNetworkProxyApplySettingsDeviceWide(value *bool)() {
    m.networkProxyApplySettingsDeviceWide = value
}
// SetNetworkProxyAutomaticConfigurationUrl sets the networkProxyAutomaticConfigurationUrl property value. Address to the proxy auto-config (PAC) script you want to use.
func (m *Windows10GeneralConfiguration) SetNetworkProxyAutomaticConfigurationUrl(value *string)() {
    m.networkProxyAutomaticConfigurationUrl = value
}
// SetNetworkProxyDisableAutoDetect sets the networkProxyDisableAutoDetect property value. Disable automatic detection of settings. If enabled, the system will try to find the path to a proxy auto-config (PAC) script.
func (m *Windows10GeneralConfiguration) SetNetworkProxyDisableAutoDetect(value *bool)() {
    m.networkProxyDisableAutoDetect = value
}
// SetNetworkProxyServer sets the networkProxyServer property value. Specifies manual proxy server settings.
func (m *Windows10GeneralConfiguration) SetNetworkProxyServer(value Windows10NetworkProxyServerable)() {
    m.networkProxyServer = value
}
// SetNfcBlocked sets the nfcBlocked property value. Indicates whether or not to Block the user from using near field communication.
func (m *Windows10GeneralConfiguration) SetNfcBlocked(value *bool)() {
    m.nfcBlocked = value
}
// SetOneDriveDisableFileSync sets the oneDriveDisableFileSync property value. Gets or sets a value allowing IT admins to prevent apps and features from working with files on OneDrive.
func (m *Windows10GeneralConfiguration) SetOneDriveDisableFileSync(value *bool)() {
    m.oneDriveDisableFileSync = value
}
// SetPasswordBlockSimple sets the passwordBlockSimple property value. Specify whether PINs or passwords such as '1111' or '1234' are allowed. For Windows 10 desktops, it also controls the use of picture passwords.
func (m *Windows10GeneralConfiguration) SetPasswordBlockSimple(value *bool)() {
    m.passwordBlockSimple = value
}
// SetPasswordExpirationDays sets the passwordExpirationDays property value. The password expiration in days. Valid values 0 to 730
func (m *Windows10GeneralConfiguration) SetPasswordExpirationDays(value *int32)() {
    m.passwordExpirationDays = value
}
// SetPasswordMinimumCharacterSetCount sets the passwordMinimumCharacterSetCount property value. The number of character sets required in the password.
func (m *Windows10GeneralConfiguration) SetPasswordMinimumCharacterSetCount(value *int32)() {
    m.passwordMinimumCharacterSetCount = value
}
// SetPasswordMinimumLength sets the passwordMinimumLength property value. The minimum password length. Valid values 4 to 16
func (m *Windows10GeneralConfiguration) SetPasswordMinimumLength(value *int32)() {
    m.passwordMinimumLength = value
}
// SetPasswordMinutesOfInactivityBeforeScreenTimeout sets the passwordMinutesOfInactivityBeforeScreenTimeout property value. The minutes of inactivity before the screen times out.
func (m *Windows10GeneralConfiguration) SetPasswordMinutesOfInactivityBeforeScreenTimeout(value *int32)() {
    m.passwordMinutesOfInactivityBeforeScreenTimeout = value
}
// SetPasswordPreviousPasswordBlockCount sets the passwordPreviousPasswordBlockCount property value. The number of previous passwords to prevent reuse of. Valid values 0 to 50
func (m *Windows10GeneralConfiguration) SetPasswordPreviousPasswordBlockCount(value *int32)() {
    m.passwordPreviousPasswordBlockCount = value
}
// SetPasswordRequired sets the passwordRequired property value. Indicates whether or not to require the user to have a password.
func (m *Windows10GeneralConfiguration) SetPasswordRequired(value *bool)() {
    m.passwordRequired = value
}
// SetPasswordRequiredType sets the passwordRequiredType property value. Possible values of required passwords.
func (m *Windows10GeneralConfiguration) SetPasswordRequiredType(value *RequiredPasswordType)() {
    m.passwordRequiredType = value
}
// SetPasswordRequireWhenResumeFromIdleState sets the passwordRequireWhenResumeFromIdleState property value. Indicates whether or not to require a password upon resuming from an idle state.
func (m *Windows10GeneralConfiguration) SetPasswordRequireWhenResumeFromIdleState(value *bool)() {
    m.passwordRequireWhenResumeFromIdleState = value
}
// SetPasswordSignInFailureCountBeforeFactoryReset sets the passwordSignInFailureCountBeforeFactoryReset property value. The number of sign in failures before factory reset. Valid values 0 to 999
func (m *Windows10GeneralConfiguration) SetPasswordSignInFailureCountBeforeFactoryReset(value *int32)() {
    m.passwordSignInFailureCountBeforeFactoryReset = value
}
// SetPersonalizationDesktopImageUrl sets the personalizationDesktopImageUrl property value. A http or https Url to a jpg, jpeg or png image that needs to be downloaded and used as the Desktop Image or a file Url to a local image on the file system that needs to used as the Desktop Image.
func (m *Windows10GeneralConfiguration) SetPersonalizationDesktopImageUrl(value *string)() {
    m.personalizationDesktopImageUrl = value
}
// SetPersonalizationLockScreenImageUrl sets the personalizationLockScreenImageUrl property value. A http or https Url to a jpg, jpeg or png image that neeeds to be downloaded and used as the Lock Screen Image or a file Url to a local image on the file system that needs to be used as the Lock Screen Image.
func (m *Windows10GeneralConfiguration) SetPersonalizationLockScreenImageUrl(value *string)() {
    m.personalizationLockScreenImageUrl = value
}
// SetPrivacyAdvertisingId sets the privacyAdvertisingId property value. State Management Setting.
func (m *Windows10GeneralConfiguration) SetPrivacyAdvertisingId(value *StateManagementSetting)() {
    m.privacyAdvertisingId = value
}
// SetPrivacyAutoAcceptPairingAndConsentPrompts sets the privacyAutoAcceptPairingAndConsentPrompts property value. Indicates whether or not to allow the automatic acceptance of the pairing and privacy user consent dialog when launching apps.
func (m *Windows10GeneralConfiguration) SetPrivacyAutoAcceptPairingAndConsentPrompts(value *bool)() {
    m.privacyAutoAcceptPairingAndConsentPrompts = value
}
// SetPrivacyBlockInputPersonalization sets the privacyBlockInputPersonalization property value. Indicates whether or not to block the usage of cloud based speech services for Cortana, Dictation, or Store applications.
func (m *Windows10GeneralConfiguration) SetPrivacyBlockInputPersonalization(value *bool)() {
    m.privacyBlockInputPersonalization = value
}
// SetResetProtectionModeBlocked sets the resetProtectionModeBlocked property value. Indicates whether or not to Block the user from reset protection mode.
func (m *Windows10GeneralConfiguration) SetResetProtectionModeBlocked(value *bool)() {
    m.resetProtectionModeBlocked = value
}
// SetSafeSearchFilter sets the safeSearchFilter property value. Specifies what level of safe search (filtering adult content) is required
func (m *Windows10GeneralConfiguration) SetSafeSearchFilter(value *SafeSearchFilterType)() {
    m.safeSearchFilter = value
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether or not to Block the user from taking Screenshots.
func (m *Windows10GeneralConfiguration) SetScreenCaptureBlocked(value *bool)() {
    m.screenCaptureBlocked = value
}
// SetSearchBlockDiacritics sets the searchBlockDiacritics property value. Specifies if search can use diacritics.
func (m *Windows10GeneralConfiguration) SetSearchBlockDiacritics(value *bool)() {
    m.searchBlockDiacritics = value
}
// SetSearchDisableAutoLanguageDetection sets the searchDisableAutoLanguageDetection property value. Specifies whether to use automatic language detection when indexing content and properties.
func (m *Windows10GeneralConfiguration) SetSearchDisableAutoLanguageDetection(value *bool)() {
    m.searchDisableAutoLanguageDetection = value
}
// SetSearchDisableIndexerBackoff sets the searchDisableIndexerBackoff property value. Indicates whether or not to disable the search indexer backoff feature.
func (m *Windows10GeneralConfiguration) SetSearchDisableIndexerBackoff(value *bool)() {
    m.searchDisableIndexerBackoff = value
}
// SetSearchDisableIndexingEncryptedItems sets the searchDisableIndexingEncryptedItems property value. Indicates whether or not to block indexing of WIP-protected items to prevent them from appearing in search results for Cortana or Explorer.
func (m *Windows10GeneralConfiguration) SetSearchDisableIndexingEncryptedItems(value *bool)() {
    m.searchDisableIndexingEncryptedItems = value
}
// SetSearchDisableIndexingRemovableDrive sets the searchDisableIndexingRemovableDrive property value. Indicates whether or not to allow users to add locations on removable drives to libraries and to be indexed.
func (m *Windows10GeneralConfiguration) SetSearchDisableIndexingRemovableDrive(value *bool)() {
    m.searchDisableIndexingRemovableDrive = value
}
// SetSearchEnableAutomaticIndexSizeManangement sets the searchEnableAutomaticIndexSizeManangement property value. Specifies minimum amount of hard drive space on the same drive as the index location before indexing stops.
func (m *Windows10GeneralConfiguration) SetSearchEnableAutomaticIndexSizeManangement(value *bool)() {
    m.searchEnableAutomaticIndexSizeManangement = value
}
// SetSearchEnableRemoteQueries sets the searchEnableRemoteQueries property value. Indicates whether or not to block remote queries of this computer’s index.
func (m *Windows10GeneralConfiguration) SetSearchEnableRemoteQueries(value *bool)() {
    m.searchEnableRemoteQueries = value
}
// SetSettingsBlockAccountsPage sets the settingsBlockAccountsPage property value. Indicates whether or not to block access to Accounts in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockAccountsPage(value *bool)() {
    m.settingsBlockAccountsPage = value
}
// SetSettingsBlockAddProvisioningPackage sets the settingsBlockAddProvisioningPackage property value. Indicates whether or not to block the user from installing provisioning packages.
func (m *Windows10GeneralConfiguration) SetSettingsBlockAddProvisioningPackage(value *bool)() {
    m.settingsBlockAddProvisioningPackage = value
}
// SetSettingsBlockAppsPage sets the settingsBlockAppsPage property value. Indicates whether or not to block access to Apps in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockAppsPage(value *bool)() {
    m.settingsBlockAppsPage = value
}
// SetSettingsBlockChangeLanguage sets the settingsBlockChangeLanguage property value. Indicates whether or not to block the user from changing the language settings.
func (m *Windows10GeneralConfiguration) SetSettingsBlockChangeLanguage(value *bool)() {
    m.settingsBlockChangeLanguage = value
}
// SetSettingsBlockChangePowerSleep sets the settingsBlockChangePowerSleep property value. Indicates whether or not to block the user from changing power and sleep settings.
func (m *Windows10GeneralConfiguration) SetSettingsBlockChangePowerSleep(value *bool)() {
    m.settingsBlockChangePowerSleep = value
}
// SetSettingsBlockChangeRegion sets the settingsBlockChangeRegion property value. Indicates whether or not to block the user from changing the region settings.
func (m *Windows10GeneralConfiguration) SetSettingsBlockChangeRegion(value *bool)() {
    m.settingsBlockChangeRegion = value
}
// SetSettingsBlockChangeSystemTime sets the settingsBlockChangeSystemTime property value. Indicates whether or not to block the user from changing date and time settings.
func (m *Windows10GeneralConfiguration) SetSettingsBlockChangeSystemTime(value *bool)() {
    m.settingsBlockChangeSystemTime = value
}
// SetSettingsBlockDevicesPage sets the settingsBlockDevicesPage property value. Indicates whether or not to block access to Devices in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockDevicesPage(value *bool)() {
    m.settingsBlockDevicesPage = value
}
// SetSettingsBlockEaseOfAccessPage sets the settingsBlockEaseOfAccessPage property value. Indicates whether or not to block access to Ease of Access in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockEaseOfAccessPage(value *bool)() {
    m.settingsBlockEaseOfAccessPage = value
}
// SetSettingsBlockEditDeviceName sets the settingsBlockEditDeviceName property value. Indicates whether or not to block the user from editing the device name.
func (m *Windows10GeneralConfiguration) SetSettingsBlockEditDeviceName(value *bool)() {
    m.settingsBlockEditDeviceName = value
}
// SetSettingsBlockGamingPage sets the settingsBlockGamingPage property value. Indicates whether or not to block access to Gaming in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockGamingPage(value *bool)() {
    m.settingsBlockGamingPage = value
}
// SetSettingsBlockNetworkInternetPage sets the settingsBlockNetworkInternetPage property value. Indicates whether or not to block access to Network & Internet in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockNetworkInternetPage(value *bool)() {
    m.settingsBlockNetworkInternetPage = value
}
// SetSettingsBlockPersonalizationPage sets the settingsBlockPersonalizationPage property value. Indicates whether or not to block access to Personalization in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockPersonalizationPage(value *bool)() {
    m.settingsBlockPersonalizationPage = value
}
// SetSettingsBlockPrivacyPage sets the settingsBlockPrivacyPage property value. Indicates whether or not to block access to Privacy in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockPrivacyPage(value *bool)() {
    m.settingsBlockPrivacyPage = value
}
// SetSettingsBlockRemoveProvisioningPackage sets the settingsBlockRemoveProvisioningPackage property value. Indicates whether or not to block the runtime configuration agent from removing provisioning packages.
func (m *Windows10GeneralConfiguration) SetSettingsBlockRemoveProvisioningPackage(value *bool)() {
    m.settingsBlockRemoveProvisioningPackage = value
}
// SetSettingsBlockSettingsApp sets the settingsBlockSettingsApp property value. Indicates whether or not to block access to Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockSettingsApp(value *bool)() {
    m.settingsBlockSettingsApp = value
}
// SetSettingsBlockSystemPage sets the settingsBlockSystemPage property value. Indicates whether or not to block access to System in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockSystemPage(value *bool)() {
    m.settingsBlockSystemPage = value
}
// SetSettingsBlockTimeLanguagePage sets the settingsBlockTimeLanguagePage property value. Indicates whether or not to block access to Time & Language in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockTimeLanguagePage(value *bool)() {
    m.settingsBlockTimeLanguagePage = value
}
// SetSettingsBlockUpdateSecurityPage sets the settingsBlockUpdateSecurityPage property value. Indicates whether or not to block access to Update & Security in Settings app.
func (m *Windows10GeneralConfiguration) SetSettingsBlockUpdateSecurityPage(value *bool)() {
    m.settingsBlockUpdateSecurityPage = value
}
// SetSharedUserAppDataAllowed sets the sharedUserAppDataAllowed property value. Indicates whether or not to block multiple users of the same app to share data.
func (m *Windows10GeneralConfiguration) SetSharedUserAppDataAllowed(value *bool)() {
    m.sharedUserAppDataAllowed = value
}
// SetSmartScreenBlockPromptOverride sets the smartScreenBlockPromptOverride property value. Indicates whether or not users can override SmartScreen Filter warnings about potentially malicious websites.
func (m *Windows10GeneralConfiguration) SetSmartScreenBlockPromptOverride(value *bool)() {
    m.smartScreenBlockPromptOverride = value
}
// SetSmartScreenBlockPromptOverrideForFiles sets the smartScreenBlockPromptOverrideForFiles property value. Indicates whether or not users can override the SmartScreen Filter warnings about downloading unverified files
func (m *Windows10GeneralConfiguration) SetSmartScreenBlockPromptOverrideForFiles(value *bool)() {
    m.smartScreenBlockPromptOverrideForFiles = value
}
// SetSmartScreenEnableAppInstallControl sets the smartScreenEnableAppInstallControl property value. This property will be deprecated in July 2019 and will be replaced by property SmartScreenAppInstallControl. Allows IT Admins to control whether users are allowed to install apps from places other than the Store.
func (m *Windows10GeneralConfiguration) SetSmartScreenEnableAppInstallControl(value *bool)() {
    m.smartScreenEnableAppInstallControl = value
}
// SetStartBlockUnpinningAppsFromTaskbar sets the startBlockUnpinningAppsFromTaskbar property value. Indicates whether or not to block the user from unpinning apps from taskbar.
func (m *Windows10GeneralConfiguration) SetStartBlockUnpinningAppsFromTaskbar(value *bool)() {
    m.startBlockUnpinningAppsFromTaskbar = value
}
// SetStartMenuAppListVisibility sets the startMenuAppListVisibility property value. Type of start menu app list visibility.
func (m *Windows10GeneralConfiguration) SetStartMenuAppListVisibility(value *WindowsStartMenuAppListVisibilityType)() {
    m.startMenuAppListVisibility = value
}
// SetStartMenuHideChangeAccountSettings sets the startMenuHideChangeAccountSettings property value. Enabling this policy hides the change account setting from appearing in the user tile in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideChangeAccountSettings(value *bool)() {
    m.startMenuHideChangeAccountSettings = value
}
// SetStartMenuHideFrequentlyUsedApps sets the startMenuHideFrequentlyUsedApps property value. Enabling this policy hides the most used apps from appearing on the start menu and disables the corresponding toggle in the Settings app.
func (m *Windows10GeneralConfiguration) SetStartMenuHideFrequentlyUsedApps(value *bool)() {
    m.startMenuHideFrequentlyUsedApps = value
}
// SetStartMenuHideHibernate sets the startMenuHideHibernate property value. Enabling this policy hides hibernate from appearing in the power button in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideHibernate(value *bool)() {
    m.startMenuHideHibernate = value
}
// SetStartMenuHideLock sets the startMenuHideLock property value. Enabling this policy hides lock from appearing in the user tile in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideLock(value *bool)() {
    m.startMenuHideLock = value
}
// SetStartMenuHidePowerButton sets the startMenuHidePowerButton property value. Enabling this policy hides the power button from appearing in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHidePowerButton(value *bool)() {
    m.startMenuHidePowerButton = value
}
// SetStartMenuHideRecentJumpLists sets the startMenuHideRecentJumpLists property value. Enabling this policy hides recent jump lists from appearing on the start menu/taskbar and disables the corresponding toggle in the Settings app.
func (m *Windows10GeneralConfiguration) SetStartMenuHideRecentJumpLists(value *bool)() {
    m.startMenuHideRecentJumpLists = value
}
// SetStartMenuHideRecentlyAddedApps sets the startMenuHideRecentlyAddedApps property value. Enabling this policy hides recently added apps from appearing on the start menu and disables the corresponding toggle in the Settings app.
func (m *Windows10GeneralConfiguration) SetStartMenuHideRecentlyAddedApps(value *bool)() {
    m.startMenuHideRecentlyAddedApps = value
}
// SetStartMenuHideRestartOptions sets the startMenuHideRestartOptions property value. Enabling this policy hides 'Restart/Update and Restart' from appearing in the power button in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideRestartOptions(value *bool)() {
    m.startMenuHideRestartOptions = value
}
// SetStartMenuHideShutDown sets the startMenuHideShutDown property value. Enabling this policy hides shut down/update and shut down from appearing in the power button in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideShutDown(value *bool)() {
    m.startMenuHideShutDown = value
}
// SetStartMenuHideSignOut sets the startMenuHideSignOut property value. Enabling this policy hides sign out from appearing in the user tile in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideSignOut(value *bool)() {
    m.startMenuHideSignOut = value
}
// SetStartMenuHideSleep sets the startMenuHideSleep property value. Enabling this policy hides sleep from appearing in the power button in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideSleep(value *bool)() {
    m.startMenuHideSleep = value
}
// SetStartMenuHideSwitchAccount sets the startMenuHideSwitchAccount property value. Enabling this policy hides switch account from appearing in the user tile in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideSwitchAccount(value *bool)() {
    m.startMenuHideSwitchAccount = value
}
// SetStartMenuHideUserTile sets the startMenuHideUserTile property value. Enabling this policy hides the user tile from appearing in the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuHideUserTile(value *bool)() {
    m.startMenuHideUserTile = value
}
// SetStartMenuLayoutEdgeAssetsXml sets the startMenuLayoutEdgeAssetsXml property value. This policy setting allows you to import Edge assets to be used with startMenuLayoutXml policy. Start layout can contain secondary tile from Edge app which looks for Edge local asset file. Edge local asset would not exist and cause Edge secondary tile to appear empty in this case. This policy only gets applied when startMenuLayoutXml policy is modified. The value should be a UTF-8 Base64 encoded byte array.
func (m *Windows10GeneralConfiguration) SetStartMenuLayoutEdgeAssetsXml(value []byte)() {
    m.startMenuLayoutEdgeAssetsXml = value
}
// SetStartMenuLayoutXml sets the startMenuLayoutXml property value. Allows admins to override the default Start menu layout and prevents the user from changing it. The layout is modified by specifying an XML file based on a layout modification schema. XML needs to be in a UTF8 encoded byte array format.
func (m *Windows10GeneralConfiguration) SetStartMenuLayoutXml(value []byte)() {
    m.startMenuLayoutXml = value
}
// SetStartMenuMode sets the startMenuMode property value. Type of display modes for the start menu.
func (m *Windows10GeneralConfiguration) SetStartMenuMode(value *WindowsStartMenuModeType)() {
    m.startMenuMode = value
}
// SetStartMenuPinnedFolderDocuments sets the startMenuPinnedFolderDocuments property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderDocuments(value *VisibilitySetting)() {
    m.startMenuPinnedFolderDocuments = value
}
// SetStartMenuPinnedFolderDownloads sets the startMenuPinnedFolderDownloads property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderDownloads(value *VisibilitySetting)() {
    m.startMenuPinnedFolderDownloads = value
}
// SetStartMenuPinnedFolderFileExplorer sets the startMenuPinnedFolderFileExplorer property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderFileExplorer(value *VisibilitySetting)() {
    m.startMenuPinnedFolderFileExplorer = value
}
// SetStartMenuPinnedFolderHomeGroup sets the startMenuPinnedFolderHomeGroup property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderHomeGroup(value *VisibilitySetting)() {
    m.startMenuPinnedFolderHomeGroup = value
}
// SetStartMenuPinnedFolderMusic sets the startMenuPinnedFolderMusic property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderMusic(value *VisibilitySetting)() {
    m.startMenuPinnedFolderMusic = value
}
// SetStartMenuPinnedFolderNetwork sets the startMenuPinnedFolderNetwork property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderNetwork(value *VisibilitySetting)() {
    m.startMenuPinnedFolderNetwork = value
}
// SetStartMenuPinnedFolderPersonalFolder sets the startMenuPinnedFolderPersonalFolder property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderPersonalFolder(value *VisibilitySetting)() {
    m.startMenuPinnedFolderPersonalFolder = value
}
// SetStartMenuPinnedFolderPictures sets the startMenuPinnedFolderPictures property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderPictures(value *VisibilitySetting)() {
    m.startMenuPinnedFolderPictures = value
}
// SetStartMenuPinnedFolderSettings sets the startMenuPinnedFolderSettings property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderSettings(value *VisibilitySetting)() {
    m.startMenuPinnedFolderSettings = value
}
// SetStartMenuPinnedFolderVideos sets the startMenuPinnedFolderVideos property value. Generic visibility state.
func (m *Windows10GeneralConfiguration) SetStartMenuPinnedFolderVideos(value *VisibilitySetting)() {
    m.startMenuPinnedFolderVideos = value
}
// SetStorageBlockRemovableStorage sets the storageBlockRemovableStorage property value. Indicates whether or not to Block the user from using removable storage.
func (m *Windows10GeneralConfiguration) SetStorageBlockRemovableStorage(value *bool)() {
    m.storageBlockRemovableStorage = value
}
// SetStorageRequireMobileDeviceEncryption sets the storageRequireMobileDeviceEncryption property value. Indicating whether or not to require encryption on a mobile device.
func (m *Windows10GeneralConfiguration) SetStorageRequireMobileDeviceEncryption(value *bool)() {
    m.storageRequireMobileDeviceEncryption = value
}
// SetStorageRestrictAppDataToSystemVolume sets the storageRestrictAppDataToSystemVolume property value. Indicates whether application data is restricted to the system drive.
func (m *Windows10GeneralConfiguration) SetStorageRestrictAppDataToSystemVolume(value *bool)() {
    m.storageRestrictAppDataToSystemVolume = value
}
// SetStorageRestrictAppInstallToSystemVolume sets the storageRestrictAppInstallToSystemVolume property value. Indicates whether the installation of applications is restricted to the system drive.
func (m *Windows10GeneralConfiguration) SetStorageRestrictAppInstallToSystemVolume(value *bool)() {
    m.storageRestrictAppInstallToSystemVolume = value
}
// SetTenantLockdownRequireNetworkDuringOutOfBoxExperience sets the tenantLockdownRequireNetworkDuringOutOfBoxExperience property value. Whether the device is required to connect to the network.
func (m *Windows10GeneralConfiguration) SetTenantLockdownRequireNetworkDuringOutOfBoxExperience(value *bool)() {
    m.tenantLockdownRequireNetworkDuringOutOfBoxExperience = value
}
// SetUsbBlocked sets the usbBlocked property value. Indicates whether or not to Block the user from USB connection.
func (m *Windows10GeneralConfiguration) SetUsbBlocked(value *bool)() {
    m.usbBlocked = value
}
// SetVoiceRecordingBlocked sets the voiceRecordingBlocked property value. Indicates whether or not to Block the user from voice recording.
func (m *Windows10GeneralConfiguration) SetVoiceRecordingBlocked(value *bool)() {
    m.voiceRecordingBlocked = value
}
// SetWebRtcBlockLocalhostIpAddress sets the webRtcBlockLocalhostIpAddress property value. Indicates whether or not user's localhost IP address is displayed while making phone calls using the WebRTC
func (m *Windows10GeneralConfiguration) SetWebRtcBlockLocalhostIpAddress(value *bool)() {
    m.webRtcBlockLocalhostIpAddress = value
}
// SetWiFiBlockAutomaticConnectHotspots sets the wiFiBlockAutomaticConnectHotspots property value. Indicating whether or not to block automatically connecting to Wi-Fi hotspots. Has no impact if Wi-Fi is blocked.
func (m *Windows10GeneralConfiguration) SetWiFiBlockAutomaticConnectHotspots(value *bool)() {
    m.wiFiBlockAutomaticConnectHotspots = value
}
// SetWiFiBlocked sets the wiFiBlocked property value. Indicates whether or not to Block the user from using Wi-Fi.
func (m *Windows10GeneralConfiguration) SetWiFiBlocked(value *bool)() {
    m.wiFiBlocked = value
}
// SetWiFiBlockManualConfiguration sets the wiFiBlockManualConfiguration property value. Indicates whether or not to Block the user from using Wi-Fi manual configuration.
func (m *Windows10GeneralConfiguration) SetWiFiBlockManualConfiguration(value *bool)() {
    m.wiFiBlockManualConfiguration = value
}
// SetWiFiScanInterval sets the wiFiScanInterval property value. Specify how often devices scan for Wi-Fi networks. Supported values are 1-500, where 100 = default, and 500 = low frequency. Valid values 1 to 500
func (m *Windows10GeneralConfiguration) SetWiFiScanInterval(value *int32)() {
    m.wiFiScanInterval = value
}
// SetWindowsSpotlightBlockConsumerSpecificFeatures sets the windowsSpotlightBlockConsumerSpecificFeatures property value. Allows IT admins to block experiences that are typically for consumers only, such as Start suggestions, Membership notifications, Post-OOBE app install and redirect tiles.
func (m *Windows10GeneralConfiguration) SetWindowsSpotlightBlockConsumerSpecificFeatures(value *bool)() {
    m.windowsSpotlightBlockConsumerSpecificFeatures = value
}
// SetWindowsSpotlightBlocked sets the windowsSpotlightBlocked property value. Allows IT admins to turn off all Windows Spotlight features
func (m *Windows10GeneralConfiguration) SetWindowsSpotlightBlocked(value *bool)() {
    m.windowsSpotlightBlocked = value
}
// SetWindowsSpotlightBlockOnActionCenter sets the windowsSpotlightBlockOnActionCenter property value. Block suggestions from Microsoft that show after each OS clean install, upgrade or in an on-going basis to introduce users to what is new or changed
func (m *Windows10GeneralConfiguration) SetWindowsSpotlightBlockOnActionCenter(value *bool)() {
    m.windowsSpotlightBlockOnActionCenter = value
}
// SetWindowsSpotlightBlockTailoredExperiences sets the windowsSpotlightBlockTailoredExperiences property value. Block personalized content in Windows spotlight based on user’s device usage.
func (m *Windows10GeneralConfiguration) SetWindowsSpotlightBlockTailoredExperiences(value *bool)() {
    m.windowsSpotlightBlockTailoredExperiences = value
}
// SetWindowsSpotlightBlockThirdPartyNotifications sets the windowsSpotlightBlockThirdPartyNotifications property value. Block third party content delivered via Windows Spotlight
func (m *Windows10GeneralConfiguration) SetWindowsSpotlightBlockThirdPartyNotifications(value *bool)() {
    m.windowsSpotlightBlockThirdPartyNotifications = value
}
// SetWindowsSpotlightBlockWelcomeExperience sets the windowsSpotlightBlockWelcomeExperience property value. Block Windows Spotlight Windows welcome experience
func (m *Windows10GeneralConfiguration) SetWindowsSpotlightBlockWelcomeExperience(value *bool)() {
    m.windowsSpotlightBlockWelcomeExperience = value
}
// SetWindowsSpotlightBlockWindowsTips sets the windowsSpotlightBlockWindowsTips property value. Allows IT admins to turn off the popup of Windows Tips.
func (m *Windows10GeneralConfiguration) SetWindowsSpotlightBlockWindowsTips(value *bool)() {
    m.windowsSpotlightBlockWindowsTips = value
}
// SetWindowsSpotlightConfigureOnLockScreen sets the windowsSpotlightConfigureOnLockScreen property value. Allows IT admind to set a predefined default search engine for MDM-Controlled devices
func (m *Windows10GeneralConfiguration) SetWindowsSpotlightConfigureOnLockScreen(value *WindowsSpotlightEnablementSettings)() {
    m.windowsSpotlightConfigureOnLockScreen = value
}
// SetWindowsStoreBlockAutoUpdate sets the windowsStoreBlockAutoUpdate property value. Indicates whether or not to block automatic update of apps from Windows Store.
func (m *Windows10GeneralConfiguration) SetWindowsStoreBlockAutoUpdate(value *bool)() {
    m.windowsStoreBlockAutoUpdate = value
}
// SetWindowsStoreBlocked sets the windowsStoreBlocked property value. Indicates whether or not to Block the user from using the Windows store.
func (m *Windows10GeneralConfiguration) SetWindowsStoreBlocked(value *bool)() {
    m.windowsStoreBlocked = value
}
// SetWindowsStoreEnablePrivateStoreOnly sets the windowsStoreEnablePrivateStoreOnly property value. Indicates whether or not to enable Private Store Only.
func (m *Windows10GeneralConfiguration) SetWindowsStoreEnablePrivateStoreOnly(value *bool)() {
    m.windowsStoreEnablePrivateStoreOnly = value
}
// SetWirelessDisplayBlockProjectionToThisDevice sets the wirelessDisplayBlockProjectionToThisDevice property value. Indicates whether or not to allow other devices from discovering this PC for projection.
func (m *Windows10GeneralConfiguration) SetWirelessDisplayBlockProjectionToThisDevice(value *bool)() {
    m.wirelessDisplayBlockProjectionToThisDevice = value
}
// SetWirelessDisplayBlockUserInputFromReceiver sets the wirelessDisplayBlockUserInputFromReceiver property value. Indicates whether or not to allow user input from wireless display receiver.
func (m *Windows10GeneralConfiguration) SetWirelessDisplayBlockUserInputFromReceiver(value *bool)() {
    m.wirelessDisplayBlockUserInputFromReceiver = value
}
// SetWirelessDisplayRequirePinForPairing sets the wirelessDisplayRequirePinForPairing property value. Indicates whether or not to require a PIN for new devices to initiate pairing.
func (m *Windows10GeneralConfiguration) SetWirelessDisplayRequirePinForPairing(value *bool)() {
    m.wirelessDisplayRequirePinForPairing = value
}
