package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessTemplateable 
type ConditionalAccessTemplateable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDetails()(ConditionalAccessPolicyDetailable)
    GetName()(*string)
    GetScenarios()(*TemplateScenarios)
    SetDescription(value *string)()
    SetDetails(value ConditionalAccessPolicyDetailable)()
    SetName(value *string)()
    SetScenarios(value *TemplateScenarios)()
}
