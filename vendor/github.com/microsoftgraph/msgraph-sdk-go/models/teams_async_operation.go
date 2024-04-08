package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAsyncOperation 
type TeamsAsyncOperation struct {
    Entity
    // Number of times the operation was attempted before being marked successful or failed.
    attemptsCount *int32
    // Time when the operation was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Any error that causes the async operation to fail.
    error OperationErrorable
    // Time when the async operation was last updated.
    lastActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The operationType property
    operationType *TeamsAsyncOperationType
    // The status property
    status *TeamsAsyncOperationStatus
    // The ID of the object that's created or modified as result of this async operation, typically a team.
    targetResourceId *string
    // The location of the object that's created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths.
    targetResourceLocation *string
}
// NewTeamsAsyncOperation instantiates a new teamsAsyncOperation and sets the default values.
func NewTeamsAsyncOperation()(*TeamsAsyncOperation) {
    m := &TeamsAsyncOperation{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsAsyncOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsAsyncOperationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamsAsyncOperation(), nil
}
// GetAttemptsCount gets the attemptsCount property value. Number of times the operation was attempted before being marked successful or failed.
func (m *TeamsAsyncOperation) GetAttemptsCount()(*int32) {
    return m.attemptsCount
}
// GetCreatedDateTime gets the createdDateTime property value. Time when the operation was created.
func (m *TeamsAsyncOperation) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetError gets the error property value. Any error that causes the async operation to fail.
func (m *TeamsAsyncOperation) GetError()(OperationErrorable) {
    return m.error
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsAsyncOperation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attemptsCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttemptsCount(val)
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
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOperationErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(OperationErrorable))
        }
        return nil
    }
    res["lastActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActionDateTime(val)
        }
        return nil
    }
    res["operationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAsyncOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*TeamsAsyncOperationType))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamsAsyncOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TeamsAsyncOperationStatus))
        }
        return nil
    }
    res["targetResourceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetResourceId(val)
        }
        return nil
    }
    res["targetResourceLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetResourceLocation(val)
        }
        return nil
    }
    return res
}
// GetLastActionDateTime gets the lastActionDateTime property value. Time when the async operation was last updated.
func (m *TeamsAsyncOperation) GetLastActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActionDateTime
}
// GetOperationType gets the operationType property value. The operationType property
func (m *TeamsAsyncOperation) GetOperationType()(*TeamsAsyncOperationType) {
    return m.operationType
}
// GetStatus gets the status property value. The status property
func (m *TeamsAsyncOperation) GetStatus()(*TeamsAsyncOperationStatus) {
    return m.status
}
// GetTargetResourceId gets the targetResourceId property value. The ID of the object that's created or modified as result of this async operation, typically a team.
func (m *TeamsAsyncOperation) GetTargetResourceId()(*string) {
    return m.targetResourceId
}
// GetTargetResourceLocation gets the targetResourceLocation property value. The location of the object that's created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths.
func (m *TeamsAsyncOperation) GetTargetResourceLocation()(*string) {
    return m.targetResourceLocation
}
// Serialize serializes information the current object
func (m *TeamsAsyncOperation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("attemptsCount", m.GetAttemptsCount())
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
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastActionDateTime", m.GetLastActionDateTime())
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
    {
        err = writer.WriteStringValue("targetResourceId", m.GetTargetResourceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("targetResourceLocation", m.GetTargetResourceLocation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttemptsCount sets the attemptsCount property value. Number of times the operation was attempted before being marked successful or failed.
func (m *TeamsAsyncOperation) SetAttemptsCount(value *int32)() {
    m.attemptsCount = value
}
// SetCreatedDateTime sets the createdDateTime property value. Time when the operation was created.
func (m *TeamsAsyncOperation) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetError sets the error property value. Any error that causes the async operation to fail.
func (m *TeamsAsyncOperation) SetError(value OperationErrorable)() {
    m.error = value
}
// SetLastActionDateTime sets the lastActionDateTime property value. Time when the async operation was last updated.
func (m *TeamsAsyncOperation) SetLastActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActionDateTime = value
}
// SetOperationType sets the operationType property value. The operationType property
func (m *TeamsAsyncOperation) SetOperationType(value *TeamsAsyncOperationType)() {
    m.operationType = value
}
// SetStatus sets the status property value. The status property
func (m *TeamsAsyncOperation) SetStatus(value *TeamsAsyncOperationStatus)() {
    m.status = value
}
// SetTargetResourceId sets the targetResourceId property value. The ID of the object that's created or modified as result of this async operation, typically a team.
func (m *TeamsAsyncOperation) SetTargetResourceId(value *string)() {
    m.targetResourceId = value
}
// SetTargetResourceLocation sets the targetResourceLocation property value. The location of the object that's created or modified as result of this async operation. This URL should be treated as an opaque value and not parsed into its component paths.
func (m *TeamsAsyncOperation) SetTargetResourceLocation(value *string)() {
    m.targetResourceLocation = value
}
