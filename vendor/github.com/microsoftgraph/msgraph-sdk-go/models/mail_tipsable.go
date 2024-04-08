package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailTipsable 
type MailTipsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAutomaticReplies()(AutomaticRepliesMailTipsable)
    GetCustomMailTip()(*string)
    GetDeliveryRestricted()(*bool)
    GetEmailAddress()(EmailAddressable)
    GetError()(MailTipsErrorable)
    GetExternalMemberCount()(*int32)
    GetIsModerated()(*bool)
    GetMailboxFull()(*bool)
    GetMaxMessageSize()(*int32)
    GetOdataType()(*string)
    GetRecipientScope()(*RecipientScopeType)
    GetRecipientSuggestions()([]Recipientable)
    GetTotalMemberCount()(*int32)
    SetAutomaticReplies(value AutomaticRepliesMailTipsable)()
    SetCustomMailTip(value *string)()
    SetDeliveryRestricted(value *bool)()
    SetEmailAddress(value EmailAddressable)()
    SetError(value MailTipsErrorable)()
    SetExternalMemberCount(value *int32)()
    SetIsModerated(value *bool)()
    SetMailboxFull(value *bool)()
    SetMaxMessageSize(value *int32)()
    SetOdataType(value *string)()
    SetRecipientScope(value *RecipientScopeType)()
    SetRecipientSuggestions(value []Recipientable)()
    SetTotalMemberCount(value *int32)()
}
