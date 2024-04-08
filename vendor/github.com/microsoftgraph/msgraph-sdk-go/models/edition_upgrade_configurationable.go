package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EditionUpgradeConfigurationable 
type EditionUpgradeConfigurationable interface {
    DeviceConfigurationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLicense()(*string)
    GetLicenseType()(*EditionUpgradeLicenseType)
    GetProductKey()(*string)
    GetTargetEdition()(*Windows10EditionType)
    SetLicense(value *string)()
    SetLicenseType(value *EditionUpgradeLicenseType)()
    SetProductKey(value *string)()
    SetTargetEdition(value *Windows10EditionType)()
}
