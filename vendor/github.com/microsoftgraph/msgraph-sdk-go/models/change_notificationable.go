package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChangeNotificationable 
type ChangeNotificationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChangeType()(*ChangeType)
    GetClientState()(*string)
    GetEncryptedContent()(ChangeNotificationEncryptedContentable)
    GetId()(*string)
    GetLifecycleEvent()(*LifecycleEventType)
    GetOdataType()(*string)
    GetResource()(*string)
    GetResourceData()(ResourceDataable)
    GetSubscriptionExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubscriptionId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetTenantId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetChangeType(value *ChangeType)()
    SetClientState(value *string)()
    SetEncryptedContent(value ChangeNotificationEncryptedContentable)()
    SetId(value *string)()
    SetLifecycleEvent(value *LifecycleEventType)()
    SetOdataType(value *string)()
    SetResource(value *string)()
    SetResourceData(value ResourceDataable)()
    SetSubscriptionExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubscriptionId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetTenantId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
