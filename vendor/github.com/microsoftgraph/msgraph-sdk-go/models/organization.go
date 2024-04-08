package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Organization 
type Organization struct {
    DirectoryObject
    // The collection of service plans associated with the tenant. Not nullable.
    assignedPlans []AssignedPlanable
    // Branding for the organization. Nullable.
    branding OrganizationalBrandingable
    // Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
    businessPhones []string
    // Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
    certificateBasedAuthConfiguration []CertificateBasedAuthConfigurationable
    // City name of the address for the organization.
    city *string
    // Country/region name of the address for the organization.
    country *string
    // Country or region abbreviation for the organization in ISO 3166-2 format.
    countryLetterCode *string
    // Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Two-letter ISO 3166 country code indicating the default service usage location of an organization.
    defaultUsageLocation *string
    // The display name for the tenant.
    displayName *string
    // The collection of open extensions defined for the organization. Read-only. Nullable.
    extensions []Extensionable
    // Not nullable.
    marketingNotificationEmails []string
    // Mobile device management authority.
    mobileDeviceManagementAuthority *MdmAuthority
    // The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced. Nullable. null if this object has never been synced from an on-premises directory (default).
    onPremisesSyncEnabled *bool
    // The type of partnership this tenant has with Microsoft. The possible values are: microsoftSupport, syndicatePartner, breadthPartner, breadthPartnerDelegatedAdmin, resellerPartnerDelegatedAdmin, valueAddedResellerPartnerDelegatedAdmin, unknownFutureValue. Nullable. For more information about the possible types, see partnerTenantType values.
    partnerTenantType *PartnerTenantType
    // Postal code of the address for the organization.
    postalCode *string
    // The preferred language for the organization. Should follow ISO 639-1 Code; for example, en.
    preferredLanguage *string
    // The privacy profile of an organization.
    privacyProfile PrivacyProfileable
    // Not nullable.
    provisionedPlans []ProvisionedPlanable
    // The securityComplianceNotificationMails property
    securityComplianceNotificationMails []string
    // The securityComplianceNotificationPhones property
    securityComplianceNotificationPhones []string
    // State name of the address for the organization.
    state *string
    // Street name of the address for organization.
    street *string
    // Not nullable.
    technicalNotificationMails []string
    // Not nullable. The tenant type option that was selected when the tenant was created. The possible values are:  AAD - An enterprise identity access management (IAM) service that serves business-to-employee and business-to-business (B2B) scenarios.  AAD B2C A customer identity access management (CIAM) service that serves business-to-consumer (B2C) scenarios.
    tenantType *string
    // The collection of domains associated with this tenant. Not nullable.
    verifiedDomains []VerifiedDomainable
}
// NewOrganization instantiates a new Organization and sets the default values.
func NewOrganization()(*Organization) {
    m := &Organization{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.organization"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOrganizationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrganizationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganization(), nil
}
// GetAssignedPlans gets the assignedPlans property value. The collection of service plans associated with the tenant. Not nullable.
func (m *Organization) GetAssignedPlans()([]AssignedPlanable) {
    return m.assignedPlans
}
// GetBranding gets the branding property value. Branding for the organization. Nullable.
func (m *Organization) GetBranding()(OrganizationalBrandingable) {
    return m.branding
}
// GetBusinessPhones gets the businessPhones property value. Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
func (m *Organization) GetBusinessPhones()([]string) {
    return m.businessPhones
}
// GetCertificateBasedAuthConfiguration gets the certificateBasedAuthConfiguration property value. Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
func (m *Organization) GetCertificateBasedAuthConfiguration()([]CertificateBasedAuthConfigurationable) {
    return m.certificateBasedAuthConfiguration
}
// GetCity gets the city property value. City name of the address for the organization.
func (m *Organization) GetCity()(*string) {
    return m.city
}
// GetCountry gets the country property value. Country/region name of the address for the organization.
func (m *Organization) GetCountry()(*string) {
    return m.country
}
// GetCountryLetterCode gets the countryLetterCode property value. Country or region abbreviation for the organization in ISO 3166-2 format.
func (m *Organization) GetCountryLetterCode()(*string) {
    return m.countryLetterCode
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDefaultUsageLocation gets the defaultUsageLocation property value. Two-letter ISO 3166 country code indicating the default service usage location of an organization.
func (m *Organization) GetDefaultUsageLocation()(*string) {
    return m.defaultUsageLocation
}
// GetDisplayName gets the displayName property value. The display name for the tenant.
func (m *Organization) GetDisplayName()(*string) {
    return m.displayName
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the organization. Read-only. Nullable.
func (m *Organization) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Organization) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["assignedPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedPlanable, len(val))
            for i, v := range val {
                res[i] = v.(AssignedPlanable)
            }
            m.SetAssignedPlans(res)
        }
        return nil
    }
    res["branding"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationalBrandingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBranding(val.(OrganizationalBrandingable))
        }
        return nil
    }
    res["businessPhones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBusinessPhones(res)
        }
        return nil
    }
    res["certificateBasedAuthConfiguration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCertificateBasedAuthConfigurationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CertificateBasedAuthConfigurationable, len(val))
            for i, v := range val {
                res[i] = v.(CertificateBasedAuthConfigurationable)
            }
            m.SetCertificateBasedAuthConfiguration(res)
        }
        return nil
    }
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["country"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountry(val)
        }
        return nil
    }
    res["countryLetterCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryLetterCode(val)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["defaultUsageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultUsageLocation(val)
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
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["marketingNotificationEmails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetMarketingNotificationEmails(res)
        }
        return nil
    }
    res["mobileDeviceManagementAuthority"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMdmAuthority)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobileDeviceManagementAuthority(val.(*MdmAuthority))
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["partnerTenantType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePartnerTenantType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPartnerTenantType(val.(*PartnerTenantType))
        }
        return nil
    }
    res["postalCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["preferredLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreferredLanguage(val)
        }
        return nil
    }
    res["privacyProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePrivacyProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivacyProfile(val.(PrivacyProfileable))
        }
        return nil
    }
    res["provisionedPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisionedPlanFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisionedPlanable, len(val))
            for i, v := range val {
                res[i] = v.(ProvisionedPlanable)
            }
            m.SetProvisionedPlans(res)
        }
        return nil
    }
    res["securityComplianceNotificationMails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSecurityComplianceNotificationMails(res)
        }
        return nil
    }
    res["securityComplianceNotificationPhones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSecurityComplianceNotificationPhones(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["street"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreet(val)
        }
        return nil
    }
    res["technicalNotificationMails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTechnicalNotificationMails(res)
        }
        return nil
    }
    res["tenantType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantType(val)
        }
        return nil
    }
    res["verifiedDomains"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateVerifiedDomainFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VerifiedDomainable, len(val))
            for i, v := range val {
                res[i] = v.(VerifiedDomainable)
            }
            m.SetVerifiedDomains(res)
        }
        return nil
    }
    return res
}
// GetMarketingNotificationEmails gets the marketingNotificationEmails property value. Not nullable.
func (m *Organization) GetMarketingNotificationEmails()([]string) {
    return m.marketingNotificationEmails
}
// GetMobileDeviceManagementAuthority gets the mobileDeviceManagementAuthority property value. Mobile device management authority.
func (m *Organization) GetMobileDeviceManagementAuthority()(*MdmAuthority) {
    return m.mobileDeviceManagementAuthority
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced. Nullable. null if this object has never been synced from an on-premises directory (default).
func (m *Organization) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetPartnerTenantType gets the partnerTenantType property value. The type of partnership this tenant has with Microsoft. The possible values are: microsoftSupport, syndicatePartner, breadthPartner, breadthPartnerDelegatedAdmin, resellerPartnerDelegatedAdmin, valueAddedResellerPartnerDelegatedAdmin, unknownFutureValue. Nullable. For more information about the possible types, see partnerTenantType values.
func (m *Organization) GetPartnerTenantType()(*PartnerTenantType) {
    return m.partnerTenantType
}
// GetPostalCode gets the postalCode property value. Postal code of the address for the organization.
func (m *Organization) GetPostalCode()(*string) {
    return m.postalCode
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for the organization. Should follow ISO 639-1 Code; for example, en.
func (m *Organization) GetPreferredLanguage()(*string) {
    return m.preferredLanguage
}
// GetPrivacyProfile gets the privacyProfile property value. The privacy profile of an organization.
func (m *Organization) GetPrivacyProfile()(PrivacyProfileable) {
    return m.privacyProfile
}
// GetProvisionedPlans gets the provisionedPlans property value. Not nullable.
func (m *Organization) GetProvisionedPlans()([]ProvisionedPlanable) {
    return m.provisionedPlans
}
// GetSecurityComplianceNotificationMails gets the securityComplianceNotificationMails property value. The securityComplianceNotificationMails property
func (m *Organization) GetSecurityComplianceNotificationMails()([]string) {
    return m.securityComplianceNotificationMails
}
// GetSecurityComplianceNotificationPhones gets the securityComplianceNotificationPhones property value. The securityComplianceNotificationPhones property
func (m *Organization) GetSecurityComplianceNotificationPhones()([]string) {
    return m.securityComplianceNotificationPhones
}
// GetState gets the state property value. State name of the address for the organization.
func (m *Organization) GetState()(*string) {
    return m.state
}
// GetStreet gets the street property value. Street name of the address for organization.
func (m *Organization) GetStreet()(*string) {
    return m.street
}
// GetTechnicalNotificationMails gets the technicalNotificationMails property value. Not nullable.
func (m *Organization) GetTechnicalNotificationMails()([]string) {
    return m.technicalNotificationMails
}
// GetTenantType gets the tenantType property value. Not nullable. The tenant type option that was selected when the tenant was created. The possible values are:  AAD - An enterprise identity access management (IAM) service that serves business-to-employee and business-to-business (B2B) scenarios.  AAD B2C A customer identity access management (CIAM) service that serves business-to-consumer (B2C) scenarios.
func (m *Organization) GetTenantType()(*string) {
    return m.tenantType
}
// GetVerifiedDomains gets the verifiedDomains property value. The collection of domains associated with this tenant. Not nullable.
func (m *Organization) GetVerifiedDomains()([]VerifiedDomainable) {
    return m.verifiedDomains
}
// Serialize serializes information the current object
func (m *Organization) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedPlans()))
        for i, v := range m.GetAssignedPlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignedPlans", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("branding", m.GetBranding())
        if err != nil {
            return err
        }
    }
    if m.GetBusinessPhones() != nil {
        err = writer.WriteCollectionOfStringValues("businessPhones", m.GetBusinessPhones())
        if err != nil {
            return err
        }
    }
    if m.GetCertificateBasedAuthConfiguration() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCertificateBasedAuthConfiguration()))
        for i, v := range m.GetCertificateBasedAuthConfiguration() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("certificateBasedAuthConfiguration", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("country", m.GetCountry())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("countryLetterCode", m.GetCountryLetterCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultUsageLocation", m.GetDefaultUsageLocation())
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
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMarketingNotificationEmails() != nil {
        err = writer.WriteCollectionOfStringValues("marketingNotificationEmails", m.GetMarketingNotificationEmails())
        if err != nil {
            return err
        }
    }
    if m.GetMobileDeviceManagementAuthority() != nil {
        cast := (*m.GetMobileDeviceManagementAuthority()).String()
        err = writer.WriteStringValue("mobileDeviceManagementAuthority", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPartnerTenantType() != nil {
        cast := (*m.GetPartnerTenantType()).String()
        err = writer.WriteStringValue("partnerTenantType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preferredLanguage", m.GetPreferredLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("privacyProfile", m.GetPrivacyProfile())
        if err != nil {
            return err
        }
    }
    if m.GetProvisionedPlans() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisionedPlans()))
        for i, v := range m.GetProvisionedPlans() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("provisionedPlans", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSecurityComplianceNotificationMails() != nil {
        err = writer.WriteCollectionOfStringValues("securityComplianceNotificationMails", m.GetSecurityComplianceNotificationMails())
        if err != nil {
            return err
        }
    }
    if m.GetSecurityComplianceNotificationPhones() != nil {
        err = writer.WriteCollectionOfStringValues("securityComplianceNotificationPhones", m.GetSecurityComplianceNotificationPhones())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("street", m.GetStreet())
        if err != nil {
            return err
        }
    }
    if m.GetTechnicalNotificationMails() != nil {
        err = writer.WriteCollectionOfStringValues("technicalNotificationMails", m.GetTechnicalNotificationMails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantType", m.GetTenantType())
        if err != nil {
            return err
        }
    }
    if m.GetVerifiedDomains() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetVerifiedDomains()))
        for i, v := range m.GetVerifiedDomains() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("verifiedDomains", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedPlans sets the assignedPlans property value. The collection of service plans associated with the tenant. Not nullable.
func (m *Organization) SetAssignedPlans(value []AssignedPlanable)() {
    m.assignedPlans = value
}
// SetBranding sets the branding property value. Branding for the organization. Nullable.
func (m *Organization) SetBranding(value OrganizationalBrandingable)() {
    m.branding = value
}
// SetBusinessPhones sets the businessPhones property value. Telephone number for the organization. Although this is a string collection, only one number can be set for this property.
func (m *Organization) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// SetCertificateBasedAuthConfiguration sets the certificateBasedAuthConfiguration property value. Navigation property to manage certificate-based authentication configuration. Only a single instance of certificateBasedAuthConfiguration can be created in the collection.
func (m *Organization) SetCertificateBasedAuthConfiguration(value []CertificateBasedAuthConfigurationable)() {
    m.certificateBasedAuthConfiguration = value
}
// SetCity sets the city property value. City name of the address for the organization.
func (m *Organization) SetCity(value *string)() {
    m.city = value
}
// SetCountry sets the country property value. Country/region name of the address for the organization.
func (m *Organization) SetCountry(value *string)() {
    m.country = value
}
// SetCountryLetterCode sets the countryLetterCode property value. Country or region abbreviation for the organization in ISO 3166-2 format.
func (m *Organization) SetCountryLetterCode(value *string)() {
    m.countryLetterCode = value
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp of when the organization was created. The value cannot be modified and is automatically populated when the organization is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDefaultUsageLocation sets the defaultUsageLocation property value. Two-letter ISO 3166 country code indicating the default service usage location of an organization.
func (m *Organization) SetDefaultUsageLocation(value *string)() {
    m.defaultUsageLocation = value
}
// SetDisplayName sets the displayName property value. The display name for the tenant.
func (m *Organization) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the organization. Read-only. Nullable.
func (m *Organization) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetMarketingNotificationEmails sets the marketingNotificationEmails property value. Not nullable.
func (m *Organization) SetMarketingNotificationEmails(value []string)() {
    m.marketingNotificationEmails = value
}
// SetMobileDeviceManagementAuthority sets the mobileDeviceManagementAuthority property value. Mobile device management authority.
func (m *Organization) SetMobileDeviceManagementAuthority(value *MdmAuthority)() {
    m.mobileDeviceManagementAuthority = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. The time and date at which the tenant was last synced with the on-premises directory. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *Organization) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. true if this object is synced from an on-premises directory; false if this object was originally synced from an on-premises directory but is no longer synced. Nullable. null if this object has never been synced from an on-premises directory (default).
func (m *Organization) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetPartnerTenantType sets the partnerTenantType property value. The type of partnership this tenant has with Microsoft. The possible values are: microsoftSupport, syndicatePartner, breadthPartner, breadthPartnerDelegatedAdmin, resellerPartnerDelegatedAdmin, valueAddedResellerPartnerDelegatedAdmin, unknownFutureValue. Nullable. For more information about the possible types, see partnerTenantType values.
func (m *Organization) SetPartnerTenantType(value *PartnerTenantType)() {
    m.partnerTenantType = value
}
// SetPostalCode sets the postalCode property value. Postal code of the address for the organization.
func (m *Organization) SetPostalCode(value *string)() {
    m.postalCode = value
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for the organization. Should follow ISO 639-1 Code; for example, en.
func (m *Organization) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// SetPrivacyProfile sets the privacyProfile property value. The privacy profile of an organization.
func (m *Organization) SetPrivacyProfile(value PrivacyProfileable)() {
    m.privacyProfile = value
}
// SetProvisionedPlans sets the provisionedPlans property value. Not nullable.
func (m *Organization) SetProvisionedPlans(value []ProvisionedPlanable)() {
    m.provisionedPlans = value
}
// SetSecurityComplianceNotificationMails sets the securityComplianceNotificationMails property value. The securityComplianceNotificationMails property
func (m *Organization) SetSecurityComplianceNotificationMails(value []string)() {
    m.securityComplianceNotificationMails = value
}
// SetSecurityComplianceNotificationPhones sets the securityComplianceNotificationPhones property value. The securityComplianceNotificationPhones property
func (m *Organization) SetSecurityComplianceNotificationPhones(value []string)() {
    m.securityComplianceNotificationPhones = value
}
// SetState sets the state property value. State name of the address for the organization.
func (m *Organization) SetState(value *string)() {
    m.state = value
}
// SetStreet sets the street property value. Street name of the address for organization.
func (m *Organization) SetStreet(value *string)() {
    m.street = value
}
// SetTechnicalNotificationMails sets the technicalNotificationMails property value. Not nullable.
func (m *Organization) SetTechnicalNotificationMails(value []string)() {
    m.technicalNotificationMails = value
}
// SetTenantType sets the tenantType property value. Not nullable. The tenant type option that was selected when the tenant was created. The possible values are:  AAD - An enterprise identity access management (IAM) service that serves business-to-employee and business-to-business (B2B) scenarios.  AAD B2C A customer identity access management (CIAM) service that serves business-to-consumer (B2C) scenarios.
func (m *Organization) SetTenantType(value *string)() {
    m.tenantType = value
}
// SetVerifiedDomains sets the verifiedDomains property value. The collection of domains associated with this tenant. Not nullable.
func (m *Organization) SetVerifiedDomains(value []VerifiedDomainable)() {
    m.verifiedDomains = value
}
