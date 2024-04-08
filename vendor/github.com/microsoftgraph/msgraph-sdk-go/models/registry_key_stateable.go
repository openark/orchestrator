package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegistryKeyStateable 
type RegistryKeyStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetHive()(*RegistryHive)
    GetKey()(*string)
    GetOdataType()(*string)
    GetOldKey()(*string)
    GetOldValueData()(*string)
    GetOldValueName()(*string)
    GetOperation()(*RegistryOperation)
    GetProcessId()(*int32)
    GetValueData()(*string)
    GetValueName()(*string)
    GetValueType()(*RegistryValueType)
    SetHive(value *RegistryHive)()
    SetKey(value *string)()
    SetOdataType(value *string)()
    SetOldKey(value *string)()
    SetOldValueData(value *string)()
    SetOldValueName(value *string)()
    SetOperation(value *RegistryOperation)()
    SetProcessId(value *int32)()
    SetValueData(value *string)()
    SetValueName(value *string)()
    SetValueType(value *RegistryValueType)()
}
