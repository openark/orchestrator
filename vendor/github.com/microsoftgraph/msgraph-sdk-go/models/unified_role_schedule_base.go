package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleScheduleBase 
type UnifiedRoleScheduleBase struct {
    Entity
    // Read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
    appScope AppScopeable
    // Identifier of the app-specific scope when the assignment or eligibility is scoped to an app. The scope of an assignment or eligibility determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
    appScopeId *string
    // When the schedule was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Identifier of the object through which this schedule was created.
    createdUsing *string
    // The directory object that is the scope of the role eligibility or assignment. Read-only.
    directoryScope DirectoryObjectable
    // Identifier of the directory object representing the scope of the assignment or eligibility. The scope of an assignment or eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
    directoryScopeId *string
    // When the schedule was last modified.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The principal that's getting a role assignment or that's eligible for a role through the request.
    principal DirectoryObjectable
    // Identifier of the principal that has been granted the role assignment or eligibility.
    principalId *string
    // Detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
    roleDefinition UnifiedRoleDefinitionable
    // Identifier of the unifiedRoleDefinition object that is being assigned to the principal or that a principal is eligible for.
    roleDefinitionId *string
    // The status of the role assignment or eligibility request.
    status *string
}
// NewUnifiedRoleScheduleBase instantiates a new unifiedRoleScheduleBase and sets the default values.
func NewUnifiedRoleScheduleBase()(*UnifiedRoleScheduleBase) {
    m := &UnifiedRoleScheduleBase{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUnifiedRoleScheduleBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleScheduleBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.unifiedRoleAssignmentSchedule":
                        return NewUnifiedRoleAssignmentSchedule(), nil
                    case "#microsoft.graph.unifiedRoleEligibilitySchedule":
                        return NewUnifiedRoleEligibilitySchedule(), nil
                }
            }
        }
    }
    return NewUnifiedRoleScheduleBase(), nil
}
// GetAppScope gets the appScope property value. Read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
func (m *UnifiedRoleScheduleBase) GetAppScope()(AppScopeable) {
    return m.appScope
}
// GetAppScopeId gets the appScopeId property value. Identifier of the app-specific scope when the assignment or eligibility is scoped to an app. The scope of an assignment or eligibility determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
func (m *UnifiedRoleScheduleBase) GetAppScopeId()(*string) {
    return m.appScopeId
}
// GetCreatedDateTime gets the createdDateTime property value. When the schedule was created.
func (m *UnifiedRoleScheduleBase) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreatedUsing gets the createdUsing property value. Identifier of the object through which this schedule was created.
func (m *UnifiedRoleScheduleBase) GetCreatedUsing()(*string) {
    return m.createdUsing
}
// GetDirectoryScope gets the directoryScope property value. The directory object that is the scope of the role eligibility or assignment. Read-only.
func (m *UnifiedRoleScheduleBase) GetDirectoryScope()(DirectoryObjectable) {
    return m.directoryScope
}
// GetDirectoryScopeId gets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment or eligibility. The scope of an assignment or eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
func (m *UnifiedRoleScheduleBase) GetDirectoryScopeId()(*string) {
    return m.directoryScopeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleScheduleBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["createdUsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedUsing(val)
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
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. When the schedule was last modified.
func (m *UnifiedRoleScheduleBase) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetPrincipal gets the principal property value. The principal that's getting a role assignment or that's eligible for a role through the request.
func (m *UnifiedRoleScheduleBase) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. Identifier of the principal that has been granted the role assignment or eligibility.
func (m *UnifiedRoleScheduleBase) GetPrincipalId()(*string) {
    return m.principalId
}
// GetRoleDefinition gets the roleDefinition property value. Detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
func (m *UnifiedRoleScheduleBase) GetRoleDefinition()(UnifiedRoleDefinitionable) {
    return m.roleDefinition
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition object that is being assigned to the principal or that a principal is eligible for.
func (m *UnifiedRoleScheduleBase) GetRoleDefinitionId()(*string) {
    return m.roleDefinitionId
}
// GetStatus gets the status property value. The status of the role assignment or eligibility request.
func (m *UnifiedRoleScheduleBase) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *UnifiedRoleScheduleBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdUsing", m.GetCreatedUsing())
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
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
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
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppScope sets the appScope property value. Read-only property with details of the app-specific scope when the role eligibility or assignment is scoped to an app. Nullable.
func (m *UnifiedRoleScheduleBase) SetAppScope(value AppScopeable)() {
    m.appScope = value
}
// SetAppScopeId sets the appScopeId property value. Identifier of the app-specific scope when the assignment or eligibility is scoped to an app. The scope of an assignment or eligibility determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units.
func (m *UnifiedRoleScheduleBase) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
// SetCreatedDateTime sets the createdDateTime property value. When the schedule was created.
func (m *UnifiedRoleScheduleBase) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreatedUsing sets the createdUsing property value. Identifier of the object through which this schedule was created.
func (m *UnifiedRoleScheduleBase) SetCreatedUsing(value *string)() {
    m.createdUsing = value
}
// SetDirectoryScope sets the directoryScope property value. The directory object that is the scope of the role eligibility or assignment. Read-only.
func (m *UnifiedRoleScheduleBase) SetDirectoryScope(value DirectoryObjectable)() {
    m.directoryScope = value
}
// SetDirectoryScopeId sets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment or eligibility. The scope of an assignment or eligibility determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only.
func (m *UnifiedRoleScheduleBase) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. When the schedule was last modified.
func (m *UnifiedRoleScheduleBase) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetPrincipal sets the principal property value. The principal that's getting a role assignment or that's eligible for a role through the request.
func (m *UnifiedRoleScheduleBase) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. Identifier of the principal that has been granted the role assignment or eligibility.
func (m *UnifiedRoleScheduleBase) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetRoleDefinition sets the roleDefinition property value. Detailed information for the roleDefinition object that is referenced through the roleDefinitionId property.
func (m *UnifiedRoleScheduleBase) SetRoleDefinition(value UnifiedRoleDefinitionable)() {
    m.roleDefinition = value
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition object that is being assigned to the principal or that a principal is eligible for.
func (m *UnifiedRoleScheduleBase) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// SetStatus sets the status property value. The status of the role assignment or eligibility request.
func (m *UnifiedRoleScheduleBase) SetStatus(value *string)() {
    m.status = value
}
