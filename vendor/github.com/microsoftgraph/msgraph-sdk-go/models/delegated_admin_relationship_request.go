package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DelegatedAdminRelationshipRequest 
type DelegatedAdminRelationshipRequest struct {
    Entity
    // The action property
    action *DelegatedAdminRelationshipRequestAction
    // The date and time in ISO 8601 format and in UTC time when the relationship request was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date and time in ISO 8601 format and UTC time when this relationship request was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The status of the request. Read-only. The possible values are: created, pending, succeeded, failed, unknownFutureValue.
    status *DelegatedAdminRelationshipRequestStatus
}
// NewDelegatedAdminRelationshipRequest instantiates a new delegatedAdminRelationshipRequest and sets the default values.
func NewDelegatedAdminRelationshipRequest()(*DelegatedAdminRelationshipRequest) {
    m := &DelegatedAdminRelationshipRequest{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDelegatedAdminRelationshipRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDelegatedAdminRelationshipRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDelegatedAdminRelationshipRequest(), nil
}
// GetAction gets the action property value. The action property
func (m *DelegatedAdminRelationshipRequest) GetAction()(*DelegatedAdminRelationshipRequestAction) {
    return m.action
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time in ISO 8601 format and in UTC time when the relationship request was created. Read-only.
func (m *DelegatedAdminRelationshipRequest) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DelegatedAdminRelationshipRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipRequestAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*DelegatedAdminRelationshipRequestAction))
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
        val, err := n.GetEnumValue(ParseDelegatedAdminRelationshipRequestStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*DelegatedAdminRelationshipRequestStatus))
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time in ISO 8601 format and UTC time when this relationship request was last modified. Read-only.
func (m *DelegatedAdminRelationshipRequest) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetStatus gets the status property value. The status of the request. Read-only. The possible values are: created, pending, succeeded, failed, unknownFutureValue.
func (m *DelegatedAdminRelationshipRequest) GetStatus()(*DelegatedAdminRelationshipRequestStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *DelegatedAdminRelationshipRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
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
// SetAction sets the action property value. The action property
func (m *DelegatedAdminRelationshipRequest) SetAction(value *DelegatedAdminRelationshipRequestAction)() {
    m.action = value
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time in ISO 8601 format and in UTC time when the relationship request was created. Read-only.
func (m *DelegatedAdminRelationshipRequest) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time in ISO 8601 format and UTC time when this relationship request was last modified. Read-only.
func (m *DelegatedAdminRelationshipRequest) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetStatus sets the status property value. The status of the request. Read-only. The possible values are: created, pending, succeeded, failed, unknownFutureValue.
func (m *DelegatedAdminRelationshipRequest) SetStatus(value *DelegatedAdminRelationshipRequestStatus)() {
    m.status = value
}
