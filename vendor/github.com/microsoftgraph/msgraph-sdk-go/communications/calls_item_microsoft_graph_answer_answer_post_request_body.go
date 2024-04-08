package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallsItemMicrosoftGraphAnswerAnswerPostRequestBody 
type CallsItemMicrosoftGraphAnswerAnswerPostRequestBody struct {
    // The acceptedModalities property
    acceptedModalities []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Modality
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The callbackUri property
    callbackUri *string
    // The callOptions property
    callOptions iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable
    // The mediaConfig property
    mediaConfig iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable
    // The participantCapacity property
    participantCapacity *int32
}
// NewCallsItemMicrosoftGraphAnswerAnswerPostRequestBody instantiates a new CallsItemMicrosoftGraphAnswerAnswerPostRequestBody and sets the default values.
func NewCallsItemMicrosoftGraphAnswerAnswerPostRequestBody()(*CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) {
    m := &CallsItemMicrosoftGraphAnswerAnswerPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallsItemMicrosoftGraphAnswerAnswerPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsItemMicrosoftGraphAnswerAnswerPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsItemMicrosoftGraphAnswerAnswerPostRequestBody(), nil
}
// GetAcceptedModalities gets the acceptedModalities property value. The acceptedModalities property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) GetAcceptedModalities()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Modality) {
    return m.acceptedModalities
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCallbackUri gets the callbackUri property value. The callbackUri property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetCallOptions gets the callOptions property value. The callOptions property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) GetCallOptions()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable) {
    return m.callOptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["acceptedModalities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Modality, len(val))
            for i, v := range val {
                res[i] = *(v.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Modality))
            }
            m.SetAcceptedModalities(res)
        }
        return nil
    }
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
    res["callOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateIncomingCallOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallOptions(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable))
        }
        return nil
    }
    res["mediaConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMediaConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaConfig(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable))
        }
        return nil
    }
    res["participantCapacity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipantCapacity(val)
        }
        return nil
    }
    return res
}
// GetMediaConfig gets the mediaConfig property value. The mediaConfig property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) GetMediaConfig()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable) {
    return m.mediaConfig
}
// GetParticipantCapacity gets the participantCapacity property value. The participantCapacity property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) GetParticipantCapacity()(*int32) {
    return m.participantCapacity
}
// Serialize serializes information the current object
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAcceptedModalities() != nil {
        err := writer.WriteCollectionOfStringValues("acceptedModalities", iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SerializeModality(m.GetAcceptedModalities()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("callOptions", m.GetCallOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mediaConfig", m.GetMediaConfig())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("participantCapacity", m.GetParticipantCapacity())
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
// SetAcceptedModalities sets the acceptedModalities property value. The acceptedModalities property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) SetAcceptedModalities(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Modality)() {
    m.acceptedModalities = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCallbackUri sets the callbackUri property value. The callbackUri property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetCallOptions sets the callOptions property value. The callOptions property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) SetCallOptions(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable)() {
    m.callOptions = value
}
// SetMediaConfig sets the mediaConfig property value. The mediaConfig property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) SetMediaConfig(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable)() {
    m.mediaConfig = value
}
// SetParticipantCapacity sets the participantCapacity property value. The participantCapacity property
func (m *CallsItemMicrosoftGraphAnswerAnswerPostRequestBody) SetParticipantCapacity(value *int32)() {
    m.participantCapacity = value
}
