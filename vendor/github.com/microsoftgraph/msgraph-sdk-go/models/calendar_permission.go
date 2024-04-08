package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CalendarPermission 
type CalendarPermission struct {
    Entity
    // List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
    allowedRoles []CalendarRoleType
    // Represents a sharee or delegate who has access to the calendar. For the 'My Organization' sharee, the address property is null. Read-only.
    emailAddress EmailAddressable
    // True if the user in context (sharee or delegate) is inside the same organization as the calendar owner.
    isInsideOrganization *bool
    // True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The 'My organization' user determines the permissions other people within your organization have to the given calendar. You cannot remove 'My organization' as a sharee to a calendar.
    isRemovable *bool
    // Current permission level of the calendar sharee or delegate.
    role *CalendarRoleType
}
// NewCalendarPermission instantiates a new calendarPermission and sets the default values.
func NewCalendarPermission()(*CalendarPermission) {
    m := &CalendarPermission{
        Entity: *NewEntity(),
    }
    return m
}
// CreateCalendarPermissionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCalendarPermissionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCalendarPermission(), nil
}
// GetAllowedRoles gets the allowedRoles property value. List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
func (m *CalendarPermission) GetAllowedRoles()([]CalendarRoleType) {
    return m.allowedRoles
}
// GetEmailAddress gets the emailAddress property value. Represents a sharee or delegate who has access to the calendar. For the 'My Organization' sharee, the address property is null. Read-only.
func (m *CalendarPermission) GetEmailAddress()(EmailAddressable) {
    return m.emailAddress
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CalendarPermission) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowedRoles"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseCalendarRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]CalendarRoleType, len(val))
            for i, v := range val {
                res[i] = *(v.(*CalendarRoleType))
            }
            m.SetAllowedRoles(res)
        }
        return nil
    }
    res["emailAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val.(EmailAddressable))
        }
        return nil
    }
    res["isInsideOrganization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsInsideOrganization(val)
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
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseCalendarRoleType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*CalendarRoleType))
        }
        return nil
    }
    return res
}
// GetIsInsideOrganization gets the isInsideOrganization property value. True if the user in context (sharee or delegate) is inside the same organization as the calendar owner.
func (m *CalendarPermission) GetIsInsideOrganization()(*bool) {
    return m.isInsideOrganization
}
// GetIsRemovable gets the isRemovable property value. True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The 'My organization' user determines the permissions other people within your organization have to the given calendar. You cannot remove 'My organization' as a sharee to a calendar.
func (m *CalendarPermission) GetIsRemovable()(*bool) {
    return m.isRemovable
}
// GetRole gets the role property value. Current permission level of the calendar sharee or delegate.
func (m *CalendarPermission) GetRole()(*CalendarRoleType) {
    return m.role
}
// Serialize serializes information the current object
func (m *CalendarPermission) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedRoles() != nil {
        err = writer.WriteCollectionOfStringValues("allowedRoles", SerializeCalendarRoleType(m.GetAllowedRoles()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isInsideOrganization", m.GetIsInsideOrganization())
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
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAllowedRoles sets the allowedRoles property value. List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
func (m *CalendarPermission) SetAllowedRoles(value []CalendarRoleType)() {
    m.allowedRoles = value
}
// SetEmailAddress sets the emailAddress property value. Represents a sharee or delegate who has access to the calendar. For the 'My Organization' sharee, the address property is null. Read-only.
func (m *CalendarPermission) SetEmailAddress(value EmailAddressable)() {
    m.emailAddress = value
}
// SetIsInsideOrganization sets the isInsideOrganization property value. True if the user in context (sharee or delegate) is inside the same organization as the calendar owner.
func (m *CalendarPermission) SetIsInsideOrganization(value *bool)() {
    m.isInsideOrganization = value
}
// SetIsRemovable sets the isRemovable property value. True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The 'My organization' user determines the permissions other people within your organization have to the given calendar. You cannot remove 'My organization' as a sharee to a calendar.
func (m *CalendarPermission) SetIsRemovable(value *bool)() {
    m.isRemovable = value
}
// SetRole sets the role property value. Current permission level of the calendar sharee or delegate.
func (m *CalendarPermission) SetRole(value *CalendarRoleType)() {
    m.role = value
}
