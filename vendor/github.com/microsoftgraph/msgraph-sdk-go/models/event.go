package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Event 
type Event struct {
    OutlookItem
    // true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
    allowNewTimeProposals *bool
    // The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
    attachments []Attachmentable
    // The collection of attendees for the event.
    attendees []Attendeeable
    // The body of the message associated with the event. It can be in HTML or text format.
    body ItemBodyable
    // The preview of the message associated with the event. It is in text format.
    bodyPreview *string
    // The calendar that contains the event. Navigation property. Read-only.
    calendar Calendarable
    // The date, time, and time zone that the event ends. By default, the end time is in UTC.
    end DateTimeTimeZoneable
    // The collection of open extensions defined for the event. Nullable.
    extensions []Extensionable
    // Set to true if the event has attachments.
    hasAttachments *bool
    // When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
    hideAttendees *bool
    // A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. Read-only.
    iCalUId *string
    // The importance property
    importance *Importance
    // The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
    instances []Eventable
    // The isAllDay property
    isAllDay *bool
    // The isCancelled property
    isCancelled *bool
    // The isDraft property
    isDraft *bool
    // The isOnlineMeeting property
    isOnlineMeeting *bool
    // The isOrganizer property
    isOrganizer *bool
    // The isReminderOn property
    isReminderOn *bool
    // The location property
    location Locationable
    // The locations property
    locations []Locationable
    // The collection of multi-value extended properties defined for the event. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The onlineMeeting property
    onlineMeeting OnlineMeetingInfoable
    // The onlineMeetingProvider property
    onlineMeetingProvider *OnlineMeetingProviderType
    // The onlineMeetingUrl property
    onlineMeetingUrl *string
    // The organizer property
    organizer Recipientable
    // The originalEndTimeZone property
    originalEndTimeZone *string
    // The originalStart property
    originalStart *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The originalStartTimeZone property
    originalStartTimeZone *string
    // The recurrence property
    recurrence PatternedRecurrenceable
    // The reminderMinutesBeforeStart property
    reminderMinutesBeforeStart *int32
    // The responseRequested property
    responseRequested *bool
    // The responseStatus property
    responseStatus ResponseStatusable
    // The sensitivity property
    sensitivity *Sensitivity
    // The seriesMasterId property
    seriesMasterId *string
    // The showAs property
    showAs *FreeBusyStatus
    // The collection of single-value extended properties defined for the event. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
    // The start property
    start DateTimeTimeZoneable
    // The subject property
    subject *string
    // The transactionId property
    transactionId *string
    // The type property
    typeEscaped *EventType
    // The webLink property
    webLink *string
}
// NewEvent instantiates a new Event and sets the default values.
func NewEvent()(*Event) {
    m := &Event{
        OutlookItem: *NewOutlookItem(),
    }
    odataTypeValue := "#microsoft.graph.event"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEventFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEventFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEvent(), nil
}
// GetAllowNewTimeProposals gets the allowNewTimeProposals property value. true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
func (m *Event) GetAllowNewTimeProposals()(*bool) {
    return m.allowNewTimeProposals
}
// GetAttachments gets the attachments property value. The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (m *Event) GetAttachments()([]Attachmentable) {
    return m.attachments
}
// GetAttendees gets the attendees property value. The collection of attendees for the event.
func (m *Event) GetAttendees()([]Attendeeable) {
    return m.attendees
}
// GetBody gets the body property value. The body of the message associated with the event. It can be in HTML or text format.
func (m *Event) GetBody()(ItemBodyable) {
    return m.body
}
// GetBodyPreview gets the bodyPreview property value. The preview of the message associated with the event. It is in text format.
func (m *Event) GetBodyPreview()(*string) {
    return m.bodyPreview
}
// GetCalendar gets the calendar property value. The calendar that contains the event. Navigation property. Read-only.
func (m *Event) GetCalendar()(Calendarable) {
    return m.calendar
}
// GetEnd gets the end property value. The date, time, and time zone that the event ends. By default, the end time is in UTC.
func (m *Event) GetEnd()(DateTimeTimeZoneable) {
    return m.end
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the event. Nullable.
func (m *Event) GetExtensions()([]Extensionable) {
    return m.extensions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Event) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["allowNewTimeProposals"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowNewTimeProposals(val)
        }
        return nil
    }
    res["attachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttachmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Attachmentable, len(val))
            for i, v := range val {
                res[i] = v.(Attachmentable)
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["attendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttendeeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Attendeeable, len(val))
            for i, v := range val {
                res[i] = v.(Attendeeable)
            }
            m.SetAttendees(res)
        }
        return nil
    }
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(ItemBodyable))
        }
        return nil
    }
    res["bodyPreview"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyPreview(val)
        }
        return nil
    }
    res["calendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCalendarFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendar(val.(Calendarable))
        }
        return nil
    }
    res["end"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnd(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["hasAttachments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAttachments(val)
        }
        return nil
    }
    res["hideAttendees"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHideAttendees(val)
        }
        return nil
    }
    res["iCalUId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetICalUId(val)
        }
        return nil
    }
    res["importance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*Importance))
        }
        return nil
    }
    res["instances"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                res[i] = v.(Eventable)
            }
            m.SetInstances(res)
        }
        return nil
    }
    res["isAllDay"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAllDay(val)
        }
        return nil
    }
    res["isCancelled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsCancelled(val)
        }
        return nil
    }
    res["isDraft"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDraft(val)
        }
        return nil
    }
    res["isOnlineMeeting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOnlineMeeting(val)
        }
        return nil
    }
    res["isOrganizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizer(val)
        }
        return nil
    }
    res["isReminderOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReminderOn(val)
        }
        return nil
    }
    res["location"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocation(val.(Locationable))
        }
        return nil
    }
    res["locations"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Locationable, len(val))
            for i, v := range val {
                res[i] = v.(Locationable)
            }
            m.SetLocations(res)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(MultiValueLegacyExtendedPropertyable)
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["onlineMeeting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOnlineMeetingInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineMeeting(val.(OnlineMeetingInfoable))
        }
        return nil
    }
    res["onlineMeetingProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineMeetingProvider(val.(*OnlineMeetingProviderType))
        }
        return nil
    }
    res["onlineMeetingUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnlineMeetingUrl(val)
        }
        return nil
    }
    res["organizer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizer(val.(Recipientable))
        }
        return nil
    }
    res["originalEndTimeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalEndTimeZone(val)
        }
        return nil
    }
    res["originalStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalStart(val)
        }
        return nil
    }
    res["originalStartTimeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalStartTimeZone(val)
        }
        return nil
    }
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["reminderMinutesBeforeStart"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderMinutesBeforeStart(val)
        }
        return nil
    }
    res["responseRequested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseRequested(val)
        }
        return nil
    }
    res["responseStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateResponseStatusFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetResponseStatus(val.(ResponseStatusable))
        }
        return nil
    }
    res["sensitivity"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivity)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSensitivity(val.(*Sensitivity))
        }
        return nil
    }
    res["seriesMasterId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSeriesMasterId(val)
        }
        return nil
    }
    res["showAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowAs(val.(*FreeBusyStatus))
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(SingleValueLegacyExtendedPropertyable)
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["start"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStart(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["subject"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["transactionId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransactionId(val)
        }
        return nil
    }
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseEventType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val.(*EventType))
        }
        return nil
    }
    res["webLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebLink(val)
        }
        return nil
    }
    return res
}
// GetHasAttachments gets the hasAttachments property value. Set to true if the event has attachments.
func (m *Event) GetHasAttachments()(*bool) {
    return m.hasAttachments
}
// GetHideAttendees gets the hideAttendees property value. When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
func (m *Event) GetHideAttendees()(*bool) {
    return m.hideAttendees
}
// GetICalUId gets the iCalUId property value. A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. Read-only.
func (m *Event) GetICalUId()(*string) {
    return m.iCalUId
}
// GetImportance gets the importance property value. The importance property
func (m *Event) GetImportance()(*Importance) {
    return m.importance
}
// GetInstances gets the instances property value. The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *Event) GetInstances()([]Eventable) {
    return m.instances
}
// GetIsAllDay gets the isAllDay property value. The isAllDay property
func (m *Event) GetIsAllDay()(*bool) {
    return m.isAllDay
}
// GetIsCancelled gets the isCancelled property value. The isCancelled property
func (m *Event) GetIsCancelled()(*bool) {
    return m.isCancelled
}
// GetIsDraft gets the isDraft property value. The isDraft property
func (m *Event) GetIsDraft()(*bool) {
    return m.isDraft
}
// GetIsOnlineMeeting gets the isOnlineMeeting property value. The isOnlineMeeting property
func (m *Event) GetIsOnlineMeeting()(*bool) {
    return m.isOnlineMeeting
}
// GetIsOrganizer gets the isOrganizer property value. The isOrganizer property
func (m *Event) GetIsOrganizer()(*bool) {
    return m.isOrganizer
}
// GetIsReminderOn gets the isReminderOn property value. The isReminderOn property
func (m *Event) GetIsReminderOn()(*bool) {
    return m.isReminderOn
}
// GetLocation gets the location property value. The location property
func (m *Event) GetLocation()(Locationable) {
    return m.location
}
// GetLocations gets the locations property value. The locations property
func (m *Event) GetLocations()([]Locationable) {
    return m.locations
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the event. Read-only. Nullable.
func (m *Event) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetOnlineMeeting gets the onlineMeeting property value. The onlineMeeting property
func (m *Event) GetOnlineMeeting()(OnlineMeetingInfoable) {
    return m.onlineMeeting
}
// GetOnlineMeetingProvider gets the onlineMeetingProvider property value. The onlineMeetingProvider property
func (m *Event) GetOnlineMeetingProvider()(*OnlineMeetingProviderType) {
    return m.onlineMeetingProvider
}
// GetOnlineMeetingUrl gets the onlineMeetingUrl property value. The onlineMeetingUrl property
func (m *Event) GetOnlineMeetingUrl()(*string) {
    return m.onlineMeetingUrl
}
// GetOrganizer gets the organizer property value. The organizer property
func (m *Event) GetOrganizer()(Recipientable) {
    return m.organizer
}
// GetOriginalEndTimeZone gets the originalEndTimeZone property value. The originalEndTimeZone property
func (m *Event) GetOriginalEndTimeZone()(*string) {
    return m.originalEndTimeZone
}
// GetOriginalStart gets the originalStart property value. The originalStart property
func (m *Event) GetOriginalStart()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.originalStart
}
// GetOriginalStartTimeZone gets the originalStartTimeZone property value. The originalStartTimeZone property
func (m *Event) GetOriginalStartTimeZone()(*string) {
    return m.originalStartTimeZone
}
// GetRecurrence gets the recurrence property value. The recurrence property
func (m *Event) GetRecurrence()(PatternedRecurrenceable) {
    return m.recurrence
}
// GetReminderMinutesBeforeStart gets the reminderMinutesBeforeStart property value. The reminderMinutesBeforeStart property
func (m *Event) GetReminderMinutesBeforeStart()(*int32) {
    return m.reminderMinutesBeforeStart
}
// GetResponseRequested gets the responseRequested property value. The responseRequested property
func (m *Event) GetResponseRequested()(*bool) {
    return m.responseRequested
}
// GetResponseStatus gets the responseStatus property value. The responseStatus property
func (m *Event) GetResponseStatus()(ResponseStatusable) {
    return m.responseStatus
}
// GetSensitivity gets the sensitivity property value. The sensitivity property
func (m *Event) GetSensitivity()(*Sensitivity) {
    return m.sensitivity
}
// GetSeriesMasterId gets the seriesMasterId property value. The seriesMasterId property
func (m *Event) GetSeriesMasterId()(*string) {
    return m.seriesMasterId
}
// GetShowAs gets the showAs property value. The showAs property
func (m *Event) GetShowAs()(*FreeBusyStatus) {
    return m.showAs
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the event. Read-only. Nullable.
func (m *Event) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// GetStart gets the start property value. The start property
func (m *Event) GetStart()(DateTimeTimeZoneable) {
    return m.start
}
// GetSubject gets the subject property value. The subject property
func (m *Event) GetSubject()(*string) {
    return m.subject
}
// GetTransactionId gets the transactionId property value. The transactionId property
func (m *Event) GetTransactionId()(*string) {
    return m.transactionId
}
// GetType gets the type property value. The type property
func (m *Event) GetType()(*EventType) {
    return m.typeEscaped
}
// GetWebLink gets the webLink property value. The webLink property
func (m *Event) GetWebLink()(*string) {
    return m.webLink
}
// Serialize serializes information the current object
func (m *Event) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowNewTimeProposals", m.GetAllowNewTimeProposals())
        if err != nil {
            return err
        }
    }
    if m.GetAttachments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAttendees() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendees()))
        for i, v := range m.GetAttendees() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attendees", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bodyPreview", m.GetBodyPreview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("calendar", m.GetCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("end", m.GetEnd())
        if err != nil {
            return err
        }
    }
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hideAttendees", m.GetHideAttendees())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("iCalUId", m.GetICalUId())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInstances() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInstances()))
        for i, v := range m.GetInstances() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("instances", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAllDay", m.GetIsAllDay())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isCancelled", m.GetIsCancelled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDraft", m.GetIsDraft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOnlineMeeting", m.GetIsOnlineMeeting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizer", m.GetIsOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReminderOn", m.GetIsReminderOn())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("location", m.GetLocation())
        if err != nil {
            return err
        }
    }
    if m.GetLocations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocations()))
        for i, v := range m.GetLocations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("locations", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("onlineMeeting", m.GetOnlineMeeting())
        if err != nil {
            return err
        }
    }
    if m.GetOnlineMeetingProvider() != nil {
        cast := (*m.GetOnlineMeetingProvider()).String()
        err = writer.WriteStringValue("onlineMeetingProvider", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("onlineMeetingUrl", m.GetOnlineMeetingUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("organizer", m.GetOrganizer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalEndTimeZone", m.GetOriginalEndTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("originalStart", m.GetOriginalStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("originalStartTimeZone", m.GetOriginalStartTimeZone())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("reminderMinutesBeforeStart", m.GetReminderMinutesBeforeStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("responseRequested", m.GetResponseRequested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("responseStatus", m.GetResponseStatus())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivity() != nil {
        cast := (*m.GetSensitivity()).String()
        err = writer.WriteStringValue("sensitivity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("seriesMasterId", m.GetSeriesMasterId())
        if err != nil {
            return err
        }
    }
    if m.GetShowAs() != nil {
        cast := (*m.GetShowAs()).String()
        err = writer.WriteStringValue("showAs", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("start", m.GetStart())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("transactionId", m.GetTransactionId())
        if err != nil {
            return err
        }
    }
    if m.GetType() != nil {
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webLink", m.GetWebLink())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowNewTimeProposals sets the allowNewTimeProposals property value. true if the meeting organizer allows invitees to propose a new time when responding; otherwise, false. Optional. Default is true.
func (m *Event) SetAllowNewTimeProposals(value *bool)() {
    m.allowNewTimeProposals = value
}
// SetAttachments sets the attachments property value. The collection of FileAttachment, ItemAttachment, and referenceAttachment attachments for the event. Navigation property. Read-only. Nullable.
func (m *Event) SetAttachments(value []Attachmentable)() {
    m.attachments = value
}
// SetAttendees sets the attendees property value. The collection of attendees for the event.
func (m *Event) SetAttendees(value []Attendeeable)() {
    m.attendees = value
}
// SetBody sets the body property value. The body of the message associated with the event. It can be in HTML or text format.
func (m *Event) SetBody(value ItemBodyable)() {
    m.body = value
}
// SetBodyPreview sets the bodyPreview property value. The preview of the message associated with the event. It is in text format.
func (m *Event) SetBodyPreview(value *string)() {
    m.bodyPreview = value
}
// SetCalendar sets the calendar property value. The calendar that contains the event. Navigation property. Read-only.
func (m *Event) SetCalendar(value Calendarable)() {
    m.calendar = value
}
// SetEnd sets the end property value. The date, time, and time zone that the event ends. By default, the end time is in UTC.
func (m *Event) SetEnd(value DateTimeTimeZoneable)() {
    m.end = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the event. Nullable.
func (m *Event) SetExtensions(value []Extensionable)() {
    m.extensions = value
}
// SetHasAttachments sets the hasAttachments property value. Set to true if the event has attachments.
func (m *Event) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetHideAttendees sets the hideAttendees property value. When set to true, each attendee only sees themselves in the meeting request and meeting Tracking list. Default is false.
func (m *Event) SetHideAttendees(value *bool)() {
    m.hideAttendees = value
}
// SetICalUId sets the iCalUId property value. A unique identifier for an event across calendars. This ID is different for each occurrence in a recurring series. Read-only.
func (m *Event) SetICalUId(value *string)() {
    m.iCalUId = value
}
// SetImportance sets the importance property value. The importance property
func (m *Event) SetImportance(value *Importance)() {
    m.importance = value
}
// SetInstances sets the instances property value. The occurrences of a recurring series, if the event is a series master. This property includes occurrences that are part of the recurrence pattern, and exceptions that have been modified, but does not include occurrences that have been cancelled from the series. Navigation property. Read-only. Nullable.
func (m *Event) SetInstances(value []Eventable)() {
    m.instances = value
}
// SetIsAllDay sets the isAllDay property value. The isAllDay property
func (m *Event) SetIsAllDay(value *bool)() {
    m.isAllDay = value
}
// SetIsCancelled sets the isCancelled property value. The isCancelled property
func (m *Event) SetIsCancelled(value *bool)() {
    m.isCancelled = value
}
// SetIsDraft sets the isDraft property value. The isDraft property
func (m *Event) SetIsDraft(value *bool)() {
    m.isDraft = value
}
// SetIsOnlineMeeting sets the isOnlineMeeting property value. The isOnlineMeeting property
func (m *Event) SetIsOnlineMeeting(value *bool)() {
    m.isOnlineMeeting = value
}
// SetIsOrganizer sets the isOrganizer property value. The isOrganizer property
func (m *Event) SetIsOrganizer(value *bool)() {
    m.isOrganizer = value
}
// SetIsReminderOn sets the isReminderOn property value. The isReminderOn property
func (m *Event) SetIsReminderOn(value *bool)() {
    m.isReminderOn = value
}
// SetLocation sets the location property value. The location property
func (m *Event) SetLocation(value Locationable)() {
    m.location = value
}
// SetLocations sets the locations property value. The locations property
func (m *Event) SetLocations(value []Locationable)() {
    m.locations = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the event. Read-only. Nullable.
func (m *Event) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetOnlineMeeting sets the onlineMeeting property value. The onlineMeeting property
func (m *Event) SetOnlineMeeting(value OnlineMeetingInfoable)() {
    m.onlineMeeting = value
}
// SetOnlineMeetingProvider sets the onlineMeetingProvider property value. The onlineMeetingProvider property
func (m *Event) SetOnlineMeetingProvider(value *OnlineMeetingProviderType)() {
    m.onlineMeetingProvider = value
}
// SetOnlineMeetingUrl sets the onlineMeetingUrl property value. The onlineMeetingUrl property
func (m *Event) SetOnlineMeetingUrl(value *string)() {
    m.onlineMeetingUrl = value
}
// SetOrganizer sets the organizer property value. The organizer property
func (m *Event) SetOrganizer(value Recipientable)() {
    m.organizer = value
}
// SetOriginalEndTimeZone sets the originalEndTimeZone property value. The originalEndTimeZone property
func (m *Event) SetOriginalEndTimeZone(value *string)() {
    m.originalEndTimeZone = value
}
// SetOriginalStart sets the originalStart property value. The originalStart property
func (m *Event) SetOriginalStart(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.originalStart = value
}
// SetOriginalStartTimeZone sets the originalStartTimeZone property value. The originalStartTimeZone property
func (m *Event) SetOriginalStartTimeZone(value *string)() {
    m.originalStartTimeZone = value
}
// SetRecurrence sets the recurrence property value. The recurrence property
func (m *Event) SetRecurrence(value PatternedRecurrenceable)() {
    m.recurrence = value
}
// SetReminderMinutesBeforeStart sets the reminderMinutesBeforeStart property value. The reminderMinutesBeforeStart property
func (m *Event) SetReminderMinutesBeforeStart(value *int32)() {
    m.reminderMinutesBeforeStart = value
}
// SetResponseRequested sets the responseRequested property value. The responseRequested property
func (m *Event) SetResponseRequested(value *bool)() {
    m.responseRequested = value
}
// SetResponseStatus sets the responseStatus property value. The responseStatus property
func (m *Event) SetResponseStatus(value ResponseStatusable)() {
    m.responseStatus = value
}
// SetSensitivity sets the sensitivity property value. The sensitivity property
func (m *Event) SetSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
// SetSeriesMasterId sets the seriesMasterId property value. The seriesMasterId property
func (m *Event) SetSeriesMasterId(value *string)() {
    m.seriesMasterId = value
}
// SetShowAs sets the showAs property value. The showAs property
func (m *Event) SetShowAs(value *FreeBusyStatus)() {
    m.showAs = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the event. Read-only. Nullable.
func (m *Event) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
// SetStart sets the start property value. The start property
func (m *Event) SetStart(value DateTimeTimeZoneable)() {
    m.start = value
}
// SetSubject sets the subject property value. The subject property
func (m *Event) SetSubject(value *string)() {
    m.subject = value
}
// SetTransactionId sets the transactionId property value. The transactionId property
func (m *Event) SetTransactionId(value *string)() {
    m.transactionId = value
}
// SetType sets the type property value. The type property
func (m *Event) SetType(value *EventType)() {
    m.typeEscaped = value
}
// SetWebLink sets the webLink property value. The webLink property
func (m *Event) SetWebLink(value *string)() {
    m.webLink = value
}
