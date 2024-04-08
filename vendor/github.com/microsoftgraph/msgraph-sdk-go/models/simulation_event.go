package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SimulationEvent 
type SimulationEvent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Count of the simulation event occurrence in an attack simulation and training campaign.
    count *int32
    // Name of the simulation event in an attack simulation and training campaign.
    eventName *string
    // The OdataType property
    odataType *string
}
// NewSimulationEvent instantiates a new simulationEvent and sets the default values.
func NewSimulationEvent()(*SimulationEvent) {
    m := &SimulationEvent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSimulationEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSimulationEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSimulationEvent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SimulationEvent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCount gets the count property value. Count of the simulation event occurrence in an attack simulation and training campaign.
func (m *SimulationEvent) GetCount()(*int32) {
    return m.count
}
// GetEventName gets the eventName property value. Name of the simulation event in an attack simulation and training campaign.
func (m *SimulationEvent) GetEventName()(*string) {
    return m.eventName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SimulationEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["eventName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventName(val)
        }
        return nil
    }
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SimulationEvent) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SimulationEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("eventName", m.GetEventName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *SimulationEvent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCount sets the count property value. Count of the simulation event occurrence in an attack simulation and training campaign.
func (m *SimulationEvent) SetCount(value *int32)() {
    m.count = value
}
// SetEventName sets the eventName property value. Name of the simulation event in an attack simulation and training campaign.
func (m *SimulationEvent) SetEventName(value *string)() {
    m.eventName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SimulationEvent) SetOdataType(value *string)() {
    m.odataType = value
}
