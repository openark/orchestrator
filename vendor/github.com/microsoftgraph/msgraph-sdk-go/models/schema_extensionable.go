package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SchemaExtensionable 
type SchemaExtensionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDescription()(*string)
    GetOwner()(*string)
    GetProperties()([]ExtensionSchemaPropertyable)
    GetStatus()(*string)
    GetTargetTypes()([]string)
    SetDescription(value *string)()
    SetOwner(value *string)()
    SetProperties(value []ExtensionSchemaPropertyable)()
    SetStatus(value *string)()
    SetTargetTypes(value []string)()
}
