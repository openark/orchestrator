package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminAccessAssignment 
type DelegatedAdminAccessAssignment struct {
    Entity
    // The accessContainer property
    accessContainer DelegatedAdminAccessContainerable
    // The accessDetails property
    accessDetails DelegatedAdminAccessDetailsable
    // The date and time in ISO 8601 format and in UTC time when the access assignment was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time in ISO 8601 and in UTC time when this access assignment was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status of the access assignment. Read-only. The possible values are: pending, active, deleting, deleted, error, unknownFutureValue.
    status *DelegatedAdminAccessAssignmentStatus
}
// NewDelegatedAdminAccessAssignment instantiates a new DelegatedAdminAccessAssignment and sets the default values.
func NewDelegatedAdminAccessAssignment()(*DelegatedAdminAccessAssignment) {
    m := &DelegatedAdminAccessAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminAccessAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminAccessAssignment(), nil
}
// GetAccessContainer gets the accessContainer property value. The accessContainer property
func (m *DelegatedAdminAccessAssignment) GetAccessContainer()(DelegatedAdminAccessContainerable) {
    return m.accessContainer
}
// GetAccessDetails gets the accessDetails property value. The accessDetails property
func (m *DelegatedAdminAccessAssignment) GetAccessDetails()(DelegatedAdminAccessDetailsable) {
    return m.accessDetails
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time in ISO 8601 format and in UTC time when the access assignment was created. Read-only.
func (m *DelegatedAdminAccessAssignment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminAccessAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessContainer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDelegatedAdminAccessContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessContainer(val.(DelegatedAdminAccessContainerable))
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
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminAccessAssignmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DelegatedAdminAccessAssignmentStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time in ISO 8601 and in UTC time when this access assignment was last modified. Read-only.
func (m *DelegatedAdminAccessAssignment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetStatus gets the status property value. The status of the access assignment. Read-only. The possible values are: pending, active, deleting, deleted, error, unknownFutureValue.
func (m *DelegatedAdminAccessAssignment) GetStatus()(*DelegatedAdminAccessAssignmentStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DelegatedAdminAccessAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessContainer", m.GetAccessContainer())
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
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessContainer sets the accessContainer property value. The accessContainer property
func (m *DelegatedAdminAccessAssignment) SetAccessContainer(value DelegatedAdminAccessContainerable)() {
    m.accessContainer = value
}
// SetAccessDetails sets the accessDetails property value. The accessDetails property
func (m *DelegatedAdminAccessAssignment) SetAccessDetails(value DelegatedAdminAccessDetailsable)() {
    m.accessDetails = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time in ISO 8601 format and in UTC time when the access assignment was created. Read-only.
func (m *DelegatedAdminAccessAssignment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time in ISO 8601 and in UTC time when this access assignment was last modified. Read-only.
func (m *DelegatedAdminAccessAssignment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetStatus sets the status property value. The status of the access assignment. Read-only. The possible values are: pending, active, deleting, deleted, error, unknownFutureValue.
func (m *DelegatedAdminAccessAssignment) SetStatus(value *DelegatedAdminAccessAssignmentStatus)() {
    m.status = value
}
