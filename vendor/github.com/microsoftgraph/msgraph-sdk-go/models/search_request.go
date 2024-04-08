package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchRequest 
type SearchRequest struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The aggregationFilters property
    aggregationFilters []string
    // The aggregations property
    aggregations []AggregationOptionable
    // The contentSources property
    contentSources []string
    // The enableTopResults property
    enableTopResults *bool
    // The entityTypes property
    entityTypes []EntityType
    // The fields property
    fields []string
    // The from property
    from *int32
    // The OdataType property
    odataType *string
    // The query property
    query SearchQueryable
    // The queryAlterationOptions property
    queryAlterationOptions SearchAlterationOptionsable
    // The region property
    region *string
    // The resultTemplateOptions property
    resultTemplateOptions ResultTemplateOptionable
    // The sharePointOneDriveOptions property
    sharePointOneDriveOptions SharePointOneDriveOptionsable
    // The size property
    size *int32
    // The sortProperties property
    sortProperties []SortPropertyable
}
// NewSearchRequest instantiates a new searchRequest and sets the default values.
func NewSearchRequest()(*SearchRequest) {
    m := &SearchRequest{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchRequest(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchRequest) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAggregationFilters gets the aggregationFilters property value. The aggregationFilters property
func (m *SearchRequest) GetAggregationFilters()([]string) {
    return m.aggregationFilters
}
// GetAggregations gets the aggregations property value. The aggregations property
func (m *SearchRequest) GetAggregations()([]AggregationOptionable) {
    return m.aggregations
}
// GetContentSources gets the contentSources property value. The contentSources property
func (m *SearchRequest) GetContentSources()([]string) {
    return m.contentSources
}
// GetEnableTopResults gets the enableTopResults property value. The enableTopResults property
func (m *SearchRequest) GetEnableTopResults()(*bool) {
    return m.enableTopResults
}
// GetEntityTypes gets the entityTypes property value. The entityTypes property
func (m *SearchRequest) GetEntityTypes()([]EntityType) {
    return m.entityTypes
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aggregationFilters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAggregationFilters(res)
        }
        return nil
    }
    res["aggregations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAggregationOptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AggregationOptionable, len(val))
            for i, v := range val {
                res[i] = v.(AggregationOptionable)
            }
            m.SetAggregations(res)
        }
        return nil
    }
    res["contentSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetContentSources(res)
        }
        return nil
    }
    res["enableTopResults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableTopResults(val)
        }
        return nil
    }
    res["entityTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseEntityType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EntityType, len(val))
            for i, v := range val {
                res[i] = *(v.(*EntityType))
            }
            m.SetEntityTypes(res)
        }
        return nil
    }
    res["fields"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetFields(res)
        }
        return nil
    }
    res["from"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
    res["query"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchQueryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val.(SearchQueryable))
        }
        return nil
    }
    res["queryAlterationOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSearchAlterationOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlterationOptions(val.(SearchAlterationOptionsable))
        }
        return nil
    }
    res["region"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRegion(val)
        }
        return nil
    }
    res["resultTemplateOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResultTemplateOptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplateOptions(val.(ResultTemplateOptionable))
        }
        return nil
    }
    res["sharePointOneDriveOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointOneDriveOptionsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointOneDriveOptions(val.(SharePointOneDriveOptionsable))
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
    res["sortProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSortPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SortPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(SortPropertyable)
            }
            m.SetSortProperties(res)
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. The fields property
func (m *SearchRequest) GetFields()([]string) {
    return m.fields
}
// GetFrom gets the from property value. The from property
func (m *SearchRequest) GetFrom()(*int32) {
    return m.from
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchRequest) GetOdataType()(*string) {
    return m.odataType
}
// GetQuery gets the query property value. The query property
func (m *SearchRequest) GetQuery()(SearchQueryable) {
    return m.query
}
// GetQueryAlterationOptions gets the queryAlterationOptions property value. The queryAlterationOptions property
func (m *SearchRequest) GetQueryAlterationOptions()(SearchAlterationOptionsable) {
    return m.queryAlterationOptions
}
// GetRegion gets the region property value. The region property
func (m *SearchRequest) GetRegion()(*string) {
    return m.region
}
// GetResultTemplateOptions gets the resultTemplateOptions property value. The resultTemplateOptions property
func (m *SearchRequest) GetResultTemplateOptions()(ResultTemplateOptionable) {
    return m.resultTemplateOptions
}
// GetSharePointOneDriveOptions gets the sharePointOneDriveOptions property value. The sharePointOneDriveOptions property
func (m *SearchRequest) GetSharePointOneDriveOptions()(SharePointOneDriveOptionsable) {
    return m.sharePointOneDriveOptions
}
// GetSize gets the size property value. The size property
func (m *SearchRequest) GetSize()(*int32) {
    return m.size
}
// GetSortProperties gets the sortProperties property value. The sortProperties property
func (m *SearchRequest) GetSortProperties()([]SortPropertyable) {
    return m.sortProperties
}
// Serialize serializes information the current object
func (m *SearchRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAggregationFilters() != nil {
        err := writer.WriteCollectionOfStringValues("aggregationFilters", m.GetAggregationFilters())
        if err != nil {
            return err
        }
    }
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
    if m.GetContentSources() != nil {
        err := writer.WriteCollectionOfStringValues("contentSources", m.GetContentSources())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableTopResults", m.GetEnableTopResults())
        if err != nil {
            return err
        }
    }
    if m.GetEntityTypes() != nil {
        err := writer.WriteCollectionOfStringValues("entityTypes", SerializeEntityType(m.GetEntityTypes()))
        if err != nil {
            return err
        }
    }
    if m.GetFields() != nil {
        err := writer.WriteCollectionOfStringValues("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("from", m.GetFrom())
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
        err := writer.WriteObjectValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("queryAlterationOptions", m.GetQueryAlterationOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("region", m.GetRegion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resultTemplateOptions", m.GetResultTemplateOptions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("sharePointOneDriveOptions", m.GetSharePointOneDriveOptions())
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
    if m.GetSortProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSortProperties()))
        for i, v := range m.GetSortProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("sortProperties", cast)
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
func (m *SearchRequest) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAggregationFilters sets the aggregationFilters property value. The aggregationFilters property
func (m *SearchRequest) SetAggregationFilters(value []string)() {
    m.aggregationFilters = value
}
// SetAggregations sets the aggregations property value. The aggregations property
func (m *SearchRequest) SetAggregations(value []AggregationOptionable)() {
    m.aggregations = value
}
// SetContentSources sets the contentSources property value. The contentSources property
func (m *SearchRequest) SetContentSources(value []string)() {
    m.contentSources = value
}
// SetEnableTopResults sets the enableTopResults property value. The enableTopResults property
func (m *SearchRequest) SetEnableTopResults(value *bool)() {
    m.enableTopResults = value
}
// SetEntityTypes sets the entityTypes property value. The entityTypes property
func (m *SearchRequest) SetEntityTypes(value []EntityType)() {
    m.entityTypes = value
}
// SetFields sets the fields property value. The fields property
func (m *SearchRequest) SetFields(value []string)() {
    m.fields = value
}
// SetFrom sets the from property value. The from property
func (m *SearchRequest) SetFrom(value *int32)() {
    m.from = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchRequest) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuery sets the query property value. The query property
func (m *SearchRequest) SetQuery(value SearchQueryable)() {
    m.query = value
}
// SetQueryAlterationOptions sets the queryAlterationOptions property value. The queryAlterationOptions property
func (m *SearchRequest) SetQueryAlterationOptions(value SearchAlterationOptionsable)() {
    m.queryAlterationOptions = value
}
// SetRegion sets the region property value. The region property
func (m *SearchRequest) SetRegion(value *string)() {
    m.region = value
}
// SetResultTemplateOptions sets the resultTemplateOptions property value. The resultTemplateOptions property
func (m *SearchRequest) SetResultTemplateOptions(value ResultTemplateOptionable)() {
    m.resultTemplateOptions = value
}
// SetSharePointOneDriveOptions sets the sharePointOneDriveOptions property value. The sharePointOneDriveOptions property
func (m *SearchRequest) SetSharePointOneDriveOptions(value SharePointOneDriveOptionsable)() {
    m.sharePointOneDriveOptions = value
}
// SetSize sets the size property value. The size property
func (m *SearchRequest) SetSize(value *int32)() {
    m.size = value
}
// SetSortProperties sets the sortProperties property value. The sortProperties property
func (m *SearchRequest) SetSortProperties(value []SortPropertyable)() {
    m.sortProperties = value
}
