package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppRegistryRuleable 
type Win32LobAppRegistryRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRuleable
    GetCheck32BitOn64System()(*bool)
    GetComparisonValue()(*string)
    GetKeyPath()(*string)
    GetOperationType()(*Win32LobAppRegistryRuleOperationType)
    GetOperator()(*Win32LobAppRuleOperator)
    GetValueName()(*string)
    SetCheck32BitOn64System(value *bool)()
    SetComparisonValue(value *string)()
    SetKeyPath(value *string)()
    SetOperationType(value *Win32LobAppRegistryRuleOperationType)()
    SetOperator(value *Win32LobAppRuleOperator)()
    SetValueName(value *string)()
}
