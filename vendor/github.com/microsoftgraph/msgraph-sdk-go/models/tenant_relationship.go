package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TenantRelationship 
type TenantRelationship struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The customer who has a delegated admin relationship with a Microsoft partner.
    delegatedAdminCustomers []DelegatedAdminCustomerable
    // The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
    delegatedAdminRelationships []DelegatedAdminRelationshipable
    // The OdataType property
    odataType *string
}
// NewTenantRelationship instantiates a new TenantRelationship and sets the default values.
func NewTenantRelationship()(*TenantRelationship) {
    m := &TenantRelationship{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTenantRelationshipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTenantRelationshipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTenantRelationship(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TenantRelationship) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDelegatedAdminCustomers gets the delegatedAdminCustomers property value. The customer who has a delegated admin relationship with a Microsoft partner.
func (m *TenantRelationship) GetDelegatedAdminCustomers()([]DelegatedAdminCustomerable) {
    return m.delegatedAdminCustomers
}
// GetDelegatedAdminRelationships gets the delegatedAdminRelationships property value. The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *TenantRelationship) GetDelegatedAdminRelationships()([]DelegatedAdminRelationshipable) {
    return m.delegatedAdminRelationships
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TenantRelationship) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["delegatedAdminCustomers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminCustomerable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminCustomerable)
            }
            m.SetDelegatedAdminCustomers(res)
        }
        return nil
    }
    res["delegatedAdminRelationships"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminRelationshipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminRelationshipable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminRelationshipable)
            }
            m.SetDelegatedAdminRelationships(res)
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
func (m *TenantRelationship) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TenantRelationship) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDelegatedAdminCustomers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDelegatedAdminCustomers()))
        for i, v := range m.GetDelegatedAdminCustomers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("delegatedAdminCustomers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDelegatedAdminRelationships() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDelegatedAdminRelationships()))
        for i, v := range m.GetDelegatedAdminRelationships() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("delegatedAdminRelationships", cast)
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
func (m *TenantRelationship) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDelegatedAdminCustomers sets the delegatedAdminCustomers property value. The customer who has a delegated admin relationship with a Microsoft partner.
func (m *TenantRelationship) SetDelegatedAdminCustomers(value []DelegatedAdminCustomerable)() {
    m.delegatedAdminCustomers = value
}
// SetDelegatedAdminRelationships sets the delegatedAdminRelationships property value. The details of the delegated administrative privileges that a Microsoft partner has in a customer tenant.
func (m *TenantRelationship) SetDelegatedAdminRelationships(value []DelegatedAdminRelationshipable)() {
    m.delegatedAdminRelationships = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TenantRelationship) SetOdataType(value *string)() {
    m.odataType = value
}
