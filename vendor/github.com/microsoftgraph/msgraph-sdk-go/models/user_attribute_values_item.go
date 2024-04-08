package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAttributeValuesItem 
type UserAttributeValuesItem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Determines whether the value is set as the default.
    isDefault *bool
    // The display name of the property displayed to the user in the user flow.
    name *string
    // The OdataType property
    odataType *string
    // The value that is set when this item is selected.
    value *string
}
// NewUserAttributeValuesItem instantiates a new userAttributeValuesItem and sets the default values.
func NewUserAttributeValuesItem()(*UserAttributeValuesItem) {
    m := &UserAttributeValuesItem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserAttributeValuesItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserAttributeValuesItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserAttributeValuesItem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAttributeValuesItem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAttributeValuesItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
        }
        return nil
    }
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
// GetIsDefault gets the isDefault property value. Determines whether the value is set as the default.
func (m *UserAttributeValuesItem) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetName gets the name property value. The display name of the property displayed to the user in the user flow.
func (m *UserAttributeValuesItem) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserAttributeValuesItem) GetOdataType()(*string) {
    return m.odataType
}
// GetValue gets the value property value. The value that is set when this item is selected.
func (m *UserAttributeValuesItem) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *UserAttributeValuesItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDefault", m.GetIsDefault())
        if err != nil {
            return err
        }
    }
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
func (m *UserAttributeValuesItem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsDefault sets the isDefault property value. Determines whether the value is set as the default.
func (m *UserAttributeValuesItem) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetName sets the name property value. The display name of the property displayed to the user in the user flow.
func (m *UserAttributeValuesItem) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserAttributeValuesItem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetValue sets the value property value. The value that is set when this item is selected.
func (m *UserAttributeValuesItem) SetValue(value *string)() {
    m.value = value
}
