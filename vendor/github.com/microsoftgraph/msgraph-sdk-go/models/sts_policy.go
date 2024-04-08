package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// StsPolicy 
type StsPolicy struct {
    PolicyBase
    // The appliesTo property
    appliesTo []DirectoryObjectable
    // A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
    definition []string
    // If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
    isOrganizationDefault *bool
}
// NewStsPolicy instantiates a new StsPolicy and sets the default values.
func NewStsPolicy()(*StsPolicy) {
    m := &StsPolicy{
        PolicyBase: *NewPolicyBase(),
    }
    odataTypeValue := "#microsoft.graph.stsPolicy"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateStsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("@odata.type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "#microsoft.graph.activityBasedTimeoutPolicy":
                        return NewActivityBasedTimeoutPolicy(), nil
                    case "#microsoft.graph.claimsMappingPolicy":
                        return NewClaimsMappingPolicy(), nil
                    case "#microsoft.graph.homeRealmDiscoveryPolicy":
                        return NewHomeRealmDiscoveryPolicy(), nil
                    case "#microsoft.graph.tokenIssuancePolicy":
                        return NewTokenIssuancePolicy(), nil
                    case "#microsoft.graph.tokenLifetimePolicy":
                        return NewTokenLifetimePolicy(), nil
                }
            }
        }
    }
    return NewStsPolicy(), nil
}
// GetAppliesTo gets the appliesTo property value. The appliesTo property
func (m *StsPolicy) GetAppliesTo()([]DirectoryObjectable) {
    return m.appliesTo
}
// GetDefinition gets the definition property value. A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
func (m *StsPolicy) GetDefinition()([]string) {
    return m.definition
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.PolicyBase.GetFieldDeserializers()
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
    res["definition"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDefinition(res)
        }
        return nil
    }
    res["isOrganizationDefault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsOrganizationDefault(val)
        }
        return nil
    }
    return res
}
// GetIsOrganizationDefault gets the isOrganizationDefault property value. If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
func (m *StsPolicy) GetIsOrganizationDefault()(*bool) {
    return m.isOrganizationDefault
}
// Serialize serializes information the current object
func (m *StsPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.PolicyBase.Serialize(writer)
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
    if m.GetDefinition() != nil {
        err = writer.WriteCollectionOfStringValues("definition", m.GetDefinition())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOrganizationDefault", m.GetIsOrganizationDefault())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppliesTo sets the appliesTo property value. The appliesTo property
func (m *StsPolicy) SetAppliesTo(value []DirectoryObjectable)() {
    m.appliesTo = value
}
// SetDefinition sets the definition property value. A string collection containing a JSON string that defines the rules and settings for a policy. The syntax for the definition differs for each derived policy type. Required.
func (m *StsPolicy) SetDefinition(value []string)() {
    m.definition = value
}
// SetIsOrganizationDefault sets the isOrganizationDefault property value. If set to true, activates this policy. There can be many policies for the same policy type, but only one can be activated as the organization default. Optional, default value is false.
func (m *StsPolicy) SetIsOrganizationDefault(value *bool)() {
    m.isOrganizationDefault = value
}
