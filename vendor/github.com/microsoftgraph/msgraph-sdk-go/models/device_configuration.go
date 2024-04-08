package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfiguration device Configuration.
type DeviceConfiguration struct {
    Entity
    // The list of assignments for the device configuration profile.
    assignments []DeviceConfigurationAssignmentable
    // DateTime the object was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Admin provided description of the Device Configuration.
    description *string
    // Device Configuration Setting State Device Summary
    deviceSettingStateSummaries []SettingStateDeviceSummaryable
    // Device configuration installation status by device.
    deviceStatuses []DeviceConfigurationDeviceStatusable
    // Device Configuration devices status overview
    deviceStatusOverview DeviceConfigurationDeviceOverviewable
    // Admin provided name of the device configuration.
    displayName *string
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device configuration installation status by user.
    userStatuses []DeviceConfigurationUserStatusable
    // Device Configuration users status overview
    userStatusOverview DeviceConfigurationUserOverviewable
    // Version of the device configuration.
    version *int32
}
// NewDeviceConfiguration instantiates a new deviceConfiguration and sets the default values.
func NewDeviceConfiguration()(*DeviceConfiguration) {
    m := &DeviceConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.androidCustomConfiguration":
                        return NewAndroidCustomConfiguration(), nil
                    case "#microsoft.graph.androidGeneralDeviceConfiguration":
                        return NewAndroidGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileCustomConfiguration":
                        return NewAndroidWorkProfileCustomConfiguration(), nil
                    case "#microsoft.graph.androidWorkProfileGeneralDeviceConfiguration":
                        return NewAndroidWorkProfileGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.appleDeviceFeaturesConfigurationBase":
                        return NewAppleDeviceFeaturesConfigurationBase(), nil
                    case "#microsoft.graph.editionUpgradeConfiguration":
                        return NewEditionUpgradeConfiguration(), nil
                    case "#microsoft.graph.iosCertificateProfile":
                        return NewIosCertificateProfile(), nil
                    case "#microsoft.graph.iosCustomConfiguration":
                        return NewIosCustomConfiguration(), nil
                    case "#microsoft.graph.iosDeviceFeaturesConfiguration":
                        return NewIosDeviceFeaturesConfiguration(), nil
                    case "#microsoft.graph.iosGeneralDeviceConfiguration":
                        return NewIosGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.iosUpdateConfiguration":
                        return NewIosUpdateConfiguration(), nil
                    case "#microsoft.graph.macOSCustomConfiguration":
                        return NewMacOSCustomConfiguration(), nil
                    case "#microsoft.graph.macOSDeviceFeaturesConfiguration":
                        return NewMacOSDeviceFeaturesConfiguration(), nil
                    case "#microsoft.graph.macOSGeneralDeviceConfiguration":
                        return NewMacOSGeneralDeviceConfiguration(), nil
                    case "#microsoft.graph.sharedPCConfiguration":
                        return NewSharedPCConfiguration(), nil
                    case "#microsoft.graph.windows10CustomConfiguration":
                        return NewWindows10CustomConfiguration(), nil
                    case "#microsoft.graph.windows10EndpointProtectionConfiguration":
                        return NewWindows10EndpointProtectionConfiguration(), nil
                    case "#microsoft.graph.windows10EnterpriseModernAppManagementConfiguration":
                        return NewWindows10EnterpriseModernAppManagementConfiguration(), nil
                    case "#microsoft.graph.windows10GeneralConfiguration":
                        return NewWindows10GeneralConfiguration(), nil
                    case "#microsoft.graph.windows10SecureAssessmentConfiguration":
                        return NewWindows10SecureAssessmentConfiguration(), nil
                    case "#microsoft.graph.windows10TeamGeneralConfiguration":
                        return NewWindows10TeamGeneralConfiguration(), nil
                    case "#microsoft.graph.windows81GeneralConfiguration":
                        return NewWindows81GeneralConfiguration(), nil
                    case "#microsoft.graph.windowsDefenderAdvancedThreatProtectionConfiguration":
                        return NewWindowsDefenderAdvancedThreatProtectionConfiguration(), nil
                    case "#microsoft.graph.windowsPhone81CustomConfiguration":
                        return NewWindowsPhone81CustomConfiguration(), nil
                    case "#microsoft.graph.windowsPhone81GeneralConfiguration":
                        return NewWindowsPhone81GeneralConfiguration(), nil
                    case "#microsoft.graph.windowsUpdateForBusinessConfiguration":
                        return NewWindowsUpdateForBusinessConfiguration(), nil
                }
            }
        }
    }
    return NewDeviceConfiguration(), nil
}
// GetAssignments gets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceConfiguration) GetAssignments()([]DeviceConfigurationAssignmentable) {
    return m.assignments
}
// GetCreatedDateTime gets the createdDateTime property value. DateTime the object was created.
func (m *DeviceConfiguration) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceConfiguration) GetDescription()(*string) {
    return m.description
}
// GetDeviceSettingStateSummaries gets the deviceSettingStateSummaries property value. Device Configuration Setting State Device Summary
func (m *DeviceConfiguration) GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable) {
    return m.deviceSettingStateSummaries
}
// GetDeviceStatuses gets the deviceStatuses property value. Device configuration installation status by device.
func (m *DeviceConfiguration) GetDeviceStatuses()([]DeviceConfigurationDeviceStatusable) {
    return m.deviceStatuses
}
// GetDeviceStatusOverview gets the deviceStatusOverview property value. Device Configuration devices status overview
func (m *DeviceConfiguration) GetDeviceStatusOverview()(DeviceConfigurationDeviceOverviewable) {
    return m.deviceStatusOverview
}
// GetDisplayName gets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["deviceSettingStateSummaries"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingStateDeviceSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingStateDeviceSummaryable, len(val))
            for i, v := range val {
                res[i] = v.(SettingStateDeviceSummaryable)
            }
            m.SetDeviceSettingStateSummaries(res)
        }
        return nil
    }
    res["deviceStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationDeviceStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationDeviceStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationDeviceStatusable)
            }
            m.SetDeviceStatuses(res)
        }
        return nil
    }
    res["deviceStatusOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationDeviceOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceStatusOverview(val.(DeviceConfigurationDeviceOverviewable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["userStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationUserStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationUserStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationUserStatusable)
            }
            m.SetUserStatuses(res)
        }
        return nil
    }
    res["userStatusOverview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceConfigurationUserOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserStatusOverview(val.(DeviceConfigurationUserOverviewable))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceConfiguration) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetUserStatuses gets the userStatuses property value. Device configuration installation status by user.
func (m *DeviceConfiguration) GetUserStatuses()([]DeviceConfigurationUserStatusable) {
    return m.userStatuses
}
// GetUserStatusOverview gets the userStatusOverview property value. Device Configuration users status overview
func (m *DeviceConfiguration) GetUserStatusOverview()(DeviceConfigurationUserOverviewable) {
    return m.userStatusOverview
}
// GetVersion gets the version property value. Version of the device configuration.
func (m *DeviceConfiguration) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetDeviceSettingStateSummaries() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceSettingStateSummaries()))
        for i, v := range m.GetDeviceSettingStateSummaries() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceSettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeviceStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStatuses()))
        for i, v := range m.GetDeviceStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceStatusOverview", m.GetDeviceStatusOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetUserStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserStatuses()))
        for i, v := range m.GetUserStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userStatusOverview", m.GetUserStatusOverview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignments sets the assignments property value. The list of assignments for the device configuration profile.
