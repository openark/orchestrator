package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserFlowLanguageConfiguration 
type UserFlowLanguageConfiguration struct {
    Entity
    // Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
    defaultPages []UserFlowLanguagePageable
    // The language name to display. This property is read-only.
    displayName *string
    // Indicates whether the language is enabled within the user flow.
    isEnabled *bool
    // Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
    overridesPages []UserFlowLanguagePageable
}
// NewUserFlowLanguageConfiguration instantiates a new userFlowLanguageConfiguration and sets the default values.
func NewUserFlowLanguageConfiguration()(*UserFlowLanguageConfiguration) {
    m := &UserFlowLanguageConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserFlowLanguageConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserFlowLanguageConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserFlowLanguageConfiguration(), nil
}
// GetDefaultPages gets the defaultPages property value. Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
func (m *UserFlowLanguageConfiguration) GetDefaultPages()([]UserFlowLanguagePageable) {
    return m.defaultPages
}
// GetDisplayName gets the displayName property value. The language name to display. This property is read-only.
func (m *UserFlowLanguageConfiguration) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserFlowLanguageConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["defaultPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFlowLanguagePageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguagePageable, len(val))
            for i, v := range val {
                res[i] = v.(UserFlowLanguagePageable)
            }
            m.SetDefaultPages(res)
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
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["overridesPages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserFlowLanguagePageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserFlowLanguagePageable, len(val))
            for i, v := range val {
                res[i] = v.(UserFlowLanguagePageable)
            }
            m.SetOverridesPages(res)
        }
        return nil
    }
    return res
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the language is enabled within the user flow.
func (m *UserFlowLanguageConfiguration) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetOverridesPages gets the overridesPages property value. Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
func (m *UserFlowLanguageConfiguration) GetOverridesPages()([]UserFlowLanguagePageable) {
    return m.overridesPages
}
// Serialize serializes information the current object
func (m *UserFlowLanguageConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDefaultPages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefaultPages()))
        for i, v := range m.GetDefaultPages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("defaultPages", cast)
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
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetOverridesPages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOverridesPages()))
        for i, v := range m.GetOverridesPages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("overridesPages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultPages sets the defaultPages property value. Collection of pages with the default content to display in a user flow for a specified language. This collection does not allow any kind of modification.
func (m *UserFlowLanguageConfiguration) SetDefaultPages(value []UserFlowLanguagePageable)() {
    m.defaultPages = value
}
// SetDisplayName sets the displayName property value. The language name to display. This property is read-only.
func (m *UserFlowLanguageConfiguration) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the language is enabled within the user flow.
func (m *UserFlowLanguageConfiguration) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetOverridesPages sets the overridesPages property value. Collection of pages with the overrides messages to display in a user flow for a specified language. This collection only allows to modify the content of the page, any other modification is not allowed (creation or deletion of pages).
func (m *UserFlowLanguageConfiguration) SetOverridesPages(value []UserFlowLanguagePageable)() {
    m.overridesPages = value
}
