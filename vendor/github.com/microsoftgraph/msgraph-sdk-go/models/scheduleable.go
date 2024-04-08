package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Scheduleable 
type Scheduleable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEnabled()(*bool)
    GetOfferShiftRequests()([]OfferShiftRequestable)
    GetOfferShiftRequestsEnabled()(*bool)
    GetOpenShiftChangeRequests()([]OpenShiftChangeRequestable)
    GetOpenShifts()([]OpenShiftable)
    GetOpenShiftsEnabled()(*bool)
    GetProvisionStatus()(*OperationStatus)
    GetProvisionStatusCode()(*string)
    GetSchedulingGroups()([]SchedulingGroupable)
    GetShifts()([]Shiftable)
    GetSwapShiftsChangeRequests()([]SwapShiftsChangeRequestable)
    GetSwapShiftsRequestsEnabled()(*bool)
    GetTimeClockEnabled()(*bool)
    GetTimeOffReasons()([]TimeOffReasonable)
    GetTimeOffRequests()([]TimeOffRequestable)
    GetTimeOffRequestsEnabled()(*bool)
    GetTimesOff()([]TimeOffable)
    GetTimeZone()(*string)
    GetWorkforceIntegrationIds()([]string)
    SetEnabled(value *bool)()
    SetOfferShiftRequests(value []OfferShiftRequestable)()
    SetOfferShiftRequestsEnabled(value *bool)()
    SetOpenShiftChangeRequests(value []OpenShiftChangeRequestable)()
    SetOpenShifts(value []OpenShiftable)()
    SetOpenShiftsEnabled(value *bool)()
    SetProvisionStatus(value *OperationStatus)()
    SetProvisionStatusCode(value *string)()
    SetSchedulingGroups(value []SchedulingGroupable)()
    SetShifts(value []Shiftable)()
    SetSwapShiftsChangeRequests(value []SwapShiftsChangeRequestable)()
    SetSwapShiftsRequestsEnabled(value *bool)()
    SetTimeClockEnabled(value *bool)()
    SetTimeOffReasons(value []TimeOffReasonable)()
    SetTimeOffRequests(value []TimeOffRequestable)()
    SetTimeOffRequestsEnabled(value *bool)()
    SetTimesOff(value []TimeOffable)()
    SetTimeZone(value *string)()
    SetWorkforceIntegrationIds(value []string)()
}
