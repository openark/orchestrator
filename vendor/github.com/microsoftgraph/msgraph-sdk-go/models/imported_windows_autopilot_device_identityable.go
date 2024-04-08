package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedWindowsAutopilotDeviceIdentityable 
type ImportedWindowsAutopilotDeviceIdentityable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignedUserPrincipalName()(*string)
    GetGroupTag()(*string)
    GetHardwareIdentifier()([]byte)
    GetImportId()(*string)
    GetProductKey()(*string)
    GetSerialNumber()(*string)
    GetState()(ImportedWindowsAutopilotDeviceIdentityStateable)
    SetAssignedUserPrincipalName(value *string)()
    SetGroupTag(value *string)()
    SetHardwareIdentifier(value []byte)()
    SetImportId(value *string)()
    SetProductKey(value *string)()
    SetSerialNumber(value *string)()
    SetState(value ImportedWindowsAutopilotDeviceIdentityStateable)()
}
