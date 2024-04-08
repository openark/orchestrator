package externalconnectors

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ExternalGroup 
type ExternalGroup struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Entity
    // The description of the external group. Optional.
    description *string
    // The friendly name of the external group. Optional.
    displayName *string
    // A member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or an externalGroup as members.
    members []Identityable
}
// NewExternalGroup instantiates a new externalGroup and sets the default values.
func NewExternalGroup()(*ExternalGroup) {
    m := &ExternalGroup{
        Entity: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewEntity(),
    }
    return m
}
// CreateExternalGroupFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExternalGroupFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewExternalGroup(), nil
}
// GetDescription gets the description property value. The description of the external group. Optional.
func (m *ExternalGroup) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The friendly name of the external group. Optional.
func (m *ExternalGroup) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ExternalGroup) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Identityable, len(val))
            for i, v := range val {
                res[i] = v.(Identityable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    return res
}
// GetMembers gets the members property value. A member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or an externalGroup as members.
func (m *ExternalGroup) GetMembers()([]Identityable) {
    return m.members
}
// Serialize serializes information the current object
func (m *ExternalGroup) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. The description of the external group. Optional.
func (m *ExternalGroup) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The friendly name of the external group. Optional.
func (m *ExternalGroup) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetMembers sets the members property value. A member added to an externalGroup. You can add Azure Active Directory users, Azure Active Directory groups, or an externalGroup as members.
func (m *ExternalGroup) SetMembers(value []Identityable)() {
    m.members = value
}
