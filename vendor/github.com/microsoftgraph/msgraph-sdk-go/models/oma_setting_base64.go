package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingBase64 
type OmaSettingBase64 struct {
    OmaSetting
    // File name associated with the Value property (.cer
    fileName *string
    // Value. (Base64 encoded string)
    value *string
}
// NewOmaSettingBase64 instantiates a new OmaSettingBase64 and sets the default values.
func NewOmaSettingBase64()(*OmaSettingBase64) {
    m := &OmaSettingBase64{
        OmaSetting: *NewOmaSetting(),
    }
    odataTypeValue := "#microsoft.graph.omaSettingBase64"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOmaSettingBase64FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingBase64FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingBase64(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingBase64) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OmaSetting.GetFieldDeserializers()
    res["fileName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileName(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetFileName gets the fileName property value. File name associated with the Value property (.cer
func (m *OmaSettingBase64) GetFileName()(*string) {
    return m.fileName
}
// GetValue gets the value property value. Value. (Base64 encoded string)
func (m *OmaSettingBase64) GetValue()(*string) {
    return m.value
}
// Serialize serializes information the current object
func (m *OmaSettingBase64) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OmaSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("fileName", m.GetFileName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileName sets the fileName property value. File name associated with the Value property (.cer
func (m *OmaSettingBase64) SetFileName(value *string)() {
    m.fileName = value
}
// SetValue sets the value property value. Value. (Base64 encoded string)
func (m *OmaSettingBase64) SetValue(value *string)() {
    m.value = value
}
