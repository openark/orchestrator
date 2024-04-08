package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationUserStatus 
type DeviceConfigurationUserStatus struct {
    Entity
    // Devices count for that user.
    devicesCount *int32
    // Last modified date time of the policy report.
    lastReportedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status property
    status *ComplianceStatus
    // User name of the DevicePolicyStatus.
    userDisplayName *string
    // UserPrincipalName.
    userPrincipalName *string
}
// NewDeviceConfigurationUserStatus instantiates a new deviceConfigurationUserStatus and sets the default values.
func NewDeviceConfigurationUserStatus()(*DeviceConfigurationUserStatus) {
    m := &DeviceConfigurationUserStatus{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationUserStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationUserStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationUserStatus(), nil
}
// GetDevicesCount gets the devicesCount property value. Devices count for that user.
func (m *DeviceConfigurationUserStatus) GetDevicesCount()(*int32) {
    return m.devicesCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationUserStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["devicesCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevicesCount(val)
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
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
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
func (m *DeviceConfigurationUserStatus) GetLastReportedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastReportedDateTime
}
// GetStatus gets the status property value. The status property
func (m *DeviceConfigurationUserStatus) GetStatus()(*ComplianceStatus) {
    return m.status
}
// GetUserDisplayName gets the userDisplayName property value. User name of the DevicePolicyStatus.
func (m *DeviceConfigurationUserStatus) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserPrincipalName gets the userPrincipalName property value. UserPrincipalName.
func (m *DeviceConfigurationUserStatus) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DeviceConfigurationUserStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("devicesCount", m.GetDevicesCount())
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
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
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
// SetDevicesCount sets the devicesCount property value. Devices count for that user.
func (m *DeviceConfigurationUserStatus) SetDevicesCount(value *int32)() {
    m.devicesCount = value
}
// SetLastReportedDateTime sets the lastReportedDateTime property value. Last modified date time of the policy report.
func (m *DeviceConfigurationUserStatus) SetLastReportedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastReportedDateTime = value
}
// SetStatus sets the status property value. The status property
func (m *DeviceConfigurationUserStatus) SetStatus(value *ComplianceStatus)() {
    m.status = value
}
// SetUserDisplayName sets the userDisplayName property value. User name of the DevicePolicyStatus.
func (m *DeviceConfigurationUserStatus) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. UserPrincipalName.
func (m *DeviceConfigurationUserStatus) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
