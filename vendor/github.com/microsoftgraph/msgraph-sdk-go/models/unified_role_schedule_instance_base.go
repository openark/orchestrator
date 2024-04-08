package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleScheduleInstanceBase 
type UnifiedRoleScheduleInstanceBase struct {
    Entity
    // Read-only property with details of the app-specific scope when the assignment or role eligibility is scoped to an app. Nullable.
    appScope AppScopeable
    // Identifier of the app-specific scope when the assignment or role eligibility is scoped to an app. The scope of an assignment or role eligibility determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
    appScopeId *string
    // The directory object that is the scope of the assignment or role eligibility. Read-only.
    directoryScope DirectoryObjectable
    // Identifier of the directory object representing the scope of the assignment or role eligibility. The scope of an assignment or role eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
    directoryScopeId *string
    // The principal that's getting a role assignment or role eligibility through the request.
    principal DirectoryObjectable
    // Identifier of the principal that has been granted the role assignment or that's eligible for a role.
    principalId *string
    // Detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
    roleDefinition UnifiedRoleDefinitionable
    // Identifier of the unifiedRoleDefinition object that is being assigned to the principal or that the principal is eligible for.
    roleDefinitionId *string
}
// NewUnifiedRoleScheduleInstanceBase instantiates a new unifiedRoleScheduleInstanceBase and sets the default values.
func NewUnifiedRoleScheduleInstanceBase()(*UnifiedRoleScheduleInstanceBase) {
    m := &UnifiedRoleScheduleInstanceBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleScheduleInstanceBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleScheduleInstanceBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.unifiedRoleAssignmentScheduleInstance":
                        return NewUnifiedRoleAssignmentScheduleInstance(), nil
                    case "#microsoft.graph.unifiedRoleEligibilityScheduleInstance":
                        return NewUnifiedRoleEligibilityScheduleInstance(), nil
                }
            }
        }
    }
    return NewUnifiedRoleScheduleInstanceBase(), nil
}
// GetAppScope gets the appScope property value. Read-only property with details of the app-specific scope when the assignment or role eligibility is scoped to an app. Nullable.
func (m *UnifiedRoleScheduleInstanceBase) GetAppScope()(AppScopeable) {
    return m.appScope
}
// GetAppScopeId gets the appScopeId property value. Identifier of the app-specific scope when the assignment or role eligibility is scoped to an app. The scope of an assignment or role eligibility determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
func (m *UnifiedRoleScheduleInstanceBase) GetAppScopeId()(*string) {
    return m.appScopeId
}
// GetDirectoryScope gets the directoryScope property value. The directory object that is the scope of the assignment or role eligibility. Read-only.
func (m *UnifiedRoleScheduleInstanceBase) GetDirectoryScope()(DirectoryObjectable) {
    return m.directoryScope
}
// GetDirectoryScopeId gets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment or role eligibility. The scope of an assignment or role eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
func (m *UnifiedRoleScheduleInstanceBase) GetDirectoryScopeId()(*string) {
    return m.directoryScopeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleScheduleInstanceBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
// GetPrincipal gets the principal property value. The principal that's getting a role assignment or role eligibility through the request.
func (m *UnifiedRoleScheduleInstanceBase) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. Identifier of the principal that has been granted the role assignment or that's eligible for a role.
func (m *UnifiedRoleScheduleInstanceBase) GetPrincipalId()(*string) {
    return m.principalId
}
// GetRoleDefinition gets the roleDefinition property value. Detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
func (m *UnifiedRoleScheduleInstanceBase) GetRoleDefinition()(UnifiedRoleDefinitionable) {
    return m.roleDefinition
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition object that is being assigned to the principal or that the principal is eligible for.
func (m *UnifiedRoleScheduleInstanceBase) GetRoleDefinitionId()(*string) {
    return m.roleDefinitionId
}
// Serialize serializes information the current object
func (m *UnifiedRoleScheduleInstanceBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetAppScope sets the appScope property value. Read-only property with details of the app-specific scope when the assignment or role eligibility is scoped to an app. Nullable.
func (m *UnifiedRoleScheduleInstanceBase) SetAppScope(value AppScopeable)() {
    m.appScope = value
}
// SetAppScopeId sets the appScopeId property value. Identifier of the app-specific scope when the assignment or role eligibility is scoped to an app. The scope of an assignment or role eligibility determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
func (m *UnifiedRoleScheduleInstanceBase) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
// SetDirectoryScope sets the directoryScope property value. The directory object that is the scope of the assignment or role eligibility. Read-only.
func (m *UnifiedRoleScheduleInstanceBase) SetDirectoryScope(value DirectoryObjectable)() {
    m.directoryScope = value
}
// SetDirectoryScopeId sets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment or role eligibility. The scope of an assignment or role eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
func (m *UnifiedRoleScheduleInstanceBase) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
// SetPrincipal sets the principal property value. The principal that's getting a role assignment or role eligibility through the request.
func (m *UnifiedRoleScheduleInstanceBase) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. Identifier of the principal that has been granted the role assignment or that's eligible for a role.
func (m *UnifiedRoleScheduleInstanceBase) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetRoleDefinition sets the roleDefinition property value. Detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
func (m *UnifiedRoleScheduleInstanceBase) SetRoleDefinition(value UnifiedRoleDefinitionable)() {
    m.roleDefinition = value
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition object that is being assigned to the principal or that the principal is eligible for.
func (m *UnifiedRoleScheduleInstanceBase) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
