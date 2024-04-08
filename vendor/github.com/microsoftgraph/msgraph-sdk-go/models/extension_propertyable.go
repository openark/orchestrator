package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ExtensionPropertyable 
type ExtensionPropertyable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppDisplayName()(*string)
    GetDataType()(*string)
    GetIsSyncedFromOnPremises()(*bool)
    GetName()(*string)
    GetTargetObjects()([]string)
    SetAppDisplayName(value *string)()
    SetDataType(value *string)()
    SetIsSyncedFromOnPremises(value *bool)()
    SetName(value *string)()
    SetTargetObjects(value []string)()
}
