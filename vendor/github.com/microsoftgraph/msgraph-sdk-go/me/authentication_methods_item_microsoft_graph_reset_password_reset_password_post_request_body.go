package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody 
type AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The newPassword property
    newPassword *string
}
// NewAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody instantiates a new AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody and sets the default values.
func NewAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody()(*AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) {
    m := &AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
func (m *AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) GetNewPassword()(*string) {
    return m.newPassword
}
// Serialize serializes information the current object
func (m *AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetNewPassword sets the newPassword property value. The newPassword property
func (m *AuthenticationMethodsItemMicrosoftGraphResetPasswordResetPasswordPostRequestBody) SetNewPassword(value *string)() {
    m.newPassword = value
}
