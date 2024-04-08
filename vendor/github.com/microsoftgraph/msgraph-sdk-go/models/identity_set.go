package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentitySet 
type IdentitySet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Optional. The application associated with this action.
    application Identityable
    // Optional. The device associated with this action.
    device Identityable
    // The OdataType property
    odataType *string
    // Optional. The user associated with this action.
    user Identityable
}
// NewIdentitySet instantiates a new identitySet and sets the default values.
func NewIdentitySet()(*IdentitySet) {
    m := &IdentitySet{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.chatMessageFromIdentitySet":
                        return NewChatMessageFromIdentitySet(), nil
                    case "#microsoft.graph.chatMessageMentionedIdentitySet":
                        return NewChatMessageMentionedIdentitySet(), nil
                    case "#microsoft.graph.chatMessageReactionIdentitySet":
                        return NewChatMessageReactionIdentitySet(), nil
                    case "#microsoft.graph.sharePointIdentitySet":
                        return NewSharePointIdentitySet(), nil
                }
            }
        }
    }
    return NewIdentitySet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *IdentitySet) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApplication gets the application property value. Optional. The application associated with this action.
func (m *IdentitySet) GetApplication()(Identityable) {
    return m.application
}
// GetDevice gets the device property value. Optional. The device associated with this action.
func (m *IdentitySet) GetDevice()(Identityable) {
    return m.device
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(Identityable))
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(Identityable))
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
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(Identityable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *IdentitySet) GetOdataType()(*string) {
    return m.odataType
}
// GetUser gets the user property value. Optional. The user associated with this action.
func (m *IdentitySet) GetUser()(Identityable) {
    return m.user
}
// Serialize serializes information the current object
func (m *IdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("device", m.GetDevice())
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
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *IdentitySet) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApplication sets the application property value. Optional. The application associated with this action.
func (m *IdentitySet) SetApplication(value Identityable)() {
    m.application = value
}
// SetDevice sets the device property value. Optional. The device associated with this action.
func (m *IdentitySet) SetDevice(value Identityable)() {
    m.device = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *IdentitySet) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUser sets the user property value. Optional. The user associated with this action.
func (m *IdentitySet) SetUser(value Identityable)() {
    m.user = value
}
