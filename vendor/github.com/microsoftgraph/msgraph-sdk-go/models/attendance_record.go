package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AttendanceRecord 
type AttendanceRecord struct {
    Entity
    // List of time periods between joining and leaving a meeting.
    attendanceIntervals []AttendanceIntervalable
    // Email address of the user associated with this atttendance record.
    emailAddress *string
    // Identity of the user associated with this atttendance record.
    identity Identityable
    // Role of the attendee. Possible values are: None, Attendee, Presenter, and Organizer.
    role *string
    // Total duration of the attendances in seconds.
    totalAttendanceInSeconds *int32
}
// NewAttendanceRecord instantiates a new attendanceRecord and sets the default values.
func NewAttendanceRecord()(*AttendanceRecord) {
    m := &AttendanceRecord{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAttendanceRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttendanceRecordFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAttendanceRecord(), nil
}
// GetAttendanceIntervals gets the attendanceIntervals property value. List of time periods between joining and leaving a meeting.
func (m *AttendanceRecord) GetAttendanceIntervals()([]AttendanceIntervalable) {
    return m.attendanceIntervals
}
// GetEmailAddress gets the emailAddress property value. Email address of the user associated with this atttendance record.
func (m *AttendanceRecord) GetEmailAddress()(*string) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttendanceRecord) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attendanceIntervals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttendanceIntervalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttendanceIntervalable, len(val))
            for i, v := range val {
                res[i] = v.(AttendanceIntervalable)
            }
            m.SetAttendanceIntervals(res)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["identity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(Identityable))
        }
        return nil
    }
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["totalAttendanceInSeconds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAttendanceInSeconds(val)
        }
        return nil
    }
    return res
}
// GetIdentity gets the identity property value. Identity of the user associated with this atttendance record.
func (m *AttendanceRecord) GetIdentity()(Identityable) {
    return m.identity
}
// GetRole gets the role property value. Role of the attendee. Possible values are: None, Attendee, Presenter, and Organizer.
func (m *AttendanceRecord) GetRole()(*string) {
    return m.role
}
// GetTotalAttendanceInSeconds gets the totalAttendanceInSeconds property value. Total duration of the attendances in seconds.
func (m *AttendanceRecord) GetTotalAttendanceInSeconds()(*int32) {
    return m.totalAttendanceInSeconds
}
// Serialize serializes information the current object
func (m *AttendanceRecord) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttendanceIntervals() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendanceIntervals()))
        for i, v := range m.GetAttendanceIntervals() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attendanceIntervals", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalAttendanceInSeconds", m.GetTotalAttendanceInSeconds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttendanceIntervals sets the attendanceIntervals property value. List of time periods between joining and leaving a meeting.
func (m *AttendanceRecord) SetAttendanceIntervals(value []AttendanceIntervalable)() {
    m.attendanceIntervals = value
}
// SetEmailAddress sets the emailAddress property value. Email address of the user associated with this atttendance record.
func (m *AttendanceRecord) SetEmailAddress(value *string)() {
    m.emailAddress = value
}
// SetIdentity sets the identity property value. Identity of the user associated with this atttendance record.
func (m *AttendanceRecord) SetIdentity(value Identityable)() {
    m.identity = value
}
// SetRole sets the role property value. Role of the attendee. Possible values are: None, Attendee, Presenter, and Organizer.
func (m *AttendanceRecord) SetRole(value *string)() {
    m.role = value
}
// SetTotalAttendanceInSeconds sets the totalAttendanceInSeconds property value. Total duration of the attendances in seconds.
func (m *AttendanceRecord) SetTotalAttendanceInSeconds(value *int32)() {
    m.totalAttendanceInSeconds = value
}
