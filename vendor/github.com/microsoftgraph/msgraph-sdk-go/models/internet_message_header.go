package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// InternetMessageHeader 
type InternetMessageHeader struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Represents the key in a key-value pair.
    name *string
    // The OdataType property
    odataType *string
    // The value in a key-value pair.
    value *string
}
// NewInternetMessageHeader instantiates a new internetMessageHeader and sets the default values.
func NewInternetMessageHeader()(*InternetMessageHeader) {
    m := &InternetMessageHeader{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateInternetMessageHeaderFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInternetMessageHeaderFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewInternetMessageHeader(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InternetMessageHeader) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InternetMessageHeader) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. Represents the key in a key-value pair.
func (m *InternetMessageHeader) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *InternetMessageHeader) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The value in a key-value pair.
func (m *InternetMessageHeader) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *InternetMessageHeader) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err := writer.WriteStringValue("value", m.GetValue())
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
func (m *InternetMessageHeader) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetName sets the name property value. Represents the key in a key-value pair.
func (m *InternetMessageHeader) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *InternetMessageHeader) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The value in a key-value pair.
func (m *InternetMessageHeader) SetValue(value *string)() {
    m.value = value
}
