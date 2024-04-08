package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeleconferenceDeviceVideoQualityable 
type TeleconferenceDeviceVideoQualityable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    TeleconferenceDeviceMediaQualityable
    GetAverageInboundBitRate()(*float64)
    GetAverageInboundFrameRate()(*float64)
    GetAverageOutboundBitRate()(*float64)
    GetAverageOutboundFrameRate()(*float64)
    SetAverageInboundBitRate(value *float64)()
    SetAverageInboundFrameRate(value *float64)()
    SetAverageOutboundBitRate(value *float64)()
    SetAverageOutboundFrameRate(value *float64)()
}
