package termstore

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Set 
type Set struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
    // Children terms of set in term [store].
    children []Termable
    // Date and time of set creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description that gives details on the term usage.
    description *string
    // Name of the set for each languageTag.
    localizedNames []LocalizedNameable
    // The parentGroup property
    parentGroup Groupable
    // Custom properties for the set.
    properties []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValueable
    // Indicates which terms have been pinned or reused directly under the set.
    relations []Relationable
    // All the terms under the set.
    terms []Termable
}
// NewSet instantiates a new set and sets the default values.
func NewSet()(*Set) {
    m := &Set{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateSetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSet(), nil
}
// GetChildren gets the children property value. Children terms of set in term [store].
func (m *Set) GetChildren()([]Termable) {
    return m.children
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of set creation. Read-only.
func (m *Set) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description that gives details on the term usage.
func (m *Set) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Set) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["children"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Termable, len(val))
            for i, v := range val {
                res[i] = v.(Termable)
            }
            m.SetChildren(res)
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["localizedNames"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocalizedNameFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocalizedNameable, len(val))
            for i, v := range val {
                res[i] = v.(LocalizedNameable)
            }
            m.SetLocalizedNames(res)
        }
        return nil
    }
    res["parentGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentGroup(val.(Groupable))
        }
        return nil
    }
    res["properties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateKeyValueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValueable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValueable)
            }
            m.SetProperties(res)
        }
        return nil
    }
    res["relations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRelationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Relationable, len(val))
            for i, v := range val {
                res[i] = v.(Relationable)
            }
            m.SetRelations(res)
        }
        return nil
    }
    res["terms"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Termable, len(val))
            for i, v := range val {
                res[i] = v.(Termable)
            }
            m.SetTerms(res)
        }
        return nil
    }
    return res
}
// GetLocalizedNames gets the localizedNames property value. Name of the set for each languageTag.
func (m *Set) GetLocalizedNames()([]LocalizedNameable) {
    return m.localizedNames
}
// GetParentGroup gets the parentGroup property value. The parentGroup property
func (m *Set) GetParentGroup()(Groupable) {
    return m.parentGroup
}
// GetProperties gets the properties property value. Custom properties for the set.
func (m *Set) GetProperties()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValueable) {
    return m.properties
}
// GetRelations gets the relations property value. Indicates which terms have been pinned or reused directly under the set.
func (m *Set) GetRelations()([]Relationable) {
    return m.relations
}
// GetTerms gets the terms property value. All the terms under the set.
func (m *Set) GetTerms()([]Termable) {
    return m.terms
}
// Serialize serializes information the current object
func (m *Set) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildren() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChildren()))
        for i, v := range m.GetChildren() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("children", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizedNames() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalizedNames()))
        for i, v := range m.GetLocalizedNames() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("localizedNames", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentGroup", m.GetParentGroup())
        if err != nil {
            return err
        }
    }
    if m.GetProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRelations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRelations()))
        for i, v := range m.GetRelations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("relations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTerms() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTerms()))
        for i, v := range m.GetTerms() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("terms", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildren sets the children property value. Children terms of set in term [store].
func (m *Set) SetChildren(value []Termable)() {
    m.children = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of set creation. Read-only.
func (m *Set) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description that gives details on the term usage.
func (m *Set) SetDescription(value *string)() {
    m.description = value
}
// SetLocalizedNames sets the localizedNames property value. Name of the set for each languageTag.
func (m *Set) SetLocalizedNames(value []LocalizedNameable)() {
    m.localizedNames = value
}
// SetParentGroup sets the parentGroup property value. The parentGroup property
func (m *Set) SetParentGroup(value Groupable)() {
    m.parentGroup = value
}
// SetProperties sets the properties property value. Custom properties for the set.
func (m *Set) SetProperties(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValueable)() {
    m.properties = value
}
// SetRelations sets the relations property value. Indicates which terms have been pinned or reused directly under the set.
func (m *Set) SetRelations(value []Relationable)() {
    m.relations = value
}
// SetTerms sets the terms property value. All the terms under the set.
func (m *Set) SetTerms(value []Termable)() {
    m.terms = value
}
