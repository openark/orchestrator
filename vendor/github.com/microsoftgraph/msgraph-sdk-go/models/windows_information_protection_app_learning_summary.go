package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionAppLearningSummary windows Information Protection AppLearning Summary entity.
type WindowsInformationProtectionAppLearningSummary struct {
    Entity
    // Application Name
    applicationName *string
    // Possible types of Application
    applicationType *ApplicationType
    // Device Count
    deviceCount *int32
}
// NewWindowsInformationProtectionAppLearningSummary instantiates a new windowsInformationProtectionAppLearningSummary and sets the default values.
func NewWindowsInformationProtectionAppLearningSummary()(*WindowsInformationProtectionAppLearningSummary) {
    m := &WindowsInformationProtectionAppLearningSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionAppLearningSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionAppLearningSummary(), nil
}
// GetApplicationName gets the applicationName property value. Application Name
func (m *WindowsInformationProtectionAppLearningSummary) GetApplicationName()(*string) {
    return m.applicationName
}
// GetApplicationType gets the applicationType property value. Possible types of Application
func (m *WindowsInformationProtectionAppLearningSummary) GetApplicationType()(*ApplicationType) {
    return m.applicationType
}
// GetDeviceCount gets the deviceCount property value. Device Count
func (m *WindowsInformationProtectionAppLearningSummary) GetDeviceCount()(*int32) {
    return m.deviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionAppLearningSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applicationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationName(val)
        }
        return nil
    }
    res["applicationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationType(val.(*ApplicationType))
        }
        return nil
    }
    res["deviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceCount(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionAppLearningSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("applicationName", m.GetApplicationName())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationType() != nil {
        cast := (*m.GetApplicationType()).String()
        err = writer.WriteStringValue("applicationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationName sets the applicationName property value. Application Name
func (m *WindowsInformationProtectionAppLearningSummary) SetApplicationName(value *string)() {
    m.applicationName = value
}
// SetApplicationType sets the applicationType property value. Possible types of Application
func (m *WindowsInformationProtectionAppLearningSummary) SetApplicationType(value *ApplicationType)() {
    m.applicationType = value
}
// SetDeviceCount sets the deviceCount property value. Device Count
func (m *WindowsInformationProtectionAppLearningSummary) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
