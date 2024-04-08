package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailboxSettings 
type MailboxSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Folder ID of an archive folder for the user.
    archiveFolder *string
    // Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
    automaticRepliesSetting AutomaticRepliesSettingable
    // The date format for the user's mailbox.
    dateFormat *string
    // If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly.
    delegateMeetingMessageDeliveryOptions *DelegateMeetingMessageDeliveryOptions
    // The locale information for the user, including the preferred language and country/region.
    language LocaleInfoable
    // The OdataType property
    odataType *string
    // The time format for the user's mailbox.
    timeFormat *string
    // The default time zone for the user's mailbox.
    timeZone *string
    // The userPurpose property
    userPurpose *UserPurpose
    // The days of the week and hours in a specific time zone that the user works.
    workingHours WorkingHoursable
}
// NewMailboxSettings instantiates a new mailboxSettings and sets the default values.
func NewMailboxSettings()(*MailboxSettings) {
    m := &MailboxSettings{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMailboxSettingsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailboxSettingsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailboxSettings(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailboxSettings) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetArchiveFolder gets the archiveFolder property value. Folder ID of an archive folder for the user.
func (m *MailboxSettings) GetArchiveFolder()(*string) {
    return m.archiveFolder
}
// GetAutomaticRepliesSetting gets the automaticRepliesSetting property value. Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
func (m *MailboxSettings) GetAutomaticRepliesSetting()(AutomaticRepliesSettingable) {
    return m.automaticRepliesSetting
}
// GetDateFormat gets the dateFormat property value. The date format for the user's mailbox.
func (m *MailboxSettings) GetDateFormat()(*string) {
    return m.dateFormat
}
// GetDelegateMeetingMessageDeliveryOptions gets the delegateMeetingMessageDeliveryOptions property value. If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly.
func (m *MailboxSettings) GetDelegateMeetingMessageDeliveryOptions()(*DelegateMeetingMessageDeliveryOptions) {
    return m.delegateMeetingMessageDeliveryOptions
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailboxSettings) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["archiveFolder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArchiveFolder(val)
        }
        return nil
    }
    res["automaticRepliesSetting"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutomaticRepliesSettingFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticRepliesSetting(val.(AutomaticRepliesSettingable))
        }
        return nil
    }
    res["dateFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDateFormat(val)
        }
        return nil
    }
    res["delegateMeetingMessageDeliveryOptions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDelegateMeetingMessageDeliveryOptions)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelegateMeetingMessageDeliveryOptions(val.(*DelegateMeetingMessageDeliveryOptions))
        }
        return nil
    }
    res["language"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateLocaleInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLanguage(val.(LocaleInfoable))
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["timeFormat"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeFormat(val)
        }
        return nil
    }
    res["timeZone"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["userPurpose"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserPurpose)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPurpose(val.(*UserPurpose))
        }
        return nil
    }
    res["workingHours"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWorkingHoursFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWorkingHours(val.(WorkingHoursable))
        }
        return nil
    }
    return res
}
// GetLanguage gets the language property value. The locale information for the user, including the preferred language and country/region.
func (m *MailboxSettings) GetLanguage()(LocaleInfoable) {
    return m.language
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MailboxSettings) GetOdataType()(*string) {
    return m.odataType
}
// GetTimeFormat gets the timeFormat property value. The time format for the user's mailbox.
func (m *MailboxSettings) GetTimeFormat()(*string) {
    return m.timeFormat
}
// GetTimeZone gets the timeZone property value. The default time zone for the user's mailbox.
func (m *MailboxSettings) GetTimeZone()(*string) {
    return m.timeZone
}
// GetUserPurpose gets the userPurpose property value. The userPurpose property
func (m *MailboxSettings) GetUserPurpose()(*UserPurpose) {
    return m.userPurpose
}
// GetWorkingHours gets the workingHours property value. The days of the week and hours in a specific time zone that the user works.
func (m *MailboxSettings) GetWorkingHours()(WorkingHoursable) {
    return m.workingHours
}
// Serialize serializes information the current object
func (m *MailboxSettings) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("archiveFolder", m.GetArchiveFolder())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("automaticRepliesSetting", m.GetAutomaticRepliesSetting())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dateFormat", m.GetDateFormat())
        if err != nil {
            return err
        }
    }
    if m.GetDelegateMeetingMessageDeliveryOptions() != nil {
        cast := (*m.GetDelegateMeetingMessageDeliveryOptions()).String()
        err := writer.WriteStringValue("delegateMeetingMessageDeliveryOptions", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("language", m.GetLanguage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeFormat", m.GetTimeFormat())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    if m.GetUserPurpose() != nil {
        cast := (*m.GetUserPurpose()).String()
        err := writer.WriteStringValue("userPurpose", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("workingHours", m.GetWorkingHours())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailboxSettings) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetArchiveFolder sets the archiveFolder property value. Folder ID of an archive folder for the user.
func (m *MailboxSettings) SetArchiveFolder(value *string)() {
    m.archiveFolder = value
}
// SetAutomaticRepliesSetting sets the automaticRepliesSetting property value. Configuration settings to automatically notify the sender of an incoming email with a message from the signed-in user.
func (m *MailboxSettings) SetAutomaticRepliesSetting(value AutomaticRepliesSettingable)() {
    m.automaticRepliesSetting = value
}
// SetDateFormat sets the dateFormat property value. The date format for the user's mailbox.
func (m *MailboxSettings) SetDateFormat(value *string)() {
    m.dateFormat = value
}
// SetDelegateMeetingMessageDeliveryOptions sets the delegateMeetingMessageDeliveryOptions property value. If the user has a calendar delegate, this specifies whether the delegate, mailbox owner, or both receive meeting messages and meeting responses. Possible values are: sendToDelegateAndInformationToPrincipal, sendToDelegateAndPrincipal, sendToDelegateOnly.
func (m *MailboxSettings) SetDelegateMeetingMessageDeliveryOptions(value *DelegateMeetingMessageDeliveryOptions)() {
    m.delegateMeetingMessageDeliveryOptions = value
}
// SetLanguage sets the language property value. The locale information for the user, including the preferred language and country/region.
func (m *MailboxSettings) SetLanguage(value LocaleInfoable)() {
    m.language = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MailboxSettings) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTimeFormat sets the timeFormat property value. The time format for the user's mailbox.
func (m *MailboxSettings) SetTimeFormat(value *string)() {
    m.timeFormat = value
}
// SetTimeZone sets the timeZone property value. The default time zone for the user's mailbox.
func (m *MailboxSettings) SetTimeZone(value *string)() {
    m.timeZone = value
}
// SetUserPurpose sets the userPurpose property value. The userPurpose property
func (m *MailboxSettings) SetUserPurpose(value *UserPurpose)() {
    m.userPurpose = value
}
// SetWorkingHours sets the workingHours property value. The days of the week and hours in a specific time zone that the user works.
func (m *MailboxSettings) SetWorkingHours(value WorkingHoursable)() {
    m.workingHours = value
}
