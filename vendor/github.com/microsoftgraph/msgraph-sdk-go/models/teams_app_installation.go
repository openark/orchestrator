package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppInstallation 
type TeamsAppInstallation struct {
    Entity
    // The app that is installed.
    teamsApp TeamsAppable
    // The details of this version of the app.
    teamsAppDefinition TeamsAppDefinitionable
}
// NewTeamsAppInstallation instantiates a new teamsAppInstallation and sets the default values.
func NewTeamsAppInstallation()(*TeamsAppInstallation) {
    m := &TeamsAppInstallation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAppInstallationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAppInstallationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.userScopeTeamsAppInstallation":
                        return NewUserScopeTeamsAppInstallation(), nil
                }
            }
        }
    }
    return NewTeamsAppInstallation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAppInstallation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["teamsApp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsApp(val.(TeamsAppable))
        }
        return nil
    }
    res["teamsAppDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsAppDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppDefinition(val.(TeamsAppDefinitionable))
        }
        return nil
    }
    return res
}
// GetTeamsApp gets the teamsApp property value. The app that is installed.
func (m *TeamsAppInstallation) GetTeamsApp()(TeamsAppable) {
    return m.teamsApp
}
// GetTeamsAppDefinition gets the teamsAppDefinition property value. The details of this version of the app.
func (m *TeamsAppInstallation) GetTeamsAppDefinition()(TeamsAppDefinitionable) {
    return m.teamsAppDefinition
}
// Serialize serializes information the current object
func (m *TeamsAppInstallation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("teamsApp", m.GetTeamsApp())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teamsAppDefinition", m.GetTeamsAppDefinition())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTeamsApp sets the teamsApp property value. The app that is installed.
func (m *TeamsAppInstallation) SetTeamsApp(value TeamsAppable)() {
    m.teamsApp = value
}
// SetTeamsAppDefinition sets the teamsAppDefinition property value. The details of this version of the app.
func (m *TeamsAppInstallation) SetTeamsAppDefinition(value TeamsAppDefinitionable)() {
    m.teamsAppDefinition = value
}
