package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSLobAppable 
type MacOSLobAppable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBuildNumber()(*string)
    GetBundleId()(*string)
    GetChildApps()([]MacOSLobChildAppable)
    GetIgnoreVersionDetection()(*bool)
    GetInstallAsManaged()(*bool)
    GetMd5Hash()([]string)
    GetMd5HashChunkSize()(*int32)
    GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable)
    GetVersionNumber()(*string)
    SetBuildNumber(value *string)()
    SetBundleId(value *string)()
    SetChildApps(value []MacOSLobChildAppable)()
    SetIgnoreVersionDetection(value *bool)()
    SetInstallAsManaged(value *bool)()
    SetMd5Hash(value []string)()
    SetMd5HashChunkSize(value *int32)()
    SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)()
    SetVersionNumber(value *string)()
}