func (m *DeviceConfiguration) SetAssignments(value []DeviceConfigurationAssignmentable)() {
    m.assignments = value
}
// SetCreatedDateTime sets the createdDateTime property value. DateTime the object was created.
func (m *DeviceConfiguration) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Admin provided description of the Device Configuration.
func (m *DeviceConfiguration) SetDescription(value *string)() {
    m.description = value
}
// SetDeviceSettingStateSummaries sets the deviceSettingStateSummaries property value. Device Configuration Setting State Device Summary
func (m *DeviceConfiguration) SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)() {
    m.deviceSettingStateSummaries = value
}
// SetDeviceStatuses sets the deviceStatuses property value. Device configuration installation status by device.
func (m *DeviceConfiguration) SetDeviceStatuses(value []DeviceConfigurationDeviceStatusable)() {
    m.deviceStatuses = value
}
// SetDeviceStatusOverview sets the deviceStatusOverview property value. Device Configuration devices status overview
func (m *DeviceConfiguration) SetDeviceStatusOverview(value DeviceConfigurationDeviceOverviewable)() {
    m.deviceStatusOverview = value
}
// SetDisplayName sets the displayName property value. Admin provided name of the device configuration.
func (m *DeviceConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *DeviceConfiguration) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetUserStatuses sets the userStatuses property value. Device configuration installation status by user.
func (m *DeviceConfiguration) SetUserStatuses(value []DeviceConfigurationUserStatusable)() {
    m.userStatuses = value
}
// SetUserStatusOverview sets the userStatusOverview property value. Device Configuration users status overview
func (m *DeviceConfiguration) SetUserStatusOverview(value DeviceConfigurationUserOverviewable)() {
    m.userStatusOverview = value
}
// SetVersion sets the version property value. Version of the device configuration.
func (m *DeviceConfiguration) SetVersion(value *int32)() {
    m.version = value
}
