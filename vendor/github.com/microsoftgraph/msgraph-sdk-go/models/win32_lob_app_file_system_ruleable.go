package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppFileSystemRuleable 
type Win32LobAppFileSystemRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    Win32LobAppRuleable
    GetCheck32BitOn64System()(*bool)
    GetComparisonValue()(*string)
    GetFileOrFolderName()(*string)
    GetOperationType()(*Win32LobAppFileSystemOperationType)
    GetOperator()(*Win32LobAppRuleOperator)
    GetPath()(*string)
    SetCheck32BitOn64System(value *bool)()
    SetComparisonValue(value *string)()
    SetFileOrFolderName(value *string)()
    SetOperationType(value *Win32LobAppFileSystemOperationType)()
    SetOperator(value *Win32LobAppRuleOperator)()
    SetPath(value *string)()
}
