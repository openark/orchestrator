package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EditionUpgradeConfiguration 
type EditionUpgradeConfiguration struct {
    DeviceConfiguration
    // Edition Upgrade License File Content.
    license *string
    // Edition Upgrade License type
    licenseType *EditionUpgradeLicenseType
    // Edition Upgrade Product Key.
    productKey *string
    // Windows 10 Edition type.
    targetEdition *Windows10EditionType
}
// NewEditionUpgradeConfiguration instantiates a new EditionUpgradeConfiguration and sets the default values.
func NewEditionUpgradeConfiguration()(*EditionUpgradeConfiguration) {
    m := &EditionUpgradeConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.editionUpgradeConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEditionUpgradeConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEditionUpgradeConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEditionUpgradeConfiguration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EditionUpgradeConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["license"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicense(val)
        }
        return nil
    }
    res["licenseType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEditionUpgradeLicenseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseType(val.(*EditionUpgradeLicenseType))
        }
        return nil
    }
    res["productKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["targetEdition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWindows10EditionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetEdition(val.(*Windows10EditionType))
        }
        return nil
    }
    return res
}
// GetLicense gets the license property value. Edition Upgrade License File Content.
func (m *EditionUpgradeConfiguration) GetLicense()(*string) {
    return m.license
}
// GetLicenseType gets the licenseType property value. Edition Upgrade License type
func (m *EditionUpgradeConfiguration) GetLicenseType()(*EditionUpgradeLicenseType) {
    return m.licenseType
}
// GetProductKey gets the productKey property value. Edition Upgrade Product Key.
func (m *EditionUpgradeConfiguration) GetProductKey()(*string) {
    return m.productKey
}
// GetTargetEdition gets the targetEdition property value. Windows 10 Edition type.
func (m *EditionUpgradeConfiguration) GetTargetEdition()(*Windows10EditionType) {
    return m.targetEdition
}
// Serialize serializes information the current object
func (m *EditionUpgradeConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("license", m.GetLicense())
        if err != nil {
            return err
        }
    }
    if m.GetLicenseType() != nil {
        cast := (*m.GetLicenseType()).String()
        err = writer.WriteStringValue("licenseType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    if m.GetTargetEdition() != nil {
        cast := (*m.GetTargetEdition()).String()
        err = writer.WriteStringValue("targetEdition", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLicense sets the license property value. Edition Upgrade License File Content.
func (m *EditionUpgradeConfiguration) SetLicense(value *string)() {
    m.license = value
}
// SetLicenseType sets the licenseType property value. Edition Upgrade License type
func (m *EditionUpgradeConfiguration) SetLicenseType(value *EditionUpgradeLicenseType)() {
    m.licenseType = value
}
// SetProductKey sets the productKey property value. Edition Upgrade Product Key.
func (m *EditionUpgradeConfiguration) SetProductKey(value *string)() {
    m.productKey = value
}
// SetTargetEdition sets the targetEdition property value. Windows 10 Edition type.
func (m *EditionUpgradeConfiguration) SetTargetEdition(value *Windows10EditionType)() {
    m.targetEdition = value
}
