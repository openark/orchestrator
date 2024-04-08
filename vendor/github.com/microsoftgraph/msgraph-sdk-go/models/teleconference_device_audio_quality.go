package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TeleconferenceDeviceAudioQuality 
type TeleconferenceDeviceAudioQuality struct {
    TeleconferenceDeviceMediaQuality
}
// NewTeleconferenceDeviceAudioQuality instantiates a new TeleconferenceDeviceAudioQuality and sets the default values.
func NewTeleconferenceDeviceAudioQuality()(*TeleconferenceDeviceAudioQuality) {
    m := &TeleconferenceDeviceAudioQuality{
        TeleconferenceDeviceMediaQuality: *NewTeleconferenceDeviceMediaQuality(),
    }
    odataTypeValue := "#microsoft.graph.teleconferenceDeviceAudioQuality"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTeleconferenceDeviceAudioQualityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeleconferenceDeviceAudioQualityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTeleconferenceDeviceAudioQuality(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeleconferenceDeviceAudioQuality) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.TeleconferenceDeviceMediaQuality.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *TeleconferenceDeviceAudioQuality) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.TeleconferenceDeviceMediaQuality.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
