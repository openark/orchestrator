package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Drive 
type Drive struct {
    BaseItem
    // Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
    bundles []DriveItemable
    // Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
    driveType *string
    // The list of items the user is following. Only in OneDrive for Business.
    following []DriveItemable
    // All items contained in the drive. Read-only. Nullable.
    items []DriveItemable
    // For drives in SharePoint, the underlying document library list. Read-only. Nullable.
    list Listable
    // Optional. The user account that owns the drive. Read-only.
    owner IdentitySetable
    // Optional. Information about the drive's storage space quota. Read-only.
    quota Quotaable
    // The root folder of the drive. Read-only.
    root DriveItemable
    // The sharePointIds property
    sharePointIds SharepointIdsable
    // Collection of common folders available in OneDrive. Read-only. Nullable.
    special []DriveItemable
    // If present, indicates that this is a system-managed drive. Read-only.
    system SystemFacetable
}
// NewDrive instantiates a new drive and sets the default values.
func NewDrive()(*Drive) {
    m := &Drive{
        BaseItem: *NewBaseItem(),
    }
    odataTypeValue := "#microsoft.graph.drive"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDriveFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDriveFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDrive(), nil
}
// GetBundles gets the bundles property value. Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *Drive) GetBundles()([]DriveItemable) {
    return m.bundles
}
// GetDriveType gets the driveType property value. Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
func (m *Drive) GetDriveType()(*string) {
    return m.driveType
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Drive) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseItem.GetFieldDeserializers()
    res["bundles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
            }
            m.SetBundles(res)
        }
        return nil
    }
    res["driveType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDriveType(val)
        }
        return nil
    }
    res["following"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
            }
            m.SetFollowing(res)
        }
        return nil
    }
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
            }
            m.SetItems(res)
        }
        return nil
    }
    res["list"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateListFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetList(val.(Listable))
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(IdentitySetable))
        }
        return nil
    }
    res["quota"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateQuotaFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuota(val.(Quotaable))
        }
        return nil
    }
    res["root"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRoot(val.(DriveItemable))
        }
        return nil
    }
    res["sharePointIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharepointIdsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSharePointIds(val.(SharepointIdsable))
        }
        return nil
    }
    res["special"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDriveItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DriveItemable, len(val))
            for i, v := range val {
                res[i] = v.(DriveItemable)
            }
            m.SetSpecial(res)
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
// GetFollowing gets the following property value. The list of items the user is following. Only in OneDrive for Business.
func (m *Drive) GetFollowing()([]DriveItemable) {
    return m.following
}
// GetItems gets the items property value. All items contained in the drive. Read-only. Nullable.
func (m *Drive) GetItems()([]DriveItemable) {
    return m.items
}
// GetList gets the list property value. For drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *Drive) GetList()(Listable) {
    return m.list
}
// GetOwner gets the owner property value. Optional. The user account that owns the drive. Read-only.
func (m *Drive) GetOwner()(IdentitySetable) {
    return m.owner
}
// GetQuota gets the quota property value. Optional. Information about the drive's storage space quota. Read-only.
func (m *Drive) GetQuota()(Quotaable) {
    return m.quota
}
// GetRoot gets the root property value. The root folder of the drive. Read-only.
func (m *Drive) GetRoot()(DriveItemable) {
    return m.root
}
// GetSharePointIds gets the sharePointIds property value. The sharePointIds property
func (m *Drive) GetSharePointIds()(SharepointIdsable) {
    return m.sharePointIds
}
// GetSpecial gets the special property value. Collection of common folders available in OneDrive. Read-only. Nullable.
func (m *Drive) GetSpecial()([]DriveItemable) {
    return m.special
}
// GetSystem gets the system property value. If present, indicates that this is a system-managed drive. Read-only.
func (m *Drive) GetSystem()(SystemFacetable) {
    return m.system
}
// Serialize serializes information the current object
func (m *Drive) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBundles() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBundles()))
        for i, v := range m.GetBundles() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("bundles", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("driveType", m.GetDriveType())
        if err != nil {
            return err
        }
    }
    if m.GetFollowing() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFollowing()))
        for i, v := range m.GetFollowing() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("following", cast)
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
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("quota", m.GetQuota())
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
        err = writer.WriteObjectValue("sharePointIds", m.GetSharePointIds())
        if err != nil {
            return err
        }
    }
    if m.GetSpecial() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSpecial()))
        for i, v := range m.GetSpecial() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("special", cast)
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
// SetBundles sets the bundles property value. Collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *Drive) SetBundles(value []DriveItemable)() {
    m.bundles = value
}
// SetDriveType sets the driveType property value. Describes the type of drive represented by this resource. OneDrive personal drives will return personal. OneDrive for Business will return business. SharePoint document libraries will return documentLibrary. Read-only.
func (m *Drive) SetDriveType(value *string)() {
    m.driveType = value
}
// SetFollowing sets the following property value. The list of items the user is following. Only in OneDrive for Business.
func (m *Drive) SetFollowing(value []DriveItemable)() {
    m.following = value
}
// SetItems sets the items property value. All items contained in the drive. Read-only. Nullable.
func (m *Drive) SetItems(value []DriveItemable)() {
    m.items = value
}
// SetList sets the list property value. For drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *Drive) SetList(value Listable)() {
    m.list = value
}
// SetOwner sets the owner property value. Optional. The user account that owns the drive. Read-only.
func (m *Drive) SetOwner(value IdentitySetable)() {
    m.owner = value
}
// SetQuota sets the quota property value. Optional. Information about the drive's storage space quota. Read-only.
func (m *Drive) SetQuota(value Quotaable)() {
    m.quota = value
}
// SetRoot sets the root property value. The root folder of the drive. Read-only.
func (m *Drive) SetRoot(value DriveItemable)() {
    m.root = value
}
// SetSharePointIds sets the sharePointIds property value. The sharePointIds property
func (m *Drive) SetSharePointIds(value SharepointIdsable)() {
    m.sharePointIds = value
}
// SetSpecial sets the special property value. Collection of common folders available in OneDrive. Read-only. Nullable.
func (m *Drive) SetSpecial(value []DriveItemable)() {
    m.special = value
}
// SetSystem sets the system property value. If present, indicates that this is a system-managed drive. Read-only.
func (m *Drive) SetSystem(value SystemFacetable)() {
    m.system = value
}
