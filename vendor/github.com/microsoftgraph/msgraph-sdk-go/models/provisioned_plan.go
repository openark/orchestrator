package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisionedPlan 
type ProvisionedPlan struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // For example, 'Enabled'.
    capabilityStatus *string
    // The OdataType property
    odataType *string
    // For example, 'Success'.
    provisioningStatus *string
    // The name of the service; for example, 'AccessControlS2S'
    service *string
}
// NewProvisionedPlan instantiates a new provisionedPlan and sets the default values.
func NewProvisionedPlan()(*ProvisionedPlan) {
    m := &ProvisionedPlan{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProvisionedPlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisionedPlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisionedPlan(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisionedPlan) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCapabilityStatus gets the capabilityStatus property value. For example, 'Enabled'.
func (m *ProvisionedPlan) GetCapabilityStatus()(*string) {
    return m.capabilityStatus
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisionedPlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["capabilityStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilityStatus(val)
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
    res["provisioningStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStatus(val)
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ProvisionedPlan) GetOdataType()(*string) {
    return m.odataType
}
// GetProvisioningStatus gets the provisioningStatus property value. For example, 'Success'.
func (m *ProvisionedPlan) GetProvisioningStatus()(*string) {
    return m.provisioningStatus
}
// GetService gets the service property value. The name of the service; for example, 'AccessControlS2S'
func (m *ProvisionedPlan) GetService()(*string) {
    return m.service
}
// Serialize serializes information the current object
func (m *ProvisionedPlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("capabilityStatus", m.GetCapabilityStatus())
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
        err := writer.WriteStringValue("provisioningStatus", m.GetProvisioningStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("service", m.GetService())
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
func (m *ProvisionedPlan) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCapabilityStatus sets the capabilityStatus property value. For example, 'Enabled'.
func (m *ProvisionedPlan) SetCapabilityStatus(value *string)() {
    m.capabilityStatus = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ProvisionedPlan) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProvisioningStatus sets the provisioningStatus property value. For example, 'Success'.
func (m *ProvisionedPlan) SetProvisioningStatus(value *string)() {
    m.provisioningStatus = value
}
// SetService sets the service property value. The name of the service; for example, 'AccessControlS2S'
func (m *ProvisionedPlan) SetService(value *string)() {
    m.service = value
}
