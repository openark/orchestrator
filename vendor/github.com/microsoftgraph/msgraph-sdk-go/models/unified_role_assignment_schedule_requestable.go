package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignmentScheduleRequestable 
type UnifiedRoleAssignmentScheduleRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Requestable
    GetAction()(*UnifiedRoleScheduleRequestActions)
    GetActivatedUsing()(UnifiedRoleEligibilityScheduleable)
    GetAppScope()(AppScopeable)
    GetAppScopeId()(*string)
    GetDirectoryScope()(DirectoryObjectable)
    GetDirectoryScopeId()(*string)
    GetIsValidationOnly()(*bool)
    GetJustification()(*string)
    GetPrincipal()(DirectoryObjectable)
    GetPrincipalId()(*string)
    GetRoleDefinition()(UnifiedRoleDefinitionable)
    GetRoleDefinitionId()(*string)
    GetScheduleInfo()(RequestScheduleable)
    GetTargetSchedule()(UnifiedRoleAssignmentScheduleable)
    GetTargetScheduleId()(*string)
    GetTicketInfo()(TicketInfoable)
    SetAction(value *UnifiedRoleScheduleRequestActions)()
    SetActivatedUsing(value UnifiedRoleEligibilityScheduleable)()
    SetAppScope(value AppScopeable)()
    SetAppScopeId(value *string)()
    SetDirectoryScope(value DirectoryObjectable)()
    SetDirectoryScopeId(value *string)()
    SetIsValidationOnly(value *bool)()
    SetJustification(value *string)()
    SetPrincipal(value DirectoryObjectable)()
    SetPrincipalId(value *string)()
    SetRoleDefinition(value UnifiedRoleDefinitionable)()
    SetRoleDefinitionId(value *string)()
    SetScheduleInfo(value RequestScheduleable)()
    SetTargetSchedule(value UnifiedRoleAssignmentScheduleable)()
    SetTargetScheduleId(value *string)()
    SetTicketInfo(value TicketInfoable)()
}
