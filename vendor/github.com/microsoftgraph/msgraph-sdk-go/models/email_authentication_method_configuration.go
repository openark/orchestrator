package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailAuthenticationMethodConfiguration 
type EmailAuthenticationMethodConfiguration struct {
    AuthenticationMethodConfiguration
    // Determines whether email OTP is usable by external users for authentication. Possible values are: default, enabled, disabled, unknownFutureValue. Tenants in the default state who did not use public preview will automatically have email OTP enabled beginning in October 2021.
    allowExternalIdToUseEmailOtp *ExternalEmailOtpState
    // A collection of groups that are enabled to use the authentication method.
    includeTargets []AuthenticationMethodTargetable
}
// NewEmailAuthenticationMethodConfiguration instantiates a new EmailAuthenticationMethodConfiguration and sets the default values.
func NewEmailAuthenticationMethodConfiguration()(*EmailAuthenticationMethodConfiguration) {
    m := &EmailAuthenticationMethodConfiguration{
        AuthenticationMethodConfiguration: *NewAuthenticationMethodConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.emailAuthenticationMethodConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEmailAuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailAuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailAuthenticationMethodConfiguration(), nil
}
// GetAllowExternalIdToUseEmailOtp gets the allowExternalIdToUseEmailOtp property value. Determines whether email OTP is usable by external users for authentication. Possible values are: default, enabled, disabled, unknownFutureValue. Tenants in the default state who did not use public preview will automatically have email OTP enabled beginning in October 2021.
func (m *EmailAuthenticationMethodConfiguration) GetAllowExternalIdToUseEmailOtp()(*ExternalEmailOtpState) {
    return m.allowExternalIdToUseEmailOtp
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailAuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodConfiguration.GetFieldDeserializers()
    res["allowExternalIdToUseEmailOtp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalEmailOtpState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowExternalIdToUseEmailOtp(val.(*ExternalEmailOtpState))
        }
        return nil
    }
    res["includeTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthenticationMethodTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthenticationMethodTargetable, len(val))
            for i, v := range val {
                res[i] = v.(AuthenticationMethodTargetable)
            }
            m.SetIncludeTargets(res)
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. A collection of groups that are enabled to use the authentication method.
func (m *EmailAuthenticationMethodConfiguration) GetIncludeTargets()([]AuthenticationMethodTargetable) {
    return m.includeTargets
}
// Serialize serializes information the current object
func (m *EmailAuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowExternalIdToUseEmailOtp() != nil {
        cast := (*m.GetAllowExternalIdToUseEmailOtp()).String()
        err = writer.WriteStringValue("allowExternalIdToUseEmailOtp", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIncludeTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIncludeTargets()))
        for i, v := range m.GetIncludeTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("includeTargets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowExternalIdToUseEmailOtp sets the allowExternalIdToUseEmailOtp property value. Determines whether email OTP is usable by external users for authentication. Possible values are: default, enabled, disabled, unknownFutureValue. Tenants in the default state who did not use public preview will automatically have email OTP enabled beginning in October 2021.
func (m *EmailAuthenticationMethodConfiguration) SetAllowExternalIdToUseEmailOtp(value *ExternalEmailOtpState)() {
    m.allowExternalIdToUseEmailOtp = value
}
// SetIncludeTargets sets the includeTargets property value. A collection of groups that are enabled to use the authentication method.
func (m *EmailAuthenticationMethodConfiguration) SetIncludeTargets(value []AuthenticationMethodTargetable)() {
    m.includeTargets = value
}
