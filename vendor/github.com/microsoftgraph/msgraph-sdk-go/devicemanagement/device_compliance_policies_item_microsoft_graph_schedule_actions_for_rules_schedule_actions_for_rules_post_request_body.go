package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody 
type DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The deviceComplianceScheduledActionForRules property
    deviceComplianceScheduledActionForRules []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable
}
// NewDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody instantiates a new DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody and sets the default values.
func NewDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody()(*DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) {
    m := &DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDeviceComplianceScheduledActionForRules gets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) GetDeviceComplianceScheduledActionForRules()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable) {
    return m.deviceComplianceScheduledActionForRules
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["deviceComplianceScheduledActionForRules"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceScheduledActionForRuleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)
            }
            m.SetDeviceComplianceScheduledActionForRules(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDeviceComplianceScheduledActionForRules() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceComplianceScheduledActionForRules()))
        for i, v := range m.GetDeviceComplianceScheduledActionForRules() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("deviceComplianceScheduledActionForRules", cast)
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
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDeviceComplianceScheduledActionForRules sets the deviceComplianceScheduledActionForRules property value. The deviceComplianceScheduledActionForRules property
func (m *DeviceCompliancePoliciesItemMicrosoftGraphScheduleActionsForRulesScheduleActionsForRulesPostRequestBody) SetDeviceComplianceScheduledActionForRules(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceScheduledActionForRuleable)() {
    m.deviceComplianceScheduledActionForRules = value
}
