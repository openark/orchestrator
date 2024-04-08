package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceAppManagement 
type DeviceAppManagement struct {
    Entity
    // Android managed app policies.
    androidManagedAppProtections []AndroidManagedAppProtectionable
    // Default managed app policies.
    defaultManagedAppProtections []DefaultManagedAppProtectionable
    // iOS managed app policies.
    iosManagedAppProtections []IosManagedAppProtectionable
    // Whether the account is enabled for syncing applications from the Microsoft Store for Business.
    isEnabledForMicrosoftStoreForBusiness *bool
    // Managed app policies.
    managedAppPolicies []ManagedAppPolicyable
    // The managed app registrations.
    managedAppRegistrations []ManagedAppRegistrationable
    // The managed app statuses.
    managedAppStatuses []ManagedAppStatusable
    // The Managed eBook.
    managedEBooks []ManagedEBookable
    // Windows information protection for apps running on devices which are MDM enrolled.
    mdmWindowsInformationProtectionPolicies []MdmWindowsInformationProtectionPolicyable
    // The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
    microsoftStoreForBusinessLanguage *string
    // The last time an application sync from the Microsoft Store for Business was completed.
    microsoftStoreForBusinessLastCompletedApplicationSyncTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The last time the apps from the Microsoft Store for Business were synced successfully for the account.
    microsoftStoreForBusinessLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The mobile app categories.
    mobileAppCategories []MobileAppCategoryable
    // The Managed Device Mobile Application Configurations.
    mobileAppConfigurations []ManagedDeviceMobileAppConfigurationable
    // The mobile apps.
    mobileApps []MobileAppable
    // Targeted managed app configurations.
    targetedManagedAppConfigurations []TargetedManagedAppConfigurationable
    // List of Vpp tokens for this organization.
    vppTokens []VppTokenable
    // Windows information protection for apps running on devices which are not MDM enrolled.
    windowsInformationProtectionPolicies []WindowsInformationProtectionPolicyable
}
// NewDeviceAppManagement instantiates a new DeviceAppManagement and sets the default values.
func NewDeviceAppManagement()(*DeviceAppManagement) {
    m := &DeviceAppManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceAppManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceAppManagementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceAppManagement(), nil
}
// GetAndroidManagedAppProtections gets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) GetAndroidManagedAppProtections()([]AndroidManagedAppProtectionable) {
    return m.androidManagedAppProtections
}
// GetDefaultManagedAppProtections gets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) GetDefaultManagedAppProtections()([]DefaultManagedAppProtectionable) {
    return m.defaultManagedAppProtections
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAndroidManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedAppProtectionable, len(val))
            for i, v := range val {
                res[i] = v.(AndroidManagedAppProtectionable)
            }
            m.SetAndroidManagedAppProtections(res)
        }
        return nil
    }
    res["defaultManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDefaultManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DefaultManagedAppProtectionable, len(val))
            for i, v := range val {
                res[i] = v.(DefaultManagedAppProtectionable)
            }
            m.SetDefaultManagedAppProtections(res)
        }
        return nil
    }
    res["iosManagedAppProtections"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIosManagedAppProtectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosManagedAppProtectionable, len(val))
            for i, v := range val {
                res[i] = v.(IosManagedAppProtectionable)
            }
            m.SetIosManagedAppProtections(res)
        }
        return nil
    }
    res["isEnabledForMicrosoftStoreForBusiness"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForMicrosoftStoreForBusiness(val)
        }
        return nil
    }
    res["managedAppPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppPolicyable)
            }
            m.SetManagedAppPolicies(res)
        }
        return nil
    }
    res["managedAppRegistrations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppRegistrationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppRegistrationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppRegistrationable)
            }
            m.SetManagedAppRegistrations(res)
        }
        return nil
    }
    res["managedAppStatuses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedAppStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppStatusable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedAppStatusable)
            }
            m.SetManagedAppStatuses(res)
        }
        return nil
    }
    res["managedEBooks"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedEBookFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBookable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedEBookable)
            }
            m.SetManagedEBooks(res)
        }
        return nil
    }
    res["mdmWindowsInformationProtectionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMdmWindowsInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MdmWindowsInformationProtectionPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(MdmWindowsInformationProtectionPolicyable)
            }
            m.SetMdmWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    res["microsoftStoreForBusinessLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLanguage(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastCompletedApplicationSyncTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastSuccessfulSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["mobileAppCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppCategoryable)
            }
            m.SetMobileAppCategories(res)
        }
        return nil
    }
    res["mobileAppConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedDeviceMobileAppConfigurationable)
            }
            m.SetMobileAppConfigurations(res)
        }
        return nil
    }
    res["mobileApps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(MobileAppable)
            }
            m.SetMobileApps(res)
        }
        return nil
    }
    res["targetedManagedAppConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetedManagedAppConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(TargetedManagedAppConfigurationable)
            }
            m.SetTargetedManagedAppConfigurations(res)
        }
        return nil
    }
    res["vppTokens"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVppTokenFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VppTokenable, len(val))
            for i, v := range val {
                res[i] = v.(VppTokenable)
            }
            m.SetVppTokens(res)
        }
        return nil
    }
    res["windowsInformationProtectionPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWindowsInformationProtectionPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(WindowsInformationProtectionPolicyable)
            }
            m.SetWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    return res
}
// GetIosManagedAppProtections gets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) GetIosManagedAppProtections()([]IosManagedAppProtectionable) {
    return m.iosManagedAppProtections
}
// GetIsEnabledForMicrosoftStoreForBusiness gets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness()(*bool) {
    return m.isEnabledForMicrosoftStoreForBusiness
}
// GetManagedAppPolicies gets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) GetManagedAppPolicies()([]ManagedAppPolicyable) {
    return m.managedAppPolicies
}
// GetManagedAppRegistrations gets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) GetManagedAppRegistrations()([]ManagedAppRegistrationable) {
    return m.managedAppRegistrations
}
// GetManagedAppStatuses gets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) GetManagedAppStatuses()([]ManagedAppStatusable) {
    return m.managedAppStatuses
}
// GetManagedEBooks gets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) GetManagedEBooks()([]ManagedEBookable) {
    return m.managedEBooks
}
// GetMdmWindowsInformationProtectionPolicies gets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicyable) {
    return m.mdmWindowsInformationProtectionPolicies
}
// GetMicrosoftStoreForBusinessLanguage gets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage()(*string) {
    return m.microsoftStoreForBusinessLanguage
}
// GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime gets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.microsoftStoreForBusinessLastCompletedApplicationSyncTime
}
// GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime gets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.microsoftStoreForBusinessLastSuccessfulSyncDateTime
}
// GetMobileAppCategories gets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) GetMobileAppCategories()([]MobileAppCategoryable) {
    return m.mobileAppCategories
}
// GetMobileAppConfigurations gets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfigurationable) {
    return m.mobileAppConfigurations
}
// GetMobileApps gets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) GetMobileApps()([]MobileAppable) {
    return m.mobileApps
}
// GetTargetedManagedAppConfigurations gets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfigurationable) {
    return m.targetedManagedAppConfigurations
}
// GetVppTokens gets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) GetVppTokens()([]VppTokenable) {
    return m.vppTokens
}
// GetWindowsInformationProtectionPolicies gets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicyable) {
    return m.windowsInformationProtectionPolicies
}
// Serialize serializes information the current object
func (m *DeviceAppManagement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAndroidManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAndroidManagedAppProtections()))
        for i, v := range m.GetAndroidManagedAppProtections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("androidManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDefaultManagedAppProtections()))
        for i, v := range m.GetDefaultManagedAppProtections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("defaultManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIosManagedAppProtections() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIosManagedAppProtections()))
        for i, v := range m.GetIosManagedAppProtections() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("iosManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledForMicrosoftStoreForBusiness", m.GetIsEnabledForMicrosoftStoreForBusiness())
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppPolicies()))
        for i, v := range m.GetManagedAppPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppRegistrations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppRegistrations()))
        for i, v := range m.GetManagedAppRegistrations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedAppStatuses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedAppStatuses()))
        for i, v := range m.GetManagedAppStatuses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppStatuses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetManagedEBooks() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetManagedEBooks()))
        for i, v := range m.GetManagedEBooks() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("managedEBooks", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMdmWindowsInformationProtectionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMdmWindowsInformationProtectionPolicies()))
        for i, v := range m.GetMdmWindowsInformationProtectionPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mdmWindowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftStoreForBusinessLanguage", m.GetMicrosoftStoreForBusinessLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("microsoftStoreForBusinessLastCompletedApplicationSyncTime", m.GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("microsoftStoreForBusinessLastSuccessfulSyncDateTime", m.GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppCategories()))
        for i, v := range m.GetMobileAppCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppCategories", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileAppConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileAppConfigurations()))
        for i, v := range m.GetMobileAppConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMobileApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMobileApps()))
        for i, v := range m.GetMobileApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mobileApps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTargetedManagedAppConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargetedManagedAppConfigurations()))
        for i, v := range m.GetTargetedManagedAppConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("targetedManagedAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetVppTokens() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVppTokens()))
        for i, v := range m.GetVppTokens() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("vppTokens", cast)
        if err != nil {
            return err
        }
    }
    if m.GetWindowsInformationProtectionPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWindowsInformationProtectionPolicies()))
        for i, v := range m.GetWindowsInformationProtectionPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidManagedAppProtections sets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) SetAndroidManagedAppProtections(value []AndroidManagedAppProtectionable)() {
    m.androidManagedAppProtections = value
}
// SetDefaultManagedAppProtections sets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) SetDefaultManagedAppProtections(value []DefaultManagedAppProtectionable)() {
    m.defaultManagedAppProtections = value
}
// SetIosManagedAppProtections sets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) SetIosManagedAppProtections(value []IosManagedAppProtectionable)() {
    m.iosManagedAppProtections = value
}
// SetIsEnabledForMicrosoftStoreForBusiness sets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(value *bool)() {
    m.isEnabledForMicrosoftStoreForBusiness = value
}
// SetManagedAppPolicies sets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) SetManagedAppPolicies(value []ManagedAppPolicyable)() {
    m.managedAppPolicies = value
}
// SetManagedAppRegistrations sets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) SetManagedAppRegistrations(value []ManagedAppRegistrationable)() {
    m.managedAppRegistrations = value
}
// SetManagedAppStatuses sets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) SetManagedAppStatuses(value []ManagedAppStatusable)() {
    m.managedAppStatuses = value
}
// SetManagedEBooks sets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) SetManagedEBooks(value []ManagedEBookable)() {
    m.managedEBooks = value
}
// SetMdmWindowsInformationProtectionPolicies sets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicyable)() {
    m.mdmWindowsInformationProtectionPolicies = value
}
// SetMicrosoftStoreForBusinessLanguage sets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(value *string)() {
    m.microsoftStoreForBusinessLanguage = value
}
// SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime sets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastCompletedApplicationSyncTime = value
}
// SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime sets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastSuccessfulSyncDateTime = value
}
// SetMobileAppCategories sets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) SetMobileAppCategories(value []MobileAppCategoryable)() {
    m.mobileAppCategories = value
}
// SetMobileAppConfigurations sets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfigurationable)() {
    m.mobileAppConfigurations = value
}
// SetMobileApps sets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) SetMobileApps(value []MobileAppable)() {
    m.mobileApps = value
}
// SetTargetedManagedAppConfigurations sets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfigurationable)() {
    m.targetedManagedAppConfigurations = value
}
// SetVppTokens sets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) SetVppTokens(value []VppTokenable)() {
    m.vppTokens = value
}
// SetWindowsInformationProtectionPolicies sets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicyable)() {
    m.windowsInformationProtectionPolicies = value
}
