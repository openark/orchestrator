package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationDeviceStateSummary 
type DeviceConfigurationDeviceStateSummary struct {
    Entity
    // Number of compliant devices
    compliantDeviceCount *int32
    // Number of conflict devices
    conflictDeviceCount *int32
    // Number of error devices
    errorDeviceCount *int32
    // Number of NonCompliant devices
    nonCompliantDeviceCount *int32
    // Number of not applicable devices
    notApplicableDeviceCount *int32
    // Number of remediated devices
    remediatedDeviceCount *int32
    // Number of unknown devices
    unknownDeviceCount *int32
}
// NewDeviceConfigurationDeviceStateSummary instantiates a new deviceConfigurationDeviceStateSummary and sets the default values.
func NewDeviceConfigurationDeviceStateSummary()(*DeviceConfigurationDeviceStateSummary) {
    m := &DeviceConfigurationDeviceStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationDeviceStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationDeviceStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationDeviceStateSummary(), nil
}
// GetCompliantDeviceCount gets the compliantDeviceCount property value. Number of compliant devices
func (m *DeviceConfigurationDeviceStateSummary) GetCompliantDeviceCount()(*int32) {
    return m.compliantDeviceCount
}
// GetConflictDeviceCount gets the conflictDeviceCount property value. Number of conflict devices
func (m *DeviceConfigurationDeviceStateSummary) GetConflictDeviceCount()(*int32) {
    return m.conflictDeviceCount
}
// GetErrorDeviceCount gets the errorDeviceCount property value. Number of error devices
func (m *DeviceConfigurationDeviceStateSummary) GetErrorDeviceCount()(*int32) {
    return m.errorDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationDeviceStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetNonCompliantDeviceCount gets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *DeviceConfigurationDeviceStateSummary) GetNonCompliantDeviceCount()(*int32) {
    return m.nonCompliantDeviceCount
}
// GetNotApplicableDeviceCount gets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *DeviceConfigurationDeviceStateSummary) GetNotApplicableDeviceCount()(*int32) {
    return m.notApplicableDeviceCount
}
// GetRemediatedDeviceCount gets the remediatedDeviceCount property value. Number of remediated devices
func (m *DeviceConfigurationDeviceStateSummary) GetRemediatedDeviceCount()(*int32) {
    return m.remediatedDeviceCount
}
// GetUnknownDeviceCount gets the unknownDeviceCount property value. Number of unknown devices
func (m *DeviceConfigurationDeviceStateSummary) GetUnknownDeviceCount()(*int32) {
    return m.unknownDeviceCount
}
// Serialize serializes information the current object
func (m *DeviceConfigurationDeviceStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteInt32Value("conflictDeviceCount", m.GetConflictDeviceCount())
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
        err = writer.WriteInt32Value("nonCompliantDeviceCount", m.GetNonCompliantDeviceCount())
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
        err = writer.WriteInt32Value("remediatedDeviceCount", m.GetRemediatedDeviceCount())
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
    return nil
}
// SetCompliantDeviceCount sets the compliantDeviceCount property value. Number of compliant devices
func (m *DeviceConfigurationDeviceStateSummary) SetCompliantDeviceCount(value *int32)() {
    m.compliantDeviceCount = value
}
// SetConflictDeviceCount sets the conflictDeviceCount property value. Number of conflict devices
func (m *DeviceConfigurationDeviceStateSummary) SetConflictDeviceCount(value *int32)() {
    m.conflictDeviceCount = value
}
// SetErrorDeviceCount sets the errorDeviceCount property value. Number of error devices
func (m *DeviceConfigurationDeviceStateSummary) SetErrorDeviceCount(value *int32)() {
    m.errorDeviceCount = value
}
// SetNonCompliantDeviceCount sets the nonCompliantDeviceCount property value. Number of NonCompliant devices
func (m *DeviceConfigurationDeviceStateSummary) SetNonCompliantDeviceCount(value *int32)() {
    m.nonCompliantDeviceCount = value
}
// SetNotApplicableDeviceCount sets the notApplicableDeviceCount property value. Number of not applicable devices
func (m *DeviceConfigurationDeviceStateSummary) SetNotApplicableDeviceCount(value *int32)() {
    m.notApplicableDeviceCount = value
}
// SetRemediatedDeviceCount sets the remediatedDeviceCount property value. Number of remediated devices
func (m *DeviceConfigurationDeviceStateSummary) SetRemediatedDeviceCount(value *int32)() {
    m.remediatedDeviceCount = value
}
// SetUnknownDeviceCount sets the unknownDeviceCount property value. Number of unknown devices
func (m *DeviceConfigurationDeviceStateSummary) SetUnknownDeviceCount(value *int32)() {
    m.unknownDeviceCount = value
}
