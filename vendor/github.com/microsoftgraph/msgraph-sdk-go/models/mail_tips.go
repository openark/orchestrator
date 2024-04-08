package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MailTips 
type MailTips struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // Mail tips for automatic reply if it has been set up by the recipient.
    automaticReplies AutomaticRepliesMailTipsable
    // A custom mail tip that can be set on the recipient's mailbox.
    customMailTip *string
    // Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
    deliveryRestricted *bool
    // The email address of the recipient to get mailtips for.
    emailAddress EmailAddressable
    // Errors that occur during the getMailTips action.
    error MailTipsErrorable
    // The number of external members if the recipient is a distribution list.
    externalMemberCount *int32
    // Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
    isModerated *bool
    // The mailbox full status of the recipient.
    mailboxFull *bool
    // The maximum message size that has been configured for the recipient's organization or mailbox.
    maxMessageSize *int32
    // The OdataType property
    odataType *string
    // The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
    recipientScope *RecipientScopeType
    // Recipients suggested based on previous contexts where they appear in the same message.
    recipientSuggestions []Recipientable
    // The number of members if the recipient is a distribution list.
    totalMemberCount *int32
}
// NewMailTips instantiates a new mailTips and sets the default values.
func NewMailTips()(*MailTips) {
    m := &MailTips{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMailTipsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMailTipsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMailTips(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MailTips) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAutomaticReplies gets the automaticReplies property value. Mail tips for automatic reply if it has been set up by the recipient.
func (m *MailTips) GetAutomaticReplies()(AutomaticRepliesMailTipsable) {
    return m.automaticReplies
}
// GetCustomMailTip gets the customMailTip property value. A custom mail tip that can be set on the recipient's mailbox.
func (m *MailTips) GetCustomMailTip()(*string) {
    return m.customMailTip
}
// GetDeliveryRestricted gets the deliveryRestricted property value. Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
func (m *MailTips) GetDeliveryRestricted()(*bool) {
    return m.deliveryRestricted
}
// GetEmailAddress gets the emailAddress property value. The email address of the recipient to get mailtips for.
func (m *MailTips) GetEmailAddress()(EmailAddressable) {
    return m.emailAddress
}
// GetError gets the error property value. Errors that occur during the getMailTips action.
func (m *MailTips) GetError()(MailTipsErrorable) {
    return m.error
}
// GetExternalMemberCount gets the externalMemberCount property value. The number of external members if the recipient is a distribution list.
func (m *MailTips) GetExternalMemberCount()(*int32) {
    return m.externalMemberCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MailTips) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["automaticReplies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAutomaticRepliesMailTipsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAutomaticReplies(val.(AutomaticRepliesMailTipsable))
        }
        return nil
    }
    res["customMailTip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomMailTip(val)
        }
        return nil
    }
    res["deliveryRestricted"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeliveryRestricted(val)
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
    res["error"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateMailTipsErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(MailTipsErrorable))
        }
        return nil
    }
    res["externalMemberCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalMemberCount(val)
        }
        return nil
    }
    res["isModerated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsModerated(val)
        }
        return nil
    }
    res["mailboxFull"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMailboxFull(val)
        }
        return nil
    }
    res["maxMessageSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxMessageSize(val)
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
    res["recipientScope"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRecipientScopeType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientScope(val.(*RecipientScopeType))
        }
        return nil
    }
    res["recipientSuggestions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetRecipientSuggestions(res)
        }
        return nil
    }
    res["totalMemberCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalMemberCount(val)
        }
        return nil
    }
    return res
}
// GetIsModerated gets the isModerated property value. Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
func (m *MailTips) GetIsModerated()(*bool) {
    return m.isModerated
}
// GetMailboxFull gets the mailboxFull property value. The mailbox full status of the recipient.
func (m *MailTips) GetMailboxFull()(*bool) {
    return m.mailboxFull
}
// GetMaxMessageSize gets the maxMessageSize property value. The maximum message size that has been configured for the recipient's organization or mailbox.
func (m *MailTips) GetMaxMessageSize()(*int32) {
    return m.maxMessageSize
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MailTips) GetOdataType()(*string) {
    return m.odataType
}
// GetRecipientScope gets the recipientScope property value. The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
func (m *MailTips) GetRecipientScope()(*RecipientScopeType) {
    return m.recipientScope
}
// GetRecipientSuggestions gets the recipientSuggestions property value. Recipients suggested based on previous contexts where they appear in the same message.
func (m *MailTips) GetRecipientSuggestions()([]Recipientable) {
    return m.recipientSuggestions
}
// GetTotalMemberCount gets the totalMemberCount property value. The number of members if the recipient is a distribution list.
func (m *MailTips) GetTotalMemberCount()(*int32) {
    return m.totalMemberCount
}
// Serialize serializes information the current object
func (m *MailTips) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("automaticReplies", m.GetAutomaticReplies())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customMailTip", m.GetCustomMailTip())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("deliveryRestricted", m.GetDeliveryRestricted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("error", m.GetError())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("externalMemberCount", m.GetExternalMemberCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isModerated", m.GetIsModerated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("mailboxFull", m.GetMailboxFull())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("maxMessageSize", m.GetMaxMessageSize())
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
    if m.GetRecipientScope() != nil {
        cast := (*m.GetRecipientScope()).String()
        err := writer.WriteStringValue("recipientScope", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetRecipientSuggestions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipientSuggestions()))
        for i, v := range m.GetRecipientSuggestions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("recipientSuggestions", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("totalMemberCount", m.GetTotalMemberCount())
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
func (m *MailTips) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAutomaticReplies sets the automaticReplies property value. Mail tips for automatic reply if it has been set up by the recipient.
func (m *MailTips) SetAutomaticReplies(value AutomaticRepliesMailTipsable)() {
    m.automaticReplies = value
}
// SetCustomMailTip sets the customMailTip property value. A custom mail tip that can be set on the recipient's mailbox.
func (m *MailTips) SetCustomMailTip(value *string)() {
    m.customMailTip = value
}
// SetDeliveryRestricted sets the deliveryRestricted property value. Whether the recipient's mailbox is restricted, for example, accepting messages from only a predefined list of senders, rejecting messages from a predefined list of senders, or accepting messages from only authenticated senders.
func (m *MailTips) SetDeliveryRestricted(value *bool)() {
    m.deliveryRestricted = value
}
// SetEmailAddress sets the emailAddress property value. The email address of the recipient to get mailtips for.
func (m *MailTips) SetEmailAddress(value EmailAddressable)() {
    m.emailAddress = value
}
// SetError sets the error property value. Errors that occur during the getMailTips action.
func (m *MailTips) SetError(value MailTipsErrorable)() {
    m.error = value
}
// SetExternalMemberCount sets the externalMemberCount property value. The number of external members if the recipient is a distribution list.
func (m *MailTips) SetExternalMemberCount(value *int32)() {
    m.externalMemberCount = value
}
// SetIsModerated sets the isModerated property value. Whether sending messages to the recipient requires approval. For example, if the recipient is a large distribution list and a moderator has been set up to approve messages sent to that distribution list, or if sending messages to a recipient requires approval of the recipient's manager.
func (m *MailTips) SetIsModerated(value *bool)() {
    m.isModerated = value
}
// SetMailboxFull sets the mailboxFull property value. The mailbox full status of the recipient.
func (m *MailTips) SetMailboxFull(value *bool)() {
    m.mailboxFull = value
}
// SetMaxMessageSize sets the maxMessageSize property value. The maximum message size that has been configured for the recipient's organization or mailbox.
func (m *MailTips) SetMaxMessageSize(value *int32)() {
    m.maxMessageSize = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MailTips) SetOdataType(value *string)() {
    m.odataType = value
}
// SetRecipientScope sets the recipientScope property value. The scope of the recipient. Possible values are: none, internal, external, externalPartner, externalNonParther. For example, an administrator can set another organization to be its 'partner'. The scope is useful if an administrator wants certain mailtips to be accessible to certain scopes. It's also useful to senders to inform them that their message may leave the organization, helping them make the correct decisions about wording, tone and content.
func (m *MailTips) SetRecipientScope(value *RecipientScopeType)() {
    m.recipientScope = value
}
// SetRecipientSuggestions sets the recipientSuggestions property value. Recipients suggested based on previous contexts where they appear in the same message.
func (m *MailTips) SetRecipientSuggestions(value []Recipientable)() {
    m.recipientSuggestions = value
}
// SetTotalMemberCount sets the totalMemberCount property value. The number of members if the recipient is a distribution list.
func (m *MailTips) SetTotalMemberCount(value *int32)() {
    m.totalMemberCount = value
}
