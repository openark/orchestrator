package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnLink 
type ColumnLink struct {
    Entity
    // The name of the column  in this content type.
    name *string
}
// NewColumnLink instantiates a new columnLink and sets the default values.
func NewColumnLink()(*ColumnLink) {
    m := &ColumnLink{
        Entity: *NewEntity(),
    }
    return m
}
// CreateColumnLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnLinkFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnLink(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnLink) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the column  in this content type.
func (m *ColumnLink) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *ColumnLink) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name of the column  in this content type.
func (m *ColumnLink) SetName(value *string)() {
    m.name = value
}
