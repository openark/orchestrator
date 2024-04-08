package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceHealthAttestationState 
type DeviceHealthAttestationState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate.
    attestationIdentityKey *string
    // On or Off of BitLocker Drive Encryption
    bitLockerStatus *string
    // The security version number of the Boot Application
    bootAppSecurityVersion *string
    // When bootDebugging is enabled, the device is used in development and testing
    bootDebugging *string
    // The security version number of the Boot Application
    bootManagerSecurityVersion *string
    // The version of the Boot Manager
    bootManagerVersion *string
    // The Boot Revision List that was loaded during initial boot on the attested device
    bootRevisionListInfo *string
    // When code integrity is enabled, code execution is restricted to integrity verified code
    codeIntegrity *string
    // The version of the Boot Manager
    codeIntegrityCheckVersion *string
    // The Code Integrity policy that is controlling the security of the boot environment
    codeIntegrityPolicy *string
    // The DHA report version. (Namespace version)
    contentNamespaceUrl *string
    // The HealthAttestation state schema version
    contentVersion *string
    // DEP Policy defines a set of hardware and software technologies that perform additional checks on memory
    dataExcutionPolicy *string
    // The DHA report version. (Namespace version)
    deviceHealthAttestationStatus *string
    // ELAM provides protection for the computers in your network when they start up
    earlyLaunchAntiMalwareDriverProtection *string
    // This attribute indicates if DHA is supported for the device
    healthAttestationSupportedStatus *string
    // This attribute appears if DHA-Service detects an integrity issue
    healthStatusMismatchInfo *string
    // The DateTime when device was evaluated or issued to MDM
    issuedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Timestamp of the last update.
    lastUpdateDateTime *string
    // The OdataType property
    odataType *string
    // When operatingSystemKernelDebugging is enabled, the device is used in development and testing
    operatingSystemKernelDebugging *string
    // The Operating System Revision List that was loaded during initial boot on the attested device
    operatingSystemRevListInfo *string
    // The measurement that is captured in PCR[0]
    pcr0 *string
    // Informational attribute that identifies the HASH algorithm that was used by TPM
    pcrHashAlgorithm *string
    // The number of times a PC device has hibernated or resumed
    resetCount *int64
    // The number of times a PC device has rebooted
    restartCount *int64
    // Safe mode is a troubleshooting option for Windows that starts your computer in a limited state
    safeMode *string
    // When Secure Boot is enabled, the core components must have the correct cryptographic signatures
    secureBoot *string
    // Fingerprint of the Custom Secure Boot Configuration Policy
    secureBootConfigurationPolicyFingerPrint *string
    // When test signing is allowed, the device does not enforce signature validation during boot
    testSigning *string
    // The security version number of the Boot Application
    tpmVersion *string
    // VSM is a container that protects high value assets from a compromised kernel
    virtualSecureMode *string
    // Operating system running with limited services that is used to prepare a computer for Windows
    windowsPE *string
}
// NewDeviceHealthAttestationState instantiates a new deviceHealthAttestationState and sets the default values.
func NewDeviceHealthAttestationState()(*DeviceHealthAttestationState) {
    m := &DeviceHealthAttestationState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceHealthAttestationStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceHealthAttestationStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceHealthAttestationState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthAttestationState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttestationIdentityKey gets the attestationIdentityKey property value. TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate.
func (m *DeviceHealthAttestationState) GetAttestationIdentityKey()(*string) {
    return m.attestationIdentityKey
}
// GetBitLockerStatus gets the bitLockerStatus property value. On or Off of BitLocker Drive Encryption
func (m *DeviceHealthAttestationState) GetBitLockerStatus()(*string) {
    return m.bitLockerStatus
}
// GetBootAppSecurityVersion gets the bootAppSecurityVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) GetBootAppSecurityVersion()(*string) {
    return m.bootAppSecurityVersion
}
// GetBootDebugging gets the bootDebugging property value. When bootDebugging is enabled, the device is used in development and testing
func (m *DeviceHealthAttestationState) GetBootDebugging()(*string) {
    return m.bootDebugging
}
// GetBootManagerSecurityVersion gets the bootManagerSecurityVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) GetBootManagerSecurityVersion()(*string) {
    return m.bootManagerSecurityVersion
}
// GetBootManagerVersion gets the bootManagerVersion property value. The version of the Boot Manager
func (m *DeviceHealthAttestationState) GetBootManagerVersion()(*string) {
    return m.bootManagerVersion
}
// GetBootRevisionListInfo gets the bootRevisionListInfo property value. The Boot Revision List that was loaded during initial boot on the attested device
func (m *DeviceHealthAttestationState) GetBootRevisionListInfo()(*string) {
    return m.bootRevisionListInfo
}
// GetCodeIntegrity gets the codeIntegrity property value. When code integrity is enabled, code execution is restricted to integrity verified code
func (m *DeviceHealthAttestationState) GetCodeIntegrity()(*string) {
    return m.codeIntegrity
}
// GetCodeIntegrityCheckVersion gets the codeIntegrityCheckVersion property value. The version of the Boot Manager
func (m *DeviceHealthAttestationState) GetCodeIntegrityCheckVersion()(*string) {
    return m.codeIntegrityCheckVersion
}
// GetCodeIntegrityPolicy gets the codeIntegrityPolicy property value. The Code Integrity policy that is controlling the security of the boot environment
func (m *DeviceHealthAttestationState) GetCodeIntegrityPolicy()(*string) {
    return m.codeIntegrityPolicy
}
// GetContentNamespaceUrl gets the contentNamespaceUrl property value. The DHA report version. (Namespace version)
func (m *DeviceHealthAttestationState) GetContentNamespaceUrl()(*string) {
    return m.contentNamespaceUrl
}
// GetContentVersion gets the contentVersion property value. The HealthAttestation state schema version
func (m *DeviceHealthAttestationState) GetContentVersion()(*string) {
    return m.contentVersion
}
// GetDataExcutionPolicy gets the dataExcutionPolicy property value. DEP Policy defines a set of hardware and software technologies that perform additional checks on memory
func (m *DeviceHealthAttestationState) GetDataExcutionPolicy()(*string) {
    return m.dataExcutionPolicy
}
// GetDeviceHealthAttestationStatus gets the deviceHealthAttestationStatus property value. The DHA report version. (Namespace version)
func (m *DeviceHealthAttestationState) GetDeviceHealthAttestationStatus()(*string) {
    return m.deviceHealthAttestationStatus
}
// GetEarlyLaunchAntiMalwareDriverProtection gets the earlyLaunchAntiMalwareDriverProtection property value. ELAM provides protection for the computers in your network when they start up
func (m *DeviceHealthAttestationState) GetEarlyLaunchAntiMalwareDriverProtection()(*string) {
    return m.earlyLaunchAntiMalwareDriverProtection
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceHealthAttestationState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attestationIdentityKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttestationIdentityKey(val)
        }
        return nil
    }
    res["bitLockerStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerStatus(val)
        }
        return nil
    }
    res["bootAppSecurityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootAppSecurityVersion(val)
        }
        return nil
    }
    res["bootDebugging"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootDebugging(val)
        }
        return nil
    }
    res["bootManagerSecurityVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootManagerSecurityVersion(val)
        }
        return nil
    }
    res["bootManagerVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootManagerVersion(val)
        }
        return nil
    }
    res["bootRevisionListInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBootRevisionListInfo(val)
        }
        return nil
    }
    res["codeIntegrity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeIntegrity(val)
        }
        return nil
    }
    res["codeIntegrityCheckVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeIntegrityCheckVersion(val)
        }
        return nil
    }
    res["codeIntegrityPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCodeIntegrityPolicy(val)
        }
        return nil
    }
    res["contentNamespaceUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentNamespaceUrl(val)
        }
        return nil
    }
    res["contentVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentVersion(val)
        }
        return nil
    }
    res["dataExcutionPolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataExcutionPolicy(val)
        }
        return nil
    }
    res["deviceHealthAttestationStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceHealthAttestationStatus(val)
        }
        return nil
    }
    res["earlyLaunchAntiMalwareDriverProtection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEarlyLaunchAntiMalwareDriverProtection(val)
        }
        return nil
    }
    res["healthAttestationSupportedStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthAttestationSupportedStatus(val)
        }
        return nil
    }
    res["healthStatusMismatchInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatusMismatchInfo(val)
        }
        return nil
    }
    res["issuedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuedDateTime(val)
        }
        return nil
    }
    res["lastUpdateDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdateDateTime(val)
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
    res["operatingSystemKernelDebugging"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemKernelDebugging(val)
        }
        return nil
    }
    res["operatingSystemRevListInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperatingSystemRevListInfo(val)
        }
        return nil
    }
    res["pcr0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPcr0(val)
        }
        return nil
    }
    res["pcrHashAlgorithm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPcrHashAlgorithm(val)
        }
        return nil
    }
    res["resetCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResetCount(val)
        }
        return nil
    }
    res["restartCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestartCount(val)
        }
        return nil
    }
    res["safeMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSafeMode(val)
        }
        return nil
    }
    res["secureBoot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBoot(val)
        }
        return nil
    }
    res["secureBootConfigurationPolicyFingerPrint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSecureBootConfigurationPolicyFingerPrint(val)
        }
        return nil
    }
    res["testSigning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTestSigning(val)
        }
        return nil
    }
    res["tpmVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTpmVersion(val)
        }
        return nil
    }
    res["virtualSecureMode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVirtualSecureMode(val)
        }
        return nil
    }
    res["windowsPE"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWindowsPE(val)
        }
        return nil
    }
    return res
}
// GetHealthAttestationSupportedStatus gets the healthAttestationSupportedStatus property value. This attribute indicates if DHA is supported for the device
func (m *DeviceHealthAttestationState) GetHealthAttestationSupportedStatus()(*string) {
    return m.healthAttestationSupportedStatus
}
// GetHealthStatusMismatchInfo gets the healthStatusMismatchInfo property value. This attribute appears if DHA-Service detects an integrity issue
func (m *DeviceHealthAttestationState) GetHealthStatusMismatchInfo()(*string) {
    return m.healthStatusMismatchInfo
}
// GetIssuedDateTime gets the issuedDateTime property value. The DateTime when device was evaluated or issued to MDM
func (m *DeviceHealthAttestationState) GetIssuedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.issuedDateTime
}
// GetLastUpdateDateTime gets the lastUpdateDateTime property value. The Timestamp of the last update.
func (m *DeviceHealthAttestationState) GetLastUpdateDateTime()(*string) {
    return m.lastUpdateDateTime
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceHealthAttestationState) GetOdataType()(*string) {
    return m.odataType
}
// GetOperatingSystemKernelDebugging gets the operatingSystemKernelDebugging property value. When operatingSystemKernelDebugging is enabled, the device is used in development and testing
func (m *DeviceHealthAttestationState) GetOperatingSystemKernelDebugging()(*string) {
    return m.operatingSystemKernelDebugging
}
// GetOperatingSystemRevListInfo gets the operatingSystemRevListInfo property value. The Operating System Revision List that was loaded during initial boot on the attested device
func (m *DeviceHealthAttestationState) GetOperatingSystemRevListInfo()(*string) {
    return m.operatingSystemRevListInfo
}
// GetPcr0 gets the pcr0 property value. The measurement that is captured in PCR[0]
func (m *DeviceHealthAttestationState) GetPcr0()(*string) {
    return m.pcr0
}
// GetPcrHashAlgorithm gets the pcrHashAlgorithm property value. Informational attribute that identifies the HASH algorithm that was used by TPM
func (m *DeviceHealthAttestationState) GetPcrHashAlgorithm()(*string) {
    return m.pcrHashAlgorithm
}
// GetResetCount gets the resetCount property value. The number of times a PC device has hibernated or resumed
func (m *DeviceHealthAttestationState) GetResetCount()(*int64) {
    return m.resetCount
}
// GetRestartCount gets the restartCount property value. The number of times a PC device has rebooted
func (m *DeviceHealthAttestationState) GetRestartCount()(*int64) {
    return m.restartCount
}
// GetSafeMode gets the safeMode property value. Safe mode is a troubleshooting option for Windows that starts your computer in a limited state
func (m *DeviceHealthAttestationState) GetSafeMode()(*string) {
    return m.safeMode
}
// GetSecureBoot gets the secureBoot property value. When Secure Boot is enabled, the core components must have the correct cryptographic signatures
func (m *DeviceHealthAttestationState) GetSecureBoot()(*string) {
    return m.secureBoot
}
// GetSecureBootConfigurationPolicyFingerPrint gets the secureBootConfigurationPolicyFingerPrint property value. Fingerprint of the Custom Secure Boot Configuration Policy
func (m *DeviceHealthAttestationState) GetSecureBootConfigurationPolicyFingerPrint()(*string) {
    return m.secureBootConfigurationPolicyFingerPrint
}
// GetTestSigning gets the testSigning property value. When test signing is allowed, the device does not enforce signature validation during boot
func (m *DeviceHealthAttestationState) GetTestSigning()(*string) {
    return m.testSigning
}
// GetTpmVersion gets the tpmVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) GetTpmVersion()(*string) {
    return m.tpmVersion
}
// GetVirtualSecureMode gets the virtualSecureMode property value. VSM is a container that protects high value assets from a compromised kernel
func (m *DeviceHealthAttestationState) GetVirtualSecureMode()(*string) {
    return m.virtualSecureMode
}
// GetWindowsPE gets the windowsPE property value. Operating system running with limited services that is used to prepare a computer for Windows
func (m *DeviceHealthAttestationState) GetWindowsPE()(*string) {
    return m.windowsPE
}
// Serialize serializes information the current object
func (m *DeviceHealthAttestationState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("attestationIdentityKey", m.GetAttestationIdentityKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bitLockerStatus", m.GetBitLockerStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootAppSecurityVersion", m.GetBootAppSecurityVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootDebugging", m.GetBootDebugging())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootManagerSecurityVersion", m.GetBootManagerSecurityVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootManagerVersion", m.GetBootManagerVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootRevisionListInfo", m.GetBootRevisionListInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeIntegrity", m.GetCodeIntegrity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeIntegrityCheckVersion", m.GetCodeIntegrityCheckVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeIntegrityPolicy", m.GetCodeIntegrityPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentNamespaceUrl", m.GetContentNamespaceUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentVersion", m.GetContentVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dataExcutionPolicy", m.GetDataExcutionPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceHealthAttestationStatus", m.GetDeviceHealthAttestationStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("earlyLaunchAntiMalwareDriverProtection", m.GetEarlyLaunchAntiMalwareDriverProtection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("healthAttestationSupportedStatus", m.GetHealthAttestationSupportedStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("healthStatusMismatchInfo", m.GetHealthStatusMismatchInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("issuedDateTime", m.GetIssuedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
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
        err := writer.WriteStringValue("operatingSystemKernelDebugging", m.GetOperatingSystemKernelDebugging())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemRevListInfo", m.GetOperatingSystemRevListInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pcr0", m.GetPcr0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pcrHashAlgorithm", m.GetPcrHashAlgorithm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("resetCount", m.GetResetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("restartCount", m.GetRestartCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("safeMode", m.GetSafeMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secureBoot", m.GetSecureBoot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secureBootConfigurationPolicyFingerPrint", m.GetSecureBootConfigurationPolicyFingerPrint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("testSigning", m.GetTestSigning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmVersion", m.GetTpmVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("virtualSecureMode", m.GetVirtualSecureMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("windowsPE", m.GetWindowsPE())
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
func (m *DeviceHealthAttestationState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttestationIdentityKey sets the attestationIdentityKey property value. TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate.
func (m *DeviceHealthAttestationState) SetAttestationIdentityKey(value *string)() {
    m.attestationIdentityKey = value
}
// SetBitLockerStatus sets the bitLockerStatus property value. On or Off of BitLocker Drive Encryption
func (m *DeviceHealthAttestationState) SetBitLockerStatus(value *string)() {
    m.bitLockerStatus = value
}
// SetBootAppSecurityVersion sets the bootAppSecurityVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) SetBootAppSecurityVersion(value *string)() {
    m.bootAppSecurityVersion = value
}
// SetBootDebugging sets the bootDebugging property value. When bootDebugging is enabled, the device is used in development and testing
func (m *DeviceHealthAttestationState) SetBootDebugging(value *string)() {
    m.bootDebugging = value
}
// SetBootManagerSecurityVersion sets the bootManagerSecurityVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) SetBootManagerSecurityVersion(value *string)() {
    m.bootManagerSecurityVersion = value
}
// SetBootManagerVersion sets the bootManagerVersion property value. The version of the Boot Manager
func (m *DeviceHealthAttestationState) SetBootManagerVersion(value *string)() {
    m.bootManagerVersion = value
}
// SetBootRevisionListInfo sets the bootRevisionListInfo property value. The Boot Revision List that was loaded during initial boot on the attested device
func (m *DeviceHealthAttestationState) SetBootRevisionListInfo(value *string)() {
    m.bootRevisionListInfo = value
}
// SetCodeIntegrity sets the codeIntegrity property value. When code integrity is enabled, code execution is restricted to integrity verified code
func (m *DeviceHealthAttestationState) SetCodeIntegrity(value *string)() {
    m.codeIntegrity = value
}
// SetCodeIntegrityCheckVersion sets the codeIntegrityCheckVersion property value. The version of the Boot Manager
func (m *DeviceHealthAttestationState) SetCodeIntegrityCheckVersion(value *string)() {
    m.codeIntegrityCheckVersion = value
}
// SetCodeIntegrityPolicy sets the codeIntegrityPolicy property value. The Code Integrity policy that is controlling the security of the boot environment
func (m *DeviceHealthAttestationState) SetCodeIntegrityPolicy(value *string)() {
    m.codeIntegrityPolicy = value
}
// SetContentNamespaceUrl sets the contentNamespaceUrl property value. The DHA report version. (Namespace version)
func (m *DeviceHealthAttestationState) SetContentNamespaceUrl(value *string)() {
    m.contentNamespaceUrl = value
}
// SetContentVersion sets the contentVersion property value. The HealthAttestation state schema version
func (m *DeviceHealthAttestationState) SetContentVersion(value *string)() {
    m.contentVersion = value
}
// SetDataExcutionPolicy sets the dataExcutionPolicy property value. DEP Policy defines a set of hardware and software technologies that perform additional checks on memory
func (m *DeviceHealthAttestationState) SetDataExcutionPolicy(value *string)() {
    m.dataExcutionPolicy = value
}
// SetDeviceHealthAttestationStatus sets the deviceHealthAttestationStatus property value. The DHA report version. (Namespace version)
func (m *DeviceHealthAttestationState) SetDeviceHealthAttestationStatus(value *string)() {
    m.deviceHealthAttestationStatus = value
}
// SetEarlyLaunchAntiMalwareDriverProtection sets the earlyLaunchAntiMalwareDriverProtection property value. ELAM provides protection for the computers in your network when they start up
func (m *DeviceHealthAttestationState) SetEarlyLaunchAntiMalwareDriverProtection(value *string)() {
    m.earlyLaunchAntiMalwareDriverProtection = value
}
// SetHealthAttestationSupportedStatus sets the healthAttestationSupportedStatus property value. This attribute indicates if DHA is supported for the device
func (m *DeviceHealthAttestationState) SetHealthAttestationSupportedStatus(value *string)() {
    m.healthAttestationSupportedStatus = value
}
// SetHealthStatusMismatchInfo sets the healthStatusMismatchInfo property value. This attribute appears if DHA-Service detects an integrity issue
func (m *DeviceHealthAttestationState) SetHealthStatusMismatchInfo(value *string)() {
    m.healthStatusMismatchInfo = value
}
// SetIssuedDateTime sets the issuedDateTime property value. The DateTime when device was evaluated or issued to MDM
func (m *DeviceHealthAttestationState) SetIssuedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.issuedDateTime = value
}
// SetLastUpdateDateTime sets the lastUpdateDateTime property value. The Timestamp of the last update.
func (m *DeviceHealthAttestationState) SetLastUpdateDateTime(value *string)() {
    m.lastUpdateDateTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceHealthAttestationState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperatingSystemKernelDebugging sets the operatingSystemKernelDebugging property value. When operatingSystemKernelDebugging is enabled, the device is used in development and testing
func (m *DeviceHealthAttestationState) SetOperatingSystemKernelDebugging(value *string)() {
    m.operatingSystemKernelDebugging = value
}
// SetOperatingSystemRevListInfo sets the operatingSystemRevListInfo property value. The Operating System Revision List that was loaded during initial boot on the attested device
func (m *DeviceHealthAttestationState) SetOperatingSystemRevListInfo(value *string)() {
    m.operatingSystemRevListInfo = value
}
// SetPcr0 sets the pcr0 property value. The measurement that is captured in PCR[0]
func (m *DeviceHealthAttestationState) SetPcr0(value *string)() {
    m.pcr0 = value
}
// SetPcrHashAlgorithm sets the pcrHashAlgorithm property value. Informational attribute that identifies the HASH algorithm that was used by TPM
func (m *DeviceHealthAttestationState) SetPcrHashAlgorithm(value *string)() {
    m.pcrHashAlgorithm = value
}
// SetResetCount sets the resetCount property value. The number of times a PC device has hibernated or resumed
func (m *DeviceHealthAttestationState) SetResetCount(value *int64)() {
    m.resetCount = value
}
// SetRestartCount sets the restartCount property value. The number of times a PC device has rebooted
func (m *DeviceHealthAttestationState) SetRestartCount(value *int64)() {
    m.restartCount = value
}
// SetSafeMode sets the safeMode property value. Safe mode is a troubleshooting option for Windows that starts your computer in a limited state
func (m *DeviceHealthAttestationState) SetSafeMode(value *string)() {
    m.safeMode = value
}
// SetSecureBoot sets the secureBoot property value. When Secure Boot is enabled, the core components must have the correct cryptographic signatures
func (m *DeviceHealthAttestationState) SetSecureBoot(value *string)() {
    m.secureBoot = value
}
// SetSecureBootConfigurationPolicyFingerPrint sets the secureBootConfigurationPolicyFingerPrint property value. Fingerprint of the Custom Secure Boot Configuration Policy
func (m *DeviceHealthAttestationState) SetSecureBootConfigurationPolicyFingerPrint(value *string)() {
    m.secureBootConfigurationPolicyFingerPrint = value
}
// SetTestSigning sets the testSigning property value. When test signing is allowed, the device does not enforce signature validation during boot
func (m *DeviceHealthAttestationState) SetTestSigning(value *string)() {
    m.testSigning = value
}
// SetTpmVersion sets the tpmVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) SetTpmVersion(value *string)() {
    m.tpmVersion = value
}
// SetVirtualSecureMode sets the virtualSecureMode property value. VSM is a container that protects high value assets from a compromised kernel
func (m *DeviceHealthAttestationState) SetVirtualSecureMode(value *string)() {
    m.virtualSecureMode = value
}
// SetWindowsPE sets the windowsPE property value. Operating system running with limited services that is used to prepare a computer for Windows
func (m *DeviceHealthAttestationState) SetWindowsPE(value *string)() {
    m.windowsPE = value
}
