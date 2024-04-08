package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SchedulingGroup 
type SchedulingGroup struct {
    ChangeTrackedEntity
    // The display name for the schedulingGroup. Required.
    displayName *string
    // Indicates whether the schedulingGroup can be used when creating new entities or updating existing ones. Required.
    isActive *bool
    // The list of user IDs that are a member of the schedulingGroup. Required.
    userIds []string
}
// NewSchedulingGroup instantiates a new SchedulingGroup and sets the default values.
func NewSchedulingGroup()(*SchedulingGroup) {
    m := &SchedulingGroup{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.schedulingGroup"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSchedulingGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSchedulingGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSchedulingGroup(), nil
}
// GetDisplayName gets the displayName property value. The display name for the schedulingGroup. Required.
func (m *SchedulingGroup) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SchedulingGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
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
    res["isActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsActive(val)
        }
        return nil
    }
    res["userIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUserIds(res)
        }
        return nil
    }
    return res
}
// GetIsActive gets the isActive property value. Indicates whether the schedulingGroup can be used when creating new entities or updating existing ones. Required.
func (m *SchedulingGroup) GetIsActive()(*bool) {
    return m.isActive
}
// GetUserIds gets the userIds property value. The list of user IDs that are a member of the schedulingGroup. Required.
func (m *SchedulingGroup) GetUserIds()([]string) {
    return m.userIds
}
// Serialize serializes information the current object
func (m *SchedulingGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetUserIds() != nil {
        err = writer.WriteCollectionOfStringValues("userIds", m.GetUserIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name for the schedulingGroup. Required.
func (m *SchedulingGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsActive sets the isActive property value. Indicates whether the schedulingGroup can be used when creating new entities or updating existing ones. Required.
func (m *SchedulingGroup) SetIsActive(value *bool)() {
    m.isActive = value
}
// SetUserIds sets the userIds property value. The list of user IDs that are a member of the schedulingGroup. Required.
func (m *SchedulingGroup) SetUserIds(value []string)() {
    m.userIds = value
}
