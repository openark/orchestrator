package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationUserStatusCollectionResponse 
type DeviceConfigurationUserStatusCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DeviceConfigurationUserStatusable
}
// NewDeviceConfigurationUserStatusCollectionResponse instantiates a new DeviceConfigurationUserStatusCollectionResponse and sets the default values.
func NewDeviceConfigurationUserStatusCollectionResponse()(*DeviceConfigurationUserStatusCollectionResponse) {
    m := &DeviceConfigurationUserStatusCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceConfigurationUserStatusCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationUserStatusCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationUserStatusCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationUserStatusCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationUserStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationUserStatusable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationUserStatusable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceConfigurationUserStatusCollectionResponse) GetValue()([]DeviceConfigurationUserStatusable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceConfigurationUserStatusCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceConfigurationUserStatusCollectionResponse) SetValue(value []DeviceConfigurationUserStatusable)() {
    m.value = value
}
