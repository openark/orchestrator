package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SortProperty 
type SortProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // True if the sort order is descending. Default is false, with the sort order as ascending. Optional.
    isDescending *bool
    // The name of the property to sort on. Required.
    name *string
    // The OdataType property
    odataType *string
}
// NewSortProperty instantiates a new sortProperty and sets the default values.
func NewSortProperty()(*SortProperty) {
    m := &SortProperty{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSortPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSortPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSortProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SortProperty) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SortProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["isDescending"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDescending(val)
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
    return res
}
// GetIsDescending gets the isDescending property value. True if the sort order is descending. Default is false, with the sort order as ascending. Optional.
func (m *SortProperty) GetIsDescending()(*bool) {
    return m.isDescending
}
// GetName gets the name property value. The name of the property to sort on. Required.
func (m *SortProperty) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SortProperty) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SortProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDescending", m.GetIsDescending())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SortProperty) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetIsDescending sets the isDescending property value. True if the sort order is descending. Default is false, with the sort order as ascending. Optional.
func (m *SortProperty) SetIsDescending(value *bool)() {
    m.isDescending = value
}
// SetName sets the name property value. The name of the property to sort on. Required.
func (m *SortProperty) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SortProperty) SetOdataType(value *string)() {
    m.odataType = value
}
