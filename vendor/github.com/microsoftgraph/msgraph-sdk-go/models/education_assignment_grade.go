package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentGrade 
type EducationAssignmentGrade struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // User who did the grading.
    gradedBy IdentitySetable
    // Moment in time when the grade was applied to this submission object. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    gradedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The OdataType property
    odataType *string
}
// NewEducationAssignmentGrade instantiates a new educationAssignmentGrade and sets the default values.
func NewEducationAssignmentGrade()(*EducationAssignmentGrade) {
    m := &EducationAssignmentGrade{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationAssignmentGradeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentGradeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.educationAssignmentPointsGrade":
                        return NewEducationAssignmentPointsGrade(), nil
                }
            }
        }
    }
    return NewEducationAssignmentGrade(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationAssignmentGrade) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentGrade) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["gradedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGradedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["gradedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGradedDateTime(val)
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
    return res
}
// GetGradedBy gets the gradedBy property value. User who did the grading.
func (m *EducationAssignmentGrade) GetGradedBy()(IdentitySetable) {
    return m.gradedBy
}
// GetGradedDateTime gets the gradedDateTime property value. Moment in time when the grade was applied to this submission object. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignmentGrade) GetGradedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.gradedDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationAssignmentGrade) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *EducationAssignmentGrade) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("gradedBy", m.GetGradedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("gradedDateTime", m.GetGradedDateTime())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationAssignmentGrade) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGradedBy sets the gradedBy property value. User who did the grading.
func (m *EducationAssignmentGrade) SetGradedBy(value IdentitySetable)() {
    m.gradedBy = value
}
// SetGradedDateTime sets the gradedDateTime property value. Moment in time when the grade was applied to this submission object. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignmentGrade) SetGradedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.gradedDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationAssignmentGrade) SetOdataType(value *string)() {
    m.odataType = value
}
