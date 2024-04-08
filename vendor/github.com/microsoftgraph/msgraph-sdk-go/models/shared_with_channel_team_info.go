package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedWithChannelTeamInfo 
type SharedWithChannelTeamInfo struct {
    TeamInfo
    // A collection of team members who have access to the shared channel.
    allowedMembers []ConversationMemberable
    // Indicates whether the team is the host of the channel.
    isHostTeam *bool
}
// NewSharedWithChannelTeamInfo instantiates a new SharedWithChannelTeamInfo and sets the default values.
func NewSharedWithChannelTeamInfo()(*SharedWithChannelTeamInfo) {
    m := &SharedWithChannelTeamInfo{
        TeamInfo: *NewTeamInfo(),
    }
    return m
}
// CreateSharedWithChannelTeamInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedWithChannelTeamInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedWithChannelTeamInfo(), nil
}
// GetAllowedMembers gets the allowedMembers property value. A collection of team members who have access to the shared channel.
func (m *SharedWithChannelTeamInfo) GetAllowedMembers()([]ConversationMemberable) {
    return m.allowedMembers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedWithChannelTeamInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamInfo.GetFieldDeserializers()
    res["allowedMembers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConversationMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationMemberable, len(val))
            for i, v := range val {
                res[i] = v.(ConversationMemberable)
            }
            m.SetAllowedMembers(res)
        }
        return nil
    }
    res["isHostTeam"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHostTeam(val)
        }
        return nil
    }
    return res
}
// GetIsHostTeam gets the isHostTeam property value. Indicates whether the team is the host of the channel.
func (m *SharedWithChannelTeamInfo) GetIsHostTeam()(*bool) {
    return m.isHostTeam
}
// Serialize serializes information the current object
func (m *SharedWithChannelTeamInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamInfo.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAllowedMembers()))
        for i, v := range m.GetAllowedMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("allowedMembers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isHostTeam", m.GetIsHostTeam())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedMembers sets the allowedMembers property value. A collection of team members who have access to the shared channel.
func (m *SharedWithChannelTeamInfo) SetAllowedMembers(value []ConversationMemberable)() {
    m.allowedMembers = value
}
// SetIsHostTeam sets the isHostTeam property value. Indicates whether the team is the host of the channel.
func (m *SharedWithChannelTeamInfo) SetIsHostTeam(value *bool)() {
    m.isHostTeam = value
}
