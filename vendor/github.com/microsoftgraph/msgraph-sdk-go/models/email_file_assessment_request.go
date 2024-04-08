package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EmailFileAssessmentRequest 
type EmailFileAssessmentRequest struct {
    ThreatAssessmentRequest
    // Base64 encoded .eml email file content. The file content cannot fetch back because it isn't stored.
    contentData *string
    // The reason for mail routed to its destination. Possible values are: none, mailFlowRule, safeSender, blockedSender, advancedSpamFiltering, domainAllowList, domainBlockList, notInAddressBook, firstTimeSender, autoPurgeToInbox, autoPurgeToJunk, autoPurgeToDeleted, outbound, notJunk, junk.
    destinationRoutingReason *MailDestinationRoutingReason
    // The mail recipient whose policies are used to assess the mail.
    recipientEmail *string
}
// NewEmailFileAssessmentRequest instantiates a new EmailFileAssessmentRequest and sets the default values.
func NewEmailFileAssessmentRequest()(*EmailFileAssessmentRequest) {
    m := &EmailFileAssessmentRequest{
        ThreatAssessmentRequest: *NewThreatAssessmentRequest(),
    }
    odataTypeValue := "#microsoft.graph.emailFileAssessmentRequest"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateEmailFileAssessmentRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEmailFileAssessmentRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEmailFileAssessmentRequest(), nil
}
// GetContentData gets the contentData property value. Base64 encoded .eml email file content. The file content cannot fetch back because it isn't stored.
func (m *EmailFileAssessmentRequest) GetContentData()(*string) {
    return m.contentData
}
// GetDestinationRoutingReason gets the destinationRoutingReason property value. The reason for mail routed to its destination. Possible values are: none, mailFlowRule, safeSender, blockedSender, advancedSpamFiltering, domainAllowList, domainBlockList, notInAddressBook, firstTimeSender, autoPurgeToInbox, autoPurgeToJunk, autoPurgeToDeleted, outbound, notJunk, junk.
func (m *EmailFileAssessmentRequest) GetDestinationRoutingReason()(*MailDestinationRoutingReason) {
    return m.destinationRoutingReason
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EmailFileAssessmentRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ThreatAssessmentRequest.GetFieldDeserializers()
    res["contentData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContentData(val)
        }
        return nil
    }
    res["destinationRoutingReason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseMailDestinationRoutingReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDestinationRoutingReason(val.(*MailDestinationRoutingReason))
        }
        return nil
    }
    res["recipientEmail"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecipientEmail(val)
        }
        return nil
    }
    return res
}
// GetRecipientEmail gets the recipientEmail property value. The mail recipient whose policies are used to assess the mail.
func (m *EmailFileAssessmentRequest) GetRecipientEmail()(*string) {
    return m.recipientEmail
}
// Serialize serializes information the current object
func (m *EmailFileAssessmentRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ThreatAssessmentRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("contentData", m.GetContentData())
        if err != nil {
            return err
        }
    }
    if m.GetDestinationRoutingReason() != nil {
        cast := (*m.GetDestinationRoutingReason()).String()
        err = writer.WriteStringValue("destinationRoutingReason", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("recipientEmail", m.GetRecipientEmail())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetContentData sets the contentData property value. Base64 encoded .eml email file content. The file content cannot fetch back because it isn't stored.
func (m *EmailFileAssessmentRequest) SetContentData(value *string)() {
    m.contentData = value
}
// SetDestinationRoutingReason sets the destinationRoutingReason property value. The reason for mail routed to its destination. Possible values are: none, mailFlowRule, safeSender, blockedSender, advancedSpamFiltering, domainAllowList, domainBlockList, notInAddressBook, firstTimeSender, autoPurgeToInbox, autoPurgeToJunk, autoPurgeToDeleted, outbound, notJunk, junk.
func (m *EmailFileAssessmentRequest) SetDestinationRoutingReason(value *MailDestinationRoutingReason)() {
    m.destinationRoutingReason = value
}
// SetRecipientEmail sets the recipientEmail property value. The mail recipient whose policies are used to assess the mail.
func (m *EmailFileAssessmentRequest) SetRecipientEmail(value *string)() {
    m.recipientEmail = value
}
