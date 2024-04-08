package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEnrollmentPlatformRestriction platform specific enrollment restrictions
type DeviceEnrollmentPlatformRestriction struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Max OS version supported
    osMaximumVersion *string
    // Min OS version supported
    osMinimumVersion *string
    // Block personally owned devices from enrolling
    personalDeviceEnrollmentBlocked *bool
    // Block the platform from enrolling
    platformBlocked *bool
}
// NewDeviceEnrollmentPlatformRestriction instantiates a new deviceEnrollmentPlatformRestriction and sets the default values.
func NewDeviceEnrollmentPlatformRestriction()(*DeviceEnrollmentPlatformRestriction) {
    m := &DeviceEnrollmentPlatformRestriction{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEnrollmentPlatformRestrictionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEnrollmentPlatformRestriction(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceEnrollmentPlatformRestriction) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEnrollmentPlatformRestriction) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["osMaximumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMaximumVersion(val)
        }
        return nil
    }
    res["osMinimumVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsMinimumVersion(val)
        }
        return nil
    }
    res["personalDeviceEnrollmentBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonalDeviceEnrollmentBlocked(val)
        }
        return nil
    }
    res["platformBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformBlocked(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceEnrollmentPlatformRestriction) GetOdataType()(*string) {
    return m.odataType
}
// GetOsMaximumVersion gets the osMaximumVersion property value. Max OS version supported
func (m *DeviceEnrollmentPlatformRestriction) GetOsMaximumVersion()(*string) {
    return m.osMaximumVersion
}
// GetOsMinimumVersion gets the osMinimumVersion property value. Min OS version supported
func (m *DeviceEnrollmentPlatformRestriction) GetOsMinimumVersion()(*string) {
    return m.osMinimumVersion
}
// GetPersonalDeviceEnrollmentBlocked gets the personalDeviceEnrollmentBlocked property value. Block personally owned devices from enrolling
func (m *DeviceEnrollmentPlatformRestriction) GetPersonalDeviceEnrollmentBlocked()(*bool) {
    return m.personalDeviceEnrollmentBlocked
}
// GetPlatformBlocked gets the platformBlocked property value. Block the platform from enrolling
func (m *DeviceEnrollmentPlatformRestriction) GetPlatformBlocked()(*bool) {
    return m.platformBlocked
}
// Serialize serializes information the current object
func (m *DeviceEnrollmentPlatformRestriction) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("osMaximumVersion", m.GetOsMaximumVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("osMinimumVersion", m.GetOsMinimumVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("personalDeviceEnrollmentBlocked", m.GetPersonalDeviceEnrollmentBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("platformBlocked", m.GetPlatformBlocked())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceEnrollmentPlatformRestriction) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceEnrollmentPlatformRestriction) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOsMaximumVersion sets the osMaximumVersion property value. Max OS version supported
func (m *DeviceEnrollmentPlatformRestriction) SetOsMaximumVersion(value *string)() {
    m.osMaximumVersion = value
}
// SetOsMinimumVersion sets the osMinimumVersion property value. Min OS version supported
func (m *DeviceEnrollmentPlatformRestriction) SetOsMinimumVersion(value *string)() {
    m.osMinimumVersion = value
}
// SetPersonalDeviceEnrollmentBlocked sets the personalDeviceEnrollmentBlocked property value. Block personally owned devices from enrolling
func (m *DeviceEnrollmentPlatformRestriction) SetPersonalDeviceEnrollmentBlocked(value *bool)() {
    m.personalDeviceEnrollmentBlocked = value
}
// SetPlatformBlocked sets the platformBlocked property value. Block the platform from enrolling
func (m *DeviceEnrollmentPlatformRestriction) SetPlatformBlocked(value *bool)() {
    m.platformBlocked = value
}
