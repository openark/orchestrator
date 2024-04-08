package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidLobApp 
type AndroidLobApp struct {
    MobileLobApp
    // The value for the minimum applicable operating system.
    minimumSupportedOperatingSystem AndroidMinimumOperatingSystemable
    // The package identifier.
    packageId *string
    // The version code of Android Line of Business (LoB) app.
    versionCode *string
    // The version name of Android Line of Business (LoB) app.
    versionName *string
}
// NewAndroidLobApp instantiates a new AndroidLobApp and sets the default values.
func NewAndroidLobApp()(*AndroidLobApp) {
    m := &AndroidLobApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.androidLobApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidLobApp(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
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
    res["versionCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionCode(val)
        }
        return nil
    }
    res["versionName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionName(val)
        }
        return nil
    }
    return res
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *AndroidLobApp) GetMinimumSupportedOperatingSystem()(AndroidMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetPackageId gets the packageId property value. The package identifier.
func (m *AndroidLobApp) GetPackageId()(*string) {
    return m.packageId
}
// GetVersionCode gets the versionCode property value. The version code of Android Line of Business (LoB) app.
func (m *AndroidLobApp) GetVersionCode()(*string) {
    return m.versionCode
}
// GetVersionName gets the versionName property value. The version name of Android Line of Business (LoB) app.
func (m *AndroidLobApp) GetVersionName()(*string) {
    return m.versionName
}
// Serialize serializes information the current object
func (m *AndroidLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
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
    {
        err = writer.WriteStringValue("versionCode", m.GetVersionCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("versionName", m.GetVersionName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. The value for the minimum applicable operating system.
func (m *AndroidLobApp) SetMinimumSupportedOperatingSystem(value AndroidMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetPackageId sets the packageId property value. The package identifier.
func (m *AndroidLobApp) SetPackageId(value *string)() {
    m.packageId = value
}
// SetVersionCode sets the versionCode property value. The version code of Android Line of Business (LoB) app.
func (m *AndroidLobApp) SetVersionCode(value *string)() {
    m.versionCode = value
}
// SetVersionName sets the versionName property value. The version name of Android Line of Business (LoB) app.
func (m *AndroidLobApp) SetVersionName(value *string)() {
    m.versionName = value
}
