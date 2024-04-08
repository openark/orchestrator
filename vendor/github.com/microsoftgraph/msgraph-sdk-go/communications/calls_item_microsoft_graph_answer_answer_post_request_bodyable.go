package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallsItemMicrosoftGraphAnswerAnswerPostRequestBodyable 
type CallsItemMicrosoftGraphAnswerAnswerPostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcceptedModalities()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Modality)
    GetCallbackUri()(*string)
    GetCallOptions()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable)
    GetMediaConfig()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable)
    GetParticipantCapacity()(*int32)
    SetAcceptedModalities(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Modality)()
    SetCallbackUri(value *string)()
    SetCallOptions(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.IncomingCallOptionsable)()
    SetMediaConfig(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MediaConfigable)()
    SetParticipantCapacity(value *int32)()
}
