package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookRangeFill 
type WorkbookRangeFill struct {
    Entity
    // HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange')
    color *string
}
// NewWorkbookRangeFill instantiates a new workbookRangeFill and sets the default values.
func NewWorkbookRangeFill()(*WorkbookRangeFill) {
    m := &WorkbookRangeFill{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWorkbookRangeFillFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookRangeFillFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookRangeFill(), nil
}
// GetColor gets the color property value. HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange')
func (m *WorkbookRangeFill) GetColor()(*string) {
    return m.color
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookRangeFill) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WorkbookRangeFill) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColor sets the color property value. HTML color code representing the color of the border line, of the form #RRGGBB (e.g. 'FFA500') or as a named HTML color (e.g. 'orange')
func (m *WorkbookRangeFill) SetColor(value *string)() {
    m.color = value
}
