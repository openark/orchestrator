package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppUpgradedEventMessageDetail 
type TeamsAppUpgradedEventMessageDetail struct {
    EventMessageDetail
    // Initiator of the event.
    initiator IdentitySetable
    // Display name of the teamsApp.
    teamsAppDisplayName *string
    // Unique identifier of the teamsApp.
    teamsAppId *string
}
// NewTeamsAppUpgradedEventMessageDetail instantiates a new TeamsAppUpgradedEventMessageDetail and sets the default values.
func NewTeamsAppUpgradedEventMessageDetail()(*TeamsAppUpgradedEventMessageDetail) {
    m := &TeamsAppUpgradedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.teamsAppUpgradedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeamsAppUpgradedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppUpgradedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAppUpgradedEventMessageDetail(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppUpgradedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["teamsAppDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppDisplayName(val)
        }
        return nil
    }
    res["teamsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *TeamsAppUpgradedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// GetTeamsAppDisplayName gets the teamsAppDisplayName property value. Display name of the teamsApp.
func (m *TeamsAppUpgradedEventMessageDetail) GetTeamsAppDisplayName()(*string) {
    return m.teamsAppDisplayName
}
// GetTeamsAppId gets the teamsAppId property value. Unique identifier of the teamsApp.
func (m *TeamsAppUpgradedEventMessageDetail) GetTeamsAppId()(*string) {
    return m.teamsAppId
}
// Serialize serializes information the current object
func (m *TeamsAppUpgradedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("teamsAppDisplayName", m.GetTeamsAppDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *TeamsAppUpgradedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
// SetTeamsAppDisplayName sets the teamsAppDisplayName property value. Display name of the teamsApp.
func (m *TeamsAppUpgradedEventMessageDetail) SetTeamsAppDisplayName(value *string)() {
    m.teamsAppDisplayName = value
}
// SetTeamsAppId sets the teamsAppId property value. Unique identifier of the teamsApp.
func (m *TeamsAppUpgradedEventMessageDetail) SetTeamsAppId(value *string)() {
    m.teamsAppId = value
}
