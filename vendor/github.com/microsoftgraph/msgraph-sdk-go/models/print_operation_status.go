package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintOperationStatus 
type PrintOperationStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A human-readable description of the printOperation's current processing state. Read-only.
    description *string
    // The OdataType property
    odataType *string
    // The state property
    state *PrintOperationProcessingState
}
// NewPrintOperationStatus instantiates a new printOperationStatus and sets the default values.
func NewPrintOperationStatus()(*PrintOperationStatus) {
    m := &PrintOperationStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrintOperationStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintOperationStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintOperationStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrintOperationStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. A human-readable description of the printOperation's current processing state. Read-only.
func (m *PrintOperationStatus) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintOperationStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePrintOperationProcessingState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*PrintOperationProcessingState))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrintOperationStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. The state property
func (m *PrintOperationStatus) GetState()(*PrintOperationProcessingState) {
    return m.state
}
// Serialize serializes information the current object
func (m *PrintOperationStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *PrintOperationStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. A human-readable description of the printOperation's current processing state. Read-only.
func (m *PrintOperationStatus) SetDescription(value *string)() {
    m.description = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrintOperationStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. The state property
func (m *PrintOperationStatus) SetState(value *PrintOperationProcessingState)() {
    m.state = value
}
