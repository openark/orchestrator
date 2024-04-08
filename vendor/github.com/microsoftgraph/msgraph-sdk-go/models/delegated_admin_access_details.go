package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminAccessDetails 
type DelegatedAdminAccessDetails struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The directory roles that the Microsoft partner is assigned in the customer tenant.
    unifiedRoles []UnifiedRoleable
}
// NewDelegatedAdminAccessDetails instantiates a new delegatedAdminAccessDetails and sets the default values.
func NewDelegatedAdminAccessDetails()(*DelegatedAdminAccessDetails) {
    m := &DelegatedAdminAccessDetails{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDelegatedAdminAccessDetailsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminAccessDetailsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminAccessDetails(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DelegatedAdminAccessDetails) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminAccessDetails) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["unifiedRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnifiedRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnifiedRoleable, len(val))
            for i, v := range val {
                res[i] = v.(UnifiedRoleable)
            }
            m.SetUnifiedRoles(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DelegatedAdminAccessDetails) GetOdataType()(*string) {
    return m.odataType
}
// GetUnifiedRoles gets the unifiedRoles property value. The directory roles that the Microsoft partner is assigned in the customer tenant.
func (m *DelegatedAdminAccessDetails) GetUnifiedRoles()([]UnifiedRoleable) {
    return m.unifiedRoles
}
// Serialize serializes information the current object
func (m *DelegatedAdminAccessDetails) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetUnifiedRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUnifiedRoles()))
        for i, v := range m.GetUnifiedRoles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("unifiedRoles", cast)
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
func (m *DelegatedAdminAccessDetails) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DelegatedAdminAccessDetails) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUnifiedRoles sets the unifiedRoles property value. The directory roles that the Microsoft partner is assigned in the customer tenant.
func (m *DelegatedAdminAccessDetails) SetUnifiedRoles(value []UnifiedRoleable)() {
    m.unifiedRoles = value
}
