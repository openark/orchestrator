package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Phone 
type Phone struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The language property
    language *string
    // The phone number.
    number *string
    // The OdataType property
    odataType *string
    // The region property
    region *string
    // The type of phone number. The possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio.
    typeEscaped *PhoneType
}
// NewPhone instantiates a new phone and sets the default values.
func NewPhone()(*Phone) {
    m := &Phone{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePhoneFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePhoneFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPhone(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Phone) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Phone) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val)
        }
        return nil
    }
    res["number"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumber(val)
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
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePhoneType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*PhoneType))
        }
        return nil
    }
    return res
}
// GetLanguage gets the language property value. The language property
func (m *Phone) GetLanguage()(*string) {
    return m.language
}
// GetNumber gets the number property value. The phone number.
func (m *Phone) GetNumber()(*string) {
    return m.number
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Phone) GetOdataType()(*string) {
    return m.odataType
}
// GetRegion gets the region property value. The region property
func (m *Phone) GetRegion()(*string) {
    return m.region
}
// GetType gets the type property value. The type of phone number. The possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio.
func (m *Phone) GetType()(*PhoneType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Phone) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("number", m.GetNumber())
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
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *Phone) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLanguage sets the language property value. The language property
func (m *Phone) SetLanguage(value *string)() {
    m.language = value
}
// SetNumber sets the number property value. The phone number.
func (m *Phone) SetNumber(value *string)() {
    m.number = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Phone) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRegion sets the region property value. The region property
func (m *Phone) SetRegion(value *string)() {
    m.region = value
}
// SetType sets the type property value. The type of phone number. The possible values are: home, business, mobile, other, assistant, homeFax, businessFax, otherFax, pager, radio.
func (m *Phone) SetType(value *PhoneType)() {
    m.typeEscaped = value
}
