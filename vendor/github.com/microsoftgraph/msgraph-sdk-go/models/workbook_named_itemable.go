package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookNamedItemable 
type WorkbookNamedItemable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetComment()(*string)
    GetName()(*string)
    GetScope()(*string)
    GetType()(*string)
    GetValue()(Jsonable)
    GetVisible()(*bool)
    GetWorksheet()(WorkbookWorksheetable)
    SetComment(value *string)()
    SetName(value *string)()
    SetScope(value *string)()
    SetType(value *string)()
    SetValue(value Jsonable)()
    SetVisible(value *bool)()
    SetWorksheet(value WorkbookWorksheetable)()
}
