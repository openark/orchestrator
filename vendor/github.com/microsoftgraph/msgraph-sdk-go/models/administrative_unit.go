package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AdministrativeUnit 
type AdministrativeUnit struct {
    DirectoryObject
    // An optional description for the administrative unit. Supports $filter (eq, ne, in, startsWith), $search.
    description *string
    // Display name for the administrative unit. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string
    // The collection of open extensions defined for this administrative unit. Nullable.
    extensions []Extensionable
    // Users and groups that are members of this administrative unit. Supports $expand.
    members []DirectoryObjectable
    // Scoped-role members of this administrative unit.
    scopedRoleMembers []ScopedRoleMembershipable
    // Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership. If not set (value is null), the default behavior is public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
    visibility *string
}
// NewAdministrativeUnit instantiates a new AdministrativeUnit and sets the default values.
func NewAdministrativeUnit()(*AdministrativeUnit) {
    m := &AdministrativeUnit{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.administrativeUnit"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAdministrativeUnitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAdministrativeUnitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAdministrativeUnit(), nil
}
// GetDescription gets the description property value. An optional description for the administrative unit. Supports $filter (eq, ne, in, startsWith), $search.
func (m *AdministrativeUnit) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name for the administrative unit. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *AdministrativeUnit) GetDisplayName()(*string) {
    return m.displayName
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for this administrative unit. Nullable.
func (m *AdministrativeUnit) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AdministrativeUnit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["scopedRoleMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScopedRoleMembershipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScopedRoleMembershipable, len(val))
            for i, v := range val {
                res[i] = v.(ScopedRoleMembershipable)
            }
            m.SetScopedRoleMembers(res)
        }
        return nil
    }
    res["visibility"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibility(val)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. Users and groups that are members of this administrative unit. Supports $expand.
func (m *AdministrativeUnit) GetMembers()([]DirectoryObjectable) {
    return m.members
}
// GetScopedRoleMembers gets the scopedRoleMembers property value. Scoped-role members of this administrative unit.
func (m *AdministrativeUnit) GetScopedRoleMembers()([]ScopedRoleMembershipable) {
    return m.scopedRoleMembers
}
// GetVisibility gets the visibility property value. Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership. If not set (value is null), the default behavior is public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
func (m *AdministrativeUnit) GetVisibility()(*string) {
    return m.visibility
}
// Serialize serializes information the current object
func (m *AdministrativeUnit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetScopedRoleMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScopedRoleMembers()))
        for i, v := range m.GetScopedRoleMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("scopedRoleMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("visibility", m.GetVisibility())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. An optional description for the administrative unit. Supports $filter (eq, ne, in, startsWith), $search.
func (m *AdministrativeUnit) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name for the administrative unit. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *AdministrativeUnit) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for this administrative unit. Nullable.
func (m *AdministrativeUnit) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetMembers sets the members property value. Users and groups that are members of this administrative unit. Supports $expand.
func (m *AdministrativeUnit) SetMembers(value []DirectoryObjectable)() {
    m.members = value
}
// SetScopedRoleMembers sets the scopedRoleMembers property value. Scoped-role members of this administrative unit.
func (m *AdministrativeUnit) SetScopedRoleMembers(value []ScopedRoleMembershipable)() {
    m.scopedRoleMembers = value
}
// SetVisibility sets the visibility property value. Controls whether the administrative unit and its members are hidden or public. Can be set to HiddenMembership. If not set (value is null), the default behavior is public. When set to HiddenMembership, only members of the administrative unit can list other members of the administrative unit.
func (m *AdministrativeUnit) SetVisibility(value *string)() {
    m.visibility = value
}
