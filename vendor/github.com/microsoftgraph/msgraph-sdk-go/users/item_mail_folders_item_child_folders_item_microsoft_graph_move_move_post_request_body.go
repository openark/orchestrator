package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody 
type ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The DestinationId property
    destinationId *string
}
// NewItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody instantiates a new ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody and sets the default values.
func NewItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody()(*ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody) {
    m := &ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDestinationId gets the destinationId property value. The DestinationId property
func (m *ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody) GetDestinationId()(*string) {
    return m.destinationId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["destinationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationId(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("destinationId", m.GetDestinationId())
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
func (m *ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDestinationId sets the destinationId property value. The DestinationId property
func (m *ItemMailFoldersItemChildFoldersItemMicrosoftGraphMoveMovePostRequestBody) SetDestinationId(value *string)() {
    m.destinationId = value
}
