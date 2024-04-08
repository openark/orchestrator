package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftStoreForBusinessAppable 
type MicrosoftStoreForBusinessAppable interface {
    MobileAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLicenseType()(*MicrosoftStoreForBusinessLicenseType)
    GetPackageIdentityName()(*string)
    GetProductKey()(*string)
    GetTotalLicenseCount()(*int32)
    GetUsedLicenseCount()(*int32)
    SetLicenseType(value *MicrosoftStoreForBusinessLicenseType)()
    SetPackageIdentityName(value *string)()
    SetProductKey(value *string)()
    SetTotalLicenseCount(value *int32)()
    SetUsedLicenseCount(value *int32)()
}
