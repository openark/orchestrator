package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Workbookable 
type Workbookable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApplication()(WorkbookApplicationable)
    GetComments()([]WorkbookCommentable)
    GetFunctions()(WorkbookFunctionsable)
    GetNames()([]WorkbookNamedItemable)
    GetOperations()([]WorkbookOperationable)
    GetTables()([]WorkbookTableable)
    GetWorksheets()([]WorkbookWorksheetable)
    SetApplication(value WorkbookApplicationable)()
    SetComments(value []WorkbookCommentable)()
    SetFunctions(value WorkbookFunctionsable)()
    SetNames(value []WorkbookNamedItemable)()
    SetOperations(value []WorkbookOperationable)()
    SetTables(value []WorkbookTableable)()
    SetWorksheets(value []WorkbookWorksheetable)()
}
