package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeamFunSettingsable 
type TeamFunSettingsable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAllowCustomMemes()(*bool)
    GetAllowGiphy()(*bool)
    GetAllowStickersAndMemes()(*bool)
    GetGiphyContentRating()(*GiphyRatingType)
    GetOdataType()(*string)
    SetAllowCustomMemes(value *bool)()
    SetAllowGiphy(value *bool)()
    SetAllowStickersAndMemes(value *bool)()
    SetGiphyContentRating(value *GiphyRatingType)()
    SetOdataType(value *string)()
}
