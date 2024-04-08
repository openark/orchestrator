package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemPreviewInfo 
type ItemPreviewInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The getUrl property
    getUrl *string
    // The OdataType property
    odataType *string
    // The postParameters property
    postParameters *string
    // The postUrl property
    postUrl *string
}
// NewItemPreviewInfo instantiates a new itemPreviewInfo and sets the default values.
func NewItemPreviewInfo()(*ItemPreviewInfo) {
    m := &ItemPreviewInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemPreviewInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemPreviewInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemPreviewInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemPreviewInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemPreviewInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["getUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGetUrl(val)
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
    res["postParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostParameters(val)
        }
        return nil
    }
    res["postUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostUrl(val)
        }
        return nil
    }
    return res
}
// GetGetUrl gets the getUrl property value. The getUrl property
func (m *ItemPreviewInfo) GetGetUrl()(*string) {
    return m.getUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ItemPreviewInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetPostParameters gets the postParameters property value. The postParameters property
func (m *ItemPreviewInfo) GetPostParameters()(*string) {
    return m.postParameters
}
// GetPostUrl gets the postUrl property value. The postUrl property
func (m *ItemPreviewInfo) GetPostUrl()(*string) {
    return m.postUrl
}
// Serialize serializes information the current object
func (m *ItemPreviewInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("getUrl", m.GetGetUrl())
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
        err := writer.WriteStringValue("postParameters", m.GetPostParameters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postUrl", m.GetPostUrl())
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
func (m *ItemPreviewInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetGetUrl sets the getUrl property value. The getUrl property
func (m *ItemPreviewInfo) SetGetUrl(value *string)() {
    m.getUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ItemPreviewInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPostParameters sets the postParameters property value. The postParameters property
func (m *ItemPreviewInfo) SetPostParameters(value *string)() {
    m.postParameters = value
}
// SetPostUrl sets the postUrl property value. The postUrl property
func (m *ItemPreviewInfo) SetPostUrl(value *string)() {
    m.postUrl = value
}
