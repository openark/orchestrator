package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceManagementPartner entity which represents a connection to device management partner.
type DeviceManagementPartner struct {
    Entity
    // Partner display name
    displayName *string
    // User groups that specifies whether enrollment is through partner.
    groupsRequiringPartnerEnrollment []DeviceManagementPartnerAssignmentable
    // Whether device management partner is configured or not
    isConfigured *bool
    // Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
    lastHeartbeatDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Partner App Type.
    partnerAppType *DeviceManagementPartnerAppType
    // Partner state of this tenant.
    partnerState *DeviceManagementPartnerTenantState
    // Partner Single tenant App id
    singleTenantAppId *string
    // DateTime in UTC when PartnerDevices will be marked as NonCompliant
    whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // DateTime in UTC when PartnerDevices will be removed
    whenPartnerDevicesWillBeRemovedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewDeviceManagementPartner instantiates a new deviceManagementPartner and sets the default values.
func NewDeviceManagementPartner()(*DeviceManagementPartner) {
    m := &DeviceManagementPartner{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceManagementPartnerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceManagementPartnerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceManagementPartner(), nil
}
// GetDisplayName gets the displayName property value. Partner display name
func (m *DeviceManagementPartner) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceManagementPartner) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["groupsRequiringPartnerEnrollment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceManagementPartnerAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceManagementPartnerAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceManagementPartnerAssignmentable)
            }
            m.SetGroupsRequiringPartnerEnrollment(res)
        }
        return nil
    }
    res["isConfigured"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsConfigured(val)
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
    res["partnerAppType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerAppType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerAppType(val.(*DeviceManagementPartnerAppType))
        }
        return nil
    }
    res["partnerState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementPartnerTenantState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerState(val.(*DeviceManagementPartnerTenantState))
        }
        return nil
    }
    res["singleTenantAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSingleTenantAppId(val)
        }
        return nil
    }
    res["whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime(val)
        }
        return nil
    }
    res["whenPartnerDevicesWillBeRemovedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWhenPartnerDevicesWillBeRemovedDateTime(val)
        }
        return nil
    }
    return res
}
// GetGroupsRequiringPartnerEnrollment gets the groupsRequiringPartnerEnrollment property value. User groups that specifies whether enrollment is through partner.
func (m *DeviceManagementPartner) GetGroupsRequiringPartnerEnrollment()([]DeviceManagementPartnerAssignmentable) {
    return m.groupsRequiringPartnerEnrollment
}
// GetIsConfigured gets the isConfigured property value. Whether device management partner is configured or not
func (m *DeviceManagementPartner) GetIsConfigured()(*bool) {
    return m.isConfigured
}
// GetLastHeartbeatDateTime gets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
func (m *DeviceManagementPartner) GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastHeartbeatDateTime
}
// GetPartnerAppType gets the partnerAppType property value. Partner App Type.
func (m *DeviceManagementPartner) GetPartnerAppType()(*DeviceManagementPartnerAppType) {
    return m.partnerAppType
}
// GetPartnerState gets the partnerState property value. Partner state of this tenant.
func (m *DeviceManagementPartner) GetPartnerState()(*DeviceManagementPartnerTenantState) {
    return m.partnerState
}
// GetSingleTenantAppId gets the singleTenantAppId property value. Partner Single tenant App id
func (m *DeviceManagementPartner) GetSingleTenantAppId()(*string) {
    return m.singleTenantAppId
}
// GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime gets the whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime property value. DateTime in UTC when PartnerDevices will be marked as NonCompliant
func (m *DeviceManagementPartner) GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime
}
// GetWhenPartnerDevicesWillBeRemovedDateTime gets the whenPartnerDevicesWillBeRemovedDateTime property value. DateTime in UTC when PartnerDevices will be removed
func (m *DeviceManagementPartner) GetWhenPartnerDevicesWillBeRemovedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.whenPartnerDevicesWillBeRemovedDateTime
}
// Serialize serializes information the current object
func (m *DeviceManagementPartner) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetGroupsRequiringPartnerEnrollment() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroupsRequiringPartnerEnrollment()))
        for i, v := range m.GetGroupsRequiringPartnerEnrollment() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("groupsRequiringPartnerEnrollment", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isConfigured", m.GetIsConfigured())
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
    if m.GetPartnerAppType() != nil {
        cast := (*m.GetPartnerAppType()).String()
        err = writer.WriteStringValue("partnerAppType", &cast)
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
        err = writer.WriteStringValue("singleTenantAppId", m.GetSingleTenantAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime", m.GetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("whenPartnerDevicesWillBeRemovedDateTime", m.GetWhenPartnerDevicesWillBeRemovedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Partner display name
func (m *DeviceManagementPartner) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroupsRequiringPartnerEnrollment sets the groupsRequiringPartnerEnrollment property value. User groups that specifies whether enrollment is through partner.
func (m *DeviceManagementPartner) SetGroupsRequiringPartnerEnrollment(value []DeviceManagementPartnerAssignmentable)() {
    m.groupsRequiringPartnerEnrollment = value
}
// SetIsConfigured sets the isConfigured property value. Whether device management partner is configured or not
func (m *DeviceManagementPartner) SetIsConfigured(value *bool)() {
    m.isConfigured = value
}
// SetLastHeartbeatDateTime sets the lastHeartbeatDateTime property value. Timestamp of last heartbeat after admin enabled option Connect to Device management Partner
func (m *DeviceManagementPartner) SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastHeartbeatDateTime = value
}
// SetPartnerAppType sets the partnerAppType property value. Partner App Type.
func (m *DeviceManagementPartner) SetPartnerAppType(value *DeviceManagementPartnerAppType)() {
    m.partnerAppType = value
}
// SetPartnerState sets the partnerState property value. Partner state of this tenant.
func (m *DeviceManagementPartner) SetPartnerState(value *DeviceManagementPartnerTenantState)() {
    m.partnerState = value
}
// SetSingleTenantAppId sets the singleTenantAppId property value. Partner Single tenant App id
func (m *DeviceManagementPartner) SetSingleTenantAppId(value *string)() {
    m.singleTenantAppId = value
}
// SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime sets the whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime property value. DateTime in UTC when PartnerDevices will be marked as NonCompliant
func (m *DeviceManagementPartner) SetWhenPartnerDevicesWillBeMarkedAsNonCompliantDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.whenPartnerDevicesWillBeMarkedAsNonCompliantDateTime = value
}
// SetWhenPartnerDevicesWillBeRemovedDateTime sets the whenPartnerDevicesWillBeRemovedDateTime property value. DateTime in UTC when PartnerDevices will be removed
func (m *DeviceManagementPartner) SetWhenPartnerDevicesWillBeRemovedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.whenPartnerDevicesWillBeRemovedDateTime = value
}
