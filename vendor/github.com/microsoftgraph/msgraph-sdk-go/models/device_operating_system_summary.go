package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceOperatingSystemSummary device operating system summary.
type DeviceOperatingSystemSummary struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid values -1 to 2147483647
    androidCorporateWorkProfileCount *int32
    // Number of android device count.
    androidCount *int32
    // Number of dedicated Android devices.
    androidDedicatedCount *int32
    // Number of device admin Android devices.
    androidDeviceAdminCount *int32
    // Number of fully managed Android devices.
    androidFullyManagedCount *int32
    // Number of work profile Android devices.
    androidWorkProfileCount *int32
    // Number of ConfigMgr managed devices.
    configMgrDeviceCount *int32
    // Number of iOS device count.
    iosCount *int32
    // Number of Mac OS X device count.
    macOSCount *int32
    // The OdataType property
    odataType *string
    // Number of unknown device count.
    unknownCount *int32
    // Number of Windows device count.
    windowsCount *int32
    // Number of Windows mobile device count.
    windowsMobileCount *int32
}
// NewDeviceOperatingSystemSummary instantiates a new deviceOperatingSystemSummary and sets the default values.
func NewDeviceOperatingSystemSummary()(*DeviceOperatingSystemSummary) {
    m := &DeviceOperatingSystemSummary{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceOperatingSystemSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceOperatingSystemSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceOperatingSystemSummary(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceOperatingSystemSummary) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAndroidCorporateWorkProfileCount gets the androidCorporateWorkProfileCount property value. The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid values -1 to 2147483647
func (m *DeviceOperatingSystemSummary) GetAndroidCorporateWorkProfileCount()(*int32) {
    return m.androidCorporateWorkProfileCount
}
// GetAndroidCount gets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) GetAndroidCount()(*int32) {
    return m.androidCount
}
// GetAndroidDedicatedCount gets the androidDedicatedCount property value. Number of dedicated Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidDedicatedCount()(*int32) {
    return m.androidDedicatedCount
}
// GetAndroidDeviceAdminCount gets the androidDeviceAdminCount property value. Number of device admin Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidDeviceAdminCount()(*int32) {
    return m.androidDeviceAdminCount
}
// GetAndroidFullyManagedCount gets the androidFullyManagedCount property value. Number of fully managed Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidFullyManagedCount()(*int32) {
    return m.androidFullyManagedCount
}
// GetAndroidWorkProfileCount gets the androidWorkProfileCount property value. Number of work profile Android devices.
func (m *DeviceOperatingSystemSummary) GetAndroidWorkProfileCount()(*int32) {
    return m.androidWorkProfileCount
}
// GetConfigMgrDeviceCount gets the configMgrDeviceCount property value. Number of ConfigMgr managed devices.
func (m *DeviceOperatingSystemSummary) GetConfigMgrDeviceCount()(*int32) {
    return m.configMgrDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceOperatingSystemSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["androidCorporateWorkProfileCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidCorporateWorkProfileCount(val)
        }
        return nil
    }
    res["androidCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidCount(val)
        }
        return nil
    }
    res["androidDedicatedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDedicatedCount(val)
        }
        return nil
    }
    res["androidDeviceAdminCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidDeviceAdminCount(val)
        }
        return nil
    }
    res["androidFullyManagedCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidFullyManagedCount(val)
        }
        return nil
    }
    res["androidWorkProfileCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAndroidWorkProfileCount(val)
        }
        return nil
    }
    res["configMgrDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfigMgrDeviceCount(val)
        }
        return nil
    }
    res["iosCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIosCount(val)
        }
        return nil
    }
    res["macOSCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacOSCount(val)
        }
        return nil
    }
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
    res["unknownCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnknownCount(val)
        }
        return nil
    }
    res["windowsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsCount(val)
        }
        return nil
    }
    res["windowsMobileCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsMobileCount(val)
        }
        return nil
    }
    return res
}
// GetIosCount gets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) GetIosCount()(*int32) {
    return m.iosCount
}
// GetMacOSCount gets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) GetMacOSCount()(*int32) {
    return m.macOSCount
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceOperatingSystemSummary) GetOdataType()(*string) {
    return m.odataType
}
// GetUnknownCount gets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) GetUnknownCount()(*int32) {
    return m.unknownCount
}
// GetWindowsCount gets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) GetWindowsCount()(*int32) {
    return m.windowsCount
}
// GetWindowsMobileCount gets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) GetWindowsMobileCount()(*int32) {
    return m.windowsMobileCount
}
// Serialize serializes information the current object
func (m *DeviceOperatingSystemSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("androidCorporateWorkProfileCount", m.GetAndroidCorporateWorkProfileCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidCount", m.GetAndroidCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidDedicatedCount", m.GetAndroidDedicatedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidDeviceAdminCount", m.GetAndroidDeviceAdminCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidFullyManagedCount", m.GetAndroidFullyManagedCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("androidWorkProfileCount", m.GetAndroidWorkProfileCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("configMgrDeviceCount", m.GetConfigMgrDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("iosCount", m.GetIosCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("macOSCount", m.GetMacOSCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("unknownCount", m.GetUnknownCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsCount", m.GetWindowsCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("windowsMobileCount", m.GetWindowsMobileCount())
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
func (m *DeviceOperatingSystemSummary) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAndroidCorporateWorkProfileCount sets the androidCorporateWorkProfileCount property value. The count of Corporate work profile Android devices. Also known as Corporate Owned Personally Enabled (COPE). Valid values -1 to 2147483647
func (m *DeviceOperatingSystemSummary) SetAndroidCorporateWorkProfileCount(value *int32)() {
    m.androidCorporateWorkProfileCount = value
}
// SetAndroidCount sets the androidCount property value. Number of android device count.
func (m *DeviceOperatingSystemSummary) SetAndroidCount(value *int32)() {
    m.androidCount = value
}
// SetAndroidDedicatedCount sets the androidDedicatedCount property value. Number of dedicated Android devices.
func (m *DeviceOperatingSystemSummary) SetAndroidDedicatedCount(value *int32)() {
    m.androidDedicatedCount = value
}
// SetAndroidDeviceAdminCount sets the androidDeviceAdminCount property value. Number of device admin Android devices.
func (m *DeviceOperatingSystemSummary) SetAndroidDeviceAdminCount(value *int32)() {
    m.androidDeviceAdminCount = value
}
// SetAndroidFullyManagedCount sets the androidFullyManagedCount property value. Number of fully managed Android devices.
func (m *DeviceOperatingSystemSummary) SetAndroidFullyManagedCount(value *int32)() {
    m.androidFullyManagedCount = value
}
// SetAndroidWorkProfileCount sets the androidWorkProfileCount property value. Number of work profile Android devices.
func (m *DeviceOperatingSystemSummary) SetAndroidWorkProfileCount(value *int32)() {
    m.androidWorkProfileCount = value
}
// SetConfigMgrDeviceCount sets the configMgrDeviceCount property value. Number of ConfigMgr managed devices.
func (m *DeviceOperatingSystemSummary) SetConfigMgrDeviceCount(value *int32)() {
    m.configMgrDeviceCount = value
}
// SetIosCount sets the iosCount property value. Number of iOS device count.
func (m *DeviceOperatingSystemSummary) SetIosCount(value *int32)() {
    m.iosCount = value
}
// SetMacOSCount sets the macOSCount property value. Number of Mac OS X device count.
func (m *DeviceOperatingSystemSummary) SetMacOSCount(value *int32)() {
    m.macOSCount = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceOperatingSystemSummary) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUnknownCount sets the unknownCount property value. Number of unknown device count.
func (m *DeviceOperatingSystemSummary) SetUnknownCount(value *int32)() {
    m.unknownCount = value
}
// SetWindowsCount sets the windowsCount property value. Number of Windows device count.
func (m *DeviceOperatingSystemSummary) SetWindowsCount(value *int32)() {
    m.windowsCount = value
}
// SetWindowsMobileCount sets the windowsMobileCount property value. Number of Windows mobile device count.
func (m *DeviceOperatingSystemSummary) SetWindowsMobileCount(value *int32)() {
    m.windowsMobileCount = value
}
