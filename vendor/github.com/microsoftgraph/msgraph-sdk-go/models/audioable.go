package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Audioable 
type Audioable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAlbum()(*string)
    GetAlbumArtist()(*string)
    GetArtist()(*string)
    GetBitrate()(*int64)
    GetComposers()(*string)
    GetCopyright()(*string)
    GetDisc()(*int32)
    GetDiscCount()(*int32)
    GetDuration()(*int64)
    GetGenre()(*string)
    GetHasDrm()(*bool)
    GetIsVariableBitrate()(*bool)
    GetOdataType()(*string)
    GetTitle()(*string)
    GetTrack()(*int32)
    GetTrackCount()(*int32)
    GetYear()(*int32)
    SetAlbum(value *string)()
    SetAlbumArtist(value *string)()
    SetArtist(value *string)()
    SetBitrate(value *int64)()
    SetComposers(value *string)()
    SetCopyright(value *string)()
    SetDisc(value *int32)()
    SetDiscCount(value *int32)()
    SetDuration(value *int64)()
    SetGenre(value *string)()
    SetHasDrm(value *bool)()
    SetIsVariableBitrate(value *bool)()
    SetOdataType(value *string)()
    SetTitle(value *string)()
    SetTrack(value *int32)()
    SetTrackCount(value *int32)()
    SetYear(value *int32)()
}
