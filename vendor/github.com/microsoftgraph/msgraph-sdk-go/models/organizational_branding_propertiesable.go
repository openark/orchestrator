package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrganizationalBrandingPropertiesable 
type OrganizationalBrandingPropertiesable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackgroundColor()(*string)
    GetBackgroundImage()([]byte)
    GetBackgroundImageRelativeUrl()(*string)
    GetBannerLogo()([]byte)
    GetBannerLogoRelativeUrl()(*string)
    GetCdnList()([]string)
    GetSignInPageText()(*string)
    GetSquareLogo()([]byte)
    GetSquareLogoRelativeUrl()(*string)
    GetUsernameHintText()(*string)
    SetBackgroundColor(value *string)()
    SetBackgroundImage(value []byte)()
    SetBackgroundImageRelativeUrl(value *string)()
    SetBannerLogo(value []byte)()
    SetBannerLogoRelativeUrl(value *string)()
    SetCdnList(value []string)()
    SetSignInPageText(value *string)()
    SetSquareLogo(value []byte)()
    SetSquareLogoRelativeUrl(value *string)()
    SetUsernameHintText(value *string)()
}
