package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceHealth 
type ServiceHealth struct {
    Entity
    // A collection of issues that happened on the service, with detailed information for each issue.
    issues []ServiceHealthIssueable
    // The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
    service *string
    // The status property
    status *ServiceHealthStatus
}
// NewServiceHealth instantiates a new serviceHealth and sets the default values.
func NewServiceHealth()(*ServiceHealth) {
    m := &ServiceHealth{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceHealthFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceHealthFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceHealth(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceHealth) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["issues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceHealthIssueFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealthIssueable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceHealthIssueable)
            }
            m.SetIssues(res)
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ServiceHealthStatus))
        }
        return nil
    }
    return res
}
// GetIssues gets the issues property value. A collection of issues that happened on the service, with detailed information for each issue.
func (m *ServiceHealth) GetIssues()([]ServiceHealthIssueable) {
    return m.issues
}
// GetService gets the service property value. The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
func (m *ServiceHealth) GetService()(*string) {
    return m.service
}
// GetStatus gets the status property value. The status property
func (m *ServiceHealth) GetStatus()(*ServiceHealthStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ServiceHealth) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetIssues() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetIssues()))
        for i, v := range m.GetIssues() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("issues", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err = writer.WriteStringValue("status", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetIssues sets the issues property value. A collection of issues that happened on the service, with detailed information for each issue.
func (m *ServiceHealth) SetIssues(value []ServiceHealthIssueable)() {
    m.issues = value
}
// SetService sets the service property value. The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
func (m *ServiceHealth) SetService(value *string)() {
    m.service = value
}
// SetStatus sets the status property value. The status property
func (m *ServiceHealth) SetStatus(value *ServiceHealthStatus)() {
    m.status = value
}
