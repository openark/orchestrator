package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserAccount 
type UserAccount struct {
    // The user account's displayed name.
    accountName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The user object identifier in Azure AD.
    azureAdUserId *string
    // The name of the Active Directory domain of which the user is a member.
    domainName *string
    // The OdataType property
    odataType *string
    // The user principal name of the account in Azure AD.
    userPrincipalName *string
    // The local security identifier of the user account.
    userSid *string
}
// NewUserAccount instantiates a new userAccount and sets the default values.
func NewUserAccount()(*UserAccount) {
    m := &UserAccount{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserAccount(), nil
}
// GetAccountName gets the accountName property value. The user account's displayed name.
func (m *UserAccount) GetAccountName()(*string) {
    return m.accountName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAccount) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAzureAdUserId gets the azureAdUserId property value. The user object identifier in Azure AD.
func (m *UserAccount) GetAzureAdUserId()(*string) {
    return m.azureAdUserId
}
// GetDomainName gets the domainName property value. The name of the Active Directory domain of which the user is a member.
func (m *UserAccount) GetDomainName()(*string) {
    return m.domainName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["azureAdUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdUserId(val)
        }
        return nil
    }
    res["domainName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainName(val)
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
    res["userSid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserSid(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserAccount) GetOdataType()(*string) {
    return m.odataType
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name of the account in Azure AD.
func (m *UserAccount) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetUserSid gets the userSid property value. The local security identifier of the user account.
func (m *UserAccount) GetUserSid()(*string) {
    return m.userSid
}
// Serialize serializes information the current object
func (m *UserAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("azureAdUserId", m.GetAzureAdUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("domainName", m.GetDomainName())
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
        err := writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("userSid", m.GetUserSid())
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
// SetAccountName sets the accountName property value. The user account's displayed name.
func (m *UserAccount) SetAccountName(value *string)() {
    m.accountName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserAccount) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAzureAdUserId sets the azureAdUserId property value. The user object identifier in Azure AD.
func (m *UserAccount) SetAzureAdUserId(value *string)() {
    m.azureAdUserId = value
}
// SetDomainName sets the domainName property value. The name of the Active Directory domain of which the user is a member.
func (m *UserAccount) SetDomainName(value *string)() {
    m.domainName = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserAccount) SetOdataType(value *string)() {
    m.odataType = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name of the account in Azure AD.
func (m *UserAccount) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetUserSid sets the userSid property value. The local security identifier of the user account.
func (m *UserAccount) SetUserSid(value *string)() {
    m.userSid = value
}
