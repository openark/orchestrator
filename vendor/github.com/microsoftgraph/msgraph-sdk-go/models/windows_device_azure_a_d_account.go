package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsDeviceAzureADAccount 
type WindowsDeviceAzureADAccount struct {
    WindowsDeviceAccount
    // Not yet documented
    userPrincipalName *string
}
// NewWindowsDeviceAzureADAccount instantiates a new WindowsDeviceAzureADAccount and sets the default values.
func NewWindowsDeviceAzureADAccount()(*WindowsDeviceAzureADAccount) {
    m := &WindowsDeviceAzureADAccount{
        WindowsDeviceAccount: *NewWindowsDeviceAccount(),
    }
    odataTypeValue := "#microsoft.graph.windowsDeviceAzureADAccount"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindowsDeviceAzureADAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsDeviceAzureADAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsDeviceAzureADAccount(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsDeviceAzureADAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.WindowsDeviceAccount.GetFieldDeserializers()
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
// GetUserPrincipalName gets the userPrincipalName property value. Not yet documented
func (m *WindowsDeviceAzureADAccount) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *WindowsDeviceAzureADAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.WindowsDeviceAccount.Serialize(writer)
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
// SetUserPrincipalName sets the userPrincipalName property value. Not yet documented
func (m *WindowsDeviceAzureADAccount) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
