package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskDetection 
type RiskDetection struct {
    Entity
    // Indicates the activity type the detected risk is linked to. Possible values are: signin, user, unknownFutureValue.
    activity *ActivityType
    // Date and time that the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Additional information associated with the risk detection in JSON format. For example, '[{/'Key/':/'userAgent/',/'Value/':/'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36/'}]'. Possible keys in the additionalInfo JSON string are: userAgent, alertUrl, relatedEventTimeInUtc, relatedUserAgent, deviceInformation, relatedLocation, requestId, correlationId, lastActivityTimeInUtc, malwareName, clientLocation, clientIp, riskReasons. For more information about riskReasons and possible values, see riskReasons values.
    additionalInfo *string
    // Correlation ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
    correlationId *string
    // Date and time that the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: 2014-01-01T00:00:00Z
    detectedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timing of the detected risk (real-time/offline). Possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
    detectionTimingType *RiskDetectionTimingType
    // Provides the IP address of the client from where the risk occurred.
    ipAddress *string
    // Date and time that the risk detection was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Location of the sign-in.
    location SignInLocationable
    // Request ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
    requestId *string
    // Details of the detected risk. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue, m365DAdminDismissedDetection. Note that you must use the Prefer: include - unknown -enum-members request header to get the following value(s) in this evolvable enum: m365DAdminDismissedDetection.
    riskDetail *RiskDetail
    // The type of risk event detected. The possible values are unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence, generic,adminConfirmedUserCompromised, passwordSpray, impossibleTravel, newCountry, anomalousToken, tokenIssuerAnomaly,suspiciousBrowser, riskyIPAddress, mcasSuspiciousInboxManipulationRules, suspiciousInboxForwarding, and anomalousUserActivity. If the risk detection is a premium detection, will show generic. For more information about each value, see riskEventType values.
    riskEventType *string
    // Level of the detected risk. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
    riskLevel *RiskLevel
    // The state of a detected risky user or sign-in. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
    riskState *RiskState
    // Source of the risk detection. For example, activeDirectory.
    source *string
    // Indicates the type of token issuer for the detected sign-in risk. Possible values are: AzureAD, ADFederationServices, UnknownFutureValue.
    tokenIssuerType *TokenIssuerType
    // The user principal name (UPN) of the user.
    userDisplayName *string
    // Unique ID of the user.
    userId *string
    // The user principal name (UPN) of the user.
    userPrincipalName *string
}
// NewRiskDetection instantiates a new RiskDetection and sets the default values.
func NewRiskDetection()(*RiskDetection) {
    m := &RiskDetection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRiskDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRiskDetection(), nil
}
// GetActivity gets the activity property value. Indicates the activity type the detected risk is linked to. Possible values are: signin, user, unknownFutureValue.
func (m *RiskDetection) GetActivity()(*ActivityType) {
    return m.activity
}
// GetActivityDateTime gets the activityDateTime property value. Date and time that the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activityDateTime
}
// GetAdditionalInfo gets the additionalInfo property value. Additional information associated with the risk detection in JSON format. For example, '[{/'Key/':/'userAgent/',/'Value/':/'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36/'}]'. Possible keys in the additionalInfo JSON string are: userAgent, alertUrl, relatedEventTimeInUtc, relatedUserAgent, deviceInformation, relatedLocation, requestId, correlationId, lastActivityTimeInUtc, malwareName, clientLocation, clientIp, riskReasons. For more information about riskReasons and possible values, see riskReasons values.
func (m *RiskDetection) GetAdditionalInfo()(*string) {
    return m.additionalInfo
}
// GetCorrelationId gets the correlationId property value. Correlation ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
func (m *RiskDetection) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetDetectedDateTime gets the detectedDateTime property value. Date and time that the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.detectedDateTime
}
// GetDetectionTimingType gets the detectionTimingType property value. Timing of the detected risk (real-time/offline). Possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
func (m *RiskDetection) GetDetectionTimingType()(*RiskDetectionTimingType) {
    return m.detectionTimingType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseActivityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivity(val.(*ActivityType))
        }
        return nil
    }
    res["activityDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["additionalInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalInfo(val)
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
    res["detectedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectedDateTime(val)
        }
        return nil
    }
    res["detectionTimingType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskDetectionTimingType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionTimingType(val.(*RiskDetectionTimingType))
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
    res["lastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUpdatedDateTime(val)
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
    res["requestId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRequestId(val)
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
    res["riskEventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskEventType(val)
        }
        return nil
    }
    res["riskLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRiskLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLevel(val.(*RiskLevel))
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
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val)
        }
        return nil
    }
    res["tokenIssuerType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTokenIssuerType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTokenIssuerType(val.(*TokenIssuerType))
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
// GetIpAddress gets the ipAddress property value. Provides the IP address of the client from where the risk occurred.
func (m *RiskDetection) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Date and time that the risk detection was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetLocation gets the location property value. Location of the sign-in.
func (m *RiskDetection) GetLocation()(SignInLocationable) {
    return m.location
}
// GetRequestId gets the requestId property value. Request ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
func (m *RiskDetection) GetRequestId()(*string) {
    return m.requestId
}
// GetRiskDetail gets the riskDetail property value. Details of the detected risk. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue, m365DAdminDismissedDetection. Note that you must use the Prefer: include - unknown -enum-members request header to get the following value(s) in this evolvable enum: m365DAdminDismissedDetection.
func (m *RiskDetection) GetRiskDetail()(*RiskDetail) {
    return m.riskDetail
}
// GetRiskEventType gets the riskEventType property value. The type of risk event detected. The possible values are unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence, generic,adminConfirmedUserCompromised, passwordSpray, impossibleTravel, newCountry, anomalousToken, tokenIssuerAnomaly,suspiciousBrowser, riskyIPAddress, mcasSuspiciousInboxManipulationRules, suspiciousInboxForwarding, and anomalousUserActivity. If the risk detection is a premium detection, will show generic. For more information about each value, see riskEventType values.
func (m *RiskDetection) GetRiskEventType()(*string) {
    return m.riskEventType
}
// GetRiskLevel gets the riskLevel property value. Level of the detected risk. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
func (m *RiskDetection) GetRiskLevel()(*RiskLevel) {
    return m.riskLevel
}
// GetRiskState gets the riskState property value. The state of a detected risky user or sign-in. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *RiskDetection) GetRiskState()(*RiskState) {
    return m.riskState
}
// GetSource gets the source property value. Source of the risk detection. For example, activeDirectory.
func (m *RiskDetection) GetSource()(*string) {
    return m.source
}
// GetTokenIssuerType gets the tokenIssuerType property value. Indicates the type of token issuer for the detected sign-in risk. Possible values are: AzureAD, ADFederationServices, UnknownFutureValue.
func (m *RiskDetection) GetTokenIssuerType()(*TokenIssuerType) {
    return m.tokenIssuerType
}
// GetUserDisplayName gets the userDisplayName property value. The user principal name (UPN) of the user.
func (m *RiskDetection) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. Unique ID of the user.
func (m *RiskDetection) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the user.
func (m *RiskDetection) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *RiskDetection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetActivity() != nil {
        cast := (*m.GetActivity()).String()
        err = writer.WriteStringValue("activity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("additionalInfo", m.GetAdditionalInfo())
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
        err = writer.WriteTimeValue("detectedDateTime", m.GetDetectedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetDetectionTimingType() != nil {
        cast := (*m.GetDetectionTimingType()).String()
        err = writer.WriteStringValue("detectionTimingType", &cast)
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
        err = writer.WriteTimeValue("lastUpdatedDateTime", m.GetLastUpdatedDateTime())
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
        err = writer.WriteStringValue("requestId", m.GetRequestId())
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
    {
        err = writer.WriteStringValue("riskEventType", m.GetRiskEventType())
        if err != nil {
            return err
        }
    }
    if m.GetRiskLevel() != nil {
        cast := (*m.GetRiskLevel()).String()
        err = writer.WriteStringValue("riskLevel", &cast)
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
        err = writer.WriteStringValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    if m.GetTokenIssuerType() != nil {
        cast := (*m.GetTokenIssuerType()).String()
        err = writer.WriteStringValue("tokenIssuerType", &cast)
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
// SetActivity sets the activity property value. Indicates the activity type the detected risk is linked to. Possible values are: signin, user, unknownFutureValue.
func (m *RiskDetection) SetActivity(value *ActivityType)() {
    m.activity = value
}
// SetActivityDateTime sets the activityDateTime property value. Date and time that the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// SetAdditionalInfo sets the additionalInfo property value. Additional information associated with the risk detection in JSON format. For example, '[{/'Key/':/'userAgent/',/'Value/':/'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36/'}]'. Possible keys in the additionalInfo JSON string are: userAgent, alertUrl, relatedEventTimeInUtc, relatedUserAgent, deviceInformation, relatedLocation, requestId, correlationId, lastActivityTimeInUtc, malwareName, clientLocation, clientIp, riskReasons. For more information about riskReasons and possible values, see riskReasons values.
func (m *RiskDetection) SetAdditionalInfo(value *string)() {
    m.additionalInfo = value
}
// SetCorrelationId sets the correlationId property value. Correlation ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
func (m *RiskDetection) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetDetectedDateTime sets the detectedDateTime property value. Date and time that the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.detectedDateTime = value
}
// SetDetectionTimingType sets the detectionTimingType property value. Timing of the detected risk (real-time/offline). Possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
func (m *RiskDetection) SetDetectionTimingType(value *RiskDetectionTimingType)() {
    m.detectionTimingType = value
}
// SetIpAddress sets the ipAddress property value. Provides the IP address of the client from where the risk occurred.
func (m *RiskDetection) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Date and time that the risk detection was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is look like this: 2014-01-01T00:00:00Z
func (m *RiskDetection) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetLocation sets the location property value. Location of the sign-in.
func (m *RiskDetection) SetLocation(value SignInLocationable)() {
    m.location = value
}
// SetRequestId sets the requestId property value. Request ID of the sign-in associated with the risk detection. This property is null if the risk detection is not associated with a sign-in.
func (m *RiskDetection) SetRequestId(value *string)() {
    m.requestId = value
}
// SetRiskDetail sets the riskDetail property value. Details of the detected risk. The possible values are: none, adminGeneratedTemporaryPassword, userPerformedSecuredPasswordChange, userPerformedSecuredPasswordReset, adminConfirmedSigninSafe, aiConfirmedSigninSafe, userPassedMFADrivenByRiskBasedPolicy, adminDismissedAllRiskForUser, adminConfirmedSigninCompromised, hidden, adminConfirmedUserCompromised, unknownFutureValue, m365DAdminDismissedDetection. Note that you must use the Prefer: include - unknown -enum-members request header to get the following value(s) in this evolvable enum: m365DAdminDismissedDetection.
func (m *RiskDetection) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
// SetRiskEventType sets the riskEventType property value. The type of risk event detected. The possible values are unlikelyTravel, anonymizedIPAddress, maliciousIPAddress, unfamiliarFeatures, malwareInfectedIPAddress, suspiciousIPAddress, leakedCredentials, investigationsThreatIntelligence, generic,adminConfirmedUserCompromised, passwordSpray, impossibleTravel, newCountry, anomalousToken, tokenIssuerAnomaly,suspiciousBrowser, riskyIPAddress, mcasSuspiciousInboxManipulationRules, suspiciousInboxForwarding, and anomalousUserActivity. If the risk detection is a premium detection, will show generic. For more information about each value, see riskEventType values.
func (m *RiskDetection) SetRiskEventType(value *string)() {
    m.riskEventType = value
}
// SetRiskLevel sets the riskLevel property value. Level of the detected risk. Possible values are: low, medium, high, hidden, none, unknownFutureValue.
func (m *RiskDetection) SetRiskLevel(value *RiskLevel)() {
    m.riskLevel = value
}
// SetRiskState sets the riskState property value. The state of a detected risky user or sign-in. Possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *RiskDetection) SetRiskState(value *RiskState)() {
    m.riskState = value
}
// SetSource sets the source property value. Source of the risk detection. For example, activeDirectory.
func (m *RiskDetection) SetSource(value *string)() {
    m.source = value
}
// SetTokenIssuerType sets the tokenIssuerType property value. Indicates the type of token issuer for the detected sign-in risk. Possible values are: AzureAD, ADFederationServices, UnknownFutureValue.
func (m *RiskDetection) SetTokenIssuerType(value *TokenIssuerType)() {
    m.tokenIssuerType = value
}
// SetUserDisplayName sets the userDisplayName property value. The user principal name (UPN) of the user.
func (m *RiskDetection) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. Unique ID of the user.
func (m *RiskDetection) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the user.
func (m *RiskDetection) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
