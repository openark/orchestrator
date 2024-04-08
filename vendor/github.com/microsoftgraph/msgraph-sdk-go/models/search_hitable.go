package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchHitable 
type SearchHitable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetContentSource()(*string)
    GetHitId()(*string)
    GetOdataType()(*string)
    GetRank()(*int32)
    GetResource()(Entityable)
    GetResultTemplateId()(*string)
    GetSummary()(*string)
    SetContentSource(value *string)()
    SetHitId(value *string)()
    SetOdataType(value *string)()
    SetRank(value *int32)()
    SetResource(value Entityable)()
    SetResultTemplateId(value *string)()
    SetSummary(value *string)()
}
