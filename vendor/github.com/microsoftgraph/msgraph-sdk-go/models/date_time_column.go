package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DateTimeColumn 
type DateTimeColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more details. If unspecified, treated as default.
    displayAs *string
    // Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or dateTime
    format *string
    // The OdataType property
    odataType *string
}
// NewDateTimeColumn instantiates a new dateTimeColumn and sets the default values.
func NewDateTimeColumn()(*DateTimeColumn) {
    m := &DateTimeColumn{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreateDateTimeColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDateTimeColumnFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDateTimeColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DateTimeColumn) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetDisplayAs gets the displayAs property value. How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more details. If unspecified, treated as default.
func (m *DateTimeColumn) GetDisplayAs()(*string) {
    return m.displayAs
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DateTimeColumn) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAs(val)
        }
        return nil
    }
    res["format"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFormat(val)
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
// GetFormat gets the format property value. Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or dateTime
func (m *DateTimeColumn) GetFormat()(*string) {
    return m.format
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *DateTimeColumn) GetOdataType()(*string) {
    return m.odataType
}
// Serialize serializes information the current object
func (m *DateTimeColumn) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("format", m.GetFormat())
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
func (m *DateTimeColumn) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetDisplayAs sets the displayAs property value. How the value should be presented in the UX. Must be one of default, friendly, or standard. See below for more details. If unspecified, treated as default.
func (m *DateTimeColumn) SetDisplayAs(value *string)() {
    m.displayAs = value
}
// SetFormat sets the format property value. Indicates whether the value should be presented as a date only or a date and time. Must be one of dateOnly or dateTime
func (m *DateTimeColumn) SetFormat(value *string)() {
    m.format = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *DateTimeColumn) SetOdataType(value *string)() {
    m.odataType = value
}
