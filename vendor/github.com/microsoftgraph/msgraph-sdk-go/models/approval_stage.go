package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApprovalStage 
type ApprovalStage struct {
    Entity
    // Indicates whether the stage is assigned to the calling user to review. Read-only.
    assignedToMe *bool
    // The label provided by the policy creator to identify an approval stage. Read-only.
    displayName *string
    // The justification associated with the approval stage decision.
    justification *string
    // The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Read-only.
    reviewedBy Identityable
    // The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
    reviewedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
    reviewResult *string
    // The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
    status *string
}
// NewApprovalStage instantiates a new approvalStage and sets the default values.
func NewApprovalStage()(*ApprovalStage) {
    m := &ApprovalStage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateApprovalStageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApprovalStageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewApprovalStage(), nil
}
// GetAssignedToMe gets the assignedToMe property value. Indicates whether the stage is assigned to the calling user to review. Read-only.
func (m *ApprovalStage) GetAssignedToMe()(*bool) {
    return m.assignedToMe
}
// GetDisplayName gets the displayName property value. The label provided by the policy creator to identify an approval stage. Read-only.
func (m *ApprovalStage) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApprovalStage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedToMe"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedToMe(val)
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
    res["reviewedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewedBy(val.(Identityable))
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
    res["reviewResult"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReviewResult(val)
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
// GetJustification gets the justification property value. The justification associated with the approval stage decision.
func (m *ApprovalStage) GetJustification()(*string) {
    return m.justification
}
// GetReviewedBy gets the reviewedBy property value. The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Read-only.
func (m *ApprovalStage) GetReviewedBy()(Identityable) {
    return m.reviewedBy
}
// GetReviewedDateTime gets the reviewedDateTime property value. The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ApprovalStage) GetReviewedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.reviewedDateTime
}
// GetReviewResult gets the reviewResult property value. The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
func (m *ApprovalStage) GetReviewResult()(*string) {
    return m.reviewResult
}
// GetStatus gets the status property value. The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
func (m *ApprovalStage) GetStatus()(*string) {
    return m.status
}
// Serialize serializes information the current object
func (m *ApprovalStage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("assignedToMe", m.GetAssignedToMe())
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
        err = writer.WriteStringValue("justification", m.GetJustification())
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
    {
        err = writer.WriteStringValue("reviewResult", m.GetReviewResult())
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
// SetAssignedToMe sets the assignedToMe property value. Indicates whether the stage is assigned to the calling user to review. Read-only.
func (m *ApprovalStage) SetAssignedToMe(value *bool)() {
    m.assignedToMe = value
}
// SetDisplayName sets the displayName property value. The label provided by the policy creator to identify an approval stage. Read-only.
func (m *ApprovalStage) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetJustification sets the justification property value. The justification associated with the approval stage decision.
func (m *ApprovalStage) SetJustification(value *string)() {
    m.justification = value
}
// SetReviewedBy sets the reviewedBy property value. The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed. Read-only.
func (m *ApprovalStage) SetReviewedBy(value Identityable)() {
    m.reviewedBy = value
}
// SetReviewedDateTime sets the reviewedDateTime property value. The date and time when a decision was recorded. The date and time information uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
func (m *ApprovalStage) SetReviewedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.reviewedDateTime = value
}
// SetReviewResult sets the reviewResult property value. The result of this approval record. Possible values include: NotReviewed, Approved, Denied.
func (m *ApprovalStage) SetReviewResult(value *string)() {
    m.reviewResult = value
}
// SetStatus sets the status property value. The stage status. Possible values: InProgress, Initializing, Completed, Expired. Read-only.
func (m *ApprovalStage) SetStatus(value *string)() {
    m.status = value
}
