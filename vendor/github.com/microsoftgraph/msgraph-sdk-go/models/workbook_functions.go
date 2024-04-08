package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookFunctions 
type WorkbookFunctions struct {
    Entity
}
// NewWorkbookFunctions instantiates a new workbookFunctions and sets the default values.
func NewWorkbookFunctions()(*WorkbookFunctions) {
    m := &WorkbookFunctions{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookFunctionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFunctionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookFunctions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFunctions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WorkbookFunctions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
