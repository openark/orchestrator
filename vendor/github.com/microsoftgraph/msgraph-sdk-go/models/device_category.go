package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceCategory 
type DeviceCategory struct {
    Entity
    // Optional description for the device category.
    description *string
    // Display name for the device category.
    displayName *string
}
// NewDeviceCategory instantiates a new deviceCategory and sets the default values.
func NewDeviceCategory()(*DeviceCategory) {
    m := &DeviceCategory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceCategoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceCategoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceCategory(), nil
}
// GetDescription gets the description property value. Optional description for the device category.
func (m *DeviceCategory) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name for the device category.
func (m *DeviceCategory) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceCategory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *DeviceCategory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Optional description for the device category.
func (m *DeviceCategory) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name for the device category.
func (m *DeviceCategory) SetDisplayName(value *string)() {
    m.displayName = value
}
