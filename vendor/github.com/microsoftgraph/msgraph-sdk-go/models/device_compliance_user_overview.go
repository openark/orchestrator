package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceUserOverview 
type DeviceComplianceUserOverview struct {
    Entity
    // Version of the policy for that overview
    configurationVersion *int32
    // Number of error Users
    errorCount *int32
    // Number of failed Users
    failedCount *int32
    // Last update time
    lastUpdateDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Number of not applicable users
    notApplicableCount *int32
    // Number of pending Users
    pendingCount *int32
    // Number of succeeded Users
    successCount *int32
}
// NewDeviceComplianceUserOverview instantiates a new deviceComplianceUserOverview and sets the default values.
func NewDeviceComplianceUserOverview()(*DeviceComplianceUserOverview) {
    m := &DeviceComplianceUserOverview{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceComplianceUserOverviewFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceUserOverviewFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceUserOverview(), nil
}
// GetConfigurationVersion gets the configurationVersion property value. Version of the policy for that overview
func (m *DeviceComplianceUserOverview) GetConfigurationVersion()(*int32) {
    return m.configurationVersion
}
// GetErrorCount gets the errorCount property value. Number of error Users
func (m *DeviceComplianceUserOverview) GetErrorCount()(*int32) {
    return m.errorCount
}
// GetFailedCount gets the failedCount property value. Number of failed Users
func (m *DeviceComplianceUserOverview) GetFailedCount()(*int32) {
    return m.failedCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceUserOverview) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configurationVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigurationVersion(val)
        }
        return nil
    }
    res["errorCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCount(val)
        }
        return nil
    }
    res["failedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedCount(val)
        }
        return nil
    }
    res["lastUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
        }
        return nil
    }
    res["notApplicableCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotApplicableCount(val)
        }
        return nil
    }
    res["pendingCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPendingCount(val)
        }
        return nil
    }
    res["successCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessCount(val)
        }
        return nil
    }
    return res
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. Last update time
func (m *DeviceComplianceUserOverview) GetLastUpdateDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdateDateTime
}
// GetNotApplicableCount gets the notApplicableCount property value. Number of not applicable users
func (m *DeviceComplianceUserOverview) GetNotApplicableCount()(*int32) {
    return m.notApplicableCount
}
// GetPendingCount gets the pendingCount property value. Number of pending Users
func (m *DeviceComplianceUserOverview) GetPendingCount()(*int32) {
    return m.pendingCount
}
// GetSuccessCount gets the successCount property value. Number of succeeded Users
func (m *DeviceComplianceUserOverview) GetSuccessCount()(*int32) {
    return m.successCount
}
// Serialize serializes information the current object
func (m *DeviceComplianceUserOverview) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("configurationVersion", m.GetConfigurationVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("errorCount", m.GetErrorCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedCount", m.GetFailedCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notApplicableCount", m.GetNotApplicableCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("pendingCount", m.GetPendingCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("successCount", m.GetSuccessCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConfigurationVersion sets the configurationVersion property value. Version of the policy for that overview
func (m *DeviceComplianceUserOverview) SetConfigurationVersion(value *int32)() {
    m.configurationVersion = value
}
// SetErrorCount sets the errorCount property value. Number of error Users
func (m *DeviceComplianceUserOverview) SetErrorCount(value *int32)() {
    m.errorCount = value
}
// SetFailedCount sets the failedCount property value. Number of failed Users
func (m *DeviceComplianceUserOverview) SetFailedCount(value *int32)() {
    m.failedCount = value
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. Last update time
func (m *DeviceComplianceUserOverview) SetLastUpdateDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdateDateTime = value
}
// SetNotApplicableCount sets the notApplicableCount property value. Number of not applicable users
func (m *DeviceComplianceUserOverview) SetNotApplicableCount(value *int32)() {
    m.notApplicableCount = value
}
// SetPendingCount sets the pendingCount property value. Number of pending Users
func (m *DeviceComplianceUserOverview) SetPendingCount(value *int32)() {
    m.pendingCount = value
}
// SetSuccessCount sets the successCount property value. Number of succeeded Users
func (m *DeviceComplianceUserOverview) SetSuccessCount(value *int32)() {
    m.successCount = value
}
