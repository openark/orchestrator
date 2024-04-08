package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SignIn 
type SignIn struct {
    Entity
    // App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
    appDisplayName *string
    // Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only).
    appId *string
    // Provides a list of conditional access policies that are triggered by the corresponding sign-in activity.
    appliedConditionalAccessPolicies []AppliedConditionalAccessPolicyable
    // Identifies the client used for the sign-in activity. Modern authentication clients include Browser and modern clients. Legacy authentication clients include Exchange ActiveSync, IMAP, MAPI, SMTP, POP, and other clients. Supports $filter (eq operator only).
    clientAppUsed *string
    // Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
    conditionalAccessStatus *ConditionalAccessStatus
    // The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
    correlationId *string
    // Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
    deviceDetail DeviceDetailable
    // IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
    ipAddress *string
    // Indicates if a sign-in is interactive or not.
    isInteractive *bool
    // Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
    location SignInLocationable
    // Name of the resource the user signed into. Supports $filter (eq operator only).
    resourceDisplayName *string
    // ID of the resource that the user signed into. Supports $filter (eq operator only).
    resourceId *string
    // Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden.
    riskDetail *RiskDetail
    // Risk event types associated with the sign-in. The possible values are: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, and unknownFutureValue. Supports $filter (eq operator only).
    riskEventTypes []RiskEventType
    // The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
    riskEventTypes_v2 []string
    // Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
    riskLevelAggregated *RiskLevel
    // Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
    riskLevelDuringSignIn *RiskLevel
    // Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only).
    riskState *RiskState
    // Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
    status SignInStatusable
    // Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
    userDisplayName *string
    // ID of the user that initiated the sign-in. Supports $filter (eq operator only).
    userId *string
    // User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
    userPrincipalName *string
}
// NewSignIn instantiates a new signIn and sets the default values.
func NewSignIn()(*SignIn) {
    m := &SignIn{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSignInFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSignInFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSignIn(), nil
}
// GetAppDisplayName gets the appDisplayName property value. App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetAppDisplayName()(*string) {
    return m.appDisplayName
}
// GetAppId gets the appId property value. Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only).
func (m *SignIn) GetAppId()(*string) {
    return m.appId
}
// GetAppliedConditionalAccessPolicies gets the appliedConditionalAccessPolicies property value. Provides a list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) GetAppliedConditionalAccessPolicies()([]AppliedConditionalAccessPolicyable) {
    return m.appliedConditionalAccessPolicies
}
// GetClientAppUsed gets the clientAppUsed property value. Identifies the client used for the sign-in activity. Modern authentication clients include Browser and modern clients. Legacy authentication clients include Exchange ActiveSync, IMAP, MAPI, SMTP, POP, and other clients. Supports $filter (eq operator only).
func (m *SignIn) GetClientAppUsed()(*string) {
    return m.clientAppUsed
}
// GetConditionalAccessStatus gets the conditionalAccessStatus property value. Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetConditionalAccessStatus()(*ConditionalAccessStatus) {
    return m.conditionalAccessStatus
}
// GetCorrelationId gets the correlationId property value. The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
func (m *SignIn) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
func (m *SignIn) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDeviceDetail gets the deviceDetail property value. Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
func (m *SignIn) GetDeviceDetail()(DeviceDetailable) {
    return m.deviceDetail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SignIn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDisplayName(val)
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
    res["appliedConditionalAccessPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAppliedConditionalAccessPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AppliedConditionalAccessPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AppliedConditionalAccessPolicyable)
            }
            m.SetAppliedConditionalAccessPolicies(res)
        }
        return nil
    }
    res["clientAppUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientAppUsed(val)
        }
        return nil
    }
    res["conditionalAccessStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseConditionalAccessStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConditionalAccessStatus(val.(*ConditionalAccessStatus))
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
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
    res["deviceDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDetail(val.(DeviceDetailable))
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["isInteractive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInteractive(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(SignInLocationable))
        }
        return nil
    }
    res["resourceDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceDisplayName(val)
        }
        return nil
    }
    res["resourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceId(val)
        }
        return nil
    }
    res["riskDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetail)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskDetail(val.(*RiskDetail))
        }
        return nil
    }
    res["riskEventTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseRiskEventType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskEventType, len(val))
            for i, v := range val {
                res[i] = *(v.(*RiskEventType))
            }
            m.SetRiskEventTypes(res)
        }
        return nil
    }
    res["riskEventTypes_v2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRiskEventTypesV2(res)
        }
        return nil
    }
    res["riskLevelAggregated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevelAggregated(val.(*RiskLevel))
        }
        return nil
    }
    res["riskLevelDuringSignIn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevelDuringSignIn(val.(*RiskLevel))
        }
        return nil
    }
    res["riskState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskState(val.(*RiskState))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSignInStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(SignInStatusable))
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
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
    return res
}
// GetIpAddress gets the ipAddress property value. IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetIsInteractive gets the isInteractive property value. Indicates if a sign-in is interactive or not.
func (m *SignIn) GetIsInteractive()(*bool) {
    return m.isInteractive
}
// GetLocation gets the location property value. Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
func (m *SignIn) GetLocation()(SignInLocationable) {
    return m.location
}
// GetResourceDisplayName gets the resourceDisplayName property value. Name of the resource the user signed into. Supports $filter (eq operator only).
func (m *SignIn) GetResourceDisplayName()(*string) {
    return m.resourceDisplayName
}
// GetResourceId gets the resourceId property value. ID of the resource that the user signed into. Supports $filter (eq operator only).
func (m *SignIn) GetResourceId()(*string) {
    return m.resourceId
}
// GetRiskDetail gets the riskDetail property value. Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden.
func (m *SignIn) GetRiskDetail()(*RiskDetail) {
    return m.riskDetail
}
// GetRiskEventTypes gets the riskEventTypes property value. Risk event types associated with the sign-in. The possible values are: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, and unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetRiskEventTypes()([]RiskEventType) {
    return m.riskEventTypes
}
// GetRiskEventTypesV2 gets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetRiskEventTypesV2()([]string) {
    return m.riskEventTypes_v2
}
// GetRiskLevelAggregated gets the riskLevelAggregated property value. Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) GetRiskLevelAggregated()(*RiskLevel) {
    return m.riskLevelAggregated
}
// GetRiskLevelDuringSignIn gets the riskLevelDuringSignIn property value. Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) GetRiskLevelDuringSignIn()(*RiskLevel) {
    return m.riskLevelDuringSignIn
}
// GetRiskState gets the riskState property value. Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) GetRiskState()(*RiskState) {
    return m.riskState
}
// GetStatus gets the status property value. Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
func (m *SignIn) GetStatus()(SignInStatusable) {
    return m.status
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. ID of the user that initiated the sign-in. Supports $filter (eq operator only).
func (m *SignIn) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *SignIn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appDisplayName", m.GetAppDisplayName())
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
    if m.GetAppliedConditionalAccessPolicies() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliedConditionalAccessPolicies()))
        for i, v := range m.GetAppliedConditionalAccessPolicies() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliedConditionalAccessPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientAppUsed", m.GetClientAppUsed())
        if err != nil {
            return err
        }
    }
    if m.GetConditionalAccessStatus() != nil {
        cast := (*m.GetConditionalAccessStatus()).String()
        err = writer.WriteStringValue("conditionalAccessStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
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
        err = writer.WriteObjectValue("deviceDetail", m.GetDeviceDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInteractive", m.GetIsInteractive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceDisplayName", m.GetResourceDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceId", m.GetResourceId())
        if err != nil {
            return err
        }
    }
    if m.GetRiskDetail() != nil {
        cast := (*m.GetRiskDetail()).String()
        err = writer.WriteStringValue("riskDetail", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskEventTypes() != nil {
        err = writer.WriteCollectionOfStringValues("riskEventTypes", SerializeRiskEventType(m.GetRiskEventTypes()))
        if err != nil {
            return err
        }
    }
    if m.GetRiskEventTypesV2() != nil {
        err = writer.WriteCollectionOfStringValues("riskEventTypes_v2", m.GetRiskEventTypesV2())
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevelAggregated() != nil {
        cast := (*m.GetRiskLevelAggregated()).String()
        err = writer.WriteStringValue("riskLevelAggregated", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevelDuringSignIn() != nil {
        cast := (*m.GetRiskLevelDuringSignIn()).String()
        err = writer.WriteStringValue("riskLevelDuringSignIn", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRiskState() != nil {
        cast := (*m.GetRiskState()).String()
        err = writer.WriteStringValue("riskState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
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
    return nil
}
// SetAppDisplayName sets the appDisplayName property value. App name displayed in the Azure Portal. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetAppDisplayName(value *string)() {
    m.appDisplayName = value
}
// SetAppId sets the appId property value. Unique GUID representing the app ID in the Azure Active Directory. Supports $filter (eq operator only).
func (m *SignIn) SetAppId(value *string)() {
    m.appId = value
}
// SetAppliedConditionalAccessPolicies sets the appliedConditionalAccessPolicies property value. Provides a list of conditional access policies that are triggered by the corresponding sign-in activity.
func (m *SignIn) SetAppliedConditionalAccessPolicies(value []AppliedConditionalAccessPolicyable)() {
    m.appliedConditionalAccessPolicies = value
}
// SetClientAppUsed sets the clientAppUsed property value. Identifies the client used for the sign-in activity. Modern authentication clients include Browser and modern clients. Legacy authentication clients include Exchange ActiveSync, IMAP, MAPI, SMTP, POP, and other clients. Supports $filter (eq operator only).
func (m *SignIn) SetClientAppUsed(value *string)() {
    m.clientAppUsed = value
}
// SetConditionalAccessStatus sets the conditionalAccessStatus property value. Reports status of an activated conditional access policy. Possible values are: success, failure, notApplied, and unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) SetConditionalAccessStatus(value *ConditionalAccessStatus)() {
    m.conditionalAccessStatus = value
}
// SetCorrelationId sets the correlationId property value. The request ID sent from the client when the sign-in is initiated; used to troubleshoot sign-in activity. Supports $filter (eq operator only).
func (m *SignIn) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time (UTC) the sign-in was initiated. Example: midnight on Jan 1, 2014 is reported as 2014-01-01T00:00:00Z. Supports $orderby and $filter (eq, le, and ge operators only).
func (m *SignIn) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDeviceDetail sets the deviceDetail property value. Device information from where the sign-in occurred; includes device ID, operating system, and browser. Supports $filter (eq and startsWith operators only) on browser and operatingSytem properties.
func (m *SignIn) SetDeviceDetail(value DeviceDetailable)() {
    m.deviceDetail = value
}
// SetIpAddress sets the ipAddress property value. IP address of the client used to sign in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetIsInteractive sets the isInteractive property value. Indicates if a sign-in is interactive or not.
func (m *SignIn) SetIsInteractive(value *bool)() {
    m.isInteractive = value
}
// SetLocation sets the location property value. Provides the city, state, and country code where the sign-in originated. Supports $filter (eq and startsWith operators only) on city, state, and countryOrRegion properties.
func (m *SignIn) SetLocation(value SignInLocationable)() {
    m.location = value
}
// SetResourceDisplayName sets the resourceDisplayName property value. Name of the resource the user signed into. Supports $filter (eq operator only).
func (m *SignIn) SetResourceDisplayName(value *string)() {
    m.resourceDisplayName = value
}
// SetResourceId sets the resourceId property value. ID of the resource that the user signed into. Supports $filter (eq operator only).
func (m *SignIn) SetResourceId(value *string)() {
    m.resourceId = value
}
// SetRiskDetail sets the riskDetail property value. Provides the 'reason' behind a specific state of a risky user, sign-in or a risk event. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, unknownFutureValue. The value none means that no action has been performed on the user or sign-in so far.  Supports $filter (eq operator only).Note: Details for this property require an Azure AD Premium P2 license. Other licenses return the value hidden.
func (m *SignIn) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
// SetRiskEventTypes sets the riskEventTypes property value. Risk event types associated with the sign-in. The possible values are: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, and unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) SetRiskEventTypes(value []RiskEventType)() {
    m.riskEventTypes = value
}
// SetRiskEventTypesV2 sets the riskEventTypes_v2 property value. The list of risk event types associated with the sign-in. Possible values: unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence,  generic, or unknownFutureValue. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetRiskEventTypesV2(value []string)() {
    m.riskEventTypes_v2 = value
}
// SetRiskLevelAggregated sets the riskLevelAggregated property value. Aggregated risk level. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection. Supports $filter (eq operator only).  Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) SetRiskLevelAggregated(value *RiskLevel)() {
    m.riskLevelAggregated = value
}
// SetRiskLevelDuringSignIn sets the riskLevelDuringSignIn property value. Risk level during sign-in. The possible values are: none, low, medium, high, hidden, and unknownFutureValue. The value hidden means the user or sign-in was not enabled for Azure AD Identity Protection.  Supports $filter (eq operator only). Note: Details for this property are only available for Azure AD Premium P2 customers. All other customers will be returned hidden.
func (m *SignIn) SetRiskLevelDuringSignIn(value *RiskLevel)() {
    m.riskLevelDuringSignIn = value
}
// SetRiskState sets the riskState property value. Reports status of the risky user, sign-in, or a risk event. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue. Supports $filter (eq operator only).
func (m *SignIn) SetRiskState(value *RiskState)() {
    m.riskState = value
}
// SetStatus sets the status property value. Sign-in status. Includes the error code and description of the error (in case of a sign-in failure). Supports $filter (eq operator only) on errorCode property.
func (m *SignIn) SetStatus(value SignInStatusable)() {
    m.status = value
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. ID of the user that initiated the sign-in. Supports $filter (eq operator only).
func (m *SignIn) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User principal name of the user that initiated the sign-in. Supports $filter (eq and startsWith operators only).
func (m *SignIn) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
