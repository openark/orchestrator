package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ResponseStatus 
type ResponseStatus struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The response type. Possible values are: none, organizer, tentativelyAccepted, accepted, declined, notResponded.To differentiate between none and notResponded:  none – from organizer's perspective. This value is used when the status of an attendee/participant is reported to the organizer of a meeting.  notResponded – from attendee's perspective. Indicates the attendee has not responded to the meeting request.  Clients can treat notResponded == none.  As an example, if attendee Alex hasn't responded to a meeting request, getting Alex' response status for that event in Alex' calendar returns notResponded. Getting Alex' response from the calendar of any other attendee or the organizer's returns none. Getting the organizer's response for the event in anybody's calendar also returns none.
    response *ResponseType
    // The date and time when the response was returned. It uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    time *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewResponseStatus instantiates a new responseStatus and sets the default values.
func NewResponseStatus()(*ResponseStatus) {
    m := &ResponseStatus{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateResponseStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResponseStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewResponseStatus(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResponseStatus) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResponseStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["response"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseResponseType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponse(val.(*ResponseType))
        }
        return nil
    }
    res["time"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTime(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ResponseStatus) GetOdataType()(*string) {
    return m.odataType
}
// GetResponse gets the response property value. The response type. Possible values are: none, organizer, tentativelyAccepted, accepted, declined, notResponded.To differentiate between none and notResponded:  none – from organizer's perspective. This value is used when the status of an attendee/participant is reported to the organizer of a meeting.  notResponded – from attendee's perspective. Indicates the attendee has not responded to the meeting request.  Clients can treat notResponded == none.  As an example, if attendee Alex hasn't responded to a meeting request, getting Alex' response status for that event in Alex' calendar returns notResponded. Getting Alex' response from the calendar of any other attendee or the organizer's returns none. Getting the organizer's response for the event in anybody's calendar also returns none.
func (m *ResponseStatus) GetResponse()(*ResponseType) {
    return m.response
}
// GetTime gets the time property value. The date and time when the response was returned. It uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ResponseStatus) GetTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.time
}
// Serialize serializes information the current object
func (m *ResponseStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetResponse() != nil {
        cast := (*m.GetResponse()).String()
        err := writer.WriteStringValue("response", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("time", m.GetTime())
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
func (m *ResponseStatus) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ResponseStatus) SetOdataType(value *string)() {
    m.odataType = value
}
// SetResponse sets the response property value. The response type. Possible values are: none, organizer, tentativelyAccepted, accepted, declined, notResponded.To differentiate between none and notResponded:  none – from organizer's perspective. This value is used when the status of an attendee/participant is reported to the organizer of a meeting.  notResponded – from attendee's perspective. Indicates the attendee has not responded to the meeting request.  Clients can treat notResponded == none.  As an example, if attendee Alex hasn't responded to a meeting request, getting Alex' response status for that event in Alex' calendar returns notResponded. Getting Alex' response from the calendar of any other attendee or the organizer's returns none. Getting the organizer's response for the event in anybody's calendar also returns none.
func (m *ResponseStatus) SetResponse(value *ResponseType)() {
    m.response = value
}
// SetTime sets the time property value. The date and time when the response was returned. It uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ResponseStatus) SetTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.time = value
}
