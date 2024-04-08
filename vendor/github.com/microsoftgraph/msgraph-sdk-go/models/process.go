package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Process 
type Process struct {
    // User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
    accountName *string
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The full process invocation commandline including all parameters.
    commandLine *string
    // Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Complex type containing file hashes (cryptographic and location-sensitive).
    fileHash FileHashable
    // The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
    integrityLevel *ProcessIntegrityLevel
    // True if the process is elevated.
    isElevated *bool
    // The name of the process' Image file.
    name *string
    // The OdataType property
    odataType *string
    // DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    parentProcessCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The Process ID (PID) of the parent process.
    parentProcessId *int32
    // The name of the image file of the parent process.
    parentProcessName *string
    // Full path, including filename.
    path *string
    // The Process ID (PID) of the process.
    processId *int32
}
// NewProcess instantiates a new process and sets the default values.
func NewProcess()(*Process) {
    m := &Process{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateProcessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProcessFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProcess(), nil
}
// GetAccountName gets the accountName property value. User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
func (m *Process) GetAccountName()(*string) {
    return m.accountName
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Process) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetCommandLine gets the commandLine property value. The full process invocation commandline including all parameters.
func (m *Process) GetCommandLine()(*string) {
    return m.commandLine
}
// GetCreatedDateTime gets the createdDateTime property value. Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.createdDateTime
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Process) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["accountName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["commandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommandLine(val)
        }
        return nil
    }
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
    res["fileHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileHashFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHash(val.(FileHashable))
        }
        return nil
    }
    res["integrityLevel"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProcessIntegrityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrityLevel(val.(*ProcessIntegrityLevel))
        }
        return nil
    }
    res["isElevated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsElevated(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
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
    res["parentProcessCreatedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessCreatedDateTime(val)
        }
        return nil
    }
    res["parentProcessId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessId(val)
        }
        return nil
    }
    res["parentProcessName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessName(val)
        }
        return nil
    }
    res["path"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
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
    return res
}
// GetFileHash gets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
func (m *Process) GetFileHash()(FileHashable) {
    return m.fileHash
}
// GetIntegrityLevel gets the integrityLevel property value. The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
func (m *Process) GetIntegrityLevel()(*ProcessIntegrityLevel) {
    return m.integrityLevel
}
// GetIsElevated gets the isElevated property value. True if the process is elevated.
func (m *Process) GetIsElevated()(*bool) {
    return m.isElevated
}
// GetName gets the name property value. The name of the process' Image file.
func (m *Process) GetName()(*string) {
    return m.name
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *Process) GetOdataType()(*string) {
    return m.odataType
}
// GetParentProcessCreatedDateTime gets the parentProcessCreatedDateTime property value. DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) GetParentProcessCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.parentProcessCreatedDateTime
}
// GetParentProcessId gets the parentProcessId property value. The Process ID (PID) of the parent process.
func (m *Process) GetParentProcessId()(*int32) {
    return m.parentProcessId
}
// GetParentProcessName gets the parentProcessName property value. The name of the image file of the parent process.
func (m *Process) GetParentProcessName()(*string) {
    return m.parentProcessName
}
// GetPath gets the path property value. Full path, including filename.
func (m *Process) GetPath()(*string) {
    return m.path
}
// GetProcessId gets the processId property value. The Process ID (PID) of the process.
func (m *Process) GetProcessId()(*int32) {
    return m.processId
}
// Serialize serializes information the current object
func (m *Process) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commandLine", m.GetCommandLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fileHash", m.GetFileHash())
        if err != nil {
            return err
        }
    }
    if m.GetIntegrityLevel() != nil {
        cast := (*m.GetIntegrityLevel()).String()
        err := writer.WriteStringValue("integrityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isElevated", m.GetIsElevated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
        err := writer.WriteTimeValue("parentProcessCreatedDateTime", m.GetParentProcessCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("parentProcessId", m.GetParentProcessId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("parentProcessName", m.GetParentProcessName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountName sets the accountName property value. User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
func (m *Process) SetAccountName(value *string)() {
    m.accountName = value
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Process) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetCommandLine sets the commandLine property value. The full process invocation commandline including all parameters.
func (m *Process) SetCommandLine(value *string)() {
    m.commandLine = value
}
// SetCreatedDateTime sets the createdDateTime property value. Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// SetFileHash sets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
func (m *Process) SetFileHash(value FileHashable)() {
    m.fileHash = value
}
// SetIntegrityLevel sets the integrityLevel property value. The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
func (m *Process) SetIntegrityLevel(value *ProcessIntegrityLevel)() {
    m.integrityLevel = value
}
// SetIsElevated sets the isElevated property value. True if the process is elevated.
func (m *Process) SetIsElevated(value *bool)() {
    m.isElevated = value
}
// SetName sets the name property value. The name of the process' Image file.
func (m *Process) SetName(value *string)() {
    m.name = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *Process) SetOdataType(value *string)() {
    m.odataType = value
}
// SetParentProcessCreatedDateTime sets the parentProcessCreatedDateTime property value. DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) SetParentProcessCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.parentProcessCreatedDateTime = value
}
// SetParentProcessId sets the parentProcessId property value. The Process ID (PID) of the parent process.
func (m *Process) SetParentProcessId(value *int32)() {
    m.parentProcessId = value
}
// SetParentProcessName sets the parentProcessName property value. The name of the image file of the parent process.
func (m *Process) SetParentProcessName(value *string)() {
    m.parentProcessName = value
}
// SetPath sets the path property value. Full path, including filename.
func (m *Process) SetPath(value *string)() {
    m.path = value
}
// SetProcessId sets the processId property value. The Process ID (PID) of the process.
func (m *Process) SetProcessId(value *int32)() {
    m.processId = value
}
