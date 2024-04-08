package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimulationReport 
type SimulationReport struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Overview of an attack simulation and training campaign.
    overview SimulationReportOverviewable
    // The tenant users and their online actions in an attack simulation and training campaign.
    simulationUsers []UserSimulationDetailsable
}
// NewSimulationReport instantiates a new simulationReport and sets the default values.
func NewSimulationReport()(*SimulationReport) {
    m := &SimulationReport{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSimulationReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimulationReportFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulationReport(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReport) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationReport) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["overview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSimulationReportOverviewFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverview(val.(SimulationReportOverviewable))
        }
        return nil
    }
    res["simulationUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserSimulationDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserSimulationDetailsable, len(val))
            for i, v := range val {
                res[i] = v.(UserSimulationDetailsable)
            }
            m.SetSimulationUsers(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SimulationReport) GetOdataType()(*string) {
    return m.odataType
}
// GetOverview gets the overview property value. Overview of an attack simulation and training campaign.
func (m *SimulationReport) GetOverview()(SimulationReportOverviewable) {
    return m.overview
}
// GetSimulationUsers gets the simulationUsers property value. The tenant users and their online actions in an attack simulation and training campaign.
func (m *SimulationReport) GetSimulationUsers()([]UserSimulationDetailsable) {
    return m.simulationUsers
}
// Serialize serializes information the current object
func (m *SimulationReport) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("overview", m.GetOverview())
        if err != nil {
            return err
        }
    }
    if m.GetSimulationUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSimulationUsers()))
        for i, v := range m.GetSimulationUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("simulationUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationReport) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SimulationReport) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOverview sets the overview property value. Overview of an attack simulation and training campaign.
func (m *SimulationReport) SetOverview(value SimulationReportOverviewable)() {
    m.overview = value
}
// SetSimulationUsers sets the simulationUsers property value. The tenant users and their online actions in an attack simulation and training campaign.
func (m *SimulationReport) SetSimulationUsers(value []UserSimulationDetailsable)() {
    m.simulationUsers = value
}
