package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidManagedAppRegistration 
type AndroidManagedAppRegistration struct {
    ManagedAppRegistration
}
// NewAndroidManagedAppRegistration instantiates a new AndroidManagedAppRegistration and sets the default values.
func NewAndroidManagedAppRegistration()(*AndroidManagedAppRegistration) {
    m := &AndroidManagedAppRegistration{
        ManagedAppRegistration: *NewManagedAppRegistration(),
    }
    odataTypeValue := "#microsoft.graph.androidManagedAppRegistration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateAndroidManagedAppRegistrationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidManagedAppRegistrationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidManagedAppRegistration(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidManagedAppRegistration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppRegistration.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *AndroidManagedAppRegistration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppRegistration.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
