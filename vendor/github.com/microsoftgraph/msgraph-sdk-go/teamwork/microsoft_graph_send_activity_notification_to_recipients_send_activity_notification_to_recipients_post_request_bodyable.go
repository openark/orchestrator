package teamwork

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBodyable 
type MicrosoftGraphSendActivityNotificationToRecipientsSendActivityNotificationToRecipientsPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivityType()(*string)
    GetChainId()(*int64)
    GetPreviewText()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable)
    GetRecipients()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)
    GetTeamsAppId()(*string)
    GetTemplateParameters()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)
    GetTopic()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable)
    SetActivityType(value *string)()
    SetChainId(value *int64)()
    SetPreviewText(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemBodyable)()
    SetRecipients(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkNotificationRecipientable)()
    SetTeamsAppId(value *string)()
    SetTemplateParameters(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.KeyValuePairable)()
    SetTopic(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamworkActivityTopicable)()
}
