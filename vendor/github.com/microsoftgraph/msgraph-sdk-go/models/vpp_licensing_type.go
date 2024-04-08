package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// VppLicensingType contains properties for iOS Volume-Purchased Program (Vpp) Licensing Type.
type VppLicensingType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The OdataType property
    odataType *string
    // Whether the program supports the device licensing type.
    supportsDeviceLicensing *bool
    // Whether the program supports the user licensing type.
    supportsUserLicensing *bool
}
// NewVppLicensingType instantiates a new vppLicensingType and sets the default values.
func NewVppLicensingType()(*VppLicensingType) {
    m := &VppLicensingType{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateVppLicensingTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVppLicensingTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVppLicensingType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VppLicensingType) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VppLicensingType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
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
    res["supportsDeviceLicensing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsDeviceLicensing(val)
        }
        return nil
    }
    res["supportsUserLicensing"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSupportsUserLicensing(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *VppLicensingType) GetOdataType()(*string) {
    return m.odataType
}
// GetSupportsDeviceLicensing gets the supportsDeviceLicensing property value. Whether the program supports the device licensing type.
func (m *VppLicensingType) GetSupportsDeviceLicensing()(*bool) {
    return m.supportsDeviceLicensing
}
// GetSupportsUserLicensing gets the supportsUserLicensing property value. Whether the program supports the user licensing type.
func (m *VppLicensingType) GetSupportsUserLicensing()(*bool) {
    return m.supportsUserLicensing
}
// Serialize serializes information the current object
func (m *VppLicensingType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("supportsDeviceLicensing", m.GetSupportsDeviceLicensing())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("supportsUserLicensing", m.GetSupportsUserLicensing())
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
func (m *VppLicensingType) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *VppLicensingType) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSupportsDeviceLicensing sets the supportsDeviceLicensing property value. Whether the program supports the device licensing type.
func (m *VppLicensingType) SetSupportsDeviceLicensing(value *bool)() {
    m.supportsDeviceLicensing = value
}
// SetSupportsUserLicensing sets the supportsUserLicensing property value. Whether the program supports the user licensing type.
func (m *VppLicensingType) SetSupportsUserLicensing(value *bool)() {
    m.supportsUserLicensing = value
}
