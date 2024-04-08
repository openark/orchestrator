package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessGrantControls 
type ConditionalAccessGrantControls struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.
    builtInControls []ConditionalAccessGrantControl
    // List of custom controls IDs required by the policy. For more information, see Custom controls.
    customAuthenticationFactors []string
    // The OdataType property
    odataType *string
    // Defines the relationship of the grant controls. Possible values: AND, OR.
    operator *string
    // List of terms of use IDs required by the policy.
    termsOfUse []string
}
// NewConditionalAccessGrantControls instantiates a new conditionalAccessGrantControls and sets the default values.
func NewConditionalAccessGrantControls()(*ConditionalAccessGrantControls) {
    m := &ConditionalAccessGrantControls{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateConditionalAccessGrantControlsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessGrantControlsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessGrantControls(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessGrantControls) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuiltInControls gets the builtInControls property value. List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.
func (m *ConditionalAccessGrantControls) GetBuiltInControls()([]ConditionalAccessGrantControl) {
    return m.builtInControls
}
// GetCustomAuthenticationFactors gets the customAuthenticationFactors property value. List of custom controls IDs required by the policy. For more information, see Custom controls.
func (m *ConditionalAccessGrantControls) GetCustomAuthenticationFactors()([]string) {
    return m.customAuthenticationFactors
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessGrantControls) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["builtInControls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseConditionalAccessGrantControl)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessGrantControl, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessGrantControl))
            }
            m.SetBuiltInControls(res)
        }
        return nil
    }
    res["customAuthenticationFactors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCustomAuthenticationFactors(res)
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
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val)
        }
        return nil
    }
    res["termsOfUse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTermsOfUse(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ConditionalAccessGrantControls) GetOdataType()(*string) {
    return m.odataType
}
// GetOperator gets the operator property value. Defines the relationship of the grant controls. Possible values: AND, OR.
func (m *ConditionalAccessGrantControls) GetOperator()(*string) {
    return m.operator
}
// GetTermsOfUse gets the termsOfUse property value. List of terms of use IDs required by the policy.
func (m *ConditionalAccessGrantControls) GetTermsOfUse()([]string) {
    return m.termsOfUse
}
// Serialize serializes information the current object
func (m *ConditionalAccessGrantControls) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBuiltInControls() != nil {
        err := writer.WriteCollectionOfStringValues("builtInControls", SerializeConditionalAccessGrantControl(m.GetBuiltInControls()))
        if err != nil {
            return err
        }
    }
    if m.GetCustomAuthenticationFactors() != nil {
        err := writer.WriteCollectionOfStringValues("customAuthenticationFactors", m.GetCustomAuthenticationFactors())
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
        err := writer.WriteStringValue("operator", m.GetOperator())
        if err != nil {
            return err
        }
    }
    if m.GetTermsOfUse() != nil {
        err := writer.WriteCollectionOfStringValues("termsOfUse", m.GetTermsOfUse())
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
func (m *ConditionalAccessGrantControls) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuiltInControls sets the builtInControls property value. List of values of built-in controls required by the policy. Possible values: block, mfa, compliantDevice, domainJoinedDevice, approvedApplication, compliantApplication, passwordChange, unknownFutureValue.
func (m *ConditionalAccessGrantControls) SetBuiltInControls(value []ConditionalAccessGrantControl)() {
    m.builtInControls = value
}
// SetCustomAuthenticationFactors sets the customAuthenticationFactors property value. List of custom controls IDs required by the policy. For more information, see Custom controls.
func (m *ConditionalAccessGrantControls) SetCustomAuthenticationFactors(value []string)() {
    m.customAuthenticationFactors = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ConditionalAccessGrantControls) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperator sets the operator property value. Defines the relationship of the grant controls. Possible values: AND, OR.
func (m *ConditionalAccessGrantControls) SetOperator(value *string)() {
    m.operator = value
}
// SetTermsOfUse sets the termsOfUse property value. List of terms of use IDs required by the policy.
func (m *ConditionalAccessGrantControls) SetTermsOfUse(value []string)() {
    m.termsOfUse = value
}
