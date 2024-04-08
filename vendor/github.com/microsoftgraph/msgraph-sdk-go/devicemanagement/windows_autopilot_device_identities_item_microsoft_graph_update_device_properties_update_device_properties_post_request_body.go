package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody 
type WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The addressableUserName property
    addressableUserName *string
    // The displayName property
    displayName *string
    // The groupTag property
    groupTag *string
    // The userPrincipalName property
    userPrincipalName *string
}
// NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody instantiates a new WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody and sets the default values.
func NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody()(*WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) {
    m := &WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddressableUserName gets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) GetAddressableUserName()(*string) {
    return m.addressableUserName
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["groupTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupTag(val)
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
// GetGroupTag gets the groupTag property value. The groupTag property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) GetGroupTag()(*string) {
    return m.groupTag
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("addressableUserName", m.GetAddressableUserName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("groupTag", m.GetGroupTag())
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
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddressableUserName sets the addressableUserName property value. The addressableUserName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) SetAddressableUserName(value *string)() {
    m.addressableUserName = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGroupTag sets the groupTag property value. The groupTag property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) SetGroupTag(value *string)() {
    m.groupTag = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *WindowsAutopilotDeviceIdentitiesItemMicrosoftGraphUpdateDevicePropertiesUpdateDevicePropertiesPostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
