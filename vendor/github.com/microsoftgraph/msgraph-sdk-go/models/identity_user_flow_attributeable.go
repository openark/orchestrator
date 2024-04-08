package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IdentityUserFlowAttributeable 
type IdentityUserFlowAttributeable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDataType()(*IdentityUserFlowAttributeDataType)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetUserFlowAttributeType()(*IdentityUserFlowAttributeType)
    SetDataType(value *IdentityUserFlowAttributeDataType)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetUserFlowAttributeType(value *IdentityUserFlowAttributeType)()
}
