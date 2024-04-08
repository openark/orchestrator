package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Win32LobAppPowerShellScriptRule 
type Win32LobAppPowerShellScriptRule struct {
    Win32LobAppRule
    // The script output comparison value. Do not specify a value if the rule is used for detection.
    comparisonValue *string
    // The display name for the rule. Do not specify this value if the rule is used for detection.
    displayName *string
    // A value indicating whether a signature check is enforced.
    enforceSignatureCheck *bool
    // Contains all supported Powershell Script output detection type.
    operationType *Win32LobAppPowerShellScriptRuleOperationType
    // Contains properties for detection operator.
    operator *Win32LobAppRuleOperator
    // A value indicating whether the script should run as 32-bit.
    runAs32Bit *bool
    // The execution context of the script. Do not specify this value if the rule is used for detection. Script detection rules will run in the same context as the associated app install context. Possible values are: system, user.
    runAsAccount *RunAsAccountType
    // The base64-encoded script content.
    scriptContent *string
}
// NewWin32LobAppPowerShellScriptRule instantiates a new Win32LobAppPowerShellScriptRule and sets the default values.
func NewWin32LobAppPowerShellScriptRule()(*Win32LobAppPowerShellScriptRule) {
    m := &Win32LobAppPowerShellScriptRule{
        Win32LobAppRule: *NewWin32LobAppRule(),
    }
    odataTypeValue := "#microsoft.graph.win32LobAppPowerShellScriptRule"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateWin32LobAppPowerShellScriptRuleFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWin32LobAppPowerShellScriptRuleFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWin32LobAppPowerShellScriptRule(), nil
}
// GetComparisonValue gets the comparisonValue property value. The script output comparison value. Do not specify a value if the rule is used for detection.
func (m *Win32LobAppPowerShellScriptRule) GetComparisonValue()(*string) {
    return m.comparisonValue
}
// GetDisplayName gets the displayName property value. The display name for the rule. Do not specify this value if the rule is used for detection.
func (m *Win32LobAppPowerShellScriptRule) GetDisplayName()(*string) {
    return m.displayName
}
// GetEnforceSignatureCheck gets the enforceSignatureCheck property value. A value indicating whether a signature check is enforced.
func (m *Win32LobAppPowerShellScriptRule) GetEnforceSignatureCheck()(*bool) {
    return m.enforceSignatureCheck
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Win32LobAppPowerShellScriptRule) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Win32LobAppRule.GetFieldDeserializers()
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
    res["enforceSignatureCheck"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforceSignatureCheck(val)
        }
        return nil
    }
    res["operationType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseWin32LobAppPowerShellScriptRuleOperationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperationType(val.(*Win32LobAppPowerShellScriptRuleOperationType))
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
    res["runAs32Bit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAs32Bit(val)
        }
        return nil
    }
    res["runAsAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRunAsAccountType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRunAsAccount(val.(*RunAsAccountType))
        }
        return nil
    }
    res["scriptContent"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScriptContent(val)
        }
        return nil
    }
    return res
}
// GetOperationType gets the operationType property value. Contains all supported Powershell Script output detection type.
func (m *Win32LobAppPowerShellScriptRule) GetOperationType()(*Win32LobAppPowerShellScriptRuleOperationType) {
    return m.operationType
}
// GetOperator gets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppPowerShellScriptRule) GetOperator()(*Win32LobAppRuleOperator) {
    return m.operator
}
// GetRunAs32Bit gets the runAs32Bit property value. A value indicating whether the script should run as 32-bit.
func (m *Win32LobAppPowerShellScriptRule) GetRunAs32Bit()(*bool) {
    return m.runAs32Bit
}
// GetRunAsAccount gets the runAsAccount property value. The execution context of the script. Do not specify this value if the rule is used for detection. Script detection rules will run in the same context as the associated app install context. Possible values are: system, user.
func (m *Win32LobAppPowerShellScriptRule) GetRunAsAccount()(*RunAsAccountType) {
    return m.runAsAccount
}
// GetScriptContent gets the scriptContent property value. The base64-encoded script content.
func (m *Win32LobAppPowerShellScriptRule) GetScriptContent()(*string) {
    return m.scriptContent
}
// Serialize serializes information the current object
func (m *Win32LobAppPowerShellScriptRule) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Win32LobAppRule.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("comparisonValue", m.GetComparisonValue())
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
        err = writer.WriteBoolValue("enforceSignatureCheck", m.GetEnforceSignatureCheck())
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
        err = writer.WriteBoolValue("runAs32Bit", m.GetRunAs32Bit())
        if err != nil {
            return err
        }
    }
    if m.GetRunAsAccount() != nil {
        cast := (*m.GetRunAsAccount()).String()
        err = writer.WriteStringValue("runAsAccount", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("scriptContent", m.GetScriptContent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComparisonValue sets the comparisonValue property value. The script output comparison value. Do not specify a value if the rule is used for detection.
func (m *Win32LobAppPowerShellScriptRule) SetComparisonValue(value *string)() {
    m.comparisonValue = value
}
// SetDisplayName sets the displayName property value. The display name for the rule. Do not specify this value if the rule is used for detection.
func (m *Win32LobAppPowerShellScriptRule) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEnforceSignatureCheck sets the enforceSignatureCheck property value. A value indicating whether a signature check is enforced.
func (m *Win32LobAppPowerShellScriptRule) SetEnforceSignatureCheck(value *bool)() {
    m.enforceSignatureCheck = value
}
// SetOperationType sets the operationType property value. Contains all supported Powershell Script output detection type.
func (m *Win32LobAppPowerShellScriptRule) SetOperationType(value *Win32LobAppPowerShellScriptRuleOperationType)() {
    m.operationType = value
}
// SetOperator sets the operator property value. Contains properties for detection operator.
func (m *Win32LobAppPowerShellScriptRule) SetOperator(value *Win32LobAppRuleOperator)() {
    m.operator = value
}
// SetRunAs32Bit sets the runAs32Bit property value. A value indicating whether the script should run as 32-bit.
func (m *Win32LobAppPowerShellScriptRule) SetRunAs32Bit(value *bool)() {
    m.runAs32Bit = value
}
// SetRunAsAccount sets the runAsAccount property value. The execution context of the script. Do not specify this value if the rule is used for detection. Script detection rules will run in the same context as the associated app install context. Possible values are: system, user.
func (m *Win32LobAppPowerShellScriptRule) SetRunAsAccount(value *RunAsAccountType)() {
    m.runAsAccount = value
}
// SetScriptContent sets the scriptContent property value. The base64-encoded script content.
func (m *Win32LobAppPowerShellScriptRule) SetScriptContent(value *string)() {
    m.scriptContent = value
}
