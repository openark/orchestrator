package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenotePage 
type OnenotePage struct {
    OnenoteEntitySchemaObjectModel
    // The page's HTML content.
    content []byte
    // The URL for the page's HTML content.  Read-only.
    contentUrl *string
    // The unique identifier of the application that created the page. Read-only.
    createdByAppId *string
    // The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The indentation level of the page. Read-only.
    level *int32
    // Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it 's installed. The oneNoteWebUrl link opens the page in OneNote on the web. Read-only.
    links PageLinksable
    // The order of the page within its parent section. Read-only.
    order *int32
    // The notebook that contains the page.  Read-only.
    parentNotebook Notebookable
    // The section that contains the page. Read-only.
    parentSection OnenoteSectionable
    // The title of the page.
    title *string
    // The userTags property
    userTags []string
}
// NewOnenotePage instantiates a new onenotePage and sets the default values.
func NewOnenotePage()(*OnenotePage) {
    m := &OnenotePage{
        OnenoteEntitySchemaObjectModel: *NewOnenoteEntitySchemaObjectModel(),
    }
    odataTypeValue := "#microsoft.graph.onenotePage"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnenotePageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenotePageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenotePage(), nil
}
// GetContent gets the content property value. The page's HTML content.
func (m *OnenotePage) GetContent()([]byte) {
    return m.content
}
// GetContentUrl gets the contentUrl property value. The URL for the page's HTML content.  Read-only.
func (m *OnenotePage) GetContentUrl()(*string) {
    return m.contentUrl
}
// GetCreatedByAppId gets the createdByAppId property value. The unique identifier of the application that created the page. Read-only.
func (m *OnenotePage) GetCreatedByAppId()(*string) {
    return m.createdByAppId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenotePage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnenoteEntitySchemaObjectModel.GetFieldDeserializers()
    res["content"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContent(val)
        }
        return nil
    }
    res["contentUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentUrl(val)
        }
        return nil
    }
    res["createdByAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByAppId(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["level"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
        }
        return nil
    }
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePageLinksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(PageLinksable))
        }
        return nil
    }
    res["order"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrder(val)
        }
        return nil
    }
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
    res["parentSection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnenoteSectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSection(val.(OnenoteSectionable))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["userTags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUserTags(res)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *OnenotePage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLevel gets the level property value. The indentation level of the page. Read-only.
func (m *OnenotePage) GetLevel()(*int32) {
    return m.level
}
// GetLinks gets the links property value. Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it 's installed. The oneNoteWebUrl link opens the page in OneNote on the web. Read-only.
func (m *OnenotePage) GetLinks()(PageLinksable) {
    return m.links
}
// GetOrder gets the order property value. The order of the page within its parent section. Read-only.
func (m *OnenotePage) GetOrder()(*int32) {
    return m.order
}
// GetParentNotebook gets the parentNotebook property value. The notebook that contains the page.  Read-only.
func (m *OnenotePage) GetParentNotebook()(Notebookable) {
    return m.parentNotebook
}
// GetParentSection gets the parentSection property value. The section that contains the page. Read-only.
func (m *OnenotePage) GetParentSection()(OnenoteSectionable) {
    return m.parentSection
}
// GetTitle gets the title property value. The title of the page.
func (m *OnenotePage) GetTitle()(*string) {
    return m.title
}
// GetUserTags gets the userTags property value. The userTags property
func (m *OnenotePage) GetUserTags()([]string) {
    return m.userTags
}
// Serialize serializes information the current object
func (m *OnenotePage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnenoteEntitySchemaObjectModel.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteByteArrayValue("content", m.GetContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("contentUrl", m.GetContentUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("createdByAppId", m.GetCreatedByAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("level", m.GetLevel())
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
    {
        err = writer.WriteInt32Value("order", m.GetOrder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentNotebook", m.GetParentNotebook())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentSection", m.GetParentSection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    if m.GetUserTags() != nil {
        err = writer.WriteCollectionOfStringValues("userTags", m.GetUserTags())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContent sets the content property value. The page's HTML content.
func (m *OnenotePage) SetContent(value []byte)() {
    m.content = value
}
// SetContentUrl sets the contentUrl property value. The URL for the page's HTML content.  Read-only.
func (m *OnenotePage) SetContentUrl(value *string)() {
    m.contentUrl = value
}
// SetCreatedByAppId sets the createdByAppId property value. The unique identifier of the application that created the page. Read-only.
func (m *OnenotePage) SetCreatedByAppId(value *string)() {
    m.createdByAppId = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *OnenotePage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLevel sets the level property value. The indentation level of the page. Read-only.
func (m *OnenotePage) SetLevel(value *int32)() {
    m.level = value
}
// SetLinks sets the links property value. Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it 's installed. The oneNoteWebUrl link opens the page in OneNote on the web. Read-only.
func (m *OnenotePage) SetLinks(value PageLinksable)() {
    m.links = value
}
// SetOrder sets the order property value. The order of the page within its parent section. Read-only.
func (m *OnenotePage) SetOrder(value *int32)() {
    m.order = value
}
// SetParentNotebook sets the parentNotebook property value. The notebook that contains the page.  Read-only.
func (m *OnenotePage) SetParentNotebook(value Notebookable)() {
    m.parentNotebook = value
}
// SetParentSection sets the parentSection property value. The section that contains the page. Read-only.
func (m *OnenotePage) SetParentSection(value OnenoteSectionable)() {
    m.parentSection = value
}
// SetTitle sets the title property value. The title of the page.
func (m *OnenotePage) SetTitle(value *string)() {
    m.title = value
}
// SetUserTags sets the userTags property value. The userTags property
func (m *OnenotePage) SetUserTags(value []string)() {
    m.userTags = value
}
