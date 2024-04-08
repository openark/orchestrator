package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationUser 
type EducationUser struct {
    Entity
    // True if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter.
    accountEnabled *bool
    // The licenses that are assigned to the user. Not nullable.
    assignedLicenses []AssignedLicenseable
    // The plans that are assigned to the user. Read-only. Not nullable.
    assignedPlans []AssignedPlanable
    // Assignments belonging to the user.
    assignments []EducationAssignmentable
    // The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
    businessPhones []string
    // Classes to which the user belongs. Nullable.
    classes []EducationClassable
    // The entity who created the user.
    createdBy IdentitySetable
    // The name for the department in which the user works. Supports $filter.
    department *string
    // The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Supports $filter and $orderby.
    displayName *string
    // Where this user was created from. Possible values are: sis, manual.
    externalSource *EducationExternalSource
    // The name of the external source this resource was generated from.
    externalSourceDetail *string
    // The given name (first name) of the user. Supports $filter.
    givenName *string
    // The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com. Read-Only. Supports $filter.
    mail *string
    // The mail address of the user.
    mailingAddress PhysicalAddressable
    // The mail alias for the user. This property must be specified when a user is created. Supports $filter.
    mailNickname *string
    // The middle name of the user.
    middleName *string
    // The primary cellular telephone number for the user.
    mobilePhone *string
    // The officeLocation property
    officeLocation *string
    // Additional information used to associate the Azure Active Directory user with its Active Directory counterpart.
    onPremisesInfo EducationOnPremisesInfoable
    // Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two can be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.
    passwordPolicies *string
    // Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required.
    passwordProfile PasswordProfileable
    // The preferred language for the user that should follow the ISO 639-1 code, for example, en-US.
    preferredLanguage *string
    // The primaryRole property
    primaryRole *EducationUserRole
    // The plans that are provisioned for the user. Read-only. Not nullable.
    provisionedPlans []ProvisionedPlanable
    // The refreshTokensValidFromDateTime property
    refreshTokensValidFromDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Related records associated with the user. Read-only.
    relatedContacts []RelatedContactable
    // The address where the user lives.
    residenceAddress PhysicalAddressable
    // When set, the grading rubric attached to the assignment.
    rubrics []EducationRubricable
    // Schools to which the user belongs. Nullable.
    schools []EducationSchoolable
    // True if the Outlook Global Address List should contain this user; otherwise, false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
    showInAddressList *bool
    // If the primary role is student, this block will contain student specific data.
    student EducationStudentable
    // The user's surname (family name or last name). Supports $filter.
    surname *string
    // Classes for which the user is a teacher.
    taughtClasses []EducationClassable
    // If the primary role is teacher, this block will contain teacher specific data.
    teacher EducationTeacherable
    // A two-letter country code (ISO standard 3166). Required for users who will be assigned licenses due to a legal requirement to check for availability of services in countries or regions. Examples include: US, JP, and GB. Not nullable. Supports $filter.
    usageLocation *string
    // The directory user that corresponds to this user.
    user Userable
    // The user principal name (UPN) of the user. The UPN is an internet-style login name for the user based on the internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of the organization. Supports $filter and $orderby.
    userPrincipalName *string
    // A string value that can be used to classify user types in your directory, such as Member and Guest. Supports $filter.
    userType *string
}
// NewEducationUser instantiates a new educationUser and sets the default values.
func NewEducationUser()(*EducationUser) {
    m := &EducationUser{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationUserFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationUserFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationUser(), nil
}
// GetAccountEnabled gets the accountEnabled property value. True if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter.
func (m *EducationUser) GetAccountEnabled()(*bool) {
    return m.accountEnabled
}
// GetAssignedLicenses gets the assignedLicenses property value. The licenses that are assigned to the user. Not nullable.
func (m *EducationUser) GetAssignedLicenses()([]AssignedLicenseable) {
    return m.assignedLicenses
}
// GetAssignedPlans gets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.
func (m *EducationUser) GetAssignedPlans()([]AssignedPlanable) {
    return m.assignedPlans
}
// GetAssignments gets the assignments property value. Assignments belonging to the user.
func (m *EducationUser) GetAssignments()([]EducationAssignmentable) {
    return m.assignments
}
// GetBusinessPhones gets the businessPhones property value. The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
func (m *EducationUser) GetBusinessPhones()([]string) {
    return m.businessPhones
}
// GetClasses gets the classes property value. Classes to which the user belongs. Nullable.
func (m *EducationUser) GetClasses()([]EducationClassable) {
    return m.classes
}
// GetCreatedBy gets the createdBy property value. The entity who created the user.
func (m *EducationUser) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetDepartment gets the department property value. The name for the department in which the user works. Supports $filter.
func (m *EducationUser) GetDepartment()(*string) {
    return m.department
}
// GetDisplayName gets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Supports $filter and $orderby.
func (m *EducationUser) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalSource gets the externalSource property value. Where this user was created from. Possible values are: sis, manual.
func (m *EducationUser) GetExternalSource()(*EducationExternalSource) {
    return m.externalSource
}
// GetExternalSourceDetail gets the externalSourceDetail property value. The name of the external source this resource was generated from.
func (m *EducationUser) GetExternalSourceDetail()(*string) {
    return m.externalSourceDetail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationUser) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accountEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountEnabled(val)
        }
        return nil
    }
    res["assignedLicenses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AssignedLicenseable, len(val))
            for i, v := range val {
                res[i] = v.(AssignedLicenseable)
            }
            m.SetAssignedLicenses(res)
        }
        return nil
    }
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
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(EducationAssignmentable)
            }
            m.SetAssignments(res)
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
    res["classes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClassable, len(val))
            for i, v := range val {
                res[i] = v.(EducationClassable)
            }
            m.SetClasses(res)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["department"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
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
    res["externalSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationExternalSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSource(val.(*EducationExternalSource))
        }
        return nil
    }
    res["externalSourceDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSourceDetail(val)
        }
        return nil
    }
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["mail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMail(val)
        }
        return nil
    }
    res["mailingAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailingAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["mailNickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["middleName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMiddleName(val)
        }
        return nil
    }
    res["mobilePhone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMobilePhone(val)
        }
        return nil
    }
    res["officeLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["onPremisesInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationOnPremisesInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesInfo(val.(EducationOnPremisesInfoable))
        }
        return nil
    }
    res["passwordPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordPolicies(val)
        }
        return nil
    }
    res["passwordProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePasswordProfileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordProfile(val.(PasswordProfileable))
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
    res["primaryRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationUserRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrimaryRole(val.(*EducationUserRole))
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
    res["refreshTokensValidFromDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRefreshTokensValidFromDateTime(val)
        }
        return nil
    }
    res["relatedContacts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRelatedContactFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RelatedContactable, len(val))
            for i, v := range val {
                res[i] = v.(RelatedContactable)
            }
            m.SetRelatedContacts(res)
        }
        return nil
    }
    res["residenceAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResidenceAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["rubrics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationRubricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationRubricable, len(val))
            for i, v := range val {
                res[i] = v.(EducationRubricable)
            }
            m.SetRubrics(res)
        }
        return nil
    }
    res["schools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSchoolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSchoolable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSchoolable)
            }
            m.SetSchools(res)
        }
        return nil
    }
    res["showInAddressList"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowInAddressList(val)
        }
        return nil
    }
    res["student"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationStudentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStudent(val.(EducationStudentable))
        }
        return nil
    }
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["taughtClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClassable, len(val))
            for i, v := range val {
                res[i] = v.(EducationClassable)
            }
            m.SetTaughtClasses(res)
        }
        return nil
    }
    res["teacher"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationTeacherFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeacher(val.(EducationTeacherable))
        }
        return nil
    }
    res["usageLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageLocation(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(Userable))
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    res["userType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserType(val)
        }
        return nil
    }
    return res
}
// GetGivenName gets the givenName property value. The given name (first name) of the user. Supports $filter.
func (m *EducationUser) GetGivenName()(*string) {
    return m.givenName
}
// GetMail gets the mail property value. The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com. Read-Only. Supports $filter.
func (m *EducationUser) GetMail()(*string) {
    return m.mail
}
// GetMailingAddress gets the mailingAddress property value. The mail address of the user.
func (m *EducationUser) GetMailingAddress()(PhysicalAddressable) {
    return m.mailingAddress
}
// GetMailNickname gets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Supports $filter.
func (m *EducationUser) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetMiddleName gets the middleName property value. The middle name of the user.
func (m *EducationUser) GetMiddleName()(*string) {
    return m.middleName
}
// GetMobilePhone gets the mobilePhone property value. The primary cellular telephone number for the user.
func (m *EducationUser) GetMobilePhone()(*string) {
    return m.mobilePhone
}
// GetOfficeLocation gets the officeLocation property value. The officeLocation property
func (m *EducationUser) GetOfficeLocation()(*string) {
    return m.officeLocation
}
// GetOnPremisesInfo gets the onPremisesInfo property value. Additional information used to associate the Azure Active Directory user with its Active Directory counterpart.
func (m *EducationUser) GetOnPremisesInfo()(EducationOnPremisesInfoable) {
    return m.onPremisesInfo
}
// GetPasswordPolicies gets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two can be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.
func (m *EducationUser) GetPasswordPolicies()(*string) {
    return m.passwordPolicies
}
// GetPasswordProfile gets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required.
func (m *EducationUser) GetPasswordProfile()(PasswordProfileable) {
    return m.passwordProfile
}
// GetPreferredLanguage gets the preferredLanguage property value. The preferred language for the user that should follow the ISO 639-1 code, for example, en-US.
func (m *EducationUser) GetPreferredLanguage()(*string) {
    return m.preferredLanguage
}
// GetPrimaryRole gets the primaryRole property value. The primaryRole property
func (m *EducationUser) GetPrimaryRole()(*EducationUserRole) {
    return m.primaryRole
}
// GetProvisionedPlans gets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable.
func (m *EducationUser) GetProvisionedPlans()([]ProvisionedPlanable) {
    return m.provisionedPlans
}
// GetRefreshTokensValidFromDateTime gets the refreshTokensValidFromDateTime property value. The refreshTokensValidFromDateTime property
func (m *EducationUser) GetRefreshTokensValidFromDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.refreshTokensValidFromDateTime
}
// GetRelatedContacts gets the relatedContacts property value. Related records associated with the user. Read-only.
func (m *EducationUser) GetRelatedContacts()([]RelatedContactable) {
    return m.relatedContacts
}
// GetResidenceAddress gets the residenceAddress property value. The address where the user lives.
func (m *EducationUser) GetResidenceAddress()(PhysicalAddressable) {
    return m.residenceAddress
}
// GetRubrics gets the rubrics property value. When set, the grading rubric attached to the assignment.
func (m *EducationUser) GetRubrics()([]EducationRubricable) {
    return m.rubrics
}
// GetSchools gets the schools property value. Schools to which the user belongs. Nullable.
func (m *EducationUser) GetSchools()([]EducationSchoolable) {
    return m.schools
}
// GetShowInAddressList gets the showInAddressList property value. True if the Outlook Global Address List should contain this user; otherwise, false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
func (m *EducationUser) GetShowInAddressList()(*bool) {
    return m.showInAddressList
}
// GetStudent gets the student property value. If the primary role is student, this block will contain student specific data.
func (m *EducationUser) GetStudent()(EducationStudentable) {
    return m.student
}
// GetSurname gets the surname property value. The user's surname (family name or last name). Supports $filter.
func (m *EducationUser) GetSurname()(*string) {
    return m.surname
}
// GetTaughtClasses gets the taughtClasses property value. Classes for which the user is a teacher.
func (m *EducationUser) GetTaughtClasses()([]EducationClassable) {
    return m.taughtClasses
}
// GetTeacher gets the teacher property value. If the primary role is teacher, this block will contain teacher specific data.
func (m *EducationUser) GetTeacher()(EducationTeacherable) {
    return m.teacher
}
// GetUsageLocation gets the usageLocation property value. A two-letter country code (ISO standard 3166). Required for users who will be assigned licenses due to a legal requirement to check for availability of services in countries or regions. Examples include: US, JP, and GB. Not nullable. Supports $filter.
func (m *EducationUser) GetUsageLocation()(*string) {
    return m.usageLocation
}
// GetUser gets the user property value. The directory user that corresponds to this user.
func (m *EducationUser) GetUser()(Userable) {
    return m.user
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an internet-style login name for the user based on the internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of the organization. Supports $filter and $orderby.
func (m *EducationUser) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserType gets the userType property value. A string value that can be used to classify user types in your directory, such as Member and Guest. Supports $filter.
func (m *EducationUser) GetUserType()(*string) {
    return m.userType
}
// Serialize serializes information the current object
func (m *EducationUser) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("accountEnabled", m.GetAccountEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetAssignedLicenses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignedLicenses()))
        for i, v := range m.GetAssignedLicenses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignedLicenses", cast)
        if err != nil {
            return err
        }
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
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
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
    if m.GetClasses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClasses()))
        for i, v := range m.GetClasses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("classes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
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
    if m.GetExternalSource() != nil {
        cast := (*m.GetExternalSource()).String()
        err = writer.WriteStringValue("externalSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalSourceDetail", m.GetExternalSourceDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mailingAddress", m.GetMailingAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("middleName", m.GetMiddleName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mobilePhone", m.GetMobilePhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onPremisesInfo", m.GetOnPremisesInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("passwordPolicies", m.GetPasswordPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("passwordProfile", m.GetPasswordProfile())
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
    if m.GetPrimaryRole() != nil {
        cast := (*m.GetPrimaryRole()).String()
        err = writer.WriteStringValue("primaryRole", &cast)
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
    {
        err = writer.WriteTimeValue("refreshTokensValidFromDateTime", m.GetRefreshTokensValidFromDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetRelatedContacts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelatedContacts()))
        for i, v := range m.GetRelatedContacts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("relatedContacts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("residenceAddress", m.GetResidenceAddress())
        if err != nil {
            return err
        }
    }
    if m.GetRubrics() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRubrics()))
        for i, v := range m.GetRubrics() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("rubrics", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchools() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchools()))
        for i, v := range m.GetSchools() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("schools", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("showInAddressList", m.GetShowInAddressList())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("student", m.GetStudent())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    if m.GetTaughtClasses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTaughtClasses()))
        for i, v := range m.GetTaughtClasses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("taughtClasses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("teacher", m.GetTeacher())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("usageLocation", m.GetUsageLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userType", m.GetUserType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountEnabled sets the accountEnabled property value. True if the account is enabled; otherwise, false. This property is required when a user is created. Supports $filter.
func (m *EducationUser) SetAccountEnabled(value *bool)() {
    m.accountEnabled = value
}
// SetAssignedLicenses sets the assignedLicenses property value. The licenses that are assigned to the user. Not nullable.
func (m *EducationUser) SetAssignedLicenses(value []AssignedLicenseable)() {
    m.assignedLicenses = value
}
// SetAssignedPlans sets the assignedPlans property value. The plans that are assigned to the user. Read-only. Not nullable.
func (m *EducationUser) SetAssignedPlans(value []AssignedPlanable)() {
    m.assignedPlans = value
}
// SetAssignments sets the assignments property value. Assignments belonging to the user.
func (m *EducationUser) SetAssignments(value []EducationAssignmentable)() {
    m.assignments = value
}
// SetBusinessPhones sets the businessPhones property value. The telephone numbers for the user. Note: Although this is a string collection, only one number can be set for this property.
func (m *EducationUser) SetBusinessPhones(value []string)() {
    m.businessPhones = value
}
// SetClasses sets the classes property value. Classes to which the user belongs. Nullable.
func (m *EducationUser) SetClasses(value []EducationClassable)() {
    m.classes = value
}
// SetCreatedBy sets the createdBy property value. The entity who created the user.
func (m *EducationUser) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetDepartment sets the department property value. The name for the department in which the user works. Supports $filter.
func (m *EducationUser) SetDepartment(value *string)() {
    m.department = value
}
// SetDisplayName sets the displayName property value. The name displayed in the address book for the user. This is usually the combination of the user's first name, middle initial, and last name. This property is required when a user is created and it cannot be cleared during updates. Supports $filter and $orderby.
func (m *EducationUser) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalSource sets the externalSource property value. Where this user was created from. Possible values are: sis, manual.
func (m *EducationUser) SetExternalSource(value *EducationExternalSource)() {
    m.externalSource = value
}
// SetExternalSourceDetail sets the externalSourceDetail property value. The name of the external source this resource was generated from.
func (m *EducationUser) SetExternalSourceDetail(value *string)() {
    m.externalSourceDetail = value
}
// SetGivenName sets the givenName property value. The given name (first name) of the user. Supports $filter.
func (m *EducationUser) SetGivenName(value *string)() {
    m.givenName = value
}
// SetMail sets the mail property value. The SMTP address for the user, for example, jeff@contoso.onmicrosoft.com. Read-Only. Supports $filter.
func (m *EducationUser) SetMail(value *string)() {
    m.mail = value
}
// SetMailingAddress sets the mailingAddress property value. The mail address of the user.
func (m *EducationUser) SetMailingAddress(value PhysicalAddressable)() {
    m.mailingAddress = value
}
// SetMailNickname sets the mailNickname property value. The mail alias for the user. This property must be specified when a user is created. Supports $filter.
func (m *EducationUser) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetMiddleName sets the middleName property value. The middle name of the user.
func (m *EducationUser) SetMiddleName(value *string)() {
    m.middleName = value
}
// SetMobilePhone sets the mobilePhone property value. The primary cellular telephone number for the user.
func (m *EducationUser) SetMobilePhone(value *string)() {
    m.mobilePhone = value
}
// SetOfficeLocation sets the officeLocation property value. The officeLocation property
func (m *EducationUser) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// SetOnPremisesInfo sets the onPremisesInfo property value. Additional information used to associate the Azure Active Directory user with its Active Directory counterpart.
func (m *EducationUser) SetOnPremisesInfo(value EducationOnPremisesInfoable)() {
    m.onPremisesInfo = value
}
// SetPasswordPolicies sets the passwordPolicies property value. Specifies password policies for the user. This value is an enumeration with one possible value being DisableStrongPassword, which allows weaker passwords than the default policy to be specified. DisablePasswordExpiration can also be specified. The two can be specified together; for example: DisablePasswordExpiration, DisableStrongPassword.
func (m *EducationUser) SetPasswordPolicies(value *string)() {
    m.passwordPolicies = value
}
// SetPasswordProfile sets the passwordProfile property value. Specifies the password profile for the user. The profile contains the user's password. This property is required when a user is created. The password in the profile must satisfy minimum requirements as specified by the passwordPolicies property. By default, a strong password is required.
func (m *EducationUser) SetPasswordProfile(value PasswordProfileable)() {
    m.passwordProfile = value
}
// SetPreferredLanguage sets the preferredLanguage property value. The preferred language for the user that should follow the ISO 639-1 code, for example, en-US.
func (m *EducationUser) SetPreferredLanguage(value *string)() {
    m.preferredLanguage = value
}
// SetPrimaryRole sets the primaryRole property value. The primaryRole property
func (m *EducationUser) SetPrimaryRole(value *EducationUserRole)() {
    m.primaryRole = value
}
// SetProvisionedPlans sets the provisionedPlans property value. The plans that are provisioned for the user. Read-only. Not nullable.
func (m *EducationUser) SetProvisionedPlans(value []ProvisionedPlanable)() {
    m.provisionedPlans = value
}
// SetRefreshTokensValidFromDateTime sets the refreshTokensValidFromDateTime property value. The refreshTokensValidFromDateTime property
func (m *EducationUser) SetRefreshTokensValidFromDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.refreshTokensValidFromDateTime = value
}
// SetRelatedContacts sets the relatedContacts property value. Related records associated with the user. Read-only.
func (m *EducationUser) SetRelatedContacts(value []RelatedContactable)() {
    m.relatedContacts = value
}
// SetResidenceAddress sets the residenceAddress property value. The address where the user lives.
func (m *EducationUser) SetResidenceAddress(value PhysicalAddressable)() {
    m.residenceAddress = value
}
// SetRubrics sets the rubrics property value. When set, the grading rubric attached to the assignment.
func (m *EducationUser) SetRubrics(value []EducationRubricable)() {
    m.rubrics = value
}
// SetSchools sets the schools property value. Schools to which the user belongs. Nullable.
func (m *EducationUser) SetSchools(value []EducationSchoolable)() {
    m.schools = value
}
// SetShowInAddressList sets the showInAddressList property value. True if the Outlook Global Address List should contain this user; otherwise, false. If not set, this will be treated as true. For users invited through the invitation manager, this property will be set to false.
func (m *EducationUser) SetShowInAddressList(value *bool)() {
    m.showInAddressList = value
}
// SetStudent sets the student property value. If the primary role is student, this block will contain student specific data.
func (m *EducationUser) SetStudent(value EducationStudentable)() {
    m.student = value
}
// SetSurname sets the surname property value. The user's surname (family name or last name). Supports $filter.
func (m *EducationUser) SetSurname(value *string)() {
    m.surname = value
}
// SetTaughtClasses sets the taughtClasses property value. Classes for which the user is a teacher.
func (m *EducationUser) SetTaughtClasses(value []EducationClassable)() {
    m.taughtClasses = value
}
// SetTeacher sets the teacher property value. If the primary role is teacher, this block will contain teacher specific data.
func (m *EducationUser) SetTeacher(value EducationTeacherable)() {
    m.teacher = value
}
// SetUsageLocation sets the usageLocation property value. A two-letter country code (ISO standard 3166). Required for users who will be assigned licenses due to a legal requirement to check for availability of services in countries or regions. Examples include: US, JP, and GB. Not nullable. Supports $filter.
func (m *EducationUser) SetUsageLocation(value *string)() {
    m.usageLocation = value
}
// SetUser sets the user property value. The directory user that corresponds to this user.
func (m *EducationUser) SetUser(value Userable)() {
    m.user = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user. The UPN is an internet-style login name for the user based on the internet standard RFC 822. By convention, this should map to the user's email name. The general format is alias@domain, where domain must be present in the tenant's collection of verified domains. This property is required when a user is created. The verified domains for the tenant can be accessed from the verifiedDomains property of the organization. Supports $filter and $orderby.
func (m *EducationUser) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserType sets the userType property value. A string value that can be used to classify user types in your directory, such as Member and Guest. Supports $filter.
func (m *EducationUser) SetUserType(value *string)() {
    m.userType = value
}
