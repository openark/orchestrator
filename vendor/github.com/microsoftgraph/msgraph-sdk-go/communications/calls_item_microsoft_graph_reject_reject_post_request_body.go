package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallsItemMicrosoftGraphRejectRejectPostRequestBody 
type CallsItemMicrosoftGraphRejectRejectPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The callbackUri property
    callbackUri *string
    // The reason property
    reason *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RejectReason
}
// NewCallsItemMicrosoftGraphRejectRejectPostRequestBody instantiates a new CallsItemMicrosoftGraphRejectRejectPostRequestBody and sets the default values.
func NewCallsItemMicrosoftGraphRejectRejectPostRequestBody()(*CallsItemMicrosoftGraphRejectRejectPostRequestBody) {
    m := &CallsItemMicrosoftGraphRejectRejectPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallsItemMicrosoftGraphRejectRejectPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemMicrosoftGraphRejectRejectPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemMicrosoftGraphRejectRejectPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemMicrosoftGraphRejectRejectPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *CallsItemMicrosoftGraphRejectRejectPostRequestBody) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemMicrosoftGraphRejectRejectPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callbackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseRejectReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RejectReason))
        }
        return nil
    }
    return res
}
// GetReason gets the reason property value. The reason property
func (m *CallsItemMicrosoftGraphRejectRejectPostRequestBody) GetReason()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RejectReason) {
    return m.reason
}
// Serialize serializes information the current object
func (m *CallsItemMicrosoftGraphRejectRejectPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    if m.GetReason() != nil {
        cast := (*m.GetReason()).String()
        err := writer.WriteStringValue("reason", &cast)
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
func (m *CallsItemMicrosoftGraphRejectRejectPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *CallsItemMicrosoftGraphRejectRejectPostRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetReason sets the reason property value. The reason property
func (m *CallsItemMicrosoftGraphRejectRejectPostRequestBody) SetReason(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RejectReason)() {
    m.reason = value
}
