package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RubricQualityFeedbackModel 
type RubricQualityFeedbackModel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specific feedback for one quality of this rubric.
    feedback EducationItemBodyable
    // The OdataType property
    odataType *string
    // The ID of the rubricQuality that this feedback is related to.
    qualityId *string
}
// NewRubricQualityFeedbackModel instantiates a new rubricQualityFeedbackModel and sets the default values.
func NewRubricQualityFeedbackModel()(*RubricQualityFeedbackModel) {
    m := &RubricQualityFeedbackModel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRubricQualityFeedbackModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRubricQualityFeedbackModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRubricQualityFeedbackModel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricQualityFeedbackModel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFeedback gets the feedback property value. Specific feedback for one quality of this rubric.
func (m *RubricQualityFeedbackModel) GetFeedback()(EducationItemBodyable) {
    return m.feedback
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RubricQualityFeedbackModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["feedback"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedback(val.(EducationItemBodyable))
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
    res["qualityId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RubricQualityFeedbackModel) GetOdataType()(*string) {
    return m.odataType
}
// GetQualityId gets the qualityId property value. The ID of the rubricQuality that this feedback is related to.
func (m *RubricQualityFeedbackModel) GetQualityId()(*string) {
    return m.qualityId
}
// Serialize serializes information the current object
func (m *RubricQualityFeedbackModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("feedback", m.GetFeedback())
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
        err := writer.WriteStringValue("qualityId", m.GetQualityId())
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
func (m *RubricQualityFeedbackModel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFeedback sets the feedback property value. Specific feedback for one quality of this rubric.
func (m *RubricQualityFeedbackModel) SetFeedback(value EducationItemBodyable)() {
    m.feedback = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RubricQualityFeedbackModel) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQualityId sets the qualityId property value. The ID of the rubricQuality that this feedback is related to.
func (m *RubricQualityFeedbackModel) SetQualityId(value *string)() {
    m.qualityId = value
}
