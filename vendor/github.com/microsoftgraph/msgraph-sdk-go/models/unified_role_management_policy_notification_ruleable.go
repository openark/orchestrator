package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyNotificationRuleable 
type UnifiedRoleManagementPolicyNotificationRuleable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    UnifiedRoleManagementPolicyRuleable
    GetIsDefaultRecipientsEnabled()(*bool)
    GetNotificationLevel()(*string)
    GetNotificationRecipients()([]string)
    GetNotificationType()(*string)
    GetRecipientType()(*string)
    SetIsDefaultRecipientsEnabled(value *bool)()
    SetNotificationLevel(value *string)()
    SetNotificationRecipients(value []string)()
    SetNotificationType(value *string)()
    SetRecipientType(value *string)()
}
