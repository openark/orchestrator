package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminRelationship 
type DelegatedAdminRelationship struct {
    Entity
    // The access assignments associated with the delegated admin relationship.
    accessAssignments []DelegatedAdminAccessAssignmentable
    // The accessDetails property
    accessDetails DelegatedAdminAccessDetailsable
    // The date and time in ISO 8601 format and in UTC time when the relationship became active. Read-only.
    activatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time in ISO 8601 format and in UTC time when the relationship was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The display name and unique identifier of the customer of the relationship. This is configured either by the partner at the time the relationship is created or by the system after the customer approves the relationship. Cannot be changed by the customer.
    customer DelegatedAdminRelationshipCustomerParticipantable
    // The display name of the relationship used for ease of identification. Must be unique across all delegated admin relationships of the partner. This is set by the partner only when the relationship is in the created status and cannot be changed by the customer.
    displayName *string
    // The duration of the relationship in ISO 8601 format. Must be a value between P1D and P2Y inclusive. This is set by the partner only when the relationship is in the created status and cannot be changed by the customer.
    duration *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration
    // The date and time in ISO 8601 format and in UTC time when the status of relationship changes to either terminated or expired. Calculated as endDateTime = activatedDateTime + duration. Read-only.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time in ISO 8601 format and in UTC time when the relationship was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The long running operations associated with the delegated admin relationship.
    operations []DelegatedAdminRelationshipOperationable
    // The requests associated with the delegated admin relationship.
    requests []DelegatedAdminRelationshipRequestable
    // The status of the relationship. Read Only. The possible values are: activating, active, approvalPending, approved, created, expired, expiring, terminated, terminating, terminationRequested, unknownFutureValue. Supports $orderBy.
    status *DelegatedAdminRelationshipStatus
}
// NewDelegatedAdminRelationship instantiates a new DelegatedAdminRelationship and sets the default values.
func NewDelegatedAdminRelationship()(*DelegatedAdminRelationship) {
    m := &DelegatedAdminRelationship{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedAdminRelationshipFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminRelationshipFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminRelationship(), nil
}
// GetAccessAssignments gets the accessAssignments property value. The access assignments associated with the delegated admin relationship.
func (m *DelegatedAdminRelationship) GetAccessAssignments()([]DelegatedAdminAccessAssignmentable) {
    return m.accessAssignments
}
// GetAccessDetails gets the accessDetails property value. The accessDetails property
func (m *DelegatedAdminRelationship) GetAccessDetails()(DelegatedAdminAccessDetailsable) {
    return m.accessDetails
}
// GetActivatedDateTime gets the activatedDateTime property value. The date and time in ISO 8601 format and in UTC time when the relationship became active. Read-only.
func (m *DelegatedAdminRelationship) GetActivatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.activatedDateTime
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time in ISO 8601 format and in UTC time when the relationship was created. Read-only.
func (m *DelegatedAdminRelationship) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetCustomer gets the customer property value. The display name and unique identifier of the customer of the relationship. This is configured either by the partner at the time the relationship is created or by the system after the customer approves the relationship. Cannot be changed by the customer.
func (m *DelegatedAdminRelationship) GetCustomer()(DelegatedAdminRelationshipCustomerParticipantable) {
    return m.customer
}
// GetDisplayName gets the displayName property value. The display name of the relationship used for ease of identification. Must be unique across all delegated admin relationships of the partner. This is set by the partner only when the relationship is in the created status and cannot be changed by the customer.
func (m *DelegatedAdminRelationship) GetDisplayName()(*string) {
    return m.displayName
}
// GetDuration gets the duration property value. The duration of the relationship in ISO 8601 format. Must be a value between P1D and P2Y inclusive. This is set by the partner only when the relationship is in the created status and cannot be changed by the customer.
func (m *DelegatedAdminRelationship) GetDuration()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration) {
    return m.duration
}
// GetEndDateTime gets the endDateTime property value. The date and time in ISO 8601 format and in UTC time when the status of relationship changes to either terminated or expired. Calculated as endDateTime = activatedDateTime + duration. Read-only.
func (m *DelegatedAdminRelationship) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminRelationship) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessAssignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminAccessAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminAccessAssignmentable)
            }
            m.SetAccessAssignments(res)
        }
        return nil
    }
    res["accessDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDelegatedAdminAccessDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessDetails(val.(DelegatedAdminAccessDetailsable))
        }
        return nil
    }
    res["activatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivatedDateTime(val)
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
    res["customer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDelegatedAdminRelationshipCustomerParticipantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomer(val.(DelegatedAdminRelationshipCustomerParticipantable))
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
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
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
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminRelationshipOperationable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminRelationshipOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["requests"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDelegatedAdminRelationshipRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DelegatedAdminRelationshipRequestable, len(val))
            for i, v := range val {
                res[i] = v.(DelegatedAdminRelationshipRequestable)
            }
            m.SetRequests(res)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DelegatedAdminRelationshipStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time in ISO 8601 format and in UTC time when the relationship was last modified. Read-only.
func (m *DelegatedAdminRelationship) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOperations gets the operations property value. The long running operations associated with the delegated admin relationship.
func (m *DelegatedAdminRelationship) GetOperations()([]DelegatedAdminRelationshipOperationable) {
    return m.operations
}
// GetRequests gets the requests property value. The requests associated with the delegated admin relationship.
func (m *DelegatedAdminRelationship) GetRequests()([]DelegatedAdminRelationshipRequestable) {
    return m.requests
}
// GetStatus gets the status property value. The status of the relationship. Read Only. The possible values are: activating, active, approvalPending, approved, created, expired, expiring, terminated, terminating, terminationRequested, unknownFutureValue. Supports $orderBy.
func (m *DelegatedAdminRelationship) GetStatus()(*DelegatedAdminRelationshipStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DelegatedAdminRelationship) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAccessAssignments()))
        for i, v := range m.GetAccessAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("accessDetails", m.GetAccessDetails())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("activatedDateTime", m.GetActivatedDateTime())
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
        err = writer.WriteObjectValue("customer", m.GetCustomer())
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
        err = writer.WriteISODurationValue("duration", m.GetDuration())
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
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequests() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRequests()))
        for i, v := range m.GetRequests() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("requests", cast)
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
// SetAccessAssignments sets the accessAssignments property value. The access assignments associated with the delegated admin relationship.
func (m *DelegatedAdminRelationship) SetAccessAssignments(value []DelegatedAdminAccessAssignmentable)() {
    m.accessAssignments = value
}
// SetAccessDetails sets the accessDetails property value. The accessDetails property
func (m *DelegatedAdminRelationship) SetAccessDetails(value DelegatedAdminAccessDetailsable)() {
    m.accessDetails = value
}
// SetActivatedDateTime sets the activatedDateTime property value. The date and time in ISO 8601 format and in UTC time when the relationship became active. Read-only.
func (m *DelegatedAdminRelationship) SetActivatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activatedDateTime = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time in ISO 8601 format and in UTC time when the relationship was created. Read-only.
func (m *DelegatedAdminRelationship) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetCustomer sets the customer property value. The display name and unique identifier of the customer of the relationship. This is configured either by the partner at the time the relationship is created or by the system after the customer approves the relationship. Cannot be changed by the customer.
func (m *DelegatedAdminRelationship) SetCustomer(value DelegatedAdminRelationshipCustomerParticipantable)() {
    m.customer = value
}
// SetDisplayName sets the displayName property value. The display name of the relationship used for ease of identification. Must be unique across all delegated admin relationships of the partner. This is set by the partner only when the relationship is in the created status and cannot be changed by the customer.
func (m *DelegatedAdminRelationship) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDuration sets the duration property value. The duration of the relationship in ISO 8601 format. Must be a value between P1D and P2Y inclusive. This is set by the partner only when the relationship is in the created status and cannot be changed by the customer.
func (m *DelegatedAdminRelationship) SetDuration(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ISODuration)() {
    m.duration = value
}
// SetEndDateTime sets the endDateTime property value. The date and time in ISO 8601 format and in UTC time when the status of relationship changes to either terminated or expired. Calculated as endDateTime = activatedDateTime + duration. Read-only.
func (m *DelegatedAdminRelationship) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time in ISO 8601 format and in UTC time when the relationship was last modified. Read-only.
func (m *DelegatedAdminRelationship) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOperations sets the operations property value. The long running operations associated with the delegated admin relationship.
func (m *DelegatedAdminRelationship) SetOperations(value []DelegatedAdminRelationshipOperationable)() {
    m.operations = value
}
// SetRequests sets the requests property value. The requests associated with the delegated admin relationship.
func (m *DelegatedAdminRelationship) SetRequests(value []DelegatedAdminRelationshipRequestable)() {
    m.requests = value
}
// SetStatus sets the status property value. The status of the relationship. Read Only. The possible values are: activating, active, approvalPending, approved, created, expired, expiring, terminated, terminating, terminationRequested, unknownFutureValue. Supports $orderBy.
func (m *DelegatedAdminRelationship) SetStatus(value *DelegatedAdminRelationshipStatus)() {
    m.status = value
}
