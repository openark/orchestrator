package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ManagedAppPolicy the ManagedAppPolicy resource represents a base type for platform specific policies.
type ManagedAppPolicy struct {
    Entity
    // The date and time the policy was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The policy's description.
    description *string
    // Policy display name.
    displayName *string
    // Last time the policy was modified.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Version of the entity.
    version *string
}
// NewManagedAppPolicy instantiates a new managedAppPolicy and sets the default values.
func NewManagedAppPolicy()(*ManagedAppPolicy) {
    m := &ManagedAppPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// CreateManagedAppPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateManagedAppPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "#microsoft.graph.androidManagedAppProtection":
                        return NewAndroidManagedAppProtection(), nil
                    case "#microsoft.graph.defaultManagedAppProtection":
                        return NewDefaultManagedAppProtection(), nil
                    case "#microsoft.graph.iosManagedAppProtection":
                        return NewIosManagedAppProtection(), nil
                    case "#microsoft.graph.managedAppConfiguration":
                        return NewManagedAppConfiguration(), nil
                    case "#microsoft.graph.managedAppProtection":
                        return NewManagedAppProtection(), nil
                    case "#microsoft.graph.mdmWindowsInformationProtectionPolicy":
                        return NewMdmWindowsInformationProtectionPolicy(), nil
                    case "#microsoft.graph.targetedManagedAppConfiguration":
                        return NewTargetedManagedAppConfiguration(), nil
                    case "#microsoft.graph.targetedManagedAppProtection":
                        return NewTargetedManagedAppProtection(), nil
                    case "#microsoft.graph.windowsInformationProtection":
                        return NewWindowsInformationProtection(), nil
                    case "#microsoft.graph.windowsInformationProtectionPolicy":
                        return NewWindowsInformationProtectionPolicy(), nil
                }
            }
        }
    }
    return NewManagedAppPolicy(), nil
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time the policy was created.
func (m *ManagedAppPolicy) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetDescription gets the description property value. The policy's description.
func (m *ManagedAppPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Policy display name.
func (m *ManagedAppPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
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
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Last time the policy was modified.
func (m *ManagedAppPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedDateTime
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedAppPolicy) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *ManagedAppPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
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
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time the policy was created.
func (m *ManagedAppPolicy) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetDescription sets the description property value. The policy's description.
func (m *ManagedAppPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Policy display name.
func (m *ManagedAppPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Last time the policy was modified.
func (m *ManagedAppPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedAppPolicy) SetVersion(value *string)() {
    m.version = value
}
