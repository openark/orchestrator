package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// FeatureRolloutPolicy 
type FeatureRolloutPolicy struct {
    Entity
    // Nullable. Specifies a list of directoryObjects that feature is enabled for.
    appliesTo []DirectoryObjectable
    // A description for this feature rollout policy.
    description *string
    // The display name for this  feature rollout policy.
    displayName *string
    // The feature property
    feature *StagedFeatureName
    // Indicates whether this feature rollout policy should be applied to the entire organization.
    isAppliedToOrganization *bool
    // Indicates whether the feature rollout is enabled.
    isEnabled *bool
}
// NewFeatureRolloutPolicy instantiates a new featureRolloutPolicy and sets the default values.
func NewFeatureRolloutPolicy()(*FeatureRolloutPolicy) {
    m := &FeatureRolloutPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateFeatureRolloutPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFeatureRolloutPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewFeatureRolloutPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. Nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *FeatureRolloutPolicy) GetAppliesTo()([]DirectoryObjectable) {
    return m.appliesTo
}
// GetDescription gets the description property value. A description for this feature rollout policy.
func (m *FeatureRolloutPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. The display name for this  feature rollout policy.
func (m *FeatureRolloutPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFeature gets the feature property value. The feature property
func (m *FeatureRolloutPolicy) GetFeature()(*StagedFeatureName) {
    return m.feature
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FeatureRolloutPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["appliesTo"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDirectoryObjectFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DirectoryObjectable, len(val))
            for i, v := range val {
                res[i] = v.(DirectoryObjectable)
            }
            m.SetAppliesTo(res)
        }
        return nil
    }
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
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["feature"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseStagedFeatureName)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFeature(val.(*StagedFeatureName))
        }
        return nil
    }
    res["isAppliedToOrganization"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAppliedToOrganization(val)
        }
        return nil
    }
    res["isEnabled"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabled(val)
        }
        return nil
    }
    return res
}
// GetIsAppliedToOrganization gets the isAppliedToOrganization property value. Indicates whether this feature rollout policy should be applied to the entire organization.
func (m *FeatureRolloutPolicy) GetIsAppliedToOrganization()(*bool) {
    return m.isAppliedToOrganization
}
// GetIsEnabled gets the isEnabled property value. Indicates whether the feature rollout is enabled.
func (m *FeatureRolloutPolicy) GetIsEnabled()(*bool) {
    return m.isEnabled
}
// Serialize serializes information the current object
func (m *FeatureRolloutPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppliesTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAppliesTo()))
        for i, v := range m.GetAppliesTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("appliesTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetFeature() != nil {
        cast := (*m.GetFeature()).String()
        err = writer.WriteStringValue("feature", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAppliedToOrganization", m.GetIsAppliedToOrganization())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabled", m.GetIsEnabled())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. Nullable. Specifies a list of directoryObjects that feature is enabled for.
func (m *FeatureRolloutPolicy) SetAppliesTo(value []DirectoryObjectable)() {
    m.appliesTo = value
}
// SetDescription sets the description property value. A description for this feature rollout policy.
func (m *FeatureRolloutPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. The display name for this  feature rollout policy.
func (m *FeatureRolloutPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFeature sets the feature property value. The feature property
func (m *FeatureRolloutPolicy) SetFeature(value *StagedFeatureName)() {
    m.feature = value
}
// SetIsAppliedToOrganization sets the isAppliedToOrganization property value. Indicates whether this feature rollout policy should be applied to the entire organization.
func (m *FeatureRolloutPolicy) SetIsAppliedToOrganization(value *bool)() {
    m.isAppliedToOrganization = value
}
// SetIsEnabled sets the isEnabled property value. Indicates whether the feature rollout is enabled.
func (m *FeatureRolloutPolicy) SetIsEnabled(value *bool)() {
    m.isEnabled = value
}
