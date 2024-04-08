package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SectionGroup 
type SectionGroup struct {
    OnenoteEntityHierarchyModel
    // The notebook that contains the section group. Read-only.
    parentNotebook Notebookable
    // The section group that contains the section group. Read-only.
    parentSectionGroup SectionGroupable
    // The section groups in the section. Read-only. Nullable.
    sectionGroups []SectionGroupable
    // The URL for the sectionGroups navigation property, which returns all the section groups in the section group. Read-only.
    sectionGroupsUrl *string
    // The sections in the section group. Read-only. Nullable.
    sections []OnenoteSectionable
    // The URL for the sections navigation property, which returns all the sections in the section group. Read-only.
    sectionsUrl *string
}
// NewSectionGroup instantiates a new sectionGroup and sets the default values.
func NewSectionGroup()(*SectionGroup) {
    m := &SectionGroup{
        OnenoteEntityHierarchyModel: *NewOnenoteEntityHierarchyModel(),
    }
    odataTypeValue := "#microsoft.graph.sectionGroup"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSectionGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSectionGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSectionGroup(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SectionGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnenoteEntityHierarchyModel.GetFieldDeserializers()
    res["parentNotebook"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNotebookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentNotebook(val.(Notebookable))
        }
        return nil
    }
    res["parentSectionGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSectionGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSectionGroup(val.(SectionGroupable))
        }
        return nil
    }
    res["sectionGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSectionGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SectionGroupable, len(val))
            for i, v := range val {
                res[i] = v.(SectionGroupable)
            }
            m.SetSectionGroups(res)
        }
        return nil
    }
    res["sectionGroupsUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSectionGroupsUrl(val)
        }
        return nil
    }
    res["sections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnenoteSectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnenoteSectionable, len(val))
            for i, v := range val {
                res[i] = v.(OnenoteSectionable)
            }
            m.SetSections(res)
        }
        return nil
    }
    res["sectionsUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSectionsUrl(val)
        }
        return nil
    }
    return res
}
// GetParentNotebook gets the parentNotebook property value. The notebook that contains the section group. Read-only.
func (m *SectionGroup) GetParentNotebook()(Notebookable) {
    return m.parentNotebook
}
// GetParentSectionGroup gets the parentSectionGroup property value. The section group that contains the section group. Read-only.
func (m *SectionGroup) GetParentSectionGroup()(SectionGroupable) {
    return m.parentSectionGroup
}
// GetSectionGroups gets the sectionGroups property value. The section groups in the section. Read-only. Nullable.
func (m *SectionGroup) GetSectionGroups()([]SectionGroupable) {
    return m.sectionGroups
}
// GetSectionGroupsUrl gets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the section group. Read-only.
func (m *SectionGroup) GetSectionGroupsUrl()(*string) {
    return m.sectionGroupsUrl
}
// GetSections gets the sections property value. The sections in the section group. Read-only. Nullable.
func (m *SectionGroup) GetSections()([]OnenoteSectionable) {
    return m.sections
}
// GetSectionsUrl gets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the section group. Read-only.
func (m *SectionGroup) GetSectionsUrl()(*string) {
    return m.sectionsUrl
}
// Serialize serializes information the current object
func (m *SectionGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnenoteEntityHierarchyModel.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("parentNotebook", m.GetParentNotebook())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentSectionGroup", m.GetParentSectionGroup())
        if err != nil {
            return err
        }
    }
    if m.GetSectionGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSectionGroups()))
        for i, v := range m.GetSectionGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sectionGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sectionGroupsUrl", m.GetSectionGroupsUrl())
        if err != nil {
            return err
        }
    }
    if m.GetSections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSections()))
        for i, v := range m.GetSections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("sectionsUrl", m.GetSectionsUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetParentNotebook sets the parentNotebook property value. The notebook that contains the section group. Read-only.
func (m *SectionGroup) SetParentNotebook(value Notebookable)() {
    m.parentNotebook = value
}
// SetParentSectionGroup sets the parentSectionGroup property value. The section group that contains the section group. Read-only.
func (m *SectionGroup) SetParentSectionGroup(value SectionGroupable)() {
    m.parentSectionGroup = value
}
// SetSectionGroups sets the sectionGroups property value. The section groups in the section. Read-only. Nullable.
func (m *SectionGroup) SetSectionGroups(value []SectionGroupable)() {
    m.sectionGroups = value
}
// SetSectionGroupsUrl sets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the section group. Read-only.
func (m *SectionGroup) SetSectionGroupsUrl(value *string)() {
    m.sectionGroupsUrl = value
}
// SetSections sets the sections property value. The sections in the section group. Read-only. Nullable.
func (m *SectionGroup) SetSections(value []OnenoteSectionable)() {
    m.sections = value
}
// SetSectionsUrl sets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the section group. Read-only.
func (m *SectionGroup) SetSectionsUrl(value *string)() {
    m.sectionsUrl = value
}
