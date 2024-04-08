package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Media 
type Media struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Device information associated with the callee endpoint of this media.
    calleeDevice DeviceInfoable
    // Network information associated with the callee endpoint of this media.
    calleeNetwork NetworkInfoable
    // Device information associated with the caller endpoint of this media.
    callerDevice DeviceInfoable
    // Network information associated with the caller endpoint of this media.
    callerNetwork NetworkInfoable
    // How the media was identified during media negotiation stage.
    label *string
    // The OdataType property
    odataType *string
    // Network streams associated with this media.
    streams []MediaStreamable
}
// NewMedia instantiates a new media and sets the default values.
func NewMedia()(*Media) {
    m := &Media{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMediaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMediaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMedia(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Media) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCalleeDevice gets the calleeDevice property value. Device information associated with the callee endpoint of this media.
func (m *Media) GetCalleeDevice()(DeviceInfoable) {
    return m.calleeDevice
}
// GetCalleeNetwork gets the calleeNetwork property value. Network information associated with the callee endpoint of this media.
func (m *Media) GetCalleeNetwork()(NetworkInfoable) {
    return m.calleeNetwork
}
// GetCallerDevice gets the callerDevice property value. Device information associated with the caller endpoint of this media.
func (m *Media) GetCallerDevice()(DeviceInfoable) {
    return m.callerDevice
}
// GetCallerNetwork gets the callerNetwork property value. Network information associated with the caller endpoint of this media.
func (m *Media) GetCallerNetwork()(NetworkInfoable) {
    return m.callerNetwork
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Media) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["calleeDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalleeDevice(val.(DeviceInfoable))
        }
        return nil
    }
    res["calleeNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNetworkInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalleeNetwork(val.(NetworkInfoable))
        }
        return nil
    }
    res["callerDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDeviceInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallerDevice(val.(DeviceInfoable))
        }
        return nil
    }
    res["callerNetwork"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNetworkInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCallerNetwork(val.(NetworkInfoable))
        }
        return nil
    }
    res["label"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLabel(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["streams"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMediaStreamFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MediaStreamable, len(val))
            for i, v := range val {
                res[i] = v.(MediaStreamable)
            }
            m.SetStreams(res)
        }
        return nil
    }
    return res
}
// GetLabel gets the label property value. How the media was identified during media negotiation stage.
func (m *Media) GetLabel()(*string) {
    return m.label
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Media) GetOdataType()(*string) {
    return m.odataType
}
// GetStreams gets the streams property value. Network streams associated with this media.
func (m *Media) GetStreams()([]MediaStreamable) {
    return m.streams
}
// Serialize serializes information the current object
func (m *Media) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("calleeDevice", m.GetCalleeDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("calleeNetwork", m.GetCalleeNetwork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("callerDevice", m.GetCallerDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("callerNetwork", m.GetCallerNetwork())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("label", m.GetLabel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    if m.GetStreams() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetStreams()))
        for i, v := range m.GetStreams() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("streams", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Media) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCalleeDevice sets the calleeDevice property value. Device information associated with the callee endpoint of this media.
func (m *Media) SetCalleeDevice(value DeviceInfoable)() {
    m.calleeDevice = value
}
// SetCalleeNetwork sets the calleeNetwork property value. Network information associated with the callee endpoint of this media.
func (m *Media) SetCalleeNetwork(value NetworkInfoable)() {
    m.calleeNetwork = value
}
// SetCallerDevice sets the callerDevice property value. Device information associated with the caller endpoint of this media.
func (m *Media) SetCallerDevice(value DeviceInfoable)() {
    m.callerDevice = value
}
// SetCallerNetwork sets the callerNetwork property value. Network information associated with the caller endpoint of this media.
func (m *Media) SetCallerNetwork(value NetworkInfoable)() {
    m.callerNetwork = value
}
// SetLabel sets the label property value. How the media was identified during media negotiation stage.
func (m *Media) SetLabel(value *string)() {
    m.label = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Media) SetOdataType(value *string)() {
    m.odataType = value
}
// SetStreams sets the streams property value. Network streams associated with this media.
func (m *Media) SetStreams(value []MediaStreamable)() {
    m.streams = value
}
