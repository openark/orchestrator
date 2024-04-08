package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChannelDescriptionUpdatedEventMessageDetailable 
type ChannelDescriptionUpdatedEventMessageDetailable interface {
    EventMessageDetailable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetChannelDescription()(*string)
    GetChannelId()(*string)
    GetInitiator()(IdentitySetable)
    SetChannelDescription(value *string)()
    SetChannelId(value *string)()
    SetInitiator(value IdentitySetable)()
}
