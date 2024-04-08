package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Videoable 
type Videoable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAudioBitsPerSample()(*int32)
    GetAudioChannels()(*int32)
    GetAudioFormat()(*string)
    GetAudioSamplesPerSecond()(*int32)
    GetBitrate()(*int32)
    GetDuration()(*int64)
    GetFourCC()(*string)
    GetFrameRate()(*float64)
    GetHeight()(*int32)
    GetOdataType()(*string)
    GetWidth()(*int32)
    SetAudioBitsPerSample(value *int32)()
    SetAudioChannels(value *int32)()
    SetAudioFormat(value *string)()
    SetAudioSamplesPerSecond(value *int32)()
    SetBitrate(value *int32)()
    SetDuration(value *int64)()
    SetFourCC(value *string)()
    SetFrameRate(value *float64)()
    SetHeight(value *int32)()
    SetOdataType(value *string)()
    SetWidth(value *int32)()
}
