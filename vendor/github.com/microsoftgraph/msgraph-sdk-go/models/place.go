package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Place 
type Place struct {
    Entity
    // The street address of the place.
    address PhysicalAddressable
    // The name associated with the place.
    displayName *string
    // Specifies the place location in latitude, longitude and (optionally) altitude coordinates.
    geoCoordinates OutlookGeoCoordinatesable
    // The phone number of the place.
    phone *string
}
// NewPlace instantiates a new place and sets the default values.
func NewPlace()(*Place) {
    m := &Place{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePlaceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlaceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.room":
                        return NewRoom(), nil
                    case "#microsoft.graph.roomList":
                        return NewRoomList(), nil
                }
            }
        }
    }
    return NewPlace(), nil
}
// GetAddress gets the address property value. The street address of the place.
func (m *Place) GetAddress()(PhysicalAddressable) {
    return m.address
}
// GetDisplayName gets the displayName property value. The name associated with the place.
func (m *Place) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Place) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["geoCoordinates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOutlookGeoCoordinatesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGeoCoordinates(val.(OutlookGeoCoordinatesable))
        }
        return nil
    }
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    return res
}
// GetGeoCoordinates gets the geoCoordinates property value. Specifies the place location in latitude, longitude and (optionally) altitude coordinates.
func (m *Place) GetGeoCoordinates()(OutlookGeoCoordinatesable) {
    return m.geoCoordinates
}
// GetPhone gets the phone property value. The phone number of the place.
func (m *Place) GetPhone()(*string) {
    return m.phone
}
// Serialize serializes information the current object
func (m *Place) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("geoCoordinates", m.GetGeoCoordinates())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. The street address of the place.
func (m *Place) SetAddress(value PhysicalAddressable)() {
    m.address = value
}
// SetDisplayName sets the displayName property value. The name associated with the place.
func (m *Place) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGeoCoordinates sets the geoCoordinates property value. Specifies the place location in latitude, longitude and (optionally) altitude coordinates.
func (m *Place) SetGeoCoordinates(value OutlookGeoCoordinatesable)() {
    m.geoCoordinates = value
}
// SetPhone sets the phone property value. The phone number of the place.
func (m *Place) SetPhone(value *string)() {
    m.phone = value
}
