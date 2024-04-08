package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatInfoable 
type ChatInfoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMessageId()(*string)
    GetOdataType()(*string)
    GetReplyChainMessageId()(*string)
    GetThreadId()(*string)
    SetMessageId(value *string)()
    SetOdataType(value *string)()
    SetReplyChainMessageId(value *string)()
    SetThreadId(value *string)()
}
