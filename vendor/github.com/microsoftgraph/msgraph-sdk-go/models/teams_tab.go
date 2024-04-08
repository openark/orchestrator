package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsTab 
type TeamsTab struct {
    Entity
    // Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
    configuration TeamsTabConfigurationable
    // Name of the tab.
    displayName *string
    // The application that is linked to the tab. This cannot be changed after tab creation.
    teamsApp TeamsAppable
    // Deep link URL of the tab instance. Read only.
    webUrl *string
}
// NewTeamsTab instantiates a new teamsTab and sets the default values.
func NewTeamsTab()(*TeamsTab) {
    m := &TeamsTab{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsTabFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsTabFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsTab(), nil
}
// GetConfiguration gets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) GetConfiguration()(TeamsTabConfigurationable) {
    return m.configuration
}
// GetDisplayName gets the displayName property value. Name of the tab.
func (m *TeamsTab) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsTab) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTeamsTabConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(TeamsTabConfigurationable))
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
// GetTeamsApp gets the teamsApp property value. The application that is linked to the tab. This cannot be changed after tab creation.
func (m *TeamsTab) GetTeamsApp()(TeamsAppable) {
    return m.teamsApp
}
// GetWebUrl gets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *TeamsTab) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
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
        err = writer.WriteObjectValue("teamsApp", m.GetTeamsApp())
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
// SetConfiguration sets the configuration property value. Container for custom settings applied to a tab. The tab is considered configured only once this property is set.
func (m *TeamsTab) SetConfiguration(value TeamsTabConfigurationable)() {
    m.configuration = value
}
// SetDisplayName sets the displayName property value. Name of the tab.
func (m *TeamsTab) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetTeamsApp sets the teamsApp property value. The application that is linked to the tab. This cannot be changed after tab creation.
func (m *TeamsTab) SetTeamsApp(value TeamsAppable)() {
    m.teamsApp = value
}
// SetWebUrl sets the webUrl property value. Deep link URL of the tab instance. Read only.
func (m *TeamsTab) SetWebUrl(value *string)() {
    m.webUrl = value
}
