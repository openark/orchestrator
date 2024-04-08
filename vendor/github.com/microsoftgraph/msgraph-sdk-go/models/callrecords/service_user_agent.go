package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceUserAgent 
type ServiceUserAgent struct {
    UserAgent
    // The role property
    role *ServiceRole
}
// NewServiceUserAgent instantiates a new ServiceUserAgent and sets the default values.
func NewServiceUserAgent()(*ServiceUserAgent) {
    m := &ServiceUserAgent{
        UserAgent: *NewUserAgent(),
    }
    odataTypeValue := "#microsoft.graph.callRecords.serviceUserAgent"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServiceUserAgentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceUserAgentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceUserAgent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceUserAgent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserAgent.GetFieldDeserializers()
    res["role"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceRole)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val.(*ServiceRole))
        }
        return nil
    }
    return res
}
// GetRole gets the role property value. The role property
func (m *ServiceUserAgent) GetRole()(*ServiceRole) {
    return m.role
}
// Serialize serializes information the current object
func (m *ServiceUserAgent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserAgent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetRole() != nil {
        cast := (*m.GetRole()).String()
        err = writer.WriteStringValue("role", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetRole sets the role property value. The role property
func (m *ServiceUserAgent) SetRole(value *ServiceRole)() {
    m.role = value
}
