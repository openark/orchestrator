package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningObjectSummary 
type ProvisioningObjectSummary struct {
    Entity
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique ID of this change in this cycle.
    changeId *string
    // Unique ID per job iteration.
    cycleId *string
    // Indicates how long this provisioning action took to finish. Measured in milliseconds.
    durationInMilliseconds *int32
    // Details of who initiated this provisioning.
    initiatedBy Initiatorable
    // The unique ID for the whole provisioning job.
    jobId *string
    // Details of each property that was modified in this provisioning action on this object.
    modifiedProperties []ModifiedPropertyable
    // Indicates the activity name or the operation name. Possible values are: create, update, delete, stageddelete, disable, other and unknownFutureValue. For a list of activities logged, refer to Azure AD activity list.
    provisioningAction *ProvisioningAction
    // Details of provisioning status.
    provisioningStatusInfo ProvisioningStatusInfoable
    // Details of each step in provisioning.
    provisioningSteps []ProvisioningStepable
    // Represents the service principal used for provisioning.
    servicePrincipal ProvisioningServicePrincipalable
    // Details of source object being provisioned.
    sourceIdentity ProvisionedIdentityable
    // Details of source system of the object being provisioned.
    sourceSystem ProvisioningSystemable
    // Details of target object being provisioned.
    targetIdentity ProvisionedIdentityable
    // Details of target system of the object being provisioned.
    targetSystem ProvisioningSystemable
    // Unique Azure AD tenant ID.
    tenantId *string
}
// NewProvisioningObjectSummary instantiates a new provisioningObjectSummary and sets the default values.
func NewProvisioningObjectSummary()(*ProvisioningObjectSummary) {
    m := &ProvisioningObjectSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateProvisioningObjectSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningObjectSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningObjectSummary(), nil
}
// GetActivityDateTime gets the activityDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ProvisioningObjectSummary) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activityDateTime
}
// GetChangeId gets the changeId property value. Unique ID of this change in this cycle.
func (m *ProvisioningObjectSummary) GetChangeId()(*string) {
    return m.changeId
}
// GetCycleId gets the cycleId property value. Unique ID per job iteration.
func (m *ProvisioningObjectSummary) GetCycleId()(*string) {
    return m.cycleId
}
// GetDurationInMilliseconds gets the durationInMilliseconds property value. Indicates how long this provisioning action took to finish. Measured in milliseconds.
func (m *ProvisioningObjectSummary) GetDurationInMilliseconds()(*int32) {
    return m.durationInMilliseconds
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningObjectSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["changeId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeId(val)
        }
        return nil
    }
    res["cycleId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCycleId(val)
        }
        return nil
    }
    res["durationInMilliseconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInMilliseconds(val)
        }
        return nil
    }
    res["initiatedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateInitiatorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedBy(val.(Initiatorable))
        }
        return nil
    }
    res["jobId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobId(val)
        }
        return nil
    }
    res["modifiedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateModifiedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ModifiedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(ModifiedPropertyable)
            }
            m.SetModifiedProperties(res)
        }
        return nil
    }
    res["provisioningAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningAction(val.(*ProvisioningAction))
        }
        return nil
    }
    res["provisioningStatusInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningStatusInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStatusInfo(val.(ProvisioningStatusInfoable))
        }
        return nil
    }
    res["provisioningSteps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateProvisioningStepFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningStepable, len(val))
            for i, v := range val {
                res[i] = v.(ProvisioningStepable)
            }
            m.SetProvisioningSteps(res)
        }
        return nil
    }
    res["servicePrincipal"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningServicePrincipalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipal(val.(ProvisioningServicePrincipalable))
        }
        return nil
    }
    res["sourceIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisionedIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIdentity(val.(ProvisionedIdentityable))
        }
        return nil
    }
    res["sourceSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceSystem(val.(ProvisioningSystemable))
        }
        return nil
    }
    res["targetIdentity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisionedIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetIdentity(val.(ProvisionedIdentityable))
        }
        return nil
    }
    res["targetSystem"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningSystemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSystem(val.(ProvisioningSystemable))
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
// GetInitiatedBy gets the initiatedBy property value. Details of who initiated this provisioning.
func (m *ProvisioningObjectSummary) GetInitiatedBy()(Initiatorable) {
    return m.initiatedBy
}
// GetJobId gets the jobId property value. The unique ID for the whole provisioning job.
func (m *ProvisioningObjectSummary) GetJobId()(*string) {
    return m.jobId
}
// GetModifiedProperties gets the modifiedProperties property value. Details of each property that was modified in this provisioning action on this object.
func (m *ProvisioningObjectSummary) GetModifiedProperties()([]ModifiedPropertyable) {
    return m.modifiedProperties
}
// GetProvisioningAction gets the provisioningAction property value. Indicates the activity name or the operation name. Possible values are: create, update, delete, stageddelete, disable, other and unknownFutureValue. For a list of activities logged, refer to Azure AD activity list.
func (m *ProvisioningObjectSummary) GetProvisioningAction()(*ProvisioningAction) {
    return m.provisioningAction
}
// GetProvisioningStatusInfo gets the provisioningStatusInfo property value. Details of provisioning status.
func (m *ProvisioningObjectSummary) GetProvisioningStatusInfo()(ProvisioningStatusInfoable) {
    return m.provisioningStatusInfo
}
// GetProvisioningSteps gets the provisioningSteps property value. Details of each step in provisioning.
func (m *ProvisioningObjectSummary) GetProvisioningSteps()([]ProvisioningStepable) {
    return m.provisioningSteps
}
// GetServicePrincipal gets the servicePrincipal property value. Represents the service principal used for provisioning.
func (m *ProvisioningObjectSummary) GetServicePrincipal()(ProvisioningServicePrincipalable) {
    return m.servicePrincipal
}
// GetSourceIdentity gets the sourceIdentity property value. Details of source object being provisioned.
func (m *ProvisioningObjectSummary) GetSourceIdentity()(ProvisionedIdentityable) {
    return m.sourceIdentity
}
// GetSourceSystem gets the sourceSystem property value. Details of source system of the object being provisioned.
func (m *ProvisioningObjectSummary) GetSourceSystem()(ProvisioningSystemable) {
    return m.sourceSystem
}
// GetTargetIdentity gets the targetIdentity property value. Details of target object being provisioned.
func (m *ProvisioningObjectSummary) GetTargetIdentity()(ProvisionedIdentityable) {
    return m.targetIdentity
}
// GetTargetSystem gets the targetSystem property value. Details of target system of the object being provisioned.
func (m *ProvisioningObjectSummary) GetTargetSystem()(ProvisioningSystemable) {
    return m.targetSystem
}
// GetTenantId gets the tenantId property value. Unique Azure AD tenant ID.
func (m *ProvisioningObjectSummary) GetTenantId()(*string) {
    return m.tenantId
}
// Serialize serializes information the current object
func (m *ProvisioningObjectSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("changeId", m.GetChangeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cycleId", m.GetCycleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("durationInMilliseconds", m.GetDurationInMilliseconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobId", m.GetJobId())
        if err != nil {
            return err
        }
    }
    if m.GetModifiedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetModifiedProperties()))
        for i, v := range m.GetModifiedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("modifiedProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningAction() != nil {
        cast := (*m.GetProvisioningAction()).String()
        err = writer.WriteStringValue("provisioningAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("provisioningStatusInfo", m.GetProvisioningStatusInfo())
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningSteps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProvisioningSteps()))
        for i, v := range m.GetProvisioningSteps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("provisioningSteps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("servicePrincipal", m.GetServicePrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceIdentity", m.GetSourceIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceSystem", m.GetSourceSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetIdentity", m.GetTargetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetSystem", m.GetTargetSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActivityDateTime sets the activityDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ProvisioningObjectSummary) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// SetChangeId sets the changeId property value. Unique ID of this change in this cycle.
func (m *ProvisioningObjectSummary) SetChangeId(value *string)() {
    m.changeId = value
}
// SetCycleId sets the cycleId property value. Unique ID per job iteration.
func (m *ProvisioningObjectSummary) SetCycleId(value *string)() {
    m.cycleId = value
}
// SetDurationInMilliseconds sets the durationInMilliseconds property value. Indicates how long this provisioning action took to finish. Measured in milliseconds.
func (m *ProvisioningObjectSummary) SetDurationInMilliseconds(value *int32)() {
    m.durationInMilliseconds = value
}
// SetInitiatedBy sets the initiatedBy property value. Details of who initiated this provisioning.
func (m *ProvisioningObjectSummary) SetInitiatedBy(value Initiatorable)() {
    m.initiatedBy = value
}
// SetJobId sets the jobId property value. The unique ID for the whole provisioning job.
func (m *ProvisioningObjectSummary) SetJobId(value *string)() {
    m.jobId = value
}
// SetModifiedProperties sets the modifiedProperties property value. Details of each property that was modified in this provisioning action on this object.
func (m *ProvisioningObjectSummary) SetModifiedProperties(value []ModifiedPropertyable)() {
    m.modifiedProperties = value
}
// SetProvisioningAction sets the provisioningAction property value. Indicates the activity name or the operation name. Possible values are: create, update, delete, stageddelete, disable, other and unknownFutureValue. For a list of activities logged, refer to Azure AD activity list.
func (m *ProvisioningObjectSummary) SetProvisioningAction(value *ProvisioningAction)() {
    m.provisioningAction = value
}
// SetProvisioningStatusInfo sets the provisioningStatusInfo property value. Details of provisioning status.
func (m *ProvisioningObjectSummary) SetProvisioningStatusInfo(value ProvisioningStatusInfoable)() {
    m.provisioningStatusInfo = value
}
// SetProvisioningSteps sets the provisioningSteps property value. Details of each step in provisioning.
func (m *ProvisioningObjectSummary) SetProvisioningSteps(value []ProvisioningStepable)() {
    m.provisioningSteps = value
}
// SetServicePrincipal sets the servicePrincipal property value. Represents the service principal used for provisioning.
func (m *ProvisioningObjectSummary) SetServicePrincipal(value ProvisioningServicePrincipalable)() {
    m.servicePrincipal = value
}
// SetSourceIdentity sets the sourceIdentity property value. Details of source object being provisioned.
func (m *ProvisioningObjectSummary) SetSourceIdentity(value ProvisionedIdentityable)() {
    m.sourceIdentity = value
}
// SetSourceSystem sets the sourceSystem property value. Details of source system of the object being provisioned.
func (m *ProvisioningObjectSummary) SetSourceSystem(value ProvisioningSystemable)() {
    m.sourceSystem = value
}
// SetTargetIdentity sets the targetIdentity property value. Details of target object being provisioned.
func (m *ProvisioningObjectSummary) SetTargetIdentity(value ProvisionedIdentityable)() {
    m.targetIdentity = value
}
// SetTargetSystem sets the targetSystem property value. Details of target system of the object being provisioned.
func (m *ProvisioningObjectSummary) SetTargetSystem(value ProvisioningSystemable)() {
    m.targetSystem = value
}
// SetTenantId sets the tenantId property value. Unique Azure AD tenant ID.
func (m *ProvisioningObjectSummary) SetTenantId(value *string)() {
    m.tenantId = value
}
