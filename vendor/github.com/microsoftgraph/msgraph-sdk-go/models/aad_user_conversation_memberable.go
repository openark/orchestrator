package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AadUserConversationMemberable 
type AadUserConversationMemberable interface {
    ConversationMemberable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetTenantId()(*string)
    GetUser()(Userable)
    GetUserId()(*string)
    SetEmail(value *string)()
    SetTenantId(value *string)()
    SetUser(value Userable)()
    SetUserId(value *string)()
}
