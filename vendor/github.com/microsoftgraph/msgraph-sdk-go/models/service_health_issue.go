package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ServiceHealthIssue 
type ServiceHealthIssue struct {
    ServiceAnnouncementBase
    // The classification property
    classification *ServiceHealthClassificationType
    // The feature name of the service issue.
    feature *string
    // The feature group name of the service issue.
    featureGroup *string
    // The description of the service issue impact.
    impactDescription *string
    // Indicates whether the issue is resolved.
    isResolved *bool
    // The origin property
    origin *ServiceHealthOrigin
    // Collection of historical posts for the service issue.
    posts []ServiceHealthIssuePostable
    // Indicates the service affected by the issue.
    service *string
    // The status property
    status *ServiceHealthStatus
}
// NewServiceHealthIssue instantiates a new ServiceHealthIssue and sets the default values.
func NewServiceHealthIssue()(*ServiceHealthIssue) {
    m := &ServiceHealthIssue{
        ServiceAnnouncementBase: *NewServiceAnnouncementBase(),
    }
    odataTypeValue := "#microsoft.graph.serviceHealthIssue"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateServiceHealthIssueFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateServiceHealthIssueFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceHealthIssue(), nil
}
// GetClassification gets the classification property value. The classification property
func (m *ServiceHealthIssue) GetClassification()(*ServiceHealthClassificationType) {
    return m.classification
}
// GetFeature gets the feature property value. The feature name of the service issue.
func (m *ServiceHealthIssue) GetFeature()(*string) {
    return m.feature
}
// GetFeatureGroup gets the featureGroup property value. The feature group name of the service issue.
func (m *ServiceHealthIssue) GetFeatureGroup()(*string) {
    return m.featureGroup
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ServiceHealthIssue) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ServiceAnnouncementBase.GetFieldDeserializers()
    res["classification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClassification(val.(*ServiceHealthClassificationType))
        }
        return nil
    }
    res["feature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeature(val)
        }
        return nil
    }
    res["featureGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeatureGroup(val)
        }
        return nil
    }
    res["impactDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImpactDescription(val)
        }
        return nil
    }
    res["isResolved"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsResolved(val)
        }
        return nil
    }
    res["origin"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceHealthOrigin)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrigin(val.(*ServiceHealthOrigin))
        }
        return nil
    }
    res["posts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateServiceHealthIssuePostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ServiceHealthIssuePostable, len(val))
            for i, v := range val {
                res[i] = v.(ServiceHealthIssuePostable)
            }
            m.SetPosts(res)
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
// GetImpactDescription gets the impactDescription property value. The description of the service issue impact.
func (m *ServiceHealthIssue) GetImpactDescription()(*string) {
    return m.impactDescription
}
// GetIsResolved gets the isResolved property value. Indicates whether the issue is resolved.
func (m *ServiceHealthIssue) GetIsResolved()(*bool) {
    return m.isResolved
}
// GetOrigin gets the origin property value. The origin property
func (m *ServiceHealthIssue) GetOrigin()(*ServiceHealthOrigin) {
    return m.origin
}
// GetPosts gets the posts property value. Collection of historical posts for the service issue.
func (m *ServiceHealthIssue) GetPosts()([]ServiceHealthIssuePostable) {
    return m.posts
}
// GetService gets the service property value. Indicates the service affected by the issue.
func (m *ServiceHealthIssue) GetService()(*string) {
    return m.service
}
// GetStatus gets the status property value. The status property
func (m *ServiceHealthIssue) GetStatus()(*ServiceHealthStatus) {
    return m.status
}
// Serialize serializes information the current object
func (m *ServiceHealthIssue) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ServiceAnnouncementBase.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassification() != nil {
        cast := (*m.GetClassification()).String()
        err = writer.WriteStringValue("classification", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("feature", m.GetFeature())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("featureGroup", m.GetFeatureGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("impactDescription", m.GetImpactDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isResolved", m.GetIsResolved())
        if err != nil {
            return err
        }
    }
    if m.GetOrigin() != nil {
        cast := (*m.GetOrigin()).String()
        err = writer.WriteStringValue("origin", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetPosts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPosts()))
        for i, v := range m.GetPosts() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("posts", cast)
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
// SetClassification sets the classification property value. The classification property
func (m *ServiceHealthIssue) SetClassification(value *ServiceHealthClassificationType)() {
    m.classification = value
}
// SetFeature sets the feature property value. The feature name of the service issue.
func (m *ServiceHealthIssue) SetFeature(value *string)() {
    m.feature = value
}
// SetFeatureGroup sets the featureGroup property value. The feature group name of the service issue.
func (m *ServiceHealthIssue) SetFeatureGroup(value *string)() {
    m.featureGroup = value
}
// SetImpactDescription sets the impactDescription property value. The description of the service issue impact.
func (m *ServiceHealthIssue) SetImpactDescription(value *string)() {
    m.impactDescription = value
}
// SetIsResolved sets the isResolved property value. Indicates whether the issue is resolved.
func (m *ServiceHealthIssue) SetIsResolved(value *bool)() {
    m.isResolved = value
}
// SetOrigin sets the origin property value. The origin property
func (m *ServiceHealthIssue) SetOrigin(value *ServiceHealthOrigin)() {
    m.origin = value
}
// SetPosts sets the posts property value. Collection of historical posts for the service issue.
func (m *ServiceHealthIssue) SetPosts(value []ServiceHealthIssuePostable)() {
    m.posts = value
}
// SetService sets the service property value. Indicates the service affected by the issue.
func (m *ServiceHealthIssue) SetService(value *string)() {
    m.service = value
}
// SetStatus sets the status property value. The status property
func (m *ServiceHealthIssue) SetStatus(value *ServiceHealthStatus)() {
    m.status = value
}
