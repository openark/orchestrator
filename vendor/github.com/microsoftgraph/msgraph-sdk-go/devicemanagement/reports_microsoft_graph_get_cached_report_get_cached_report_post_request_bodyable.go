package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBodyable 
type ReportsMicrosoftGraphGetCachedReportGetCachedReportPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroupBy()([]string)
    GetId()(*string)
    GetOrderBy()([]string)
    GetSearch()(*string)
    GetSelect()([]string)
    GetSkip()(*int32)
    GetTop()(*int32)
    SetGroupBy(value []string)()
    SetId(value *string)()
    SetOrderBy(value []string)()
    SetSearch(value *string)()
    SetSelect(value []string)()
    SetSkip(value *int32)()
    SetTop(value *int32)()
}
