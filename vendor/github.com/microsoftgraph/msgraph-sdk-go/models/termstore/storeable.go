package termstore

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Storeable 
type Storeable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDefaultLanguageTag()(*string)
    GetGroups()([]Groupable)
    GetLanguageTags()([]string)
    GetSets()([]Setable)
    SetDefaultLanguageTag(value *string)()
    SetGroups(value []Groupable)()
    SetLanguageTags(value []string)()
    SetSets(value []Setable)()
}
