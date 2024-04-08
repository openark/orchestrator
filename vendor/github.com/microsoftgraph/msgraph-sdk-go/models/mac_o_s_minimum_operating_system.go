package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MacOSMinimumOperatingSystem the minimum operating system required for a macOS app.
type MacOSMinimumOperatingSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // When TRUE, indicates OS X 10.10 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_10 *bool
    // When TRUE, indicates OS X 10.11 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_11 *bool
    // When TRUE, indicates macOS 10.12 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_12 *bool
    // When TRUE, indicates macOS 10.13 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_13 *bool
    // When TRUE, indicates macOS 10.14 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_14 *bool
    // When TRUE, indicates macOS 10.15 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_15 *bool
    // When TRUE, indicates Mac OS X 10.7 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_7 *bool
    // When TRUE, indicates OS X 10.8 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_8 *bool
    // When TRUE, indicates OS X 10.9 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v10_9 *bool
    // When TRUE, indicates macOS 11.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v11_0 *bool
    // When TRUE, indicates macOS 12.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v12_0 *bool
    // When TRUE, indicates macOS 13.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
    v13_0 *bool
}
// NewMacOSMinimumOperatingSystem instantiates a new macOSMinimumOperatingSystem and sets the default values.
func NewMacOSMinimumOperatingSystem()(*MacOSMinimumOperatingSystem) {
    m := &MacOSMinimumOperatingSystem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateMacOSMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMacOSMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMacOSMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MacOSMinimumOperatingSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MacOSMinimumOperatingSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
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
    res["v10_10"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV1010(val)
        }
        return nil
    }
    res["v10_11"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV1011(val)
        }
        return nil
    }
    res["v10_12"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV1012(val)
        }
        return nil
    }
    res["v10_13"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV1013(val)
        }
        return nil
    }
    res["v10_14"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV1014(val)
        }
        return nil
    }
    res["v10_15"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV1015(val)
        }
        return nil
    }
    res["v10_7"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV107(val)
        }
        return nil
    }
    res["v10_8"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV108(val)
        }
        return nil
    }
    res["v10_9"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV109(val)
        }
        return nil
    }
    res["v11_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV110(val)
        }
        return nil
    }
    res["v12_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV120(val)
        }
        return nil
    }
    res["v13_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV130(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *MacOSMinimumOperatingSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetV1010 gets the v10_10 property value. When TRUE, indicates OS X 10.10 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV1010()(*bool) {
    return m.v10_10
}
// GetV1011 gets the v10_11 property value. When TRUE, indicates OS X 10.11 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV1011()(*bool) {
    return m.v10_11
}
// GetV1012 gets the v10_12 property value. When TRUE, indicates macOS 10.12 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV1012()(*bool) {
    return m.v10_12
}
// GetV1013 gets the v10_13 property value. When TRUE, indicates macOS 10.13 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV1013()(*bool) {
    return m.v10_13
}
// GetV1014 gets the v10_14 property value. When TRUE, indicates macOS 10.14 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV1014()(*bool) {
    return m.v10_14
}
// GetV1015 gets the v10_15 property value. When TRUE, indicates macOS 10.15 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV1015()(*bool) {
    return m.v10_15
}
// GetV107 gets the v10_7 property value. When TRUE, indicates Mac OS X 10.7 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV107()(*bool) {
    return m.v10_7
}
// GetV108 gets the v10_8 property value. When TRUE, indicates OS X 10.8 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV108()(*bool) {
    return m.v10_8
}
// GetV109 gets the v10_9 property value. When TRUE, indicates OS X 10.9 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV109()(*bool) {
    return m.v10_9
}
// GetV110 gets the v11_0 property value. When TRUE, indicates macOS 11.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV110()(*bool) {
    return m.v11_0
}
// GetV120 gets the v12_0 property value. When TRUE, indicates macOS 12.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV120()(*bool) {
    return m.v12_0
}
// GetV130 gets the v13_0 property value. When TRUE, indicates macOS 13.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) GetV130()(*bool) {
    return m.v13_0
}
// Serialize serializes information the current object
func (m *MacOSMinimumOperatingSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_10", m.GetV1010())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_11", m.GetV1011())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_12", m.GetV1012())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_13", m.GetV1013())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_14", m.GetV1014())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_15", m.GetV1015())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_7", m.GetV107())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_8", m.GetV108())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_9", m.GetV109())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v11_0", m.GetV110())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v12_0", m.GetV120())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v13_0", m.GetV130())
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
func (m *MacOSMinimumOperatingSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *MacOSMinimumOperatingSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetV1010 sets the v10_10 property value. When TRUE, indicates OS X 10.10 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV1010(value *bool)() {
    m.v10_10 = value
}
// SetV1011 sets the v10_11 property value. When TRUE, indicates OS X 10.11 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV1011(value *bool)() {
    m.v10_11 = value
}
// SetV1012 sets the v10_12 property value. When TRUE, indicates macOS 10.12 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV1012(value *bool)() {
    m.v10_12 = value
}
// SetV1013 sets the v10_13 property value. When TRUE, indicates macOS 10.13 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV1013(value *bool)() {
    m.v10_13 = value
}
// SetV1014 sets the v10_14 property value. When TRUE, indicates macOS 10.14 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV1014(value *bool)() {
    m.v10_14 = value
}
// SetV1015 sets the v10_15 property value. When TRUE, indicates macOS 10.15 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV1015(value *bool)() {
    m.v10_15 = value
}
// SetV107 sets the v10_7 property value. When TRUE, indicates Mac OS X 10.7 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV107(value *bool)() {
    m.v10_7 = value
}
// SetV108 sets the v10_8 property value. When TRUE, indicates OS X 10.8 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV108(value *bool)() {
    m.v10_8 = value
}
// SetV109 sets the v10_9 property value. When TRUE, indicates OS X 10.9 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV109(value *bool)() {
    m.v10_9 = value
}
// SetV110 sets the v11_0 property value. When TRUE, indicates macOS 11.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV110(value *bool)() {
    m.v11_0 = value
}
// SetV120 sets the v12_0 property value. When TRUE, indicates macOS 12.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV120(value *bool)() {
    m.v12_0 = value
}
// SetV130 sets the v13_0 property value. When TRUE, indicates macOS 13.0 or later is required to install the app. When FALSE, indicates some other OS version is the minimum OS to install the app. Default value is FALSE.
func (m *MacOSMinimumOperatingSystem) SetV130(value *bool)() {
    m.v13_0 = value
}
