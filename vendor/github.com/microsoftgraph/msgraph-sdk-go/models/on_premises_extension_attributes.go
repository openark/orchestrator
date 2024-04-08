package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OnPremisesExtensionAttributes 
type OnPremisesExtensionAttributes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // First customizable extension attribute.
    extensionAttribute1 *string
    // Tenth customizable extension attribute.
    extensionAttribute10 *string
    // Eleventh customizable extension attribute.
    extensionAttribute11 *string
    // Twelfth customizable extension attribute.
    extensionAttribute12 *string
    // Thirteenth customizable extension attribute.
    extensionAttribute13 *string
    // Fourteenth customizable extension attribute.
    extensionAttribute14 *string
    // Fifteenth customizable extension attribute.
    extensionAttribute15 *string
    // Second customizable extension attribute.
    extensionAttribute2 *string
    // Third customizable extension attribute.
    extensionAttribute3 *string
    // Fourth customizable extension attribute.
    extensionAttribute4 *string
    // Fifth customizable extension attribute.
    extensionAttribute5 *string
    // Sixth customizable extension attribute.
    extensionAttribute6 *string
    // Seventh customizable extension attribute.
    extensionAttribute7 *string
    // Eighth customizable extension attribute.
    extensionAttribute8 *string
    // Ninth customizable extension attribute.
    extensionAttribute9 *string
    // The OdataType property
    odataType *string
}
// NewOnPremisesExtensionAttributes instantiates a new onPremisesExtensionAttributes and sets the default values.
func NewOnPremisesExtensionAttributes()(*OnPremisesExtensionAttributes) {
    m := &OnPremisesExtensionAttributes{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateOnPremisesExtensionAttributesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnPremisesExtensionAttributesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnPremisesExtensionAttributes(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnPremisesExtensionAttributes) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExtensionAttribute1 gets the extensionAttribute1 property value. First customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute1()(*string) {
    return m.extensionAttribute1
}
// GetExtensionAttribute10 gets the extensionAttribute10 property value. Tenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute10()(*string) {
    return m.extensionAttribute10
}
// GetExtensionAttribute11 gets the extensionAttribute11 property value. Eleventh customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute11()(*string) {
    return m.extensionAttribute11
}
// GetExtensionAttribute12 gets the extensionAttribute12 property value. Twelfth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute12()(*string) {
    return m.extensionAttribute12
}
// GetExtensionAttribute13 gets the extensionAttribute13 property value. Thirteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute13()(*string) {
    return m.extensionAttribute13
}
// GetExtensionAttribute14 gets the extensionAttribute14 property value. Fourteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute14()(*string) {
    return m.extensionAttribute14
}
// GetExtensionAttribute15 gets the extensionAttribute15 property value. Fifteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute15()(*string) {
    return m.extensionAttribute15
}
// GetExtensionAttribute2 gets the extensionAttribute2 property value. Second customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute2()(*string) {
    return m.extensionAttribute2
}
// GetExtensionAttribute3 gets the extensionAttribute3 property value. Third customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute3()(*string) {
    return m.extensionAttribute3
}
// GetExtensionAttribute4 gets the extensionAttribute4 property value. Fourth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute4()(*string) {
    return m.extensionAttribute4
}
// GetExtensionAttribute5 gets the extensionAttribute5 property value. Fifth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute5()(*string) {
    return m.extensionAttribute5
}
// GetExtensionAttribute6 gets the extensionAttribute6 property value. Sixth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute6()(*string) {
    return m.extensionAttribute6
}
// GetExtensionAttribute7 gets the extensionAttribute7 property value. Seventh customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute7()(*string) {
    return m.extensionAttribute7
}
// GetExtensionAttribute8 gets the extensionAttribute8 property value. Eighth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute8()(*string) {
    return m.extensionAttribute8
}
// GetExtensionAttribute9 gets the extensionAttribute9 property value. Ninth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) GetExtensionAttribute9()(*string) {
    return m.extensionAttribute9
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnPremisesExtensionAttributes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["extensionAttribute1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute1(val)
        }
        return nil
    }
    res["extensionAttribute10"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute10(val)
        }
        return nil
    }
    res["extensionAttribute11"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute11(val)
        }
        return nil
    }
    res["extensionAttribute12"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute12(val)
        }
        return nil
    }
    res["extensionAttribute13"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute13(val)
        }
        return nil
    }
    res["extensionAttribute14"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute14(val)
        }
        return nil
    }
    res["extensionAttribute15"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute15(val)
        }
        return nil
    }
    res["extensionAttribute2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute2(val)
        }
        return nil
    }
    res["extensionAttribute3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute3(val)
        }
        return nil
    }
    res["extensionAttribute4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute4(val)
        }
        return nil
    }
    res["extensionAttribute5"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute5(val)
        }
        return nil
    }
    res["extensionAttribute6"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute6(val)
        }
        return nil
    }
    res["extensionAttribute7"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute7(val)
        }
        return nil
    }
    res["extensionAttribute8"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute8(val)
        }
        return nil
    }
    res["extensionAttribute9"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExtensionAttribute9(val)
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
func (m *OnPremisesExtensionAttributes) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *OnPremisesExtensionAttributes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("extensionAttribute1", m.GetExtensionAttribute1())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute10", m.GetExtensionAttribute10())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute11", m.GetExtensionAttribute11())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute12", m.GetExtensionAttribute12())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute13", m.GetExtensionAttribute13())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute14", m.GetExtensionAttribute14())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute15", m.GetExtensionAttribute15())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute2", m.GetExtensionAttribute2())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute3", m.GetExtensionAttribute3())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute4", m.GetExtensionAttribute4())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute5", m.GetExtensionAttribute5())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute6", m.GetExtensionAttribute6())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute7", m.GetExtensionAttribute7())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute8", m.GetExtensionAttribute8())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("extensionAttribute9", m.GetExtensionAttribute9())
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
func (m *OnPremisesExtensionAttributes) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExtensionAttribute1 sets the extensionAttribute1 property value. First customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute1(value *string)() {
    m.extensionAttribute1 = value
}
// SetExtensionAttribute10 sets the extensionAttribute10 property value. Tenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute10(value *string)() {
    m.extensionAttribute10 = value
}
// SetExtensionAttribute11 sets the extensionAttribute11 property value. Eleventh customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute11(value *string)() {
    m.extensionAttribute11 = value
}
// SetExtensionAttribute12 sets the extensionAttribute12 property value. Twelfth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute12(value *string)() {
    m.extensionAttribute12 = value
}
// SetExtensionAttribute13 sets the extensionAttribute13 property value. Thirteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute13(value *string)() {
    m.extensionAttribute13 = value
}
// SetExtensionAttribute14 sets the extensionAttribute14 property value. Fourteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute14(value *string)() {
    m.extensionAttribute14 = value
}
// SetExtensionAttribute15 sets the extensionAttribute15 property value. Fifteenth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute15(value *string)() {
    m.extensionAttribute15 = value
}
// SetExtensionAttribute2 sets the extensionAttribute2 property value. Second customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute2(value *string)() {
    m.extensionAttribute2 = value
}
// SetExtensionAttribute3 sets the extensionAttribute3 property value. Third customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute3(value *string)() {
    m.extensionAttribute3 = value
}
// SetExtensionAttribute4 sets the extensionAttribute4 property value. Fourth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute4(value *string)() {
    m.extensionAttribute4 = value
}
// SetExtensionAttribute5 sets the extensionAttribute5 property value. Fifth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute5(value *string)() {
    m.extensionAttribute5 = value
}
// SetExtensionAttribute6 sets the extensionAttribute6 property value. Sixth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute6(value *string)() {
    m.extensionAttribute6 = value
}
// SetExtensionAttribute7 sets the extensionAttribute7 property value. Seventh customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute7(value *string)() {
    m.extensionAttribute7 = value
}
// SetExtensionAttribute8 sets the extensionAttribute8 property value. Eighth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute8(value *string)() {
    m.extensionAttribute8 = value
}
// SetExtensionAttribute9 sets the extensionAttribute9 property value. Ninth customizable extension attribute.
func (m *OnPremisesExtensionAttributes) SetExtensionAttribute9(value *string)() {
    m.extensionAttribute9 = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *OnPremisesExtensionAttributes) SetOdataType(value *string)() {
    m.odataType = value
}
