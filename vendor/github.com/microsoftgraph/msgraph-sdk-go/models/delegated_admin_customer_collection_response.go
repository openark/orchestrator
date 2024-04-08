package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminCustomerCollectionResponse 
type DelegatedAdminCustomerCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DelegatedAdminCustomerable
}
// NewDelegatedAdminCustomerCollectionResponse instantiates a new DelegatedAdminCustomerCollectionResponse and sets the default values.
func NewDelegatedAdminCustomerCollectionResponse()(*DelegatedAdminCustomerCollectionResponse) {
    m := &DelegatedAdminCustomerCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDelegatedAdminCustomerCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminCustomerCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminCustomerCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminCustomerCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminCustomerable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminCustomerable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DelegatedAdminCustomerCollectionResponse) GetValue()([]DelegatedAdminCustomerable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DelegatedAdminCustomerCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DelegatedAdminCustomerCollectionResponse) SetValue(value []DelegatedAdminCustomerable)() {
    m.value = value
}
