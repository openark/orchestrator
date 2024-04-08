package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecentNotebook 
type RecentNotebook struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the notebook.
    displayName *string
    // The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    lastAccessedTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
    links RecentNotebookLinksable
    // The OdataType property
    odataType *string
    // The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive.
    sourceService *OnenoteSourceService
}
// NewRecentNotebook instantiates a new recentNotebook and sets the default values.
func NewRecentNotebook()(*RecentNotebook) {
    m := &RecentNotebook{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecentNotebookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecentNotebookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecentNotebook(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecentNotebook) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. The name of the notebook.
func (m *RecentNotebook) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecentNotebook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["lastAccessedTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastAccessedTime(val)
        }
        return nil
    }
    res["links"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecentNotebookLinksFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinks(val.(RecentNotebookLinksable))
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
    res["sourceService"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnenoteSourceService)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceService(val.(*OnenoteSourceService))
        }
        return nil
    }
    return res
}
// GetLastAccessedTime gets the lastAccessedTime property value. The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *RecentNotebook) GetLastAccessedTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastAccessedTime
}
// GetLinks gets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *RecentNotebook) GetLinks()(RecentNotebookLinksable) {
    return m.links
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RecentNotebook) GetOdataType()(*string) {
    return m.odataType
}
// GetSourceService gets the sourceService property value. The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive.
func (m *RecentNotebook) GetSourceService()(*OnenoteSourceService) {
    return m.sourceService
}
// Serialize serializes information the current object
func (m *RecentNotebook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastAccessedTime", m.GetLastAccessedTime())
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetSourceService() != nil {
        cast := (*m.GetSourceService()).String()
        err := writer.WriteStringValue("sourceService", &cast)
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
func (m *RecentNotebook) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. The name of the notebook.
func (m *RecentNotebook) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastAccessedTime sets the lastAccessedTime property value. The date and time when the notebook was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *RecentNotebook) SetLastAccessedTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastAccessedTime = value
}
// SetLinks sets the links property value. Links for opening the notebook. The oneNoteClientURL link opens the notebook in the OneNote client, if it's installed. The oneNoteWebURL link opens the notebook in OneNote on the web.
func (m *RecentNotebook) SetLinks(value RecentNotebookLinksable)() {
    m.links = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RecentNotebook) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSourceService sets the sourceService property value. The backend store where the Notebook resides, either OneDriveForBusiness or OneDrive.
func (m *RecentNotebook) SetSourceService(value *OnenoteSourceService)() {
    m.sourceService = value
}
