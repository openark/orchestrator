package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatViewpoint 
type ChatViewpoint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether the chat is hidden for the current user.
    isHidden *bool
    // Represents the dateTime up until which the current user has read chatMessages in a specific chat.
    lastMessageReadDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
}
// NewChatViewpoint instantiates a new chatViewpoint and sets the default values.
func NewChatViewpoint()(*ChatViewpoint) {
    m := &ChatViewpoint{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChatViewpointFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatViewpointFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatViewpoint(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatViewpoint) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatViewpoint) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isHidden"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsHidden(val)
        }
        return nil
    }
    res["lastMessageReadDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastMessageReadDateTime(val)
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
// GetIsHidden gets the isHidden property value. Indicates whether the chat is hidden for the current user.
func (m *ChatViewpoint) GetIsHidden()(*bool) {
    return m.isHidden
}
// GetLastMessageReadDateTime gets the lastMessageReadDateTime property value. Represents the dateTime up until which the current user has read chatMessages in a specific chat.
func (m *ChatViewpoint) GetLastMessageReadDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastMessageReadDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ChatViewpoint) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ChatViewpoint) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isHidden", m.GetIsHidden())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastMessageReadDateTime", m.GetLastMessageReadDateTime())
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
func (m *ChatViewpoint) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsHidden sets the isHidden property value. Indicates whether the chat is hidden for the current user.
func (m *ChatViewpoint) SetIsHidden(value *bool)() {
    m.isHidden = value
}
// SetLastMessageReadDateTime sets the lastMessageReadDateTime property value. Represents the dateTime up until which the current user has read chatMessages in a specific chat.
func (m *ChatViewpoint) SetLastMessageReadDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastMessageReadDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChatViewpoint) SetOdataType(value *string)() {
    m.odataType = value
}
