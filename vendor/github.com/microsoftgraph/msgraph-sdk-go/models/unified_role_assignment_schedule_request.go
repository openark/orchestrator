package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignmentScheduleRequest 
type UnifiedRoleAssignmentScheduleRequest struct {
    Request
    // Represents the type of the operation on the role assignment request. The possible values are: adminAssign, adminUpdate, adminRemove, selfActivate, selfDeactivate, adminExtend, adminRenew, selfExtend, selfRenew, unknownFutureValue. adminAssign: For administrators to assign roles to principals.adminRemove: For administrators to remove principals from roles. adminUpdate: For administrators to change existing role assignments.adminExtend: For administrators to extend expiring assignments.adminRenew: For administrators to renew expired assignments.selfActivate: For principals to activate their assignments.selfDeactivate: For principals to deactivate their active assignments.selfExtend: For principals to request to extend their expiring assignments.selfRenew: For principals to request to renew their expired assignments.
    action *UnifiedRoleScheduleRequestActions
    // If the request is from an eligible administrator to activate a role, this parameter will show the related eligible assignment for that activation. Otherwise, it's null. Supports $expand.
    activatedUsing UnifiedRoleEligibilityScheduleable
    // Read-only property with details of the app-specific scope when the assignment is scoped to an app. Nullable. Supports $expand.
    appScope AppScopeable
    // Identifier of the app-specific scope when the assignment is scoped to an app. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, ne, and on null values).
    appScopeId *string
    // The directory object that is the scope of the assignment. Read-only. Supports $expand.
    directoryScope DirectoryObjectable
    // Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, ne, and on null values).
    directoryScopeId *string
    // Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
    isValidationOnly *bool
    // A message provided by users and administrators when create they create the unifiedRoleAssignmentScheduleRequest object.
    justification *string
    // The principal that's getting a role assignment through the request. Supports $expand.
    principal DirectoryObjectable
    // Identifier of the principal that has been granted the assignment. Can be a user, role-assignable group, or a service principal. Supports $filter (eq, ne).
    principalId *string
    // Detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
    roleDefinition UnifiedRoleDefinitionable
    // Identifier of the unifiedRoleDefinition object that is being assigned to the principal. Supports $filter (eq, ne).
    roleDefinitionId *string
    // The period of the role assignment. Recurring schedules are currently unsupported.
    scheduleInfo RequestScheduleable
    // The schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand.
    targetSchedule UnifiedRoleAssignmentScheduleable
    // Identifier of the schedule object that's linked to the assignment request. Supports $filter (eq, ne).
    targetScheduleId *string
    // Ticket details linked to the role assignment request including details of the ticket number and ticket system.
    ticketInfo TicketInfoable
}
// NewUnifiedRoleAssignmentScheduleRequest instantiates a new UnifiedRoleAssignmentScheduleRequest and sets the default values.
func NewUnifiedRoleAssignmentScheduleRequest()(*UnifiedRoleAssignmentScheduleRequest) {
    m := &UnifiedRoleAssignmentScheduleRequest{
        Request: *NewRequest(),
    }
    return m
}
// CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleAssignmentScheduleRequest(), nil
}
// GetAction gets the action property value. Represents the type of the operation on the role assignment request. The possible values are: adminAssign, adminUpdate, adminRemove, selfActivate, selfDeactivate, adminExtend, adminRenew, selfExtend, selfRenew, unknownFutureValue. adminAssign: For administrators to assign roles to principals.adminRemove: For administrators to remove principals from roles. adminUpdate: For administrators to change existing role assignments.adminExtend: For administrators to extend expiring assignments.adminRenew: For administrators to renew expired assignments.selfActivate: For principals to activate their assignments.selfDeactivate: For principals to deactivate their active assignments.selfExtend: For principals to request to extend their expiring assignments.selfRenew: For principals to request to renew their expired assignments.
func (m *UnifiedRoleAssignmentScheduleRequest) GetAction()(*UnifiedRoleScheduleRequestActions) {
    return m.action
}
// GetActivatedUsing gets the activatedUsing property value. If the request is from an eligible administrator to activate a role, this parameter will show the related eligible assignment for that activation. Otherwise, it's null. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) GetActivatedUsing()(UnifiedRoleEligibilityScheduleable) {
    return m.activatedUsing
}
// GetAppScope gets the appScope property value. Read-only property with details of the app-specific scope when the assignment is scoped to an app. Nullable. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) GetAppScope()(AppScopeable) {
    return m.appScope
}
// GetAppScopeId gets the appScopeId property value. Identifier of the app-specific scope when the assignment is scoped to an app. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, ne, and on null values).
func (m *UnifiedRoleAssignmentScheduleRequest) GetAppScopeId()(*string) {
    return m.appScopeId
}
// GetDirectoryScope gets the directoryScope property value. The directory object that is the scope of the assignment. Read-only. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) GetDirectoryScope()(DirectoryObjectable) {
    return m.directoryScope
}
// GetDirectoryScopeId gets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, ne, and on null values).
func (m *UnifiedRoleAssignmentScheduleRequest) GetDirectoryScopeId()(*string) {
    return m.directoryScopeId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleAssignmentScheduleRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Request.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUnifiedRoleScheduleRequestActions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*UnifiedRoleScheduleRequestActions))
        }
        return nil
    }
    res["activatedUsing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedUsing(val.(UnifiedRoleEligibilityScheduleable))
        }
        return nil
    }
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
    res["isValidationOnly"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsValidationOnly(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
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
    res["scheduleInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleInfo(val.(RequestScheduleable))
        }
        return nil
    }
    res["targetSchedule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSchedule(val.(UnifiedRoleAssignmentScheduleable))
        }
        return nil
    }
    res["targetScheduleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetScheduleId(val)
        }
        return nil
    }
    res["ticketInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTicketInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTicketInfo(val.(TicketInfoable))
        }
        return nil
    }
    return res
}
// GetIsValidationOnly gets the isValidationOnly property value. Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
func (m *UnifiedRoleAssignmentScheduleRequest) GetIsValidationOnly()(*bool) {
    return m.isValidationOnly
}
// GetJustification gets the justification property value. A message provided by users and administrators when create they create the unifiedRoleAssignmentScheduleRequest object.
func (m *UnifiedRoleAssignmentScheduleRequest) GetJustification()(*string) {
    return m.justification
}
// GetPrincipal gets the principal property value. The principal that's getting a role assignment through the request. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) GetPrincipal()(DirectoryObjectable) {
    return m.principal
}
// GetPrincipalId gets the principalId property value. Identifier of the principal that has been granted the assignment. Can be a user, role-assignable group, or a service principal. Supports $filter (eq, ne).
func (m *UnifiedRoleAssignmentScheduleRequest) GetPrincipalId()(*string) {
    return m.principalId
}
// GetRoleDefinition gets the roleDefinition property value. Detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) GetRoleDefinition()(UnifiedRoleDefinitionable) {
    return m.roleDefinition
}
// GetRoleDefinitionId gets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition object that is being assigned to the principal. Supports $filter (eq, ne).
func (m *UnifiedRoleAssignmentScheduleRequest) GetRoleDefinitionId()(*string) {
    return m.roleDefinitionId
}
// GetScheduleInfo gets the scheduleInfo property value. The period of the role assignment. Recurring schedules are currently unsupported.
func (m *UnifiedRoleAssignmentScheduleRequest) GetScheduleInfo()(RequestScheduleable) {
    return m.scheduleInfo
}
// GetTargetSchedule gets the targetSchedule property value. The schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) GetTargetSchedule()(UnifiedRoleAssignmentScheduleable) {
    return m.targetSchedule
}
// GetTargetScheduleId gets the targetScheduleId property value. Identifier of the schedule object that's linked to the assignment request. Supports $filter (eq, ne).
func (m *UnifiedRoleAssignmentScheduleRequest) GetTargetScheduleId()(*string) {
    return m.targetScheduleId
}
// GetTicketInfo gets the ticketInfo property value. Ticket details linked to the role assignment request including details of the ticket number and ticket system.
func (m *UnifiedRoleAssignmentScheduleRequest) GetTicketInfo()(TicketInfoable) {
    return m.ticketInfo
}
// Serialize serializes information the current object
func (m *UnifiedRoleAssignmentScheduleRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Request.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("activatedUsing", m.GetActivatedUsing())
        if err != nil {
            return err
        }
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
        err = writer.WriteBoolValue("isValidationOnly", m.GetIsValidationOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
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
        err = writer.WriteObjectValue("scheduleInfo", m.GetScheduleInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetSchedule", m.GetTargetSchedule())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetScheduleId", m.GetTargetScheduleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("ticketInfo", m.GetTicketInfo())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAction sets the action property value. Represents the type of the operation on the role assignment request. The possible values are: adminAssign, adminUpdate, adminRemove, selfActivate, selfDeactivate, adminExtend, adminRenew, selfExtend, selfRenew, unknownFutureValue. adminAssign: For administrators to assign roles to principals.adminRemove: For administrators to remove principals from roles. adminUpdate: For administrators to change existing role assignments.adminExtend: For administrators to extend expiring assignments.adminRenew: For administrators to renew expired assignments.selfActivate: For principals to activate their assignments.selfDeactivate: For principals to deactivate their active assignments.selfExtend: For principals to request to extend their expiring assignments.selfRenew: For principals to request to renew their expired assignments.
func (m *UnifiedRoleAssignmentScheduleRequest) SetAction(value *UnifiedRoleScheduleRequestActions)() {
    m.action = value
}
// SetActivatedUsing sets the activatedUsing property value. If the request is from an eligible administrator to activate a role, this parameter will show the related eligible assignment for that activation. Otherwise, it's null. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) SetActivatedUsing(value UnifiedRoleEligibilityScheduleable)() {
    m.activatedUsing = value
}
// SetAppScope sets the appScope property value. Read-only property with details of the app-specific scope when the assignment is scoped to an app. Nullable. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) SetAppScope(value AppScopeable)() {
    m.appScope = value
}
// SetAppScopeId sets the appScopeId property value. Identifier of the app-specific scope when the assignment is scoped to an app. The scope of an assignment determines the set of resources for which the principal has been granted access. App scopes are scopes that are defined and understood by this application only. Use / for tenant-wide app scopes. Use directoryScopeId to limit the scope to particular directory objects, for example, administrative units. Supports $filter (eq, ne, and on null values).
func (m *UnifiedRoleAssignmentScheduleRequest) SetAppScopeId(value *string)() {
    m.appScopeId = value
}
// SetDirectoryScope sets the directoryScope property value. The directory object that is the scope of the assignment. Read-only. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) SetDirectoryScope(value DirectoryObjectable)() {
    m.directoryScope = value
}
// SetDirectoryScopeId sets the directoryScopeId property value. Identifier of the directory object representing the scope of the assignment. The scope of an assignment determines the set of resources for which the principal has been granted access. Directory scopes are shared scopes stored in the directory that are understood by multiple applications. Use / for tenant-wide scope. Use appScopeId to limit the scope to an application only. Supports $filter (eq, ne, and on null values).
func (m *UnifiedRoleAssignmentScheduleRequest) SetDirectoryScopeId(value *string)() {
    m.directoryScopeId = value
}
// SetIsValidationOnly sets the isValidationOnly property value. Determines whether the call is a validation or an actual call. Only set this property if you want to check whether an activation is subject to additional rules like MFA before actually submitting the request.
func (m *UnifiedRoleAssignmentScheduleRequest) SetIsValidationOnly(value *bool)() {
    m.isValidationOnly = value
}
// SetJustification sets the justification property value. A message provided by users and administrators when create they create the unifiedRoleAssignmentScheduleRequest object.
func (m *UnifiedRoleAssignmentScheduleRequest) SetJustification(value *string)() {
    m.justification = value
}
// SetPrincipal sets the principal property value. The principal that's getting a role assignment through the request. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) SetPrincipal(value DirectoryObjectable)() {
    m.principal = value
}
// SetPrincipalId sets the principalId property value. Identifier of the principal that has been granted the assignment. Can be a user, role-assignable group, or a service principal. Supports $filter (eq, ne).
func (m *UnifiedRoleAssignmentScheduleRequest) SetPrincipalId(value *string)() {
    m.principalId = value
}
// SetRoleDefinition sets the roleDefinition property value. Detailed information for the unifiedRoleDefinition object that is referenced through the roleDefinitionId property. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) SetRoleDefinition(value UnifiedRoleDefinitionable)() {
    m.roleDefinition = value
}
// SetRoleDefinitionId sets the roleDefinitionId property value. Identifier of the unifiedRoleDefinition object that is being assigned to the principal. Supports $filter (eq, ne).
func (m *UnifiedRoleAssignmentScheduleRequest) SetRoleDefinitionId(value *string)() {
    m.roleDefinitionId = value
}
// SetScheduleInfo sets the scheduleInfo property value. The period of the role assignment. Recurring schedules are currently unsupported.
func (m *UnifiedRoleAssignmentScheduleRequest) SetScheduleInfo(value RequestScheduleable)() {
    m.scheduleInfo = value
}
// SetTargetSchedule sets the targetSchedule property value. The schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand.
func (m *UnifiedRoleAssignmentScheduleRequest) SetTargetSchedule(value UnifiedRoleAssignmentScheduleable)() {
    m.targetSchedule = value
}
// SetTargetScheduleId sets the targetScheduleId property value. Identifier of the schedule object that's linked to the assignment request. Supports $filter (eq, ne).
func (m *UnifiedRoleAssignmentScheduleRequest) SetTargetScheduleId(value *string)() {
    m.targetScheduleId = value
}
// SetTicketInfo sets the ticketInfo property value. Ticket details linked to the role assignment request including details of the ticket number and ticket system.
func (m *UnifiedRoleAssignmentScheduleRequest) SetTicketInfo(value TicketInfoable)() {
    m.ticketInfo = value
}
