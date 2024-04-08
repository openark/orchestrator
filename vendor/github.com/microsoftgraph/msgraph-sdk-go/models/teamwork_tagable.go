package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamworkTagable 
type TeamworkTagable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetMemberCount()(*int32)
    GetMembers()([]TeamworkTagMemberable)
    GetTagType()(*TeamworkTagType)
    GetTeamId()(*string)
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetMemberCount(value *int32)()
    SetMembers(value []TeamworkTagMemberable)()
    SetTagType(value *TeamworkTagType)()
    SetTeamId(value *string)()
}
