package callrecords

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PstnCallLogRow 
type PstnCallLogRow struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The source of the call duration data. If the call uses a third-party telecommunications operator via the Operator Connect Program, the operator may provide their own call duration data. In this case, the property value is operator. Otherwise, the value is microsoft.
    callDurationSource *PstnCallDurationSource
    // Number dialed in E.164 format.
    calleeNumber *string
    // Number that received the call for inbound calls or the number dialed for outbound calls. E.164 format.
    callerNumber *string
    // Call identifier. Not guaranteed to be unique.
    callId *string
    // Whether the call was a PSTN outbound or inbound call and the type of call such as a call placed by a user or an audio conference.
    callType *string
    // Amount of money or cost of the call that is charged to your account.
    charge *float64
    // ID of the audio conference.
    conferenceId *string
    // Connection fee price.
    connectionCharge *float64
    // Type of currency used to calculate the cost of the call. For details, see (ISO 4217.
    currency *string
    // Whether the call was domestic (within a country or region) or international (outside a country or region) based on the user's location.
    destinationContext *string
    // Country or region dialed.
    destinationName *string
    // How long the call was connected, in seconds.
    duration *int32
    // Call end time.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Unique call identifier. GUID.
    id *string
    // User's phone number type, such as a service of toll-free number.
    inventoryType *string
    // The license used for the call.
    licenseCapability *string
    // The OdataType property
    odataType *string
    // The telecommunications operator which provided PSTN services for this call. This may be Microsoft, or it may be a third-party operator via the Operator Connect Program.
    operator *string
    // Call start time.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Country code of the tenant. For details, see ISO 3166-1 alpha-2.
    tenantCountryCode *string
    // Country code of the user. For details, see ISO 3166-1 alpha-2.
    usageCountryCode *string
    // Display name of the user.
    userDisplayName *string
    // Calling user's ID in Graph. GUID. This and other user info will be null/empty for bot call types (ucap_in, ucap_out).
    userId *string
    // The user principal name (sign-in name) in Azure Active Directory. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
    userPrincipalName *string
}
// NewPstnCallLogRow instantiates a new pstnCallLogRow and sets the default values.
func NewPstnCallLogRow()(*PstnCallLogRow) {
    m := &PstnCallLogRow{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePstnCallLogRowFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePstnCallLogRowFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPstnCallLogRow(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PstnCallLogRow) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCallDurationSource gets the callDurationSource property value. The source of the call duration data. If the call uses a third-party telecommunications operator via the Operator Connect Program, the operator may provide their own call duration data. In this case, the property value is operator. Otherwise, the value is microsoft.
func (m *PstnCallLogRow) GetCallDurationSource()(*PstnCallDurationSource) {
    return m.callDurationSource
}
// GetCalleeNumber gets the calleeNumber property value. Number dialed in E.164 format.
func (m *PstnCallLogRow) GetCalleeNumber()(*string) {
    return m.calleeNumber
}
// GetCallerNumber gets the callerNumber property value. Number that received the call for inbound calls or the number dialed for outbound calls. E.164 format.
func (m *PstnCallLogRow) GetCallerNumber()(*string) {
    return m.callerNumber
}
// GetCallId gets the callId property value. Call identifier. Not guaranteed to be unique.
func (m *PstnCallLogRow) GetCallId()(*string) {
    return m.callId
}
// GetCallType gets the callType property value. Whether the call was a PSTN outbound or inbound call and the type of call such as a call placed by a user or an audio conference.
func (m *PstnCallLogRow) GetCallType()(*string) {
    return m.callType
}
// GetCharge gets the charge property value. Amount of money or cost of the call that is charged to your account.
func (m *PstnCallLogRow) GetCharge()(*float64) {
    return m.charge
}
// GetConferenceId gets the conferenceId property value. ID of the audio conference.
func (m *PstnCallLogRow) GetConferenceId()(*string) {
    return m.conferenceId
}
// GetConnectionCharge gets the connectionCharge property value. Connection fee price.
func (m *PstnCallLogRow) GetConnectionCharge()(*float64) {
    return m.connectionCharge
}
// GetCurrency gets the currency property value. Type of currency used to calculate the cost of the call. For details, see (ISO 4217.
func (m *PstnCallLogRow) GetCurrency()(*string) {
    return m.currency
}
// GetDestinationContext gets the destinationContext property value. Whether the call was domestic (within a country or region) or international (outside a country or region) based on the user's location.
func (m *PstnCallLogRow) GetDestinationContext()(*string) {
    return m.destinationContext
}
// GetDestinationName gets the destinationName property value. Country or region dialed.
func (m *PstnCallLogRow) GetDestinationName()(*string) {
    return m.destinationName
}
// GetDuration gets the duration property value. How long the call was connected, in seconds.
func (m *PstnCallLogRow) GetDuration()(*int32) {
    return m.duration
}
// GetEndDateTime gets the endDateTime property value. Call end time.
func (m *PstnCallLogRow) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PstnCallLogRow) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["callDurationSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePstnCallDurationSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallDurationSource(val.(*PstnCallDurationSource))
        }
        return nil
    }
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
    res["callId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallId(val)
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
    res["charge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCharge(val)
        }
        return nil
    }
    res["conferenceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConferenceId(val)
        }
        return nil
    }
    res["connectionCharge"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectionCharge(val)
        }
        return nil
    }
    res["currency"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrency(val)
        }
        return nil
    }
    res["destinationContext"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationContext(val)
        }
        return nil
    }
    res["destinationName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationName(val)
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
    res["inventoryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInventoryType(val)
        }
        return nil
    }
    res["licenseCapability"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLicenseCapability(val)
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
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val)
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
    res["tenantCountryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantCountryCode(val)
        }
        return nil
    }
    res["usageCountryCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsageCountryCode(val)
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
// GetId gets the id property value. Unique call identifier. GUID.
func (m *PstnCallLogRow) GetId()(*string) {
    return m.id
}
// GetInventoryType gets the inventoryType property value. User's phone number type, such as a service of toll-free number.
func (m *PstnCallLogRow) GetInventoryType()(*string) {
    return m.inventoryType
}
// GetLicenseCapability gets the licenseCapability property value. The license used for the call.
func (m *PstnCallLogRow) GetLicenseCapability()(*string) {
    return m.licenseCapability
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PstnCallLogRow) GetOdataType()(*string) {
    return m.odataType
}
// GetOperator gets the operator property value. The telecommunications operator which provided PSTN services for this call. This may be Microsoft, or it may be a third-party operator via the Operator Connect Program.
func (m *PstnCallLogRow) GetOperator()(*string) {
    return m.operator
}
// GetStartDateTime gets the startDateTime property value. Call start time.
func (m *PstnCallLogRow) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetTenantCountryCode gets the tenantCountryCode property value. Country code of the tenant. For details, see ISO 3166-1 alpha-2.
func (m *PstnCallLogRow) GetTenantCountryCode()(*string) {
    return m.tenantCountryCode
}
// GetUsageCountryCode gets the usageCountryCode property value. Country code of the user. For details, see ISO 3166-1 alpha-2.
func (m *PstnCallLogRow) GetUsageCountryCode()(*string) {
    return m.usageCountryCode
}
// GetUserDisplayName gets the userDisplayName property value. Display name of the user.
func (m *PstnCallLogRow) GetUserDisplayName()(*string) {
    return m.userDisplayName
}
// GetUserId gets the userId property value. Calling user's ID in Graph. GUID. This and other user info will be null/empty for bot call types (ucap_in, ucap_out).
func (m *PstnCallLogRow) GetUserId()(*string) {
    return m.userId
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (sign-in name) in Azure Active Directory. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
func (m *PstnCallLogRow) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *PstnCallLogRow) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCallDurationSource() != nil {
        cast := (*m.GetCallDurationSource()).String()
        err := writer.WriteStringValue("callDurationSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("calleeNumber", m.GetCalleeNumber())
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
        err := writer.WriteStringValue("callId", m.GetCallId())
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
        err := writer.WriteFloat64Value("charge", m.GetCharge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("conferenceId", m.GetConferenceId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("connectionCharge", m.GetConnectionCharge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("currency", m.GetCurrency())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationContext", m.GetDestinationContext())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("destinationName", m.GetDestinationName())
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
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("inventoryType", m.GetInventoryType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("licenseCapability", m.GetLicenseCapability())
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
        err := writer.WriteStringValue("operator", m.GetOperator())
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
        err := writer.WriteStringValue("tenantCountryCode", m.GetTenantCountryCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("usageCountryCode", m.GetUsageCountryCode())
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
func (m *PstnCallLogRow) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCallDurationSource sets the callDurationSource property value. The source of the call duration data. If the call uses a third-party telecommunications operator via the Operator Connect Program, the operator may provide their own call duration data. In this case, the property value is operator. Otherwise, the value is microsoft.
func (m *PstnCallLogRow) SetCallDurationSource(value *PstnCallDurationSource)() {
    m.callDurationSource = value
}
// SetCalleeNumber sets the calleeNumber property value. Number dialed in E.164 format.
func (m *PstnCallLogRow) SetCalleeNumber(value *string)() {
    m.calleeNumber = value
}
// SetCallerNumber sets the callerNumber property value. Number that received the call for inbound calls or the number dialed for outbound calls. E.164 format.
func (m *PstnCallLogRow) SetCallerNumber(value *string)() {
    m.callerNumber = value
}
// SetCallId sets the callId property value. Call identifier. Not guaranteed to be unique.
func (m *PstnCallLogRow) SetCallId(value *string)() {
    m.callId = value
}
// SetCallType sets the callType property value. Whether the call was a PSTN outbound or inbound call and the type of call such as a call placed by a user or an audio conference.
func (m *PstnCallLogRow) SetCallType(value *string)() {
    m.callType = value
}
// SetCharge sets the charge property value. Amount of money or cost of the call that is charged to your account.
func (m *PstnCallLogRow) SetCharge(value *float64)() {
    m.charge = value
}
// SetConferenceId sets the conferenceId property value. ID of the audio conference.
func (m *PstnCallLogRow) SetConferenceId(value *string)() {
    m.conferenceId = value
}
// SetConnectionCharge sets the connectionCharge property value. Connection fee price.
func (m *PstnCallLogRow) SetConnectionCharge(value *float64)() {
    m.connectionCharge = value
}
// SetCurrency sets the currency property value. Type of currency used to calculate the cost of the call. For details, see (ISO 4217.
func (m *PstnCallLogRow) SetCurrency(value *string)() {
    m.currency = value
}
// SetDestinationContext sets the destinationContext property value. Whether the call was domestic (within a country or region) or international (outside a country or region) based on the user's location.
func (m *PstnCallLogRow) SetDestinationContext(value *string)() {
    m.destinationContext = value
}
// SetDestinationName sets the destinationName property value. Country or region dialed.
func (m *PstnCallLogRow) SetDestinationName(value *string)() {
    m.destinationName = value
}
// SetDuration sets the duration property value. How long the call was connected, in seconds.
func (m *PstnCallLogRow) SetDuration(value *int32)() {
    m.duration = value
}
// SetEndDateTime sets the endDateTime property value. Call end time.
func (m *PstnCallLogRow) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetId sets the id property value. Unique call identifier. GUID.
func (m *PstnCallLogRow) SetId(value *string)() {
    m.id = value
}
// SetInventoryType sets the inventoryType property value. User's phone number type, such as a service of toll-free number.
func (m *PstnCallLogRow) SetInventoryType(value *string)() {
    m.inventoryType = value
}
// SetLicenseCapability sets the licenseCapability property value. The license used for the call.
func (m *PstnCallLogRow) SetLicenseCapability(value *string)() {
    m.licenseCapability = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PstnCallLogRow) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperator sets the operator property value. The telecommunications operator which provided PSTN services for this call. This may be Microsoft, or it may be a third-party operator via the Operator Connect Program.
func (m *PstnCallLogRow) SetOperator(value *string)() {
    m.operator = value
}
// SetStartDateTime sets the startDateTime property value. Call start time.
func (m *PstnCallLogRow) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetTenantCountryCode sets the tenantCountryCode property value. Country code of the tenant. For details, see ISO 3166-1 alpha-2.
func (m *PstnCallLogRow) SetTenantCountryCode(value *string)() {
    m.tenantCountryCode = value
}
// SetUsageCountryCode sets the usageCountryCode property value. Country code of the user. For details, see ISO 3166-1 alpha-2.
func (m *PstnCallLogRow) SetUsageCountryCode(value *string)() {
    m.usageCountryCode = value
}
// SetUserDisplayName sets the userDisplayName property value. Display name of the user.
func (m *PstnCallLogRow) SetUserDisplayName(value *string)() {
    m.userDisplayName = value
}
// SetUserId sets the userId property value. Calling user's ID in Graph. GUID. This and other user info will be null/empty for bot call types (ucap_in, ucap_out).
func (m *PstnCallLogRow) SetUserId(value *string)() {
    m.userId = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (sign-in name) in Azure Active Directory. This is usually the same as the user's SIP address, and can be same as the user's e-mail address.
func (m *PstnCallLogRow) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
