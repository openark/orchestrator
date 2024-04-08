package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DomainDnsRecordable 
type DomainDnsRecordable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsOptional()(*bool)
    GetLabel()(*string)
    GetRecordType()(*string)
    GetSupportedService()(*string)
    GetTtl()(*int32)
    SetIsOptional(value *bool)()
    SetLabel(value *string)()
    SetRecordType(value *string)()
    SetSupportedService(value *string)()
    SetTtl(value *int32)()
}
