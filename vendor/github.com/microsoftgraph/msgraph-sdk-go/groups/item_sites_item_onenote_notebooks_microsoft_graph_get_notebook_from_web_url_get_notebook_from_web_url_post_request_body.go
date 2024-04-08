package groups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody 
type ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The webUrl property
    webUrl *string
}
// NewItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody instantiates a new ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody and sets the default values.
func NewItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody()(*ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) {
    m := &ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetWebUrl gets the webUrl property value. The webUrl property
func (m *ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("webUrl", m.GetWebUrl())
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
func (m *ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWebUrl sets the webUrl property value. The webUrl property
func (m *ItemSitesItemOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) SetWebUrl(value *string)() {
    m.webUrl = value
}
