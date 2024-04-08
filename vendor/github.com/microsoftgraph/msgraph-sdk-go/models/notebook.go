package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Notebook 
type Notebook struct {
    OnenoteEntityHierarchyModel
    // Indicates whether this is the user's default notebook. Read-only.
    isDefault *bool
    // Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
    isShared *bool
    // Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
    links NotebookLinksable
    // The section groups in the notebook. Read-only. Nullable.
    sectionGroups []SectionGroupable
    // The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
    sectionGroupsUrl *string
    // The sections in the notebook. Read-only. Nullable.
    sections []OnenoteSectionable
    // The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
    sectionsUrl *string
    // Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
    userRole *OnenoteUserRole
}
// NewNotebook instantiates a new Notebook and sets the default values.
func NewNotebook()(*Notebook) {
    m := &Notebook{
        OnenoteEntityHierarchyModel: *NewOnenoteEntityHierarchyModel(),
    }
    odataTypeValue := "#microsoft.graph.notebook"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateNotebookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotebookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotebook(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Notebook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnenoteEntityHierarchyModel.GetFieldDeserializers()
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
    res["isShared"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsShared(val)
        }
        return nil
    }
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNotebookLinksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(NotebookLinksable))
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
    res["userRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenoteUserRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserRole(val.(*OnenoteUserRole))
        }
        return nil
    }
    return res
}
// GetIsDefault gets the isDefault property value. Indicates whether this is the user's default notebook. Read-only.
func (m *Notebook) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetIsShared gets the isShared property value. Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
func (m *Notebook) GetIsShared()(*bool) {
    return m.isShared
}
// GetLinks gets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *Notebook) GetLinks()(NotebookLinksable) {
    return m.links
}
// GetSectionGroups gets the sectionGroups property value. The section groups in the notebook. Read-only. Nullable.
func (m *Notebook) GetSectionGroups()([]SectionGroupable) {
    return m.sectionGroups
}
// GetSectionGroupsUrl gets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
func (m *Notebook) GetSectionGroupsUrl()(*string) {
    return m.sectionGroupsUrl
}
// GetSections gets the sections property value. The sections in the notebook. Read-only. Nullable.
func (m *Notebook) GetSections()([]OnenoteSectionable) {
    return m.sections
}
// GetSectionsUrl gets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
func (m *Notebook) GetSectionsUrl()(*string) {
    return m.sectionsUrl
}
// GetUserRole gets the userRole property value. Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
func (m *Notebook) GetUserRole()(*OnenoteUserRole) {
    return m.userRole
}
// Serialize serializes information the current object
func (m *Notebook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnenoteEntityHierarchyModel.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isShared", m.GetIsShared())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("links", m.GetLinks())
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
    if m.GetUserRole() != nil {
        cast := (*m.GetUserRole()).String()
        err = writer.WriteStringValue("userRole", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefault sets the isDefault property value. Indicates whether this is the user's default notebook. Read-only.
func (m *Notebook) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetIsShared sets the isShared property value. Indicates whether the notebook is shared. If true, the contents of the notebook can be seen by people other than the owner. Read-only.
func (m *Notebook) SetIsShared(value *bool)() {
    m.isShared = value
}
// SetLinks sets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote native client if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *Notebook) SetLinks(value NotebookLinksable)() {
    m.links = value
}
// SetSectionGroups sets the sectionGroups property value. The section groups in the notebook. Read-only. Nullable.
func (m *Notebook) SetSectionGroups(value []SectionGroupable)() {
    m.sectionGroups = value
}
// SetSectionGroupsUrl sets the sectionGroupsUrl property value. The URL for the sectionGroups navigation property, which returns all the section groups in the notebook. Read-only.
func (m *Notebook) SetSectionGroupsUrl(value *string)() {
    m.sectionGroupsUrl = value
}
// SetSections sets the sections property value. The sections in the notebook. Read-only. Nullable.
func (m *Notebook) SetSections(value []OnenoteSectionable)() {
    m.sections = value
}
// SetSectionsUrl sets the sectionsUrl property value. The URL for the sections navigation property, which returns all the sections in the notebook. Read-only.
func (m *Notebook) SetSectionsUrl(value *string)() {
    m.sectionsUrl = value
}
// SetUserRole sets the userRole property value. Possible values are: Owner, Contributor, Reader, None. Owner represents owner-level access to the notebook. Contributor represents read/write access to the notebook. Reader represents read-only access to the notebook. Read-only.
func (m *Notebook) SetUserRole(value *OnenoteUserRole)() {
    m.userRole = value
}
