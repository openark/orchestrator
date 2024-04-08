package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AutomaticRepliesSetting 
type AutomaticRepliesSetting struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The set of audience external to the signed-in user's organization who will receive the ExternalReplyMessage, if Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all.
    externalAudience *ExternalAudienceScope
    // The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
    externalReplyMessage *string
    // The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
    internalReplyMessage *string
    // The OdataType property
    odataType *string
    // The date and time that automatic replies are set to end, if Status is set to Scheduled.
    scheduledEndDateTime DateTimeTimeZoneable
    // The date and time that automatic replies are set to begin, if Status is set to Scheduled.
    scheduledStartDateTime DateTimeTimeZoneable
    // Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled.
    status *AutomaticRepliesStatus
}
// NewAutomaticRepliesSetting instantiates a new automaticRepliesSetting and sets the default values.
func NewAutomaticRepliesSetting()(*AutomaticRepliesSetting) {
    m := &AutomaticRepliesSetting{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAutomaticRepliesSettingFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAutomaticRepliesSettingFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAutomaticRepliesSetting(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AutomaticRepliesSetting) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetExternalAudience gets the externalAudience property value. The set of audience external to the signed-in user's organization who will receive the ExternalReplyMessage, if Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all.
func (m *AutomaticRepliesSetting) GetExternalAudience()(*ExternalAudienceScope) {
    return m.externalAudience
}
// GetExternalReplyMessage gets the externalReplyMessage property value. The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) GetExternalReplyMessage()(*string) {
    return m.externalReplyMessage
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AutomaticRepliesSetting) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["externalAudience"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseExternalAudienceScope)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalAudience(val.(*ExternalAudienceScope))
        }
        return nil
    }
    res["externalReplyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalReplyMessage(val)
        }
        return nil
    }
    res["internalReplyMessage"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternalReplyMessage(val)
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
    res["scheduledEndDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledEndDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["scheduledStartDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScheduledStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAutomaticRepliesStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*AutomaticRepliesStatus))
        }
        return nil
    }
    return res
}
// GetInternalReplyMessage gets the internalReplyMessage property value. The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) GetInternalReplyMessage()(*string) {
    return m.internalReplyMessage
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AutomaticRepliesSetting) GetOdataType()(*string) {
    return m.odataType
}
// GetScheduledEndDateTime gets the scheduledEndDateTime property value. The date and time that automatic replies are set to end, if Status is set to Scheduled.
func (m *AutomaticRepliesSetting) GetScheduledEndDateTime()(DateTimeTimeZoneable) {
    return m.scheduledEndDateTime
}
// GetScheduledStartDateTime gets the scheduledStartDateTime property value. The date and time that automatic replies are set to begin, if Status is set to Scheduled.
func (m *AutomaticRepliesSetting) GetScheduledStartDateTime()(DateTimeTimeZoneable) {
    return m.scheduledStartDateTime
}
// GetStatus gets the status property value. Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled.
func (m *AutomaticRepliesSetting) GetStatus()(*AutomaticRepliesStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *AutomaticRepliesSetting) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetExternalAudience() != nil {
        cast := (*m.GetExternalAudience()).String()
        err := writer.WriteStringValue("externalAudience", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalReplyMessage", m.GetExternalReplyMessage())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("internalReplyMessage", m.GetInternalReplyMessage())
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
        err := writer.WriteObjectValue("scheduledEndDateTime", m.GetScheduledEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("scheduledStartDateTime", m.GetScheduledStartDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *AutomaticRepliesSetting) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetExternalAudience sets the externalAudience property value. The set of audience external to the signed-in user's organization who will receive the ExternalReplyMessage, if Status is AlwaysEnabled or Scheduled. The possible values are: none, contactsOnly, all.
func (m *AutomaticRepliesSetting) SetExternalAudience(value *ExternalAudienceScope)() {
    m.externalAudience = value
}
// SetExternalReplyMessage sets the externalReplyMessage property value. The automatic reply to send to the specified external audience, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) SetExternalReplyMessage(value *string)() {
    m.externalReplyMessage = value
}
// SetInternalReplyMessage sets the internalReplyMessage property value. The automatic reply to send to the audience internal to the signed-in user's organization, if Status is AlwaysEnabled or Scheduled.
func (m *AutomaticRepliesSetting) SetInternalReplyMessage(value *string)() {
    m.internalReplyMessage = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AutomaticRepliesSetting) SetOdataType(value *string)() {
    m.odataType = value
}
// SetScheduledEndDateTime sets the scheduledEndDateTime property value. The date and time that automatic replies are set to end, if Status is set to Scheduled.
func (m *AutomaticRepliesSetting) SetScheduledEndDateTime(value DateTimeTimeZoneable)() {
    m.scheduledEndDateTime = value
}
// SetScheduledStartDateTime sets the scheduledStartDateTime property value. The date and time that automatic replies are set to begin, if Status is set to Scheduled.
func (m *AutomaticRepliesSetting) SetScheduledStartDateTime(value DateTimeTimeZoneable)() {
    m.scheduledStartDateTime = value
}
// SetStatus sets the status property value. Configurations status for automatic replies. The possible values are: disabled, alwaysEnabled, scheduled.
func (m *AutomaticRepliesSetting) SetStatus(value *AutomaticRepliesStatus)() {
    m.status = value
}
