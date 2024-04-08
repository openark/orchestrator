package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosVppEBookable 
type IosVppEBookable interface {
    ManagedEBookable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleId()(*string)
    GetGenres()([]string)
    GetLanguage()(*string)
    GetSeller()(*string)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    GetVppOrganizationName()(*string)
    GetVppTokenId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetAppleId(value *string)()
    SetGenres(value []string)()
    SetLanguage(value *string)()
    SetSeller(value *string)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
    SetVppOrganizationName(value *string)()
    SetVppTokenId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
