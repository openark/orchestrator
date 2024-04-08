package users

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemMicrosoftGraphAssignLicenseAssignLicensePostRequestBodyable 
type ItemMicrosoftGraphAssignLicenseAssignLicensePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddLicenses()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable)
    GetRemoveLicenses()([]i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetAddLicenses(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable)()
    SetRemoveLicenses(value []i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
