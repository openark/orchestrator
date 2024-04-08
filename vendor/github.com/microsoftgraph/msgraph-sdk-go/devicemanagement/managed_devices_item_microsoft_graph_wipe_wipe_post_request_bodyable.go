package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevicesItemMicrosoftGraphWipeWipePostRequestBodyable 
type ManagedDevicesItemMicrosoftGraphWipeWipePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetKeepEnrollmentData()(*bool)
    GetKeepUserData()(*bool)
    GetMacOsUnlockCode()(*string)
    GetPersistEsimDataPlan()(*bool)
    SetKeepEnrollmentData(value *bool)()
    SetKeepUserData(value *bool)()
    SetMacOsUnlockCode(value *string)()
    SetPersistEsimDataPlan(value *bool)()
}
