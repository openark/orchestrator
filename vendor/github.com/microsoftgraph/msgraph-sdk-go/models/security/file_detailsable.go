package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FileDetailsable 
type FileDetailsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFileName()(*string)
    GetFilePath()(*string)
    GetFilePublisher()(*string)
    GetFileSize()(*int64)
    GetIssuer()(*string)
    GetOdataType()(*string)
    GetSha1()(*string)
    GetSha256()(*string)
    GetSigner()(*string)
    SetFileName(value *string)()
    SetFilePath(value *string)()
    SetFilePublisher(value *string)()
    SetFileSize(value *int64)()
    SetIssuer(value *string)()
    SetOdataType(value *string)()
    SetSha1(value *string)()
    SetSha256(value *string)()
    SetSigner(value *string)()
}
