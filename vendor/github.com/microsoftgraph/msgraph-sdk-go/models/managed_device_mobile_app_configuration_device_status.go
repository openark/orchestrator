package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDeviceMobileAppConfigurationDeviceStatus contains properties, inherited properties and actions for an MDM mobile app configuration status for a device.
type ManagedDeviceMobileAppConfigurationDeviceStatus struct {
    Entity
    // The DateTime when device compliance grace period expires
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device name of the DevicePolicyStatus.
    deviceDisplayName *string
    // The device model that is being reported
    deviceModel *string
    // Last modified date time of the policy report.
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *ComplianceStatus
    // The User Name that is being reported
    userName *string
    // UserPrincipalName.
    userPrincipalName *string
}
// NewManagedDeviceMobileAppConfigurationDeviceStatus instantiates a new managedDeviceMobileAppConfigurationDeviceStatus and sets the default values.
func NewManagedDeviceMobileAppConfigurationDeviceStatus()(*ManagedDeviceMobileAppConfigurationDeviceStatus) {
    m := &ManagedDeviceMobileAppConfigurationDeviceStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedDeviceMobileAppConfigurationDeviceStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDeviceMobileAppConfigurationDeviceStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDeviceMobileAppConfigurationDeviceStatus(), nil
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.complianceGracePeriodExpirationDateTime
}
// GetDeviceDisplayName gets the deviceDisplayName property value. Device name of the DevicePolicyStatus.
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) GetDeviceDisplayName()(*string) {
    return m.deviceDisplayName
}
// GetDeviceModel gets the deviceModel property value. The device model that is being reported
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) GetDeviceModel()(*string) {
    return m.deviceModel
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceGracePeriodExpirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["deviceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDisplayName(val)
        }
        return nil
    }
    res["deviceModel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["lastReportedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastReportedDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ComplianceStatus))
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetLastReportedDateTime gets the lastReportedDateTime property value. Last modified date time of the policy report.
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetStatus gets the status property value. The status property
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) GetStatus()(*ComplianceStatus) {
    return m.status
}
// GetUserName gets the userName property value. The User Name that is being reported
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) GetUserName()(*string) {
    return m.userName
}
// GetUserPrincipalName gets the userPrincipalName property value. UserPrincipalName.
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("complianceGracePeriodExpirationDateTime", m.GetComplianceGracePeriodExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceDisplayName", m.GetDeviceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastReportedDateTime", m.GetLastReportedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.complianceGracePeriodExpirationDateTime = value
}
// SetDeviceDisplayName sets the deviceDisplayName property value. Device name of the DevicePolicyStatus.
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) SetDeviceDisplayName(value *string)() {
    m.deviceDisplayName = value
}
// SetDeviceModel sets the deviceModel property value. The device model that is being reported
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) SetDeviceModel(value *string)() {
    m.deviceModel = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last modified date time of the policy report.
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) SetStatus(value *ComplianceStatus)() {
    m.status = value
}
// SetUserName sets the userName property value. The User Name that is being reported
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) SetUserName(value *string)() {
    m.userName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. UserPrincipalName.
func (m *ManagedDeviceMobileAppConfigurationDeviceStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
