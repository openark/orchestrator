package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChoiceColumn 
type ChoiceColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // If true, allows custom values that aren't in the configured choices.
    allowTextEntry *bool
    // The list of values available for this column.
    choices []string
    // How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons
    displayAs *string
    // The OdataType property
    odataType *string
}
// NewChoiceColumn instantiates a new choiceColumn and sets the default values.
func NewChoiceColumn()(*ChoiceColumn) {
    m := &ChoiceColumn{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChoiceColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChoiceColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChoiceColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChoiceColumn) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowTextEntry gets the allowTextEntry property value. If true, allows custom values that aren't in the configured choices.
func (m *ChoiceColumn) GetAllowTextEntry()(*bool) {
    return m.allowTextEntry
}
// GetChoices gets the choices property value. The list of values available for this column.
func (m *ChoiceColumn) GetChoices()([]string) {
    return m.choices
}
// GetDisplayAs gets the displayAs property value. How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons
func (m *ChoiceColumn) GetDisplayAs()(*string) {
    return m.displayAs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChoiceColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowTextEntry"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTextEntry(val)
        }
        return nil
    }
    res["choices"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetChoices(res)
        }
        return nil
    }
    res["displayAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAs(val)
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
func (m *ChoiceColumn) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *ChoiceColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowTextEntry", m.GetAllowTextEntry())
        if err != nil {
            return err
        }
    }
    if m.GetChoices() != nil {
        err := writer.WriteCollectionOfStringValues("choices", m.GetChoices())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
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
func (m *ChoiceColumn) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowTextEntry sets the allowTextEntry property value. If true, allows custom values that aren't in the configured choices.
func (m *ChoiceColumn) SetAllowTextEntry(value *bool)() {
    m.allowTextEntry = value
}
// SetChoices sets the choices property value. The list of values available for this column.
func (m *ChoiceColumn) SetChoices(value []string)() {
    m.choices = value
}
// SetDisplayAs sets the displayAs property value. How the choices are to be presented in the UX. Must be one of checkBoxes, dropDownMenu, or radioButtons
func (m *ChoiceColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChoiceColumn) SetOdataType(value *string)() {
    m.odataType = value
}
