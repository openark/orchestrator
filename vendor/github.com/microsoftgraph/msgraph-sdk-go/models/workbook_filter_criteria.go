package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbookFilterCriteria 
type WorkbookFilterCriteria struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The color property
    color *string
    // The criterion1 property
    criterion1 *string
    // The criterion2 property
    criterion2 *string
    // The dynamicCriteria property
    dynamicCriteria *string
    // The filterOn property
    filterOn *string
    // The icon property
    icon WorkbookIconable
    // The OdataType property
    odataType *string
    // The operator property
    operator *string
    // The values property
    values Jsonable
}
// NewWorkbookFilterCriteria instantiates a new workbookFilterCriteria and sets the default values.
func NewWorkbookFilterCriteria()(*WorkbookFilterCriteria) {
    m := &WorkbookFilterCriteria{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateWorkbookFilterCriteriaFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbookFilterCriteriaFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbookFilterCriteria(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbookFilterCriteria) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColor gets the color property value. The color property
func (m *WorkbookFilterCriteria) GetColor()(*string) {
    return m.color
}
// GetCriterion1 gets the criterion1 property value. The criterion1 property
func (m *WorkbookFilterCriteria) GetCriterion1()(*string) {
    return m.criterion1
}
// GetCriterion2 gets the criterion2 property value. The criterion2 property
func (m *WorkbookFilterCriteria) GetCriterion2()(*string) {
    return m.criterion2
}
// GetDynamicCriteria gets the dynamicCriteria property value. The dynamicCriteria property
func (m *WorkbookFilterCriteria) GetDynamicCriteria()(*string) {
    return m.dynamicCriteria
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbookFilterCriteria) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val)
        }
        return nil
    }
    res["criterion1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriterion1(val)
        }
        return nil
    }
    res["criterion2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCriterion2(val)
        }
        return nil
    }
    res["dynamicCriteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDynamicCriteria(val)
        }
        return nil
    }
    res["filterOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilterOn(val)
        }
        return nil
    }
    res["icon"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkbookIconFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIcon(val.(WorkbookIconable))
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
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val)
        }
        return nil
    }
    res["values"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateJsonFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValues(val.(Jsonable))
        }
        return nil
    }
    return res
}
// GetFilterOn gets the filterOn property value. The filterOn property
func (m *WorkbookFilterCriteria) GetFilterOn()(*string) {
    return m.filterOn
}
// GetIcon gets the icon property value. The icon property
func (m *WorkbookFilterCriteria) GetIcon()(WorkbookIconable) {
    return m.icon
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *WorkbookFilterCriteria) GetOdataType()(*string) {
    return m.odataType
}
// GetOperator gets the operator property value. The operator property
func (m *WorkbookFilterCriteria) GetOperator()(*string) {
    return m.operator
}
// GetValues gets the values property value. The values property
func (m *WorkbookFilterCriteria) GetValues()(Jsonable) {
    return m.values
}
// Serialize serializes information the current object
func (m *WorkbookFilterCriteria) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("color", m.GetColor())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criterion1", m.GetCriterion1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("criterion2", m.GetCriterion2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dynamicCriteria", m.GetDynamicCriteria())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("filterOn", m.GetFilterOn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("icon", m.GetIcon())
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
        err := writer.WriteStringValue("operator", m.GetOperator())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("values", m.GetValues())
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
func (m *WorkbookFilterCriteria) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColor sets the color property value. The color property
func (m *WorkbookFilterCriteria) SetColor(value *string)() {
    m.color = value
}
// SetCriterion1 sets the criterion1 property value. The criterion1 property
func (m *WorkbookFilterCriteria) SetCriterion1(value *string)() {
    m.criterion1 = value
}
// SetCriterion2 sets the criterion2 property value. The criterion2 property
func (m *WorkbookFilterCriteria) SetCriterion2(value *string)() {
    m.criterion2 = value
}
// SetDynamicCriteria sets the dynamicCriteria property value. The dynamicCriteria property
func (m *WorkbookFilterCriteria) SetDynamicCriteria(value *string)() {
    m.dynamicCriteria = value
}
// SetFilterOn sets the filterOn property value. The filterOn property
func (m *WorkbookFilterCriteria) SetFilterOn(value *string)() {
    m.filterOn = value
}
// SetIcon sets the icon property value. The icon property
func (m *WorkbookFilterCriteria) SetIcon(value WorkbookIconable)() {
    m.icon = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *WorkbookFilterCriteria) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOperator sets the operator property value. The operator property
func (m *WorkbookFilterCriteria) SetOperator(value *string)() {
    m.operator = value
}
// SetValues sets the values property value. The values property
func (m *WorkbookFilterCriteria) SetValues(value Jsonable)() {
    m.values = value
}
