package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SingleServicePrincipal 
type SingleServicePrincipal struct {
    SubjectSet
    // Description of this service principal.
    description *string
    // ID of the servicePrincipal.
    servicePrincipalId *string
}
// NewSingleServicePrincipal instantiates a new SingleServicePrincipal and sets the default values.
func NewSingleServicePrincipal()(*SingleServicePrincipal) {
    m := &SingleServicePrincipal{
        SubjectSet: *NewSubjectSet(),
    }
    odataTypeValue := "#microsoft.graph.singleServicePrincipal"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateSingleServicePrincipalFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSingleServicePrincipalFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSingleServicePrincipal(), nil
}
// GetDescription gets the description property value. Description of this service principal.
func (m *SingleServicePrincipal) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SingleServicePrincipal) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.SubjectSet.GetFieldDeserializers()
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["servicePrincipalId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipalId(val)
        }
        return nil
    }
    return res
}
// GetServicePrincipalId gets the servicePrincipalId property value. ID of the servicePrincipal.
func (m *SingleServicePrincipal) GetServicePrincipalId()(*string) {
    return m.servicePrincipalId
}
// Serialize serializes information the current object
func (m *SingleServicePrincipal) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.SubjectSet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("servicePrincipalId", m.GetServicePrincipalId())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description of this service principal.
func (m *SingleServicePrincipal) SetDescription(value *string)() {
    m.description = value
}
// SetServicePrincipalId sets the servicePrincipalId property value. ID of the servicePrincipal.
func (m *SingleServicePrincipal) SetServicePrincipalId(value *string)() {
    m.servicePrincipalId = value
}
