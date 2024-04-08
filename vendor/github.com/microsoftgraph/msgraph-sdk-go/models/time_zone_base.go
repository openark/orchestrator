package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TimeZoneBase 
type TimeZoneBase struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
    name *string
    // The OdataType property
    odataType *string
}
// NewTimeZoneBase instantiates a new timeZoneBase and sets the default values.
func NewTimeZoneBase()(*TimeZoneBase) {
    m := &TimeZoneBase{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTimeZoneBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTimeZoneBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.customTimeZone":
                        return NewCustomTimeZone(), nil
                }
            }
        }
    }
    return NewTimeZoneBase(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TimeZoneBase) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TimeZoneBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
// GetName gets the name property value. The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
func (m *TimeZoneBase) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TimeZoneBase) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *TimeZoneBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *TimeZoneBase) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. The name of a time zone. It can be a standard time zone name such as 'Hawaii-Aleutian Standard Time', or 'Customized Time Zone' for a custom time zone.
func (m *TimeZoneBase) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TimeZoneBase) SetOdataType(value *string)() {
    m.odataType = value
}
