package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAccountable 
type UserAccountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountName()(*string)
    GetAzureAdUserId()(*string)
    GetDomainName()(*string)
    GetOdataType()(*string)
    GetUserPrincipalName()(*string)
    GetUserSid()(*string)
    SetAccountName(value *string)()
    SetAzureAdUserId(value *string)()
    SetDomainName(value *string)()
    SetOdataType(value *string)()
    SetUserPrincipalName(value *string)()
    SetUserSid(value *string)()
}
