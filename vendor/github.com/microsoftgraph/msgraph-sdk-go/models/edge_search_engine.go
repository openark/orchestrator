package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdgeSearchEngine 
type EdgeSearchEngine struct {
    EdgeSearchEngineBase
    // Allows IT admind to set a predefined default search engine for MDM-Controlled devices
    edgeSearchEngineType *EdgeSearchEngineType
}
// NewEdgeSearchEngine instantiates a new EdgeSearchEngine and sets the default values.
func NewEdgeSearchEngine()(*EdgeSearchEngine) {
    m := &EdgeSearchEngine{
        EdgeSearchEngineBase: *NewEdgeSearchEngineBase(),
    }
    odataTypeValue := "#microsoft.graph.edgeSearchEngine"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEdgeSearchEngineFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdgeSearchEngineFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdgeSearchEngine(), nil
}
// GetEdgeSearchEngineType gets the edgeSearchEngineType property value. Allows IT admind to set a predefined default search engine for MDM-Controlled devices
func (m *EdgeSearchEngine) GetEdgeSearchEngineType()(*EdgeSearchEngineType) {
    return m.edgeSearchEngineType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdgeSearchEngine) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EdgeSearchEngineBase.GetFieldDeserializers()
    res["edgeSearchEngineType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEdgeSearchEngineType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEdgeSearchEngineType(val.(*EdgeSearchEngineType))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *EdgeSearchEngine) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EdgeSearchEngineBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEdgeSearchEngineType() != nil {
        cast := (*m.GetEdgeSearchEngineType()).String()
        err = writer.WriteStringValue("edgeSearchEngineType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEdgeSearchEngineType sets the edgeSearchEngineType property value. Allows IT admind to set a predefined default search engine for MDM-Controlled devices
func (m *EdgeSearchEngine) SetEdgeSearchEngineType(value *EdgeSearchEngineType)() {
    m.edgeSearchEngineType = value
}
