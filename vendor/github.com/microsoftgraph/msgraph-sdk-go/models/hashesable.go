package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Hashesable 
type Hashesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCrc32Hash()(*string)
    GetOdataType()(*string)
    GetQuickXorHash()(*string)
    GetSha1Hash()(*string)
    GetSha256Hash()(*string)
    SetCrc32Hash(value *string)()
    SetOdataType(value *string)()
    SetQuickXorHash(value *string)()
    SetSha1Hash(value *string)()
    SetSha256Hash(value *string)()
}
