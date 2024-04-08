package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEvidence 
type DeviceEvidence struct {
    AlertEvidence
    // A unique identifier assigned to a device by Azure Active Directory (Azure AD) when device is Azure AD-joined.
    azureAdDeviceId *string
    // State of the Defender AntiMalware engine. The possible values are: notReporting, disabled, notUpdated, updated, unknown, notSupported, unknownFutureValue.
    defenderAvStatus *DefenderAvStatus
    // The fully qualified domain name (FQDN) for the device.
    deviceDnsName *string
    // The date and time when the device was first seen.
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The health state of the device.The possible values are: active, inactive, impairedCommunication, noSensorData, noSensorDataImpairedCommunication, unknown, unknownFutureValue.
    healthStatus *DeviceHealthStatus
    // Users that were logged on the machine during the time of the alert.
    loggedOnUsers []LoggedOnUserable
    // A unique identifier assigned to a device by Microsoft Defender for Endpoint.
    mdeDeviceId *string
    // The status of the machine onboarding to Microsoft Defender for Endpoint.The possible values are: insufficientInfo, onboarded, canBeOnboarded, unsupported, unknownFutureValue.
    onboardingStatus *OnboardingStatus
    // The build version for the operating system the device is running.
    osBuild *int64
    // The operating system platform the device is running.
    osPlatform *string
    // The ID of the role-based access control (RBAC) device group.
    rbacGroupId *int32
    // The name of the RBAC device group.
    rbacGroupName *string
    // Risk score as evaluated by Microsoft Defender for Endpoint. The possible values are: none, informational, low, medium, high, unknownFutureValue.
    riskScore *DeviceRiskScore
    // The version of the operating system platform.
    version *string
    // Metadata of the virtual machine (VM) on which Microsoft Defender for Endpoint is running.
    vmMetadata VmMetadataable
}
// NewDeviceEvidence instantiates a new DeviceEvidence and sets the default values.
func NewDeviceEvidence()(*DeviceEvidence) {
    m := &DeviceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateDeviceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEvidence(), nil
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. A unique identifier assigned to a device by Azure Active Directory (Azure AD) when device is Azure AD-joined.
func (m *DeviceEvidence) GetAzureAdDeviceId()(*string) {
    return m.azureAdDeviceId
}
// GetDefenderAvStatus gets the defenderAvStatus property value. State of the Defender AntiMalware engine. The possible values are: notReporting, disabled, notUpdated, updated, unknown, notSupported, unknownFutureValue.
func (m *DeviceEvidence) GetDefenderAvStatus()(*DefenderAvStatus) {
    return m.defenderAvStatus
}
// GetDeviceDnsName gets the deviceDnsName property value. The fully qualified domain name (FQDN) for the device.
func (m *DeviceEvidence) GetDeviceDnsName()(*string) {
    return m.deviceDnsName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    res["defenderAvStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAvStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAvStatus(val.(*DefenderAvStatus))
        }
        return nil
    }
    res["deviceDnsName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDnsName(val)
        }
        return nil
    }
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*DeviceHealthStatus))
        }
        return nil
    }
    res["loggedOnUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLoggedOnUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LoggedOnUserable, len(val))
            for i, v := range val {
                res[i] = v.(LoggedOnUserable)
            }
            m.SetLoggedOnUsers(res)
        }
        return nil
    }
    res["mdeDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdeDeviceId(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*OnboardingStatus))
        }
        return nil
    }
    res["osBuild"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuild(val)
        }
        return nil
    }
    res["osPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsPlatform(val)
        }
        return nil
    }
    res["rbacGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRbacGroupId(val)
        }
        return nil
    }
    res["rbacGroupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRbacGroupName(val)
        }
        return nil
    }
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRiskScore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val.(*DeviceRiskScore))
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
    res["vmMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVmMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmMetadata(val.(VmMetadataable))
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The date and time when the device was first seen.
func (m *DeviceEvidence) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetHealthStatus gets the healthStatus property value. The health state of the device.The possible values are: active, inactive, impairedCommunication, noSensorData, noSensorDataImpairedCommunication, unknown, unknownFutureValue.
func (m *DeviceEvidence) GetHealthStatus()(*DeviceHealthStatus) {
    return m.healthStatus
}
// GetLoggedOnUsers gets the loggedOnUsers property value. Users that were logged on the machine during the time of the alert.
func (m *DeviceEvidence) GetLoggedOnUsers()([]LoggedOnUserable) {
    return m.loggedOnUsers
}
// GetMdeDeviceId gets the mdeDeviceId property value. A unique identifier assigned to a device by Microsoft Defender for Endpoint.
func (m *DeviceEvidence) GetMdeDeviceId()(*string) {
    return m.mdeDeviceId
}
// GetOnboardingStatus gets the onboardingStatus property value. The status of the machine onboarding to Microsoft Defender for Endpoint.The possible values are: insufficientInfo, onboarded, canBeOnboarded, unsupported, unknownFutureValue.
func (m *DeviceEvidence) GetOnboardingStatus()(*OnboardingStatus) {
    return m.onboardingStatus
}
// GetOsBuild gets the osBuild property value. The build version for the operating system the device is running.
func (m *DeviceEvidence) GetOsBuild()(*int64) {
    return m.osBuild
}
// GetOsPlatform gets the osPlatform property value. The operating system platform the device is running.
func (m *DeviceEvidence) GetOsPlatform()(*string) {
    return m.osPlatform
}
// GetRbacGroupId gets the rbacGroupId property value. The ID of the role-based access control (RBAC) device group.
func (m *DeviceEvidence) GetRbacGroupId()(*int32) {
    return m.rbacGroupId
}
// GetRbacGroupName gets the rbacGroupName property value. The name of the RBAC device group.
func (m *DeviceEvidence) GetRbacGroupName()(*string) {
    return m.rbacGroupName
}
// GetRiskScore gets the riskScore property value. Risk score as evaluated by Microsoft Defender for Endpoint. The possible values are: none, informational, low, medium, high, unknownFutureValue.
func (m *DeviceEvidence) GetRiskScore()(*DeviceRiskScore) {
    return m.riskScore
}
// GetVersion gets the version property value. The version of the operating system platform.
func (m *DeviceEvidence) GetVersion()(*string) {
    return m.version
}
// GetVmMetadata gets the vmMetadata property value. Metadata of the virtual machine (VM) on which Microsoft Defender for Endpoint is running.
func (m *DeviceEvidence) GetVmMetadata()(VmMetadataable) {
    return m.vmMetadata
}
// Serialize serializes information the current object
func (m *DeviceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAvStatus() != nil {
        cast := (*m.GetDefenderAvStatus()).String()
        err = writer.WriteStringValue("defenderAvStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceDnsName", m.GetDeviceDnsName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLoggedOnUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLoggedOnUsers()))
        for i, v := range m.GetLoggedOnUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("loggedOnUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdeDeviceId", m.GetMdeDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetOnboardingStatus() != nil {
        cast := (*m.GetOnboardingStatus()).String()
        err = writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("osBuild", m.GetOsBuild())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osPlatform", m.GetOsPlatform())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rbacGroupId", m.GetRbacGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rbacGroupName", m.GetRbacGroupName())
        if err != nil {
            return err
        }
    }
    if m.GetRiskScore() != nil {
        cast := (*m.GetRiskScore()).String()
        err = writer.WriteStringValue("riskScore", &cast)
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
    {
        err = writer.WriteObjectValue("vmMetadata", m.GetVmMetadata())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. A unique identifier assigned to a device by Azure Active Directory (Azure AD) when device is Azure AD-joined.
func (m *DeviceEvidence) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
// SetDefenderAvStatus sets the defenderAvStatus property value. State of the Defender AntiMalware engine. The possible values are: notReporting, disabled, notUpdated, updated, unknown, notSupported, unknownFutureValue.
func (m *DeviceEvidence) SetDefenderAvStatus(value *DefenderAvStatus)() {
    m.defenderAvStatus = value
}
// SetDeviceDnsName sets the deviceDnsName property value. The fully qualified domain name (FQDN) for the device.
func (m *DeviceEvidence) SetDeviceDnsName(value *string)() {
    m.deviceDnsName = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The date and time when the device was first seen.
func (m *DeviceEvidence) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetHealthStatus sets the healthStatus property value. The health state of the device.The possible values are: active, inactive, impairedCommunication, noSensorData, noSensorDataImpairedCommunication, unknown, unknownFutureValue.
func (m *DeviceEvidence) SetHealthStatus(value *DeviceHealthStatus)() {
    m.healthStatus = value
}
// SetLoggedOnUsers sets the loggedOnUsers property value. Users that were logged on the machine during the time of the alert.
func (m *DeviceEvidence) SetLoggedOnUsers(value []LoggedOnUserable)() {
    m.loggedOnUsers = value
}
// SetMdeDeviceId sets the mdeDeviceId property value. A unique identifier assigned to a device by Microsoft Defender for Endpoint.
func (m *DeviceEvidence) SetMdeDeviceId(value *string)() {
    m.mdeDeviceId = value
}
// SetOnboardingStatus sets the onboardingStatus property value. The status of the machine onboarding to Microsoft Defender for Endpoint.The possible values are: insufficientInfo, onboarded, canBeOnboarded, unsupported, unknownFutureValue.
func (m *DeviceEvidence) SetOnboardingStatus(value *OnboardingStatus)() {
    m.onboardingStatus = value
}
// SetOsBuild sets the osBuild property value. The build version for the operating system the device is running.
func (m *DeviceEvidence) SetOsBuild(value *int64)() {
    m.osBuild = value
}
// SetOsPlatform sets the osPlatform property value. The operating system platform the device is running.
func (m *DeviceEvidence) SetOsPlatform(value *string)() {
    m.osPlatform = value
}
// SetRbacGroupId sets the rbacGroupId property value. The ID of the role-based access control (RBAC) device group.
func (m *DeviceEvidence) SetRbacGroupId(value *int32)() {
    m.rbacGroupId = value
}
// SetRbacGroupName sets the rbacGroupName property value. The name of the RBAC device group.
func (m *DeviceEvidence) SetRbacGroupName(value *string)() {
    m.rbacGroupName = value
}
// SetRiskScore sets the riskScore property value. Risk score as evaluated by Microsoft Defender for Endpoint. The possible values are: none, informational, low, medium, high, unknownFutureValue.
func (m *DeviceEvidence) SetRiskScore(value *DeviceRiskScore)() {
    m.riskScore = value
}
// SetVersion sets the version property value. The version of the operating system platform.
func (m *DeviceEvidence) SetVersion(value *string)() {
    m.version = value
}
// SetVmMetadata sets the vmMetadata property value. Metadata of the virtual machine (VM) on which Microsoft Defender for Endpoint is running.
func (m *DeviceEvidence) SetVmMetadata(value VmMetadataable)() {
    m.vmMetadata = value
}
