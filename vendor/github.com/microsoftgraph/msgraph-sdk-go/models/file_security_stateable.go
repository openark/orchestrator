package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileSecurityStateable 
type FileSecurityStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileHash()(FileHashable)
    GetName()(*string)
    GetOdataType()(*string)
    GetPath()(*string)
    GetRiskScore()(*string)
    SetFileHash(value FileHashable)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPath(value *string)()
    SetRiskScore(value *string)()
}
