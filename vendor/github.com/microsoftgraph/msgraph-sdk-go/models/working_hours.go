package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkingHours 
type WorkingHours struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The days of the week on which the user works.
    daysOfWeek []DayOfWeek
    // The time of the day that the user stops working.
    endTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The OdataType property
    odataType *string
    // The time of the day that the user starts working.
    startTime *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly
    // The time zone to which the working hours apply.
    timeZone TimeZoneBaseable
}
// NewWorkingHours instantiates a new workingHours and sets the default values.
func NewWorkingHours()(*WorkingHours) {
    m := &WorkingHours{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkingHoursFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkingHoursFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkingHours(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkingHours) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDaysOfWeek gets the daysOfWeek property value. The days of the week on which the user works.
func (m *WorkingHours) GetDaysOfWeek()([]DayOfWeek) {
    return m.daysOfWeek
}
// GetEndTime gets the endTime property value. The time of the day that the user stops working.
func (m *WorkingHours) GetEndTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.endTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkingHours) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["daysOfWeek"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DayOfWeek, len(val))
            for i, v := range val {
                res[i] = *(v.(*DayOfWeek))
            }
            m.SetDaysOfWeek(res)
        }
        return nil
    }
    res["endTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndTime(val)
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
    res["startTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartTime(val)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeZoneBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val.(TimeZoneBaseable))
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkingHours) GetOdataType()(*string) {
    return m.odataType
}
// GetStartTime gets the startTime property value. The time of the day that the user starts working.
func (m *WorkingHours) GetStartTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly) {
    return m.startTime
}
// GetTimeZone gets the timeZone property value. The time zone to which the working hours apply.
func (m *WorkingHours) GetTimeZone()(TimeZoneBaseable) {
    return m.timeZone
}
// Serialize serializes information the current object
func (m *WorkingHours) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDaysOfWeek() != nil {
        err := writer.WriteCollectionOfStringValues("daysOfWeek", SerializeDayOfWeek(m.GetDaysOfWeek()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeOnlyValue("endTime", m.GetEndTime())
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
        err := writer.WriteTimeOnlyValue("startTime", m.GetStartTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("timeZone", m.GetTimeZone())
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
func (m *WorkingHours) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDaysOfWeek sets the daysOfWeek property value. The days of the week on which the user works.
func (m *WorkingHours) SetDaysOfWeek(value []DayOfWeek)() {
    m.daysOfWeek = value
}
// SetEndTime sets the endTime property value. The time of the day that the user stops working.
func (m *WorkingHours) SetEndTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.endTime = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkingHours) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStartTime sets the startTime property value. The time of the day that the user starts working.
func (m *WorkingHours) SetStartTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)() {
    m.startTime = value
}
// SetTimeZone sets the timeZone property value. The time zone to which the working hours apply.
func (m *WorkingHours) SetTimeZone(value TimeZoneBaseable)() {
    m.timeZone = value
}
