package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsTabConfigurationable 
type TeamsTabConfigurationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentUrl()(*string)
    GetEntityId()(*string)
    GetOdataType()(*string)
    GetRemoveUrl()(*string)
    GetWebsiteUrl()(*string)
    SetContentUrl(value *string)()
    SetEntityId(value *string)()
    SetOdataType(value *string)()
    SetRemoveUrl(value *string)()
    SetWebsiteUrl(value *string)()
}
