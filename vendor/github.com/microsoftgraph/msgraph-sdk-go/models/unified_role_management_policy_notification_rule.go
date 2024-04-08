package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnifiedRoleManagementPolicyNotificationRule 
type UnifiedRoleManagementPolicyNotificationRule struct {
    UnifiedRoleManagementPolicyRule
    // Indicates whether a default recipient will receive the notification email.
    isDefaultRecipientsEnabled *bool
    // The level of notification. The possible values are None, Critical, All.
    notificationLevel *string
    // The list of recipients of the email notifications.
    notificationRecipients []string
    // The type of notification. Only Email is supported.
    notificationType *string
    // The type of recipient of the notification. The possible values are Requestor, Approver, Admin.
    recipientType *string
}
// NewUnifiedRoleManagementPolicyNotificationRule instantiates a new UnifiedRoleManagementPolicyNotificationRule and sets the default values.
func NewUnifiedRoleManagementPolicyNotificationRule()(*UnifiedRoleManagementPolicyNotificationRule) {
    m := &UnifiedRoleManagementPolicyNotificationRule{
        UnifiedRoleManagementPolicyRule: *NewUnifiedRoleManagementPolicyRule(),
    }
    odataTypeValue := "#microsoft.graph.unifiedRoleManagementPolicyNotificationRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateUnifiedRoleManagementPolicyNotificationRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnifiedRoleManagementPolicyNotificationRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnifiedRoleManagementPolicyNotificationRule(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnifiedRoleManagementPolicyNotificationRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UnifiedRoleManagementPolicyRule.GetFieldDeserializers()
    res["isDefaultRecipientsEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDefaultRecipientsEnabled(val)
        }
        return nil
    }
    res["notificationLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationLevel(val)
        }
        return nil
    }
    res["notificationRecipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetNotificationRecipients(res)
        }
        return nil
    }
    res["notificationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationType(val)
        }
        return nil
    }
    res["recipientType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientType(val)
        }
        return nil
    }
    return res
}
// GetIsDefaultRecipientsEnabled gets the isDefaultRecipientsEnabled property value. Indicates whether a default recipient will receive the notification email.
func (m *UnifiedRoleManagementPolicyNotificationRule) GetIsDefaultRecipientsEnabled()(*bool) {
    return m.isDefaultRecipientsEnabled
}
// GetNotificationLevel gets the notificationLevel property value. The level of notification. The possible values are None, Critical, All.
func (m *UnifiedRoleManagementPolicyNotificationRule) GetNotificationLevel()(*string) {
    return m.notificationLevel
}
// GetNotificationRecipients gets the notificationRecipients property value. The list of recipients of the email notifications.
func (m *UnifiedRoleManagementPolicyNotificationRule) GetNotificationRecipients()([]string) {
    return m.notificationRecipients
}
// GetNotificationType gets the notificationType property value. The type of notification. Only Email is supported.
func (m *UnifiedRoleManagementPolicyNotificationRule) GetNotificationType()(*string) {
    return m.notificationType
}
// GetRecipientType gets the recipientType property value. The type of recipient of the notification. The possible values are Requestor, Approver, Admin.
func (m *UnifiedRoleManagementPolicyNotificationRule) GetRecipientType()(*string) {
    return m.recipientType
}
// Serialize serializes information the current object
func (m *UnifiedRoleManagementPolicyNotificationRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UnifiedRoleManagementPolicyRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("isDefaultRecipientsEnabled", m.GetIsDefaultRecipientsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationLevel", m.GetNotificationLevel())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationRecipients() != nil {
        err = writer.WriteCollectionOfStringValues("notificationRecipients", m.GetNotificationRecipients())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationType", m.GetNotificationType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientType", m.GetRecipientType())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIsDefaultRecipientsEnabled sets the isDefaultRecipientsEnabled property value. Indicates whether a default recipient will receive the notification email.
func (m *UnifiedRoleManagementPolicyNotificationRule) SetIsDefaultRecipientsEnabled(value *bool)() {
    m.isDefaultRecipientsEnabled = value
}
// SetNotificationLevel sets the notificationLevel property value. The level of notification. The possible values are None, Critical, All.
func (m *UnifiedRoleManagementPolicyNotificationRule) SetNotificationLevel(value *string)() {
    m.notificationLevel = value
}
// SetNotificationRecipients sets the notificationRecipients property value. The list of recipients of the email notifications.
func (m *UnifiedRoleManagementPolicyNotificationRule) SetNotificationRecipients(value []string)() {
    m.notificationRecipients = value
}
// SetNotificationType sets the notificationType property value. The type of notification. Only Email is supported.
func (m *UnifiedRoleManagementPolicyNotificationRule) SetNotificationType(value *string)() {
    m.notificationType = value
}
// SetRecipientType sets the recipientType property value. The type of recipient of the notification. The possible values are Requestor, Approver, Admin.
func (m *UnifiedRoleManagementPolicyNotificationRule) SetRecipientType(value *string)() {
    m.recipientType = value
}
