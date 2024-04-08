package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedIOSLobAppable 
type ManagedIOSLobAppable interface {
    ManagedMobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicableDeviceType()(IosDeviceTypeable)
    GetBuildNumber()(*string)
    GetBundleId()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMinimumSupportedOperatingSystem()(IosMinimumOperatingSystemable)
    GetVersionNumber()(*string)
    SetApplicableDeviceType(value IosDeviceTypeable)()
    SetBuildNumber(value *string)()
    SetBundleId(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMinimumSupportedOperatingSystem(value IosMinimumOperatingSystemable)()
    SetVersionNumber(value *string)()
}
