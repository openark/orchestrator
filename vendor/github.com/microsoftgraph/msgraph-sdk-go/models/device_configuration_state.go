package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceConfigurationState device Configuration State for a given device.
type DeviceConfigurationState struct {
    Entity
    // The name of the policy for this policyBase
    displayName *string
    // Supported platform types for policies.
    platformType *PolicyPlatformType
    // Count of how many setting a policy holds
    settingCount *int32
    // The settingStates property
    settingStates []DeviceConfigurationSettingStateable
    // The state property
    state *ComplianceStatus
    // The version of the policy
    version *int32
}
// NewDeviceConfigurationState instantiates a new deviceConfigurationState and sets the default values.
func NewDeviceConfigurationState()(*DeviceConfigurationState) {
    m := &DeviceConfigurationState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceConfigurationStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceConfigurationStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceConfigurationState(), nil
}
// GetDisplayName gets the displayName property value. The name of the policy for this policyBase
func (m *DeviceConfigurationState) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceConfigurationState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["platformType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParsePolicyPlatformType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatformType(val.(*PolicyPlatformType))
        }
        return nil
    }
    res["settingCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingCount(val)
        }
        return nil
    }
    res["settingStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceConfigurationSettingStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceConfigurationSettingStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceConfigurationSettingStateable)
            }
            m.SetSettingStates(res)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
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
// GetPlatformType gets the platformType property value. Supported platform types for policies.
func (m *DeviceConfigurationState) GetPlatformType()(*PolicyPlatformType) {
    return m.platformType
}
// GetSettingCount gets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceConfigurationState) GetSettingCount()(*int32) {
    return m.settingCount
}
// GetSettingStates gets the settingStates property value. The settingStates property
func (m *DeviceConfigurationState) GetSettingStates()([]DeviceConfigurationSettingStateable) {
    return m.settingStates
}
// GetState gets the state property value. The state property
func (m *DeviceConfigurationState) GetState()(*ComplianceStatus) {
    return m.state
}
// GetVersion gets the version property value. The version of the policy
func (m *DeviceConfigurationState) GetVersion()(*int32) {
    return m.version
}
// Serialize serializes information the current object
func (m *DeviceConfigurationState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetPlatformType() != nil {
        cast := (*m.GetPlatformType()).String()
        err = writer.WriteStringValue("platformType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("settingCount", m.GetSettingCount())
        if err != nil {
            return err
        }
    }
    if m.GetSettingStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSettingStates()))
        for i, v := range m.GetSettingStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("settingStates", cast)
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. The name of the policy for this policyBase
func (m *DeviceConfigurationState) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetPlatformType sets the platformType property value. Supported platform types for policies.
func (m *DeviceConfigurationState) SetPlatformType(value *PolicyPlatformType)() {
    m.platformType = value
}
// SetSettingCount sets the settingCount property value. Count of how many setting a policy holds
func (m *DeviceConfigurationState) SetSettingCount(value *int32)() {
    m.settingCount = value
}
// SetSettingStates sets the settingStates property value. The settingStates property
func (m *DeviceConfigurationState) SetSettingStates(value []DeviceConfigurationSettingStateable)() {
    m.settingStates = value
}
// SetState sets the state property value. The state property
func (m *DeviceConfigurationState) SetState(value *ComplianceStatus)() {
    m.state = value
}
// SetVersion sets the version property value. The version of the policy
func (m *DeviceConfigurationState) SetVersion(value *int32)() {
    m.version = value
}
