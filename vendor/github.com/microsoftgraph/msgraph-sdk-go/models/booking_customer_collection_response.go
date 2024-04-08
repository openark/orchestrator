package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingCustomerCollectionResponse 
type BookingCustomerCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []BookingCustomerable
}
// NewBookingCustomerCollectionResponse instantiates a new BookingCustomerCollectionResponse and sets the default values.
func NewBookingCustomerCollectionResponse()(*BookingCustomerCollectionResponse) {
    m := &BookingCustomerCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateBookingCustomerCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingCustomerCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingCustomerCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingCustomerCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingCustomerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingCustomerable, len(val))
            for i, v := range val {
                res[i] = v.(BookingCustomerable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *BookingCustomerCollectionResponse) GetValue()([]BookingCustomerable) {
    return m.value
}
// Serialize serializes information the current object
func (m *BookingCustomerCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *BookingCustomerCollectionResponse) SetValue(value []BookingCustomerable)() {
    m.value = value
}
