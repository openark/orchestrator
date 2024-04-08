package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageSecurityState 
type MessageSecurityState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The connectingIP property
    connectingIP *string
    // The deliveryAction property
    deliveryAction *string
    // The deliveryLocation property
    deliveryLocation *string
    // The directionality property
    directionality *string
    // The internetMessageId property
    internetMessageId *string
    // The messageFingerprint property
    messageFingerprint *string
    // The messageReceivedDateTime property
    messageReceivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The messageSubject property
    messageSubject *string
    // The networkMessageId property
    networkMessageId *string
    // The OdataType property
    odataType *string
}
// NewMessageSecurityState instantiates a new messageSecurityState and sets the default values.
func NewMessageSecurityState()(*MessageSecurityState) {
    m := &MessageSecurityState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMessageSecurityStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageSecurityStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessageSecurityState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageSecurityState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetConnectingIP gets the connectingIP property value. The connectingIP property
func (m *MessageSecurityState) GetConnectingIP()(*string) {
    return m.connectingIP
}
// GetDeliveryAction gets the deliveryAction property value. The deliveryAction property
func (m *MessageSecurityState) GetDeliveryAction()(*string) {
    return m.deliveryAction
}
// GetDeliveryLocation gets the deliveryLocation property value. The deliveryLocation property
func (m *MessageSecurityState) GetDeliveryLocation()(*string) {
    return m.deliveryLocation
}
// GetDirectionality gets the directionality property value. The directionality property
func (m *MessageSecurityState) GetDirectionality()(*string) {
    return m.directionality
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MessageSecurityState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["connectingIP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectingIP(val)
        }
        return nil
    }
    res["deliveryAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryAction(val)
        }
        return nil
    }
    res["deliveryLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryLocation(val)
        }
        return nil
    }
    res["directionality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectionality(val)
        }
        return nil
    }
    res["internetMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
        }
        return nil
    }
    res["messageFingerprint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageFingerprint(val)
        }
        return nil
    }
    res["messageReceivedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageReceivedDateTime(val)
        }
        return nil
    }
    res["messageSubject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageSubject(val)
        }
        return nil
    }
    res["networkMessageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNetworkMessageId(val)
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
// GetInternetMessageId gets the internetMessageId property value. The internetMessageId property
func (m *MessageSecurityState) GetInternetMessageId()(*string) {
    return m.internetMessageId
}
// GetMessageFingerprint gets the messageFingerprint property value. The messageFingerprint property
func (m *MessageSecurityState) GetMessageFingerprint()(*string) {
    return m.messageFingerprint
}
// GetMessageReceivedDateTime gets the messageReceivedDateTime property value. The messageReceivedDateTime property
func (m *MessageSecurityState) GetMessageReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.messageReceivedDateTime
}
// GetMessageSubject gets the messageSubject property value. The messageSubject property
func (m *MessageSecurityState) GetMessageSubject()(*string) {
    return m.messageSubject
}
// GetNetworkMessageId gets the networkMessageId property value. The networkMessageId property
func (m *MessageSecurityState) GetNetworkMessageId()(*string) {
    return m.networkMessageId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MessageSecurityState) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MessageSecurityState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("connectingIP", m.GetConnectingIP())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deliveryAction", m.GetDeliveryAction())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deliveryLocation", m.GetDeliveryLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("directionality", m.GetDirectionality())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("messageFingerprint", m.GetMessageFingerprint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("messageReceivedDateTime", m.GetMessageReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("messageSubject", m.GetMessageSubject())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("networkMessageId", m.GetNetworkMessageId())
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
func (m *MessageSecurityState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetConnectingIP sets the connectingIP property value. The connectingIP property
func (m *MessageSecurityState) SetConnectingIP(value *string)() {
    m.connectingIP = value
}
// SetDeliveryAction sets the deliveryAction property value. The deliveryAction property
func (m *MessageSecurityState) SetDeliveryAction(value *string)() {
    m.deliveryAction = value
}
// SetDeliveryLocation sets the deliveryLocation property value. The deliveryLocation property
func (m *MessageSecurityState) SetDeliveryLocation(value *string)() {
    m.deliveryLocation = value
}
// SetDirectionality sets the directionality property value. The directionality property
func (m *MessageSecurityState) SetDirectionality(value *string)() {
    m.directionality = value
}
// SetInternetMessageId sets the internetMessageId property value. The internetMessageId property
func (m *MessageSecurityState) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// SetMessageFingerprint sets the messageFingerprint property value. The messageFingerprint property
func (m *MessageSecurityState) SetMessageFingerprint(value *string)() {
    m.messageFingerprint = value
}
// SetMessageReceivedDateTime sets the messageReceivedDateTime property value. The messageReceivedDateTime property
func (m *MessageSecurityState) SetMessageReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.messageReceivedDateTime = value
}
// SetMessageSubject sets the messageSubject property value. The messageSubject property
func (m *MessageSecurityState) SetMessageSubject(value *string)() {
    m.messageSubject = value
}
// SetNetworkMessageId sets the networkMessageId property value. The networkMessageId property
func (m *MessageSecurityState) SetNetworkMessageId(value *string)() {
    m.networkMessageId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MessageSecurityState) SetOdataType(value *string)() {
    m.odataType = value
}
