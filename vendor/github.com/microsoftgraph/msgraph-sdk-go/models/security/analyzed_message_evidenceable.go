package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AnalyzedMessageEvidenceable 
type AnalyzedMessageEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAntiSpamDirection()(*string)
    GetAttachmentsCount()(*int64)
    GetDeliveryAction()(*string)
    GetDeliveryLocation()(*string)
    GetInternetMessageId()(*string)
    GetLanguage()(*string)
    GetNetworkMessageId()(*string)
    GetP1Sender()(EmailSenderable)
    GetP2Sender()(EmailSenderable)
    GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRecipientEmailAddress()(*string)
    GetSenderIp()(*string)
    GetSubject()(*string)
    GetThreatDetectionMethods()([]string)
    GetThreats()([]string)
    GetUrlCount()(*int64)
    GetUrls()([]string)
    GetUrn()(*string)
    SetAntiSpamDirection(value *string)()
    SetAttachmentsCount(value *int64)()
    SetDeliveryAction(value *string)()
    SetDeliveryLocation(value *string)()
    SetInternetMessageId(value *string)()
    SetLanguage(value *string)()
    SetNetworkMessageId(value *string)()
    SetP1Sender(value EmailSenderable)()
    SetP2Sender(value EmailSenderable)()
    SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRecipientEmailAddress(value *string)()
    SetSenderIp(value *string)()
    SetSubject(value *string)()
    SetThreatDetectionMethods(value []string)()
    SetThreats(value []string)()
    SetUrlCount(value *int64)()
    SetUrls(value []string)()
    SetUrn(value *string)()
}
