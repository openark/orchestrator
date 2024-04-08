package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GeoCoordinates 
type GeoCoordinates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
    altitude *float64
    // Optional. The latitude, in decimal, for the item. Read-only.
    latitude *float64
    // Optional. The longitude, in decimal, for the item. Read-only.
    longitude *float64
    // The OdataType property
    odataType *string
}
// NewGeoCoordinates instantiates a new geoCoordinates and sets the default values.
func NewGeoCoordinates()(*GeoCoordinates) {
    m := &GeoCoordinates{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateGeoCoordinatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGeoCoordinatesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGeoCoordinates(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GeoCoordinates) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAltitude gets the altitude property value. Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
func (m *GeoCoordinates) GetAltitude()(*float64) {
    return m.altitude
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GeoCoordinates) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    return res
}
// GetLatitude gets the latitude property value. Optional. The latitude, in decimal, for the item. Read-only.
func (m *GeoCoordinates) GetLatitude()(*float64) {
    return m.latitude
}
// GetLongitude gets the longitude property value. Optional. The longitude, in decimal, for the item. Read-only.
func (m *GeoCoordinates) GetLongitude()(*float64) {
    return m.longitude
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *GeoCoordinates) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *GeoCoordinates) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("altitude", m.GetAltitude())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GeoCoordinates) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAltitude sets the altitude property value. Optional. The altitude (height), in feet,  above sea level for the item. Read-only.
func (m *GeoCoordinates) SetAltitude(value *float64)() {
    m.altitude = value
}
// SetLatitude sets the latitude property value. Optional. The latitude, in decimal, for the item. Read-only.
func (m *GeoCoordinates) SetLatitude(value *float64)() {
    m.latitude = value
}
// SetLongitude sets the longitude property value. Optional. The longitude, in decimal, for the item. Read-only.
func (m *GeoCoordinates) SetLongitude(value *float64)() {
    m.longitude = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *GeoCoordinates) SetOdataType(value *string)() {
    m.odataType = value
}
