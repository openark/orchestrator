package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamCreatedEventMessageDetail 
type TeamCreatedEventMessageDetail struct {
    EventMessageDetail
    // Initiator of the event.
    initiator IdentitySetable
    // Description for the team.
    teamDescription *string
    // Display name of the team.
    teamDisplayName *string
    // Unique identifier of the team.
    teamId *string
}
// NewTeamCreatedEventMessageDetail instantiates a new TeamCreatedEventMessageDetail and sets the default values.
func NewTeamCreatedEventMessageDetail()(*TeamCreatedEventMessageDetail) {
    m := &TeamCreatedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.teamCreatedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeamCreatedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamCreatedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamCreatedEventMessageDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamCreatedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["initiator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
        }
        return nil
    }
    res["teamDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamDescription(val)
        }
        return nil
    }
    res["teamDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamDisplayName(val)
        }
        return nil
    }
    res["teamId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamId(val)
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *TeamCreatedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// GetTeamDescription gets the teamDescription property value. Description for the team.
func (m *TeamCreatedEventMessageDetail) GetTeamDescription()(*string) {
    return m.teamDescription
}
// GetTeamDisplayName gets the teamDisplayName property value. Display name of the team.
func (m *TeamCreatedEventMessageDetail) GetTeamDisplayName()(*string) {
    return m.teamDisplayName
}
// GetTeamId gets the teamId property value. Unique identifier of the team.
func (m *TeamCreatedEventMessageDetail) GetTeamId()(*string) {
    return m.teamId
}
// Serialize serializes information the current object
func (m *TeamCreatedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamDescription", m.GetTeamDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamDisplayName", m.GetTeamDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *TeamCreatedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
// SetTeamDescription sets the teamDescription property value. Description for the team.
func (m *TeamCreatedEventMessageDetail) SetTeamDescription(value *string)() {
    m.teamDescription = value
}
// SetTeamDisplayName sets the teamDisplayName property value. Display name of the team.
func (m *TeamCreatedEventMessageDetail) SetTeamDisplayName(value *string)() {
    m.teamDisplayName = value
}
// SetTeamId sets the teamId property value. Unique identifier of the team.
func (m *TeamCreatedEventMessageDetail) SetTeamId(value *string)() {
    m.teamId = value
}
