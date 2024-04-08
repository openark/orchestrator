package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RedundancyDetectionSettingsable 
type RedundancyDetectionSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetIsEnabled()(*bool)
    GetMaxWords()(*int32)
    GetMinWords()(*int32)
    GetOdataType()(*string)
    GetSimilarityThreshold()(*int32)
    SetIsEnabled(value *bool)()
    SetMaxWords(value *int32)()
    SetMinWords(value *int32)()
    SetOdataType(value *string)()
    SetSimilarityThreshold(value *int32)()
}
