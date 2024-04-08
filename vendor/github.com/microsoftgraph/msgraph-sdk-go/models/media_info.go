package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MediaInfo 
type MediaInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Optional. Used to uniquely identity the resource. If passed in, the prompt uri will be cached against this resourceId as a key.
    resourceId *string
    // Path to the prompt that will be played. Currently supports only Wave file (.wav) format, single-channel, 16-bit samples with a 16,000 (16KHz) sampling rate.
    uri *string
}
// NewMediaInfo instantiates a new mediaInfo and sets the default values.
func NewMediaInfo()(*MediaInfo) {
    m := &MediaInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMediaInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMediaInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MediaInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MediaInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MediaInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetResourceId gets the resourceId property value. Optional. Used to uniquely identity the resource. If passed in, the prompt uri will be cached against this resourceId as a key.
func (m *MediaInfo) GetResourceId()(*string) {
    return m.resourceId
}
// GetUri gets the uri property value. Path to the prompt that will be played. Currently supports only Wave file (.wav) format, single-channel, 16-bit samples with a 16,000 (16KHz) sampling rate.
func (m *MediaInfo) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *MediaInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("uri", m.GetUri())
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
func (m *MediaInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MediaInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResourceId sets the resourceId property value. Optional. Used to uniquely identity the resource. If passed in, the prompt uri will be cached against this resourceId as a key.
func (m *MediaInfo) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetUri sets the uri property value. Path to the prompt that will be played. Currently supports only Wave file (.wav) format, single-channel, 16-bit samples with a 16,000 (16KHz) sampling rate.
func (m *MediaInfo) SetUri(value *string)() {
    m.uri = value
}
