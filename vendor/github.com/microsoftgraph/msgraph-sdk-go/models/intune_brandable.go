package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IntuneBrandable 
type IntuneBrandable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContactITEmailAddress()(*string)
    GetContactITName()(*string)
    GetContactITNotes()(*string)
    GetContactITPhoneNumber()(*string)
    GetDarkBackgroundLogo()(MimeContentable)
    GetDisplayName()(*string)
    GetLightBackgroundLogo()(MimeContentable)
    GetOdataType()(*string)
    GetOnlineSupportSiteName()(*string)
    GetOnlineSupportSiteUrl()(*string)
    GetPrivacyUrl()(*string)
    GetShowDisplayNameNextToLogo()(*bool)
    GetShowLogo()(*bool)
    GetShowNameNextToLogo()(*bool)
    GetThemeColor()(RgbColorable)
    SetContactITEmailAddress(value *string)()
    SetContactITName(value *string)()
    SetContactITNotes(value *string)()
    SetContactITPhoneNumber(value *string)()
    SetDarkBackgroundLogo(value MimeContentable)()
    SetDisplayName(value *string)()
    SetLightBackgroundLogo(value MimeContentable)()
    SetOdataType(value *string)()
    SetOnlineSupportSiteName(value *string)()
    SetOnlineSupportSiteUrl(value *string)()
    SetPrivacyUrl(value *string)()
    SetShowDisplayNameNextToLogo(value *bool)()
    SetShowLogo(value *bool)()
    SetShowNameNextToLogo(value *bool)()
    SetThemeColor(value RgbColorable)()
}
