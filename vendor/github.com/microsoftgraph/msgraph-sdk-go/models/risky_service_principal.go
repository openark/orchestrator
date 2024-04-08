package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RiskyServicePrincipal 
type RiskyServicePrincipal struct {
    Entity
    // The globally unique identifier for the associated application (its appId property), if any.
    appId *string
    // The display name for the service principal.
    displayName *string
    // Represents the risk history of Azure AD service principals.
    history []RiskyServicePrincipalHistoryItemable
    // true if the service principal account is enabled; otherwise, false.
    isEnabled *bool
    // Indicates whether Azure AD is currently processing the service principal's risky state.
    isProcessing *bool
    // Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: none, hidden,  unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
    riskDetail *RiskDetail
    // The date and time that the risk state was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2021 is 2021-01-01T00:00:00Z. Supports $filter (eq).
    riskLastUpdatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Level of the detected risky workload identity. The possible values are: low, medium, high, hidden, none, unknownFutureValue. Supports $filter (eq).
    riskLevel *RiskLevel
    // State of the service principal's risk. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
    riskState *RiskState
    // Identifies whether the service principal represents an Application, a ManagedIdentity, or a legacy application (socialIdp). This is set by Azure AD internally and is inherited from servicePrincipal.
    servicePrincipalType *string
}
// NewRiskyServicePrincipal instantiates a new riskyServicePrincipal and sets the default values.
func NewRiskyServicePrincipal()(*RiskyServicePrincipal) {
    m := &RiskyServicePrincipal{
        Entity: *NewEntity(),
    }
    return m
}
// CreateRiskyServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRiskyServicePrincipalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.riskyServicePrincipalHistoryItem":
                        return NewRiskyServicePrincipalHistoryItem(), nil
                }
            }
        }
    }
    return NewRiskyServicePrincipal(), nil
}
// GetAppId gets the appId property value. The globally unique identifier for the associated application (its appId property), if any.
func (m *RiskyServicePrincipal) GetAppId()(*string) {
    return m.appId
}
// GetDisplayName gets the displayName property value. The display name for the service principal.
func (m *RiskyServicePrincipal) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RiskyServicePrincipal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["history"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRiskyServicePrincipalHistoryItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RiskyServicePrincipalHistoryItemable, len(val))
            for i, v := range val {
                res[i] = v.(RiskyServicePrincipalHistoryItemable)
            }
            m.SetHistory(res)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["isProcessing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsProcessing(val)
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
    res["riskLastUpdatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskLastUpdatedDateTime(val)
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
    res["servicePrincipalType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalType(val)
        }
        return nil
    }
    return res
}
// GetHistory gets the history property value. Represents the risk history of Azure AD service principals.
func (m *RiskyServicePrincipal) GetHistory()([]RiskyServicePrincipalHistoryItemable) {
    return m.history
}
// GetIsEnabled gets the isEnabled property value. true if the service principal account is enabled; otherwise, false.
func (m *RiskyServicePrincipal) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetIsProcessing gets the isProcessing property value. Indicates whether Azure AD is currently processing the service principal's risky state.
func (m *RiskyServicePrincipal) GetIsProcessing()(*bool) {
    return m.isProcessing
}
// GetRiskDetail gets the riskDetail property value. Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: none, hidden,  unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *RiskyServicePrincipal) GetRiskDetail()(*RiskDetail) {
    return m.riskDetail
}
// GetRiskLastUpdatedDateTime gets the riskLastUpdatedDateTime property value. The date and time that the risk state was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2021 is 2021-01-01T00:00:00Z. Supports $filter (eq).
func (m *RiskyServicePrincipal) GetRiskLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.riskLastUpdatedDateTime
}
// GetRiskLevel gets the riskLevel property value. Level of the detected risky workload identity. The possible values are: low, medium, high, hidden, none, unknownFutureValue. Supports $filter (eq).
func (m *RiskyServicePrincipal) GetRiskLevel()(*RiskLevel) {
    return m.riskLevel
}
// GetRiskState gets the riskState property value. State of the service principal's risk. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *RiskyServicePrincipal) GetRiskState()(*RiskState) {
    return m.riskState
}
// GetServicePrincipalType gets the servicePrincipalType property value. Identifies whether the service principal represents an Application, a ManagedIdentity, or a legacy application (socialIdp). This is set by Azure AD internally and is inherited from servicePrincipal.
func (m *RiskyServicePrincipal) GetServicePrincipalType()(*string) {
    return m.servicePrincipalType
}
// Serialize serializes information the current object
func (m *RiskyServicePrincipal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("appId", m.GetAppId())
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
    if m.GetHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHistory()))
        for i, v := range m.GetHistory() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("history", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isProcessing", m.GetIsProcessing())
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
        err = writer.WriteTimeValue("riskLastUpdatedDateTime", m.GetRiskLastUpdatedDateTime())
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
        err = writer.WriteStringValue("servicePrincipalType", m.GetServicePrincipalType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppId sets the appId property value. The globally unique identifier for the associated application (its appId property), if any.
func (m *RiskyServicePrincipal) SetAppId(value *string)() {
    m.appId = value
}
// SetDisplayName sets the displayName property value. The display name for the service principal.
func (m *RiskyServicePrincipal) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetHistory sets the history property value. Represents the risk history of Azure AD service principals.
func (m *RiskyServicePrincipal) SetHistory(value []RiskyServicePrincipalHistoryItemable)() {
    m.history = value
}
// SetIsEnabled sets the isEnabled property value. true if the service principal account is enabled; otherwise, false.
func (m *RiskyServicePrincipal) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetIsProcessing sets the isProcessing property value. Indicates whether Azure AD is currently processing the service principal's risky state.
func (m *RiskyServicePrincipal) SetIsProcessing(value *bool)() {
    m.isProcessing = value
}
// SetRiskDetail sets the riskDetail property value. Details of the detected risk. Note: Details for this property are only available for Workload Identities Premium customers. Events in tenants without this license will be returned hidden. The possible values are: none, hidden,  unknownFutureValue, adminConfirmedServicePrincipalCompromised, adminDismissedAllRiskForServicePrincipal. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: adminConfirmedServicePrincipalCompromised , adminDismissedAllRiskForServicePrincipal.
func (m *RiskyServicePrincipal) SetRiskDetail(value *RiskDetail)() {
    m.riskDetail = value
}
// SetRiskLastUpdatedDateTime sets the riskLastUpdatedDateTime property value. The date and time that the risk state was last updated. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2021 is 2021-01-01T00:00:00Z. Supports $filter (eq).
func (m *RiskyServicePrincipal) SetRiskLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.riskLastUpdatedDateTime = value
}
// SetRiskLevel sets the riskLevel property value. Level of the detected risky workload identity. The possible values are: low, medium, high, hidden, none, unknownFutureValue. Supports $filter (eq).
func (m *RiskyServicePrincipal) SetRiskLevel(value *RiskLevel)() {
    m.riskLevel = value
}
// SetRiskState sets the riskState property value. State of the service principal's risk. The possible values are: none, confirmedSafe, remediated, dismissed, atRisk, confirmedCompromised, unknownFutureValue.
func (m *RiskyServicePrincipal) SetRiskState(value *RiskState)() {
    m.riskState = value
}
// SetServicePrincipalType sets the servicePrincipalType property value. Identifies whether the service principal represents an Application, a ManagedIdentity, or a legacy application (socialIdp). This is set by Azure AD internally and is inherited from servicePrincipal.
func (m *RiskyServicePrincipal) SetServicePrincipalType(value *string)() {
    m.servicePrincipalType = value
}
