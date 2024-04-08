package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EducationAssignmentDefaultsable 
type EducationAssignmentDefaultsable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAddedStudentAction()(*EducationAddedStudentAction)
    GetAddToCalendarAction()(*EducationAddToCalendarOptions)
    GetDueTime()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)
    GetNotificationChannelUrl()(*string)
    SetAddedStudentAction(value *EducationAddedStudentAction)()
    SetAddToCalendarAction(value *EducationAddToCalendarOptions)()
    SetDueTime(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.TimeOnly)()
    SetNotificationChannelUrl(value *string)()
}
