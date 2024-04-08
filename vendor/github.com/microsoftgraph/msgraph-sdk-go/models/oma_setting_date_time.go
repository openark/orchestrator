package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// OmaSettingDateTime 
type OmaSettingDateTime struct {
    OmaSetting
    // Value.
    value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
}
// NewOmaSettingDateTime instantiates a new OmaSettingDateTime and sets the default values.
func NewOmaSettingDateTime()(*OmaSettingDateTime) {
    m := &OmaSettingDateTime{
        OmaSetting: *NewOmaSetting(),
    }
    odataTypeValue := "#microsoft.graph.omaSettingDateTime"
    m.SetOdataType(&odataTypeValue)
    return m
}
// CreateOmaSettingDateTimeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOmaSettingDateTimeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOmaSettingDateTime(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OmaSettingDateTime) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OmaSetting.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
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
// GetValue gets the value property value. Value.
func (m *OmaSettingDateTime) GetValue()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.value
}
// Serialize serializes information the current object
func (m *OmaSettingDateTime) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OmaSetting.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. Value.
func (m *OmaSettingDateTime) SetValue(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.value = value
}
