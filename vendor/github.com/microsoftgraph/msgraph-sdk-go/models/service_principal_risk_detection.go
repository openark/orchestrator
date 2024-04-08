package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalRiskDetection 
type ServicePrincipalRiskDetection struct {
    Entity
    // Indicates the activity type the detected risk is linked to.  The possible values are: signin, servicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: servicePrincipal.
    activity *ActivityType
    // Date and time when the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Additional information associated with the risk detection. This string value is represented as a JSON object with the quotations escaped.
    additionalInfo *string
    // The unique identifier for the associated application.
    appId *string
    // Correlation ID of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity.
    correlationId *string
    // Date and time when the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    detectedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Timing of the detected risk , whether real-time or offline. The possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
    detectionTimingType *RiskDetectionTimingType
    // Provides the IP address of the client from where the risk occurred.
    ipAddress *string
    // The unique identifier for the key credential associated with the risk detection.
    keyIds []string
    // Date and time when the risk detection was last updated.
    lastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Location from where the sign-in was initiated.
    location SignInLocationable
    // Request identifier of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity. Supports $filter (eq).
    requestId *string
    // Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: none, hidden, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
    riskDetail *RiskDetail
    // The type of risk event detected. The possible values are: investigationsThreatIntelligence, generic, adminConfirmedServicePrincipalCompromised, suspiciousSignins, leakedCredentials, anomalousServicePrincipalActivity, maliciousApplication, suspiciousApplication.
    riskEventType *string
    // Level of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: low, medium, high, hidden, none.
    riskLevel *RiskLevel
    // The state of a detected risky service principal or sign-in activity. The possible values are: none, dismissed, atRisk, confirmedCompromised.
    riskState *RiskState
    // The display name for the service principal.
    servicePrincipalDisplayName *string
    // The unique identifier for the service principal. Supports $filter (eq).
    servicePrincipalId *string
    // Source of the risk detection. For example, identityProtection.
    source *string
    // Indicates the type of token issuer for the detected sign-in risk. The possible values are: AzureAD.
    tokenIssuerType *TokenIssuerType
}
// NewServicePrincipalRiskDetection instantiates a new ServicePrincipalRiskDetection and sets the default values.
func NewServicePrincipalRiskDetection()(*ServicePrincipalRiskDetection) {
    m := &ServicePrincipalRiskDetection{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServicePrincipalRiskDetectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalRiskDetectionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalRiskDetection(), nil
}
// GetActivity gets the activity property value. Indicates the activity type the detected risk is linked to.  The possible values are: signin, servicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: servicePrincipal.
func (m *ServicePrincipalRiskDetection) GetActivity()(*ActivityType) {
    return m.activity
}
// GetActivityDateTime gets the activityDateTime property value. Date and time when the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ServicePrincipalRiskDetection) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activityDateTime
}
// GetAdditionalInfo gets the additionalInfo property value. Additional information associated with the risk detection. This string value is represented as a JSON object with the quotations escaped.
func (m *ServicePrincipalRiskDetection) GetAdditionalInfo()(*string) {
    return m.additionalInfo
}
// GetAppId gets the appId property value. The unique identifier for the associated application.
func (m *ServicePrincipalRiskDetection) GetAppId()(*string) {
    return m.appId
}
// GetCorrelationId gets the correlationId property value. Correlation ID of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity.
func (m *ServicePrincipalRiskDetection) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetDetectedDateTime gets the detectedDateTime property value. Date and time when the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ServicePrincipalRiskDetection) GetDetectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.detectedDateTime
}
// GetDetectionTimingType gets the detectionTimingType property value. Timing of the detected risk , whether real-time or offline. The possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
func (m *ServicePrincipalRiskDetection) GetDetectionTimingType()(*RiskDetectionTimingType) {
    return m.detectionTimingType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalRiskDetection) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["keyIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetKeyIds(res)
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
    res["servicePrincipalDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalDisplayName(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
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
    return res
}
// GetIpAddress gets the ipAddress property value. Provides the IP address of the client from where the risk occurred.
func (m *ServicePrincipalRiskDetection) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetKeyIds gets the keyIds property value. The unique identifier for the key credential associated with the risk detection.
func (m *ServicePrincipalRiskDetection) GetKeyIds()([]string) {
    return m.keyIds
}
// GetLastUpdatedDateTime gets the lastUpdatedDateTime property value. Date and time when the risk detection was last updated.
func (m *ServicePrincipalRiskDetection) GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastUpdatedDateTime
}
// GetLocation gets the location property value. Location from where the sign-in was initiated.
func (m *ServicePrincipalRiskDetection) GetLocation()(SignInLocationable) {
    return m.location
}
// GetRequestId gets the requestId property value. Request identifier of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) GetRequestId()(*string) {
    return m.requestId
}
// GetRiskDetail gets the riskDetail property value. Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: none, hidden, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *ServicePrincipalRiskDetection) GetRiskDetail()(*RiskDetail) {
    return m.riskDetail
}
// GetRiskEventType gets the riskEventType property value. The type of risk event detected. The possible values are: investigationsThreatIntelligence, generic, adminConfirmedServicePrincipalCompromised, suspiciousSignins, leakedCredentials, anomalousServicePrincipalActivity, maliciousApplication, suspiciousApplication.
func (m *ServicePrincipalRiskDetection) GetRiskEventType()(*string) {
    return m.riskEventType
}
// GetRiskLevel gets the riskLevel property value. Level of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: low, medium, high, hidden, none.
func (m *ServicePrincipalRiskDetection) GetRiskLevel()(*RiskLevel) {
    return m.riskLevel
}
// GetRiskState gets the riskState property value. The state of a detected risky service principal or sign-in activity. The possible values are: none, dismissed, atRisk, confirmedCompromised.
func (m *ServicePrincipalRiskDetection) GetRiskState()(*RiskState) {
    return m.riskState
}
// GetServicePrincipalDisplayName gets the servicePrincipalDisplayName property value. The display name for the service principal.
func (m *ServicePrincipalRiskDetection) GetServicePrincipalDisplayName()(*string) {
    return m.servicePrincipalDisplayName
}
// GetServicePrincipalId gets the servicePrincipalId property value. The unique identifier for the service principal. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) GetServicePrincipalId()(*string) {
    return m.servicePrincipalId
}
// GetSource gets the source property value. Source of the risk detection. For example, identityProtection.
func (m *ServicePrincipalRiskDetection) GetSource()(*string) {
    return m.source
}
// GetTokenIssuerType gets the tokenIssuerType property value. Indicates the type of token issuer for the detected sign-in risk. The possible values are: AzureAD.
func (m *ServicePrincipalRiskDetection) GetTokenIssuerType()(*TokenIssuerType) {
    return m.tokenIssuerType
}
// Serialize serializes information the current object
func (m *ServicePrincipalRiskDetection) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("appId", m.GetAppId())
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
    if m.GetKeyIds() != nil {
        err = writer.WriteCollectionOfStringValues("keyIds", m.GetKeyIds())
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
        err = writer.WriteStringValue("servicePrincipalDisplayName", m.GetServicePrincipalDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
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
    return nil
}
// SetActivity sets the activity property value. Indicates the activity type the detected risk is linked to.  The possible values are: signin, servicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: servicePrincipal.
func (m *ServicePrincipalRiskDetection) SetActivity(value *ActivityType)() {
    m.activity = value
}
// SetActivityDateTime sets the activityDateTime property value. Date and time when the risky activity occurred. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ServicePrincipalRiskDetection) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// SetAdditionalInfo sets the additionalInfo property value. Additional information associated with the risk detection. This string value is represented as a JSON object with the quotations escaped.
func (m *ServicePrincipalRiskDetection) SetAdditionalInfo(value *string)() {
    m.additionalInfo = value
}
// SetAppId sets the appId property value. The unique identifier for the associated application.
func (m *ServicePrincipalRiskDetection) SetAppId(value *string)() {
    m.appId = value
}
// SetCorrelationId sets the correlationId property value. Correlation ID of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity.
func (m *ServicePrincipalRiskDetection) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetDetectedDateTime sets the detectedDateTime property value. Date and time when the risk was detected. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ServicePrincipalRiskDetection) SetDetectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.detectedDateTime = value
}
// SetDetectionTimingType sets the detectionTimingType property value. Timing of the detected risk , whether real-time or offline. The possible values are: notDefined, realtime, nearRealtime, offline, unknownFutureValue.
func (m *ServicePrincipalRiskDetection) SetDetectionTimingType(value *RiskDetectionTimingType)() {
    m.detectionTimingType = value
}
// SetIpAddress sets the ipAddress property value. Provides the IP address of the client from where the risk occurred.
func (m *ServicePrincipalRiskDetection) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetKeyIds sets the keyIds property value. The unique identifier for the key credential associated with the risk detection.
func (m *ServicePrincipalRiskDetection) SetKeyIds(value []string)() {
    m.keyIds = value
}
// SetLastUpdatedDateTime sets the lastUpdatedDateTime property value. Date and time when the risk detection was last updated.
func (m *ServicePrincipalRiskDetection) SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastUpdatedDateTime = value
}
// SetLocation sets the location property value. Location from where the sign-in was initiated.
func (m *ServicePrincipalRiskDetection) SetLocation(value SignInLocationable)() {
    m.location = value
}
// SetRequestId sets the requestId property value. Request identifier of the sign-in activity associated with the risk detection. This property is null if the risk detection is not associated with a sign-in activity. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) SetRequestId(value *string)() {
    m.requestId = value
}
// SetRiskDetail sets the riskDetail property value. Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: none, hidden, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *ServicePrincipalRiskDetection) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
// SetRiskEventType sets the riskEventType property value. The type of risk event detected. The possible values are: investigationsThreatIntelligence, generic, adminConfirmedServicePrincipalCompromised, suspiciousSignins, leakedCredentials, anomalousServicePrincipalActivity, maliciousApplication, suspiciousApplication.
func (m *ServicePrincipalRiskDetection) SetRiskEventType(value *string)() {
    m.riskEventType = value
}
// SetRiskLevel sets the riskLevel property value. Level of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: low, medium, high, hidden, none.
func (m *ServicePrincipalRiskDetection) SetRiskLevel(value *RiskLevel)() {
    m.riskLevel = value
}
// SetRiskState sets the riskState property value. The state of a detected risky service principal or sign-in activity. The possible values are: none, dismissed, atRisk, confirmedCompromised.
func (m *ServicePrincipalRiskDetection) SetRiskState(value *RiskState)() {
    m.riskState = value
}
// SetServicePrincipalDisplayName sets the servicePrincipalDisplayName property value. The display name for the service principal.
func (m *ServicePrincipalRiskDetection) SetServicePrincipalDisplayName(value *string)() {
    m.servicePrincipalDisplayName = value
}
// SetServicePrincipalId sets the servicePrincipalId property value. The unique identifier for the service principal. Supports $filter (eq).
func (m *ServicePrincipalRiskDetection) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
// SetSource sets the source property value. Source of the risk detection. For example, identityProtection.
func (m *ServicePrincipalRiskDetection) SetSource(value *string)() {
    m.source = value
}
// SetTokenIssuerType sets the tokenIssuerType property value. Indicates the type of token issuer for the detected sign-in risk. The possible values are: AzureAD.
func (m *ServicePrincipalRiskDetection) SetTokenIssuerType(value *TokenIssuerType)() {
    m.tokenIssuerType = value
}
