package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserExperienceAnalyticsDevicePerformanceable 
type UserExperienceAnalyticsDevicePerformanceable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAverageBlueScreens()(*float64)
    GetAverageRestarts()(*float64)
    GetBlueScreenCount()(*int32)
    GetBootScore()(*int32)
    GetCoreBootTimeInMs()(*int32)
    GetCoreLoginTimeInMs()(*int32)
    GetDeviceCount()(*int64)
    GetDeviceName()(*string)
    GetDiskType()(*DiskType)
    GetGroupPolicyBootTimeInMs()(*int32)
    GetGroupPolicyLoginTimeInMs()(*int32)
    GetHealthStatus()(*UserExperienceAnalyticsHealthState)
    GetLoginScore()(*int32)
    GetManufacturer()(*string)
    GetModel()(*string)
    GetModelStartupPerformanceScore()(*float64)
    GetOperatingSystemVersion()(*string)
    GetResponsiveDesktopTimeInMs()(*int32)
    GetRestartCount()(*int32)
    GetStartupPerformanceScore()(*float64)
    SetAverageBlueScreens(value *float64)()
    SetAverageRestarts(value *float64)()
    SetBlueScreenCount(value *int32)()
    SetBootScore(value *int32)()
    SetCoreBootTimeInMs(value *int32)()
    SetCoreLoginTimeInMs(value *int32)()
    SetDeviceCount(value *int64)()
    SetDeviceName(value *string)()
    SetDiskType(value *DiskType)()
    SetGroupPolicyBootTimeInMs(value *int32)()
    SetGroupPolicyLoginTimeInMs(value *int32)()
    SetHealthStatus(value *UserExperienceAnalyticsHealthState)()
    SetLoginScore(value *int32)()
    SetManufacturer(value *string)()
    SetModel(value *string)()
    SetModelStartupPerformanceScore(value *float64)()
    SetOperatingSystemVersion(value *string)()
    SetResponsiveDesktopTimeInMs(value *int32)()
    SetRestartCount(value *int32)()
    SetStartupPerformanceScore(value *float64)()
}
