package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeOffItem 
type TimeOffItem struct {
    ScheduleEntity
    // ID of the timeOffReason for this timeOffItem. Required.
    timeOffReasonId *string
}
// NewTimeOffItem instantiates a new TimeOffItem and sets the default values.
func NewTimeOffItem()(*TimeOffItem) {
    m := &TimeOffItem{
        ScheduleEntity: *NewScheduleEntity(),
    }
    return m
}
// CreateTimeOffItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeOffItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeOffItem(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeOffItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ScheduleEntity.GetFieldDeserializers()
    res["timeOffReasonId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeOffReasonId(val)
        }
        return nil
    }
    return res
}
// GetTimeOffReasonId gets the timeOffReasonId property value. ID of the timeOffReason for this timeOffItem. Required.
func (m *TimeOffItem) GetTimeOffReasonId()(*string) {
    return m.timeOffReasonId
}
// Serialize serializes information the current object
func (m *TimeOffItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ScheduleEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("timeOffReasonId", m.GetTimeOffReasonId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTimeOffReasonId sets the timeOffReasonId property value. ID of the timeOffReason for this timeOffItem. Required.
func (m *TimeOffItem) SetTimeOffReasonId(value *string)() {
    m.timeOffReasonId = value
}
