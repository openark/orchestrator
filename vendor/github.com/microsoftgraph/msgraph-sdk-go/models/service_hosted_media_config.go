package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceHostedMediaConfig 
type ServiceHostedMediaConfig struct {
    MediaConfig
    // The list of media to pre-fetch.
    preFetchMedia []MediaInfoable
}
// NewServiceHostedMediaConfig instantiates a new ServiceHostedMediaConfig and sets the default values.
func NewServiceHostedMediaConfig()(*ServiceHostedMediaConfig) {
    m := &ServiceHostedMediaConfig{
        MediaConfig: *NewMediaConfig(),
    }
    odataTypeValue := "#microsoft.graph.serviceHostedMediaConfig"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServiceHostedMediaConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceHostedMediaConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceHostedMediaConfig(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceHostedMediaConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.MediaConfig.GetFieldDeserializers()
    res["preFetchMedia"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMediaInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MediaInfoable, len(val))
            for i, v := range val {
                res[i] = v.(MediaInfoable)
            }
            m.SetPreFetchMedia(res)
        }
        return nil
    }
    return res
}
// GetPreFetchMedia gets the preFetchMedia property value. The list of media to pre-fetch.
func (m *ServiceHostedMediaConfig) GetPreFetchMedia()([]MediaInfoable) {
    return m.preFetchMedia
}
// Serialize serializes information the current object
func (m *ServiceHostedMediaConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.MediaConfig.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPreFetchMedia() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPreFetchMedia()))
        for i, v := range m.GetPreFetchMedia() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("preFetchMedia", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPreFetchMedia sets the preFetchMedia property value. The list of media to pre-fetch.
func (m *ServiceHostedMediaConfig) SetPreFetchMedia(value []MediaInfoable)() {
    m.preFetchMedia = value
}
