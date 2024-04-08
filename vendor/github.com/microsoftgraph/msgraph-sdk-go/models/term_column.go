package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TermColumn 
type TermColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Specifies whether the column will allow more than one value.
    allowMultipleValues *bool
    // The OdataType property
    odataType *string
    // Specifies whether to display the entire term path or only the term label.
    showFullyQualifiedName *bool
}
// NewTermColumn instantiates a new termColumn and sets the default values.
func NewTermColumn()(*TermColumn) {
    m := &TermColumn{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateTermColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTermColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTermColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TermColumn) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAllowMultipleValues gets the allowMultipleValues property value. Specifies whether the column will allow more than one value.
func (m *TermColumn) GetAllowMultipleValues()(*bool) {
    return m.allowMultipleValues
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["allowMultipleValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleValues(val)
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
    res["showFullyQualifiedName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowFullyQualifiedName(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *TermColumn) GetOdataType()(*string) {
    return m.odataType
}
// GetShowFullyQualifiedName gets the showFullyQualifiedName property value. Specifies whether to display the entire term path or only the term label.
func (m *TermColumn) GetShowFullyQualifiedName()(*bool) {
    return m.showFullyQualifiedName
}
// Serialize serializes information the current object
func (m *TermColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleValues", m.GetAllowMultipleValues())
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
        err := writer.WriteBoolValue("showFullyQualifiedName", m.GetShowFullyQualifiedName())
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
func (m *TermColumn) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAllowMultipleValues sets the allowMultipleValues property value. Specifies whether the column will allow more than one value.
func (m *TermColumn) SetAllowMultipleValues(value *bool)() {
    m.allowMultipleValues = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *TermColumn) SetOdataType(value *string)() {
    m.odataType = value
}
// SetShowFullyQualifiedName sets the showFullyQualifiedName property value. Specifies whether to display the entire term path or only the term label.
func (m *TermColumn) SetShowFullyQualifiedName(value *bool)() {
    m.showFullyQualifiedName = value
}
