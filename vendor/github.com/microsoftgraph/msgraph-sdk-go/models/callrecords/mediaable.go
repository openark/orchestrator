package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Mediaable 
type Mediaable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCalleeDevice()(DeviceInfoable)
    GetCalleeNetwork()(NetworkInfoable)
    GetCallerDevice()(DeviceInfoable)
    GetCallerNetwork()(NetworkInfoable)
    GetLabel()(*string)
    GetOdataType()(*string)
    GetStreams()([]MediaStreamable)
    SetCalleeDevice(value DeviceInfoable)()
    SetCalleeNetwork(value NetworkInfoable)()
    SetCallerDevice(value DeviceInfoable)()
    SetCallerNetwork(value NetworkInfoable)()
    SetLabel(value *string)()
    SetOdataType(value *string)()
    SetStreams(value []MediaStreamable)()
}
