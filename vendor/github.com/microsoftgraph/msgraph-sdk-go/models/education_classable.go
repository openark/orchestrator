package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationClassable 
type EducationClassable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignmentCategories()([]EducationCategoryable)
    GetAssignmentDefaults()(EducationAssignmentDefaultsable)
    GetAssignments()([]EducationAssignmentable)
    GetAssignmentSettings()(EducationAssignmentSettingsable)
    GetClassCode()(*string)
    GetCourse()(EducationCourseable)
    GetCreatedBy()(IdentitySetable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExternalId()(*string)
    GetExternalName()(*string)
    GetExternalSource()(*EducationExternalSource)
    GetExternalSourceDetail()(*string)
    GetGrade()(*string)
    GetGroup()(Groupable)
    GetMailNickname()(*string)
    GetMembers()([]EducationUserable)
    GetSchools()([]EducationSchoolable)
    GetTeachers()([]EducationUserable)
    GetTerm()(EducationTermable)
    SetAssignmentCategories(value []EducationCategoryable)()
    SetAssignmentDefaults(value EducationAssignmentDefaultsable)()
    SetAssignments(value []EducationAssignmentable)()
    SetAssignmentSettings(value EducationAssignmentSettingsable)()
    SetClassCode(value *string)()
    SetCourse(value EducationCourseable)()
    SetCreatedBy(value IdentitySetable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExternalId(value *string)()
    SetExternalName(value *string)()
    SetExternalSource(value *EducationExternalSource)()
    SetExternalSourceDetail(value *string)()
    SetGrade(value *string)()
    SetGroup(value Groupable)()
    SetMailNickname(value *string)()
    SetMembers(value []EducationUserable)()
    SetSchools(value []EducationSchoolable)()
    SetTeachers(value []EducationUserable)()
    SetTerm(value EducationTermable)()
}
