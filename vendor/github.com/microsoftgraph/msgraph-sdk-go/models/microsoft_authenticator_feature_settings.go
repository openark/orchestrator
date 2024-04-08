package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MicrosoftAuthenticatorFeatureSettings 
type MicrosoftAuthenticatorFeatureSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Determines whether the user's Authenticator app will show them the client app they are signing into.
    displayAppInformationRequiredState AuthenticationMethodFeatureConfigurationable
    // Determines whether the user's Authenticator app will show them the geographic location of where the authentication request originated from.
    displayLocationInformationRequiredState AuthenticationMethodFeatureConfigurationable
    // The OdataType property
    odataType *string
}
// NewMicrosoftAuthenticatorFeatureSettings instantiates a new microsoftAuthenticatorFeatureSettings and sets the default values.
func NewMicrosoftAuthenticatorFeatureSettings()(*MicrosoftAuthenticatorFeatureSettings) {
    m := &MicrosoftAuthenticatorFeatureSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftAuthenticatorFeatureSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftAuthenticatorFeatureSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftAuthenticatorFeatureSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftAuthenticatorFeatureSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayAppInformationRequiredState gets the displayAppInformationRequiredState property value. Determines whether the user's Authenticator app will show them the client app they are signing into.
func (m *MicrosoftAuthenticatorFeatureSettings) GetDisplayAppInformationRequiredState()(AuthenticationMethodFeatureConfigurationable) {
    return m.displayAppInformationRequiredState
}
// GetDisplayLocationInformationRequiredState gets the displayLocationInformationRequiredState property value. Determines whether the user's Authenticator app will show them the geographic location of where the authentication request originated from.
func (m *MicrosoftAuthenticatorFeatureSettings) GetDisplayLocationInformationRequiredState()(AuthenticationMethodFeatureConfigurationable) {
    return m.displayLocationInformationRequiredState
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftAuthenticatorFeatureSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayAppInformationRequiredState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodFeatureConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAppInformationRequiredState(val.(AuthenticationMethodFeatureConfigurationable))
        }
        return nil
    }
    res["displayLocationInformationRequiredState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodFeatureConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayLocationInformationRequiredState(val.(AuthenticationMethodFeatureConfigurationable))
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
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MicrosoftAuthenticatorFeatureSettings) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MicrosoftAuthenticatorFeatureSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("displayAppInformationRequiredState", m.GetDisplayAppInformationRequiredState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("displayLocationInformationRequiredState", m.GetDisplayLocationInformationRequiredState())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftAuthenticatorFeatureSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayAppInformationRequiredState sets the displayAppInformationRequiredState property value. Determines whether the user's Authenticator app will show them the client app they are signing into.
func (m *MicrosoftAuthenticatorFeatureSettings) SetDisplayAppInformationRequiredState(value AuthenticationMethodFeatureConfigurationable)() {
    m.displayAppInformationRequiredState = value
}
// SetDisplayLocationInformationRequiredState sets the displayLocationInformationRequiredState property value. Determines whether the user's Authenticator app will show them the geographic location of where the authentication request originated from.
func (m *MicrosoftAuthenticatorFeatureSettings) SetDisplayLocationInformationRequiredState(value AuthenticationMethodFeatureConfigurationable)() {
    m.displayLocationInformationRequiredState = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MicrosoftAuthenticatorFeatureSettings) SetOdataType(value *string)() {
    m.odataType = value
}
