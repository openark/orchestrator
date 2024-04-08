package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AlterationResponseable 
type AlterationResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetOdataType()(*string)
    GetOriginalQueryString()(*string)
    GetQueryAlteration()(SearchAlterationable)
    GetQueryAlterationType()(*SearchAlterationType)
    SetOdataType(value *string)()
    SetOriginalQueryString(value *string)()
    SetQueryAlteration(value SearchAlterationable)()
    SetQueryAlterationType(value *SearchAlterationType)()
}
