package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceGeoLocation device location
type DeviceGeoLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Altitude, given in meters above sea level
    altitude *float64
    // Heading in degrees from true north
    heading *float64
    // Accuracy of longitude and latitude in meters
    horizontalAccuracy *float64
    // Time at which location was recorded, relative to UTC
    lastCollectedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Latitude coordinate of the device's location
    latitude *float64
    // Longitude coordinate of the device's location
    longitude *float64
    // The OdataType property
    odataType *string
    // Speed the device is traveling in meters per second
    speed *float64
    // Accuracy of altitude in meters
    verticalAccuracy *float64
}
// NewDeviceGeoLocation instantiates a new deviceGeoLocation and sets the default values.
func NewDeviceGeoLocation()(*DeviceGeoLocation) {
    m := &DeviceGeoLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDeviceGeoLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceGeoLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceGeoLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceGeoLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAltitude gets the altitude property value. Altitude, given in meters above sea level
func (m *DeviceGeoLocation) GetAltitude()(*float64) {
    return m.altitude
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceGeoLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["altitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltitude(val)
        }
        return nil
    }
    res["heading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHeading(val)
        }
        return nil
    }
    res["horizontalAccuracy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHorizontalAccuracy(val)
        }
        return nil
    }
    res["lastCollectedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastCollectedDateTime(val)
        }
        return nil
    }
    res["latitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLatitude(val)
        }
        return nil
    }
    res["longitude"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLongitude(val)
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
    res["speed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSpeed(val)
        }
        return nil
    }
    res["verticalAccuracy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerticalAccuracy(val)
        }
        return nil
    }
    return res
}
// GetHeading gets the heading property value. Heading in degrees from true north
func (m *DeviceGeoLocation) GetHeading()(*float64) {
    return m.heading
}
// GetHorizontalAccuracy gets the horizontalAccuracy property value. Accuracy of longitude and latitude in meters
func (m *DeviceGeoLocation) GetHorizontalAccuracy()(*float64) {
    return m.horizontalAccuracy
}
// GetLastCollectedDateTime gets the lastCollectedDateTime property value. Time at which location was recorded, relative to UTC
func (m *DeviceGeoLocation) GetLastCollectedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastCollectedDateTime
}
// GetLatitude gets the latitude property value. Latitude coordinate of the device's location
func (m *DeviceGeoLocation) GetLatitude()(*float64) {
    return m.latitude
}
// GetLongitude gets the longitude property value. Longitude coordinate of the device's location
func (m *DeviceGeoLocation) GetLongitude()(*float64) {
    return m.longitude
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DeviceGeoLocation) GetOdataType()(*string) {
    return m.odataType
}
// GetSpeed gets the speed property value. Speed the device is traveling in meters per second
func (m *DeviceGeoLocation) GetSpeed()(*float64) {
    return m.speed
}
// GetVerticalAccuracy gets the verticalAccuracy property value. Accuracy of altitude in meters
func (m *DeviceGeoLocation) GetVerticalAccuracy()(*float64) {
    return m.verticalAccuracy
}
// Serialize serializes information the current object
func (m *DeviceGeoLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("altitude", m.GetAltitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("heading", m.GetHeading())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("horizontalAccuracy", m.GetHorizontalAccuracy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("lastCollectedDateTime", m.GetLastCollectedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("latitude", m.GetLatitude())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("longitude", m.GetLongitude())
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
    {
        err := writer.WriteFloat64Value("speed", m.GetSpeed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("verticalAccuracy", m.GetVerticalAccuracy())
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
func (m *DeviceGeoLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAltitude sets the altitude property value. Altitude, given in meters above sea level
func (m *DeviceGeoLocation) SetAltitude(value *float64)() {
    m.altitude = value
}
// SetHeading sets the heading property value. Heading in degrees from true north
func (m *DeviceGeoLocation) SetHeading(value *float64)() {
    m.heading = value
}
// SetHorizontalAccuracy sets the horizontalAccuracy property value. Accuracy of longitude and latitude in meters
func (m *DeviceGeoLocation) SetHorizontalAccuracy(value *float64)() {
    m.horizontalAccuracy = value
}
// SetLastCollectedDateTime sets the lastCollectedDateTime property value. Time at which location was recorded, relative to UTC
func (m *DeviceGeoLocation) SetLastCollectedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastCollectedDateTime = value
}
// SetLatitude sets the latitude property value. Latitude coordinate of the device's location
func (m *DeviceGeoLocation) SetLatitude(value *float64)() {
    m.latitude = value
}
// SetLongitude sets the longitude property value. Longitude coordinate of the device's location
func (m *DeviceGeoLocation) SetLongitude(value *float64)() {
    m.longitude = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DeviceGeoLocation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSpeed sets the speed property value. Speed the device is traveling in meters per second
func (m *DeviceGeoLocation) SetSpeed(value *float64)() {
    m.speed = value
}
// SetVerticalAccuracy sets the verticalAccuracy property value. Accuracy of altitude in meters
func (m *DeviceGeoLocation) SetVerticalAccuracy(value *float64)() {
    m.verticalAccuracy = value
}
