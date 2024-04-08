package models

import (
    i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22 "github.com/google/uuid"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PermissionScope 
type PermissionScope struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A description of the delegated permissions, intended to be read by an administrator granting the permission on behalf of all users. This text appears in tenant-wide admin consent experiences.
    adminConsentDescription *string
    // The permission's title, intended to be read by an administrator granting the permission on behalf of all users.
    adminConsentDisplayName *string
    // Unique delegated permission identifier inside the collection of delegated permissions defined for a resource application.
    id *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID
    // When creating or updating a permission, this property must be set to true (which is the default). To delete a permission, this property must first be set to false.  At that point, in a subsequent call, the permission may be removed.
    isEnabled *bool
    // The OdataType property
    odataType *string
    // The origin property
    origin *string
    // The possible values are: User and Admin. Specifies whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator consent should always be required. While Microsoft Graph defines the default consent requirement for each permission, the tenant administrator may override the behavior in their organization (by allowing, restricting, or limiting user consent to this delegated permission). For more information, see Configure how users consent to applications.
    typeEscaped *string
    // A description of the delegated permissions, intended to be read by a user granting the permission on their own behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves.
    userConsentDescription *string
    // A title for the permission, intended to be read by a user granting the permission on their own behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves.
    userConsentDisplayName *string
    // Specifies the value to include in the scp (scope) claim in access tokens. Must not exceed 120 characters in length. Allowed characters are : ! # $ % & ' ( ) * + , - . / : ;  =  ? @ [ ] ^ + _  {  } ~, as well as characters in the ranges 0-9, A-Z and a-z. Any other character, including the space character, are not allowed. May not begin with ..
    value *string
}
// NewPermissionScope instantiates a new permissionScope and sets the default values.
func NewPermissionScope()(*PermissionScope) {
    m := &PermissionScope{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePermissionScopeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePermissionScopeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPermissionScope(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PermissionScope) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAdminConsentDescription gets the adminConsentDescription property value. A description of the delegated permissions, intended to be read by an administrator granting the permission on behalf of all users. This text appears in tenant-wide admin consent experiences.
func (m *PermissionScope) GetAdminConsentDescription()(*string) {
    return m.adminConsentDescription
}
// GetAdminConsentDisplayName gets the adminConsentDisplayName property value. The permission's title, intended to be read by an administrator granting the permission on behalf of all users.
func (m *PermissionScope) GetAdminConsentDisplayName()(*string) {
    return m.adminConsentDisplayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PermissionScope) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["adminConsentDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsentDescription(val)
        }
        return nil
    }
    res["adminConsentDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdminConsentDisplayName(val)
        }
        return nil
    }
    res["id"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["origin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrigin(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    res["userConsentDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserConsentDescription(val)
        }
        return nil
    }
    res["userConsentDisplayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserConsentDisplayName(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Unique delegated permission identifier inside the collection of delegated permissions defined for a resource application.
func (m *PermissionScope) GetId()(*i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID) {
    return m.id
}
// GetIsEnabled gets the isEnabled property value. When creating or updating a permission, this property must be set to true (which is the default). To delete a permission, this property must first be set to false.  At that point, in a subsequent call, the permission may be removed.
func (m *PermissionScope) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *PermissionScope) GetOdataType()(*string) {
    return m.odataType
}
// GetOrigin gets the origin property value. The origin property
func (m *PermissionScope) GetOrigin()(*string) {
    return m.origin
}
// GetType gets the type property value. The possible values are: User and Admin. Specifies whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator consent should always be required. While Microsoft Graph defines the default consent requirement for each permission, the tenant administrator may override the behavior in their organization (by allowing, restricting, or limiting user consent to this delegated permission). For more information, see Configure how users consent to applications.
func (m *PermissionScope) GetType()(*string) {
    return m.typeEscaped
}
// GetUserConsentDescription gets the userConsentDescription property value. A description of the delegated permissions, intended to be read by a user granting the permission on their own behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves.
func (m *PermissionScope) GetUserConsentDescription()(*string) {
    return m.userConsentDescription
}
// GetUserConsentDisplayName gets the userConsentDisplayName property value. A title for the permission, intended to be read by a user granting the permission on their own behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves.
func (m *PermissionScope) GetUserConsentDisplayName()(*string) {
    return m.userConsentDisplayName
}
// GetValue gets the value property value. Specifies the value to include in the scp (scope) claim in access tokens. Must not exceed 120 characters in length. Allowed characters are : ! # $ % & ' ( ) * + , - . / : ;  =  ? @ [ ] ^ + _  {  } ~, as well as characters in the ranges 0-9, A-Z and a-z. Any other character, including the space character, are not allowed. May not begin with ..
func (m *PermissionScope) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *PermissionScope) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("adminConsentDescription", m.GetAdminConsentDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("adminConsentDisplayName", m.GetAdminConsentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("origin", m.GetOrigin())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userConsentDescription", m.GetUserConsentDescription())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userConsentDisplayName", m.GetUserConsentDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PermissionScope) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAdminConsentDescription sets the adminConsentDescription property value. A description of the delegated permissions, intended to be read by an administrator granting the permission on behalf of all users. This text appears in tenant-wide admin consent experiences.
func (m *PermissionScope) SetAdminConsentDescription(value *string)() {
    m.adminConsentDescription = value
}
// SetAdminConsentDisplayName sets the adminConsentDisplayName property value. The permission's title, intended to be read by an administrator granting the permission on behalf of all users.
func (m *PermissionScope) SetAdminConsentDisplayName(value *string)() {
    m.adminConsentDisplayName = value
}
// SetId sets the id property value. Unique delegated permission identifier inside the collection of delegated permissions defined for a resource application.
func (m *PermissionScope) SetId(value *i561e97a8befe7661a44c8f54600992b4207a3a0cf6770e5559949bc276de2e22.UUID)() {
    m.id = value
}
// SetIsEnabled sets the isEnabled property value. When creating or updating a permission, this property must be set to true (which is the default). To delete a permission, this property must first be set to false.  At that point, in a subsequent call, the permission may be removed.
func (m *PermissionScope) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *PermissionScope) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOrigin sets the origin property value. The origin property
func (m *PermissionScope) SetOrigin(value *string)() {
    m.origin = value
}
// SetType sets the type property value. The possible values are: User and Admin. Specifies whether this delegated permission should be considered safe for non-admin users to consent to on behalf of themselves, or whether an administrator consent should always be required. While Microsoft Graph defines the default consent requirement for each permission, the tenant administrator may override the behavior in their organization (by allowing, restricting, or limiting user consent to this delegated permission). For more information, see Configure how users consent to applications.
func (m *PermissionScope) SetType(value *string)() {
    m.typeEscaped = value
}
// SetUserConsentDescription sets the userConsentDescription property value. A description of the delegated permissions, intended to be read by a user granting the permission on their own behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves.
func (m *PermissionScope) SetUserConsentDescription(value *string)() {
    m.userConsentDescription = value
}
// SetUserConsentDisplayName sets the userConsentDisplayName property value. A title for the permission, intended to be read by a user granting the permission on their own behalf. This text appears in consent experiences where the user is consenting only on behalf of themselves.
func (m *PermissionScope) SetUserConsentDisplayName(value *string)() {
    m.userConsentDisplayName = value
}
// SetValue sets the value property value. Specifies the value to include in the scp (scope) claim in access tokens. Must not exceed 120 characters in length. Allowed characters are : ! # $ % & ' ( ) * + , - . / : ;  =  ? @ [ ] ^ + _  {  } ~, as well as characters in the ranges 0-9, A-Z and a-z. Any other character, including the space character, are not allowed. May not begin with ..
func (m *PermissionScope) SetValue(value *string)() {
    m.value = value
}
