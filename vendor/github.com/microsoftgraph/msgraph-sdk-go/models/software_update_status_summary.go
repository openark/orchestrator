package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SoftwareUpdateStatusSummary 
type SoftwareUpdateStatusSummary struct {
    Entity
    // Number of compliant devices.
    compliantDeviceCount *int32
    // Number of compliant users.
    compliantUserCount *int32
    // Number of conflict devices.
    conflictDeviceCount *int32
    // Number of conflict users.
    conflictUserCount *int32
    // The name of the policy.
    displayName *string
    // Number of devices had error.
    errorDeviceCount *int32
    // Number of users had error.
    errorUserCount *int32
    // Number of non compliant devices.
    nonCompliantDeviceCount *int32
    // Number of non compliant users.
    nonCompliantUserCount *int32
    // Number of not applicable devices.
    notApplicableDeviceCount *int32
    // Number of not applicable users.
    notApplicableUserCount *int32
    // Number of remediated devices.
    remediatedDeviceCount *int32
    // Number of remediated users.
    remediatedUserCount *int32
    // Number of unknown devices.
    unknownDeviceCount *int32
    // Number of unknown users.
    unknownUserCount *int32
}
// NewSoftwareUpdateStatusSummary instantiates a new softwareUpdateStatusSummary and sets the default values.
func NewSoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummary) {
    m := &SoftwareUpdateStatusSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSoftwareUpdateStatusSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSoftwareUpdateStatusSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSoftwareUpdateStatusSummary(), nil
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. Number of compliant devices.
func (m *SoftwareUpdateStatusSummary) GetCompliantDeviceCount()(*int32) {
    return m.compliantDeviceCount
}
// GetCompliantUserCount gets the compliantUserCount property value. Number of compliant users.
func (m *SoftwareUpdateStatusSummary) GetCompliantUserCount()(*int32) {
    return m.compliantUserCount
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. Number of conflict devices.
func (m *SoftwareUpdateStatusSummary) GetConflictDeviceCount()(*int32) {
    return m.conflictDeviceCount
}
// GetConflictUserCount gets the conflictUserCount property value. Number of conflict users.
func (m *SoftwareUpdateStatusSummary) GetConflictUserCount()(*int32) {
    return m.conflictUserCount
}
// GetDisplayName gets the displayName property value. The name of the policy.
func (m *SoftwareUpdateStatusSummary) GetDisplayName()(*string) {
    return m.displayName
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Number of devices had error.
func (m *SoftwareUpdateStatusSummary) GetErrorDeviceCount()(*int32) {
    return m.errorDeviceCount
}
// GetErrorUserCount gets the errorUserCount property value. Number of users had error.
func (m *SoftwareUpdateStatusSummary) GetErrorUserCount()(*int32) {
    return m.errorUserCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SoftwareUpdateStatusSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["compliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantDeviceCount(val)
        }
        return nil
    }
    res["compliantUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompliantUserCount(val)
        }
        return nil
    }
    res["conflictDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictDeviceCount(val)
        }
        return nil
    }
    res["conflictUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConflictUserCount(val)
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
    res["errorDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorDeviceCount(val)
        }
        return nil
    }
    res["errorUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorUserCount(val)
        }
        return nil
    }
    res["nonCompliantDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantDeviceCount(val)
        }
        return nil
    }
    res["nonCompliantUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNonCompliantUserCount(val)
        }
        return nil
    }
    res["notApplicableDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableDeviceCount(val)
        }
        return nil
    }
    res["notApplicableUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableUserCount(val)
        }
        return nil
    }
    res["remediatedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedDeviceCount(val)
        }
        return nil
    }
    res["remediatedUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediatedUserCount(val)
        }
        return nil
    }
    res["unknownDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownDeviceCount(val)
        }
        return nil
    }
    res["unknownUserCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownUserCount(val)
        }
        return nil
    }
    return res
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. Number of non compliant devices.
func (m *SoftwareUpdateStatusSummary) GetNonCompliantDeviceCount()(*int32) {
    return m.nonCompliantDeviceCount
}
// GetNonCompliantUserCount gets the nonCompliantUserCount property value. Number of non compliant users.
func (m *SoftwareUpdateStatusSummary) GetNonCompliantUserCount()(*int32) {
    return m.nonCompliantUserCount
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of not applicable devices.
func (m *SoftwareUpdateStatusSummary) GetNotApplicableDeviceCount()(*int32) {
    return m.notApplicableDeviceCount
}
// GetNotApplicableUserCount gets the notApplicableUserCount property value. Number of not applicable users.
func (m *SoftwareUpdateStatusSummary) GetNotApplicableUserCount()(*int32) {
    return m.notApplicableUserCount
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. Number of remediated devices.
func (m *SoftwareUpdateStatusSummary) GetRemediatedDeviceCount()(*int32) {
    return m.remediatedDeviceCount
}
// GetRemediatedUserCount gets the remediatedUserCount property value. Number of remediated users.
func (m *SoftwareUpdateStatusSummary) GetRemediatedUserCount()(*int32) {
    return m.remediatedUserCount
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Number of unknown devices.
func (m *SoftwareUpdateStatusSummary) GetUnknownDeviceCount()(*int32) {
    return m.unknownDeviceCount
}
// GetUnknownUserCount gets the unknownUserCount property value. Number of unknown users.
func (m *SoftwareUpdateStatusSummary) GetUnknownUserCount()(*int32) {
    return m.unknownUserCount
}
// Serialize serializes information the current object
func (m *SoftwareUpdateStatusSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("compliantDeviceCount", m.GetCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("compliantUserCount", m.GetCompliantUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("conflictDeviceCount", m.GetConflictDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("conflictUserCount", m.GetConflictUserCount())
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
        err = writer.WriteInt32Value("errorDeviceCount", m.GetErrorDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorUserCount", m.GetErrorUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("nonCompliantDeviceCount", m.GetNonCompliantDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("nonCompliantUserCount", m.GetNonCompliantUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableDeviceCount", m.GetNotApplicableDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableUserCount", m.GetNotApplicableUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("remediatedUserCount", m.GetRemediatedUserCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownDeviceCount", m.GetUnknownDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("unknownUserCount", m.GetUnknownUserCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. Number of compliant devices.
func (m *SoftwareUpdateStatusSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// SetCompliantUserCount sets the compliantUserCount property value. Number of compliant users.
func (m *SoftwareUpdateStatusSummary) SetCompliantUserCount(value *int32)() {
    m.compliantUserCount = value
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. Number of conflict devices.
func (m *SoftwareUpdateStatusSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
// SetConflictUserCount sets the conflictUserCount property value. Number of conflict users.
func (m *SoftwareUpdateStatusSummary) SetConflictUserCount(value *int32)() {
    m.conflictUserCount = value
}
// SetDisplayName sets the displayName property value. The name of the policy.
func (m *SoftwareUpdateStatusSummary) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Number of devices had error.
func (m *SoftwareUpdateStatusSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// SetErrorUserCount sets the errorUserCount property value. Number of users had error.
func (m *SoftwareUpdateStatusSummary) SetErrorUserCount(value *int32)() {
    m.errorUserCount = value
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. Number of non compliant devices.
func (m *SoftwareUpdateStatusSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// SetNonCompliantUserCount sets the nonCompliantUserCount property value. Number of non compliant users.
func (m *SoftwareUpdateStatusSummary) SetNonCompliantUserCount(value *int32)() {
    m.nonCompliantUserCount = value
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of not applicable devices.
func (m *SoftwareUpdateStatusSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// SetNotApplicableUserCount sets the notApplicableUserCount property value. Number of not applicable users.
func (m *SoftwareUpdateStatusSummary) SetNotApplicableUserCount(value *int32)() {
    m.notApplicableUserCount = value
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. Number of remediated devices.
func (m *SoftwareUpdateStatusSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// SetRemediatedUserCount sets the remediatedUserCount property value. Number of remediated users.
func (m *SoftwareUpdateStatusSummary) SetRemediatedUserCount(value *int32)() {
    m.remediatedUserCount = value
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Number of unknown devices.
func (m *SoftwareUpdateStatusSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
// SetUnknownUserCount sets the unknownUserCount property value. Number of unknown users.
func (m *SoftwareUpdateStatusSummary) SetUnknownUserCount(value *int32)() {
    m.unknownUserCount = value
}
