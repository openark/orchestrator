package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedMobileLobAppable 
type ManagedMobileLobAppable interface {
    ManagedAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommittedContentVersion()(*string)
    GetContentVersions()([]MobileAppContentable)
    GetFileName()(*string)
    GetSize()(*int64)
    SetCommittedContentVersion(value *string)()
    SetContentVersions(value []MobileAppContentable)()
    SetFileName(value *string)()
    SetSize(value *int64)()
}
