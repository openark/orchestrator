package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ImportedWindowsAutopilotDeviceIdentity imported windows autopilot devices.
type ImportedWindowsAutopilotDeviceIdentity struct {
    Entity
    // UPN of the user the device will be assigned
    assignedUserPrincipalName *string
    // Group Tag of the Windows autopilot device.
    groupTag *string
    // Hardware Blob of the Windows autopilot device.
    hardwareIdentifier []byte
    // The Import Id of the Windows autopilot device.
    importId *string
    // Product Key of the Windows autopilot device.
    productKey *string
    // Serial number of the Windows autopilot device.
    serialNumber *string
    // Current state of the imported device.
    state ImportedWindowsAutopilotDeviceIdentityStateable
}
// NewImportedWindowsAutopilotDeviceIdentity instantiates a new importedWindowsAutopilotDeviceIdentity and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentity()(*ImportedWindowsAutopilotDeviceIdentity) {
    m := &ImportedWindowsAutopilotDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateImportedWindowsAutopilotDeviceIdentityFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedWindowsAutopilotDeviceIdentity(), nil
}
// GetAssignedUserPrincipalName gets the assignedUserPrincipalName property value. UPN of the user the device will be assigned
func (m *ImportedWindowsAutopilotDeviceIdentity) GetAssignedUserPrincipalName()(*string) {
    return m.assignedUserPrincipalName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedWindowsAutopilotDeviceIdentity) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedUserPrincipalName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedUserPrincipalName(val)
        }
        return nil
    }
    res["groupTag"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupTag(val)
        }
        return nil
    }
    res["hardwareIdentifier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareIdentifier(val)
        }
        return nil
    }
    res["importId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportId(val)
        }
        return nil
    }
    res["productKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["serialNumber"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateImportedWindowsAutopilotDeviceIdentityStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(ImportedWindowsAutopilotDeviceIdentityStateable))
        }
        return nil
    }
    return res
}
// GetGroupTag gets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetGroupTag()(*string) {
    return m.groupTag
}
// GetHardwareIdentifier gets the hardwareIdentifier property value. Hardware Blob of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetHardwareIdentifier()([]byte) {
    return m.hardwareIdentifier
}
// GetImportId gets the importId property value. The Import Id of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetImportId()(*string) {
    return m.importId
}
// GetProductKey gets the productKey property value. Product Key of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetProductKey()(*string) {
    return m.productKey
}
// GetSerialNumber gets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetSerialNumber()(*string) {
    return m.serialNumber
}
// GetState gets the state property value. Current state of the imported device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetState()(ImportedWindowsAutopilotDeviceIdentityStateable) {
    return m.state
}
// Serialize serializes information the current object
func (m *ImportedWindowsAutopilotDeviceIdentity) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedUserPrincipalName", m.GetAssignedUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupTag", m.GetGroupTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("hardwareIdentifier", m.GetHardwareIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("importId", m.GetImportId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAssignedUserPrincipalName sets the assignedUserPrincipalName property value. UPN of the user the device will be assigned
func (m *ImportedWindowsAutopilotDeviceIdentity) SetAssignedUserPrincipalName(value *string)() {
    m.assignedUserPrincipalName = value
}
// SetGroupTag sets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetGroupTag(value *string)() {
    m.groupTag = value
}
// SetHardwareIdentifier sets the hardwareIdentifier property value. Hardware Blob of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetHardwareIdentifier(value []byte)() {
    m.hardwareIdentifier = value
}
// SetImportId sets the importId property value. The Import Id of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetImportId(value *string)() {
    m.importId = value
}
// SetProductKey sets the productKey property value. Product Key of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetProductKey(value *string)() {
    m.productKey = value
}
// SetSerialNumber sets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// SetState sets the state property value. Current state of the imported device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetState(value ImportedWindowsAutopilotDeviceIdentityStateable)() {
    m.state = value
}
