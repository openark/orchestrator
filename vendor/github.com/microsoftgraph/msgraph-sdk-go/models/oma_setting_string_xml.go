package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingStringXml 
type OmaSettingStringXml struct {
    OmaSetting
    // File name associated with the Value property (.xml).
    fileName *string
    // Value. (UTF8 encoded byte array)
    value []byte
}
// NewOmaSettingStringXml instantiates a new OmaSettingStringXml and sets the default values.
func NewOmaSettingStringXml()(*OmaSettingStringXml) {
    m := &OmaSettingStringXml{
        OmaSetting: *NewOmaSetting(),
    }
    odataTypeValue := "#microsoft.graph.omaSettingStringXml"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOmaSettingStringXmlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingStringXmlFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingStringXml(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingStringXml) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
        val, err := n.GetByteArrayValue()
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
// GetFileName gets the fileName property value. File name associated with the Value property (.xml).
func (m *OmaSettingStringXml) GetFileName()(*string) {
    return m.fileName
}
// GetValue gets the value property value. Value. (UTF8 encoded byte array)
func (m *OmaSettingStringXml) GetValue()([]byte) {
    return m.value
}
// Serialize serializes information the current object
func (m *OmaSettingStringXml) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        err = writer.WriteByteArrayValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFileName sets the fileName property value. File name associated with the Value property (.xml).
func (m *OmaSettingStringXml) SetFileName(value *string)() {
    m.fileName = value
}
// SetValue sets the value property value. Value. (UTF8 encoded byte array)
func (m *OmaSettingStringXml) SetValue(value []byte)() {
    m.value = value
}
