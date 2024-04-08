package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UsedInsight 
type UsedInsight struct {
    Entity
    // Information about when the item was last viewed or modified by the user. Read only.
    lastUsed UsageDetailsable
    // Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
    resource Entityable
    // Reference properties of the used document, such as the url and type of the document. Read-only
    resourceReference ResourceReferenceable
    // Properties that you can use to visualize the document in your experience. Read-only
    resourceVisualization ResourceVisualizationable
}
// NewUsedInsight instantiates a new usedInsight and sets the default values.
func NewUsedInsight()(*UsedInsight) {
    m := &UsedInsight{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUsedInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUsedInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUsedInsight(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UsedInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUsageDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val.(UsageDetailsable))
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
    return res
}
// GetLastUsed gets the lastUsed property value. Information about when the item was last viewed or modified by the user. Read only.
func (m *UsedInsight) GetLastUsed()(UsageDetailsable) {
    return m.lastUsed
}
// GetResource gets the resource property value. Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *UsedInsight) GetResource()(Entityable) {
    return m.resource
}
// GetResourceReference gets the resourceReference property value. Reference properties of the used document, such as the url and type of the document. Read-only
func (m *UsedInsight) GetResourceReference()(ResourceReferenceable) {
    return m.resourceReference
}
// GetResourceVisualization gets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *UsedInsight) GetResourceVisualization()(ResourceVisualizationable) {
    return m.resourceVisualization
}
// Serialize serializes information the current object
func (m *UsedInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("lastUsed", m.GetLastUsed())
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
    return nil
}
// SetLastUsed sets the lastUsed property value. Information about when the item was last viewed or modified by the user. Read only.
func (m *UsedInsight) SetLastUsed(value UsageDetailsable)() {
    m.lastUsed = value
}
// SetResource sets the resource property value. Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *UsedInsight) SetResource(value Entityable)() {
    m.resource = value
}
// SetResourceReference sets the resourceReference property value. Reference properties of the used document, such as the url and type of the document. Read-only
func (m *UsedInsight) SetResourceReference(value ResourceReferenceable)() {
    m.resourceReference = value
}
// SetResourceVisualization sets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *UsedInsight) SetResourceVisualization(value ResourceVisualizationable)() {
    m.resourceVisualization = value
}
