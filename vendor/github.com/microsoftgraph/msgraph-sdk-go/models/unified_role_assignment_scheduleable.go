package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleAssignmentScheduleable 
type UnifiedRoleAssignmentScheduleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleScheduleBaseable
    GetActivatedUsing()(UnifiedRoleEligibilityScheduleable)
    GetAssignmentType()(*string)
    GetMemberType()(*string)
    GetScheduleInfo()(RequestScheduleable)
    SetActivatedUsing(value UnifiedRoleEligibilityScheduleable)()
    SetAssignmentType(value *string)()
    SetMemberType(value *string)()
    SetScheduleInfo(value RequestScheduleable)()
}
