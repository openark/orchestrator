package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Site 
type Site struct {
    BaseItem
    // Analytics about the view activities that took place in this site.
    analytics ItemAnalyticsable
    // The collection of column definitions reusable across lists under this site.
    columns []ColumnDefinitionable
    // The collection of content types defined for this site.
    contentTypes []ContentTypeable
    // The full title for the site. Read-only.
    displayName *string
    // The default drive (document library) for this site.
    drive Driveable
    // The collection of drives (document libraries) under this site.
    drives []Driveable
    // The error property
    error PublicErrorable
    // The externalColumns property
    externalColumns []ColumnDefinitionable
    // Used to address any item contained in this site. This collection can't be enumerated.
    items []BaseItemable
    // The collection of lists under this site.
    lists []Listable
    // Calls the OneNote service for notebook related operations.
    onenote Onenoteable
    // The collection of long-running operations on the site.
    operations []RichLongRunningOperationable
    // The permissions associated with the site. Nullable.
    permissions []Permissionable
    // If present, indicates that this is the root site in the site collection. Read-only.
    root Rootable
    // Returns identifiers useful for SharePoint REST compatibility. Read-only.
    sharepointIds SharepointIdsable
    // Provides details about the site's site collection. Available only on the root site. Read-only.
    siteCollection SiteCollectionable
    // The collection of the sub-sites under this site.
    sites []Siteable
}
// NewSite instantiates a new site and sets the default values.
func NewSite()(*Site) {
    m := &Site{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.site"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSiteFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSiteFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSite(), nil
}
// GetAnalytics gets the analytics property value. Analytics about the view activities that took place in this site.
func (m *Site) GetAnalytics()(ItemAnalyticsable) {
    return m.analytics
}
// GetColumns gets the columns property value. The collection of column definitions reusable across lists under this site.
func (m *Site) GetColumns()([]ColumnDefinitionable) {
    return m.columns
}
// GetContentTypes gets the contentTypes property value. The collection of content types defined for this site.
func (m *Site) GetContentTypes()([]ContentTypeable) {
    return m.contentTypes
}
// GetDisplayName gets the displayName property value. The full title for the site. Read-only.
func (m *Site) GetDisplayName()(*string) {
    return m.displayName
}
// GetDrive gets the drive property value. The default drive (document library) for this site.
func (m *Site) GetDrive()(Driveable) {
    return m.drive
}
// GetDrives gets the drives property value. The collection of drives (document libraries) under this site.
func (m *Site) GetDrives()([]Driveable) {
    return m.drives
}
// GetError gets the error property value. The error property
func (m *Site) GetError()(PublicErrorable) {
    return m.error
}
// GetExternalColumns gets the externalColumns property value. The externalColumns property
func (m *Site) GetExternalColumns()([]ColumnDefinitionable) {
    return m.externalColumns
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Site) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["analytics"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemAnalyticsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAnalytics(val.(ItemAnalyticsable))
        }
        return nil
    }
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
    res["drives"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Driveable, len(val))
            for i, v := range val {
                res[i] = v.(Driveable)
            }
            m.SetDrives(res)
        }
        return nil
    }
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(PublicErrorable))
        }
        return nil
    }
    res["externalColumns"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateColumnDefinitionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ColumnDefinitionable, len(val))
            for i, v := range val {
                res[i] = v.(ColumnDefinitionable)
            }
            m.SetExternalColumns(res)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateBaseItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]BaseItemable, len(val))
            for i, v := range val {
                res[i] = v.(BaseItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["lists"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateListFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Listable, len(val))
            for i, v := range val {
                res[i] = v.(Listable)
            }
            m.SetLists(res)
        }
        return nil
    }
    res["onenote"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnenoteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnenote(val.(Onenoteable))
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
    res["permissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Permissionable, len(val))
            for i, v := range val {
                res[i] = v.(Permissionable)
            }
            m.SetPermissions(res)
        }
        return nil
    }
    res["root"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRootFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(Rootable))
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
    res["siteCollection"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSiteCollectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteCollection(val.(SiteCollectionable))
        }
        return nil
    }
    res["sites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Siteable, len(val))
            for i, v := range val {
                res[i] = v.(Siteable)
            }
            m.SetSites(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. Used to address any item contained in this site. This collection can't be enumerated.
func (m *Site) GetItems()([]BaseItemable) {
    return m.items
}
// GetLists gets the lists property value. The collection of lists under this site.
func (m *Site) GetLists()([]Listable) {
    return m.lists
}
// GetOnenote gets the onenote property value. Calls the OneNote service for notebook related operations.
func (m *Site) GetOnenote()(Onenoteable) {
    return m.onenote
}
// GetOperations gets the operations property value. The collection of long-running operations on the site.
func (m *Site) GetOperations()([]RichLongRunningOperationable) {
    return m.operations
}
// GetPermissions gets the permissions property value. The permissions associated with the site. Nullable.
func (m *Site) GetPermissions()([]Permissionable) {
    return m.permissions
}
// GetRoot gets the root property value. If present, indicates that this is the root site in the site collection. Read-only.
func (m *Site) GetRoot()(Rootable) {
    return m.root
}
// GetSharepointIds gets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *Site) GetSharepointIds()(SharepointIdsable) {
    return m.sharepointIds
}
// GetSiteCollection gets the siteCollection property value. Provides details about the site's site collection. Available only on the root site. Read-only.
func (m *Site) GetSiteCollection()(SiteCollectionable) {
    return m.siteCollection
}
// GetSites gets the sites property value. The collection of the sub-sites under this site.
func (m *Site) GetSites()([]Siteable) {
    return m.sites
}
// Serialize serializes information the current object
func (m *Site) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("analytics", m.GetAnalytics())
        if err != nil {
            return err
        }
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
    if m.GetDrives() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDrives()))
        for i, v := range m.GetDrives() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("drives", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    if m.GetExternalColumns() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExternalColumns()))
        for i, v := range m.GetExternalColumns() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("externalColumns", cast)
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
    if m.GetLists() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLists()))
        for i, v := range m.GetLists() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("lists", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onenote", m.GetOnenote())
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
    if m.GetPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPermissions()))
        for i, v := range m.GetPermissions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("permissions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("root", m.GetRoot())
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
    {
        err = writer.WriteObjectValue("siteCollection", m.GetSiteCollection())
        if err != nil {
            return err
        }
    }
    if m.GetSites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSites()))
        for i, v := range m.GetSites() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("sites", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAnalytics sets the analytics property value. Analytics about the view activities that took place in this site.
func (m *Site) SetAnalytics(value ItemAnalyticsable)() {
    m.analytics = value
}
// SetColumns sets the columns property value. The collection of column definitions reusable across lists under this site.
func (m *Site) SetColumns(value []ColumnDefinitionable)() {
    m.columns = value
}
// SetContentTypes sets the contentTypes property value. The collection of content types defined for this site.
func (m *Site) SetContentTypes(value []ContentTypeable)() {
    m.contentTypes = value
}
// SetDisplayName sets the displayName property value. The full title for the site. Read-only.
func (m *Site) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDrive sets the drive property value. The default drive (document library) for this site.
func (m *Site) SetDrive(value Driveable)() {
    m.drive = value
}
// SetDrives sets the drives property value. The collection of drives (document libraries) under this site.
func (m *Site) SetDrives(value []Driveable)() {
    m.drives = value
}
// SetError sets the error property value. The error property
func (m *Site) SetError(value PublicErrorable)() {
    m.error = value
}
// SetExternalColumns sets the externalColumns property value. The externalColumns property
func (m *Site) SetExternalColumns(value []ColumnDefinitionable)() {
    m.externalColumns = value
}
// SetItems sets the items property value. Used to address any item contained in this site. This collection can't be enumerated.
func (m *Site) SetItems(value []BaseItemable)() {
    m.items = value
}
// SetLists sets the lists property value. The collection of lists under this site.
func (m *Site) SetLists(value []Listable)() {
    m.lists = value
}
// SetOnenote sets the onenote property value. Calls the OneNote service for notebook related operations.
func (m *Site) SetOnenote(value Onenoteable)() {
    m.onenote = value
}
// SetOperations sets the operations property value. The collection of long-running operations on the site.
func (m *Site) SetOperations(value []RichLongRunningOperationable)() {
    m.operations = value
}
// SetPermissions sets the permissions property value. The permissions associated with the site. Nullable.
func (m *Site) SetPermissions(value []Permissionable)() {
    m.permissions = value
}
// SetRoot sets the root property value. If present, indicates that this is the root site in the site collection. Read-only.
func (m *Site) SetRoot(value Rootable)() {
    m.root = value
}
// SetSharepointIds sets the sharepointIds property value. Returns identifiers useful for SharePoint REST compatibility. Read-only.
func (m *Site) SetSharepointIds(value SharepointIdsable)() {
    m.sharepointIds = value
}
// SetSiteCollection sets the siteCollection property value. Provides details about the site's site collection. Available only on the root site. Read-only.
func (m *Site) SetSiteCollection(value SiteCollectionable)() {
    m.siteCollection = value
}
// SetSites sets the sites property value. The collection of the sub-sites under this site.
func (m *Site) SetSites(value []Siteable)() {
    m.sites = value
}
