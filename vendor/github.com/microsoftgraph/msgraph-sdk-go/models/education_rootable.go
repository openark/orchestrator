package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationRootable 
type EducationRootable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetClasses()([]EducationClassable)
    GetMe()(EducationUserable)
    GetOdataType()(*string)
    GetSchools()([]EducationSchoolable)
    GetUsers()([]EducationUserable)
    SetClasses(value []EducationClassable)()
    SetMe(value EducationUserable)()
    SetOdataType(value *string)()
    SetSchools(value []EducationSchoolable)()
    SetUsers(value []EducationUserable)()
}
