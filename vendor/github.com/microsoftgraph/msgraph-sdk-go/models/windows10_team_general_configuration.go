package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10TeamGeneralConfiguration 
type Windows10TeamGeneralConfiguration struct {
    DeviceConfiguration
    // Indicates whether or not to Block Azure Operational Insights.
    azureOperationalInsightsBlockTelemetry *bool
    // The Azure Operational Insights workspace id.
    azureOperationalInsightsWorkspaceId *string
    // The Azure Operational Insights Workspace key.
    azureOperationalInsightsWorkspaceKey *string
    // Specifies whether to automatically launch the Connect app whenever a projection is initiated.
    connectAppBlockAutoLaunch *bool
    // Indicates whether or not to Block setting a maintenance window for device updates.
    maintenanceWindowBlocked *bool
    // Maintenance window duration for device updates. Valid values 0 to 5
    maintenanceWindowDurationInHours *int32
    // Maintenance window start time for device updates.
    maintenanceWindowStartTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // Indicates whether or not to Block wireless projection.
    miracastBlocked *bool
    // Possible values for Miracast channel.
    miracastChannel *MiracastChannel
    // Indicates whether or not to require a pin for wireless projection.
    miracastRequirePin *bool
    // Specifies whether to disable the 'My meetings and files' feature in the Start menu, which shows the signed-in user's meetings and files from Office 365.
    settingsBlockMyMeetingsAndFiles *bool
    // Specifies whether to allow the ability to resume a session when the session times out.
    settingsBlockSessionResume *bool
    // Specifies whether to disable auto-populating of the sign-in dialog with invitees from scheduled meetings.
    settingsBlockSigninSuggestions *bool
    // Specifies the default volume value for a new session. Permitted values are 0-100. The default is 45. Valid values 0 to 100
    settingsDefaultVolume *int32
    // Specifies the number of minutes until the Hub screen turns off.
    settingsScreenTimeoutInMinutes *int32
    // Specifies the number of minutes until the session times out.
    settingsSessionTimeoutInMinutes *int32
    // Specifies the number of minutes until the Hub enters sleep mode.
    settingsSleepTimeoutInMinutes *int32
    // The welcome screen background image URL. The URL must use the HTTPS protocol and return a PNG image.
    welcomeScreenBackgroundImageUrl *string
    // Indicates whether or not to Block the welcome screen from waking up automatically when someone enters the room.
    welcomeScreenBlockAutomaticWakeUp *bool
    // Possible values for welcome screen meeting information.
    welcomeScreenMeetingInformation *WelcomeScreenMeetingInformation
}
// NewWindows10TeamGeneralConfiguration instantiates a new Windows10TeamGeneralConfiguration and sets the default values.
func NewWindows10TeamGeneralConfiguration()(*Windows10TeamGeneralConfiguration) {
    m := &Windows10TeamGeneralConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10TeamGeneralConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10TeamGeneralConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10TeamGeneralConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10TeamGeneralConfiguration(), nil
}
// GetAzureOperationalInsightsBlockTelemetry gets the azureOperationalInsightsBlockTelemetry property value. Indicates whether or not to Block Azure Operational Insights.
func (m *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsBlockTelemetry()(*bool) {
    return m.azureOperationalInsightsBlockTelemetry
}
// GetAzureOperationalInsightsWorkspaceId gets the azureOperationalInsightsWorkspaceId property value. The Azure Operational Insights workspace id.
func (m *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceId()(*string) {
    return m.azureOperationalInsightsWorkspaceId
}
// GetAzureOperationalInsightsWorkspaceKey gets the azureOperationalInsightsWorkspaceKey property value. The Azure Operational Insights Workspace key.
func (m *Windows10TeamGeneralConfiguration) GetAzureOperationalInsightsWorkspaceKey()(*string) {
    return m.azureOperationalInsightsWorkspaceKey
}
// GetConnectAppBlockAutoLaunch gets the connectAppBlockAutoLaunch property value. Specifies whether to automatically launch the Connect app whenever a projection is initiated.
func (m *Windows10TeamGeneralConfiguration) GetConnectAppBlockAutoLaunch()(*bool) {
    return m.connectAppBlockAutoLaunch
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10TeamGeneralConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["azureOperationalInsightsBlockTelemetry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureOperationalInsightsBlockTelemetry(val)
        }
        return nil
    }
    res["azureOperationalInsightsWorkspaceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureOperationalInsightsWorkspaceId(val)
        }
        return nil
    }
    res["azureOperationalInsightsWorkspaceKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureOperationalInsightsWorkspaceKey(val)
        }
        return nil
    }
    res["connectAppBlockAutoLaunch"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectAppBlockAutoLaunch(val)
        }
        return nil
    }
    res["maintenanceWindowBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintenanceWindowBlocked(val)
        }
        return nil
    }
    res["maintenanceWindowDurationInHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintenanceWindowDurationInHours(val)
        }
        return nil
    }
    res["maintenanceWindowStartTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaintenanceWindowStartTime(val)
        }
        return nil
    }
    res["miracastBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiracastBlocked(val)
        }
        return nil
    }
    res["miracastChannel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMiracastChannel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiracastChannel(val.(*MiracastChannel))
        }
        return nil
    }
    res["miracastRequirePin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiracastRequirePin(val)
        }
        return nil
    }
    res["settingsBlockMyMeetingsAndFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockMyMeetingsAndFiles(val)
        }
        return nil
    }
    res["settingsBlockSessionResume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockSessionResume(val)
        }
        return nil
    }
    res["settingsBlockSigninSuggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsBlockSigninSuggestions(val)
        }
        return nil
    }
    res["settingsDefaultVolume"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsDefaultVolume(val)
        }
        return nil
    }
    res["settingsScreenTimeoutInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsScreenTimeoutInMinutes(val)
        }
        return nil
    }
    res["settingsSessionTimeoutInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsSessionTimeoutInMinutes(val)
        }
        return nil
    }
    res["settingsSleepTimeoutInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingsSleepTimeoutInMinutes(val)
        }
        return nil
    }
    res["welcomeScreenBackgroundImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWelcomeScreenBackgroundImageUrl(val)
        }
        return nil
    }
    res["welcomeScreenBlockAutomaticWakeUp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWelcomeScreenBlockAutomaticWakeUp(val)
        }
        return nil
    }
    res["welcomeScreenMeetingInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWelcomeScreenMeetingInformation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWelcomeScreenMeetingInformation(val.(*WelcomeScreenMeetingInformation))
        }
        return nil
    }
    return res
}
// GetMaintenanceWindowBlocked gets the maintenanceWindowBlocked property value. Indicates whether or not to Block setting a maintenance window for device updates.
func (m *Windows10TeamGeneralConfiguration) GetMaintenanceWindowBlocked()(*bool) {
    return m.maintenanceWindowBlocked
}
// GetMaintenanceWindowDurationInHours gets the maintenanceWindowDurationInHours property value. Maintenance window duration for device updates. Valid values 0 to 5
func (m *Windows10TeamGeneralConfiguration) GetMaintenanceWindowDurationInHours()(*int32) {
    return m.maintenanceWindowDurationInHours
}
// GetMaintenanceWindowStartTime gets the maintenanceWindowStartTime property value. Maintenance window start time for device updates.
func (m *Windows10TeamGeneralConfiguration) GetMaintenanceWindowStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.maintenanceWindowStartTime
}
// GetMiracastBlocked gets the miracastBlocked property value. Indicates whether or not to Block wireless projection.
func (m *Windows10TeamGeneralConfiguration) GetMiracastBlocked()(*bool) {
    return m.miracastBlocked
}
// GetMiracastChannel gets the miracastChannel property value. Possible values for Miracast channel.
func (m *Windows10TeamGeneralConfiguration) GetMiracastChannel()(*MiracastChannel) {
    return m.miracastChannel
}
// GetMiracastRequirePin gets the miracastRequirePin property value. Indicates whether or not to require a pin for wireless projection.
func (m *Windows10TeamGeneralConfiguration) GetMiracastRequirePin()(*bool) {
    return m.miracastRequirePin
}
// GetSettingsBlockMyMeetingsAndFiles gets the settingsBlockMyMeetingsAndFiles property value. Specifies whether to disable the 'My meetings and files' feature in the Start menu, which shows the signed-in user's meetings and files from Office 365.
func (m *Windows10TeamGeneralConfiguration) GetSettingsBlockMyMeetingsAndFiles()(*bool) {
    return m.settingsBlockMyMeetingsAndFiles
}
// GetSettingsBlockSessionResume gets the settingsBlockSessionResume property value. Specifies whether to allow the ability to resume a session when the session times out.
func (m *Windows10TeamGeneralConfiguration) GetSettingsBlockSessionResume()(*bool) {
    return m.settingsBlockSessionResume
}
// GetSettingsBlockSigninSuggestions gets the settingsBlockSigninSuggestions property value. Specifies whether to disable auto-populating of the sign-in dialog with invitees from scheduled meetings.
func (m *Windows10TeamGeneralConfiguration) GetSettingsBlockSigninSuggestions()(*bool) {
    return m.settingsBlockSigninSuggestions
}
// GetSettingsDefaultVolume gets the settingsDefaultVolume property value. Specifies the default volume value for a new session. Permitted values are 0-100. The default is 45. Valid values 0 to 100
func (m *Windows10TeamGeneralConfiguration) GetSettingsDefaultVolume()(*int32) {
    return m.settingsDefaultVolume
}
// GetSettingsScreenTimeoutInMinutes gets the settingsScreenTimeoutInMinutes property value. Specifies the number of minutes until the Hub screen turns off.
func (m *Windows10TeamGeneralConfiguration) GetSettingsScreenTimeoutInMinutes()(*int32) {
    return m.settingsScreenTimeoutInMinutes
}
// GetSettingsSessionTimeoutInMinutes gets the settingsSessionTimeoutInMinutes property value. Specifies the number of minutes until the session times out.
func (m *Windows10TeamGeneralConfiguration) GetSettingsSessionTimeoutInMinutes()(*int32) {
    return m.settingsSessionTimeoutInMinutes
}
// GetSettingsSleepTimeoutInMinutes gets the settingsSleepTimeoutInMinutes property value. Specifies the number of minutes until the Hub enters sleep mode.
func (m *Windows10TeamGeneralConfiguration) GetSettingsSleepTimeoutInMinutes()(*int32) {
    return m.settingsSleepTimeoutInMinutes
}
// GetWelcomeScreenBackgroundImageUrl gets the welcomeScreenBackgroundImageUrl property value. The welcome screen background image URL. The URL must use the HTTPS protocol and return a PNG image.
func (m *Windows10TeamGeneralConfiguration) GetWelcomeScreenBackgroundImageUrl()(*string) {
    return m.welcomeScreenBackgroundImageUrl
}
// GetWelcomeScreenBlockAutomaticWakeUp gets the welcomeScreenBlockAutomaticWakeUp property value. Indicates whether or not to Block the welcome screen from waking up automatically when someone enters the room.
func (m *Windows10TeamGeneralConfiguration) GetWelcomeScreenBlockAutomaticWakeUp()(*bool) {
    return m.welcomeScreenBlockAutomaticWakeUp
}
// GetWelcomeScreenMeetingInformation gets the welcomeScreenMeetingInformation property value. Possible values for welcome screen meeting information.
func (m *Windows10TeamGeneralConfiguration) GetWelcomeScreenMeetingInformation()(*WelcomeScreenMeetingInformation) {
    return m.welcomeScreenMeetingInformation
}
// Serialize serializes information the current object
func (m *Windows10TeamGeneralConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("azureOperationalInsightsBlockTelemetry", m.GetAzureOperationalInsightsBlockTelemetry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureOperationalInsightsWorkspaceId", m.GetAzureOperationalInsightsWorkspaceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureOperationalInsightsWorkspaceKey", m.GetAzureOperationalInsightsWorkspaceKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("connectAppBlockAutoLaunch", m.GetConnectAppBlockAutoLaunch())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("maintenanceWindowBlocked", m.GetMaintenanceWindowBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maintenanceWindowDurationInHours", m.GetMaintenanceWindowDurationInHours())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeOnlyValue("maintenanceWindowStartTime", m.GetMaintenanceWindowStartTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("miracastBlocked", m.GetMiracastBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetMiracastChannel() != nil {
        cast := (*m.GetMiracastChannel()).String()
        err = writer.WriteStringValue("miracastChannel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("miracastRequirePin", m.GetMiracastRequirePin())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockMyMeetingsAndFiles", m.GetSettingsBlockMyMeetingsAndFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockSessionResume", m.GetSettingsBlockSessionResume())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("settingsBlockSigninSuggestions", m.GetSettingsBlockSigninSuggestions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("settingsDefaultVolume", m.GetSettingsDefaultVolume())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("settingsScreenTimeoutInMinutes", m.GetSettingsScreenTimeoutInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("settingsSessionTimeoutInMinutes", m.GetSettingsSessionTimeoutInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("settingsSleepTimeoutInMinutes", m.GetSettingsSleepTimeoutInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("welcomeScreenBackgroundImageUrl", m.GetWelcomeScreenBackgroundImageUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("welcomeScreenBlockAutomaticWakeUp", m.GetWelcomeScreenBlockAutomaticWakeUp())
        if err != nil {
            return err
        }
    }
    if m.GetWelcomeScreenMeetingInformation() != nil {
        cast := (*m.GetWelcomeScreenMeetingInformation()).String()
        err = writer.WriteStringValue("welcomeScreenMeetingInformation", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureOperationalInsightsBlockTelemetry sets the azureOperationalInsightsBlockTelemetry property value. Indicates whether or not to Block Azure Operational Insights.
func (m *Windows10TeamGeneralConfiguration) SetAzureOperationalInsightsBlockTelemetry(value *bool)() {
    m.azureOperationalInsightsBlockTelemetry = value
}
// SetAzureOperationalInsightsWorkspaceId sets the azureOperationalInsightsWorkspaceId property value. The Azure Operational Insights workspace id.
func (m *Windows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceId(value *string)() {
    m.azureOperationalInsightsWorkspaceId = value
}
// SetAzureOperationalInsightsWorkspaceKey sets the azureOperationalInsightsWorkspaceKey property value. The Azure Operational Insights Workspace key.
func (m *Windows10TeamGeneralConfiguration) SetAzureOperationalInsightsWorkspaceKey(value *string)() {
    m.azureOperationalInsightsWorkspaceKey = value
}
// SetConnectAppBlockAutoLaunch sets the connectAppBlockAutoLaunch property value. Specifies whether to automatically launch the Connect app whenever a projection is initiated.
func (m *Windows10TeamGeneralConfiguration) SetConnectAppBlockAutoLaunch(value *bool)() {
    m.connectAppBlockAutoLaunch = value
}
// SetMaintenanceWindowBlocked sets the maintenanceWindowBlocked property value. Indicates whether or not to Block setting a maintenance window for device updates.
func (m *Windows10TeamGeneralConfiguration) SetMaintenanceWindowBlocked(value *bool)() {
    m.maintenanceWindowBlocked = value
}
// SetMaintenanceWindowDurationInHours sets the maintenanceWindowDurationInHours property value. Maintenance window duration for device updates. Valid values 0 to 5
func (m *Windows10TeamGeneralConfiguration) SetMaintenanceWindowDurationInHours(value *int32)() {
    m.maintenanceWindowDurationInHours = value
}
// SetMaintenanceWindowStartTime sets the maintenanceWindowStartTime property value. Maintenance window start time for device updates.
func (m *Windows10TeamGeneralConfiguration) SetMaintenanceWindowStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.maintenanceWindowStartTime = value
}
// SetMiracastBlocked sets the miracastBlocked property value. Indicates whether or not to Block wireless projection.
func (m *Windows10TeamGeneralConfiguration) SetMiracastBlocked(value *bool)() {
    m.miracastBlocked = value
}
// SetMiracastChannel sets the miracastChannel property value. Possible values for Miracast channel.
func (m *Windows10TeamGeneralConfiguration) SetMiracastChannel(value *MiracastChannel)() {
    m.miracastChannel = value
}
// SetMiracastRequirePin sets the miracastRequirePin property value. Indicates whether or not to require a pin for wireless projection.
func (m *Windows10TeamGeneralConfiguration) SetMiracastRequirePin(value *bool)() {
    m.miracastRequirePin = value
}
// SetSettingsBlockMyMeetingsAndFiles sets the settingsBlockMyMeetingsAndFiles property value. Specifies whether to disable the 'My meetings and files' feature in the Start menu, which shows the signed-in user's meetings and files from Office 365.
func (m *Windows10TeamGeneralConfiguration) SetSettingsBlockMyMeetingsAndFiles(value *bool)() {
    m.settingsBlockMyMeetingsAndFiles = value
}
// SetSettingsBlockSessionResume sets the settingsBlockSessionResume property value. Specifies whether to allow the ability to resume a session when the session times out.
func (m *Windows10TeamGeneralConfiguration) SetSettingsBlockSessionResume(value *bool)() {
    m.settingsBlockSessionResume = value
}
// SetSettingsBlockSigninSuggestions sets the settingsBlockSigninSuggestions property value. Specifies whether to disable auto-populating of the sign-in dialog with invitees from scheduled meetings.
func (m *Windows10TeamGeneralConfiguration) SetSettingsBlockSigninSuggestions(value *bool)() {
    m.settingsBlockSigninSuggestions = value
}
// SetSettingsDefaultVolume sets the settingsDefaultVolume property value. Specifies the default volume value for a new session. Permitted values are 0-100. The default is 45. Valid values 0 to 100
func (m *Windows10TeamGeneralConfiguration) SetSettingsDefaultVolume(value *int32)() {
    m.settingsDefaultVolume = value
}
// SetSettingsScreenTimeoutInMinutes sets the settingsScreenTimeoutInMinutes property value. Specifies the number of minutes until the Hub screen turns off.
func (m *Windows10TeamGeneralConfiguration) SetSettingsScreenTimeoutInMinutes(value *int32)() {
    m.settingsScreenTimeoutInMinutes = value
}
// SetSettingsSessionTimeoutInMinutes sets the settingsSessionTimeoutInMinutes property value. Specifies the number of minutes until the session times out.
func (m *Windows10TeamGeneralConfiguration) SetSettingsSessionTimeoutInMinutes(value *int32)() {
    m.settingsSessionTimeoutInMinutes = value
}
// SetSettingsSleepTimeoutInMinutes sets the settingsSleepTimeoutInMinutes property value. Specifies the number of minutes until the Hub enters sleep mode.
func (m *Windows10TeamGeneralConfiguration) SetSettingsSleepTimeoutInMinutes(value *int32)() {
    m.settingsSleepTimeoutInMinutes = value
}
// SetWelcomeScreenBackgroundImageUrl sets the welcomeScreenBackgroundImageUrl property value. The welcome screen background image URL. The URL must use the HTTPS protocol and return a PNG image.
func (m *Windows10TeamGeneralConfiguration) SetWelcomeScreenBackgroundImageUrl(value *string)() {
    m.welcomeScreenBackgroundImageUrl = value
}
// SetWelcomeScreenBlockAutomaticWakeUp sets the welcomeScreenBlockAutomaticWakeUp property value. Indicates whether or not to Block the welcome screen from waking up automatically when someone enters the room.
func (m *Windows10TeamGeneralConfiguration) SetWelcomeScreenBlockAutomaticWakeUp(value *bool)() {
    m.welcomeScreenBlockAutomaticWakeUp = value
}
// SetWelcomeScreenMeetingInformation sets the welcomeScreenMeetingInformation property value. Possible values for welcome screen meeting information.
func (m *Windows10TeamGeneralConfiguration) SetWelcomeScreenMeetingInformation(value *WelcomeScreenMeetingInformation)() {
    m.welcomeScreenMeetingInformation = value
}
