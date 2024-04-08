package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignment 
type EducationAssignment struct {
    Entity
    // Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none. Supported values are: none, assignIfOpen. For example, a teacher can use assignIfOpen to indicate that an assignment should be assigned to any new student who joins the class while the assignment is still open, and none to indicate that an assignment should not be assigned to new students.
    addedStudentAction *EducationAddedStudentAction
    // Optional field to control the assignment behavior  for adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
    addToCalendarAction *EducationAddToCalendarOptions
    // Identifies whether students can submit after the due date. If this property isn't specified during create, it defaults to true.
    allowLateSubmissions *bool
    // Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
    allowStudentsToAddResourcesToSubmission *bool
    // The date when the assignment should become active.  If in the future, the assignment isn't shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    assignDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    assignedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Which users, or whole class should receive a submission object once the assignment is published.
    assignTo EducationAssignmentRecipientable
    // When set, enables users to easily find assignments of a given type.  Read-only. Nullable.
    categories []EducationCategoryable
    // Class which this assignment belongs.
    classId *string
    // Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    closeDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Who created the assignment.
    createdBy IdentitySetable
    // Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Name of the assignment.
    displayName *string
    // Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    dueDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Folder URL where all the feedback file resources for this assignment are stored.
    feedbackResourcesFolderUrl *string
    // How the assignment will be graded.
    grading EducationAssignmentGradeTypeable
    // Instructions for the assignment.  This along with the display name tell the student what to do.
    instructions EducationItemBodyable
    // Who last modified the assignment.
    lastModifiedBy IdentitySetable
    // Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl isn't allowed after the assignment has been published.
    notificationChannelUrl *string
    // Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
    resources []EducationAssignmentResourceable
    // Folder URL where all the file resources for this assignment are stored.
    resourcesFolderUrl *string
    // When set, the grading rubric attached to this assignment.
    rubric EducationRubricable
    // Status of the Assignment.  You can't PATCH this value.  Possible values are: draft, scheduled, published, assigned.
    status *EducationAssignmentStatus
    // Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
    submissions []EducationSubmissionable
    // The deep link URL for the given assignment.
    webUrl *string
}
// NewEducationAssignment instantiates a new EducationAssignment and sets the default values.
func NewEducationAssignment()(*EducationAssignment) {
    m := &EducationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEducationAssignmentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEducationAssignmentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEducationAssignment(), nil
}
// GetAddedStudentAction gets the addedStudentAction property value. Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none. Supported values are: none, assignIfOpen. For example, a teacher can use assignIfOpen to indicate that an assignment should be assigned to any new student who joins the class while the assignment is still open, and none to indicate that an assignment should not be assigned to new students.
func (m *EducationAssignment) GetAddedStudentAction()(*EducationAddedStudentAction) {
    return m.addedStudentAction
}
// GetAddToCalendarAction gets the addToCalendarAction property value. Optional field to control the assignment behavior  for adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
func (m *EducationAssignment) GetAddToCalendarAction()(*EducationAddToCalendarOptions) {
    return m.addToCalendarAction
}
// GetAllowLateSubmissions gets the allowLateSubmissions property value. Identifies whether students can submit after the due date. If this property isn't specified during create, it defaults to true.
func (m *EducationAssignment) GetAllowLateSubmissions()(*bool) {
    return m.allowLateSubmissions
}
// GetAllowStudentsToAddResourcesToSubmission gets the allowStudentsToAddResourcesToSubmission property value. Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
func (m *EducationAssignment) GetAllowStudentsToAddResourcesToSubmission()(*bool) {
    return m.allowStudentsToAddResourcesToSubmission
}
// GetAssignDateTime gets the assignDateTime property value. The date when the assignment should become active.  If in the future, the assignment isn't shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetAssignDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.assignDateTime
}
// GetAssignedDateTime gets the assignedDateTime property value. The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetAssignedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.assignedDateTime
}
// GetAssignTo gets the assignTo property value. Which users, or whole class should receive a submission object once the assignment is published.
func (m *EducationAssignment) GetAssignTo()(EducationAssignmentRecipientable) {
    return m.assignTo
}
// GetCategories gets the categories property value. When set, enables users to easily find assignments of a given type.  Read-only. Nullable.
func (m *EducationAssignment) GetCategories()([]EducationCategoryable) {
    return m.categories
}
// GetClassId gets the classId property value. Class which this assignment belongs.
func (m *EducationAssignment) GetClassId()(*string) {
    return m.classId
}
// GetCloseDateTime gets the closeDateTime property value. Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetCloseDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.closeDateTime
}
// GetCreatedBy gets the createdBy property value. Who created the assignment.
func (m *EducationAssignment) GetCreatedBy()(IdentitySetable) {
    return m.createdBy
}
// GetCreatedDateTime gets the createdDateTime property value. Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDisplayName gets the displayName property value. Name of the assignment.
func (m *EducationAssignment) GetDisplayName()(*string) {
    return m.displayName
}
// GetDueDateTime gets the dueDateTime property value. Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetDueDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dueDateTime
}
// GetFeedbackResourcesFolderUrl gets the feedbackResourcesFolderUrl property value. Folder URL where all the feedback file resources for this assignment are stored.
func (m *EducationAssignment) GetFeedbackResourcesFolderUrl()(*string) {
    return m.feedbackResourcesFolderUrl
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignment) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedStudentAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddedStudentAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddedStudentAction(val.(*EducationAddedStudentAction))
        }
        return nil
    }
    res["addToCalendarAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddToCalendarOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddToCalendarAction(val.(*EducationAddToCalendarOptions))
        }
        return nil
    }
    res["allowLateSubmissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowLateSubmissions(val)
        }
        return nil
    }
    res["allowStudentsToAddResourcesToSubmission"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowStudentsToAddResourcesToSubmission(val)
        }
        return nil
    }
    res["assignDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignDateTime(val)
        }
        return nil
    }
    res["assignedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedDateTime(val)
        }
        return nil
    }
    res["assignTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignTo(val.(EducationAssignmentRecipientable))
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationCategoryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationCategoryable, len(val))
            for i, v := range val {
                res[i] = v.(EducationCategoryable)
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["classId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassId(val)
        }
        return nil
    }
    res["closeDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCloseDateTime(val)
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
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val)
        }
        return nil
    }
    res["feedbackResourcesFolderUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeedbackResourcesFolderUrl(val)
        }
        return nil
    }
    res["grading"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationAssignmentGradeTypeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGrading(val.(EducationAssignmentGradeTypeable))
        }
        return nil
    }
    res["instructions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstructions(val.(EducationItemBodyable))
        }
        return nil
    }
    res["lastModifiedBy"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedBy(val.(IdentitySetable))
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["notificationChannelUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationChannelUrl(val)
        }
        return nil
    }
    res["resources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationAssignmentResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationAssignmentResourceable, len(val))
            for i, v := range val {
                res[i] = v.(EducationAssignmentResourceable)
            }
            m.SetResources(res)
        }
        return nil
    }
    res["resourcesFolderUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResourcesFolderUrl(val)
        }
        return nil
    }
    res["rubric"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEducationRubricFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRubric(val.(EducationRubricable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAssignmentStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*EducationAssignmentStatus))
        }
        return nil
    }
    res["submissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEducationSubmissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]EducationSubmissionable, len(val))
            for i, v := range val {
                res[i] = v.(EducationSubmissionable)
            }
            m.SetSubmissions(res)
        }
        return nil
    }
    res["webUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetGrading gets the grading property value. How the assignment will be graded.
func (m *EducationAssignment) GetGrading()(EducationAssignmentGradeTypeable) {
    return m.grading
}
// GetInstructions gets the instructions property value. Instructions for the assignment.  This along with the display name tell the student what to do.
func (m *EducationAssignment) GetInstructions()(EducationItemBodyable) {
    return m.instructions
}
// GetLastModifiedBy gets the lastModifiedBy property value. Who last modified the assignment.
func (m *EducationAssignment) GetLastModifiedBy()(IdentitySetable) {
    return m.lastModifiedBy
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetNotificationChannelUrl gets the notificationChannelUrl property value. Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl isn't allowed after the assignment has been published.
func (m *EducationAssignment) GetNotificationChannelUrl()(*string) {
    return m.notificationChannelUrl
}
// GetResources gets the resources property value. Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
func (m *EducationAssignment) GetResources()([]EducationAssignmentResourceable) {
    return m.resources
}
// GetResourcesFolderUrl gets the resourcesFolderUrl property value. Folder URL where all the file resources for this assignment are stored.
func (m *EducationAssignment) GetResourcesFolderUrl()(*string) {
    return m.resourcesFolderUrl
}
// GetRubric gets the rubric property value. When set, the grading rubric attached to this assignment.
func (m *EducationAssignment) GetRubric()(EducationRubricable) {
    return m.rubric
}
// GetStatus gets the status property value. Status of the Assignment.  You can't PATCH this value.  Possible values are: draft, scheduled, published, assigned.
func (m *EducationAssignment) GetStatus()(*EducationAssignmentStatus) {
    return m.status
}
// GetSubmissions gets the submissions property value. Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationAssignment) GetSubmissions()([]EducationSubmissionable) {
    return m.submissions
}
// GetWebUrl gets the webUrl property value. The deep link URL for the given assignment.
func (m *EducationAssignment) GetWebUrl()(*string) {
    return m.webUrl
}
// Serialize serializes information the current object
func (m *EducationAssignment) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddedStudentAction() != nil {
        cast := (*m.GetAddedStudentAction()).String()
        err = writer.WriteStringValue("addedStudentAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAddToCalendarAction() != nil {
        cast := (*m.GetAddToCalendarAction()).String()
        err = writer.WriteStringValue("addToCalendarAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowLateSubmissions", m.GetAllowLateSubmissions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowStudentsToAddResourcesToSubmission", m.GetAllowStudentsToAddResourcesToSubmission())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("assignTo", m.GetAssignTo())
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCategories()))
        for i, v := range m.GetCategories() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("categories", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("classId", m.GetClassId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("closeDateTime", m.GetCloseDateTime())
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
        err = writer.WriteTimeValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("grading", m.GetGrading())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("instructions", m.GetInstructions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationChannelUrl", m.GetNotificationChannelUrl())
        if err != nil {
            return err
        }
    }
    if m.GetResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetResources()))
        for i, v := range m.GetResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("resources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("rubric", m.GetRubric())
        if err != nil {
            return err
        }
    }
    if m.GetSubmissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSubmissions()))
        for i, v := range m.GetSubmissions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("submissions", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedStudentAction sets the addedStudentAction property value. Optional field to control the assignment behavior for students who are added after the assignment is published. If not specified, defaults to none. Supported values are: none, assignIfOpen. For example, a teacher can use assignIfOpen to indicate that an assignment should be assigned to any new student who joins the class while the assignment is still open, and none to indicate that an assignment should not be assigned to new students.
func (m *EducationAssignment) SetAddedStudentAction(value *EducationAddedStudentAction)() {
    m.addedStudentAction = value
}
// SetAddToCalendarAction sets the addToCalendarAction property value. Optional field to control the assignment behavior  for adding assignments to students' and teachers' calendars when the assignment is published. The possible values are: none, studentsAndPublisher, studentsAndTeamOwners, unknownFutureValue, and studentsOnly. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: studentsOnly. The default value is none.
func (m *EducationAssignment) SetAddToCalendarAction(value *EducationAddToCalendarOptions)() {
    m.addToCalendarAction = value
}
// SetAllowLateSubmissions sets the allowLateSubmissions property value. Identifies whether students can submit after the due date. If this property isn't specified during create, it defaults to true.
func (m *EducationAssignment) SetAllowLateSubmissions(value *bool)() {
    m.allowLateSubmissions = value
}
// SetAllowStudentsToAddResourcesToSubmission sets the allowStudentsToAddResourcesToSubmission property value. Identifies whether students can add their own resources to a submission or if they can only modify resources added by the teacher.
func (m *EducationAssignment) SetAllowStudentsToAddResourcesToSubmission(value *bool)() {
    m.allowStudentsToAddResourcesToSubmission = value
}
// SetAssignDateTime sets the assignDateTime property value. The date when the assignment should become active.  If in the future, the assignment isn't shown to the student until this date.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetAssignDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignDateTime = value
}
// SetAssignedDateTime sets the assignedDateTime property value. The moment that the assignment was published to students and the assignment shows up on the students timeline.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetAssignedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.assignedDateTime = value
}
// SetAssignTo sets the assignTo property value. Which users, or whole class should receive a submission object once the assignment is published.
func (m *EducationAssignment) SetAssignTo(value EducationAssignmentRecipientable)() {
    m.assignTo = value
}
// SetCategories sets the categories property value. When set, enables users to easily find assignments of a given type.  Read-only. Nullable.
func (m *EducationAssignment) SetCategories(value []EducationCategoryable)() {
    m.categories = value
}
// SetClassId sets the classId property value. Class which this assignment belongs.
func (m *EducationAssignment) SetClassId(value *string)() {
    m.classId = value
}
// SetCloseDateTime sets the closeDateTime property value. Date when the assignment will be closed for submissions. This is an optional field that can be null if the assignment does not allowLateSubmissions or when the closeDateTime is the same as the dueDateTime. But if specified, then the closeDateTime must be greater than or equal to the dueDateTime. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetCloseDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.closeDateTime = value
}
// SetCreatedBy sets the createdBy property value. Who created the assignment.
func (m *EducationAssignment) SetCreatedBy(value IdentitySetable)() {
    m.createdBy = value
}
// SetCreatedDateTime sets the createdDateTime property value. Moment when the assignment was created.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDisplayName sets the displayName property value. Name of the assignment.
func (m *EducationAssignment) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetDueDateTime sets the dueDateTime property value. Date when the students assignment is due.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetDueDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDateTime = value
}
// SetFeedbackResourcesFolderUrl sets the feedbackResourcesFolderUrl property value. Folder URL where all the feedback file resources for this assignment are stored.
func (m *EducationAssignment) SetFeedbackResourcesFolderUrl(value *string)() {
    m.feedbackResourcesFolderUrl = value
}
// SetGrading sets the grading property value. How the assignment will be graded.
func (m *EducationAssignment) SetGrading(value EducationAssignmentGradeTypeable)() {
    m.grading = value
}
// SetInstructions sets the instructions property value. Instructions for the assignment.  This along with the display name tell the student what to do.
func (m *EducationAssignment) SetInstructions(value EducationItemBodyable)() {
    m.instructions = value
}
// SetLastModifiedBy sets the lastModifiedBy property value. Who last modified the assignment.
func (m *EducationAssignment) SetLastModifiedBy(value IdentitySetable)() {
    m.lastModifiedBy = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Moment when the assignment was last modified.  The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *EducationAssignment) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetNotificationChannelUrl sets the notificationChannelUrl property value. Optional field to specify the URL of the channel to post the assignment publish notification. If not specified or null, defaults to the General channel. This field only applies to assignments where the assignTo value is educationAssignmentClassRecipient. Updating the notificationChannelUrl isn't allowed after the assignment has been published.
