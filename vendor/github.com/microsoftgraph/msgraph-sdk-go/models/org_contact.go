package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OrgContact 
type OrgContact struct {
    DirectoryObject
    // The addresses property
    addresses []PhysicalOfficeAddressable
    // The companyName property
    companyName *string
    // The department property
    department *string
    // The directReports property
    directReports []DirectoryObjectable
    // The displayName property
    displayName *string
    // The givenName property
    givenName *string
    // The jobTitle property
    jobTitle *string
    // The mail property
    mail *string
    // The mailNickname property
    mailNickname *string
    // The manager property
    manager DirectoryObjectable
    // The memberOf property
    memberOf []DirectoryObjectable
    // The onPremisesLastSyncDateTime property
    onPremisesLastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The onPremisesProvisioningErrors property
    onPremisesProvisioningErrors []OnPremisesProvisioningErrorable
    // The onPremisesSyncEnabled property
    onPremisesSyncEnabled *bool
    // The phones property
    phones []Phoneable
    // The proxyAddresses property
    proxyAddresses []string
    // The surname property
    surname *string
    // The transitiveMemberOf property
    transitiveMemberOf []DirectoryObjectable
}
// NewOrgContact instantiates a new OrgContact and sets the default values.
func NewOrgContact()(*OrgContact) {
    m := &OrgContact{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.orgContact"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOrgContactFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOrgContactFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrgContact(), nil
}
// GetAddresses gets the addresses property value. The addresses property
func (m *OrgContact) GetAddresses()([]PhysicalOfficeAddressable) {
    return m.addresses
}
// GetCompanyName gets the companyName property value. The companyName property
func (m *OrgContact) GetCompanyName()(*string) {
    return m.companyName
}
// GetDepartment gets the department property value. The department property
func (m *OrgContact) GetDepartment()(*string) {
    return m.department
}
// GetDirectReports gets the directReports property value. The directReports property
func (m *OrgContact) GetDirectReports()([]DirectoryObjectable) {
    return m.directReports
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *OrgContact) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrgContact) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
    res["addresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhysicalOfficeAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PhysicalOfficeAddressable, len(val))
            for i, v := range val {
                res[i] = v.(PhysicalOfficeAddressable)
            }
            m.SetAddresses(res)
        }
        return nil
    }
    res["companyName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompanyName(val)
        }
        return nil
    }
    res["department"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDepartment(val)
        }
        return nil
    }
    res["directReports"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetDirectReports(res)
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
    res["givenName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGivenName(val)
        }
        return nil
    }
    res["jobTitle"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobTitle(val)
        }
        return nil
    }
    res["mail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMail(val)
        }
        return nil
    }
    res["mailNickname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailNickname(val)
        }
        return nil
    }
    res["manager"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManager(val.(DirectoryObjectable))
        }
        return nil
    }
    res["memberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetMemberOf(res)
        }
        return nil
    }
    res["onPremisesLastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesLastSyncDateTime(val)
        }
        return nil
    }
    res["onPremisesProvisioningErrors"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateOnPremisesProvisioningErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnPremisesProvisioningErrorable, len(val))
            for i, v := range val {
                res[i] = v.(OnPremisesProvisioningErrorable)
            }
            m.SetOnPremisesProvisioningErrors(res)
        }
        return nil
    }
    res["onPremisesSyncEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnPremisesSyncEnabled(val)
        }
        return nil
    }
    res["phones"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePhoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Phoneable, len(val))
            for i, v := range val {
                res[i] = v.(Phoneable)
            }
            m.SetPhones(res)
        }
        return nil
    }
    res["proxyAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetProxyAddresses(res)
        }
        return nil
    }
    res["surname"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSurname(val)
        }
        return nil
    }
    res["transitiveMemberOf"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetTransitiveMemberOf(res)
        }
        return nil
    }
    return res
}
// GetGivenName gets the givenName property value. The givenName property
func (m *OrgContact) GetGivenName()(*string) {
    return m.givenName
}
// GetJobTitle gets the jobTitle property value. The jobTitle property
func (m *OrgContact) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetMail gets the mail property value. The mail property
func (m *OrgContact) GetMail()(*string) {
    return m.mail
}
// GetMailNickname gets the mailNickname property value. The mailNickname property
func (m *OrgContact) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetManager gets the manager property value. The manager property
func (m *OrgContact) GetManager()(DirectoryObjectable) {
    return m.manager
}
// GetMemberOf gets the memberOf property value. The memberOf property
func (m *OrgContact) GetMemberOf()([]DirectoryObjectable) {
    return m.memberOf
}
// GetOnPremisesLastSyncDateTime gets the onPremisesLastSyncDateTime property value. The onPremisesLastSyncDateTime property
func (m *OrgContact) GetOnPremisesLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.onPremisesLastSyncDateTime
}
// GetOnPremisesProvisioningErrors gets the onPremisesProvisioningErrors property value. The onPremisesProvisioningErrors property
func (m *OrgContact) GetOnPremisesProvisioningErrors()([]OnPremisesProvisioningErrorable) {
    return m.onPremisesProvisioningErrors
}
// GetOnPremisesSyncEnabled gets the onPremisesSyncEnabled property value. The onPremisesSyncEnabled property
func (m *OrgContact) GetOnPremisesSyncEnabled()(*bool) {
    return m.onPremisesSyncEnabled
}
// GetPhones gets the phones property value. The phones property
func (m *OrgContact) GetPhones()([]Phoneable) {
    return m.phones
}
// GetProxyAddresses gets the proxyAddresses property value. The proxyAddresses property
func (m *OrgContact) GetProxyAddresses()([]string) {
    return m.proxyAddresses
}
// GetSurname gets the surname property value. The surname property
func (m *OrgContact) GetSurname()(*string) {
    return m.surname
}
// GetTransitiveMemberOf gets the transitiveMemberOf property value. The transitiveMemberOf property
func (m *OrgContact) GetTransitiveMemberOf()([]DirectoryObjectable) {
    return m.transitiveMemberOf
}
// Serialize serializes information the current object
func (m *OrgContact) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddresses()))
        for i, v := range m.GetAddresses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("addresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("companyName", m.GetCompanyName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("department", m.GetDepartment())
        if err != nil {
            return err
        }
    }
    if m.GetDirectReports() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDirectReports()))
        for i, v := range m.GetDirectReports() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("directReports", cast)
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
        err = writer.WriteStringValue("givenName", m.GetGivenName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobTitle", m.GetJobTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mail", m.GetMail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mailNickname", m.GetMailNickname())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("manager", m.GetManager())
        if err != nil {
            return err
        }
    }
    if m.GetMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMemberOf()))
        for i, v := range m.GetMemberOf() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("memberOf", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("onPremisesLastSyncDateTime", m.GetOnPremisesLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetOnPremisesProvisioningErrors() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetOnPremisesProvisioningErrors()))
        for i, v := range m.GetOnPremisesProvisioningErrors() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("onPremisesProvisioningErrors", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("onPremisesSyncEnabled", m.GetOnPremisesSyncEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetPhones() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPhones()))
        for i, v := range m.GetPhones() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("phones", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProxyAddresses() != nil {
        err = writer.WriteCollectionOfStringValues("proxyAddresses", m.GetProxyAddresses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("surname", m.GetSurname())
        if err != nil {
            return err
        }
    }
    if m.GetTransitiveMemberOf() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTransitiveMemberOf()))
        for i, v := range m.GetTransitiveMemberOf() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("transitiveMemberOf", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddresses sets the addresses property value. The addresses property
func (m *OrgContact) SetAddresses(value []PhysicalOfficeAddressable)() {
    m.addresses = value
}
// SetCompanyName sets the companyName property value. The companyName property
func (m *OrgContact) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetDepartment sets the department property value. The department property
func (m *OrgContact) SetDepartment(value *string)() {
    m.department = value
}
// SetDirectReports sets the directReports property value. The directReports property
func (m *OrgContact) SetDirectReports(value []DirectoryObjectable)() {
    m.directReports = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *OrgContact) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGivenName sets the givenName property value. The givenName property
func (m *OrgContact) SetGivenName(value *string)() {
    m.givenName = value
}
// SetJobTitle sets the jobTitle property value. The jobTitle property
func (m *OrgContact) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetMail sets the mail property value. The mail property
func (m *OrgContact) SetMail(value *string)() {
    m.mail = value
}
// SetMailNickname sets the mailNickname property value. The mailNickname property
func (m *OrgContact) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetManager sets the manager property value. The manager property
func (m *OrgContact) SetManager(value DirectoryObjectable)() {
    m.manager = value
}
// SetMemberOf sets the memberOf property value. The memberOf property
func (m *OrgContact) SetMemberOf(value []DirectoryObjectable)() {
    m.memberOf = value
}
// SetOnPremisesLastSyncDateTime sets the onPremisesLastSyncDateTime property value. The onPremisesLastSyncDateTime property
func (m *OrgContact) SetOnPremisesLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.onPremisesLastSyncDateTime = value
}
// SetOnPremisesProvisioningErrors sets the onPremisesProvisioningErrors property value. The onPremisesProvisioningErrors property
func (m *OrgContact) SetOnPremisesProvisioningErrors(value []OnPremisesProvisioningErrorable)() {
    m.onPremisesProvisioningErrors = value
}
// SetOnPremisesSyncEnabled sets the onPremisesSyncEnabled property value. The onPremisesSyncEnabled property
func (m *OrgContact) SetOnPremisesSyncEnabled(value *bool)() {
    m.onPremisesSyncEnabled = value
}
// SetPhones sets the phones property value. The phones property
func (m *OrgContact) SetPhones(value []Phoneable)() {
    m.phones = value
}
// SetProxyAddresses sets the proxyAddresses property value. The proxyAddresses property
func (m *OrgContact) SetProxyAddresses(value []string)() {
    m.proxyAddresses = value
}
// SetSurname sets the surname property value. The surname property
func (m *OrgContact) SetSurname(value *string)() {
    m.surname = value
}
// SetTransitiveMemberOf sets the transitiveMemberOf property value. The transitiveMemberOf property
func (m *OrgContact) SetTransitiveMemberOf(value []DirectoryObjectable)() {
    m.transitiveMemberOf = value
}
