package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppProductCodeRuleable 
type Win32LobAppProductCodeRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRuleable
    GetProductCode()(*string)
    GetProductVersion()(*string)
    GetProductVersionOperator()(*Win32LobAppRuleOperator)
    SetProductCode(value *string)()
    SetProductVersion(value *string)()
    SetProductVersionOperator(value *Win32LobAppRuleOperator)()
}
