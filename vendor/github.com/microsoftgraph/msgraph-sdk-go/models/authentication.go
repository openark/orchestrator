package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Authentication 
type Authentication struct {
    Entity
    // The email address registered to a user for authentication.
    emailMethods []EmailAuthenticationMethodable
    // Represents the FIDO2 security keys registered to a user for authentication.
    fido2Methods []Fido2AuthenticationMethodable
    // Represents all authentication methods registered to a user.
    methods []AuthenticationMethodable
    // The details of the Microsoft Authenticator app registered to a user for authentication.
    microsoftAuthenticatorMethods []MicrosoftAuthenticatorAuthenticationMethodable
    // Represents the status of a long-running operation.
    operations []LongRunningOperationable
    // Represents the password that's registered to a user for authentication. For security, the password itself will never be returned in the object, but action can be taken to reset a password.
    passwordMethods []PasswordAuthenticationMethodable
    // The phone numbers registered to a user for authentication.
    phoneMethods []PhoneAuthenticationMethodable
    // The software OATH TOTP applications registered to a user for authentication.
    softwareOathMethods []SoftwareOathAuthenticationMethodable
    // Represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
    temporaryAccessPassMethods []TemporaryAccessPassAuthenticationMethodable
    // Represents the Windows Hello for Business authentication method registered to a user for authentication.
    windowsHelloForBusinessMethods []WindowsHelloForBusinessAuthenticationMethodable
}
// NewAuthentication instantiates a new authentication and sets the default values.
func NewAuthentication()(*Authentication) {
    m := &Authentication{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAuthenticationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthentication(), nil
}
// GetEmailMethods gets the emailMethods property value. The email address registered to a user for authentication.
func (m *Authentication) GetEmailMethods()([]EmailAuthenticationMethodable) {
    return m.emailMethods
}
// GetFido2Methods gets the fido2Methods property value. Represents the FIDO2 security keys registered to a user for authentication.
func (m *Authentication) GetFido2Methods()([]Fido2AuthenticationMethodable) {
    return m.fido2Methods
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Authentication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["emailMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEmailAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EmailAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(EmailAuthenticationMethodable)
            }
            m.SetEmailMethods(res)
        }
        return nil
    }
    res["fido2Methods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFido2AuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Fido2AuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(Fido2AuthenticationMethodable)
            }
            m.SetFido2Methods(res)
        }
        return nil
    }
    res["methods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodable)
            }
            m.SetMethods(res)
        }
        return nil
    }
    res["microsoftAuthenticatorMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MicrosoftAuthenticatorAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(MicrosoftAuthenticatorAuthenticationMethodable)
            }
            m.SetMicrosoftAuthenticatorMethods(res)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLongRunningOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LongRunningOperationable, len(val))
            for i, v := range val {
                res[i] = v.(LongRunningOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["passwordMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(PasswordAuthenticationMethodable)
            }
            m.SetPasswordMethods(res)
        }
        return nil
    }
    res["phoneMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhoneAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(PhoneAuthenticationMethodable)
            }
            m.SetPhoneMethods(res)
        }
        return nil
    }
    res["softwareOathMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSoftwareOathAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SoftwareOathAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(SoftwareOathAuthenticationMethodable)
            }
            m.SetSoftwareOathMethods(res)
        }
        return nil
    }
    res["temporaryAccessPassMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTemporaryAccessPassAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TemporaryAccessPassAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(TemporaryAccessPassAuthenticationMethodable)
            }
            m.SetTemporaryAccessPassMethods(res)
        }
        return nil
    }
    res["windowsHelloForBusinessMethods"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsHelloForBusinessAuthenticationMethodFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsHelloForBusinessAuthenticationMethodable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsHelloForBusinessAuthenticationMethodable)
            }
            m.SetWindowsHelloForBusinessMethods(res)
        }
        return nil
    }
    return res
}
// GetMethods gets the methods property value. Represents all authentication methods registered to a user.
func (m *Authentication) GetMethods()([]AuthenticationMethodable) {
    return m.methods
}
// GetMicrosoftAuthenticatorMethods gets the microsoftAuthenticatorMethods property value. The details of the Microsoft Authenticator app registered to a user for authentication.
func (m *Authentication) GetMicrosoftAuthenticatorMethods()([]MicrosoftAuthenticatorAuthenticationMethodable) {
    return m.microsoftAuthenticatorMethods
}
// GetOperations gets the operations property value. Represents the status of a long-running operation.
func (m *Authentication) GetOperations()([]LongRunningOperationable) {
    return m.operations
}
// GetPasswordMethods gets the passwordMethods property value. Represents the password that's registered to a user for authentication. For security, the password itself will never be returned in the object, but action can be taken to reset a password.
func (m *Authentication) GetPasswordMethods()([]PasswordAuthenticationMethodable) {
    return m.passwordMethods
}
// GetPhoneMethods gets the phoneMethods property value. The phone numbers registered to a user for authentication.
func (m *Authentication) GetPhoneMethods()([]PhoneAuthenticationMethodable) {
    return m.phoneMethods
}
// GetSoftwareOathMethods gets the softwareOathMethods property value. The software OATH TOTP applications registered to a user for authentication.
func (m *Authentication) GetSoftwareOathMethods()([]SoftwareOathAuthenticationMethodable) {
    return m.softwareOathMethods
}
// GetTemporaryAccessPassMethods gets the temporaryAccessPassMethods property value. Represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
func (m *Authentication) GetTemporaryAccessPassMethods()([]TemporaryAccessPassAuthenticationMethodable) {
    return m.temporaryAccessPassMethods
}
// GetWindowsHelloForBusinessMethods gets the windowsHelloForBusinessMethods property value. Represents the Windows Hello for Business authentication method registered to a user for authentication.
func (m *Authentication) GetWindowsHelloForBusinessMethods()([]WindowsHelloForBusinessAuthenticationMethodable) {
    return m.windowsHelloForBusinessMethods
}
// Serialize serializes information the current object
func (m *Authentication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEmailMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEmailMethods()))
        for i, v := range m.GetEmailMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("emailMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFido2Methods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFido2Methods()))
        for i, v := range m.GetFido2Methods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("fido2Methods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMethods()))
        for i, v := range m.GetMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("methods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMicrosoftAuthenticatorMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMicrosoftAuthenticatorMethods()))
        for i, v := range m.GetMicrosoftAuthenticatorMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("microsoftAuthenticatorMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPasswordMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPasswordMethods()))
        for i, v := range m.GetPasswordMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("passwordMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPhoneMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhoneMethods()))
        for i, v := range m.GetPhoneMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("phoneMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSoftwareOathMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSoftwareOathMethods()))
        for i, v := range m.GetSoftwareOathMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("softwareOathMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTemporaryAccessPassMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemporaryAccessPassMethods()))
        for i, v := range m.GetTemporaryAccessPassMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("temporaryAccessPassMethods", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsHelloForBusinessMethods() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsHelloForBusinessMethods()))
        for i, v := range m.GetWindowsHelloForBusinessMethods() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsHelloForBusinessMethods", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmailMethods sets the emailMethods property value. The email address registered to a user for authentication.
func (m *Authentication) SetEmailMethods(value []EmailAuthenticationMethodable)() {
    m.emailMethods = value
}
// SetFido2Methods sets the fido2Methods property value. Represents the FIDO2 security keys registered to a user for authentication.
func (m *Authentication) SetFido2Methods(value []Fido2AuthenticationMethodable)() {
    m.fido2Methods = value
}
// SetMethods sets the methods property value. Represents all authentication methods registered to a user.
func (m *Authentication) SetMethods(value []AuthenticationMethodable)() {
    m.methods = value
}
// SetMicrosoftAuthenticatorMethods sets the microsoftAuthenticatorMethods property value. The details of the Microsoft Authenticator app registered to a user for authentication.
func (m *Authentication) SetMicrosoftAuthenticatorMethods(value []MicrosoftAuthenticatorAuthenticationMethodable)() {
    m.microsoftAuthenticatorMethods = value
}
// SetOperations sets the operations property value. Represents the status of a long-running operation.
func (m *Authentication) SetOperations(value []LongRunningOperationable)() {
    m.operations = value
}
// SetPasswordMethods sets the passwordMethods property value. Represents the password that's registered to a user for authentication. For security, the password itself will never be returned in the object, but action can be taken to reset a password.
func (m *Authentication) SetPasswordMethods(value []PasswordAuthenticationMethodable)() {
    m.passwordMethods = value
}
// SetPhoneMethods sets the phoneMethods property value. The phone numbers registered to a user for authentication.
func (m *Authentication) SetPhoneMethods(value []PhoneAuthenticationMethodable)() {
    m.phoneMethods = value
}
// SetSoftwareOathMethods sets the softwareOathMethods property value. The software OATH TOTP applications registered to a user for authentication.
func (m *Authentication) SetSoftwareOathMethods(value []SoftwareOathAuthenticationMethodable)() {
    m.softwareOathMethods = value
}
// SetTemporaryAccessPassMethods sets the temporaryAccessPassMethods property value. Represents a Temporary Access Pass registered to a user for authentication through time-limited passcodes.
func (m *Authentication) SetTemporaryAccessPassMethods(value []TemporaryAccessPassAuthenticationMethodable)() {
    m.temporaryAccessPassMethods = value
}
// SetWindowsHelloForBusinessMethods sets the windowsHelloForBusinessMethods property value. Represents the Windows Hello for Business authentication method registered to a user for authentication.
func (m *Authentication) SetWindowsHelloForBusinessMethods(value []WindowsHelloForBusinessAuthenticationMethodable)() {
    m.windowsHelloForBusinessMethods = value
}
