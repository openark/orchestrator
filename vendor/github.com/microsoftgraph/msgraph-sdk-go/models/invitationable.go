package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Invitationable 
type Invitationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetInvitedUser()(Userable)
    GetInvitedUserDisplayName()(*string)
    GetInvitedUserEmailAddress()(*string)
    GetInvitedUserMessageInfo()(InvitedUserMessageInfoable)
    GetInvitedUserType()(*string)
    GetInviteRedeemUrl()(*string)
    GetInviteRedirectUrl()(*string)
    GetResetRedemption()(*bool)
    GetSendInvitationMessage()(*bool)
    GetStatus()(*string)
    SetInvitedUser(value Userable)()
    SetInvitedUserDisplayName(value *string)()
    SetInvitedUserEmailAddress(value *string)()
    SetInvitedUserMessageInfo(value InvitedUserMessageInfoable)()
    SetInvitedUserType(value *string)()
    SetInviteRedeemUrl(value *string)()
    SetInviteRedirectUrl(value *string)()
    SetResetRedemption(value *bool)()
    SetSendInvitationMessage(value *bool)()
    SetStatus(value *string)()
}
