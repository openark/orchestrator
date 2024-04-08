package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserSecurityState 
type UserSecurityState struct {
    // AAD User object identifier (GUID) - represents the physical/multi-account user entity.
    aadUserId *string
    // Account name of user account (without Active Directory domain or DNS domain) - (also called mailNickName).
    accountName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // NetBIOS/Active Directory domain of user account (that is, domain/account format).
    domainName *string
    // For email-related alerts - user account's email 'role'. Possible values are: unknown, sender, recipient.
    emailRole *EmailRole
    // Indicates whether the user logged on through a VPN.
    isVpn *bool
    // Time at which the sign-in occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    logonDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // User sign-in ID.
    logonId *string
    // IP Address the sign-in request originated from.
    logonIp *string
    // Location (by IP address mapping) associated with a user sign-in event by this user.
    logonLocation *string
    // Method of user sign in. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
    logonType *LogonType
    // The OdataType property
    odataType *string
    // Active Directory (on-premises) Security Identifier (SID) of the user.
    onPremisesSecurityIdentifier *string
    // Provider-generated/calculated risk score of the user account. Recommended value range of 0-1, which equates to a percentage.
    riskScore *string
    // User account type (group membership), per Windows definition. Possible values are: unknown, standard, power, administrator.
    userAccountType *UserAccountSecurityType
    // User sign-in name - internet format: (user account name)@(user account DNS domain name).
    userPrincipalName *string
}
// NewUserSecurityState instantiates a new userSecurityState and sets the default values.
func NewUserSecurityState()(*UserSecurityState) {
    m := &UserSecurityState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateUserSecurityStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserSecurityStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserSecurityState(), nil
}
// GetAadUserId gets the aadUserId property value. AAD User object identifier (GUID) - represents the physical/multi-account user entity.
func (m *UserSecurityState) GetAadUserId()(*string) {
    return m.aadUserId
}
// GetAccountName gets the accountName property value. Account name of user account (without Active Directory domain or DNS domain) - (also called mailNickName).
func (m *UserSecurityState) GetAccountName()(*string) {
    return m.accountName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSecurityState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDomainName gets the domainName property value. NetBIOS/Active Directory domain of user account (that is, domain/account format).
func (m *UserSecurityState) GetDomainName()(*string) {
    return m.domainName
}
// GetEmailRole gets the emailRole property value. For email-related alerts - user account's email 'role'. Possible values are: unknown, sender, recipient.
func (m *UserSecurityState) GetEmailRole()(*EmailRole) {
    return m.emailRole
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserSecurityState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["aadUserId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAadUserId(val)
        }
        return nil
    }
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
    res["emailRole"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEmailRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailRole(val.(*EmailRole))
        }
        return nil
    }
    res["isVpn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVpn(val)
        }
        return nil
    }
    res["logonDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonDateTime(val)
        }
        return nil
    }
    res["logonId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonId(val)
        }
        return nil
    }
    res["logonIp"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonIp(val)
        }
        return nil
    }
    res["logonLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonLocation(val)
        }
        return nil
    }
    res["logonType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseLogonType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogonType(val.(*LogonType))
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
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val)
        }
        return nil
    }
    res["userAccountType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserAccountSecurityType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccountType(val.(*UserAccountSecurityType))
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
// GetIsVpn gets the isVpn property value. Indicates whether the user logged on through a VPN.
func (m *UserSecurityState) GetIsVpn()(*bool) {
    return m.isVpn
}
// GetLogonDateTime gets the logonDateTime property value. Time at which the sign-in occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *UserSecurityState) GetLogonDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.logonDateTime
}
// GetLogonId gets the logonId property value. User sign-in ID.
func (m *UserSecurityState) GetLogonId()(*string) {
    return m.logonId
}
// GetLogonIp gets the logonIp property value. IP Address the sign-in request originated from.
func (m *UserSecurityState) GetLogonIp()(*string) {
    return m.logonIp
}
// GetLogonLocation gets the logonLocation property value. Location (by IP address mapping) associated with a user sign-in event by this user.
func (m *UserSecurityState) GetLogonLocation()(*string) {
    return m.logonLocation
}
// GetLogonType gets the logonType property value. Method of user sign in. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
func (m *UserSecurityState) GetLogonType()(*LogonType) {
    return m.logonType
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *UserSecurityState) GetOdataType()(*string) {
    return m.odataType
}
// GetOnPremisesSecurityIdentifier gets the onPremisesSecurityIdentifier property value. Active Directory (on-premises) Security Identifier (SID) of the user.
func (m *UserSecurityState) GetOnPremisesSecurityIdentifier()(*string) {
    return m.onPremisesSecurityIdentifier
}
// GetRiskScore gets the riskScore property value. Provider-generated/calculated risk score of the user account. Recommended value range of 0-1, which equates to a percentage.
func (m *UserSecurityState) GetRiskScore()(*string) {
    return m.riskScore
}
// GetUserAccountType gets the userAccountType property value. User account type (group membership), per Windows definition. Possible values are: unknown, standard, power, administrator.
func (m *UserSecurityState) GetUserAccountType()(*UserAccountSecurityType) {
    return m.userAccountType
}
// GetUserPrincipalName gets the userPrincipalName property value. User sign-in name - internet format: (user account name)@(user account DNS domain name).
func (m *UserSecurityState) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// Serialize serializes information the current object
func (m *UserSecurityState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("aadUserId", m.GetAadUserId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
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
    if m.GetEmailRole() != nil {
        cast := (*m.GetEmailRole()).String()
        err := writer.WriteStringValue("emailRole", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVpn", m.GetIsVpn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("logonDateTime", m.GetLogonDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logonId", m.GetLogonId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logonIp", m.GetLogonIp())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("logonLocation", m.GetLogonLocation())
        if err != nil {
            return err
        }
    }
    if m.GetLogonType() != nil {
        cast := (*m.GetLogonType()).String()
        err := writer.WriteStringValue("logonType", &cast)
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
        err := writer.WriteStringValue("onPremisesSecurityIdentifier", m.GetOnPremisesSecurityIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
        if err != nil {
            return err
        }
    }
    if m.GetUserAccountType() != nil {
        cast := (*m.GetUserAccountType()).String()
        err := writer.WriteStringValue("userAccountType", &cast)
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAadUserId sets the aadUserId property value. AAD User object identifier (GUID) - represents the physical/multi-account user entity.
func (m *UserSecurityState) SetAadUserId(value *string)() {
    m.aadUserId = value
}
// SetAccountName sets the accountName property value. Account name of user account (without Active Directory domain or DNS domain) - (also called mailNickName).
func (m *UserSecurityState) SetAccountName(value *string)() {
    m.accountName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UserSecurityState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDomainName sets the domainName property value. NetBIOS/Active Directory domain of user account (that is, domain/account format).
func (m *UserSecurityState) SetDomainName(value *string)() {
    m.domainName = value
}
// SetEmailRole sets the emailRole property value. For email-related alerts - user account's email 'role'. Possible values are: unknown, sender, recipient.
func (m *UserSecurityState) SetEmailRole(value *EmailRole)() {
    m.emailRole = value
}
// SetIsVpn sets the isVpn property value. Indicates whether the user logged on through a VPN.
func (m *UserSecurityState) SetIsVpn(value *bool)() {
    m.isVpn = value
}
// SetLogonDateTime sets the logonDateTime property value. Time at which the sign-in occurred. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *UserSecurityState) SetLogonDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.logonDateTime = value
}
// SetLogonId sets the logonId property value. User sign-in ID.
func (m *UserSecurityState) SetLogonId(value *string)() {
    m.logonId = value
}
// SetLogonIp sets the logonIp property value. IP Address the sign-in request originated from.
func (m *UserSecurityState) SetLogonIp(value *string)() {
    m.logonIp = value
}
// SetLogonLocation sets the logonLocation property value. Location (by IP address mapping) associated with a user sign-in event by this user.
func (m *UserSecurityState) SetLogonLocation(value *string)() {
    m.logonLocation = value
}
// SetLogonType sets the logonType property value. Method of user sign in. Possible values are: unknown, interactive, remoteInteractive, network, batch, service.
func (m *UserSecurityState) SetLogonType(value *LogonType)() {
    m.logonType = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *UserSecurityState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOnPremisesSecurityIdentifier sets the onPremisesSecurityIdentifier property value. Active Directory (on-premises) Security Identifier (SID) of the user.
func (m *UserSecurityState) SetOnPremisesSecurityIdentifier(value *string)() {
    m.onPremisesSecurityIdentifier = value
}
// SetRiskScore sets the riskScore property value. Provider-generated/calculated risk score of the user account. Recommended value range of 0-1, which equates to a percentage.
func (m *UserSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
// SetUserAccountType sets the userAccountType property value. User account type (group membership), per Windows definition. Possible values are: unknown, standard, power, administrator.
func (m *UserSecurityState) SetUserAccountType(value *UserAccountSecurityType)() {
    m.userAccountType = value
}
// SetUserPrincipalName sets the userPrincipalName property value. User sign-in name - internet format: (user account name)@(user account DNS domain name).
func (m *UserSecurityState) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
