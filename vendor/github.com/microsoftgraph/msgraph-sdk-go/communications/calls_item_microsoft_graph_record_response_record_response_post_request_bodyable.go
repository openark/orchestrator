package communications

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CallsItemMicrosoftGraphRecordResponseRecordResponsePostRequestBodyable 
type CallsItemMicrosoftGraphRecordResponseRecordResponsePostRequestBodyable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBargeInAllowed()(*bool)
    GetClientContext()(*string)
    GetInitialSilenceTimeoutInSeconds()(*int32)
    GetMaxRecordDurationInSeconds()(*int32)
    GetMaxSilenceTimeoutInSeconds()(*int32)
    GetPlayBeep()(*bool)
    GetPrompts()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable)
    GetStopTones()([]string)
    SetBargeInAllowed(value *bool)()
    SetClientContext(value *string)()
    SetInitialSilenceTimeoutInSeconds(value *int32)()
    SetMaxRecordDurationInSeconds(value *int32)()
    SetMaxSilenceTimeoutInSeconds(value *int32)()
    SetPlayBeep(value *bool)()
    SetPrompts(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Promptable)()
    SetStopTones(value []string)()
}
