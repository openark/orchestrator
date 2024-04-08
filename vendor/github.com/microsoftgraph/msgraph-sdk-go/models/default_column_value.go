package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DefaultColumnValue 
type DefaultColumnValue struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The formula used to compute the default value for this column.
    formula *string
    // The OdataType property
    odataType *string
    // The direct value to use as the default value for this column.
    value *string
}
// NewDefaultColumnValue instantiates a new defaultColumnValue and sets the default values.
func NewDefaultColumnValue()(*DefaultColumnValue) {
    m := &DefaultColumnValue{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDefaultColumnValueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultColumnValueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDefaultColumnValue(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DefaultColumnValue) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultColumnValue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["formula"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormula(val)
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
// GetFormula gets the formula property value. The formula used to compute the default value for this column.
func (m *DefaultColumnValue) GetFormula()(*string) {
    return m.formula
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DefaultColumnValue) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The direct value to use as the default value for this column.
func (m *DefaultColumnValue) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *DefaultColumnValue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("formula", m.GetFormula())
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
func (m *DefaultColumnValue) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFormula sets the formula property value. The formula used to compute the default value for this column.
func (m *DefaultColumnValue) SetFormula(value *string)() {
    m.formula = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DefaultColumnValue) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The direct value to use as the default value for this column.
func (m *DefaultColumnValue) SetValue(value *string)() {
    m.value = value
}
