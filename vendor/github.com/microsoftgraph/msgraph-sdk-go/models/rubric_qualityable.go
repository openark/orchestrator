package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RubricQualityable 
type RubricQualityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCriteria()([]RubricCriterionable)
    GetDescription()(EducationItemBodyable)
    GetDisplayName()(*string)
    GetOdataType()(*string)
    GetQualityId()(*string)
    GetWeight()(*float32)
    SetCriteria(value []RubricCriterionable)()
    SetDescription(value EducationItemBodyable)()
    SetDisplayName(value *string)()
    SetOdataType(value *string)()
    SetQualityId(value *string)()
    SetWeight(value *float32)()
}
