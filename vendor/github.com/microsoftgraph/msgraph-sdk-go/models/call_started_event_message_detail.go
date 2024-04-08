package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CallStartedEventMessageDetail 
type CallStartedEventMessageDetail struct {
    EventMessageDetail
    // Represents the call event type. Possible values are: call, meeting, screenShare, unknownFutureValue.
    callEventType *TeamworkCallEventType
    // Unique identifier of the call.
    callId *string
    // Initiator of the event.
    initiator IdentitySetable
}
// NewCallStartedEventMessageDetail instantiates a new CallStartedEventMessageDetail and sets the default values.
func NewCallStartedEventMessageDetail()(*CallStartedEventMessageDetail) {
    m := &CallStartedEventMessageDetail{
        EventMessageDetail: *NewEventMessageDetail(),
    }
    odataTypeValue := "#microsoft.graph.callStartedEventMessageDetail"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateCallStartedEventMessageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallStartedEventMessageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCallStartedEventMessageDetail(), nil
}
// GetCallEventType gets the callEventType property value. Represents the call event type. Possible values are: call, meeting, screenShare, unknownFutureValue.
func (m *CallStartedEventMessageDetail) GetCallEventType()(*TeamworkCallEventType) {
    return m.callEventType
}
// GetCallId gets the callId property value. Unique identifier of the call.
func (m *CallStartedEventMessageDetail) GetCallId()(*string) {
    return m.callId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CallStartedEventMessageDetail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EventMessageDetail.GetFieldDeserializers()
    res["callEventType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkCallEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallEventType(val.(*TeamworkCallEventType))
        }
        return nil
    }
    res["callId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallId(val)
        }
        return nil
    }
    res["initiator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiator(val.(IdentitySetable))
        }
        return nil
    }
    return res
}
// GetInitiator gets the initiator property value. Initiator of the event.
func (m *CallStartedEventMessageDetail) GetInitiator()(IdentitySetable) {
    return m.initiator
}
// Serialize serializes information the current object
func (m *CallStartedEventMessageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EventMessageDetail.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCallEventType() != nil {
        cast := (*m.GetCallEventType()).String()
        err = writer.WriteStringValue("callEventType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callId", m.GetCallId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("initiator", m.GetInitiator())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallEventType sets the callEventType property value. Represents the call event type. Possible values are: call, meeting, screenShare, unknownFutureValue.
func (m *CallStartedEventMessageDetail) SetCallEventType(value *TeamworkCallEventType)() {
    m.callEventType = value
}
// SetCallId sets the callId property value. Unique identifier of the call.
func (m *CallStartedEventMessageDetail) SetCallId(value *string)() {
    m.callId = value
}
// SetInitiator sets the initiator property value. Initiator of the event.
func (m *CallStartedEventMessageDetail) SetInitiator(value IdentitySetable)() {
    m.initiator = value
}
