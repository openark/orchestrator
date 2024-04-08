package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeleteUserFromSharedAppleDeviceActionResult 
type DeleteUserFromSharedAppleDeviceActionResult struct {
    DeviceActionResult
    // User principal name of the user to be deleted
    userPrincipalName *string
}
// NewDeleteUserFromSharedAppleDeviceActionResult instantiates a new DeleteUserFromSharedAppleDeviceActionResult and sets the default values.
func NewDeleteUserFromSharedAppleDeviceActionResult()(*DeleteUserFromSharedAppleDeviceActionResult) {
    m := &DeleteUserFromSharedAppleDeviceActionResult{
        DeviceActionResult: *NewDeviceActionResult(),
    }
    return m
}
// CreateDeleteUserFromSharedAppleDeviceActionResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeleteUserFromSharedAppleDeviceActionResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeleteUserFromSharedAppleDeviceActionResult(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeleteUserFromSharedAppleDeviceActionResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceActionResult.GetFieldDeserializers()
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
// GetUserPrincipalName gets the userPrincipalName property value. User principal name of the user to be deleted
func (m *DeleteUserFromSharedAppleDeviceActionResult) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DeleteUserFromSharedAppleDeviceActionResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceActionResult.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name of the user to be deleted
func (m *DeleteUserFromSharedAppleDeviceActionResult) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
