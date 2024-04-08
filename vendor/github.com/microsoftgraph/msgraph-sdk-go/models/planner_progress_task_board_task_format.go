package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerProgressTaskBoardTaskFormat 
type PlannerProgressTaskBoardTaskFormat struct {
    Entity
    // Hint value used to order the task on the progress view of the task board. For details about the supported format, see Using order hints in Planner.
    orderHint *string
}
// NewPlannerProgressTaskBoardTaskFormat instantiates a new plannerProgressTaskBoardTaskFormat and sets the default values.
func NewPlannerProgressTaskBoardTaskFormat()(*PlannerProgressTaskBoardTaskFormat) {
    m := &PlannerProgressTaskBoardTaskFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerProgressTaskBoardTaskFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerProgressTaskBoardTaskFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerProgressTaskBoardTaskFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerProgressTaskBoardTaskFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["orderHint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderHint(val)
        }
        return nil
    }
    return res
}
// GetOrderHint gets the orderHint property value. Hint value used to order the task on the progress view of the task board. For details about the supported format, see Using order hints in Planner.
func (m *PlannerProgressTaskBoardTaskFormat) GetOrderHint()(*string) {
    return m.orderHint
}
// Serialize serializes information the current object
func (m *PlannerProgressTaskBoardTaskFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("orderHint", m.GetOrderHint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrderHint sets the orderHint property value. Hint value used to order the task on the progress view of the task board. For details about the supported format, see Using order hints in Planner.
func (m *PlannerProgressTaskBoardTaskFormat) SetOrderHint(value *string)() {
    m.orderHint = value
}
