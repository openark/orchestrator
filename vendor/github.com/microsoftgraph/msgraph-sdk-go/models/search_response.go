package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchResponse 
type SearchResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A collection of search results.
    hitsContainers []SearchHitsContainerable
    // The OdataType property
    odataType *string
    // Provides information related to spelling corrections in the alteration response.
    queryAlterationResponse AlterationResponseable
    // A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
    resultTemplates ResultTemplateDictionaryable
    // Contains the search terms sent in the initial search query.
    searchTerms []string
}
// NewSearchResponse instantiates a new searchResponse and sets the default values.
func NewSearchResponse()(*SearchResponse) {
    m := &SearchResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hitsContainers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSearchHitsContainerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SearchHitsContainerable, len(val))
            for i, v := range val {
                res[i] = v.(SearchHitsContainerable)
            }
            m.SetHitsContainers(res)
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
    res["queryAlterationResponse"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAlterationResponseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlterationResponse(val.(AlterationResponseable))
        }
        return nil
    }
    res["resultTemplates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResultTemplateDictionaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplates(val.(ResultTemplateDictionaryable))
        }
        return nil
    }
    res["searchTerms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSearchTerms(res)
        }
        return nil
    }
    return res
}
// GetHitsContainers gets the hitsContainers property value. A collection of search results.
func (m *SearchResponse) GetHitsContainers()([]SearchHitsContainerable) {
    return m.hitsContainers
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchResponse) GetOdataType()(*string) {
    return m.odataType
}
// GetQueryAlterationResponse gets the queryAlterationResponse property value. Provides information related to spelling corrections in the alteration response.
func (m *SearchResponse) GetQueryAlterationResponse()(AlterationResponseable) {
    return m.queryAlterationResponse
}
// GetResultTemplates gets the resultTemplates property value. A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
func (m *SearchResponse) GetResultTemplates()(ResultTemplateDictionaryable) {
    return m.resultTemplates
}
// GetSearchTerms gets the searchTerms property value. Contains the search terms sent in the initial search query.
func (m *SearchResponse) GetSearchTerms()([]string) {
    return m.searchTerms
}
// Serialize serializes information the current object
func (m *SearchResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHitsContainers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHitsContainers()))
        for i, v := range m.GetHitsContainers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("hitsContainers", cast)
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
        err := writer.WriteObjectValue("queryAlterationResponse", m.GetQueryAlterationResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resultTemplates", m.GetResultTemplates())
        if err != nil {
            return err
        }
    }
    if m.GetSearchTerms() != nil {
        err := writer.WriteCollectionOfStringValues("searchTerms", m.GetSearchTerms())
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
func (m *SearchResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHitsContainers sets the hitsContainers property value. A collection of search results.
func (m *SearchResponse) SetHitsContainers(value []SearchHitsContainerable)() {
    m.hitsContainers = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchResponse) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQueryAlterationResponse sets the queryAlterationResponse property value. Provides information related to spelling corrections in the alteration response.
func (m *SearchResponse) SetQueryAlterationResponse(value AlterationResponseable)() {
    m.queryAlterationResponse = value
}
// SetResultTemplates sets the resultTemplates property value. A dictionary of resultTemplateIds and associated values, which include the name and JSON schema of the result templates.
func (m *SearchResponse) SetResultTemplates(value ResultTemplateDictionaryable)() {
    m.resultTemplates = value
}
// SetSearchTerms sets the searchTerms property value. Contains the search terms sent in the initial search query.
func (m *SearchResponse) SetSearchTerms(value []string)() {
    m.searchTerms = value
}
