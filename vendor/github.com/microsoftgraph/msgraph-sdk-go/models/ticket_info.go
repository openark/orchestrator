package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TicketInfo 
type TicketInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The ticket number.
    ticketNumber *string
    // The description of the ticket system.
    ticketSystem *string
}
// NewTicketInfo instantiates a new ticketInfo and sets the default values.
func NewTicketInfo()(*TicketInfo) {
    m := &TicketInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTicketInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTicketInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTicketInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TicketInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TicketInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["ticketNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketNumber(val)
        }
        return nil
    }
    res["ticketSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketSystem(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TicketInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetTicketNumber gets the ticketNumber property value. The ticket number.
func (m *TicketInfo) GetTicketNumber()(*string) {
    return m.ticketNumber
}
// GetTicketSystem gets the ticketSystem property value. The description of the ticket system.
func (m *TicketInfo) GetTicketSystem()(*string) {
    return m.ticketSystem
}
// Serialize serializes information the current object
func (m *TicketInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ticketNumber", m.GetTicketNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ticketSystem", m.GetTicketSystem())
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
func (m *TicketInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TicketInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTicketNumber sets the ticketNumber property value. The ticket number.
func (m *TicketInfo) SetTicketNumber(value *string)() {
    m.ticketNumber = value
}
// SetTicketSystem sets the ticketSystem property value. The description of the ticket system.
func (m *TicketInfo) SetTicketSystem(value *string)() {
    m.ticketSystem = value
}
