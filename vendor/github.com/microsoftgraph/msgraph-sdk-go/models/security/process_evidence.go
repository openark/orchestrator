package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProcessEvidence 
type ProcessEvidence struct {
    AlertEvidence
    // The status of the detection.The possible values are: detected, blocked, prevented, unknownFutureValue.
    detectionStatus *DetectionStatus
    // Image file details.
    imageFile FileDetailsable
    // A unique identifier assigned to a device by Microsoft Defender for Endpoint.
    mdeDeviceId *string
    // Date and time when the parent of the process was created.
    parentProcessCreationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Process ID (PID) of the parent process that spawned the process.
    parentProcessId *int64
    // Parent process image file details.
    parentProcessImageFile FileDetailsable
    // Command line used to create the new process.
    processCommandLine *string
    // Date and time the process was created.
    processCreationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // Process ID (PID) of the newly created process.
    processId *int64
    // User details of the user that ran the process.
    userAccount UserAccountable
}
// NewProcessEvidence instantiates a new ProcessEvidence and sets the default values.
func NewProcessEvidence()(*ProcessEvidence) {
    m := &ProcessEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateProcessEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProcessEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProcessEvidence(), nil
}
// GetDetectionStatus gets the detectionStatus property value. The status of the detection.The possible values are: detected, blocked, prevented, unknownFutureValue.
func (m *ProcessEvidence) GetDetectionStatus()(*DetectionStatus) {
    return m.detectionStatus
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProcessEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["detectionStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDetectionStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDetectionStatus(val.(*DetectionStatus))
        }
        return nil
    }
    res["imageFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImageFile(val.(FileDetailsable))
        }
        return nil
    }
    res["mdeDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdeDeviceId(val)
        }
        return nil
    }
    res["parentProcessCreationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessCreationDateTime(val)
        }
        return nil
    }
    res["parentProcessId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessId(val)
        }
        return nil
    }
    res["parentProcessImageFile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileDetailsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessImageFile(val.(FileDetailsable))
        }
        return nil
    }
    res["processCommandLine"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessCommandLine(val)
        }
        return nil
    }
    res["processCreationDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessCreationDateTime(val)
        }
        return nil
    }
    res["processId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessId(val)
        }
        return nil
    }
    res["userAccount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateUserAccountFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAccount(val.(UserAccountable))
        }
        return nil
    }
    return res
}
// GetImageFile gets the imageFile property value. Image file details.
func (m *ProcessEvidence) GetImageFile()(FileDetailsable) {
    return m.imageFile
}
// GetMdeDeviceId gets the mdeDeviceId property value. A unique identifier assigned to a device by Microsoft Defender for Endpoint.
func (m *ProcessEvidence) GetMdeDeviceId()(*string) {
    return m.mdeDeviceId
}
// GetParentProcessCreationDateTime gets the parentProcessCreationDateTime property value. Date and time when the parent of the process was created.
func (m *ProcessEvidence) GetParentProcessCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.parentProcessCreationDateTime
}
// GetParentProcessId gets the parentProcessId property value. Process ID (PID) of the parent process that spawned the process.
func (m *ProcessEvidence) GetParentProcessId()(*int64) {
    return m.parentProcessId
}
// GetParentProcessImageFile gets the parentProcessImageFile property value. Parent process image file details.
func (m *ProcessEvidence) GetParentProcessImageFile()(FileDetailsable) {
    return m.parentProcessImageFile
}
// GetProcessCommandLine gets the processCommandLine property value. Command line used to create the new process.
func (m *ProcessEvidence) GetProcessCommandLine()(*string) {
    return m.processCommandLine
}
// GetProcessCreationDateTime gets the processCreationDateTime property value. Date and time the process was created.
func (m *ProcessEvidence) GetProcessCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.processCreationDateTime
}
// GetProcessId gets the processId property value. Process ID (PID) of the newly created process.
func (m *ProcessEvidence) GetProcessId()(*int64) {
    return m.processId
}
// GetUserAccount gets the userAccount property value. User details of the user that ran the process.
func (m *ProcessEvidence) GetUserAccount()(UserAccountable) {
    return m.userAccount
}
// Serialize serializes information the current object
func (m *ProcessEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDetectionStatus() != nil {
        cast := (*m.GetDetectionStatus()).String()
        err = writer.WriteStringValue("detectionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("imageFile", m.GetImageFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdeDeviceId", m.GetMdeDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("parentProcessCreationDateTime", m.GetParentProcessCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("parentProcessId", m.GetParentProcessId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parentProcessImageFile", m.GetParentProcessImageFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("processCommandLine", m.GetProcessCommandLine())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("processCreationDateTime", m.GetProcessCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("processId", m.GetProcessId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userAccount", m.GetUserAccount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDetectionStatus sets the detectionStatus property value. The status of the detection.The possible values are: detected, blocked, prevented, unknownFutureValue.
func (m *ProcessEvidence) SetDetectionStatus(value *DetectionStatus)() {
    m.detectionStatus = value
}
// SetImageFile sets the imageFile property value. Image file details.
func (m *ProcessEvidence) SetImageFile(value FileDetailsable)() {
    m.imageFile = value
}
// SetMdeDeviceId sets the mdeDeviceId property value. A unique identifier assigned to a device by Microsoft Defender for Endpoint.
func (m *ProcessEvidence) SetMdeDeviceId(value *string)() {
    m.mdeDeviceId = value
}
// SetParentProcessCreationDateTime sets the parentProcessCreationDateTime property value. Date and time when the parent of the process was created.
func (m *ProcessEvidence) SetParentProcessCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.parentProcessCreationDateTime = value
}
// SetParentProcessId sets the parentProcessId property value. Process ID (PID) of the parent process that spawned the process.
func (m *ProcessEvidence) SetParentProcessId(value *int64)() {
    m.parentProcessId = value
}
// SetParentProcessImageFile sets the parentProcessImageFile property value. Parent process image file details.
func (m *ProcessEvidence) SetParentProcessImageFile(value FileDetailsable)() {
    m.parentProcessImageFile = value
}
// SetProcessCommandLine sets the processCommandLine property value. Command line used to create the new process.
func (m *ProcessEvidence) SetProcessCommandLine(value *string)() {
    m.processCommandLine = value
}
// SetProcessCreationDateTime sets the processCreationDateTime property value. Date and time the process was created.
func (m *ProcessEvidence) SetProcessCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.processCreationDateTime = value
}
// SetProcessId sets the processId property value. Process ID (PID) of the newly created process.
func (m *ProcessEvidence) SetProcessId(value *int64)() {
    m.processId = value
}
// SetUserAccount sets the userAccount property value. User details of the user that ran the process.
func (m *ProcessEvidence) SetUserAccount(value UserAccountable)() {
    m.userAccount = value
}
