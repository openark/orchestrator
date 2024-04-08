package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingTimeSuggestionsResult 
type MeetingTimeSuggestionsResult struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A reason for not returning any meeting suggestions. The possible values are: attendeesUnavailable, attendeesUnavailableOrUnknown, locationsUnavailable, organizerUnavailable, or unknown. This property is an empty string if the meetingTimeSuggestions property does include any meeting suggestions.
    emptySuggestionsReason *string
    // An array of meeting suggestions.
    meetingTimeSuggestions []MeetingTimeSuggestionable
    // The OdataType property
    odataType *string
}
// NewMeetingTimeSuggestionsResult instantiates a new meetingTimeSuggestionsResult and sets the default values.
func NewMeetingTimeSuggestionsResult()(*MeetingTimeSuggestionsResult) {
    m := &MeetingTimeSuggestionsResult{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMeetingTimeSuggestionsResultFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingTimeSuggestionsResultFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingTimeSuggestionsResult(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingTimeSuggestionsResult) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetEmptySuggestionsReason gets the emptySuggestionsReason property value. A reason for not returning any meeting suggestions. The possible values are: attendeesUnavailable, attendeesUnavailableOrUnknown, locationsUnavailable, organizerUnavailable, or unknown. This property is an empty string if the meetingTimeSuggestions property does include any meeting suggestions.
func (m *MeetingTimeSuggestionsResult) GetEmptySuggestionsReason()(*string) {
    return m.emptySuggestionsReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingTimeSuggestionsResult) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["emptySuggestionsReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmptySuggestionsReason(val)
        }
        return nil
    }
    res["meetingTimeSuggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingTimeSuggestionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingTimeSuggestionable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingTimeSuggestionable)
            }
            m.SetMeetingTimeSuggestions(res)
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
    return res
}
// GetMeetingTimeSuggestions gets the meetingTimeSuggestions property value. An array of meeting suggestions.
func (m *MeetingTimeSuggestionsResult) GetMeetingTimeSuggestions()([]MeetingTimeSuggestionable) {
    return m.meetingTimeSuggestions
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MeetingTimeSuggestionsResult) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *MeetingTimeSuggestionsResult) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("emptySuggestionsReason", m.GetEmptySuggestionsReason())
        if err != nil {
            return err
        }
    }
    if m.GetMeetingTimeSuggestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMeetingTimeSuggestions()))
        for i, v := range m.GetMeetingTimeSuggestions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("meetingTimeSuggestions", cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingTimeSuggestionsResult) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetEmptySuggestionsReason sets the emptySuggestionsReason property value. A reason for not returning any meeting suggestions. The possible values are: attendeesUnavailable, attendeesUnavailableOrUnknown, locationsUnavailable, organizerUnavailable, or unknown. This property is an empty string if the meetingTimeSuggestions property does include any meeting suggestions.
func (m *MeetingTimeSuggestionsResult) SetEmptySuggestionsReason(value *string)() {
    m.emptySuggestionsReason = value
}
// SetMeetingTimeSuggestions sets the meetingTimeSuggestions property value. An array of meeting suggestions.
func (m *MeetingTimeSuggestionsResult) SetMeetingTimeSuggestions(value []MeetingTimeSuggestionable)() {
    m.meetingTimeSuggestions = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MeetingTimeSuggestionsResult) SetOdataType(value *string)() {
    m.odataType = value
}
