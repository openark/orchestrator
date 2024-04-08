package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentPointsGrade 
type EducationAssignmentPointsGrade struct {
    EducationAssignmentGrade
    // Number of points a teacher is giving this submission object.
    points *float32
}
// NewEducationAssignmentPointsGrade instantiates a new EducationAssignmentPointsGrade and sets the default values.
func NewEducationAssignmentPointsGrade()(*EducationAssignmentPointsGrade) {
    m := &EducationAssignmentPointsGrade{
        EducationAssignmentGrade: *NewEducationAssignmentGrade(),
    }
    odataTypeValue := "#microsoft.graph.educationAssignmentPointsGrade"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationAssignmentPointsGradeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentPointsGradeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignmentPointsGrade(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentPointsGrade) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationAssignmentGrade.GetFieldDeserializers()
    res["points"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPoints(val)
        }
        return nil
    }
    return res
}
// GetPoints gets the points property value. Number of points a teacher is giving this submission object.
func (m *EducationAssignmentPointsGrade) GetPoints()(*float32) {
    return m.points
}
// Serialize serializes information the current object
func (m *EducationAssignmentPointsGrade) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationAssignmentGrade.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat32Value("points", m.GetPoints())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPoints sets the points property value. Number of points a teacher is giving this submission object.
func (m *EducationAssignmentPointsGrade) SetPoints(value *float32)() {
    m.points = value
}
