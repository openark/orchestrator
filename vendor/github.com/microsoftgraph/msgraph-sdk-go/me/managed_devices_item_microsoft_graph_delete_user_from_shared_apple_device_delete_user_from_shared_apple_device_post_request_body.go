package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody 
type ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The userPrincipalName property
    userPrincipalName *string
}
// NewManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody instantiates a new ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody and sets the default values.
func NewManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody()(*ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody) {
    m := &ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
func (m *ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName property
func (m *ManagedDevicesItemMicrosoftGraphDeleteUserFromSharedAppleDeviceDeleteUserFromSharedAppleDevicePostRequestBody) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
