package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupSettingTemplateable 
type GroupSettingTemplateable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetValues()([]SettingTemplateValueable)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetValues(value []SettingTemplateValueable)()
}
