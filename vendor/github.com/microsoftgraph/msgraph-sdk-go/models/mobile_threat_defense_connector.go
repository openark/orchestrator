package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileThreatDefenseConnector entity which represents a connection to Mobile threat defense partner.
type MobileThreatDefenseConnector struct {
    Entity
    // When TRUE, indicates the data sync partner may collect metadata about installed applications from Intune for IOS devices. When FALSE, indicates the data sync partner may not collect metadata about installed applications from Intune for IOS devices. Default value is FALSE.
    allowPartnerToCollectIOSApplicationMetadata *bool
    // When TRUE, indicates the data sync partner may collect metadata about personally installed applications from Intune for IOS devices. When FALSE, indicates the data sync partner may not collect metadata about personally installed applications from Intune for IOS devices. Default value is FALSE.
    allowPartnerToCollectIOSPersonalApplicationMetadata *bool
    // For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    androidDeviceBlockedOnMissingPartnerData *bool
    // For Android, set whether data from the data sync partner should be used during compliance evaluations
    androidEnabled *bool
    // When TRUE, inidicates that data from the data sync partner can be used during Mobile Application Management (MAM) evaluations for Android devices. When FALSE, inidicates that data from the data sync partner should not be used during Mobile Application Management (MAM) evaluations for Android devices. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
    androidMobileApplicationManagementEnabled *bool
    // For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
    iosDeviceBlockedOnMissingPartnerData *bool
    // For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
    iosEnabled *bool
    // When TRUE, inidicates that data from the data sync partner can be used during Mobile Application Management (MAM) evaluations for IOS devices. When FALSE, inidicates that data from the data sync partner should not be used during Mobile Application Management (MAM) evaluations for IOS devices. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
    iosMobileApplicationManagementEnabled *bool
    // DateTime of last Heartbeat recieved from the Data Sync Partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // When TRUE, inidicates that configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, inidicates that configuration profile management via Microsoft Defender for Endpoint is disabled. Default value is FALSE.
    microsoftDefenderForEndpointAttachEnabled *bool
    // Partner state of this tenant.
    partnerState *MobileThreatPartnerTenantState
    // Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
    partnerUnresponsivenessThresholdInDays *int32
    // Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
    partnerUnsupportedOsVersionBlocked *bool
    // When TRUE, inidicates that Intune must receive data from the data sync partner prior to marking a device compliant for Windows. When FALSE, inidicates that Intune may make a device compliant without receiving data from the data sync partner for Windows. Default value is FALSE.
    windowsDeviceBlockedOnMissingPartnerData *bool
    // When TRUE, inidicates that data from the data sync partner can be used during compliance evaluations for Windows. When FALSE, inidicates that data from the data sync partner should not be used during compliance evaluations for Windows. Default value is FALSE.
    windowsEnabled *bool
}
// NewMobileThreatDefenseConnector instantiates a new mobileThreatDefenseConnector and sets the default values.
func NewMobileThreatDefenseConnector()(*MobileThreatDefenseConnector) {
    m := &MobileThreatDefenseConnector{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMobileThreatDefenseConnectorFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMobileThreatDefenseConnectorFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMobileThreatDefenseConnector(), nil
}
// GetAllowPartnerToCollectIOSApplicationMetadata gets the allowPartnerToCollectIOSApplicationMetadata property value. When TRUE, indicates the data sync partner may collect metadata about installed applications from Intune for IOS devices. When FALSE, indicates the data sync partner may not collect metadata about installed applications from Intune for IOS devices. Default value is FALSE.
func (m *MobileThreatDefenseConnector) GetAllowPartnerToCollectIOSApplicationMetadata()(*bool) {
    return m.allowPartnerToCollectIOSApplicationMetadata
}
// GetAllowPartnerToCollectIOSPersonalApplicationMetadata gets the allowPartnerToCollectIOSPersonalApplicationMetadata property value. When TRUE, indicates the data sync partner may collect metadata about personally installed applications from Intune for IOS devices. When FALSE, indicates the data sync partner may not collect metadata about personally installed applications from Intune for IOS devices. Default value is FALSE.
func (m *MobileThreatDefenseConnector) GetAllowPartnerToCollectIOSPersonalApplicationMetadata()(*bool) {
    return m.allowPartnerToCollectIOSPersonalApplicationMetadata
}
// GetAndroidDeviceBlockedOnMissingPartnerData gets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetAndroidDeviceBlockedOnMissingPartnerData()(*bool) {
    return m.androidDeviceBlockedOnMissingPartnerData
}
// GetAndroidEnabled gets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetAndroidEnabled()(*bool) {
    return m.androidEnabled
}
// GetAndroidMobileApplicationManagementEnabled gets the androidMobileApplicationManagementEnabled property value. When TRUE, inidicates that data from the data sync partner can be used during Mobile Application Management (MAM) evaluations for Android devices. When FALSE, inidicates that data from the data sync partner should not be used during Mobile Application Management (MAM) evaluations for Android devices. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
func (m *MobileThreatDefenseConnector) GetAndroidMobileApplicationManagementEnabled()(*bool) {
    return m.androidMobileApplicationManagementEnabled
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MobileThreatDefenseConnector) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowPartnerToCollectIOSApplicationMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPartnerToCollectIOSApplicationMetadata(val)
        }
        return nil
    }
    res["allowPartnerToCollectIOSPersonalApplicationMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowPartnerToCollectIOSPersonalApplicationMetadata(val)
        }
        return nil
    }
    res["androidDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["androidEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidEnabled(val)
        }
        return nil
    }
    res["androidMobileApplicationManagementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidMobileApplicationManagementEnabled(val)
        }
        return nil
    }
    res["iosDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["iosEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosEnabled(val)
        }
        return nil
    }
    res["iosMobileApplicationManagementEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosMobileApplicationManagementEnabled(val)
        }
        return nil
    }
    res["lastHeartbeatDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastHeartbeatDateTime(val)
        }
        return nil
    }
    res["microsoftDefenderForEndpointAttachEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftDefenderForEndpointAttachEnabled(val)
        }
        return nil
    }
    res["partnerState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMobileThreatPartnerTenantState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerState(val.(*MobileThreatPartnerTenantState))
        }
        return nil
    }
    res["partnerUnresponsivenessThresholdInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerUnresponsivenessThresholdInDays(val)
        }
        return nil
    }
    res["partnerUnsupportedOsVersionBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerUnsupportedOsVersionBlocked(val)
        }
        return nil
    }
    res["windowsDeviceBlockedOnMissingPartnerData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsDeviceBlockedOnMissingPartnerData(val)
        }
        return nil
    }
    res["windowsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsEnabled(val)
        }
        return nil
    }
    return res
}
// GetIosDeviceBlockedOnMissingPartnerData gets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) GetIosDeviceBlockedOnMissingPartnerData()(*bool) {
    return m.iosDeviceBlockedOnMissingPartnerData
}
// GetIosEnabled gets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) GetIosEnabled()(*bool) {
    return m.iosEnabled
}
// GetIosMobileApplicationManagementEnabled gets the iosMobileApplicationManagementEnabled property value. When TRUE, inidicates that data from the data sync partner can be used during Mobile Application Management (MAM) evaluations for IOS devices. When FALSE, inidicates that data from the data sync partner should not be used during Mobile Application Management (MAM) evaluations for IOS devices. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
func (m *MobileThreatDefenseConnector) GetIosMobileApplicationManagementEnabled()(*bool) {
    return m.iosMobileApplicationManagementEnabled
}
// GetLastHeartbeatDateTime gets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastHeartbeatDateTime
}
// GetMicrosoftDefenderForEndpointAttachEnabled gets the microsoftDefenderForEndpointAttachEnabled property value. When TRUE, inidicates that configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, inidicates that configuration profile management via Microsoft Defender for Endpoint is disabled. Default value is FALSE.
func (m *MobileThreatDefenseConnector) GetMicrosoftDefenderForEndpointAttachEnabled()(*bool) {
    return m.microsoftDefenderForEndpointAttachEnabled
}
// GetPartnerState gets the partnerState property value. Partner state of this tenant.
func (m *MobileThreatDefenseConnector) GetPartnerState()(*MobileThreatPartnerTenantState) {
    return m.partnerState
}
// GetPartnerUnresponsivenessThresholdInDays gets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) GetPartnerUnresponsivenessThresholdInDays()(*int32) {
    return m.partnerUnresponsivenessThresholdInDays
}
// GetPartnerUnsupportedOsVersionBlocked gets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) GetPartnerUnsupportedOsVersionBlocked()(*bool) {
    return m.partnerUnsupportedOsVersionBlocked
}
// GetWindowsDeviceBlockedOnMissingPartnerData gets the windowsDeviceBlockedOnMissingPartnerData property value. When TRUE, inidicates that Intune must receive data from the data sync partner prior to marking a device compliant for Windows. When FALSE, inidicates that Intune may make a device compliant without receiving data from the data sync partner for Windows. Default value is FALSE.
func (m *MobileThreatDefenseConnector) GetWindowsDeviceBlockedOnMissingPartnerData()(*bool) {
    return m.windowsDeviceBlockedOnMissingPartnerData
}
// GetWindowsEnabled gets the windowsEnabled property value. When TRUE, inidicates that data from the data sync partner can be used during compliance evaluations for Windows. When FALSE, inidicates that data from the data sync partner should not be used during compliance evaluations for Windows. Default value is FALSE.
func (m *MobileThreatDefenseConnector) GetWindowsEnabled()(*bool) {
    return m.windowsEnabled
}
// Serialize serializes information the current object
func (m *MobileThreatDefenseConnector) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowPartnerToCollectIOSApplicationMetadata", m.GetAllowPartnerToCollectIOSApplicationMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowPartnerToCollectIOSPersonalApplicationMetadata", m.GetAllowPartnerToCollectIOSPersonalApplicationMetadata())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidDeviceBlockedOnMissingPartnerData", m.GetAndroidDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidEnabled", m.GetAndroidEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("androidMobileApplicationManagementEnabled", m.GetAndroidMobileApplicationManagementEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosDeviceBlockedOnMissingPartnerData", m.GetIosDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosEnabled", m.GetIosEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("iosMobileApplicationManagementEnabled", m.GetIosMobileApplicationManagementEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastHeartbeatDateTime", m.GetLastHeartbeatDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("microsoftDefenderForEndpointAttachEnabled", m.GetMicrosoftDefenderForEndpointAttachEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerState() != nil {
        cast := (*m.GetPartnerState()).String()
        err = writer.WriteStringValue("partnerState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("partnerUnresponsivenessThresholdInDays", m.GetPartnerUnresponsivenessThresholdInDays())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("partnerUnsupportedOsVersionBlocked", m.GetPartnerUnsupportedOsVersionBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsDeviceBlockedOnMissingPartnerData", m.GetWindowsDeviceBlockedOnMissingPartnerData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("windowsEnabled", m.GetWindowsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowPartnerToCollectIOSApplicationMetadata sets the allowPartnerToCollectIOSApplicationMetadata property value. When TRUE, indicates the data sync partner may collect metadata about installed applications from Intune for IOS devices. When FALSE, indicates the data sync partner may not collect metadata about installed applications from Intune for IOS devices. Default value is FALSE.
func (m *MobileThreatDefenseConnector) SetAllowPartnerToCollectIOSApplicationMetadata(value *bool)() {
    m.allowPartnerToCollectIOSApplicationMetadata = value
}
// SetAllowPartnerToCollectIOSPersonalApplicationMetadata sets the allowPartnerToCollectIOSPersonalApplicationMetadata property value. When TRUE, indicates the data sync partner may collect metadata about personally installed applications from Intune for IOS devices. When FALSE, indicates the data sync partner may not collect metadata about personally installed applications from Intune for IOS devices. Default value is FALSE.
func (m *MobileThreatDefenseConnector) SetAllowPartnerToCollectIOSPersonalApplicationMetadata(value *bool)() {
    m.allowPartnerToCollectIOSPersonalApplicationMetadata = value
}
// SetAndroidDeviceBlockedOnMissingPartnerData sets the androidDeviceBlockedOnMissingPartnerData property value. For Android, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetAndroidDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.androidDeviceBlockedOnMissingPartnerData = value
}
// SetAndroidEnabled sets the androidEnabled property value. For Android, set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetAndroidEnabled(value *bool)() {
    m.androidEnabled = value
}
// SetAndroidMobileApplicationManagementEnabled sets the androidMobileApplicationManagementEnabled property value. When TRUE, inidicates that data from the data sync partner can be used during Mobile Application Management (MAM) evaluations for Android devices. When FALSE, inidicates that data from the data sync partner should not be used during Mobile Application Management (MAM) evaluations for Android devices. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
func (m *MobileThreatDefenseConnector) SetAndroidMobileApplicationManagementEnabled(value *bool)() {
    m.androidMobileApplicationManagementEnabled = value
}
// SetIosDeviceBlockedOnMissingPartnerData sets the iosDeviceBlockedOnMissingPartnerData property value. For IOS, set whether Intune must receive data from the data sync partner prior to marking a device compliant
func (m *MobileThreatDefenseConnector) SetIosDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.iosDeviceBlockedOnMissingPartnerData = value
}
// SetIosEnabled sets the iosEnabled property value. For IOS, get or set whether data from the data sync partner should be used during compliance evaluations
func (m *MobileThreatDefenseConnector) SetIosEnabled(value *bool)() {
    m.iosEnabled = value
}
// SetIosMobileApplicationManagementEnabled sets the iosMobileApplicationManagementEnabled property value. When TRUE, inidicates that data from the data sync partner can be used during Mobile Application Management (MAM) evaluations for IOS devices. When FALSE, inidicates that data from the data sync partner should not be used during Mobile Application Management (MAM) evaluations for IOS devices. Only one partner per platform may be enabled for Mobile Application Management (MAM) evaluation. Default value is FALSE.
func (m *MobileThreatDefenseConnector) SetIosMobileApplicationManagementEnabled(value *bool)() {
    m.iosMobileApplicationManagementEnabled = value
}
// SetLastHeartbeatDateTime sets the lastHeartbeatDateTime property value. DateTime of last Heartbeat recieved from the Data Sync Partner
func (m *MobileThreatDefenseConnector) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
// SetMicrosoftDefenderForEndpointAttachEnabled sets the microsoftDefenderForEndpointAttachEnabled property value. When TRUE, inidicates that configuration profile management via Microsoft Defender for Endpoint is enabled. When FALSE, inidicates that configuration profile management via Microsoft Defender for Endpoint is disabled. Default value is FALSE.
func (m *MobileThreatDefenseConnector) SetMicrosoftDefenderForEndpointAttachEnabled(value *bool)() {
    m.microsoftDefenderForEndpointAttachEnabled = value
}
// SetPartnerState sets the partnerState property value. Partner state of this tenant.
func (m *MobileThreatDefenseConnector) SetPartnerState(value *MobileThreatPartnerTenantState)() {
    m.partnerState = value
}
// SetPartnerUnresponsivenessThresholdInDays sets the partnerUnresponsivenessThresholdInDays property value. Get or Set days the per tenant tolerance to unresponsiveness for this partner integration
func (m *MobileThreatDefenseConnector) SetPartnerUnresponsivenessThresholdInDays(value *int32)() {
    m.partnerUnresponsivenessThresholdInDays = value
}
// SetPartnerUnsupportedOsVersionBlocked sets the partnerUnsupportedOsVersionBlocked property value. Get or set whether to block devices on the enabled platforms that do not meet the minimum version requirements of the Data Sync Partner
func (m *MobileThreatDefenseConnector) SetPartnerUnsupportedOsVersionBlocked(value *bool)() {
    m.partnerUnsupportedOsVersionBlocked = value
}
// SetWindowsDeviceBlockedOnMissingPartnerData sets the windowsDeviceBlockedOnMissingPartnerData property value. When TRUE, inidicates that Intune must receive data from the data sync partner prior to marking a device compliant for Windows. When FALSE, inidicates that Intune may make a device compliant without receiving data from the data sync partner for Windows. Default value is FALSE.
func (m *MobileThreatDefenseConnector) SetWindowsDeviceBlockedOnMissingPartnerData(value *bool)() {
    m.windowsDeviceBlockedOnMissingPartnerData = value
}
// SetWindowsEnabled sets the windowsEnabled property value. When TRUE, inidicates that data from the data sync partner can be used during compliance evaluations for Windows. When FALSE, inidicates that data from the data sync partner should not be used during compliance evaluations for Windows. Default value is FALSE.
func (m *MobileThreatDefenseConnector) SetWindowsEnabled(value *bool)() {
    m.windowsEnabled = value
}
