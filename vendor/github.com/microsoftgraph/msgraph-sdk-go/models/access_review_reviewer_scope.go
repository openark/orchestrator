package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessReviewReviewerScope 
type AccessReviewReviewerScope struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // The query specifying who will be the reviewer.
    query *string
    // In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
    queryRoot *string
    // The type of query. Examples include MicrosoftGraph and ARM.
    queryType *string
}
// NewAccessReviewReviewerScope instantiates a new accessReviewReviewerScope and sets the default values.
func NewAccessReviewReviewerScope()(*AccessReviewReviewerScope) {
    m := &AccessReviewReviewerScope{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAccessReviewReviewerScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessReviewReviewerScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessReviewReviewerScope(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AccessReviewReviewerScope) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessReviewReviewerScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuery(val)
        }
        return nil
    }
    res["queryRoot"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryRoot(val)
        }
        return nil
    }
    res["queryType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryType(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AccessReviewReviewerScope) GetOdataType()(*string) {
    return m.odataType
}
// GetQuery gets the query property value. The query specifying who will be the reviewer.
func (m *AccessReviewReviewerScope) GetQuery()(*string) {
    return m.query
}
// GetQueryRoot gets the queryRoot property value. In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
func (m *AccessReviewReviewerScope) GetQueryRoot()(*string) {
    return m.queryRoot
}
// GetQueryType gets the queryType property value. The type of query. Examples include MicrosoftGraph and ARM.
func (m *AccessReviewReviewerScope) GetQueryType()(*string) {
    return m.queryType
}
// Serialize serializes information the current object
func (m *AccessReviewReviewerScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("query", m.GetQuery())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryRoot", m.GetQueryRoot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("queryType", m.GetQueryType())
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
func (m *AccessReviewReviewerScope) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AccessReviewReviewerScope) SetOdataType(value *string)() {
    m.odataType = value
}
// SetQuery sets the query property value. The query specifying who will be the reviewer.
func (m *AccessReviewReviewerScope) SetQuery(value *string)() {
    m.query = value
}
// SetQueryRoot sets the queryRoot property value. In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions.
func (m *AccessReviewReviewerScope) SetQueryRoot(value *string)() {
    m.queryRoot = value
}
// SetQueryType sets the queryType property value. The type of query. Examples include MicrosoftGraph and ARM.
func (m *AccessReviewReviewerScope) SetQueryType(value *string)() {
    m.queryType = value
}
