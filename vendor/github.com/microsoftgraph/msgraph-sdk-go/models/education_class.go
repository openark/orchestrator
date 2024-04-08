package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationClass 
type EducationClass struct {
    Entity
    // All categories associated with this class. Nullable.
    assignmentCategories []EducationCategoryable
    // Specifies class-level defaults respected by new assignments created in the class.
    assignmentDefaults EducationAssignmentDefaultsable
    // All assignments associated with this class. Nullable.
    assignments []EducationAssignmentable
    // Specifies class-level assignments settings.
    assignmentSettings EducationAssignmentSettingsable
    // Class code used by the school to identify the class.
    classCode *string
    // The course property
    course EducationCourseable
    // Entity who created the class
    createdBy IdentitySetable
    // Description of the class.
    description *string
    // Name of the class.
    displayName *string
    // ID of the class from the syncing system.
    externalId *string
    // Name of the class in the syncing system.
    externalName *string
    // How this class was created. Possible values are: sis, manual.
    externalSource *EducationExternalSource
    // The name of the external source this resources was generated from.
    externalSourceDetail *string
    // Grade level of the class.
    grade *string
    // The underlying Microsoft 365 group object.
    group Groupable
    // Mail name for sending email to all members, if this is enabled.
    mailNickname *string
    // All users in the class. Nullable.
    members []EducationUserable
    // All schools that this class is associated with. Nullable.
    schools []EducationSchoolable
    // All teachers in the class. Nullable.
    teachers []EducationUserable
    // Term for this class.
    term EducationTermable
}
// NewEducationClass instantiates a new EducationClass and sets the default values.
func NewEducationClass()(*EducationClass) {
    m := &EducationClass{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationClassFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationClassFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationClass(), nil
}
// GetAssignmentCategories gets the assignmentCategories property value. All categories associated with this class. Nullable.
func (m *EducationClass) GetAssignmentCategories()([]EducationCategoryable) {
    return m.assignmentCategories
}
// GetAssignmentDefaults gets the assignmentDefaults property value. Specifies class-level defaults respected by new assignments created in the class.
func (m *EducationClass) GetAssignmentDefaults()(EducationAssignmentDefaultsable) {
    return m.assignmentDefaults
}
// GetAssignments gets the assignments property value. All assignments associated with this class. Nullable.
func (m *EducationClass) GetAssignments()([]EducationAssignmentable) {
    return m.assignments
}
// GetAssignmentSettings gets the assignmentSettings property value. Specifies class-level assignments settings.
func (m *EducationClass) GetAssignmentSettings()(EducationAssignmentSettingsable) {
    return m.assignmentSettings
}
// GetClassCode gets the classCode property value. Class code used by the school to identify the class.
func (m *EducationClass) GetClassCode()(*string) {
    return m.classCode
}
// GetCourse gets the course property value. The course property
func (m *EducationClass) GetCourse()(EducationCourseable) {
    return m.course
}
// GetCreatedBy gets the createdBy property value. Entity who created the class
func (m *EducationClass) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetDescription gets the description property value. Description of the class.
func (m *EducationClass) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Name of the class.
func (m *EducationClass) GetDisplayName()(*string) {
    return m.displayName
}
// GetExternalId gets the externalId property value. ID of the class from the syncing system.
func (m *EducationClass) GetExternalId()(*string) {
    return m.externalId
}
// GetExternalName gets the externalName property value. Name of the class in the syncing system.
func (m *EducationClass) GetExternalName()(*string) {
    return m.externalName
}
// GetExternalSource gets the externalSource property value. How this class was created. Possible values are: sis, manual.
func (m *EducationClass) GetExternalSource()(*EducationExternalSource) {
    return m.externalSource
}
// GetExternalSourceDetail gets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationClass) GetExternalSourceDetail()(*string) {
    return m.externalSourceDetail
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationClass) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignmentCategories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(EducationCategoryable)
            }
            m.SetAssignmentCategories(res)
        }
        return nil
    }
    res["assignmentDefaults"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentDefaultsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentDefaults(val.(EducationAssignmentDefaultsable))
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(EducationAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["assignmentSettings"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignmentSettings(val.(EducationAssignmentSettingsable))
        }
        return nil
    }
    res["classCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassCode(val)
        }
        return nil
    }
    res["course"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationCourseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCourse(val.(EducationCourseable))
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
    res["externalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalName(val)
        }
        return nil
    }
    res["externalSource"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationExternalSource)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSource(val.(*EducationExternalSource))
        }
        return nil
    }
    res["externalSourceDetail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalSourceDetail(val)
        }
        return nil
    }
    res["grade"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrade(val)
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(Groupable))
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
    res["members"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationUserable, len(val))
            for i, v := range val {
                res[i] = v.(EducationUserable)
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["schools"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSchoolFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSchoolable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSchoolable)
            }
            m.SetSchools(res)
        }
        return nil
    }
    res["teachers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationUserable, len(val))
            for i, v := range val {
                res[i] = v.(EducationUserable)
            }
            m.SetTeachers(res)
        }
        return nil
    }
    res["term"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationTermFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTerm(val.(EducationTermable))
        }
        return nil
    }
    return res
}
// GetGrade gets the grade property value. Grade level of the class.
func (m *EducationClass) GetGrade()(*string) {
    return m.grade
}
// GetGroup gets the group property value. The underlying Microsoft 365 group object.
func (m *EducationClass) GetGroup()(Groupable) {
    return m.group
}
// GetMailNickname gets the mailNickname property value. Mail name for sending email to all members, if this is enabled.
func (m *EducationClass) GetMailNickname()(*string) {
    return m.mailNickname
}
// GetMembers gets the members property value. All users in the class. Nullable.
func (m *EducationClass) GetMembers()([]EducationUserable) {
    return m.members
}
// GetSchools gets the schools property value. All schools that this class is associated with. Nullable.
func (m *EducationClass) GetSchools()([]EducationSchoolable) {
    return m.schools
}
// GetTeachers gets the teachers property value. All teachers in the class. Nullable.
func (m *EducationClass) GetTeachers()([]EducationUserable) {
    return m.teachers
}
// GetTerm gets the term property value. Term for this class.
func (m *EducationClass) GetTerm()(EducationTermable) {
    return m.term
}
// Serialize serializes information the current object
func (m *EducationClass) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAssignmentCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignmentCategories()))
        for i, v := range m.GetAssignmentCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignmentDefaults", m.GetAssignmentDefaults())
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignmentSettings", m.GetAssignmentSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("classCode", m.GetClassCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("course", m.GetCourse())
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
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalName", m.GetExternalName())
        if err != nil {
            return err
        }
    }
    if m.GetExternalSource() != nil {
        cast := (*m.GetExternalSource()).String()
        err = writer.WriteStringValue("externalSource", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalSourceDetail", m.GetExternalSourceDetail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("grade", m.GetGrade())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
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
    if m.GetMembers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSchools() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSchools()))
        for i, v := range m.GetSchools() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("schools", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTeachers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTeachers()))
        for i, v := range m.GetTeachers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("teachers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("term", m.GetTerm())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignmentCategories sets the assignmentCategories property value. All categories associated with this class. Nullable.
func (m *EducationClass) SetAssignmentCategories(value []EducationCategoryable)() {
    m.assignmentCategories = value
}
// SetAssignmentDefaults sets the assignmentDefaults property value. Specifies class-level defaults respected by new assignments created in the class.
func (m *EducationClass) SetAssignmentDefaults(value EducationAssignmentDefaultsable)() {
    m.assignmentDefaults = value
}
// SetAssignments sets the assignments property value. All assignments associated with this class. Nullable.
func (m *EducationClass) SetAssignments(value []EducationAssignmentable)() {
    m.assignments = value
}
// SetAssignmentSettings sets the assignmentSettings property value. Specifies class-level assignments settings.
func (m *EducationClass) SetAssignmentSettings(value EducationAssignmentSettingsable)() {
    m.assignmentSettings = value
}
// SetClassCode sets the classCode property value. Class code used by the school to identify the class.
func (m *EducationClass) SetClassCode(value *string)() {
    m.classCode = value
}
// SetCourse sets the course property value. The course property
func (m *EducationClass) SetCourse(value EducationCourseable)() {
    m.course = value
}
// SetCreatedBy sets the createdBy property value. Entity who created the class
func (m *EducationClass) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetDescription sets the description property value. Description of the class.
func (m *EducationClass) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Name of the class.
func (m *EducationClass) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetExternalId sets the externalId property value. ID of the class from the syncing system.
func (m *EducationClass) SetExternalId(value *string)() {
    m.externalId = value
}
// SetExternalName sets the externalName property value. Name of the class in the syncing system.
func (m *EducationClass) SetExternalName(value *string)() {
    m.externalName = value
}
// SetExternalSource sets the externalSource property value. How this class was created. Possible values are: sis, manual.
func (m *EducationClass) SetExternalSource(value *EducationExternalSource)() {
    m.externalSource = value
}
// SetExternalSourceDetail sets the externalSourceDetail property value. The name of the external source this resources was generated from.
func (m *EducationClass) SetExternalSourceDetail(value *string)() {
    m.externalSourceDetail = value
}
// SetGrade sets the grade property value. Grade level of the class.
func (m *EducationClass) SetGrade(value *string)() {
    m.grade = value
}
// SetGroup sets the group property value. The underlying Microsoft 365 group object.
func (m *EducationClass) SetGroup(value Groupable)() {
    m.group = value
}
// SetMailNickname sets the mailNickname property value. Mail name for sending email to all members, if this is enabled.
func (m *EducationClass) SetMailNickname(value *string)() {
    m.mailNickname = value
}
// SetMembers sets the members property value. All users in the class. Nullable.
func (m *EducationClass) SetMembers(value []EducationUserable)() {
    m.members = value
}
// SetSchools sets the schools property value. All schools that this class is associated with. Nullable.
func (m *EducationClass) SetSchools(value []EducationSchoolable)() {
    m.schools = value
}
// SetTeachers sets the teachers property value. All teachers in the class. Nullable.
func (m *EducationClass) SetTeachers(value []EducationUserable)() {
    m.teachers = value
}
// SetTerm sets the term property value. Term for this class.
func (m *EducationClass) SetTerm(value EducationTermable)() {
    m.term = value
}
