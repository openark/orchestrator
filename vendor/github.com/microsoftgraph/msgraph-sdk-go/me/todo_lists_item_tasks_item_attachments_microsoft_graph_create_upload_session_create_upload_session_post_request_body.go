package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody 
type TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The attachmentInfo property
    attachmentInfo iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable
}
// NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody instantiates a new TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody and sets the default values.
func NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody()(*TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) {
    m := &TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttachmentInfo gets the attachmentInfo property value. The attachmentInfo property
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetAttachmentInfo()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable) {
    return m.attachmentInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attachmentInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAttachmentInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttachmentInfo(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attachmentInfo", m.GetAttachmentInfo())
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
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttachmentInfo sets the attachmentInfo property value. The attachmentInfo property
func (m *TodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBody) SetAttachmentInfo(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AttachmentInfoable)() {
    m.attachmentInfo = value
}
