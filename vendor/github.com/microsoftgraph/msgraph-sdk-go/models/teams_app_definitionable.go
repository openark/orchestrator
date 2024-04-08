package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamsAppDefinitionable 
type TeamsAppDefinitionable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBot()(TeamworkBotable)
    GetCreatedBy()(IdentitySetable)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPublishingState()(*TeamsAppPublishingState)
    GetShortDescription()(*string)
    GetTeamsAppId()(*string)
    GetVersion()(*string)
    SetBot(value TeamworkBotable)()
    SetCreatedBy(value IdentitySetable)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPublishingState(value *TeamsAppPublishingState)()
    SetShortDescription(value *string)()
    SetTeamsAppId(value *string)()
    SetVersion(value *string)()
}
