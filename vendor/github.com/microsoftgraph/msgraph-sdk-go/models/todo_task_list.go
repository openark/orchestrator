package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TodoTaskList 
type TodoTaskList struct {
    Entity
    // The name of the task list.
    displayName *string
    // The collection of open extensions defined for the task list. Nullable.
    extensions []Extensionable
    // True if the user is owner of the given task list.
    isOwner *bool
    // True if the task list is shared with other users
    isShared *bool
    // The tasks in this task list. Read-only. Nullable.
    tasks []TodoTaskable
    // The wellknownListName property
    wellknownListName *WellknownListName
}
// NewTodoTaskList instantiates a new todoTaskList and sets the default values.
func NewTodoTaskList()(*TodoTaskList) {
    m := &TodoTaskList{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTodoTaskListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTodoTaskListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTodoTaskList(), nil
}
// GetDisplayName gets the displayName property value. The name of the task list.
func (m *TodoTaskList) GetDisplayName()(*string) {
    return m.displayName
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *TodoTaskList) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TodoTaskList) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["isOwner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOwner(val)
        }
        return nil
    }
    res["isShared"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsShared(val)
        }
        return nil
    }
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTodoTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TodoTaskable, len(val))
            for i, v := range val {
                res[i] = v.(TodoTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    res["wellknownListName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWellknownListName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWellknownListName(val.(*WellknownListName))
        }
        return nil
    }
    return res
}
// GetIsOwner gets the isOwner property value. True if the user is owner of the given task list.
func (m *TodoTaskList) GetIsOwner()(*bool) {
    return m.isOwner
}
// GetIsShared gets the isShared property value. True if the task list is shared with other users
func (m *TodoTaskList) GetIsShared()(*bool) {
    return m.isShared
}
// GetTasks gets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *TodoTaskList) GetTasks()([]TodoTaskable) {
    return m.tasks
}
// GetWellknownListName gets the wellknownListName property value. The wellknownListName property
func (m *TodoTaskList) GetWellknownListName()(*WellknownListName) {
    return m.wellknownListName
}
// Serialize serializes information the current object
func (m *TodoTaskList) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOwner", m.GetIsOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isShared", m.GetIsShared())
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
    if m.GetWellknownListName() != nil {
        cast := (*m.GetWellknownListName()).String()
        err = writer.WriteStringValue("wellknownListName", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the task list.
func (m *TodoTaskList) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the task list. Nullable.
func (m *TodoTaskList) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetIsOwner sets the isOwner property value. True if the user is owner of the given task list.
func (m *TodoTaskList) SetIsOwner(value *bool)() {
    m.isOwner = value
}
// SetIsShared sets the isShared property value. True if the task list is shared with other users
func (m *TodoTaskList) SetIsShared(value *bool)() {
    m.isShared = value
}
// SetTasks sets the tasks property value. The tasks in this task list. Read-only. Nullable.
func (m *TodoTaskList) SetTasks(value []TodoTaskable)() {
    m.tasks = value
}
// SetWellknownListName sets the wellknownListName property value. The wellknownListName property
func (m *TodoTaskList) SetWellknownListName(value *WellknownListName)() {
    m.wellknownListName = value
}
