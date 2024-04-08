package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TodoTaskable 
type TodoTaskable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttachments()([]AttachmentBaseable)
    GetAttachmentSessions()([]AttachmentSessionable)
    GetBody()(ItemBodyable)
    GetBodyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetCategories()([]string)
    GetChecklistItems()([]ChecklistItemable)
    GetCompletedDateTime()(DateTimeTimeZoneable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDueDateTime()(DateTimeTimeZoneable)
    GetExtensions()([]Extensionable)
    GetHasAttachments()(*bool)
    GetImportance()(*Importance)
    GetIsReminderOn()(*bool)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLinkedResources()([]LinkedResourceable)
    GetRecurrence()(PatternedRecurrenceable)
    GetReminderDateTime()(DateTimeTimeZoneable)
    GetStartDateTime()(DateTimeTimeZoneable)
    GetStatus()(*TaskStatus)
    GetTitle()(*string)
    SetAttachments(value []AttachmentBaseable)()
    SetAttachmentSessions(value []AttachmentSessionable)()
    SetBody(value ItemBodyable)()
    SetBodyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetCategories(value []string)()
    SetChecklistItems(value []ChecklistItemable)()
    SetCompletedDateTime(value DateTimeTimeZoneable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDueDateTime(value DateTimeTimeZoneable)()
    SetExtensions(value []Extensionable)()
    SetHasAttachments(value *bool)()
    SetImportance(value *Importance)()
    SetIsReminderOn(value *bool)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLinkedResources(value []LinkedResourceable)()
    SetRecurrence(value PatternedRecurrenceable)()
    SetReminderDateTime(value DateTimeTimeZoneable)()
    SetStartDateTime(value DateTimeTimeZoneable)()
    SetStatus(value *TaskStatus)()
    SetTitle(value *string)()
}
