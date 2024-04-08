package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserFeedback 
type UserFeedback struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The rating property
    rating *UserFeedbackRating
    // The feedback text provided by the user of this endpoint for the session.
    text *string
    // The set of feedback tokens provided by the user of this endpoint for the session. This is a set of Boolean properties. The property names should not be relied upon since they may change depending on what tokens are offered to the user.
    tokens FeedbackTokenSetable
}
// NewUserFeedback instantiates a new userFeedback and sets the default values.
func NewUserFeedback()(*UserFeedback) {
    m := &UserFeedback{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserFeedbackFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserFeedbackFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserFeedback(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserFeedback) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserFeedback) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["rating"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserFeedbackRating)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRating(val.(*UserFeedbackRating))
        }
        return nil
    }
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val)
        }
        return nil
    }
    res["tokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFeedbackTokenSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokens(val.(FeedbackTokenSetable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserFeedback) GetOdataType()(*string) {
    return m.odataType
}
// GetRating gets the rating property value. The rating property
func (m *UserFeedback) GetRating()(*UserFeedbackRating) {
    return m.rating
}
// GetText gets the text property value. The feedback text provided by the user of this endpoint for the session.
func (m *UserFeedback) GetText()(*string) {
    return m.text
}
// GetTokens gets the tokens property value. The set of feedback tokens provided by the user of this endpoint for the session. This is a set of Boolean properties. The property names should not be relied upon since they may change depending on what tokens are offered to the user.
func (m *UserFeedback) GetTokens()(FeedbackTokenSetable) {
    return m.tokens
}
// Serialize serializes information the current object
func (m *UserFeedback) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetRating() != nil {
        cast := (*m.GetRating()).String()
        err := writer.WriteStringValue("rating", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("text", m.GetText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("tokens", m.GetTokens())
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
func (m *UserFeedback) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserFeedback) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRating sets the rating property value. The rating property
func (m *UserFeedback) SetRating(value *UserFeedbackRating)() {
    m.rating = value
}
// SetText sets the text property value. The feedback text provided by the user of this endpoint for the session.
func (m *UserFeedback) SetText(value *string)() {
    m.text = value
}
// SetTokens sets the tokens property value. The set of feedback tokens provided by the user of this endpoint for the session. This is a set of Boolean properties. The property names should not be relied upon since they may change depending on what tokens are offered to the user.
func (m *UserFeedback) SetTokens(value FeedbackTokenSetable)() {
    m.tokens = value
}
