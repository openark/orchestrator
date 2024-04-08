package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Endpointable 
type Endpointable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCapability()(*string)
    GetProviderId()(*string)
    GetProviderName()(*string)
    GetProviderResourceId()(*string)
    GetUri()(*string)
    SetCapability(value *string)()
    SetProviderId(value *string)()
    SetProviderName(value *string)()
    SetProviderResourceId(value *string)()
    SetUri(value *string)()
}
