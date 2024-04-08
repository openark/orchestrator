package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentRequestorSettingsable 
type AccessPackageAssignmentRequestorSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowCustomAssignmentSchedule()(*bool)
    GetEnableOnBehalfRequestorsToAddAccess()(*bool)
    GetEnableOnBehalfRequestorsToRemoveAccess()(*bool)
    GetEnableOnBehalfRequestorsToUpdateAccess()(*bool)
    GetEnableTargetsToSelfAddAccess()(*bool)
    GetEnableTargetsToSelfRemoveAccess()(*bool)
    GetEnableTargetsToSelfUpdateAccess()(*bool)
    GetOdataType()(*string)
    GetOnBehalfRequestors()([]SubjectSetable)
    SetAllowCustomAssignmentSchedule(value *bool)()
    SetEnableOnBehalfRequestorsToAddAccess(value *bool)()
    SetEnableOnBehalfRequestorsToRemoveAccess(value *bool)()
    SetEnableOnBehalfRequestorsToUpdateAccess(value *bool)()
    SetEnableTargetsToSelfAddAccess(value *bool)()
    SetEnableTargetsToSelfRemoveAccess(value *bool)()
    SetEnableTargetsToSelfUpdateAccess(value *bool)()
    SetOdataType(value *string)()
    SetOnBehalfRequestors(value []SubjectSetable)()
}
