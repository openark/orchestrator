package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentTroubleshootingEventable 
type EnrollmentTroubleshootingEventable interface {
    DeviceManagementTroubleshootingEventable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceId()(*string)
    GetEnrollmentType()(*DeviceEnrollmentType)
    GetFailureCategory()(*DeviceEnrollmentFailureReason)
    GetFailureReason()(*string)
    GetManagedDeviceIdentifier()(*string)
    GetOperatingSystem()(*string)
    GetOsVersion()(*string)
    GetUserId()(*string)
    SetDeviceId(value *string)()
    SetEnrollmentType(value *DeviceEnrollmentType)()
    SetFailureCategory(value *DeviceEnrollmentFailureReason)()
    SetFailureReason(value *string)()
    SetManagedDeviceIdentifier(value *string)()
    SetOperatingSystem(value *string)()
    SetOsVersion(value *string)()
    SetUserId(value *string)()
}
