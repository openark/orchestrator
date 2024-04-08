package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse 
type DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []DeviceEnrollmentPlatformRestrictionsConfigurationable
}
// NewDeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse instantiates a new DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse and sets the default values.
func NewDeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse()(*DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse) {
    m := &DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceEnrollmentPlatformRestrictionsConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceEnrollmentPlatformRestrictionsConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceEnrollmentPlatformRestrictionsConfigurationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse) GetValue()([]DeviceEnrollmentPlatformRestrictionsConfigurationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *DeviceEnrollmentPlatformRestrictionsConfigurationCollectionResponse) SetValue(value []DeviceEnrollmentPlatformRestrictionsConfigurationable)() {
    m.value = value
}
