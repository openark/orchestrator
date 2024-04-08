package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComplianceManagementPartnerable 
type ComplianceManagementPartnerable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable)
    GetAndroidOnboarded()(*bool)
    GetDisplayName()(*string)
    GetIosEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable)
    GetIosOnboarded()(*bool)
    GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMacOsEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable)
    GetMacOsOnboarded()(*bool)
    GetPartnerState()(*DeviceManagementPartnerTenantState)
    SetAndroidEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)()
    SetAndroidOnboarded(value *bool)()
    SetDisplayName(value *string)()
    SetIosEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)()
    SetIosOnboarded(value *bool)()
    SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMacOsEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)()
    SetMacOsOnboarded(value *bool)()
    SetPartnerState(value *DeviceManagementPartnerTenantState)()
}
