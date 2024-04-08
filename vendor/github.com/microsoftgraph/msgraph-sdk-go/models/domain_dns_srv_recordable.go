package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainDnsSrvRecordable 
type DomainDnsSrvRecordable interface {
    DomainDnsRecordable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNameTarget()(*string)
    GetPort()(*int32)
    GetPriority()(*int32)
    GetProtocol()(*string)
    GetService()(*string)
    GetWeight()(*int32)
    SetNameTarget(value *string)()
    SetPort(value *int32)()
    SetPriority(value *int32)()
    SetProtocol(value *string)()
    SetService(value *string)()
    SetWeight(value *int32)()
}
