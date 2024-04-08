package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSourceable 
type UserSourceable interface {
    DataSourceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetEmail()(*string)
    GetIncludedSources()(*SourceType)
    GetSiteWebUrl()(*string)
    SetEmail(value *string)()
    SetIncludedSources(value *SourceType)()
    SetSiteWebUrl(value *string)()
}
