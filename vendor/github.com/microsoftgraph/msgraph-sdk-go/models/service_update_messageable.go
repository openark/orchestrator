package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceUpdateMessageable 
type ServiceUpdateMessageable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceAnnouncementBaseable
    GetActionRequiredByDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetAttachments()([]ServiceAnnouncementAttachmentable)
    GetAttachmentsArchive()([]byte)
    GetBody()(ItemBodyable)
    GetCategory()(*ServiceUpdateCategory)
    GetHasAttachments()(*bool)
    GetIsMajorChange()(*bool)
    GetServices()([]string)
    GetSeverity()(*ServiceUpdateSeverity)
    GetTags()([]string)
    GetViewPoint()(ServiceUpdateMessageViewpointable)
    SetActionRequiredByDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetAttachments(value []ServiceAnnouncementAttachmentable)()
    SetAttachmentsArchive(value []byte)()
    SetBody(value ItemBodyable)()
    SetCategory(value *ServiceUpdateCategory)()
    SetHasAttachments(value *bool)()
    SetIsMajorChange(value *bool)()
    SetServices(value []string)()
    SetSeverity(value *ServiceUpdateSeverity)()
    SetTags(value []string)()
    SetViewPoint(value ServiceUpdateMessageViewpointable)()
}
