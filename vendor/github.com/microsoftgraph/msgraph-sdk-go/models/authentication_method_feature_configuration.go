package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationMethodFeatureConfiguration 
type AuthenticationMethodFeatureConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A single entity that is excluded from this feature.
    excludeTarget FeatureTargetable
    // A single entity that is included in this feature.
    includeTarget FeatureTargetable
    // The OdataType property
    odataType *string
    // Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure AD for the setting. The default value is disabled.
    state *AdvancedConfigState
}
// NewAuthenticationMethodFeatureConfiguration instantiates a new authenticationMethodFeatureConfiguration and sets the default values.
func NewAuthenticationMethodFeatureConfiguration()(*AuthenticationMethodFeatureConfiguration) {
    m := &AuthenticationMethodFeatureConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAuthenticationMethodFeatureConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationMethodFeatureConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationMethodFeatureConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AuthenticationMethodFeatureConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExcludeTarget gets the excludeTarget property value. A single entity that is excluded from this feature.
func (m *AuthenticationMethodFeatureConfiguration) GetExcludeTarget()(FeatureTargetable) {
    return m.excludeTarget
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodFeatureConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["excludeTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFeatureTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExcludeTarget(val.(FeatureTargetable))
        }
        return nil
    }
    res["includeTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFeatureTargetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeTarget(val.(FeatureTargetable))
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAdvancedConfigState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*AdvancedConfigState))
        }
        return nil
    }
    return res
}
// GetIncludeTarget gets the includeTarget property value. A single entity that is included in this feature.
func (m *AuthenticationMethodFeatureConfiguration) GetIncludeTarget()(FeatureTargetable) {
    return m.includeTarget
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AuthenticationMethodFeatureConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// GetState gets the state property value. Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure AD for the setting. The default value is disabled.
func (m *AuthenticationMethodFeatureConfiguration) GetState()(*AdvancedConfigState) {
    return m.state
}
// Serialize serializes information the current object
func (m *AuthenticationMethodFeatureConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("excludeTarget", m.GetExcludeTarget())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("includeTarget", m.GetIncludeTarget())
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
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err := writer.WriteStringValue("state", &cast)
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
func (m *AuthenticationMethodFeatureConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExcludeTarget sets the excludeTarget property value. A single entity that is excluded from this feature.
func (m *AuthenticationMethodFeatureConfiguration) SetExcludeTarget(value FeatureTargetable)() {
    m.excludeTarget = value
}
// SetIncludeTarget sets the includeTarget property value. A single entity that is included in this feature.
func (m *AuthenticationMethodFeatureConfiguration) SetIncludeTarget(value FeatureTargetable)() {
    m.includeTarget = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AuthenticationMethodFeatureConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
// SetState sets the state property value. Enable or disable the feature. Possible values are: default, enabled, disabled, unknownFutureValue. The default value is used when the configuration hasn't been explicitly set and uses the default behavior of Azure AD for the setting. The default value is disabled.
func (m *AuthenticationMethodFeatureConfiguration) SetState(value *AdvancedConfigState)() {
    m.state = value
}
