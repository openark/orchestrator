package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AccessPackageSubject 
type AccessPackageSubject struct {
    Entity
    // The connected organization of the subject. Read-only. Nullable.
    connectedOrganization ConnectedOrganizationable
    // The display name of the subject.
    displayName *string
    // The email address of the subject.
    email *string
    // The object identifier of the subject. null if the subject is not yet a user in the tenant.
    objectId *string
    // A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
    onPremisesSecurityIdentifier *string
    // The principal name, if known, of the subject.
    principalName *string
    // The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue.
    subjectType *AccessPackageSubjectType
}
// NewAccessPackageSubject instantiates a new accessPackageSubject and sets the default values.
func NewAccessPackageSubject()(*AccessPackageSubject) {
    m := &AccessPackageSubject{
        Entity: *NewEntity(),
    }
    return m
}
// CreateAccessPackageSubjectFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAccessPackageSubjectFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAccessPackageSubject(), nil
}
// GetConnectedOrganization gets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
func (m *AccessPackageSubject) GetConnectedOrganization()(ConnectedOrganizationable) {
    return m.connectedOrganization
}
// GetDisplayName gets the displayName property value. The display name of the subject.
func (m *AccessPackageSubject) GetDisplayName()(*string) {
    return m.displayName
}
// GetEmail gets the email property value. The email address of the subject.
func (m *AccessPackageSubject) GetEmail()(*string) {
    return m.email
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AccessPackageSubject) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["connectedOrganization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateConnectedOrganizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectedOrganization(val.(ConnectedOrganizationable))
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
    res["objectId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectId(val)
        }
        return nil
    }
    res["onPremisesSecurityIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSecurityIdentifier(val)
        }
        return nil
    }
    res["principalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalName(val)
        }
        return nil
    }
    res["subjectType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageSubjectType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubjectType(val.(*AccessPackageSubjectType))
        }
        return nil
    }
    return res
}
// GetObjectId gets the objectId property value. The object identifier of the subject. null if the subject is not yet a user in the tenant.
func (m *AccessPackageSubject) GetObjectId()(*string) {
    return m.objectId
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
func (m *AccessPackageSubject) GetOnPremisesSecurityIdentifier()(*string) {
    return m.onPremisesSecurityIdentifier
}
// GetPrincipalName gets the principalName property value. The principal name, if known, of the subject.
func (m *AccessPackageSubject) GetPrincipalName()(*string) {
    return m.principalName
}
// GetSubjectType gets the subjectType property value. The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue.
func (m *AccessPackageSubject) GetSubjectType()(*AccessPackageSubjectType) {
    return m.subjectType
}
// Serialize serializes information the current object
func (m *AccessPackageSubject) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("connectedOrganization", m.GetConnectedOrganization())
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
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("objectId", m.GetObjectId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalName", m.GetPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetSubjectType() != nil {
        cast := (*m.GetSubjectType()).String()
        err = writer.WriteStringValue("subjectType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConnectedOrganization sets the connectedOrganization property value. The connected organization of the subject. Read-only. Nullable.
func (m *AccessPackageSubject) SetConnectedOrganization(value ConnectedOrganizationable)() {
    m.connectedOrganization = value
}
// SetDisplayName sets the displayName property value. The display name of the subject.
func (m *AccessPackageSubject) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEmail sets the email property value. The email address of the subject.
func (m *AccessPackageSubject) SetEmail(value *string)() {
    m.email = value
}
// SetObjectId sets the objectId property value. The object identifier of the subject. null if the subject is not yet a user in the tenant.
func (m *AccessPackageSubject) SetObjectId(value *string)() {
    m.objectId = value
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
func (m *AccessPackageSubject) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// SetPrincipalName sets the principalName property value. The principal name, if known, of the subject.
func (m *AccessPackageSubject) SetPrincipalName(value *string)() {
    m.principalName = value
}
// SetSubjectType sets the subjectType property value. The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue.
func (m *AccessPackageSubject) SetSubjectType(value *AccessPackageSubjectType)() {
    m.subjectType = value
}
