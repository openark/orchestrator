package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AndroidMinimumOperatingSystem contains properties for the minimum operating system required for an Android mobile app.
type AndroidMinimumOperatingSystem struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v10_0 *bool
    // When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v11_0 *bool
    // When TRUE, only Version 4.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_0 *bool
    // When TRUE, only Version 4.0.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_0_3 *bool
    // When TRUE, only Version 4.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_1 *bool
    // When TRUE, only Version 4.2 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_2 *bool
    // When TRUE, only Version 4.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_3 *bool
    // When TRUE, only Version 4.4 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v4_4 *bool
    // When TRUE, only Version 5.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v5_0 *bool
    // When TRUE, only Version 5.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v5_1 *bool
    // When TRUE, only Version 6.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v6_0 *bool
    // When TRUE, only Version 7.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v7_0 *bool
    // When TRUE, only Version 7.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v7_1 *bool
    // When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v8_0 *bool
    // When TRUE, only Version 8.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v8_1 *bool
    // When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
    v9_0 *bool
}
// NewAndroidMinimumOperatingSystem instantiates a new androidMinimumOperatingSystem and sets the default values.
func NewAndroidMinimumOperatingSystem()(*AndroidMinimumOperatingSystem) {
    m := &AndroidMinimumOperatingSystem{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateAndroidMinimumOperatingSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAndroidMinimumOperatingSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAndroidMinimumOperatingSystem(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AndroidMinimumOperatingSystem) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AndroidMinimumOperatingSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["v10_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV100(val)
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
    res["v4_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV40(val)
        }
        return nil
    }
    res["v4_0_3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV403(val)
        }
        return nil
    }
    res["v4_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV41(val)
        }
        return nil
    }
    res["v4_2"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV42(val)
        }
        return nil
    }
    res["v4_3"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV43(val)
        }
        return nil
    }
    res["v4_4"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV44(val)
        }
        return nil
    }
    res["v5_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV50(val)
        }
        return nil
    }
    res["v5_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV51(val)
        }
        return nil
    }
    res["v6_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV60(val)
        }
        return nil
    }
    res["v7_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV70(val)
        }
        return nil
    }
    res["v7_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV71(val)
        }
        return nil
    }
    res["v8_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV80(val)
        }
        return nil
    }
    res["v8_1"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV81(val)
        }
        return nil
    }
    res["v9_0"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetV90(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AndroidMinimumOperatingSystem) GetOdataType()(*string) {
    return m.odataType
}
// GetV100 gets the v10_0 property value. When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV100()(*bool) {
    return m.v10_0
}
// GetV110 gets the v11_0 property value. When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV110()(*bool) {
    return m.v11_0
}
// GetV40 gets the v4_0 property value. When TRUE, only Version 4.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV40()(*bool) {
    return m.v4_0
}
// GetV403 gets the v4_0_3 property value. When TRUE, only Version 4.0.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV403()(*bool) {
    return m.v4_0_3
}
// GetV41 gets the v4_1 property value. When TRUE, only Version 4.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV41()(*bool) {
    return m.v4_1
}
// GetV42 gets the v4_2 property value. When TRUE, only Version 4.2 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV42()(*bool) {
    return m.v4_2
}
// GetV43 gets the v4_3 property value. When TRUE, only Version 4.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV43()(*bool) {
    return m.v4_3
}
// GetV44 gets the v4_4 property value. When TRUE, only Version 4.4 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV44()(*bool) {
    return m.v4_4
}
// GetV50 gets the v5_0 property value. When TRUE, only Version 5.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV50()(*bool) {
    return m.v5_0
}
// GetV51 gets the v5_1 property value. When TRUE, only Version 5.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV51()(*bool) {
    return m.v5_1
}
// GetV60 gets the v6_0 property value. When TRUE, only Version 6.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV60()(*bool) {
    return m.v6_0
}
// GetV70 gets the v7_0 property value. When TRUE, only Version 7.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV70()(*bool) {
    return m.v7_0
}
// GetV71 gets the v7_1 property value. When TRUE, only Version 7.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV71()(*bool) {
    return m.v7_1
}
// GetV80 gets the v8_0 property value. When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV80()(*bool) {
    return m.v8_0
}
// GetV81 gets the v8_1 property value. When TRUE, only Version 8.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV81()(*bool) {
    return m.v8_1
}
// GetV90 gets the v9_0 property value. When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) GetV90()(*bool) {
    return m.v9_0
}
// Serialize serializes information the current object
func (m *AndroidMinimumOperatingSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v10_0", m.GetV100())
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
        err := writer.WriteBoolValue("v4_0", m.GetV40())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_0_3", m.GetV403())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_1", m.GetV41())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_2", m.GetV42())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_3", m.GetV43())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v4_4", m.GetV44())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v5_0", m.GetV50())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v5_1", m.GetV51())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v6_0", m.GetV60())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v7_0", m.GetV70())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v7_1", m.GetV71())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_0", m.GetV80())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v8_1", m.GetV81())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("v9_0", m.GetV90())
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
func (m *AndroidMinimumOperatingSystem) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AndroidMinimumOperatingSystem) SetOdataType(value *string)() {
    m.odataType = value
}
// SetV100 sets the v10_0 property value. When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV100(value *bool)() {
    m.v10_0 = value
}
// SetV110 sets the v11_0 property value. When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV110(value *bool)() {
    m.v11_0 = value
}
// SetV40 sets the v4_0 property value. When TRUE, only Version 4.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV40(value *bool)() {
    m.v4_0 = value
}
// SetV403 sets the v4_0_3 property value. When TRUE, only Version 4.0.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV403(value *bool)() {
    m.v4_0_3 = value
}
// SetV41 sets the v4_1 property value. When TRUE, only Version 4.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV41(value *bool)() {
    m.v4_1 = value
}
// SetV42 sets the v4_2 property value. When TRUE, only Version 4.2 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV42(value *bool)() {
    m.v4_2 = value
}
// SetV43 sets the v4_3 property value. When TRUE, only Version 4.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV43(value *bool)() {
    m.v4_3 = value
}
// SetV44 sets the v4_4 property value. When TRUE, only Version 4.4 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV44(value *bool)() {
    m.v4_4 = value
}
// SetV50 sets the v5_0 property value. When TRUE, only Version 5.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV50(value *bool)() {
    m.v5_0 = value
}
// SetV51 sets the v5_1 property value. When TRUE, only Version 5.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV51(value *bool)() {
    m.v5_1 = value
}
// SetV60 sets the v6_0 property value. When TRUE, only Version 6.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV60(value *bool)() {
    m.v6_0 = value
}
// SetV70 sets the v7_0 property value. When TRUE, only Version 7.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV70(value *bool)() {
    m.v7_0 = value
}
// SetV71 sets the v7_1 property value. When TRUE, only Version 7.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV71(value *bool)() {
    m.v7_1 = value
}
// SetV80 sets the v8_0 property value. When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV80(value *bool)() {
    m.v8_0 = value
}
// SetV81 sets the v8_1 property value. When TRUE, only Version 8.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV81(value *bool)() {
    m.v8_1 = value
}
// SetV90 sets the v9_0 property value. When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating system boolean values will be TRUE.
func (m *AndroidMinimumOperatingSystem) SetV90(value *bool)() {
    m.v9_0 = value
}
