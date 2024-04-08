package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocationConstraintItemCollectionResponse 
type LocationConstraintItemCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []LocationConstraintItemable
}
// NewLocationConstraintItemCollectionResponse instantiates a new LocationConstraintItemCollectionResponse and sets the default values.
func NewLocationConstraintItemCollectionResponse()(*LocationConstraintItemCollectionResponse) {
    m := &LocationConstraintItemCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateLocationConstraintItemCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocationConstraintItemCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocationConstraintItemCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocationConstraintItemCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocationConstraintItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocationConstraintItemable, len(val))
            for i, v := range val {
                res[i] = v.(LocationConstraintItemable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *LocationConstraintItemCollectionResponse) GetValue()([]LocationConstraintItemable) {
    return m.value
}
// Serialize serializes information the current object
func (m *LocationConstraintItemCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *LocationConstraintItemCollectionResponse) SetValue(value []LocationConstraintItemable)() {
    m.value = value
}
