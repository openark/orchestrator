package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Approval 
type Approval struct {
    Entity
    // A collection of stages in the approval decision.
    stages []ApprovalStageable
}
// NewApproval instantiates a new approval and sets the default values.
func NewApproval()(*Approval) {
    m := &Approval{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApproval(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Approval) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["stages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalStageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ApprovalStageable, len(val))
            for i, v := range val {
                res[i] = v.(ApprovalStageable)
            }
            m.SetStages(res)
        }
        return nil
    }
    return res
}
// GetStages gets the stages property value. A collection of stages in the approval decision.
func (m *Approval) GetStages()([]ApprovalStageable) {
    return m.stages
}
// Serialize serializes information the current object
func (m *Approval) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetStages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStages()))
        for i, v := range m.GetStages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("stages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetStages sets the stages property value. A collection of stages in the approval decision.
func (m *Approval) SetStages(value []ApprovalStageable)() {
    m.stages = value
}
