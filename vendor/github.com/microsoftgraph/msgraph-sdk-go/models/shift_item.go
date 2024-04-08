package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ShiftItem 
type ShiftItem struct {
    ScheduleEntity
    // An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
    activities []ShiftActivityable
    // The shift label of the shiftItem.
    displayName *string
    // The shift notes for the shiftItem.
    notes *string
}
// NewShiftItem instantiates a new ShiftItem and sets the default values.
func NewShiftItem()(*ShiftItem) {
    m := &ShiftItem{
        ScheduleEntity: *NewScheduleEntity(),
    }
    return m
}
// CreateShiftItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateShiftItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.openShiftItem":
                        return NewOpenShiftItem(), nil
                }
            }
        }
    }
    return NewShiftItem(), nil
}
// GetActivities gets the activities property value. An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
func (m *ShiftItem) GetActivities()([]ShiftActivityable) {
    return m.activities
}
// GetDisplayName gets the displayName property value. The shift label of the shiftItem.
func (m *ShiftItem) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ShiftItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ScheduleEntity.GetFieldDeserializers()
    res["activities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateShiftActivityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ShiftActivityable, len(val))
            for i, v := range val {
                res[i] = v.(ShiftActivityable)
            }
            m.SetActivities(res)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    return res
}
// GetNotes gets the notes property value. The shift notes for the shiftItem.
func (m *ShiftItem) GetNotes()(*string) {
    return m.notes
}
// Serialize serializes information the current object
func (m *ShiftItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ScheduleEntity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetActivities()))
        for i, v := range m.GetActivities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("activities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivities sets the activities property value. An incremental part of a shift which can cover details of when and where an employee is during their shift. For example, an assignment or a scheduled break or lunch. Required.
func (m *ShiftItem) SetActivities(value []ShiftActivityable)() {
    m.activities = value
}
// SetDisplayName sets the displayName property value. The shift label of the shiftItem.
func (m *ShiftItem) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetNotes sets the notes property value. The shift notes for the shiftItem.
func (m *ShiftItem) SetNotes(value *string)() {
    m.notes = value
}
