package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstance 
type AccessReviewInstance struct {
    Entity
    // Returns the collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers, contactedReviewers returns their individual identities. Supports $select. Read-only.
    contactedReviewers []AccessReviewReviewerable
    // Each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
    decisions []AccessReviewInstanceDecisionItemable
    // DateTime when review instance is scheduled to end.The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. Supports $select.
    fallbackReviewers []AccessReviewReviewerScopeable
    // This collection of access review scopes is used to define who the reviewers are. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
    reviewers []AccessReviewReviewerScopeable
    // Created based on scope and instanceEnumerationScope at the accessReviewScheduleDefinition level. Defines the scope of users reviewed in a group. Supports $select and $filter (contains only). Read-only.
    scope AccessReviewScopeable
    // If the instance has multiple stages, this returns the collection of stages. A new stage will only be created when the previous stage ends. The existence, number, and settings of stages on a review instance are created based on the accessReviewStageSettings on the parent accessReviewScheduleDefinition.
    stages []AccessReviewStageable
    // DateTime when review instance is scheduled to start. May be in the future. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Specifies the status of an accessReview. Possible values: Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed. Supports $select, $orderby, and $filter (eq only). Read-only.
    status *string
}
// NewAccessReviewInstance instantiates a new AccessReviewInstance and sets the default values.
func NewAccessReviewInstance()(*AccessReviewInstance) {
    m := &AccessReviewInstance{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewInstanceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewInstance(), nil
}
// GetContactedReviewers gets the contactedReviewers property value. Returns the collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers, contactedReviewers returns their individual identities. Supports $select. Read-only.
func (m *AccessReviewInstance) GetContactedReviewers()([]AccessReviewReviewerable) {
    return m.contactedReviewers
}
// GetDecisions gets the decisions property value. Each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstance) GetDecisions()([]AccessReviewInstanceDecisionItemable) {
    return m.decisions
}
// GetEndDateTime gets the endDateTime property value. DateTime when review instance is scheduled to end.The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only.
func (m *AccessReviewInstance) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFallbackReviewers gets the fallbackReviewers property value. This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. Supports $select.
func (m *AccessReviewInstance) GetFallbackReviewers()([]AccessReviewReviewerScopeable) {
    return m.fallbackReviewers
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstance) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["contactedReviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewReviewerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewReviewerable)
            }
            m.SetContactedReviewers(res)
        }
        return nil
    }
    res["decisions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewInstanceDecisionItemable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewInstanceDecisionItemable)
            }
            m.SetDecisions(res)
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["fallbackReviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewReviewerScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerScopeable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewReviewerScopeable)
            }
            m.SetFallbackReviewers(res)
        }
        return nil
    }
    res["reviewers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewReviewerScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewReviewerScopeable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewReviewerScopeable)
            }
            m.SetReviewers(res)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewScopeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(AccessReviewScopeable))
        }
        return nil
    }
    res["stages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessReviewStageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessReviewStageable, len(val))
            for i, v := range val {
                res[i] = v.(AccessReviewStageable)
            }
            m.SetStages(res)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val)
        }
        return nil
    }
    return res
}
// GetReviewers gets the reviewers property value. This collection of access review scopes is used to define who the reviewers are. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
func (m *AccessReviewInstance) GetReviewers()([]AccessReviewReviewerScopeable) {
    return m.reviewers
}
// GetScope gets the scope property value. Created based on scope and instanceEnumerationScope at the accessReviewScheduleDefinition level. Defines the scope of users reviewed in a group. Supports $select and $filter (contains only). Read-only.
func (m *AccessReviewInstance) GetScope()(AccessReviewScopeable) {
    return m.scope
}
// GetStages gets the stages property value. If the instance has multiple stages, this returns the collection of stages. A new stage will only be created when the previous stage ends. The existence, number, and settings of stages on a review instance are created based on the accessReviewStageSettings on the parent accessReviewScheduleDefinition.
func (m *AccessReviewInstance) GetStages()([]AccessReviewStageable) {
    return m.stages
}
// GetStartDateTime gets the startDateTime property value. DateTime when review instance is scheduled to start. May be in the future. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only.
func (m *AccessReviewInstance) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetStatus gets the status property value. Specifies the status of an accessReview. Possible values: Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstance) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *AccessReviewInstance) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetContactedReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContactedReviewers()))
        for i, v := range m.GetContactedReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("contactedReviewers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDecisions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDecisions()))
        for i, v := range m.GetDecisions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("decisions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetFallbackReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFallbackReviewers()))
        for i, v := range m.GetFallbackReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("fallbackReviewers", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReviewers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReviewers()))
        for i, v := range m.GetReviewers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reviewers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("scope", m.GetScope())
        if err != nil {
            return err
        }
    }
    if m.GetStages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStages()))
        for i, v := range m.GetStages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("stages", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContactedReviewers sets the contactedReviewers property value. Returns the collection of reviewers who were contacted to complete this review. While the reviewers and fallbackReviewers properties of the accessReviewScheduleDefinition might specify group owners or managers as reviewers, contactedReviewers returns their individual identities. Supports $select. Read-only.
func (m *AccessReviewInstance) SetContactedReviewers(value []AccessReviewReviewerable)() {
    m.contactedReviewers = value
}
// SetDecisions sets the decisions property value. Each user reviewed in an accessReviewInstance has a decision item representing if they were approved, denied, or not yet reviewed.
func (m *AccessReviewInstance) SetDecisions(value []AccessReviewInstanceDecisionItemable)() {
    m.decisions = value
}
// SetEndDateTime sets the endDateTime property value. DateTime when review instance is scheduled to end.The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only.
func (m *AccessReviewInstance) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetFallbackReviewers sets the fallbackReviewers property value. This collection of reviewer scopes is used to define the list of fallback reviewers. These fallback reviewers will be notified to take action if no users are found from the list of reviewers specified. This could occur when either the group owner is specified as the reviewer but the group owner does not exist, or manager is specified as reviewer but a user's manager does not exist. Supports $select.
func (m *AccessReviewInstance) SetFallbackReviewers(value []AccessReviewReviewerScopeable)() {
    m.fallbackReviewers = value
}
// SetReviewers sets the reviewers property value. This collection of access review scopes is used to define who the reviewers are. Supports $select. For examples of options for assigning reviewers, see Assign reviewers to your access review definition using the Microsoft Graph API.
func (m *AccessReviewInstance) SetReviewers(value []AccessReviewReviewerScopeable)() {
    m.reviewers = value
}
// SetScope sets the scope property value. Created based on scope and instanceEnumerationScope at the accessReviewScheduleDefinition level. Defines the scope of users reviewed in a group. Supports $select and $filter (contains only). Read-only.
func (m *AccessReviewInstance) SetScope(value AccessReviewScopeable)() {
    m.scope = value
}
// SetStages sets the stages property value. If the instance has multiple stages, this returns the collection of stages. A new stage will only be created when the previous stage ends. The existence, number, and settings of stages on a review instance are created based on the accessReviewStageSettings on the parent accessReviewScheduleDefinition.
func (m *AccessReviewInstance) SetStages(value []AccessReviewStageable)() {
    m.stages = value
}
// SetStartDateTime sets the startDateTime property value. DateTime when review instance is scheduled to start. May be in the future. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $select. Read-only.
func (m *AccessReviewInstance) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetStatus sets the status property value. Specifies the status of an accessReview. Possible values: Initializing, NotStarted, Starting, InProgress, Completing, Completed, AutoReviewing, and AutoReviewed. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstance) SetStatus(value *string)() {
    m.status = value
}
