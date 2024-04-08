package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ScheduleItem 
type ScheduleItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The date, time, and time zone that the corresponding event ends.
    end DateTimeTimeZoneable
    // The sensitivity of the corresponding event. True if the event is marked private, false otherwise. Optional.
    isPrivate *bool
    // The location where the corresponding event is held or attended from. Optional.
    location *string
    // The OdataType property
    odataType *string
    // The date, time, and time zone that the corresponding event starts.
    start DateTimeTimeZoneable
    // The availability status of the user or resource during the corresponding event. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
    status *FreeBusyStatus
    // The corresponding event's subject line. Optional.
    subject *string
}
// NewScheduleItem instantiates a new scheduleItem and sets the default values.
func NewScheduleItem()(*ScheduleItem) {
    m := &ScheduleItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateScheduleItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateScheduleItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewScheduleItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ScheduleItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEnd gets the end property value. The date, time, and time zone that the corresponding event ends.
func (m *ScheduleItem) GetEnd()(DateTimeTimeZoneable) {
    return m.end
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ScheduleItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["isPrivate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPrivate(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val)
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
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*FreeBusyStatus))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    return res
}
// GetIsPrivate gets the isPrivate property value. The sensitivity of the corresponding event. True if the event is marked private, false otherwise. Optional.
func (m *ScheduleItem) GetIsPrivate()(*bool) {
    return m.isPrivate
}
// GetLocation gets the location property value. The location where the corresponding event is held or attended from. Optional.
func (m *ScheduleItem) GetLocation()(*string) {
    return m.location
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ScheduleItem) GetOdataType()(*string) {
    return m.odataType
}
// GetStart gets the start property value. The date, time, and time zone that the corresponding event starts.
func (m *ScheduleItem) GetStart()(DateTimeTimeZoneable) {
    return m.start
}
// GetStatus gets the status property value. The availability status of the user or resource during the corresponding event. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *ScheduleItem) GetStatus()(*FreeBusyStatus) {
    return m.status
}
// GetSubject gets the subject property value. The corresponding event's subject line. Optional.
func (m *ScheduleItem) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *ScheduleItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPrivate", m.GetIsPrivate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("location", m.GetLocation())
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
        err := writer.WriteObjectValue("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subject", m.GetSubject())
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
func (m *ScheduleItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEnd sets the end property value. The date, time, and time zone that the corresponding event ends.
func (m *ScheduleItem) SetEnd(value DateTimeTimeZoneable)() {
    m.end = value
}
// SetIsPrivate sets the isPrivate property value. The sensitivity of the corresponding event. True if the event is marked private, false otherwise. Optional.
func (m *ScheduleItem) SetIsPrivate(value *bool)() {
    m.isPrivate = value
}
// SetLocation sets the location property value. The location where the corresponding event is held or attended from. Optional.
func (m *ScheduleItem) SetLocation(value *string)() {
    m.location = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ScheduleItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStart sets the start property value. The date, time, and time zone that the corresponding event starts.
func (m *ScheduleItem) SetStart(value DateTimeTimeZoneable)() {
    m.start = value
}
// SetStatus sets the status property value. The availability status of the user or resource during the corresponding event. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *ScheduleItem) SetStatus(value *FreeBusyStatus)() {
    m.status = value
}
// SetSubject sets the subject property value. The corresponding event's subject line. Optional.
func (m *ScheduleItem) SetSubject(value *string)() {
    m.subject = value
}
