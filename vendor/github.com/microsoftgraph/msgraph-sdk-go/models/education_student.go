package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationStudent 
type EducationStudent struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Birth date of the student.
    birthDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // ID of the student in the source system.
    externalId *string
    // The possible values are: female, male, other, unknownFutureValue.
    gender *EducationGender
    // Current grade level of the student.
    grade *string
    // Year the student is graduating from the school.
    graduationYear *string
    // The OdataType property
    odataType *string
    // Student Number.
    studentNumber *string
}
// NewEducationStudent instantiates a new educationStudent and sets the default values.
func NewEducationStudent()(*EducationStudent) {
    m := &EducationStudent{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateEducationStudentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationStudentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationStudent(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationStudent) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBirthDate gets the birthDate property value. Birth date of the student.
func (m *EducationStudent) GetBirthDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.birthDate
}
// GetExternalId gets the externalId property value. ID of the student in the source system.
func (m *EducationStudent) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationStudent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["birthDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthDate(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["gender"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationGender)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGender(val.(*EducationGender))
        }
        return nil
    }
    res["grade"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrade(val)
        }
        return nil
    }
    res["graduationYear"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGraduationYear(val)
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
    res["studentNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudentNumber(val)
        }
        return nil
    }
    return res
}
// GetGender gets the gender property value. The possible values are: female, male, other, unknownFutureValue.
func (m *EducationStudent) GetGender()(*EducationGender) {
    return m.gender
}
// GetGrade gets the grade property value. Current grade level of the student.
func (m *EducationStudent) GetGrade()(*string) {
    return m.grade
}
// GetGraduationYear gets the graduationYear property value. Year the student is graduating from the school.
func (m *EducationStudent) GetGraduationYear()(*string) {
    return m.graduationYear
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *EducationStudent) GetOdataType()(*string) {
    return m.odataType
}
// GetStudentNumber gets the studentNumber property value. Student Number.
func (m *EducationStudent) GetStudentNumber()(*string) {
    return m.studentNumber
}
// Serialize serializes information the current object
func (m *EducationStudent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteDateOnlyValue("birthDate", m.GetBirthDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    if m.GetGender() != nil {
        cast := (*m.GetGender()).String()
        err := writer.WriteStringValue("gender", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("grade", m.GetGrade())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("graduationYear", m.GetGraduationYear())
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
        err := writer.WriteStringValue("studentNumber", m.GetStudentNumber())
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
func (m *EducationStudent) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBirthDate sets the birthDate property value. Birth date of the student.
func (m *EducationStudent) SetBirthDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.birthDate = value
}
// SetExternalId sets the externalId property value. ID of the student in the source system.
func (m *EducationStudent) SetExternalId(value *string)() {
    m.externalId = value
}
// SetGender sets the gender property value. The possible values are: female, male, other, unknownFutureValue.
func (m *EducationStudent) SetGender(value *EducationGender)() {
    m.gender = value
}
// SetGrade sets the grade property value. Current grade level of the student.
func (m *EducationStudent) SetGrade(value *string)() {
    m.grade = value
}
// SetGraduationYear sets the graduationYear property value. Year the student is graduating from the school.
func (m *EducationStudent) SetGraduationYear(value *string)() {
    m.graduationYear = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *EducationStudent) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStudentNumber sets the studentNumber property value. Student Number.
func (m *EducationStudent) SetStudentNumber(value *string)() {
    m.studentNumber = value
}
