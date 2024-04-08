package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserFlowLanguageConfigurationable 
type UserFlowLanguageConfigurationable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultPages()([]UserFlowLanguagePageable)
    GetDisplayName()(*string)
    GetIsEnabled()(*bool)
    GetOverridesPages()([]UserFlowLanguagePageable)
    SetDefaultPages(value []UserFlowLanguagePageable)()
    SetDisplayName(value *string)()
    SetIsEnabled(value *bool)()
    SetOverridesPages(value []UserFlowLanguagePageable)()
}
