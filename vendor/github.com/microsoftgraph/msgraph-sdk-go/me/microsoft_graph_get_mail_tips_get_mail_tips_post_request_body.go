package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody 
type MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The EmailAddresses property
    emailAddresses []string
    // The MailTipsOptions property
    mailTipsOptions *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType
}
// NewMicrosoftGraphGetMailTipsGetMailTipsPostRequestBody instantiates a new MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody and sets the default values.
func NewMicrosoftGraphGetMailTipsGetMailTipsPostRequestBody()(*MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) {
    m := &MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphGetMailTipsGetMailTipsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphGetMailTipsGetMailTipsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphGetMailTipsGetMailTipsPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmailAddresses gets the emailAddresses property value. The EmailAddresses property
func (m *MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) GetEmailAddresses()([]string) {
    return m.emailAddresses
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetEmailAddresses(res)
        }
        return nil
    }
    res["mailTipsOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParseMailTipsType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailTipsOptions(val.(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType))
        }
        return nil
    }
    return res
}
// GetMailTipsOptions gets the mailTipsOptions property value. The MailTipsOptions property
func (m *MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) GetMailTipsOptions()(*iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType) {
    return m.mailTipsOptions
}
// Serialize serializes information the current object
func (m *MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetEmailAddresses() != nil {
        err := writer.WriteCollectionOfStringValues("emailAddresses", m.GetEmailAddresses())
        if err != nil {
            return err
        }
    }
    if m.GetMailTipsOptions() != nil {
        cast := (*m.GetMailTipsOptions()).String()
        err := writer.WriteStringValue("mailTipsOptions", &cast)
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
func (m *MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmailAddresses sets the emailAddresses property value. The EmailAddresses property
func (m *MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) SetEmailAddresses(value []string)() {
    m.emailAddresses = value
}
// SetMailTipsOptions sets the mailTipsOptions property value. The MailTipsOptions property
func (m *MicrosoftGraphGetMailTipsGetMailTipsPostRequestBody) SetMailTipsOptions(value *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MailTipsType)() {
    m.mailTipsOptions = value
}
