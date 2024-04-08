package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegistrationEnforcement 
type RegistrationEnforcement struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Run campaigns to remind users to set up targeted authentication methods.
    authenticationMethodsRegistrationCampaign AuthenticationMethodsRegistrationCampaignable
    // The OdataType property
    odataType *string
}
// NewRegistrationEnforcement instantiates a new registrationEnforcement and sets the default values.
func NewRegistrationEnforcement()(*RegistrationEnforcement) {
    m := &RegistrationEnforcement{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRegistrationEnforcementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegistrationEnforcementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegistrationEnforcement(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegistrationEnforcement) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAuthenticationMethodsRegistrationCampaign gets the authenticationMethodsRegistrationCampaign property value. Run campaigns to remind users to set up targeted authentication methods.
func (m *RegistrationEnforcement) GetAuthenticationMethodsRegistrationCampaign()(AuthenticationMethodsRegistrationCampaignable) {
    return m.authenticationMethodsRegistrationCampaign
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegistrationEnforcement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["authenticationMethodsRegistrationCampaign"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthenticationMethodsRegistrationCampaignFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationMethodsRegistrationCampaign(val.(AuthenticationMethodsRegistrationCampaignable))
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
func (m *RegistrationEnforcement) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *RegistrationEnforcement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("authenticationMethodsRegistrationCampaign", m.GetAuthenticationMethodsRegistrationCampaign())
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
func (m *RegistrationEnforcement) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAuthenticationMethodsRegistrationCampaign sets the authenticationMethodsRegistrationCampaign property value. Run campaigns to remind users to set up targeted authentication methods.
func (m *RegistrationEnforcement) SetAuthenticationMethodsRegistrationCampaign(value AuthenticationMethodsRegistrationCampaignable)() {
    m.authenticationMethodsRegistrationCampaign = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RegistrationEnforcement) SetOdataType(value *string)() {
    m.odataType = value
}
