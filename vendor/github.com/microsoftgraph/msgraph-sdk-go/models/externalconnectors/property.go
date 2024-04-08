package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Property 
type Property struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A set of aliases or a friendly names for the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^. Optional.
    aliases []string
    // Specifies if the property is queryable. Queryable properties can be used in Keyword Query Language (KQL) queries. Optional.
    isQueryable *bool
    // Specifies if the property is refinable.  Refinable properties can be used to filter search results in the Search API and add a refiner control in the Microsoft Search user experience. Optional.
    isRefinable *bool
    // Specifies if the property is retrievable. Retrievable properties are returned in the result set when items are returned by the search API. Retrievable properties are also available to add to the display template used to render search results. Optional.
    isRetrievable *bool
    // Specifies if the property is searchable. Only properties of type String or StringCollection can be searchable. Non-searchable properties are not added to the search index. Optional.
    isSearchable *bool
    // Specifies one or more well-known tags added against a property. Labels help Microsoft Search understand the semantics of the data in the connection. Adding appropriate labels would result in an enhanced search experience (e.g. better relevance). The possible values are: title, url, createdBy, lastModifiedBy, authors, createdDateTime, lastModifiedDateTime, fileName, fileExtension, unknownFutureValue. Optional.
    labels []Label
    // The name of the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^.  Required.
    name *string
    // The OdataType property
    odataType *string
    // The type property
    typeEscaped *PropertyType
}
// NewProperty instantiates a new property and sets the default values.
func NewProperty()(*Property) {
    m := &Property{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Property) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAliases gets the aliases property value. A set of aliases or a friendly names for the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^. Optional.
func (m *Property) GetAliases()([]string) {
    return m.aliases
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Property) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aliases"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetAliases(res)
        }
        return nil
    }
    res["isQueryable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsQueryable(val)
        }
        return nil
    }
    res["isRefinable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRefinable(val)
        }
        return nil
    }
    res["isRetrievable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRetrievable(val)
        }
        return nil
    }
    res["isSearchable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSearchable(val)
        }
        return nil
    }
    res["labels"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseLabel)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Label, len(val))
            for i, v := range val {
                res[i] = *(v.(*Label))
            }
            m.SetLabels(res)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePropertyType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*PropertyType))
        }
        return nil
    }
    return res
}
// GetIsQueryable gets the isQueryable property value. Specifies if the property is queryable. Queryable properties can be used in Keyword Query Language (KQL) queries. Optional.
func (m *Property) GetIsQueryable()(*bool) {
    return m.isQueryable
}
// GetIsRefinable gets the isRefinable property value. Specifies if the property is refinable.  Refinable properties can be used to filter search results in the Search API and add a refiner control in the Microsoft Search user experience. Optional.
func (m *Property) GetIsRefinable()(*bool) {
    return m.isRefinable
}
// GetIsRetrievable gets the isRetrievable property value. Specifies if the property is retrievable. Retrievable properties are returned in the result set when items are returned by the search API. Retrievable properties are also available to add to the display template used to render search results. Optional.
func (m *Property) GetIsRetrievable()(*bool) {
    return m.isRetrievable
}
// GetIsSearchable gets the isSearchable property value. Specifies if the property is searchable. Only properties of type String or StringCollection can be searchable. Non-searchable properties are not added to the search index. Optional.
func (m *Property) GetIsSearchable()(*bool) {
    return m.isSearchable
}
// GetLabels gets the labels property value. Specifies one or more well-known tags added against a property. Labels help Microsoft Search understand the semantics of the data in the connection. Adding appropriate labels would result in an enhanced search experience (e.g. better relevance). The possible values are: title, url, createdBy, lastModifiedBy, authors, createdDateTime, lastModifiedDateTime, fileName, fileExtension, unknownFutureValue. Optional.
func (m *Property) GetLabels()([]Label) {
    return m.labels
}
// GetName gets the name property value. The name of the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^.  Required.
func (m *Property) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Property) GetOdataType()(*string) {
    return m.odataType
}
// GetType gets the type property value. The type property
func (m *Property) GetType()(*PropertyType) {
    return m.typeEscaped
}
// Serialize serializes information the current object
func (m *Property) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAliases() != nil {
        err := writer.WriteCollectionOfStringValues("aliases", m.GetAliases())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isQueryable", m.GetIsQueryable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRefinable", m.GetIsRefinable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRetrievable", m.GetIsRetrievable())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSearchable", m.GetIsSearchable())
        if err != nil {
            return err
        }
    }
    if m.GetLabels() != nil {
        err := writer.WriteCollectionOfStringValues("labels", SerializeLabel(m.GetLabels()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err := writer.WriteStringValue("type", &cast)
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
func (m *Property) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAliases sets the aliases property value. A set of aliases or a friendly names for the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^. Optional.
func (m *Property) SetAliases(value []string)() {
    m.aliases = value
}
// SetIsQueryable sets the isQueryable property value. Specifies if the property is queryable. Queryable properties can be used in Keyword Query Language (KQL) queries. Optional.
func (m *Property) SetIsQueryable(value *bool)() {
    m.isQueryable = value
}
// SetIsRefinable sets the isRefinable property value. Specifies if the property is refinable.  Refinable properties can be used to filter search results in the Search API and add a refiner control in the Microsoft Search user experience. Optional.
func (m *Property) SetIsRefinable(value *bool)() {
    m.isRefinable = value
}
// SetIsRetrievable sets the isRetrievable property value. Specifies if the property is retrievable. Retrievable properties are returned in the result set when items are returned by the search API. Retrievable properties are also available to add to the display template used to render search results. Optional.
func (m *Property) SetIsRetrievable(value *bool)() {
    m.isRetrievable = value
}
// SetIsSearchable sets the isSearchable property value. Specifies if the property is searchable. Only properties of type String or StringCollection can be searchable. Non-searchable properties are not added to the search index. Optional.
func (m *Property) SetIsSearchable(value *bool)() {
    m.isSearchable = value
}
// SetLabels sets the labels property value. Specifies one or more well-known tags added against a property. Labels help Microsoft Search understand the semantics of the data in the connection. Adding appropriate labels would result in an enhanced search experience (e.g. better relevance). The possible values are: title, url, createdBy, lastModifiedBy, authors, createdDateTime, lastModifiedDateTime, fileName, fileExtension, unknownFutureValue. Optional.
func (m *Property) SetLabels(value []Label)() {
    m.labels = value
}
// SetName sets the name property value. The name of the property. Maximum 32 characters. Only alphanumeric characters allowed. For example, each string may not contain control characters, whitespace, or any of the following: :, ;, ,, (, ), [, ], {, }, %, $, +, !, *, =, &, ?, @, #, /, ~, ', ', <, >, `, ^.  Required.
func (m *Property) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Property) SetOdataType(value *string)() {
    m.odataType = value
}
// SetType sets the type property value. The type property
func (m *Property) SetType(value *PropertyType)() {
    m.typeEscaped = value
}
