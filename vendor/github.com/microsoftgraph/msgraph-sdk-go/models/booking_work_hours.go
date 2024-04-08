package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BookingWorkHours this type represents the set of working hours in a single day of the week.
type BookingWorkHours struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The day property
    day *DayOfWeek
    // The OdataType property
    odataType *string
    // A list of start/end times during a day.
    timeSlots []BookingWorkTimeSlotable
}
// NewBookingWorkHours instantiates a new bookingWorkHours and sets the default values.
func NewBookingWorkHours()(*BookingWorkHours) {
    m := &BookingWorkHours{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBookingWorkHoursFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBookingWorkHoursFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBookingWorkHours(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BookingWorkHours) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDay gets the day property value. The day property
func (m *BookingWorkHours) GetDay()(*DayOfWeek) {
    return m.day
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BookingWorkHours) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["day"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDay(val.(*DayOfWeek))
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
    res["timeSlots"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBookingWorkTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BookingWorkTimeSlotable, len(val))
            for i, v := range val {
                res[i] = v.(BookingWorkTimeSlotable)
            }
            m.SetTimeSlots(res)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BookingWorkHours) GetOdataType()(*string) {
    return m.odataType
}
// GetTimeSlots gets the timeSlots property value. A list of start/end times during a day.
func (m *BookingWorkHours) GetTimeSlots()([]BookingWorkTimeSlotable) {
    return m.timeSlots
}
// Serialize serializes information the current object
func (m *BookingWorkHours) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDay() != nil {
        cast := (*m.GetDay()).String()
        err := writer.WriteStringValue("day", &cast)
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
    if m.GetTimeSlots() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTimeSlots()))
        for i, v := range m.GetTimeSlots() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("timeSlots", cast)
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
func (m *BookingWorkHours) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDay sets the day property value. The day property
func (m *BookingWorkHours) SetDay(value *DayOfWeek)() {
    m.day = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BookingWorkHours) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTimeSlots sets the timeSlots property value. A list of start/end times during a day.
func (m *BookingWorkHours) SetTimeSlots(value []BookingWorkTimeSlotable)() {
    m.timeSlots = value
}
