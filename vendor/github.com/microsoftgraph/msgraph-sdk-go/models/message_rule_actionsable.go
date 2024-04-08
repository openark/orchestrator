package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MessageRuleActionsable 
type MessageRuleActionsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAssignCategories()([]string)
    GetCopyToFolder()(*string)
    GetDelete()(*bool)
    GetForwardAsAttachmentTo()([]Recipientable)
    GetForwardTo()([]Recipientable)
    GetMarkAsRead()(*bool)
    GetMarkImportance()(*Importance)
    GetMoveToFolder()(*string)
    GetOdataType()(*string)
    GetPermanentDelete()(*bool)
    GetRedirectTo()([]Recipientable)
    GetStopProcessingRules()(*bool)
    SetAssignCategories(value []string)()
    SetCopyToFolder(value *string)()
    SetDelete(value *bool)()
    SetForwardAsAttachmentTo(value []Recipientable)()
    SetForwardTo(value []Recipientable)()
    SetMarkAsRead(value *bool)()
    SetMarkImportance(value *Importance)()
    SetMoveToFolder(value *string)()
    SetOdataType(value *string)()
    SetPermanentDelete(value *bool)()
    SetRedirectTo(value []Recipientable)()
    SetStopProcessingRules(value *bool)()
}
