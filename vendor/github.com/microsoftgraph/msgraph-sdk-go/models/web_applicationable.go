package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WebApplicationable 
type WebApplicationable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHomePageUrl()(*string)
    GetImplicitGrantSettings()(ImplicitGrantSettingsable)
    GetLogoutUrl()(*string)
    GetOdataType()(*string)
    GetRedirectUris()([]string)
    GetRedirectUriSettings()([]RedirectUriSettingsable)
    SetHomePageUrl(value *string)()
    SetImplicitGrantSettings(value ImplicitGrantSettingsable)()
    SetLogoutUrl(value *string)()
    SetOdataType(value *string)()
    SetRedirectUris(value []string)()
    SetRedirectUriSettings(value []RedirectUriSettingsable)()
}
