package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentitySourceCollectionResponse 
type IdentitySourceCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []IdentitySourceable
}
// NewIdentitySourceCollectionResponse instantiates a new IdentitySourceCollectionResponse and sets the default values.
func NewIdentitySourceCollectionResponse()(*IdentitySourceCollectionResponse) {
    m := &IdentitySourceCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateIdentitySourceCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIdentitySourceCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIdentitySourceCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IdentitySourceCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentitySourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentitySourceable, len(val))
            for i, v := range val {
                res[i] = v.(IdentitySourceable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *IdentitySourceCollectionResponse) GetValue()([]IdentitySourceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *IdentitySourceCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *IdentitySourceCollectionResponse) SetValue(value []IdentitySourceable)() {
    m.value = value
}
