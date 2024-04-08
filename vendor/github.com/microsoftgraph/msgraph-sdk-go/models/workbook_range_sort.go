package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRangeSort 
type WorkbookRangeSort struct {
    Entity
}
// NewWorkbookRangeSort instantiates a new workbookRangeSort and sets the default values.
func NewWorkbookRangeSort()(*WorkbookRangeSort) {
    m := &WorkbookRangeSort{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookRangeSortFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookRangeSortFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRangeSort(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRangeSort) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *WorkbookRangeSort) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
