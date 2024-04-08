package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// UnifiedGroupSource 
type UnifiedGroupSource struct {
    DataSource
    // The group property
    group iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable
    // Specifies which sources are included in this group. Possible values are: mailbox, site.
    includedSources *SourceType
}
// NewUnifiedGroupSource instantiates a new UnifiedGroupSource and sets the default values.
func NewUnifiedGroupSource()(*UnifiedGroupSource) {
    m := &UnifiedGroupSource{
        DataSource: *NewDataSource(),
    }
    odataTypeValue := "#microsoft.graph.security.unifiedGroupSource"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUnifiedGroupSourceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedGroupSourceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedGroupSource(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedGroupSource) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DataSource.GetFieldDeserializers()
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable))
        }
        return nil
    }
    res["includedSources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIncludedSources(val.(*SourceType))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
func (m *UnifiedGroupSource) GetGroup()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable) {
    return m.group
}
// GetIncludedSources gets the includedSources property value. Specifies which sources are included in this group. Possible values are: mailbox, site.
func (m *UnifiedGroupSource) GetIncludedSources()(*SourceType) {
    return m.includedSources
}
// Serialize serializes information the current object
func (m *UnifiedGroupSource) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DataSource.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    if m.GetIncludedSources() != nil {
        cast := (*m.GetIncludedSources()).String()
        err = writer.WriteStringValue("includedSources", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. The group property
func (m *UnifiedGroupSource) SetGroup(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Groupable)() {
    m.group = value
}
// SetIncludedSources sets the includedSources property value. Specifies which sources are included in this group. Possible values are: mailbox, site.
func (m *UnifiedGroupSource) SetIncludedSources(value *SourceType)() {
    m.includedSources = value
}
