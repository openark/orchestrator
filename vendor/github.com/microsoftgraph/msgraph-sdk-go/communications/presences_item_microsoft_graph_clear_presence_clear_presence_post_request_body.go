package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody 
type PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The sessionId property
    sessionId *string
}
// NewPresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody instantiates a new PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody and sets the default values.
func NewPresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody()(*PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody) {
    m := &PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["sessionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionId(val)
        }
        return nil
    }
    return res
}
// GetSessionId gets the sessionId property value. The sessionId property
func (m *PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody) GetSessionId()(*string) {
    return m.sessionId
}
// Serialize serializes information the current object
func (m *PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("sessionId", m.GetSessionId())
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
func (m *PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetSessionId sets the sessionId property value. The sessionId property
func (m *PresencesItemMicrosoftGraphClearPresenceClearPresencePostRequestBody) SetSessionId(value *string)() {
    m.sessionId = value
}
