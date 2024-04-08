package shares

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse 
type ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The value property
    value *bool
}
// NewItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse instantiates a new ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse and sets the default values.
func NewItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse()(*ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse) {
    m := &ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse) GetValue()(*bool) {
    return m.value
}
// Serialize serializes information the current object
func (m *ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("value", m.GetValue())
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
func (m *ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *ItemListContentTypesItemMicrosoftGraphIsPublishedIsPublishedResponse) SetValue(value *bool)() {
    m.value = value
}
