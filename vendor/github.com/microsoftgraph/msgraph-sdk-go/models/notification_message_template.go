package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// NotificationMessageTemplate notification messages are messages that are sent to end users who are determined to be not-compliant with the compliance policies defined by the administrator. Administrators choose notifications and configure them in the Intune Admin Console using the compliance policy creation page under the “Actions for non-compliance” section. Use the notificationMessageTemplate object to create your own custom notifications for administrators to choose while configuring actions for non-compliance.
type NotificationMessageTemplate struct {
    Entity
    // Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
    brandingOptions *NotificationTemplateBrandingOptions
    // The default locale to fallback onto when the requested locale is not available.
    defaultLocale *string
    // Display name for the Notification Message Template.
    displayName *string
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The list of localized messages for this Notification Message Template.
    localizedNotificationMessages []LocalizedNotificationMessageable
    // List of Scope Tags for this Entity instance.
    roleScopeTagIds []string
}
// NewNotificationMessageTemplate instantiates a new notificationMessageTemplate and sets the default values.
func NewNotificationMessageTemplate()(*NotificationMessageTemplate) {
    m := &NotificationMessageTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateNotificationMessageTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateNotificationMessageTemplateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationMessageTemplate(), nil
}
// GetBrandingOptions gets the brandingOptions property value. Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
func (m *NotificationMessageTemplate) GetBrandingOptions()(*NotificationTemplateBrandingOptions) {
    return m.brandingOptions
}
// GetDefaultLocale gets the defaultLocale property value. The default locale to fallback onto when the requested locale is not available.
func (m *NotificationMessageTemplate) GetDefaultLocale()(*string) {
    return m.defaultLocale
}
// GetDisplayName gets the displayName property value. Display name for the Notification Message Template.
func (m *NotificationMessageTemplate) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NotificationMessageTemplate) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["brandingOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseNotificationTemplateBrandingOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBrandingOptions(val.(*NotificationTemplateBrandingOptions))
        }
        return nil
    }
    res["defaultLocale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultLocale(val)
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
    res["localizedNotificationMessages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocalizedNotificationMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LocalizedNotificationMessageable, len(val))
            for i, v := range val {
                res[i] = v.(LocalizedNotificationMessageable)
            }
            m.SetLocalizedNotificationMessages(res)
        }
        return nil
    }
    res["roleScopeTagIds"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRoleScopeTagIds(res)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *NotificationMessageTemplate) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLocalizedNotificationMessages gets the localizedNotificationMessages property value. The list of localized messages for this Notification Message Template.
func (m *NotificationMessageTemplate) GetLocalizedNotificationMessages()([]LocalizedNotificationMessageable) {
    return m.localizedNotificationMessages
}
// GetRoleScopeTagIds gets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *NotificationMessageTemplate) GetRoleScopeTagIds()([]string) {
    return m.roleScopeTagIds
}
// Serialize serializes information the current object
func (m *NotificationMessageTemplate) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetBrandingOptions() != nil {
        cast := (*m.GetBrandingOptions()).String()
        err = writer.WriteStringValue("brandingOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("defaultLocale", m.GetDefaultLocale())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetLocalizedNotificationMessages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocalizedNotificationMessages()))
        for i, v := range m.GetLocalizedNotificationMessages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("localizedNotificationMessages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRoleScopeTagIds() != nil {
        err = writer.WriteCollectionOfStringValues("roleScopeTagIds", m.GetRoleScopeTagIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBrandingOptions sets the brandingOptions property value. Branding Options for the Message Template. Branding is defined in the Intune Admin Console.
func (m *NotificationMessageTemplate) SetBrandingOptions(value *NotificationTemplateBrandingOptions)() {
    m.brandingOptions = value
}
// SetDefaultLocale sets the defaultLocale property value. The default locale to fallback onto when the requested locale is not available.
func (m *NotificationMessageTemplate) SetDefaultLocale(value *string)() {
    m.defaultLocale = value
}
// SetDisplayName sets the displayName property value. Display name for the Notification Message Template.
func (m *NotificationMessageTemplate) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *NotificationMessageTemplate) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLocalizedNotificationMessages sets the localizedNotificationMessages property value. The list of localized messages for this Notification Message Template.
func (m *NotificationMessageTemplate) SetLocalizedNotificationMessages(value []LocalizedNotificationMessageable)() {
    m.localizedNotificationMessages = value
}
// SetRoleScopeTagIds sets the roleScopeTagIds property value. List of Scope Tags for this Entity instance.
func (m *NotificationMessageTemplate) SetRoleScopeTagIds(value []string)() {
    m.roleScopeTagIds = value
}
