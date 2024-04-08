package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SelfServiceSignUpAuthenticationFlowConfiguration 
type SelfServiceSignUpAuthenticationFlowConfiguration struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
    isEnabled *bool
    // The OdataType property
    odataType *string
}
// NewSelfServiceSignUpAuthenticationFlowConfiguration instantiates a new selfServiceSignUpAuthenticationFlowConfiguration and sets the default values.
func NewSelfServiceSignUpAuthenticationFlowConfiguration()(*SelfServiceSignUpAuthenticationFlowConfiguration) {
    m := &SelfServiceSignUpAuthenticationFlowConfiguration{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSelfServiceSignUpAuthenticationFlowConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSelfServiceSignUpAuthenticationFlowConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSelfServiceSignUpAuthenticationFlowConfiguration(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
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
// GetIsEnabled gets the isEnabled property value. Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
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
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether self-service sign-up flow is enabled or disabled. The default value is false. This property is not a key. Required.
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SelfServiceSignUpAuthenticationFlowConfiguration) SetOdataType(value *string)() {
    m.odataType = value
}
