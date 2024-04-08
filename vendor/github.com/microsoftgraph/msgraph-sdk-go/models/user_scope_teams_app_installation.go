package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserScopeTeamsAppInstallation 
type UserScopeTeamsAppInstallation struct {
    TeamsAppInstallation
    // The chat between the user and Teams app.
    chat Chatable
}
// NewUserScopeTeamsAppInstallation instantiates a new UserScopeTeamsAppInstallation and sets the default values.
func NewUserScopeTeamsAppInstallation()(*UserScopeTeamsAppInstallation) {
    m := &UserScopeTeamsAppInstallation{
        TeamsAppInstallation: *NewTeamsAppInstallation(),
    }
    odataTypeValue := "#microsoft.graph.userScopeTeamsAppInstallation"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserScopeTeamsAppInstallationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserScopeTeamsAppInstallationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserScopeTeamsAppInstallation(), nil
}
// GetChat gets the chat property value. The chat between the user and Teams app.
func (m *UserScopeTeamsAppInstallation) GetChat()(Chatable) {
    return m.chat
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserScopeTeamsAppInstallation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeamsAppInstallation.GetFieldDeserializers()
    res["chat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChat(val.(Chatable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *UserScopeTeamsAppInstallation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeamsAppInstallation.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("chat", m.GetChat())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChat sets the chat property value. The chat between the user and Teams app.
func (m *UserScopeTeamsAppInstallation) SetChat(value Chatable)() {
    m.chat = value
}
