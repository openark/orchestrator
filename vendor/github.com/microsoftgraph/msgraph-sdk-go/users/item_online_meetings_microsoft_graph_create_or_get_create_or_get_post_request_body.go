package users

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody 
type ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The chatInfo property
    chatInfo iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatInfoable
    // The endDateTime property
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The externalId property
    externalId *string
    // The participants property
    participants iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingParticipantsable
    // The startDateTime property
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The subject property
    subject *string
}
// NewItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody instantiates a new ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody and sets the default values.
func NewItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody()(*ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) {
    m := &ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetChatInfo gets the chatInfo property value. The chatInfo property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) GetChatInfo()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatInfoable) {
    return m.chatInfo
}
// GetEndDateTime gets the endDateTime property value. The endDateTime property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.endDateTime
}
// GetExternalId gets the externalId property value. The externalId property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) GetExternalId()(*string) {
    return m.externalId
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["chatInfo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatInfoable))
        }
        return nil
    }
    res["endDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["externalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["participants"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMeetingParticipantsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipants(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingParticipantsable))
        }
        return nil
    }
    res["startDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
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
// GetParticipants gets the participants property value. The participants property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) GetParticipants()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingParticipantsable) {
    return m.participants
}
// GetStartDateTime gets the startDateTime property value. The startDateTime property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.startDateTime
}
// GetSubject gets the subject property value. The subject property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) GetSubject()(*string) {
    return m.subject
}
// Serialize serializes information the current object
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("chatInfo", m.GetChatInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subject", m.GetSubject())
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
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetChatInfo sets the chatInfo property value. The chatInfo property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) SetChatInfo(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChatInfoable)() {
    m.chatInfo = value
}
// SetEndDateTime sets the endDateTime property value. The endDateTime property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// SetExternalId sets the externalId property value. The externalId property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) SetExternalId(value *string)() {
    m.externalId = value
}
// SetParticipants sets the participants property value. The participants property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) SetParticipants(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MeetingParticipantsable)() {
    m.participants = value
}
// SetStartDateTime sets the startDateTime property value. The startDateTime property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// SetSubject sets the subject property value. The subject property
func (m *ItemOnlineMeetingsMicrosoftGraphCreateOrGetCreateOrGetPostRequestBody) SetSubject(value *string)() {
    m.subject = value
}
