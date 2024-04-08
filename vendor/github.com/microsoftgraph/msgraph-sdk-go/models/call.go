package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Call 
type Call struct {
    Entity
    // The audioRoutingGroups property
    audioRoutingGroups []AudioRoutingGroupable
    // The callback URL on which callbacks will be delivered. Must be https.
    callbackUri *string
    // A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in a P2P call.  This needs to be copied over from Microsoft.Graph.Call.CallChainId.
    callChainId *string
    // Contains the optional features for the call.
    callOptions CallOptionsable
    // The routing information on how the call was retargeted. Read-only.
    callRoutes []CallRouteable
    // The chat information. Required information for joining a meeting.
    chatInfo ChatInfoable
    // The contentSharingSessions property
    contentSharingSessions []ContentSharingSessionable
    // The direction of the call. The possible value are incoming or outgoing. Read-only.
    direction *CallDirection
    // Call context associated with an incoming call.
    incomingContext IncomingContextable
    // The media configuration. Required.
    mediaConfig MediaConfigable
    // Read-only. The call media state.
    mediaState CallMediaStateable
    // The meeting information. Required information for meeting scenarios.
    meetingInfo MeetingInfoable
    // The myParticipantId property
    myParticipantId *string
    // The operations property
    operations []CommsOperationable
    // The participants property
    participants []Participantable
    // The list of requested modalities. Possible values are: unknown, audio, video, videoBasedScreenSharing, data.
    requestedModalities []Modality
    // The result information. For example can hold termination reason. Read-only.
    resultInfo ResultInfoable
    // The originator of the call.
    source ParticipantInfoable
    // The call state. Possible values are: incoming, establishing, ringing, established, hold, transferring, transferAccepted, redirecting, terminating, terminated. Read-only.
    state *CallState
    // The subject of the conversation.
    subject *string
    // The targets of the call. Required information for creating peer to peer call.
    targets []InvitationParticipantInfoable
    // The tenantId property
    tenantId *string
    // The toneInfo property
    toneInfo ToneInfoable
    // The transcription information for the call. Read-only.
    transcription CallTranscriptionInfoable
}
// NewCall instantiates a new call and sets the default values.
func NewCall()(*Call) {
    m := &Call{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCallFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCallFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCall(), nil
}
// GetAudioRoutingGroups gets the audioRoutingGroups property value. The audioRoutingGroups property
func (m *Call) GetAudioRoutingGroups()([]AudioRoutingGroupable) {
    return m.audioRoutingGroups
}
// GetCallbackUri gets the callbackUri property value. The callback URL on which callbacks will be delivered. Must be https.
func (m *Call) GetCallbackUri()(*string) {
    return m.callbackUri
}
// GetCallChainId gets the callChainId property value. A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in a P2P call.  This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *Call) GetCallChainId()(*string) {
    return m.callChainId
}
// GetCallOptions gets the callOptions property value. Contains the optional features for the call.
func (m *Call) GetCallOptions()(CallOptionsable) {
    return m.callOptions
}
// GetCallRoutes gets the callRoutes property value. The routing information on how the call was retargeted. Read-only.
func (m *Call) GetCallRoutes()([]CallRouteable) {
    return m.callRoutes
}
// GetChatInfo gets the chatInfo property value. The chat information. Required information for joining a meeting.
func (m *Call) GetChatInfo()(ChatInfoable) {
    return m.chatInfo
}
// GetContentSharingSessions gets the contentSharingSessions property value. The contentSharingSessions property
func (m *Call) GetContentSharingSessions()([]ContentSharingSessionable) {
    return m.contentSharingSessions
}
// GetDirection gets the direction property value. The direction of the call. The possible value are incoming or outgoing. Read-only.
func (m *Call) GetDirection()(*CallDirection) {
    return m.direction
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Call) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["audioRoutingGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAudioRoutingGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AudioRoutingGroupable, len(val))
            for i, v := range val {
                res[i] = v.(AudioRoutingGroupable)
            }
            m.SetAudioRoutingGroups(res)
        }
        return nil
    }
    res["callbackUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallbackUri(val)
        }
        return nil
    }
    res["callChainId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallChainId(val)
        }
        return nil
    }
    res["callOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCallOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallOptions(val.(CallOptionsable))
        }
        return nil
    }
    res["callRoutes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCallRouteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CallRouteable, len(val))
            for i, v := range val {
                res[i] = v.(CallRouteable)
            }
            m.SetCallRoutes(res)
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
    res["contentSharingSessions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContentSharingSessionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContentSharingSessionable, len(val))
            for i, v := range val {
                res[i] = v.(ContentSharingSessionable)
            }
            m.SetContentSharingSessions(res)
        }
        return nil
    }
    res["direction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallDirection)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirection(val.(*CallDirection))
        }
        return nil
    }
    res["incomingContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIncomingContextFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncomingContext(val.(IncomingContextable))
        }
        return nil
    }
    res["mediaConfig"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMediaConfigFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaConfig(val.(MediaConfigable))
        }
        return nil
    }
    res["mediaState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCallMediaStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaState(val.(CallMediaStateable))
        }
        return nil
    }
    res["meetingInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMeetingInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingInfo(val.(MeetingInfoable))
        }
        return nil
    }
    res["myParticipantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMyParticipantId(val)
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCommsOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CommsOperationable, len(val))
            for i, v := range val {
                res[i] = v.(CommsOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateParticipantFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Participantable, len(val))
            for i, v := range val {
                res[i] = v.(Participantable)
            }
            m.SetParticipants(res)
        }
        return nil
    }
    res["requestedModalities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseModality)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Modality, len(val))
            for i, v := range val {
                res[i] = *(v.(*Modality))
            }
            m.SetRequestedModalities(res)
        }
        return nil
    }
    res["resultInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResultInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultInfo(val.(ResultInfoable))
        }
        return nil
    }
    res["source"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSource(val.(ParticipantInfoable))
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCallState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*CallState))
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
    res["targets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInvitationParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InvitationParticipantInfoable, len(val))
            for i, v := range val {
                res[i] = v.(InvitationParticipantInfoable)
            }
            m.SetTargets(res)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["toneInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateToneInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetToneInfo(val.(ToneInfoable))
        }
        return nil
    }
    res["transcription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCallTranscriptionInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTranscription(val.(CallTranscriptionInfoable))
        }
        return nil
    }
    return res
}
// GetIncomingContext gets the incomingContext property value. Call context associated with an incoming call.
func (m *Call) GetIncomingContext()(IncomingContextable) {
    return m.incomingContext
}
// GetMediaConfig gets the mediaConfig property value. The media configuration. Required.
func (m *Call) GetMediaConfig()(MediaConfigable) {
    return m.mediaConfig
}
// GetMediaState gets the mediaState property value. Read-only. The call media state.
func (m *Call) GetMediaState()(CallMediaStateable) {
    return m.mediaState
}
// GetMeetingInfo gets the meetingInfo property value. The meeting information. Required information for meeting scenarios.
func (m *Call) GetMeetingInfo()(MeetingInfoable) {
    return m.meetingInfo
}
// GetMyParticipantId gets the myParticipantId property value. The myParticipantId property
func (m *Call) GetMyParticipantId()(*string) {
    return m.myParticipantId
}
// GetOperations gets the operations property value. The operations property
func (m *Call) GetOperations()([]CommsOperationable) {
    return m.operations
}
// GetParticipants gets the participants property value. The participants property
func (m *Call) GetParticipants()([]Participantable) {
    return m.participants
}
// GetRequestedModalities gets the requestedModalities property value. The list of requested modalities. Possible values are: unknown, audio, video, videoBasedScreenSharing, data.
func (m *Call) GetRequestedModalities()([]Modality) {
    return m.requestedModalities
}
// GetResultInfo gets the resultInfo property value. The result information. For example can hold termination reason. Read-only.
func (m *Call) GetResultInfo()(ResultInfoable) {
    return m.resultInfo
}
// GetSource gets the source property value. The originator of the call.
func (m *Call) GetSource()(ParticipantInfoable) {
    return m.source
}
// GetState gets the state property value. The call state. Possible values are: incoming, establishing, ringing, established, hold, transferring, transferAccepted, redirecting, terminating, terminated. Read-only.
func (m *Call) GetState()(*CallState) {
    return m.state
}
// GetSubject gets the subject property value. The subject of the conversation.
func (m *Call) GetSubject()(*string) {
    return m.subject
}
// GetTargets gets the targets property value. The targets of the call. Required information for creating peer to peer call.
func (m *Call) GetTargets()([]InvitationParticipantInfoable) {
    return m.targets
}
// GetTenantId gets the tenantId property value. The tenantId property
func (m *Call) GetTenantId()(*string) {
    return m.tenantId
}
// GetToneInfo gets the toneInfo property value. The toneInfo property
func (m *Call) GetToneInfo()(ToneInfoable) {
    return m.toneInfo
}
// GetTranscription gets the transcription property value. The transcription information for the call. Read-only.
func (m *Call) GetTranscription()(CallTranscriptionInfoable) {
    return m.transcription
}
// Serialize serializes information the current object
func (m *Call) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAudioRoutingGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAudioRoutingGroups()))
        for i, v := range m.GetAudioRoutingGroups() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("audioRoutingGroups", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callbackUri", m.GetCallbackUri())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("callChainId", m.GetCallChainId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("callOptions", m.GetCallOptions())
        if err != nil {
            return err
        }
    }
    if m.GetCallRoutes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCallRoutes()))
        for i, v := range m.GetCallRoutes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("callRoutes", cast)
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
    if m.GetContentSharingSessions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentSharingSessions()))
        for i, v := range m.GetContentSharingSessions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("contentSharingSessions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDirection() != nil {
        cast := (*m.GetDirection()).String()
        err = writer.WriteStringValue("direction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("incomingContext", m.GetIncomingContext())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaConfig", m.GetMediaConfig())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("mediaState", m.GetMediaState())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("meetingInfo", m.GetMeetingInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("myParticipantId", m.GetMyParticipantId())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetParticipants() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetParticipants()))
        for i, v := range m.GetParticipants() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("participants", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRequestedModalities() != nil {
        err = writer.WriteCollectionOfStringValues("requestedModalities", SerializeModality(m.GetRequestedModalities()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resultInfo", m.GetResultInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("source", m.GetSource())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
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
    if m.GetTargets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTargets()))
        for i, v := range m.GetTargets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("targets", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("toneInfo", m.GetToneInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("transcription", m.GetTranscription())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAudioRoutingGroups sets the audioRoutingGroups property value. The audioRoutingGroups property
func (m *Call) SetAudioRoutingGroups(value []AudioRoutingGroupable)() {
    m.audioRoutingGroups = value
}
// SetCallbackUri sets the callbackUri property value. The callback URL on which callbacks will be delivered. Must be https.
func (m *Call) SetCallbackUri(value *string)() {
    m.callbackUri = value
}
// SetCallChainId sets the callChainId property value. A unique identifier for all the participant calls in a conference or a unique identifier for two participant calls in a P2P call.  This needs to be copied over from Microsoft.Graph.Call.CallChainId.
func (m *Call) SetCallChainId(value *string)() {
    m.callChainId = value
}
// SetCallOptions sets the callOptions property value. Contains the optional features for the call.
func (m *Call) SetCallOptions(value CallOptionsable)() {
    m.callOptions = value
}
// SetCallRoutes sets the callRoutes property value. The routing information on how the call was retargeted. Read-only.
func (m *Call) SetCallRoutes(value []CallRouteable)() {
    m.callRoutes = value
}
// SetChatInfo sets the chatInfo property value. The chat information. Required information for joining a meeting.
func (m *Call) SetChatInfo(value ChatInfoable)() {
    m.chatInfo = value
}
// SetContentSharingSessions sets the contentSharingSessions property value. The contentSharingSessions property
func (m *Call) SetContentSharingSessions(value []ContentSharingSessionable)() {
    m.contentSharingSessions = value
}
// SetDirection sets the direction property value. The direction of the call. The possible value are incoming or outgoing. Read-only.
func (m *Call) SetDirection(value *CallDirection)() {
    m.direction = value
}
// SetIncomingContext sets the incomingContext property value. Call context associated with an incoming call.
func (m *Call) SetIncomingContext(value IncomingContextable)() {
    m.incomingContext = value
}
// SetMediaConfig sets the mediaConfig property value. The media configuration. Required.
func (m *Call) SetMediaConfig(value MediaConfigable)() {
    m.mediaConfig = value
}
// SetMediaState sets the mediaState property value. Read-only. The call media state.
func (m *Call) SetMediaState(value CallMediaStateable)() {
    m.mediaState = value
}
// SetMeetingInfo sets the meetingInfo property value. The meeting information. Required information for meeting scenarios.
func (m *Call) SetMeetingInfo(value MeetingInfoable)() {
    m.meetingInfo = value
}
// SetMyParticipantId sets the myParticipantId property value. The myParticipantId property
func (m *Call) SetMyParticipantId(value *string)() {
    m.myParticipantId = value
}
// SetOperations sets the operations property value. The operations property
func (m *Call) SetOperations(value []CommsOperationable)() {
    m.operations = value
}
// SetParticipants sets the participants property value. The participants property
func (m *Call) SetParticipants(value []Participantable)() {
    m.participants = value
}
// SetRequestedModalities sets the requestedModalities property value. The list of requested modalities. Possible values are: unknown, audio, video, videoBasedScreenSharing, data.
func (m *Call) SetRequestedModalities(value []Modality)() {
    m.requestedModalities = value
}
// SetResultInfo sets the resultInfo property value. The result information. For example can hold termination reason. Read-only.
func (m *Call) SetResultInfo(value ResultInfoable)() {
    m.resultInfo = value
}
// SetSource sets the source property value. The originator of the call.
func (m *Call) SetSource(value ParticipantInfoable)() {
    m.source = value
}
// SetState sets the state property value. The call state. Possible values are: incoming, establishing, ringing, established, hold, transferring, transferAccepted, redirecting, terminating, terminated. Read-only.
func (m *Call) SetState(value *CallState)() {
    m.state = value
}
// SetSubject sets the subject property value. The subject of the conversation.
func (m *Call) SetSubject(value *string)() {
    m.subject = value
}
// SetTargets sets the targets property value. The targets of the call. Required information for creating peer to peer call.
func (m *Call) SetTargets(value []InvitationParticipantInfoable)() {
    m.targets = value
}
// SetTenantId sets the tenantId property value. The tenantId property
func (m *Call) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetToneInfo sets the toneInfo property value. The toneInfo property
func (m *Call) SetToneInfo(value ToneInfoable)() {
    m.toneInfo = value
}
// SetTranscription sets the transcription property value. The transcription information for the call. Read-only.
func (m *Call) SetTranscription(value CallTranscriptionInfoable)() {
    m.transcription = value
}
