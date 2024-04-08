package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewHistoryDefinition 
type AccessReviewHistoryDefinition struct {
    Entity
    // The createdBy property
    createdBy UserIdentityable
    // Timestamp when the access review definition was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Determines which review decisions will be included in the fetched review history data if specified. Optional on create. All decisions will be included by default if no decisions are provided on create. Possible values are: approve, deny, dontKnow, notReviewed, and notNotified.
    decisions []AccessReviewHistoryDecisionFilter
    // Name for the access review history data collection. Required.
    displayName *string
    // If the accessReviewHistoryDefinition is a recurring definition, instances represent each recurrence. A definition that does not recur will have exactly one instance.
    instances []AccessReviewHistoryInstanceable
    // A timestamp. Reviews ending on or before this date will be included in the fetched history data. Only required if scheduleSettings is not defined.
    reviewHistoryPeriodEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A timestamp. Reviews starting on or before this date will be included in the fetched history data. Only required if scheduleSettings is not defined.
    reviewHistoryPeriodStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The settings for a recurring access review history definition series. Only required if reviewHistoryPeriodStartDateTime or reviewHistoryPeriodEndDateTime are not defined. Not supported yet.
    scheduleSettings AccessReviewHistoryScheduleSettingsable
    // Used to scope what reviews are included in the fetched history data. Fetches reviews whose scope matches with this provided scope. Required.
    scopes []AccessReviewScopeable
    // Represents the status of the review history data collection. The possible values are: done, inProgress, error, requested, unknownFutureValue.
    status *AccessReviewHistoryStatus
}
// NewAccessReviewHistoryDefinition instantiates a new accessReviewHistoryDefinition and sets the default values.
func NewAccessReviewHistoryDefinition()(*AccessReviewHistoryDefinition) {
    m := &AccessReviewHistoryDefinition{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewHistoryDefinitionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewHistoryDefinitionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewHistoryDefinition(), nil
}
// GetCreatedBy gets the createdBy property value. The createdBy property
func (m *AccessReviewHistoryDefinition) GetCreatedBy()(UserIdentityable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp when the access review definition was created.
func (m *AccessReviewHistoryDefinition) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDecisions gets the decisions property value. Determines which review decisions will be included in the fetched review history data if specified. Optional on create. All decisions will be included by default if no decisions are provided on create. Possible values are: approve, deny, dontKnow, notReviewed, and notNotified.
func (m *AccessReviewHistoryDefinition) GetDecisions()([]AccessReviewHistoryDecisionFilter) {
    return m.decisions
}
// GetDisplayName gets the displayName property value. Name for the access review history data collection. Required.
func (m *AccessReviewHistoryDefinition) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewHistoryDefinition) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(UserIdentityable))
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
    res["decisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAccessReviewHistoryDecisionFilter)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewHistoryDecisionFilter, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessReviewHistoryDecisionFilter))
            }
            m.SetDecisions(res)
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
    res["instances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewHistoryInstanceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewHistoryInstanceable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewHistoryInstanceable)
            }
            m.SetInstances(res)
        }
        return nil
    }
    res["reviewHistoryPeriodEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewHistoryPeriodEndDateTime(val)
        }
        return nil
    }
    res["reviewHistoryPeriodStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewHistoryPeriodStartDateTime(val)
        }
        return nil
    }
    res["scheduleSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewHistoryScheduleSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduleSettings(val.(AccessReviewHistoryScheduleSettingsable))
        }
        return nil
    }
    res["scopes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewScopeable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewScopeable)
            }
            m.SetScopes(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessReviewHistoryStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AccessReviewHistoryStatus))
        }
        return nil
    }
    return res
}
// GetInstances gets the instances property value. If the accessReviewHistoryDefinition is a recurring definition, instances represent each recurrence. A definition that does not recur will have exactly one instance.
func (m *AccessReviewHistoryDefinition) GetInstances()([]AccessReviewHistoryInstanceable) {
    return m.instances
}
// GetReviewHistoryPeriodEndDateTime gets the reviewHistoryPeriodEndDateTime property value. A timestamp. Reviews ending on or before this date will be included in the fetched history data. Only required if scheduleSettings is not defined.
func (m *AccessReviewHistoryDefinition) GetReviewHistoryPeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewHistoryPeriodEndDateTime
}
// GetReviewHistoryPeriodStartDateTime gets the reviewHistoryPeriodStartDateTime property value. A timestamp. Reviews starting on or before this date will be included in the fetched history data. Only required if scheduleSettings is not defined.
func (m *AccessReviewHistoryDefinition) GetReviewHistoryPeriodStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewHistoryPeriodStartDateTime
}
// GetScheduleSettings gets the scheduleSettings property value. The settings for a recurring access review history definition series. Only required if reviewHistoryPeriodStartDateTime or reviewHistoryPeriodEndDateTime are not defined. Not supported yet.
func (m *AccessReviewHistoryDefinition) GetScheduleSettings()(AccessReviewHistoryScheduleSettingsable) {
    return m.scheduleSettings
}
// GetScopes gets the scopes property value. Used to scope what reviews are included in the fetched history data. Fetches reviews whose scope matches with this provided scope. Required.
func (m *AccessReviewHistoryDefinition) GetScopes()([]AccessReviewScopeable) {
    return m.scopes
}
// GetStatus gets the status property value. Represents the status of the review history data collection. The possible values are: done, inProgress, error, requested, unknownFutureValue.
func (m *AccessReviewHistoryDefinition) GetStatus()(*AccessReviewHistoryStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *AccessReviewHistoryDefinition) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
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
    if m.GetDecisions() != nil {
        err = writer.WriteCollectionOfStringValues("decisions", SerializeAccessReviewHistoryDecisionFilter(m.GetDecisions()))
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
    if m.GetInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewHistoryPeriodEndDateTime", m.GetReviewHistoryPeriodEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewHistoryPeriodStartDateTime", m.GetReviewHistoryPeriodStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scheduleSettings", m.GetScheduleSettings())
        if err != nil {
            return err
        }
    }
    if m.GetScopes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScopes()))
        for i, v := range m.GetScopes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("scopes", cast)
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. The createdBy property
func (m *AccessReviewHistoryDefinition) SetCreatedBy(value UserIdentityable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp when the access review definition was created.
func (m *AccessReviewHistoryDefinition) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDecisions sets the decisions property value. Determines which review decisions will be included in the fetched review history data if specified. Optional on create. All decisions will be included by default if no decisions are provided on create. Possible values are: approve, deny, dontKnow, notReviewed, and notNotified.
func (m *AccessReviewHistoryDefinition) SetDecisions(value []AccessReviewHistoryDecisionFilter)() {
    m.decisions = value
}
// SetDisplayName sets the displayName property value. Name for the access review history data collection. Required.
func (m *AccessReviewHistoryDefinition) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetInstances sets the instances property value. If the accessReviewHistoryDefinition is a recurring definition, instances represent each recurrence. A definition that does not recur will have exactly one instance.
func (m *AccessReviewHistoryDefinition) SetInstances(value []AccessReviewHistoryInstanceable)() {
    m.instances = value
}
// SetReviewHistoryPeriodEndDateTime sets the reviewHistoryPeriodEndDateTime property value. A timestamp. Reviews ending on or before this date will be included in the fetched history data. Only required if scheduleSettings is not defined.
func (m *AccessReviewHistoryDefinition) SetReviewHistoryPeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewHistoryPeriodEndDateTime = value
}
// SetReviewHistoryPeriodStartDateTime sets the reviewHistoryPeriodStartDateTime property value. A timestamp. Reviews starting on or before this date will be included in the fetched history data. Only required if scheduleSettings is not defined.
func (m *AccessReviewHistoryDefinition) SetReviewHistoryPeriodStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewHistoryPeriodStartDateTime = value
}
// SetScheduleSettings sets the scheduleSettings property value. The settings for a recurring access review history definition series. Only required if reviewHistoryPeriodStartDateTime or reviewHistoryPeriodEndDateTime are not defined. Not supported yet.
func (m *AccessReviewHistoryDefinition) SetScheduleSettings(value AccessReviewHistoryScheduleSettingsable)() {
    m.scheduleSettings = value
}
// SetScopes sets the scopes property value. Used to scope what reviews are included in the fetched history data. Fetches reviews whose scope matches with this provided scope. Required.
func (m *AccessReviewHistoryDefinition) SetScopes(value []AccessReviewScopeable)() {
    m.scopes = value
}
// SetStatus sets the status property value. Represents the status of the review history data collection. The possible values are: done, inProgress, error, requested, unknownFutureValue.
func (m *AccessReviewHistoryDefinition) SetStatus(value *AccessReviewHistoryStatus)() {
    m.status = value
}
