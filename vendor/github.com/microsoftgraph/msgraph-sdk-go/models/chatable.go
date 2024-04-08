package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Chatable 
type Chatable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChatType()(*ChatType)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInstalledApps()([]TeamsAppInstallationable)
    GetLastMessagePreview()(ChatMessageInfoable)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMembers()([]ConversationMemberable)
    GetMessages()([]ChatMessageable)
    GetOnlineMeetingInfo()(TeamworkOnlineMeetingInfoable)
    GetPinnedMessages()([]PinnedChatMessageInfoable)
    GetTabs()([]TeamsTabable)
    GetTenantId()(*string)
    GetTopic()(*string)
    GetViewpoint()(ChatViewpointable)
    GetWebUrl()(*string)
    SetChatType(value *ChatType)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInstalledApps(value []TeamsAppInstallationable)()
    SetLastMessagePreview(value ChatMessageInfoable)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMembers(value []ConversationMemberable)()
    SetMessages(value []ChatMessageable)()
    SetOnlineMeetingInfo(value TeamworkOnlineMeetingInfoable)()
    SetPinnedMessages(value []PinnedChatMessageInfoable)()
    SetTabs(value []TeamsTabable)()
    SetTenantId(value *string)()
    SetTopic(value *string)()
    SetViewpoint(value ChatViewpointable)()
    SetWebUrl(value *string)()
}
