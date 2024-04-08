package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageAssignmentPolicy 
type AccessPackageAssignmentPolicy struct {
    Entity
    // Access package containing this policy. Read-only.
    accessPackage AccessPackageable
    // Principals that can be assigned the access package through this policy. The possible values are: notSpecified, specificDirectoryUsers, specificConnectedOrganizationUsers, specificDirectoryServicePrincipals, allMemberUsers, allDirectoryUsers, allDirectoryServicePrincipals, allConfiguredConnectedOrganizationUsers, allExternalUsers, unknownFutureValue.
    allowedTargetScope *AllowedTargetScope
    // This property is only present for an auto assignment policy; if absent, this is a request-based policy.
    automaticRequestSettings AccessPackageAutomaticRequestSettingsable
    // Catalog of the access package containing this policy. Read-only.
    catalog AccessPackageCatalogable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The description of the policy.
    description *string
    // The display name of the policy.
    displayName *string
    // The expiration date for assignments created in this policy.
    expiration ExpirationPatternable
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    modifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The questions property
    questions []AccessPackageQuestionable
    // Specifies the settings for approval of requests for an access package assignment through this policy. For example, if approval is required for new requests.
    requestApprovalSettings AccessPackageAssignmentApprovalSettingsable
    // Provides additional settings to select who can create a request for an access package assignment through this policy, and what they can include in their request.
    requestorSettings AccessPackageAssignmentRequestorSettingsable
    // Settings for access reviews of assignments through this policy.
    reviewSettings AccessPackageAssignmentReviewSettingsable
    // The principals that can be assigned access from an access package through this policy.
    specificAllowedTargets []SubjectSetable
}
// NewAccessPackageAssignmentPolicy instantiates a new accessPackageAssignmentPolicy and sets the default values.
func NewAccessPackageAssignmentPolicy()(*AccessPackageAssignmentPolicy) {
    m := &AccessPackageAssignmentPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageAssignmentPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageAssignmentPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageAssignmentPolicy(), nil
}
// GetAccessPackage gets the accessPackage property value. Access package containing this policy. Read-only.
func (m *AccessPackageAssignmentPolicy) GetAccessPackage()(AccessPackageable) {
    return m.accessPackage
}
// GetAllowedTargetScope gets the allowedTargetScope property value. Principals that can be assigned the access package through this policy. The possible values are: notSpecified, specificDirectoryUsers, specificConnectedOrganizationUsers, specificDirectoryServicePrincipals, allMemberUsers, allDirectoryUsers, allDirectoryServicePrincipals, allConfiguredConnectedOrganizationUsers, allExternalUsers, unknownFutureValue.
func (m *AccessPackageAssignmentPolicy) GetAllowedTargetScope()(*AllowedTargetScope) {
    return m.allowedTargetScope
}
// GetAutomaticRequestSettings gets the automaticRequestSettings property value. This property is only present for an auto assignment policy; if absent, this is a request-based policy.
func (m *AccessPackageAssignmentPolicy) GetAutomaticRequestSettings()(AccessPackageAutomaticRequestSettingsable) {
    return m.automaticRequestSettings
}
// GetCatalog gets the catalog property value. Catalog of the access package containing this policy. Read-only.
func (m *AccessPackageAssignmentPolicy) GetCatalog()(AccessPackageCatalogable) {
    return m.catalog
}
// GetCreatedDateTime gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageAssignmentPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name of the policy.
func (m *AccessPackageAssignmentPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetExpiration gets the expiration property value. The expiration date for assignments created in this policy.
func (m *AccessPackageAssignmentPolicy) GetExpiration()(ExpirationPatternable) {
    return m.expiration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageAssignmentPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessPackage(val.(AccessPackageable))
        }
        return nil
    }
    res["allowedTargetScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAllowedTargetScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedTargetScope(val.(*AllowedTargetScope))
        }
        return nil
    }
    res["automaticRequestSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAutomaticRequestSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticRequestSettings(val.(AccessPackageAutomaticRequestSettingsable))
        }
        return nil
    }
    res["catalog"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCatalog(val.(AccessPackageCatalogable))
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
    res["expiration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateExpirationPatternFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpiration(val.(ExpirationPatternable))
        }
        return nil
    }
    res["modifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModifiedDateTime(val)
        }
        return nil
    }
    res["questions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageQuestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageQuestionable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageQuestionable)
            }
            m.SetQuestions(res)
        }
        return nil
    }
    res["requestApprovalSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentApprovalSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestApprovalSettings(val.(AccessPackageAssignmentApprovalSettingsable))
        }
        return nil
    }
    res["requestorSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentRequestorSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestorSettings(val.(AccessPackageAssignmentRequestorSettingsable))
        }
        return nil
    }
    res["reviewSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessPackageAssignmentReviewSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewSettings(val.(AccessPackageAssignmentReviewSettingsable))
        }
        return nil
    }
    res["specificAllowedTargets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetSpecificAllowedTargets(res)
        }
        return nil
    }
    return res
}
// GetModifiedDateTime gets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageAssignmentPolicy) GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.modifiedDateTime
}
// GetQuestions gets the questions property value. The questions property
func (m *AccessPackageAssignmentPolicy) GetQuestions()([]AccessPackageQuestionable) {
    return m.questions
}
// GetRequestApprovalSettings gets the requestApprovalSettings property value. Specifies the settings for approval of requests for an access package assignment through this policy. For example, if approval is required for new requests.
func (m *AccessPackageAssignmentPolicy) GetRequestApprovalSettings()(AccessPackageAssignmentApprovalSettingsable) {
    return m.requestApprovalSettings
}
// GetRequestorSettings gets the requestorSettings property value. Provides additional settings to select who can create a request for an access package assignment through this policy, and what they can include in their request.
func (m *AccessPackageAssignmentPolicy) GetRequestorSettings()(AccessPackageAssignmentRequestorSettingsable) {
    return m.requestorSettings
}
// GetReviewSettings gets the reviewSettings property value. Settings for access reviews of assignments through this policy.
func (m *AccessPackageAssignmentPolicy) GetReviewSettings()(AccessPackageAssignmentReviewSettingsable) {
    return m.reviewSettings
}
// GetSpecificAllowedTargets gets the specificAllowedTargets property value. The principals that can be assigned access from an access package through this policy.
func (m *AccessPackageAssignmentPolicy) GetSpecificAllowedTargets()([]SubjectSetable) {
    return m.specificAllowedTargets
}
// Serialize serializes information the current object
func (m *AccessPackageAssignmentPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessPackage", m.GetAccessPackage())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedTargetScope() != nil {
        cast := (*m.GetAllowedTargetScope()).String()
        err = writer.WriteStringValue("allowedTargetScope", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("automaticRequestSettings", m.GetAutomaticRequestSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("catalog", m.GetCatalog())
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
        err = writer.WriteStringValue("description", m.GetDescription())
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
    {
        err = writer.WriteObjectValue("expiration", m.GetExpiration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("modifiedDateTime", m.GetModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetQuestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetQuestions()))
        for i, v := range m.GetQuestions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("questions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestApprovalSettings", m.GetRequestApprovalSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("requestorSettings", m.GetRequestorSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewSettings", m.GetReviewSettings())
        if err != nil {
            return err
        }
    }
    if m.GetSpecificAllowedTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSpecificAllowedTargets()))
        for i, v := range m.GetSpecificAllowedTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("specificAllowedTargets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackage sets the accessPackage property value. Access package containing this policy. Read-only.
func (m *AccessPackageAssignmentPolicy) SetAccessPackage(value AccessPackageable)() {
    m.accessPackage = value
}
// SetAllowedTargetScope sets the allowedTargetScope property value. Principals that can be assigned the access package through this policy. The possible values are: notSpecified, specificDirectoryUsers, specificConnectedOrganizationUsers, specificDirectoryServicePrincipals, allMemberUsers, allDirectoryUsers, allDirectoryServicePrincipals, allConfiguredConnectedOrganizationUsers, allExternalUsers, unknownFutureValue.
func (m *AccessPackageAssignmentPolicy) SetAllowedTargetScope(value *AllowedTargetScope)() {
    m.allowedTargetScope = value
}
// SetAutomaticRequestSettings sets the automaticRequestSettings property value. This property is only present for an auto assignment policy; if absent, this is a request-based policy.
func (m *AccessPackageAssignmentPolicy) SetAutomaticRequestSettings(value AccessPackageAutomaticRequestSettingsable)() {
    m.automaticRequestSettings = value
}
// SetCatalog sets the catalog property value. Catalog of the access package containing this policy. Read-only.
func (m *AccessPackageAssignmentPolicy) SetCatalog(value AccessPackageCatalogable)() {
    m.catalog = value
}
// SetCreatedDateTime sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageAssignmentPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The description of the policy.
func (m *AccessPackageAssignmentPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name of the policy.
func (m *AccessPackageAssignmentPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExpiration sets the expiration property value. The expiration date for assignments created in this policy.
func (m *AccessPackageAssignmentPolicy) SetExpiration(value ExpirationPatternable)() {
    m.expiration = value
}
// SetModifiedDateTime sets the modifiedDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AccessPackageAssignmentPolicy) SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.modifiedDateTime = value
}
// SetQuestions sets the questions property value. The questions property
func (m *AccessPackageAssignmentPolicy) SetQuestions(value []AccessPackageQuestionable)() {
    m.questions = value
}
// SetRequestApprovalSettings sets the requestApprovalSettings property value. Specifies the settings for approval of requests for an access package assignment through this policy. For example, if approval is required for new requests.
func (m *AccessPackageAssignmentPolicy) SetRequestApprovalSettings(value AccessPackageAssignmentApprovalSettingsable)() {
    m.requestApprovalSettings = value
}
// SetRequestorSettings sets the requestorSettings property value. Provides additional settings to select who can create a request for an access package assignment through this policy, and what they can include in their request.
func (m *AccessPackageAssignmentPolicy) SetRequestorSettings(value AccessPackageAssignmentRequestorSettingsable)() {
    m.requestorSettings = value
}
// SetReviewSettings sets the reviewSettings property value. Settings for access reviews of assignments through this policy.
func (m *AccessPackageAssignmentPolicy) SetReviewSettings(value AccessPackageAssignmentReviewSettingsable)() {
    m.reviewSettings = value
}
// SetSpecificAllowedTargets sets the specificAllowedTargets property value. The principals that can be assigned access from an access package through this policy.
func (m *AccessPackageAssignmentPolicy) SetSpecificAllowedTargets(value []SubjectSetable)() {
    m.specificAllowedTargets = value
}
