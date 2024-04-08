package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PlannerCategoryDescriptions 
type PlannerCategoryDescriptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The label associated with Category 1
    category1 *string
    // The label associated with Category 10
    category10 *string
    // The label associated with Category 11
    category11 *string
    // The label associated with Category 12
    category12 *string
    // The label associated with Category 13
    category13 *string
    // The label associated with Category 14
    category14 *string
    // The label associated with Category 15
    category15 *string
    // The label associated with Category 16
    category16 *string
    // The label associated with Category 17
    category17 *string
    // The label associated with Category 18
    category18 *string
    // The label associated with Category 19
    category19 *string
    // The label associated with Category 2
    category2 *string
    // The label associated with Category 20
    category20 *string
    // The label associated with Category 21
    category21 *string
    // The label associated with Category 22
    category22 *string
    // The label associated with Category 23
    category23 *string
    // The label associated with Category 24
    category24 *string
    // The label associated with Category 25
    category25 *string
    // The label associated with Category 3
    category3 *string
    // The label associated with Category 4
    category4 *string
    // The label associated with Category 5
    category5 *string
    // The label associated with Category 6
    category6 *string
    // The label associated with Category 7
    category7 *string
    // The label associated with Category 8
    category8 *string
    // The label associated with Category 9
    category9 *string
    // The OdataType property
    odataType *string
}
// NewPlannerCategoryDescriptions instantiates a new plannerCategoryDescriptions and sets the default values.
func NewPlannerCategoryDescriptions()(*PlannerCategoryDescriptions) {
    m := &PlannerCategoryDescriptions{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePlannerCategoryDescriptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlannerCategoryDescriptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPlannerCategoryDescriptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PlannerCategoryDescriptions) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCategory1 gets the category1 property value. The label associated with Category 1
func (m *PlannerCategoryDescriptions) GetCategory1()(*string) {
    return m.category1
}
// GetCategory10 gets the category10 property value. The label associated with Category 10
func (m *PlannerCategoryDescriptions) GetCategory10()(*string) {
    return m.category10
}
// GetCategory11 gets the category11 property value. The label associated with Category 11
func (m *PlannerCategoryDescriptions) GetCategory11()(*string) {
    return m.category11
}
// GetCategory12 gets the category12 property value. The label associated with Category 12
func (m *PlannerCategoryDescriptions) GetCategory12()(*string) {
    return m.category12
}
// GetCategory13 gets the category13 property value. The label associated with Category 13
func (m *PlannerCategoryDescriptions) GetCategory13()(*string) {
    return m.category13
}
// GetCategory14 gets the category14 property value. The label associated with Category 14
func (m *PlannerCategoryDescriptions) GetCategory14()(*string) {
    return m.category14
}
// GetCategory15 gets the category15 property value. The label associated with Category 15
func (m *PlannerCategoryDescriptions) GetCategory15()(*string) {
    return m.category15
}
// GetCategory16 gets the category16 property value. The label associated with Category 16
func (m *PlannerCategoryDescriptions) GetCategory16()(*string) {
    return m.category16
}
// GetCategory17 gets the category17 property value. The label associated with Category 17
func (m *PlannerCategoryDescriptions) GetCategory17()(*string) {
    return m.category17
}
// GetCategory18 gets the category18 property value. The label associated with Category 18
func (m *PlannerCategoryDescriptions) GetCategory18()(*string) {
    return m.category18
}
// GetCategory19 gets the category19 property value. The label associated with Category 19
func (m *PlannerCategoryDescriptions) GetCategory19()(*string) {
    return m.category19
}
// GetCategory2 gets the category2 property value. The label associated with Category 2
func (m *PlannerCategoryDescriptions) GetCategory2()(*string) {
    return m.category2
}
// GetCategory20 gets the category20 property value. The label associated with Category 20
func (m *PlannerCategoryDescriptions) GetCategory20()(*string) {
    return m.category20
}
// GetCategory21 gets the category21 property value. The label associated with Category 21
func (m *PlannerCategoryDescriptions) GetCategory21()(*string) {
    return m.category21
}
// GetCategory22 gets the category22 property value. The label associated with Category 22
func (m *PlannerCategoryDescriptions) GetCategory22()(*string) {
    return m.category22
}
// GetCategory23 gets the category23 property value. The label associated with Category 23
func (m *PlannerCategoryDescriptions) GetCategory23()(*string) {
    return m.category23
}
// GetCategory24 gets the category24 property value. The label associated with Category 24
func (m *PlannerCategoryDescriptions) GetCategory24()(*string) {
    return m.category24
}
// GetCategory25 gets the category25 property value. The label associated with Category 25
func (m *PlannerCategoryDescriptions) GetCategory25()(*string) {
    return m.category25
}
// GetCategory3 gets the category3 property value. The label associated with Category 3
func (m *PlannerCategoryDescriptions) GetCategory3()(*string) {
    return m.category3
}
// GetCategory4 gets the category4 property value. The label associated with Category 4
func (m *PlannerCategoryDescriptions) GetCategory4()(*string) {
    return m.category4
}
// GetCategory5 gets the category5 property value. The label associated with Category 5
func (m *PlannerCategoryDescriptions) GetCategory5()(*string) {
    return m.category5
}
// GetCategory6 gets the category6 property value. The label associated with Category 6
func (m *PlannerCategoryDescriptions) GetCategory6()(*string) {
    return m.category6
}
// GetCategory7 gets the category7 property value. The label associated with Category 7
func (m *PlannerCategoryDescriptions) GetCategory7()(*string) {
    return m.category7
}
// GetCategory8 gets the category8 property value. The label associated with Category 8
func (m *PlannerCategoryDescriptions) GetCategory8()(*string) {
    return m.category8
}
// GetCategory9 gets the category9 property value. The label associated with Category 9
func (m *PlannerCategoryDescriptions) GetCategory9()(*string) {
    return m.category9
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlannerCategoryDescriptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["category1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory1(val)
        }
        return nil
    }
    res["category10"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory10(val)
        }
        return nil
    }
    res["category11"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory11(val)
        }
        return nil
    }
    res["category12"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory12(val)
        }
        return nil
    }
    res["category13"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory13(val)
        }
        return nil
    }
    res["category14"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory14(val)
        }
        return nil
    }
    res["category15"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory15(val)
        }
        return nil
    }
    res["category16"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory16(val)
        }
        return nil
    }
    res["category17"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory17(val)
        }
        return nil
    }
    res["category18"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory18(val)
        }
        return nil
    }
    res["category19"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory19(val)
        }
        return nil
    }
    res["category2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory2(val)
        }
        return nil
    }
    res["category20"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory20(val)
        }
        return nil
    }
    res["category21"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory21(val)
        }
        return nil
    }
    res["category22"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory22(val)
        }
        return nil
    }
    res["category23"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory23(val)
        }
        return nil
    }
    res["category24"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory24(val)
        }
        return nil
    }
    res["category25"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory25(val)
        }
        return nil
    }
    res["category3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory3(val)
        }
        return nil
    }
    res["category4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory4(val)
        }
        return nil
    }
    res["category5"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory5(val)
        }
        return nil
    }
    res["category6"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory6(val)
        }
        return nil
    }
    res["category7"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory7(val)
        }
        return nil
    }
    res["category8"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory8(val)
        }
        return nil
    }
    res["category9"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCategory9(val)
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
func (m *PlannerCategoryDescriptions) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *PlannerCategoryDescriptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("category1", m.GetCategory1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category10", m.GetCategory10())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category11", m.GetCategory11())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category12", m.GetCategory12())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category13", m.GetCategory13())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category14", m.GetCategory14())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category15", m.GetCategory15())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category16", m.GetCategory16())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category17", m.GetCategory17())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category18", m.GetCategory18())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category19", m.GetCategory19())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category2", m.GetCategory2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category20", m.GetCategory20())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category21", m.GetCategory21())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category22", m.GetCategory22())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category23", m.GetCategory23())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category24", m.GetCategory24())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category25", m.GetCategory25())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category3", m.GetCategory3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category4", m.GetCategory4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category5", m.GetCategory5())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category6", m.GetCategory6())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category7", m.GetCategory7())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category8", m.GetCategory8())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("category9", m.GetCategory9())
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
func (m *PlannerCategoryDescriptions) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCategory1 sets the category1 property value. The label associated with Category 1
func (m *PlannerCategoryDescriptions) SetCategory1(value *string)() {
    m.category1 = value
}
// SetCategory10 sets the category10 property value. The label associated with Category 10
func (m *PlannerCategoryDescriptions) SetCategory10(value *string)() {
    m.category10 = value
}
// SetCategory11 sets the category11 property value. The label associated with Category 11
func (m *PlannerCategoryDescriptions) SetCategory11(value *string)() {
    m.category11 = value
}
// SetCategory12 sets the category12 property value. The label associated with Category 12
func (m *PlannerCategoryDescriptions) SetCategory12(value *string)() {
    m.category12 = value
}
// SetCategory13 sets the category13 property value. The label associated with Category 13
func (m *PlannerCategoryDescriptions) SetCategory13(value *string)() {
    m.category13 = value
}
// SetCategory14 sets the category14 property value. The label associated with Category 14
func (m *PlannerCategoryDescriptions) SetCategory14(value *string)() {
    m.category14 = value
}
// SetCategory15 sets the category15 property value. The label associated with Category 15
func (m *PlannerCategoryDescriptions) SetCategory15(value *string)() {
    m.category15 = value
}
// SetCategory16 sets the category16 property value. The label associated with Category 16
func (m *PlannerCategoryDescriptions) SetCategory16(value *string)() {
    m.category16 = value
}
// SetCategory17 sets the category17 property value. The label associated with Category 17
func (m *PlannerCategoryDescriptions) SetCategory17(value *string)() {
    m.category17 = value
}
// SetCategory18 sets the category18 property value. The label associated with Category 18
func (m *PlannerCategoryDescriptions) SetCategory18(value *string)() {
    m.category18 = value
}
// SetCategory19 sets the category19 property value. The label associated with Category 19
func (m *PlannerCategoryDescriptions) SetCategory19(value *string)() {
    m.category19 = value
}
// SetCategory2 sets the category2 property value. The label associated with Category 2
func (m *PlannerCategoryDescriptions) SetCategory2(value *string)() {
    m.category2 = value
}
// SetCategory20 sets the category20 property value. The label associated with Category 20
func (m *PlannerCategoryDescriptions) SetCategory20(value *string)() {
    m.category20 = value
}
// SetCategory21 sets the category21 property value. The label associated with Category 21
func (m *PlannerCategoryDescriptions) SetCategory21(value *string)() {
    m.category21 = value
}
// SetCategory22 sets the category22 property value. The label associated with Category 22
func (m *PlannerCategoryDescriptions) SetCategory22(value *string)() {
    m.category22 = value
}
// SetCategory23 sets the category23 property value. The label associated with Category 23
func (m *PlannerCategoryDescriptions) SetCategory23(value *string)() {
    m.category23 = value
}
// SetCategory24 sets the category24 property value. The label associated with Category 24
func (m *PlannerCategoryDescriptions) SetCategory24(value *string)() {
    m.category24 = value
}
// SetCategory25 sets the category25 property value. The label associated with Category 25
func (m *PlannerCategoryDescriptions) SetCategory25(value *string)() {
    m.category25 = value
}
// SetCategory3 sets the category3 property value. The label associated with Category 3
func (m *PlannerCategoryDescriptions) SetCategory3(value *string)() {
    m.category3 = value
}
// SetCategory4 sets the category4 property value. The label associated with Category 4
func (m *PlannerCategoryDescriptions) SetCategory4(value *string)() {
    m.category4 = value
}
// SetCategory5 sets the category5 property value. The label associated with Category 5
func (m *PlannerCategoryDescriptions) SetCategory5(value *string)() {
    m.category5 = value
}
// SetCategory6 sets the category6 property value. The label associated with Category 6
func (m *PlannerCategoryDescriptions) SetCategory6(value *string)() {
    m.category6 = value
}
// SetCategory7 sets the category7 property value. The label associated with Category 7
func (m *PlannerCategoryDescriptions) SetCategory7(value *string)() {
    m.category7 = value
}
// SetCategory8 sets the category8 property value. The label associated with Category 8
func (m *PlannerCategoryDescriptions) SetCategory8(value *string)() {
    m.category8 = value
}
// SetCategory9 sets the category9 property value. The label associated with Category 9
func (m *PlannerCategoryDescriptions) SetCategory9(value *string)() {
    m.category9 = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PlannerCategoryDescriptions) SetOdataType(value *string)() {
    m.odataType = value
}
