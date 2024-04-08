package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeleconferenceDeviceVideoQuality 
type TeleconferenceDeviceVideoQuality struct {
    TeleconferenceDeviceMediaQuality
    // The average inbound stream video bit rate per second.
    averageInboundBitRate *float64
    // The average inbound stream video frame rate per second.
    averageInboundFrameRate *float64
    // The average outbound stream video bit rate per second.
    averageOutboundBitRate *float64
    // The average outbound stream video frame rate per second.
    averageOutboundFrameRate *float64
}
// NewTeleconferenceDeviceVideoQuality instantiates a new TeleconferenceDeviceVideoQuality and sets the default values.
func NewTeleconferenceDeviceVideoQuality()(*TeleconferenceDeviceVideoQuality) {
    m := &TeleconferenceDeviceVideoQuality{
        TeleconferenceDeviceMediaQuality: *NewTeleconferenceDeviceMediaQuality(),
    }
    odataTypeValue := "#microsoft.graph.teleconferenceDeviceVideoQuality"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeleconferenceDeviceVideoQualityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeleconferenceDeviceVideoQualityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.teleconferenceDeviceScreenSharingQuality":
                        return NewTeleconferenceDeviceScreenSharingQuality(), nil
                }
            }
        }
    }
    return NewTeleconferenceDeviceVideoQuality(), nil
}
// GetAverageInboundBitRate gets the averageInboundBitRate property value. The average inbound stream video bit rate per second.
func (m *TeleconferenceDeviceVideoQuality) GetAverageInboundBitRate()(*float64) {
    return m.averageInboundBitRate
}
// GetAverageInboundFrameRate gets the averageInboundFrameRate property value. The average inbound stream video frame rate per second.
func (m *TeleconferenceDeviceVideoQuality) GetAverageInboundFrameRate()(*float64) {
    return m.averageInboundFrameRate
}
// GetAverageOutboundBitRate gets the averageOutboundBitRate property value. The average outbound stream video bit rate per second.
func (m *TeleconferenceDeviceVideoQuality) GetAverageOutboundBitRate()(*float64) {
    return m.averageOutboundBitRate
}
// GetAverageOutboundFrameRate gets the averageOutboundFrameRate property value. The average outbound stream video frame rate per second.
func (m *TeleconferenceDeviceVideoQuality) GetAverageOutboundFrameRate()(*float64) {
    return m.averageOutboundFrameRate
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeleconferenceDeviceVideoQuality) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeleconferenceDeviceMediaQuality.GetFieldDeserializers()
    res["averageInboundBitRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundBitRate(val)
        }
        return nil
    }
    res["averageInboundFrameRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageInboundFrameRate(val)
        }
        return nil
    }
    res["averageOutboundBitRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundBitRate(val)
        }
        return nil
    }
    res["averageOutboundFrameRate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageOutboundFrameRate(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *TeleconferenceDeviceVideoQuality) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeleconferenceDeviceMediaQuality.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteFloat64Value("averageInboundBitRate", m.GetAverageInboundBitRate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("averageInboundFrameRate", m.GetAverageInboundFrameRate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("averageOutboundBitRate", m.GetAverageOutboundBitRate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("averageOutboundFrameRate", m.GetAverageOutboundFrameRate())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAverageInboundBitRate sets the averageInboundBitRate property value. The average inbound stream video bit rate per second.
func (m *TeleconferenceDeviceVideoQuality) SetAverageInboundBitRate(value *float64)() {
    m.averageInboundBitRate = value
}
// SetAverageInboundFrameRate sets the averageInboundFrameRate property value. The average inbound stream video frame rate per second.
func (m *TeleconferenceDeviceVideoQuality) SetAverageInboundFrameRate(value *float64)() {
    m.averageInboundFrameRate = value
}
// SetAverageOutboundBitRate sets the averageOutboundBitRate property value. The average outbound stream video bit rate per second.
func (m *TeleconferenceDeviceVideoQuality) SetAverageOutboundBitRate(value *float64)() {
    m.averageOutboundBitRate = value
}
// SetAverageOutboundFrameRate sets the averageOutboundFrameRate property value. The average outbound stream video frame rate per second.
func (m *TeleconferenceDeviceVideoQuality) SetAverageOutboundFrameRate(value *float64)() {
    m.averageOutboundFrameRate = value
}
