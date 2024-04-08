package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IosManagedAppProtection 
type IosManagedAppProtection struct {
    TargetedManagedAppProtection
    // Represents the level to which app data is encrypted for managed apps
    appDataEncryptionType *ManagedAppDataEncryptionType
    // List of apps to which the policy is deployed.
    apps []ManagedMobileAppable
    // A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
    customBrowserProtocol *string
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32
    // Navigation property to deployment summary of the configuration.
    deploymentSummary ManagedAppPolicyDeploymentSummaryable
    // Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
    faceIdBlocked *bool
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredSdkVersion *string
}
// NewIosManagedAppProtection instantiates a new IosManagedAppProtection and sets the default values.
func NewIosManagedAppProtection()(*IosManagedAppProtection) {
    m := &IosManagedAppProtection{
        TargetedManagedAppProtection: *NewTargetedManagedAppProtection(),
    }
    odataTypeValue := "#microsoft.graph.iosManagedAppProtection"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateIosManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIosManagedAppProtectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIosManagedAppProtection(), nil
}
// GetAppDataEncryptionType gets the appDataEncryptionType property value. Represents the level to which app data is encrypted for managed apps
func (m *IosManagedAppProtection) GetAppDataEncryptionType()(*ManagedAppDataEncryptionType) {
    return m.appDataEncryptionType
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *IosManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    return m.apps
}
// GetCustomBrowserProtocol gets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *IosManagedAppProtection) GetCustomBrowserProtocol()(*string) {
    return m.customBrowserProtocol
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *IosManagedAppProtection) GetDeployedAppCount()(*int32) {
    return m.deployedAppCount
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *IosManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    return m.deploymentSummary
}
// GetFaceIdBlocked gets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
func (m *IosManagedAppProtection) GetFaceIdBlocked()(*bool) {
    return m.faceIdBlocked
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IosManagedAppProtection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TargetedManagedAppProtection.GetFieldDeserializers()
    res["appDataEncryptionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataEncryptionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDataEncryptionType(val.(*ManagedAppDataEncryptionType))
        }
        return nil
    }
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    res["customBrowserProtocol"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomBrowserProtocol(val)
        }
        return nil
    }
    res["deployedAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["deploymentSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(ManagedAppPolicyDeploymentSummaryable))
        }
        return nil
    }
    res["faceIdBlocked"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaceIdBlocked(val)
        }
        return nil
    }
    res["minimumRequiredSdkVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredSdkVersion(val)
        }
        return nil
    }
    return res
}
// GetMinimumRequiredSdkVersion gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    return m.minimumRequiredSdkVersion
}
// Serialize serializes information the current object
func (m *IosManagedAppProtection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TargetedManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppDataEncryptionType() != nil {
        cast := (*m.GetAppDataEncryptionType()).String()
        err = writer.WriteStringValue("appDataEncryptionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customBrowserProtocol", m.GetCustomBrowserProtocol())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deployedAppCount", m.GetDeployedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSummary", m.GetDeploymentSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("faceIdBlocked", m.GetFaceIdBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredSdkVersion", m.GetMinimumRequiredSdkVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDataEncryptionType sets the appDataEncryptionType property value. Represents the level to which app data is encrypted for managed apps
func (m *IosManagedAppProtection) SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)() {
    m.appDataEncryptionType = value
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *IosManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    m.apps = value
}
// SetCustomBrowserProtocol sets the customBrowserProtocol property value. A custom browser protocol to open weblink on iOS. When this property is configured, ManagedBrowserToOpenLinksRequired should be true.
func (m *IosManagedAppProtection) SetCustomBrowserProtocol(value *string)() {
    m.customBrowserProtocol = value
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *IosManagedAppProtection) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *IosManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    m.deploymentSummary = value
}
// SetFaceIdBlocked sets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True.
func (m *IosManagedAppProtection) SetFaceIdBlocked(value *bool)() {
    m.faceIdBlocked = value
}
// SetMinimumRequiredSdkVersion sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *IosManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    m.minimumRequiredSdkVersion = value
}
