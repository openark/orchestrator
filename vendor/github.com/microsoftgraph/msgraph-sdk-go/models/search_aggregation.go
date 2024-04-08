package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchAggregation 
type SearchAggregation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The buckets property
    buckets []SearchBucketable
    // The field property
    field *string
    // The OdataType property
    odataType *string
}
// NewSearchAggregation instantiates a new searchAggregation and sets the default values.
func NewSearchAggregation()(*SearchAggregation) {
    m := &SearchAggregation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchAggregationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchAggregationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchAggregation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAggregation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBuckets gets the buckets property value. The buckets property
func (m *SearchAggregation) GetBuckets()([]SearchBucketable) {
    return m.buckets
}
// GetField gets the field property value. The field property
func (m *SearchAggregation) GetField()(*string) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchAggregation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["buckets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSearchBucketFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchBucketable, len(val))
            for i, v := range val {
                res[i] = v.(SearchBucketable)
            }
            m.SetBuckets(res)
        }
        return nil
    }
    res["field"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetField(val)
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
func (m *SearchAggregation) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SearchAggregation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetBuckets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBuckets()))
        for i, v := range m.GetBuckets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("buckets", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("field", m.GetField())
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
func (m *SearchAggregation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBuckets sets the buckets property value. The buckets property
func (m *SearchAggregation) SetBuckets(value []SearchBucketable)() {
    m.buckets = value
}
// SetField sets the field property value. The field property
func (m *SearchAggregation) SetField(value *string)() {
    m.field = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchAggregation) SetOdataType(value *string)() {
    m.odataType = value
}
