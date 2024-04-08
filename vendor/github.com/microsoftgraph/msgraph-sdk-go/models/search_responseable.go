package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchResponseable 
type SearchResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHitsContainers()([]SearchHitsContainerable)
    GetOdataType()(*string)
    GetQueryAlterationResponse()(AlterationResponseable)
    GetResultTemplates()(ResultTemplateDictionaryable)
    GetSearchTerms()([]string)
    SetHitsContainers(value []SearchHitsContainerable)()
    SetOdataType(value *string)()
    SetQueryAlterationResponse(value AlterationResponseable)()
    SetResultTemplates(value ResultTemplateDictionaryable)()
    SetSearchTerms(value []string)()
}
