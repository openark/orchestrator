package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// LocalizedNotificationMessage the text content of a Notification Message Template for the specified locale.
type LocalizedNotificationMessage struct {
    Entity
    // Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
    isDefault *bool
    // DateTime the object was last modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Locale for which this message is destined.
    locale *string
    // The Message Template content.
    messageTemplate *string
    // The Message Template Subject.
    subject *string
}
// NewLocalizedNotificationMessage instantiates a new localizedNotificationMessage and sets the default values.
func NewLocalizedNotificationMessage()(*LocalizedNotificationMessage) {
    m := &LocalizedNotificationMessage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateLocalizedNotificationMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateLocalizedNotificationMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLocalizedNotificationMessage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *LocalizedNotificationMessage) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["isDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefault(val)
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
    res["locale"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    res["messageTemplate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageTemplate(val)
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
    return res
}
// GetIsDefault gets the isDefault property value. Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
func (m *LocalizedNotificationMessage) GetIsDefault()(*bool) {
    return m.isDefault
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *LocalizedNotificationMessage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetLocale gets the locale property value. The Locale for which this message is destined.
func (m *LocalizedNotificationMessage) GetLocale()(*string) {
    return m.locale
}
// GetMessageTemplate gets the messageTemplate property value. The Message Template content.
func (m *LocalizedNotificationMessage) GetMessageTemplate()(*string) {
    return m.messageTemplate
}
// GetSubject gets the subject property value. The Message Template Subject.
func (m *LocalizedNotificationMessage) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *LocalizedNotificationMessage) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefault", m.GetIsDefault())
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
    {
        err = writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("messageTemplate", m.GetMessageTemplate())
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
    return nil
}
// SetIsDefault sets the isDefault property value. Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
func (m *LocalizedNotificationMessage) SetIsDefault(value *bool)() {
    m.isDefault = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. DateTime the object was last modified.
func (m *LocalizedNotificationMessage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetLocale sets the locale property value. The Locale for which this message is destined.
func (m *LocalizedNotificationMessage) SetLocale(value *string)() {
    m.locale = value
}
// SetMessageTemplate sets the messageTemplate property value. The Message Template content.
func (m *LocalizedNotificationMessage) SetMessageTemplate(value *string)() {
    m.messageTemplate = value
}
// SetSubject sets the subject property value. The Message Template Subject.
func (m *LocalizedNotificationMessage) SetSubject(value *string)() {
    m.subject = value
}
