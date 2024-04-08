package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenotePagePreviewLinks 
type OnenotePagePreviewLinks struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The previewImageUrl property
    previewImageUrl ExternalLinkable
}
// NewOnenotePagePreviewLinks instantiates a new onenotePagePreviewLinks and sets the default values.
func NewOnenotePagePreviewLinks()(*OnenotePagePreviewLinks) {
    m := &OnenotePagePreviewLinks{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnenotePagePreviewLinksFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenotePagePreviewLinksFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenotePagePreviewLinks(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePagePreviewLinks) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenotePagePreviewLinks) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["previewImageUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExternalLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewImageUrl(val.(ExternalLinkable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *OnenotePagePreviewLinks) GetOdataType()(*string) {
    return m.odataType
}
// GetPreviewImageUrl gets the previewImageUrl property value. The previewImageUrl property
func (m *OnenotePagePreviewLinks) GetPreviewImageUrl()(ExternalLinkable) {
    return m.previewImageUrl
}
// Serialize serializes information the current object
func (m *OnenotePagePreviewLinks) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewImageUrl", m.GetPreviewImageUrl())
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
func (m *OnenotePagePreviewLinks) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnenotePagePreviewLinks) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPreviewImageUrl sets the previewImageUrl property value. The previewImageUrl property
func (m *OnenotePagePreviewLinks) SetPreviewImageUrl(value ExternalLinkable)() {
    m.previewImageUrl = value
}
