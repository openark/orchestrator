package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosStoreApp 
type IosStoreApp struct {
    MobileApp
    // Contains properties of the possible iOS device types the mobile app can run on.
    applicableDeviceType IosDeviceTypeable
    // The Apple App Store URL
    appStoreUrl *string
    // The Identity Name.
    bundleId *string
    // The value for the minimum applicable operating system.
    minimumSupportedOperatingSystem IosMinimumOperatingSystemable
}
// NewIosStoreApp instantiates a new IosStoreApp and sets the default values.
func NewIosStoreApp()(*IosStoreApp) {
    m := &IosStoreApp{
        MobileApp: *NewMobileApp(),
    }
    odataTypeValue := "#microsoft.graph.iosStoreApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosStoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosStoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosStoreApp(), nil
}
// GetApplicableDeviceType gets the applicableDeviceType property value. Contains properties of the possible iOS device types the mobile app can run on.
func (m *IosStoreApp) GetApplicableDeviceType()(IosDeviceTypeable) {
    return m.applicableDeviceType
}
// GetAppStoreUrl gets the appStoreUrl property value. The Apple App Store URL
func (m *IosStoreApp) GetAppStoreUrl()(*string) {
    return m.appStoreUrl
}
// GetBundleId gets the bundleId property value. The Identity Name.
func (m *IosStoreApp) GetBundleId()(*string) {
    return m.bundleId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileApp.GetFieldDeserializers()
    res["applicableDeviceType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosDeviceTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicableDeviceType(val.(IosDeviceTypeable))
        }
        return nil
    }
    res["appStoreUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppStoreUrl(val)
        }
        return nil
    }
    res["bundleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBundleId(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIosMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(IosMinimumOperatingSystemable))
        }
        return nil
    }
    return res
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *IosStoreApp) GetMinimumSupportedOperatingSystem()(IosMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// Serialize serializes information the current object
func (m *IosStoreApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applicableDeviceType", m.GetApplicableDeviceType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bundleId", m.GetBundleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("minimumSupportedOperatingSystem", m.GetMinimumSupportedOperatingSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicableDeviceType sets the applicableDeviceType property value. Contains properties of the possible iOS device types the mobile app can run on.
func (m *IosStoreApp) SetApplicableDeviceType(value IosDeviceTypeable)() {
    m.applicableDeviceType = value
}
// SetAppStoreUrl sets the appStoreUrl property value. The Apple App Store URL
func (m *IosStoreApp) SetAppStoreUrl(value *string)() {
    m.appStoreUrl = value
}
// SetBundleId sets the bundleId property value. The Identity Name.
func (m *IosStoreApp) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *IosStoreApp) SetMinimumSupportedOperatingSystem(value IosMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
