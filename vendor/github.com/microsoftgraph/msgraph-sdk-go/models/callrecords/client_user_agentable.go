package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientUserAgentable 
type ClientUserAgentable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UserAgentable
    GetAzureADAppId()(*string)
    GetCommunicationServiceId()(*string)
    GetPlatform()(*ClientPlatform)
    GetProductFamily()(*ProductFamily)
    SetAzureADAppId(value *string)()
    SetCommunicationServiceId(value *string)()
    SetPlatform(value *ClientPlatform)()
    SetProductFamily(value *ProductFamily)()
}
