package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationStudentable 
type EducationStudentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBirthDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetExternalId()(*string)
    GetGender()(*EducationGender)
    GetGrade()(*string)
    GetGraduationYear()(*string)
    GetOdataType()(*string)
    GetStudentNumber()(*string)
    SetBirthDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetExternalId(value *string)()
    SetGender(value *EducationGender)()
    SetGrade(value *string)()
    SetGraduationYear(value *string)()
    SetOdataType(value *string)()
    SetStudentNumber(value *string)()
}
