package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatRenamedEventMessageDetailable 
type ChatRenamedEventMessageDetailable interface {
    EventMessageDetailable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChatDisplayName()(*string)
    GetChatId()(*string)
    GetInitiator()(IdentitySetable)
    SetChatDisplayName(value *string)()
    SetChatId(value *string)()
    SetInitiator(value IdentitySetable)()
}
