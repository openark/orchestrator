package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RemoteItemable 
type RemoteItemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCreatedBy()(IdentitySetable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFile()(Fileable)
    GetFileSystemInfo()(FileSystemInfoable)
    GetFolder()(Folderable)
    GetId()(*string)
    GetImage()(Imageable)
    GetLastModifiedBy()(IdentitySetable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetName()(*string)
    GetOdataType()(*string)
    GetPackage()(PackageEscapedable)
    GetParentReference()(ItemReferenceable)
    GetShared()(Sharedable)
    GetSharepointIds()(SharepointIdsable)
    GetSize()(*int64)
    GetSpecialFolder()(SpecialFolderable)
    GetVideo()(Videoable)
    GetWebDavUrl()(*string)
    GetWebUrl()(*string)
    SetCreatedBy(value IdentitySetable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFile(value Fileable)()
    SetFileSystemInfo(value FileSystemInfoable)()
    SetFolder(value Folderable)()
    SetId(value *string)()
    SetImage(value Imageable)()
    SetLastModifiedBy(value IdentitySetable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetName(value *string)()
    SetOdataType(value *string)()
    SetPackage(value PackageEscapedable)()
    SetParentReference(value ItemReferenceable)()
    SetShared(value Sharedable)()
    SetSharepointIds(value SharepointIdsable)()
    SetSize(value *int64)()
    SetSpecialFolder(value SpecialFolderable)()
    SetVideo(value Videoable)()
    SetWebDavUrl(value *string)()
    SetWebUrl(value *string)()
}
