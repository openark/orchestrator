package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchHitsContainer 
type SearchHitsContainer struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The aggregations property
    aggregations []SearchAggregationable
    // A collection of the search results.
    hits []SearchHitable
    // Provides information if more results are available. Based on this information, you can adjust the from and size properties of the searchRequest accordingly.
    moreResultsAvailable *bool
    // The OdataType property
    odataType *string
    // The total number of results. Note this is not the number of results on the page, but the total number of results satisfying the query.
    total *int32
}
// NewSearchHitsContainer instantiates a new searchHitsContainer and sets the default values.
func NewSearchHitsContainer()(*SearchHitsContainer) {
    m := &SearchHitsContainer{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchHitsContainerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchHitsContainerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchHitsContainer(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHitsContainer) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAggregations gets the aggregations property value. The aggregations property
func (m *SearchHitsContainer) GetAggregations()([]SearchAggregationable) {
    return m.aggregations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchHitsContainer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSearchAggregationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchAggregationable, len(val))
            for i, v := range val {
                res[i] = v.(SearchAggregationable)
            }
            m.SetAggregations(res)
        }
        return nil
    }
    res["hits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSearchHitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchHitable, len(val))
            for i, v := range val {
                res[i] = v.(SearchHitable)
            }
            m.SetHits(res)
        }
        return nil
    }
    res["moreResultsAvailable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMoreResultsAvailable(val)
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
    res["total"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotal(val)
        }
        return nil
    }
    return res
}
// GetHits gets the hits property value. A collection of the search results.
func (m *SearchHitsContainer) GetHits()([]SearchHitable) {
    return m.hits
}
// GetMoreResultsAvailable gets the moreResultsAvailable property value. Provides information if more results are available. Based on this information, you can adjust the from and size properties of the searchRequest accordingly.
func (m *SearchHitsContainer) GetMoreResultsAvailable()(*bool) {
    return m.moreResultsAvailable
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchHitsContainer) GetOdataType()(*string) {
    return m.odataType
}
// GetTotal gets the total property value. The total number of results. Note this is not the number of results on the page, but the total number of results satisfying the query.
func (m *SearchHitsContainer) GetTotal()(*int32) {
    return m.total
}
// Serialize serializes information the current object
func (m *SearchHitsContainer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAggregations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAggregations()))
        for i, v := range m.GetAggregations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("aggregations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetHits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHits()))
        for i, v := range m.GetHits() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("hits", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("moreResultsAvailable", m.GetMoreResultsAvailable())
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
        err := writer.WriteInt32Value("total", m.GetTotal())
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
func (m *SearchHitsContainer) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAggregations sets the aggregations property value. The aggregations property
func (m *SearchHitsContainer) SetAggregations(value []SearchAggregationable)() {
    m.aggregations = value
}
// SetHits sets the hits property value. A collection of the search results.
func (m *SearchHitsContainer) SetHits(value []SearchHitable)() {
    m.hits = value
}
// SetMoreResultsAvailable sets the moreResultsAvailable property value. Provides information if more results are available. Based on this information, you can adjust the from and size properties of the searchRequest accordingly.
func (m *SearchHitsContainer) SetMoreResultsAvailable(value *bool)() {
    m.moreResultsAvailable = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchHitsContainer) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTotal sets the total property value. The total number of results. Note this is not the number of results on the page, but the total number of results satisfying the query.
func (m *SearchHitsContainer) SetTotal(value *int32)() {
    m.total = value
}
