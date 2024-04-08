package termstore

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// Group 
type Group struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
    // Date and time of the group creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Description that gives details on the term usage.
    description *string
    // Name of the group.
    displayName *string
    // ID of the parent site of this group.
    parentSiteId *string
    // Returns the type of the group. Possible values are: global, system, and siteCollection.
    scope *TermGroupScope
    // All sets under the group in a term [store].
    sets []Setable
}
// NewGroup instantiates a new group and sets the default values.
func NewGroup()(*Group) {
    m := &Group{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroup(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of the group creation. Read-only.
func (m *Group) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Description that gives details on the term usage.
func (m *Group) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the group.
func (m *Group) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Group) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["parentSiteId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentSiteId(val)
        }
        return nil
    }
    res["scope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTermGroupScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScope(val.(*TermGroupScope))
        }
        return nil
    }
    res["sets"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Setable, len(val))
            for i, v := range val {
                res[i] = v.(Setable)
            }
            m.SetSets(res)
        }
        return nil
    }
    return res
}
// GetParentSiteId gets the parentSiteId property value. ID of the parent site of this group.
func (m *Group) GetParentSiteId()(*string) {
    return m.parentSiteId
}
// GetScope gets the scope property value. Returns the type of the group. Possible values are: global, system, and siteCollection.
func (m *Group) GetScope()(*TermGroupScope) {
    return m.scope
}
// GetSets gets the sets property value. All sets under the group in a term [store].
func (m *Group) GetSets()([]Setable) {
    return m.sets
}
// Serialize serializes information the current object
func (m *Group) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentSiteId", m.GetParentSiteId())
        if err != nil {
            return err
        }
    }
    if m.GetScope() != nil {
        cast := (*m.GetScope()).String()
        err = writer.WriteStringValue("scope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSets() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSets()))
        for i, v := range m.GetSets() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sets", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of the group creation. Read-only.
func (m *Group) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Description that gives details on the term usage.
func (m *Group) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the group.
func (m *Group) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetParentSiteId sets the parentSiteId property value. ID of the parent site of this group.
func (m *Group) SetParentSiteId(value *string)() {
    m.parentSiteId = value
}
// SetScope sets the scope property value. Returns the type of the group. Possible values are: global, system, and siteCollection.
func (m *Group) SetScope(value *TermGroupScope)() {
    m.scope = value
}
// SetSets sets the sets property value. All sets under the group in a term [store].
func (m *Group) SetSets(value []Setable)() {
    m.sets = value
}
