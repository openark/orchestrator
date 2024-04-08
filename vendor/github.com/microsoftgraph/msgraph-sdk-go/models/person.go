package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Person 
type Person struct {
    Entity
    // The person's birthday.
    birthday *string
    // The name of the person's company.
    companyName *string
    // The person's department.
    department *string
    // The person's display name.
    displayName *string
    // The person's given name.
    givenName *string
    // The instant message voice over IP (VOIP) session initiation protocol (SIP) address for the user. Read-only.
    imAddress *string
    // true if the user has flagged this person as a favorite.
    isFavorite *bool
    // The person's job title.
    jobTitle *string
    // The location of the person's office.
    officeLocation *string
    // Free-form notes that the user has taken about this person.
    personNotes *string
    // The type of person.
    personType PersonTypeable
    // The person's phone numbers.
    phones []Phoneable
    // The person's addresses.
    postalAddresses []Locationable
    // The person's profession.
    profession *string
    // The person's email addresses.
    scoredEmailAddresses []ScoredEmailAddressable
    // The person's surname.
    surname *string
    // The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
    userPrincipalName *string
    // The person's websites.
    websites []Websiteable
    // The phonetic Japanese name of the person's company.
    yomiCompany *string
}
// NewPerson instantiates a new person and sets the default values.
func NewPerson()(*Person) {
    m := &Person{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePersonFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPerson(), nil
}
// GetBirthday gets the birthday property value. The person's birthday.
func (m *Person) GetBirthday()(*string) {
    return m.birthday
}
// GetCompanyName gets the companyName property value. The name of the person's company.
func (m *Person) GetCompanyName()(*string) {
    return m.companyName
}
// GetDepartment gets the department property value. The person's department.
func (m *Person) GetDepartment()(*string) {
    return m.department
}
// GetDisplayName gets the displayName property value. The person's display name.
func (m *Person) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Person) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["birthday"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBirthday(val)
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
    res["imAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImAddress(val)
        }
        return nil
    }
    res["isFavorite"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavorite(val)
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
    res["officeLocation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfficeLocation(val)
        }
        return nil
    }
    res["personNotes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonNotes(val)
        }
        return nil
    }
    res["personType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePersonTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPersonType(val.(PersonTypeable))
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
    res["postalAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Locationable, len(val))
            for i, v := range val {
                res[i] = v.(Locationable)
            }
            m.SetPostalAddresses(res)
        }
        return nil
    }
    res["profession"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProfession(val)
        }
        return nil
    }
    res["scoredEmailAddresses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateScoredEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ScoredEmailAddressable, len(val))
            for i, v := range val {
                res[i] = v.(ScoredEmailAddressable)
            }
            m.SetScoredEmailAddresses(res)
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
    res["websites"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateWebsiteFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Websiteable, len(val))
            for i, v := range val {
                res[i] = v.(Websiteable)
            }
            m.SetWebsites(res)
        }
        return nil
    }
    res["yomiCompany"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYomiCompany(val)
        }
        return nil
    }
    return res
}
// GetGivenName gets the givenName property value. The person's given name.
func (m *Person) GetGivenName()(*string) {
    return m.givenName
}
// GetImAddress gets the imAddress property value. The instant message voice over IP (VOIP) session initiation protocol (SIP) address for the user. Read-only.
func (m *Person) GetImAddress()(*string) {
    return m.imAddress
}
// GetIsFavorite gets the isFavorite property value. true if the user has flagged this person as a favorite.
func (m *Person) GetIsFavorite()(*bool) {
    return m.isFavorite
}
// GetJobTitle gets the jobTitle property value. The person's job title.
func (m *Person) GetJobTitle()(*string) {
    return m.jobTitle
}
// GetOfficeLocation gets the officeLocation property value. The location of the person's office.
func (m *Person) GetOfficeLocation()(*string) {
    return m.officeLocation
}
// GetPersonNotes gets the personNotes property value. Free-form notes that the user has taken about this person.
func (m *Person) GetPersonNotes()(*string) {
    return m.personNotes
}
// GetPersonType gets the personType property value. The type of person.
func (m *Person) GetPersonType()(PersonTypeable) {
    return m.personType
}
// GetPhones gets the phones property value. The person's phone numbers.
func (m *Person) GetPhones()([]Phoneable) {
    return m.phones
}
// GetPostalAddresses gets the postalAddresses property value. The person's addresses.
func (m *Person) GetPostalAddresses()([]Locationable) {
    return m.postalAddresses
}
// GetProfession gets the profession property value. The person's profession.
func (m *Person) GetProfession()(*string) {
    return m.profession
}
// GetScoredEmailAddresses gets the scoredEmailAddresses property value. The person's email addresses.
func (m *Person) GetScoredEmailAddresses()([]ScoredEmailAddressable) {
    return m.scoredEmailAddresses
}
// GetSurname gets the surname property value. The person's surname.
func (m *Person) GetSurname()(*string) {
    return m.surname
}
// GetUserPrincipalName gets the userPrincipalName property value. The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
func (m *Person) GetUserPrincipalName()(*string) {
    return m.userPrincipalName
}
// GetWebsites gets the websites property value. The person's websites.
func (m *Person) GetWebsites()([]Websiteable) {
    return m.websites
}
// GetYomiCompany gets the yomiCompany property value. The phonetic Japanese name of the person's company.
func (m *Person) GetYomiCompany()(*string) {
    return m.yomiCompany
}
// Serialize serializes information the current object
func (m *Person) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("birthday", m.GetBirthday())
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
        err = writer.WriteStringValue("imAddress", m.GetImAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFavorite", m.GetIsFavorite())
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
        err = writer.WriteStringValue("officeLocation", m.GetOfficeLocation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("personNotes", m.GetPersonNotes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("personType", m.GetPersonType())
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
    if m.GetPostalAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPostalAddresses()))
        for i, v := range m.GetPostalAddresses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("postalAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("profession", m.GetProfession())
        if err != nil {
            return err
        }
    }
    if m.GetScoredEmailAddresses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetScoredEmailAddresses()))
        for i, v := range m.GetScoredEmailAddresses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("scoredEmailAddresses", cast)
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
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    if m.GetWebsites() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetWebsites()))
        for i, v := range m.GetWebsites() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("websites", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("yomiCompany", m.GetYomiCompany())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBirthday sets the birthday property value. The person's birthday.
func (m *Person) SetBirthday(value *string)() {
    m.birthday = value
}
// SetCompanyName sets the companyName property value. The name of the person's company.
func (m *Person) SetCompanyName(value *string)() {
    m.companyName = value
}
// SetDepartment sets the department property value. The person's department.
func (m *Person) SetDepartment(value *string)() {
    m.department = value
}
// SetDisplayName sets the displayName property value. The person's display name.
func (m *Person) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetGivenName sets the givenName property value. The person's given name.
func (m *Person) SetGivenName(value *string)() {
    m.givenName = value
}
// SetImAddress sets the imAddress property value. The instant message voice over IP (VOIP) session initiation protocol (SIP) address for the user. Read-only.
func (m *Person) SetImAddress(value *string)() {
    m.imAddress = value
}
// SetIsFavorite sets the isFavorite property value. true if the user has flagged this person as a favorite.
func (m *Person) SetIsFavorite(value *bool)() {
    m.isFavorite = value
}
// SetJobTitle sets the jobTitle property value. The person's job title.
func (m *Person) SetJobTitle(value *string)() {
    m.jobTitle = value
}
// SetOfficeLocation sets the officeLocation property value. The location of the person's office.
func (m *Person) SetOfficeLocation(value *string)() {
    m.officeLocation = value
}
// SetPersonNotes sets the personNotes property value. Free-form notes that the user has taken about this person.
func (m *Person) SetPersonNotes(value *string)() {
    m.personNotes = value
}
// SetPersonType sets the personType property value. The type of person.
func (m *Person) SetPersonType(value PersonTypeable)() {
    m.personType = value
}
// SetPhones sets the phones property value. The person's phone numbers.
func (m *Person) SetPhones(value []Phoneable)() {
    m.phones = value
}
// SetPostalAddresses sets the postalAddresses property value. The person's addresses.
func (m *Person) SetPostalAddresses(value []Locationable)() {
    m.postalAddresses = value
}
// SetProfession sets the profession property value. The person's profession.
func (m *Person) SetProfession(value *string)() {
    m.profession = value
}
// SetScoredEmailAddresses sets the scoredEmailAddresses property value. The person's email addresses.
func (m *Person) SetScoredEmailAddresses(value []ScoredEmailAddressable)() {
    m.scoredEmailAddresses = value
}
// SetSurname sets the surname property value. The person's surname.
func (m *Person) SetSurname(value *string)() {
    m.surname = value
}
// SetUserPrincipalName sets the userPrincipalName property value. The user principal name (UPN) of the person. The UPN is an Internet-style login name for the person based on the Internet standard RFC 822. By convention, this should map to the person's email name. The general format is alias@domain.
func (m *Person) SetUserPrincipalName(value *string)() {
    m.userPrincipalName = value
}
// SetWebsites sets the websites property value. The person's websites.
func (m *Person) SetWebsites(value []Websiteable)() {
    m.websites = value
}
// SetYomiCompany sets the yomiCompany property value. The phonetic Japanese name of the person's company.
func (m *Person) SetYomiCompany(value *string)() {
    m.yomiCompany = value
}
