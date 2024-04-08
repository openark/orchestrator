package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody 
type CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The clientContext property
    clientContext *string
}
// NewCallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody instantiates a new CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody and sets the default values.
func NewCallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody()(*CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody) {
    m := &CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetClientContext gets the clientContext property value. The clientContext property
func (m *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody) GetClientContext()(*string) {
    return m.clientContext
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["clientContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientContext(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("clientContext", m.GetClientContext())
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
func (m *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetClientContext sets the clientContext property value. The clientContext property
func (m *CallsItemParticipantsItemMicrosoftGraphStopHoldMusicStopHoldMusicPostRequestBody) SetClientContext(value *string)() {
    m.clientContext = value
}
