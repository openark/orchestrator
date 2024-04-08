package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScopedRoleMembership 
type ScopedRoleMembership struct {
    Entity
    // Unique identifier for the administrative unit that the directory role is scoped to
    administrativeUnitId *string
    // Unique identifier for the directory role that the member is in.
    roleId *string
    // The roleMemberInfo property
    roleMemberInfo Identityable
}
// NewScopedRoleMembership instantiates a new scopedRoleMembership and sets the default values.
func NewScopedRoleMembership()(*ScopedRoleMembership) {
    m := &ScopedRoleMembership{
        Entity: *NewEntity(),
    }
    return m
}
// CreateScopedRoleMembershipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScopedRoleMembershipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScopedRoleMembership(), nil
}
// GetAdministrativeUnitId gets the administrativeUnitId property value. Unique identifier for the administrative unit that the directory role is scoped to
func (m *ScopedRoleMembership) GetAdministrativeUnitId()(*string) {
    return m.administrativeUnitId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScopedRoleMembership) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnitId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdministrativeUnitId(val)
        }
        return nil
    }
    res["roleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleId(val)
        }
        return nil
    }
    res["roleMemberInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleMemberInfo(val.(Identityable))
        }
        return nil
    }
    return res
}
// GetRoleId gets the roleId property value. Unique identifier for the directory role that the member is in.
func (m *ScopedRoleMembership) GetRoleId()(*string) {
    return m.roleId
}
// GetRoleMemberInfo gets the roleMemberInfo property value. The roleMemberInfo property
func (m *ScopedRoleMembership) GetRoleMemberInfo()(Identityable) {
    return m.roleMemberInfo
}
// Serialize serializes information the current object
func (m *ScopedRoleMembership) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("administrativeUnitId", m.GetAdministrativeUnitId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleId", m.GetRoleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleMemberInfo", m.GetRoleMemberInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnitId sets the administrativeUnitId property value. Unique identifier for the administrative unit that the directory role is scoped to
func (m *ScopedRoleMembership) SetAdministrativeUnitId(value *string)() {
    m.administrativeUnitId = value
}
// SetRoleId sets the roleId property value. Unique identifier for the directory role that the member is in.
func (m *ScopedRoleMembership) SetRoleId(value *string)() {
    m.roleId = value
}
// SetRoleMemberInfo sets the roleMemberInfo property value. The roleMemberInfo property
func (m *ScopedRoleMembership) SetRoleMemberInfo(value Identityable)() {
    m.roleMemberInfo = value
}
