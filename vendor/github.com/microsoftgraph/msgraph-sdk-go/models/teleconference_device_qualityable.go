package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeleconferenceDeviceQualityable 
type TeleconferenceDeviceQualityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCallChainId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetCloudServiceDeploymentEnvironment()(*string)
    GetCloudServiceDeploymentId()(*string)
    GetCloudServiceInstanceName()(*string)
    GetCloudServiceName()(*string)
    GetDeviceDescription()(*string)
    GetDeviceName()(*string)
    GetMediaLegId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    GetMediaQualityList()([]TeleconferenceDeviceMediaQualityable)
    GetOdataType()(*string)
    GetParticipantId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)
    SetCallChainId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetCloudServiceDeploymentEnvironment(value *string)()
    SetCloudServiceDeploymentId(value *string)()
    SetCloudServiceInstanceName(value *string)()
    SetCloudServiceName(value *string)()
    SetDeviceDescription(value *string)()
    SetDeviceName(value *string)()
    SetMediaLegId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
    SetMediaQualityList(value []TeleconferenceDeviceMediaQualityable)()
    SetOdataType(value *string)()
    SetParticipantId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)()
}
