package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeZoneInformation 
type TimeZoneInformation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // An identifier for the time zone.
    alias *string
    // A display string that represents the time zone.
    displayName *string
    // The OdataType property
    odataType *string
}
// NewTimeZoneInformation instantiates a new timeZoneInformation and sets the default values.
func NewTimeZoneInformation()(*TimeZoneInformation) {
    m := &TimeZoneInformation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimeZoneInformationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeZoneInformationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTimeZoneInformation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeZoneInformation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlias gets the alias property value. An identifier for the time zone.
func (m *TimeZoneInformation) GetAlias()(*string) {
    return m.alias
}
// GetDisplayName gets the displayName property value. A display string that represents the time zone.
func (m *TimeZoneInformation) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeZoneInformation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["alias"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlias(val)
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
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TimeZoneInformation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TimeZoneInformation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("alias", m.GetAlias())
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
func (m *TimeZoneInformation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlias sets the alias property value. An identifier for the time zone.
func (m *TimeZoneInformation) SetAlias(value *string)() {
    m.alias = value
}
// SetDisplayName sets the displayName property value. A display string that represents the time zone.
func (m *TimeZoneInformation) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TimeZoneInformation) SetOdataType(value *string)() {
    m.odataType = value
}
