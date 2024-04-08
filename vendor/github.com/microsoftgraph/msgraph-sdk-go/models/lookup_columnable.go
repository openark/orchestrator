package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LookupColumnable 
type LookupColumnable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowMultipleValues()(*bool)
    GetAllowUnlimitedLength()(*bool)
    GetColumnName()(*string)
    GetListId()(*string)
    GetOdataType()(*string)
    GetPrimaryLookupColumnId()(*string)
    SetAllowMultipleValues(value *bool)()
    SetAllowUnlimitedLength(value *bool)()
    SetColumnName(value *string)()
    SetListId(value *string)()
    SetOdataType(value *string)()
    SetPrimaryLookupColumnId(value *string)()
}
