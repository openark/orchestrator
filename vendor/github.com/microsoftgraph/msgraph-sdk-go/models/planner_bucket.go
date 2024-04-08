package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerBucket 
type PlannerBucket struct {
    Entity
    // Name of the bucket.
    name *string
    // Hint used to order items of this type in a list view. For details about the supported format, see Using order hints in Planner.
    orderHint *string
    // Plan ID to which the bucket belongs.
    planId *string
    // Read-only. Nullable. The collection of tasks in the bucket.
    tasks []PlannerTaskable
}
// NewPlannerBucket instantiates a new plannerBucket and sets the default values.
func NewPlannerBucket()(*PlannerBucket) {
    m := &PlannerBucket{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlannerBucketFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerBucketFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerBucket(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerBucket) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
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
    res["planId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlanId(val)
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePlannerTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PlannerTaskable, len(val))
            for i, v := range val {
                res[i] = v.(PlannerTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Name of the bucket.
func (m *PlannerBucket) GetName()(*string) {
    return m.name
}
// GetOrderHint gets the orderHint property value. Hint used to order items of this type in a list view. For details about the supported format, see Using order hints in Planner.
func (m *PlannerBucket) GetOrderHint()(*string) {
    return m.orderHint
}
// GetPlanId gets the planId property value. Plan ID to which the bucket belongs.
func (m *PlannerBucket) GetPlanId()(*string) {
    return m.planId
}
// GetTasks gets the tasks property value. Read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerBucket) GetTasks()([]PlannerTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *PlannerBucket) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("orderHint", m.GetOrderHint())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("planId", m.GetPlanId())
        if err != nil {
            return err
        }
    }
    if m.GetTasks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. Name of the bucket.
func (m *PlannerBucket) SetName(value *string)() {
    m.name = value
}
// SetOrderHint sets the orderHint property value. Hint used to order items of this type in a list view. For details about the supported format, see Using order hints in Planner.
func (m *PlannerBucket) SetOrderHint(value *string)() {
    m.orderHint = value
}
// SetPlanId sets the planId property value. Plan ID to which the bucket belongs.
func (m *PlannerBucket) SetPlanId(value *string)() {
    m.planId = value
}
// SetTasks sets the tasks property value. Read-only. Nullable. The collection of tasks in the bucket.
func (m *PlannerBucket) SetTasks(value []PlannerTaskable)() {
    m.tasks = value
}
