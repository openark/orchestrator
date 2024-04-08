package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServicePrincipalRiskDetectionCollectionResponse 
type ServicePrincipalRiskDetectionCollectionResponse struct {
    BaseCollectionPaginationCountResponse
    // The value property
    value []ServicePrincipalRiskDetectionable
}
// NewServicePrincipalRiskDetectionCollectionResponse instantiates a new ServicePrincipalRiskDetectionCollectionResponse and sets the default values.
func NewServicePrincipalRiskDetectionCollectionResponse()(*ServicePrincipalRiskDetectionCollectionResponse) {
    m := &ServicePrincipalRiskDetectionCollectionResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateServicePrincipalRiskDetectionCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServicePrincipalRiskDetectionCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServicePrincipalRiskDetectionCollectionResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServicePrincipalRiskDetectionCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServicePrincipalRiskDetectionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServicePrincipalRiskDetectionable, len(val))
            for i, v := range val {
                res[i] = v.(ServicePrincipalRiskDetectionable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
func (m *ServicePrincipalRiskDetectionCollectionResponse) GetValue()([]ServicePrincipalRiskDetectionable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ServicePrincipalRiskDetectionCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *ServicePrincipalRiskDetectionCollectionResponse) SetValue(value []ServicePrincipalRiskDetectionable)() {
    m.value = value
}
