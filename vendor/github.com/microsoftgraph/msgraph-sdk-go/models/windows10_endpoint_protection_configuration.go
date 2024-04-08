package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Windows10EndpointProtectionConfiguration 
type Windows10EndpointProtectionConfiguration struct {
    DeviceConfiguration
    // Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
    applicationGuardAllowPersistence *bool
    // Allow printing to Local Printers from Container
    applicationGuardAllowPrintToLocalPrinters *bool
    // Allow printing to Network Printers from Container
    applicationGuardAllowPrintToNetworkPrinters *bool
    // Allow printing to PDF from Container
    applicationGuardAllowPrintToPDF *bool
    // Allow printing to XPS from Container
    applicationGuardAllowPrintToXPS *bool
    // Possible values for applicationGuardBlockClipboardSharingType
    applicationGuardBlockClipboardSharing *ApplicationGuardBlockClipboardSharingType
    // Possible values for applicationGuardBlockFileTransfer
    applicationGuardBlockFileTransfer *ApplicationGuardBlockFileTransferType
    // Block enterprise sites to load non-enterprise content, such as third party plug-ins
    applicationGuardBlockNonEnterpriseContent *bool
    // Enable Windows Defender Application Guard
    applicationGuardEnabled *bool
    // Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
    applicationGuardForceAuditing *bool
    // Possible values of AppLocker Application Control Types
    appLockerApplicationControl *AppLockerApplicationControlType
    // Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
    bitLockerDisableWarningForOtherDiskEncryption *bool
    // Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
    bitLockerEnableStorageCardEncryptionOnMobile *bool
    // Allows the admin to require encryption to be turned on using BitLocker.
    bitLockerEncryptDevice *bool
    // BitLocker Removable Drive Policy.
    bitLockerRemovableDrivePolicy BitLockerRemovableDrivePolicyable
    // List of folder paths to be added to the list of protected folders
    defenderAdditionalGuardedFolders []string
    // List of exe files and folders to be excluded from attack surface reduction rules
    defenderAttackSurfaceReductionExcludedPaths []string
    // Xml content containing information regarding exploit protection details.
    defenderExploitProtectionXml []byte
    // Name of the file from which DefenderExploitProtectionXml was obtained.
    defenderExploitProtectionXmlFileName *string
    // List of paths to exe that are allowed to access protected folders
    defenderGuardedFoldersAllowedAppPaths []string
    // Indicates whether or not to block user from overriding Exploit Protection settings.
    defenderSecurityCenterBlockExploitProtectionOverride *bool
    // Blocks stateful FTP connections to the device
    firewallBlockStatefulFTP *bool
    // Possible values for firewallCertificateRevocationListCheckMethod
    firewallCertificateRevocationListCheckMethod *FirewallCertificateRevocationListCheckMethodType
    // Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
    firewallIdleTimeoutForSecurityAssociationInSeconds *int32
    // Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
    firewallIPSecExemptionsAllowDHCP *bool
    // Configures IPSec exemptions to allow ICMP
    firewallIPSecExemptionsAllowICMP *bool
    // Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
    firewallIPSecExemptionsAllowNeighborDiscovery *bool
    // Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
    firewallIPSecExemptionsAllowRouterDiscovery *bool
    // If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
    firewallMergeKeyingModuleSettings *bool
    // Possible values for firewallPacketQueueingMethod
    firewallPacketQueueingMethod *FirewallPacketQueueingMethodType
    // Possible values for firewallPreSharedKeyEncodingMethod
    firewallPreSharedKeyEncodingMethod *FirewallPreSharedKeyEncodingMethodType
    // Configures the firewall profile settings for domain networks
    firewallProfileDomain WindowsFirewallNetworkProfileable
    // Configures the firewall profile settings for private networks
    firewallProfilePrivate WindowsFirewallNetworkProfileable
    // Configures the firewall profile settings for public networks
    firewallProfilePublic WindowsFirewallNetworkProfileable
    // Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
    smartScreenBlockOverrideForFiles *bool
    // Allows IT Admins to configure SmartScreen for Windows.
    smartScreenEnableInShell *bool
}
// NewWindows10EndpointProtectionConfiguration instantiates a new Windows10EndpointProtectionConfiguration and sets the default values.
func NewWindows10EndpointProtectionConfiguration()(*Windows10EndpointProtectionConfiguration) {
    m := &Windows10EndpointProtectionConfiguration{
        DeviceConfiguration: *NewDeviceConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.windows10EndpointProtectionConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWindows10EndpointProtectionConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindows10EndpointProtectionConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindows10EndpointProtectionConfiguration(), nil
}
// GetApplicationGuardAllowPersistence gets the applicationGuardAllowPersistence property value. Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPersistence()(*bool) {
    return m.applicationGuardAllowPersistence
}
// GetApplicationGuardAllowPrintToLocalPrinters gets the applicationGuardAllowPrintToLocalPrinters property value. Allow printing to Local Printers from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToLocalPrinters()(*bool) {
    return m.applicationGuardAllowPrintToLocalPrinters
}
// GetApplicationGuardAllowPrintToNetworkPrinters gets the applicationGuardAllowPrintToNetworkPrinters property value. Allow printing to Network Printers from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToNetworkPrinters()(*bool) {
    return m.applicationGuardAllowPrintToNetworkPrinters
}
// GetApplicationGuardAllowPrintToPDF gets the applicationGuardAllowPrintToPDF property value. Allow printing to PDF from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToPDF()(*bool) {
    return m.applicationGuardAllowPrintToPDF
}
// GetApplicationGuardAllowPrintToXPS gets the applicationGuardAllowPrintToXPS property value. Allow printing to XPS from Container
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardAllowPrintToXPS()(*bool) {
    return m.applicationGuardAllowPrintToXPS
}
// GetApplicationGuardBlockClipboardSharing gets the applicationGuardBlockClipboardSharing property value. Possible values for applicationGuardBlockClipboardSharingType
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockClipboardSharing()(*ApplicationGuardBlockClipboardSharingType) {
    return m.applicationGuardBlockClipboardSharing
}
// GetApplicationGuardBlockFileTransfer gets the applicationGuardBlockFileTransfer property value. Possible values for applicationGuardBlockFileTransfer
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockFileTransfer()(*ApplicationGuardBlockFileTransferType) {
    return m.applicationGuardBlockFileTransfer
}
// GetApplicationGuardBlockNonEnterpriseContent gets the applicationGuardBlockNonEnterpriseContent property value. Block enterprise sites to load non-enterprise content, such as third party plug-ins
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardBlockNonEnterpriseContent()(*bool) {
    return m.applicationGuardBlockNonEnterpriseContent
}
// GetApplicationGuardEnabled gets the applicationGuardEnabled property value. Enable Windows Defender Application Guard
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardEnabled()(*bool) {
    return m.applicationGuardEnabled
}
// GetApplicationGuardForceAuditing gets the applicationGuardForceAuditing property value. Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
func (m *Windows10EndpointProtectionConfiguration) GetApplicationGuardForceAuditing()(*bool) {
    return m.applicationGuardForceAuditing
}
// GetAppLockerApplicationControl gets the appLockerApplicationControl property value. Possible values of AppLocker Application Control Types
func (m *Windows10EndpointProtectionConfiguration) GetAppLockerApplicationControl()(*AppLockerApplicationControlType) {
    return m.appLockerApplicationControl
}
// GetBitLockerDisableWarningForOtherDiskEncryption gets the bitLockerDisableWarningForOtherDiskEncryption property value. Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerDisableWarningForOtherDiskEncryption()(*bool) {
    return m.bitLockerDisableWarningForOtherDiskEncryption
}
// GetBitLockerEnableStorageCardEncryptionOnMobile gets the bitLockerEnableStorageCardEncryptionOnMobile property value. Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerEnableStorageCardEncryptionOnMobile()(*bool) {
    return m.bitLockerEnableStorageCardEncryptionOnMobile
}
// GetBitLockerEncryptDevice gets the bitLockerEncryptDevice property value. Allows the admin to require encryption to be turned on using BitLocker.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerEncryptDevice()(*bool) {
    return m.bitLockerEncryptDevice
}
// GetBitLockerRemovableDrivePolicy gets the bitLockerRemovableDrivePolicy property value. BitLocker Removable Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) GetBitLockerRemovableDrivePolicy()(BitLockerRemovableDrivePolicyable) {
    return m.bitLockerRemovableDrivePolicy
}
// GetDefenderAdditionalGuardedFolders gets the defenderAdditionalGuardedFolders property value. List of folder paths to be added to the list of protected folders
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAdditionalGuardedFolders()([]string) {
    return m.defenderAdditionalGuardedFolders
}
// GetDefenderAttackSurfaceReductionExcludedPaths gets the defenderAttackSurfaceReductionExcludedPaths property value. List of exe files and folders to be excluded from attack surface reduction rules
func (m *Windows10EndpointProtectionConfiguration) GetDefenderAttackSurfaceReductionExcludedPaths()([]string) {
    return m.defenderAttackSurfaceReductionExcludedPaths
}
// GetDefenderExploitProtectionXml gets the defenderExploitProtectionXml property value. Xml content containing information regarding exploit protection details.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXml()([]byte) {
    return m.defenderExploitProtectionXml
}
// GetDefenderExploitProtectionXmlFileName gets the defenderExploitProtectionXmlFileName property value. Name of the file from which DefenderExploitProtectionXml was obtained.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderExploitProtectionXmlFileName()(*string) {
    return m.defenderExploitProtectionXmlFileName
}
// GetDefenderGuardedFoldersAllowedAppPaths gets the defenderGuardedFoldersAllowedAppPaths property value. List of paths to exe that are allowed to access protected folders
func (m *Windows10EndpointProtectionConfiguration) GetDefenderGuardedFoldersAllowedAppPaths()([]string) {
    return m.defenderGuardedFoldersAllowedAppPaths
}
// GetDefenderSecurityCenterBlockExploitProtectionOverride gets the defenderSecurityCenterBlockExploitProtectionOverride property value. Indicates whether or not to block user from overriding Exploit Protection settings.
func (m *Windows10EndpointProtectionConfiguration) GetDefenderSecurityCenterBlockExploitProtectionOverride()(*bool) {
    return m.defenderSecurityCenterBlockExploitProtectionOverride
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Windows10EndpointProtectionConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DeviceConfiguration.GetFieldDeserializers()
    res["applicationGuardAllowPersistence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPersistence(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToLocalPrinters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToLocalPrinters(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToNetworkPrinters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToNetworkPrinters(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToPDF"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToPDF(val)
        }
        return nil
    }
    res["applicationGuardAllowPrintToXPS"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardAllowPrintToXPS(val)
        }
        return nil
    }
    res["applicationGuardBlockClipboardSharing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationGuardBlockClipboardSharingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockClipboardSharing(val.(*ApplicationGuardBlockClipboardSharingType))
        }
        return nil
    }
    res["applicationGuardBlockFileTransfer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseApplicationGuardBlockFileTransferType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockFileTransfer(val.(*ApplicationGuardBlockFileTransferType))
        }
        return nil
    }
    res["applicationGuardBlockNonEnterpriseContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardBlockNonEnterpriseContent(val)
        }
        return nil
    }
    res["applicationGuardEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardEnabled(val)
        }
        return nil
    }
    res["applicationGuardForceAuditing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationGuardForceAuditing(val)
        }
        return nil
    }
    res["appLockerApplicationControl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppLockerApplicationControlType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppLockerApplicationControl(val.(*AppLockerApplicationControlType))
        }
        return nil
    }
    res["bitLockerDisableWarningForOtherDiskEncryption"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerDisableWarningForOtherDiskEncryption(val)
        }
        return nil
    }
    res["bitLockerEnableStorageCardEncryptionOnMobile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerEnableStorageCardEncryptionOnMobile(val)
        }
        return nil
    }
    res["bitLockerEncryptDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerEncryptDevice(val)
        }
        return nil
    }
    res["bitLockerRemovableDrivePolicy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBitLockerRemovableDrivePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitLockerRemovableDrivePolicy(val.(BitLockerRemovableDrivePolicyable))
        }
        return nil
    }
    res["defenderAdditionalGuardedFolders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderAdditionalGuardedFolders(res)
        }
        return nil
    }
    res["defenderAttackSurfaceReductionExcludedPaths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderAttackSurfaceReductionExcludedPaths(res)
        }
        return nil
    }
    res["defenderExploitProtectionXml"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderExploitProtectionXml(val)
        }
        return nil
    }
    res["defenderExploitProtectionXmlFileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderExploitProtectionXmlFileName(val)
        }
        return nil
    }
    res["defenderGuardedFoldersAllowedAppPaths"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefenderGuardedFoldersAllowedAppPaths(res)
        }
        return nil
    }
    res["defenderSecurityCenterBlockExploitProtectionOverride"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderSecurityCenterBlockExploitProtectionOverride(val)
        }
        return nil
    }
    res["firewallBlockStatefulFTP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallBlockStatefulFTP(val)
        }
        return nil
    }
    res["firewallCertificateRevocationListCheckMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallCertificateRevocationListCheckMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallCertificateRevocationListCheckMethod(val.(*FirewallCertificateRevocationListCheckMethodType))
        }
        return nil
    }
    res["firewallIdleTimeoutForSecurityAssociationInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIdleTimeoutForSecurityAssociationInSeconds(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowDHCP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowDHCP(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowICMP"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowICMP(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowNeighborDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowNeighborDiscovery(val)
        }
        return nil
    }
    res["firewallIPSecExemptionsAllowRouterDiscovery"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallIPSecExemptionsAllowRouterDiscovery(val)
        }
        return nil
    }
    res["firewallMergeKeyingModuleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallMergeKeyingModuleSettings(val)
        }
        return nil
    }
    res["firewallPacketQueueingMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallPacketQueueingMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallPacketQueueingMethod(val.(*FirewallPacketQueueingMethodType))
        }
        return nil
    }
    res["firewallPreSharedKeyEncodingMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFirewallPreSharedKeyEncodingMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallPreSharedKeyEncodingMethod(val.(*FirewallPreSharedKeyEncodingMethodType))
        }
        return nil
    }
    res["firewallProfileDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfileDomain(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["firewallProfilePrivate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfilePrivate(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["firewallProfilePublic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWindowsFirewallNetworkProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirewallProfilePublic(val.(WindowsFirewallNetworkProfileable))
        }
        return nil
    }
    res["smartScreenBlockOverrideForFiles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenBlockOverrideForFiles(val)
        }
        return nil
    }
    res["smartScreenEnableInShell"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSmartScreenEnableInShell(val)
        }
        return nil
    }
    return res
}
// GetFirewallBlockStatefulFTP gets the firewallBlockStatefulFTP property value. Blocks stateful FTP connections to the device
func (m *Windows10EndpointProtectionConfiguration) GetFirewallBlockStatefulFTP()(*bool) {
    return m.firewallBlockStatefulFTP
}
// GetFirewallCertificateRevocationListCheckMethod gets the firewallCertificateRevocationListCheckMethod property value. Possible values for firewallCertificateRevocationListCheckMethod
func (m *Windows10EndpointProtectionConfiguration) GetFirewallCertificateRevocationListCheckMethod()(*FirewallCertificateRevocationListCheckMethodType) {
    return m.firewallCertificateRevocationListCheckMethod
}
// GetFirewallIdleTimeoutForSecurityAssociationInSeconds gets the firewallIdleTimeoutForSecurityAssociationInSeconds property value. Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIdleTimeoutForSecurityAssociationInSeconds()(*int32) {
    return m.firewallIdleTimeoutForSecurityAssociationInSeconds
}
// GetFirewallIPSecExemptionsAllowDHCP gets the firewallIPSecExemptionsAllowDHCP property value. Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowDHCP()(*bool) {
    return m.firewallIPSecExemptionsAllowDHCP
}
// GetFirewallIPSecExemptionsAllowICMP gets the firewallIPSecExemptionsAllowICMP property value. Configures IPSec exemptions to allow ICMP
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowICMP()(*bool) {
    return m.firewallIPSecExemptionsAllowICMP
}
// GetFirewallIPSecExemptionsAllowNeighborDiscovery gets the firewallIPSecExemptionsAllowNeighborDiscovery property value. Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowNeighborDiscovery()(*bool) {
    return m.firewallIPSecExemptionsAllowNeighborDiscovery
}
// GetFirewallIPSecExemptionsAllowRouterDiscovery gets the firewallIPSecExemptionsAllowRouterDiscovery property value. Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) GetFirewallIPSecExemptionsAllowRouterDiscovery()(*bool) {
    return m.firewallIPSecExemptionsAllowRouterDiscovery
}
// GetFirewallMergeKeyingModuleSettings gets the firewallMergeKeyingModuleSettings property value. If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
func (m *Windows10EndpointProtectionConfiguration) GetFirewallMergeKeyingModuleSettings()(*bool) {
    return m.firewallMergeKeyingModuleSettings
}
// GetFirewallPacketQueueingMethod gets the firewallPacketQueueingMethod property value. Possible values for firewallPacketQueueingMethod
func (m *Windows10EndpointProtectionConfiguration) GetFirewallPacketQueueingMethod()(*FirewallPacketQueueingMethodType) {
    return m.firewallPacketQueueingMethod
}
// GetFirewallPreSharedKeyEncodingMethod gets the firewallPreSharedKeyEncodingMethod property value. Possible values for firewallPreSharedKeyEncodingMethod
func (m *Windows10EndpointProtectionConfiguration) GetFirewallPreSharedKeyEncodingMethod()(*FirewallPreSharedKeyEncodingMethodType) {
    return m.firewallPreSharedKeyEncodingMethod
}
// GetFirewallProfileDomain gets the firewallProfileDomain property value. Configures the firewall profile settings for domain networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfileDomain()(WindowsFirewallNetworkProfileable) {
    return m.firewallProfileDomain
}
// GetFirewallProfilePrivate gets the firewallProfilePrivate property value. Configures the firewall profile settings for private networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfilePrivate()(WindowsFirewallNetworkProfileable) {
    return m.firewallProfilePrivate
}
// GetFirewallProfilePublic gets the firewallProfilePublic property value. Configures the firewall profile settings for public networks
func (m *Windows10EndpointProtectionConfiguration) GetFirewallProfilePublic()(WindowsFirewallNetworkProfileable) {
    return m.firewallProfilePublic
}
// GetSmartScreenBlockOverrideForFiles gets the smartScreenBlockOverrideForFiles property value. Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
func (m *Windows10EndpointProtectionConfiguration) GetSmartScreenBlockOverrideForFiles()(*bool) {
    return m.smartScreenBlockOverrideForFiles
}
// GetSmartScreenEnableInShell gets the smartScreenEnableInShell property value. Allows IT Admins to configure SmartScreen for Windows.
func (m *Windows10EndpointProtectionConfiguration) GetSmartScreenEnableInShell()(*bool) {
    return m.smartScreenEnableInShell
}
// Serialize serializes information the current object
func (m *Windows10EndpointProtectionConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DeviceConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPersistence", m.GetApplicationGuardAllowPersistence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToLocalPrinters", m.GetApplicationGuardAllowPrintToLocalPrinters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToNetworkPrinters", m.GetApplicationGuardAllowPrintToNetworkPrinters())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToPDF", m.GetApplicationGuardAllowPrintToPDF())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardAllowPrintToXPS", m.GetApplicationGuardAllowPrintToXPS())
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardBlockClipboardSharing() != nil {
        cast := (*m.GetApplicationGuardBlockClipboardSharing()).String()
        err = writer.WriteStringValue("applicationGuardBlockClipboardSharing", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApplicationGuardBlockFileTransfer() != nil {
        cast := (*m.GetApplicationGuardBlockFileTransfer()).String()
        err = writer.WriteStringValue("applicationGuardBlockFileTransfer", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardBlockNonEnterpriseContent", m.GetApplicationGuardBlockNonEnterpriseContent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardEnabled", m.GetApplicationGuardEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("applicationGuardForceAuditing", m.GetApplicationGuardForceAuditing())
        if err != nil {
            return err
        }
    }
    if m.GetAppLockerApplicationControl() != nil {
        cast := (*m.GetAppLockerApplicationControl()).String()
        err = writer.WriteStringValue("appLockerApplicationControl", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerDisableWarningForOtherDiskEncryption", m.GetBitLockerDisableWarningForOtherDiskEncryption())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerEnableStorageCardEncryptionOnMobile", m.GetBitLockerEnableStorageCardEncryptionOnMobile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("bitLockerEncryptDevice", m.GetBitLockerEncryptDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("bitLockerRemovableDrivePolicy", m.GetBitLockerRemovableDrivePolicy())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAdditionalGuardedFolders() != nil {
        err = writer.WriteCollectionOfStringValues("defenderAdditionalGuardedFolders", m.GetDefenderAdditionalGuardedFolders())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAttackSurfaceReductionExcludedPaths() != nil {
        err = writer.WriteCollectionOfStringValues("defenderAttackSurfaceReductionExcludedPaths", m.GetDefenderAttackSurfaceReductionExcludedPaths())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("defenderExploitProtectionXml", m.GetDefenderExploitProtectionXml())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defenderExploitProtectionXmlFileName", m.GetDefenderExploitProtectionXmlFileName())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderGuardedFoldersAllowedAppPaths() != nil {
        err = writer.WriteCollectionOfStringValues("defenderGuardedFoldersAllowedAppPaths", m.GetDefenderGuardedFoldersAllowedAppPaths())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defenderSecurityCenterBlockExploitProtectionOverride", m.GetDefenderSecurityCenterBlockExploitProtectionOverride())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallBlockStatefulFTP", m.GetFirewallBlockStatefulFTP())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallCertificateRevocationListCheckMethod() != nil {
        cast := (*m.GetFirewallCertificateRevocationListCheckMethod()).String()
        err = writer.WriteStringValue("firewallCertificateRevocationListCheckMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("firewallIdleTimeoutForSecurityAssociationInSeconds", m.GetFirewallIdleTimeoutForSecurityAssociationInSeconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowDHCP", m.GetFirewallIPSecExemptionsAllowDHCP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowICMP", m.GetFirewallIPSecExemptionsAllowICMP())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowNeighborDiscovery", m.GetFirewallIPSecExemptionsAllowNeighborDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallIPSecExemptionsAllowRouterDiscovery", m.GetFirewallIPSecExemptionsAllowRouterDiscovery())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("firewallMergeKeyingModuleSettings", m.GetFirewallMergeKeyingModuleSettings())
        if err != nil {
            return err
        }
    }
    if m.GetFirewallPacketQueueingMethod() != nil {
        cast := (*m.GetFirewallPacketQueueingMethod()).String()
        err = writer.WriteStringValue("firewallPacketQueueingMethod", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetFirewallPreSharedKeyEncodingMethod() != nil {
        cast := (*m.GetFirewallPreSharedKeyEncodingMethod()).String()
        err = writer.WriteStringValue("firewallPreSharedKeyEncodingMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfileDomain", m.GetFirewallProfileDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfilePrivate", m.GetFirewallProfilePrivate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("firewallProfilePublic", m.GetFirewallProfilePublic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenBlockOverrideForFiles", m.GetSmartScreenBlockOverrideForFiles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("smartScreenEnableInShell", m.GetSmartScreenEnableInShell())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplicationGuardAllowPersistence sets the applicationGuardAllowPersistence property value. Allow persisting user generated data inside the App Guard Containter (favorites, cookies, web passwords, etc.)
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPersistence(value *bool)() {
    m.applicationGuardAllowPersistence = value
}
// SetApplicationGuardAllowPrintToLocalPrinters sets the applicationGuardAllowPrintToLocalPrinters property value. Allow printing to Local Printers from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToLocalPrinters(value *bool)() {
    m.applicationGuardAllowPrintToLocalPrinters = value
}
// SetApplicationGuardAllowPrintToNetworkPrinters sets the applicationGuardAllowPrintToNetworkPrinters property value. Allow printing to Network Printers from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToNetworkPrinters(value *bool)() {
    m.applicationGuardAllowPrintToNetworkPrinters = value
}
// SetApplicationGuardAllowPrintToPDF sets the applicationGuardAllowPrintToPDF property value. Allow printing to PDF from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToPDF(value *bool)() {
    m.applicationGuardAllowPrintToPDF = value
}
// SetApplicationGuardAllowPrintToXPS sets the applicationGuardAllowPrintToXPS property value. Allow printing to XPS from Container
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardAllowPrintToXPS(value *bool)() {
    m.applicationGuardAllowPrintToXPS = value
}
// SetApplicationGuardBlockClipboardSharing sets the applicationGuardBlockClipboardSharing property value. Possible values for applicationGuardBlockClipboardSharingType
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockClipboardSharing(value *ApplicationGuardBlockClipboardSharingType)() {
    m.applicationGuardBlockClipboardSharing = value
}
// SetApplicationGuardBlockFileTransfer sets the applicationGuardBlockFileTransfer property value. Possible values for applicationGuardBlockFileTransfer
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockFileTransfer(value *ApplicationGuardBlockFileTransferType)() {
    m.applicationGuardBlockFileTransfer = value
}
// SetApplicationGuardBlockNonEnterpriseContent sets the applicationGuardBlockNonEnterpriseContent property value. Block enterprise sites to load non-enterprise content, such as third party plug-ins
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardBlockNonEnterpriseContent(value *bool)() {
    m.applicationGuardBlockNonEnterpriseContent = value
}
// SetApplicationGuardEnabled sets the applicationGuardEnabled property value. Enable Windows Defender Application Guard
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardEnabled(value *bool)() {
    m.applicationGuardEnabled = value
}
// SetApplicationGuardForceAuditing sets the applicationGuardForceAuditing property value. Force auditing will persist Windows logs and events to meet security/compliance criteria (sample events are user login-logoff, use of privilege rights, software installation, system changes, etc.)
func (m *Windows10EndpointProtectionConfiguration) SetApplicationGuardForceAuditing(value *bool)() {
    m.applicationGuardForceAuditing = value
}
// SetAppLockerApplicationControl sets the appLockerApplicationControl property value. Possible values of AppLocker Application Control Types
func (m *Windows10EndpointProtectionConfiguration) SetAppLockerApplicationControl(value *AppLockerApplicationControlType)() {
    m.appLockerApplicationControl = value
}
// SetBitLockerDisableWarningForOtherDiskEncryption sets the bitLockerDisableWarningForOtherDiskEncryption property value. Allows the Admin to disable the warning prompt for other disk encryption on the user machines.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerDisableWarningForOtherDiskEncryption(value *bool)() {
    m.bitLockerDisableWarningForOtherDiskEncryption = value
}
// SetBitLockerEnableStorageCardEncryptionOnMobile sets the bitLockerEnableStorageCardEncryptionOnMobile property value. Allows the admin to require encryption to be turned on using BitLocker. This policy is valid only for a mobile SKU.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerEnableStorageCardEncryptionOnMobile(value *bool)() {
    m.bitLockerEnableStorageCardEncryptionOnMobile = value
}
// SetBitLockerEncryptDevice sets the bitLockerEncryptDevice property value. Allows the admin to require encryption to be turned on using BitLocker.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerEncryptDevice(value *bool)() {
    m.bitLockerEncryptDevice = value
}
// SetBitLockerRemovableDrivePolicy sets the bitLockerRemovableDrivePolicy property value. BitLocker Removable Drive Policy.
func (m *Windows10EndpointProtectionConfiguration) SetBitLockerRemovableDrivePolicy(value BitLockerRemovableDrivePolicyable)() {
    m.bitLockerRemovableDrivePolicy = value
}
// SetDefenderAdditionalGuardedFolders sets the defenderAdditionalGuardedFolders property value. List of folder paths to be added to the list of protected folders
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAdditionalGuardedFolders(value []string)() {
    m.defenderAdditionalGuardedFolders = value
}
// SetDefenderAttackSurfaceReductionExcludedPaths sets the defenderAttackSurfaceReductionExcludedPaths property value. List of exe files and folders to be excluded from attack surface reduction rules
func (m *Windows10EndpointProtectionConfiguration) SetDefenderAttackSurfaceReductionExcludedPaths(value []string)() {
    m.defenderAttackSurfaceReductionExcludedPaths = value
}
// SetDefenderExploitProtectionXml sets the defenderExploitProtectionXml property value. Xml content containing information regarding exploit protection details.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXml(value []byte)() {
    m.defenderExploitProtectionXml = value
}
// SetDefenderExploitProtectionXmlFileName sets the defenderExploitProtectionXmlFileName property value. Name of the file from which DefenderExploitProtectionXml was obtained.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderExploitProtectionXmlFileName(value *string)() {
    m.defenderExploitProtectionXmlFileName = value
}
// SetDefenderGuardedFoldersAllowedAppPaths sets the defenderGuardedFoldersAllowedAppPaths property value. List of paths to exe that are allowed to access protected folders
func (m *Windows10EndpointProtectionConfiguration) SetDefenderGuardedFoldersAllowedAppPaths(value []string)() {
    m.defenderGuardedFoldersAllowedAppPaths = value
}
// SetDefenderSecurityCenterBlockExploitProtectionOverride sets the defenderSecurityCenterBlockExploitProtectionOverride property value. Indicates whether or not to block user from overriding Exploit Protection settings.
func (m *Windows10EndpointProtectionConfiguration) SetDefenderSecurityCenterBlockExploitProtectionOverride(value *bool)() {
    m.defenderSecurityCenterBlockExploitProtectionOverride = value
}
// SetFirewallBlockStatefulFTP sets the firewallBlockStatefulFTP property value. Blocks stateful FTP connections to the device
func (m *Windows10EndpointProtectionConfiguration) SetFirewallBlockStatefulFTP(value *bool)() {
    m.firewallBlockStatefulFTP = value
}
// SetFirewallCertificateRevocationListCheckMethod sets the firewallCertificateRevocationListCheckMethod property value. Possible values for firewallCertificateRevocationListCheckMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallCertificateRevocationListCheckMethod(value *FirewallCertificateRevocationListCheckMethodType)() {
    m.firewallCertificateRevocationListCheckMethod = value
}
// SetFirewallIdleTimeoutForSecurityAssociationInSeconds sets the firewallIdleTimeoutForSecurityAssociationInSeconds property value. Configures the idle timeout for security associations, in seconds, from 300 to 3600 inclusive. This is the period after which security associations will expire and be deleted. Valid values 300 to 3600
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIdleTimeoutForSecurityAssociationInSeconds(value *int32)() {
    m.firewallIdleTimeoutForSecurityAssociationInSeconds = value
}
// SetFirewallIPSecExemptionsAllowDHCP sets the firewallIPSecExemptionsAllowDHCP property value. Configures IPSec exemptions to allow both IPv4 and IPv6 DHCP traffic
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowDHCP(value *bool)() {
    m.firewallIPSecExemptionsAllowDHCP = value
}
// SetFirewallIPSecExemptionsAllowICMP sets the firewallIPSecExemptionsAllowICMP property value. Configures IPSec exemptions to allow ICMP
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowICMP(value *bool)() {
    m.firewallIPSecExemptionsAllowICMP = value
}
// SetFirewallIPSecExemptionsAllowNeighborDiscovery sets the firewallIPSecExemptionsAllowNeighborDiscovery property value. Configures IPSec exemptions to allow neighbor discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowNeighborDiscovery(value *bool)() {
    m.firewallIPSecExemptionsAllowNeighborDiscovery = value
}
// SetFirewallIPSecExemptionsAllowRouterDiscovery sets the firewallIPSecExemptionsAllowRouterDiscovery property value. Configures IPSec exemptions to allow router discovery IPv6 ICMP type-codes
func (m *Windows10EndpointProtectionConfiguration) SetFirewallIPSecExemptionsAllowRouterDiscovery(value *bool)() {
    m.firewallIPSecExemptionsAllowRouterDiscovery = value
}
// SetFirewallMergeKeyingModuleSettings sets the firewallMergeKeyingModuleSettings property value. If an authentication set is not fully supported by a keying module, direct the module to ignore only unsupported authentication suites rather than the entire set
func (m *Windows10EndpointProtectionConfiguration) SetFirewallMergeKeyingModuleSettings(value *bool)() {
    m.firewallMergeKeyingModuleSettings = value
}
// SetFirewallPacketQueueingMethod sets the firewallPacketQueueingMethod property value. Possible values for firewallPacketQueueingMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallPacketQueueingMethod(value *FirewallPacketQueueingMethodType)() {
    m.firewallPacketQueueingMethod = value
}
// SetFirewallPreSharedKeyEncodingMethod sets the firewallPreSharedKeyEncodingMethod property value. Possible values for firewallPreSharedKeyEncodingMethod
func (m *Windows10EndpointProtectionConfiguration) SetFirewallPreSharedKeyEncodingMethod(value *FirewallPreSharedKeyEncodingMethodType)() {
    m.firewallPreSharedKeyEncodingMethod = value
}
// SetFirewallProfileDomain sets the firewallProfileDomain property value. Configures the firewall profile settings for domain networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfileDomain(value WindowsFirewallNetworkProfileable)() {
    m.firewallProfileDomain = value
}
// SetFirewallProfilePrivate sets the firewallProfilePrivate property value. Configures the firewall profile settings for private networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfilePrivate(value WindowsFirewallNetworkProfileable)() {
    m.firewallProfilePrivate = value
}
// SetFirewallProfilePublic sets the firewallProfilePublic property value. Configures the firewall profile settings for public networks
func (m *Windows10EndpointProtectionConfiguration) SetFirewallProfilePublic(value WindowsFirewallNetworkProfileable)() {
    m.firewallProfilePublic = value
}
// SetSmartScreenBlockOverrideForFiles sets the smartScreenBlockOverrideForFiles property value. Allows IT Admins to control whether users can can ignore SmartScreen warnings and run malicious files.
func (m *Windows10EndpointProtectionConfiguration) SetSmartScreenBlockOverrideForFiles(value *bool)() {
    m.smartScreenBlockOverrideForFiles = value
}
// SetSmartScreenEnableInShell sets the smartScreenEnableInShell property value. Allows IT Admins to configure SmartScreen for Windows.
func (m *Windows10EndpointProtectionConfiguration) SetSmartScreenEnableInShell(value *bool)() {
    m.smartScreenEnableInShell = value
}
