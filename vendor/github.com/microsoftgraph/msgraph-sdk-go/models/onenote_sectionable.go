package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteSectionable 
type OnenoteSectionable interface {
    OnenoteEntityHierarchyModelable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsDefault()(*bool)
    GetLinks()(SectionLinksable)
    GetPages()([]OnenotePageable)
    GetPagesUrl()(*string)
    GetParentNotebook()(Notebookable)
    GetParentSectionGroup()(SectionGroupable)
    SetIsDefault(value *bool)()
    SetLinks(value SectionLinksable)()
    SetPages(value []OnenotePageable)()
    SetPagesUrl(value *string)()
    SetParentNotebook(value Notebookable)()
    SetParentSectionGroup(value SectionGroupable)()
}
