package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsMobileMSIable 
type WindowsMobileMSIable interface {
    MobileLobAppable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCommandLine()(*string)
    GetIgnoreVersionDetection()(*bool)
    GetProductCode()(*string)
    GetProductVersion()(*string)
    SetCommandLine(value *string)()
    SetIgnoreVersionDetection(value *bool)()
    SetProductCode(value *string)()
    SetProductVersion(value *string)()
}
