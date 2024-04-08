package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RubricQuality 
type RubricQuality struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The collection of criteria for this rubric quality.
    criteria []RubricCriterionable
    // The description of this rubric quality.
    description EducationItemBodyable
    // The name of this rubric quality.
    displayName *string
    // The OdataType property
    odataType *string
    // The ID of this resource.
    qualityId *string
    // If present, a numerical weight for this quality.  Weights must add up to 100.
    weight *float32
}
// NewRubricQuality instantiates a new rubricQuality and sets the default values.
func NewRubricQuality()(*RubricQuality) {
    m := &RubricQuality{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRubricQualityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRubricQualityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRubricQuality(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RubricQuality) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCriteria gets the criteria property value. The collection of criteria for this rubric quality.
func (m *RubricQuality) GetCriteria()([]RubricCriterionable) {
    return m.criteria
}
// GetDescription gets the description property value. The description of this rubric quality.
func (m *RubricQuality) GetDescription()(EducationItemBodyable) {
    return m.description
}
// GetDisplayName gets the displayName property value. The name of this rubric quality.
func (m *RubricQuality) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RubricQuality) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["criteria"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRubricCriterionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RubricCriterionable, len(val))
            for i, v := range val {
                res[i] = v.(RubricCriterionable)
            }
            m.SetCriteria(res)
        }
        return nil
    }
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["weight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RubricQuality) GetOdataType()(*string) {
    return m.odataType
}
// GetQualityId gets the qualityId property value. The ID of this resource.
func (m *RubricQuality) GetQualityId()(*string) {
    return m.qualityId
}
// GetWeight gets the weight property value. If present, a numerical weight for this quality.  Weights must add up to 100.
func (m *RubricQuality) GetWeight()(*float32) {
    return m.weight
}
// Serialize serializes information the current object
func (m *RubricQuality) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCriteria() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCriteria()))
        for i, v := range m.GetCriteria() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("criteria", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
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
        err := writer.WriteFloat32Value("weight", m.GetWeight())
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
func (m *RubricQuality) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCriteria sets the criteria property value. The collection of criteria for this rubric quality.
func (m *RubricQuality) SetCriteria(value []RubricCriterionable)() {
    m.criteria = value
}
// SetDescription sets the description property value. The description of this rubric quality.
func (m *RubricQuality) SetDescription(value EducationItemBodyable)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The name of this rubric quality.
func (m *RubricQuality) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RubricQuality) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQualityId sets the qualityId property value. The ID of this resource.
func (m *RubricQuality) SetQualityId(value *string)() {
    m.qualityId = value
}
// SetWeight sets the weight property value. If present, a numerical weight for this quality.  Weights must add up to 100.
func (m *RubricQuality) SetWeight(value *float32)() {
    m.weight = value
}
