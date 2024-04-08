package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UnmuteParticipantOperationCollectionResponse 
type UnmuteParticipantOperationCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []UnmuteParticipantOperationable
}
// NewUnmuteParticipantOperationCollectionResponse instantiates a new UnmuteParticipantOperationCollectionResponse and sets the default values.
func NewUnmuteParticipantOperationCollectionResponse()(*UnmuteParticipantOperationCollectionResponse) {
    m := &UnmuteParticipantOperationCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateUnmuteParticipantOperationCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnmuteParticipantOperationCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUnmuteParticipantOperationCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnmuteParticipantOperationCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateUnmuteParticipantOperationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UnmuteParticipantOperationable, len(val))
            for i, v := range val {
                res[i] = v.(UnmuteParticipantOperationable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *UnmuteParticipantOperationCollectionResponse) GetValue()([]UnmuteParticipantOperationable) {
    return m.value
}
// Serialize serializes information the current object
func (m *UnmuteParticipantOperationCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *UnmuteParticipantOperationCollectionResponse) SetValue(value []UnmuteParticipantOperationable)() {
    m.value = value
}
