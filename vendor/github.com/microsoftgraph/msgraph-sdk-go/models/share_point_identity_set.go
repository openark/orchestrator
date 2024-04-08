package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SharePointIdentitySet 
type SharePointIdentitySet struct {
    IdentitySet
    // The group associated with this action. Optional.
    group Identityable
    // The SharePoint group associated with this action. Optional.
    siteGroup SharePointIdentityable
    // The SharePoint user associated with this action. Optional.
    siteUser SharePointIdentityable
}
// NewSharePointIdentitySet instantiates a new SharePointIdentitySet and sets the default values.
func NewSharePointIdentitySet()(*SharePointIdentitySet) {
    m := &SharePointIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    odataTypeValue := "#microsoft.graph.sharePointIdentitySet"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSharePointIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSharePointIdentitySetFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSharePointIdentitySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SharePointIdentitySet) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Identityable))
        }
        return nil
    }
    res["siteGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteGroup(val.(SharePointIdentityable))
        }
        return nil
    }
    res["siteUser"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSharePointIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSiteUser(val.(SharePointIdentityable))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group associated with this action. Optional.
func (m *SharePointIdentitySet) GetGroup()(Identityable) {
    return m.group
}
// GetSiteGroup gets the siteGroup property value. The SharePoint group associated with this action. Optional.
func (m *SharePointIdentitySet) GetSiteGroup()(SharePointIdentityable) {
    return m.siteGroup
}
// GetSiteUser gets the siteUser property value. The SharePoint user associated with this action. Optional.
func (m *SharePointIdentitySet) GetSiteUser()(SharePointIdentityable) {
    return m.siteUser
}
// Serialize serializes information the current object
func (m *SharePointIdentitySet) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteGroup", m.GetSiteGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("siteUser", m.GetSiteUser())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. The group associated with this action. Optional.
func (m *SharePointIdentitySet) SetGroup(value Identityable)() {
    m.group = value
}
// SetSiteGroup sets the siteGroup property value. The SharePoint group associated with this action. Optional.
func (m *SharePointIdentitySet) SetSiteGroup(value SharePointIdentityable)() {
    m.siteGroup = value
}
// SetSiteUser sets the siteUser property value. The SharePoint user associated with this action. Optional.
func (m *SharePointIdentitySet) SetSiteUser(value SharePointIdentityable)() {
    m.siteUser = value
}
