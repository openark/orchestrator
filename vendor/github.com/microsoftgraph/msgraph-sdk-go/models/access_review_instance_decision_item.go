package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewInstanceDecisionItem 
type AccessReviewInstanceDecisionItem struct {
    Entity
    // The identifier of the accessReviewInstance parent. Supports $select. Read-only.
    accessReviewId *string
    // The identifier of the user who applied the decision. Read-only.
    appliedBy UserIdentityable
    // The timestamp when the approval decision was applied.00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't applied the decision or it was automatically applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $select. Read-only.
    appliedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure, AppliedSuccessfullyButObjectNotFound and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only). Read-only.
    applyResult *string
    // Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and $filter (eq only).
    decision *string
    // Justification left by the reviewer when they made the decision.
    justification *string
    // Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity. Supports $select. Read-only.
    principal Identityable
    // A link to the principal object. For example, https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only.
    principalLink *string
    // A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend approve if sign-in is within thirty days of start of review. Recommend deny if sign-in is greater than thirty days of start of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports $select, $orderby, and $filter (eq only). Read-only.
    recommendation *string
    // Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource. Read-only.
    resource AccessReviewInstanceDecisionItemResourceable
    // A link to the resource. For example, https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only.
    resourceLink *string
    // The identifier of the reviewer.00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Supports $select. Read-only.
    reviewedBy UserIdentityable
    // The timestamp when the review decision occurred. Supports $select. Read-only.
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewAccessReviewInstanceDecisionItem instantiates a new accessReviewInstanceDecisionItem and sets the default values.
func NewAccessReviewInstanceDecisionItem()(*AccessReviewInstanceDecisionItem) {
    m := &AccessReviewInstanceDecisionItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewInstanceDecisionItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewInstanceDecisionItem(), nil
}
// GetAccessReviewId gets the accessReviewId property value. The identifier of the accessReviewInstance parent. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetAccessReviewId()(*string) {
    return m.accessReviewId
}
// GetAppliedBy gets the appliedBy property value. The identifier of the user who applied the decision. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetAppliedBy()(UserIdentityable) {
    return m.appliedBy
}
// GetAppliedDateTime gets the appliedDateTime property value. The timestamp when the approval decision was applied.00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't applied the decision or it was automatically applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetAppliedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.appliedDateTime
}
// GetApplyResult gets the applyResult property value. The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure, AppliedSuccessfullyButObjectNotFound and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstanceDecisionItem) GetApplyResult()(*string) {
    return m.applyResult
}
// GetDecision gets the decision property value. Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and $filter (eq only).
func (m *AccessReviewInstanceDecisionItem) GetDecision()(*string) {
    return m.decision
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewInstanceDecisionItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessReviewId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessReviewId(val)
        }
        return nil
    }
    res["appliedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedBy(val.(UserIdentityable))
        }
        return nil
    }
    res["appliedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppliedDateTime(val)
        }
        return nil
    }
    res["applyResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplyResult(val)
        }
        return nil
    }
    res["decision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDecision(val)
        }
        return nil
    }
    res["justification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustification(val)
        }
        return nil
    }
    res["principal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipal(val.(Identityable))
        }
        return nil
    }
    res["principalLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalLink(val)
        }
        return nil
    }
    res["recommendation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecommendation(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAccessReviewInstanceDecisionItemResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(AccessReviewInstanceDecisionItemResourceable))
        }
        return nil
    }
    res["resourceLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceLink(val)
        }
        return nil
    }
    res["reviewedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedBy(val.(UserIdentityable))
        }
        return nil
    }
    res["reviewedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedDateTime(val)
        }
        return nil
    }
    return res
}
// GetJustification gets the justification property value. Justification left by the reviewer when they made the decision.
func (m *AccessReviewInstanceDecisionItem) GetJustification()(*string) {
    return m.justification
}
// GetPrincipal gets the principal property value. Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetPrincipal()(Identityable) {
    return m.principal
}
// GetPrincipalLink gets the principalLink property value. A link to the principal object. For example, https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetPrincipalLink()(*string) {
    return m.principalLink
}
// GetRecommendation gets the recommendation property value. A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend approve if sign-in is within thirty days of start of review. Recommend deny if sign-in is greater than thirty days of start of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstanceDecisionItem) GetRecommendation()(*string) {
    return m.recommendation
}
// GetResource gets the resource property value. Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetResource()(AccessReviewInstanceDecisionItemResourceable) {
    return m.resource
}
// GetResourceLink gets the resourceLink property value. A link to the resource. For example, https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetResourceLink()(*string) {
    return m.resourceLink
}
// GetReviewedBy gets the reviewedBy property value. The identifier of the reviewer.00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetReviewedBy()(UserIdentityable) {
    return m.reviewedBy
}
// GetReviewedDateTime gets the reviewedDateTime property value. The timestamp when the review decision occurred. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewedDateTime
}
// Serialize serializes information the current object
func (m *AccessReviewInstanceDecisionItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accessReviewId", m.GetAccessReviewId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("appliedBy", m.GetAppliedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("appliedDateTime", m.GetAppliedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("applyResult", m.GetApplyResult())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("decision", m.GetDecision())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("justification", m.GetJustification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("principal", m.GetPrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalLink", m.GetPrincipalLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recommendation", m.GetRecommendation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceLink", m.GetResourceLink())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reviewedBy", m.GetReviewedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("reviewedDateTime", m.GetReviewedDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessReviewId sets the accessReviewId property value. The identifier of the accessReviewInstance parent. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetAccessReviewId(value *string)() {
    m.accessReviewId = value
}
// SetAppliedBy sets the appliedBy property value. The identifier of the user who applied the decision. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetAppliedBy(value UserIdentityable)() {
    m.appliedBy = value
}
// SetAppliedDateTime sets the appliedDateTime property value. The timestamp when the approval decision was applied.00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't applied the decision or it was automatically applied. The DatetimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.  Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetAppliedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.appliedDateTime = value
}
// SetApplyResult sets the applyResult property value. The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure, AppliedSuccessfullyButObjectNotFound and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstanceDecisionItem) SetApplyResult(value *string)() {
    m.applyResult = value
}
// SetDecision sets the decision property value. Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and $filter (eq only).
func (m *AccessReviewInstanceDecisionItem) SetDecision(value *string)() {
    m.decision = value
}
// SetJustification sets the justification property value. Justification left by the reviewer when they made the decision.
func (m *AccessReviewInstanceDecisionItem) SetJustification(value *string)() {
    m.justification = value
}
// SetPrincipal sets the principal property value. Every decision item in an access review represents a principal's access to a resource. This property represents details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and servicePrincipalIdentity. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetPrincipal(value Identityable)() {
    m.principal = value
}
// SetPrincipalLink sets the principalLink property value. A link to the principal object. For example, https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetPrincipalLink(value *string)() {
    m.principalLink = value
}
// SetRecommendation sets the recommendation property value. A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend approve if sign-in is within thirty days of start of review. Recommend deny if sign-in is greater than thirty days of start of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports $select, $orderby, and $filter (eq only). Read-only.
func (m *AccessReviewInstanceDecisionItem) SetRecommendation(value *string)() {
    m.recommendation = value
}
// SetResource sets the resource property value. Every decision item in an access review represents a principal's access to a resource. This property represents details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See accessReviewInstanceDecisionItemResource. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetResource(value AccessReviewInstanceDecisionItemResourceable)() {
    m.resource = value
}
// SetResourceLink sets the resourceLink property value. A link to the resource. For example, https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetResourceLink(value *string)() {
    m.resourceLink = value
}
// SetReviewedBy sets the reviewedBy property value. The identifier of the reviewer.00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetReviewedBy(value UserIdentityable)() {
    m.reviewedBy = value
}
// SetReviewedDateTime sets the reviewedDateTime property value. The timestamp when the review decision occurred. Supports $select. Read-only.
func (m *AccessReviewInstanceDecisionItem) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
