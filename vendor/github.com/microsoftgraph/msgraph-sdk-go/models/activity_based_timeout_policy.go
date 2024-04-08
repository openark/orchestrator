package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ActivityBasedTimeoutPolicy 
type ActivityBasedTimeoutPolicy struct {
    StsPolicy
}
// NewActivityBasedTimeoutPolicy instantiates a new ActivityBasedTimeoutPolicy and sets the default values.
func NewActivityBasedTimeoutPolicy()(*ActivityBasedTimeoutPolicy) {
    m := &ActivityBasedTimeoutPolicy{
        StsPolicy: *NewStsPolicy(),
    }
    odataTypeValue := "#microsoft.graph.activityBasedTimeoutPolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateActivityBasedTimeoutPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateActivityBasedTimeoutPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewActivityBasedTimeoutPolicy(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ActivityBasedTimeoutPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.StsPolicy.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ActivityBasedTimeoutPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.StsPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
