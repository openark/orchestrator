package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSchoolable 
type EducationSchoolable interface {
    EducationOrganizationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddress()(PhysicalAddressable)
    GetAdministrativeUnit()(AdministrativeUnitable)
    GetClasses()([]EducationClassable)
    GetCreatedBy()(IdentitySetable)
    GetExternalId()(*string)
    GetExternalPrincipalId()(*string)
    GetFax()(*string)
    GetHighestGrade()(*string)
    GetLowestGrade()(*string)
    GetPhone()(*string)
    GetPrincipalEmail()(*string)
    GetPrincipalName()(*string)
    GetSchoolNumber()(*string)
    GetUsers()([]EducationUserable)
    SetAddress(value PhysicalAddressable)()
    SetAdministrativeUnit(value AdministrativeUnitable)()
    SetClasses(value []EducationClassable)()
    SetCreatedBy(value IdentitySetable)()
    SetExternalId(value *string)()
    SetExternalPrincipalId(value *string)()
    SetFax(value *string)()
    SetHighestGrade(value *string)()
    SetLowestGrade(value *string)()
    SetPhone(value *string)()
    SetPrincipalEmail(value *string)()
    SetPrincipalName(value *string)()
    SetSchoolNumber(value *string)()
    SetUsers(value []EducationUserable)()
}
