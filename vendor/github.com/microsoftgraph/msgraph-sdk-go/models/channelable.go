package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Channelable 
type Channelable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetFilesFolder()(DriveItemable)
    GetIsFavoriteByDefault()(*bool)
    GetMembers()([]ConversationMemberable)
    GetMembershipType()(*ChannelMembershipType)
    GetMessages()([]ChatMessageable)
    GetSharedWithTeams()([]SharedWithChannelTeamInfoable)
    GetTabs()([]TeamsTabable)
    GetTenantId()(*string)
    GetWebUrl()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetFilesFolder(value DriveItemable)()
    SetIsFavoriteByDefault(value *bool)()
    SetMembers(value []ConversationMemberable)()
    SetMembershipType(value *ChannelMembershipType)()
    SetMessages(value []ChatMessageable)()
    SetSharedWithTeams(value []SharedWithChannelTeamInfoable)()
    SetTabs(value []TeamsTabable)()
    SetTenantId(value *string)()
    SetWebUrl(value *string)()
}
