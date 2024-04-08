package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody 
type MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody struct {
    // The activityType property
    activityType *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The chainId property
    chainId *int64
    // The previewText property
    previewText iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable
    // The recipients property
    recipients []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable
    // The teamsAppId property
    teamsAppId *string
    // The templateParameters property
    templateParameters []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable
    // The topic property
    topic iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable
}
// NewMicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody instantiates a new MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody and sets the default values.
func NewMicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody()(*MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) {
    m := &MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody(), nil
}
// GetActivityType gets the activityType property value. The activityType property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetActivityType()(*string) {
    return m.activityType
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChainId gets the chainId property value. The chainId property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetChainId()(*int64) {
    return m.chainId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["activityType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityType(val)
        }
        return nil
    }
    res["chainId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChainId(val)
        }
        return nil
    }
    res["previewText"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreviewText(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable))
        }
        return nil
    }
    res["recipients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamworkNotificationRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)
            }
            m.SetRecipients(res)
        }
        return nil
    }
    res["teamsAppId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTeamsAppId(val)
        }
        return nil
    }
    res["templateParameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)
            }
            m.SetTemplateParameters(res)
        }
        return nil
    }
    res["topic"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamworkActivityTopicFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable))
        }
        return nil
    }
    return res
}
// GetPreviewText gets the previewText property value. The previewText property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetPreviewText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable) {
    return m.previewText
}
// GetRecipients gets the recipients property value. The recipients property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetRecipients()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable) {
    return m.recipients
}
// GetTeamsAppId gets the teamsAppId property value. The teamsAppId property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetTeamsAppId()(*string) {
    return m.teamsAppId
}
// GetTemplateParameters gets the templateParameters property value. The templateParameters property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetTemplateParameters()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable) {
    return m.templateParameters
}
// GetTopic gets the topic property value. The topic property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) GetTopic()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable) {
    return m.topic
}
// Serialize serializes information the current object
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("activityType", m.GetActivityType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("chainId", m.GetChainId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("previewText", m.GetPreviewText())
        if err != nil {
            return err
        }
    }
    if m.GetRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetRecipients()))
        for i, v := range m.GetRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("recipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("teamsAppId", m.GetTeamsAppId())
        if err != nil {
            return err
        }
    }
    if m.GetTemplateParameters() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetTemplateParameters()))
        for i, v := range m.GetTemplateParameters() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("templateParameters", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("topic", m.GetTopic())
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
// SetActivityType sets the activityType property value. The activityType property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) SetActivityType(value *string)() {
    m.activityType = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChainId sets the chainId property value. The chainId property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) SetChainId(value *int64)() {
    m.chainId = value
}
// SetPreviewText sets the previewText property value. The previewText property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) SetPreviewText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable)() {
    m.previewText = value
}
// SetRecipients sets the recipients property value. The recipients property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) SetRecipients(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)() {
    m.recipients = value
}
// SetTeamsAppId sets the teamsAppId property value. The teamsAppId property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) SetTeamsAppId(value *string)() {
    m.teamsAppId = value
}
// SetTemplateParameters sets the templateParameters property value. The templateParameters property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) SetTemplateParameters(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)() {
    m.templateParameters = value
}
// SetTopic sets the topic property value. The topic property
func (m *MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBody) SetTopic(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable)() {
    m.topic = value
}
