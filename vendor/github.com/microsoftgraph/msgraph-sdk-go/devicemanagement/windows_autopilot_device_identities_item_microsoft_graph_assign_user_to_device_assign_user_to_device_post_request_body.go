package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody 
type WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The addressableUserName property
    addressableUserName *string
    // The userPrincipalName property
    userPrincipalName *string
}
// NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody instantiates a new WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody()(*WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) {
    m := &WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddressableUserName gets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) GetAddressableUserName()(*string) {
    return m.addressableUserName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addressableUserName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddressableUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddressableUserName sets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphAssignUserToDeviceAssignUserToDevicePostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
