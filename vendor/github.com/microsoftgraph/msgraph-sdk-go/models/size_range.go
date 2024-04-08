package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SizeRange 
type SizeRange struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
    maximumSize *int32
    // The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
    minimumSize *int32
    // The OdataType property
    odataType *string
}
// NewSizeRange instantiates a new sizeRange and sets the default values.
func NewSizeRange()(*SizeRange) {
    m := &SizeRange{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateSizeRangeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSizeRangeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSizeRange(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SizeRange) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SizeRange) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["maximumSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumSize(val)
        }
        return nil
    }
    res["minimumSize"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumSize(val)
        }
        return nil
    }
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
    return res
}
// GetMaximumSize gets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) GetMaximumSize()(*int32) {
    return m.maximumSize
}
// GetMinimumSize gets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) GetMinimumSize()(*int32) {
    return m.minimumSize
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *SizeRange) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *SizeRange) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("maximumSize", m.GetMaximumSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("minimumSize", m.GetMinimumSize())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
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
func (m *SizeRange) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetMaximumSize sets the maximumSize property value. The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) SetMaximumSize(value *int32)() {
    m.maximumSize = value
}
// SetMinimumSize sets the minimumSize property value. The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
func (m *SizeRange) SetMinimumSize(value *int32)() {
    m.minimumSize = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *SizeRange) SetOdataType(value *string)() {
    m.odataType = value
}
