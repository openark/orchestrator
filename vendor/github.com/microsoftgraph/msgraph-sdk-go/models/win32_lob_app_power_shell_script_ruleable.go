package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppPowerShellScriptRuleable 
type Win32LobAppPowerShellScriptRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRuleable
    GetComparisonValue()(*string)
    GetDisplayName()(*string)
    GetEnforceSignatureCheck()(*bool)
    GetOperationType()(*Win32LobAppPowerShellScriptRuleOperationType)
    GetOperator()(*Win32LobAppRuleOperator)
    GetRunAs32Bit()(*bool)
    GetRunAsAccount()(*RunAsAccountType)
    GetScriptContent()(*string)
    SetComparisonValue(value *string)()
    SetDisplayName(value *string)()
    SetEnforceSignatureCheck(value *bool)()
    SetOperationType(value *Win32LobAppPowerShellScriptRuleOperationType)()
    SetOperator(value *Win32LobAppRuleOperator)()
    SetRunAs32Bit(value *bool)()
    SetRunAsAccount(value *RunAsAccountType)()
    SetScriptContent(value *string)()
}
