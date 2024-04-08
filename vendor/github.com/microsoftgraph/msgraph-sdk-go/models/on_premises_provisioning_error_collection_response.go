package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesProvisioningErrorCollectionResponse 
type OnPremisesProvisioningErrorCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []OnPremisesProvisioningErrorable
}
// NewOnPremisesProvisioningErrorCollectionResponse instantiates a new OnPremisesProvisioningErrorCollectionResponse and sets the default values.
func NewOnPremisesProvisioningErrorCollectionResponse()(*OnPremisesProvisioningErrorCollectionResponse) {
    m := &OnPremisesProvisioningErrorCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateOnPremisesProvisioningErrorCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesProvisioningErrorCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesProvisioningErrorCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesProvisioningErrorCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesProvisioningErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesProvisioningErrorable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesProvisioningErrorable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *OnPremisesProvisioningErrorCollectionResponse) GetValue()([]OnPremisesProvisioningErrorable) {
    return m.value
}
// Serialize serializes information the current object
func (m *OnPremisesProvisioningErrorCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *OnPremisesProvisioningErrorCollectionResponse) SetValue(value []OnPremisesProvisioningErrorable)() {
    m.value = value
}
