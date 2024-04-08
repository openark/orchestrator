package admin

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody 
type ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The messageIds property
    messageIds []string
}
// NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody instantiates a new ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody and sets the default values.
func NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody()(*ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody) {
    m := &ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["messageIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMessageIds(res)
        }
        return nil
    }
    return res
}
// GetMessageIds gets the messageIds property value. The messageIds property
func (m *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody) GetMessageIds()([]string) {
    return m.messageIds
}
// Serialize serializes information the current object
func (m *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMessageIds() != nil {
        err := writer.WriteCollectionOfStringValues("messageIds", m.GetMessageIds())
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
func (m *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMessageIds sets the messageIds property value. The messageIds property
func (m *ServiceAnnouncementMessagesMicrosoftGraphUnfavoriteUnfavoritePostRequestBody) SetMessageIds(value []string)() {
    m.messageIds = value
}
