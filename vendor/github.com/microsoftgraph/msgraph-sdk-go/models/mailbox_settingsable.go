package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailboxSettingsable 
type MailboxSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetArchiveFolder()(*string)
    GetAutomaticRepliesSetting()(AutomaticRepliesSettingable)
    GetDateFormat()(*string)
    GetDelegateMeetingMessageDeliveryOptions()(*DelegateMeetingMessageDeliveryOptions)
    GetLanguage()(LocaleInfoable)
    GetOdataType()(*string)
    GetTimeFormat()(*string)
    GetTimeZone()(*string)
    GetUserPurpose()(*UserPurpose)
    GetWorkingHours()(WorkingHoursable)
    SetArchiveFolder(value *string)()
    SetAutomaticRepliesSetting(value AutomaticRepliesSettingable)()
    SetDateFormat(value *string)()
    SetDelegateMeetingMessageDeliveryOptions(value *DelegateMeetingMessageDeliveryOptions)()
    SetLanguage(value LocaleInfoable)()
    SetOdataType(value *string)()
    SetTimeFormat(value *string)()
    SetTimeZone(value *string)()
    SetUserPurpose(value *UserPurpose)()
    SetWorkingHours(value WorkingHoursable)()
}
