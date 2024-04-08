package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OutlookCategory 
type OutlookCategory struct {
    Entity
    // A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
    color *CategoryColor
    // A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
    displayName *string
}
// NewOutlookCategory instantiates a new outlookCategory and sets the default values.
func NewOutlookCategory()(*OutlookCategory) {
    m := &OutlookCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOutlookCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOutlookCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOutlookCategory(), nil
}
// GetColor gets the color property value. A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
func (m *OutlookCategory) GetColor()(*CategoryColor) {
    return m.color
}
// GetDisplayName gets the displayName property value. A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
func (m *OutlookCategory) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OutlookCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCategoryColor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*CategoryColor))
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
    return res
}
// Serialize serializes information the current object
func (m *OutlookCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColor() != nil {
        cast := (*m.GetColor()).String()
        err = writer.WriteStringValue("color", &cast)
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
    return nil
}
// SetColor sets the color property value. A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below.
func (m *OutlookCategory) SetColor(value *CategoryColor)() {
    m.color = value
}
// SetDisplayName sets the displayName property value. A unique name that identifies a category in the user's mailbox. After a category is created, the name cannot be changed. Read-only.
func (m *OutlookCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
