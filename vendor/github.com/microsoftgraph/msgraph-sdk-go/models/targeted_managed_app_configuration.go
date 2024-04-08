package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TargetedManagedAppConfiguration 
type TargetedManagedAppConfiguration struct {
    ManagedAppConfiguration
    // List of apps to which the policy is deployed.
    apps []ManagedMobileAppable
    // Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
    assignments []TargetedManagedAppPolicyAssignmentable
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32
    // Navigation property to deployment summary of the configuration.
    deploymentSummary ManagedAppPolicyDeploymentSummaryable
    // Indicates if the policy is deployed to any inclusion groups or not.
    isAssigned *bool
}
// NewTargetedManagedAppConfiguration instantiates a new TargetedManagedAppConfiguration and sets the default values.
func NewTargetedManagedAppConfiguration()(*TargetedManagedAppConfiguration) {
    m := &TargetedManagedAppConfiguration{
        ManagedAppConfiguration: *NewManagedAppConfiguration(),
    }
    odataTypeValue := "#microsoft.graph.targetedManagedAppConfiguration"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateTargetedManagedAppConfigurationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTargetedManagedAppConfigurationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTargetedManagedAppConfiguration(), nil
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *TargetedManagedAppConfiguration) GetApps()([]ManagedMobileAppable) {
    return m.apps
}
// GetAssignments gets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *TargetedManagedAppConfiguration) GetAssignments()([]TargetedManagedAppPolicyAssignmentable) {
    return m.assignments
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *TargetedManagedAppConfiguration) GetDeployedAppCount()(*int32) {
    return m.deployedAppCount
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *TargetedManagedAppConfiguration) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    return m.deploymentSummary
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TargetedManagedAppConfiguration) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ManagedAppConfiguration.GetFieldDeserializers()
    res["apps"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    res["assignments"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateTargetedManagedAppPolicyAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppPolicyAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(TargetedManagedAppPolicyAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["deployedAppCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["deploymentSummary"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(ManagedAppPolicyDeploymentSummaryable))
        }
        return nil
    }
    res["isAssigned"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAssigned(val)
        }
        return nil
    }
    return res
}
// GetIsAssigned gets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *TargetedManagedAppConfiguration) GetIsAssigned()(*bool) {
    return m.isAssigned
}
// Serialize serializes information the current object
func (m *TargetedManagedAppConfiguration) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ManagedAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetApps() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deployedAppCount", m.GetDeployedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSummary", m.GetDeploymentSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *TargetedManagedAppConfiguration) SetApps(value []ManagedMobileAppable)() {
    m.apps = value
}
// SetAssignments sets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *TargetedManagedAppConfiguration) SetAssignments(value []TargetedManagedAppPolicyAssignmentable)() {
    m.assignments = value
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *TargetedManagedAppConfiguration) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *TargetedManagedAppConfiguration) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    m.deploymentSummary = value
}
// SetIsAssigned sets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *TargetedManagedAppConfiguration) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
