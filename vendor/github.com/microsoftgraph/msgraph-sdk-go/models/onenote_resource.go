package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteResource 
type OnenoteResource struct {
    OnenoteEntityBaseModel
    // The content stream
    content []byte
    // The URL for downloading the content
    contentUrl *string
}
// NewOnenoteResource instantiates a new OnenoteResource and sets the default values.
func NewOnenoteResource()(*OnenoteResource) {
    m := &OnenoteResource{
        OnenoteEntityBaseModel: *NewOnenoteEntityBaseModel(),
    }
    odataTypeValue := "#microsoft.graph.onenoteResource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOnenoteResourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteResourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenoteResource(), nil
}
// GetContent gets the content property value. The content stream
func (m *OnenoteResource) GetContent()([]byte) {
    return m.content
}
// GetContentUrl gets the contentUrl property value. The URL for downloading the content
func (m *OnenoteResource) GetContentUrl()(*string) {
    return m.contentUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteResource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OnenoteEntityBaseModel.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *OnenoteResource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OnenoteEntityBaseModel.Serialize(writer)
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
    return nil
}
// SetContent sets the content property value. The content stream
func (m *OnenoteResource) SetContent(value []byte)() {
    m.content = value
}
// SetContentUrl sets the contentUrl property value. The URL for downloading the content
func (m *OnenoteResource) SetContentUrl(value *string)() {
    m.contentUrl = value
}
