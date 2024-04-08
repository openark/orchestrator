package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RubricLevel 
type RubricLevel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of this rubric level.
    description EducationItemBodyable
    // The name of this rubric level.
    displayName *string
    // Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric.
    grading EducationAssignmentGradeTypeable
    // The ID of this resource.
    levelId *string
    // The OdataType property
    odataType *string
}
// NewRubricLevel instantiates a new rubricLevel and sets the default values.
func NewRubricLevel()(*RubricLevel) {
    m := &RubricLevel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRubricLevelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRubricLevelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRubricLevel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricLevel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of this rubric level.
func (m *RubricLevel) GetDescription()(EducationItemBodyable) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of this rubric level.
func (m *RubricLevel) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RubricLevel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(EducationItemBodyable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["grading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentGradeTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrading(val.(EducationAssignmentGradeTypeable))
        }
        return nil
    }
    res["levelId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevelId(val)
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
// GetGrading gets the grading property value. Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric.
func (m *RubricLevel) GetGrading()(EducationAssignmentGradeTypeable) {
    return m.grading
}
// GetLevelId gets the levelId property value. The ID of this resource.
func (m *RubricLevel) GetLevelId()(*string) {
    return m.levelId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RubricLevel) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *RubricLevel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("grading", m.GetGrading())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("levelId", m.GetLevelId())
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
func (m *RubricLevel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of this rubric level.
func (m *RubricLevel) SetDescription(value EducationItemBodyable)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of this rubric level.
func (m *RubricLevel) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGrading sets the grading property value. Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric.
func (m *RubricLevel) SetGrading(value EducationAssignmentGradeTypeable)() {
    m.grading = value
}
// SetLevelId sets the levelId property value. The ID of this resource.
func (m *RubricLevel) SetLevelId(value *string)() {
    m.levelId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RubricLevel) SetOdataType(value *string)() {
    m.odataType = value
}
