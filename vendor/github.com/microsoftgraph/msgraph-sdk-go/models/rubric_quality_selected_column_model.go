package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RubricQualitySelectedColumnModel 
type RubricQualitySelectedColumnModel struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // ID of the selected level for this quality.
    columnId *string
    // The OdataType property
    odataType *string
    // ID of the associated quality.
    qualityId *string
}
// NewRubricQualitySelectedColumnModel instantiates a new rubricQualitySelectedColumnModel and sets the default values.
func NewRubricQualitySelectedColumnModel()(*RubricQualitySelectedColumnModel) {
    m := &RubricQualitySelectedColumnModel{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRubricQualitySelectedColumnModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRubricQualitySelectedColumnModelFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRubricQualitySelectedColumnModel(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricQualitySelectedColumnModel) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetColumnId gets the columnId property value. ID of the selected level for this quality.
func (m *RubricQualitySelectedColumnModel) GetColumnId()(*string) {
    return m.columnId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RubricQualitySelectedColumnModel) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["columnId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColumnId(val)
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
    res["qualityId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQualityId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RubricQualitySelectedColumnModel) GetOdataType()(*string) {
    return m.odataType
}
// GetQualityId gets the qualityId property value. ID of the associated quality.
func (m *RubricQualitySelectedColumnModel) GetQualityId()(*string) {
    return m.qualityId
}
// Serialize serializes information the current object
func (m *RubricQualitySelectedColumnModel) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("columnId", m.GetColumnId())
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
        err := writer.WriteStringValue("qualityId", m.GetQualityId())
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
func (m *RubricQualitySelectedColumnModel) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetColumnId sets the columnId property value. ID of the selected level for this quality.
func (m *RubricQualitySelectedColumnModel) SetColumnId(value *string)() {
    m.columnId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RubricQualitySelectedColumnModel) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQualityId sets the qualityId property value. ID of the associated quality.
func (m *RubricQualitySelectedColumnModel) SetQualityId(value *string)() {
    m.qualityId = value
}
