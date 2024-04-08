package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityResourceCollectionResponse 
type SecurityResourceCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []SecurityResourceable
}
// NewSecurityResourceCollectionResponse instantiates a new SecurityResourceCollectionResponse and sets the default values.
func NewSecurityResourceCollectionResponse()(*SecurityResourceCollectionResponse) {
    m := &SecurityResourceCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSecurityResourceCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityResourceCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityResourceCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityResourceCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecurityResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecurityResourceable, len(val))
            for i, v := range val {
                res[i] = v.(SecurityResourceable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *SecurityResourceCollectionResponse) GetValue()([]SecurityResourceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *SecurityResourceCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *SecurityResourceCollectionResponse) SetValue(value []SecurityResourceable)() {
    m.value = value
}
