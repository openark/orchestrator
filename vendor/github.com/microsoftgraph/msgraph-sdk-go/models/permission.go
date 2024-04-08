package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Permission 
type Permission struct {
    Entity
    // A format of yyyy-MM-ddTHH:mm:ssZ of DateTimeOffset indicates the expiration time of the permission. DateTime.MinValue indicates there is no expiration set for this permission. Optional.
    expirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The grantedTo property
    grantedTo IdentitySetable
    // The grantedToIdentities property
    grantedToIdentities []IdentitySetable
    // For link type permissions, the details of the users to whom permission was granted. Read-only.
    grantedToIdentitiesV2 []SharePointIdentitySetable
    // For user type permissions, the details of the users and applications for this permission. Read-only.
    grantedToV2 SharePointIdentitySetable
    // Indicates whether the password is set for this permission. This property only appears in the response. Optional. Read-only. For OneDrive Personal only..
    hasPassword *bool
    // Provides a reference to the ancestor of the current permission, if it is inherited from an ancestor. Read-only.
    inheritedFrom ItemReferenceable
    // Details of any associated sharing invitation for this permission. Read-only.
    invitation SharingInvitationable
    // Provides the link details of the current permission, if it is a link type permissions. Read-only.
    link SharingLinkable
    // The type of permission, for example, read. See below for the full list of roles. Read-only.
    roles []string
    // A unique token that can be used to access this shared item via the **shares** API. Read-only.
    shareId *string
}
// NewPermission instantiates a new permission and sets the default values.
func NewPermission()(*Permission) {
    m := &Permission{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermission(), nil
}
// GetExpirationDateTime gets the expirationDateTime property value. A format of yyyy-MM-ddTHH:mm:ssZ of DateTimeOffset indicates the expiration time of the permission. DateTime.MinValue indicates there is no expiration set for this permission. Optional.
func (m *Permission) GetExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.expirationDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Permission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["expirationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExpirationDateTime(val)
        }
        return nil
    }
    res["grantedTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrantedTo(val.(IdentitySetable))
        }
        return nil
    }
    res["grantedToIdentities"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentitySetable, len(val))
            for i, v := range val {
                res[i] = v.(IdentitySetable)
            }
            m.SetGrantedToIdentities(res)
        }
        return nil
    }
    res["grantedToIdentitiesV2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSharePointIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SharePointIdentitySetable, len(val))
            for i, v := range val {
                res[i] = v.(SharePointIdentitySetable)
            }
            m.SetGrantedToIdentitiesV2(res)
        }
        return nil
    }
    res["grantedToV2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrantedToV2(val.(SharePointIdentitySetable))
        }
        return nil
    }
    res["hasPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasPassword(val)
        }
        return nil
    }
    res["inheritedFrom"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemReferenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInheritedFrom(val.(ItemReferenceable))
        }
        return nil
    }
    res["invitation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingInvitationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInvitation(val.(SharingInvitationable))
        }
        return nil
    }
    res["link"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharingLinkFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLink(val.(SharingLinkable))
        }
        return nil
    }
    res["roles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoles(res)
        }
        return nil
    }
    res["shareId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShareId(val)
        }
        return nil
    }
    return res
}
// GetGrantedTo gets the grantedTo property value. The grantedTo property
func (m *Permission) GetGrantedTo()(IdentitySetable) {
    return m.grantedTo
}
// GetGrantedToIdentities gets the grantedToIdentities property value. The grantedToIdentities property
func (m *Permission) GetGrantedToIdentities()([]IdentitySetable) {
    return m.grantedToIdentities
}
// GetGrantedToIdentitiesV2 gets the grantedToIdentitiesV2 property value. For link type permissions, the details of the users to whom permission was granted. Read-only.
func (m *Permission) GetGrantedToIdentitiesV2()([]SharePointIdentitySetable) {
    return m.grantedToIdentitiesV2
}
// GetGrantedToV2 gets the grantedToV2 property value. For user type permissions, the details of the users and applications for this permission. Read-only.
func (m *Permission) GetGrantedToV2()(SharePointIdentitySetable) {
    return m.grantedToV2
}
// GetHasPassword gets the hasPassword property value. Indicates whether the password is set for this permission. This property only appears in the response. Optional. Read-only. For OneDrive Personal only..
func (m *Permission) GetHasPassword()(*bool) {
    return m.hasPassword
}
// GetInheritedFrom gets the inheritedFrom property value. Provides a reference to the ancestor of the current permission, if it is inherited from an ancestor. Read-only.
func (m *Permission) GetInheritedFrom()(ItemReferenceable) {
    return m.inheritedFrom
}
// GetInvitation gets the invitation property value. Details of any associated sharing invitation for this permission. Read-only.
func (m *Permission) GetInvitation()(SharingInvitationable) {
    return m.invitation
}
// GetLink gets the link property value. Provides the link details of the current permission, if it is a link type permissions. Read-only.
func (m *Permission) GetLink()(SharingLinkable) {
    return m.link
}
// GetRoles gets the roles property value. The type of permission, for example, read. See below for the full list of roles. Read-only.
func (m *Permission) GetRoles()([]string) {
    return m.roles
}
// GetShareId gets the shareId property value. A unique token that can be used to access this shared item via the **shares** API. Read-only.
func (m *Permission) GetShareId()(*string) {
    return m.shareId
}
// Serialize serializes information the current object
func (m *Permission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("expirationDateTime", m.GetExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("grantedTo", m.GetGrantedTo())
        if err != nil {
            return err
        }
    }
    if m.GetGrantedToIdentities() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGrantedToIdentities()))
        for i, v := range m.GetGrantedToIdentities() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("grantedToIdentities", cast)
        if err != nil {
            return err
        }
    }
    if m.GetGrantedToIdentitiesV2() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGrantedToIdentitiesV2()))
        for i, v := range m.GetGrantedToIdentitiesV2() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("grantedToIdentitiesV2", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("grantedToV2", m.GetGrantedToV2())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasPassword", m.GetHasPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("inheritedFrom", m.GetInheritedFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("invitation", m.GetInvitation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("link", m.GetLink())
        if err != nil {
            return err
        }
    }
    if m.GetRoles() != nil {
        err = writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("shareId", m.GetShareId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetExpirationDateTime sets the expirationDateTime property value. A format of yyyy-MM-ddTHH:mm:ssZ of DateTimeOffset indicates the expiration time of the permission. DateTime.MinValue indicates there is no expiration set for this permission. Optional.
func (m *Permission) SetExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.expirationDateTime = value
}
// SetGrantedTo sets the grantedTo property value. The grantedTo property
func (m *Permission) SetGrantedTo(value IdentitySetable)() {
    m.grantedTo = value
}
// SetGrantedToIdentities sets the grantedToIdentities property value. The grantedToIdentities property
func (m *Permission) SetGrantedToIdentities(value []IdentitySetable)() {
    m.grantedToIdentities = value
}
// SetGrantedToIdentitiesV2 sets the grantedToIdentitiesV2 property value. For link type permissions, the details of the users to whom permission was granted. Read-only.
func (m *Permission) SetGrantedToIdentitiesV2(value []SharePointIdentitySetable)() {
    m.grantedToIdentitiesV2 = value
}
// SetGrantedToV2 sets the grantedToV2 property value. For user type permissions, the details of the users and applications for this permission. Read-only.
func (m *Permission) SetGrantedToV2(value SharePointIdentitySetable)() {
    m.grantedToV2 = value
}
// SetHasPassword sets the hasPassword property value. Indicates whether the password is set for this permission. This property only appears in the response. Optional. Read-only. For OneDrive Personal only..
func (m *Permission) SetHasPassword(value *bool)() {
    m.hasPassword = value
}
// SetInheritedFrom sets the inheritedFrom property value. Provides a reference to the ancestor of the current permission, if it is inherited from an ancestor. Read-only.
func (m *Permission) SetInheritedFrom(value ItemReferenceable)() {
    m.inheritedFrom = value
}
// SetInvitation sets the invitation property value. Details of any associated sharing invitation for this permission. Read-only.
func (m *Permission) SetInvitation(value SharingInvitationable)() {
    m.invitation = value
}
// SetLink sets the link property value. Provides the link details of the current permission, if it is a link type permissions. Read-only.
func (m *Permission) SetLink(value SharingLinkable)() {
    m.link = value
}
// SetRoles sets the roles property value. The type of permission, for example, read. See below for the full list of roles. Read-only.
func (m *Permission) SetRoles(value []string)() {
    m.roles = value
}
// SetShareId sets the shareId property value. A unique token that can be used to access this shared item via the **shares** API. Read-only.
func (m *Permission) SetShareId(value *string)() {
    m.shareId = value
}
