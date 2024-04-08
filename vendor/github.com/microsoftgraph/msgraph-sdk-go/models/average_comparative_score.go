package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AverageComparativeScore 
type AverageComparativeScore struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Average score within specified basis.
    averageScore *float64
    // Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
    basis *string
    // The OdataType property
    odataType *string
}
// NewAverageComparativeScore instantiates a new averageComparativeScore and sets the default values.
func NewAverageComparativeScore()(*AverageComparativeScore) {
    m := &AverageComparativeScore{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAverageComparativeScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAverageComparativeScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAverageComparativeScore(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AverageComparativeScore) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAverageScore gets the averageScore property value. Average score within specified basis.
func (m *AverageComparativeScore) GetAverageScore()(*float64) {
    return m.averageScore
}
// GetBasis gets the basis property value. Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
func (m *AverageComparativeScore) GetBasis()(*string) {
    return m.basis
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AverageComparativeScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["averageScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAverageScore(val)
        }
        return nil
    }
    res["basis"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasis(val)
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
func (m *AverageComparativeScore) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *AverageComparativeScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteFloat64Value("averageScore", m.GetAverageScore())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("basis", m.GetBasis())
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
func (m *AverageComparativeScore) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAverageScore sets the averageScore property value. Average score within specified basis.
func (m *AverageComparativeScore) SetAverageScore(value *float64)() {
    m.averageScore = value
}
// SetBasis sets the basis property value. Scope type. The possible values are: AllTenants, TotalSeats, IndustryTypes.
func (m *AverageComparativeScore) SetBasis(value *string)() {
    m.basis = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AverageComparativeScore) SetOdataType(value *string)() {
    m.odataType = value
}
