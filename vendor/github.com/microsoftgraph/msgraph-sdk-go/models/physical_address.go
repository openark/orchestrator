package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PhysicalAddress 
type PhysicalAddress struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The city.
    city *string
    // The country or region. It's a free-format string value, for example, 'United States'.
    countryOrRegion *string
    // The OdataType property
    odataType *string
    // The postal code.
    postalCode *string
    // The state.
    state *string
    // The street.
    street *string
}
// NewPhysicalAddress instantiates a new physicalAddress and sets the default values.
func NewPhysicalAddress()(*PhysicalAddress) {
    m := &PhysicalAddress{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePhysicalAddressFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePhysicalAddressFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPhysicalAddress(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PhysicalAddress) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCity gets the city property value. The city.
func (m *PhysicalAddress) GetCity()(*string) {
    return m.city
}
// GetCountryOrRegion gets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalAddress) GetCountryOrRegion()(*string) {
    return m.countryOrRegion
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PhysicalAddress) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val)
        }
        return nil
    }
    res["street"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStreet(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PhysicalAddress) GetOdataType()(*string) {
    return m.odataType
}
// GetPostalCode gets the postalCode property value. The postal code.
func (m *PhysicalAddress) GetPostalCode()(*string) {
    return m.postalCode
}
// GetState gets the state property value. The state.
func (m *PhysicalAddress) GetState()(*string) {
    return m.state
}
// GetStreet gets the street property value. The street.
func (m *PhysicalAddress) GetStreet()(*string) {
    return m.street
}
// Serialize serializes information the current object
func (m *PhysicalAddress) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
        err := writer.WriteStringValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("street", m.GetStreet())
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
func (m *PhysicalAddress) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCity sets the city property value. The city.
func (m *PhysicalAddress) SetCity(value *string)() {
    m.city = value
}
// SetCountryOrRegion sets the countryOrRegion property value. The country or region. It's a free-format string value, for example, 'United States'.
func (m *PhysicalAddress) SetCountryOrRegion(value *string)() {
    m.countryOrRegion = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PhysicalAddress) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPostalCode sets the postalCode property value. The postal code.
func (m *PhysicalAddress) SetPostalCode(value *string)() {
    m.postalCode = value
}
// SetState sets the state property value. The state.
func (m *PhysicalAddress) SetState(value *string)() {
    m.state = value
}
// SetStreet sets the street property value. The street.
func (m *PhysicalAddress) SetStreet(value *string)() {
    m.street = value
}
