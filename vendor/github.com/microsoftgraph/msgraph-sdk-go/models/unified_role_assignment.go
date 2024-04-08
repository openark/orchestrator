package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignment 
type UnifiedRoleAssignment struct {
    Entity
    // Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity. Supports $expand.
    appScope AppScopeable
    // Identifier of the app-specific scope when the assignment scope is app-specific.  Either this property or directoryScopeId is required. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, in).
    appScopeId *string
    // The condition property
    condition *string
    // The directory object that is the scope of the assignment. Read-only. Supports $expand.
    directoryScope DirectoryObjectable
    // Identifier of the directory object representing the scope of the assignment.  Either this property or appScopeId is required. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, in).
    directoryScopeId *string
    // Referencing the assigned principal. Read-only. Supports $expand.
    principal DirectoryObjectable
    // Identifier of the principal to which the assignment is granted. Supports $filter (eq, in).
    principalId *string
    // The roleDefinition the assignment is for.  Supports $expand. roleDefinition.Id will be auto expanded.
    roleDefinition UnifiedRoleDefinitionable
    // Identifier of the role definition the assignment is for. Read only. Supports $filter (eq, in).
    roleDefinitionId *string
}
// NewUnifiedRoleAssignment instantiates a new unifiedRoleAssignment and sets the default values.
func NewUnifiedRoleAssignment()(*UnifiedRoleAssignment) {
    m := &UnifiedRoleAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleAssignment(), nil
}
// GetAppScope gets the appScope property value. Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity. Supports $expand.
func (m *UnifiedRoleAssignment) GetAppScope()(AppScopeable) {
    return m.appScope
}
// GetAppScopeId gets the appScopeId property value. Identifier of the app-specific scope when the assignment scope is app-specific.  Either this property or directoryScopeId is required. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, in).
func (m *UnifiedRoleAssignment) GetAppScopeId()(*string) {
    return m.appScopeId
}
// GetCondition gets the condition property value. The condition property
func (m *UnifiedRoleAssignment) GetCondition()(*string) {
    return m.condition
}
// GetDirectoryScope gets the directoryScope property value. The directory object that is the scope of the assignment. Read-only. Supports $expand.
func (m *UnifiedRoleAssignment) GetDirectoryScope()(DirectoryObjectable) {
    return m.directoryScope
}
// GetDirectoryScopeId gets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment.  Either this property or appScopeId is required. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, in).
func (m *UnifiedRoleAssignment) GetDirectoryScopeId()(*string) {
    return m.directoryScopeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAppScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppScope(val.(AppScopeable))
        }
        return nil
    }
    res["appScopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppScopeId(val)
        }
        return nil
    }
    res["condition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCondition(val)
        }
        return nil
    }
    res["directoryScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryScope(val.(DirectoryObjectable))
        }
        return nil
    }
    res["directoryScopeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectoryScopeId(val)
        }
        return nil
    }
    res["principal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipal(val.(DirectoryObjectable))
        }
        return nil
    }
    res["principalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalId(val)
        }
        return nil
    }
    res["roleDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinition(val.(UnifiedRoleDefinitionable))
        }
        return nil
    }
    res["roleDefinitionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoleDefinitionId(val)
        }
        return nil
    }
    return res
}
// GetPrincipal gets the principal property value. Referencing the assigned principal. Read-only. Supports $expand.
func (m *UnifiedRoleAssignment) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. Identifier of the principal to which the assignment is granted. Supports $filter (eq, in).
func (m *UnifiedRoleAssignment) GetPrincipalId()(*string) {
    return m.principalId
}
// GetRoleDefinition gets the roleDefinition property value. The roleDefinition the assignment is for.  Supports $expand. roleDefinition.Id will be auto expanded.
func (m *UnifiedRoleAssignment) GetRoleDefinition()(UnifiedRoleDefinitionable) {
    return m.roleDefinition
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Identifier of the role definition the assignment is for. Read only. Supports $filter (eq, in).
func (m *UnifiedRoleAssignment) GetRoleDefinitionId()(*string) {
    return m.roleDefinitionId
}
// Serialize serializes information the current object
func (m *UnifiedRoleAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("appScope", m.GetAppScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appScopeId", m.GetAppScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("condition", m.GetCondition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("directoryScope", m.GetDirectoryScope())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("directoryScopeId", m.GetDirectoryScopeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalId", m.GetPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("roleDefinition", m.GetRoleDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("roleDefinitionId", m.GetRoleDefinitionId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppScope sets the appScope property value. Read-only property with details of the app specific scope when the assignment scope is app specific. Containment entity. Supports $expand.
func (m *UnifiedRoleAssignment) SetAppScope(value AppScopeable)() {
    m.appScope = value
}
// SetAppScopeId sets the appScopeId property value. Identifier of the app-specific scope when the assignment scope is app-specific.  Either this property or directoryScopeId is required. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, in).
func (m *UnifiedRoleAssignment) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
// SetCondition sets the condition property value. The condition property
func (m *UnifiedRoleAssignment) SetCondition(value *string)() {
    m.condition = value
}
// SetDirectoryScope sets the directoryScope property value. The directory object that is the scope of the assignment. Read-only. Supports $expand.
func (m *UnifiedRoleAssignment) SetDirectoryScope(value DirectoryObjectable)() {
    m.directoryScope = value
}
// SetDirectoryScopeId sets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment.  Either this property or appScopeId is required. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, in).
func (m *UnifiedRoleAssignment) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
// SetPrincipal sets the principal property value. Referencing the assigned principal. Read-only. Supports $expand.
func (m *UnifiedRoleAssignment) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. Identifier of the principal to which the assignment is granted. Supports $filter (eq, in).
func (m *UnifiedRoleAssignment) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetRoleDefinition sets the roleDefinition property value. The roleDefinition the assignment is for.  Supports $expand. roleDefinition.Id will be auto expanded.
func (m *UnifiedRoleAssignment) SetRoleDefinition(value UnifiedRoleDefinitionable)() {
    m.roleDefinition = value
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Identifier of the role definition the assignment is for. Read only. Supports $filter (eq, in).
func (m *UnifiedRoleAssignment) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
