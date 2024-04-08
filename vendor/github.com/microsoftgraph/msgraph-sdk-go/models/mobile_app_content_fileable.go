package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MobileAppContentFileable 
type MobileAppContentFileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureStorageUri()(*string)
    GetAzureStorageUriExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIsCommitted()(*bool)
    GetManifest()([]byte)
    GetName()(*string)
    GetSize()(*int64)
    GetSizeEncrypted()(*int64)
    GetUploadState()(*MobileAppContentFileUploadState)
    SetAzureStorageUri(value *string)()
    SetAzureStorageUriExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIsCommitted(value *bool)()
    SetManifest(value []byte)()
    SetName(value *string)()
    SetSize(value *int64)()
    SetSizeEncrypted(value *int64)()
    SetUploadState(value *MobileAppContentFileUploadState)()
}
