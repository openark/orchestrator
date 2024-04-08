package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrinterLocation 
type PrinterLocation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The altitude, in meters, that the printer is located at.
    altitudeInMeters *int32
    // The building that the printer is located in.
    building *string
    // The city that the printer is located in.
    city *string
    // The country or region that the printer is located in.
    countryOrRegion *string
    // The floor that the printer is located on. Only numerical values are supported right now.
    floor *string
    // The description of the floor that the printer is located on.
    floorDescription *string
    // The latitude that the printer is located at.
    latitude *float64
    // The longitude that the printer is located at.
    longitude *float64
    // The OdataType property
    odataType *string
    // The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
    organization []string
    // The postal code that the printer is located in.
    postalCode *string
    // The description of the room that the printer is located in.
    roomDescription *string
    // The room that the printer is located in. Only numerical values are supported right now.
    roomName *string
    // The site that the printer is located in.
    site *string
    // The state or province that the printer is located in.
    stateOrProvince *string
    // The street address where the printer is located.
    streetAddress *string
    // The subdivision that the printer is located in. The elements should be in hierarchical order.
    subdivision []string
    // The subunit property
    subunit []string
}
// NewPrinterLocation instantiates a new printerLocation and sets the default values.
func NewPrinterLocation()(*PrinterLocation) {
    m := &PrinterLocation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePrinterLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrinterLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrinterLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PrinterLocation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAltitudeInMeters gets the altitudeInMeters property value. The altitude, in meters, that the printer is located at.
func (m *PrinterLocation) GetAltitudeInMeters()(*int32) {
    return m.altitudeInMeters
}
// GetBuilding gets the building property value. The building that the printer is located in.
func (m *PrinterLocation) GetBuilding()(*string) {
    return m.building
}
// GetCity gets the city property value. The city that the printer is located in.
func (m *PrinterLocation) GetCity()(*string) {
    return m.city
}
// GetCountryOrRegion gets the countryOrRegion property value. The country or region that the printer is located in.
func (m *PrinterLocation) GetCountryOrRegion()(*string) {
    return m.countryOrRegion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrinterLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["altitudeInMeters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAltitudeInMeters(val)
        }
        return nil
    }
    res["building"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBuilding(val)
        }
        return nil
    }
    res["city"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCity(val)
        }
        return nil
    }
    res["countryOrRegion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryOrRegion(val)
        }
        return nil
    }
    res["floor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloor(val)
        }
        return nil
    }
    res["floorDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFloorDescription(val)
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
    res["organization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetOrganization(res)
        }
        return nil
    }
    res["postalCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPostalCode(val)
        }
        return nil
    }
    res["roomDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomDescription(val)
        }
        return nil
    }
    res["roomName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoomName(val)
        }
        return nil
    }
    res["site"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSite(val)
        }
        return nil
    }
    res["stateOrProvince"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStateOrProvince(val)
        }
        return nil
    }
    res["streetAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreetAddress(val)
        }
        return nil
    }
    res["subdivision"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSubdivision(res)
        }
        return nil
    }
    res["subunit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSubunit(res)
        }
        return nil
    }
    return res
}
// GetFloor gets the floor property value. The floor that the printer is located on. Only numerical values are supported right now.
func (m *PrinterLocation) GetFloor()(*string) {
    return m.floor
}
// GetFloorDescription gets the floorDescription property value. The description of the floor that the printer is located on.
func (m *PrinterLocation) GetFloorDescription()(*string) {
    return m.floorDescription
}
// GetLatitude gets the latitude property value. The latitude that the printer is located at.
func (m *PrinterLocation) GetLatitude()(*float64) {
    return m.latitude
}
// GetLongitude gets the longitude property value. The longitude that the printer is located at.
func (m *PrinterLocation) GetLongitude()(*float64) {
    return m.longitude
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PrinterLocation) GetOdataType()(*string) {
    return m.odataType
}
// GetOrganization gets the organization property value. The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
func (m *PrinterLocation) GetOrganization()([]string) {
    return m.organization
}
// GetPostalCode gets the postalCode property value. The postal code that the printer is located in.
func (m *PrinterLocation) GetPostalCode()(*string) {
    return m.postalCode
}
// GetRoomDescription gets the roomDescription property value. The description of the room that the printer is located in.
func (m *PrinterLocation) GetRoomDescription()(*string) {
    return m.roomDescription
}
// GetRoomName gets the roomName property value. The room that the printer is located in. Only numerical values are supported right now.
func (m *PrinterLocation) GetRoomName()(*string) {
    return m.roomName
}
// GetSite gets the site property value. The site that the printer is located in.
func (m *PrinterLocation) GetSite()(*string) {
    return m.site
}
// GetStateOrProvince gets the stateOrProvince property value. The state or province that the printer is located in.
func (m *PrinterLocation) GetStateOrProvince()(*string) {
    return m.stateOrProvince
}
// GetStreetAddress gets the streetAddress property value. The street address where the printer is located.
func (m *PrinterLocation) GetStreetAddress()(*string) {
    return m.streetAddress
}
// GetSubdivision gets the subdivision property value. The subdivision that the printer is located in. The elements should be in hierarchical order.
func (m *PrinterLocation) GetSubdivision()([]string) {
    return m.subdivision
}
// GetSubunit gets the subunit property value. The subunit property
func (m *PrinterLocation) GetSubunit()([]string) {
    return m.subunit
}
// Serialize serializes information the current object
func (m *PrinterLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("altitudeInMeters", m.GetAltitudeInMeters())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("building", m.GetBuilding())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("city", m.GetCity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("countryOrRegion", m.GetCountryOrRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("floor", m.GetFloor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("floorDescription", m.GetFloorDescription())
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
    if m.GetOrganization() != nil {
        err := writer.WriteCollectionOfStringValues("organization", m.GetOrganization())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("postalCode", m.GetPostalCode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roomDescription", m.GetRoomDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("roomName", m.GetRoomName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("site", m.GetSite())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("stateOrProvince", m.GetStateOrProvince())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("streetAddress", m.GetStreetAddress())
        if err != nil {
            return err
        }
    }
    if m.GetSubdivision() != nil {
        err := writer.WriteCollectionOfStringValues("subdivision", m.GetSubdivision())
        if err != nil {
            return err
        }
    }
    if m.GetSubunit() != nil {
        err := writer.WriteCollectionOfStringValues("subunit", m.GetSubunit())
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
func (m *PrinterLocation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAltitudeInMeters sets the altitudeInMeters property value. The altitude, in meters, that the printer is located at.
func (m *PrinterLocation) SetAltitudeInMeters(value *int32)() {
    m.altitudeInMeters = value
}
// SetBuilding sets the building property value. The building that the printer is located in.
func (m *PrinterLocation) SetBuilding(value *string)() {
    m.building = value
}
// SetCity sets the city property value. The city that the printer is located in.
func (m *PrinterLocation) SetCity(value *string)() {
    m.city = value
}
// SetCountryOrRegion sets the countryOrRegion property value. The country or region that the printer is located in.
func (m *PrinterLocation) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// SetFloor sets the floor property value. The floor that the printer is located on. Only numerical values are supported right now.
func (m *PrinterLocation) SetFloor(value *string)() {
    m.floor = value
}
// SetFloorDescription sets the floorDescription property value. The description of the floor that the printer is located on.
func (m *PrinterLocation) SetFloorDescription(value *string)() {
    m.floorDescription = value
}
// SetLatitude sets the latitude property value. The latitude that the printer is located at.
func (m *PrinterLocation) SetLatitude(value *float64)() {
    m.latitude = value
}
// SetLongitude sets the longitude property value. The longitude that the printer is located at.
func (m *PrinterLocation) SetLongitude(value *float64)() {
    m.longitude = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PrinterLocation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrganization sets the organization property value. The organizational hierarchy that the printer belongs to. The elements should be in hierarchical order.
func (m *PrinterLocation) SetOrganization(value []string)() {
    m.organization = value
}
// SetPostalCode sets the postalCode property value. The postal code that the printer is located in.
func (m *PrinterLocation) SetPostalCode(value *string)() {
    m.postalCode = value
}
// SetRoomDescription sets the roomDescription property value. The description of the room that the printer is located in.
func (m *PrinterLocation) SetRoomDescription(value *string)() {
    m.roomDescription = value
}
// SetRoomName sets the roomName property value. The room that the printer is located in. Only numerical values are supported right now.
func (m *PrinterLocation) SetRoomName(value *string)() {
    m.roomName = value
}
// SetSite sets the site property value. The site that the printer is located in.
func (m *PrinterLocation) SetSite(value *string)() {
    m.site = value
}
// SetStateOrProvince sets the stateOrProvince property value. The state or province that the printer is located in.
func (m *PrinterLocation) SetStateOrProvince(value *string)() {
    m.stateOrProvince = value
}
// SetStreetAddress sets the streetAddress property value. The street address where the printer is located.
func (m *PrinterLocation) SetStreetAddress(value *string)() {
    m.streetAddress = value
}
// SetSubdivision sets the subdivision property value. The subdivision that the printer is located in. The elements should be in hierarchical order.
func (m *PrinterLocation) SetSubdivision(value []string)() {
    m.subdivision = value
}
// SetSubunit sets the subunit property value. The subunit property
func (m *PrinterLocation) SetSubunit(value []string)() {
    m.subunit = value
}
