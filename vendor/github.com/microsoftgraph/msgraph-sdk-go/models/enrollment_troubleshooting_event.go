package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EnrollmentTroubleshootingEvent 
type EnrollmentTroubleshootingEvent struct {
    DeviceManagementTroubleshootingEvent
    // Azure AD device identifier.
    deviceId *string
    // Possible ways of adding a mobile device to management.
    enrollmentType *DeviceEnrollmentType
    // Top level failure categories for enrollment.
    failureCategory *DeviceEnrollmentFailureReason
    // Detailed failure reason.
    failureReason *string
    // Device identifier created or collected by Intune.
    managedDeviceIdentifier *string
    // Operating System.
    operatingSystem *string
    // OS Version.
    osVersion *string
    // Identifier for the user that tried to enroll the device.
    userId *string
}
// NewEnrollmentTroubleshootingEvent instantiates a new EnrollmentTroubleshootingEvent and sets the default values.
func NewEnrollmentTroubleshootingEvent()(*EnrollmentTroubleshootingEvent) {
    m := &EnrollmentTroubleshootingEvent{
        DeviceManagementTroubleshootingEvent: *NewDeviceManagementTroubleshootingEvent(),
    }
    return m
}
// CreateEnrollmentTroubleshootingEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEnrollmentTroubleshootingEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEnrollmentTroubleshootingEvent(), nil
}
// GetDeviceId gets the deviceId property value. Azure AD device identifier.
func (m *EnrollmentTroubleshootingEvent) GetDeviceId()(*string) {
    return m.deviceId
}
// GetEnrollmentType gets the enrollmentType property value. Possible ways of adding a mobile device to management.
func (m *EnrollmentTroubleshootingEvent) GetEnrollmentType()(*DeviceEnrollmentType) {
    return m.enrollmentType
}
// GetFailureCategory gets the failureCategory property value. Top level failure categories for enrollment.
func (m *EnrollmentTroubleshootingEvent) GetFailureCategory()(*DeviceEnrollmentFailureReason) {
    return m.failureCategory
}
// GetFailureReason gets the failureReason property value. Detailed failure reason.
func (m *EnrollmentTroubleshootingEvent) GetFailureReason()(*string) {
    return m.failureReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EnrollmentTroubleshootingEvent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceManagementTroubleshootingEvent.GetFieldDeserializers()
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["enrollmentType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnrollmentType(val.(*DeviceEnrollmentType))
        }
        return nil
    }
    res["failureCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceEnrollmentFailureReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureCategory(val.(*DeviceEnrollmentFailureReason))
        }
        return nil
    }
    res["failureReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureReason(val)
        }
        return nil
    }
    res["managedDeviceIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedDeviceIdentifier(val)
        }
        return nil
    }
    res["operatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystem(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetManagedDeviceIdentifier gets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *EnrollmentTroubleshootingEvent) GetManagedDeviceIdentifier()(*string) {
    return m.managedDeviceIdentifier
}
// GetOperatingSystem gets the operatingSystem property value. Operating System.
func (m *EnrollmentTroubleshootingEvent) GetOperatingSystem()(*string) {
    return m.operatingSystem
}
// GetOsVersion gets the osVersion property value. OS Version.
func (m *EnrollmentTroubleshootingEvent) GetOsVersion()(*string) {
    return m.osVersion
}
// GetUserId gets the userId property value. Identifier for the user that tried to enroll the device.
func (m *EnrollmentTroubleshootingEvent) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *EnrollmentTroubleshootingEvent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceManagementTroubleshootingEvent.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetEnrollmentType() != nil {
        cast := (*m.GetEnrollmentType()).String()
        err = writer.WriteStringValue("enrollmentType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFailureCategory() != nil {
        cast := (*m.GetFailureCategory()).String()
        err = writer.WriteStringValue("failureCategory", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("failureReason", m.GetFailureReason())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managedDeviceIdentifier", m.GetManagedDeviceIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operatingSystem", m.GetOperatingSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceId sets the deviceId property value. Azure AD device identifier.
func (m *EnrollmentTroubleshootingEvent) SetDeviceId(value *string)() {
    m.deviceId = value
}
// SetEnrollmentType sets the enrollmentType property value. Possible ways of adding a mobile device to management.
func (m *EnrollmentTroubleshootingEvent) SetEnrollmentType(value *DeviceEnrollmentType)() {
    m.enrollmentType = value
}
// SetFailureCategory sets the failureCategory property value. Top level failure categories for enrollment.
func (m *EnrollmentTroubleshootingEvent) SetFailureCategory(value *DeviceEnrollmentFailureReason)() {
    m.failureCategory = value
}
// SetFailureReason sets the failureReason property value. Detailed failure reason.
func (m *EnrollmentTroubleshootingEvent) SetFailureReason(value *string)() {
    m.failureReason = value
}
// SetManagedDeviceIdentifier sets the managedDeviceIdentifier property value. Device identifier created or collected by Intune.
func (m *EnrollmentTroubleshootingEvent) SetManagedDeviceIdentifier(value *string)() {
    m.managedDeviceIdentifier = value
}
// SetOperatingSystem sets the operatingSystem property value. Operating System.
func (m *EnrollmentTroubleshootingEvent) SetOperatingSystem(value *string)() {
    m.operatingSystem = value
}
// SetOsVersion sets the osVersion property value. OS Version.
func (m *EnrollmentTroubleshootingEvent) SetOsVersion(value *string)() {
    m.osVersion = value
}
// SetUserId sets the userId property value. Identifier for the user that tried to enroll the device.
func (m *EnrollmentTroubleshootingEvent) SetUserId(value *string)() {
    m.userId = value
}
