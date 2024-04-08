package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotebookLinks 
type NotebookLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Opens the notebook in the OneNote native client if it's installed.
    oneNoteClientUrl ExternalLinkable
    // Opens the notebook in OneNote on the web.
    oneNoteWebUrl ExternalLinkable
}
// NewNotebookLinks instantiates a new notebookLinks and sets the default values.
func NewNotebookLinks()(*NotebookLinks) {
    m := &NotebookLinks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateNotebookLinksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotebookLinksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotebookLinks(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NotebookLinks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotebookLinks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["oneNoteClientUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNoteClientUrl(val.(ExternalLinkable))
        }
        return nil
    }
    res["oneNoteWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOneNoteWebUrl(val.(ExternalLinkable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *NotebookLinks) GetOdataType()(*string) {
    return m.odataType
}
// GetOneNoteClientUrl gets the oneNoteClientUrl property value. Opens the notebook in the OneNote native client if it's installed.
func (m *NotebookLinks) GetOneNoteClientUrl()(ExternalLinkable) {
    return m.oneNoteClientUrl
}
// GetOneNoteWebUrl gets the oneNoteWebUrl property value. Opens the notebook in OneNote on the web.
func (m *NotebookLinks) GetOneNoteWebUrl()(ExternalLinkable) {
    return m.oneNoteWebUrl
}
// Serialize serializes information the current object
func (m *NotebookLinks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("oneNoteClientUrl", m.GetOneNoteClientUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("oneNoteWebUrl", m.GetOneNoteWebUrl())
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
func (m *NotebookLinks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *NotebookLinks) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOneNoteClientUrl sets the oneNoteClientUrl property value. Opens the notebook in the OneNote native client if it's installed.
func (m *NotebookLinks) SetOneNoteClientUrl(value ExternalLinkable)() {
    m.oneNoteClientUrl = value
}
// SetOneNoteWebUrl sets the oneNoteWebUrl property value. Opens the notebook in OneNote on the web.
func (m *NotebookLinks) SetOneNoteWebUrl(value ExternalLinkable)() {
    m.oneNoteWebUrl = value
}
