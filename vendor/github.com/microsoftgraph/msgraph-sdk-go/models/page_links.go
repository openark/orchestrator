package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PageLinks 
type PageLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Opens the page in the OneNote native client if it's installed.
    oneNoteClientUrl ExternalLinkable
    // Opens the page in OneNote on the web.
    oneNoteWebUrl ExternalLinkable
}
// NewPageLinks instantiates a new pageLinks and sets the default values.
func NewPageLinks()(*PageLinks) {
    m := &PageLinks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePageLinksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePageLinksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPageLinks(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PageLinks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PageLinks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *PageLinks) GetOdataType()(*string) {
    return m.odataType
}
// GetOneNoteClientUrl gets the oneNoteClientUrl property value. Opens the page in the OneNote native client if it's installed.
func (m *PageLinks) GetOneNoteClientUrl()(ExternalLinkable) {
    return m.oneNoteClientUrl
}
// GetOneNoteWebUrl gets the oneNoteWebUrl property value. Opens the page in OneNote on the web.
func (m *PageLinks) GetOneNoteWebUrl()(ExternalLinkable) {
    return m.oneNoteWebUrl
}
// Serialize serializes information the current object
func (m *PageLinks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *PageLinks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PageLinks) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOneNoteClientUrl sets the oneNoteClientUrl property value. Opens the page in the OneNote native client if it's installed.
func (m *PageLinks) SetOneNoteClientUrl(value ExternalLinkable)() {
    m.oneNoteClientUrl = value
}
// SetOneNoteWebUrl sets the oneNoteWebUrl property value. Opens the page in OneNote on the web.
func (m *PageLinks) SetOneNoteWebUrl(value ExternalLinkable)() {
    m.oneNoteWebUrl = value
}
