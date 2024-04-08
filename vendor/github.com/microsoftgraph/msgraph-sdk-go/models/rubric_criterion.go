package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RubricCriterion 
type RubricCriterion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The description of this criterion.
    description EducationItemBodyable
    // The OdataType property
    odataType *string
}
// NewRubricCriterion instantiates a new rubricCriterion and sets the default values.
func NewRubricCriterion()(*RubricCriterion) {
    m := &RubricCriterion{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRubricCriterionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRubricCriterionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRubricCriterion(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricCriterion) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDescription gets the description property value. The description of this criterion.
func (m *RubricCriterion) GetDescription()(EducationItemBodyable) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RubricCriterion) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val.(EducationItemBodyable))
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
func (m *RubricCriterion) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *RubricCriterion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("description", m.GetDescription())
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
func (m *RubricCriterion) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDescription sets the description property value. The description of this criterion.
func (m *RubricCriterion) SetDescription(value EducationItemBodyable)() {
    m.description = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RubricCriterion) SetOdataType(value *string)() {
    m.odataType = value
}
