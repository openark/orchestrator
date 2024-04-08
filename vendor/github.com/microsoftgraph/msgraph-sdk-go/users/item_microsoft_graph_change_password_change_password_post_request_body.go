package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody 
type ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The currentPassword property
    currentPassword *string
    // The newPassword property
    newPassword *string
}
// NewItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody instantiates a new ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody and sets the default values.
func NewItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody()(*ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) {
    m := &ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphChangePasswordChangePasswordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphChangePasswordChangePasswordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCurrentPassword gets the currentPassword property value. The currentPassword property
func (m *ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) GetCurrentPassword()(*string) {
    return m.currentPassword
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["currentPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentPassword(val)
        }
        return nil
    }
    res["newPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewPassword(val)
        }
        return nil
    }
    return res
}
// GetNewPassword gets the newPassword property value. The newPassword property
func (m *ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) GetNewPassword()(*string) {
    return m.newPassword
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("currentPassword", m.GetCurrentPassword())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newPassword", m.GetNewPassword())
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
func (m *ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCurrentPassword sets the currentPassword property value. The currentPassword property
func (m *ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) SetCurrentPassword(value *string)() {
    m.currentPassword = value
}
// SetNewPassword sets the newPassword property value. The newPassword property
func (m *ItemMicrosoftGraphChangePasswordChangePasswordPostRequestBody) SetNewPassword(value *string)() {
    m.newPassword = value
}
