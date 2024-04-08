package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewNotificationRecipientQueryScopeable 
type AccessReviewNotificationRecipientQueryScopeable interface {
    AccessReviewNotificationRecipientScopeable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetQuery()(*string)
    GetQueryRoot()(*string)
    GetQueryType()(*string)
    SetQuery(value *string)()
    SetQueryRoot(value *string)()
    SetQueryType(value *string)()
}
