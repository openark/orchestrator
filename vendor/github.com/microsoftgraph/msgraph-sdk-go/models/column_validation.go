package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ColumnValidation 
type ColumnValidation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Default BCP 47 language tag for the description.
    defaultLanguage *string
    // Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
    descriptions []DisplayNameLocalizationable
    // The formula to validate column value. For examples, see Examples of common formulas in lists.
    formula *string
    // The OdataType property
    odataType *string
}
// NewColumnValidation instantiates a new columnValidation and sets the default values.
func NewColumnValidation()(*ColumnValidation) {
    m := &ColumnValidation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateColumnValidationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnValidationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewColumnValidation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnValidation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDefaultLanguage gets the defaultLanguage property value. Default BCP 47 language tag for the description.
func (m *ColumnValidation) GetDefaultLanguage()(*string) {
    return m.defaultLanguage
}
// GetDescriptions gets the descriptions property value. Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
func (m *ColumnValidation) GetDescriptions()([]DisplayNameLocalizationable) {
    return m.descriptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnValidation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["defaultLanguage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLanguage(val)
        }
        return nil
    }
    res["descriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDisplayNameLocalizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DisplayNameLocalizationable, len(val))
            for i, v := range val {
                res[i] = v.(DisplayNameLocalizationable)
            }
            m.SetDescriptions(res)
        }
        return nil
    }
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
    return res
}
// GetFormula gets the formula property value. The formula to validate column value. For examples, see Examples of common formulas in lists.
func (m *ColumnValidation) GetFormula()(*string) {
    return m.formula
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ColumnValidation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ColumnValidation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("defaultLanguage", m.GetDefaultLanguage())
        if err != nil {
            return err
        }
    }
    if m.GetDescriptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDescriptions()))
        for i, v := range m.GetDescriptions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("descriptions", cast)
        if err != nil {
            return err
        }
    }
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ColumnValidation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDefaultLanguage sets the defaultLanguage property value. Default BCP 47 language tag for the description.
func (m *ColumnValidation) SetDefaultLanguage(value *string)() {
    m.defaultLanguage = value
}
// SetDescriptions sets the descriptions property value. Localized messages that explain what is needed for this column's value to be considered valid. User will be prompted with this message if validation fails.
func (m *ColumnValidation) SetDescriptions(value []DisplayNameLocalizationable)() {
    m.descriptions = value
}
// SetFormula sets the formula property value. The formula to validate column value. For examples, see Examples of common formulas in lists.
func (m *ColumnValidation) SetFormula(value *string)() {
    m.formula = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ColumnValidation) SetOdataType(value *string)() {
    m.odataType = value
}
