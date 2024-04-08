package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAuthenticatorAuthenticationMethod 
type MicrosoftAuthenticatorAuthenticationMethod struct {
    AuthenticationMethod
    // The date and time that this app was registered. This property is null if the device is not registered for passwordless Phone Sign-In.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
    device Deviceable
    // Tags containing app metadata.
    deviceTag *string
    // The name of the device on which this app is registered.
    displayName *string
    // Numerical version of this instance of the Authenticator app.
    phoneAppVersion *string
}
// NewMicrosoftAuthenticatorAuthenticationMethod instantiates a new MicrosoftAuthenticatorAuthenticationMethod and sets the default values.
func NewMicrosoftAuthenticatorAuthenticationMethod()(*MicrosoftAuthenticatorAuthenticationMethod) {
    m := &MicrosoftAuthenticatorAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    odataTypeValue := "#microsoft.graph.microsoftAuthenticatorAuthenticationMethod"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAuthenticatorAuthenticationMethod(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that this app was registered. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDevice gets the device property value. The registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDevice()(Deviceable) {
    return m.device
}
// GetDeviceTag gets the deviceTag property value. Tags containing app metadata.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDeviceTag()(*string) {
    return m.deviceTag
}
// GetDisplayName gets the displayName property value. The name of the device on which this app is registered.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["device"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(Deviceable))
        }
        return nil
    }
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
    res["phoneAppVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneAppVersion(val)
        }
        return nil
    }
    return res
}
// GetPhoneAppVersion gets the phoneAppVersion property value. Numerical version of this instance of the Authenticator app.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetPhoneAppVersion()(*string) {
    return m.phoneAppVersion
}
// Serialize serializes information the current object
func (m *MicrosoftAuthenticatorAuthenticationMethod) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceTag", m.GetDeviceTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneAppVersion", m.GetPhoneAppVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that this app was registered. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDevice sets the device property value. The registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDevice(value Deviceable)() {
    m.device = value
}
// SetDeviceTag sets the deviceTag property value. Tags containing app metadata.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDeviceTag(value *string)() {
    m.deviceTag = value
}
// SetDisplayName sets the displayName property value. The name of the device on which this app is registered.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPhoneAppVersion sets the phoneAppVersion property value. Numerical version of this instance of the Authenticator app.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetPhoneAppVersion(value *string)() {
    m.phoneAppVersion = value
}
