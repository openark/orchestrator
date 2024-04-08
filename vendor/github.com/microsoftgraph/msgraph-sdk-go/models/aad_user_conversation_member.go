package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AadUserConversationMember 
type AadUserConversationMember struct {
    ConversationMember
    // The email address of the user.
    email *string
    // TenantId which the Azure AD user belongs to.
    tenantId *string
    // The user property
    user Userable
    // The guid of the user.
    userId *string
}
// NewAadUserConversationMember instantiates a new AadUserConversationMember and sets the default values.
func NewAadUserConversationMember()(*AadUserConversationMember) {
    m := &AadUserConversationMember{
        ConversationMember: *NewConversationMember(),
    }
    odataTypeValue := "#microsoft.graph.aadUserConversationMember"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAadUserConversationMemberFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAadUserConversationMemberFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAadUserConversationMember(), nil
}
// GetEmail gets the email property value. The email address of the user.
func (m *AadUserConversationMember) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AadUserConversationMember) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ConversationMember.GetFieldDeserializers()
    res["email"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["tenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    res["user"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUser(val.(Userable))
        }
        return nil
    }
    res["userId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    return res
}
// GetTenantId gets the tenantId property value. TenantId which the Azure AD user belongs to.
func (m *AadUserConversationMember) GetTenantId()(*string) {
    return m.tenantId
}
// GetUser gets the user property value. The user property
func (m *AadUserConversationMember) GetUser()(Userable) {
    return m.user
}
// GetUserId gets the userId property value. The guid of the user.
func (m *AadUserConversationMember) GetUserId()(*string) {
    return m.userId
}
// Serialize serializes information the current object
func (m *AadUserConversationMember) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ConversationMember.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("user", m.GetUser())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEmail sets the email property value. The email address of the user.
func (m *AadUserConversationMember) SetEmail(value *string)() {
    m.email = value
}
// SetTenantId sets the tenantId property value. TenantId which the Azure AD user belongs to.
func (m *AadUserConversationMember) SetTenantId(value *string)() {
    m.tenantId = value
}
// SetUser sets the user property value. The user property
func (m *AadUserConversationMember) SetUser(value Userable)() {
    m.user = value
}
// SetUserId sets the userId property value. The guid of the user.
func (m *AadUserConversationMember) SetUserId(value *string)() {
    m.userId = value
}
