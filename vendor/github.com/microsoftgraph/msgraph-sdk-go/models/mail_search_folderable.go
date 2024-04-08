package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailSearchFolderable 
type MailSearchFolderable interface {
    MailFolderable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFilterQuery()(*string)
    GetIncludeNestedFolders()(*bool)
    GetIsSupported()(*bool)
    GetSourceFolderIds()([]string)
    SetFilterQuery(value *string)()
    SetIncludeNestedFolders(value *bool)()
    SetIsSupported(value *bool)()
    SetSourceFolderIds(value []string)()
}
