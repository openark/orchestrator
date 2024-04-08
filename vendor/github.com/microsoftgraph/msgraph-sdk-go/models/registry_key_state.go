package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// RegistryKeyState 
type RegistryKeyState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // A Windows registry hive : HKEY_CURRENT_CONFIG HKEY_CURRENT_USER HKEY_LOCAL_MACHINE/SAM HKEY_LOCAL_MACHINE/Security HKEY_LOCAL_MACHINE/Software HKEY_LOCAL_MACHINE/System HKEY_USERS/.Default. Possible values are: unknown, currentConfig, currentUser, localMachineSam, localMachineSecurity, localMachineSoftware, localMachineSystem, usersDefault.
    hive *RegistryHive
    // Current (i.e. changed) registry key (excludes HIVE).
    key *string
    // The OdataType property
    odataType *string
    // Previous (i.e. before changed) registry key (excludes HIVE).
    oldKey *string
    // Previous (i.e. before changed) registry key value data (contents).
    oldValueData *string
    // Previous (i.e. before changed) registry key value name.
    oldValueName *string
    // Operation that changed the registry key name and/or value. Possible values are: unknown, create, modify, delete.
    operation *RegistryOperation
    // Process ID (PID) of the process that modified the registry key (process details will appear in the alert 'processes' collection).
    processId *int32
    // Current (i.e. changed) registry key value data (contents).
    valueData *string
    // Current (i.e. changed) registry key value name
    valueName *string
    // Registry key value type REG_BINARY REG_DWORD REG_DWORD_LITTLE_ENDIAN REG_DWORD_BIG_ENDIANREG_EXPAND_SZ REG_LINK REG_MULTI_SZ REG_NONE REG_QWORD REG_QWORD_LITTLE_ENDIAN REG_SZ Possible values are: unknown, binary, dword, dwordLittleEndian, dwordBigEndian, expandSz, link, multiSz, none, qword, qwordlittleEndian, sz.
    valueType *RegistryValueType
}
// NewRegistryKeyState instantiates a new registryKeyState and sets the default values.
func NewRegistryKeyState()(*RegistryKeyState) {
    m := &RegistryKeyState{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateRegistryKeyStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateRegistryKeyStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRegistryKeyState(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegistryKeyState) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegistryKeyState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["hive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryHive)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHive(val.(*RegistryHive))
        }
        return nil
    }
    res["key"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["oldKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldKey(val)
        }
        return nil
    }
    res["oldValueData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldValueData(val)
        }
        return nil
    }
    res["oldValueName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldValueName(val)
        }
        return nil
    }
    res["operation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryOperation)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOperation(val.(*RegistryOperation))
        }
        return nil
    }
    res["processId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessId(val)
        }
        return nil
    }
    res["valueData"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueData(val)
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
    res["valueType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryValueType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueType(val.(*RegistryValueType))
        }
        return nil
    }
    return res
}
// GetHive gets the hive property value. A Windows registry hive : HKEY_CURRENT_CONFIG HKEY_CURRENT_USER HKEY_LOCAL_MACHINE/SAM HKEY_LOCAL_MACHINE/Security HKEY_LOCAL_MACHINE/Software HKEY_LOCAL_MACHINE/System HKEY_USERS/.Default. Possible values are: unknown, currentConfig, currentUser, localMachineSam, localMachineSecurity, localMachineSoftware, localMachineSystem, usersDefault.
func (m *RegistryKeyState) GetHive()(*RegistryHive) {
    return m.hive
}
// GetKey gets the key property value. Current (i.e. changed) registry key (excludes HIVE).
func (m *RegistryKeyState) GetKey()(*string) {
    return m.key
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *RegistryKeyState) GetOdataType()(*string) {
    return m.odataType
}
// GetOldKey gets the oldKey property value. Previous (i.e. before changed) registry key (excludes HIVE).
func (m *RegistryKeyState) GetOldKey()(*string) {
    return m.oldKey
}
// GetOldValueData gets the oldValueData property value. Previous (i.e. before changed) registry key value data (contents).
func (m *RegistryKeyState) GetOldValueData()(*string) {
    return m.oldValueData
}
// GetOldValueName gets the oldValueName property value. Previous (i.e. before changed) registry key value name.
func (m *RegistryKeyState) GetOldValueName()(*string) {
    return m.oldValueName
}
// GetOperation gets the operation property value. Operation that changed the registry key name and/or value. Possible values are: unknown, create, modify, delete.
func (m *RegistryKeyState) GetOperation()(*RegistryOperation) {
    return m.operation
}
// GetProcessId gets the processId property value. Process ID (PID) of the process that modified the registry key (process details will appear in the alert 'processes' collection).
func (m *RegistryKeyState) GetProcessId()(*int32) {
    return m.processId
}
// GetValueData gets the valueData property value. Current (i.e. changed) registry key value data (contents).
func (m *RegistryKeyState) GetValueData()(*string) {
    return m.valueData
}
// GetValueName gets the valueName property value. Current (i.e. changed) registry key value name
func (m *RegistryKeyState) GetValueName()(*string) {
    return m.valueName
}
// GetValueType gets the valueType property value. Registry key value type REG_BINARY REG_DWORD REG_DWORD_LITTLE_ENDIAN REG_DWORD_BIG_ENDIANREG_EXPAND_SZ REG_LINK REG_MULTI_SZ REG_NONE REG_QWORD REG_QWORD_LITTLE_ENDIAN REG_SZ Possible values are: unknown, binary, dword, dwordLittleEndian, dwordBigEndian, expandSz, link, multiSz, none, qword, qwordlittleEndian, sz.
func (m *RegistryKeyState) GetValueType()(*RegistryValueType) {
    return m.valueType
}
// Serialize serializes information the current object
func (m *RegistryKeyState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetHive() != nil {
        cast := (*m.GetHive()).String()
        err := writer.WriteStringValue("hive", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldKey", m.GetOldKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldValueData", m.GetOldValueData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldValueName", m.GetOldValueName())
        if err != nil {
            return err
        }
    }
    if m.GetOperation() != nil {
        cast := (*m.GetOperation()).String()
        err := writer.WriteStringValue("operation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("processId", m.GetProcessId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("valueData", m.GetValueData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("valueName", m.GetValueName())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := (*m.GetValueType()).String()
        err := writer.WriteStringValue("valueType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegistryKeyState) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetHive sets the hive property value. A Windows registry hive : HKEY_CURRENT_CONFIG HKEY_CURRENT_USER HKEY_LOCAL_MACHINE/SAM HKEY_LOCAL_MACHINE/Security HKEY_LOCAL_MACHINE/Software HKEY_LOCAL_MACHINE/System HKEY_USERS/.Default. Possible values are: unknown, currentConfig, currentUser, localMachineSam, localMachineSecurity, localMachineSoftware, localMachineSystem, usersDefault.
func (m *RegistryKeyState) SetHive(value *RegistryHive)() {
    m.hive = value
}
// SetKey sets the key property value. Current (i.e. changed) registry key (excludes HIVE).
func (m *RegistryKeyState) SetKey(value *string)() {
    m.key = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *RegistryKeyState) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOldKey sets the oldKey property value. Previous (i.e. before changed) registry key (excludes HIVE).
func (m *RegistryKeyState) SetOldKey(value *string)() {
    m.oldKey = value
}
// SetOldValueData sets the oldValueData property value. Previous (i.e. before changed) registry key value data (contents).
func (m *RegistryKeyState) SetOldValueData(value *string)() {
    m.oldValueData = value
}
// SetOldValueName sets the oldValueName property value. Previous (i.e. before changed) registry key value name.
func (m *RegistryKeyState) SetOldValueName(value *string)() {
    m.oldValueName = value
}
// SetOperation sets the operation property value. Operation that changed the registry key name and/or value. Possible values are: unknown, create, modify, delete.
func (m *RegistryKeyState) SetOperation(value *RegistryOperation)() {
    m.operation = value
}
// SetProcessId sets the processId property value. Process ID (PID) of the process that modified the registry key (process details will appear in the alert 'processes' collection).
func (m *RegistryKeyState) SetProcessId(value *int32)() {
    m.processId = value
}
// SetValueData sets the valueData property value. Current (i.e. changed) registry key value data (contents).
func (m *RegistryKeyState) SetValueData(value *string)() {
    m.valueData = value
}
// SetValueName sets the valueName property value. Current (i.e. changed) registry key value name
func (m *RegistryKeyState) SetValueName(value *string)() {
    m.valueName = value
}
// SetValueType sets the valueType property value. Registry key value type REG_BINARY REG_DWORD REG_DWORD_LITTLE_ENDIAN REG_DWORD_BIG_ENDIANREG_EXPAND_SZ REG_LINK REG_MULTI_SZ REG_NONE REG_QWORD REG_QWORD_LITTLE_ENDIAN REG_SZ Possible values are: unknown, binary, dword, dwordLittleEndian, dwordBigEndian, expandSz, link, multiSz, none, qword, qwordlittleEndian, sz.
func (m *RegistryKeyState) SetValueType(value *RegistryValueType)() {
    m.valueType = value
}
