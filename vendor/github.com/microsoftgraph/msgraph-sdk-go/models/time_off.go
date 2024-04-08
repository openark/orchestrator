package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeOff 
type TimeOff struct {
    ChangeTrackedEntity
    // The draft version of this timeOff that is viewable by managers. Required.
    draftTimeOff TimeOffItemable
    // The shared version of this timeOff that is viewable by both employees and managers. Required.
    sharedTimeOff TimeOffItemable
    // ID of the user assigned to the timeOff. Required.
    userId *string
}
// NewTimeOff instantiates a new TimeOff and sets the default values.
func NewTimeOff()(*TimeOff) {
    m := &TimeOff{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.timeOff"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTimeOffFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeOffFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeOff(), nil
}
// GetDraftTimeOff gets the draftTimeOff property value. The draft version of this timeOff that is viewable by managers. Required.
func (m *TimeOff) GetDraftTimeOff()(TimeOffItemable) {
    return m.draftTimeOff
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeOff) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["draftTimeOff"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeOffItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDraftTimeOff(val.(TimeOffItemable))
        }
        return nil
    }
    res["sharedTimeOff"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeOffItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharedTimeOff(val.(TimeOffItemable))
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
// GetSharedTimeOff gets the sharedTimeOff property value. The shared version of this timeOff that is viewable by both employees and managers. Required.
func (m *TimeOff) GetSharedTimeOff()(TimeOffItemable) {
    return m.sharedTimeOff
}
// GetUserId gets the userId property value. ID of the user assigned to the timeOff. Required.
func (m *TimeOff) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *TimeOff) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("draftTimeOff", m.GetDraftTimeOff())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharedTimeOff", m.GetSharedTimeOff())
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
// SetDraftTimeOff sets the draftTimeOff property value. The draft version of this timeOff that is viewable by managers. Required.
func (m *TimeOff) SetDraftTimeOff(value TimeOffItemable)() {
    m.draftTimeOff = value
}
// SetSharedTimeOff sets the sharedTimeOff property value. The shared version of this timeOff that is viewable by both employees and managers. Required.
func (m *TimeOff) SetSharedTimeOff(value TimeOffItemable)() {
    m.sharedTimeOff = value
}
// SetUserId sets the userId property value. ID of the user assigned to the timeOff. Required.
func (m *TimeOff) SetUserId(value *string)() {
    m.userId = value
}
