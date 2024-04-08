package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ExternalConnectionable 
type ExternalConnectionable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfiguration()(Configurationable)
    GetDescription()(*string)
    GetGroups()([]ExternalGroupable)
    GetItems()([]ExternalItemable)
    GetName()(*string)
    GetOperations()([]ConnectionOperationable)
    GetSchema()(Schemaable)
    GetState()(*ConnectionState)
    SetConfiguration(value Configurationable)()
    SetDescription(value *string)()
    SetGroups(value []ExternalGroupable)()
    SetItems(value []ExternalItemable)()
    SetName(value *string)()
    SetOperations(value []ConnectionOperationable)()
    SetSchema(value Schemaable)()
    SetState(value *ConnectionState)()
}
