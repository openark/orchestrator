package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignedPlan 
type AssignedPlan struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date and time at which the plan was assigned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    assignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut. See a detailed description of each value.
    capabilityStatus *string
    // The OdataType property
    odataType *string
    // The name of the service; for example, exchange.
    service *string
    // A GUID that identifies the service plan. For a complete list of GUIDs and their equivalent friendly service names, see Product names and service plan identifiers for licensing.
    servicePlanId *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
}
// NewAssignedPlan instantiates a new assignedPlan and sets the default values.
func NewAssignedPlan()(*AssignedPlan) {
    m := &AssignedPlan{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAssignedPlanFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignedPlanFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignedPlan(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedPlan) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAssignedDateTime gets the assignedDateTime property value. The date and time at which the plan was assigned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AssignedPlan) GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.assignedDateTime
}
// GetCapabilityStatus gets the capabilityStatus property value. Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut. See a detailed description of each value.
func (m *AssignedPlan) GetCapabilityStatus()(*string) {
    return m.capabilityStatus
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignedPlan) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["assignedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedDateTime(val)
        }
        return nil
    }
    res["capabilityStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCapabilityStatus(val)
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
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["servicePlanId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePlanId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignedPlan) GetOdataType()(*string) {
    return m.odataType
}
// GetService gets the service property value. The name of the service; for example, exchange.
func (m *AssignedPlan) GetService()(*string) {
    return m.service
}
// GetServicePlanId gets the servicePlanId property value. A GUID that identifies the service plan. For a complete list of GUIDs and their equivalent friendly service names, see Product names and service plan identifiers for licensing.
func (m *AssignedPlan) GetServicePlanId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.servicePlanId
}
// Serialize serializes information the current object
func (m *AssignedPlan) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("assignedDateTime", m.GetAssignedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("capabilityStatus", m.GetCapabilityStatus())
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
    {
        err := writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("servicePlanId", m.GetServicePlanId())
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
func (m *AssignedPlan) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAssignedDateTime sets the assignedDateTime property value. The date and time at which the plan was assigned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *AssignedPlan) SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignedDateTime = value
}
// SetCapabilityStatus sets the capabilityStatus property value. Condition of the capability assignment. The possible values are Enabled, Warning, Suspended, Deleted, LockedOut. See a detailed description of each value.
func (m *AssignedPlan) SetCapabilityStatus(value *string)() {
    m.capabilityStatus = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignedPlan) SetOdataType(value *string)() {
    m.odataType = value
}
// SetService sets the service property value. The name of the service; for example, exchange.
func (m *AssignedPlan) SetService(value *string)() {
    m.service = value
}
// SetServicePlanId sets the servicePlanId property value. A GUID that identifies the service plan. For a complete list of GUIDs and their equivalent friendly service names, see Product names and service plan identifiers for licensing.
func (m *AssignedPlan) SetServicePlanId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.servicePlanId = value
}
