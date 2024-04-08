package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Channel 
type Channel struct {
    Entity
    // Read only. Timestamp at which the channel was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Optional textual description for the channel.
    description *string
    // Channel name as it will appear to the user in Microsoft Teams. The maximum length is 50 characters.
    displayName *string
    // The email address for sending messages to the channel. Read-only.
    email *string
    // Metadata for the location where the channel's files are stored.
    filesFolder DriveItemable
    // Indicates whether the channel should automatically be marked 'favorite' for all members of the team. Can only be set programmatically with Create team. Default: false.
    isFavoriteByDefault *bool
    // A collection of membership records associated with the channel.
    members []ConversationMemberable
    // The type of the channel. Can be set during creation and can't be changed. The possible values are: standard, private, unknownFutureValue, shared. The default value is standard. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: shared.
    membershipType *ChannelMembershipType
    // A collection of all the messages in the channel. A navigation property. Nullable.
    messages []ChatMessageable
    // A collection of teams with which a channel is shared.
    sharedWithTeams []SharedWithChannelTeamInfoable
    // A collection of all the tabs in the channel. A navigation property.
    tabs []TeamsTabable
    // The ID of the Azure Active Directory tenant.
    tenantId *string
    // A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only.
    webUrl *string
}
// NewChannel instantiates a new channel and sets the default values.
func NewChannel()(*Channel) {
    m := &Channel{
        Entity: *NewEntity(),
    }
    return m
}
// CreateChannelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChannelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChannel(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Read only. Timestamp at which the channel was created.
func (m *Channel) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Optional textual description for the channel.
func (m *Channel) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Channel name as it will appear to the user in Microsoft Teams. The maximum length is 50 characters.
func (m *Channel) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email address for sending messages to the channel. Read-only.
func (m *Channel) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Channel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["filesFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilesFolder(val.(DriveItemable))
        }
        return nil
    }
    res["isFavoriteByDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavoriteByDefault(val)
        }
        return nil
    }
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConversationMemberFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationMemberable, len(val))
            for i, v := range val {
                res[i] = v.(ConversationMemberable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["membershipType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChannelMembershipType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMembershipType(val.(*ChannelMembershipType))
        }
        return nil
    }
    res["messages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageable)
            }
            m.SetMessages(res)
        }
        return nil
    }
    res["sharedWithTeams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharedWithChannelTeamInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharedWithChannelTeamInfoable, len(val))
            for i, v := range val {
                res[i] = v.(SharedWithChannelTeamInfoable)
            }
            m.SetSharedWithTeams(res)
        }
        return nil
    }
    res["tabs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTeamsTabFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsTabable, len(val))
            for i, v := range val {
                res[i] = v.(TeamsTabable)
            }
            m.SetTabs(res)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetFilesFolder gets the filesFolder property value. Metadata for the location where the channel's files are stored.
func (m *Channel) GetFilesFolder()(DriveItemable) {
    return m.filesFolder
}
// GetIsFavoriteByDefault gets the isFavoriteByDefault property value. Indicates whether the channel should automatically be marked 'favorite' for all members of the team. Can only be set programmatically with Create team. Default: false.
func (m *Channel) GetIsFavoriteByDefault()(*bool) {
    return m.isFavoriteByDefault
}
// GetMembers gets the members property value. A collection of membership records associated with the channel.
func (m *Channel) GetMembers()([]ConversationMemberable) {
    return m.members
}
// GetMembershipType gets the membershipType property value. The type of the channel. Can be set during creation and can't be changed. The possible values are: standard, private, unknownFutureValue, shared. The default value is standard. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: shared.
func (m *Channel) GetMembershipType()(*ChannelMembershipType) {
    return m.membershipType
}
// GetMessages gets the messages property value. A collection of all the messages in the channel. A navigation property. Nullable.
func (m *Channel) GetMessages()([]ChatMessageable) {
    return m.messages
}
// GetSharedWithTeams gets the sharedWithTeams property value. A collection of teams with which a channel is shared.
func (m *Channel) GetSharedWithTeams()([]SharedWithChannelTeamInfoable) {
    return m.sharedWithTeams
}
// GetTabs gets the tabs property value. A collection of all the tabs in the channel. A navigation property.
func (m *Channel) GetTabs()([]TeamsTabable) {
    return m.tabs
}
// GetTenantId gets the tenantId property value. The ID of the Azure Active Directory tenant.
func (m *Channel) GetTenantId()(*string) {
    return m.tenantId
}
// GetWebUrl gets the webUrl property value. A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only.
func (m *Channel) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *Channel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("filesFolder", m.GetFilesFolder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFavoriteByDefault", m.GetIsFavoriteByDefault())
        if err != nil {
            return err
        }
    }
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMembershipType() != nil {
        cast := (*m.GetMembershipType()).String()
        err = writer.WriteStringValue("membershipType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetMessages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSharedWithTeams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharedWithTeams()))
        for i, v := range m.GetSharedWithTeams() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sharedWithTeams", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTabs() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTabs()))
        for i, v := range m.GetTabs() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tabs", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Read only. Timestamp at which the channel was created.
func (m *Channel) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Optional textual description for the channel.
func (m *Channel) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Channel name as it will appear to the user in Microsoft Teams. The maximum length is 50 characters.
func (m *Channel) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email address for sending messages to the channel. Read-only.
func (m *Channel) SetEmail(value *string)() {
    m.email = value
}
// SetFilesFolder sets the filesFolder property value. Metadata for the location where the channel's files are stored.
func (m *Channel) SetFilesFolder(value DriveItemable)() {
    m.filesFolder = value
}
// SetIsFavoriteByDefault sets the isFavoriteByDefault property value. Indicates whether the channel should automatically be marked 'favorite' for all members of the team. Can only be set programmatically with Create team. Default: false.
func (m *Channel) SetIsFavoriteByDefault(value *bool)() {
    m.isFavoriteByDefault = value
}
// SetMembers sets the members property value. A collection of membership records associated with the channel.
func (m *Channel) SetMembers(value []ConversationMemberable)() {
    m.members = value
}
// SetMembershipType sets the membershipType property value. The type of the channel. Can be set during creation and can't be changed. The possible values are: standard, private, unknownFutureValue, shared. The default value is standard. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: shared.
func (m *Channel) SetMembershipType(value *ChannelMembershipType)() {
    m.membershipType = value
}
// SetMessages sets the messages property value. A collection of all the messages in the channel. A navigation property. Nullable.
func (m *Channel) SetMessages(value []ChatMessageable)() {
    m.messages = value
}
// SetSharedWithTeams sets the sharedWithTeams property value. A collection of teams with which a channel is shared.
func (m *Channel) SetSharedWithTeams(value []SharedWithChannelTeamInfoable)() {
    m.sharedWithTeams = value
}
// SetTabs sets the tabs property value. A collection of all the tabs in the channel. A navigation property.
func (m *Channel) SetTabs(value []TeamsTabable)() {
    m.tabs = value
}
// SetTenantId sets the tenantId property value. The ID of the Azure Active Directory tenant.
func (m *Channel) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetWebUrl sets the webUrl property value. A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only.
func (m *Channel) SetWebUrl(value *string)() {
    m.webUrl = value
}
