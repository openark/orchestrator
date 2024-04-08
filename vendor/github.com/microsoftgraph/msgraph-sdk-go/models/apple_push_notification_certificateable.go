package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplePushNotificationCertificateable 
type ApplePushNotificationCertificateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppleIdentifier()(*string)
    GetCertificate()(*string)
    GetCertificateSerialNumber()(*string)
    GetCertificateUploadFailureReason()(*string)
    GetCertificateUploadStatus()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetTopicIdentifier()(*string)
    SetAppleIdentifier(value *string)()
    SetCertificate(value *string)()
    SetCertificateSerialNumber(value *string)()
    SetCertificateUploadFailureReason(value *string)()
    SetCertificateUploadStatus(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetTopicIdentifier(value *string)()
}
