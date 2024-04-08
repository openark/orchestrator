package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PolicyBase 
type PolicyBase struct {
    DirectoryObject
    // Description for this policy. Required.
    description *string
    // Display name for this policy. Required.
    displayName *string
}
// NewPolicyBase instantiates a new policyBase and sets the default values.
func NewPolicyBase()(*PolicyBase) {
    m := &PolicyBase{
        DirectoryObject: *NewDirectoryObject(),
    }
    odataTypeValue := "#microsoft.graph.policyBase"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreatePolicyBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePolicyBaseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.authorizationPolicy":
                        return NewAuthorizationPolicy(), nil
                    case "#microsoft.graph.claimsMappingPolicy":
                        return NewClaimsMappingPolicy(), nil
                    case "#microsoft.graph.crossTenantAccessPolicy":
                        return NewCrossTenantAccessPolicy(), nil
                    case "#microsoft.graph.homeRealmDiscoveryPolicy":
                        return NewHomeRealmDiscoveryPolicy(), nil
                    case "#microsoft.graph.identitySecurityDefaultsEnforcementPolicy":
                        return NewIdentitySecurityDefaultsEnforcementPolicy(), nil
                    case "#microsoft.graph.permissionGrantPolicy":
                        return NewPermissionGrantPolicy(), nil
                    case "#microsoft.graph.stsPolicy":
                        return NewStsPolicy(), nil
                    case "#microsoft.graph.tokenIssuancePolicy":
                        return NewTokenIssuancePolicy(), nil
                    case "#microsoft.graph.tokenLifetimePolicy":
                        return NewTokenLifetimePolicy(), nil
                }
            }
        }
    }
    return NewPolicyBase(), nil
}
// GetDescription gets the description property value. Description for this policy. Required.
func (m *PolicyBase) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Display name for this policy. Required.
func (m *PolicyBase) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PolicyBase) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryObject.GetFieldDeserializers()
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
    return res
}
// Serialize serializes information the current object
func (m *PolicyBase) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryObject.Serialize(writer)
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Description for this policy. Required.
func (m *PolicyBase) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Display name for this policy. Required.
func (m *PolicyBase) SetDisplayName(value *string)() {
    m.displayName = value
}