func (m *EducationAssignment) SetNotificationChannelUrl(value *string)() {
    m.notificationChannelUrl = value
}
// SetResources sets the resources property value. Learning objects that are associated with this assignment.  Only teachers can modify this list. Nullable.
func (m *EducationAssignment) SetResources(value []EducationAssignmentResourceable)() {
    m.resources = value
}
// SetResourcesFolderUrl sets the resourcesFolderUrl property value. Folder URL where all the file resources for this assignment are stored.
func (m *EducationAssignment) SetResourcesFolderUrl(value *string)() {
    m.resourcesFolderUrl = value
}
// SetRubric sets the rubric property value. When set, the grading rubric attached to this assignment.
func (m *EducationAssignment) SetRubric(value EducationRubricable)() {
    m.rubric = value
}
// SetStatus sets the status property value. Status of the Assignment.  You can't PATCH this value.  Possible values are: draft, scheduled, published, assigned.
func (m *EducationAssignment) SetStatus(value *EducationAssignmentStatus)() {
    m.status = value
}
// SetSubmissions sets the submissions property value. Once published, there is a submission object for each student representing their work and grade.  Read-only. Nullable.
func (m *EducationAssignment) SetSubmissions(value []EducationSubmissionable)() {
    m.submissions = value
}
// SetWebUrl sets the webUrl property value. The deep link URL for the given assignment.
func (m *EducationAssignment) SetWebUrl(value *string)() {
    m.webUrl = value
}
