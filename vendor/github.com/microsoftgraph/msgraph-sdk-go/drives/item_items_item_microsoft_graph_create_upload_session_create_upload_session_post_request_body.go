package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody 
type ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The item property
    item iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemUploadablePropertiesable
}
// NewItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody instantiates a new ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody and sets the default values.
func NewItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody()(*ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) {
    m := &ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["item"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemUploadablePropertiesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetItem(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemUploadablePropertiesable))
        }
        return nil
    }
    return res
}
// GetItem gets the item property value. The item property
func (m *ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetItem()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemUploadablePropertiesable) {
    return m.item
}
// Serialize serializes information the current object
func (m *ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("item", m.GetItem())
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
func (m *ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetItem sets the item property value. The item property
func (m *ItemItemsItemMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) SetItem(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemUploadablePropertiesable)() {
    m.item = value
}
