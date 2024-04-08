package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerAssignedToTaskBoardTaskFormat 
type PlannerAssignedToTaskBoardTaskFormat struct {
    Entity
    // Dictionary of hints used to order tasks on the AssignedTo view of the Task Board. The key of each entry is one of the users the task is assigned to and the value is the order hint. The format of each value is defined as outlined here.
    orderHintsByAssignee PlannerOrderHintsByAssigneeable
    // Hint value used to order the task on the AssignedTo view of the Task Board when the task is not assigned to anyone, or if the orderHintsByAssignee dictionary does not provide an order hint for the user the task is assigned to. The format is defined as outlined here.
    unassignedOrderHint *string
}
// NewPlannerAssignedToTaskBoardTaskFormat instantiates a new plannerAssignedToTaskBoardTaskFormat and sets the default values.
func NewPlannerAssignedToTaskBoardTaskFormat()(*PlannerAssignedToTaskBoardTaskFormat) {
    m := &PlannerAssignedToTaskBoardTaskFormat{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerAssignedToTaskBoardTaskFormatFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerAssignedToTaskBoardTaskFormatFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerAssignedToTaskBoardTaskFormat(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerAssignedToTaskBoardTaskFormat) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["orderHintsByAssignee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePlannerOrderHintsByAssigneeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrderHintsByAssignee(val.(PlannerOrderHintsByAssigneeable))
        }
        return nil
    }
    res["unassignedOrderHint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnassignedOrderHint(val)
        }
        return nil
    }
    return res
}
// GetOrderHintsByAssignee gets the orderHintsByAssignee property value. Dictionary of hints used to order tasks on the AssignedTo view of the Task Board. The key of each entry is one of the users the task is assigned to and the value is the order hint. The format of each value is defined as outlined here.
func (m *PlannerAssignedToTaskBoardTaskFormat) GetOrderHintsByAssignee()(PlannerOrderHintsByAssigneeable) {
    return m.orderHintsByAssignee
}
// GetUnassignedOrderHint gets the unassignedOrderHint property value. Hint value used to order the task on the AssignedTo view of the Task Board when the task is not assigned to anyone, or if the orderHintsByAssignee dictionary does not provide an order hint for the user the task is assigned to. The format is defined as outlined here.
func (m *PlannerAssignedToTaskBoardTaskFormat) GetUnassignedOrderHint()(*string) {
    return m.unassignedOrderHint
}
// Serialize serializes information the current object
func (m *PlannerAssignedToTaskBoardTaskFormat) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("orderHintsByAssignee", m.GetOrderHintsByAssignee())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("unassignedOrderHint", m.GetUnassignedOrderHint())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetOrderHintsByAssignee sets the orderHintsByAssignee property value. Dictionary of hints used to order tasks on the AssignedTo view of the Task Board. The key of each entry is one of the users the task is assigned to and the value is the order hint. The format of each value is defined as outlined here.
func (m *PlannerAssignedToTaskBoardTaskFormat) SetOrderHintsByAssignee(value PlannerOrderHintsByAssigneeable)() {
    m.orderHintsByAssignee = value
}
// SetUnassignedOrderHint sets the unassignedOrderHint property value. Hint value used to order the task on the AssignedTo view of the Task Board when the task is not assigned to anyone, or if the orderHintsByAssignee dictionary does not provide an order hint for the user the task is assigned to. The format is defined as outlined here.
func (m *PlannerAssignedToTaskBoardTaskFormat) SetUnassignedOrderHint(value *string)() {
    m.unassignedOrderHint = value
}
