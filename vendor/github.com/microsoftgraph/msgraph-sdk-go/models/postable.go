package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Postable 
type Postable interface {
    OutlookItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachments()([]Attachmentable)
    GetBody()(ItemBodyable)
    GetConversationId()(*string)
    GetConversationThreadId()(*string)
    GetExtensions()([]Extensionable)
    GetFrom()(Recipientable)
    GetHasAttachments()(*bool)
    GetInReplyTo()(Postable)
    GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable)
    GetNewParticipants()([]Recipientable)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSender()(Recipientable)
    GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable)
    SetAttachments(value []Attachmentable)()
    SetBody(value ItemBodyable)()
    SetConversationId(value *string)()
    SetConversationThreadId(value *string)()
    SetExtensions(value []Extensionable)()
    SetFrom(value Recipientable)()
    SetHasAttachments(value *bool)()
    SetInReplyTo(value Postable)()
    SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)()
    SetNewParticipants(value []Recipientable)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSender(value Recipientable)()
    SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)()
}
