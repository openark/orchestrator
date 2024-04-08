package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionAppLearningSummaryable 
type WindowsInformationProtectionAppLearningSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationName()(*string)
    GetApplicationType()(*ApplicationType)
    GetDeviceCount()(*int32)
    SetApplicationName(value *string)()
    SetApplicationType(value *ApplicationType)()
    SetDeviceCount(value *int32)()
}
