package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Subscriptionable 
type Subscriptionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplicationId()(*string)
    GetChangeType()(*string)
    GetClientState()(*string)
    GetCreatorId()(*string)
    GetEncryptionCertificate()(*string)
    GetEncryptionCertificateId()(*string)
    GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetIncludeResourceData()(*bool)
    GetLatestSupportedTlsVersion()(*string)
    GetLifecycleNotificationUrl()(*string)
    GetNotificationQueryOptions()(*string)
    GetNotificationUrl()(*string)
    GetNotificationUrlAppId()(*string)
    GetResource()(*string)
    SetApplicationId(value *string)()
    SetChangeType(value *string)()
    SetClientState(value *string)()
    SetCreatorId(value *string)()
    SetEncryptionCertificate(value *string)()
    SetEncryptionCertificateId(value *string)()
    SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetIncludeResourceData(value *bool)()
    SetLatestSupportedTlsVersion(value *string)()
    SetLifecycleNotificationUrl(value *string)()
    SetNotificationQueryOptions(value *string)()
    SetNotificationUrl(value *string)()
    SetNotificationUrlAppId(value *string)()
    SetResource(value *string)()
}
