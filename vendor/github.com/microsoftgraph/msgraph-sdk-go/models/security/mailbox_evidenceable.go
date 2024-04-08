package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailboxEvidenceable 
type MailboxEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetPrimaryAddress()(*string)
    GetUserAccount()(UserAccountable)
    SetDisplayName(value *string)()
    SetPrimaryAddress(value *string)()
    SetUserAccount(value UserAccountable)()
}
