package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedApprovalStage 
type UnifiedApprovalStage struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The number of days that a request can be pending a response before it is automatically denied.
    approvalStageTimeOutInDays *int32
    // The escalation approvers for this stage when the primary approvers don't respond.
    escalationApprovers []SubjectSetable
    // The time a request can be pending a response from a primary approver before it can be escalated to the escalation approvers.
    escalationTimeInMinutes *int32
    // Indicates whether the approver must provide justification for their reponse.
    isApproverJustificationRequired *bool
    // Indicates whether escalation if enabled.
    isEscalationEnabled *bool
    // The OdataType property
    odataType *string
    // The primary approvers of this stage.
    primaryApprovers []SubjectSetable
}
// NewUnifiedApprovalStage instantiates a new unifiedApprovalStage and sets the default values.
func NewUnifiedApprovalStage()(*UnifiedApprovalStage) {
    m := &UnifiedApprovalStage{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUnifiedApprovalStageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedApprovalStageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedApprovalStage(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedApprovalStage) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetApprovalStageTimeOutInDays gets the approvalStageTimeOutInDays property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *UnifiedApprovalStage) GetApprovalStageTimeOutInDays()(*int32) {
    return m.approvalStageTimeOutInDays
}
// GetEscalationApprovers gets the escalationApprovers property value. The escalation approvers for this stage when the primary approvers don't respond.
func (m *UnifiedApprovalStage) GetEscalationApprovers()([]SubjectSetable) {
    return m.escalationApprovers
}
// GetEscalationTimeInMinutes gets the escalationTimeInMinutes property value. The time a request can be pending a response from a primary approver before it can be escalated to the escalation approvers.
func (m *UnifiedApprovalStage) GetEscalationTimeInMinutes()(*int32) {
    return m.escalationTimeInMinutes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedApprovalStage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["approvalStageTimeOutInDays"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApprovalStageTimeOutInDays(val)
        }
        return nil
    }
    res["escalationApprovers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetEscalationApprovers(res)
        }
        return nil
    }
    res["escalationTimeInMinutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEscalationTimeInMinutes(val)
        }
        return nil
    }
    res["isApproverJustificationRequired"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApproverJustificationRequired(val)
        }
        return nil
    }
    res["isEscalationEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEscalationEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["primaryApprovers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubjectSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SubjectSetable, len(val))
            for i, v := range val {
                res[i] = v.(SubjectSetable)
            }
            m.SetPrimaryApprovers(res)
        }
        return nil
    }
    return res
}
// GetIsApproverJustificationRequired gets the isApproverJustificationRequired property value. Indicates whether the approver must provide justification for their reponse.
func (m *UnifiedApprovalStage) GetIsApproverJustificationRequired()(*bool) {
    return m.isApproverJustificationRequired
}
// GetIsEscalationEnabled gets the isEscalationEnabled property value. Indicates whether escalation if enabled.
func (m *UnifiedApprovalStage) GetIsEscalationEnabled()(*bool) {
    return m.isEscalationEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UnifiedApprovalStage) GetOdataType()(*string) {
    return m.odataType
}
// GetPrimaryApprovers gets the primaryApprovers property value. The primary approvers of this stage.
func (m *UnifiedApprovalStage) GetPrimaryApprovers()([]SubjectSetable) {
    return m.primaryApprovers
}
// Serialize serializes information the current object
func (m *UnifiedApprovalStage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("approvalStageTimeOutInDays", m.GetApprovalStageTimeOutInDays())
        if err != nil {
            return err
        }
    }
    if m.GetEscalationApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEscalationApprovers()))
        for i, v := range m.GetEscalationApprovers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("escalationApprovers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("escalationTimeInMinutes", m.GetEscalationTimeInMinutes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApproverJustificationRequired", m.GetIsApproverJustificationRequired())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEscalationEnabled", m.GetIsEscalationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetPrimaryApprovers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPrimaryApprovers()))
        for i, v := range m.GetPrimaryApprovers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("primaryApprovers", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UnifiedApprovalStage) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetApprovalStageTimeOutInDays sets the approvalStageTimeOutInDays property value. The number of days that a request can be pending a response before it is automatically denied.
func (m *UnifiedApprovalStage) SetApprovalStageTimeOutInDays(value *int32)() {
    m.approvalStageTimeOutInDays = value
}
// SetEscalationApprovers sets the escalationApprovers property value. The escalation approvers for this stage when the primary approvers don't respond.
func (m *UnifiedApprovalStage) SetEscalationApprovers(value []SubjectSetable)() {
    m.escalationApprovers = value
}
// SetEscalationTimeInMinutes sets the escalationTimeInMinutes property value. The time a request can be pending a response from a primary approver before it can be escalated to the escalation approvers.
func (m *UnifiedApprovalStage) SetEscalationTimeInMinutes(value *int32)() {
    m.escalationTimeInMinutes = value
}
// SetIsApproverJustificationRequired sets the isApproverJustificationRequired property value. Indicates whether the approver must provide justification for their reponse.
func (m *UnifiedApprovalStage) SetIsApproverJustificationRequired(value *bool)() {
    m.isApproverJustificationRequired = value
}
// SetIsEscalationEnabled sets the isEscalationEnabled property value. Indicates whether escalation if enabled.
func (m *UnifiedApprovalStage) SetIsEscalationEnabled(value *bool)() {
    m.isEscalationEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UnifiedApprovalStage) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPrimaryApprovers sets the primaryApprovers property value. The primary approvers of this stage.
func (m *UnifiedApprovalStage) SetPrimaryApprovers(value []SubjectSetable)() {
    m.primaryApprovers = value
}
