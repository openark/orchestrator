package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookSortFieldable 
type WorkbookSortFieldable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAscending()(*bool)
    GetColor()(*string)
    GetDataOption()(*string)
    GetIcon()(WorkbookIconable)
    GetKey()(*int32)
    GetOdataType()(*string)
    GetSortOn()(*string)
    SetAscending(value *bool)()
    SetColor(value *string)()
    SetDataOption(value *string)()
    SetIcon(value WorkbookIconable)()
    SetKey(value *int32)()
    SetOdataType(value *string)()
    SetSortOn(value *string)()
}
