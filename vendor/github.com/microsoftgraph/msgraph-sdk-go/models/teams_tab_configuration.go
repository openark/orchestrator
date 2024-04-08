package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsTabConfiguration 
type TeamsTabConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Url used for rendering tab contents in Teams. Required.
    contentUrl *string
    // Identifier for the entity hosted by the tab provider.
    entityId *string
    // The OdataType property
    odataType *string
    // Url called by Teams client when a Tab is removed using the Teams Client.
    removeUrl *string
    // Url for showing tab contents outside of Teams.
    websiteUrl *string
}
// NewTeamsTabConfiguration instantiates a new teamsTabConfiguration and sets the default values.
func NewTeamsTabConfiguration()(*TeamsTabConfiguration) {
    m := &TeamsTabConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamsTabConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsTabConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsTabConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamsTabConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentUrl gets the contentUrl property value. Url used for rendering tab contents in Teams. Required.
func (m *TeamsTabConfiguration) GetContentUrl()(*string) {
    return m.contentUrl
}
// GetEntityId gets the entityId property value. Identifier for the entity hosted by the tab provider.
func (m *TeamsTabConfiguration) GetEntityId()(*string) {
    return m.entityId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsTabConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
        }
        return nil
    }
    res["entityId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEntityId(val)
        }
        return nil
    }
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
    res["removeUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemoveUrl(val)
        }
        return nil
    }
    res["websiteUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebsiteUrl(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamsTabConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetRemoveUrl gets the removeUrl property value. Url called by Teams client when a Tab is removed using the Teams Client.
func (m *TeamsTabConfiguration) GetRemoveUrl()(*string) {
    return m.removeUrl
}
// GetWebsiteUrl gets the websiteUrl property value. Url for showing tab contents outside of Teams.
func (m *TeamsTabConfiguration) GetWebsiteUrl()(*string) {
    return m.websiteUrl
}
// Serialize serializes information the current object
func (m *TeamsTabConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("entityId", m.GetEntityId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("removeUrl", m.GetRemoveUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("websiteUrl", m.GetWebsiteUrl())
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
func (m *TeamsTabConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentUrl sets the contentUrl property value. Url used for rendering tab contents in Teams. Required.
func (m *TeamsTabConfiguration) SetContentUrl(value *string)() {
    m.contentUrl = value
}
// SetEntityId sets the entityId property value. Identifier for the entity hosted by the tab provider.
func (m *TeamsTabConfiguration) SetEntityId(value *string)() {
    m.entityId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamsTabConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRemoveUrl sets the removeUrl property value. Url called by Teams client when a Tab is removed using the Teams Client.
func (m *TeamsTabConfiguration) SetRemoveUrl(value *string)() {
    m.removeUrl = value
}
// SetWebsiteUrl sets the websiteUrl property value. Url for showing tab contents outside of Teams.
func (m *TeamsTabConfiguration) SetWebsiteUrl(value *string)() {
    m.websiteUrl = value
}
