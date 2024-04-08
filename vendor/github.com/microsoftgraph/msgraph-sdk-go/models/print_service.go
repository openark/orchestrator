package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PrintService 
type PrintService struct {
    Entity
    // Endpoints that can be used to access the service. Read-only. Nullable.
    endpoints []PrintServiceEndpointable
}
// NewPrintService instantiates a new PrintService and sets the default values.
func NewPrintService()(*PrintService) {
    m := &PrintService{
        Entity: *NewEntity(),
    }
    return m
}
// CreatePrintServiceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePrintServiceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPrintService(), nil
}
// GetEndpoints gets the endpoints property value. Endpoints that can be used to access the service. Read-only. Nullable.
func (m *PrintService) GetEndpoints()([]PrintServiceEndpointable) {
    return m.endpoints
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintService) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["endpoints"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePrintServiceEndpointFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintServiceEndpointable, len(val))
            for i, v := range val {
                res[i] = v.(PrintServiceEndpointable)
            }
            m.SetEndpoints(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *PrintService) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetEndpoints() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetEndpoints()))
        for i, v := range m.GetEndpoints() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("endpoints", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEndpoints sets the endpoints property value. Endpoints that can be used to access the service. Read-only. Nullable.
func (m *PrintService) SetEndpoints(value []PrintServiceEndpointable)() {
    m.endpoints = value
}
