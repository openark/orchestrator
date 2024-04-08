package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BucketAggregationRange 
type BucketAggregationRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Defines the lower bound from which to compute the aggregation. This can be a numeric value or a string representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required.
    from *string
    // The OdataType property
    odataType *string
    // Defines the upper bound up to which to compute the aggregation. This can be a numeric value or a string representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required.
    to *string
}
// NewBucketAggregationRange instantiates a new bucketAggregationRange and sets the default values.
func NewBucketAggregationRange()(*BucketAggregationRange) {
    m := &BucketAggregationRange{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateBucketAggregationRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBucketAggregationRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewBucketAggregationRange(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BucketAggregationRange) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BucketAggregationRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["from"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val)
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
    res["to"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTo(val)
        }
        return nil
    }
    return res
}
// GetFrom gets the from property value. Defines the lower bound from which to compute the aggregation. This can be a numeric value or a string representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required.
func (m *BucketAggregationRange) GetFrom()(*string) {
    return m.from
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *BucketAggregationRange) GetOdataType()(*string) {
    return m.odataType
}
// GetTo gets the to property value. Defines the upper bound up to which to compute the aggregation. This can be a numeric value or a string representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required.
func (m *BucketAggregationRange) GetTo()(*string) {
    return m.to
}
// Serialize serializes information the current object
func (m *BucketAggregationRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("from", m.GetFrom())
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
        err := writer.WriteStringValue("to", m.GetTo())
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
func (m *BucketAggregationRange) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetFrom sets the from property value. Defines the lower bound from which to compute the aggregation. This can be a numeric value or a string representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required.
func (m *BucketAggregationRange) SetFrom(value *string)() {
    m.from = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *BucketAggregationRange) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTo sets the to property value. Defines the upper bound up to which to compute the aggregation. This can be a numeric value or a string representation of a date using the YYYY-MM-DDTHH:mm:ss.sssZ format. Required.
func (m *BucketAggregationRange) SetTo(value *string)() {
    m.to = value
}
