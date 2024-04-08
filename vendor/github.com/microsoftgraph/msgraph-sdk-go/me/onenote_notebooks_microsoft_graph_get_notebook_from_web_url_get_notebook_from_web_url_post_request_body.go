package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody 
type OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The webUrl property
    webUrl *string
}
// NewOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody instantiates a new OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody and sets the default values.
func NewOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody()(*OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) {
    m := &OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetWebUrl sets the webUrl property value. The webUrl property
func (m *OnenoteNotebooksMicrosoftGraphGetNotebookFromWebUrlGetNotebookFromWebUrlPostRequestBody) SetWebUrl(value *string)() {
    m.webUrl = value
}
