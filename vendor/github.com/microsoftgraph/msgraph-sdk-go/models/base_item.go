package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// BaseItem 
type BaseItem struct {
    Entity
    // Identity of the user, device, or application which created the item. Read-only.
    createdBy IdentitySetable
    // Identity of the user who created the item. Read-only.
    createdByUser Userable
    // Date and time of item creation. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Provides a user-visible description of the item. Optional.
    description *string
    // ETag for the item. Read-only.
    eTag *string
    // Identity of the user, device, and application which last modified the item. Read-only.
    lastModifiedBy IdentitySetable
    // Identity of the user who last modified the item. Read-only.
    lastModifiedByUser Userable
    // Date and time the item was last modified. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The name of the item. Read-write.
    name *string
    // Parent information, if the item has a parent. Read-write.
    parentReference ItemReferenceable
    // URL that displays the resource in the browser. Read-only.
    webUrl *string
}
// NewBaseItem instantiates a new baseItem and sets the default values.
func NewBaseItem()(*BaseItem) {
    m := &BaseItem{
        Entity: *NewEntity(),
    }
    return m
}
// CreateBaseItemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateBaseItemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.drive":
                        return NewDrive(), nil
                    case "#microsoft.graph.driveItem":
                        return NewDriveItem(), nil
                    case "#microsoft.graph.list":
                        return NewList(), nil
                    case "#microsoft.graph.listItem":
                        return NewListItem(), nil
                    case "#microsoft.graph.sharedDriveItem":
                        return NewSharedDriveItem(), nil
                    case "#microsoft.graph.site":
                        return NewSite(), nil
                }
            }
        }
    }
    return NewBaseItem(), nil
}
// GetCreatedBy gets the createdBy property value. Identity of the user, device, or application which created the item. Read-only.
func (m *BaseItem) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedByUser gets the createdByUser property value. Identity of the user who created the item. Read-only.
func (m *BaseItem) GetCreatedByUser()(Userable) {
    return m.createdByUser
}
// GetCreatedDateTime gets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *BaseItem) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. Provides a user-visible description of the item. Optional.
func (m *BaseItem) GetDescription()(*string) {
    return m.description
}
// GetETag gets the eTag property value. ETag for the item. Read-only.
func (m *BaseItem) GetETag()(*string) {
    return m.eTag
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BaseItem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["createdByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedByUser(val.(Userable))
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
    res["eTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetETag(val)
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedByUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedByUser(val.(Userable))
        }
        return nil
    }
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
    res["parentReference"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentReference(val.(ItemReferenceable))
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedBy gets the lastModifiedBy property value. Identity of the user, device, and application which last modified the item. Read-only.
func (m *BaseItem) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedByUser gets the lastModifiedByUser property value. Identity of the user who last modified the item. Read-only.
func (m *BaseItem) GetLastModifiedByUser()(Userable) {
    return m.lastModifiedByUser
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
func (m *BaseItem) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetName gets the name property value. The name of the item. Read-write.
func (m *BaseItem) GetName()(*string) {
    return m.name
}
// GetParentReference gets the parentReference property value. Parent information, if the item has a parent. Read-write.
func (m *BaseItem) GetParentReference()(ItemReferenceable) {
    return m.parentReference
}
// GetWebUrl gets the webUrl property value. URL that displays the resource in the browser. Read-only.
func (m *BaseItem) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *BaseItem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdByUser", m.GetCreatedByUser())
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
    {
        err = writer.WriteStringValue("eTag", m.GetETag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedBy", m.GetLastModifiedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lastModifiedByUser", m.GetLastModifiedByUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentReference", m.GetParentReference())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedBy sets the createdBy property value. Identity of the user, device, or application which created the item. Read-only.
func (m *BaseItem) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedByUser sets the createdByUser property value. Identity of the user who created the item. Read-only.
func (m *BaseItem) SetCreatedByUser(value Userable)() {
    m.createdByUser = value
}
// SetCreatedDateTime sets the createdDateTime property value. Date and time of item creation. Read-only.
func (m *BaseItem) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. Provides a user-visible description of the item. Optional.
func (m *BaseItem) SetDescription(value *string)() {
    m.description = value
}
// SetETag sets the eTag property value. ETag for the item. Read-only.
func (m *BaseItem) SetETag(value *string)() {
    m.eTag = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Identity of the user, device, and application which last modified the item. Read-only.
func (m *BaseItem) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedByUser sets the lastModifiedByUser property value. Identity of the user who last modified the item. Read-only.
func (m *BaseItem) SetLastModifiedByUser(value Userable)() {
    m.lastModifiedByUser = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Date and time the item was last modified. Read-only.
func (m *BaseItem) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetName sets the name property value. The name of the item. Read-write.
func (m *BaseItem) SetName(value *string)() {
    m.name = value
}
// SetParentReference sets the parentReference property value. Parent information, if the item has a parent. Read-write.
func (m *BaseItem) SetParentReference(value ItemReferenceable)() {
    m.parentReference = value
}
// SetWebUrl sets the webUrl property value. URL that displays the resource in the browser. Read-only.
func (m *BaseItem) SetWebUrl(value *string)() {
    m.webUrl = value
}
