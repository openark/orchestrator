package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessTemplate 
type ConditionalAccessTemplate struct {
    Entity
    // The user-friendly name of the template.
    description *string
    // The details property
    details ConditionalAccessPolicyDetailable
    // The user-friendly name of the template.
    name *string
    // The scenarios property
    scenarios *TemplateScenarios
}
// NewConditionalAccessTemplate instantiates a new conditionalAccessTemplate and sets the default values.
func NewConditionalAccessTemplate()(*ConditionalAccessTemplate) {
    m := &ConditionalAccessTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateConditionalAccessTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessTemplate(), nil
}
// GetDescription gets the description property value. The user-friendly name of the template.
func (m *ConditionalAccessTemplate) GetDescription()(*string) {
    return m.description
}
// GetDetails gets the details property value. The details property
func (m *ConditionalAccessTemplate) GetDetails()(ConditionalAccessPolicyDetailable) {
    return m.details
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["details"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConditionalAccessPolicyDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetails(val.(ConditionalAccessPolicyDetailable))
        }
        return nil
    }
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
    res["scenarios"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTemplateScenarios)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScenarios(val.(*TemplateScenarios))
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The user-friendly name of the template.
func (m *ConditionalAccessTemplate) GetName()(*string) {
    return m.name
}
// GetScenarios gets the scenarios property value. The scenarios property
func (m *ConditionalAccessTemplate) GetScenarios()(*TemplateScenarios) {
    return m.scenarios
}
// Serialize serializes information the current object
func (m *ConditionalAccessTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("details", m.GetDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    if m.GetScenarios() != nil {
        cast := (*m.GetScenarios()).String()
        err = writer.WriteStringValue("scenarios", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The user-friendly name of the template.
func (m *ConditionalAccessTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDetails sets the details property value. The details property
func (m *ConditionalAccessTemplate) SetDetails(value ConditionalAccessPolicyDetailable)() {
    m.details = value
}
// SetName sets the name property value. The user-friendly name of the template.
func (m *ConditionalAccessTemplate) SetName(value *string)() {
    m.name = value
}
// SetScenarios sets the scenarios property value. The scenarios property
func (m *ConditionalAccessTemplate) SetScenarios(value *TemplateScenarios)() {
    m.scenarios = value
}
