package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TemporaryAccessPassAuthenticationMethodConfiguration 
type TemporaryAccessPassAuthenticationMethodConfiguration struct {
    AuthenticationMethodConfiguration
    // Default length in characters of a Temporary Access Pass object. Must be between 8 and 48 characters.
    defaultLength *int32
    // Default lifetime in minutes for a Temporary Access Pass. Value can be any integer between the minimumLifetimeInMinutes and maximumLifetimeInMinutes.
    defaultLifetimeInMinutes *int32
    // A collection of groups that are enabled to use the authentication method.
    includeTargets []AuthenticationMethodTargetable
    // If true, all the passes in the tenant will be restricted to one-time use. If false, passes in the tenant can be created to be either one-time use or reusable.
    isUsableOnce *bool
    // Maximum lifetime in minutes for any Temporary Access Pass created in the tenant. Value can be between 10 and 43200 minutes (equivalent to 30 days).
    maximumLifetimeInMinutes *int32
    // Minimum lifetime in minutes for any Temporary Access Pass created in the tenant. Value can be between 10 and 43200 minutes (equivalent to 30 days).
    minimumLifetimeInMinutes *int32
}
// NewTemporaryAccessPassAuthenticationMethodConfiguration instantiates a new TemporaryAccessPassAuthenticationMethodConfiguration and sets the default values.
func NewTemporaryAccessPassAuthenticationMethodConfiguration()(*TemporaryAccessPassAuthenticationMethodConfiguration) {
    m := &TemporaryAccessPassAuthenticationMethodConfiguration{
        AuthenticationMethodConfiguration: *NewAuthenticationMethodConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.temporaryAccessPassAuthenticationMethodConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTemporaryAccessPassAuthenticationMethodConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTemporaryAccessPassAuthenticationMethodConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTemporaryAccessPassAuthenticationMethodConfiguration(), nil
}
// GetDefaultLength gets the defaultLength property value. Default length in characters of a Temporary Access Pass object. Must be between 8 and 48 characters.
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) GetDefaultLength()(*int32) {
    return m.defaultLength
}
// GetDefaultLifetimeInMinutes gets the defaultLifetimeInMinutes property value. Default lifetime in minutes for a Temporary Access Pass. Value can be any integer between the minimumLifetimeInMinutes and maximumLifetimeInMinutes.
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) GetDefaultLifetimeInMinutes()(*int32) {
    return m.defaultLifetimeInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthenticationMethodConfiguration.GetFieldDeserializers()
    res["defaultLength"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLength(val)
        }
        return nil
    }
    res["defaultLifetimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLifetimeInMinutes(val)
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
    res["isUsableOnce"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsUsableOnce(val)
        }
        return nil
    }
    res["maximumLifetimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumLifetimeInMinutes(val)
        }
        return nil
    }
    res["minimumLifetimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumLifetimeInMinutes(val)
        }
        return nil
    }
    return res
}
// GetIncludeTargets gets the includeTargets property value. A collection of groups that are enabled to use the authentication method.
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) GetIncludeTargets()([]AuthenticationMethodTargetable) {
    return m.includeTargets
}
// GetIsUsableOnce gets the isUsableOnce property value. If true, all the passes in the tenant will be restricted to one-time use. If false, passes in the tenant can be created to be either one-time use or reusable.
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) GetIsUsableOnce()(*bool) {
    return m.isUsableOnce
}
// GetMaximumLifetimeInMinutes gets the maximumLifetimeInMinutes property value. Maximum lifetime in minutes for any Temporary Access Pass created in the tenant. Value can be between 10 and 43200 minutes (equivalent to 30 days).
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) GetMaximumLifetimeInMinutes()(*int32) {
    return m.maximumLifetimeInMinutes
}
// GetMinimumLifetimeInMinutes gets the minimumLifetimeInMinutes property value. Minimum lifetime in minutes for any Temporary Access Pass created in the tenant. Value can be between 10 and 43200 minutes (equivalent to 30 days).
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) GetMinimumLifetimeInMinutes()(*int32) {
    return m.minimumLifetimeInMinutes
}
// Serialize serializes information the current object
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthenticationMethodConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("defaultLength", m.GetDefaultLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("defaultLifetimeInMinutes", m.GetDefaultLifetimeInMinutes())
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
    {
        err = writer.WriteBoolValue("isUsableOnce", m.GetIsUsableOnce())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumLifetimeInMinutes", m.GetMaximumLifetimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumLifetimeInMinutes", m.GetMinimumLifetimeInMinutes())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDefaultLength sets the defaultLength property value. Default length in characters of a Temporary Access Pass object. Must be between 8 and 48 characters.
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) SetDefaultLength(value *int32)() {
    m.defaultLength = value
}
// SetDefaultLifetimeInMinutes sets the defaultLifetimeInMinutes property value. Default lifetime in minutes for a Temporary Access Pass. Value can be any integer between the minimumLifetimeInMinutes and maximumLifetimeInMinutes.
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) SetDefaultLifetimeInMinutes(value *int32)() {
    m.defaultLifetimeInMinutes = value
}
// SetIncludeTargets sets the includeTargets property value. A collection of groups that are enabled to use the authentication method.
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) SetIncludeTargets(value []AuthenticationMethodTargetable)() {
    m.includeTargets = value
}
// SetIsUsableOnce sets the isUsableOnce property value. If true, all the passes in the tenant will be restricted to one-time use. If false, passes in the tenant can be created to be either one-time use or reusable.
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) SetIsUsableOnce(value *bool)() {
    m.isUsableOnce = value
}
// SetMaximumLifetimeInMinutes sets the maximumLifetimeInMinutes property value. Maximum lifetime in minutes for any Temporary Access Pass created in the tenant. Value can be between 10 and 43200 minutes (equivalent to 30 days).
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) SetMaximumLifetimeInMinutes(value *int32)() {
    m.maximumLifetimeInMinutes = value
}
// SetMinimumLifetimeInMinutes sets the minimumLifetimeInMinutes property value. Minimum lifetime in minutes for any Temporary Access Pass created in the tenant. Value can be between 10 and 43200 minutes (equivalent to 30 days).
func (m *TemporaryAccessPassAuthenticationMethodConfiguration) SetMinimumLifetimeInMinutes(value *int32)() {
    m.minimumLifetimeInMinutes = value
}
