package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CurrencyColumn 
type CurrencyColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies the locale from which to infer the currency symbol.
    locale *string
    // The OdataType property
    odataType *string
}
// NewCurrencyColumn instantiates a new currencyColumn and sets the default values.
func NewCurrencyColumn()(*CurrencyColumn) {
    m := &CurrencyColumn{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateCurrencyColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCurrencyColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCurrencyColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CurrencyColumn) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CurrencyColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
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
// GetLocale gets the locale property value. Specifies the locale from which to infer the currency symbol.
func (m *CurrencyColumn) GetLocale()(*string) {
    return m.locale
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *CurrencyColumn) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *CurrencyColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("locale", m.GetLocale())
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
func (m *CurrencyColumn) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetLocale sets the locale property value. Specifies the locale from which to infer the currency symbol.
func (m *CurrencyColumn) SetLocale(value *string)() {
    m.locale = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *CurrencyColumn) SetOdataType(value *string)() {
    m.odataType = value
}
