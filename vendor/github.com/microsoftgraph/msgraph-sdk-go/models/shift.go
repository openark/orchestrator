package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Shift 
type Shift struct {
    ChangeTrackedEntity
    // The draft version of this shift that is viewable by managers. Required.
    draftShift ShiftItemable
    // ID of the scheduling group the shift is part of. Required.
    schedulingGroupId *string
    // The shared version of this shift that is viewable by both employees and managers. Required.
    sharedShift ShiftItemable
    // ID of the user assigned to the shift. Required.
    userId *string
}
// NewShift instantiates a new Shift and sets the default values.
func NewShift()(*Shift) {
    m := &Shift{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.shift"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateShiftFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateShiftFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewShift(), nil
}
// GetDraftShift gets the draftShift property value. The draft version of this shift that is viewable by managers. Required.
func (m *Shift) GetDraftShift()(ShiftItemable) {
    return m.draftShift
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Shift) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftShift(val.(ShiftItemable))
        }
        return nil
    }
    res["schedulingGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchedulingGroupId(val)
        }
        return nil
    }
    res["sharedShift"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateShiftItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedShift(val.(ShiftItemable))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetSchedulingGroupId gets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
func (m *Shift) GetSchedulingGroupId()(*string) {
    return m.schedulingGroupId
}
// GetSharedShift gets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Required.
func (m *Shift) GetSharedShift()(ShiftItemable) {
    return m.sharedShift
}
// GetUserId gets the userId property value. ID of the user assigned to the shift. Required.
func (m *Shift) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *Shift) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftShift", m.GetDraftShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("schedulingGroupId", m.GetSchedulingGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedShift", m.GetSharedShift())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDraftShift sets the draftShift property value. The draft version of this shift that is viewable by managers. Required.
func (m *Shift) SetDraftShift(value ShiftItemable)() {
    m.draftShift = value
}
// SetSchedulingGroupId sets the schedulingGroupId property value. ID of the scheduling group the shift is part of. Required.
func (m *Shift) SetSchedulingGroupId(value *string)() {
    m.schedulingGroupId = value
}
// SetSharedShift sets the sharedShift property value. The shared version of this shift that is viewable by both employees and managers. Required.
func (m *Shift) SetSharedShift(value ShiftItemable)() {
    m.sharedShift = value
}
// SetUserId sets the userId property value. ID of the user assigned to the shift. Required.
func (m *Shift) SetUserId(value *string)() {
    m.userId = value
}
