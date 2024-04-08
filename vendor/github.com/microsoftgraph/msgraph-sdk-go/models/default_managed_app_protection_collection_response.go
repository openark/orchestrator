package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultManagedAppProtectionCollectionResponse 
type DefaultManagedAppProtectionCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DefaultManagedAppProtectionable
}
// NewDefaultManagedAppProtectionCollectionResponse instantiates a new DefaultManagedAppProtectionCollectionResponse and sets the default values.
func NewDefaultManagedAppProtectionCollectionResponse()(*DefaultManagedAppProtectionCollectionResponse) {
    m := &DefaultManagedAppProtectionCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDefaultManagedAppProtectionCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultManagedAppProtectionCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultManagedAppProtectionCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultManagedAppProtectionCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDefaultManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DefaultManagedAppProtectionable, len(val))
            for i, v := range val {
                res[i] = v.(DefaultManagedAppProtectionable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DefaultManagedAppProtectionCollectionResponse) GetValue()([]DefaultManagedAppProtectionable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DefaultManagedAppProtectionCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DefaultManagedAppProtectionCollectionResponse) SetValue(value []DefaultManagedAppProtectionable)() {
    m.value = value
}
