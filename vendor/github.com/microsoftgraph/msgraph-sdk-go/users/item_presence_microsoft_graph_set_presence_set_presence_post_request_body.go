package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody 
type ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody struct {
    // The activity property
    activity *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The availability property
    availability *string
    // The expirationDuration property
    expirationDuration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The sessionId property
    sessionId *string
}
// NewItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody instantiates a new ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody and sets the default values.
func NewItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody()(*ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) {
    m := &ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody(), nil
}
// GetActivity gets the activity property value. The activity property
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) GetActivity()(*string) {
    return m.activity
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAvailability gets the availability property value. The availability property
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) GetAvailability()(*string) {
    return m.availability
}
// GetExpirationDuration gets the expirationDuration property value. The expirationDuration property
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) GetExpirationDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.expirationDuration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val)
        }
        return nil
    }
    res["availability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAvailability(val)
        }
        return nil
    }
    res["expirationDuration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDuration(val)
        }
        return nil
    }
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
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) GetSessionId()(*string) {
    return m.sessionId
}
// Serialize serializes information the current object
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activity", m.GetActivity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("availability", m.GetAvailability())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteISODurationValue("expirationDuration", m.GetExpirationDuration())
        if err != nil {
            return err
        }
    }
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
// SetActivity sets the activity property value. The activity property
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) SetActivity(value *string)() {
    m.activity = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAvailability sets the availability property value. The availability property
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) SetAvailability(value *string)() {
    m.availability = value
}
// SetExpirationDuration sets the expirationDuration property value. The expirationDuration property
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) SetExpirationDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.expirationDuration = value
}
// SetSessionId sets the sessionId property value. The sessionId property
func (m *ItemPresenceMicrosoftGraphSetPresenceSetPresencePostRequestBody) SetSessionId(value *string)() {
    m.sessionId = value
}
