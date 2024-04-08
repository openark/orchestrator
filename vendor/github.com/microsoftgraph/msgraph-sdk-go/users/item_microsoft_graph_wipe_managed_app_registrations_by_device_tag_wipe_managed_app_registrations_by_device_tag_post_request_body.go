package users

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody 
type ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceTag property
    deviceTag *string
}
// NewItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody instantiates a new ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody and sets the default values.
func NewItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody()(*ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) {
    m := &ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceTag gets the deviceTag property value. The deviceTag property
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) GetDeviceTag()(*string) {
    return m.deviceTag
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("deviceTag", m.GetDeviceTag())
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
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceTag sets the deviceTag property value. The deviceTag property
func (m *ItemMicrosoftGraphWipeManagedAppRegistrationsByDeviceTagWipeManagedAppRegistrationsByDeviceTagPostRequestBody) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
