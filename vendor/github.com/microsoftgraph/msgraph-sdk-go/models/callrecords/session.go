package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Session 
type Session struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
    // Endpoint that answered the session.
    callee Endpointable
    // Endpoint that initiated the session.
    caller Endpointable
    // UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Failure information associated with the session if the session failed.
    failureInfo FailureInfoable
    // List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
    modalities []Modality
    // The list of segments involved in the session. Read-only. Nullable.
    segments []Segmentable
    // UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewSession instantiates a new session and sets the default values.
func NewSession()(*Session) {
    m := &Session{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateSessionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSessionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSession(), nil
}
// GetCallee gets the callee property value. Endpoint that answered the session.
func (m *Session) GetCallee()(Endpointable) {
    return m.callee
}
// GetCaller gets the caller property value. Endpoint that initiated the session.
func (m *Session) GetCaller()(Endpointable) {
    return m.caller
}
// GetEndDateTime gets the endDateTime property value. UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Session) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFailureInfo gets the failureInfo property value. Failure information associated with the session if the session failed.
func (m *Session) GetFailureInfo()(FailureInfoable) {
    return m.failureInfo
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Session) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["callee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallee(val.(Endpointable))
        }
        return nil
    }
    res["caller"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCaller(val.(Endpointable))
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["failureInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFailureInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureInfo(val.(FailureInfoable))
        }
        return nil
    }
    res["modalities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Modality, len(val))
            for i, v := range val {
                res[i] = *(v.(*Modality))
            }
            m.SetModalities(res)
        }
        return nil
    }
    res["segments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSegmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Segmentable, len(val))
            for i, v := range val {
                res[i] = v.(Segmentable)
            }
            m.SetSegments(res)
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    return res
}
// GetModalities gets the modalities property value. List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
func (m *Session) GetModalities()([]Modality) {
    return m.modalities
}
// GetSegments gets the segments property value. The list of segments involved in the session. Read-only. Nullable.
func (m *Session) GetSegments()([]Segmentable) {
    return m.segments
}
// GetStartDateTime gets the startDateTime property value. UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Session) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// Serialize serializes information the current object
func (m *Session) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("callee", m.GetCallee())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("caller", m.GetCaller())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("failureInfo", m.GetFailureInfo())
        if err != nil {
            return err
        }
    }
    if m.GetModalities() != nil {
        err = writer.WriteCollectionOfStringValues("modalities", SerializeModality(m.GetModalities()))
        if err != nil {
            return err
        }
    }
    if m.GetSegments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSegments()))
        for i, v := range m.GetSegments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("segments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCallee sets the callee property value. Endpoint that answered the session.
func (m *Session) SetCallee(value Endpointable)() {
    m.callee = value
}
// SetCaller sets the caller property value. Endpoint that initiated the session.
func (m *Session) SetCaller(value Endpointable)() {
    m.caller = value
}
// SetEndDateTime sets the endDateTime property value. UTC time when the last user left the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Session) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetFailureInfo sets the failureInfo property value. Failure information associated with the session if the session failed.
func (m *Session) SetFailureInfo(value FailureInfoable)() {
    m.failureInfo = value
}
// SetModalities sets the modalities property value. List of modalities present in the session. Possible values are: unknown, audio, video, videoBasedScreenSharing, data, screenSharing, unknownFutureValue.
func (m *Session) SetModalities(value []Modality)() {
    m.modalities = value
}
// SetSegments sets the segments property value. The list of segments involved in the session. Read-only. Nullable.
func (m *Session) SetSegments(value []Segmentable)() {
    m.segments = value
}
// SetStartDateTime sets the startDateTime property value. UTC time when the first user joined the session. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Session) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
