package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimulationAutomationRun 
type SimulationAutomationRun struct {
    Entity
    // Date and time when the run ends in an attack simulation automation.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique identifier for the attack simulation campaign initiated in the attack simulation automation run.
    simulationId *string
    // Date and time when the run starts in an attack simulation automation.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Status of the attack simulation automation run. The possible values are: unknown, running, succeeded, failed, skipped, unknownFutureValue.
    status *SimulationAutomationRunStatus
}
// NewSimulationAutomationRun instantiates a new simulationAutomationRun and sets the default values.
func NewSimulationAutomationRun()(*SimulationAutomationRun) {
    m := &SimulationAutomationRun{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSimulationAutomationRunFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimulationAutomationRunFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulationAutomationRun(), nil
}
// GetEndDateTime gets the endDateTime property value. Date and time when the run ends in an attack simulation automation.
func (m *SimulationAutomationRun) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationAutomationRun) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["simulationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimulationId(val)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSimulationAutomationRunStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SimulationAutomationRunStatus))
        }
        return nil
    }
    return res
}
// GetSimulationId gets the simulationId property value. Unique identifier for the attack simulation campaign initiated in the attack simulation automation run.
func (m *SimulationAutomationRun) GetSimulationId()(*string) {
    return m.simulationId
}
// GetStartDateTime gets the startDateTime property value. Date and time when the run starts in an attack simulation automation.
func (m *SimulationAutomationRun) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. Status of the attack simulation automation run. The possible values are: unknown, running, succeeded, failed, skipped, unknownFutureValue.
func (m *SimulationAutomationRun) GetStatus()(*SimulationAutomationRunStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *SimulationAutomationRun) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("simulationId", m.GetSimulationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
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
// SetEndDateTime sets the endDateTime property value. Date and time when the run ends in an attack simulation automation.
func (m *SimulationAutomationRun) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetSimulationId sets the simulationId property value. Unique identifier for the attack simulation campaign initiated in the attack simulation automation run.
func (m *SimulationAutomationRun) SetSimulationId(value *string)() {
    m.simulationId = value
}
// SetStartDateTime sets the startDateTime property value. Date and time when the run starts in an attack simulation automation.
func (m *SimulationAutomationRun) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. Status of the attack simulation automation run. The possible values are: unknown, running, succeeded, failed, skipped, unknownFutureValue.
func (m *SimulationAutomationRun) SetStatus(value *SimulationAutomationRunStatus)() {
    m.status = value
}
