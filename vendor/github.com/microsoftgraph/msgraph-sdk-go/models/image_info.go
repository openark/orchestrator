package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImageInfo 
type ImageInfo struct {
    // Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
    addImageQuery *bool
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Optional; alt-text accessible content for the image
    alternateText *string
    // The alternativeText property
    alternativeText *string
    // Optional; URI that points to an icon which represents the application used to generate the activity
    iconUrl *string
    // The OdataType property
    odataType *string
}
// NewImageInfo instantiates a new imageInfo and sets the default values.
func NewImageInfo()(*ImageInfo) {
    m := &ImageInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateImageInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImageInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImageInfo(), nil
}
// GetAddImageQuery gets the addImageQuery property value. Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
func (m *ImageInfo) GetAddImageQuery()(*bool) {
    return m.addImageQuery
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImageInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlternateText gets the alternateText property value. Optional; alt-text accessible content for the image
func (m *ImageInfo) GetAlternateText()(*string) {
    return m.alternateText
}
// GetAlternativeText gets the alternativeText property value. The alternativeText property
func (m *ImageInfo) GetAlternativeText()(*string) {
    return m.alternativeText
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImageInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addImageQuery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddImageQuery(val)
        }
        return nil
    }
    res["alternateText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternateText(val)
        }
        return nil
    }
    res["alternativeText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlternativeText(val)
        }
        return nil
    }
    res["iconUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIconUrl(val)
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
    return res
}
// GetIconUrl gets the iconUrl property value. Optional; URI that points to an icon which represents the application used to generate the activity
func (m *ImageInfo) GetIconUrl()(*string) {
    return m.iconUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ImageInfo) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ImageInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("addImageQuery", m.GetAddImageQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alternateText", m.GetAlternateText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("alternativeText", m.GetAlternativeText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("iconUrl", m.GetIconUrl())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddImageQuery sets the addImageQuery property value. Optional; parameter used to indicate the server is able to render image dynamically in response to parameterization. For example – a high contrast image
func (m *ImageInfo) SetAddImageQuery(value *bool)() {
    m.addImageQuery = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ImageInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlternateText sets the alternateText property value. Optional; alt-text accessible content for the image
func (m *ImageInfo) SetAlternateText(value *string)() {
    m.alternateText = value
}
// SetAlternativeText sets the alternativeText property value. The alternativeText property
func (m *ImageInfo) SetAlternativeText(value *string)() {
    m.alternativeText = value
}
// SetIconUrl sets the iconUrl property value. Optional; URI that points to an icon which represents the application used to generate the activity
func (m *ImageInfo) SetIconUrl(value *string)() {
    m.iconUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ImageInfo) SetOdataType(value *string)() {
    m.odataType = value
}
