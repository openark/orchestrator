package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedPermissionClassification 
type DelegatedPermissionClassification struct {
    Entity
    // The classification value being given. Possible value: low. Does not support $filter.
    classification *PermissionClassificationType
    // The unique identifier (id) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Required on create. Does not support $filter.
    permissionId *string
    // The claim value (value) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Does not support $filter.
    permissionName *string
}
// NewDelegatedPermissionClassification instantiates a new delegatedPermissionClassification and sets the default values.
func NewDelegatedPermissionClassification()(*DelegatedPermissionClassification) {
    m := &DelegatedPermissionClassification{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedPermissionClassificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedPermissionClassificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedPermissionClassification(), nil
}
// GetClassification gets the classification property value. The classification value being given. Possible value: low. Does not support $filter.
func (m *DelegatedPermissionClassification) GetClassification()(*PermissionClassificationType) {
    return m.classification
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedPermissionClassification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*PermissionClassificationType))
        }
        return nil
    }
    res["permissionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionId(val)
        }
        return nil
    }
    res["permissionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPermissionName(val)
        }
        return nil
    }
    return res
}
// GetPermissionId gets the permissionId property value. The unique identifier (id) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Required on create. Does not support $filter.
func (m *DelegatedPermissionClassification) GetPermissionId()(*string) {
    return m.permissionId
}
// GetPermissionName gets the permissionName property value. The claim value (value) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Does not support $filter.
func (m *DelegatedPermissionClassification) GetPermissionName()(*string) {
    return m.permissionName
}
// Serialize serializes information the current object
func (m *DelegatedPermissionClassification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionId", m.GetPermissionId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionName", m.GetPermissionName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassification sets the classification property value. The classification value being given. Possible value: low. Does not support $filter.
func (m *DelegatedPermissionClassification) SetClassification(value *PermissionClassificationType)() {
    m.classification = value
}
// SetPermissionId sets the permissionId property value. The unique identifier (id) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Required on create. Does not support $filter.
func (m *DelegatedPermissionClassification) SetPermissionId(value *string)() {
    m.permissionId = value
}
// SetPermissionName sets the permissionName property value. The claim value (value) for the delegated permission listed in the oauth2PermissionScopes collection of the servicePrincipal. Does not support $filter.
func (m *DelegatedPermissionClassification) SetPermissionName(value *string)() {
    m.permissionName = value
}
