package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityUserFlowAttributeAssignment 
type IdentityUserFlowAttributeAssignment struct {
    Entity
    // The display name of the identityUserFlowAttribute within a user flow.
    displayName *string
    // Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value. false means the user cannot complete sign-up without providing a value.
    isOptional *bool
    // Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user's phone number or email address.
    requiresVerification *bool
    // The user attribute that you want to add to your user flow.
    userAttribute IdentityUserFlowAttributeable
    // The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect.
    userAttributeValues []UserAttributeValuesItemable
    // The userInputType property
    userInputType *IdentityUserFlowAttributeInputType
}
// NewIdentityUserFlowAttributeAssignment instantiates a new identityUserFlowAttributeAssignment and sets the default values.
func NewIdentityUserFlowAttributeAssignment()(*IdentityUserFlowAttributeAssignment) {
    m := &IdentityUserFlowAttributeAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentityUserFlowAttributeAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentityUserFlowAttributeAssignment(), nil
}
// GetDisplayName gets the displayName property value. The display name of the identityUserFlowAttribute within a user flow.
func (m *IdentityUserFlowAttributeAssignment) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentityUserFlowAttributeAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["isOptional"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOptional(val)
        }
        return nil
    }
    res["requiresVerification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequiresVerification(val)
        }
        return nil
    }
    res["userAttribute"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityUserFlowAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAttribute(val.(IdentityUserFlowAttributeable))
        }
        return nil
    }
    res["userAttributeValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUserAttributeValuesItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UserAttributeValuesItemable, len(val))
            for i, v := range val {
                res[i] = v.(UserAttributeValuesItemable)
            }
            m.SetUserAttributeValues(res)
        }
        return nil
    }
    res["userInputType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityUserFlowAttributeInputType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserInputType(val.(*IdentityUserFlowAttributeInputType))
        }
        return nil
    }
    return res
}
// GetIsOptional gets the isOptional property value. Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value. false means the user cannot complete sign-up without providing a value.
func (m *IdentityUserFlowAttributeAssignment) GetIsOptional()(*bool) {
    return m.isOptional
}
// GetRequiresVerification gets the requiresVerification property value. Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user's phone number or email address.
func (m *IdentityUserFlowAttributeAssignment) GetRequiresVerification()(*bool) {
    return m.requiresVerification
}
// GetUserAttribute gets the userAttribute property value. The user attribute that you want to add to your user flow.
func (m *IdentityUserFlowAttributeAssignment) GetUserAttribute()(IdentityUserFlowAttributeable) {
    return m.userAttribute
}
// GetUserAttributeValues gets the userAttributeValues property value. The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect.
func (m *IdentityUserFlowAttributeAssignment) GetUserAttributeValues()([]UserAttributeValuesItemable) {
    return m.userAttributeValues
}
// GetUserInputType gets the userInputType property value. The userInputType property
func (m *IdentityUserFlowAttributeAssignment) GetUserInputType()(*IdentityUserFlowAttributeInputType) {
    return m.userInputType
}
// Serialize serializes information the current object
func (m *IdentityUserFlowAttributeAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteBoolValue("isOptional", m.GetIsOptional())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requiresVerification", m.GetRequiresVerification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userAttribute", m.GetUserAttribute())
        if err != nil {
            return err
        }
    }
    if m.GetUserAttributeValues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUserAttributeValues()))
        for i, v := range m.GetUserAttributeValues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("userAttributeValues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserInputType() != nil {
        cast := (*m.GetUserInputType()).String()
        err = writer.WriteStringValue("userInputType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The display name of the identityUserFlowAttribute within a user flow.
func (m *IdentityUserFlowAttributeAssignment) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetIsOptional sets the isOptional property value. Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value. false means the user cannot complete sign-up without providing a value.
func (m *IdentityUserFlowAttributeAssignment) SetIsOptional(value *bool)() {
    m.isOptional = value
}
// SetRequiresVerification sets the requiresVerification property value. Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user's phone number or email address.
func (m *IdentityUserFlowAttributeAssignment) SetRequiresVerification(value *bool)() {
    m.requiresVerification = value
}
// SetUserAttribute sets the userAttribute property value. The user attribute that you want to add to your user flow.
func (m *IdentityUserFlowAttributeAssignment) SetUserAttribute(value IdentityUserFlowAttributeable)() {
    m.userAttribute = value
}
// SetUserAttributeValues sets the userAttributeValues property value. The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect.
func (m *IdentityUserFlowAttributeAssignment) SetUserAttributeValues(value []UserAttributeValuesItemable)() {
    m.userAttributeValues = value
}
// SetUserInputType sets the userInputType property value. The userInputType property
func (m *IdentityUserFlowAttributeAssignment) SetUserInputType(value *IdentityUserFlowAttributeInputType)() {
    m.userInputType = value
}
