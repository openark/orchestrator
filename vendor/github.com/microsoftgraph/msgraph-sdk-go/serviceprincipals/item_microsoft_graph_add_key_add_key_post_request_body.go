package serviceprincipals

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemMicrosoftGraphAddKeyAddKeyPostRequestBody 
type ItemMicrosoftGraphAddKeyAddKeyPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The keyCredential property
    keyCredential iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable
    // The passwordCredential property
    passwordCredential iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable
    // The proof property
    proof *string
}
// NewItemMicrosoftGraphAddKeyAddKeyPostRequestBody instantiates a new ItemMicrosoftGraphAddKeyAddKeyPostRequestBody and sets the default values.
func NewItemMicrosoftGraphAddKeyAddKeyPostRequestBody()(*ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) {
    m := &ItemMicrosoftGraphAddKeyAddKeyPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphAddKeyAddKeyPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphAddKeyAddKeyPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphAddKeyAddKeyPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["keyCredential"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyCredential(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable))
        }
        return nil
    }
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
    res["proof"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProof(val)
        }
        return nil
    }
    return res
}
// GetKeyCredential gets the keyCredential property value. The keyCredential property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetKeyCredential()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable) {
    return m.keyCredential
}
// GetPasswordCredential gets the passwordCredential property value. The passwordCredential property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetPasswordCredential()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable) {
    return m.passwordCredential
}
// GetProof gets the proof property value. The proof property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) GetProof()(*string) {
    return m.proof
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("keyCredential", m.GetKeyCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("passwordCredential", m.GetPasswordCredential())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("proof", m.GetProof())
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
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetKeyCredential sets the keyCredential property value. The keyCredential property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) SetKeyCredential(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyCredentialable)() {
    m.keyCredential = value
}
// SetPasswordCredential sets the passwordCredential property value. The passwordCredential property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) SetPasswordCredential(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PasswordCredentialable)() {
    m.passwordCredential = value
}
// SetProof sets the proof property value. The proof property
func (m *ItemMicrosoftGraphAddKeyAddKeyPostRequestBody) SetProof(value *string)() {
    m.proof = value
}
