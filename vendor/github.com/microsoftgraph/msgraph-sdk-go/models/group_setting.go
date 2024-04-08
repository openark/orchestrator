package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupSetting 
type GroupSetting struct {
    Entity
    // Display name of this group of settings, which comes from the associated template.
    displayName *string
    // Unique identifier for the tenant-level groupSettingTemplates object that's been customized for this group-level settings object. Read-only.
    templateId *string
    // Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced groupSettingTemplates object.
    values []SettingValueable
}
// NewGroupSetting instantiates a new groupSetting and sets the default values.
func NewGroupSetting()(*GroupSetting) {
    m := &GroupSetting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateGroupSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupSetting(), nil
}
// GetDisplayName gets the displayName property value. Display name of this group of settings, which comes from the associated template.
func (m *GroupSetting) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["templateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTemplateId(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSettingValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SettingValueable, len(val))
            for i, v := range val {
                res[i] = v.(SettingValueable)
            }
            m.SetValues(res)
        }
        return nil
    }
    return res
}
// GetTemplateId gets the templateId property value. Unique identifier for the tenant-level groupSettingTemplates object that's been customized for this group-level settings object. Read-only.
func (m *GroupSetting) GetTemplateId()(*string) {
    return m.templateId
}
// GetValues gets the values property value. Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced groupSettingTemplates object.
func (m *GroupSetting) GetValues()([]SettingValueable) {
    return m.values
}
// Serialize serializes information the current object
func (m *GroupSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteStringValue("templateId", m.GetTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValues()))
        for i, v := range m.GetValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("values", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Display name of this group of settings, which comes from the associated template.
func (m *GroupSetting) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTemplateId sets the templateId property value. Unique identifier for the tenant-level groupSettingTemplates object that's been customized for this group-level settings object. Read-only.
func (m *GroupSetting) SetTemplateId(value *string)() {
    m.templateId = value
}
// SetValues sets the values property value. Collection of name-value pairs corresponding to the name and defaultValue properties in the referenced groupSettingTemplates object.
func (m *GroupSetting) SetValues(value []SettingValueable)() {
    m.values = value
}
