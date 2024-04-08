package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// HostSecurityStateable 
type HostSecurityStateable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFqdn()(*string)
    GetIsAzureAdJoined()(*bool)
    GetIsAzureAdRegistered()(*bool)
    GetIsHybridAzureDomainJoined()(*bool)
    GetNetBiosName()(*string)
    GetOdataType()(*string)
    GetOs()(*string)
    GetPrivateIpAddress()(*string)
    GetPublicIpAddress()(*string)
    GetRiskScore()(*string)
    SetFqdn(value *string)()
    SetIsAzureAdJoined(value *bool)()
    SetIsAzureAdRegistered(value *bool)()
    SetIsHybridAzureDomainJoined(value *bool)()
    SetNetBiosName(value *string)()
    SetOdataType(value *string)()
    SetOs(value *string)()
    SetPrivateIpAddress(value *string)()
    SetPublicIpAddress(value *string)()
    SetRiskScore(value *string)()
}
