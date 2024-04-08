package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserIdentity 
type UserIdentity struct {
    Identity
    // Indicates the client IP address used by user performing the activity (audit log only).
    ipAddress *string
    // The userPrincipalName attribute of the user.
    userPrincipalName *string
}
// NewUserIdentity instantiates a new UserIdentity and sets the default values.
func NewUserIdentity()(*UserIdentity) {
    m := &UserIdentity{
        Identity: *NewIdentity(),
    }
    odataTypeValue := "#microsoft.graph.userIdentity"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUserIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserIdentity(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. Indicates the client IP address used by user performing the activity (audit log only).
func (m *UserIdentity) GetIpAddress()(*string) {
    return m.ipAddress
}
// GetUserPrincipalName gets the userPrincipalName property value. The userPrincipalName attribute of the user.
func (m *UserIdentity) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *UserIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIpAddress sets the ipAddress property value. Indicates the client IP address used by user performing the activity (audit log only).
func (m *UserIdentity) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The userPrincipalName attribute of the user.
func (m *UserIdentity) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
