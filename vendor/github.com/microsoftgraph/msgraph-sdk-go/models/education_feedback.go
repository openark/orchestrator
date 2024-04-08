package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationFeedback 
type EducationFeedback struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // User who created the feedback.
    feedbackBy IdentitySetable
    // Moment in time when the feedback was given. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    feedbackDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
    // Feedback.
    text EducationItemBodyable
}
// NewEducationFeedback instantiates a new educationFeedback and sets the default values.
func NewEducationFeedback()(*EducationFeedback) {
    m := &EducationFeedback{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationFeedbackFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationFeedbackFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationFeedback(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationFeedback) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFeedbackBy gets the feedbackBy property value. User who created the feedback.
func (m *EducationFeedback) GetFeedbackBy()(IdentitySetable) {
    return m.feedbackBy
}
// GetFeedbackDateTime gets the feedbackDateTime property value. Moment in time when the feedback was given. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationFeedback) GetFeedbackDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.feedbackDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationFeedback) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["feedbackBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedbackBy(val.(IdentitySetable))
        }
        return nil
    }
    res["feedbackDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedbackDateTime(val)
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
    res["text"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetText(val.(EducationItemBodyable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationFeedback) GetOdataType()(*string) {
    return m.odataType
}
// GetText gets the text property value. Feedback.
func (m *EducationFeedback) GetText()(EducationItemBodyable) {
    return m.text
}
// Serialize serializes information the current object
func (m *EducationFeedback) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("feedbackBy", m.GetFeedbackBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("feedbackDateTime", m.GetFeedbackDateTime())
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
        err := writer.WriteObjectValue("text", m.GetText())
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
func (m *EducationFeedback) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFeedbackBy sets the feedbackBy property value. User who created the feedback.
func (m *EducationFeedback) SetFeedbackBy(value IdentitySetable)() {
    m.feedbackBy = value
}
// SetFeedbackDateTime sets the feedbackDateTime property value. Moment in time when the feedback was given. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationFeedback) SetFeedbackDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.feedbackDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationFeedback) SetOdataType(value *string)() {
    m.odataType = value
}
// SetText sets the text property value. Feedback.
func (m *EducationFeedback) SetText(value EducationItemBodyable)() {
    m.text = value
}
