package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Notebookable 
type Notebookable interface {
    OnenoteEntityHierarchyModelable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsDefault()(*bool)
    GetIsShared()(*bool)
    GetLinks()(NotebookLinksable)
    GetSectionGroups()([]SectionGroupable)
    GetSectionGroupsUrl()(*string)
    GetSections()([]OnenoteSectionable)
    GetSectionsUrl()(*string)
    GetUserRole()(*OnenoteUserRole)
    SetIsDefault(value *bool)()
    SetIsShared(value *bool)()
    SetLinks(value NotebookLinksable)()
    SetSectionGroups(value []SectionGroupable)()
    SetSectionGroupsUrl(value *string)()
    SetSections(value []OnenoteSectionable)()
    SetSectionsUrl(value *string)()
    SetUserRole(value *OnenoteUserRole)()
}
