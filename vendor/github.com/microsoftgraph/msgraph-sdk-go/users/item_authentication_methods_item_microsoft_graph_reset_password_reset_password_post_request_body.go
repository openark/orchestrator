package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody 
type ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The newPassword property
    newPassword *string
}
// NewItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody instantiates a new ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody and sets the default values.
func NewItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody()(*ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) {
    m := &ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) GetNewPassword()(*string) {
    return m.newPassword
}
// Serialize serializes information the current object
func (m *ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewPassword sets the newPassword property value. The newPassword property
func (m *ItemAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) SetNewPassword(value *string)() {
    m.newPassword = value
}
