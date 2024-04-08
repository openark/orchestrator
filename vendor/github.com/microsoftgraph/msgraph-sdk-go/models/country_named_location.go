package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CountryNamedLocation 
type CountryNamedLocation struct {
    NamedLocation
    // List of countries and/or regions in two-letter format specified by ISO 3166-2. Required.
    countriesAndRegions []string
    // Determines what method is used to decide which country the user is located in. Possible values are clientIpAddress(default) and authenticatorAppGps. Note: authenticatorAppGps is not yet supported in the Microsoft Cloud for US Government.
    countryLookupMethod *CountryLookupMethodType
    // true if IP addresses that don't map to a country or region should be included in the named location. Optional. Default value is false.
    includeUnknownCountriesAndRegions *bool
}
// NewCountryNamedLocation instantiates a new CountryNamedLocation and sets the default values.
func NewCountryNamedLocation()(*CountryNamedLocation) {
    m := &CountryNamedLocation{
        NamedLocation: *NewNamedLocation(),
    }
    return m
}
// CreateCountryNamedLocationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCountryNamedLocationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCountryNamedLocation(), nil
}
// GetCountriesAndRegions gets the countriesAndRegions property value. List of countries and/or regions in two-letter format specified by ISO 3166-2. Required.
func (m *CountryNamedLocation) GetCountriesAndRegions()([]string) {
    return m.countriesAndRegions
}
// GetCountryLookupMethod gets the countryLookupMethod property value. Determines what method is used to decide which country the user is located in. Possible values are clientIpAddress(default) and authenticatorAppGps. Note: authenticatorAppGps is not yet supported in the Microsoft Cloud for US Government.
func (m *CountryNamedLocation) GetCountryLookupMethod()(*CountryLookupMethodType) {
    return m.countryLookupMethod
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CountryNamedLocation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NamedLocation.GetFieldDeserializers()
    res["countriesAndRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCountriesAndRegions(res)
        }
        return nil
    }
    res["countryLookupMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCountryLookupMethodType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryLookupMethod(val.(*CountryLookupMethodType))
        }
        return nil
    }
    res["includeUnknownCountriesAndRegions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludeUnknownCountriesAndRegions(val)
        }
        return nil
    }
    return res
}
// GetIncludeUnknownCountriesAndRegions gets the includeUnknownCountriesAndRegions property value. true if IP addresses that don't map to a country or region should be included in the named location. Optional. Default value is false.
func (m *CountryNamedLocation) GetIncludeUnknownCountriesAndRegions()(*bool) {
    return m.includeUnknownCountriesAndRegions
}
// Serialize serializes information the current object
func (m *CountryNamedLocation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NamedLocation.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCountriesAndRegions() != nil {
        err = writer.WriteCollectionOfStringValues("countriesAndRegions", m.GetCountriesAndRegions())
        if err != nil {
            return err
        }
    }
    if m.GetCountryLookupMethod() != nil {
        cast := (*m.GetCountryLookupMethod()).String()
        err = writer.WriteStringValue("countryLookupMethod", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("includeUnknownCountriesAndRegions", m.GetIncludeUnknownCountriesAndRegions())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCountriesAndRegions sets the countriesAndRegions property value. List of countries and/or regions in two-letter format specified by ISO 3166-2. Required.
func (m *CountryNamedLocation) SetCountriesAndRegions(value []string)() {
    m.countriesAndRegions = value
}
// SetCountryLookupMethod sets the countryLookupMethod property value. Determines what method is used to decide which country the user is located in. Possible values are clientIpAddress(default) and authenticatorAppGps. Note: authenticatorAppGps is not yet supported in the Microsoft Cloud for US Government.
func (m *CountryNamedLocation) SetCountryLookupMethod(value *CountryLookupMethodType)() {
    m.countryLookupMethod = value
}
// SetIncludeUnknownCountriesAndRegions sets the includeUnknownCountriesAndRegions property value. true if IP addresses that don't map to a country or region should be included in the named location. Optional. Default value is false.
func (m *CountryNamedLocation) SetIncludeUnknownCountriesAndRegions(value *bool)() {
    m.includeUnknownCountriesAndRegions = value
}
