package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScheduleChangeRequest 
type ScheduleChangeRequest struct {
    ChangeTrackedEntity
    // The assignedTo property
    assignedTo *ScheduleChangeRequestActor
    // The managerActionDateTime property
    managerActionDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The managerActionMessage property
    managerActionMessage *string
    // The managerUserId property
    managerUserId *string
    // The senderDateTime property
    senderDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The senderMessage property
    senderMessage *string
    // The senderUserId property
    senderUserId *string
    // The state property
    state *ScheduleChangeState
}
// NewScheduleChangeRequest instantiates a new ScheduleChangeRequest and sets the default values.
func NewScheduleChangeRequest()(*ScheduleChangeRequest) {
    m := &ScheduleChangeRequest{
        ChangeTrackedEntity: *NewChangeTrackedEntity(),
    }
    odataTypeValue := "#microsoft.graph.scheduleChangeRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateScheduleChangeRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScheduleChangeRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.offerShiftRequest":
                        return NewOfferShiftRequest(), nil
                    case "#microsoft.graph.openShiftChangeRequest":
                        return NewOpenShiftChangeRequest(), nil
                    case "#microsoft.graph.swapShiftsChangeRequest":
                        return NewSwapShiftsChangeRequest(), nil
                    case "#microsoft.graph.timeOffRequest":
                        return NewTimeOffRequest(), nil
                }
            }
        }
    }
    return NewScheduleChangeRequest(), nil
}
// GetAssignedTo gets the assignedTo property value. The assignedTo property
func (m *ScheduleChangeRequest) GetAssignedTo()(*ScheduleChangeRequestActor) {
    return m.assignedTo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScheduleChangeRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ChangeTrackedEntity.GetFieldDeserializers()
    res["assignedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScheduleChangeRequestActor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedTo(val.(*ScheduleChangeRequestActor))
        }
        return nil
    }
    res["managerActionDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagerActionDateTime(val)
        }
        return nil
    }
    res["managerActionMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagerActionMessage(val)
        }
        return nil
    }
    res["managerUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagerUserId(val)
        }
        return nil
    }
    res["senderDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderDateTime(val)
        }
        return nil
    }
    res["senderMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderMessage(val)
        }
        return nil
    }
    res["senderUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderUserId(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseScheduleChangeState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ScheduleChangeState))
        }
        return nil
    }
    return res
}
// GetManagerActionDateTime gets the managerActionDateTime property value. The managerActionDateTime property
func (m *ScheduleChangeRequest) GetManagerActionDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.managerActionDateTime
}
// GetManagerActionMessage gets the managerActionMessage property value. The managerActionMessage property
func (m *ScheduleChangeRequest) GetManagerActionMessage()(*string) {
    return m.managerActionMessage
}
// GetManagerUserId gets the managerUserId property value. The managerUserId property
func (m *ScheduleChangeRequest) GetManagerUserId()(*string) {
    return m.managerUserId
}
// GetSenderDateTime gets the senderDateTime property value. The senderDateTime property
func (m *ScheduleChangeRequest) GetSenderDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.senderDateTime
}
// GetSenderMessage gets the senderMessage property value. The senderMessage property
func (m *ScheduleChangeRequest) GetSenderMessage()(*string) {
    return m.senderMessage
}
// GetSenderUserId gets the senderUserId property value. The senderUserId property
func (m *ScheduleChangeRequest) GetSenderUserId()(*string) {
    return m.senderUserId
}
// GetState gets the state property value. The state property
func (m *ScheduleChangeRequest) GetState()(*ScheduleChangeState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ScheduleChangeRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ChangeTrackedEntity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignedTo() != nil {
        cast := (*m.GetAssignedTo()).String()
        err = writer.WriteStringValue("assignedTo", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("managerActionMessage", m.GetManagerActionMessage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("senderMessage", m.GetSenderMessage())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedTo sets the assignedTo property value. The assignedTo property
func (m *ScheduleChangeRequest) SetAssignedTo(value *ScheduleChangeRequestActor)() {
    m.assignedTo = value
}
// SetManagerActionDateTime sets the managerActionDateTime property value. The managerActionDateTime property
func (m *ScheduleChangeRequest) SetManagerActionDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.managerActionDateTime = value
}
// SetManagerActionMessage sets the managerActionMessage property value. The managerActionMessage property
func (m *ScheduleChangeRequest) SetManagerActionMessage(value *string)() {
    m.managerActionMessage = value
}
// SetManagerUserId sets the managerUserId property value. The managerUserId property
func (m *ScheduleChangeRequest) SetManagerUserId(value *string)() {
    m.managerUserId = value
}
// SetSenderDateTime sets the senderDateTime property value. The senderDateTime property
func (m *ScheduleChangeRequest) SetSenderDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.senderDateTime = value
}
// SetSenderMessage sets the senderMessage property value. The senderMessage property
func (m *ScheduleChangeRequest) SetSenderMessage(value *string)() {
    m.senderMessage = value
}
// SetSenderUserId sets the senderUserId property value. The senderUserId property
func (m *ScheduleChangeRequest) SetSenderUserId(value *string)() {
    m.senderUserId = value
}
// SetState sets the state property value. The state property
func (m *ScheduleChangeRequest) SetState(value *ScheduleChangeState)() {
    m.state = value
}
