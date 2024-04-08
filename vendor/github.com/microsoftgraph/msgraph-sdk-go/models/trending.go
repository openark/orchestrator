package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Trending 
type Trending struct {
    Entity
    // The lastModifiedDateTime property
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Used for navigating to the trending document.
    resource Entityable
    // Reference properties of the trending document, such as the url and type of the document.
    resourceReference ResourceReferenceable
    // Properties that you can use to visualize the document in your experience.
    resourceVisualization ResourceVisualizationable
    // Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value.
    weight *float64
}
// NewTrending instantiates a new trending and sets the default values.
func NewTrending()(*Trending) {
    m := &Trending{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTrendingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTrendingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTrending(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Trending) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
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
    res["resourceReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceReference(val.(ResourceReferenceable))
        }
        return nil
    }
    res["resourceVisualization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResourceVisualizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourceVisualization(val.(ResourceVisualizationable))
        }
        return nil
    }
    res["weight"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWeight(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Trending) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetResource gets the resource property value. Used for navigating to the trending document.
func (m *Trending) GetResource()(Entityable) {
    return m.resource
}
// GetResourceReference gets the resourceReference property value. Reference properties of the trending document, such as the url and type of the document.
func (m *Trending) GetResourceReference()(ResourceReferenceable) {
    return m.resourceReference
}
// GetResourceVisualization gets the resourceVisualization property value. Properties that you can use to visualize the document in your experience.
func (m *Trending) GetResourceVisualization()(ResourceVisualizationable) {
    return m.resourceVisualization
}
// GetWeight gets the weight property value. Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value.
func (m *Trending) GetWeight()(*float64) {
    return m.weight
}
// Serialize serializes information the current object
func (m *Trending) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("resource", m.GetResource())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat64Value("weight", m.GetWeight())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The lastModifiedDateTime property
func (m *Trending) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetResource sets the resource property value. Used for navigating to the trending document.
func (m *Trending) SetResource(value Entityable)() {
    m.resource = value
}
// SetResourceReference sets the resourceReference property value. Reference properties of the trending document, such as the url and type of the document.
func (m *Trending) SetResourceReference(value ResourceReferenceable)() {
    m.resourceReference = value
}
// SetResourceVisualization sets the resourceVisualization property value. Properties that you can use to visualize the document in your experience.
func (m *Trending) SetResourceVisualization(value ResourceVisualizationable)() {
    m.resourceVisualization = value
}
// SetWeight sets the weight property value. Value indicating how much the document is currently trending. The larger the number, the more the document is currently trending around the user (the more relevant it is). Returned documents are sorted by this value.
func (m *Trending) SetWeight(value *float64)() {
    m.weight = value
}
