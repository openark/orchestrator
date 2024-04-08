package drives

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody 
type ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The checkInAs property
    checkInAs *string
    // The comment property
    comment *string
}
// NewItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody instantiates a new ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody and sets the default values.
func NewItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody()(*ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) {
    m := &ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCheckInAs gets the checkInAs property value. The checkInAs property
func (m *ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) GetCheckInAs()(*string) {
    return m.checkInAs
}
// GetComment gets the comment property value. The comment property
func (m *ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checkInAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckInAs(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("checkInAs", m.GetCheckInAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comment", m.GetComment())
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
func (m *ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCheckInAs sets the checkInAs property value. The checkInAs property
func (m *ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) SetCheckInAs(value *string)() {
    m.checkInAs = value
}
// SetComment sets the comment property value. The comment property
func (m *ItemItemsItemMicrosoftGraphCheckinCheckinPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
