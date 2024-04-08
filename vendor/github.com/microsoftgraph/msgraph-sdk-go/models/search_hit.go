package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchHit 
type SearchHit struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The name of the content source that the externalItem is part of.
    contentSource *string
    // The internal identifier for the item. The format of the identifier varies based on the entity type. For details, see hitId format.
    hitId *string
    // The OdataType property
    odataType *string
    // The rank or the order of the result.
    rank *int32
    // The resource property
    resource Entityable
    // ID of the result template used to render the search result. This ID must map to a display layout in the resultTemplates dictionary that is also included in the searchResponse.
    resultTemplateId *string
    // A summary of the result, if a summary is available.
    summary *string
}
// NewSearchHit instantiates a new searchHit and sets the default values.
func NewSearchHit()(*SearchHit) {
    m := &SearchHit{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSearchHitFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchHitFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchHit(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchHit) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetContentSource gets the contentSource property value. The name of the content source that the externalItem is part of.
func (m *SearchHit) GetContentSource()(*string) {
    return m.contentSource
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchHit) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["contentSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentSource(val)
        }
        return nil
    }
    res["hitId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHitId(val)
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
    res["rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["resource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResource(val.(Entityable))
        }
        return nil
    }
    res["resultTemplateId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResultTemplateId(val)
        }
        return nil
    }
    res["summary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    return res
}
// GetHitId gets the hitId property value. The internal identifier for the item. The format of the identifier varies based on the entity type. For details, see hitId format.
func (m *SearchHit) GetHitId()(*string) {
    return m.hitId
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SearchHit) GetOdataType()(*string) {
    return m.odataType
}
// GetRank gets the rank property value. The rank or the order of the result.
func (m *SearchHit) GetRank()(*int32) {
    return m.rank
}
// GetResource gets the resource property value. The resource property
func (m *SearchHit) GetResource()(Entityable) {
    return m.resource
}
// GetResultTemplateId gets the resultTemplateId property value. ID of the result template used to render the search result. This ID must map to a display layout in the resultTemplates dictionary that is also included in the searchResponse.
func (m *SearchHit) GetResultTemplateId()(*string) {
    return m.resultTemplateId
}
// GetSummary gets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) GetSummary()(*string) {
    return m.summary
}
// Serialize serializes information the current object
func (m *SearchHit) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("contentSource", m.GetContentSource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hitId", m.GetHitId())
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
        err := writer.WriteInt32Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("resultTemplateId", m.GetResultTemplateId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("summary", m.GetSummary())
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
func (m *SearchHit) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetContentSource sets the contentSource property value. The name of the content source that the externalItem is part of.
func (m *SearchHit) SetContentSource(value *string)() {
    m.contentSource = value
}
// SetHitId sets the hitId property value. The internal identifier for the item. The format of the identifier varies based on the entity type. For details, see hitId format.
func (m *SearchHit) SetHitId(value *string)() {
    m.hitId = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SearchHit) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRank sets the rank property value. The rank or the order of the result.
func (m *SearchHit) SetRank(value *int32)() {
    m.rank = value
}
// SetResource sets the resource property value. The resource property
func (m *SearchHit) SetResource(value Entityable)() {
    m.resource = value
}
// SetResultTemplateId sets the resultTemplateId property value. ID of the result template used to render the search result. This ID must map to a display layout in the resultTemplates dictionary that is also included in the searchResponse.
func (m *SearchHit) SetResultTemplateId(value *string)() {
    m.resultTemplateId = value
}
// SetSummary sets the summary property value. A summary of the result, if a summary is available.
func (m *SearchHit) SetSummary(value *string)() {
    m.summary = value
}
