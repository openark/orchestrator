package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CopyNotebookModel 
type CopyNotebookModel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The createdBy property
    createdBy *string
    // The createdByIdentity property
    createdByIdentity IdentitySetable
    // The createdTime property
    createdTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The id property
    id *string
    // The isDefault property
    isDefault *bool
    // The isShared property
    isShared *bool
    // The lastModifiedBy property
    lastModifiedBy *string
    // The lastModifiedByIdentity property
    lastModifiedByIdentity IdentitySetable
    // The lastModifiedTime property
    lastModifiedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The links property
    links NotebookLinksable
    // The name property
    name *string
    // The OdataType property
    odataType *string
    // The sectionGroupsUrl property
    sectionGroupsUrl *string
    // The sectionsUrl property
    sectionsUrl *string
    // The self property
    self *string
    // The userRole property
    userRole *OnenoteUserRole
}
// NewCopyNotebookModel instantiates a new CopyNotebookModel and sets the default values.
func NewCopyNotebookModel()(*CopyNotebookModel) {
    m := &CopyNotebookModel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCopyNotebookModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCopyNotebookModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCopyNotebookModel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyNotebookModel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *CopyNotebookModel) GetCreatedBy()(*string) {
    return m.createdBy
}
// GetCreatedByIdentity gets the createdByIdentity property value. The createdByIdentity property
func (m *CopyNotebookModel) GetCreatedByIdentity()(IdentitySetable) {
    return m.createdByIdentity
}
// GetCreatedTime gets the createdTime property value. The createdTime property
func (m *CopyNotebookModel) GetCreatedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CopyNotebookModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val)
        }
        return nil
    }
    res["createdByIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByIdentity(val.(IdentitySetable))
        }
        return nil
    }
    res["createdTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedTime(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
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
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val)
        }
        return nil
    }
    res["lastModifiedByIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedByIdentity(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedTime(val)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
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
    res["self"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val)
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
// GetId gets the id property value. The id property
func (m *CopyNotebookModel) GetId()(*string) {
    return m.id
}
// GetIsDefault gets the isDefault property value. The isDefault property
func (m *CopyNotebookModel) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetIsShared gets the isShared property value. The isShared property
func (m *CopyNotebookModel) GetIsShared()(*bool) {
    return m.isShared
}
// GetLastModifiedBy gets the lastModifiedBy property value. The lastModifiedBy property
func (m *CopyNotebookModel) GetLastModifiedBy()(*string) {
    return m.lastModifiedBy
}
// GetLastModifiedByIdentity gets the lastModifiedByIdentity property value. The lastModifiedByIdentity property
func (m *CopyNotebookModel) GetLastModifiedByIdentity()(IdentitySetable) {
    return m.lastModifiedByIdentity
}
// GetLastModifiedTime gets the lastModifiedTime property value. The lastModifiedTime property
func (m *CopyNotebookModel) GetLastModifiedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedTime
}
// GetLinks gets the links property value. The links property
func (m *CopyNotebookModel) GetLinks()(NotebookLinksable) {
    return m.links
}
// GetName gets the name property value. The name property
func (m *CopyNotebookModel) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CopyNotebookModel) GetOdataType()(*string) {
    return m.odataType
}
// GetSectionGroupsUrl gets the sectionGroupsUrl property value. The sectionGroupsUrl property
func (m *CopyNotebookModel) GetSectionGroupsUrl()(*string) {
    return m.sectionGroupsUrl
}
// GetSectionsUrl gets the sectionsUrl property value. The sectionsUrl property
func (m *CopyNotebookModel) GetSectionsUrl()(*string) {
    return m.sectionsUrl
}
// GetSelf gets the self property value. The self property
func (m *CopyNotebookModel) GetSelf()(*string) {
    return m.self
}
// GetUserRole gets the userRole property value. The userRole property
func (m *CopyNotebookModel) GetUserRole()(*OnenoteUserRole) {
    return m.userRole
}
// Serialize serializes information the current object
func (m *CopyNotebookModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("createdByIdentity", m.GetCreatedByIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdTime", m.GetCreatedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isShared", m.GetIsShared())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("lastModifiedByIdentity", m.GetLastModifiedByIdentity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastModifiedTime", m.GetLastModifiedTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("links", m.GetLinks())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sectionGroupsUrl", m.GetSectionGroupsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sectionsUrl", m.GetSectionsUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("self", m.GetSelf())
        if err != nil {
            return err
        }
    }
    if m.GetUserRole() != nil {
        cast := (*m.GetUserRole()).String()
        err := writer.WriteStringValue("userRole", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CopyNotebookModel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *CopyNotebookModel) SetCreatedBy(value *string)() {
    m.createdBy = value
}
// SetCreatedByIdentity sets the createdByIdentity property value. The createdByIdentity property
func (m *CopyNotebookModel) SetCreatedByIdentity(value IdentitySetable)() {
    m.createdByIdentity = value
}
// SetCreatedTime sets the createdTime property value. The createdTime property
func (m *CopyNotebookModel) SetCreatedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdTime = value
}
// SetId sets the id property value. The id property
func (m *CopyNotebookModel) SetId(value *string)() {
    m.id = value
}
// SetIsDefault sets the isDefault property value. The isDefault property
func (m *CopyNotebookModel) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetIsShared sets the isShared property value. The isShared property
func (m *CopyNotebookModel) SetIsShared(value *bool)() {
    m.isShared = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. The lastModifiedBy property
func (m *CopyNotebookModel) SetLastModifiedBy(value *string)() {
    m.lastModifiedBy = value
}
// SetLastModifiedByIdentity sets the lastModifiedByIdentity property value. The lastModifiedByIdentity property
func (m *CopyNotebookModel) SetLastModifiedByIdentity(value IdentitySetable)() {
    m.lastModifiedByIdentity = value
}
// SetLastModifiedTime sets the lastModifiedTime property value. The lastModifiedTime property
func (m *CopyNotebookModel) SetLastModifiedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedTime = value
}
// SetLinks sets the links property value. The links property
func (m *CopyNotebookModel) SetLinks(value NotebookLinksable)() {
    m.links = value
}
// SetName sets the name property value. The name property
func (m *CopyNotebookModel) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CopyNotebookModel) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSectionGroupsUrl sets the sectionGroupsUrl property value. The sectionGroupsUrl property
func (m *CopyNotebookModel) SetSectionGroupsUrl(value *string)() {
    m.sectionGroupsUrl = value
}
// SetSectionsUrl sets the sectionsUrl property value. The sectionsUrl property
func (m *CopyNotebookModel) SetSectionsUrl(value *string)() {
    m.sectionsUrl = value
}
// SetSelf sets the self property value. The self property
func (m *CopyNotebookModel) SetSelf(value *string)() {
    m.self = value
}
// SetUserRole sets the userRole property value. The userRole property
func (m *CopyNotebookModel) SetUserRole(value *OnenoteUserRole)() {
    m.userRole = value
}
