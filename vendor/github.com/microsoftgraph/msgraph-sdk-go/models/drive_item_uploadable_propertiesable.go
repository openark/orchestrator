package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DriveItemUploadablePropertiesable 
type DriveItemUploadablePropertiesable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetFileSize()(*int64)
    GetFileSystemInfo()(FileSystemInfoable)
    GetName()(*string)
    GetOdataType()(*string)
    SetDescription(value *string)()
    SetFileSize(value *int64)()
    SetFileSystemInfo(value FileSystemInfoable)()
    SetName(value *string)()
    SetOdataType(value *string)()
}
