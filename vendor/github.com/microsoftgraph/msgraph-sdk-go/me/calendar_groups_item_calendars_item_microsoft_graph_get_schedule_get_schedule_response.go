package me

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse 
type CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
    // The value property
    value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScheduleInformationable
}
// NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse instantiates a new CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse and sets the default values.
func NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse()(*CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse) {
    m := &CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScheduleInformationable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScheduleInformationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse) GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScheduleInformationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *CalendarGroupsItemCalendarsItemMicrosoftGraphGetScheduleGetScheduleResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ScheduleInformationable)() {
    m.value = value
}
