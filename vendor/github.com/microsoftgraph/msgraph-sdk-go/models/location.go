package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Location 
type Location struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The street address of the location.
    address PhysicalAddressable
    // The geographic coordinates and elevation of the location.
    coordinates OutlookGeoCoordinatesable
    // The name associated with the location.
    displayName *string
    // Optional email address of the location.
    locationEmailAddress *string
    // The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
    locationType *LocationType
    // Optional URI representing the location.
    locationUri *string
    // The OdataType property
    odataType *string
    // For internal use only.
    uniqueId *string
    // For internal use only.
    uniqueIdType *LocationUniqueIdType
}
// NewLocation instantiates a new location and sets the default values.
func NewLocation()(*Location) {
    m := &Location{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.locationConstraintItem":
                        return NewLocationConstraintItem(), nil
                }
            }
        }
    }
    return NewLocation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Location) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAddress gets the address property value. The street address of the location.
func (m *Location) GetAddress()(PhysicalAddressable) {
    return m.address
}
// GetCoordinates gets the coordinates property value. The geographic coordinates and elevation of the location.
func (m *Location) GetCoordinates()(OutlookGeoCoordinatesable) {
    return m.coordinates
}
// GetDisplayName gets the displayName property value. The name associated with the location.
func (m *Location) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Location) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["coordinates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCoordinates(val.(OutlookGeoCoordinatesable))
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["locationEmailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationEmailAddress(val)
        }
        return nil
    }
    res["locationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationType(val.(*LocationType))
        }
        return nil
    }
    res["locationUri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocationUri(val)
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
    res["uniqueId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueId(val)
        }
        return nil
    }
    res["uniqueIdType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLocationUniqueIdType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueIdType(val.(*LocationUniqueIdType))
        }
        return nil
    }
    return res
}
// GetLocationEmailAddress gets the locationEmailAddress property value. Optional email address of the location.
func (m *Location) GetLocationEmailAddress()(*string) {
    return m.locationEmailAddress
}
// GetLocationType gets the locationType property value. The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
func (m *Location) GetLocationType()(*LocationType) {
    return m.locationType
}
// GetLocationUri gets the locationUri property value. Optional URI representing the location.
func (m *Location) GetLocationUri()(*string) {
    return m.locationUri
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Location) GetOdataType()(*string) {
    return m.odataType
}
// GetUniqueId gets the uniqueId property value. For internal use only.
func (m *Location) GetUniqueId()(*string) {
    return m.uniqueId
}
// GetUniqueIdType gets the uniqueIdType property value. For internal use only.
func (m *Location) GetUniqueIdType()(*LocationUniqueIdType) {
    return m.uniqueIdType
}
// Serialize serializes information the current object
func (m *Location) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("coordinates", m.GetCoordinates())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locationEmailAddress", m.GetLocationEmailAddress())
        if err != nil {
            return err
        }
    }
    if m.GetLocationType() != nil {
        cast := (*m.GetLocationType()).String()
        err := writer.WriteStringValue("locationType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("locationUri", m.GetLocationUri())
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
        err := writer.WriteStringValue("uniqueId", m.GetUniqueId())
        if err != nil {
            return err
        }
    }
    if m.GetUniqueIdType() != nil {
        cast := (*m.GetUniqueIdType()).String()
        err := writer.WriteStringValue("uniqueIdType", &cast)
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
func (m *Location) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAddress sets the address property value. The street address of the location.
func (m *Location) SetAddress(value PhysicalAddressable)() {
    m.address = value
}
// SetCoordinates sets the coordinates property value. The geographic coordinates and elevation of the location.
func (m *Location) SetCoordinates(value OutlookGeoCoordinatesable)() {
    m.coordinates = value
}
// SetDisplayName sets the displayName property value. The name associated with the location.
func (m *Location) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLocationEmailAddress sets the locationEmailAddress property value. Optional email address of the location.
func (m *Location) SetLocationEmailAddress(value *string)() {
    m.locationEmailAddress = value
}
// SetLocationType sets the locationType property value. The type of location. The possible values are: default, conferenceRoom, homeAddress, businessAddress,geoCoordinates, streetAddress, hotel, restaurant, localBusiness, postalAddress. Read-only.
func (m *Location) SetLocationType(value *LocationType)() {
    m.locationType = value
}
// SetLocationUri sets the locationUri property value. Optional URI representing the location.
func (m *Location) SetLocationUri(value *string)() {
    m.locationUri = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Location) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUniqueId sets the uniqueId property value. For internal use only.
func (m *Location) SetUniqueId(value *string)() {
    m.uniqueId = value
}
// SetUniqueIdType sets the uniqueIdType property value. For internal use only.
func (m *Location) SetUniqueIdType(value *LocationUniqueIdType)() {
    m.uniqueIdType = value
}
