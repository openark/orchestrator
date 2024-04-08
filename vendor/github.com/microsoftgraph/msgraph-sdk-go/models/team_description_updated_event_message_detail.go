package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamDescriptionUpdatedEventMessageDetail 
type TeamDescriptionUpdatedEventMessageDetail struct {
    EventMessageDetail
    // Initiator of the event.
    initiator IdentitySetable
    // The updated description for the team.
    teamDescription *string
    // Unique identifier of the team.
    teamId *string
}
// NewTeamDescriptionUpdatedEventMessageDetail instantiates a new TeamDescriptionUpdatedEventMessageDetail and sets the default values.
func NewTeamDescriptionUpdatedEventMessageDetail()(*TeamDescriptionUpdatedEventMessageDetail) {
    m := &TeamDescriptionUpdatedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.teamDescriptionUpdatedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeamDescriptionUpdatedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamDescriptionUpdatedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamDescriptionUpdatedEventMessageDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamDescriptionUpdatedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *TeamDescriptionUpdatedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// GetTeamDescription gets the teamDescription property value. The updated description for the team.
func (m *TeamDescriptionUpdatedEventMessageDetail) GetTeamDescription()(*string) {
    return m.teamDescription
}
// GetTeamId gets the teamId property value. Unique identifier of the team.
func (m *TeamDescriptionUpdatedEventMessageDetail) GetTeamId()(*string) {
    return m.teamId
}
// Serialize serializes information the current object
func (m *TeamDescriptionUpdatedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *TeamDescriptionUpdatedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
// SetTeamDescription sets the teamDescription property value. The updated description for the team.
func (m *TeamDescriptionUpdatedEventMessageDetail) SetTeamDescription(value *string)() {
    m.teamDescription = value
}
// SetTeamId sets the teamId property value. Unique identifier of the team.
func (m *TeamDescriptionUpdatedEventMessageDetail) SetTeamId(value *string)() {
    m.teamId = value
}
