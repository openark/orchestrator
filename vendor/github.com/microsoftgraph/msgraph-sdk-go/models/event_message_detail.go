package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EventMessageDetail 
type EventMessageDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
}
// NewEventMessageDetail instantiates a new eventMessageDetail and sets the default values.
func NewEventMessageDetail()(*EventMessageDetail) {
    m := &EventMessageDetail{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.callEndedEventMessageDetail":
                        return NewCallEndedEventMessageDetail(), nil
                    case "#microsoft.graph.callRecordingEventMessageDetail":
                        return NewCallRecordingEventMessageDetail(), nil
                    case "#microsoft.graph.callStartedEventMessageDetail":
                        return NewCallStartedEventMessageDetail(), nil
                    case "#microsoft.graph.callTranscriptEventMessageDetail":
                        return NewCallTranscriptEventMessageDetail(), nil
                    case "#microsoft.graph.channelAddedEventMessageDetail":
                        return NewChannelAddedEventMessageDetail(), nil
                    case "#microsoft.graph.channelDeletedEventMessageDetail":
                        return NewChannelDeletedEventMessageDetail(), nil
                    case "#microsoft.graph.channelDescriptionUpdatedEventMessageDetail":
                        return NewChannelDescriptionUpdatedEventMessageDetail(), nil
                    case "#microsoft.graph.channelRenamedEventMessageDetail":
                        return NewChannelRenamedEventMessageDetail(), nil
                    case "#microsoft.graph.channelSetAsFavoriteByDefaultEventMessageDetail":
                        return NewChannelSetAsFavoriteByDefaultEventMessageDetail(), nil
                    case "#microsoft.graph.channelUnsetAsFavoriteByDefaultEventMessageDetail":
                        return NewChannelUnsetAsFavoriteByDefaultEventMessageDetail(), nil
                    case "#microsoft.graph.chatRenamedEventMessageDetail":
                        return NewChatRenamedEventMessageDetail(), nil
                    case "#microsoft.graph.conversationMemberRoleUpdatedEventMessageDetail":
                        return NewConversationMemberRoleUpdatedEventMessageDetail(), nil
                    case "#microsoft.graph.meetingPolicyUpdatedEventMessageDetail":
                        return NewMeetingPolicyUpdatedEventMessageDetail(), nil
                    case "#microsoft.graph.membersAddedEventMessageDetail":
                        return NewMembersAddedEventMessageDetail(), nil
                    case "#microsoft.graph.membersDeletedEventMessageDetail":
                        return NewMembersDeletedEventMessageDetail(), nil
                    case "#microsoft.graph.membersJoinedEventMessageDetail":
                        return NewMembersJoinedEventMessageDetail(), nil
                    case "#microsoft.graph.membersLeftEventMessageDetail":
                        return NewMembersLeftEventMessageDetail(), nil
                    case "#microsoft.graph.messagePinnedEventMessageDetail":
                        return NewMessagePinnedEventMessageDetail(), nil
                    case "#microsoft.graph.messageUnpinnedEventMessageDetail":
                        return NewMessageUnpinnedEventMessageDetail(), nil
                    case "#microsoft.graph.tabUpdatedEventMessageDetail":
                        return NewTabUpdatedEventMessageDetail(), nil
                    case "#microsoft.graph.teamArchivedEventMessageDetail":
                        return NewTeamArchivedEventMessageDetail(), nil
                    case "#microsoft.graph.teamCreatedEventMessageDetail":
                        return NewTeamCreatedEventMessageDetail(), nil
                    case "#microsoft.graph.teamDescriptionUpdatedEventMessageDetail":
                        return NewTeamDescriptionUpdatedEventMessageDetail(), nil
                    case "#microsoft.graph.teamJoiningDisabledEventMessageDetail":
                        return NewTeamJoiningDisabledEventMessageDetail(), nil
                    case "#microsoft.graph.teamJoiningEnabledEventMessageDetail":
                        return NewTeamJoiningEnabledEventMessageDetail(), nil
                    case "#microsoft.graph.teamRenamedEventMessageDetail":
                        return NewTeamRenamedEventMessageDetail(), nil
                    case "#microsoft.graph.teamsAppInstalledEventMessageDetail":
                        return NewTeamsAppInstalledEventMessageDetail(), nil
                    case "#microsoft.graph.teamsAppRemovedEventMessageDetail":
                        return NewTeamsAppRemovedEventMessageDetail(), nil
                    case "#microsoft.graph.teamsAppUpgradedEventMessageDetail":
                        return NewTeamsAppUpgradedEventMessageDetail(), nil
                    case "#microsoft.graph.teamUnarchivedEventMessageDetail":
                        return NewTeamUnarchivedEventMessageDetail(), nil
                }
            }
        }
    }
    return NewEventMessageDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EventMessageDetail) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EventMessageDetail) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *EventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EventMessageDetail) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EventMessageDetail) SetOdataType(value *string)() {
    m.odataType = value
}
