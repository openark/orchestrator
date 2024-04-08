package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Audio 
type Audio struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The title of the album for this audio file.
    album *string
    // The artist named on the album for the audio file.
    albumArtist *string
    // The performing artist for the audio file.
    artist *string
    // Bitrate expressed in kbps.
    bitrate *int64
    // The name of the composer of the audio file.
    composers *string
    // Copyright information for the audio file.
    copyright *string
    // The number of the disc this audio file came from.
    disc *int32
    // The total number of discs in this album.
    discCount *int32
    // Duration of the audio file, expressed in milliseconds
    duration *int64
    // The genre of this audio file.
    genre *string
    // Indicates if the file is protected with digital rights management.
    hasDrm *bool
    // Indicates if the file is encoded with a variable bitrate.
    isVariableBitrate *bool
    // The OdataType property
    odataType *string
    // The title of the audio file.
    title *string
    // The number of the track on the original disc for this audio file.
    track *int32
    // The total number of tracks on the original disc for this audio file.
    trackCount *int32
    // The year the audio file was recorded.
    year *int32
}
// NewAudio instantiates a new audio and sets the default values.
func NewAudio()(*Audio) {
    m := &Audio{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAudioFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAudioFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAudio(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Audio) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetAlbum gets the album property value. The title of the album for this audio file.
func (m *Audio) GetAlbum()(*string) {
    return m.album
}
// GetAlbumArtist gets the albumArtist property value. The artist named on the album for the audio file.
func (m *Audio) GetAlbumArtist()(*string) {
    return m.albumArtist
}
// GetArtist gets the artist property value. The performing artist for the audio file.
func (m *Audio) GetArtist()(*string) {
    return m.artist
}
// GetBitrate gets the bitrate property value. Bitrate expressed in kbps.
func (m *Audio) GetBitrate()(*int64) {
    return m.bitrate
}
// GetComposers gets the composers property value. The name of the composer of the audio file.
func (m *Audio) GetComposers()(*string) {
    return m.composers
}
// GetCopyright gets the copyright property value. Copyright information for the audio file.
func (m *Audio) GetCopyright()(*string) {
    return m.copyright
}
// GetDisc gets the disc property value. The number of the disc this audio file came from.
func (m *Audio) GetDisc()(*int32) {
    return m.disc
}
// GetDiscCount gets the discCount property value. The total number of discs in this album.
func (m *Audio) GetDiscCount()(*int32) {
    return m.discCount
}
// GetDuration gets the duration property value. Duration of the audio file, expressed in milliseconds
func (m *Audio) GetDuration()(*int64) {
    return m.duration
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Audio) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["album"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlbum(val)
        }
        return nil
    }
    res["albumArtist"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAlbumArtist(val)
        }
        return nil
    }
    res["artist"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetArtist(val)
        }
        return nil
    }
    res["bitrate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBitrate(val)
        }
        return nil
    }
    res["composers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComposers(val)
        }
        return nil
    }
    res["copyright"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCopyright(val)
        }
        return nil
    }
    res["disc"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisc(val)
        }
        return nil
    }
    res["discCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDiscCount(val)
        }
        return nil
    }
    res["duration"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDuration(val)
        }
        return nil
    }
    res["genre"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGenre(val)
        }
        return nil
    }
    res["hasDrm"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasDrm(val)
        }
        return nil
    }
    res["isVariableBitrate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVariableBitrate(val)
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
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["track"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrack(val)
        }
        return nil
    }
    res["trackCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrackCount(val)
        }
        return nil
    }
    res["year"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetYear(val)
        }
        return nil
    }
    return res
}
// GetGenre gets the genre property value. The genre of this audio file.
func (m *Audio) GetGenre()(*string) {
    return m.genre
}
// GetHasDrm gets the hasDrm property value. Indicates if the file is protected with digital rights management.
func (m *Audio) GetHasDrm()(*bool) {
    return m.hasDrm
}
// GetIsVariableBitrate gets the isVariableBitrate property value. Indicates if the file is encoded with a variable bitrate.
func (m *Audio) GetIsVariableBitrate()(*bool) {
    return m.isVariableBitrate
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Audio) GetOdataType()(*string) {
    return m.odataType
}
// GetTitle gets the title property value. The title of the audio file.
func (m *Audio) GetTitle()(*string) {
    return m.title
}
// GetTrack gets the track property value. The number of the track on the original disc for this audio file.
func (m *Audio) GetTrack()(*int32) {
    return m.track
}
// GetTrackCount gets the trackCount property value. The total number of tracks on the original disc for this audio file.
func (m *Audio) GetTrackCount()(*int32) {
    return m.trackCount
}
// GetYear gets the year property value. The year the audio file was recorded.
func (m *Audio) GetYear()(*int32) {
    return m.year
}
// Serialize serializes information the current object
func (m *Audio) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("album", m.GetAlbum())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("albumArtist", m.GetAlbumArtist())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("artist", m.GetArtist())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("bitrate", m.GetBitrate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("composers", m.GetComposers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("copyright", m.GetCopyright())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("disc", m.GetDisc())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("discCount", m.GetDiscCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("duration", m.GetDuration())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("genre", m.GetGenre())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasDrm", m.GetHasDrm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVariableBitrate", m.GetIsVariableBitrate())
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
        err := writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("track", m.GetTrack())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("trackCount", m.GetTrackCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("year", m.GetYear())
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
func (m *Audio) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetAlbum sets the album property value. The title of the album for this audio file.
func (m *Audio) SetAlbum(value *string)() {
    m.album = value
}
// SetAlbumArtist sets the albumArtist property value. The artist named on the album for the audio file.
func (m *Audio) SetAlbumArtist(value *string)() {
    m.albumArtist = value
}
// SetArtist sets the artist property value. The performing artist for the audio file.
func (m *Audio) SetArtist(value *string)() {
    m.artist = value
}
// SetBitrate sets the bitrate property value. Bitrate expressed in kbps.
func (m *Audio) SetBitrate(value *int64)() {
    m.bitrate = value
}
// SetComposers sets the composers property value. The name of the composer of the audio file.
func (m *Audio) SetComposers(value *string)() {
    m.composers = value
}
// SetCopyright sets the copyright property value. Copyright information for the audio file.
func (m *Audio) SetCopyright(value *string)() {
    m.copyright = value
}
// SetDisc sets the disc property value. The number of the disc this audio file came from.
func (m *Audio) SetDisc(value *int32)() {
    m.disc = value
}
// SetDiscCount sets the discCount property value. The total number of discs in this album.
func (m *Audio) SetDiscCount(value *int32)() {
    m.discCount = value
}
// SetDuration sets the duration property value. Duration of the audio file, expressed in milliseconds
func (m *Audio) SetDuration(value *int64)() {
    m.duration = value
}
// SetGenre sets the genre property value. The genre of this audio file.
func (m *Audio) SetGenre(value *string)() {
    m.genre = value
}
// SetHasDrm sets the hasDrm property value. Indicates if the file is protected with digital rights management.
func (m *Audio) SetHasDrm(value *bool)() {
    m.hasDrm = value
}
// SetIsVariableBitrate sets the isVariableBitrate property value. Indicates if the file is encoded with a variable bitrate.
func (m *Audio) SetIsVariableBitrate(value *bool)() {
    m.isVariableBitrate = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Audio) SetOdataType(value *string)() {
    m.odataType = value
}
// SetTitle sets the title property value. The title of the audio file.
func (m *Audio) SetTitle(value *string)() {
    m.title = value
}
// SetTrack sets the track property value. The number of the track on the original disc for this audio file.
func (m *Audio) SetTrack(value *int32)() {
    m.track = value
}
// SetTrackCount sets the trackCount property value. The total number of tracks on the original disc for this audio file.
func (m *Audio) SetTrackCount(value *int32)() {
    m.trackCount = value
}
// SetYear sets the year property value. The year the audio file was recorded.
func (m *Audio) SetYear(value *int32)() {
    m.year = value
}
