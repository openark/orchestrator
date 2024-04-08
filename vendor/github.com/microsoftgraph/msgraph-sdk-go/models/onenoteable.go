package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Onenoteable 
type Onenoteable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNotebooks()([]Notebookable)
    GetOperations()([]OnenoteOperationable)
    GetPages()([]OnenotePageable)
    GetResources()([]OnenoteResourceable)
    GetSectionGroups()([]SectionGroupable)
    GetSections()([]OnenoteSectionable)
    SetNotebooks(value []Notebookable)()
    SetOperations(value []OnenoteOperationable)()
    SetPages(value []OnenotePageable)()
    SetResources(value []OnenoteResourceable)()
    SetSectionGroups(value []SectionGroupable)()
    SetSections(value []OnenoteSectionable)()
}
