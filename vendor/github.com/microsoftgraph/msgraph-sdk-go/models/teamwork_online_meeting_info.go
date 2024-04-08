package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkOnlineMeetingInfo 
type TeamworkOnlineMeetingInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The identifier of the calendar event associated with the meeting.
    calendarEventId *string
    // The URL that users click to join or uniquely identify the meeting.
    joinWebUrl *string
    // The OdataType property
    odataType *string
    // The organizer of the meeting.
    organizer TeamworkUserIdentityable
}
// NewTeamworkOnlineMeetingInfo instantiates a new teamworkOnlineMeetingInfo and sets the default values.
func NewTeamworkOnlineMeetingInfo()(*TeamworkOnlineMeetingInfo) {
    m := &TeamworkOnlineMeetingInfo{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTeamworkOnlineMeetingInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamworkOnlineMeetingInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeamworkOnlineMeetingInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TeamworkOnlineMeetingInfo) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCalendarEventId gets the calendarEventId property value. The identifier of the calendar event associated with the meeting.
func (m *TeamworkOnlineMeetingInfo) GetCalendarEventId()(*string) {
    return m.calendarEventId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkOnlineMeetingInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calendarEventId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendarEventId(val)
        }
        return nil
    }
    res["joinWebUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
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
        val, err := n.GetObjectValue(CreateTeamworkUserIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(TeamworkUserIdentityable))
        }
        return nil
    }
    return res
}
// GetJoinWebUrl gets the joinWebUrl property value. The URL that users click to join or uniquely identify the meeting.
func (m *TeamworkOnlineMeetingInfo) GetJoinWebUrl()(*string) {
    return m.joinWebUrl
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TeamworkOnlineMeetingInfo) GetOdataType()(*string) {
    return m.odataType
}
// GetOrganizer gets the organizer property value. The organizer of the meeting.
func (m *TeamworkOnlineMeetingInfo) GetOrganizer()(TeamworkUserIdentityable) {
    return m.organizer
}
// Serialize serializes information the current object
func (m *TeamworkOnlineMeetingInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("calendarEventId", m.GetCalendarEventId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
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
func (m *TeamworkOnlineMeetingInfo) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCalendarEventId sets the calendarEventId property value. The identifier of the calendar event associated with the meeting.
func (m *TeamworkOnlineMeetingInfo) SetCalendarEventId(value *string)() {
    m.calendarEventId = value
}
// SetJoinWebUrl sets the joinWebUrl property value. The URL that users click to join or uniquely identify the meeting.
func (m *TeamworkOnlineMeetingInfo) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TeamworkOnlineMeetingInfo) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrganizer sets the organizer property value. The organizer of the meeting.
func (m *TeamworkOnlineMeetingInfo) SetOrganizer(value TeamworkUserIdentityable)() {
    m.organizer = value
}
