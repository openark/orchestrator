package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegistryValueEvidenceable 
type RegistryValueEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetRegistryHive()(*string)
    GetRegistryKey()(*string)
    GetRegistryValue()(*string)
    GetRegistryValueName()(*string)
    GetRegistryValueType()(*string)
    SetRegistryHive(value *string)()
    SetRegistryKey(value *string)()
    SetRegistryValue(value *string)()
    SetRegistryValueName(value *string)()
    SetRegistryValueType(value *string)()
}
