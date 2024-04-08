package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingParticipants 
type MeetingParticipants struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The attendees property
    attendees []MeetingParticipantInfoable
    // The OdataType property
    odataType *string
    // The organizer property
    organizer MeetingParticipantInfoable
}
// NewMeetingParticipants instantiates a new meetingParticipants and sets the default values.
func NewMeetingParticipants()(*MeetingParticipants) {
    m := &MeetingParticipants{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMeetingParticipantsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingParticipantsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingParticipants(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingParticipants) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAttendees gets the attendees property value. The attendees property
func (m *MeetingParticipants) GetAttendees()([]MeetingParticipantInfoable) {
    return m.attendees
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingParticipants) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingParticipantInfoable)
            }
            m.SetAttendees(res)
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
    res["organizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMeetingParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(MeetingParticipantInfoable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MeetingParticipants) GetOdataType()(*string) {
    return m.odataType
}
// GetOrganizer gets the organizer property value. The organizer property
func (m *MeetingParticipants) GetOrganizer()(MeetingParticipantInfoable) {
    return m.organizer
}
// Serialize serializes information the current object
func (m *MeetingParticipants) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("attendees", cast)
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
        err := writer.WriteObjectValue("organizer", m.GetOrganizer())
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
func (m *MeetingParticipants) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAttendees sets the attendees property value. The attendees property
func (m *MeetingParticipants) SetAttendees(value []MeetingParticipantInfoable)() {
    m.attendees = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingParticipants) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrganizer sets the organizer property value. The organizer property
func (m *MeetingParticipants) SetOrganizer(value MeetingParticipantInfoable)() {
    m.organizer = value
}
