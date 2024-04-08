package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationSchool 
type EducationSchool struct {
    EducationOrganization
    // Address of the school.
    address PhysicalAddressable
    // The underlying administrativeUnit for this school.
    administrativeUnit AdministrativeUnitable
    // Classes taught at the school. Nullable.
    classes []EducationClassable
    // Entity who created the school.
    createdBy IdentitySetable
    // ID of school in syncing system.
    externalId *string
    // ID of principal in syncing system.
    externalPrincipalId *string
    // The fax property
    fax *string
    // Highest grade taught.
    highestGrade *string
    // Lowest grade taught.
    lowestGrade *string
    // Phone number of school.
    phone *string
    // Email address of the principal.
    principalEmail *string
    // Name of the principal.
    principalName *string
    // School Number.
    schoolNumber *string
    // Users in the school. Nullable.
    users []EducationUserable
}
// NewEducationSchool instantiates a new EducationSchool and sets the default values.
func NewEducationSchool()(*EducationSchool) {
    m := &EducationSchool{
        EducationOrganization: *NewEducationOrganization(),
    }
    odataTypeValue := "#microsoft.graph.educationSchool"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEducationSchoolFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationSchoolFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationSchool(), nil
}
// GetAddress gets the address property value. Address of the school.
func (m *EducationSchool) GetAddress()(PhysicalAddressable) {
    return m.address
}
// GetAdministrativeUnit gets the administrativeUnit property value. The underlying administrativeUnit for this school.
func (m *EducationSchool) GetAdministrativeUnit()(AdministrativeUnitable) {
    return m.administrativeUnit
}
// GetClasses gets the classes property value. Classes taught at the school. Nullable.
func (m *EducationSchool) GetClasses()([]EducationClassable) {
    return m.classes
}
// GetCreatedBy gets the createdBy property value. Entity who created the school.
func (m *EducationSchool) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetExternalId gets the externalId property value. ID of school in syncing system.
func (m *EducationSchool) GetExternalId()(*string) {
    return m.externalId
}
// GetExternalPrincipalId gets the externalPrincipalId property value. ID of principal in syncing system.
func (m *EducationSchool) GetExternalPrincipalId()(*string) {
    return m.externalPrincipalId
}
// GetFax gets the fax property value. The fax property
func (m *EducationSchool) GetFax()(*string) {
    return m.fax
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationSchool) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.EducationOrganization.GetFieldDeserializers()
    res["address"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePhysicalAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val.(PhysicalAddressable))
        }
        return nil
    }
    res["administrativeUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAdministrativeUnitFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdministrativeUnit(val.(AdministrativeUnitable))
        }
        return nil
    }
    res["classes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationClassFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationClassable, len(val))
            for i, v := range val {
                res[i] = v.(EducationClassable)
            }
            m.SetClasses(res)
        }
        return nil
    }
    res["createdBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["externalPrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalPrincipalId(val)
        }
        return nil
    }
    res["fax"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFax(val)
        }
        return nil
    }
    res["highestGrade"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHighestGrade(val)
        }
        return nil
    }
    res["lowestGrade"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLowestGrade(val)
        }
        return nil
    }
    res["phone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhone(val)
        }
        return nil
    }
    res["principalEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrincipalEmail(val)
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
    res["schoolNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSchoolNumber(val)
        }
        return nil
    }
    res["users"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationUserable, len(val))
            for i, v := range val {
                res[i] = v.(EducationUserable)
            }
            m.SetUsers(res)
        }
        return nil
    }
    return res
}
// GetHighestGrade gets the highestGrade property value. Highest grade taught.
func (m *EducationSchool) GetHighestGrade()(*string) {
    return m.highestGrade
}
// GetLowestGrade gets the lowestGrade property value. Lowest grade taught.
func (m *EducationSchool) GetLowestGrade()(*string) {
    return m.lowestGrade
}
// GetPhone gets the phone property value. Phone number of school.
func (m *EducationSchool) GetPhone()(*string) {
    return m.phone
}
// GetPrincipalEmail gets the principalEmail property value. Email address of the principal.
func (m *EducationSchool) GetPrincipalEmail()(*string) {
    return m.principalEmail
}
// GetPrincipalName gets the principalName property value. Name of the principal.
func (m *EducationSchool) GetPrincipalName()(*string) {
    return m.principalName
}
// GetSchoolNumber gets the schoolNumber property value. School Number.
func (m *EducationSchool) GetSchoolNumber()(*string) {
    return m.schoolNumber
}
// GetUsers gets the users property value. Users in the school. Nullable.
func (m *EducationSchool) GetUsers()([]EducationUserable) {
    return m.users
}
// Serialize serializes information the current object
func (m *EducationSchool) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.EducationOrganization.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("administrativeUnit", m.GetAdministrativeUnit())
        if err != nil {
            return err
        }
    }
    if m.GetClasses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClasses()))
        for i, v := range m.GetClasses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("classes", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalPrincipalId", m.GetExternalPrincipalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fax", m.GetFax())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("highestGrade", m.GetHighestGrade())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("lowestGrade", m.GetLowestGrade())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phone", m.GetPhone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("principalEmail", m.GetPrincipalEmail())
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
    {
        err = writer.WriteStringValue("schoolNumber", m.GetSchoolNumber())
        if err != nil {
            return err
        }
    }
    if m.GetUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetUsers()))
        for i, v := range m.GetUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("users", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddress sets the address property value. Address of the school.
func (m *EducationSchool) SetAddress(value PhysicalAddressable)() {
    m.address = value
}
// SetAdministrativeUnit sets the administrativeUnit property value. The underlying administrativeUnit for this school.
func (m *EducationSchool) SetAdministrativeUnit(value AdministrativeUnitable)() {
    m.administrativeUnit = value
}
// SetClasses sets the classes property value. Classes taught at the school. Nullable.
func (m *EducationSchool) SetClasses(value []EducationClassable)() {
    m.classes = value
}
// SetCreatedBy sets the createdBy property value. Entity who created the school.
func (m *EducationSchool) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetExternalId sets the externalId property value. ID of school in syncing system.
func (m *EducationSchool) SetExternalId(value *string)() {
    m.externalId = value
}
// SetExternalPrincipalId sets the externalPrincipalId property value. ID of principal in syncing system.
func (m *EducationSchool) SetExternalPrincipalId(value *string)() {
    m.externalPrincipalId = value
}
// SetFax sets the fax property value. The fax property
func (m *EducationSchool) SetFax(value *string)() {
    m.fax = value
}
// SetHighestGrade sets the highestGrade property value. Highest grade taught.
func (m *EducationSchool) SetHighestGrade(value *string)() {
    m.highestGrade = value
}
// SetLowestGrade sets the lowestGrade property value. Lowest grade taught.
func (m *EducationSchool) SetLowestGrade(value *string)() {
    m.lowestGrade = value
}
// SetPhone sets the phone property value. Phone number of school.
func (m *EducationSchool) SetPhone(value *string)() {
    m.phone = value
}
// SetPrincipalEmail sets the principalEmail property value. Email address of the principal.
func (m *EducationSchool) SetPrincipalEmail(value *string)() {
    m.principalEmail = value
}
// SetPrincipalName sets the principalName property value. Name of the principal.
func (m *EducationSchool) SetPrincipalName(value *string)() {
    m.principalName = value
}
// SetSchoolNumber sets the schoolNumber property value. School Number.
func (m *EducationSchool) SetSchoolNumber(value *string)() {
    m.schoolNumber = value
}
// SetUsers sets the users property value. Users in the school. Nullable.
func (m *EducationSchool) SetUsers(value []EducationUserable)() {
    m.users = value
}
