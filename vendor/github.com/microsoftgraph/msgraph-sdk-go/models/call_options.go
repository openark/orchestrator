package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallOptions 
type CallOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether to hide the app after the call is escalated.
    hideBotAfterEscalation *bool
    // Indicates whether content sharing notifications should be enabled for the call.
    isContentSharingNotificationEnabled *bool
    // The OdataType property
    odataType *string
}
// NewCallOptions instantiates a new callOptions and sets the default values.
func NewCallOptions()(*CallOptions) {
    m := &CallOptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.incomingCallOptions":
                        return NewIncomingCallOptions(), nil
                    case "#microsoft.graph.outgoingCallOptions":
                        return NewOutgoingCallOptions(), nil
                }
            }
        }
    }
    return NewCallOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallOptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hideBotAfterEscalation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideBotAfterEscalation(val)
        }
        return nil
    }
    res["isContentSharingNotificationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsContentSharingNotificationEnabled(val)
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
// GetHideBotAfterEscalation gets the hideBotAfterEscalation property value. Indicates whether to hide the app after the call is escalated.
func (m *CallOptions) GetHideBotAfterEscalation()(*bool) {
    return m.hideBotAfterEscalation
}
// GetIsContentSharingNotificationEnabled gets the isContentSharingNotificationEnabled property value. Indicates whether content sharing notifications should be enabled for the call.
func (m *CallOptions) GetIsContentSharingNotificationEnabled()(*bool) {
    return m.isContentSharingNotificationEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CallOptions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CallOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("hideBotAfterEscalation", m.GetHideBotAfterEscalation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isContentSharingNotificationEnabled", m.GetIsContentSharingNotificationEnabled())
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
func (m *CallOptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHideBotAfterEscalation sets the hideBotAfterEscalation property value. Indicates whether to hide the app after the call is escalated.
func (m *CallOptions) SetHideBotAfterEscalation(value *bool)() {
    m.hideBotAfterEscalation = value
}
// SetIsContentSharingNotificationEnabled sets the isContentSharingNotificationEnabled property value. Indicates whether content sharing notifications should be enabled for the call.
func (m *CallOptions) SetIsContentSharingNotificationEnabled(value *bool)() {
    m.isContentSharingNotificationEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CallOptions) SetOdataType(value *string)() {
    m.odataType = value
}
