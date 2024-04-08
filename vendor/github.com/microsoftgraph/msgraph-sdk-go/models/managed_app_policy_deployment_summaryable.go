package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppPolicyDeploymentSummaryable 
type ManagedAppPolicyDeploymentSummaryable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetConfigurationDeployedUserCount()(*int32)
    GetConfigurationDeploymentSummaryPerApp()([]ManagedAppPolicyDeploymentSummaryPerAppable)
    GetDisplayName()(*string)
    GetLastRefreshTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersion()(*string)
    SetConfigurationDeployedUserCount(value *int32)()
    SetConfigurationDeploymentSummaryPerApp(value []ManagedAppPolicyDeploymentSummaryPerAppable)()
    SetDisplayName(value *string)()
    SetLastRefreshTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersion(value *string)()
}
