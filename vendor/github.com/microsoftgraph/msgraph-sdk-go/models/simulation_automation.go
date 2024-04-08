package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimulationAutomation 
type SimulationAutomation struct {
    Entity
    // Identity of the user who created the attack simulation automation.
    createdBy EmailIdentityable
    // Date and time when the attack simulation automation was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description of the attack simulation automation.
    description *string
    // Display name of the attack simulation automation. Supports $filter and $orderby.
    displayName *string
    // Identity of the user who most recently modified the attack simulation automation.
    lastModifiedBy EmailIdentityable
    // Date and time when the attack simulation automation was most recently modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time of the latest run of the attack simulation automation.
    lastRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Date and time of the upcoming run of the attack simulation automation.
    nextRunDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of simulation automation runs.
    runs []SimulationAutomationRunable
    // Status of the attack simulation automation. Supports $filter and $orderby. The possible values are: unknown, draft, notRunning, running, completed, unknownFutureValue.
    status *SimulationAutomationStatus
}
// NewSimulationAutomation instantiates a new simulationAutomation and sets the default values.
func NewSimulationAutomation()(*SimulationAutomation) {
    m := &SimulationAutomation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSimulationAutomationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimulationAutomationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulationAutomation(), nil
}
// GetCreatedBy gets the createdBy property value. Identity of the user who created the attack simulation automation.
func (m *SimulationAutomation) GetCreatedBy()(EmailIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time when the attack simulation automation was created.
func (m *SimulationAutomation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description of the attack simulation automation.
func (m *SimulationAutomation) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name of the attack simulation automation. Supports $filter and $orderby.
func (m *SimulationAutomation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationAutomation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(EmailIdentityable))
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
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(EmailIdentityable))
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
    res["lastRunDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastRunDateTime(val)
        }
        return nil
    }
    res["nextRunDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNextRunDateTime(val)
        }
        return nil
    }
    res["runs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSimulationAutomationRunFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SimulationAutomationRunable, len(val))
            for i, v := range val {
                res[i] = v.(SimulationAutomationRunable)
            }
            m.SetRuns(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAutomationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SimulationAutomationStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation automation.
func (m *SimulationAutomation) GetLastModifiedBy()(EmailIdentityable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time when the attack simulation automation was most recently modified.
func (m *SimulationAutomation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLastRunDateTime gets the lastRunDateTime property value. Date and time of the latest run of the attack simulation automation.
func (m *SimulationAutomation) GetLastRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastRunDateTime
}
// GetNextRunDateTime gets the nextRunDateTime property value. Date and time of the upcoming run of the attack simulation automation.
func (m *SimulationAutomation) GetNextRunDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.nextRunDateTime
}
// GetRuns gets the runs property value. A collection of simulation automation runs.
func (m *SimulationAutomation) GetRuns()([]SimulationAutomationRunable) {
    return m.runs
}
// GetStatus gets the status property value. Status of the attack simulation automation. Supports $filter and $orderby. The possible values are: unknown, draft, notRunning, running, completed, unknownFutureValue.
func (m *SimulationAutomation) GetStatus()(*SimulationAutomationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *SimulationAutomation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
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
    {
        err = writer.WriteTimeValue("lastRunDateTime", m.GetLastRunDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("nextRunDateTime", m.GetNextRunDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRuns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRuns()))
        for i, v := range m.GetRuns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("runs", cast)
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
    return nil
}
// SetCreatedBy sets the createdBy property value. Identity of the user who created the attack simulation automation.
func (m *SimulationAutomation) SetCreatedBy(value EmailIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time when the attack simulation automation was created.
func (m *SimulationAutomation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description of the attack simulation automation.
func (m *SimulationAutomation) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name of the attack simulation automation. Supports $filter and $orderby.
func (m *SimulationAutomation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user who most recently modified the attack simulation automation.
func (m *SimulationAutomation) SetLastModifiedBy(value EmailIdentityable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time when the attack simulation automation was most recently modified.
func (m *SimulationAutomation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLastRunDateTime sets the lastRunDateTime property value. Date and time of the latest run of the attack simulation automation.
func (m *SimulationAutomation) SetLastRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastRunDateTime = value
}
// SetNextRunDateTime sets the nextRunDateTime property value. Date and time of the upcoming run of the attack simulation automation.
func (m *SimulationAutomation) SetNextRunDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.nextRunDateTime = value
}
// SetRuns sets the runs property value. A collection of simulation automation runs.
func (m *SimulationAutomation) SetRuns(value []SimulationAutomationRunable)() {
    m.runs = value
}
// SetStatus sets the status property value. Status of the attack simulation automation. Supports $filter and $orderby. The possible values are: unknown, draft, notRunning, running, completed, unknownFutureValue.
func (m *SimulationAutomation) SetStatus(value *SimulationAutomationStatus)() {
    m.status = value
}
