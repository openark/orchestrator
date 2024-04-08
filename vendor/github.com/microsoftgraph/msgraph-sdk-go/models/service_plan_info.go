package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePlanInfo 
type ServicePlanInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The object the service plan can be assigned to. The possible values are:User - service plan can be assigned to individual users.Company - service plan can be assigned to the entire tenant.
    appliesTo *string
    // The OdataType property
    odataType *string
    // The provisioning status of the service plan. The possible values are:Success - Service is fully provisioned.Disabled - Service has been disabled.ErrorStatus - The service plan has not been provisioned and is in an error state.PendingInput - Service is not yet provisioned; awaiting service confirmation.PendingActivation - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)PendingProvisioning - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
    provisioningStatus *string
    // The unique identifier of the service plan.
    servicePlanId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The name of the service plan.
    servicePlanName *string
}
// NewServicePlanInfo instantiates a new servicePlanInfo and sets the default values.
func NewServicePlanInfo()(*ServicePlanInfo) {
    m := &ServicePlanInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateServicePlanInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePlanInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePlanInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServicePlanInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAppliesTo gets the appliesTo property value. The object the service plan can be assigned to. The possible values are:User - service plan can be assigned to individual users.Company - service plan can be assigned to the entire tenant.
func (m *ServicePlanInfo) GetAppliesTo()(*string) {
    return m.appliesTo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePlanInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliesTo(val)
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
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    res["servicePlanName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ServicePlanInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetProvisioningStatus gets the provisioningStatus property value. The provisioning status of the service plan. The possible values are:Success - Service is fully provisioned.Disabled - Service has been disabled.ErrorStatus - The service plan has not been provisioned and is in an error state.PendingInput - Service is not yet provisioned; awaiting service confirmation.PendingActivation - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)PendingProvisioning - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
func (m *ServicePlanInfo) GetProvisioningStatus()(*string) {
    return m.provisioningStatus
}
// GetServicePlanId gets the servicePlanId property value. The unique identifier of the service plan.
func (m *ServicePlanInfo) GetServicePlanId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.servicePlanId
}
// GetServicePlanName gets the servicePlanName property value. The name of the service plan.
func (m *ServicePlanInfo) GetServicePlanName()(*string) {
    return m.servicePlanName
}
// Serialize serializes information the current object
func (m *ServicePlanInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appliesTo", m.GetAppliesTo())
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
        err := writer.WriteUUIDValue("servicePlanId", m.GetServicePlanId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("servicePlanName", m.GetServicePlanName())
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
func (m *ServicePlanInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAppliesTo sets the appliesTo property value. The object the service plan can be assigned to. The possible values are:User - service plan can be assigned to individual users.Company - service plan can be assigned to the entire tenant.
func (m *ServicePlanInfo) SetAppliesTo(value *string)() {
    m.appliesTo = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ServicePlanInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetProvisioningStatus sets the provisioningStatus property value. The provisioning status of the service plan. The possible values are:Success - Service is fully provisioned.Disabled - Service has been disabled.ErrorStatus - The service plan has not been provisioned and is in an error state.PendingInput - Service is not yet provisioned; awaiting service confirmation.PendingActivation - Service is provisioned but requires explicit activation by administrator (for example, Intune_O365 service plan)PendingProvisioning - Microsoft has added a new service to the product SKU and it has not been activated in the tenant, yet.
func (m *ServicePlanInfo) SetProvisioningStatus(value *string)() {
    m.provisioningStatus = value
}
// SetServicePlanId sets the servicePlanId property value. The unique identifier of the service plan.
func (m *ServicePlanInfo) SetServicePlanId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.servicePlanId = value
}
// SetServicePlanName sets the servicePlanName property value. The name of the service plan.
func (m *ServicePlanInfo) SetServicePlanName(value *string)() {
    m.servicePlanName = value
}
