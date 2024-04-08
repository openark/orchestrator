package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ChatMessagePolicyViolation 
type ChatMessagePolicyViolation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The action taken by the DLP provider on the message with sensitive content. Supported values are: NoneNotifySender -- Inform the sender of the violation but allow readers to read the message.BlockAccess -- Block readers from reading the message.BlockAccessExternal -- Block users outside the organization from reading the message, while allowing users within the organization to read the message.
    dlpAction *ChatMessagePolicyViolationDlpActionTypes
    // Justification text provided by the sender of the message when overriding a policy violation.
    justificationText *string
    // The OdataType property
    odataType *string
    // Information to display to the message sender about why the message was flagged as a violation.
    policyTip ChatMessagePolicyViolationPolicyTipable
    // Indicates the action taken by the user on a message blocked by the DLP provider. Supported values are: NoneOverrideReportFalsePositiveWhen the DLP provider is updating the message for blocking sensitive content, userAction is not required.
    userAction *ChatMessagePolicyViolationUserActionTypes
    // Indicates what actions the sender may take in response to the policy violation. Supported values are: NoneAllowFalsePositiveOverride -- Allows the sender to declare the policyViolation to be an error in the DLP app and its rules, and allow readers to see the message again if the dlpAction had hidden it.AllowOverrideWithoutJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, without needing to provide an explanation for doing so. AllowOverrideWithJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, after providing an explanation for doing so.AllowOverrideWithoutJustification and AllowOverrideWithJustification are mutually exclusive.
    verdictDetails *ChatMessagePolicyViolationVerdictDetailsTypes
}
// NewChatMessagePolicyViolation instantiates a new chatMessagePolicyViolation and sets the default values.
func NewChatMessagePolicyViolation()(*ChatMessagePolicyViolation) {
    m := &ChatMessagePolicyViolation{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateChatMessagePolicyViolationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessagePolicyViolationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewChatMessagePolicyViolation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessagePolicyViolation) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDlpAction gets the dlpAction property value. The action taken by the DLP provider on the message with sensitive content. Supported values are: NoneNotifySender -- Inform the sender of the violation but allow readers to read the message.BlockAccess -- Block readers from reading the message.BlockAccessExternal -- Block users outside the organization from reading the message, while allowing users within the organization to read the message.
func (m *ChatMessagePolicyViolation) GetDlpAction()(*ChatMessagePolicyViolationDlpActionTypes) {
    return m.dlpAction
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessagePolicyViolation) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["dlpAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationDlpActionTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDlpAction(val.(*ChatMessagePolicyViolationDlpActionTypes))
        }
        return nil
    }
    res["justificationText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationText(val)
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
    res["policyTip"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessagePolicyViolationPolicyTipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTip(val.(ChatMessagePolicyViolationPolicyTipable))
        }
        return nil
    }
    res["userAction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationUserActionTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAction(val.(*ChatMessagePolicyViolationUserActionTypes))
        }
        return nil
    }
    res["verdictDetails"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationVerdictDetailsTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerdictDetails(val.(*ChatMessagePolicyViolationVerdictDetailsTypes))
        }
        return nil
    }
    return res
}
// GetJustificationText gets the justificationText property value. Justification text provided by the sender of the message when overriding a policy violation.
func (m *ChatMessagePolicyViolation) GetJustificationText()(*string) {
    return m.justificationText
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ChatMessagePolicyViolation) GetOdataType()(*string) {
    return m.odataType
}
// GetPolicyTip gets the policyTip property value. Information to display to the message sender about why the message was flagged as a violation.
func (m *ChatMessagePolicyViolation) GetPolicyTip()(ChatMessagePolicyViolationPolicyTipable) {
    return m.policyTip
}
// GetUserAction gets the userAction property value. Indicates the action taken by the user on a message blocked by the DLP provider. Supported values are: NoneOverrideReportFalsePositiveWhen the DLP provider is updating the message for blocking sensitive content, userAction is not required.
func (m *ChatMessagePolicyViolation) GetUserAction()(*ChatMessagePolicyViolationUserActionTypes) {
    return m.userAction
}
// GetVerdictDetails gets the verdictDetails property value. Indicates what actions the sender may take in response to the policy violation. Supported values are: NoneAllowFalsePositiveOverride -- Allows the sender to declare the policyViolation to be an error in the DLP app and its rules, and allow readers to see the message again if the dlpAction had hidden it.AllowOverrideWithoutJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, without needing to provide an explanation for doing so. AllowOverrideWithJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, after providing an explanation for doing so.AllowOverrideWithoutJustification and AllowOverrideWithJustification are mutually exclusive.
func (m *ChatMessagePolicyViolation) GetVerdictDetails()(*ChatMessagePolicyViolationVerdictDetailsTypes) {
    return m.verdictDetails
}
// Serialize serializes information the current object
func (m *ChatMessagePolicyViolation) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDlpAction() != nil {
        cast := (*m.GetDlpAction()).String()
        err := writer.WriteStringValue("dlpAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justificationText", m.GetJustificationText())
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
        err := writer.WriteObjectValue("policyTip", m.GetPolicyTip())
        if err != nil {
            return err
        }
    }
    if m.GetUserAction() != nil {
        cast := (*m.GetUserAction()).String()
        err := writer.WriteStringValue("userAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVerdictDetails() != nil {
        cast := (*m.GetVerdictDetails()).String()
        err := writer.WriteStringValue("verdictDetails", &cast)
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
func (m *ChatMessagePolicyViolation) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDlpAction sets the dlpAction property value. The action taken by the DLP provider on the message with sensitive content. Supported values are: NoneNotifySender -- Inform the sender of the violation but allow readers to read the message.BlockAccess -- Block readers from reading the message.BlockAccessExternal -- Block users outside the organization from reading the message, while allowing users within the organization to read the message.
func (m *ChatMessagePolicyViolation) SetDlpAction(value *ChatMessagePolicyViolationDlpActionTypes)() {
    m.dlpAction = value
}
// SetJustificationText sets the justificationText property value. Justification text provided by the sender of the message when overriding a policy violation.
func (m *ChatMessagePolicyViolation) SetJustificationText(value *string)() {
    m.justificationText = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ChatMessagePolicyViolation) SetOdataType(value *string)() {
    m.odataType = value
}
// SetPolicyTip sets the policyTip property value. Information to display to the message sender about why the message was flagged as a violation.
func (m *ChatMessagePolicyViolation) SetPolicyTip(value ChatMessagePolicyViolationPolicyTipable)() {
    m.policyTip = value
}
// SetUserAction sets the userAction property value. Indicates the action taken by the user on a message blocked by the DLP provider. Supported values are: NoneOverrideReportFalsePositiveWhen the DLP provider is updating the message for blocking sensitive content, userAction is not required.
func (m *ChatMessagePolicyViolation) SetUserAction(value *ChatMessagePolicyViolationUserActionTypes)() {
    m.userAction = value
}
// SetVerdictDetails sets the verdictDetails property value. Indicates what actions the sender may take in response to the policy violation. Supported values are: NoneAllowFalsePositiveOverride -- Allows the sender to declare the policyViolation to be an error in the DLP app and its rules, and allow readers to see the message again if the dlpAction had hidden it.AllowOverrideWithoutJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, without needing to provide an explanation for doing so. AllowOverrideWithJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, after providing an explanation for doing so.AllowOverrideWithoutJustification and AllowOverrideWithJustification are mutually exclusive.
func (m *ChatMessagePolicyViolation) SetVerdictDetails(value *ChatMessagePolicyViolationVerdictDetailsTypes)() {
    m.verdictDetails = value
}
