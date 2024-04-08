package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// UnifiedGroupSourceable 
type UnifiedGroupSourceable interface {
    DataSourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroup()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable)
    GetIncludedSources()(*SourceType)
    SetGroup(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable)()
    SetIncludedSources(value *SourceType)()
}
