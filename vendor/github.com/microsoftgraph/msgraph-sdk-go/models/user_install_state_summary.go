package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// UserInstallStateSummary contains properties for the installation state summary for a user.
type UserInstallStateSummary struct {
    Entity
    // The install state of the eBook.
    deviceStates []DeviceInstallStateable
    // Failed Device Count.
    failedDeviceCount *int32
    // Installed Device Count.
    installedDeviceCount *int32
    // Not installed device count.
    notInstalledDeviceCount *int32
    // User name.
    userName *string
}
// NewUserInstallStateSummary instantiates a new userInstallStateSummary and sets the default values.
func NewUserInstallStateSummary()(*UserInstallStateSummary) {
    m := &UserInstallStateSummary{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserInstallStateSummaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserInstallStateSummaryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewUserInstallStateSummary(), nil
}
// GetDeviceStates gets the deviceStates property value. The install state of the eBook.
func (m *UserInstallStateSummary) GetDeviceStates()([]DeviceInstallStateable) {
    return m.deviceStates
}
// GetFailedDeviceCount gets the failedDeviceCount property value. Failed Device Count.
func (m *UserInstallStateSummary) GetFailedDeviceCount()(*int32) {
    return m.failedDeviceCount
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserInstallStateSummary) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceStates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateDeviceInstallStateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DeviceInstallStateable, len(val))
            for i, v := range val {
                res[i] = v.(DeviceInstallStateable)
            }
            m.SetDeviceStates(res)
        }
        return nil
    }
    res["failedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailedDeviceCount(val)
        }
        return nil
    }
    res["installedDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstalledDeviceCount(val)
        }
        return nil
    }
    res["notInstalledDeviceCount"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotInstalledDeviceCount(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
// GetInstalledDeviceCount gets the installedDeviceCount property value. Installed Device Count.
func (m *UserInstallStateSummary) GetInstalledDeviceCount()(*int32) {
    return m.installedDeviceCount
}
// GetNotInstalledDeviceCount gets the notInstalledDeviceCount property value. Not installed device count.
func (m *UserInstallStateSummary) GetNotInstalledDeviceCount()(*int32) {
    return m.notInstalledDeviceCount
}
// GetUserName gets the userName property value. User name.
func (m *UserInstallStateSummary) GetUserName()(*string) {
    return m.userName
}
// Serialize serializes information the current object
func (m *UserInstallStateSummary) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetDeviceStates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetDeviceStates()))
        for i, v := range m.GetDeviceStates() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("deviceStates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("failedDeviceCount", m.GetFailedDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("installedDeviceCount", m.GetInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("notInstalledDeviceCount", m.GetNotInstalledDeviceCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDeviceStates sets the deviceStates property value. The install state of the eBook.
func (m *UserInstallStateSummary) SetDeviceStates(value []DeviceInstallStateable)() {
    m.deviceStates = value
}
// SetFailedDeviceCount sets the failedDeviceCount property value. Failed Device Count.
func (m *UserInstallStateSummary) SetFailedDeviceCount(value *int32)() {
    m.failedDeviceCount = value
}
// SetInstalledDeviceCount sets the installedDeviceCount property value. Installed Device Count.
func (m *UserInstallStateSummary) SetInstalledDeviceCount(value *int32)() {
    m.installedDeviceCount = value
}
// SetNotInstalledDeviceCount sets the notInstalledDeviceCount property value. Not installed device count.
func (m *UserInstallStateSummary) SetNotInstalledDeviceCount(value *int32)() {
    m.notInstalledDeviceCount = value
}
// SetUserName sets the userName property value. User name.
func (m *UserInstallStateSummary) SetUserName(value *string)() {
    m.userName = value
}
