package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubscribedSku 
type SubscribedSku struct {
    Entity
    // For example, 'User' or 'Company'.
    appliesTo *string
    // Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription.
    capabilityStatus *string
    // The number of licenses that have been assigned.
    consumedUnits *int32
    // Information about the number and status of prepaid licenses.
    prepaidUnits LicenseUnitsDetailable
    // Information about the service plans that are available with the SKU. Not nullable
    servicePlans []ServicePlanInfoable
    // The unique identifier (GUID) for the service SKU.
    skuId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The SKU part number; for example: 'AAD_PREMIUM' or 'RMSBASIC'. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
    skuPartNumber *string
}
// NewSubscribedSku instantiates a new SubscribedSku and sets the default values.
func NewSubscribedSku()(*SubscribedSku) {
    m := &SubscribedSku{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSubscribedSkuFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubscribedSkuFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubscribedSku(), nil
}
// GetAppliesTo gets the appliesTo property value. For example, 'User' or 'Company'.
func (m *SubscribedSku) GetAppliesTo()(*string) {
    return m.appliesTo
}
// GetCapabilityStatus gets the capabilityStatus property value. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription.
func (m *SubscribedSku) GetCapabilityStatus()(*string) {
    return m.capabilityStatus
}
// GetConsumedUnits gets the consumedUnits property value. The number of licenses that have been assigned.
func (m *SubscribedSku) GetConsumedUnits()(*int32) {
    return m.consumedUnits
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubscribedSku) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["consumedUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConsumedUnits(val)
        }
        return nil
    }
    res["prepaidUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLicenseUnitsDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrepaidUnits(val.(LicenseUnitsDetailable))
        }
        return nil
    }
    res["servicePlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePlanInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePlanInfoable, len(val))
            for i, v := range val {
                res[i] = v.(ServicePlanInfoable)
            }
            m.SetServicePlans(res)
        }
        return nil
    }
    res["skuId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    res["skuPartNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuPartNumber(val)
        }
        return nil
    }
    return res
}
// GetPrepaidUnits gets the prepaidUnits property value. Information about the number and status of prepaid licenses.
func (m *SubscribedSku) GetPrepaidUnits()(LicenseUnitsDetailable) {
    return m.prepaidUnits
}
// GetServicePlans gets the servicePlans property value. Information about the service plans that are available with the SKU. Not nullable
func (m *SubscribedSku) GetServicePlans()([]ServicePlanInfoable) {
    return m.servicePlans
}
// GetSkuId gets the skuId property value. The unique identifier (GUID) for the service SKU.
func (m *SubscribedSku) GetSkuId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.skuId
}
// GetSkuPartNumber gets the skuPartNumber property value. The SKU part number; for example: 'AAD_PREMIUM' or 'RMSBASIC'. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
func (m *SubscribedSku) GetSkuPartNumber()(*string) {
    return m.skuPartNumber
}
// Serialize serializes information the current object
func (m *SubscribedSku) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appliesTo", m.GetAppliesTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("capabilityStatus", m.GetCapabilityStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("consumedUnits", m.GetConsumedUnits())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("prepaidUnits", m.GetPrepaidUnits())
        if err != nil {
            return err
        }
    }
    if m.GetServicePlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetServicePlans()))
        for i, v := range m.GetServicePlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("servicePlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("skuId", m.GetSkuId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("skuPartNumber", m.GetSkuPartNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. For example, 'User' or 'Company'.
func (m *SubscribedSku) SetAppliesTo(value *string)() {
    m.appliesTo = value
}
// SetCapabilityStatus sets the capabilityStatus property value. Possible values are: Enabled, Warning, Suspended, Deleted, LockedOut. The capabilityStatus is Enabled if the prepaidUnits property has at least 1 unit that is enabled, and LockedOut if the customer cancelled their subscription.
func (m *SubscribedSku) SetCapabilityStatus(value *string)() {
    m.capabilityStatus = value
}
// SetConsumedUnits sets the consumedUnits property value. The number of licenses that have been assigned.
func (m *SubscribedSku) SetConsumedUnits(value *int32)() {
    m.consumedUnits = value
}
// SetPrepaidUnits sets the prepaidUnits property value. Information about the number and status of prepaid licenses.
func (m *SubscribedSku) SetPrepaidUnits(value LicenseUnitsDetailable)() {
    m.prepaidUnits = value
}
// SetServicePlans sets the servicePlans property value. Information about the service plans that are available with the SKU. Not nullable
func (m *SubscribedSku) SetServicePlans(value []ServicePlanInfoable)() {
    m.servicePlans = value
}
// SetSkuId sets the skuId property value. The unique identifier (GUID) for the service SKU.
func (m *SubscribedSku) SetSkuId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.skuId = value
}
// SetSkuPartNumber sets the skuPartNumber property value. The SKU part number; for example: 'AAD_PREMIUM' or 'RMSBASIC'. To get a list of commercial subscriptions that an organization has acquired, see List subscribedSkus.
func (m *SubscribedSku) SetSkuPartNumber(value *string)() {
    m.skuPartNumber = value
}
