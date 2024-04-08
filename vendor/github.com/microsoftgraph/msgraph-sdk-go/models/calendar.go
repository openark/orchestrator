package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Calendar 
type Calendar struct {
    Entity
    // Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
    allowedOnlineMeetingProviders []OnlineMeetingProviderType
    // The permissions of the users with whom the calendar is shared.
    calendarPermissions []CalendarPermissionable
    // The calendar view for the calendar. Navigation property. Read-only.
    calendarView []Eventable
    // true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access.
    canEdit *bool
    // true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.
    canShare *bool
    // true if the user can read calendar items that have been marked private, false otherwise.
    canViewPrivateItems *bool
    // Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
    changeKey *string
    // Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
    color *CalendarColor
    // The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
    defaultOnlineMeetingProvider *OnlineMeetingProviderType
    // The events in the calendar. Navigation property. Read-only.
    events []Eventable
    // The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only.
    hexColor *string
    // true if this is the default calendar where new events are created by default, false otherwise.
    isDefaultCalendar *bool
    // Indicates whether this user calendar can be deleted from the user mailbox.
    isRemovable *bool
    // Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
    isTallyingResponses *bool
    // The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable
    // The calendar name.
    name *string
    // If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user.
    owner EmailAddressable
    // The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable
}
// NewCalendar instantiates a new calendar and sets the default values.
func NewCalendar()(*Calendar) {
    m := &Calendar{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCalendarFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalendarFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalendar(), nil
}
// GetAllowedOnlineMeetingProviders gets the allowedOnlineMeetingProviders property value. Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) GetAllowedOnlineMeetingProviders()([]OnlineMeetingProviderType) {
    return m.allowedOnlineMeetingProviders
}
// GetCalendarPermissions gets the calendarPermissions property value. The permissions of the users with whom the calendar is shared.
func (m *Calendar) GetCalendarPermissions()([]CalendarPermissionable) {
    return m.calendarPermissions
}
// GetCalendarView gets the calendarView property value. The calendar view for the calendar. Navigation property. Read-only.
func (m *Calendar) GetCalendarView()([]Eventable) {
    return m.calendarView
}
// GetCanEdit gets the canEdit property value. true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access.
func (m *Calendar) GetCanEdit()(*bool) {
    return m.canEdit
}
// GetCanShare gets the canShare property value. true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.
func (m *Calendar) GetCanShare()(*bool) {
    return m.canShare
}
// GetCanViewPrivateItems gets the canViewPrivateItems property value. true if the user can read calendar items that have been marked private, false otherwise.
func (m *Calendar) GetCanViewPrivateItems()(*bool) {
    return m.canViewPrivateItems
}
// GetChangeKey gets the changeKey property value. Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
func (m *Calendar) GetChangeKey()(*string) {
    return m.changeKey
}
// GetColor gets the color property value. Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
func (m *Calendar) GetColor()(*CalendarColor) {
    return m.color
}
// GetDefaultOnlineMeetingProvider gets the defaultOnlineMeetingProvider property value. The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) GetDefaultOnlineMeetingProvider()(*OnlineMeetingProviderType) {
    return m.defaultOnlineMeetingProvider
}
// GetEvents gets the events property value. The events in the calendar. Navigation property. Read-only.
func (m *Calendar) GetEvents()([]Eventable) {
    return m.events
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Calendar) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedOnlineMeetingProviders"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseOnlineMeetingProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OnlineMeetingProviderType, len(val))
            for i, v := range val {
                res[i] = *(v.(*OnlineMeetingProviderType))
            }
            m.SetAllowedOnlineMeetingProviders(res)
        }
        return nil
    }
    res["calendarPermissions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateCalendarPermissionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CalendarPermissionable, len(val))
            for i, v := range val {
                res[i] = v.(CalendarPermissionable)
            }
            m.SetCalendarPermissions(res)
        }
        return nil
    }
    res["calendarView"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                res[i] = v.(Eventable)
            }
            m.SetCalendarView(res)
        }
        return nil
    }
    res["canEdit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanEdit(val)
        }
        return nil
    }
    res["canShare"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanShare(val)
        }
        return nil
    }
    res["canViewPrivateItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCanViewPrivateItems(val)
        }
        return nil
    }
    res["changeKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeKey(val)
        }
        return nil
    }
    res["color"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarColor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetColor(val.(*CalendarColor))
        }
        return nil
    }
    res["defaultOnlineMeetingProvider"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingProviderType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultOnlineMeetingProvider(val.(*OnlineMeetingProviderType))
        }
        return nil
    }
    res["events"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateEventFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Eventable, len(val))
            for i, v := range val {
                res[i] = v.(Eventable)
            }
            m.SetEvents(res)
        }
        return nil
    }
    res["hexColor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHexColor(val)
        }
        return nil
    }
    res["isDefaultCalendar"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultCalendar(val)
        }
        return nil
    }
    res["isRemovable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRemovable(val)
        }
        return nil
    }
    res["isTallyingResponses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsTallyingResponses(val)
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
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(EmailAddressable))
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
    return res
}
// GetHexColor gets the hexColor property value. The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only.
func (m *Calendar) GetHexColor()(*string) {
    return m.hexColor
}
// GetIsDefaultCalendar gets the isDefaultCalendar property value. true if this is the default calendar where new events are created by default, false otherwise.
func (m *Calendar) GetIsDefaultCalendar()(*bool) {
    return m.isDefaultCalendar
}
// GetIsRemovable gets the isRemovable property value. Indicates whether this user calendar can be deleted from the user mailbox.
func (m *Calendar) GetIsRemovable()(*bool) {
    return m.isRemovable
}
// GetIsTallyingResponses gets the isTallyingResponses property value. Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
func (m *Calendar) GetIsTallyingResponses()(*bool) {
    return m.isTallyingResponses
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    return m.multiValueExtendedProperties
}
// GetName gets the name property value. The calendar name.
func (m *Calendar) GetName()(*string) {
    return m.name
}
// GetOwner gets the owner property value. If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user.
func (m *Calendar) GetOwner()(EmailAddressable) {
    return m.owner
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    return m.singleValueExtendedProperties
}
// Serialize serializes information the current object
func (m *Calendar) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedOnlineMeetingProviders() != nil {
        err = writer.WriteCollectionOfStringValues("allowedOnlineMeetingProviders", SerializeOnlineMeetingProviderType(m.GetAllowedOnlineMeetingProviders()))
        if err != nil {
            return err
        }
    }
    if m.GetCalendarPermissions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalendarPermissions()))
        for i, v := range m.GetCalendarPermissions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("calendarPermissions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCalendarView() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCalendarView()))
        for i, v := range m.GetCalendarView() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("calendarView", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("canEdit", m.GetCanEdit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("canShare", m.GetCanShare())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("canViewPrivateItems", m.GetCanViewPrivateItems())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("changeKey", m.GetChangeKey())
        if err != nil {
            return err
        }
    }
    if m.GetColor() != nil {
        cast := (*m.GetColor()).String()
        err = writer.WriteStringValue("color", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetDefaultOnlineMeetingProvider() != nil {
        cast := (*m.GetDefaultOnlineMeetingProvider()).String()
        err = writer.WriteStringValue("defaultOnlineMeetingProvider", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetEvents() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEvents()))
        for i, v := range m.GetEvents() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("events", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("hexColor", m.GetHexColor())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDefaultCalendar", m.GetIsDefaultCalendar())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRemovable", m.GetIsRemovable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isTallyingResponses", m.GetIsTallyingResponses())
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
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
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
    return nil
}
// SetAllowedOnlineMeetingProviders sets the allowedOnlineMeetingProviders property value. Represent the online meeting service providers that can be used to create online meetings in this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) SetAllowedOnlineMeetingProviders(value []OnlineMeetingProviderType)() {
    m.allowedOnlineMeetingProviders = value
}
// SetCalendarPermissions sets the calendarPermissions property value. The permissions of the users with whom the calendar is shared.
func (m *Calendar) SetCalendarPermissions(value []CalendarPermissionable)() {
    m.calendarPermissions = value
}
// SetCalendarView sets the calendarView property value. The calendar view for the calendar. Navigation property. Read-only.
func (m *Calendar) SetCalendarView(value []Eventable)() {
    m.calendarView = value
}
// SetCanEdit sets the canEdit property value. true if the user can write to the calendar, false otherwise. This property is true for the user who created the calendar. This property is also true for a user who has been shared a calendar and granted write access.
func (m *Calendar) SetCanEdit(value *bool)() {
    m.canEdit = value
}
// SetCanShare sets the canShare property value. true if the user has the permission to share the calendar, false otherwise. Only the user who created the calendar can share it.
func (m *Calendar) SetCanShare(value *bool)() {
    m.canShare = value
}
// SetCanViewPrivateItems sets the canViewPrivateItems property value. true if the user can read calendar items that have been marked private, false otherwise.
func (m *Calendar) SetCanViewPrivateItems(value *bool)() {
    m.canViewPrivateItems = value
}
// SetChangeKey sets the changeKey property value. Identifies the version of the calendar object. Every time the calendar is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
func (m *Calendar) SetChangeKey(value *string)() {
    m.changeKey = value
}
// SetColor sets the color property value. Specifies the color theme to distinguish the calendar from other calendars in a UI. The property values are: auto, lightBlue, lightGreen, lightOrange, lightGray, lightYellow, lightTeal, lightPink, lightBrown, lightRed, maxColor.
func (m *Calendar) SetColor(value *CalendarColor)() {
    m.color = value
}
// SetDefaultOnlineMeetingProvider sets the defaultOnlineMeetingProvider property value. The default online meeting provider for meetings sent from this calendar. Possible values are: unknown, skypeForBusiness, skypeForConsumer, teamsForBusiness.
func (m *Calendar) SetDefaultOnlineMeetingProvider(value *OnlineMeetingProviderType)() {
    m.defaultOnlineMeetingProvider = value
}
// SetEvents sets the events property value. The events in the calendar. Navigation property. Read-only.
func (m *Calendar) SetEvents(value []Eventable)() {
    m.events = value
}
// SetHexColor sets the hexColor property value. The calendar color, expressed in a hex color code of three hexadecimal values, each ranging from 00 to FF and representing the red, green, or blue components of the color in the RGB color space. If the user has never explicitly set a color for the calendar, this property is empty. Read-only.
func (m *Calendar) SetHexColor(value *string)() {
    m.hexColor = value
}
// SetIsDefaultCalendar sets the isDefaultCalendar property value. true if this is the default calendar where new events are created by default, false otherwise.
func (m *Calendar) SetIsDefaultCalendar(value *bool)() {
    m.isDefaultCalendar = value
}
// SetIsRemovable sets the isRemovable property value. Indicates whether this user calendar can be deleted from the user mailbox.
func (m *Calendar) SetIsRemovable(value *bool)() {
    m.isRemovable = value
}
// SetIsTallyingResponses sets the isTallyingResponses property value. Indicates whether this user calendar supports tracking of meeting responses. Only meeting invites sent from users' primary calendars support tracking of meeting responses.
func (m *Calendar) SetIsTallyingResponses(value *bool)() {
    m.isTallyingResponses = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    m.multiValueExtendedProperties = value
}
// SetName sets the name property value. The calendar name.
func (m *Calendar) SetName(value *string)() {
    m.name = value
}
// SetOwner sets the owner property value. If set, this represents the user who created or added the calendar. For a calendar that the user created or added, the owner property is set to the user. For a calendar shared with the user, the owner property is set to the person who shared that calendar with the user.
func (m *Calendar) SetOwner(value EmailAddressable)() {
    m.owner = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the calendar. Read-only. Nullable.
func (m *Calendar) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    m.singleValueExtendedProperties = value
}
