package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// MicrosoftGraphSendMailSendMailPostRequestBody 
type MicrosoftGraphSendMailSendMailPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The Message property
    message iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable
    // The SaveToSentItems property
    saveToSentItems *bool
}
// NewMicrosoftGraphSendMailSendMailPostRequestBody instantiates a new MicrosoftGraphSendMailSendMailPostRequestBody and sets the default values.
func NewMicrosoftGraphSendMailSendMailPostRequestBody()(*MicrosoftGraphSendMailSendMailPostRequestBody) {
    m := &MicrosoftGraphSendMailSendMailPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphSendMailSendMailPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphSendMailSendMailPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphSendMailSendMailPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphSendMailSendMailPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphSendMailSendMailPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["message"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessage(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable))
        }
        return nil
    }
    res["saveToSentItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaveToSentItems(val)
        }
        return nil
    }
    return res
}
// GetMessage gets the message property value. The Message property
func (m *MicrosoftGraphSendMailSendMailPostRequestBody) GetMessage()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable) {
    return m.message
}
// GetSaveToSentItems gets the saveToSentItems property value. The SaveToSentItems property
func (m *MicrosoftGraphSendMailSendMailPostRequestBody) GetSaveToSentItems()(*bool) {
    return m.saveToSentItems
}
// Serialize serializes information the current object
func (m *MicrosoftGraphSendMailSendMailPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("saveToSentItems", m.GetSaveToSentItems())
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
func (m *MicrosoftGraphSendMailSendMailPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessage sets the message property value. The Message property
func (m *MicrosoftGraphSendMailSendMailPostRequestBody) SetMessage(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Messageable)() {
    m.message = value
}
// SetSaveToSentItems sets the saveToSentItems property value. The SaveToSentItems property
func (m *MicrosoftGraphSendMailSendMailPostRequestBody) SetSaveToSentItems(value *bool)() {
    m.saveToSentItems = value
}
