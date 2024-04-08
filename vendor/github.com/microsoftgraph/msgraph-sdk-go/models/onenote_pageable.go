package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenotePageable 
type OnenotePageable interface {
    OnenoteEntitySchemaObjectModelable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContent()([]byte)
    GetContentUrl()(*string)
    GetCreatedByAppId()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLevel()(*int32)
    GetLinks()(PageLinksable)
    GetOrder()(*int32)
    GetParentNotebook()(Notebookable)
    GetParentSection()(OnenoteSectionable)
    GetTitle()(*string)
    GetUserTags()([]string)
    SetContent(value []byte)()
    SetContentUrl(value *string)()
    SetCreatedByAppId(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLevel(value *int32)()
    SetLinks(value PageLinksable)()
    SetOrder(value *int32)()
    SetParentNotebook(value Notebookable)()
    SetParentSection(value OnenoteSectionable)()
    SetTitle(value *string)()
    SetUserTags(value []string)()
}
