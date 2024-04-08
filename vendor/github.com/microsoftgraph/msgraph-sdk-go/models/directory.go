package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Directory 
type Directory struct {
    Entity
    // Conceptual container for user and group directory objects.
    administrativeUnits []AdministrativeUnitable
    // Recently deleted items. Read-only. Nullable.
    deletedItems []DirectoryObjectable
    // Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
    federationConfigurations []IdentityProviderBaseable
}
// NewDirectory instantiates a new Directory and sets the default values.
func NewDirectory()(*Directory) {
    m := &Directory{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectory(), nil
}
// GetAdministrativeUnits gets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) GetAdministrativeUnits()([]AdministrativeUnitable) {
    return m.administrativeUnits
}
// GetDeletedItems gets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) GetDeletedItems()([]DirectoryObjectable) {
    return m.deletedItems
}
// GetFederationConfigurations gets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) GetFederationConfigurations()([]IdentityProviderBaseable) {
    return m.federationConfigurations
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Directory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["administrativeUnits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAdministrativeUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AdministrativeUnitable, len(val))
            for i, v := range val {
                res[i] = v.(AdministrativeUnitable)
            }
            m.SetAdministrativeUnits(res)
        }
        return nil
    }
    res["deletedItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetDeletedItems(res)
        }
        return nil
    }
    res["federationConfigurations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityProviderBaseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IdentityProviderBaseable, len(val))
            for i, v := range val {
                res[i] = v.(IdentityProviderBaseable)
            }
            m.SetFederationConfigurations(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *Directory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAdministrativeUnits() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAdministrativeUnits()))
        for i, v := range m.GetAdministrativeUnits() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("administrativeUnits", cast)
        if err != nil {
            return err
        }
    }
    if m.GetDeletedItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeletedItems()))
        for i, v := range m.GetDeletedItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deletedItems", cast)
        if err != nil {
            return err
        }
    }
    if m.GetFederationConfigurations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetFederationConfigurations()))
        for i, v := range m.GetFederationConfigurations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("federationConfigurations", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdministrativeUnits sets the administrativeUnits property value. Conceptual container for user and group directory objects.
func (m *Directory) SetAdministrativeUnits(value []AdministrativeUnitable)() {
    m.administrativeUnits = value
}
// SetDeletedItems sets the deletedItems property value. Recently deleted items. Read-only. Nullable.
func (m *Directory) SetDeletedItems(value []DirectoryObjectable)() {
    m.deletedItems = value
}
// SetFederationConfigurations sets the federationConfigurations property value. Configure domain federation with organizations whose identity provider (IdP) supports either the SAML or WS-Fed protocol.
func (m *Directory) SetFederationConfigurations(value []IdentityProviderBaseable)() {
    m.federationConfigurations = value
}
