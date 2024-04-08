package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DirectoryRoleTemplate 
type DirectoryRoleTemplate struct {
    DirectoryObject
    // The description to set for the directory role. Read-only.
    description *string
    // The display name to set for the directory role. Read-only.
    displayName *string
}
// NewDirectoryRoleTemplate instantiates a new DirectoryRoleTemplate and sets the default values.
func NewDirectoryRoleTemplate()(*DirectoryRoleTemplate) {
    m := &DirectoryRoleTemplate{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.directoryRoleTemplate"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateDirectoryRoleTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDirectoryRoleTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleTemplate(), nil
}
// GetDescription gets the description property value. The description to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DirectoryRoleTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *DirectoryRoleTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
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
    return nil
}
// SetDescription sets the description property value. The description to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name to set for the directory role. Read-only.
func (m *DirectoryRoleTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
