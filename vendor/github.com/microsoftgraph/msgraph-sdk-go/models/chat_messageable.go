package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessageable 
type ChatMessageable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachments()([]ChatMessageAttachmentable)
    GetBody()(ItemBodyable)
    GetChannelIdentity()(ChannelIdentityable)
    GetChatId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEtag()(*string)
    GetEventDetail()(EventMessageDetailable)
    GetFrom()(ChatMessageFromIdentitySetable)
    GetHostedContents()([]ChatMessageHostedContentable)
    GetImportance()(*ChatMessageImportance)
    GetLastEditedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLocale()(*string)
    GetMentions()([]ChatMessageMentionable)
    GetMessageType()(*ChatMessageType)
    GetPolicyViolation()(ChatMessagePolicyViolationable)
    GetReactions()([]ChatMessageReactionable)
    GetReplies()([]ChatMessageable)
    GetReplyToId()(*string)
    GetSubject()(*string)
    GetSummary()(*string)
    GetWebUrl()(*string)
    SetAttachments(value []ChatMessageAttachmentable)()
    SetBody(value ItemBodyable)()
    SetChannelIdentity(value ChannelIdentityable)()
    SetChatId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEtag(value *string)()
    SetEventDetail(value EventMessageDetailable)()
    SetFrom(value ChatMessageFromIdentitySetable)()
    SetHostedContents(value []ChatMessageHostedContentable)()
    SetImportance(value *ChatMessageImportance)()
    SetLastEditedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLocale(value *string)()
    SetMentions(value []ChatMessageMentionable)()
    SetMessageType(value *ChatMessageType)()
    SetPolicyViolation(value ChatMessagePolicyViolationable)()
    SetReactions(value []ChatMessageReactionable)()
    SetReplies(value []ChatMessageable)()
    SetReplyToId(value *string)()
    SetSubject(value *string)()
    SetSummary(value *string)()
    SetWebUrl(value *string)()
}
