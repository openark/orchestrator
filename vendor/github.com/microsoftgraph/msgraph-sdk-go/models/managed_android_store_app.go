package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAndroidStoreApp 
type ManagedAndroidStoreApp struct {
    ManagedApp
    // The Android AppStoreUrl.
    appStoreUrl *string
    // Contains properties for the minimum operating system required for an Android mobile app.
    minimumSupportedOperatingSystem AndroidMinimumOperatingSystemable
    // The app's package ID.
    packageId *string
}
// NewManagedAndroidStoreApp instantiates a new ManagedAndroidStoreApp and sets the default values.
func NewManagedAndroidStoreApp()(*ManagedAndroidStoreApp) {
    m := &ManagedAndroidStoreApp{
        ManagedApp: *NewManagedApp(),
    }
    odataTypeValue := "#microsoft.graph.managedAndroidStoreApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateManagedAndroidStoreAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAndroidStoreAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewManagedAndroidStoreApp(), nil
}
// GetAppStoreUrl gets the appStoreUrl property value. The Android AppStoreUrl.
func (m *ManagedAndroidStoreApp) GetAppStoreUrl()(*string) {
    return m.appStoreUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAndroidStoreApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedApp.GetFieldDeserializers()
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
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAndroidMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(AndroidMinimumOperatingSystemable))
        }
        return nil
    }
    res["packageId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPackageId(val)
        }
        return nil
    }
    return res
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. Contains properties for the minimum operating system required for an Android mobile app.
func (m *ManagedAndroidStoreApp) GetMinimumSupportedOperatingSystem()(AndroidMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetPackageId gets the packageId property value. The app's package ID.
func (m *ManagedAndroidStoreApp) GetPackageId()(*string) {
    return m.packageId
}
// Serialize serializes information the current object
func (m *ManagedAndroidStoreApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appStoreUrl", m.GetAppStoreUrl())
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
    {
        err = writer.WriteStringValue("packageId", m.GetPackageId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppStoreUrl sets the appStoreUrl property value. The Android AppStoreUrl.
func (m *ManagedAndroidStoreApp) SetAppStoreUrl(value *string)() {
    m.appStoreUrl = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. Contains properties for the minimum operating system required for an Android mobile app.
func (m *ManagedAndroidStoreApp) SetMinimumSupportedOperatingSystem(value AndroidMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetPackageId sets the packageId property value. The app's package ID.
func (m *ManagedAndroidStoreApp) SetPackageId(value *string)() {
    m.packageId = value
}
