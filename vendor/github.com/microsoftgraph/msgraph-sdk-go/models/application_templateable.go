package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ApplicationTemplateable 
type ApplicationTemplateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCategories()([]string)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetHomePageUrl()(*string)
    GetLogoUrl()(*string)
    GetPublisher()(*string)
    GetSupportedProvisioningTypes()([]string)
    GetSupportedSingleSignOnModes()([]string)
    SetCategories(value []string)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetHomePageUrl(value *string)()
    SetLogoUrl(value *string)()
    SetPublisher(value *string)()
    SetSupportedProvisioningTypes(value []string)()
    SetSupportedSingleSignOnModes(value []string)()
}
