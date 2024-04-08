package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceAnnouncement 
type ServiceAnnouncement struct {
    Entity
    // A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
    healthOverviews []ServiceHealthable
    // A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
    issues []ServiceHealthIssueable
    // A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
    messages []ServiceUpdateMessageable
}
// NewServiceAnnouncement instantiates a new serviceAnnouncement and sets the default values.
func NewServiceAnnouncement()(*ServiceAnnouncement) {
    m := &ServiceAnnouncement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateServiceAnnouncementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceAnnouncementFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceAnnouncement(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceAnnouncement) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["healthOverviews"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceHealthFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealthable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceHealthable)
            }
            m.SetHealthOverviews(res)
        }
        return nil
    }
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
    res["messages"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceUpdateMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceUpdateMessageable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceUpdateMessageable)
            }
            m.SetMessages(res)
        }
        return nil
    }
    return res
}
// GetHealthOverviews gets the healthOverviews property value. A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetHealthOverviews()([]ServiceHealthable) {
    return m.healthOverviews
}
// GetIssues gets the issues property value. A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetIssues()([]ServiceHealthIssueable) {
    return m.issues
}
// GetMessages gets the messages property value. A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) GetMessages()([]ServiceUpdateMessageable) {
    return m.messages
}
// Serialize serializes information the current object
func (m *ServiceAnnouncement) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetHealthOverviews() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetHealthOverviews()))
        for i, v := range m.GetHealthOverviews() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("healthOverviews", cast)
        if err != nil {
            return err
        }
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
    if m.GetMessages() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetHealthOverviews sets the healthOverviews property value. A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetHealthOverviews(value []ServiceHealthable)() {
    m.healthOverviews = value
}
// SetIssues sets the issues property value. A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetIssues(value []ServiceHealthIssueable)() {
    m.issues = value
}
// SetMessages sets the messages property value. A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
func (m *ServiceAnnouncement) SetMessages(value []ServiceUpdateMessageable)() {
    m.messages = value
}
