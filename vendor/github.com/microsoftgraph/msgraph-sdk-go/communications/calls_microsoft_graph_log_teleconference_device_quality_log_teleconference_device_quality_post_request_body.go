package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody 
type CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The quality property
    quality iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable
}
// NewCallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody instantiates a new CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody and sets the default values.
func NewCallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody()(*CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody) {
    m := &CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["quality"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeleconferenceDeviceQualityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuality(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable))
        }
        return nil
    }
    return res
}
// GetQuality gets the quality property value. The quality property
func (m *CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody) GetQuality()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable) {
    return m.quality
}
// Serialize serializes information the current object
func (m *CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("quality", m.GetQuality())
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
func (m *CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetQuality sets the quality property value. The quality property
func (m *CallsMicrosoftGraphLogTeleconferenceDeviceQualityLogTeleconferenceDeviceQualityPostRequestBody) SetQuality(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeleconferenceDeviceQualityable)() {
    m.quality = value
}
