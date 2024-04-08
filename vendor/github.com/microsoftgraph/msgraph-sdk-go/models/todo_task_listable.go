package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TodoTaskListable 
type TodoTaskListable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDisplayName()(*string)
    GetExtensions()([]Extensionable)
    GetIsOwner()(*bool)
    GetIsShared()(*bool)
    GetTasks()([]TodoTaskable)
    GetWellknownListName()(*WellknownListName)
    SetDisplayName(value *string)()
    SetExtensions(value []Extensionable)()
    SetIsOwner(value *bool)()
    SetIsShared(value *bool)()
    SetTasks(value []TodoTaskable)()
    SetWellknownListName(value *WellknownListName)()
}
