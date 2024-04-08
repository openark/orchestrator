package applications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody 
type ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The passwordCredential property
    passwordCredential iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable
}
// NewItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody instantiates a new ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody and sets the default values.
func NewItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody()(*ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) {
    m := &ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphAddPasswordAddPasswordPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphAddPasswordAddPasswordPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["passwordCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordCredential(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable))
        }
        return nil
    }
    return res
}
// GetPasswordCredential gets the passwordCredential property value. The passwordCredential property
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) GetPasswordCredential()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable) {
    return m.passwordCredential
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("passwordCredential", m.GetPasswordCredential())
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
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPasswordCredential sets the passwordCredential property value. The passwordCredential property
func (m *ItemMicrosoftGraphAddPasswordAddPasswordPostRequestBody) SetPasswordCredential(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable)() {
    m.passwordCredential = value
}
