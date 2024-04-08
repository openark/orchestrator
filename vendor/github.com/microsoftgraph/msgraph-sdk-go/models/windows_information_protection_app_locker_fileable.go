package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionAppLockerFileable 
type WindowsInformationProtectionAppLockerFileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetFile()([]byte)
    GetFileHash()(*string)
    GetVersion()(*string)
    SetDisplayName(value *string)()
    SetFile(value []byte)()
    SetFileHash(value *string)()
    SetVersion(value *string)()
}
