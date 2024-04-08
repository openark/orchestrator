package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Application 
type Application struct {
    DirectoryObject
    // Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Office 365 call the application in the context of a document the user is working on.
    addIns []AddInable
    // Specifies settings for an application that implements a web API.
    api ApiApplicationable
    // The unique identifier for the application that is assigned to an application by Azure AD. Not nullable. Read-only. Supports $filter (eq).
    appId *string
    // Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne).
    applicationTemplateId *string
    // The collection of roles defined for the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
    appRoles []AppRoleable
    // Specifies the certification status of the application.
    certification Certificationable
    // The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Supports $filter (/$count eq 0, /$count ne 0). Read-only.
    createdOnBehalfOf DirectoryObjectable
    // The defaultRedirectUri property
    defaultRedirectUri *string
    // Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
    description *string
    // Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
    disabledByMicrosoftStatus *string
    // The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
    displayName *string
    // Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0).
    extensionProperties []ExtensionPropertyable
    // Federated identities for applications. Supports $expand and $filter (startsWith, /$count eq 0, /$count ne 0).
    federatedIdentityCredentials []FederatedIdentityCredentialable
    // Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following valid string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all of the security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
    groupMembershipClaims *string
    // The homeRealmDiscoveryPolicies property
    homeRealmDiscoveryPolicies []HomeRealmDiscoveryPolicyable
    // Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
    identifierUris []string
    // Basic profile information of the application such as  app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
    info InformationalUrlable
    // Specifies whether this application supports device authentication without a user. The default is false.
    isDeviceOnlyAuthSupported *bool
    // Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where it is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
    isFallbackPublicClient *bool
    // The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
    keyCredentials []KeyCredentialable
    // The main logo for the application. Not nullable.
    logo []byte
    // Notes relevant for the management of the application.
    notes *string
    // The oauth2RequirePostResponse property
    oauth2RequirePostResponse *bool
    // Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
    optionalClaims OptionalClaimsable
    // Directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
    owners []DirectoryObjectable
    // Specifies parental control settings for an application.
    parentalControlSettings ParentalControlSettingsable
    // The collection of password credentials associated with the application. Not nullable.
    passwordCredentials []PasswordCredentialable
    // Specifies settings for installed clients such as desktop or mobile devices.
    publicClient PublicClientApplicationable
    // The verified publisher domain for the application. Read-only. For more information, see How to: Configure an application's publisher domain. Supports $filter (eq, ne, ge, le, startsWith).
    publisherDomain *string
    // The requestSignatureVerification property
    requestSignatureVerification RequestSignatureVerificationable
    // Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. For more information, see Limits on requested permissions per app. Not nullable. Supports $filter (eq, not, ge, le).
    requiredResourceAccess []RequiredResourceAccessable
    // The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
    samlMetadataUrl *string
    // References application or service contact information from a Service or Asset Management database. Nullable.
    serviceManagementReference *string
    // Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table. The value of this object also limits the number of permissions an app can request. For more information, see Limits on requested permissions per app. The value for this property has implications on other app object properties. As a result, if you change this property, you may need to change other properties first. For more information, see Validation differences for signInAudience.Supports $filter (eq, ne, not).
    signInAudience *string
    // Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
    spa SpaApplicationable
    // Custom strings that can be used to categorize and identify the application. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
    tags []string
    // Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
    tokenEncryptionKeyId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // The tokenIssuancePolicies property
    tokenIssuancePolicies []TokenIssuancePolicyable
    // The tokenLifetimePolicies property
    tokenLifetimePolicies []TokenLifetimePolicyable
    // Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
    verifiedPublisher VerifiedPublisherable
    // Specifies settings for a web application.
    web WebApplicationable
}
// NewApplication instantiates a new Application and sets the default values.
func NewApplication()(*Application) {
    m := &Application{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.application"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApplication(), nil
}
// GetAddIns gets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Office 365 call the application in the context of a document the user is working on.
func (m *Application) GetAddIns()([]AddInable) {
    return m.addIns
}
// GetApi gets the api property value. Specifies settings for an application that implements a web API.
func (m *Application) GetApi()(ApiApplicationable) {
    return m.api
}
// GetAppId gets the appId property value. The unique identifier for the application that is assigned to an application by Azure AD. Not nullable. Read-only. Supports $filter (eq).
func (m *Application) GetAppId()(*string) {
    return m.appId
}
// GetApplicationTemplateId gets the applicationTemplateId property value. Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne).
func (m *Application) GetApplicationTemplateId()(*string) {
    return m.applicationTemplateId
}
// GetAppRoles gets the appRoles property value. The collection of roles defined for the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
func (m *Application) GetAppRoles()([]AppRoleable) {
    return m.appRoles
}
// GetCertification gets the certification property value. Specifies the certification status of the application.
func (m *Application) GetCertification()(Certificationable) {
    return m.certification
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
func (m *Application) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCreatedOnBehalfOf gets the createdOnBehalfOf property value. Supports $filter (/$count eq 0, /$count ne 0). Read-only.
func (m *Application) GetCreatedOnBehalfOf()(DirectoryObjectable) {
    return m.createdOnBehalfOf
}
// GetDefaultRedirectUri gets the defaultRedirectUri property value. The defaultRedirectUri property
func (m *Application) GetDefaultRedirectUri()(*string) {
    return m.defaultRedirectUri
}
// GetDescription gets the description property value. Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *Application) GetDescription()(*string) {
    return m.description
}
// GetDisabledByMicrosoftStatus gets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *Application) GetDisabledByMicrosoftStatus()(*string) {
    return m.disabledByMicrosoftStatus
}
// GetDisplayName gets the displayName property value. The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Application) GetDisplayName()(*string) {
    return m.displayName
}
// GetExtensionProperties gets the extensionProperties property value. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0).
func (m *Application) GetExtensionProperties()([]ExtensionPropertyable) {
    return m.extensionProperties
}
// GetFederatedIdentityCredentials gets the federatedIdentityCredentials property value. Federated identities for applications. Supports $expand and $filter (startsWith, /$count eq 0, /$count ne 0).
func (m *Application) GetFederatedIdentityCredentials()([]FederatedIdentityCredentialable) {
    return m.federatedIdentityCredentials
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Application) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["addIns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAddInFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AddInable, len(val))
            for i, v := range val {
                res[i] = v.(AddInable)
            }
            m.SetAddIns(res)
        }
        return nil
    }
    res["api"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateApiApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApi(val.(ApiApplicationable))
        }
        return nil
    }
    res["appId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["applicationTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplicationTemplateId(val)
        }
        return nil
    }
    res["appRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppRoleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppRoleable, len(val))
            for i, v := range val {
                res[i] = v.(AppRoleable)
            }
            m.SetAppRoles(res)
        }
        return nil
    }
    res["certification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertification(val.(Certificationable))
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
    res["createdOnBehalfOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedOnBehalfOf(val.(DirectoryObjectable))
        }
        return nil
    }
    res["defaultRedirectUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultRedirectUri(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["disabledByMicrosoftStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisabledByMicrosoftStatus(val)
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
    res["extensionProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ExtensionPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(ExtensionPropertyable)
            }
            m.SetExtensionProperties(res)
        }
        return nil
    }
    res["federatedIdentityCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateFederatedIdentityCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]FederatedIdentityCredentialable, len(val))
            for i, v := range val {
                res[i] = v.(FederatedIdentityCredentialable)
            }
            m.SetFederatedIdentityCredentials(res)
        }
        return nil
    }
    res["groupMembershipClaims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupMembershipClaims(val)
        }
        return nil
    }
    res["homeRealmDiscoveryPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]HomeRealmDiscoveryPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(HomeRealmDiscoveryPolicyable)
            }
            m.SetHomeRealmDiscoveryPolicies(res)
        }
        return nil
    }
    res["identifierUris"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetIdentifierUris(res)
        }
        return nil
    }
    res["info"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInformationalUrlFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInfo(val.(InformationalUrlable))
        }
        return nil
    }
    res["isDeviceOnlyAuthSupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeviceOnlyAuthSupported(val)
        }
        return nil
    }
    res["isFallbackPublicClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFallbackPublicClient(val)
        }
        return nil
    }
    res["keyCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyCredentialable, len(val))
            for i, v := range val {
                res[i] = v.(KeyCredentialable)
            }
            m.SetKeyCredentials(res)
        }
        return nil
    }
    res["logo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogo(val)
        }
        return nil
    }
    res["notes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotes(val)
        }
        return nil
    }
    res["oauth2RequirePostResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOauth2RequirePostResponse(val)
        }
        return nil
    }
    res["optionalClaims"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOptionalClaimsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOptionalClaims(val.(OptionalClaimsable))
        }
        return nil
    }
    res["owners"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetOwners(res)
        }
        return nil
    }
    res["parentalControlSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParentalControlSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentalControlSettings(val.(ParentalControlSettingsable))
        }
        return nil
    }
    res["passwordCredentials"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePasswordCredentialFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PasswordCredentialable, len(val))
            for i, v := range val {
                res[i] = v.(PasswordCredentialable)
            }
            m.SetPasswordCredentials(res)
        }
        return nil
    }
    res["publicClient"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicClientApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublicClient(val.(PublicClientApplicationable))
        }
        return nil
    }
    res["publisherDomain"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPublisherDomain(val)
        }
        return nil
    }
    res["requestSignatureVerification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRequestSignatureVerificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestSignatureVerification(val.(RequestSignatureVerificationable))
        }
        return nil
    }
    res["requiredResourceAccess"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRequiredResourceAccessFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RequiredResourceAccessable, len(val))
            for i, v := range val {
                res[i] = v.(RequiredResourceAccessable)
            }
            m.SetRequiredResourceAccess(res)
        }
        return nil
    }
    res["samlMetadataUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSamlMetadataUrl(val)
        }
        return nil
    }
    res["serviceManagementReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceManagementReference(val)
        }
        return nil
    }
    res["signInAudience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignInAudience(val)
        }
        return nil
    }
    res["spa"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSpaApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpa(val.(SpaApplicationable))
        }
        return nil
    }
    res["tags"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetTags(res)
        }
        return nil
    }
    res["tokenEncryptionKeyId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenEncryptionKeyId(val)
        }
        return nil
    }
    res["tokenIssuancePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenIssuancePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenIssuancePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TokenIssuancePolicyable)
            }
            m.SetTokenIssuancePolicies(res)
        }
        return nil
    }
    res["tokenLifetimePolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTokenLifetimePolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TokenLifetimePolicyable, len(val))
            for i, v := range val {
                res[i] = v.(TokenLifetimePolicyable)
            }
            m.SetTokenLifetimePolicies(res)
        }
        return nil
    }
    res["verifiedPublisher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVerifiedPublisherFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerifiedPublisher(val.(VerifiedPublisherable))
        }
        return nil
    }
    res["web"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebApplicationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeb(val.(WebApplicationable))
        }
        return nil
    }
    return res
}
// GetGroupMembershipClaims gets the groupMembershipClaims property value. Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following valid string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all of the security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
func (m *Application) GetGroupMembershipClaims()(*string) {
    return m.groupMembershipClaims
}
// GetHomeRealmDiscoveryPolicies gets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies property
func (m *Application) GetHomeRealmDiscoveryPolicies()([]HomeRealmDiscoveryPolicyable) {
    return m.homeRealmDiscoveryPolicies
}
// GetIdentifierUris gets the identifierUris property value. Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) GetIdentifierUris()([]string) {
    return m.identifierUris
}
// GetInfo gets the info property value. Basic profile information of the application such as  app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *Application) GetInfo()(InformationalUrlable) {
    return m.info
}
// GetIsDeviceOnlyAuthSupported gets the isDeviceOnlyAuthSupported property value. Specifies whether this application supports device authentication without a user. The default is false.
func (m *Application) GetIsDeviceOnlyAuthSupported()(*bool) {
    return m.isDeviceOnlyAuthSupported
}
// GetIsFallbackPublicClient gets the isFallbackPublicClient property value. Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where it is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
func (m *Application) GetIsFallbackPublicClient()(*bool) {
    return m.isFallbackPublicClient
}
// GetKeyCredentials gets the keyCredentials property value. The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) GetKeyCredentials()([]KeyCredentialable) {
    return m.keyCredentials
}
// GetLogo gets the logo property value. The main logo for the application. Not nullable.
func (m *Application) GetLogo()([]byte) {
    return m.logo
}
// GetNotes gets the notes property value. Notes relevant for the management of the application.
func (m *Application) GetNotes()(*string) {
    return m.notes
}
// GetOauth2RequirePostResponse gets the oauth2RequirePostResponse property value. The oauth2RequirePostResponse property
func (m *Application) GetOauth2RequirePostResponse()(*bool) {
    return m.oauth2RequirePostResponse
}
// GetOptionalClaims gets the optionalClaims property value. Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
func (m *Application) GetOptionalClaims()(OptionalClaimsable) {
    return m.optionalClaims
}
// GetOwners gets the owners property value. Directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
func (m *Application) GetOwners()([]DirectoryObjectable) {
    return m.owners
}
// GetParentalControlSettings gets the parentalControlSettings property value. Specifies parental control settings for an application.
func (m *Application) GetParentalControlSettings()(ParentalControlSettingsable) {
    return m.parentalControlSettings
}
// GetPasswordCredentials gets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *Application) GetPasswordCredentials()([]PasswordCredentialable) {
    return m.passwordCredentials
}
// GetPublicClient gets the publicClient property value. Specifies settings for installed clients such as desktop or mobile devices.
func (m *Application) GetPublicClient()(PublicClientApplicationable) {
    return m.publicClient
}
// GetPublisherDomain gets the publisherDomain property value. The verified publisher domain for the application. Read-only. For more information, see How to: Configure an application's publisher domain. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) GetPublisherDomain()(*string) {
    return m.publisherDomain
}
// GetRequestSignatureVerification gets the requestSignatureVerification property value. The requestSignatureVerification property
func (m *Application) GetRequestSignatureVerification()(RequestSignatureVerificationable) {
    return m.requestSignatureVerification
}
// GetRequiredResourceAccess gets the requiredResourceAccess property value. Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. For more information, see Limits on requested permissions per app. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) GetRequiredResourceAccess()([]RequiredResourceAccessable) {
    return m.requiredResourceAccess
}
// GetSamlMetadataUrl gets the samlMetadataUrl property value. The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
func (m *Application) GetSamlMetadataUrl()(*string) {
    return m.samlMetadataUrl
}
// GetServiceManagementReference gets the serviceManagementReference property value. References application or service contact information from a Service or Asset Management database. Nullable.
func (m *Application) GetServiceManagementReference()(*string) {
    return m.serviceManagementReference
}
// GetSignInAudience gets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table. The value of this object also limits the number of permissions an app can request. For more information, see Limits on requested permissions per app. The value for this property has implications on other app object properties. As a result, if you change this property, you may need to change other properties first. For more information, see Validation differences for signInAudience.Supports $filter (eq, ne, not).
func (m *Application) GetSignInAudience()(*string) {
    return m.signInAudience
}
// GetSpa gets the spa property value. Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
func (m *Application) GetSpa()(SpaApplicationable) {
    return m.spa
}
// GetTags gets the tags property value. Custom strings that can be used to categorize and identify the application. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
func (m *Application) GetTags()([]string) {
    return m.tags
}
// GetTokenEncryptionKeyId gets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *Application) GetTokenEncryptionKeyId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.tokenEncryptionKeyId
}
// GetTokenIssuancePolicies gets the tokenIssuancePolicies property value. The tokenIssuancePolicies property
func (m *Application) GetTokenIssuancePolicies()([]TokenIssuancePolicyable) {
    return m.tokenIssuancePolicies
}
// GetTokenLifetimePolicies gets the tokenLifetimePolicies property value. The tokenLifetimePolicies property
func (m *Application) GetTokenLifetimePolicies()([]TokenLifetimePolicyable) {
    return m.tokenLifetimePolicies
}
// GetVerifiedPublisher gets the verifiedPublisher property value. Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
func (m *Application) GetVerifiedPublisher()(VerifiedPublisherable) {
    return m.verifiedPublisher
}
// GetWeb gets the web property value. Specifies settings for a web application.
func (m *Application) GetWeb()(WebApplicationable) {
    return m.web
}
// Serialize serializes information the current object
func (m *Application) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddIns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddIns()))
        for i, v := range m.GetAddIns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("addIns", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("api", m.GetApi())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applicationTemplateId", m.GetApplicationTemplateId())
        if err != nil {
            return err
        }
    }
    if m.GetAppRoles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppRoles()))
        for i, v := range m.GetAppRoles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appRoles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("certification", m.GetCertification())
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
        err = writer.WriteObjectValue("createdOnBehalfOf", m.GetCreatedOnBehalfOf())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultRedirectUri", m.GetDefaultRedirectUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("disabledByMicrosoftStatus", m.GetDisabledByMicrosoftStatus())
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
    if m.GetExtensionProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensionProperties()))
        for i, v := range m.GetExtensionProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensionProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederatedIdentityCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFederatedIdentityCredentials()))
        for i, v := range m.GetFederatedIdentityCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("federatedIdentityCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupMembershipClaims", m.GetGroupMembershipClaims())
        if err != nil {
            return err
        }
    }
    if m.GetHomeRealmDiscoveryPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHomeRealmDiscoveryPolicies()))
        for i, v := range m.GetHomeRealmDiscoveryPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("homeRealmDiscoveryPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetIdentifierUris() != nil {
        err = writer.WriteCollectionOfStringValues("identifierUris", m.GetIdentifierUris())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("info", m.GetInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeviceOnlyAuthSupported", m.GetIsDeviceOnlyAuthSupported())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFallbackPublicClient", m.GetIsFallbackPublicClient())
        if err != nil {
            return err
        }
    }
    if m.GetKeyCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetKeyCredentials()))
        for i, v := range m.GetKeyCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("keyCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("logo", m.GetLogo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notes", m.GetNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("oauth2RequirePostResponse", m.GetOauth2RequirePostResponse())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("optionalClaims", m.GetOptionalClaims())
        if err != nil {
            return err
        }
    }
    if m.GetOwners() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOwners()))
        for i, v := range m.GetOwners() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("owners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentalControlSettings", m.GetParentalControlSettings())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordCredentials() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPasswordCredentials()))
        for i, v := range m.GetPasswordCredentials() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("passwordCredentials", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("publicClient", m.GetPublicClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("publisherDomain", m.GetPublisherDomain())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestSignatureVerification", m.GetRequestSignatureVerification())
        if err != nil {
            return err
        }
    }
    if m.GetRequiredResourceAccess() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequiredResourceAccess()))
        for i, v := range m.GetRequiredResourceAccess() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("requiredResourceAccess", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("samlMetadataUrl", m.GetSamlMetadataUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceManagementReference", m.GetServiceManagementReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("signInAudience", m.GetSignInAudience())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("spa", m.GetSpa())
        if err != nil {
            return err
        }
    }
    if m.GetTags() != nil {
        err = writer.WriteCollectionOfStringValues("tags", m.GetTags())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteUUIDValue("tokenEncryptionKeyId", m.GetTokenEncryptionKeyId())
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuancePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenIssuancePolicies()))
        for i, v := range m.GetTokenIssuancePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tokenIssuancePolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTokenLifetimePolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTokenLifetimePolicies()))
        for i, v := range m.GetTokenLifetimePolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("tokenLifetimePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("verifiedPublisher", m.GetVerifiedPublisher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("web", m.GetWeb())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddIns sets the addIns property value. Defines custom behavior that a consuming service can use to call an app in specific contexts. For example, applications that can render file streams may set the addIns property for its 'FileHandler' functionality. This will let services like Office 365 call the application in the context of a document the user is working on.
func (m *Application) SetAddIns(value []AddInable)() {
    m.addIns = value
}
// SetApi sets the api property value. Specifies settings for an application that implements a web API.
func (m *Application) SetApi(value ApiApplicationable)() {
    m.api = value
}
// SetAppId sets the appId property value. The unique identifier for the application that is assigned to an application by Azure AD. Not nullable. Read-only. Supports $filter (eq).
func (m *Application) SetAppId(value *string)() {
    m.appId = value
}
// SetApplicationTemplateId sets the applicationTemplateId property value. Unique identifier of the applicationTemplate. Supports $filter (eq, not, ne).
func (m *Application) SetApplicationTemplateId(value *string)() {
    m.applicationTemplateId = value
}
// SetAppRoles sets the appRoles property value. The collection of roles defined for the application. With app role assignments, these roles can be assigned to users, groups, or service principals associated with other applications. Not nullable.
func (m *Application) SetAppRoles(value []AppRoleable)() {
    m.appRoles = value
}
// SetCertification sets the certification property value. Specifies the certification status of the application.
func (m *Application) SetCertification(value Certificationable)() {
    m.certification = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the application was registered. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.  Supports $filter (eq, ne, not, ge, le, in, and eq on null values) and $orderBy.
func (m *Application) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCreatedOnBehalfOf sets the createdOnBehalfOf property value. Supports $filter (/$count eq 0, /$count ne 0). Read-only.
func (m *Application) SetCreatedOnBehalfOf(value DirectoryObjectable)() {
    m.createdOnBehalfOf = value
}
// SetDefaultRedirectUri sets the defaultRedirectUri property value. The defaultRedirectUri property
func (m *Application) SetDefaultRedirectUri(value *string)() {
    m.defaultRedirectUri = value
}
// SetDescription sets the description property value. Free text field to provide a description of the application object to end users. The maximum allowed size is 1024 characters. Supports $filter (eq, ne, not, ge, le, startsWith) and $search.
func (m *Application) SetDescription(value *string)() {
    m.description = value
}
// SetDisabledByMicrosoftStatus sets the disabledByMicrosoftStatus property value. Specifies whether Microsoft has disabled the registered application. Possible values are: null (default value), NotDisabled, and DisabledDueToViolationOfServicesAgreement (reasons may include suspicious, abusive, or malicious activity, or a violation of the Microsoft Services Agreement).  Supports $filter (eq, ne, not).
func (m *Application) SetDisabledByMicrosoftStatus(value *string)() {
    m.disabledByMicrosoftStatus = value
}
// SetDisplayName sets the displayName property value. The display name for the application. Supports $filter (eq, ne, not, ge, le, in, startsWith, and eq on null values), $search, and $orderBy.
func (m *Application) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExtensionProperties sets the extensionProperties property value. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0).
func (m *Application) SetExtensionProperties(value []ExtensionPropertyable)() {
    m.extensionProperties = value
}
// SetFederatedIdentityCredentials sets the federatedIdentityCredentials property value. Federated identities for applications. Supports $expand and $filter (startsWith, /$count eq 0, /$count ne 0).
func (m *Application) SetFederatedIdentityCredentials(value []FederatedIdentityCredentialable)() {
    m.federatedIdentityCredentials = value
}
// SetGroupMembershipClaims sets the groupMembershipClaims property value. Configures the groups claim issued in a user or OAuth 2.0 access token that the application expects. To set this attribute, use one of the following valid string values: None, SecurityGroup (for security groups and Azure AD roles), All (this gets all of the security groups, distribution groups, and Azure AD directory roles that the signed-in user is a member of).
func (m *Application) SetGroupMembershipClaims(value *string)() {
    m.groupMembershipClaims = value
}
// SetHomeRealmDiscoveryPolicies sets the homeRealmDiscoveryPolicies property value. The homeRealmDiscoveryPolicies property
func (m *Application) SetHomeRealmDiscoveryPolicies(value []HomeRealmDiscoveryPolicyable)() {
    m.homeRealmDiscoveryPolicies = value
}
// SetIdentifierUris sets the identifierUris property value. Also known as App ID URI, this value is set when an application is used as a resource app. The identifierUris acts as the prefix for the scopes you'll reference in your API's code, and it must be globally unique. You can use the default value provided, which is in the form api://<application-client-id>, or specify a more readable URI like https://contoso.com/api. For more information on valid identifierUris patterns and best practices, see Azure AD application registration security best practices. Not nullable. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) SetIdentifierUris(value []string)() {
    m.identifierUris = value
}
// SetInfo sets the info property value. Basic profile information of the application such as  app's marketing, support, terms of service and privacy statement URLs. The terms of service and privacy statement are surfaced to users through the user consent experience. For more info, see How to: Add Terms of service and privacy statement for registered Azure AD apps. Supports $filter (eq, ne, not, ge, le, and eq on null values).
func (m *Application) SetInfo(value InformationalUrlable)() {
    m.info = value
}
// SetIsDeviceOnlyAuthSupported sets the isDeviceOnlyAuthSupported property value. Specifies whether this application supports device authentication without a user. The default is false.
func (m *Application) SetIsDeviceOnlyAuthSupported(value *bool)() {
    m.isDeviceOnlyAuthSupported = value
}
// SetIsFallbackPublicClient sets the isFallbackPublicClient property value. Specifies the fallback application type as public client, such as an installed application running on a mobile device. The default value is false which means the fallback application type is confidential client such as a web app. There are certain scenarios where Azure AD cannot determine the client application type. For example, the ROPC flow where it is configured without specifying a redirect URI. In those cases Azure AD interprets the application type based on the value of this property.
func (m *Application) SetIsFallbackPublicClient(value *bool)() {
    m.isFallbackPublicClient = value
}
// SetKeyCredentials sets the keyCredentials property value. The collection of key credentials associated with the application. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) SetKeyCredentials(value []KeyCredentialable)() {
    m.keyCredentials = value
}
// SetLogo sets the logo property value. The main logo for the application. Not nullable.
func (m *Application) SetLogo(value []byte)() {
    m.logo = value
}
// SetNotes sets the notes property value. Notes relevant for the management of the application.
func (m *Application) SetNotes(value *string)() {
    m.notes = value
}
// SetOauth2RequirePostResponse sets the oauth2RequirePostResponse property value. The oauth2RequirePostResponse property
func (m *Application) SetOauth2RequirePostResponse(value *bool)() {
    m.oauth2RequirePostResponse = value
}
// SetOptionalClaims sets the optionalClaims property value. Application developers can configure optional claims in their Azure AD applications to specify the claims that are sent to their application by the Microsoft security token service. For more information, see How to: Provide optional claims to your app.
func (m *Application) SetOptionalClaims(value OptionalClaimsable)() {
    m.optionalClaims = value
}
// SetOwners sets the owners property value. Directory objects that are owners of the application. Read-only. Nullable. Supports $expand and $filter (/$count eq 0, /$count ne 0, /$count eq 1, /$count ne 1).
func (m *Application) SetOwners(value []DirectoryObjectable)() {
    m.owners = value
}
// SetParentalControlSettings sets the parentalControlSettings property value. Specifies parental control settings for an application.
func (m *Application) SetParentalControlSettings(value ParentalControlSettingsable)() {
    m.parentalControlSettings = value
}
// SetPasswordCredentials sets the passwordCredentials property value. The collection of password credentials associated with the application. Not nullable.
func (m *Application) SetPasswordCredentials(value []PasswordCredentialable)() {
    m.passwordCredentials = value
}
// SetPublicClient sets the publicClient property value. Specifies settings for installed clients such as desktop or mobile devices.
func (m *Application) SetPublicClient(value PublicClientApplicationable)() {
    m.publicClient = value
}
// SetPublisherDomain sets the publisherDomain property value. The verified publisher domain for the application. Read-only. For more information, see How to: Configure an application's publisher domain. Supports $filter (eq, ne, ge, le, startsWith).
func (m *Application) SetPublisherDomain(value *string)() {
    m.publisherDomain = value
}
// SetRequestSignatureVerification sets the requestSignatureVerification property value. The requestSignatureVerification property
func (m *Application) SetRequestSignatureVerification(value RequestSignatureVerificationable)() {
    m.requestSignatureVerification = value
}
// SetRequiredResourceAccess sets the requiredResourceAccess property value. Specifies the resources that the application needs to access. This property also specifies the set of delegated permissions and application roles that it needs for each of those resources. This configuration of access to the required resources drives the consent experience. No more than 50 resource services (APIs) can be configured. Beginning mid-October 2021, the total number of required permissions must not exceed 400. For more information, see Limits on requested permissions per app. Not nullable. Supports $filter (eq, not, ge, le).
func (m *Application) SetRequiredResourceAccess(value []RequiredResourceAccessable)() {
    m.requiredResourceAccess = value
}
// SetSamlMetadataUrl sets the samlMetadataUrl property value. The URL where the service exposes SAML metadata for federation. This property is valid only for single-tenant applications. Nullable.
func (m *Application) SetSamlMetadataUrl(value *string)() {
    m.samlMetadataUrl = value
}
// SetServiceManagementReference sets the serviceManagementReference property value. References application or service contact information from a Service or Asset Management database. Nullable.
func (m *Application) SetServiceManagementReference(value *string)() {
    m.serviceManagementReference = value
}
// SetSignInAudience sets the signInAudience property value. Specifies the Microsoft accounts that are supported for the current application. The possible values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount (default), and PersonalMicrosoftAccount. See more in the table. The value of this object also limits the number of permissions an app can request. For more information, see Limits on requested permissions per app. The value for this property has implications on other app object properties. As a result, if you change this property, you may need to change other properties first. For more information, see Validation differences for signInAudience.Supports $filter (eq, ne, not).
func (m *Application) SetSignInAudience(value *string)() {
    m.signInAudience = value
}
// SetSpa sets the spa property value. Specifies settings for a single-page application, including sign out URLs and redirect URIs for authorization codes and access tokens.
func (m *Application) SetSpa(value SpaApplicationable)() {
    m.spa = value
}
// SetTags sets the tags property value. Custom strings that can be used to categorize and identify the application. Not nullable. Supports $filter (eq, not, ge, le, startsWith).
func (m *Application) SetTags(value []string)() {
    m.tags = value
}
// SetTokenEncryptionKeyId sets the tokenEncryptionKeyId property value. Specifies the keyId of a public key from the keyCredentials collection. When configured, Azure AD encrypts all the tokens it emits by using the key this property points to. The application code that receives the encrypted token must use the matching private key to decrypt the token before it can be used for the signed-in user.
func (m *Application) SetTokenEncryptionKeyId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.tokenEncryptionKeyId = value
}
// SetTokenIssuancePolicies sets the tokenIssuancePolicies property value. The tokenIssuancePolicies property
func (m *Application) SetTokenIssuancePolicies(value []TokenIssuancePolicyable)() {
    m.tokenIssuancePolicies = value
}
// SetTokenLifetimePolicies sets the tokenLifetimePolicies property value. The tokenLifetimePolicies property
func (m *Application) SetTokenLifetimePolicies(value []TokenLifetimePolicyable)() {
    m.tokenLifetimePolicies = value
}
// SetVerifiedPublisher sets the verifiedPublisher property value. Specifies the verified publisher of the application. For more information about how publisher verification helps support application security, trustworthiness, and compliance, see Publisher verification.
func (m *Application) SetVerifiedPublisher(value VerifiedPublisherable)() {
    m.verifiedPublisher = value
}
// SetWeb sets the web property value. Specifies settings for a web application.
func (m *Application) SetWeb(value WebApplicationable)() {
    m.web = value
}
