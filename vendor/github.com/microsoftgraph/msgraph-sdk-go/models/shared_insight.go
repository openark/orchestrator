package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharedInsight 
type SharedInsight struct {
    Entity
    // Details about the shared item. Read only.
    lastShared SharingDetailable
    // The lastSharedMethod property
    lastSharedMethod Entityable
    // Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
    resource Entityable
    // Reference properties of the shared document, such as the url and type of the document. Read-only
    resourceReference ResourceReferenceable
    // Properties that you can use to visualize the document in your experience. Read-only
    resourceVisualization ResourceVisualizationable
    // The sharingHistory property
    sharingHistory []SharingDetailable
}
// NewSharedInsight instantiates a new sharedInsight and sets the default values.
func NewSharedInsight()(*SharedInsight) {
    m := &SharedInsight{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSharedInsightFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharedInsightFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharedInsight(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharedInsight) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["lastShared"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastShared(val.(SharingDetailable))
        }
        return nil
    }
    res["lastSharedMethod"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSharedMethod(val.(Entityable))
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
    res["sharingHistory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharingDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharingDetailable, len(val))
            for i, v := range val {
                res[i] = v.(SharingDetailable)
            }
            m.SetSharingHistory(res)
        }
        return nil
    }
    return res
}
// GetLastShared gets the lastShared property value. Details about the shared item. Read only.
func (m *SharedInsight) GetLastShared()(SharingDetailable) {
    return m.lastShared
}
// GetLastSharedMethod gets the lastSharedMethod property value. The lastSharedMethod property
func (m *SharedInsight) GetLastSharedMethod()(Entityable) {
    return m.lastSharedMethod
}
// GetResource gets the resource property value. Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *SharedInsight) GetResource()(Entityable) {
    return m.resource
}
// GetResourceReference gets the resourceReference property value. Reference properties of the shared document, such as the url and type of the document. Read-only
func (m *SharedInsight) GetResourceReference()(ResourceReferenceable) {
    return m.resourceReference
}
// GetResourceVisualization gets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *SharedInsight) GetResourceVisualization()(ResourceVisualizationable) {
    return m.resourceVisualization
}
// GetSharingHistory gets the sharingHistory property value. The sharingHistory property
func (m *SharedInsight) GetSharingHistory()([]SharingDetailable) {
    return m.sharingHistory
}
// Serialize serializes information the current object
func (m *SharedInsight) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("lastShared", m.GetLastShared())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastSharedMethod", m.GetLastSharedMethod())
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
    if m.GetSharingHistory() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSharingHistory()))
        for i, v := range m.GetSharingHistory() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sharingHistory", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetLastShared sets the lastShared property value. Details about the shared item. Read only.
func (m *SharedInsight) SetLastShared(value SharingDetailable)() {
    m.lastShared = value
}
// SetLastSharedMethod sets the lastSharedMethod property value. The lastSharedMethod property
func (m *SharedInsight) SetLastSharedMethod(value Entityable)() {
    m.lastSharedMethod = value
}
// SetResource sets the resource property value. Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked attachments, the type is driveItem.
func (m *SharedInsight) SetResource(value Entityable)() {
    m.resource = value
}
// SetResourceReference sets the resourceReference property value. Reference properties of the shared document, such as the url and type of the document. Read-only
func (m *SharedInsight) SetResourceReference(value ResourceReferenceable)() {
    m.resourceReference = value
}
// SetResourceVisualization sets the resourceVisualization property value. Properties that you can use to visualize the document in your experience. Read-only
func (m *SharedInsight) SetResourceVisualization(value ResourceVisualizationable)() {
    m.resourceVisualization = value
}
// SetSharingHistory sets the sharingHistory property value. The sharingHistory property
func (m *SharedInsight) SetSharingHistory(value []SharingDetailable)() {
    m.sharingHistory = value
}
