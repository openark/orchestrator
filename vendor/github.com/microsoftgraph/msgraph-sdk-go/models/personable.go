package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Personable 
type Personable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBirthday()(*string)
    GetCompanyName()(*string)
    GetDepartment()(*string)
    GetDisplayName()(*string)
    GetGivenName()(*string)
    GetImAddress()(*string)
    GetIsFavorite()(*bool)
    GetJobTitle()(*string)
    GetOfficeLocation()(*string)
    GetPersonNotes()(*string)
    GetPersonType()(PersonTypeable)
    GetPhones()([]Phoneable)
    GetPostalAddresses()([]Locationable)
    GetProfession()(*string)
    GetScoredEmailAddresses()([]ScoredEmailAddressable)
    GetSurname()(*string)
    GetUserPrincipalName()(*string)
    GetWebsites()([]Websiteable)
    GetYomiCompany()(*string)
    SetBirthday(value *string)()
    SetCompanyName(value *string)()
    SetDepartment(value *string)()
    SetDisplayName(value *string)()
    SetGivenName(value *string)()
    SetImAddress(value *string)()
    SetIsFavorite(value *bool)()
    SetJobTitle(value *string)()
    SetOfficeLocation(value *string)()
    SetPersonNotes(value *string)()
    SetPersonType(value PersonTypeable)()
    SetPhones(value []Phoneable)()
    SetPostalAddresses(value []Locationable)()
    SetProfession(value *string)()
    SetScoredEmailAddresses(value []ScoredEmailAddressable)()
    SetSurname(value *string)()
    SetUserPrincipalName(value *string)()
    SetWebsites(value []Websiteable)()
    SetYomiCompany(value *string)()
}
