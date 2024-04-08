package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedWindowsAutopilotDeviceIdentityStateable 
type ImportedWindowsAutopilotDeviceIdentityStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDeviceErrorCode()(*int32)
    GetDeviceErrorName()(*string)
    GetDeviceImportStatus()(*ImportedWindowsAutopilotDeviceIdentityImportStatus)
    GetDeviceRegistrationId()(*string)
    GetOdataType()(*string)
    SetDeviceErrorCode(value *int32)()
    SetDeviceErrorName(value *string)()
    SetDeviceImportStatus(value *ImportedWindowsAutopilotDeviceIdentityImportStatus)()
    SetDeviceRegistrationId(value *string)()
    SetOdataType(value *string)()
}
