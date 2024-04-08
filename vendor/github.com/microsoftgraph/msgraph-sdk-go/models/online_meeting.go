package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnlineMeeting 
type OnlineMeeting struct {
    Entity
    // Indicates whether attendees can turn on their camera.
    allowAttendeeToEnableCamera *bool
    // Indicates whether attendees can turn on their microphone.
    allowAttendeeToEnableMic *bool
    // Indicates whether Teams reactions are enabled for the meeting.
    allowTeamworkReactions *bool
    // The attendance reports of an online meeting. Read-only.
    attendanceReports []MeetingAttendanceReportable
    // The content stream of the attendee report of a Microsoft Teams live event. Read-only.
    attendeeReport []byte
    // The phone access (dial-in) information for an online meeting. Read-only.
    audioConferencing AudioConferencingable
    // Settings related to a live event.
    broadcastSettings BroadcastMeetingSettingsable
    // The chat information associated with this online meeting.
    chatInfo ChatInfoable
    // The meeting creation time in UTC. Read-only.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The meeting end time in UTC.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The externalId property
    externalId *string
    // Indicates if this is a Teams live event.
    isBroadcast *bool
    // Indicates whether to announce when callers join or leave.
    isEntryExitAnnounced *bool
    // The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
    joinInformation ItemBodyable
    // Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is created, the joinMeetingIdSettings cannot be modified. To make any changes to this property, the meeting needs to be canceled and a new one needs to be created.
    joinMeetingIdSettings JoinMeetingIdSettingsable
    // The join URL of the online meeting. Read-only.
    joinWebUrl *string
    // Specifies which participants can bypass the meeting   lobby.
    lobbyBypassSettings LobbyBypassSettingsable
    // The participants associated with the online meeting.  This includes the organizer and the attendees.
    participants MeetingParticipantsable
    // Indicates whether to record the meeting automatically.
    recordAutomatically *bool
    // The meeting start time in UTC.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject of the online meeting.
    subject *string
    // The video teleconferencing ID. Read-only.
    videoTeleconferenceId *string
}
// NewOnlineMeeting instantiates a new onlineMeeting and sets the default values.
func NewOnlineMeeting()(*OnlineMeeting) {
    m := &OnlineMeeting{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnlineMeetingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnlineMeetingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnlineMeeting(), nil
}
// GetAllowAttendeeToEnableCamera gets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) GetAllowAttendeeToEnableCamera()(*bool) {
    return m.allowAttendeeToEnableCamera
}
// GetAllowAttendeeToEnableMic gets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) GetAllowAttendeeToEnableMic()(*bool) {
    return m.allowAttendeeToEnableMic
}
// GetAllowTeamworkReactions gets the allowTeamworkReactions property value. Indicates whether Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) GetAllowTeamworkReactions()(*bool) {
    return m.allowTeamworkReactions
}
// GetAttendanceReports gets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeeting) GetAttendanceReports()([]MeetingAttendanceReportable) {
    return m.attendanceReports
}
// GetAttendeeReport gets the attendeeReport property value. The content stream of the attendee report of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) GetAttendeeReport()([]byte) {
    return m.attendeeReport
}
// GetAudioConferencing gets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) GetAudioConferencing()(AudioConferencingable) {
    return m.audioConferencing
}
// GetBroadcastSettings gets the broadcastSettings property value. Settings related to a live event.
func (m *OnlineMeeting) GetBroadcastSettings()(BroadcastMeetingSettingsable) {
    return m.broadcastSettings
}
// GetChatInfo gets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) GetChatInfo()(ChatInfoable) {
    return m.chatInfo
}
// GetCreationDateTime gets the creationDateTime property value. The meeting creation time in UTC. Read-only.
func (m *OnlineMeeting) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.creationDateTime
}
// GetEndDateTime gets the endDateTime property value. The meeting end time in UTC.
func (m *OnlineMeeting) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetExternalId gets the externalId property value. The externalId property
func (m *OnlineMeeting) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnlineMeeting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowAttendeeToEnableCamera"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableCamera(val)
        }
        return nil
    }
    res["allowAttendeeToEnableMic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableMic(val)
        }
        return nil
    }
    res["allowTeamworkReactions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTeamworkReactions(val)
        }
        return nil
    }
    res["attendanceReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMeetingAttendanceReportFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MeetingAttendanceReportable, len(val))
            for i, v := range val {
                res[i] = v.(MeetingAttendanceReportable)
            }
            m.SetAttendanceReports(res)
        }
        return nil
    }
    res["attendeeReport"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttendeeReport(val)
        }
        return nil
    }
    res["audioConferencing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAudioConferencingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioConferencing(val.(AudioConferencingable))
        }
        return nil
    }
    res["broadcastSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBroadcastMeetingSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBroadcastSettings(val.(BroadcastMeetingSettingsable))
        }
        return nil
    }
    res["chatInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(ChatInfoable))
        }
        return nil
    }
    res["creationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
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
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["isBroadcast"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBroadcast(val)
        }
        return nil
    }
    res["isEntryExitAnnounced"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEntryExitAnnounced(val)
        }
        return nil
    }
    res["joinInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinInformation(val.(ItemBodyable))
        }
        return nil
    }
    res["joinMeetingIdSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJoinMeetingIdSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinMeetingIdSettings(val.(JoinMeetingIdSettingsable))
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
    res["lobbyBypassSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLobbyBypassSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLobbyBypassSettings(val.(LobbyBypassSettingsable))
        }
        return nil
    }
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMeetingParticipantsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipants(val.(MeetingParticipantsable))
        }
        return nil
    }
    res["recordAutomatically"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordAutomatically(val)
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
    res["videoTeleconferenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoTeleconferenceId(val)
        }
        return nil
    }
    return res
}
// GetIsBroadcast gets the isBroadcast property value. Indicates if this is a Teams live event.
func (m *OnlineMeeting) GetIsBroadcast()(*bool) {
    return m.isBroadcast
}
// GetIsEntryExitAnnounced gets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) GetIsEntryExitAnnounced()(*bool) {
    return m.isEntryExitAnnounced
}
// GetJoinInformation gets the joinInformation property value. The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
func (m *OnlineMeeting) GetJoinInformation()(ItemBodyable) {
    return m.joinInformation
}
// GetJoinMeetingIdSettings gets the joinMeetingIdSettings property value. Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is created, the joinMeetingIdSettings cannot be modified. To make any changes to this property, the meeting needs to be canceled and a new one needs to be created.
func (m *OnlineMeeting) GetJoinMeetingIdSettings()(JoinMeetingIdSettingsable) {
    return m.joinMeetingIdSettings
}
// GetJoinWebUrl gets the joinWebUrl property value. The join URL of the online meeting. Read-only.
func (m *OnlineMeeting) GetJoinWebUrl()(*string) {
    return m.joinWebUrl
}
// GetLobbyBypassSettings gets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting   lobby.
func (m *OnlineMeeting) GetLobbyBypassSettings()(LobbyBypassSettingsable) {
    return m.lobbyBypassSettings
}
// GetParticipants gets the participants property value. The participants associated with the online meeting.  This includes the organizer and the attendees.
func (m *OnlineMeeting) GetParticipants()(MeetingParticipantsable) {
    return m.participants
}
// GetRecordAutomatically gets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) GetRecordAutomatically()(*bool) {
    return m.recordAutomatically
}
// GetStartDateTime gets the startDateTime property value. The meeting start time in UTC.
func (m *OnlineMeeting) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetSubject gets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) GetSubject()(*string) {
    return m.subject
}
// GetVideoTeleconferenceId gets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) GetVideoTeleconferenceId()(*string) {
    return m.videoTeleconferenceId
}
// Serialize serializes information the current object
func (m *OnlineMeeting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAttendeeToEnableCamera", m.GetAllowAttendeeToEnableCamera())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowAttendeeToEnableMic", m.GetAllowAttendeeToEnableMic())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowTeamworkReactions", m.GetAllowTeamworkReactions())
        if err != nil {
            return err
        }
    }
    if m.GetAttendanceReports() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendanceReports()))
        for i, v := range m.GetAttendanceReports() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attendanceReports", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("attendeeReport", m.GetAttendeeReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("audioConferencing", m.GetAudioConferencing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("broadcastSettings", m.GetBroadcastSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("chatInfo", m.GetChatInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
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
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBroadcast", m.GetIsBroadcast())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEntryExitAnnounced", m.GetIsEntryExitAnnounced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("joinInformation", m.GetJoinInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("joinMeetingIdSettings", m.GetJoinMeetingIdSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lobbyBypassSettings", m.GetLobbyBypassSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("recordAutomatically", m.GetRecordAutomatically())
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
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("videoTeleconferenceId", m.GetVideoTeleconferenceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowAttendeeToEnableCamera sets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) SetAllowAttendeeToEnableCamera(value *bool)() {
    m.allowAttendeeToEnableCamera = value
}
// SetAllowAttendeeToEnableMic sets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) SetAllowAttendeeToEnableMic(value *bool)() {
    m.allowAttendeeToEnableMic = value
}
// SetAllowTeamworkReactions sets the allowTeamworkReactions property value. Indicates whether Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) SetAllowTeamworkReactions(value *bool)() {
    m.allowTeamworkReactions = value
}
// SetAttendanceReports sets the attendanceReports property value. The attendance reports of an online meeting. Read-only.
func (m *OnlineMeeting) SetAttendanceReports(value []MeetingAttendanceReportable)() {
    m.attendanceReports = value
}
// SetAttendeeReport sets the attendeeReport property value. The content stream of the attendee report of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) SetAttendeeReport(value []byte)() {
    m.attendeeReport = value
}
// SetAudioConferencing sets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) SetAudioConferencing(value AudioConferencingable)() {
    m.audioConferencing = value
}
// SetBroadcastSettings sets the broadcastSettings property value. Settings related to a live event.
func (m *OnlineMeeting) SetBroadcastSettings(value BroadcastMeetingSettingsable)() {
    m.broadcastSettings = value
}
// SetChatInfo sets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) SetChatInfo(value ChatInfoable)() {
    m.chatInfo = value
}
// SetCreationDateTime sets the creationDateTime property value. The meeting creation time in UTC. Read-only.
func (m *OnlineMeeting) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// SetEndDateTime sets the endDateTime property value. The meeting end time in UTC.
func (m *OnlineMeeting) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetExternalId sets the externalId property value. The externalId property
func (m *OnlineMeeting) SetExternalId(value *string)() {
    m.externalId = value
}
// SetIsBroadcast sets the isBroadcast property value. Indicates if this is a Teams live event.
func (m *OnlineMeeting) SetIsBroadcast(value *bool)() {
    m.isBroadcast = value
}
// SetIsEntryExitAnnounced sets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) SetIsEntryExitAnnounced(value *bool)() {
    m.isEntryExitAnnounced = value
}
// SetJoinInformation sets the joinInformation property value. The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
func (m *OnlineMeeting) SetJoinInformation(value ItemBodyable)() {
    m.joinInformation = value
}
// SetJoinMeetingIdSettings sets the joinMeetingIdSettings property value. Specifies the joinMeetingId, the meeting passcode, and the requirement for the passcode. Once an onlineMeeting is created, the joinMeetingIdSettings cannot be modified. To make any changes to this property, the meeting needs to be canceled and a new one needs to be created.
func (m *OnlineMeeting) SetJoinMeetingIdSettings(value JoinMeetingIdSettingsable)() {
    m.joinMeetingIdSettings = value
}
// SetJoinWebUrl sets the joinWebUrl property value. The join URL of the online meeting. Read-only.
func (m *OnlineMeeting) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// SetLobbyBypassSettings sets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting   lobby.
func (m *OnlineMeeting) SetLobbyBypassSettings(value LobbyBypassSettingsable)() {
    m.lobbyBypassSettings = value
}
// SetParticipants sets the participants property value. The participants associated with the online meeting.  This includes the organizer and the attendees.
func (m *OnlineMeeting) SetParticipants(value MeetingParticipantsable)() {
    m.participants = value
}
// SetRecordAutomatically sets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) SetRecordAutomatically(value *bool)() {
    m.recordAutomatically = value
}
// SetStartDateTime sets the startDateTime property value. The meeting start time in UTC.
func (m *OnlineMeeting) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetSubject sets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) SetSubject(value *string)() {
    m.subject = value
}
// SetVideoTeleconferenceId sets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) SetVideoTeleconferenceId(value *string)() {
    m.videoTeleconferenceId = value
}
