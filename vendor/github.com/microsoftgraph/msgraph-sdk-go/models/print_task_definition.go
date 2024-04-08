package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintTaskDefinition 
type PrintTaskDefinition struct {
    Entity
    // The createdBy property
    createdBy AppIdentityable
    // The name of the printTaskDefinition.
    displayName *string
    // A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
    tasks []PrintTaskable
}
// NewPrintTaskDefinition instantiates a new printTaskDefinition and sets the default values.
func NewPrintTaskDefinition()(*PrintTaskDefinition) {
    m := &PrintTaskDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintTaskDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintTaskDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintTaskDefinition(), nil
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *PrintTaskDefinition) GetCreatedBy()(AppIdentityable) {
    return m.createdBy
}
// GetDisplayName gets the displayName property value. The name of the printTaskDefinition.
func (m *PrintTaskDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintTaskDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(AppIdentityable))
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
    res["tasks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintTaskFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTaskable, len(val))
            for i, v := range val {
                res[i] = v.(PrintTaskable)
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
// GetTasks gets the tasks property value. A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
func (m *PrintTaskDefinition) GetTasks()([]PrintTaskable) {
    return m.tasks
}
// Serialize serializes information the current object
func (m *PrintTaskDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *PrintTaskDefinition) SetCreatedBy(value AppIdentityable)() {
    m.createdBy = value
}
// SetDisplayName sets the displayName property value. The name of the printTaskDefinition.
func (m *PrintTaskDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTasks sets the tasks property value. A list of tasks that have been created based on this definition. The list includes currently running tasks and recently completed tasks. Read-only.
func (m *PrintTaskDefinition) SetTasks(value []PrintTaskable)() {
    m.tasks = value
}
