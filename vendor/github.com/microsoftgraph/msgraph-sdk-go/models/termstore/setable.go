package termstore

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Setable 
type Setable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChildren()([]Termable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetLocalizedNames()([]LocalizedNameable)
    GetParentGroup()(Groupable)
    GetProperties()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValueable)
    GetRelations()([]Relationable)
    GetTerms()([]Termable)
    SetChildren(value []Termable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetLocalizedNames(value []LocalizedNameable)()
    SetParentGroup(value Groupable)()
    SetProperties(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValueable)()
    SetRelations(value []Relationable)()
    SetTerms(value []Termable)()
}
