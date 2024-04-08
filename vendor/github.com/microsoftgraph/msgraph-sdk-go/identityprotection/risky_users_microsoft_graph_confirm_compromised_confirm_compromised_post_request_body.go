package identityprotection

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody 
type RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The userIds property
    userIds []string
}
// NewRiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody instantiates a new RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody and sets the default values.
func NewRiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody()(*RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody) {
    m := &RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["userIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUserIds(res)
        }
        return nil
    }
    return res
}
// GetUserIds gets the userIds property value. The userIds property
func (m *RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody) GetUserIds()([]string) {
    return m.userIds
}
// Serialize serializes information the current object
func (m *RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetUserIds() != nil {
        err := writer.WriteCollectionOfStringValues("userIds", m.GetUserIds())
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
func (m *RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUserIds sets the userIds property value. The userIds property
func (m *RiskyUsersMicrosoftGraphConfirmCompromisedConfirmCompromisedPostRequestBody) SetUserIds(value []string)() {
    m.userIds = value
}
