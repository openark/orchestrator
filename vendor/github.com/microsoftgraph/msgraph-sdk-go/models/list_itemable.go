package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ListItemable 
type ListItemable interface {
    BaseItemable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAnalytics()(ItemAnalyticsable)
    GetContentType()(ContentTypeInfoable)
    GetDocumentSetVersions()([]DocumentSetVersionable)
    GetDriveItem()(DriveItemable)
    GetFields()(FieldValueSetable)
    GetSharepointIds()(SharepointIdsable)
    GetVersions()([]ListItemVersionable)
    SetAnalytics(value ItemAnalyticsable)()
    SetContentType(value ContentTypeInfoable)()
    SetDocumentSetVersions(value []DocumentSetVersionable)()
    SetDriveItem(value DriveItemable)()
    SetFields(value FieldValueSetable)()
    SetSharepointIds(value SharepointIdsable)()
    SetVersions(value []ListItemVersionable)()
}
