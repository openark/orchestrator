package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppRegistryRule 
type Win32LobAppRegistryRule struct {
    Win32LobAppRule
    // A value indicating whether to search the 32-bit registry on 64-bit systems.
    check32BitOn64System *bool
    // The registry comparison value.
    comparisonValue *string
    // The full path of the registry entry containing the value to detect.
    keyPath *string
    // Contains all supported registry data detection type.
    operationType *Win32LobAppRegistryRuleOperationType
    // Contains properties for detection operator.
    operator *Win32LobAppRuleOperator
    // The name of the registry value to detect.
    valueName *string
}
// NewWin32LobAppRegistryRule instantiates a new Win32LobAppRegistryRule and sets the default values.
func NewWin32LobAppRegistryRule()(*Win32LobAppRegistryRule) {
    m := &Win32LobAppRegistryRule{
        Win32LobAppRule: *NewWin32LobAppRule(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppRegistryRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppRegistryRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppRegistryRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppRegistryRule(), nil
}
// GetCheck32BitOn64System gets the check32BitOn64System property value. A value indicating whether to search the 32-bit registry on 64-bit systems.
func (m *Win32LobAppRegistryRule) GetCheck32BitOn64System()(*bool) {
    return m.check32BitOn64System
}
// GetComparisonValue gets the comparisonValue property value. The registry comparison value.
func (m *Win32LobAppRegistryRule) GetComparisonValue()(*string) {
    return m.comparisonValue
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppRegistryRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppRule.GetFieldDeserializers()
    res["check32BitOn64System"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheck32BitOn64System(val)
        }
        return nil
    }
    res["comparisonValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComparisonValue(val)
        }
        return nil
    }
    res["keyPath"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyPath(val)
        }
        return nil
    }
    res["operationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppRegistryRuleOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*Win32LobAppRegistryRuleOperationType))
        }
        return nil
    }
    res["operator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppRuleOperator)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperator(val.(*Win32LobAppRuleOperator))
        }
        return nil
    }
    res["valueName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueName(val)
        }
        return nil
    }
    return res
}
// GetKeyPath gets the keyPath property value. The full path of the registry entry containing the value to detect.
func (m *Win32LobAppRegistryRule) GetKeyPath()(*string) {
    return m.keyPath
}
// GetOperationType gets the operationType property value. Contains all supported registry data detection type.
func (m *Win32LobAppRegistryRule) GetOperationType()(*Win32LobAppRegistryRuleOperationType) {
    return m.operationType
}
// GetOperator gets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppRegistryRule) GetOperator()(*Win32LobAppRuleOperator) {
    return m.operator
}
// GetValueName gets the valueName property value. The name of the registry value to detect.
func (m *Win32LobAppRegistryRule) GetValueName()(*string) {
    return m.valueName
}
// Serialize serializes information the current object
func (m *Win32LobAppRegistryRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("check32BitOn64System", m.GetCheck32BitOn64System())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("comparisonValue", m.GetComparisonValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("keyPath", m.GetKeyPath())
        if err != nil {
            return err
        }
    }
    if m.GetOperationType() != nil {
        cast := (*m.GetOperationType()).String()
        err = writer.WriteStringValue("operationType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetOperator() != nil {
        cast := (*m.GetOperator()).String()
        err = writer.WriteStringValue("operator", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("valueName", m.GetValueName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCheck32BitOn64System sets the check32BitOn64System property value. A value indicating whether to search the 32-bit registry on 64-bit systems.
func (m *Win32LobAppRegistryRule) SetCheck32BitOn64System(value *bool)() {
    m.check32BitOn64System = value
}
// SetComparisonValue sets the comparisonValue property value. The registry comparison value.
func (m *Win32LobAppRegistryRule) SetComparisonValue(value *string)() {
    m.comparisonValue = value
}
// SetKeyPath sets the keyPath property value. The full path of the registry entry containing the value to detect.
func (m *Win32LobAppRegistryRule) SetKeyPath(value *string)() {
    m.keyPath = value
}
// SetOperationType sets the operationType property value. Contains all supported registry data detection type.
func (m *Win32LobAppRegistryRule) SetOperationType(value *Win32LobAppRegistryRuleOperationType)() {
    m.operationType = value
}
// SetOperator sets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppRegistryRule) SetOperator(value *Win32LobAppRuleOperator)() {
    m.operator = value
}
// SetValueName sets the valueName property value. The name of the registry value to detect.
func (m *Win32LobAppRegistryRule) SetValueName(value *string)() {
    m.valueName = value
}
