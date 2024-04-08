package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookFilterCriteriaable 
type WorkbookFilterCriteriaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetColor()(*string)
    GetCriterion1()(*string)
    GetCriterion2()(*string)
    GetDynamicCriteria()(*string)
    GetFilterOn()(*string)
    GetIcon()(WorkbookIconable)
    GetOdataType()(*string)
    GetOperator()(*string)
    GetValues()(Jsonable)
    SetColor(value *string)()
    SetCriterion1(value *string)()
    SetCriterion2(value *string)()
    SetDynamicCriteria(value *string)()
    SetFilterOn(value *string)()
    SetIcon(value WorkbookIconable)()
    SetOdataType(value *string)()
    SetOperator(value *string)()
    SetValues(value Jsonable)()
}
