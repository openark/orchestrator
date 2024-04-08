package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamMessagingSettings 
type TeamMessagingSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If set to true, @channel mentions are allowed.
    allowChannelMentions *bool
    // If set to true, owners can delete any message.
    allowOwnerDeleteMessages *bool
    // If set to true, @team mentions are allowed.
    allowTeamMentions *bool
    // If set to true, users can delete their messages.
    allowUserDeleteMessages *bool
    // If set to true, users can edit their messages.
    allowUserEditMessages *bool
    // The OdataType property
    odataType *string
}
// NewTeamMessagingSettings instantiates a new teamMessagingSettings and sets the default values.
func NewTeamMessagingSettings()(*TeamMessagingSettings) {
    m := &TeamMessagingSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamMessagingSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamMessagingSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamMessagingSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamMessagingSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowChannelMentions gets the allowChannelMentions property value. If set to true, @channel mentions are allowed.
func (m *TeamMessagingSettings) GetAllowChannelMentions()(*bool) {
    return m.allowChannelMentions
}
// GetAllowOwnerDeleteMessages gets the allowOwnerDeleteMessages property value. If set to true, owners can delete any message.
func (m *TeamMessagingSettings) GetAllowOwnerDeleteMessages()(*bool) {
    return m.allowOwnerDeleteMessages
}
// GetAllowTeamMentions gets the allowTeamMentions property value. If set to true, @team mentions are allowed.
func (m *TeamMessagingSettings) GetAllowTeamMentions()(*bool) {
    return m.allowTeamMentions
}
// GetAllowUserDeleteMessages gets the allowUserDeleteMessages property value. If set to true, users can delete their messages.
func (m *TeamMessagingSettings) GetAllowUserDeleteMessages()(*bool) {
    return m.allowUserDeleteMessages
}
// GetAllowUserEditMessages gets the allowUserEditMessages property value. If set to true, users can edit their messages.
func (m *TeamMessagingSettings) GetAllowUserEditMessages()(*bool) {
    return m.allowUserEditMessages
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamMessagingSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowChannelMentions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowChannelMentions(val)
        }
        return nil
    }
    res["allowOwnerDeleteMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowOwnerDeleteMessages(val)
        }
        return nil
    }
    res["allowTeamMentions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTeamMentions(val)
        }
        return nil
    }
    res["allowUserDeleteMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUserDeleteMessages(val)
        }
        return nil
    }
    res["allowUserEditMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowUserEditMessages(val)
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamMessagingSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TeamMessagingSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowChannelMentions", m.GetAllowChannelMentions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowOwnerDeleteMessages", m.GetAllowOwnerDeleteMessages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowTeamMentions", m.GetAllowTeamMentions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowUserDeleteMessages", m.GetAllowUserDeleteMessages())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("allowUserEditMessages", m.GetAllowUserEditMessages())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamMessagingSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowChannelMentions sets the allowChannelMentions property value. If set to true, @channel mentions are allowed.
func (m *TeamMessagingSettings) SetAllowChannelMentions(value *bool)() {
    m.allowChannelMentions = value
}
// SetAllowOwnerDeleteMessages sets the allowOwnerDeleteMessages property value. If set to true, owners can delete any message.
func (m *TeamMessagingSettings) SetAllowOwnerDeleteMessages(value *bool)() {
    m.allowOwnerDeleteMessages = value
}
// SetAllowTeamMentions sets the allowTeamMentions property value. If set to true, @team mentions are allowed.
func (m *TeamMessagingSettings) SetAllowTeamMentions(value *bool)() {
    m.allowTeamMentions = value
}
// SetAllowUserDeleteMessages sets the allowUserDeleteMessages property value. If set to true, users can delete their messages.
func (m *TeamMessagingSettings) SetAllowUserDeleteMessages(value *bool)() {
    m.allowUserDeleteMessages = value
}
// SetAllowUserEditMessages sets the allowUserEditMessages property value. If set to true, users can edit their messages.
func (m *TeamMessagingSettings) SetAllowUserEditMessages(value *bool)() {
    m.allowUserEditMessages = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamMessagingSettings) SetOdataType(value *string)() {
    m.odataType = value
}
