package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AggregationOption 
type AggregationOption struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The bucketDefinition property
    bucketDefinition BucketAggregationDefinitionable
    // Computes aggregation on the field while the field exists in current entity type. Required.
    field *string
    // The OdataType property
    odataType *string
    // The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
    size *int32
}
// NewAggregationOption instantiates a new aggregationOption and sets the default values.
func NewAggregationOption()(*AggregationOption) {
    m := &AggregationOption{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAggregationOptionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAggregationOptionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAggregationOption(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AggregationOption) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetBucketDefinition gets the bucketDefinition property value. The bucketDefinition property
func (m *AggregationOption) GetBucketDefinition()(BucketAggregationDefinitionable) {
    return m.bucketDefinition
}
// GetField gets the field property value. Computes aggregation on the field while the field exists in current entity type. Required.
func (m *AggregationOption) GetField()(*string) {
    return m.field
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AggregationOption) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["bucketDefinition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateBucketAggregationDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBucketDefinition(val.(BucketAggregationDefinitionable))
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
    res["size"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSize(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AggregationOption) GetOdataType()(*string) {
    return m.odataType
}
// GetSize gets the size property value. The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
func (m *AggregationOption) GetSize()(*int32) {
    return m.size
}
// Serialize serializes information the current object
func (m *AggregationOption) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("bucketDefinition", m.GetBucketDefinition())
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
        err := writer.WriteInt32Value("size", m.GetSize())
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
func (m *AggregationOption) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetBucketDefinition sets the bucketDefinition property value. The bucketDefinition property
func (m *AggregationOption) SetBucketDefinition(value BucketAggregationDefinitionable)() {
    m.bucketDefinition = value
}
// SetField sets the field property value. Computes aggregation on the field while the field exists in current entity type. Required.
func (m *AggregationOption) SetField(value *string)() {
    m.field = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AggregationOption) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSize sets the size property value. The number of searchBucket resources to be returned. This is not required when the range is provided manually in the search request. Optional.
func (m *AggregationOption) SetSize(value *int32)() {
    m.size = value
}
