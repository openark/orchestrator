package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsUniversalAppXable 
type WindowsUniversalAppXable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableArchitectures()(*WindowsArchitecture)
    GetApplicableDeviceTypes()(*WindowsDeviceType)
    GetCommittedContainedApps()([]MobileContainedAppable)
    GetIdentityName()(*string)
    GetIdentityPublisherHash()(*string)
    GetIdentityResourceIdentifier()(*string)
    GetIdentityVersion()(*string)
    GetIsBundle()(*bool)
    GetMinimumSupportedOperatingSystem()(WindowsMinimumOperatingSystemable)
    SetApplicableArchitectures(value *WindowsArchitecture)()
    SetApplicableDeviceTypes(value *WindowsDeviceType)()
    SetCommittedContainedApps(value []MobileContainedAppable)()
    SetIdentityName(value *string)()
    SetIdentityPublisherHash(value *string)()
    SetIdentityResourceIdentifier(value *string)()
    SetIdentityVersion(value *string)()
    SetIsBundle(value *bool)()
    SetMinimumSupportedOperatingSystem(value WindowsMinimumOperatingSystemable)()
}
