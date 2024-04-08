package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InformationalUrlable 
type InformationalUrlable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetLogoUrl()(*string)
    GetMarketingUrl()(*string)
    GetOdataType()(*string)
    GetPrivacyStatementUrl()(*string)
    GetSupportUrl()(*string)
    GetTermsOfServiceUrl()(*string)
    SetLogoUrl(value *string)()
    SetMarketingUrl(value *string)()
    SetOdataType(value *string)()
    SetPrivacyStatementUrl(value *string)()
    SetSupportUrl(value *string)()
    SetTermsOfServiceUrl(value *string)()
}
