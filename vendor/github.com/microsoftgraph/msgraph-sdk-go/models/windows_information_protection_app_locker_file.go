package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WindowsInformationProtectionAppLockerFile windows Information Protection AppLocker File
type WindowsInformationProtectionAppLockerFile struct {
    Entity
    // The friendly name
    displayName *string
    // File as a byte array
    file []byte
    // SHA256 hash of the file
    fileHash *string
    // Version of the entity.
    version *string
}
// NewWindowsInformationProtectionAppLockerFile instantiates a new windowsInformationProtectionAppLockerFile and sets the default values.
func NewWindowsInformationProtectionAppLockerFile()(*WindowsInformationProtectionAppLockerFile) {
    m := &WindowsInformationProtectionAppLockerFile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateWindowsInformationProtectionAppLockerFileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWindowsInformationProtectionAppLockerFileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWindowsInformationProtectionAppLockerFile(), nil
}
// GetDisplayName gets the displayName property value. The friendly name
func (m *WindowsInformationProtectionAppLockerFile) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsInformationProtectionAppLockerFile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["file"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFile(val)
        }
        return nil
    }
    res["fileHash"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHash(val)
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
// GetFile gets the file property value. File as a byte array
func (m *WindowsInformationProtectionAppLockerFile) GetFile()([]byte) {
    return m.file
}
// GetFileHash gets the fileHash property value. SHA256 hash of the file
func (m *WindowsInformationProtectionAppLockerFile) GetFileHash()(*string) {
    return m.fileHash
}
// GetVersion gets the version property value. Version of the entity.
func (m *WindowsInformationProtectionAppLockerFile) GetVersion()(*string) {
    return m.version
}
// Serialize serializes information the current object
func (m *WindowsInformationProtectionAppLockerFile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    {
        err = writer.WriteByteArrayValue("file", m.GetFile())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("fileHash", m.GetFileHash())
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
// SetDisplayName sets the displayName property value. The friendly name
func (m *WindowsInformationProtectionAppLockerFile) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetFile sets the file property value. File as a byte array
func (m *WindowsInformationProtectionAppLockerFile) SetFile(value []byte)() {
    m.file = value
}
// SetFileHash sets the fileHash property value. SHA256 hash of the file
func (m *WindowsInformationProtectionAppLockerFile) SetFileHash(value *string)() {
    m.fileHash = value
}
// SetVersion sets the version property value. Version of the entity.
func (m *WindowsInformationProtectionAppLockerFile) SetVersion(value *string)() {
    m.version = value
}
