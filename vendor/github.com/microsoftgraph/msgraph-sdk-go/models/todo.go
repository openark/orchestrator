package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Todo 
type Todo struct {
    Entity
    // The task lists in the users mailbox.
    lists []TodoTaskListable
}
// NewTodo instantiates a new todo and sets the default values.
func NewTodo()(*Todo) {
    m := &Todo{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTodoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTodoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTodo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Todo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTodoTaskListFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TodoTaskListable, len(val))
            for i, v := range val {
                res[i] = v.(TodoTaskListable)
            }
            m.SetLists(res)
        }
        return nil
    }
    return res
}
// GetLists gets the lists property value. The task lists in the users mailbox.
func (m *Todo) GetLists()([]TodoTaskListable) {
    return m.lists
}
// Serialize serializes information the current object
func (m *Todo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetLists() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLists()))
        for i, v := range m.GetLists() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("lists", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLists sets the lists property value. The task lists in the users mailbox.
func (m *Todo) SetLists(value []TodoTaskListable)() {
    m.lists = value
}
