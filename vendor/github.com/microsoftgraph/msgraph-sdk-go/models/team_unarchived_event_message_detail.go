package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamUnarchivedEventMessageDetail 
type TeamUnarchivedEventMessageDetail struct {
    EventMessageDetail
    // Initiator of the event.
    initiator IdentitySetable
    // Unique identifier of the team.
    teamId *string
}
// NewTeamUnarchivedEventMessageDetail instantiates a new TeamUnarchivedEventMessageDetail and sets the default values.
func NewTeamUnarchivedEventMessageDetail()(*TeamUnarchivedEventMessageDetail) {
    m := &TeamUnarchivedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.teamUnarchivedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeamUnarchivedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamUnarchivedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamUnarchivedEventMessageDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamUnarchivedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *TeamUnarchivedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// GetTeamId gets the teamId property value. Unique identifier of the team.
func (m *TeamUnarchivedEventMessageDetail) GetTeamId()(*string) {
    return m.teamId
}
// Serialize serializes information the current object
func (m *TeamUnarchivedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("teamId", m.GetTeamId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *TeamUnarchivedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
// SetTeamId sets the teamId property value. Unique identifier of the team.
func (m *TeamUnarchivedEventMessageDetail) SetTeamId(value *string)() {
    m.teamId = value
}
