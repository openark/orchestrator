package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RecurrencePattern 
type RecurrencePattern struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
    dayOfMonth *int32
    // A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
    daysOfWeek []DayOfWeek
    // The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
    firstDayOfWeek *DayOfWeek
    // Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
    index *WeekIndex
    // The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
    interval *int32
    // The month in which the event occurs.  This is a number from 1 to 12.
    month *int32
    // The OdataType property
    odataType *string
    // The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required. For more information, see values of type property.
    typeEscaped *RecurrencePatternType
}
// NewRecurrencePattern instantiates a new recurrencePattern and sets the default values.
func NewRecurrencePattern()(*RecurrencePattern) {
    m := &RecurrencePattern{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRecurrencePatternFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRecurrencePatternFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRecurrencePattern(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RecurrencePattern) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDayOfMonth gets the dayOfMonth property value. The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
func (m *RecurrencePattern) GetDayOfMonth()(*int32) {
    return m.dayOfMonth
}
// GetDaysOfWeek gets the daysOfWeek property value. A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
func (m *RecurrencePattern) GetDaysOfWeek()([]DayOfWeek) {
    return m.daysOfWeek
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RecurrencePattern) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dayOfMonth"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDayOfMonth(val)
        }
        return nil
    }
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
    res["firstDayOfWeek"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDayOfWeek)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstDayOfWeek(val.(*DayOfWeek))
        }
        return nil
    }
    res["index"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWeekIndex)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIndex(val.(*WeekIndex))
        }
        return nil
    }
    res["interval"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInterval(val)
        }
        return nil
    }
    res["month"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMonth(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecurrencePatternType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*RecurrencePatternType))
        }
        return nil
    }
    return res
}
// GetFirstDayOfWeek gets the firstDayOfWeek property value. The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
func (m *RecurrencePattern) GetFirstDayOfWeek()(*DayOfWeek) {
    return m.firstDayOfWeek
}
// GetIndex gets the index property value. Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
func (m *RecurrencePattern) GetIndex()(*WeekIndex) {
    return m.index
}
// GetInterval gets the interval property value. The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
func (m *RecurrencePattern) GetInterval()(*int32) {
    return m.interval
}
// GetMonth gets the month property value. The month in which the event occurs.  This is a number from 1 to 12.
func (m *RecurrencePattern) GetMonth()(*int32) {
    return m.month
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RecurrencePattern) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required. For more information, see values of type property.
func (m *RecurrencePattern) GetType()(*RecurrencePatternType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *RecurrencePattern) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("dayOfMonth", m.GetDayOfMonth())
        if err != nil {
            return err
        }
    }
    if m.GetDaysOfWeek() != nil {
        err := writer.WriteCollectionOfStringValues("daysOfWeek", SerializeDayOfWeek(m.GetDaysOfWeek()))
        if err != nil {
            return err
        }
    }
    if m.GetFirstDayOfWeek() != nil {
        cast := (*m.GetFirstDayOfWeek()).String()
        err := writer.WriteStringValue("firstDayOfWeek", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetIndex() != nil {
        cast := (*m.GetIndex()).String()
        err := writer.WriteStringValue("index", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("interval", m.GetInterval())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("month", m.GetMonth())
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *RecurrencePattern) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDayOfMonth sets the dayOfMonth property value. The day of the month on which the event occurs. Required if type is absoluteMonthly or absoluteYearly.
func (m *RecurrencePattern) SetDayOfMonth(value *int32)() {
    m.dayOfMonth = value
}
// SetDaysOfWeek sets the daysOfWeek property value. A collection of the days of the week on which the event occurs. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. If type is relativeMonthly or relativeYearly, and daysOfWeek specifies more than one day, the event falls on the first day that satisfies the pattern.  Required if type is weekly, relativeMonthly, or relativeYearly.
func (m *RecurrencePattern) SetDaysOfWeek(value []DayOfWeek)() {
    m.daysOfWeek = value
}
// SetFirstDayOfWeek sets the firstDayOfWeek property value. The first day of the week. The possible values are: sunday, monday, tuesday, wednesday, thursday, friday, saturday. Default is sunday. Required if type is weekly.
func (m *RecurrencePattern) SetFirstDayOfWeek(value *DayOfWeek)() {
    m.firstDayOfWeek = value
}
// SetIndex sets the index property value. Specifies on which instance of the allowed days specified in daysOfWeek the event occurs, counted from the first instance in the month. The possible values are: first, second, third, fourth, last. Default is first. Optional and used if type is relativeMonthly or relativeYearly.
func (m *RecurrencePattern) SetIndex(value *WeekIndex)() {
    m.index = value
}
// SetInterval sets the interval property value. The number of units between occurrences, where units can be in days, weeks, months, or years, depending on the type. Required.
func (m *RecurrencePattern) SetInterval(value *int32)() {
    m.interval = value
}
// SetMonth sets the month property value. The month in which the event occurs.  This is a number from 1 to 12.
func (m *RecurrencePattern) SetMonth(value *int32)() {
    m.month = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RecurrencePattern) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. The recurrence pattern type: daily, weekly, absoluteMonthly, relativeMonthly, absoluteYearly, relativeYearly. Required. For more information, see values of type property.
func (m *RecurrencePattern) SetType(value *RecurrencePatternType)() {
    m.typeEscaped = value
}
