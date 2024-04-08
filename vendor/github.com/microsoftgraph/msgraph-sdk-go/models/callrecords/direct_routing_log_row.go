package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectRoutingLogRow 
type DirectRoutingLogRow struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Number of the user or bot who received the call. E.164 format, but may include additional data.
    calleeNumber *string
    // In addition to the SIP codes, Microsoft has own subcodes that indicate the specific issue.
    callEndSubReason *int32
    // Number of the user or bot who made the call. E.164 format, but may include additional data.
    callerNumber *string
    // Call type and direction.
    callType *string
    // Identifier for the call that you can use when calling Microsoft Support. GUID.
    correlationId *string
    // Duration of the call in seconds.
    duration *int32
    // Only exists for successful (fully established) calls. Time when call ended.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Only exists for failed (not fully established) calls.
    failureDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The code with which the call ended, RFC 3261.
    finalSipCode *int32
    // Description of the SIP code and Microsoft subcode.
    finalSipCodePhrase *string
    // Unique call identifier. GUID.
    id *string
    // When the initial invite was sent.
    inviteDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Indicates if the trunk was enabled for media bypass or not.
    mediaBypassEnabled *bool
    // The datacenter used for media path in non-bypass call.
    mediaPathLocation *string
    // The OdataType property
    odataType *string
    // The datacenter used for signaling for both bypass and non-bypass calls.
    signalingLocation *string
    // Call start time.For failed and unanswered calls, this can be equal to invite or failure time.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Success or attempt.
    successfulCall *bool
    // Fully qualified domain name of the session border controller.
    trunkFullyQualifiedDomainName *string
    // Display name of the user.
    userDisplayName *string
    // Calling user's ID in Graph. This and other user info will be null/empty for bot call types. GUID.
    userId *string
    // UserPrincipalName (sign-in name) in Azure Active Directory. This is usually the same as user's SIP Address, and can be same as user's e-mail address.
    userPrincipalName *string
}
// NewDirectRoutingLogRow instantiates a new directRoutingLogRow and sets the default values.
func NewDirectRoutingLogRow()(*DirectRoutingLogRow) {
    m := &DirectRoutingLogRow{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDirectRoutingLogRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectRoutingLogRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectRoutingLogRow(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DirectRoutingLogRow) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCalleeNumber gets the calleeNumber property value. Number of the user or bot who received the call. E.164 format, but may include additional data.
func (m *DirectRoutingLogRow) GetCalleeNumber()(*string) {
    return m.calleeNumber
}
// GetCallEndSubReason gets the callEndSubReason property value. In addition to the SIP codes, Microsoft has own subcodes that indicate the specific issue.
func (m *DirectRoutingLogRow) GetCallEndSubReason()(*int32) {
    return m.callEndSubReason
}
// GetCallerNumber gets the callerNumber property value. Number of the user or bot who made the call. E.164 format, but may include additional data.
func (m *DirectRoutingLogRow) GetCallerNumber()(*string) {
    return m.callerNumber
}
// GetCallType gets the callType property value. Call type and direction.
func (m *DirectRoutingLogRow) GetCallType()(*string) {
    return m.callType
}
// GetCorrelationId gets the correlationId property value. Identifier for the call that you can use when calling Microsoft Support. GUID.
func (m *DirectRoutingLogRow) GetCorrelationId()(*string) {
    return m.correlationId
}
// GetDuration gets the duration property value. Duration of the call in seconds.
func (m *DirectRoutingLogRow) GetDuration()(*int32) {
    return m.duration
}
// GetEndDateTime gets the endDateTime property value. Only exists for successful (fully established) calls. Time when call ended.
func (m *DirectRoutingLogRow) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFailureDateTime gets the failureDateTime property value. Only exists for failed (not fully established) calls.
func (m *DirectRoutingLogRow) GetFailureDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.failureDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectRoutingLogRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calleeNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalleeNumber(val)
        }
        return nil
    }
    res["callEndSubReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallEndSubReason(val)
        }
        return nil
    }
    res["callerNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallerNumber(val)
        }
        return nil
    }
    res["callType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallType(val)
        }
        return nil
    }
    res["correlationId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCorrelationId(val)
        }
        return nil
    }
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
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
    res["failureDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailureDateTime(val)
        }
        return nil
    }
    res["finalSipCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinalSipCode(val)
        }
        return nil
    }
    res["finalSipCodePhrase"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFinalSipCodePhrase(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["inviteDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInviteDateTime(val)
        }
        return nil
    }
    res["mediaBypassEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaBypassEnabled(val)
        }
        return nil
    }
    res["mediaPathLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMediaPathLocation(val)
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
    res["signalingLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSignalingLocation(val)
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
    res["successfulCall"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuccessfulCall(val)
        }
        return nil
    }
    res["trunkFullyQualifiedDomainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrunkFullyQualifiedDomainName(val)
        }
        return nil
    }
    res["userDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDisplayName(val)
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetFinalSipCode gets the finalSipCode property value. The code with which the call ended, RFC 3261.
func (m *DirectRoutingLogRow) GetFinalSipCode()(*int32) {
    return m.finalSipCode
}
// GetFinalSipCodePhrase gets the finalSipCodePhrase property value. Description of the SIP code and Microsoft subcode.
func (m *DirectRoutingLogRow) GetFinalSipCodePhrase()(*string) {
    return m.finalSipCodePhrase
}
// GetId gets the id property value. Unique call identifier. GUID.
func (m *DirectRoutingLogRow) GetId()(*string) {
    return m.id
}
// GetInviteDateTime gets the inviteDateTime property value. When the initial invite was sent.
func (m *DirectRoutingLogRow) GetInviteDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.inviteDateTime
}
// GetMediaBypassEnabled gets the mediaBypassEnabled property value. Indicates if the trunk was enabled for media bypass or not.
func (m *DirectRoutingLogRow) GetMediaBypassEnabled()(*bool) {
    return m.mediaBypassEnabled
}
// GetMediaPathLocation gets the mediaPathLocation property value. The datacenter used for media path in non-bypass call.
func (m *DirectRoutingLogRow) GetMediaPathLocation()(*string) {
    return m.mediaPathLocation
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DirectRoutingLogRow) GetOdataType()(*string) {
    return m.odataType
}
// GetSignalingLocation gets the signalingLocation property value. The datacenter used for signaling for both bypass and non-bypass calls.
func (m *DirectRoutingLogRow) GetSignalingLocation()(*string) {
    return m.signalingLocation
}
// GetStartDateTime gets the startDateTime property value. Call start time.For failed and unanswered calls, this can be equal to invite or failure time.
func (m *DirectRoutingLogRow) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetSuccessfulCall gets the successfulCall property value. Success or attempt.
func (m *DirectRoutingLogRow) GetSuccessfulCall()(*bool) {
    return m.successfulCall
}
// GetTrunkFullyQualifiedDomainName gets the trunkFullyQualifiedDomainName property value. Fully qualified domain name of the session border controller.
func (m *DirectRoutingLogRow) GetTrunkFullyQualifiedDomainName()(*string) {
    return m.trunkFullyQualifiedDomainName
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user.
func (m *DirectRoutingLogRow) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. Calling user's ID in Graph. This and other user info will be null/empty for bot call types. GUID.
func (m *DirectRoutingLogRow) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. UserPrincipalName (sign-in name) in Azure Active Directory. This is usually the same as user's SIP Address, and can be same as user's e-mail address.
func (m *DirectRoutingLogRow) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *DirectRoutingLogRow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("calleeNumber", m.GetCalleeNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("callEndSubReason", m.GetCallEndSubReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("callerNumber", m.GetCallerNumber())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("callType", m.GetCallType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("correlationId", m.GetCorrelationId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("failureDateTime", m.GetFailureDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("finalSipCode", m.GetFinalSipCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("finalSipCodePhrase", m.GetFinalSipCodePhrase())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("inviteDateTime", m.GetInviteDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mediaBypassEnabled", m.GetMediaBypassEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mediaPathLocation", m.GetMediaPathLocation())
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
        err := writer.WriteStringValue("signalingLocation", m.GetSignalingLocation())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("successfulCall", m.GetSuccessfulCall())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("trunkFullyQualifiedDomainName", m.GetTrunkFullyQualifiedDomainName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userDisplayName", m.GetUserDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
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
func (m *DirectRoutingLogRow) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCalleeNumber sets the calleeNumber property value. Number of the user or bot who received the call. E.164 format, but may include additional data.
func (m *DirectRoutingLogRow) SetCalleeNumber(value *string)() {
    m.calleeNumber = value
}
// SetCallEndSubReason sets the callEndSubReason property value. In addition to the SIP codes, Microsoft has own subcodes that indicate the specific issue.
func (m *DirectRoutingLogRow) SetCallEndSubReason(value *int32)() {
    m.callEndSubReason = value
}
// SetCallerNumber sets the callerNumber property value. Number of the user or bot who made the call. E.164 format, but may include additional data.
func (m *DirectRoutingLogRow) SetCallerNumber(value *string)() {
    m.callerNumber = value
}
// SetCallType sets the callType property value. Call type and direction.
func (m *DirectRoutingLogRow) SetCallType(value *string)() {
    m.callType = value
}
// SetCorrelationId sets the correlationId property value. Identifier for the call that you can use when calling Microsoft Support. GUID.
func (m *DirectRoutingLogRow) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// SetDuration sets the duration property value. Duration of the call in seconds.
func (m *DirectRoutingLogRow) SetDuration(value *int32)() {
    m.duration = value
}
// SetEndDateTime sets the endDateTime property value. Only exists for successful (fully established) calls. Time when call ended.
func (m *DirectRoutingLogRow) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetFailureDateTime sets the failureDateTime property value. Only exists for failed (not fully established) calls.
func (m *DirectRoutingLogRow) SetFailureDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.failureDateTime = value
}
// SetFinalSipCode sets the finalSipCode property value. The code with which the call ended, RFC 3261.
func (m *DirectRoutingLogRow) SetFinalSipCode(value *int32)() {
    m.finalSipCode = value
}
// SetFinalSipCodePhrase sets the finalSipCodePhrase property value. Description of the SIP code and Microsoft subcode.
func (m *DirectRoutingLogRow) SetFinalSipCodePhrase(value *string)() {
    m.finalSipCodePhrase = value
}
// SetId sets the id property value. Unique call identifier. GUID.
func (m *DirectRoutingLogRow) SetId(value *string)() {
    m.id = value
}
// SetInviteDateTime sets the inviteDateTime property value. When the initial invite was sent.
func (m *DirectRoutingLogRow) SetInviteDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.inviteDateTime = value
}
// SetMediaBypassEnabled sets the mediaBypassEnabled property value. Indicates if the trunk was enabled for media bypass or not.
func (m *DirectRoutingLogRow) SetMediaBypassEnabled(value *bool)() {
    m.mediaBypassEnabled = value
}
// SetMediaPathLocation sets the mediaPathLocation property value. The datacenter used for media path in non-bypass call.
func (m *DirectRoutingLogRow) SetMediaPathLocation(value *string)() {
    m.mediaPathLocation = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DirectRoutingLogRow) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSignalingLocation sets the signalingLocation property value. The datacenter used for signaling for both bypass and non-bypass calls.
func (m *DirectRoutingLogRow) SetSignalingLocation(value *string)() {
    m.signalingLocation = value
}
// SetStartDateTime sets the startDateTime property value. Call start time.For failed and unanswered calls, this can be equal to invite or failure time.
func (m *DirectRoutingLogRow) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetSuccessfulCall sets the successfulCall property value. Success or attempt.
func (m *DirectRoutingLogRow) SetSuccessfulCall(value *bool)() {
    m.successfulCall = value
}
// SetTrunkFullyQualifiedDomainName sets the trunkFullyQualifiedDomainName property value. Fully qualified domain name of the session border controller.
func (m *DirectRoutingLogRow) SetTrunkFullyQualifiedDomainName(value *string)() {
    m.trunkFullyQualifiedDomainName = value
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user.
func (m *DirectRoutingLogRow) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. Calling user's ID in Graph. This and other user info will be null/empty for bot call types. GUID.
func (m *DirectRoutingLogRow) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. UserPrincipalName (sign-in name) in Azure Active Directory. This is usually the same as user's SIP Address, and can be same as user's e-mail address.
func (m *DirectRoutingLogRow) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
