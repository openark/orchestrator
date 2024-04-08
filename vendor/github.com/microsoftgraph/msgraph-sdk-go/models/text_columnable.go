package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TextColumnable 
type TextColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowMultipleLines()(*bool)
    GetAppendChangesToExistingText()(*bool)
    GetLinesForEditing()(*int32)
    GetMaxLength()(*int32)
    GetOdataType()(*string)
    GetTextType()(*string)
    SetAllowMultipleLines(value *bool)()
    SetAppendChangesToExistingText(value *bool)()
    SetLinesForEditing(value *int32)()
    SetMaxLength(value *int32)()
    SetOdataType(value *string)()
    SetTextType(value *string)()
}
