package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// List 
type List struct {
    BaseItem
    // The collection of field definitions for this list.
    columns []ColumnDefinitionable
    // The collection of content types present in this list.
    contentTypes []ContentTypeable
    // The displayable title of the list.
    displayName *string
    // Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
    drive Driveable
    // All items contained in the list.
    items []ListItemable
    // Provides additional details about the list.
    list ListInfoable
    // The collection of long-running operations on the list.
    operations []RichLongRunningOperationable
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds SharepointIdsable
    // The set of subscriptions on the list.
    subscriptions []Subscriptionable
    // If present, indicates that this is a system-managed list. Read-only.
    system SystemFacetable
}
// NewList instantiates a new list and sets the default values.
func NewList()(*List) {
    m := &List{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.list"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateListFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewList(), nil
}
// GetColumns gets the columns property value. The collection of field definitions for this list.
func (m *List) GetColumns()([]ColumnDefinitionable) {
    return m.columns
}
// GetContentTypes gets the contentTypes property value. The collection of content types present in this list.
func (m *List) GetContentTypes()([]ContentTypeable) {
    return m.contentTypes
}
// GetDisplayName gets the displayName property value. The displayable title of the list.
func (m *List) GetDisplayName()(*string) {
    return m.displayName
}
// GetDrive gets the drive property value. Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
func (m *List) GetDrive()(Driveable) {
    return m.drive
}
// GetFieldDeserializers the deserialization information for the current model
func (m *List) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["columns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateColumnDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(ColumnDefinitionable)
            }
            m.SetColumns(res)
        }
        return nil
    }
    res["contentTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateContentTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ContentTypeable, len(val))
            for i, v := range val {
                res[i] = v.(ContentTypeable)
            }
            m.SetContentTypes(res)
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
    res["drive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDrive(val.(Driveable))
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateListItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ListItemable, len(val))
            for i, v := range val {
                res[i] = v.(ListItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["list"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateListInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetList(val.(ListInfoable))
        }
        return nil
    }
    res["operations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRichLongRunningOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]RichLongRunningOperationable, len(val))
            for i, v := range val {
                res[i] = v.(RichLongRunningOperationable)
            }
            m.SetOperations(res)
        }
        return nil
    }
    res["sharepointIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharepointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["subscriptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSubscriptionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Subscriptionable, len(val))
            for i, v := range val {
                res[i] = v.(Subscriptionable)
            }
            m.SetSubscriptions(res)
        }
        return nil
    }
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSystemFacetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(SystemFacetable))
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. All items contained in the list.
func (m *List) GetItems()([]ListItemable) {
    return m.items
}
// GetList gets the list property value. Provides additional details about the list.
func (m *List) GetList()(ListInfoable) {
    return m.list
}
// GetOperations gets the operations property value. The collection of long-running operations on the list.
func (m *List) GetOperations()([]RichLongRunningOperationable) {
    return m.operations
}
// GetSharepointIds gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *List) GetSharepointIds()(SharepointIdsable) {
    return m.sharepointIds
}
// GetSubscriptions gets the subscriptions property value. The set of subscriptions on the list.
func (m *List) GetSubscriptions()([]Subscriptionable) {
    return m.subscriptions
}
// GetSystem gets the system property value. If present, indicates that this is a system-managed list. Read-only.
func (m *List) GetSystem()(SystemFacetable) {
    return m.system
}
// Serialize serializes information the current object
func (m *List) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetColumns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetColumns()))
        for i, v := range m.GetColumns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("columns", cast)
        if err != nil {
            return err
        }
    }
    if m.GetContentTypes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetContentTypes()))
        for i, v := range m.GetContentTypes() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("contentTypes", cast)
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
        err = writer.WriteObjectValue("drive", m.GetDrive())
        if err != nil {
            return err
        }
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("list", m.GetList())
        if err != nil {
            return err
        }
    }
    if m.GetOperations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOperations()))
        for i, v := range m.GetOperations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("operations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sharepointIds", m.GetSharepointIds())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubscriptions()))
        for i, v := range m.GetSubscriptions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("subscriptions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetColumns sets the columns property value. The collection of field definitions for this list.
func (m *List) SetColumns(value []ColumnDefinitionable)() {
    m.columns = value
}
// SetContentTypes sets the contentTypes property value. The collection of content types present in this list.
func (m *List) SetContentTypes(value []ContentTypeable)() {
    m.contentTypes = value
}
// SetDisplayName sets the displayName property value. The displayable title of the list.
func (m *List) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDrive sets the drive property value. Only present on document libraries. Allows access to the list as a [drive][] resource with [driveItems][driveItem].
func (m *List) SetDrive(value Driveable)() {
    m.drive = value
}
// SetItems sets the items property value. All items contained in the list.
func (m *List) SetItems(value []ListItemable)() {
    m.items = value
}
// SetList sets the list property value. Provides additional details about the list.
func (m *List) SetList(value ListInfoable)() {
    m.list = value
}
// SetOperations sets the operations property value. The collection of long-running operations on the list.
func (m *List) SetOperations(value []RichLongRunningOperationable)() {
    m.operations = value
}
// SetSharepointIds sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *List) SetSharepointIds(value SharepointIdsable)() {
    m.sharepointIds = value
}
// SetSubscriptions sets the subscriptions property value. The set of subscriptions on the list.
func (m *List) SetSubscriptions(value []Subscriptionable)() {
    m.subscriptions = value
}
// SetSystem sets the system property value. If present, indicates that this is a system-managed list. Read-only.
func (m *List) SetSystem(value SystemFacetable)() {
    m.system = value
}
