package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminRelationshipOperation 
type DelegatedAdminRelationshipOperation struct {
    Entity
    // The time in ISO 8601 format and in UTC time when the long-running operation was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The data (payload) for the operation. Read-only.
    data *string
    // The time in ISO 8601 format and in UTC time when the long-running operation was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The operationType property
    operationType *DelegatedAdminRelationshipOperationType
    // The status property
    status *LongRunningOperationStatus
}
// NewDelegatedAdminRelationshipOperation instantiates a new delegatedAdminRelationshipOperation and sets the default values.
func NewDelegatedAdminRelationshipOperation()(*DelegatedAdminRelationshipOperation) {
    m := &DelegatedAdminRelationshipOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminRelationshipOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminRelationshipOperation(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The time in ISO 8601 format and in UTC time when the long-running operation was created. Read-only.
func (m *DelegatedAdminRelationshipOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetData gets the data property value. The data (payload) for the operation. Read-only.
func (m *DelegatedAdminRelationshipOperation) GetData()(*string) {
    return m.data
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminRelationshipOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["data"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetData(val)
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
    res["operationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*DelegatedAdminRelationshipOperationType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLongRunningOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*LongRunningOperationStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The time in ISO 8601 format and in UTC time when the long-running operation was last modified. Read-only.
func (m *DelegatedAdminRelationshipOperation) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetOperationType gets the operationType property value. The operationType property
func (m *DelegatedAdminRelationshipOperation) GetOperationType()(*DelegatedAdminRelationshipOperationType) {
    return m.operationType
}
// GetStatus gets the status property value. The status property
func (m *DelegatedAdminRelationshipOperation) GetStatus()(*LongRunningOperationStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DelegatedAdminRelationshipOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("data", m.GetData())
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
    if m.GetOperationType() != nil {
        cast := (*m.GetOperationType()).String()
        err = writer.WriteStringValue("operationType", &cast)
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
// SetCreatedDateTime sets the createdDateTime property value. The time in ISO 8601 format and in UTC time when the long-running operation was created. Read-only.
func (m *DelegatedAdminRelationshipOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetData sets the data property value. The data (payload) for the operation. Read-only.
func (m *DelegatedAdminRelationshipOperation) SetData(value *string)() {
    m.data = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The time in ISO 8601 format and in UTC time when the long-running operation was last modified. Read-only.
func (m *DelegatedAdminRelationshipOperation) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetOperationType sets the operationType property value. The operationType property
func (m *DelegatedAdminRelationshipOperation) SetOperationType(value *DelegatedAdminRelationshipOperationType)() {
    m.operationType = value
}
// SetStatus sets the status property value. The status property
func (m *DelegatedAdminRelationshipOperation) SetStatus(value *LongRunningOperationStatus)() {
    m.status = value
}
