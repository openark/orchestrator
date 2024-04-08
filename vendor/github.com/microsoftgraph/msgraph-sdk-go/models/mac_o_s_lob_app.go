package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSLobApp 
type MacOSLobApp struct {
    MobileLobApp
    // The build number of the package. This should match the package CFBundleShortVersionString of the .pkg file.
    buildNumber *string
    // The primary bundleId of the package.
    bundleId *string
    // List of ComplexType macOSLobChildApp objects. Represents the apps expected to be installed by the package.
    childApps []MacOSLobChildAppable
    // When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true for apps that use a self update feature.
    ignoreVersionDetection *bool
    // When TRUE, indicates that the app will be installed as managed (requires macOS 11.0 and other managed package restrictions). When FALSE, indicates that the app will be installed as unmanaged.
    installAsManaged *bool
    // The MD5 hash codes. This is empty if the package was uploaded directly. If the Intune App Wrapping Tool is used to create a .intunemac, this value can be found inside the Detection.xml file.
    md5Hash []string
    // The chunk size for MD5 hash. This is '0' or empty if the package was uploaded directly. If the Intune App Wrapping Tool is used to create a .intunemac, this value can be found inside the Detection.xml file.
    md5HashChunkSize *int32
    // ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
    minimumSupportedOperatingSystem MacOSMinimumOperatingSystemable
    // The version number of the package. This should match the package CFBundleVersion in the packageinfo file.
    versionNumber *string
}
// NewMacOSLobApp instantiates a new MacOSLobApp and sets the default values.
func NewMacOSLobApp()(*MacOSLobApp) {
    m := &MacOSLobApp{
        MobileLobApp: *NewMobileLobApp(),
    }
    odataTypeValue := "#microsoft.graph.macOSLobApp"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateMacOSLobAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSLobAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSLobApp(), nil
}
// GetBuildNumber gets the buildNumber property value. The build number of the package. This should match the package CFBundleShortVersionString of the .pkg file.
func (m *MacOSLobApp) GetBuildNumber()(*string) {
    return m.buildNumber
}
// GetBundleId gets the bundleId property value. The primary bundleId of the package.
func (m *MacOSLobApp) GetBundleId()(*string) {
    return m.bundleId
}
// GetChildApps gets the childApps property value. List of ComplexType macOSLobChildApp objects. Represents the apps expected to be installed by the package.
func (m *MacOSLobApp) GetChildApps()([]MacOSLobChildAppable) {
    return m.childApps
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSLobApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MobileLobApp.GetFieldDeserializers()
    res["buildNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuildNumber(val)
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
    res["childApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMacOSLobChildAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MacOSLobChildAppable, len(val))
            for i, v := range val {
                res[i] = v.(MacOSLobChildAppable)
            }
            m.SetChildApps(res)
        }
        return nil
    }
    res["ignoreVersionDetection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIgnoreVersionDetection(val)
        }
        return nil
    }
    res["installAsManaged"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallAsManaged(val)
        }
        return nil
    }
    res["md5Hash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMd5Hash(res)
        }
        return nil
    }
    res["md5HashChunkSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMd5HashChunkSize(val)
        }
        return nil
    }
    res["minimumSupportedOperatingSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMacOSMinimumOperatingSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSupportedOperatingSystem(val.(MacOSMinimumOperatingSystemable))
        }
        return nil
    }
    res["versionNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionNumber(val)
        }
        return nil
    }
    return res
}
// GetIgnoreVersionDetection gets the ignoreVersionDetection property value. When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true for apps that use a self update feature.
func (m *MacOSLobApp) GetIgnoreVersionDetection()(*bool) {
    return m.ignoreVersionDetection
}
// GetInstallAsManaged gets the installAsManaged property value. When TRUE, indicates that the app will be installed as managed (requires macOS 11.0 and other managed package restrictions). When FALSE, indicates that the app will be installed as unmanaged.
func (m *MacOSLobApp) GetInstallAsManaged()(*bool) {
    return m.installAsManaged
}
// GetMd5Hash gets the md5Hash property value. The MD5 hash codes. This is empty if the package was uploaded directly. If the Intune App Wrapping Tool is used to create a .intunemac, this value can be found inside the Detection.xml file.
func (m *MacOSLobApp) GetMd5Hash()([]string) {
    return m.md5Hash
}
// GetMd5HashChunkSize gets the md5HashChunkSize property value. The chunk size for MD5 hash. This is '0' or empty if the package was uploaded directly. If the Intune App Wrapping Tool is used to create a .intunemac, this value can be found inside the Detection.xml file.
func (m *MacOSLobApp) GetMd5HashChunkSize()(*int32) {
    return m.md5HashChunkSize
}
// GetMinimumSupportedOperatingSystem gets the minimumSupportedOperatingSystem property value. ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
func (m *MacOSLobApp) GetMinimumSupportedOperatingSystem()(MacOSMinimumOperatingSystemable) {
    return m.minimumSupportedOperatingSystem
}
// GetVersionNumber gets the versionNumber property value. The version number of the package. This should match the package CFBundleVersion in the packageinfo file.
func (m *MacOSLobApp) GetVersionNumber()(*string) {
    return m.versionNumber
}
// Serialize serializes information the current object
func (m *MacOSLobApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MobileLobApp.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("buildNumber", m.GetBuildNumber())
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
    if m.GetChildApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildApps()))
        for i, v := range m.GetChildApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("childApps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("ignoreVersionDetection", m.GetIgnoreVersionDetection())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("installAsManaged", m.GetInstallAsManaged())
        if err != nil {
            return err
        }
    }
    if m.GetMd5Hash() != nil {
        err = writer.WriteCollectionOfStringValues("md5Hash", m.GetMd5Hash())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("md5HashChunkSize", m.GetMd5HashChunkSize())
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
        err = writer.WriteStringValue("versionNumber", m.GetVersionNumber())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBuildNumber sets the buildNumber property value. The build number of the package. This should match the package CFBundleShortVersionString of the .pkg file.
func (m *MacOSLobApp) SetBuildNumber(value *string)() {
    m.buildNumber = value
}
// SetBundleId sets the bundleId property value. The primary bundleId of the package.
func (m *MacOSLobApp) SetBundleId(value *string)() {
    m.bundleId = value
}
// SetChildApps sets the childApps property value. List of ComplexType macOSLobChildApp objects. Represents the apps expected to be installed by the package.
func (m *MacOSLobApp) SetChildApps(value []MacOSLobChildAppable)() {
    m.childApps = value
}
// SetIgnoreVersionDetection sets the ignoreVersionDetection property value. When TRUE, indicates that the app's version will NOT be used to detect if the app is installed on a device. When FALSE, indicates that the app's version will be used to detect if the app is installed on a device. Set this to true for apps that use a self update feature.
func (m *MacOSLobApp) SetIgnoreVersionDetection(value *bool)() {
    m.ignoreVersionDetection = value
}
// SetInstallAsManaged sets the installAsManaged property value. When TRUE, indicates that the app will be installed as managed (requires macOS 11.0 and other managed package restrictions). When FALSE, indicates that the app will be installed as unmanaged.
func (m *MacOSLobApp) SetInstallAsManaged(value *bool)() {
    m.installAsManaged = value
}
// SetMd5Hash sets the md5Hash property value. The MD5 hash codes. This is empty if the package was uploaded directly. If the Intune App Wrapping Tool is used to create a .intunemac, this value can be found inside the Detection.xml file.
func (m *MacOSLobApp) SetMd5Hash(value []string)() {
    m.md5Hash = value
}
// SetMd5HashChunkSize sets the md5HashChunkSize property value. The chunk size for MD5 hash. This is '0' or empty if the package was uploaded directly. If the Intune App Wrapping Tool is used to create a .intunemac, this value can be found inside the Detection.xml file.
func (m *MacOSLobApp) SetMd5HashChunkSize(value *int32)() {
    m.md5HashChunkSize = value
}
// SetMinimumSupportedOperatingSystem sets the minimumSupportedOperatingSystem property value. ComplexType macOSMinimumOperatingSystem that indicates the minimum operating system applicable for the application.
func (m *MacOSLobApp) SetMinimumSupportedOperatingSystem(value MacOSMinimumOperatingSystemable)() {
    m.minimumSupportedOperatingSystem = value
}
// SetVersionNumber sets the versionNumber property value. The version number of the package. This should match the package CFBundleVersion in the packageinfo file.
func (m *MacOSLobApp) SetVersionNumber(value *string)() {
    m.versionNumber = value
}
