package directory

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody 
type DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The ids property
    ids []string
}
// NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody instantiates a new DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody and sets the default values.
func NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody()(*DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) {
    m := &DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["ids"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIds(res)
        }
        return nil
    }
    return res
}
// GetIds gets the ids property value. The ids property
func (m *DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) GetIds()([]string) {
    return m.ids
}
// Serialize serializes information the current object
func (m *DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetIds() != nil {
        err := writer.WriteCollectionOfStringValues("ids", m.GetIds())
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
func (m *DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIds sets the ids property value. The ids property
func (m *DeletedItemsItemMicrosoftGraphCheckMemberObjectsCheckMemberObjectsPostRequestBody) SetIds(value []string)() {
    m.ids = value
}
