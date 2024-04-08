package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DocumentSetable 
type DocumentSetable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowedContentTypes()([]ContentTypeInfoable)
    GetDefaultContents()([]DocumentSetContentable)
    GetOdataType()(*string)
    GetPropagateWelcomePageChanges()(*bool)
    GetSharedColumns()([]ColumnDefinitionable)
    GetShouldPrefixNameToFile()(*bool)
    GetWelcomePageColumns()([]ColumnDefinitionable)
    GetWelcomePageUrl()(*string)
    SetAllowedContentTypes(value []ContentTypeInfoable)()
    SetDefaultContents(value []DocumentSetContentable)()
    SetOdataType(value *string)()
    SetPropagateWelcomePageChanges(value *bool)()
    SetSharedColumns(value []ColumnDefinitionable)()
    SetShouldPrefixNameToFile(value *bool)()
    SetWelcomePageColumns(value []ColumnDefinitionable)()
    SetWelcomePageUrl(value *string)()
}
