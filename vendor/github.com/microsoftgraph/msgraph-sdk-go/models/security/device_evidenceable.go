package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEvidenceable 
type DeviceEvidenceable interface {
    AlertEvidenceable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAzureAdDeviceId()(*string)
    GetDefenderAvStatus()(*DefenderAvStatus)
    GetDeviceDnsName()(*string)
    GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetHealthStatus()(*DeviceHealthStatus)
    GetLoggedOnUsers()([]LoggedOnUserable)
    GetMdeDeviceId()(*string)
    GetOnboardingStatus()(*OnboardingStatus)
    GetOsBuild()(*int64)
    GetOsPlatform()(*string)
    GetRbacGroupId()(*int32)
    GetRbacGroupName()(*string)
    GetRiskScore()(*DeviceRiskScore)
    GetVersion()(*string)
    GetVmMetadata()(VmMetadataable)
    SetAzureAdDeviceId(value *string)()
    SetDefenderAvStatus(value *DefenderAvStatus)()
    SetDeviceDnsName(value *string)()
    SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetHealthStatus(value *DeviceHealthStatus)()
    SetLoggedOnUsers(value []LoggedOnUserable)()
    SetMdeDeviceId(value *string)()
    SetOnboardingStatus(value *OnboardingStatus)()
    SetOsBuild(value *int64)()
    SetOsPlatform(value *string)()
    SetRbacGroupId(value *int32)()
    SetRbacGroupName(value *string)()
    SetRiskScore(value *DeviceRiskScore)()
    SetVersion(value *string)()
    SetVmMetadata(value VmMetadataable)()
}
