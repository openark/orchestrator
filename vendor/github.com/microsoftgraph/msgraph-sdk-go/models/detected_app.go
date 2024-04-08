package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DetectedApp a managed or unmanaged app that is installed on a managed device. Unmanaged apps will only appear for devices marked as corporate owned.
type DetectedApp struct {
    Entity
    // The number of devices that have installed this application
    deviceCount *int32
    // Name of the discovered application. Read-only
    displayName *string
    // The devices that have the discovered application installed
    managedDevices []ManagedDeviceable
    // Indicates the operating system / platform of the discovered application.  Some possible values are Windows, iOS, macOS. The default value is unknown (0).
    platform *DetectedAppPlatformType
    // Indicates the publisher of the discovered application. For example: 'Microsoft'.  The default value is an empty string.
    publisher *string
    // Discovered application size in bytes. Read-only
    sizeInByte *int64
    // Version of the discovered application. Read-only
    version *string
}
// NewDetectedApp instantiates a new detectedApp and sets the default values.
func NewDetectedApp()(*DetectedApp) {
    m := &DetectedApp{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDetectedAppFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDetectedAppFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDetectedApp(), nil
}
// GetDeviceCount gets the deviceCount property value. The number of devices that have installed this application
func (m *DetectedApp) GetDeviceCount()(*int32) {
    return m.deviceCount
}
// GetDisplayName gets the displayName property value. Name of the discovered application. Read-only
func (m *DetectedApp) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DetectedApp) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["managedDevices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceable)
            }
            m.SetManagedDevices(res)
        }
        return nil
    }
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDetectedAppPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*DetectedAppPlatformType))
        }
        return nil
    }
    res["publisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisher(val)
        }
        return nil
    }
    res["sizeInByte"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSizeInByte(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetManagedDevices gets the managedDevices property value. The devices that have the discovered application installed
func (m *DetectedApp) GetManagedDevices()([]ManagedDeviceable) {
    return m.managedDevices
}
// GetPlatform gets the platform property value. Indicates the operating system / platform of the discovered application.  Some possible values are Windows, iOS, macOS. The default value is unknown (0).
func (m *DetectedApp) GetPlatform()(*DetectedAppPlatformType) {
    return m.platform
}
// GetPublisher gets the publisher property value. Indicates the publisher of the discovered application. For example: 'Microsoft'.  The default value is an empty string.
func (m *DetectedApp) GetPublisher()(*string) {
    return m.publisher
}
// GetSizeInByte gets the sizeInByte property value. Discovered application size in bytes. Read-only
func (m *DetectedApp) GetSizeInByte()(*int64) {
    return m.sizeInByte
}
// GetVersion gets the version property value. Version of the discovered application. Read-only
func (m *DetectedApp) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *DetectedApp) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("deviceCount", m.GetDeviceCount())
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
    if m.GetManagedDevices() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisher", m.GetPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("sizeInByte", m.GetSizeInByte())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceCount sets the deviceCount property value. The number of devices that have installed this application
func (m *DetectedApp) SetDeviceCount(value *int32)() {
    m.deviceCount = value
}
// SetDisplayName sets the displayName property value. Name of the discovered application. Read-only
func (m *DetectedApp) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetManagedDevices sets the managedDevices property value. The devices that have the discovered application installed
func (m *DetectedApp) SetManagedDevices(value []ManagedDeviceable)() {
    m.managedDevices = value
}
// SetPlatform sets the platform property value. Indicates the operating system / platform of the discovered application.  Some possible values are Windows, iOS, macOS. The default value is unknown (0).
func (m *DetectedApp) SetPlatform(value *DetectedAppPlatformType)() {
    m.platform = value
}
// SetPublisher sets the publisher property value. Indicates the publisher of the discovered application. For example: 'Microsoft'.  The default value is an empty string.
func (m *DetectedApp) SetPublisher(value *string)() {
    m.publisher = value
}
// SetSizeInByte sets the sizeInByte property value. Discovered application size in bytes. Read-only
func (m *DetectedApp) SetSizeInByte(value *int64)() {
    m.sizeInByte = value
}
// SetVersion sets the version property value. Version of the discovered application. Read-only
func (m *DetectedApp) SetVersion(value *string)() {
    m.version = value
}
