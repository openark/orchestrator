package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PasswordProfile 
type PasswordProfile struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // true if the user must change her password on the next login; otherwise false.
    forceChangePasswordNextSignIn *bool
    // If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first perform a multi-factor authentication before password change. After a password change, this property will be automatically reset to false. If not set, default is false.
    forceChangePasswordNextSignInWithMfa *bool
    // The OdataType property
    odataType *string
    // The password for the user. This property is required when a user is created. It can be updated, but the user will be required to change the password on the next login. The password must satisfy minimum requirements as specified by the user’s passwordPolicies property. By default, a strong password is required.
    password *string
}
// NewPasswordProfile instantiates a new passwordProfile and sets the default values.
func NewPasswordProfile()(*PasswordProfile) {
    m := &PasswordProfile{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePasswordProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePasswordProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPasswordProfile(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PasswordProfile) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PasswordProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["forceChangePasswordNextSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceChangePasswordNextSignIn(val)
        }
        return nil
    }
    res["forceChangePasswordNextSignInWithMfa"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetForceChangePasswordNextSignInWithMfa(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    return res
}
// GetForceChangePasswordNextSignIn gets the forceChangePasswordNextSignIn property value. true if the user must change her password on the next login; otherwise false.
func (m *PasswordProfile) GetForceChangePasswordNextSignIn()(*bool) {
    return m.forceChangePasswordNextSignIn
}
// GetForceChangePasswordNextSignInWithMfa gets the forceChangePasswordNextSignInWithMfa property value. If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first perform a multi-factor authentication before password change. After a password change, this property will be automatically reset to false. If not set, default is false.
func (m *PasswordProfile) GetForceChangePasswordNextSignInWithMfa()(*bool) {
    return m.forceChangePasswordNextSignInWithMfa
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PasswordProfile) GetOdataType()(*string) {
    return m.odataType
}
// GetPassword gets the password property value. The password for the user. This property is required when a user is created. It can be updated, but the user will be required to change the password on the next login. The password must satisfy minimum requirements as specified by the user’s passwordPolicies property. By default, a strong password is required.
func (m *PasswordProfile) GetPassword()(*string) {
    return m.password
}
// Serialize serializes information the current object
func (m *PasswordProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("forceChangePasswordNextSignIn", m.GetForceChangePasswordNextSignIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("forceChangePasswordNextSignInWithMfa", m.GetForceChangePasswordNextSignInWithMfa())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
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
func (m *PasswordProfile) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetForceChangePasswordNextSignIn sets the forceChangePasswordNextSignIn property value. true if the user must change her password on the next login; otherwise false.
func (m *PasswordProfile) SetForceChangePasswordNextSignIn(value *bool)() {
    m.forceChangePasswordNextSignIn = value
}
// SetForceChangePasswordNextSignInWithMfa sets the forceChangePasswordNextSignInWithMfa property value. If true, at next sign-in, the user must perform a multi-factor authentication (MFA) before being forced to change their password. The behavior is identical to forceChangePasswordNextSignIn except that the user is required to first perform a multi-factor authentication before password change. After a password change, this property will be automatically reset to false. If not set, default is false.
func (m *PasswordProfile) SetForceChangePasswordNextSignInWithMfa(value *bool)() {
    m.forceChangePasswordNextSignInWithMfa = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PasswordProfile) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPassword sets the password property value. The password for the user. This property is required when a user is created. It can be updated, but the user will be required to change the password on the next login. The password must satisfy minimum requirements as specified by the user’s passwordPolicies property. By default, a strong password is required.
func (m *PasswordProfile) SetPassword(value *string)() {
    m.password = value
}
